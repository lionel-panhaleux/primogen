// Command primogen is the VTES councilor's MCP server and its data tooling.
package main

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
  primogen audit [-bundle PATH] [-archive PATH|URL] [-since YYYY-MM-DD] [-twda] [-out PATH] [CARD...]   list uncited sitting rules directors' posts by card and card pair for the rulemongers
  primogen query [-bundle PATH] [-detail] [-directors] [-draft] [-months N] TOOL ARG...   call a tool and print what the model reads:
      cards NAME...              vtes_cards
      case QUESTION [CARD...]    vtes_rules_case
      search WORDS...            vtes_rules_search
      read ID...                 vtes_rules_read
      archive WORDS...           vtes_rulings_archive (-directors: the sitting rules directors' posts only)
      deck LIST|ID|LINK|-        vtes_deck: a decklist (its lines as one or several arguments, or - for stdin), a TWDA id or a deck link (VDB, VTES Decks, Amaranth); -draft: a deck idea's cards
      stats CARD [ARCHETYPE]     vtes_card_stats: a card's play rates, companions and competitors, overall or within an archetype
      meta [ARCHETYPE]           vtes_meta: the archetypes' shares of the latest window, or one archetype's (-months: the window's length; -country, -continent: geographic filter)
      events [TOURNAMENT]        vtes_events: tournaments, standings, rankings (-rankings, -category), a member (-member), leagues (-league) or community links (-links)
      feedback QUESTION WRONG EXPECTED   vtes_feedback: file a report on an answer as a public GitHub issue, through the hosted server
  primogen stdio [-bundle PATH]                    serve the tools over stdio, for a local MCP client
  primogen serve -bundle PATH [-addr ADDR] [-knowledge DIR] [-github-app ID -github-key PEM] [-archon-client ID -archon-secret FILE]   serve the tools over HTTP at /mcp, refreshing the bundle

Without -bundle, query and stdio use the published bundle in the user cache
directory: downloaded on first run (from -from, for the tests), and again when
a newer one exists, checked at most daily. The rulings archive is downloaded
the same way (from -archive) when a tool first needs it: vtes_rulings_archive,
or vtes_rules_case given a single card or two cards no ruling names together.`

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
	case "audit":
		return audit(ctx, fs, path, args, stdout, stderr)
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

func audit(ctx context.Context, fs *flag.FlagSet, path *string, args []string, stdout, stderr io.Writer) error {
	from := fs.String("from", published, "where to download the bundle when -bundle is not given")
	archive := fs.String("archive", bundle.KRCG.Archive, "path or URL of the rulings archive")
	since := fs.String("since", "", "only posts on or after YYYY-MM-DD")
	twda := fs.Bool("twda", false, "only cards played in the TWDA window")
	outPath := fs.String("out", "", "write the report to this file instead of stdout")
	if err := fs.Parse(args); err != nil {
		return err
	}
	b, err := localBundle(ctx, *path, *from, stderr)
	if err != nil {
		return err
	}
	var a *bundle.Archive
	if strings.HasPrefix(*archive, "http://") || strings.HasPrefix(*archive, "https://") {
		p, err := installed("archive.json.gz")
		if err != nil {
			return err
		}
		a, err = cached(ctx, *archive, p, readArchive, stderr)
		if err != nil {
			return err
		}
	} else {
		a, err = readArchive(*archive)
		if err != nil {
			return err
		}
	}
	report, err := tools.Audit(b, a, tools.AuditOptions{
		Since:    *since,
		TWDAOnly: *twda,
		Cards:    fs.Args(),
	})
	if err != nil {
		return err
	}
	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(report), 0o644); err != nil {
			return err
		}
		fmt.Fprintf(stderr, "wrote %d bytes to %s\n", len(report), *outPath)
		return nil
	}
	_, err = io.WriteString(stdout, report)
	return err
}

func query(ctx context.Context, fs *flag.FlagSet, path *string, args []string, stdout, stderr io.Writer) error {
	detail := fs.Bool("detail", false, "the detailed result")
	draft := fs.Bool("draft", false, "deck: a deck idea's cards, its sizes not checked")
	librarySize := fs.Int("library-size", 0, "deck: with -draft, the library's size the odds are computed over")
	together := fs.String("together", "", `deck: recipes as JSON, e.g. '[[["Hunger of Marduk"], ["Blur", "Quickness"]]]'`)
	directors := fs.Bool("directors", false, "archive: only posts by the rules director sitting when they were written")
	from := fs.String("from", published, "where to download the bundle when -bundle is not given")
	archive := fs.String("archive", bundle.KRCG.Archive, "where to download the rulings archive")
	months := fs.Int("months", 0, "meta: the window's length in months")
	member := fs.String("member", "", "events: a member's VEKN ID or Archon uid")
	rankings := fs.Bool("rankings", false, "events: list the rating rankings")
	category := fs.String("category", "", "events: rating category (constructed_offline, constructed_online, limited_offline, limited_online)")
	league := fs.String("league", "", "events: a league's Archon uid or words of its name, or all")
	links := fs.Bool("links", false, "events: list community links")
	country := fs.String("country", "", "events/meta: ISO 3166-1 alpha-2 country code (or country name for meta)")
	continent := fs.String("continent", "", "meta: continent (Europe, North America, South America, Asia, Oceania, Africa)")
	since := fs.String("since", "", "events: start date lower bound (YYYY-MM-DD)")
	until := fs.String("until", "", "events: start date upper bound (YYYY-MM-DD)")
	rank := fs.String("rank", "", "events: championship level (NC, CC, basic)")
	format := fs.String("format", "", "events: tournament format (Standard, V5, Limited, Storyline)")
	state := fs.String("state", "", "events: tournament state (Planned, Registration, Waiting, Playing, Finished)")
	sites := fs.String("sites", "", "where the deckbuilders' APIs are read, each under its host, instead of https://")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() < 2 && fs.Arg(0) != "meta" && fs.Arg(0) != "events" {
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
		tool, in["query"], in["directors"] = "vtes_rulings_archive", strings.Join(rest, " "), *directors
	case "deck":
		list := strings.Join(rest, "\n")
		if list == "-" {
			raw, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			list = string(raw)
		}
		tool, in["deck"], in["draft"] = "vtes_deck", list, *draft
		if *librarySize != 0 {
			in["library_size"] = *librarySize
		}
		if *together != "" {
			var recipes [][][]string
			if err := json.Unmarshal([]byte(*together), &recipes); err != nil {
				return fmt.Errorf("-together: %w", err)
			}
			in["together"] = recipes
		}
	case "meta":
		tool = "vtes_meta"
		if *months != 0 {
			in["months"] = *months
		}
		if *country != "" {
			in["country"] = *country
		}
		if *continent != "" {
			in["continent"] = *continent
		}
		if len(rest) > 0 {
			in["archetype"] = strings.Join(rest, " ")
		}
	case "events":
		tool = "vtes_events"
		if len(rest) > 0 {
			in["tournament"] = strings.Join(rest, " ")
		}
		for k, v := range map[string]string{
			"member": *member, "category": *category, "league": *league,
			"country": *country, "from": *since, "to": *until,
			"rank": *rank, "format": *format, "state": *state,
		} {
			if v != "" {
				in[k] = v
			}
		}
		if *rankings {
			in["rankings"] = true
		}
		if *links {
			in["links"] = true
		}
	case "feedback":
		if len(rest) != 3 {
			return errors.New(usage)
		}
		tool, in = "vtes_feedback", map[string]any{"question": rest[0], "wrong": rest[1], "expected": rest[2], "public": true}
	case "stats":
		tool, in["card"] = "vtes_card_stats", rest[0]
		if len(rest) > 1 {
			in["archetype"] = strings.Join(rest[1:], " ")
		}
	default:
		return errors.New(usage)
	}
	live, err := localLive(ctx, *path, *from, *archive, stderr)
	if err != nil {
		return err
	}
	live.Sites = *sites
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
	live.File = forward(from)
	live.LoadArchive = func(ctx context.Context) (*bundle.Archive, error) {
		p, err := installed("archive.json.gz")
		if err != nil {
			return nil, err
		}
		return cached(ctx, archive, p, readArchive, stderr)
	}
	return live, nil
}

// forward files a report through the hosted server publishing the bundle,
// which holds the GitHub App and the limits.
func forward(published string) tools.Filer {
	return func(ctx context.Context, r tools.Report, _ string) (string, error) {
		link, err := forwarded(ctx, published, r)
		if err != nil && !strings.HasPrefix(err.Error(), "not filed: ") {
			err = fmt.Errorf("not filed: %w", err)
		}
		return link, err
	}
}

func forwarded(ctx context.Context, published string, r tools.Report) (string, error) {
	u, err := url.Parse(published)
	if err != nil {
		return "", err
	}
	u.Path, u.RawQuery = "/mcp", ""
	// a retried call could file the report twice
	cs, err := mcp.NewClient(&mcp.Implementation{Name: "primogen", Version: tools.Version()}, nil).Connect(ctx, &mcp.StreamableClientTransport{Endpoint: u.String(), MaxRetries: -1}, nil)
	if err != nil {
		return "", fmt.Errorf("%s does not answer: %w", u, err)
	}
	defer cs.Close()
	res, err := cs.CallTool(ctx, &mcp.CallToolParams{Name: "vtes_feedback", Arguments: r})
	if err != nil {
		return "", err
	}
	if res.IsError {
		var text strings.Builder
		for _, c := range res.Content {
			if t, ok := c.(*mcp.TextContent); ok {
				text.WriteString(t.Text)
			}
		}
		return "", errors.New(text.String())
	}
	var filed tools.Filed
	raw, _ := json.Marshal(res.StructuredContent)
	if err := json.Unmarshal(raw, &filed); err != nil || filed.Issue == "" {
		return "", fmt.Errorf("%s answered no issue's link: %s", u, raw)
	}
	return filed.Issue, nil
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
