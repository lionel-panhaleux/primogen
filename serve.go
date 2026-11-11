package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
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
	addr := fs.String("addr", "127.0.0.1:8040", "the address to listen on")
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

	mu       sync.Mutex
	checked  time.Time
	failure  error
	warnings []string
}

// newRefresher serves the bundle at path, built first if there is none.
func newRefresher(ctx context.Context, up bundle.Upstream, path string, log *slog.Logger) (*refresher, error) {
	r := &refresher{up: up, path: path, log: log}
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
	if b != nil {
		r.live.Swap(b)
		r.log.Info("refreshed", "data", tools.DataLine(b))
	}
	r.done(now, err)
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
	mux.HandleFunc("GET /{$}", r.status)
	return mux
}

// status is the page a person reads: what is served, and what needs them.
func (r *refresher) status(w http.ResponseWriter, _ *http.Request) {
	r.mu.Lock()
	checked, failure, warnings := r.checked, r.failure, r.warnings
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
