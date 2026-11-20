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
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/modelcontextprotocol/go-sdk/oauthex"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
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

	deckCards := map[string]int{}
	for _, d := range r.live.Bundle().Decks {
		if d.ID == "12494" {
			for _, c := range d.Cards {
				deckCards[strconv.Itoa(c.ID)] = c.Count
			}
			break
		}
	}
	cardsJSON, err := json.Marshal(deckCards)
	if err != nil {
		t.Fatal(err)
	}

	up := &archonUpstream{
		t:            t,
		clientID:     "archon-client-id",
		clientSecret: "archon-client-secret",
		veknID:       "3200340",
	}
	lines := []string{
		`{"type":"user","data":{"uid":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","vekn_id":"3200340","country":"FR","city":"Paris","roles":["IC","Prince"],"wins":["019fea18-1364-7549-8f9a-70dcc09a03d9"],"constructed_offline":{"total":180,"tournaments":[{"tournament_uid":"019fea18-1364-7549-8f9a-70dcc09a03d9","tournament_name":"Malmö Unbound","date":"2026-06-13","player_count":13,"rank":"","vp":6.5,"gw":2,"finalist_position":1,"points":180,"position":1}]},"community_links":[{"url":"https://codex-of-the-damned.org","type":"website","label":"Codex of the Damned","languages":["en","fr"],"moderation":"global"},{"url":"https://secret.example.com","type":"other","label":"Hidden Link","moderation":"hidden"}]}}`,
		`{"type":"user","data":{"uid":"019f6a5e-d5dd-74b9-9aab-00919d4de7dc","vekn_id":"1002345","country":"SE","city":"Malmö","roles":["Prince"],"wins":[],"constructed_offline":{"total":120,"tournaments":[]},"community_links":[{"url":"https://vekn.se","type":"website","label":"VEKN Sweden","country":"SE","moderation":"national"}]}}`,
		`{"type":"tournament","data":{"uid":"019fea18-1364-7549-8f9a-70dcc09a03d9","name":"Malmö Unbound","format":"Standard","rank":"","start":"2026-06-13T10:00:00","finish":"2026-06-13T18:00:00","timezone":"Europe/Stockholm","country":"SE","city":"Malmö","league_uid":"019fe07a-b107-7465-9968-122ceb6d67d1","state":"Finished","event_code":"MU-2606","external_ids":{"vekn":"13370"},"organizers_uids":["019f6a5e-d5dd-74b9-9aab-00919d4de7dc"],"winner":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","standings":[{"user_uid":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","gw":2,"vp":6.5,"tp":114,"finalist":true},{"user_uid":"019f6a5e-d5dd-74b9-9aab-00919d4de7dc","gw":1,"vp":5.0,"tp":108,"finalist":true}],"players":[{"user_uid":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","state":"Finished","result":{"gw":2,"vp":6.5,"tp":114},"finalist":true},{"user_uid":"019f6a5e-d5dd-74b9-9aab-00919d4de7dc","state":"Finished","result":{"gw":1,"vp":5.0,"tp":108},"finalist":true}],"reported_player_count":13}}`,
		`{"type":"tournament","data":{"uid":"019fe500-2222-7549-8f9a-70dcc09a03d9","name":"Paris Spring Open","format":"Standard","rank":"","start":"2026-05-10T10:00:00","finish":"2026-05-10T18:00:00","timezone":"Europe/Paris","country":"FR","city":"Paris","state":"Finished","event_code":"PS-2605","winner":"019f6a5e-d5dd-74b9-9aab-00919d4de7dc","reported_player_count":16}}`,
		`{"type":"tournament","data":{"uid":"019ff0c2-51e8-7a6d-b58c-4e9a2f7d1c03","name":"French National Championship","format":"Standard","rank":"National Championship","start":"2026-09-20T10:00:00","timezone":"Europe/Paris","country":"FR","city":"Paris","state":"Planned","event_code":"FR-NC-26","players":[]}}`,
		fmt.Sprintf(`{"type":"deck","data":{"uid":"019f6a65-01d4-7214-bcc9-4e46534b9d62","tournament_uid":"019fea18-1364-7549-8f9a-70dcc09a03d9","user_uid":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","name":"Malmö Haqim","comments":"Archon winner comment.","cards":%s,"attribution":{"kind":"Owner"},"winner":true}}`, cardsJSON),
		`{"type":"deck","data":{"uid":"019f6a65-1111-7214-bcc9-4e46534b9d62","tournament_uid":"019fea18-1364-7549-8f9a-70dcc09a03d9","user_uid":"019f6a5e-d5dd-74b9-9aab-00919d4de7dc","name":"Malmö Finalist Brew","comments":"Finalist deck.","cards":{"100001":1,"200001":1},"attribution":{"kind":"Owner"},"winner":false}}`,
		fmt.Sprintf(`{"type":"deck","data":{"uid":"019f6a65-2222-7214-bcc9-4e46534b9d62","tournament_uid":"019fe500-2222-7549-8f9a-70dcc09a03d9","user_uid":"019f6a5e-d5dd-74b9-9aab-00919d4de7dc","name":"Paris Haqim","comments":"Winner-only event deck.","cards":%s,"attribution":{"kind":"Owner"},"winner":true}}`, cardsJSON),
		`{"type":"league","data":{"uid":"019fe07a-b107-7465-9968-122ceb6d67d1","name":"Nordic Chapter 2026","kind":"League","standings_mode":"GP","format":"Standard","country":"SE","start":"2026-01-01","finish":"2026-12-31","organizers_uids":["019f6a5e-d5dd-74b9-9aab-00919d4de7dc"]}}`,
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
	if !strings.Contains(string(statusBody), "Archon export: 2026-11-20 09:15 UTC (3 tournaments, 3 decks, 2 members, 1 leagues); last refresh 2026-11-20 09:20 UTC: ok.") {
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

	callTool := func(cs *mcp.ClientSession, name string, args map[string]any) (string, bool) {
		t.Helper()
		res, err := cs.CallTool(ctx, &mcp.CallToolParams{Name: name, Arguments: args})
		if err != nil {
			t.Fatal(err)
		}
		return res.Content[0].(*mcp.TextContent).Text, res.IsError
	}

	anonCS, err := mcp.NewClient(&mcp.Implementation{Name: "anon"}, nil).Connect(ctx, &mcp.StreamableClientTransport{Endpoint: hosted.URL + "/mcp"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer anonCS.Close()
	if text, isErr := callTool(anonCS, "vtes_cards", map[string]any{"names": []string{"Freak Drive"}}); isErr || !strings.Contains(text, "## Freak Drive") {
		t.Fatalf("unauthenticated vtes_cards failed: %v %s", isErr, text)
	}
	if text, isErr := callTool(anonCS, "vtes_events", map[string]any{"links": true}); isErr || !strings.Contains(text, "Codex of the Damned") || strings.Contains(text, "Hidden Link") || strings.Contains(text, "VEKN 3200340") {
		t.Fatalf("unauthenticated community links: isErr=%v\n%s", isErr, text)
	}
	if text, isErr := callTool(anonCS, "vtes_events", map[string]any{"tournament": "MU-2606"}); !isErr || !strings.Contains(text, "logged-in VEKN members only") {
		t.Fatalf("expected unauthenticated vtes_events refusal, got isErr=%v:\n%s", isErr, text)
	}
	if text, isErr := callTool(anonCS, "vtes_deck", map[string]any{"deck": "019f6a65-01d4-7214-bcc9-4e46534b9d62"}); !isErr || !strings.Contains(text, "logged-in VEKN members only") {
		t.Fatalf("expected unauthenticated vtes_deck Archon refusal, got isErr=%v:\n%s", isErr, text)
	}
	if text, isErr := callTool(anonCS, "vtes_deck", map[string]any{"deck": "https://archon.vekn.net/tournaments/MU-2606"}); !isErr || !strings.Contains(text, "logged-in VEKN members only") {
		t.Fatalf("expected unauthenticated vtes_deck Archon link refusal, got isErr=%v:\n%s", isErr, text)
	}
	if text, isErr := callTool(anonCS, "vtes_meta", map[string]any{}); isErr ||
		!strings.Contains(text, "full-field play share against win share requires Archon login at https://mcp.krcg.org/mcp") ||
		strings.Contains(text, "## Full-field play against win share (Archon)") ||
		strings.Contains(text, "archon.vekn.net export") {
		t.Fatalf("unauthenticated vtes_meta: isErr=%v\n%s", isErr, text)
	}
	var stdout, stderr strings.Builder
	if err := run(ctx, []string{"query", "-bundle", testBundle, "-rankings", "events"}, &stdout, &stderr); err == nil || !strings.Contains(err.Error(), "logged-in VEKN members only") {
		t.Fatalf("expected local query events refusal, got err=%v", err)
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
	defer memberCS.Close()

	if text, isErr := callTool(memberCS, "vtes_events", map[string]any{"tournament": "MU-2606"}); isErr ||
		!strings.Contains(text, "# Tournament: Malmö Unbound") ||
		!strings.Contains(text, "Winner: **VEKN 3200340**") ||
		!strings.Contains(text, "019f6a65-01d4-7214-bcc9-4e46534b9d62") ||
		!strings.Contains(text, "Haqim Royalty") ||
		!strings.Contains(text, "archon.vekn.net export 2026-11-20 09:15 UTC") {
		t.Fatalf("member tournament dossier: isErr=%v\n%s", isErr, text)
	}

	if text, isErr := callTool(memberCS, "vtes_events", map[string]any{"country": "FR", "rank": "NC", "from": "2026-09-01", "to": "2026-10-01"}); isErr ||
		!strings.Contains(text, "French National Championship") ||
		strings.Contains(text, "Malmö Unbound") {
		t.Fatalf("filtered tournaments: isErr=%v\n%s", isErr, text)
	}

	if text, isErr := callTool(memberCS, "vtes_events", map[string]any{"rankings": true}); isErr ||
		!strings.Contains(text, "1. **VEKN 3200340** (Paris, FR): **180 pts**") ||
		!strings.Contains(text, "2. **VEKN 1002345** (Malmö, SE): **120 pts**") {
		t.Fatalf("rankings: isErr=%v\n%s", isErr, text)
	}

	if text, isErr := callTool(memberCS, "vtes_events", map[string]any{"member": "3200340"}); isErr ||
		!strings.Contains(text, "# Member: VEKN 3200340") ||
		!strings.Contains(text, "rank #1 of 2 worldwide, #1 of 1 in FR") ||
		!strings.Contains(text, "1 known win(s):") ||
		!strings.Contains(text, "Malmö Unbound") ||
		!strings.Contains(text, "019f6a65-01d4-7214-bcc9-4e46534b9d62") ||
		!strings.Contains(text, "Haqim Royalty") {
		t.Fatalf("member lookup: isErr=%v\n%s", isErr, text)
	}

	if text, isErr := callTool(memberCS, "vtes_events", map[string]any{"league": "Nordic"}); isErr ||
		!strings.Contains(text, "# League: Nordic Chapter 2026") ||
		!strings.Contains(text, "Malmö Unbound") {
		t.Fatalf("league lookup: isErr=%v\n%s", isErr, text)
	}

	if text, isErr := callTool(memberCS, "vtes_events", map[string]any{"links": true, "country": "SE"}); isErr ||
		!strings.Contains(text, "VEKN Sweden") ||
		!strings.Contains(text, "VEKN 1002345") {
		t.Fatalf("member community links: isErr=%v\n%s", isErr, text)
	}

	for _, target := range []string{"019f6a65-01d4-7214-bcc9-4e46534b9d62", "MU-2606", "https://archon.vekn.net/tournaments/MU-2606"} {
		if text, isErr := callTool(memberCS, "vtes_deck", map[string]any{"deck": target}); isErr ||
			!strings.Contains(text, "# Deck: Malmö Haqim") ||
			!strings.Contains(text, "Archon deck 019f6a65-01d4-7214-bcc9-4e46534b9d62: won by VEKN 3200340 at Malmö Unbound (MU-2606)") ||
			!strings.Contains(text, "Archon winner comment.") ||
			!strings.Contains(text, "Haqim Royalty") {
			t.Fatalf("vtes_deck(%q): isErr=%v\n%s", target, isErr, text)
		}
	}

	if text, isErr := callTool(memberCS, "vtes_meta", map[string]any{}); isErr ||
		!strings.Contains(text, "## Full-field play against win share (Archon)") ||
		!strings.Contains(text, "1 of 2 finished Archon tournaments (13 players in all) published decklists beyond the winner: 2 published decks (1 winners, 1 other decks)") ||
		!strings.Contains(text, "| 1 | **Haqim Royalty** (C1) | 50.0% | 1 | 100.0% | 1 | 100.0% |") ||
		!strings.Contains(text, "| | near no archetype | 50.0% | 1 | 0.0% | 0 | 0.0% |") ||
		!strings.Contains(text, "archon.vekn.net export 2026-11-20 09:15 UTC") {
		t.Fatalf("member vtes_meta: isErr=%v\n%s", isErr, text)
	}
	if text, isErr := callTool(memberCS, "vtes_meta", map[string]any{"archetype": "Haqim Royalty"}); isErr ||
		!strings.Contains(text, "Full-field (Archon, 1 events with decklists beyond the winner, 2 published decks, 1 wins): play share 50.0% (1 decks), win share 100.0% (1 wins), conversion 100.0%.") ||
		!strings.Contains(text, "Its published Archon decks in the window, of 2:") ||
		!strings.Contains(text, "019f6a65-01d4-7214-bcc9-4e46534b9d62, Malmö Haqim: won by VEKN 3200340 at Malmö Unbound (MU-2606)") ||
		!strings.Contains(text, "019f6a65-2222-7214-bcc9-4e46534b9d62, Paris Haqim: won by VEKN 1002345 at Paris Spring Open (PS-2605)") {
		t.Fatalf("member vtes_meta(Haqim Royalty): isErr=%v\n%s", isErr, text)
	}
	if text, isErr := callTool(memberCS, "vtes_meta", map[string]any{"country": "SE"}); isErr ||
		!strings.Contains(text, "(filtered to Sweden (SE), online events excluded)") ||
		!strings.Contains(text, "1 of 1 finished Archon tournaments (13 players in all) published decklists beyond the winner") {
		t.Fatalf("member vtes_meta(country=SE): isErr=%v\n%s", isErr, text)
	}
	if text, isErr := callTool(memberCS, "vtes_meta", map[string]any{"country": "France"}); isErr ||
		!strings.Contains(text, "(filtered to France (FR), online events excluded)") ||
		!strings.Contains(text, "Of 1 finished Archon tournaments in the window, none published decklists beyond the winner.") {
		t.Fatalf("member vtes_meta(country=France): isErr=%v\n%s", isErr, text)
	}
	if text, isErr := callTool(memberCS, "vtes_meta", map[string]any{"continent": "Europe"}); isErr ||
		!strings.Contains(text, "(filtered to Europe, online events excluded)") ||
		!strings.Contains(text, "1 of 2 finished Archon tournaments (13 players in all) published decklists beyond the winner") {
		t.Fatalf("member vtes_meta(continent=Europe): isErr=%v\n%s", isErr, text)
	}
	if text, isErr := callTool(memberCS, "vtes_meta", map[string]any{"continent": "Asia"}); isErr ||
		!strings.Contains(text, "No finished Archon tournaments in the window match Asia.") {
		t.Fatalf("member vtes_meta(continent=Asia): isErr=%v\n%s", isErr, text)
	}
	if text, isErr := callTool(memberCS, "vtes_meta", map[string]any{"archetype": "Haqim Royalty", "country": "SE"}); isErr ||
		!strings.Contains(text, "Its published Archon decks in the window, of 1:") ||
		!strings.Contains(text, "019f6a65-01d4-7214-bcc9-4e46534b9d62") ||
		strings.Contains(text, "019f6a65-2222-7214-bcc9-4e46534b9d62") {
		t.Fatalf("member vtes_meta(Haqim Royalty, country=SE): isErr=%v\n%s", isErr, text)
	}
	if text, isErr := callTool(memberCS, "vtes_meta", map[string]any{"archetype": "Haqim Royalty", "continent": "Asia"}); isErr ||
		!strings.Contains(text, "Full-field (Archon): no finished tournaments in the window match Asia.") {
		t.Fatalf("member vtes_meta(Haqim Royalty, continent=Asia): isErr=%v\n%s", isErr, text)
	}
	// Online Archon events are excluded when a geographic filter is set, even if country is populated.
	withOnline := append(append([]string{}, lines...),
		`{"type":"tournament","data":{"uid":"019fe999-3333-7549-8f9a-70dcc09a03d9","name":"Nordic Online Open","format":"Standard","rank":"","online":true,"start":"2026-06-20T10:00:00","finish":"2026-06-20T18:00:00","timezone":"Europe/Stockholm","country":"SE","state":"Finished","event_code":"NO-2606","winner":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","reported_player_count":15}}`,
		fmt.Sprintf(`{"type":"deck","data":{"uid":"019f6a65-3333-7214-bcc9-4e46534b9d62","tournament_uid":"019fe999-3333-7549-8f9a-70dcc09a03d9","user_uid":"019f6a5e-d1e6-7347-965e-4dac09bf03a4","name":"Online Haqim","cards":%s,"attribution":{"kind":"Owner"},"winner":true}}`, cardsJSON),
		`{"type":"deck","data":{"uid":"019f6a65-4444-7214-bcc9-4e46534b9d62","tournament_uid":"019fe999-3333-7549-8f9a-70dcc09a03d9","user_uid":"019f6a5e-d5dd-74b9-9aab-00919d4de7dc","name":"Online Finalist","cards":{"100001":1,"200001":1},"attribution":{"kind":"Owner"},"winner":false}}`,
	)
	up.setExport(2, "2026-11-20T09:16:00.000000", withOnline, len(withOnline))
	r.refreshArchon(ctx, now.Add(time.Minute))
	if text, isErr := callTool(memberCS, "vtes_meta", map[string]any{}); isErr ||
		!strings.Contains(text, "2 of 3 finished Archon tournaments (28 players in all)") {
		t.Fatalf("unfiltered vtes_meta with online event: isErr=%v\n%s", isErr, text)
	}
	if text, isErr := callTool(memberCS, "vtes_meta", map[string]any{"country": "SE"}); isErr ||
		!strings.Contains(text, "1 of 1 finished Archon tournaments (13 players in all)") {
		t.Fatalf("country=SE vtes_meta should exclude online event: isErr=%v\n%s", isErr, text)
	}
	up.setExport(2, "2026-11-20T09:15:00.000000", lines, len(lines))
	r.refreshArchon(ctx, now.Add(2*time.Minute))

	ts, err := oauthHandler.TokenSource(ctx)
	if err != nil || ts == nil {
		t.Fatalf("no token source after login: %v", err)
	}
	tok, err := ts.Token()
	if err != nil {
		t.Fatal(err)
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
