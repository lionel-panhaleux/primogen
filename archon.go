package main

import (
	"cmp"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/oauthex"

	"github.com/lionel-panhaleux/primogen/internal/tools"
)

type archonApp struct {
	siteURL, apiURL        string
	clientID, clientSecret string
	key                    [32]byte
	client                 *http.Client

	tokenMu sync.Mutex
	token   string
	expires time.Time

	mu    sync.Mutex
	flows map[string]oauthFlow
	codes map[string]oauthCode
}

type registeredClient struct {
	RedirectURIs []string `json:"uris"`
	AuthMethod   string   `json:"auth"`
}

type oauthFlow struct {
	clientID, redirectURI    string
	clientState, challenge   string
	archonVerifier, callback string
	expires                  time.Time
}

type oauthCode struct {
	clientID, redirectURI string
	challenge             string
	uid, veknID           string
	expires               time.Time
}

type tokenClaims struct {
	Type     string `json:"type"`
	ClientID string `json:"cid,omitempty"`
	UID      string `json:"sub"`
	VEKNID   string `json:"vekn"`
	Exp      int64  `json:"exp"`
}

func newArchonApp(siteURL, apiURL, clientID, secretPath string) (*archonApp, error) {
	raw, err := os.ReadFile(secretPath)
	if err != nil {
		return nil, err
	}
	secret := strings.TrimSpace(string(raw))
	if secret == "" {
		return nil, fmt.Errorf("%s: empty secret", secretPath)
	}
	return &archonApp{
		siteURL:      strings.TrimRight(siteURL, "/"),
		apiURL:       strings.TrimRight(apiURL, "/"),
		clientID:     clientID,
		clientSecret: secret,
		key:          sha256.Sum256([]byte("primogen-archon:" + clientID + ":" + secret)),
		client:       &http.Client{Timeout: 60 * time.Second},
		flows:        map[string]oauthFlow{},
		codes:        map[string]oauthCode{},
	}, nil
}

func (a *archonApp) export(ctx context.Context) (*tools.Archon, error) {
	tok, err := a.appToken(ctx)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.apiURL+"/v1/export", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tok)
	req.Header.Set("User-Agent", "primogen/"+tools.Version())
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		a.tokenMu.Lock()
		a.token = ""
		a.tokenMu.Unlock()
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Archon /v1/export answered %s", resp.Status)
	}
	return tools.ReadArchon(resp.Body)
}

func (a *archonApp) appToken(ctx context.Context) (string, error) {
	a.tokenMu.Lock()
	defer a.tokenMu.Unlock()
	if a.token != "" && time.Until(a.expires) > 5*time.Minute {
		return a.token, nil
	}
	tok, err := a.postToken(ctx, a.apiURL+"/oauth/token", url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {a.clientID},
		"client_secret": {a.clientSecret},
	})
	if err != nil {
		return "", err
	}
	a.token = tok.AccessToken
	a.expires = time.Now().Add(time.Duration(cmp.Or(tok.ExpiresIn, 3600)) * time.Second)
	return a.token, nil
}

type archonTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (a *archonApp) postToken(ctx context.Context, endpoint string, form url.Values) (archonTokenResponse, error) {
	var tok archonTokenResponse
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return tok, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "primogen/"+tools.Version())
	resp, err := a.client.Do(req)
	if err != nil {
		return tok, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var e struct {
			Detail string `json:"detail"`
		}
		json.NewDecoder(resp.Body).Decode(&e)
		return tok, fmt.Errorf("Archon %s answered %s: %s", endpoint, resp.Status, e.Detail)
	}
	err = json.NewDecoder(resp.Body).Decode(&tok)
	return tok, err
}

func (a *archonApp) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/.well-known/oauth-protected-resource", a.protectedResource)
	mux.HandleFunc("/.well-known/oauth-protected-resource/", a.protectedResource)
	mux.HandleFunc("/.well-known/oauth-authorization-server", a.authServer)
	mux.HandleFunc("/oauth/register", a.registerClient)
	mux.HandleFunc("GET /oauth/authorize", a.authorize)
	mux.HandleFunc("GET /oauth/callback", a.callback)
	mux.HandleFunc("/oauth/token", a.tokenEndpoint)
}

func (a *archonApp) wrapMCP(mcpHandler http.Handler) (root, member http.Handler) {
	member = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.requireBearer(mcpHandler, w, r)
	})
	root = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" || r.URL.Query().Get("login") == "1" {
			a.requireBearer(mcpHandler, w, r)
			return
		}
		mcpHandler.ServeHTTP(w, r)
	})
	return root, member
}

