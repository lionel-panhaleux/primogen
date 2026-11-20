package tools

import (
	"cmp"
	"context"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

type statsInput struct {
	Card      string `json:"card"`
	Archetype string `json:"archetype,omitempty"`
	Detail    bool   `json:"detail,omitempty"`
}

var statsSchema = &jsonschema.Schema{
	Type: "object",
	Properties: map[string]*jsonschema.Schema{
		"card":      {Type: "string", Description: "one card name, in English or translated, exact or approximate"},
		"archetype": {Type: "string", Description: `an archetype's name or ref, as vtes_deck names it ("Endless Dance", "C4"): the statistics within its decks; omitted, over every deck of the window`},
		"detail":    {Type: "boolean", Description: "longer lists"},
	},
	Required: []string{"card"},
}

// How many rows each list shows, by default and in detail.
var statsShown = map[bool]struct{ archetypes, with, vampires, competitors int }{
	false: {5, 10, 5, 8},
	true:  {12, 25, 12, 20},
}

func addStats(s *mcp.Server, l *Live) {
	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_card_stats",
		Title: "VTES card statistics",
		Description: "Call to survey a slot of a Vampire: The Eternal Struggle (VTES) deck before advising a card: never endorse a card from memory of how often it is played. " +
			"Returns, for one card among the tournament-winning decks (TWDA) of the archetype labels' window, overall or within an archetype: " +
			"how many decks play it and with how many copies, the archetypes playing it most, the cards and vampires played with it, " +
			"and its competitors: the cards of its type and timing that the decks like those playing it play, with and without it, each with its official text.",
		InputSchema: statsSchema,
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: new(false)},
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in statsInput) (*mcp.CallToolResult, any, error) {
		if bundle.Fold(in.Card) == "" {
			return nil, nil, fmt.Errorf(`give one card name, e.g. {"card": "Awe", "archetype": "Endless Dance"}`)
		}
		d := l.p.Load()
		r := d.rel
		if r == nil {
			return nil, nil, fmt.Errorf("this data bundle holds no archetype labels, built before it carried them: no statistics until it is rebuilt")
		}
		cards, note := d.ix.resolve(in.Card, "vtes_card_stats")
		if cards == nil {
			return nil, nil, fmt.Errorf("%s", note)
		}
		var arch *bundle.Archetype
		if in.Archetype != "" {
			var err error
			if arch, err = d.findArchetype(in.Archetype, cards[0]); err != nil {
				return nil, nil, err
			}
		}
		var out strings.Builder
		if note != "" {
			fmt.Fprintf(&out, "%s\n\n", note)
		}
		writeStats(&out, d, cards[0], arch, in.Detail)
		out.WriteString(DataLine(d.b))
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: out.String()}}}, nil, nil
	})
}

// findArchetype takes a name or a ref; for a name no archetype has, it offers
// the archetypes playing the card, then the names holding a word of it.
func (d *data) findArchetype(name string, c *bundle.Card) (*bundle.Archetype, error) {
	r := d.rel
	q := bundle.Fold(name)
	as := r.Class.Archetypes
	var found []*bundle.Archetype
	for i := range as {
		if bundle.Fold(as[i].Name) == q || bundle.Fold(as[i].Ref) == q {
			found = append(found, &as[i])
		}
	}
	if len(found) == 1 {
		return found[0], nil
	}
	if len(found) > 1 {
		names := make([]string, len(found))
		for i, a := range found {
			names[i] = archetypeName(r, a)
		}
		return nil, fmt.Errorf("several archetypes are named %q: %s. Call vtes_card_stats again with the ref of the one meant", name, strings.Join(names, "; "))
	}
	names := func(as []*bundle.Archetype) string {
		out := make([]string, len(as))
		for i, a := range as {
			out[i] = archetypeName(r, a)
		}
		return strings.Join(out, "; ")
	}
	playing, _, _, _ := d.playedBy(c)
	var named []*bundle.Archetype
	for i := range as {
		words := strings.Fields(bundle.Fold(as[i].Name))
		if len(named) < 3 && slices.ContainsFunc(strings.Fields(q), func(w string) bool { return len(w) > 2 && slices.Contains(words, w) }) {
			named = append(named, &as[i])
		}
	}
	msg := fmt.Sprintf("no archetype is named %q.", name)
	if len(playing) > 0 {
		msg += fmt.Sprintf(" The archetypes playing %s most: %s.", c.Name, names(playing[:min(5, len(playing))]))
	}
	if len(named) > 0 {
		msg += " Those whose name holds a word of it: " + names(named) + "."
	} else {
		msg += " None has a name holding a word of it."
	}
	return nil, fmt.Errorf("%s Call vtes_card_stats again with a name or ref, or vtes_deck on a list to find its archetype", msg)
}

