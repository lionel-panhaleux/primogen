package tools

import (
	"cmp"
	"context"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

const maxNames = 10

type cardsInput struct {
	Names  []string `json:"names" jsonschema:"card names, in English or translated, exact or approximate; at most 10"`
	Detail bool     `json:"detail,omitempty" jsonschema:"also list the printings and the translated texts"`
}

func addCards(s *mcp.Server, b *bundle.Bundle) {
	ix := newCardIndex(b.Cards)
	mcp.AddTool(s, &mcp.Tool{
		Name: "vtes_cards",
		Description: "VTES cards by name: the official text, the requirements and costs, legality, " +
			"every ruling of the rulings database with its rulings.krcg.org link and the archived post it cites, " +
			"and the card's entries in the history of card changes. Accepts translated and misspelled names; " +
			"an unknown name returns the closest names.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in cardsInput) (*mcp.CallToolResult, any, error) {
		if len(in.Names) == 0 || len(in.Names) > maxNames || slices.ContainsFunc(in.Names, func(n string) bool { return bundle.Fold(n) == "" }) {
			return nil, nil, fmt.Errorf(`give 1 to %d non-empty card names, e.g. {"names": ["Freak Drive", "Kiss of Ra"]}; got %q`, maxNames, in.Names)
		}
		var out strings.Builder
		for _, name := range in.Names {
			cards, note := ix.resolve(name)
			if note != "" {
				fmt.Fprintf(&out, "%s\n\n", note)
			}
			for _, c := range cards {
				writeCard(&out, c, in.Detail)
			}
		}
		out.WriteString(dataLine(b))
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: out.String()}}}, nil, nil
	})
}

// cardIndex maps folded names to cards: a vampire's printed name is shared by
// its groups. An English name wins over a translation.
type cardIndex struct {
	english, translated map[string][]*bundle.Card
	names               []string
}

func newCardIndex(cards []bundle.Card) *cardIndex {
	ix := &cardIndex{english: map[string][]*bundle.Card{}, translated: map[string][]*bundle.Card{}}
	add := func(m map[string][]*bundle.Card, name string, c *bundle.Card) {
		k := bundle.Fold(name)
		if k == "" || slices.Contains(m[k], c) {
			return
		}
		if ix.english[k] == nil && ix.translated[k] == nil {
			ix.names = append(ix.names, k)
		}
		m[k] = append(m[k], c)
	}
	for i := range cards {
		for _, n := range cards[i].Names() {
			add(ix.english, n, &cards[i])
		}
	}
	for i := range cards {
		for _, t := range cards[i].Translations {
			add(ix.translated, t.Name, &cards[i])
		}
	}
	return ix
}

func (ix *cardIndex) lookup(k string) []*bundle.Card {
	if cards := ix.english[k]; cards != nil {
		return cards
	}
	return ix.translated[k]
}

// resolve takes an exact match, else a misspelling one or two letters off a
// single card, else answers with the closest names and no card.
func (ix *cardIndex) resolve(name string) ([]*bundle.Card, string) {
	q := bundle.Fold(name)
	if cards := ix.lookup(q); cards != nil {
		return cards, ""
	}
	type candidate struct {
		name      string
		contained bool
		distance  int
	}
	cands := make([]candidate, len(ix.names))
	for i, n := range ix.names {
		cands[i] = candidate{n, strings.Contains(n, q), bundle.Distance(q, n)}
	}
	slices.SortFunc(cands, func(a, b candidate) int {
		return cmp.Or(cmp.Compare(a.distance, b.distance), cmp.Compare(len(a.name), len(b.name)), strings.Compare(a.name, b.name))
	})
	if best := cands[0]; best.distance <= 2 && utf8.RuneCountInString(q) >= 5 {
		cards := ix.lookup(best.name)
		unique := len(cards) == 1
		for _, c := range cands[1:] {
			if c.distance > best.distance {
				break
			}
			unique = unique && slices.Equal(ix.lookup(c.name), cards)
		}
		if unique {
			return cards, fmt.Sprintf("No card is named %q; this is %s, the one closest name.", name, cards[0].Name)
		}
	}
	slices.SortStableFunc(cands, func(a, b candidate) int {
		switch {
		case a.contained == b.contained:
			return 0
		case a.contained:
			return -1
		}
		return 1
	})
	var closest []string
	for _, c := range cands {
		for _, card := range ix.lookup(c.name) {
			if !slices.Contains(closest, card.Name) {
				closest = append(closest, card.Name)
			}
		}
		if len(closest) >= 5 {
			break
		}
	}
	return nil, fmt.Sprintf("No card is named %q. Closest names: %s. Call vtes_cards again with the intended one.",
		name, strings.Join(closest, "; "))
}

var reCardToken = regexp.MustCompile(`\{\d+\|([^}]*)\}`)

