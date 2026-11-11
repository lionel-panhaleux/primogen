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
	"syscall"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
	"github.com/lionel-panhaleux/primogen/internal/tools"
)

const usage = `usage:
  primogen build [-bundle PATH] [-knowledge DIR]          build the data bundle (default: data/bundle.json.gz)
  primogen query [-bundle PATH] [-detail] cards NAME...   call vtes_cards and print what the model reads
  primogen stdio [-bundle PATH]                           serve the tools over stdio, for a local MCP client
  primogen serve -bundle PATH [-addr ADDR] [-knowledge DIR]   serve the tools over HTTP at /mcp, refreshing the bundle

Without -bundle, query and stdio use the published bundle in the user cache
directory: downloaded on first run (from -from, for the tests), and again when
a newer one exists, checked at most daily.`

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
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() < 2 || fs.Arg(0) != "cards" {
		return errors.New(usage)
	}
	b, err := localBundle(ctx, *path, *from, stderr)
	if err != nil {
		return err
	}
	text, err := tools.Call(ctx, tools.NewServer(tools.NewLive(b)), "vtes_cards", map[string]any{"names": fs.Args()[1:], "detail": *detail})
	if err != nil {
		return err
	}
	fmt.Fprintln(stdout, text)
	return nil
}

func stdio(ctx context.Context, fs *flag.FlagSet, path *string, args []string, stderr io.Writer) error {
	from := fs.String("from", published, "where to download the bundle when -bundle is not given")
	if err := fs.Parse(args); err != nil {
		return err
	}
	b, err := localBundle(ctx, *path, *from, stderr)
	if err != nil {
		return err
	}
	return tools.NewServer(tools.NewLive(b)).Run(ctx, &mcp.StdioTransport{})
}

// localBundle loads the bundle given with -bundle, or else the local
// install's, checking for a newer one at most daily. Offline, it serves what
// it has.
func localBundle(ctx context.Context, flagged, from string, stderr io.Writer) (*bundle.Bundle, error) {
	if flagged != "" {
		b, err := bundle.Load(flagged)
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("no data bundle at %s: run `primogen build` first", flagged)
		}
		return b, err
	}
	p, err := installed()
	if err != nil {
		return nil, err
	}
	stamp := p + ".checked"
	if info, err := os.Stat(stamp); err == nil && time.Since(info.ModTime()) < 24*time.Hour {
		if b, err := bundle.Load(p); err == nil {
			return b, nil
		}
	}
	b, err := download(ctx, from, p)
	if err == nil {
		if err := os.WriteFile(stamp, nil, 0o644); err != nil {
			fmt.Fprintln(stderr, "primogen:", err)
		}
		if b != nil {
			return b, nil
		}
	}
	old, loadErr := bundle.Load(p)
	if loadErr != nil {
		return nil, errors.Join(fmt.Errorf("no data bundle at %s; from a checkout, run `primogen build` and pass -bundle data/bundle.json.gz", p), err, loadErr)
	}
	if err != nil {
		fmt.Fprintf(stderr, "primogen: could not check for a newer bundle, serving the one built %s: %v\n", old.Built.Format(time.DateOnly), err)
	}
	return old, nil
}

// download fetches the published bundle into path, or answers nil when the
// copy there is as new.
func download(ctx context.Context, u, path string) (*bundle.Bundle, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "primogen/"+tools.Version())
	if info, err := os.Stat(path); err == nil {
		req.Header.Set("If-Modified-Since", info.ModTime().UTC().Format(http.TimeFormat))
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotModified {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", u, resp.Status)
	}
	modified, err := http.ParseTime(resp.Header.Get("Last-Modified"))
	if err != nil {
		return nil, fmt.Errorf("%s: no Last-Modified: %w", u, err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".download-*")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmp.Name())
	_, err = io.Copy(tmp, resp.Body)
	if err := errors.Join(err, tmp.Close()); err != nil {
		return nil, err
	}
	b, err := bundle.Load(tmp.Name())
	if err != nil {
		return nil, err
	}
	// the file's time is the server's, which the next check sends back
	if err := os.Chtimes(tmp.Name(), modified, modified); err != nil {
		return nil, err
	}
	return b, os.Rename(tmp.Name(), path)
}

// installed is the local install's bundle, which only the download writes.
func installed() (string, error) {
	cache, err := os.UserCacheDir()
	return filepath.Join(cache, "primogen", "bundle.json.gz"), err
}
