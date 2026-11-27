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
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
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
	dead      string // the one link of the catalog that does not answer
}

func newUpstream() (*upstream, bundle.Upstream, error) {
	cards, err := os.ReadFile("testdata/upstream/vtes.json")
	u := &upstream{cards: cards, modified: time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC), knowledge: strings.Repeat("a", 40)}
	u.srv = httptest.NewServer(u)
	srv := u.srv
	return u, bundle.Upstream{
		Cards:         srv.URL + "/vtes.json",
		Sets:          srv.URL + "/expansions.json",
		TWDA:          srv.URL + "/twda.json",
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
	case "/link":
		if r.URL.Query().Get("u") == u.dead {
			http.NotFound(w, r)
		}
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
	if r.URL.Path == "/amaranth.vtes.co.nz/api/deck" {
		r.URL.Path += "/" + r.URL.Query().Get("id")
	}
	if name, ok := strings.CutPrefix(r.URL.Path, "/primogen/"+u.knowledge+"/knowledge/"); ok {
		http.ServeFile(w, r, filepath.Join("knowledge", name))
		return
	}
	if dir, ok := strings.CutPrefix(r.URL.Path, "/primogen-list/"); ok {
		if r.URL.Query().Get("ref") != u.knowledge {
			http.NotFound(w, r)
			return
		}
		entries, err := os.ReadDir(filepath.Join("knowledge", dir))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		var out []map[string]string
		for _, e := range entries {
			if !e.IsDir() {
				out = append(out, map[string]string{"name": e.Name(), "type": "file"})
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(out)
		return
	}
	http.FileServer(http.Dir("testdata/upstream")).ServeHTTP(w, r)
}

func (u *upstream) set(cards []byte, modified time.Time) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.cards, u.modified = cards, modified
}

// links sends every request to the upstream, which answers for every link
// of the catalog but the dead one.
func (u *upstream) links() *http.Client {
	return &http.Client{Transport: roundTrip(func(req *http.Request) (*http.Response, error) {
		out := req.Clone(req.Context())
		out.URL, _ = url.Parse(u.srv.URL + "/link?u=" + url.QueryEscape(req.URL.String()))
		out.Host = ""
		return http.DefaultTransport.RoundTrip(out)
	})}
}

func (u *upstream) kill(link string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.dead = link
}

type roundTrip func(*http.Request) (*http.Response, error)

func (f roundTrip) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

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
			"Answer from these texts and rulings, never from memory; for a rules question, quote them and cite their links.",
			"Data: cards of static.krcg.org ",
		}, not: []string{"Printings:"}},
		{args: []string{"cards", "Lord of Serenity"}, want: []string{
			"- 2018-07-19, RTR 20180719 (https://www.blackchantry.com/2018/07/18/rules-team-rulings-rtr-19-07-2018/): Gaining blood is now optional.\n",
		}},
		{args: []string{"-detail", "cards", "Freak Drive"}, want: []string{"Printings: Jyhad 1994-08-16", "Translation (fr), Pulsion monstrueuse:"}},
		// a printing only in a precon released apart from its set is dated by the precon
		{args: []string{"-detail", "cards", "Backroom Deal"}, want: []string{"> [ACTION MODIFIER] +1 stealth.", "Printings: Fifth Edition (Secrets of New York) 2026-10-31."}},
		{args: []string{"cards", "Aabbt Kindred"}, want: []string{"unless Nefertiti is ready"}},
		// krcg gives these spellings only with the group, or not at all
		{args: []string{"cards", "Lucinde", "Theo Bell (ADV)"}, want: []string{"## Lucinde, Alastor (G3)\n", "## Theo Bell (G2 ADV)\n"}, not: []string{"No card is named"}},
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
		// a card named like a Sabbat Path; The Path of Blood is named like no Path a vampire follows
		{args: []string{"cards", "Path of Death and the Soul", "The Path of Blood"}, want: []string{
			"Not a Path: Path of Death and the Soul is also the name of a Sabbat Path, which a vampire follows (its \"Path:\" line; vtes_rules_read rulebook:paths-of-enlightenment).",
		}, not: []string{"Not a Path: The Path of Blood"}},
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

// herne is TWDA deck 12947 as a player writes it: crypt columns, a Lackey
// tab, a count after the name, a line's comment, accents left out.
const herne = `Deck Name: Herne Gangrel
Author: Jozef Patrnciak

Crypt (12 cards, min=10, max=27, avg=4.92)
2x Massimiliano      7  pro ANI FOR  baron  Gangrel:6
2x Casey Snyder      6  ani cel for PRO  baron  Gangrel:6
2x Kuyén             6  ANI PRO  baron  Gangrel:6
1x Kamile Paukstys   5  ani for PRO  Gangrel:6
1x Hanna Nokelainen 4 ani for pro Gangrel:6
1 Nathan Turner
1 Indira
1x Joaquin de Cadiz
1x Ruslan Fedorenko

Library (90 cards)
Master (25)
3x Vessel
3x Villein
2x Club Illusion
2x Direct Intervention -- the only cancel
2x Garibaldi-Meucci Museum
2x Powerbase: Savannah
1x Anarch Railroad
1x Backways
1x Carfax Abbey
1x Ecoterrorists
1x Ennoia's Theater
1x Information Highway
1x Papillon
1x Powerbase: Los Angeles
1x Powerbase: Montreal
1x The Rack
1x Zoo Hunting Ground
11x Thing
1x Revolutionary Council
5	Raven Spy
1x Mr. Winthrop
Heart of Nizchetus x2
3x Earth Control
1x Monkey Wrench
4x Form of the Bat
9x Organized Resistance
7x Bait and Switch
5x Deep Ecology
3x Delaying Tactics
9x Earth Meld
4x Form of Mist
`

