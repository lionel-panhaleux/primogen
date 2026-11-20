package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/modelcontextprotocol/go-sdk/oauthex"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
	"github.com/lionel-panhaleux/primogen/internal/tools"
)

type archonUpstream struct {
	t                      *testing.T
	clientID, clientSecret string

	mu         sync.Mutex
	veknID     string
	exportRaw  []byte
	challenges map[string]string
}

func (a *archonUpstream) setExport(version int, generatedAt string, lines []string, eofCount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	var buf bytes.Buffer
	z := gzip.NewWriter(&buf)
	fmt.Fprintf(z, `{"type":"header","version":%d,"timestamp":%q,"generated_at":%q}`+"\n", version, generatedAt, generatedAt)
	for _, l := range lines {
		fmt.Fprintln(z, l)
	}
	if eofCount >= 0 {
		fmt.Fprintf(z, `{"type":"eof","count":%d}`+"\n", eofCount)
	}
	z.Close()
	a.exportRaw = buf.Bytes()
}

func (a *archonUpstream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.challenges == nil {
		a.challenges = map[string]string{}
	}
	switch r.Method + " " + r.URL.Path {
	case "GET /consent":
		q := r.URL.Query()
		if q.Get("client_id") != a.clientID || q.Get("response_type") != "code" || q.Get("scope") != "profile:read" || q.Get("code_challenge_method") != "S256" || q.Get("code_challenge") == "" {
			a.t.Errorf("bad /consent request: %s", r.URL.RawQuery)
			http.Error(w, "bad consent request", http.StatusBadRequest)
			return
		}
		code := "archon-code-" + q.Get("state")
		a.challenges[code] = q.Get("code_challenge")
		u, _ := url.Parse(q.Get("redirect_uri"))
		out := u.Query()
		out.Set("code", code)
		out.Set("state", q.Get("state"))
		u.RawQuery = out.Encode()
		http.Redirect(w, r, u.String(), http.StatusFound)
	case "POST /oauth/token":
		r.ParseForm()
		if r.PostFormValue("client_id") != a.clientID || r.PostFormValue("client_secret") != a.clientSecret {
			http.Error(w, `{"detail":"Invalid client credentials"}`, http.StatusUnauthorized)
			return
		}
		switch r.PostFormValue("grant_type") {
		case "client_credentials":
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"access_token":"archon-app-token","token_type":"Bearer","expires_in":3600,"scope":"api:read"}`))
		case "authorization_code":
			code := r.PostFormValue("code")
			challenge, ok := a.challenges[code]
			delete(a.challenges, code)
			sum := sha256.Sum256([]byte(r.PostFormValue("code_verifier")))
			got := base64.RawURLEncoding.EncodeToString(sum[:])
			if !ok || got != challenge {
				http.Error(w, `{"detail":"Invalid code_verifier (PKCE)"}`, http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"access_token":"archon-user-token","refresh_token":"archon-refresh","token_type":"Bearer","expires_in":3600,"scope":"profile:read"}`))
		default:
			http.Error(w, `{"detail":"Unsupported grant_type"}`, http.StatusBadRequest)
		}
	case "GET /oauth/userinfo":
		if r.Header.Get("Authorization") != "Bearer archon-user-token" {
			http.Error(w, `{"detail":"Invalid token"}`, http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"sub":          "019f6a5e-d1e6-7347-965e-4dac09bf03a4",
			"vekn_id":      a.veknID,
			"roles":        []string{"Prince"},
			"capabilities": []string{"create_tournament"},
		})
	case "GET /v1/export":
		if r.Header.Get("Authorization") != "Bearer archon-app-token" {
			http.Error(w, `{"detail":"Missing or invalid authorization header"}`, http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/gzip")
		w.Write(a.exportRaw)
	default:
		http.NotFound(w, r)
	}
}

