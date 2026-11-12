package tools

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

// authority labels each rules document.
var authority = map[string]string{
	bundle.DocRulebook:   "official",
	bundle.DocAdvanced:   "curated, not verified by a rules director",
	bundle.DocTournament: "official",
	bundle.DocJudges:     "official",
}

// rulesIndex finds sections by keyword, by id and by the cards they name,
// and rulings by the cards they name.
type rulesIndex struct {
	sections []bundle.Section
	byID     map[string]int
	search   *bm25
	// naming maps a card ID to the sections and the rulings that name it.
	sectionsNaming map[int][]int
	rulingsNaming  map[int][]namedRuling
	// sectionsCiting maps a reference label (LSJ 19980212) to the sections whose footnotes cite it.
	sectionsCiting map[string][]int
	// rulebook is the rulebook's text folded and terms the game's own words
	// (disciplines, clans, titles, players' shorthand): neither is
	// taken for a card's name.
	rulebook string
	terms    map[string]bool
}

type namedRuling struct {
	card *bundle.Card // the card it is filed under
	r    *bundle.Ruling
}

var (
	reBraced = regexp.MustCompile(`\{([^{}|]+)\}`)
	reCited  = regexp.MustCompile(`\[\[([A-Z]{3} [\w-]+)\]\]`)
)

func newRulesIndex(b *bundle.Bundle, cards *cardIndex) *rulesIndex {
	ix := &rulesIndex{
		sections:       b.Sections,
		byID:           map[string]int{},
		sectionsNaming: map[int][]int{},
		rulingsNaming:  map[int][]namedRuling{},
		sectionsCiting: map[string][]int{},
		terms:          map[string]bool{},
	}
	for _, t := range slices.Concat(b.GameTerms, b.Shorthand) {
		ix.terms[bundle.Fold(t)] = true
	}
	for _, c := range b.Cards {
		for _, clan := range c.Clans {
			ix.terms[bundle.Fold(clan)] = true
		}
		if c.Title != "" {
			ix.terms[bundle.Fold(c.Title)] = true
		}
		if strings.HasPrefix(c.Text, "Master: Discipline") {
			ix.terms[bundle.Fold(c.Name)] = true
		}
	}
	texts := make([]string, len(b.Sections))
	titles := make([]string, len(b.Sections))
	var rulebook strings.Builder
	for i, s := range b.Sections {
		ix.byID[s.ID] = i
		for _, n := range s.Notes {
			for _, m := range reCited.FindAllStringSubmatch(n, -1) {
				if l := ix.sectionsCiting[m[1]]; !slices.Contains(l, i) {
					ix.sectionsCiting[m[1]] = append(l, i)
				}
			}
		}
		texts[i] = s.Text
		titles[i] = s.Title
		if s.Doc == bundle.DocRulebook {
			rulebook.WriteString(" " + s.Title + " " + s.Text)
		}
		for _, m := range reBraced.FindAllStringSubmatch(s.Text+" "+strings.Join(s.Notes, " "), -1) {
			for _, c := range cards.english[bundle.Fold(m[1])] {
				if l := ix.sectionsNaming[c.ID]; len(l) == 0 || l[len(l)-1] != i {
					ix.sectionsNaming[c.ID] = append(l, i)
				}
			}
		}
	}
	ix.rulebook = proseForm(rulebook.String())
	ix.search = newBM25(texts, titles)
	for i := range b.Cards {
		c := &b.Cards[i]
		for j := range c.Rulings {
			r := &c.Rulings[j]
			for _, m := range reCard.FindAllStringSubmatch(r.Text, -1) {
				id, _ := strconv.Atoi(m[1])
				if id != c.ID {
					ix.rulingsNaming[id] = append(ix.rulingsNaming[id], namedRuling{c, r})
				}
			}
		}
	}
	return ix
}

var reCard = regexp.MustCompile(`\{(\d+)\|[^}]*\}`)

// subsections answers a section and every section under it.
func (ix *rulesIndex) subsections(i int) []bundle.Section {
	s := ix.sections[i]
	end := i + 1
	for end < len(ix.sections) && ix.sections[end].Doc == s.Doc && ix.sections[end].Level > s.Level {
		end++
	}
	return ix.sections[i:end]
}

func writeSectionHead(out *strings.Builder, s bundle.Section) {
	fmt.Fprintf(out, "### [%s] %s\n", s.ID, s.Title)
	where := bundle.DocTitles[s.Doc]
	if len(s.Path) > 0 {
		where += " › " + strings.Join(s.Path, " › ")
	}
	fmt.Fprintf(out, "%s (%s) · %s\n", where, authority[s.Doc], s.URL)
}

// writeSection writes a section, cut to the paragraphs nearest the query
// when longer than limit (0: whole, its footnotes with their card lists).
func writeSection(out *strings.Builder, s bundle.Section, q map[string]float64, limit int) {
	writeSectionHead(out, s)
	text := s.Text
	cut := false
	if limit > 0 {
		text, cut = excerpt(text, q, limit)
	}
	if text != "" {
		fmt.Fprintf(out, "%s\n", text)
	}
	if cut {
		fmt.Fprintf(out, "(excerpt: vtes_rules_read %q gives it whole)\n", s.ID)
	}
	var notes []string
	for _, n := range s.Notes {
		call, _, _ := strings.Cut(n, "]:")
		if !strings.Contains(text, call+"]") {
			continue
		}
		// excerpted, a footnote keeps its citations, not the cards it lists
		if cited, _, found := strings.Cut(n, " — "); limit > 0 && found {
			n = cited
		}
		notes = append(notes, n)
	}
	if len(notes) > 0 {
		fmt.Fprintf(out, "Its sources (footnotes):\n%s\n", strings.Join(notes, "\n"))
	}
	out.WriteString("\n")
}