// TestQueryDeck reads one deck by its TWDA id, as text and by the links of the
// deckbuilders, an Amaranth deck, and the lists it must refuse.
func TestQueryDeck(t *testing.T) {
	u, _, err := newUpstream()
	if err != nil {
		t.Fatal(err)
	}
	defer u.srv.Close()
	query := func(args ...string) (string, error) {
		var stdout, stderr strings.Builder
		err := run(t.Context(), append([]string{"query", "-bundle", testBundle, "-sites", u.srv.URL}, args...), &stdout, &stderr)
		return stdout.String(), err
	}
	// the computed facts, from the crypt to the relatives, which leave a TWDA deck out
	facts := func(out string) string {
		i, j := strings.Index(out, "## Crypt"), strings.Index(out, "## Relatives")
		if i < 0 || j < i {
			t.Fatalf("no dossier in:\n%s", out)
		}
		return strings.Replace(out[i:j], " (this deck left out of the counts)", "", 1)
	}
	byID, err := query("deck", "12947")
	if err != nil {
		t.Fatal(err)
	}
	for _, w := range []string{
		"TWDA deck 12947: won by Jozef Patrnciak at Year of the Sabbat Paths IV: Amaranth, Herne, Germany, 2026-10-24, 20 players",
		"Tournament legal today (",
		"## Crypt (12 cards, 9 distinct; capacity 2 to 7, average 4.9; groups G6)",
		"| 2 | Massimiliano (G6) | 201602 | 7 | G6 | Gangrel | pro ANI FOR | Baron |",
		"- [ani] 9: superior 4 (Massimiliano (G6) ×2, Kuyén (G6) ×2); inferior 5 (",
		"## Library (90 cards, 31 distinct)",
		"- 5× **Raven Spy** (card 101550). Requires [ani]. Cost: 1 blood.\n  Crypt cards with the [ani] it requires: 9 of 12, 4 at superior",
		// any discipline of three, and the requirement the card data does not hold
		"Crypt cards with the [ani] or [for] or [pro] it requires (each effect's symbols in the text say which): 12 of 12: Massimiliano (G6) ×2 [ANI FOR pro];",
		`Its text also requires, unchecked here: "Requires an Anarch."`,
		"- 11× **Thing** (card 101972). Clans its requirement names: Gangrel.\n  Crypt cards with the clan Gangrel it requires: all 12.",
		"  > [pro] Strike: combat ends.\n  > [PRO] As above, and unlock this vampire before combat ends.\n",
		// the library by module: a Gangrel wall, as the catalog describes it
		"- **PRO stealth** (Stealth & block denial): 20 copies, 22.2%, strong, 1.6 per hand, in 84% of opening hands; core: Earth Control ×3, Form of Mist ×4, Form of the Bat ×4, Earth Meld ×9\n" +
			"- **Baron defense** (Bleed defense): 17 copies, 18.9%, strong, 1.3 per hand, in 78% of opening hands; core: Bait and Switch ×7, Organized Resistance ×9, Powerbase: Los Angeles; not here, of its core: The Anarch Free Press\n",
		"- **Leech package** (Bloat & pool economy): 9 copies, 10.0%, standard, 0.7 per hand, in 53% of opening hands; core: Villein ×3, Vessel ×3; support: Zoo Hunting Ground, Carfax Abbey, Papillon;",
		"- **ANI intercept** (Wakes, intercept & wall kits): 5 copies, 5.6%, tactical, 0.4 per hand, in 34% of opening hands; core: Raven Spy ×5;",
		"- **Protected-permanent payload** (Wakes, intercept & wall kits): 3 copies, 3.3%, tactical, 0.2 per hand, in 22% of opening hands; core: The Rack, Powerbase: Montreal; support: Monkey Wrench;",
		"- With one library copy or none (one copy: in 8% of opening hands): **Big-cap ramp** (Information Highway); **Permanent intercept** (Mr. Winthrop)\n",
		"In no module shown (23 copies, 25.6% of the library): Thing ×11, Deep Ecology ×5, Powerbase: Savannah ×2, Anarch Railroad, Backways, Ecoterrorists, Ennoia's Theater, Revolutionary Council\n",
		// its relatives, the deck left out of them, and how well the labels classify back
		"pass of 2026-07-05, revised on 2026-11-23 by the same maintainer, curated and revisable; each of their decks left out in turn, 97% are placed back in their archetype, 99% in its family (variants merged).",
		"Nearest archetypes (95 in 100 of an archetype's own decks are at 0.49 or more from it):\n- **Gangrel Resistance** (C11): 0.92; 44 decks, its representative decks 12101.\n- **Gangrel Grinder** (C64), a variant of **Gangrel Resistance**: ",
		"### 12511, Intelligent Gangrel: 0.99\n\nWon by Milos Krstic at Shell Game in Belgrade, Belgrade, Serbia, 2025-11-15, 13 players. Event report: https://www.vekn.net/event-calendar/event/12511 Labelled **Gangrel Resistance** (C11).",
		"Only in this deck: 1× Monkey Wrench; 1× Revolutionary Council.\nOnly in 12511: 1× The Barrens; 1× The Parthenon.\nOther counts: Thing, 11 here and 9 there;",
		"Costs, every copy counted: 22 pool over 15 cards; 11 blood over 11 cards; variable: Monkey Wrench (X blood).",
		"# Procedure: a deck analysis", "# Procedure: advice on a deck or an idea",
		"TWDA of static.krcg.org ",
	} {
		if !strings.Contains(byID, w) {
			t.Errorf("missing %q in:\n%s", w, byID)
		}
	}
	asText, err := query("deck", herne)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(asText, "# Deck: Herne Gangrel\nBy Jozef Patrnciak.") ||
		strings.Replace(facts(asText), "\n  The author: the only cancel", "", 1) != facts(byID) {
		t.Errorf("the list as text reads otherwise than by id:\n%s", asText)
	}
	for _, link := range []string{"https://vdb.im/decks/12947", "https://vdb.im/decks/notwda", "https://vtesdecks.com/deck/tournament-12947", "https://vtesdecks.com/deck/notwda"} {
		linked, err := query("deck", link)
		if err != nil || facts(linked) != facts(byID) || !strings.Contains(linked, "Read from "+link+"\n") {
			t.Errorf("%s reads otherwise than by id: %v\n%s", link, err, linked)
		}
	}
	// a link pasted without its scheme, its host in capitals
	if linked, err := query("deck", "VTESDECKS.COM/deck/notwda"); err != nil || facts(linked) != facts(byID) {
		t.Errorf("a link without its scheme reads otherwise than by id: %v\n%s", err, linked)
	}
	amaranth, err := query("deck", "https://amaranth.vtes.co.nz/#deck/4d3aa426-70da-44b7-8cb7-92377a1a0dbd")
	for _, w := range []string{
		"# Deck: First Blood: Tremere\nBy BCP.\nRead from https://amaranth.vtes.co.nz/#deck/4d3aa426-70da-44b7-8cb7-92377a1a0dbd\n",
		"## Crypt (12 cards, 6 distinct; capacity 5 to 8, average 6.3; groups G4, G5)",
		"| 2 | Muhsin Samir (G4) | 201020 | 6 | G4 | Tremere | aus pot DOM THA |  |",
		"## Library (86 cards, 21 distinct)",
	} {
		if err != nil || !strings.Contains(amaranth, w) {
			t.Errorf("the Amaranth deck: %v, missing %q in:\n%s", err, w, amaranth)
		}
	}
	var stored struct{ Cards map[string]int }
	raw, _ := os.ReadFile("testdata/upstream/vdb.im/api/deck/notwda")
	json.Unmarshal(raw, &stored)
	var pairs []string
	for id, n := range stored.Cards {
		pairs = append(pairs, fmt.Sprintf("%s=%d", id, n))
	}
	inURL, err := query("deck", "https://vdb.im/decks/deck?name=Herne&author=Jozef#"+strings.Join(pairs, ";"))
	if err != nil || facts(inURL) != facts(byID) || !strings.Contains(inURL, "Read from https://vdb.im/decks/deck?name=Herne&author=Jozef\n") {
		t.Errorf("the VDB link carrying the cards reads otherwise than by id: %v\n%s", err, inURL)
	}
	if out, err := query("deck", "https://vdb.im/decks/notpinned"); err == nil || !strings.Contains(err.Error(), `VDB has no public deck "notpinned" (404 Not Found)`) {
		t.Errorf("an unknown VDB deck: %v\n%s", err, out)
	}

	for _, tc := range []struct {
		deck string
		want []string
	}{
		{strings.NewReplacer("Raven Spy", "Ravn Spy", "\n1 Indira", "\n1 Theo Bell", "3x Earth Control", "3x Elemental Stoicism\nVilein\nCombat Shotgn\nAction Modifier (4)").Replace(herne), []string{
			`No card is named "Ravn Spy"; this is Raven Spy, the one closest name.`,
			`"Theo Bell" read as Theo Bell (G6), of the group of the rest of the crypt`,
			`"Vilein" reads as no card with its count: left out`,
			`"Combat Shotgn" reads as no card with its count: left out`,
			// a combination card, one of whose effects needs [for] alone
			"Crypt cards with every discipline one of its effects names ([for]): 7 of 12: Massimiliano (G6) ×2 [FOR];",
			// a list outside the TWDA, and its relatives
			"- **Gangrel Resistance** (C11): 0.85;",
			"### 12947, (unnamed): 0.93\n\nWon by Jozef Patrnciak at Year of the Sabbat Paths IV: Amaranth, Herne, Germany, 2026-10-24, 20 players. Event report: https://archon.vekn.net/t/12947\n" +
				"Only in this deck: 3× Elemental Stoicism; 1× Theo Bell (G6).\nOnly in 12947: 3× Earth Control; 1× Indira (G6).\n",
		}},
		// a deck of an archetype the build proposed, its twin the nearest deck
		{"13019", []string{"- **Zillah + Villein ramp-bloat** (Bloat & pool economy): 6 copies, 6.7%, tactical, 0.5 per hand, in 39% of opening hands; core: Villein ×6;", "This deck is labelled **Agnieszka Wall** (G35).", "- **Matawall** (C53): 0.15;",
			"### 13163, Agnieszka: 1.00\n\nWon by Martin Weinmayer at The Vozhd of JOL, Online, 2026-03-17, 45 players. Event report: https://www.vekn.net/event-calendar/event/13163 Labelled **Agnieszka Wall** (G35).\nThe same list as this deck.\n"}},
		// a deck near no archetype
		{"13020", []string{"- **Weenie Toreador** (C70): 0.28;", "This deck is near no archetype: even the closest is under 0.49."}},
		{"12434", []string{"(cut at 1500 of 6967 characters: detail gives it whole)"}},
		// a deck the TWDA archives, illegal today
		{"zink2k1", []string{"Tournament legal today (", "): no, as archived:\n- Protect Thine Own is banned since 2008-01-01: take it out"}},
		// a card named like a Sabbat Path says so in the library
		{"11979", []string{"**Path of Death and the Soul** (card 101360). Not a Path: "}},
		// the TWDA deck with the most distinct cards keeps its texts
		{"2k6munduko", []string{"## Library (", "  > "}},
	} {
		out, err := query("deck", tc.deck)
		if err != nil {
			t.Fatal(err)
		}
		for _, w := range tc.want {
			if !strings.Contains(out, w) {
				t.Errorf("missing %q in:\n%s", w, out)
			}
		}
		if len(out) > 40000 {
			t.Errorf("%d characters, past the bound of 40,000", len(out))
		}
		// a support card of a module not shown is in no module shown; a star engine needs two of its pieces
		if i := strings.Index(out, "\nIn no module shown"); tc.deck == "2k6munduko" && (i < 0 || !strings.Contains(out[i:strings.Index(out, "## Relatives")], "Ivory Bow")) ||
			tc.deck == "13019" && strings.Contains(out, "- **Enkil Cog + Instantaneous Transformation + Homunculus**") {
			t.Errorf("%s: the modules shown or the cards in none:\n%s", tc.deck, out)
		}
	}
	// a deck idea: its cards, a count optional, no size checked and no density computed
	idea, err := query("-draft", "deck", "Sword of Nuln", "Eulogio Sánchez de los Reyes", "- Morbidity", "Hunter's Mark", "2x Villein")
	if err != nil {
		t.Fatal(err)
	}
	for _, w := range []string{
		"A deck idea: 1 crypt and 5 library cards, a line without a count read as one copy",
		"Tournament legal today (", "), its sizes aside: yes.",
		"Crypt cards with the [obl] or [tha] and path Death and the Soul it requires (each effect's symbols in the text say which): 1 of 1: Eulogio Sánchez de los Reyes (G6) [OBL THA].",
		"- 2× **Villein**",
		// a combination card's effect needing both its disciplines
		"Crypt cards with every discipline one of its effects names ([cel][tha]): none.", "- **Zillah + Villein ramp-bloat** (Bloat & pool economy): 2 copies, in 15% of opening hands; core: Villein ×2",
		"In no module shown (3 copies): Hunter's Mark, Morbidity, Sword of Nuln",
		// the combat posture: an answer to combat ends no vampire here plays at its level
		"## Combat posture against the field", "TWDA decks since 2021-07-01 carry a card giving it",
		"Answered here by Hunter's Mark, denies the strike, needs [CEL][THA]: no crypt card here has it.\n",
		"most often Immortal Grapple (", ". Not answered here.\n",
		"This deck is near no archetype", "This is a deck idea: the advice procedure applies to it", "# Procedure: advice on a deck or an idea",
		// its track record: the Sword in no winning deck since it became legal
		"  - Sword of Nuln: legal since 2003-06-18, ", " years; in no TWDA deck.\n",
		"- A deck idea's share of rarely played copies is not computed: its counts are not a list's.\n",
	} {
		if !strings.Contains(idea, w) {
			t.Errorf("the deck idea: missing %q in:\n%s", w, idea)
		}
	}
	// a Sabbat vampire with no printed Path elects one; a vampire printed with another Path, or not Sabbat, does not
	caine, err := query("-draft", "deck", "Soulgrinder", "Qawiyya el-Ghaduba", "Eulogio Sánchez de los Reyes", "Theo Bell (G2)")
	if err != nil {
		t.Fatal(err)
	}
	for _, w := range []string{
		"and path Caine it requires (each effect's symbols in the text say which): 1 of 3: Qawiyya el-Ghaduba (G5) [POT] electing the Path. Electing the Path: a Sabbat vampire with no printed Path can burn 1 blood during your influence phase to follow a Sabbat Path, and follows only one (vtes_rules_read rulebook:paths-of-enlightenment).",
		"Answered here by Soulgrinder, damage past Fortitude only, needs [POT] or [THA] and path Caine: 1 of 3 crypt cards, 1 of them electing the Path",
	} {
		if !strings.Contains(caine, w) {
			t.Errorf("a Path elected: missing %q in:\n%s", w, caine)
		}
	}
	// a recipe: a Path elected only where no held card is met without it, and one Path for every group
	for _, tc := range []struct{ recipe, want string }{
		{`[[["Soulgrinder"], ["Abuse of Power"]]]`, ": 1 of 3: Qawiyya el-Ghaduba (G5) [POT] electing the Path of Caine. Electing the Path: "},
		{`[[["Soulgrinder", "Freak Drive"]]]`, ": 1 of 3: Qawiyya el-Ghaduba (G5) [FOR POT].\n"},
		{`[[["Abuse of Power"], ["Forward Momentum"]]]`, ": none.\n"},
	} {
		out, err := query("-draft", "-together", tc.recipe, "deck", "Soulgrinder", "Abuse of Power", "Forward Momentum", "Freak Drive", "Qawiyya el-Ghaduba", "Eulogio Sánchez de los Reyes", "Theo Bell (G2)")
		if err != nil || !strings.Contains(out, "gives what)"+tc.want) {
			t.Errorf("a Path elected for %s: want %q in:\n%s (%v)", tc.recipe, tc.want, out, err)
		}
	}
	if strings.Contains(idea[strings.Index(idea, "## Library by module"):strings.Index(idea, "## Relatives")], "per hand") {
		t.Errorf("a density computed for a deck idea:\n%s", idea)
	}
	// the graded werewolf draft (calibration 2026-11-20): its nearest archetype's staples it lacks, defenses first
	entry, err := os.ReadFile("calibration/corpus/2026-11-20-werewolf-gangrel-playgroup.md")
	if err != nil {
		t.Fatal(err)
	}
	_, draft, _ := strings.Cut(string(entry), "Crypt:\n")
	draft, _, _ = strings.Cut(draft, "\n[00:31]")
	wolves, err := query("-draft", "deck", strings.Replace(draft, "\nDeck\n", "\n", 1))
	if err != nil {
		t.Fatal(err)
	}
	i, j := strings.Index(wolves, "- Bait and Switch (Reaction, Baron defense): in "), strings.Index(wolves, "- Rewilding (Action): in ")
	if !strings.Contains(wolves, "Its nearest archetype of 10 decks or more: **Gangrel Renegades** (C8), a variant of **Gangrel Resistance**, at ") ||
		!strings.Contains(wolves, `Its text requires, unchecked here: "Requires a baron."`) || i < 0 || j < i {
		t.Errorf("the werewolf draft's nearest staples:\n%s", wolves[strings.Index(wolves, "## Beside its reference"):strings.Index(wolves, "## Track record")])
	}
	// a list of an archetype too small to give staples is given no other's
	if out, err := query("deck", "12936"); err != nil || !strings.Contains(out, "This list is of **Matawall** (C53), at ") || strings.Contains(out, "Its nearest archetype of 10 decks or more") {
		t.Errorf("a list of a small archetype: %v\n%s", err, out[strings.Index(out, "## Beside its reference"):strings.Index(out, "## Track record")])
	}
	// the graded Banu list: its answers to combat ends and to prevention, with the crypt cards at their level
	banu, err := query("deck", "https://vdb.im/decks/7dcfd9fcb")
	if err != nil {
		t.Fatal(err)
	}
	for _, w := range []string{
		"Answered here by Hunter's Mark ×4, denies the strike, needs [CEL][THA]: 6 of 12 crypt cards; Psyche! ×2, undoes it, combat going on, whatever ended it, needs [CEL]: 9 of 12 crypt cards.\n",
		"Answered here by Hunger of Marduk ×7, steals blood, which is not damage, needs [tha] and clan Banu Haqim: 12 of 12 crypt cards; Blood Fury ×2, damage past Fortitude only, needs [tha]: 12 of 12 crypt cards.\n",
	} {
		if !strings.Contains(banu, w) {
			t.Errorf("the Banu list: missing %q in:\n%s", w, banu)
		}
	}
	// beside its reference decks (product.md, the vtes_deck row): the budgets, the sources, the concentration, the staples, the cards few play
	for _, w := range []string{
		"The reference: its nearest archetype, **Haqim Royalty** (C1), its decks and its variants' among the 1494 TWDA decks since 2021-07-01: 41 decks.",
		"- Cards needing a minion's action (action, political action, equipment, ally, retainer): 26, 31 per 100 library cards; action 14; political action 4; equipment 8. Reference: 6, 8, 11, 15, 18, 20, 24;",
		"- Steal (blood, life or pool taken from another Methuselah's minion or pool): 7, 8 per 100 library cards; Hunger of Marduk ×7. Reference: 0, 10, 10, 12, 14, 16, 20;",
		"- Distinct library cards: 35, 41 per 100 library cards. Reference: 18, 20, 22, 25, 29, 32, 38;",
		"How common an outlier is: of the 41 reference decks, each set against the others, 7 (17.1%)",
		"- Priority Contract: in 40 of 41 reference decks, median 3 copies, lower quartile 2; here 1.\n",
		"- Blur: in 38 of 41 reference decks, median 4 copies, lower quartile 3; here 0.\n",
		"This list holds 13 library cards (20 copies) a tenth or fewer of the reference decks play; each reference deck against the others: 1, 2, 4, 6, 7",
		"2× Blood Fury (in 0)",
	} {
		if !strings.Contains(banu, w) {
			t.Errorf("the Banu list beside its reference: missing %q in:\n%s", w, banu[strings.Index(banu, "## Beside its reference"):strings.Index(banu, "## Relatives")])
		}
	}
	if !strings.Contains(banu, "- Every card of this deck is played by at least one of these decks.\n") {
		t.Errorf("the Banu list's track record:\n%s", banu[strings.Index(banu, "## Track record"):strings.Index(banu, "## Relatives")])
	}
	// the odds of a hand: Hunger of Marduk with an additional strike, in the Banu list and in a winner of its
	// archetype; two copies of one card (1 - (C(80,7) + 10·C(80,6)) / C(90,7)); a draft's library taken as 90 cards
	marduk := `[[["Hunger of Marduk"], ["Blur", "Lightning Reflexes", "Quickness", "Side Strike", "Pursuit"]]]`
	for _, tc := range []struct {
		args []string
		want []string
	}{
		{[]string{"-together", marduk, "deck", "https://vdb.im/decks/7dcfd9fcb"}, []string{
			"1. One of each: (Hunger of Marduk ×7) + (",
			"   In an opening hand: 22.0%; seen within 14 cards 56.0%, 21 cards 79.1%, 28 cards 91.3%.\n",
			"held here requires (the requirement, not the level an effect needs: their levels shown, the texts above say which level gives what): 11 of 12: ", // Arjun Shah has no [cel]
			"Warmaksan (G6) ×2 [cel THA]",
			"- **Marduk theft-at-range** (Combat): ",
		}},
		{[]string{"-together", marduk, "deck", "12494"}, []string{
			"   In an opening hand: 36.9%; seen within 14 cards 75.1%, 21 cards 91.8%, 28 cards 97.6%.\n",
			"Not in this deck: Lightning Reflexes, Quickness, Side Strike.",
		}},
		{[]string{"-together", `[[["Hunger of Marduk"], ["Hunger of Marduk"]]]`, "deck", "12494"}, []string{
			"   In an opening hand: 17.3%; seen within 14 cards 48.4%,",
			"- **Marduk theft-at-range** (Combat): 27 copies, 30.0%, spam, 2.1 per hand, in 93% of opening hands;",
		}},
		{[]string{"-draft", "-together", `[[["Sword of Nuln"], ["Villein"]]]`, "deck", "Sword of Nuln", "2x Villein"}, []string{
			" This deck idea's library taken as 90 cards; library_size gives another.",
			"   In an opening hand: 1.0%;",
			"   No crypt given: who can play every piece is unchecked.\n",
		}},
		// a draft of a crypt alone: the library still taken as 90 cards
		{[]string{"-draft", "-together", `[[["Blur"], ["Pursuit"]]]`, "deck", "Arjun Shah"}, []string{
			" This deck idea's library taken as 90 cards; library_size gives another.",
			"   In an opening hand: 0.0%;",
		}},
		{[]string{"-draft", "-library-size", "60", "-together", `[[["Sword of Nuln"]]]`, "deck", "Sword of Nuln"}, []string{
			"Exact odds over a library of 60 cards. The library's size given for this deck idea.",
			"   In an opening hand: 11.7%;",
		}},
		// a piece the deck lacks fills no group, for the odds and for who can play it
		{[]string{"-together", `[[["Hunger of Marduk"], ["Pursuit", "Lost in Crowds"]]]`, "deck", "https://vdb.im/decks/7dcfd9fcb"}, []string{
			"Not in this deck: Lost in Crowds.", "): 11 of 12: ",
		}},
	} {
		out, err := query(tc.args...)
		if err != nil {
			t.Fatal(err)
		}
		for _, w := range tc.want {
			if !strings.Contains(out, w) {
				t.Errorf("odds %v: missing %q in:\n%s", tc.args, w, out[strings.Index(out, "## Library by module"):strings.Index(out, "## Combat posture")])
			}
		}
	}
	if _, err := query("-together", `[[["Hunger of Marduk"], ["Kassandra Tassaki"]]]`, "deck", "12494"); err == nil || !strings.Contains(err.Error(), "is a crypt card: a recipe holds library cards") {
		t.Errorf("a crypt card in a recipe: %v", err)
	}
	// two precons from the card data's set information: an old one past the window's 90th percentile of rarely
	// played cards, a recent one within its reference's range
	var precons []struct {
		PrintedName string `json:"printed_name"`
		Suffix      string `json:"suffix"`
		Prints      []struct {
			Set struct {
				Code string `json:"code"`
			} `json:"set"`
			Occurrences []struct {
				Type   string `json:"type"`
				Bundle string `json:"bundle"`
				Copies int    `json:"copies"`
			} `json:"occurrences"`
		} `json:"prints"`
	}
	cardData, err := os.ReadFile("testdata/upstream/vtes.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(cardData, &precons); err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ set, bundle, want string }{
		{"KoT", "PV", "The reference: all the 1494 TWDA decks since 2021-07-01: no archetype of 10 decks or more"},
		{"V5", "PM", "The reference: its nearest archetype, **Malk 22** (C3), its decks and its variants'"},
	} {
		var list []string
		for _, c := range precons {
			for _, p := range c.Prints {
				for _, o := range p.Occurrences {
					if p.Set.Code == tc.set && o.Type == "Precon" && o.Bundle == tc.bundle {
						list = append(list, fmt.Sprintf("%dx %s", o.Copies, strings.TrimSuffix(c.PrintedName+" ("+c.Suffix+")", " ()")))
					}
				}
			}
		}
		out, err := query(append([]string{"deck"}, list...)...)
		if err != nil {
			t.Fatal(err)
		}
		record := out[strings.Index(out, "## Track record"):strings.Index(out, "## Relatives")]
		if !strings.Contains(out, tc.want) ||
			tc.set == "KoT" && !strings.Contains(record, "rarely played cards: 43% of those not new. Beside the reference decks of the section above, each against the others, in %: 0, 0, 0, 4, 13, 27, 80 (min, 10%, 25%, median, 75%, 90%, max); 51 of the 1494 above it, 1442 below.") ||
			tc.set == "V5" && !strings.Contains(record, "rarely played cards: 0% of those not new. Beside the reference decks of the section above, each against the others, in %: 0, 0, 0, 1, 2, 7, 11 (min, 10%, 25%, median, 75%, 90%, max); 13 of the 26 above it, 0 below.") {
			t.Errorf("the %s %s precon's track record:\n%s", tc.set, tc.bundle, record)
		}
	}
	if strings.Contains(idea, "### Budgets") {
		t.Errorf("a deck idea set beside winning lists:\n%s", idea)
	}
	if _, err := query("deck", "Sword of Nuln", "Morbidity"); err == nil || !strings.Contains(err.Error(), "for a deck idea, draft reads a line without a count as one copy") {
		t.Errorf("an idea's cards without draft: %v", err)
	}
	if _, err := query("deck", "1x Sword of Nuln", "2x Villein"); err == nil || !strings.Contains(err.Error(), "For a deck idea rather than a whole list, call vtes_deck with draft") {
		t.Errorf("a short list without draft: %v", err)
	}
	// an idea without a vampire; a TWDA deck stays whole
	if out, err := query("-draft", "deck", "Sword of Nuln", "Morbidity"); err != nil ||
		!strings.Contains(out, "No crypt card given: the library's requirements are not checked against a crypt.") || strings.Contains(out, "Crypt cards with") {
		t.Errorf("an idea without a vampire: %v\n%s", err, out)
	}
	if out, err := query("-draft", "deck", "12947"); err != nil || !strings.Contains(out, "draft left aside: a tournament deck is read whole, as archived") || !strings.Contains(out, " per hand") {
		t.Errorf("a TWDA deck with draft: %v\n%s", err, out)
	}
	// a list past the bound: 20 vampires and 65 library cards, the longest texts
	b, err := bundle.Load(testBundle)
	if err != nil {
		t.Fatal(err)
	}
	cards := slices.Clone(b.Cards)
	slices.SortFunc(cards, func(x, y bundle.Card) int { return len(y.Text) - len(x.Text) })
	var long, vamps []string
	crypt := 0
	for _, c := range cards {
		vampire := slices.Contains(c.Types, "Vampire")
		if c.Banned != "" || vampire && (c.Group != "G5" || crypt == 20) || !vampire && len(long)-crypt == 65 {
			continue
		}
		if vampire {
			crypt++
			vamps = append(vamps, "1x "+c.Name)
		}
		long = append(long, "1x "+c.Name)
	}
	if out, err := query("deck", strings.Join(long, "\n")); err != nil || len(out) > 40000 || !strings.Contains(out, "Left out to fit this result: the texts of the library's cards.") || !strings.Contains(out, "This reference is loose") || !strings.Contains(out, "### Its nearest archetype's staples") {
		t.Errorf("a list past the bound: %v, %d characters", err, len(out))
	}
	// the same crypt and the 70 shortest library texts: only the relatives' details left out
	var short []string
	for _, c := range slices.Backward(cards) {
		if c.Banned == "" && !slices.Contains(c.Types, "Vampire") && !slices.Contains(c.Types, "Imbued") && len(short) < 70 {
			short = append(short, "1x "+c.Name)
		}
	}
	if out, err := query("deck", strings.Join(slices.Concat(vamps, short), "\n")); err != nil || len(out) > 40000 ||
		!strings.Contains(out, "the cards that differ from this deck are left out to fit this result.") || strings.Contains(out, "Left out to fit this result: the texts") {
		t.Errorf("a list fitting without the relatives' details: %v, %d characters", err, len(out))
	}
	if out, err := query("-detail", "deck", "12434"); err != nil || strings.Contains(out, "detail gives it whole") {
		t.Errorf("the detailed dossier cuts the comment: %v", err)
	}
	// a bundle built before it carried the archetype labels and the module tags
	b.Classification, b.Modules = bundle.Classification{}, nil
	old := filepath.Join(t.TempDir(), "old.json.gz")
	if err := b.Save(old); err != nil {
		t.Fatal(err)
	}
	var stdout, stderr strings.Builder
	if err := run(t.Context(), []string{"query", "-bundle", old, "deck", "12511"}, &stdout, &stderr); err != nil ||
		!strings.Contains(stdout.String(), "## Relatives\n\nThis data bundle holds no archetype labels: no relatives until it is rebuilt.") ||
		!strings.Contains(stdout.String(), "## Library by module\n\nThis data bundle holds no module tags: no modules until it is rebuilt.") {
		t.Errorf("a bundle without archetype labels and module tags: %v\n%s", err, stdout.String())
	}

	for _, tc := range []struct{ deck, err string }{
		{strings.Replace(herne, "Mr. Winthrop", "Mr Wintrop Supreme", 1), `No card is named "Mr Wintrop Supreme". Closest names: Mr. Winthrop;`},
		{strings.Replace(herne, "1x Ruslan Fedorenko", "1x Lucinde", 1), `"Lucinde" is several cards: Lucinde, Alastor (G3); Lucinde, Alastor (G7). Give the one meant`},
		{strings.NewReplacer("9x Earth Meld", "40x Earth Meld", "2x Club Illusion", "2x Return to Innocence", "1x Ruslan Fedorenko", "1x Anson").Replace(herne),
			"not tournament legal today"},
		{"nosuchdeck", `no deck of the TWDA has the id "nosuchdeck"`},
		{"https://vdb.im/decks/deck?name=x#100001", "not id=count pairs"},
		{"https://amaranth.vtes.co.nz/#deck/notpinned", `Amaranth has no public deck "notpinned" (Deck not found)`},
		{"https://vtesdecks.com/deck/notpinned", `VTES Decks has no public deck "notpinned" (404 Not Found)`},
		{"https://example.com/deck/12947", "primogen reads deck links from VDB (vdb.im), VTES Decks (vtesdecks.com), Amaranth (amaranth.vtes.co.nz) and Archon (archon.vekn.net), not example.com: give the list as text"},
		{strings.Join(vampires(b), "\n") + "\n60x Villein", "distinct cards does not fit a result of 40000 characters even without its texts"},
	} {
		_, err := query("deck", tc.deck)
		if err == nil || !strings.Contains(err.Error(), tc.err) {
			t.Errorf("error %v, want %q", err, tc.err)
		}
		if strings.Contains(tc.err, "legal") {
			for _, w := range []string{"the library has 121 cards: 60 to 90 (cut 31)", "The Return to Innocence is banned since 1999-04-01", "the crypt mixes groups 1, 6"} {
				if !strings.Contains(err.Error(), w) {
					t.Errorf("missing %q in %v", w, err)
				}
			}
		}
	}
	for _, w := range testWarnings {
		if strings.Contains(w, "decks of the TWDA name a card") || strings.Contains(w, "guide/module-tags.json:") || strings.Contains(w, "guide/modules.md:") {
			t.Error(w)
		}
	}
}

