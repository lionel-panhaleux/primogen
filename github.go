package main

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/lionel-panhaleux/primogen/internal/tools"
)

// githubApp files issues on one repository as a GitHub App installed there.
type githubApp struct {
	api, repo, clientID string
	key                 *rsa.PrivateKey
	client              *http.Client
	limit               *limiter

	mu           sync.Mutex
	installation int64
	token        string
	expires      time.Time
}

func newGitHubApp(api, repo, clientID, keyPath string) (*githubApp, error) {
	raw, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(raw)
	if block == nil {
		return nil, fmt.Errorf("%s: no PEM block", keyPath)
	}
	var key *rsa.PrivateKey
	if key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		parsed, err8 := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err8 != nil {
			return nil, fmt.Errorf("%s: %w", keyPath, errors.Join(err, err8))
		}
		var ok bool
		if key, ok = parsed.(*rsa.PrivateKey); !ok {
			return nil, fmt.Errorf("%s: not an RSA key", keyPath)
		}
	}
	return &githubApp{
		api: api, repo: repo, clientID: clientID, key: key,
		client: &http.Client{Timeout: 30 * time.Second},
		limit:  &limiter{perAddress: 5, all: 50, seen: map[string][]time.Time{}},
	}, nil
}

func (g *githubApp) file(ctx context.Context, r tools.Report, from, data string) (string, error) {
	if err := g.limit.allow(from, time.Now()); err != nil {
		return "", fmt.Errorf("not filed: %w", err)
	}
	link, err := g.open(ctx, r, data)
	if err != nil {
		return "", fmt.Errorf("not filed: %w", err)
	}
	g.limit.count(from, time.Now())
	return link, nil
}

func (g *githubApp) open(ctx context.Context, r tools.Report, data string) (string, error) {
	token, err := g.installationToken(ctx)
	if err != nil {
		return "", err
	}
	title, body := issueText(r, data)
	var issue struct {
		HTMLURL string `json:"html_url"`
	}
	in := map[string]any{"title": title, "body": body, "labels": []string{"feedback"}}
	if err := g.call(ctx, http.MethodPost, "/repos/"+g.repo+"/issues", "token "+token, in, http.StatusCreated, &issue); err != nil {
		return "", err
	}
	return issue.HTMLURL, nil
}

// installationToken answers a token valid for five minutes at least.
func (g *githubApp) installationToken(ctx context.Context) (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.token != "" && time.Until(g.expires) > 5*time.Minute {
		return g.token, nil
	}
	jwt, err := g.jwt(time.Now())
	if err != nil {
		return "", err
	}
	if g.installation == 0 {
		var inst struct{ ID int64 }
		if err := g.call(ctx, http.MethodGet, "/repos/"+g.repo+"/installation", "Bearer "+jwt, nil, http.StatusOK, &inst); err != nil {
			return "", err
		}
		g.installation = inst.ID
	}
	var tok struct {
		Token     string    `json:"token"`
		ExpiresAt time.Time `json:"expires_at"`
	}
	// the App is installed on other repositories too: the token reaches this one's issues only
	_, name, _ := strings.Cut(g.repo, "/")
	scope := map[string]any{"repositories": []string{name}, "permissions": map[string]string{"issues": "write"}}
	if err := g.call(ctx, http.MethodPost, fmt.Sprintf("/app/installations/%d/access_tokens", g.installation), "Bearer "+jwt, scope, http.StatusCreated, &tok); err != nil {
		return "", err
	}
	g.token, g.expires = tok.Token, tok.ExpiresAt
	return g.token, nil
}

// jwt authenticates as the App: RS256, issued a minute back for clock drift,
// GitHub refusing a lifetime over ten minutes.
func (g *githubApp) jwt(now time.Time) (string, error) {
	enc := base64.RawURLEncoding
	header := enc.EncodeToString([]byte(`{"alg":"RS256","typ":"JWT"}`))
	claims, err := json.Marshal(map[string]any{"iat": now.Add(-time.Minute).Unix(), "exp": now.Add(9 * time.Minute).Unix(), "iss": g.clientID})
	if err != nil {
		return "", err
	}
	signed := header + "." + enc.EncodeToString(claims)
	sum := sha256.Sum256([]byte(signed))
	sig, err := rsa.SignPKCS1v15(rand.Reader, g.key, crypto.SHA256, sum[:])
	if err != nil {
		return "", err
	}
	return signed + "." + enc.EncodeToString(sig), nil
}

func (g *githubApp) call(ctx context.Context, method, path, auth string, in any, want int, out any) error {
	var body bytes.Buffer
	if in != nil {
		if err := json.NewEncoder(&body).Encode(in); err != nil {
			return err
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, g.api+path, &body)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")
	req.Header.Set("Authorization", auth)
	req.Header.Set("User-Agent", "primogen/"+tools.Version())
	resp, err := g.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != want {
		var e struct{ Message string }
		json.NewDecoder(resp.Body).Decode(&e)
		return fmt.Errorf("GitHub answered %s %s with %s: %s", method, path, resp.Status, e.Message)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

// issueText is the issue a report becomes. Each field is fenced and the
// title's @ and # are full-width, so that its words mention no one and link
// no issue.
func issueText(r tools.Report, data string) (title, body string) {
	title = strings.NewReplacer("@", "＠", "#", "＃").Replace("Feedback: " + strings.Join(strings.Fields(r.Question), " "))
	if t := []rune(title); len(t) > 80 {
		title = string(t[:79]) + "…"
	}
	var b strings.Builder
	for _, f := range []struct{ head, text string }{{"Question", r.Question}, {"What was wrong", r.Wrong}, {"What was expected", r.Expected}} {
		fence := "```"
		for strings.Contains(f.text, fence) {
			fence += "`"
		}
		fmt.Fprintf(&b, "### %s\n\n%stext\n%s\n%s\n\n", f.head, fence, strings.TrimSpace(f.text), fence)
	}
	fmt.Fprintf(&b, "Filed through `vtes_feedback`, primogen %s. %s\n", tools.Version(), data)
	return title, b.String()
}

// limiter caps the issues filed for each address and for all of them over a
// day, held in memory.
type limiter struct {
	perAddress, all int
	mu              sync.Mutex
	seen            map[string][]time.Time
}

func (l *limiter) allow(from string, now time.Time) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	since := now.Add(-24 * time.Hour)
	all := 0
	for k, ts := range l.seen {
		i := 0
		for i < len(ts) && ts[i].Before(since) {
			i++
		}
		if l.seen[k] = ts[i:]; len(l.seen[k]) == 0 {
			delete(l.seen, k)
		}
		all += len(l.seen[k])
	}
	if all >= l.all {
		return fmt.Errorf("primogen filed %d reports in the last day, its limit: try again tomorrow", l.all)
	}
	if len(l.seen[from]) >= l.perAddress {
		return fmt.Errorf("%d reports filed from this address in the last day, the limit: try again tomorrow", l.perAddress)
	}
	return nil
}

func (l *limiter) count(from string, now time.Time) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.seen[from] = append(l.seen[from], now)
}