// playedBy counts the window's decks playing c by the root of their
// archetype, the roots most played first, with each root's size (its
// variants' decks counted), and those decks in no archetype.
func (d *data) playedBy(c *bundle.Card) (roots []*bundle.Archetype, n, size map[*bundle.Archetype]int, unlabelled int) {
	r := d.rel
	n, size = map[*bundle.Archetype]int{}, map[*bundle.Archetype]int{}
	for _, t := range r.Decks {
		a := r.Label(t.ID)
		if a != nil {
			a = r.Root(a)
			size[a]++
		}
		switch {
		case d.plays(t, c) == 0:
		case a == nil:
			unlabelled++
		default:
			n[a]++
		}
	}
	roots = slices.SortedFunc(maps.Keys(n), func(a, b *bundle.Archetype) int { return cmp.Or(n[b]-n[a], strings.Compare(a.Name, b.Name)) })
	return roots, n, size, unlabelled
}

// plays is how many copies of c a deck holds, a vampire counted in every
// group printed under its name.
func (d *data) plays(t *bundle.Deck, c *bundle.Card) int {
	n := 0
	for _, dc := range t.Cards {
		if bundle.Feature(d.byID[dc.ID]) == bundle.Feature(c) {
			n += dc.Count
		}
	}
	return n
}

var reUsable = regexp.MustCompile(`(?i)\bonly usable [^.]*`)

// timing is the card's "Only usable ..." clause, lowercased, or "".
func timing(c *bundle.Card) string {
	return strings.ToLower(reUsable.FindString(c.Text))
}

// slotOf is how much of c's slot a library card shares: its type and its
// timing clause, one clause opening the other ("before votes are cast"
// narrows it), or its type only.
func slotOf(c, o *bundle.Card) int {
	if o.ID == c.ID || o.IsCrypt() || !slices.ContainsFunc(o.Types, func(t string) bool { return slices.Contains(c.Types, t) }) {
		return noSlot
	}
	a, b := timing(c), timing(o)
	if a == b || a != "" && b != "" && (strings.HasPrefix(a, b) || strings.HasPrefix(b, a)) {
		return sameSlot
	}
	return typeOnly
}

const (
	noSlot = iota
	typeOnly
	sameSlot
)