func (a *archonApp) requireBearer(h http.Handler, w http.ResponseWriter, r *http.Request) {
	metaURL := baseURL(r) + "/.well-known/oauth-protected-resource" + r.URL.Path
	if r.URL.RawQuery != "" {
		metaURL += "?" + r.URL.RawQuery
	}
	auth.RequireBearerToken(a.verifyBearer, &auth.RequireBearerTokenOptions{
		ResourceMetadataURL: metaURL,
		Scopes:              []string{"member"},
	})(h).ServeHTTP(w, r)
}

func (a *archonApp) verifyBearer(_ context.Context, token string, _ *http.Request) (*auth.TokenInfo, error) {
	var claims tokenClaims
	if err := a.unseal(token, &claims); err != nil || claims.Type != "access" || claims.VEKNID == "" {
		return nil, auth.ErrInvalidToken
	}
	return &auth.TokenInfo{
		UserID:     claims.VEKNID,
		Scopes:     []string{"member"},
		Expiration: time.Unix(claims.Exp, 0),
	}, nil
}

func (a *archonApp) protectedResource(w http.ResponseWriter, r *http.Request) {
	base := baseURL(r)
	resource := base
	if p := strings.TrimPrefix(r.URL.Path, "/.well-known/oauth-protected-resource"); p != "" && p != "/" {
		resource = base + p
	}
	if r.URL.RawQuery != "" {
		resource += "?" + r.URL.RawQuery
	}
	auth.ProtectedResourceMetadataHandler(&oauthex.ProtectedResourceMetadata{
		Resource:               resource,
		AuthorizationServers:   []string{base},
		ScopesSupported:        []string{"member"},
		BearerMethodsSupported: []string{"header"},
		ResourceName:           "Primogen",
	}).ServeHTTP(w, r)
}

func (a *archonApp) authServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	base := baseURL(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"issuer":                                base,
		"authorization_endpoint":                base + "/oauth/authorize",
		"token_endpoint":                        base + "/oauth/token",
		"registration_endpoint":                 base + "/oauth/register",
		"scopes_supported":                      []string{"member"},
		"response_types_supported":              []string{"code"},
		"grant_types_supported":                 []string{"authorization_code", "refresh_token"},
		"token_endpoint_auth_methods_supported": []string{"none", "client_secret_post", "client_secret_basic"},
		"code_challenge_methods_supported":      []string{"S256"},
	})
}

func (a *archonApp) registerClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 64<<10)
	var meta oauthex.ClientRegistrationMetadata
	if err := json.NewDecoder(r.Body).Decode(&meta); err != nil || len(meta.RedirectURIs) == 0 {
		oauthError(w, http.StatusBadRequest, "invalid_client_metadata", "redirect_uris is required")
		return
	}
	for _, raw := range meta.RedirectURIs {
		if !validRedirectURI(raw) {
			oauthError(w, http.StatusBadRequest, "invalid_redirect_uri", "invalid redirect_uri")
			return
		}
	}
	meta.TokenEndpointAuthMethod = cmp.Or(meta.TokenEndpointAuthMethod, "none")
	if !slices.Contains([]string{"none", "client_secret_post", "client_secret_basic"}, meta.TokenEndpointAuthMethod) {
		oauthError(w, http.StatusBadRequest, "invalid_client_metadata", "unsupported token_endpoint_auth_method")
		return
	}
	if len(meta.GrantTypes) == 0 {
		meta.GrantTypes = []string{"authorization_code", "refresh_token"}
	}
	if len(meta.ResponseTypes) == 0 {
		meta.ResponseTypes = []string{"code"}
	}
	clientID := a.seal(registeredClient{RedirectURIs: meta.RedirectURIs, AuthMethod: meta.TokenEndpointAuthMethod})
	secret := ""
	if meta.TokenEndpointAuthMethod != "none" {
		secret = a.mac(clientID)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&oauthex.ClientRegistrationResponse{
		ClientRegistrationMetadata: meta,
		ClientID:                   clientID,
		ClientSecret:               secret,
		ClientIDIssuedAt:           time.Now(),
	})
}

func validRedirectURI(raw string) bool {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Fragment != "" {
		return false
	}
	scheme := strings.ToLower(u.Scheme)
	if slices.Contains([]string{"javascript", "data", "vbscript"}, scheme) {
		return false
	}
	if scheme == "https" {
		return u.Host != ""
	}
	if scheme == "http" {
		host := u.Hostname()
		if host == "localhost" {
			return true
		}
		ip := net.ParseIP(host)
		return ip != nil && ip.IsLoopback()
	}
	return true
}