func writeCard(out *strings.Builder, c *bundle.Card, detail bool) {
	fmt.Fprintf(out, "## %s\n", c.Name)
	fmt.Fprintf(out, "Card %d · %s · %s\n", c.ID, strings.Join(c.Types, " / "), c.Image)
	for _, line := range requirements(c) {
		fmt.Fprintf(out, "%s\n", line)
	}
	if c.Banned != "" {
		fmt.Fprintf(out, "Banned since %s.\n", c.Banned)
	} else if c.Legal != "" {
		fmt.Fprintf(out, "Tournament legal since %s.\n", c.Legal)
	}
	out.WriteString("\nCard text (official):\n")
	for line := range strings.Lines(c.Text) {
		fmt.Fprintf(out, "> %s", line)
	}
	out.WriteString("\n\n")

	if len(c.Rulings) == 0 {
		out.WriteString("Rulings: the rulings database has none for this card.\n\n")
	} else {
		fmt.Fprintf(out, "Rulings (%d, official: the rulings database):\n", len(c.Rulings))
		for i, r := range c.Rulings {
			fmt.Fprintf(out, "%d. ", i+1)
			if r.Group != "" {
				fmt.Fprintf(out, "[group: %s] ", r.Group)
			}
			if r.Reminder {
				out.WriteString("[reminder of card text or a core rule] ")
			}
			fmt.Fprintf(out, "%s\n   Ruling: %s\n", reCardToken.ReplaceAllString(r.Text, "$1"), r.Link)
			for _, ref := range r.References {
				switch {
				case ref.Unarchived:
					fmt.Fprintf(out, "   %s: %s (the rulings archive does not hold this post)\n", ref.Label, ref.URL)
				case ref.Archive == "" || ref.Archive == ref.URL:
					fmt.Fprintf(out, "   %s: %s\n", ref.Label, ref.URL)
				default:
					fmt.Fprintf(out, "   %s: %s (archived from %s)\n", ref.Label, ref.Archive, ref.URL)
				}
			}
		}
		out.WriteString("\n")
	}

	if len(c.Changes) == 0 {
		out.WriteString("Changes: no line of the history of card changes names this card.\n\n")
	} else {
		out.WriteString("Changes (the history of card changes, a digest; the linked announcement settles each):\n")
		for _, ch := range c.Changes {
			fmt.Fprintf(out, "- %s, %s (%s): %s", ch.Date, ch.Announcement, ch.URL, ch.Text)
			switch {
			case ch.Shared:
				fmt.Fprintf(out, " [the history names %q, which several cards share: it may concern another]", ch.Named)
			case bundle.Fold(ch.Named) != bundle.Fold(c.Name) && bundle.Fold(ch.Named) != bundle.Fold(c.PrintedName):
				fmt.Fprintf(out, " [the history names %q]", ch.Named)
			}
			out.WriteString("\n")
		}
		out.WriteString("\n")
	}

	if detail {
		var sets []string
		for _, p := range c.Sets {
			sets = append(sets, strings.TrimSpace(p.Set+" "+p.Date))
		}
		fmt.Fprintf(out, "Printings: %s.\n", strings.Join(sets, ", "))
		for _, t := range c.Translations {
			fmt.Fprintf(out, "Translation (%s), %s:\n", t.Lang, t.Name)
			for line := range strings.Lines(t.Text) {
				fmt.Fprintf(out, "> %s", line)
			}
			out.WriteString("\n")
		}
		out.WriteString("\n")
	}
}

func requirements(c *bundle.Card) []string {
	var lines []string
	add := func(format string, a ...any) { lines = append(lines, fmt.Sprintf(format, a...)) }
	symbols := func(ds []string) string {
		s := make([]string, len(ds))
		for i, d := range ds {
			s[i] = "[" + d + "]"
		}
		return strings.Join(s, " ")
	}
	imbued := slices.Contains(c.Types, "Imbued")
	crypt := imbued || slices.Contains(c.Types, "Vampire")
	switch {
	case len(c.Clans) == 0:
	case imbued:
		add("Creed: %s.", strings.Join(c.Clans, ", "))
	case crypt:
		add("Clan: %s.", strings.Join(c.Clans, ", "))
	default:
		add("Clans its requirement names: %s.", strings.Join(c.Clans, ", "))
	}
	switch {
	case len(c.Disciplines) == 0:
	case imbued:
		add("Virtues: %s.", symbols(c.Disciplines))
	case crypt:
		add("Disciplines: %s (capitals: superior).", symbols(c.Disciplines))
	case c.Combo:
		add("Disciplines: %s, a combination card: each effect's symbols in the text say what it requires.", symbols(c.Disciplines))
	case c.Multidisc:
		add("Disciplines: %s, a multi-discipline card: each effect's symbols in the text say what it requires.", symbols(c.Disciplines))
	default:
		add("Requires %s.", symbols(c.Disciplines))
	}
	switch {
	case imbued:
		add("Life %d, group %s.", c.Capacity, c.Group)
	case crypt:
		add("Capacity %d, group %s.", c.Capacity, c.Group)
	}
	if c.Title != "" {
		add("Title: %s.", c.Title)
	}
	if c.Path != "" {
		add("Path: %s.", c.Path)
	}
	var costs []string
	for _, cost := range [][2]string{{c.BloodCost, "blood"}, {c.PoolCost, "pool"}, {c.ConvictionCost, "conviction"}} {
		if cost[0] != "" {
			costs = append(costs, cost[0]+" "+cost[1])
		}
	}
	if len(costs) > 0 {
		add("Cost: %s.", strings.Join(costs, ", "))
	}
	if c.BurnOption {
		add("Burn option.")
	}
	return lines
}
