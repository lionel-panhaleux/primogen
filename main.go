// Command primogen is the VTES councilor's MCP server and its data tooling.
package main

import (
	"cmp"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
	"github.com/lionel-panhaleux/primogen/internal/tools"
)

const usage = `usage:
  primogen build [-bundle PATH] [-knowledge DIR]   build the data bundle (default: data/bundle.json.gz)
  primogen query [-bundle PATH] [-detail] TOOL ARG...   call a tool and print what the model reads:
      cards NAME...              vtes_cards
      case QUESTION [CARD...]    vtes_rules_case
      search WORDS...            vtes_rules_search
      read ID...                 vtes_rules_read
      archive WORDS...           vtes_rulings_archive
  primogen stdio [-bundle PATH]                    serve the tools over stdio, for a local MCP client
  primogen serve -bundle PATH [-addr ADDR] [-knowledge DIR]   serve the tools over HTTP at /mcp, refreshing the bundle

Without -bundle, query and stdio use the published bundle in the user cache
directory: downloaded on first run (from -from, for the tests), and again when
a newer one exists, checked at most daily. The rulings archive is downloaded
the same way (from -archive) when a tool first needs it: vtes_rulings_archive,
or vtes_rules_case given two cards no ruling names together.`

// published is the bundle the hosted server builds, which local installs download.
const published = "https://mcp.krcg.org/bundle.json.gz"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	err := run(ctx, os.Args[1:], os.Stdout, os.Stderr)
	if err != nil && !errors.Is(err, flag.ErrHelp) {
		fmt.Fprintln(os.Stderr, "primogen:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer) error {
	if len(args) == 0 {
		return errors.New(usage)
	}
	cmd, args := args[0], args[1:]
	fs := flag.NewFlagSet(cmd, flag.ContinueOnError)
	fs.SetOutput(stderr)
	path := fs.String("bundle", "", "the data bundle")
	switch cmd {
	case "build":
		return build(ctx, fs, path, args, stdout, stderr)
	case "query":
		return query(ctx, fs, path, args, stdout, stderr)
	case "stdio":
		return stdio(ctx, fs, path, args, stderr)
	case "serve":
		return serve(ctx, fs, path, args)
	}
	return errors.New(usage)
}

func build(ctx context.Context, fs *flag.FlagSet, path *string, args []string, stdout, stderr io.Writer) error {
	knowledge := fs.String("knowledge", "knowledge", "the knowledge directory")
	if err := fs.Parse(args); err != nil {
		return err
	}
	p := cmp.Or(*path, filepath.Join("data", "bundle.json.gz"))
	up := bundle.KRCG
	up.KnowledgeDir = *knowledge
	archive, err := bundle.FetchArchive(ctx, up)
	if err != nil {
		return err
	}
	b, warnings, err := bundle.Build(ctx, up, archive)
	if err != nil {
		return err
	}
	for _, w := range warnings {
		fmt.Fprintln(stderr, "needs a person:", w)
	}
	if err := b.Save(p); err != nil {
		return err
	}
	fmt.Fprintf(stdout, "%d cards written to %s; query it with -bundle %[2]s\n", len(b.Cards), p)
	return nil
}

func query(ctx context.Context, fs *flag.FlagSet, path *string, args []string, stdout, stderr io.Writer) error {
	detail := fs.Bool("detail", false, "the detailed result")
	from := fs.String("from", published, "where to download the bundle when -bundle is not given")
	archive := fs.String("archive", bundle.KRCG.Archive, "where to download the rulings archive")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() < 2 {
		return errors.New(usage)
	}
	rest := fs.Args()[1:]
	var tool string
	in := map[string]any{"detail": *detail}
	switch fs.Arg(0) {
	case "cards":
		tool, in["names"] = "vtes_cards", rest
	case "case":
		tool, in["question"], in["cards"] = "vtes_rules_case", rest[0], rest[1:]
	case "search":
		tool, in["query"] = "vtes_rules_search", strings.Join(rest, " ")
	case "read":
		tool, in = "vtes_rules_read", map[string]any{"ids": rest}
	case "archive":
		tool, in["query"] = "vtes_rulings_archive", strings.Join(rest, " ")
	default:
		return errors.New(usage)
	}
	live, err := localLive(ctx, *path, *from, *archive, stderr)
	if err != nil {
		return err
	}
	text, err := tools.Call(ctx, tools.NewServer(live), tool, in)
	if err != nil {
		return err
	}
	fmt.Fprintln(stdout, text)
	return nil
}

func stdio(ctx context.Context, fs *flag.FlagSet, path *string, args []string, stderr io.Writer) error {
	from := fs.String("from", published, "where to download the bundle when -bundle is not given")
	archive := fs.String("archive", bundle.KRCG.Archive, "where to download the rulings archive")
	if err := fs.Parse(args); err != nil {
		return err
	}
	live, err := localLive(ctx, *path, *from, *archive, stderr)
	if err != nil {
		return err
	}
	return tools.NewServer(live).Run(ctx, &mcp.StdioTransport{})
}

// localLive serves the bundle localBundle answers, and the rulings archive
// downloaded into the user cache directory on first use.
func localLive(ctx context.Context, flagged, from, archive string, stderr io.Writer) (*tools.Live, error) {
	b, err := localBundle(ctx, flagged, from, stderr)
	if err != nil {
		return nil, err
	}
	live := tools.NewLive(b)
	live.LoadArchive = func(ctx context.Context) (*bundle.Archive, error) {
		p, err := installed("archive.json.gz")
		if err != nil {
			return nil, err
		}
		return cached(ctx, archive, p, readArchive, stderr)
	}
	return live, nil
}

func readArchive(path string) (*bundle.Archive, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return bundle.ReadArchive(f)
}

// localBundle loads the bundle given with -bundle, or else the local
// install's.
func localBundle(ctx context.Context, flagged, from string, stderr io.Writer) (*bundle.Bundle, error) {
	if flagged != "" {
		b, err := bundle.Load(flagged)
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("no data bundle at %s: run `primogen build` first", flagged)
		}
		return b, err
	}
	p, err := installed("bundle.json.gz")
	if err != nil {
		return nil, err
	}
	b, err := cached(ctx, from, p, bundle.Load, stderr)
	if err != nil {
		return nil, fmt.Errorf("from a checkout, run `primogen build` and pass -bundle data/bundle.json.gz: %w", err)
	}
	return b, nil
}