func (a *archonApp) authorize(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	clientID := q.Get("client_id")
	redirectURI := q.Get("redirect_uri")
	var reg registeredClient
	if a.unseal(clientID, &reg) != nil || !slices.Contains(reg.RedirectURIs, redirectURI) {
		http.Error(w, "invalid client_id or redirect_uri", http.StatusBadRequest)
		return
	}
	if q.Get("response_type") != "code" || q.Get("code_challenge") == "" || q.Get("code_challenge_method") != "S256" {
		http.Error(w, "response_type=code and PKCE S256 are required", http.StatusBadRequest)
		return
	}
	verifier := randomToken()
	flowID := randomToken()
	callback := baseURL(r) + "/oauth/callback"
	now := time.Now()
	a.mu.Lock()
	for k, f := range a.flows {
		if now.After(f.expires) {
			delete(a.flows, k)
		}
	}
	a.flows[flowID] = oauthFlow{
		clientID:       clientID,
		redirectURI:    redirectURI,
		clientState:    q.Get("state"),
		challenge:      q.Get("code_challenge"),
		archonVerifier: verifier,
		callback:       callback,
		expires:        now.Add(10 * time.Minute),
	}
	a.mu.Unlock()
	archonQuery := url.Values{
		"response_type":         {"code"},
		"client_id":             {a.clientID},
		"redirect_uri":          {callback},
		"scope":                 {"profile:read"},
		"state":                 {flowID},
		"code_challenge":        {s256(verifier)},
		"code_challenge_method": {"S256"},
	}
	http.Redirect(w, r, a.siteURL+"/consent?"+archonQuery.Encode(), http.StatusFound)
}

func (a *archonApp) callback(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	now := time.Now()
	a.mu.Lock()
	flow, ok := a.flows[q.Get("state")]
	delete(a.flows, q.Get("state"))
	a.mu.Unlock()
	if !ok || now.After(flow.expires) {
		http.Error(w, "invalid or expired authorization state", http.StatusBadRequest)
		return
	}
	redirectClient := func(params url.Values) {
		u, _ := url.Parse(flow.redirectURI)
		out := u.Query()
		for k, vs := range params {
			out[k] = vs
		}
		if flow.clientState != "" {
			out.Set("state", flow.clientState)
		}
		u.RawQuery = out.Encode()
		http.Redirect(w, r, u.String(), http.StatusFound)
	}
	if errCode := q.Get("error"); errCode != "" {
		redirectClient(url.Values{"error": {"access_denied"}, "error_description": {errCode}})
		return
	}
	code := q.Get("code")
	if code == "" {
		redirectClient(url.Values{"error": {"invalid_request"}, "error_description": {"no code from Archon"}})
		return
	}
	tok, err := a.postToken(r.Context(), a.siteURL+"/oauth/token", url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {a.clientID},
		"client_secret": {a.clientSecret},
		"code":          {code},
		"redirect_uri":  {flow.callback},
		"code_verifier": {flow.archonVerifier},
	})
	if err != nil {
		redirectClient(url.Values{"error": {"server_error"}, "error_description": {err.Error()}})
		return
	}
	info, err := a.userinfo(r.Context(), tok.AccessToken)
	if err != nil {
		redirectClient(url.Values{"error": {"server_error"}, "error_description": {err.Error()}})
		return
	}
	if info.VEKNID == "" {
		redirectClient(url.Values{"error": {"access_denied"}, "error_description": {"this Archon account has no VEKN ID: VEKN membership is required"}})
		return
	}
	issuedCode := randomToken()
	a.mu.Lock()
	for k, c := range a.codes {
		if now.After(c.expires) {
			delete(a.codes, k)
		}
	}
	a.codes[issuedCode] = oauthCode{
		clientID:    flow.clientID,
		redirectURI: flow.redirectURI,
		challenge:   flow.challenge,
		uid:         info.Sub,
		veknID:      info.VEKNID,
		expires:     now.Add(time.Minute),
	}
	a.mu.Unlock()
	redirectClient(url.Values{"code": {issuedCode}})
}

type archonUserInfo struct {
	Sub    string `json:"sub"`
	VEKNID string `json:"vekn_id"`
}

func (a *archonApp) userinfo(ctx context.Context, accessToken string) (archonUserInfo, error) {
	var info archonUserInfo
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.siteURL+"/oauth/userinfo", nil)
	if err != nil {
		return info, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("User-Agent", "primogen/"+tools.Version())
	resp, err := a.client.Do(req)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return info, fmt.Errorf("Archon /oauth/userinfo answered %s", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&info)
	return info, err
}

