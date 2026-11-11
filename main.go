// Command primogen is the VTES councilor's MCP server and its data tooling.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
	"github.com/lionel-panhaleux/primogen/internal/tools"
)

const usage = `usage:
  primogen build [-bundle PATH] [-knowledge DIR]          build the data bundle from the upstream sources
  primogen query [-bundle PATH] [-detail] cards NAME...   call vtes_cards and print what the model reads
  primogen stdio [-bundle PATH]                           serve the tools over stdio, for a local MCP client`

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
	path := fs.String("bundle", "", "the data bundle (default: bundle.json.gz in the user cache directory)")
	switch cmd {
	case "build":
		return build(ctx, fs, path, args, stdout, stderr)
	case "query":
		return query(ctx, fs, path, args, stdout)
	case "stdio":
		return stdio(ctx, fs, path, args)
	}
	return errors.New(usage)
}

func build(ctx context.Context, fs *flag.FlagSet, path *string, args []string, stdout, stderr io.Writer) error {
	knowledge := fs.String("knowledge", "knowledge", "the knowledge directory")
	if err := fs.Parse(args); err != nil {
		return err
	}
	p, err := bundlePath(*path)
	if err != nil {
		return err
	}
	b, warnings, err := bundle.Build(ctx, bundle.KRCG, *knowledge)
	if err != nil {
		return err
	}
	for _, w := range warnings {
		fmt.Fprintln(stderr, "needs a person:", w)
	}
	if err := b.Save(p); err != nil {
		return err
	}
	fmt.Fprintf(stdout, "%d cards written to %s\n", len(b.Cards), p)
	return nil
}

func query(ctx context.Context, fs *flag.FlagSet, path *string, args []string, stdout io.Writer) error {
	detail := fs.Bool("detail", false, "the detailed result")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() < 2 || fs.Arg(0) != "cards" {
		return errors.New(usage)
	}
	s, err := server(*path)
	if err != nil {
		return err
	}
	text, err := tools.Call(ctx, s, "vtes_cards", map[string]any{"names": fs.Args()[1:], "detail": *detail})
	if err != nil {
		return err
	}
	fmt.Fprintln(stdout, text)
	return nil
}

func stdio(ctx context.Context, fs *flag.FlagSet, path *string, args []string) error {
	if err := fs.Parse(args); err != nil {
		return err
	}
	s, err := server(*path)
	if err != nil {
		return err
	}
	return s.Run(ctx, &mcp.StdioTransport{})
}

func server(path string) (*mcp.Server, error) {
	p, err := bundlePath(path)
	if err != nil {
		return nil, err
	}
	b, err := bundle.Load(p)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("no data bundle at %s: run `primogen build` first", p)
	}
	if err != nil {
		return nil, err
	}
	return tools.NewServer(b), nil
}

func bundlePath(flagged string) (string, error) {
	if flagged != "" {
		return flagged, nil
	}
	cache, err := os.UserCacheDir()
	return filepath.Join(cache, "primogen", "bundle.json.gz"), err
}
