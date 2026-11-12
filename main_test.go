package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

// testBundle is built once, from the real upstreams pinned in testdata/upstream.
var (
	testBundle   string
	testWarnings []string
)

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "primogen-test")
	if err == nil {
		err = buildTestBundle(filepath.Join(dir, "bundle.json.gz"))
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}

func buildTestBundle(path string) error {
	u, up, err := newUpstream()
	if err != nil {
		return err
	}
	defer u.srv.Close()
	archive, err := bundle.FetchArchive(context.Background(), up)
	if err != nil {
		return err
	}
	b, warnings, err := bundle.Build(context.Background(), up, archive)
	if err != nil {
		return err
	}
	testBundle, testWarnings = path, warnings
	return b.Save(path)
}

// upstream serves testdata/upstream as the sources do, the card data
// replaceable to stage an upstream change.
type upstream struct {
	srv       *httptest.Server
	mu        sync.Mutex
	cards     []byte
	modified  time.Time
	knowledge string // the commit of primogen's main, its files those of knowledge/
	archives  int    // fetches of the archive's bulk file
}

func newUpstream() (*upstream, bundle.Upstream, error) {
	cards, err := os.ReadFile("testdata/upstream/vtes.json")
	u := &upstream{cards: cards, modified: time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC), knowledge: strings.Repeat("a", 40)}
	u.srv = httptest.NewServer(u)
	srv := u.srv
	return u, bundle.Upstream{
		Cards:         srv.URL + "/vtes.json",
		RulingsHead:   srv.URL + "/commit",
		RulingsFile:   srv.URL + "/vtes-rulings/%s/%s",
		Archive:       srv.URL + "/archive.json.gz",
		KnowledgeDir:  "knowledge",
		ArchonCommits: srv.URL + "/archon-commits?path=%s",
		ArchonFile:    srv.URL + "/archon-vibe/%s/%s",
	}, err
}

func (u *upstream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.mu.Lock()
	defer u.mu.Unlock()
	switch r.URL.Path {
	case "/vtes.json":
		http.ServeContent(w, r, "vtes.json", u.modified, bytes.NewReader(u.cards))
		return
	case "/archive.json.gz":
		u.archives++
	case "/primogen-commit":
		io.WriteString(w, u.knowledge)
		return
	case "/archon-commits":
		// the last commit touching the documents, or the renderer
		pin := "archon-commit"
		if r.URL.Query().Get("path") == "frontend/src/lib/markdown.ts" {
			pin = "archon-renderer"
		}
		sha, err := os.ReadFile(filepath.Join("testdata/upstream", pin))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, `[{"sha": %q}]`, strings.TrimSpace(string(sha)))
		return
	}
	if name, ok := strings.CutPrefix(r.URL.Path, "/primogen/"+u.knowledge+"/knowledge/"); ok {
		http.ServeFile(w, r, filepath.Join("knowledge", name))
		return
	}
	http.FileServer(http.Dir("testdata/upstream")).ServeHTTP(w, r)
}

func (u *upstream) set(cards []byte, modified time.Time) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.cards, u.modified = cards, modified
}

func (u *upstream) push(commit string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.knowledge = commit
}