func writeStats(out *strings.Builder, d *data, c *bundle.Card, arch *bundle.Archetype, detail bool) {
	r := d.rel
	shown := statsShown[detail]
	in := func(t *bundle.Deck) bool {
		if arch == nil {
			return true
		}
		a := r.Label(t.ID)
		return a != nil && (a == arch || r.Root(a) == arch)
	}
	var scope, its []*bundle.Deck
	copies := map[int]int{}
	for _, t := range r.Decks {
		if !in(t) {
			continue
		}
		scope = append(scope, t)
		if n := d.plays(t, c); n > 0 {
			its = append(its, t)
			copies[n]++
		}
	}

	name := c.Name
	if c.IsCrypt() {
		name = c.PrintedName
	}
	fmt.Fprintf(out, "# %s among tournament-winning decks\n\n", name)
	fmt.Fprintf(out, "Card %d · %s · %s\n", c.ID, strings.Join(c.Types, " / "), c.Image)
	if c.Banned != "" {
		fmt.Fprintf(out, "Banned since %s.\n", c.Banned)
	}
	for _, line := range requirements(c) {
		fmt.Fprintf(out, "%s\n", line)
	}
	quote(out, c.Text)
	window := fmt.Sprintf("the %d TWDA decks since %s", len(r.Decks), r.Class.Since)
	if arch != nil {
		what := "its decks"
		if slices.ContainsFunc(r.Class.Archetypes, func(a bundle.Archetype) bool { return a.VariantOf == arch.Ref }) {
			what = "its decks and its variants'"
		}
		fmt.Fprintf(out, "Scope: the archetype %s, %s among %s: %d decks. ", archetypeName(r, arch), what, window, len(scope))
	} else {
		fmt.Fprintf(out, "Scope: %s. ", window)
	}
	out.WriteString("Archetypes: codex-of-the-damned's labels (codex-of-the-damned.org), curated and revisable. " +
		"These decks won a tournament: a count says what winners play, not what is strong.\n\n")

	out.WriteString("## Play rate\n\n")
	if len(its) == 0 {
		fmt.Fprintf(out, "No deck of the scope plays %s.\n\n", name)
	} else {
		total := 0
		var counts []string
		for _, n := range slices.Sorted(maps.Keys(copies)) {
			total += n * copies[n]
			counts = append(counts, fmt.Sprintf("%d× in %d", n, copies[n]))
		}
		fmt.Fprintf(out, "Played by %d of the %d decks (%.0f%%), %.1f copies on average: %s.\n\n",
			len(its), len(scope), 100*float64(len(its))/float64(len(scope)), float64(total)/float64(len(its)), strings.Join(counts, ", "))
	}
	if arch == nil && len(its) > 0 {
		roots, n, size, unlabelled := d.playedBy(c)
		out.WriteString("The archetypes playing it most, each with its variants (call again with an archetype for the statistics within it):\n")
		for _, a := range roots[:min(len(roots), shown.archetypes)] {
			fmt.Fprintf(out, "- %s: %d of its %d decks.\n", archetypeName(r, a), n[a], size[a])
		}
		if k := len(roots) - shown.archetypes; k > 0 {
			fmt.Fprintf(out, "- and %d more archetypes.\n", k)
		}
		if unlabelled > 0 {
			fmt.Fprintf(out, "- in no archetype: %d.\n", unlabelled)
		}
		out.WriteString("\n")
	}
	switch len(its) {
	case 0:
	case 1:
		fmt.Fprintf(out, "## Played with it\n\nOne deck plays it, %s: its cards are that deck's list (vtes_deck %[1]s).\n\n", its[0].ID)
	default:
		writeWith(out, d, c, scope, its, shown.with, shown.vampires)
	}
	if !c.IsCrypt() {
		writeCompetitors(out, d, c, arch, scope, its, shown.competitors)
	}
	out.WriteString("## How to answer\n\n" +
		"Call vtes_cards for the rulings of any card the advice relies on, " +
		"vtes_deck for a deck these figures should be read against.\n\n")
	out.WriteString(strings.TrimSpace(d.b.DeckAdvice))
	out.WriteString("\n\nAnswer from these figures and texts, never from memory of the cards. " +
		"Where neither a text nor a ruling settles a point, search vtes_rulings_archive (the card's name, directors only) before saying it is unsettled.\n\n")
}

// writeWith names the library cards most particular to the decks playing c,
// and their vampires.
func writeWith(out *strings.Builder, d *data, c *bundle.Card, scope, its []*bundle.Deck, nCards, nVampires int) {
	count := func(decks []*bundle.Deck, key func(*bundle.Card) string) map[string]int {
		n := map[string]int{}
		for _, t := range decks {
			seen := map[string]bool{}
			for _, dc := range t.Cards {
				if o := d.byID[dc.ID]; key(o) != "" && bundle.Feature(o) != bundle.Feature(c) && !seen[key(o)] {
					k := key(o)
					seen[k] = true
					n[k]++
				}
			}
		}
		return n
	}
	library := func(o *bundle.Card) string {
		if o.IsCrypt() {
			return ""
		}
		return o.Name
	}
	vampire := func(o *bundle.Card) string {
		if !o.IsCrypt() {
			return ""
		}
		return o.PrintedName
	}
	share := func(n, of int) string { return fmt.Sprintf("%d (%.0f%%)", n, 100*float64(n)/float64(of)) }

	here, all := count(its, library), count(scope, library)
	lift := func(k string) float64 {
		return float64(here[k]) / float64(len(its)) / (float64(all[k]) / float64(len(scope)))
	}
	var common []string
	for k, n := range here {
		if 3*n >= len(its) {
			common = append(common, k)
		}
	}
	slices.SortFunc(common, func(a, b string) int {
		return cmp.Or(cmp.Compare(lift(b), lift(a)), here[b]-here[a], strings.Compare(a, b))
	})
	fmt.Fprintf(out, "## Played with it\n\nThe library cards in a third or more of the %d decks playing it, those most particular to them first "+
		"(how many times more often they play it than the scope's decks do; near 1, a staple of the scope):\n", len(its))
	for _, k := range common[:min(len(common), nCards)] {
		fmt.Fprintf(out, "- %s: in %s, ×%.1f.\n", k, share(here[k], len(its)), lift(k))
	}
	if n := len(common) - nCards; n > 0 {
		fmt.Fprintf(out, "- and %d more.\n", n)
	}
	vamps := count(its, vampire)
	names := slices.SortedFunc(maps.Keys(vamps), func(a, b string) int { return cmp.Or(vamps[b]-vamps[a], strings.Compare(a, b)) })
	parts := make([]string, 0, nVampires)
	for _, k := range names[:min(len(names), nVampires)] {
		parts = append(parts, fmt.Sprintf("%s %s", k, share(vamps[k], len(its))))
	}
	fmt.Fprintf(out, "\nThe vampires most often in their crypts: %s.\n\n", strings.Join(parts, "; "))
}

