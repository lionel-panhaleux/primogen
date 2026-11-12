package main

import (
	"cmp"
	"encoding/csv"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.yaml.in/yaml/v3"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
	"github.com/lionel-panhaleux/primogen/internal/tools"
)

// The evidence tests: for each graded case, the pack vtes_rules_case returns
// for its question holds what settles it.

type evidenceCase struct {
	name     string
	question string
	cards    []string
	want     []evidence
}

// evidence is a section (by its document and title), a reference label, or a
// card's text, which a pack holds when it holds wantText.
type evidence struct {
	doc, title string // a section
	text       string // anything else, as the pack writes it
}

// knownMisses are the evidence keyword search does not reach, each with why:
// a miss not listed fails, and so does a listed one found again.
var knownMisses = map[string]string{
	"2026-08-26-shroud-of-decay-stalling.md ### [judges:3.5] ":                     "a lexical gap: \"failing to announce their card effects\" against \"incomplete information about publicly available game elements\"",
	"rules-qa 13.2 ### [advanced:3.2.3] ":                                          "a three-point item, cited for its third point only",
	"rules-qa 58.13 ### [advanced:3.5.2] ":                                         "a six-point item on Mask of a Thousand Faces, cited for one point: the masker takes the action's NRA marks",
	"2026-11-12-govern-superior-without-dom.md ### [judges:2.3] ":                  "a lexical gap: a judge's \"couldn't have played\" against \"Game Rule Violation\"",
	"2026-11-12-govern-superior-without-dom.md ### [judges:1.4.1] ":                "a lexical gap: \"what do I do\" against \"Rolling Back the Game State\"",
	"2026-11-12-foreshadowing-conditioning-limit.md usenet.krcg.org/t/vekn-78044/": "settled by the rules director's posts in the archive, which no ruling cites",
}

func checkEvidence(t *testing.T, b *bundle.Bundle, s *mcp.Server, c evidenceCase) []string {
	t.Helper()
	pack, err := tools.Call(t.Context(), s, "vtes_rules_case", map[string]any{"question": c.question, "cards": append([]string{}, c.cards...)})
	if err != nil {
		t.Fatalf("%s: %v", c.name, err)
	}
	var missing []string
	for _, w := range c.want {
		text := w.text
		if w.doc != "" {
			i := slices.IndexFunc(b.Sections, func(s bundle.Section) bool { return s.Doc == w.doc && sectionTitle(s) == w.title })
			if i < 0 {
				missing = append(missing, fmt.Sprintf("no section of %s is titled %q today: back to the owner", w.doc, w.title))
				continue
			}
			text = "### [" + b.Sections[i].ID + "] "
		}
		why, known := knownMisses[c.name+" "+text]
		switch found := strings.Contains(pack, text); {
		case !found && known:
			t.Logf("%s: known miss %q: %s", c.name, text, why)
		case !found:
			missing = append(missing, text)
		case known:
			t.Errorf("%s: the known miss %q is found now: remove it from knownMisses", c.name, text)
		}
	}
	return missing
}

var reSectionNumber = regexp.MustCompile(`^\d+(?:\.\d+)*[a-z]?\.?\s+`)

func sectionTitle(s bundle.Section) string { return reSectionNumber.ReplaceAllString(s.Title, "") }

func loadTestBundle(t *testing.T) (*bundle.Bundle, *mcp.Server) {
	t.Helper()
	b, err := bundle.Load(testBundle)
	if err != nil {
		t.Fatal(err)
	}
	return b, tools.NewServer(tools.NewLive(b))
}

// TestEvidenceCorpus reads the owner-graded rules entries, each with what
// settles it.
func TestEvidenceCorpus(t *testing.T) {
	b, s := loadTestBundle(t)
	docs := map[string]string{"rulebook": bundle.DocRulebook, "advanced rules": bundle.DocAdvanced, "tournament rules": bundle.DocTournament, "judges' guide": bundle.DocJudges}
	files, _ := filepath.Glob("calibration/corpus/*.md")
	seeded := 0
	for _, f := range files {
		front, body := readEntry(t, f)
		if front["mode"] != "rules" || front["status"] != "graded" {
			continue
		}
		if front["settled-by"] == "" {
			t.Logf("%s: no settled-by, not seeded", filepath.Base(f))
			continue
		}
		c := evidenceCase{name: filepath.Base(f), question: body["Question"]}
		for _, item := range strings.Split(front["settled-by"], "; ") {
			item = strings.TrimSpace(regexp.MustCompile(`\s*\(.*\)$`).ReplaceAllString(item, ""))
			if names, ok := strings.CutPrefix(item, "card text of "); ok {
				for _, n := range strings.Split(names, ", ") {
					c.cards = append(c.cards, n)
					c.want = append(c.want, evidence{text: "## " + n + "\n"})
				}
				continue
			}
			if item == "owner's grade" {
				continue
			}
			if doc, title, ok := strings.Cut(item, ", "); ok && docs[doc] != "" {
				c.want = append(c.want, evidence{doc: docs[doc], title: title})
				continue
			}
			c.want = append(c.want, evidence{text: item})
		}
		seeded++
		if missing := checkEvidence(t, b, s, c); len(missing) > 0 {
			t.Errorf("%s: the pack misses %q", c.name, missing)
		}
	}
	if seeded < 7 {
		t.Errorf("%d rules entries seeded, want the 7 graded with what settles them", seeded)
	}
}