// vampires lists every vampire of groups 5 and 6 not banned, a legal crypt
// too large for any result.
func vampires(b *bundle.Bundle) []string {
	var out []string
	for _, c := range b.Cards {
		if slices.Contains(c.Types, "Vampire") && (c.Group == "G5" || c.Group == "G6") && c.Banned == "" {
			out = append(out, "1x "+c.Name)
		}
	}
	return out
}

// TestQueryStats surveys a slot: the owner's grade of the Toreador vote entry
// (calibration/corpus/2026-08-26-toreador-vote-push.md) named Awe's
// competitors Scalpel Tongue and Closed Session.
func TestQueryStats(t *testing.T) {
	query := func(args ...string) (string, error) {
		var stdout, stderr strings.Builder
		err := run(t.Context(), append([]string{"query", "-bundle", testBundle, "stats"}, args...), &stdout, &stderr)
		return stdout.String(), err
	}
	for _, args := range [][]string{{"Awe", "Endless Dance"}, {"Awe"}} {
		out, err := query(args...)
		if err != nil {
			t.Fatal(err)
		}
		i := strings.Index(out, "## Competitors for its slot")
		if i < 0 {
			t.Fatalf("%v: no competitors in:\n%s", args, out)
		}
		for _, w := range []string{"**Scalpel Tongue**", "**Closed Session**", "**Perfect Paragon**", "**Bewitching Oration**"} {
			if !strings.Contains(out[i:], w) {
				t.Errorf("%v: missing %s among the competitors:\n%s", args, w, out[i:])
			}
		}
		if strings.Contains(out[i:], "**Voter Captivation**") {
			t.Errorf("%v: Voter Captivation, played with Awe at another timing, is no competitor:\n%s", args, out[i:])
		}
		if !strings.Contains(out, "# Procedure: advice on a deck or an idea") {
			t.Errorf("%v: the advice procedure does not close the result", args)
		}
	}
	for _, tc := range []struct {
		args []string
		err  string
	}{
		{[]string{"Awe", "Toreador vote"}, `no archetype is named "Toreador vote". The archetypes playing Awe most: **Endless Dance** (C4);`},
		{[]string{"Awe", "Lutz vote"}, "Those whose name holds a word of it: **Lutz Politics** (C15)."},
		{[]string{"Awe", "Gangrel vote"}, "Those whose name holds a word of it: **Gangrel Resistance** (C11)"},
		{[]string{"Awe", "Nothing"}, "None has a name holding a word of it."},
		{[]string{"Villein", "Girls"}, `several archetypes are named "Girls": **Girls** (C25); **Girls** (G23), a variant of **Nananimalism**. Call`},
		{[]string{"Awwwwe of Nothing"}, `No card is named "Awwwwe of Nothing"`},
	} {
		if _, err := query(tc.args...); err == nil || !strings.Contains(err.Error(), tc.err) {
			t.Errorf("error %v, want %q", err, tc.err)
		}
	}
}