func TestQueryCards(t *testing.T) {
	for _, tc := range []struct {
		args      []string // after query -bundle
		want, not []string
		err       string
	}{
		{args: []string{"cards", "Freak Drive"}, want: []string{
			"## Freak Drive",
			"> [FOR] Only usable if the action was blocked. Unlock this vampire.",
			"Requires [for].",
			"Cost: 1 blood.",
			// the rulings site's own ids for the card's four rulings, two of them group rulings
			"https://rulings.krcg.org/index.html?uid=100788#r-KHOYNPJM",
			"https://rulings.krcg.org/index.html?uid=100788#r-ROKWE7BA",
			"https://rulings.krcg.org/index.html?uid=100788#r-NKLCIFEB",
			"https://rulings.krcg.org/index.html?uid=100788#r-KFKPOACQ",
			"LSJ 19981028: https://usenet.krcg.org/t/JZHY1bTCAa0/#m1",
			"ANK 20190425: https://usenet.krcg.org/t/vekn-77581/#94652 (archived from https://www.vekn.net/",
			"Changes: no line of the history of card changes names this card.",
			"Answer from these texts and rulings, quoting them and citing their links, never from memory.",
			"Data: cards of static.krcg.org ",
		}, not: []string{"Printings:"}},
		{args: []string{"cards", "Lord of Serenity"}, want: []string{
			"- 2018-07-19, RTR 20180719 (https://www.blackchantry.com/2018/07/18/rules-team-rulings-rtr-19-07-2018/): Gaining blood is now optional.\n",
		}},
		{args: []string{"-detail", "cards", "Freak Drive"}, want: []string{"Printings: Jyhad 1994-08-16", "Translation (fr), Pulsion monstrueuse:"}},
		{args: []string{"cards", "Pulsion monstrueuse"}, want: []string{"## Freak Drive", "#r-KHOYNPJM"}},
		{args: []string{"cards", "Freak Drvie"}, want: []string{`No card is named "Freak Drvie"; this is Freak Drive`, "## Freak Drive"}},
		{args: []string{"cards", "Blood Doll Supreme"}, want: []string{`No card is named "Blood Doll Supreme". Closest names: Blood Doll;`}, not: []string{"## ", "Answer from these texts"}},
		// Misdirection is "Distraction" in French
		{args: []string{"cards", "Distraction"}, want: []string{"## Distraction"}, not: []string{"## Misdirection"}},
		{args: []string{"cards", "Theo Bell"}, want: []string{"## Theo Bell (G2)\n", "## Theo Bell (G2 ADV)\n", "## Theo Bell (G6)\n"}},
		{args: []string{"cards", "Count Germaine (G4)"}, want: []string{`[the history names "Count Germaine", which several cards share`}},
		// the 1998 entry naming Kalinda predates the G6 card
		{args: []string{"cards", "Kalinda (G6)"}, want: []string{"Changes: no line of the history of card changes names this card."}},
		{args: []string{"cards", "Constanza Vinti"}, want: []string{
			"- 1996-05-31, RTR 19960531 (https://usenet.krcg.org/t/DpvF2Peet9o/#m0): Only gets the bleed bonus if you're bleeding your prey.\n",
		}},
		// a section dated by its month only
		{args: []string{"cards", "Patagia: Flaps Allowing Limited Flight"}, want: []string{
			"- 1995-10, VTES (https://archive.org/details/Duelist7/page/n109): Can only be played once on a given vampire.\n",
		}, not: []string{"1995-11-01"}},
		{args: []string{"cards", "Anson"}, want: []string{"## Anson (G1)", "Rulings: the rulings database has none for this card."}},
		{args: []string{"cards", " "}, err: "give 1 to 10 non-empty card names"},
		{args: []string{"cards", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}, err: "give 1 to 10 non-empty card names"},
	} {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			var stdout, stderr strings.Builder
			err := run(t.Context(), append([]string{"query", "-bundle", testBundle}, tc.args...), &stdout, &stderr)
			if tc.err != "" {
				if err == nil || !strings.Contains(err.Error(), tc.err) {
					t.Fatalf("error %v, want %q", err, tc.err)
				}
				return
			}
			if err != nil {
				t.Fatal(err, stderr.String())
			}
			out := stdout.String()
			for _, w := range tc.want {
				if !strings.Contains(out, w) {
					t.Errorf("missing %q in:\n%s", w, out)
				}
			}
			for _, n := range tc.not {
				if strings.Contains(out, n) {
					t.Errorf("unexpected %q in:\n%s", n, out)
				}
			}
		})
	}
}