// readEntry answers a corpus entry's front matter and its sections by heading.
func readEntry(t *testing.T, path string) (map[string]string, map[string]string) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	parts := strings.SplitN(string(raw), "---\n", 3)
	front := map[string]string{}
	for line := range strings.Lines(parts[1]) {
		if k, v, ok := strings.Cut(strings.TrimSpace(line), ": "); ok {
			front[k] = v
		}
	}
	body := map[string]string{}
	heading := ""
	for line := range strings.Lines(parts[2]) {
		if h, ok := strings.CutPrefix(line, "## "); ok {
			heading = strings.TrimSpace(h)
			continue
		}
		body[heading] += line
	}
	return front, body
}

// TestEvidenceCodex holds the rules facts the codex strategist's graded cases
// C and F turn on (calibration/codex-cases.md).
func TestEvidenceCodex(t *testing.T) {
	b, s := loadTestBundle(t)
	for _, c := range []evidenceCase{{
		name:     "C: Entombment then Graverobbing, against combat ends and Immortal Grapple",
		question: "My vampire strikes with superior Entombment to send the opposing vampire to torpor, then steals it with Graverobbing. Does it work if the opponent strikes with combat ends, or plays Immortal Grapple?",
		cards:    []string{"Entombment", "Graverobbing", "Immortal Grapple"},
		want: []evidence{
			{doc: bundle.DocRulebook, title: "Combat Ends"},
			{text: "Strikes that are not hand strikes cannot be used this round"},
			{text: "[OBT] Strike: send the opposing vampire to torpor."},
		},
	}, {
		name:     "F: what fires CrimethInc., and who recruits Tunnel Runner",
		question: "Does an action modifier requiring an Anarch, or the Become Anarch action, let me play CrimethInc.? And can any vampire recruit Tunnel Runner?",
		cards:    []string{"CrimethInc.", "Tunnel Runner"},
		want: []evidence{
			{doc: bundle.DocRulebook, title: "ADVANCED RULES — Become Anarch"},
			{text: "Only usable after resolution of a successful action requiring an Anarch or making this vampire Anarch."},
			{text: "Clans its requirement names: Akunanse."},
		},
	}} {
		if missing := checkEvidence(t, b, s, c); len(missing) > 0 {
			t.Errorf("%s: the pack misses %q", c.name, missing)
		}
	}
}

var (
	reQAItem = regexp.MustCompile(`(?m)^\*\*(\d+\.\d+)\*\*\s+`)
	reQACite = regexp.MustCompile(`rules-qa\s+(\d+\.\d+)`)
)

// TestEvidenceRulesQA asks each item of zero-vtes's rules-qa the advanced
// rules cite as the question, cards found in its words: the pack holds a
// section citing it.
func TestEvidenceRulesQA(t *testing.T) {
	b, s := loadTestBundle(t)
	raw, err := os.ReadFile("testdata/upstream/zero-vtes/docs/rules-qa.md")
	if err != nil {
		t.Fatal(err)
	}
	items := map[string]string{}
	locs := reQAItem.FindAllStringSubmatchIndex(string(raw), -1)
	for i, l := range locs {
		end := len(raw)
		if i+1 < len(locs) {
			end = locs[i+1][0]
		}
		text, _, _ := strings.Cut(string(raw[l[1]:end]), "\n## ")
		items[string(raw[l[2]:l[3]])] = strings.TrimSpace(text)
	}
	citing := map[string][]bundle.Section{}
	for _, s := range b.Sections {
		for _, n := range s.Notes {
			for _, m := range reQACite.FindAllStringSubmatch(n, -1) {
				citing[m[1]] = append(citing[m[1]], s)
			}
		}
	}
	for _, id := range slices.Sorted(func(yield func(string) bool) {
		for k := range citing {
			if !yield(k) {
				return
			}
		}
	}) {
		q, ok := items[id]
		if !ok {
			t.Errorf("the advanced rules cite rules-qa %s, which the pinned rules-qa does not hold", id)
			continue
		}
		var want []evidence
		for _, s := range citing[id] {
			want = append(want, evidence{text: "### [" + s.ID + "] "})
		}
		if missing := checkEvidence(t, b, s, evidenceCase{name: "rules-qa " + id, question: q, want: want}); len(missing) > 0 && len(missing) == len(want) {
			t.Errorf("rules-qa %s: the pack holds none of the sections citing it, %q", id, missing)
		}
	}
}