func TestArchonLoginAndExport(t *testing.T) {
	ctx := t.Context()
	dir := t.TempDir()

	up := &archonUpstream{
		t:            t,
		clientID:     "archon-client-id",
		clientSecret: "archon-client-secret",
		veknID:       "3200340",
	}
	lines := []string{
		`{"type":"user","data":{"uid":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","vekn_id":"3200340","country":"FR","city":"Paris","roles":["IC","Prince"],"wins":["019fea18-1364-7549-8f9a-70dcc09a03d9"],"constructed_offline":{"total":180,"tournaments":[]},"community_links":[{"url":"https://codex-of-the-damned.org","type":"website","label":"Codex of the Damned","languages":["en","fr"]}]}}`,
		`{"type":"tournament","data":{"uid":"019fea18-1364-7549-8f9a-70dcc09a03d9","name":"Malmö Unbound","format":"Standard","start":"2026-06-13T10:00:00","finish":"2026-06-13T18:00:00","timezone":"Europe/Stockholm","country":"SE","city":"Malmö","state":"Finished","event_code":"MU-2606","winner":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","players":[{"user_uid":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","state":"Finished","result":{"gw":2,"vp":6.5,"tp":114},"finalist":true}]}}`,
		`{"type":"deck","data":{"uid":"019f6a65-01d4-7214-bcc9-4e46534b9d62","tournament_uid":"019fea18-1364-7549-8f9a-70dcc09a03d9","user_uid":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","name":"Haqim Royalty","cards":{"100788":2,"200191":4},"attribution":{"kind":"Owner"},"winner":true}}`,
		`{"type":"league","data":{"uid":"019fe07a-b107-7465-9968-122ceb6d67d1","name":"Nordic Chapter 2026","kind":"League","standings_mode":"GP","format":"Standard","country":"SE","start":"2026-01-01","finish":"2026-12-31"}}`,
	}
	up.setExport(2, "2026-11-20T09:15:00.000000", lines, len(lines))
	archonSrv := httptest.NewServer(up)
	defer archonSrv.Close()

	secretPath := filepath.Join(dir, "archon-secret")
	if err := os.WriteFile(secretPath, []byte(up.clientSecret+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	a, err := newArchonApp(archonSrv.URL, archonSrv.URL, up.clientID, secretPath)
	if err != nil {
		t.Fatal(err)
	}

	raw, err := os.ReadFile(testBundle)
	if err != nil {
		t.Fatal(err)
	}
	bundlePath := filepath.Join(dir, "bundle.json.gz")
	if err := os.WriteFile(bundlePath, raw, 0o644); err != nil {
		t.Fatal(err)
	}
	r, err := newRefresher(ctx, bundle.Upstream{}, bundlePath, slog.New(slog.DiscardHandler))
	if err != nil {
		t.Fatal(err)
	}
	r.withArchon(a)
	now := time.Date(2026, 11, 20, 9, 20, 0, 0, time.UTC)
	r.refreshArchon(ctx, now)

	hosted := httptest.NewServer(r.handler())
	defer hosted.Close()

	resp, err := http.Get(hosted.URL + "/status")
	if err != nil {
		t.Fatal(err)
	}
	statusBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if !strings.Contains(string(statusBody), "Archon export: 2026-11-20 09:15 UTC (1 tournaments, 1 decks, 1 members, 1 leagues); last refresh 2026-11-20 09:20 UTC: ok.") {
		t.Fatalf("unexpected /status:\n%s", statusBody)
	}
	pubResp, err := http.Get(hosted.URL + "/bundle.json.gz")
	if err != nil {
		t.Fatal(err)
	}
	pubBytes, _ := io.ReadAll(pubResp.Body)
	pubResp.Body.Close()
	if !bytes.Equal(pubBytes, raw) {
		t.Fatal("published bundle changed after Archon export refresh")
	}

	asmResp, err := http.Get(hosted.URL + "/.well-known/oauth-authorization-server")
	if err != nil {
		t.Fatal(err)
	}
	asmBytes, _ := io.ReadAll(asmResp.Body)
	asmResp.Body.Close()
	if strings.Contains(string(asmBytes), "jwks_uri") {
		t.Errorf("auth server metadata emitted empty jwks_uri: %s", asmBytes)
	}

	cs, err := mcp.NewClient(&mcp.Implementation{Name: "anon"}, nil).Connect(ctx, &mcp.StreamableClientTransport{Endpoint: hosted.URL + "/mcp"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := cs.CallTool(ctx, &mcp.CallToolParams{Name: "vtes_cards", Arguments: map[string]any{"names": []string{"Freak Drive"}}})
	cs.Close()
	if err != nil || res.IsError || !strings.Contains(res.Content[0].(*mcp.TextContent).Text, "## Freak Drive") {
		t.Fatalf("unauthenticated vtes_cards failed: %v %+v", err, res)
	}

	expect401 := func(path, authHeader string) {
		t.Helper()
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, hosted.URL+path, strings.NewReader(`{}`))
		if err != nil {
			t.Fatal(err)
		}
		if authHeader != "" {
			req.Header.Set("Authorization", authHeader)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
		www := resp.Header.Get("WWW-Authenticate")
		if resp.StatusCode != http.StatusUnauthorized || !strings.Contains(www, "resource_metadata=") || !strings.Contains(www, `scope="member"`) {
			t.Fatalf("%s (auth=%q): status %d, WWW-Authenticate=%q", path, authHeader, resp.StatusCode, www)
		}
	}
	expect401("/mcp?login=1", "")
	expect401("/mcp/member", "")
	expect401("/mcp", "Bearer not-a-valid-token")

	browserFlow := func(authMethod string) (*auth.AuthorizationCodeHandler, error) {
		noRedirect := &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
		return auth.NewAuthorizationCodeHandler(&auth.AuthorizationCodeHandlerConfig{
			DynamicClientRegistrationConfig: &auth.DynamicClientRegistrationConfig{
				Metadata: &oauthex.ClientRegistrationMetadata{
					RedirectURIs:            []string{"http://127.0.0.1:54321/callback"},
					TokenEndpointAuthMethod: authMethod,
					ClientName:              "test-mcp-client",
				},
			},
			AuthorizationCodeFetcher: func(ctx context.Context, args *auth.AuthorizationArgs) (*auth.AuthorizationResult, error) {
				r1, err := noRedirect.Get(args.URL)
				if err != nil {
					return nil, err
				}
				r1.Body.Close()
				if r1.StatusCode != http.StatusFound {
					return nil, fmt.Errorf("/oauth/authorize answered %s", r1.Status)
				}
				r2, err := noRedirect.Get(r1.Header.Get("Location"))
				if err != nil {
					return nil, err
				}
				r2.Body.Close()
				if r2.StatusCode != http.StatusFound {
					return nil, fmt.Errorf("/consent answered %s", r2.Status)
				}
				r3, err := noRedirect.Get(r2.Header.Get("Location"))
				if err != nil {
					return nil, err
				}
				r3.Body.Close()
				if r3.StatusCode != http.StatusFound {
					return nil, fmt.Errorf("/oauth/callback answered %s", r3.Status)
				}
				cb, err := url.Parse(r3.Header.Get("Location"))
				if err != nil {
					return nil, err
				}
				if errCode := cb.Query().Get("error"); errCode != "" {
					return nil, fmt.Errorf("oauth error %s: %s", errCode, cb.Query().Get("error_description"))
				}
				return &auth.AuthorizationResult{
					Code:  cb.Query().Get("code"),
					State: cb.Query().Get("state"),
				}, nil
			},
		})
	}

	oauthHandler, err := browserFlow("none")
	if err != nil {
		t.Fatal(err)
	}
	memberCS, err := mcp.NewClient(&mcp.Implementation{Name: "member"}, nil).Connect(ctx, &mcp.StreamableClientTransport{
		Endpoint:     hosted.URL + "/mcp?login=1",
		OAuthHandler: oauthHandler,
	}, nil)
	if err != nil {
		t.Fatalf("member OAuth login failed: %v", err)
	}
	res, err = memberCS.CallTool(ctx, &mcp.CallToolParams{Name: "vtes_cards", Arguments: map[string]any{"names": []string{"Freak Drive"}}})
	memberCS.Close()
	if err != nil || res.IsError {
		t.Fatalf("authenticated tool call failed: %v", err)
	}
	ts, err := oauthHandler.TokenSource(ctx)
	if err != nil || ts == nil {
		t.Fatalf("no token source after login: %v", err)
	}
	tok, err := ts.Token()
	if err != nil {
		t.Fatal(err)
	}
	info, err := a.verifyBearer(ctx, tok.AccessToken, nil)
	if err != nil || info.UserID != "3200340" {
		t.Fatalf("verified token UserID = %+v, err = %v; want 3200340", info, err)
	}
	if got := tools.MemberFrom(&mcp.CallToolRequest{Extra: &mcp.RequestExtra{TokenInfo: info}}); got != "3200340" {
		t.Errorf("MemberFrom = %q, want 3200340", got)
	}
	if tok.RefreshToken == "" {
		t.Fatal("missing refresh_token after login")
	}

	var refreshClaims tokenClaims
	if err := a.unseal(tok.RefreshToken, &refreshClaims); err != nil {
		t.Fatal(err)
	}
	refreshed, err := a.postToken(ctx, hosted.URL+"/oauth/token", url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {refreshClaims.ClientID},
		"refresh_token": {tok.RefreshToken},
	})
	if err != nil {
		t.Fatalf("refresh_token exchange failed: %v", err)
	}
	if info, err := a.verifyBearer(ctx, refreshed.AccessToken, nil); err != nil || info.UserID != "3200340" {
		t.Fatalf("refreshed token UserID = %+v, err = %v; want 3200340", info, err)
	}

	basicHandler, err := browserFlow("client_secret_basic")
	if err != nil {
		t.Fatal(err)
	}
	basicCS, err := mcp.NewClient(&mcp.Implementation{Name: "basic"}, nil).Connect(ctx, &mcp.StreamableClientTransport{
		Endpoint:     hosted.URL + "/mcp/member",
		OAuthHandler: basicHandler,
	}, nil)
	if err != nil {
		t.Fatalf("confidential client OAuth login failed: %v", err)
	}
	basicCS.Close()

	regResp, err := oauthex.RegisterClient(ctx, hosted.URL+"/oauth/register", &oauthex.ClientRegistrationMetadata{
		RedirectURIs:            []string{"https://example.com/callback"},
		TokenEndpointAuthMethod: "client_secret_post",
	}, http.DefaultClient)
	if err != nil {
		t.Fatal(err)
	}
	badSecretResp, err := http.PostForm(hosted.URL+"/oauth/token", url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {regResp.ClientID},
		"client_secret": {"wrong-secret"},
		"refresh_token": {tok.RefreshToken},
	})
	if err != nil {
		t.Fatal(err)
	}
	badSecretResp.Body.Close()
	if badSecretResp.StatusCode != http.StatusUnauthorized {
		t.Errorf("wrong client_secret status = %d, want 401", badSecretResp.StatusCode)
	}

	expiredTok := a.seal(tokenClaims{Type: "access", UID: "019f6a5e-d1e6-7347-965e-4dac09bf03a4", VEKNID: "3200340", Exp: time.Now().Add(-time.Minute).Unix()})
	expect401("/mcp", "Bearer "+expiredTok)

	up.mu.Lock()
	up.veknID = ""
	up.mu.Unlock()
	nonMemberHandler, err := browserFlow("none")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := mcp.NewClient(&mcp.Implementation{Name: "non-member"}, nil).Connect(ctx, &mcp.StreamableClientTransport{
		Endpoint:     hosted.URL + "/mcp/member",
		OAuthHandler: nonMemberHandler,
	}, nil); err == nil || !strings.Contains(err.Error(), "no VEKN ID") {
		t.Fatalf("expected non-member account to be refused, got: %v", err)
	}

	up.setExport(2, "2026-11-20T09:30:00.000000", lines, -1)
	r.refreshArchon(ctx, now.Add(15*time.Minute))
	if r.live.Archon().Date() != "2026-11-20 09:15 UTC" {
		t.Errorf("truncated export replaced the live one: %s", r.live.Archon().Date())
	}
	resp, err = http.Get(hosted.URL + "/status")
	if err != nil {
		t.Fatal(err)
	}
	statusBody, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	if !strings.Contains(string(statusBody), "Archon export: 2026-11-20 09:15 UTC") || !strings.Contains(string(statusBody), "FAILED, the export above keeps serving: truncated export: no eof") {
		t.Fatalf("unexpected /status after truncated export:\n%s", statusBody)
	}

	up.setExport(3, "2026-11-20T09:45:00.000000", lines, len(lines))
	r.refreshArchon(ctx, now.Add(30*time.Minute))
	if r.live.Archon().Date() != "2026-11-20 09:15 UTC" {
		t.Errorf("version-3 export replaced the live one: %s", r.live.Archon().Date())
	}
}