// TestQueryRules calls the rules tools as a model does.
func TestQueryRules(t *testing.T) {
	// the ten cards with the most rulings
	b, err := bundle.Load(testBundle)
	if err != nil {
		t.Fatal(err)
	}
	cards := slices.Clone(b.Cards)
	slices.SortStableFunc(cards, func(x, y bundle.Card) int { return len(y.Rulings) - len(x.Rulings) })
	var ten []string
	for _, c := range cards[:10] {
		ten = append(ten, c.Name)
	}
	u, _, err := newUpstream()
	if err != nil {
		t.Fatal(err)
	}
	defer u.srv.Close()
	t.Setenv("HOME", t.TempDir())
	t.Setenv("XDG_CACHE_HOME", "")
	for _, tc := range []struct {
		args      []string // after query -bundle
		want, not []string
		max       int // the result's bound, in characters
	}{
		{args: []string{"case", "Can I play a Freak Drive after I played superior kiss of ra?"}, want: []string{
			"Found in the question: Freak Drive; The Kiss of Ra.",
			"## Freak Drive", "## The Kiss of Ra",
			"Breakdown (computed from the text above, which governs):\n- Type: Action Modifier",
			"- Superior [FOR]: Only usable if the action was blocked. Unlock this vampire.",
			// a ruling's original post, its author the sitting rules director then
			"Original post: L. Scott Johnson, the sitting rules director, 1998-10-28",
			"# The rules sections",
			"### [advanced:3.5.3] 3.5.3 Ended and failed",
			"- No ruling names Freak Drive and The Kiss of Ra together.\n- No post by the rules director then sitting names two of these cards within a few sentences either",
			"# Procedure: a rules answer",
			"rules of archon-vibe 4bb83ba",
		}},
		// a rules director's post naming two cards no ruling names together, the archive fetched at the first call
		{args: []string{"case", "If you use Foreshadowing Destruction and your prey is over ten, can you add a Conditioning ?"}, want: []string{
			"# The rules directors' posts on cards no ruling names together",
			"Post naming Foreshadowing Destruction and Conditioning: Vincent Ripoll (Ankha), the sitting rules director, 2019-10-24, \"About bleed modifiers\", https://usenet.krcg.org/t/vekn-78044/#m3\n",
			"https://usenet.krcg.org/t/vekn-78044/#m13\n",
			"- No ruling names Foreshadowing Destruction and Conditioning together. 4 posts by the rules director then sitting name them within a few sentences: all above.\n",
		}},
		{args: []string{"case", "Can a minion block twice?"}, want: []string{"No card given or found in the question", "### [rulebook:who-may-attempt-to-block]"}},
		// a discipline is no card, and cards given are the cards
		{args: []string{"case", "Can a vampire with superior Fortitude play Freak Drive after an action blocked by a vampire with Presence and The Kiss of Ra?", "Freak Drive"},
			want: []string{"## Freak Drive", "The question also names The Kiss of Ra, not in cards"}, not: []string{"## Fortitude", "## Presence", "## The Kiss of Ra"}},
		// a card whose name the rulebook stems alike is found; a word opening a sentence is suggested
		{args: []string{"case", "Majesty against Immortal Grapple: who wins? And may my predator block after I play Conditioning?"}, want: []string{
			"Found in the question: Immortal Grapple; Conditioning.", "The question may also name Majesty: pass them in cards if meant.",
		}, not: []string{"## Majesty"}},
		// a card named by the words opening or ending its name (product.md, the vtes_rules_case row)
		{args: []string{"case", "If I use Tangle Atropos on a bleed, can I bleed again ?"}, want: []string{
			"Found in the question: Tangle Atropos' Hand.", "## Tangle Atropos' Hand\n",
		}},
		{args: []string{"case", "I played Concealed Weapon for a Magnum I had in hand but drew an Ivory Bow to replace, can I change my mind and play the Ivory Bow as the concealed weapon ?"}, want: []string{
			"Found in the question: Concealed Weapon; .44 Magnum; Ivory Bow.", "## .44 Magnum\n",
		}},
		// words opening or ending several names, or one word of a name given whole, only suggest
		{args: []string{"case", "Can I play the Jar on an ally?"}, want: []string{
			"The question may also name Canopic Jar; Jar of Skin Eaters; Jar the Soul: pass them in cards if meant.",
		}, not: []string{"## Canopic Jar"}},
		{args: []string{"case", "I play Conditioning, my predator plays Eagle Sight to block me. Is this allowed?"}, not: []string{"## Desert Eagle"}},
		// a part is weighed by its words that are no stopword, a possessive read without its 's
		{args: []string{"case", "Can the malkavian block? Can my Toreador's vote count twice? Can The Mole's ability be used twice?"}, want: []string{
			"Found in the question: The Mole.",
		}, not: []string{"## The Malkavian Seven Miseries", "## Toreador's Bane", "## Mole's Tunneling"}},
		// a word of the game names no card: players' shorthand, a trait printed on vampires, a word the rulebook uses
		{args: []string{"case", "Is that NRA?"}, want: []string{"No card given or found in the question"}, not: []string{"may also name"}},
		{args: []string{"case", "Is the Hunter's action blocked?"}, want: []string{"No card given or found in the question"}, not: []string{"may also name"}},
		{args: []string{"case", "Can the Seraph's bleed be bounced?"}, want: []string{"No card given or found in the question"}, not: []string{"may also name"}},
		{args: []string{"case", "Is the Death of my vampire final?"}, want: []string{"No card given or found in the question"}, not: []string{"may also name"}},
		{args: append([]string{"case", "How do these interact?"}, ten...), want: []string{"## " + ten[9], "less close to the question: vtes_cards gives them all"}, max: 40000},
		{args: append([]string{"-detail", "case", "How do these interact?"}, ten...), want: []string{"## " + ten[9]}, max: 80000},
		{args: []string{"read", "rulebook:4-detailed-turn-sequence", "advanced:3", "advanced:4", "advanced:1", "advanced:2"}, want: []string{"Cut here, too long to read at once"}, max: 40000},
		{args: []string{"search", "block after declining"}, want: []string{"### [advanced:3.3.2] 3.3.2 Declaring and declining", "vtes_rules_read gives a section whole"}},
		{args: []string{"search", "stalling"}, want: []string{"### [judges:4.6] 4.6. Unsportsmanlike Conduct — Stalling", "the judges' guide › 4. Unsportsmanlike Conduct (official)"}},
		{args: []string{"read", "advanced:3.5", "rulebook:nothing"}, want: []string{
			"### [advanced:3.5] 3.5 Action repetition (NRA) and canceled actions",
			"### [advanced:3.5.3] 3.5.3 Ended and failed",
			"Its sources (footnotes):\n[^3-5-",
			`No section has the id "rulebook:nothing"`,
		}, not: []string{"(excerpt:"}},
	} {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			var stdout, stderr strings.Builder
			if err := run(t.Context(), append([]string{"query", "-bundle", testBundle, "-archive", u.srv.URL + "/archive.json.gz"}, tc.args...), &stdout, &stderr); err != nil {
				t.Fatal(err, stderr.String())
			}
			out := stdout.String()
			for _, w := range tc.want {
				if !strings.Contains(out, w) {
					t.Errorf("missing %q in:\n%s", w, out)
				}
			}
			for _, n := range tc.not {
				if strings.Contains(out, n) {
					t.Errorf("unexpected %q", n)
				}
			}
			if tc.max > 0 && len(out) > tc.max {
				t.Errorf("%d characters, over %d", len(out), tc.max)
			}
		})
	}
}