// TestEvidenceLedger checks each row of zero-vtes's rulings ledger, a ruling
// filed under a card, under its group, or under another card naming it: the
// first two among the card's rulings, the third among the rulings naming it.
func TestEvidenceLedger(t *testing.T) {
	_, s := loadTestBundle(t)
	f, err := os.Open("testdata/upstream/zero-vtes/data/rulings-v5.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	type row struct{ id, via, ruling string }
	byName := map[string][]row{}
	for _, r := range rows[1:] {
		byName[r[1]] = append(byName[r[1]], row{r[0], r[2], r[4]})
	}
	reNoise := regexp.MustCompile(`\{\d+\|([^}]*)\}|\[[A-Z]{3} [\w-]+\]|\[REMINDER\]|\\n|\\u00a0|\\`)
	norm := func(s string) string {
		s = reNoise.ReplaceAllStringFunc(s, func(m string) string {
			if sub := reNoise.FindStringSubmatch(m); sub[1] != "" {
				return sub[1]
			}
			return " "
		})
		return strings.Join(strings.Fields(strings.ReplaceAll(s, "\u00a0", " ")), " ")
	}
	// what the pinned database files under each card id, read apart from the build
	commit, _ := os.ReadFile("testdata/upstream/commit")
	dir := filepath.Join("testdata/upstream/vtes-rulings", strings.TrimSpace(string(commit)))
	var rulings map[string][]yaml.Node
	var groups map[string]map[string]string
	for name, v := range map[string]any{"rulings.yaml": &rulings, "groups.yaml": &groups} {
		raw, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			t.Fatal(err)
		}
		if err := yaml.Unmarshal(raw, v); err != nil {
			t.Fatal(err)
		}
	}
	filed := map[string]map[string]bool{} // card id → its rulings, normalized
	file := func(id, text string) {
		if filed[id] == nil {
			filed[id] = map[string]bool{}
		}
		filed[id][norm(text)] = true
	}
	for key, entries := range rulings {
		target, _, _ := strings.Cut(key, "|")
		for _, e := range entries {
			var m struct {
				Text      string
				Overrides map[string]string
			}
			if e.Kind == yaml.ScalarNode {
				m.Text = e.Value
			} else if err := e.Decode(&m); err != nil {
				t.Fatal(err)
			}
			if !strings.HasPrefix(target, "G") {
				file(target, m.Text)
				continue
			}
			for member := range groups[key] {
				id, _, _ := strings.Cut(member, "|")
				file(id, cmp.Or(m.Overrides[member], m.Text))
			}
		}
	}
	anywhere := map[string]bool{}
	for _, texts := range filed {
		for text := range texts {
			anywhere[text] = true
		}
	}
	missed, drifted, checked := 0, 0, 0
	// the rulings on other cards naming one are only in the case's detail
	mentioned := func(n string) bool {
		return slices.ContainsFunc(byName[n], func(r row) bool { return strings.HasPrefix(r.via, "mention in ") })
	}
	var calls []map[string]any
	var plain []string
	for _, n := range slices.Sorted(maps.Keys(byName)) {
		if mentioned(n) {
			calls = append(calls, map[string]any{"question": "rulings", "cards": []string{n}, "detail": true})
		} else {
			plain = append(plain, n)
		}
	}
	for i := 0; i < len(plain); i += 10 {
		calls = append(calls, map[string]any{"names": plain[i:min(i+10, len(plain))]})
	}
	for _, in := range calls {
		tool, names := "vtes_cards", in["names"]
		if names == nil {
			tool, names = "vtes_rules_case", in["cards"]
		}
		out, err := tools.Call(t.Context(), s, tool, in)
		if err != nil {
			t.Fatal(err)
		}
		out = norm(out)
		for _, n := range names.([]string) {
			for _, r := range byName[n] {
				ruling := norm(r.ruling)
				if strings.HasPrefix(r.via, "mention in ") && !anywhere[ruling] || !strings.HasPrefix(r.via, "mention in ") && !filed[r.id][ruling] {
					drifted++
					continue
				}
				checked++
				if !strings.Contains(out, ruling) {
					missed++
					t.Errorf("%s (%s): the pack misses %q", n, r.via, r.ruling)
				}
			}
		}
	}
	t.Logf("%d ledger rows checked, %d missed; %d the pinned rulings database words or files otherwise", checked, missed, drifted)
}