func (a *archonApp) tokenEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Cache-Control", "no-store")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 64<<10)
	if err := r.ParseForm(); err != nil {
		oauthError(w, http.StatusBadRequest, "invalid_request", err.Error())
		return
	}
	clientID, ok := a.authenticateClient(w, r)
	if !ok {
		return
	}
	now := time.Now()
	switch r.PostFormValue("grant_type") {
	case "authorization_code":
		code := r.PostFormValue("code")
		verifier := r.PostFormValue("code_verifier")
		a.mu.Lock()
		c, ok := a.codes[code]
		delete(a.codes, code)
		a.mu.Unlock()
		if !ok || now.After(c.expires) || c.clientID != clientID || c.redirectURI != r.PostFormValue("redirect_uri") || verifier == "" || s256(verifier) != c.challenge {
			oauthError(w, http.StatusBadRequest, "invalid_grant", "invalid code, redirect_uri or PKCE verifier")
			return
		}
		a.writeTokens(w, clientID, c.uid, c.veknID, now)
	case "refresh_token":
		var claims tokenClaims
		if err := a.unseal(r.PostFormValue("refresh_token"), &claims); err != nil || claims.Type != "refresh" || claims.Exp <= now.Unix() || claims.VEKNID == "" || claims.ClientID != clientID {
			oauthError(w, http.StatusBadRequest, "invalid_grant", "invalid or expired refresh_token")
			return
		}
		a.writeTokens(w, clientID, claims.UID, claims.VEKNID, now)
	default:
		oauthError(w, http.StatusBadRequest, "unsupported_grant_type", "only authorization_code and refresh_token are supported")
	}
}

func (a *archonApp) authenticateClient(w http.ResponseWriter, r *http.Request) (string, bool) {
	hasBasic := strings.HasPrefix(strings.ToLower(r.Header.Get("Authorization")), "basic ")
	formID, formSecret := r.PostFormValue("client_id"), r.PostFormValue("client_secret")
	if hasBasic && (formID != "" || formSecret != "") {
		oauthError(w, http.StatusBadRequest, "invalid_request", "multiple client authentication methods")
		return "", false
	}
	clientID, secret := formID, formSecret
	if hasBasic {
		rawID, rawSecret, ok := r.BasicAuth()
		id, err1 := url.QueryUnescape(rawID)
		sec, err2 := url.QueryUnescape(rawSecret)
		if !ok || err1 != nil || err2 != nil {
			w.Header().Set("WWW-Authenticate", "Basic")
			oauthError(w, http.StatusUnauthorized, "invalid_client", "malformed Basic credentials")
			return "", false
		}
		clientID, secret = id, sec
	}
	var reg registeredClient
	if err := a.unseal(clientID, &reg); err != nil {
		if hasBasic {
			w.Header().Set("WWW-Authenticate", "Basic")
		}
		oauthError(w, http.StatusUnauthorized, "invalid_client", "invalid client_id")
		return "", false
	}
	switch reg.AuthMethod {
	case "none":
		if hasBasic || secret != "" {
			oauthError(w, http.StatusBadRequest, "invalid_client", "public client cannot authenticate with a secret")
			return "", false
		}
	case "client_secret_basic":
		if !hasBasic || !hmac.Equal([]byte(secret), []byte(a.mac(clientID))) {
			w.Header().Set("WWW-Authenticate", "Basic")
			oauthError(w, http.StatusUnauthorized, "invalid_client", "invalid client credentials")
			return "", false
		}
	case "client_secret_post":
		if hasBasic || !hmac.Equal([]byte(secret), []byte(a.mac(clientID))) {
			oauthError(w, http.StatusUnauthorized, "invalid_client", "invalid client credentials")
			return "", false
		}
	default:
		oauthError(w, http.StatusUnauthorized, "invalid_client", "invalid client authentication method")
		return "", false
	}
	return clientID, true
}

func (a *archonApp) writeTokens(w http.ResponseWriter, clientID, uid, veknID string, now time.Time) {
	access := a.seal(tokenClaims{Type: "access", UID: uid, VEKNID: veknID, Exp: now.Add(time.Hour).Unix()})
	refresh := a.seal(tokenClaims{Type: "refresh", ClientID: clientID, UID: uid, VEKNID: veknID, Exp: now.Add(30 * 24 * time.Hour).Unix()})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"access_token":  access,
		"refresh_token": refresh,
		"token_type":    "Bearer",
		"expires_in":    3600,
		"scope":         "member",
	})
}

func (a *archonApp) seal(v any) string {
	raw, _ := json.Marshal(v)
	payload := base64.RawURLEncoding.EncodeToString(raw)
	return payload + "." + a.mac(payload)
}

func (a *archonApp) unseal(tok string, out any) error {
	payload, sig, ok := strings.Cut(tok, ".")
	if !ok || !hmac.Equal([]byte(sig), []byte(a.mac(payload))) {
		return errors.New("bad signature")
	}
	raw, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, out)
}

func (a *archonApp) mac(s string) string {
	m := hmac.New(sha256.New, a.key[:])
	m.Write([]byte(s))
	return base64.RawURLEncoding.EncodeToString(m.Sum(nil))
}

func randomToken() string {
	var b [32]byte
	rand.Read(b[:])
	return base64.RawURLEncoding.EncodeToString(b[:])
}

func s256(verifier string) string {
	sum := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func baseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" || r.Host == "mcp.krcg.org" {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}

func oauthError(w http.ResponseWriter, status int, code, desc string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": code, "error_description": desc})
}