// TestQueryArchive searches the rulings archive as a local install does: the
// archive downloaded on first use, then kept.
func TestQueryArchive(t *testing.T) {
	u, _, err := newUpstream()
	if err != nil {
		t.Fatal(err)
	}
	defer u.srv.Close()
	t.Setenv("HOME", t.TempDir())
	t.Setenv("XDG_CACHE_HOME", "")
	for range 2 {
		var stdout, stderr strings.Builder
		args := []string{"query", "-bundle", testBundle, "-archive", u.srv.URL + "/archive.json.gz", "archive", `"Freak Drive"`, "action"}
		if err := run(t.Context(), args, &stdout, &stderr); err != nil {
			t.Fatal(err, stderr.String())
		}
		for _, w := range []string{"posts hold [\"freak drive\" \"action\"] as words of their own", "the sitting rules director", "https://usenet.krcg.org/t/", "Archive: posts to "} {
			if !strings.Contains(stdout.String(), w) {
				t.Errorf("missing %q in:\n%s", w, stdout.String())
			}
		}
	}
	if u.archives != 1 {
		t.Errorf("%d downloads of the archive by two local runs the same day, want 1", u.archives)
	}
	var stdout, stderr strings.Builder
	if err := run(t.Context(), []string{"query", "-bundle", testBundle, "-archive", u.srv.URL + "/archive.json.gz", "-detail", "archive", "action"}, &stdout, &stderr); err != nil {
		t.Fatal(err, stderr.String())
	}
	if n := stdout.Len(); n > 80000 {
		t.Errorf("a detailed search of %d characters, over 80,000", n)
	}
}

