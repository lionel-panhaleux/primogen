package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
	"github.com/lionel-panhaleux/primogen/internal/tools"
)

// nightly is the hour (UTC) after which the archive is fetched again: its
// scheduled sweep runs at 04:17, and GitHub's schedules run late.
const nightly = 6

func serve(ctx context.Context, fs *flag.FlagSet, path *string, args []string) error {
	addr := fs.String("addr", "127.0.0.1:8050", "the address to listen on")
	knowledge := fs.String("knowledge", "", "a local knowledge directory (default: knowledge/ at primogen's main branch on GitHub)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *path == "" {
		return errors.New("serve needs -bundle: the file it keeps current and publishes")
	}
	p := *path
	up := bundle.KRCG
	up.KnowledgeDir = *knowledge
	r, err := newRefresher(ctx, up, p, slog.Default())
	if err != nil {
		return err
	}
	go func() {
		t := time.NewTicker(time.Hour)
		defer t.Stop()
		for now := time.Now(); ctx.Err() == nil; {
			r.refresh(ctx, now)
			select {
			case <-ctx.Done():
			case now = <-t.C:
			}
		}
	}()
	srv := &http.Server{Addr: *addr, Handler: r.handler(), ReadHeaderTimeout: 10 * time.Second}
	go func() {
		<-ctx.Done()
		shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(shutdown)
	}()
	slog.Info("serving", "addr", *addr, "bundle", p)
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// refresher keeps the live bundle current: the cheap sources at each call,
// the archive once a night, and a new bundle only when it passes the checks.
type refresher struct {
	up      bundle.Upstream
	path    string
	log     *slog.Logger
	live    *tools.Live
	archive *bundle.Archive
	fetched time.Time
	links   *http.Client

	mu       sync.Mutex
	checked  time.Time
	failure  error
	warnings []string
	dead     []string
}

// newRefresher serves the bundle at path, built first if there is none.
func newRefresher(ctx context.Context, up bundle.Upstream, path string, log *slog.Logger) (*refresher, error) {
	r := &refresher{up: up, path: path, log: log, links: &http.Client{Timeout: 30 * time.Second}}
	b, err := bundle.Load(path)
	if errors.Is(err, os.ErrNotExist) {
		log.Info("no bundle yet: building one", "path", path)
		now := time.Now()
		b, err = r.build(ctx, now, nil)
		r.checked = now
	}
	if err != nil {
		return nil, err
	}
	r.live = tools.NewLive(b)
	return r, nil
}

func (r *refresher) refresh(ctx context.Context, now time.Time) {
	b, err := r.build(ctx, now, r.live.Bundle())
	if r.archive != nil {
		r.live.SetArchive(r.archive)
	}
	if b != nil {
		r.live.Swap(b)
		r.log.Info("refreshed", "data", tools.DataLine(b))
	}
	dead := r.deadLinks(ctx)
	for _, d := range dead {
		r.log.Warn("needs a person: " + d)
	}
	r.mu.Lock()
	r.dead = dead
	r.mu.Unlock()
	r.done(now, err)
}

var reLink = regexp.MustCompile(`\]\(([^)\s]+)\)`)

// deadLinks names the links of the live catalog that do not answer: a 405 or
// a 401 answers, an endpoint refusing a GET or wanting a token.
func (r *refresher) deadLinks(ctx context.Context) []string {
	var links []string
	for _, m := range reLink.FindAllStringSubmatch(r.live.Bundle().Catalog, -1) {
		if !slices.Contains(links, m[1]) {
			links = append(links, m[1])
		}
	}
	dead := make([]string, len(links))
	var wg sync.WaitGroup
	for i, link := range links {
		wg.Go(func() {
			if u, err := url.Parse(link); err != nil || !u.IsAbs() {
				dead[i] = "not an absolute URL"
				return
			}
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
			if err != nil {
				dead[i] = err.Error()
				return
			}
			req.Header.Set("User-Agent", "primogen")
			resp, err := r.links.Do(req)
			if err != nil {
				dead[i] = err.(*url.Error).Err.Error()
				return
			}
			resp.Body.Close()
			if c := resp.StatusCode; c == http.StatusNotFound || c == http.StatusGone || c >= 500 {
				dead[i] = resp.Status
			}
		})
	}
	wg.Wait()
	var out []string
	for i, d := range dead {
		if d != "" {
			out = append(out, fmt.Sprintf("a link of the catalog, knowledge/llms.txt, does not answer: %s: %s", links[i], d))
		}
	}
	return out
}

func (r *refresher) done(now time.Time, err error) {
	if err != nil {
		r.log.Error("refresh failed, the previous bundle keeps serving", "err", err)
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.checked, r.failure = now, err
}

// build answers a new, saved bundle, or nil when the live one is current. A
// failed fetch of the archive keeps the one held and still reports.
func (r *refresher) build(ctx context.Context, now time.Time, live *bundle.Bundle) (*bundle.Bundle, error) {
	last := time.Date(now.Year(), now.Month(), now.Day(), nightly, 0, 0, 0, time.UTC)
	if now.Before(last) {
		last = last.AddDate(0, 0, -1)
	}
	var stale error
	if r.archive == nil || r.fetched.Before(last) {
		a, err := bundle.FetchArchive(ctx, r.up)
		switch {
		case err == nil:
			r.archive, r.fetched = a, now
		case r.archive == nil:
			return nil, err
		default:
			stale = err
		}
	}
	if live != nil {
		probed, err := bundle.Probe(ctx, r.up)
		if err != nil {
			return nil, errors.Join(stale, err)
		}
		if current(live, append(probed, bundle.Source{Name: bundle.SourceArchive, Version: r.archive.LastPost})) {
			return nil, stale
		}
	}
	b, warnings, err := bundle.Build(ctx, r.up, r.archive)
	if err != nil {
		return nil, errors.Join(stale, err)
	}
	if live != nil && len(b.Cards) < len(live.Cards) {
		return nil, errors.Join(stale, fmt.Errorf("the new bundle has %d cards, the live one %d", len(b.Cards), len(live.Cards)))
	}
	if err := b.Save(r.path); err != nil {
		return nil, errors.Join(stale, err)
	}
	for _, w := range warnings {
		r.log.Warn("needs a person: " + w)
	}
	r.mu.Lock()
	r.warnings = warnings
	r.mu.Unlock()
	return b, stale
}

func current(live *bundle.Bundle, probed []bundle.Source) bool {
	for _, p := range probed {
		found := false
		for _, s := range live.Sources {
			found = found || s == p
		}
		if !found {
			return false
		}
	}
	return true
}

func (r *refresher) handler() http.Handler {
	s := tools.NewServer(r.live)
	mux := http.NewServeMux()
	mux.Handle("/mcp", mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server { return s }, &mcp.StreamableHTTPOptions{
		Stateless:    true,
		JSONResponse: true,
		// behind nginx on the loopback, every request carries the public Host
		DisableLocalhostProtection: true,
	}))
	mux.HandleFunc("GET /bundle.json.gz", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, r.path)
	})
	mux.HandleFunc("GET /{$}", r.catalog)
	mux.HandleFunc("GET /llms.txt", r.catalog)
	mux.HandleFunc("GET /status", r.status)
	return mux
}