// TestQueryMeta reads the meta: the window's shares, the share in no
// archetype stated with them, and an archetype's core and decks.
func TestQueryMeta(t *testing.T) {
	query := func(args ...string) (string, error) {
		var stdout, stderr strings.Builder
		err := run(t.Context(), append([]string{"query", "-bundle", testBundle}, args...), &stdout, &stderr)
		return stdout.String(), err
	}
	out, err := query("meta")
	if err != nil {
		t.Fatal(err)
	}
	m := regexp.MustCompile(`Of the window's players, (\d+)% saw a deck the labels hold win, (\d+)% one placed by its nearest archetype, and \*\*(\d+)% one in no archetype\*\*`).FindStringSubmatch(out)
	if m == nil {
		t.Fatalf("no shares placed and unplaced in:\n%s", out)
	}
	if sum := atoi(t, m[1]) + atoi(t, m[2]) + atoi(t, m[3]); sum < 99 || sum > 101 {
		t.Errorf("the labelled, placed and unplaced shares make %d%%:\n%s", sum, m[0])
	}
	for _, w := range []string{"| 1 | **", "| | in no archetype |", "## What rises and falls", "## The cards behind the move", "# Procedure: reading the meta"} {
		if !strings.Contains(out, w) {
			t.Errorf("missing %q in:\n%s", w, out)
		}
	}
	out, err = query("meta", "Haqim Royalty")
	if err != nil {
		t.Fatal(err)
	}
	i, j := strings.Index(out, "## Its core"), strings.Index(out, "## Its decks")
	if i < 0 || j < i || !strings.Contains(out[i:j], "- Hunger of Marduk: ") || !strings.Contains(out[j:], "The representative decks the labels pick:\n- 12494, ") {
		t.Errorf("Haqim Royalty's core without Hunger of Marduk, or its decks without 12494:\n%s", out)
	}
	for _, w := range testWarnings {
		if strings.Contains(w, "map to no country") {
			t.Errorf("unmapped TWDA place in labels' window: %s", w)
		}
	}
	outFR, err := query("-country", "FR", "meta")
	if err != nil {
		t.Fatal(err)
	}
	outFrance, err := query("-country", "France", "meta")
	if err != nil {
		t.Fatal(err)
	}
	if outFR != outFrance {
		t.Errorf("country code FR and name France differ")
	}
	if !strings.Contains(outFR, "(filtered to France (FR), online events excluded)") || !strings.Contains(outFR, "| 1 | **") {
		t.Errorf("filtered country meta missing summary or shares:\n%s", outFR)
	}
	outEU, err := query("-continent", "Europe", "meta")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(outEU, "(filtered to Europe, online events excluded)") || !strings.Contains(outEU, "| 1 | **") {
		t.Errorf("filtered continent meta missing summary or shares:\n%s", outEU)
	}
	outEmpty, err := query("-country", "IS", "meta")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(outEmpty, "(filtered to Iceland (IS), online events excluded)") ||
		!strings.Contains(outEmpty, "No tournament-winning decks in the window match Iceland (IS).") ||
		strings.Contains(outEmpty, "## The core of the leading archetypes") {
		t.Errorf("empty-window country meta unexpected:\n%s", outEmpty)
	}
	outArchEU, err := query("-continent", "Europe", "meta", "Haqim Royalty")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(outArchEU, "(filtered to Europe, online events excluded)") || !strings.Contains(outArchEU, "## Share") {
		t.Errorf("filtered archetype meta unexpected:\n%s", outArchEU)
	}
	outArchEmpty, err := query("-country", "IS", "meta", "Haqim Royalty")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(outArchEmpty, "No tournament-winning decks in the window match Iceland (IS).") ||
		!strings.Contains(outArchEmpty, "No deck of it won in the window.") {
		t.Errorf("empty-window archetype meta unexpected:\n%s", outArchEmpty)
	}
	for _, tc := range []struct {
		args []string
		err  string
	}{
		{[]string{"meta", "Toreador"}, `no archetype is named "Toreador". Those whose name holds a word of it: **Weenie Toreador** (C70).`},
		{[]string{"meta", "Nothing"}, "None has a name holding a word of it. vtes_meta without an archetype names the leading ones"},
		{[]string{"-months", "99", "meta"}, "months must be 1 to "},
		{[]string{"-country", "Nowhere", "meta"}, `unknown country "Nowhere"`},
		{[]string{"-continent", "Atlantis", "meta"}, `unknown continent "Atlantis": use one of Europe, North America, South America, Asia, Oceania, Africa`},
		{[]string{"-country", "FR", "-continent", "Asia", "meta"}, `France (FR) is in Europe, not Asia`},
	} {
		if _, err := query(tc.args...); err == nil || !strings.Contains(err.Error(), tc.err) {
			t.Errorf("%v: error %v, want %q", tc.args, err, tc.err)
		}
	}
}