// TestArchonLinks checks a section of the rulebook, the tournament rules or
// the judges' guide links an id archon.vekn.net's page has, by archon's own
// renderer, and that a section's id is its own.
func TestArchonLinks(t *testing.T) {
	b, err := bundle.Load(testBundle)
	if err != nil {
		t.Fatal(err)
	}
	commit, err := os.ReadFile("testdata/upstream/archon-commit")
	if err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join("testdata/upstream/archon-vibe", strings.TrimSpace(string(commit)), "ids.json"))
	if err != nil {
		t.Fatal(err)
	}
	var ids map[string][]string
	if err := json.Unmarshal(raw, &ids); err != nil {
		t.Fatal(err)
	}
	pages := map[string]string{"rules": "vtes-rules", "tournament-rules": "tournament-rules", "judges-guide": "judges-guide"}
	seen := map[string]bool{}
	linked := 0
	for _, s := range b.Sections {
		if page, id, ok := strings.Cut(strings.TrimPrefix(s.URL, "https://archon.vekn.net/help/"), "#"); ok && s.Doc != bundle.DocAdvanced {
			linked++
			if !slices.Contains(ids[pages[page]], id) {
				t.Errorf("%s links %s, an id archon.vekn.net's page does not have", s.ID, s.URL)
			}
		}
		if seen[s.ID] {
			t.Errorf("two sections have the id %s", s.ID)
		}
		seen[s.ID] = true
	}
	if linked < 400 {
		t.Errorf("%d sections link an id of archon.vekn.net, want most of them", linked)
	}
	for _, w := range testWarnings {
		if strings.HasPrefix(w, "archon's renderer") {
			t.Error(w)
		}
	}
}

func TestHistoryFullyFiled(t *testing.T) {
	for _, w := range testWarnings {
		if strings.HasPrefix(w, "card-changes.md") {
			t.Error(w)
		}
	}
}

