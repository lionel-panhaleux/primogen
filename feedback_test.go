package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

// githubAPI stands in for GitHub's REST API as an App installed on one
// repository sees it.
type githubAPI struct {
	t        *testing.T
	key      *rsa.PublicKey
	mu       sync.Mutex
	issues   []map[string]any
	clientID string
}

func (g *githubAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if r.Header.Get("X-GitHub-Api-Version") == "" {
		http.Error(w, `{"message":"no API version"}`, http.StatusBadRequest)
		return
	}
	switch r.Method + " " + r.URL.Path {
	case "GET /repos/lionel-panhaleux/primogen/installation":
		if g.app(w, r) {
			w.Write([]byte(`{"id": 4242}`))
		}
	case "POST /app/installations/4242/access_tokens":
		var scope struct {
			Repositories []string
			Permissions  map[string]string
		}
		json.NewDecoder(r.Body).Decode(&scope)
		if len(scope.Repositories) != 1 || scope.Repositories[0] != "primogen" || len(scope.Permissions) != 1 || scope.Permissions["issues"] != "write" {
			g.t.Errorf("a token not scoped to primogen's issues: %+v", scope)
		}
		if g.app(w, r) {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"token": "ghs_test", "expires_at": "2099-01-01T00:00:00Z"}`))
		}
	case "POST /repos/lionel-panhaleux/primogen/issues":
		if r.Header.Get("Authorization") != "token ghs_test" {
			http.Error(w, `{"message":"Bad credentials"}`, http.StatusUnauthorized)
			return
		}
		var issue map[string]any
		json.NewDecoder(r.Body).Decode(&issue)
		g.issues = append(g.issues, issue)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{"number": len(g.issues), "html_url": fmt.Sprintf("https://github.com/lionel-panhaleux/primogen/issues/%d", len(g.issues))})
	default:
		http.NotFound(w, r)
	}
}

// app checks the App's JWT as GitHub does.
func (g *githubAPI) app(w http.ResponseWriter, r *http.Request) bool {
	jwt, _ := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	parts := strings.Split(jwt, ".")
	ok := len(parts) == 3
	var claims struct {
		Iat, Exp int64
		Iss      string
	}
	if ok {
		sig, _ := base64.RawURLEncoding.DecodeString(parts[2])
		sum := sha256.Sum256([]byte(parts[0] + "." + parts[1]))
		raw, _ := base64.RawURLEncoding.DecodeString(parts[1])
		ok = rsa.VerifyPKCS1v15(g.key, crypto.SHA256, sum[:], sig) == nil && json.Unmarshal(raw, &claims) == nil
	}
	if !ok || claims.Iss != g.clientID || claims.Exp-claims.Iat > 600 {
		g.t.Errorf("a JWT GitHub refuses: %q %+v", jwt, claims)
		http.Error(w, `{"message":"A JSON web token could not be decoded"}`, http.StatusUnauthorized)
		return false
	}
	return true
}

func TestFeedback(t *testing.T) {
	ctx := t.Context()
	dir := t.TempDir()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	keyPath := filepath.Join(dir, "app.pem")
	os.WriteFile(keyPath, pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}), 0o600)
	gh := &githubAPI{t: t, key: &key.PublicKey, clientID: "Iv23test"}
	api := httptest.NewServer(gh)
	defer api.Close()

	// the hosted server, its bundle the tests' own
	raw, err := os.ReadFile(testBundle)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "bundle.json.gz")
	os.WriteFile(path, raw, 0o644)
	r, err := newRefresher(ctx, bundle.Upstream{}, path, slog.New(slog.DiscardHandler))
	if err != nil {
		t.Fatal(err)
	}
	g, err := newGitHubApp(api.URL, "lionel-panhaleux/primogen", "Iv23test", keyPath)
	if err != nil {
		t.Fatal(err)
	}
	r.fileWith(g)
	hosted := httptest.NewServer(r.handler())
	defer hosted.Close()

	call := func(from string, args map[string]any) (string, bool) {
		t.Helper()
		client := &http.Client{Transport: roundTrip(func(req *http.Request) (*http.Response, error) {
			if from != "" {
				req.Header.Set("X-Real-IP", from)
			}
			return http.DefaultTransport.RoundTrip(req)
		})}
		cs, err := mcp.NewClient(&mcp.Implementation{Name: "test"}, nil).Connect(ctx, &mcp.StreamableClientTransport{Endpoint: hosted.URL + "/mcp", HTTPClient: client}, nil)
		if err != nil {
			t.Fatal(err)
		}
		defer cs.Close()
		res, err := cs.CallTool(ctx, &mcp.CallToolParams{Name: "vtes_feedback", Arguments: args})
		if err != nil {
			t.Fatal(err)
		}
		return res.Content[0].(*mcp.TextContent).Text, res.IsError
	}
	report := map[string]any{
		"question": "Can I play a Freak Drive after superior Kiss of Ra? @someone #3",
		"wrong":    "It said no; @someone said yes in #12, see ```this```",
		"expected": "Yes: Kiss of Ra's rulings allow it.",
	}

	// the user not yet told: nothing filed
	text, isErr := call("", report)
	if !isErr || !strings.Contains(text, "public GitHub issue") || len(gh.issues) != 0 {
		t.Errorf("filed without the user's agreement: %v %s", isErr, text)
	}
	report["public"] = true
	text, isErr = call("", report)
	if isErr || !strings.Contains(text, "Filed as https://github.com/lionel-panhaleux/primogen/issues/1.") {
		t.Fatalf("not filed: %s", text)
	}
	issue := gh.issues[0]
	body := issue["body"].(string)
	if issue["title"] != "Feedback: Can I play a Freak Drive after superior Kiss of Ra? ＠someone ＃3" || len(issue["labels"].([]any)) != 1 || issue["labels"].([]any)[0] != "feedback" {
		t.Errorf("the issue: %v", issue)
	}
	for _, want := range []string{"### What was wrong\n\n````text\nIt said no; @someone said yes in #12, see ```this```\n````", "Filed through `vtes_feedback`, primogen ", "Data: "} {
		if !strings.Contains(body, want) {
			t.Errorf("missing %q in the issue's body:\n%s", want, body)
		}
	}

	// too long
	long := map[string]any{"question": strings.Repeat("a", 2001), "wrong": "w", "expected": "e", "public": true}
	if text, isErr := call("", long); !isErr || !strings.Contains(text, "at most 2000") {
		t.Errorf("a long report: %s", text)
	}

	// a local install files through the hosted server
	var stdout, stderr strings.Builder
	if err := run(ctx, []string{"query", "-bundle", testBundle, "-from", hosted.URL + "/bundle.json.gz", "feedback", "q", "w", "e"}, &stdout, &stderr); err != nil {
		t.Fatal(err, stderr.String())
	}
	if !strings.Contains(stdout.String(), "Filed as https://github.com/lionel-panhaleux/primogen/issues/2.") {
		t.Errorf("a local report: %s", stdout.String())
	}

	// the limit of an address, others unaffected
	for range 5 {
		if text, isErr := call("192.0.2.1", report); isErr {
			t.Fatal(text)
		}
	}
	if text, isErr := call("192.0.2.1", report); !isErr || !strings.Contains(text, "5 reports filed from this address in the last day") {
		t.Errorf("a sixth report from one address: %s", text)
	}
	if text, isErr := call("192.0.2.2", report); isErr {
		t.Errorf("another address refused: %s", text)
	}
}