func atoi(t *testing.T, s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		t.Fatal(err)
	}
	return n
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
		// a rules director's post naming two cards no ruling names together, the archive fetched at the first call;
		// a reply naming them in the question it answers, which it quotes; four of a thread at most
		{args: []string{"case", "If you use Foreshadowing Destruction and your prey is over ten, can you add a Conditioning ?"}, want: []string{
			"# The rules directors' posts on these cards",
			"Post naming Foreshadowing Destruction and Conditioning: Vincent Ripoll (Ankha), the sitting rules director, 2019-10-24, \"About bleed modifiers\", https://usenet.krcg.org/t/vekn-78044/#m3\n",
			"Reply to a question it quotes naming Foreshadowing Destruction and Conditioning: Vincent Ripoll (Ankha), the sitting rules director, 2019-10-25, \"About bleed modifiers\", https://usenet.krcg.org/t/vekn-78044/#m12\n| > ",
			"- No ruling names Foreshadowing Destruction and Conditioning together. 6 posts by the rules director then sitting name them within a few sentences or answer a question that does, 5 of them above.\n",
		}, not: []string{"https://usenet.krcg.org/t/vekn-78044/#m13\n"}},
		// a card asked about alone, no ruling filed under it: the rules director's replies whose own words name nothing,
		// found by the question they answer, the post before it or the lines it quotes
		{args: []string{"case", "Political Flux is in play and a vote ousts two Methuselahs at once: who gains the pool?"}, want: []string{
			"Reply, after the post before it in its thread, by Damnans, https://usenet.krcg.org/t/z_kvG6T-DTQ/#m0, which it does not quote, naming Political Flux: " +
				"L. Scott Johnson, the sitting rules director, 2001-06-12, \"[LSJ] Political Flux\", https://usenet.krcg.org/t/z_kvG6T-DTQ/#m1\n" +
				"| > What happens in a 5 player game (A, B, C, D, E) when there is a\n| > Political Flux in play",
			"| Both, yes.\n",
			"Reply to a question it quotes naming Political Flux: L. Scott Johnson, the sitting rules director, 2010-01-16, \"(LSJ) Political Flux vs Simultaneous Ousting\", https://usenet.krcg.org/t/w6jOOcMIUbM/#m1\n",
			"| A. B would get 1 as well, had xe not been ousted.\n",
			"- No ruling of the rulings database is filed under Political Flux.\n",
			"- 2 posts by the rules director then sitting name Political Flux or answer a question naming it: all above.\n",
		}},
		{args: []string{"-directors", "archive", `"Political Flux"`}, want: []string{
			"0 posts hold [\"political flux\"] as words of their own, and 3 rules directors' replies hold them only with their thread's title or the question they answer",
			"Reply, after the post before it in its thread, by Damnans, https://usenet.krcg.org/t/z_kvG6T-DTQ/#m0, which it does not quote: " +
				"L. Scott Johnson, the sitting rules director, 2001-06-12, \"[LSJ] Political Flux\", https://usenet.krcg.org/t/z_kvG6T-DTQ/#m1\n",
			"| Both, yes.\n",
			"https://usenet.krcg.org/t/w6jOOcMIUbM/#m1\n",
		}, not: []string{"a player, not a rules director"}},
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
// renderer, that a section's id is its own, and that each [RBK …] citation of
// the advanced rules links an id of the rulebook's page.
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
	reRBK := regexp.MustCompile(`\[\[RBK\s[^\]]+\]\]\(([^)]*)\)`)
	cited := 0
	for _, s := range b.Sections {
		if s.Doc != bundle.DocAdvanced {
			continue
		}
		for _, text := range append([]string{s.Text}, s.Notes...) {
			for _, m := range reRBK.FindAllStringSubmatch(text, -1) {
				cited++
				id, ok := strings.CutPrefix(m[1], "https://archon.vekn.net/help/rules#")
				if !ok || !slices.Contains(ids["vtes-rules"], id) {
					t.Errorf("%s cites %s, not an id archon.vekn.net's rulebook has", s.ID, m[0])
				}
			}
		}
	}
	if cited < 150 {
		t.Errorf("the advanced rules hold %d rulebook citations, want about 160", cited)
	}
	for _, w := range testWarnings {
		if strings.HasPrefix(w, "archon's renderer") {
			t.Error(w)
		}
	}
}