// TestServe follows the hosted server through its refreshes, as a client and
// a local install see it.
func TestServe(t *testing.T) {
	ctx := t.Context()
	u, up, err := newUpstream()
	if err != nil {
		t.Fatal(err)
	}
	defer u.srv.Close()
	// knowledge/ as the hosted server reads it, from primogen's main on GitHub
	up.KnowledgeDir = ""
	up.KnowledgeHead = u.srv.URL + "/primogen-commit"
	up.KnowledgeFile = u.srv.URL + "/primogen/%s/knowledge/%s"
	r, err := newRefresher(ctx, up, filepath.Join(t.TempDir(), "bundle.json.gz"), slog.New(slog.DiscardHandler))
	if err != nil {
		t.Fatal(err)
	}
	var downloads atomic.Int32
	h := r.handler()
	hosted := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/bundle.json.gz" {
			downloads.Add(1)
		}
		h.ServeHTTP(w, req)
	}))
	defer hosted.Close()

	connect := func() *mcp.ClientSession {
		t.Helper()
		cs, err := mcp.NewClient(&mcp.Implementation{Name: "test"}, nil).Connect(ctx, &mcp.StreamableClientTransport{Endpoint: hosted.URL + "/mcp"}, nil)
		if err != nil {
			t.Fatal(err)
		}
		return cs
	}
	card := func() string {
		t.Helper()
		cs := connect()
		defer cs.Close()
		res, err := cs.CallTool(ctx, &mcp.CallToolParams{Name: "vtes_cards", Arguments: map[string]any{"names": []string{"Freak Drive"}}})
		if err != nil || res.IsError {
			t.Fatal(err, res)
		}
		text := res.Content[0].(*mcp.TextContent).Text
		if !strings.Contains(text, "## Freak Drive") {
			t.Fatalf("no card in:\n%s", text)
		}
		return text
	}
	get := func(path string) string {
		t.Helper()
		resp, err := http.Get(hosted.URL + path)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return string(body)
	}
	expect := func(what, in string, want ...string) {
		t.Helper()
		for _, w := range want {
			if !strings.Contains(in, w) {
				t.Errorf("%s: missing %q in:\n%s", what, w, in)
			}
		}
	}
	// a local install in a fresh home, its copy and its daily stamp aged at will
	t.Setenv("HOME", t.TempDir())
	t.Setenv("XDG_CACHE_HOME", "")
	local := func(from string) (string, string) {
		t.Helper()
		var stdout, stderr strings.Builder
		if err := run(ctx, []string{"query", "-from", from, "cards", "Freak Drive"}, &stdout, &stderr); err != nil {
			t.Fatal(err, stderr.String())
		}
		return stdout.String(), stderr.String()
	}
	age := func() {
		t.Helper()
		p, _ := installed("bundle.json.gz")
		old := time.Now().Add(-48 * time.Hour)
		for _, f := range []string{p, p + ".checked"} {
			if err := os.Chtimes(f, old, old); err != nil {
				t.Fatal(err)
			}
		}
	}
	published := hosted.URL + "/bundle.json.gz"

	expect("first answer", card(), "cards of static.krcg.org 2026-11-01 00:00 UTC", "knowledge aaaaaaa")
	cs := connect()
	tools, err := cs.ListTools(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	cs.Close()
	for _, tool := range tools.Tools {
		var schema struct{ Properties map[string]struct{ Type any } }
		raw, _ := json.Marshal(tool.InputSchema)
		json.Unmarshal(raw, &schema)
		for name, p := range schema.Properties {
			if _, ok := p.Type.(string); !ok {
				t.Errorf("%s: %s has the type %v, not a single type", tool.Name, name, p.Type)
			}
		}
	}
	expect("status page", get("/"), "Data: cards of static.krcg.org 2026-11-01 00:00 UTC", "Last refresh")
	r.refresh(ctx, r.fetched)
	cs = connect()
	res, err := cs.CallTool(ctx, &mcp.CallToolParams{Name: "vtes_rulings_archive", Arguments: map[string]any{"query": "Freak Drive", "directors": true}})
	cs.Close()
	if err != nil || res.IsError {
		t.Fatal(err, res)
	}
	expect("the hosted archive", res.Content[0].(*mcp.TextContent).Text, "the sitting rules director")
	out, _ := local(published)
	expect("first local run", out, "## Freak Drive", "cards of static.krcg.org 2026-11-01 00:00 UTC")
	local(published)
	if n := downloads.Load(); n != 1 {
		t.Errorf("%d downloads by two local runs the same day, want 1", n)
	}

	// an hourly check after a change of the card data, the archive kept
	base := r.fetched
	u.set(u.cards, time.Date(2026, 11, 2, 12, 34, 0, 0, time.UTC))
	r.refresh(ctx, base)
	expect("after a change", card(), "cards of static.krcg.org 2026-11-02 12:34 UTC")
	expect("status page", get("/"), "ok.")
	if u.archives != 1 {
		t.Errorf("%d fetches of the archive before the night, want 1", u.archives)
	}
	age()
	out, _ = local(published)
	expect("local run a day later", out, "cards of static.krcg.org 2026-11-02 12:34 UTC")

	// a push to primogen's main
	u.push(strings.Repeat("b", 40))
	r.refresh(ctx, base)
	expect("after a push", card(), "knowledge bbbbbbb")

	// the nightly fetch of the archive, which did not change: nothing new is published
	night := time.Date(base.Year(), base.Month(), base.Day(), nightly, 1, 0, 0, time.UTC)
	if !night.After(base) {
		night = night.AddDate(0, 0, 1)
	}
	built := r.live.Bundle().Built
	r.refresh(ctx, night)
	if u.archives != 2 {
		t.Errorf("%d fetches of the archive after the night, want 2", u.archives)
	}
	if r.live.Bundle().Built != built {
		t.Error("a new bundle after a night that changed nothing")
	}

	// card data that lost cards is refused, the live bundle kept
	var all []json.RawMessage
	if err := json.Unmarshal(u.cards, &all); err != nil {
		t.Fatal(err)
	}
	one, _ := json.Marshal(all[:1])
	u.set(one, time.Date(2026, 11, 3, 0, 0, 0, 0, time.UTC))
	r.refresh(ctx, night.Add(time.Hour))
	expect("after a refused refresh", card(), "cards of static.krcg.org 2026-11-02 12:34 UTC")
	expect("status page", get("/"), "FAILED, the bundle above keeps serving", "cards, the live one")

	// offline, a local install serves what it has and says how old it is
	age()
	out, errs := local("http://127.0.0.1:1/bundle.json.gz")
	expect("offline local run", out, "## Freak Drive")
	expect("offline local run", errs, "could not check for a newer bundle.json.gz, serving the one of ")
}