func (r *refresher) catalog(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(r.live.Bundle().Catalog))
}

// status is the page a person reads: what is served, and what needs them.
func (r *refresher) status(w http.ResponseWriter, _ *http.Request) {
	r.mu.Lock()
	checked, failure, warnings := r.checked, r.failure, slices.Concat(r.warnings, r.dead)
	r.mu.Unlock()
	var out strings.Builder
	fmt.Fprintf(&out, "Primogen %s, an MCP server for Vampire: The Eternal Struggle: add the /mcp endpoint to your assistant.\n\n", tools.Version())
	fmt.Fprintln(&out, tools.DataLine(r.live.Bundle()))
	switch {
	case checked.IsZero():
		fmt.Fprintln(&out, "No refresh since the server started.")
	case failure != nil:
		fmt.Fprintf(&out, "Last refresh %s: FAILED, the bundle above keeps serving: %v\n", checked.UTC().Format("2006-01-02 15:04 UTC"), failure)
	default:
		fmt.Fprintf(&out, "Last refresh %s: ok.\n", checked.UTC().Format("2006-01-02 15:04 UTC"))
	}
	if len(warnings) > 0 {
		fmt.Fprintf(&out, "\nNeeds a person (%d):\n", len(warnings))
		for _, w := range warnings {
			fmt.Fprintf(&out, "- %s\n", w)
		}
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(out.String()))
}