// cached loads a file of the local install, checking for a newer one at most
// daily. Offline, it serves what it has.
func cached[T any](ctx context.Context, from, p string, load func(string) (T, error), stderr io.Writer) (T, error) {
	stamp := p + ".checked"
	if info, err := os.Stat(stamp); err == nil && time.Since(info.ModTime()) < 24*time.Hour {
		if v, err := load(p); err == nil {
			return v, nil
		}
	}
	v, fresh, err := download(ctx, from, p, load)
	if err == nil {
		if err := os.WriteFile(stamp, nil, 0o644); err != nil {
			fmt.Fprintln(stderr, "primogen:", err)
		}
		if fresh {
			return v, nil
		}
	}
	old, loadErr := load(p)
	if loadErr != nil {
		return old, errors.Join(fmt.Errorf("nothing at %s", p), err, loadErr)
	}
	if err != nil {
		info, _ := os.Stat(p)
		fmt.Fprintf(stderr, "primogen: could not check for a newer %s, serving the one of %s: %v\n", filepath.Base(p), info.ModTime().Format(time.DateOnly), err)
	}
	return old, nil
}

// download fetches a published file into path; fresh is false when the copy
// there is as new.
func download[T any](ctx context.Context, u, path string, load func(string) (T, error)) (v T, fresh bool, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return v, false, err
	}
	req.Header.Set("User-Agent", "primogen/"+tools.Version())
	if info, err := os.Stat(path); err == nil {
		req.Header.Set("If-Modified-Since", info.ModTime().UTC().Format(http.TimeFormat))
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return v, false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotModified {
		return v, false, nil
	}
	if resp.StatusCode != http.StatusOK {
		return v, false, fmt.Errorf("%s: %s", u, resp.Status)
	}
	modified, err := http.ParseTime(resp.Header.Get("Last-Modified"))
	if err != nil {
		return v, false, fmt.Errorf("%s: no Last-Modified: %w", u, err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return v, false, err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".download-*")
	if err != nil {
		return v, false, err
	}
	defer os.Remove(tmp.Name())
	_, err = io.Copy(tmp, resp.Body)
	if err := errors.Join(err, tmp.Close()); err != nil {
		return v, false, err
	}
	if v, err = load(tmp.Name()); err != nil {
		return v, false, err
	}
	// the file's time is the server's, which the next check sends back
	if err := os.Chtimes(tmp.Name(), modified, modified); err != nil {
		return v, false, err
	}
	return v, true, os.Rename(tmp.Name(), path)
}

// installed is a file of the local install, which only the download writes.
func installed(name string) (string, error) {
	cache, err := os.UserCacheDir()
	return filepath.Join(cache, "primogen", name), err
}