// writeCompetitors lists the cards of c's slot that the decks like those
// playing c play, those playing them without c first.
func writeCompetitors(out *strings.Builder, d *data, c *bundle.Card, arch *bundle.Archetype, scope, its []*bundle.Deck, n int) {
	out.WriteString("## Competitors for its slot\n\n")
	slot := "the cards sharing a type with it and its timing (" + timing(c) + ")"
	if timing(c) == "" {
		slot = "the cards sharing a type with it and, like it, no \"only usable\" clause: its type alone defines the slot, so the list is broad"
		if arch == nil {
			slot += " (within an archetype, its decks narrow it)"
		}
	}
	near := scope
	if arch == nil {
		if len(its) == 0 {
			fmt.Fprintf(out, "No deck plays it: no decks to read its competitors from. Call again with an archetype to read %s in its decks.\n\n", slot)
			return
		}
		ids := map[string]bool{}
		for _, t := range its {
			ids[t.ID] = true
		}
		near = d.rel.Around(ids, max(2*len(its), 20))
		if len(near) == len(d.rel.Decks) {
			fmt.Fprintf(out, "In every deck of the window, %s, those played without it first:\n", slot)
		} else {
			fmt.Fprintf(out, "In the %d decks nearest those playing it (by the similarity vtes_deck uses), %s, those played without it first:\n", len(near), slot)
		}
	} else {
		fmt.Fprintf(out, "In the archetype's decks, %s, those played without it first:\n", slot)
	}
	type row struct {
		card          *bundle.Card
		with, without int
	}
	rows := [3]map[int]*row{nil, {}, {}}
	nWith := 0
	for _, t := range near {
		with := d.plays(t, c) > 0
		if with {
			nWith++
		}
		seen := map[int]bool{}
		for _, dc := range t.Cards {
			o := d.byID[dc.ID]
			k := slotOf(c, o)
			if seen[o.ID] || k == noSlot {
				continue
			}
			seen[o.ID] = true
			if rows[k][o.ID] == nil {
				rows[k][o.ID] = &row{card: o}
			}
			if with {
				rows[k][o.ID].with++
			} else {
				rows[k][o.ID].without++
			}
		}
	}
	fmt.Fprintf(out, "(%d of these decks play it, %d do not.)\n\n", nWith, len(near)-nWith)
	decks := func(n int) string {
		if n == 1 {
			return "1 deck"
		}
		return fmt.Sprintf("%d decks", n)
	}
	write := func(m map[int]*row, n int) {
		list := slices.SortedFunc(maps.Values(m), func(a, b *row) int {
			return cmp.Or(b.without-a.without, b.with-a.with, strings.Compare(a.card.Name, b.card.Name))
		})
		for _, x := range list[:min(len(list), n)] {
			fmt.Fprintf(out, "- **%s** (card %d, %s): in %s without it, %s with it.", x.card.Name, x.card.ID, strings.Join(x.card.Types, " / "),
				decks(x.without), decks(x.with))
			for _, line := range requirements(x.card) {
				fmt.Fprintf(out, " %s", line)
			}
			if x.card.Banned != "" {
				fmt.Fprintf(out, " **Banned since %s.**", x.card.Banned)
			}
			out.WriteString("\n")
			quote(out, x.card.Text)
		}
		if k := len(list) - n; k > 0 {
			fmt.Fprintf(out, "- and %d more (detail lists more).\n", k)
		}
	}
	if len(rows[sameSlot]) == 0 {
		out.WriteString("None: no deck there plays another card of its type and timing.\n")
	}
	write(rows[sameSlot], n)
	// below three, the timing is too particular to make a slot alone
	if rest := n - len(rows[sameSlot]); len(rows[sameSlot]) < 3 && len(rows[typeOnly]) > 0 {
		out.WriteString("\nFew cards share its timing; those of its type with another timing, which may fill the slot another way:\n")
		write(rows[typeOnly], rest)
	}
	out.WriteString("\n")
}