// TestRulebookReferencesOpen checks every rulebook reference a ruling cites
// opens a section of the rulebook, its evidence.
func TestRulebookReferencesOpen(t *testing.T) {
	b, err := bundle.Load(testBundle)
	if err != nil {
		t.Fatal(err)
	}
	cited := 0
	for _, c := range b.Cards {
		for _, r := range c.Rulings {
			for _, ref := range r.References {
				if strings.HasPrefix(ref.Label, "RBK ") && ref.Section != "" {
					cited++
				}
			}
		}
	}
	if cited < 250 {
		t.Errorf("%d rulebook references cited with their section, want about 300, a group's ruling once per card", cited)
	}
	for _, w := range testWarnings {
		if strings.HasPrefix(w, "reference RBK ") {
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

// TestBuildProposesArchetypes undoes two of the owner's rulings: Agnieszka
// Wall unnamed is proposed again; Usurper High Council rejected, its decks all
// the owner's noise, is not.
func TestBuildProposesArchetypes(t *testing.T) {
	dir := t.TempDir()
	if err := os.CopyFS(dir, os.DirFS("knowledge")); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "classification.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var labels map[string]any
	json.Unmarshal(raw, &labels)
	var groups []any
	noise := labels["noise"].([]any)
	for _, g := range labels["groups"].([]any) {
		switch g.(map[string]any)["name"] {
		case "Agnieszka Wall":
		case "Usurper High Council":
			noise = append(noise, g.(map[string]any)["decks"].([]any)...)
		default:
			groups = append(groups, g)
		}
	}
	labels["groups"], labels["noise"] = groups, noise
	raw, _ = json.Marshal(labels)
	if err := os.WriteFile(path, raw, 0o644); err != nil {
		t.Fatal(err)
	}
	u, up, err := newUpstream()
	if err != nil {
		t.Fatal(err)
	}
	defer u.srv.Close()
	up.KnowledgeDir = dir
	archive, err := bundle.FetchArchive(context.Background(), up)
	if err != nil {
		t.Fatal(err)
	}
	_, warnings, err := bundle.Build(context.Background(), up, archive)
	if err != nil {
		t.Fatal(err)
	}
	var proposed []string
	for _, w := range warnings {
		if strings.HasPrefix(w, "proposed archetype") {
			proposed = append(proposed, w)
		}
	}
	if len(proposed) != 1 || !strings.Contains(proposed[0], `proposed archetype 1 of 1, TWDA decks since 2021-07-01 in no archetype of knowledge/classification.json `+
		`(name it there, or put its decks in its noise): 4 decks, 13460 "Agnieszka, The Unbreakable [Joab Version]" (2026-08-16), `) ||
		!strings.Contains(proposed[0], `13163 "Agnieszka" (2026-03-17); their vampires: Agnieszka, Tempter of Legions / Verrix, Naughty Boy / `) ||
		!strings.Contains(proposed[0], "nearest archetype Goratrix High Tower at 0.15") {
		t.Errorf("want Agnieszka Wall proposed alone, got:\n%s", strings.Join(proposed, "\n"))
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
	up.KnowledgeList = u.srv.URL + "/primogen-list/%s?ref=%s"
	r, err := newRefresher(ctx, up, filepath.Join(t.TempDir(), "bundle.json.gz"), slog.New(slog.DiscardHandler))
	if err != nil {
		t.Fatal(err)
	}
	r.links = u.links()
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
	if res, err := cs.CallTool(ctx, &mcp.CallToolParams{Name: "vtes_guide", Arguments: map[string]any{"topic": "examples"}}); err != nil || res.IsError {
		t.Fatalf("vtes_guide examples on hosted server: %v %v", err, res)
	} else {
		expect("hosted vtes_guide examples", res.Content[0].(*mcp.TextContent).Text, "card-comparison-trap-conquer-the-beast", "meta-reading-top-tier-decks")
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
	expect("status page", get("/status"), "Data: cards of static.krcg.org 2026-11-01 00:00 UTC", "Last refresh",
		"TWDA decks since 2021-07-01, not in the archetype labels (knowledge/classification.json), match no archetype nor proposal", `13020 "Louvre Heist!" (2026-06-17, closest Weenie Toreador at 0.28)`,
		"library cards the TWDA decks since 2021-07-01 play have no module tag (knowledge/guide/module-tags.json); the 50 most played, with their decks: ")
	for _, path := range []string{"/", "/llms.txt"} {
		expect("catalog", get(path), "# KRCG\n", "(https://mcp.krcg.org/mcp)", "(https://static.krcg.org/data/vtes.json)")
	}
	r.refresh(ctx, r.fetched)
	if status := get("/status"); strings.Contains(status, "does not answer") {
		t.Errorf("a link of the catalog does not answer:\n%s", status)
	}
	cs = connect()
	res, err := cs.CallTool(ctx, &mcp.CallToolParams{Name: "vtes_rulings_archive", Arguments: map[string]any{"query": "Freak Drive", "directors": true}})
	cs.Close()
	if err != nil || res.IsError {
		t.Fatal(err, res)
	}
	expect("the hosted archive", res.Content[0].(*mcp.TextContent).Text, "the sitting rules director")
	cs = connect()
	res, err = cs.CallTool(ctx, &mcp.CallToolParams{Name: "vtes_deck", Arguments: map[string]any{"deck": "12947"}})
	cs.Close()
	if err != nil || res.IsError {
		t.Fatal(err, res)
	}
	expect("a hosted deck", res.Content[0].(*mcp.TextContent).Text, "## Crypt (12 cards, 9 distinct", "# Procedure: a deck analysis")
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
	expect("status page", get("/status"), "ok.")
	if u.archives != 1 {
		t.Errorf("%d fetches of the archive before the night, want 1", u.archives)
	}
	age()
	out, _ = local(published)
	expect("local run a day later", out, "cards of static.krcg.org 2026-11-02 12:34 UTC")

	// a link of the catalog that died
	u.kill("https://api.krcg.org/openapi.json")
	r.refresh(ctx, base)
	expect("status page", get("/status"), "ok.", "does not answer: https://api.krcg.org/openapi.json: 404 Not Found")
	u.kill("")

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
	expect("status page", get("/status"), "FAILED, the bundle above keeps serving", "cards, the live one")

	// offline, a local install serves what it has and says how old it is
	age()
	out, errs := local("http://127.0.0.1:1/bundle.json.gz")
	expect("offline local run", out, "## Freak Drive")
	expect("offline local run", errs, "could not check for a newer bundle.json.gz, serving the one of ")
}

func TestAudit(t *testing.T) {
	runAudit := func(args ...string) (string, string, error) {
		var out, errBuf bytes.Buffer
		err := run(t.Context(), append([]string{"audit", "-bundle", testBundle, "-archive", "testdata/upstream/archive.json.gz"}, args...), &out, &errBuf)
		return out.String(), errBuf.String(), err
	}

	report, _, err := runAudit("Villein", "Conditioning", "Foreshadowing Destruction", "Political Flux")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"# Uncited rules directors' posts",
		"## By card (most played in the TWDA window first)",
		"### Villein (",
		"### Conditioning (",
		"### Foreshadowing Destruction (",
		"### Political Flux (1 TWDA decks, 2 uncited posts)",
		"Existing rulings on Political Flux: none.",
		"Link: https://usenet.krcg.org/t/z_kvG6T-DTQ/#m1",
		"Link: https://usenet.krcg.org/t/w6jOOcMIUbM/#m1",
		"| Both, yes.",
		"Card changes after 2015-05-24 (Villein):",
		"2018-05-11: The amount of blood you can move is now limited to 5.",
		"## By card pair (most played in the TWDA window first)",
		"### Conditioning + Foreshadowing Destruction (",
		"Link: https://usenet.krcg.org/t/vekn-78044/#m3",
		"Link: https://usenet.krcg.org/t/vekn-78044/#m13",
	} {
		if !strings.Contains(report, want) {
			t.Errorf("audit report misses %q", want)
		}
	}
	// Ordered by TWDA window play count: Villein > Conditioning > Foreshadowing Destruction > Political Flux.
	for _, pair := range [][2]string{
		{"### Villein (", "### Conditioning ("},
		{"### Conditioning (", "### Foreshadowing Destruction ("},
		{"### Foreshadowing Destruction (", "### Political Flux ("},
	} {
		if strings.Index(report, pair[0]) >= strings.Index(report, pair[1]) {
			t.Errorf("expected %q before %q in audit report", pair[0], pair[1])
		}
	}
	// A post already cited by a ruling on Foreshadowing Destruction (PIB 20110817, vekn-8405/#8414) is excluded.
	if strings.Contains(report, "https://usenet.krcg.org/t/vekn-8405/#m5") {
		t.Error("audit report included a post already cited by a ruling on Foreshadowing Destruction")
	}

	// -twda drops cards not played in the TWDA window (Extremis Boon has 0 TWDA decks).
	twdaOnly, _, err := runAudit("-twda", "Villein", "Extremis Boon")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(twdaOnly, "### Villein (") || strings.Contains(twdaOnly, "### Extremis Boon (") {
		t.Errorf("-twda report unexpected: %s", twdaOnly[:min(len(twdaOnly), 400)])
	}

	// -since filters older posts; -out writes to disk.
	outFile := filepath.Join(t.TempDir(), "audit.md")
	stdout, stderr, err := runAudit("-since", "2019-01-01", "-out", outFile, "Villein", "Political Flux")
	if err != nil {
		t.Fatal(err)
	}
	if stdout != "" || !strings.Contains(stderr, outFile) {
		t.Errorf("with -out, got stdout=%q stderr=%q", stdout, stderr)
	}
	raw, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatal(err)
	}
	written := string(raw)
	if !strings.Contains(written, "### Villein (") || strings.Contains(written, "### Political Flux (") || strings.Contains(written, "2015-05-24") {
		t.Errorf("filtered -out report unexpected: %s", written[:min(len(written), 400)])
	}

	if _, _, err := runAudit("-since", "not-a-date"); err == nil {
		t.Error("expected error for invalid -since")
	}

	// Non-rules forum categories (vekn-79165, deck-clinic) and tourney-report titles (M9f7z02UbzI) are skipped.
	noiseReport, _, err := runAudit("Villein", "Steam Tunnels")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(noiseReport, "vekn-74643/#m4") || strings.Contains(noiseReport, "vekn-79165") || strings.Contains(noiseReport, "M9f7z02UbzI") {
		t.Errorf("noise thread filter unexpected in audit report")
	}
}

func TestQueryGuide(t *testing.T) {
	runGuide := func(args ...string) (string, error) {
		var stdout, stderr strings.Builder
		err := run(t.Context(), append([]string{"query", "-bundle", testBundle}, args...), &stdout, &stderr)
		return stdout.String(), err
	}

	for _, tc := range []struct {
		name      string
		args      []string
		want, not []string
		err       string
	}{
		{
			name: "procedure card-comparison",
			args: []string{"guide", "card-comparison"},
			want: []string{
				"Authority: curated guide",
				"# Procedure: comparing or evaluating cards",
				"Minion Tap beats Villein",
				"never from memory",
				"Data: ",
			},
		},
		{
			name: "procedure with space alias",
			args: []string{"guide", "rules answer"},
			want: []string{"Authority: curated guide", "Data: "},
		},
		{
			name: "theory section",
			args: []string{"guide", "theory", "Combat"},
			want: []string{"## Combat", "Authority: curated guide"},
			not:  []string{"## Economy"},
		},
		{
			name: "modules whole fits bound",
			args: []string{"guide", "modules"},
			want: []string{"# Modules catalog", "## Combat defenses", "not yet canonized by the owner"},
			not:  []string{"(cut here"},
		},
		{
			name: "modules section",
			args: []string{"guide", "modules", "Combat defenses"},
			want: []string{"## Combat defenses", "not yet canonized by the owner"},
			not:  []string{"## Bloat & pool economy"},
		},
		{
			name: "terms default EN",
			args: []string{"guide", "terms"},
			want: []string{
				"# VTES game terms (EN)",
				"not an official authority",
				"Available languages: EN, ES, FR, IT, JP, LATIN, PT-BR",
				"- Bleed",
			},
		},
		{
			name: "terms FR with filter",
			args: []string{"-lang", "fr", "-filter", "Bleed", "guide", "terms"},
			want: []string{
				"# VTES game terms (FR)",
				"Translations have not been checked against the translated rulebooks.",
				"Recorded source for FR:",
				"- Bleed: Morsure",
			},
			not: []string{"- Hunt:"},
		},
		{
			name: "terms FR untranslated term",
			args: []string{"-lang", "FR", "-filter", "Stun", "guide", "terms"},
			want: []string{"- Stun: (no translation recorded)"},
		},
		{
			name: "terms category disciplines shows partial note",
			args: []string{"-lang", "FR", "guide", "terms", "disciplines"},
			want: []string{"## disciplines", "lists only some disciplines", "- Animalism: Animalisme"},
			not:  []string{"## actions"},
		},
		{
			name: "examples listing",
			args: []string{"guide", "examples"},
			want: []string{
				"# Served examples",
				"`card-comparison-trap-conquer-the-beast`",
				"`meta-reading-top-tier-decks`",
				"Call vtes_guide with topic",
			},
			not: []string{"**1. Equipment.**"},
		},
		{
			name: "examples filtered listing",
			args: []string{"-filter", "top tier", "guide", "examples"},
			want: []string{"`meta-reading-top-tier-decks`"},
			not:  []string{"`card-comparison-trap-conquer-the-beast`"},
		},
		{
			name: "example read by name",
			args: []string{"guide", "examples", "card-comparison-trap-conquer-the-beast"},
			want: []string{
				"Authority: worked example",
				"**1. Equipment.**",
				"Who owns the press",
				"Data: ",
			},
		},
		{
			name: "unknown topic lists valid options",
			args: []string{"guide", "unknown"},
			err:  "rules-answer, deck-analysis, deck-advice, card-comparison, meta-reading, theory, modules, terms, examples",
		},
		{
			name: "empty topic lists valid options",
			args: []string{"guide"},
			err:  "rules-answer, deck-analysis, deck-advice, card-comparison, meta-reading, theory, modules, terms, examples",
		},
		{
			name: "unknown language lists valid languages",
			args: []string{"-lang", "DE", "guide", "terms"},
			err:  "EN, ES, FR, IT, JP, LATIN, PT-BR",
		},
		{
			name: "unknown terms category lists categories",
			args: []string{"guide", "terms", "bogus"},
			err:  "actions",
		},
		{
			name: "unknown theory section lists sections",
			args: []string{"guide", "theory", "bogus"},
			err:  "Economy",
		},
		{
			name: "unknown example lists available examples",
			args: []string{"guide", "examples", "bogus"},
			err:  "card-comparison-trap-conquer-the-beast",
		},
		{
			name: "ambiguous example lists matching examples",
			args: []string{"guide", "examples", "Example"},
			err:  "card-comparison-trap-conquer-the-beast, meta-reading-top-tier-decks",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out, err := runGuide(tc.args...)
			if tc.err != "" {
				if err == nil || !strings.Contains(err.Error(), tc.err) {
					t.Fatalf("got err=%v, want substring %q", err, tc.err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(out) > 40000 {
				t.Errorf("result length %d exceeds 40,000-character bound", len(out))
			}
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