var docsSchema = &jsonschema.Schema{
	Type: "array", Items: &jsonschema.Schema{Type: "string", Enum: []any{bundle.DocRulebook, bundle.DocAdvanced, bundle.DocTournament, bundle.DocJudges}},
	Description: "limit to these documents: rulebook, advanced (the advanced rules), tournament (the tournament rules), judges (the judges' guide); default all",
}

type searchInput struct {
	Query  string   `json:"query"`
	Docs   []string `json:"docs,omitempty"`
	Detail bool     `json:"detail,omitempty"`
}

type readInput struct {
	IDs []string `json:"ids"`
}

const groundRules = "Answer from these sections, quoting them and citing their links, never from memory.\n\n"

func addRules(s *mcp.Server, l *Live) {
	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_rules_search",
		Title: "VTES rules search",
		Description: "Never state a rule from memory: quote the sections returned. " +
			"Keyword search over the Vampire: The Eternal Struggle (VTES) rules texts: the rulebook, the advanced rules, " +
			"the tournament rules and the judges' guide, by section. " +
			"For a question about cards, vtes_rules_case returns the cards, their rulings and the governing sections in one call.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"query":  {Type: "string", Description: "the words to look for, e.g. \"block after declining\""},
				"docs":   docsSchema,
				"detail": {Type: "boolean", Description: "more results, longer excerpts"},
			},
			Required: []string{"query"},
		},
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: new(false)},
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in searchInput) (*mcp.CallToolResult, any, error) {
		q := query(in.Query, 1, nil)
		if len(q) == 0 {
			return nil, nil, fmt.Errorf(`give words to search, e.g. {"query": "block after declining"}; got %q`, in.Query)
		}
		d := l.p.Load()
		ix := d.rules
		results := ix.search.rank(q, func(i int) bool { return len(in.Docs) == 0 || slices.Contains(in.Docs, ix.sections[i].Doc) })
		n, limit := 8, 600
		if in.Detail {
			n, limit = 20, 1500
		}
		var out strings.Builder
		if len(results) == 0 {
			fmt.Fprintf(&out, "No section of the rules texts holds these words: %q. Try other words, or fewer.\n\n", in.Query)
		} else {
			fmt.Fprintf(&out, "%d sections match %q; the best %d, each excerpted: vtes_rules_read gives a section whole by its [id].\n\n", len(results), in.Query, min(n, len(results)))
		}
		for _, r := range results[:min(n, len(results))] {
			s := ix.sections[r.i]
			s.Notes = nil
			writeSection(&out, s, q, limit)
		}
		if len(results) > n {
			fmt.Fprintf(&out, "%d more sections match: add words, limit the documents (docs), or ask for the detail.\n\n", len(results)-n)
		}
		out.WriteString(groundRules)
		out.WriteString(DataLine(d.b))
		return text(out.String()), nil, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_rules_read",
		Title: "VTES rules section",
		Description: "Never state a rule from memory: quote the sections returned. " +
			"Reads Vampire: The Eternal Struggle (VTES) rules sections whole, by the [id] vtes_rules_search or vtes_rules_case gave " +
			"(e.g. advanced:2.1.3, rulebook:combat-ends): a section with every section under it, and the footnotes linking the posts it rests on.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"ids": {Type: "array", Items: &jsonschema.Schema{Type: "string"}, Description: "section ids, at most 5"},
			},
			Required: []string{"ids"},
		},
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: new(false)},
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in readInput) (*mcp.CallToolResult, any, error) {
		if len(in.IDs) == 0 || len(in.IDs) > 5 {
			return nil, nil, fmt.Errorf(`give 1 to 5 section ids, e.g. {"ids": ["advanced:2.1.3"]}; got %q`, in.IDs)
		}
		d := l.p.Load()
		ix := d.rules
		var out strings.Builder
		var left []string // the sections the bound left out
		for _, id := range in.IDs {
			i, ok := ix.byID[strings.TrimSpace(strings.Trim(id, "[]"))]
			if !ok {
				fmt.Fprintf(&out, "No section has the id %q: ids look like advanced:2.1.3 or rulebook:combat-ends; vtes_rules_search finds them.\n\n", id)
				continue
			}
			for _, s := range ix.subsections(i) {
				var b strings.Builder
				writeSection(&b, s, nil, 0)
				if len(left) > 0 || out.Len()+b.Len() > budget-2000 { // and the closing lines
					left = append(left, s.ID)
					continue
				}
				out.WriteString(b.String())
			}
		}
		if len(left) > 0 {
			shown := "they are " + strings.Join(left, ", ")
			if len(left) > 40 {
				shown = "the first 40 are " + strings.Join(left[:40], ", ")
			}
			fmt.Fprintf(&out, "Cut here, too long to read at once: %d sections are left out; %s.\n\n", len(left), shown)
		}
		out.WriteString(groundRules)
		out.WriteString(DataLine(d.b))
		return text(out.String()), nil, nil
	})
}

func text(s string) *mcp.CallToolResult {
	return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: s}}}
}

// proseForm folds a text as the rulebook is searched for a phrase: words of
// letters, digits and apostrophes, each between spaces.
func proseForm(s string) string {
	return " " + strings.Join(strings.FieldsFunc(bundle.Fold(s), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
	}), " ") + " "
}

// inRulebook says whether the rulebook uses a phrase other than within the
// given names.
func (ix *rulesIndex) inRulebook(phrase string, names []string) bool {
	n := strings.Count(ix.rulebook, proseForm(phrase))
	for _, name := range names {
		n -= strings.Count(ix.rulebook, proseForm(name))
	}
	return n > 0
}
