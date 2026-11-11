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

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
	"github.com/lionel-panhaleux/primogen/internal/tools"
)

const usage = `usage:
  primogen build [-bundle PATH] [-knowledge DIR]          build the data bundle from the upstream sources
  primogen query [-bundle PATH] [-detail] cards NAME...   call vtes_cards and print what the model reads
  primogen stdio [-bundle PATH]                           serve the tools over stdio, for a local MCP client`

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := run(ctx, os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "primogen:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer) error {
	if len(args) == 0 {
		return errors.New(usage)
	}
	cache, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	fs.SetOutput(stderr)
	path := fs.String("bundle", filepath.Join(cache, "primogen", "bundle.json.gz"), "the data bundle")
	var knowledge *string
	var detail *bool
	switch args[0] {
	case "build":
		knowledge = fs.String("knowledge", "knowledge", "the knowledge directory")
	case "query":
		detail = fs.Bool("detail", false, "the detailed result")
	case "stdio":
	default:
		return errors.New(usage)
	}
	if err := fs.Parse(args[1:]); err != nil {
		return err
	}

	if args[0] == "build" {
		b, warnings, err := bundle.Build(ctx, *knowledge)
		if err != nil {
			return err
		}
		for _, w := range warnings {
			fmt.Fprintln(stderr, "needs a person:", w)
		}
		if err := b.Save(*path); err != nil {
			return err
		}
		fmt.Fprintf(stdout, "%d cards written to %s\n", len(b.Cards), *path)
		return nil
	}

	b, err := bundle.Load(*path)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("no data bundle at %s: run `primogen build` first", *path)
	}
	if err != nil {
		return err
	}
	s := tools.NewServer(b)
	if args[0] == "stdio" {
		return s.Run(ctx, &mcp.StdioTransport{})
	}
	rest := fs.Args()
	if len(rest) < 2 || rest[0] != "cards" {
		return errors.New(usage)
	}
	text, err := tools.Call(ctx, s, "vtes_cards", map[string]any{"names": rest[1:], "detail": *detail})
	if err != nil {
		return err
	}
	fmt.Fprintln(stdout, text)
	return nil
}
