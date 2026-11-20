package tools

import (
	"cmp"
	"context"
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

type metaInput struct {
	Archetype string `json:"archetype,omitempty"`
	Months    int    `json:"months,omitempty"`
	Detail    bool   `json:"detail,omitempty"`
}

var metaSchema = &jsonschema.Schema{
	Type: "object",
	Properties: map[string]*jsonschema.Schema{
		"archetype": {Type: "string", Description: `an archetype's name or ref, as this tool or vtes_deck names it ("Haqim Royalty", "C1"): its share over time, its core cards and its decks; omitted, the whole meta`},
		"months":    {Type: "integer", Description: "the window's length in months, ending at the newest tournament-winning deck; 12 when omitted"},
		"detail":    {Type: "boolean", Description: "longer lists"},
	},
}

const metaMonths = 12

// How many rows each list shows, by default and in detail.
var metaShown = map[bool]struct{ archetypes, moves, cards, cores, vampires, library, decks int }{
	false: {15, 5, 8, 5, 5, 8, 5},
	true:  {60, 10, 20, 12, 15, 30, 15},
}

// The bars of a card's share: in the move (of either window's players), and
// in an archetype's core (of its decks, library and crypt).
const (
	moveMin        = 0.05
	coreLibraryMin = 0.5
	coreVampireMin = 0.25
)

func addMeta(s *mcp.Server, l *Live) {
	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_meta",
		Title: "VTES meta",
		Description: "Call before any claim about the Vampire: The Eternal Struggle (VTES) meta: never say an archetype is popular, rising or dominant from memory. " +
			"Returns the archetypes' shares of the tournament-winning decks (TWDA) of the latest window, weighted by event size, " +
			"with the share in no archetype stated; what rises and falls against the previous window and the cards behind the move; " +
			"each leading archetype's core cards. With an archetype: its share over time, its core cards and its representative and newest decks.",
		InputSchema: metaSchema,
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: new(false)},
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in metaInput) (*mcp.CallToolResult, any, error) {
		d := l.p.Load()
		if d.rel == nil {
			return nil, nil, fmt.Errorf("this data bundle holds no archetype labels, built before it carried them: no meta until it is rebuilt")
		}
		m := d.meta()
		months := cmp.Or(in.Months, metaMonths)
		if months < 1 || months > m.maxMonths {
			return nil, nil, fmt.Errorf("months must be 1 to %d: the labels' window opens %s and a window needs the one before it; call vtes_meta again with months in that range",
				m.maxMonths, d.rel.Class.Since)
		}
		var arch *bundle.Archetype
		if in.Archetype != "" {
			var err error
			if arch, err = d.findArchetype(in.Archetype, nil, "vtes_meta"); err != nil {
				return nil, nil, err
			}
		}
		var out strings.Builder
		if arch == nil {
			writeMeta(&out, d, m, months, in.Detail)
		} else {
			writeArchetype(&out, d, m, arch, months, in.Detail)
		}
		out.WriteString("## How to answer\n\n" +
			"Read the meta from these figures: name the window, say the decks won their events (what wins, not what is played) and give the share in no archetype. " +
			"Call vtes_meta with an archetype for its core and its decks, vtes_deck with a deck's id for its list, vtes_card_stats for a card's play rates, " +
			"vtes_cards for any card the answer relies on.\n\n")
		if p := strings.TrimSpace(d.b.MetaReading); p != "" {
			out.WriteString(p)
		} else {
			out.WriteString("This data bundle holds no procedure for reading the meta, built before it carried one.")
		}
		out.WriteString("\n\n")
		out.WriteString("Answer from these figures, never from memory of the meta. Where they do not settle a point, say so rather than filling the gap.\n\n")
		out.WriteString(DataLine(d.b))
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: out.String()}}}, nil, nil
	})
}

// metaData places the window's decks once per bundle.
type metaData struct {
	since     time.Time // the labels' window opens after it
	newest    time.Time
	maxMonths int
	group     []int  // each window deck's archetype, -1 for none
	labelled  []bool // placed by the owner, archetype or noise
}

func (d *data) meta() *metaData {
	d.metaOnce.Do(func() {
		r := d.rel
		m := &metaData{group: make([]int, len(r.Decks)), labelled: make([]bool, len(r.Decks))}
		for i, t := range r.Decks {
			m.group[i], m.labelled[i] = r.Placement(i)
			if day, err := time.Parse(time.DateOnly, t.Date); err == nil && day.After(m.newest) {
				m.newest = day
			}
		}
		since, _ := time.Parse(time.DateOnly, r.Class.Since)
		m.since = since.AddDate(0, 0, -1)
		for !m.span(m.maxMonths+1, 1).from.Before(m.since) {
			m.maxMonths++
		}
		d.metaData = m
	})
	return d.metaData
}

// span is the dates after from and up to to.
type span struct{ from, to time.Time }

func (m *metaData) span(months, back int) span {
	return span{m.newest.AddDate(0, -months*(back+1), 0), m.newest.AddDate(0, -months*back, 0)}
}

func (s span) holds(date string) bool {
	day, err := time.Parse(time.DateOnly, date)
	return err == nil && day.After(s.from) && !day.After(s.to)
}

func (s span) String() string {
	return s.from.AddDate(0, 0, 1).Format(time.DateOnly) + " to " + s.to.Format(time.DateOnly)
}

// share is what a window holds of an archetype, or of no archetype; placed
// by similarity, not by the owner.
type share struct {
	decks, players, placed, placedPlayers int
}

func (s *share) add(players int, placed bool) {
	s.decks++
	s.players += players
	if placed {
		s.placed++
		s.placedPlayers += players
	}
}

// tally is a window's decks by archetype, a root's with its variants', and
// its cards' weights.
type tally struct {
	span
	decks, players, unweighted int
	groups, roots              map[*bundle.Archetype]*share
	noise, none                share          // the owner's noise; near no archetype
	cards                      map[string]int // players whose winner played the card, by feature
}

// of is a's share, its variants' with a root's.
func (t *tally) of(a *bundle.Archetype) share {
	s := t.groups[a]
	if a.VariantOf == "" {
		s = t.roots[a]
	}
	if s == nil {
		return share{}
	}
	return *s
}

func (t *tally) pct(players int) float64 {
	if t.players == 0 {
		return 0
	}
	return 100 * float64(players) / float64(t.players)
}

func (d *data) tally(m *metaData, s span) *tally {
	r := d.rel
	t := &tally{span: s, groups: map[*bundle.Archetype]*share{}, roots: map[*bundle.Archetype]*share{}, cards: map[string]int{}}
	for i, dk := range r.Decks {
		if !s.holds(dk.Date) {
			continue
		}
		t.decks++
		t.players += dk.Players
		if dk.Players == 0 {
			t.unweighted++
		}
		placed := !m.labelled[i]
		switch g := m.group[i]; {
		case g >= 0:
			a := &r.Class.Archetypes[g]
			for _, x := range []struct {
				shares map[*bundle.Archetype]*share
				a      *bundle.Archetype
			}{{t.groups, a}, {t.roots, r.Root(a)}} {
				if x.shares[x.a] == nil {
					x.shares[x.a] = &share{}
				}
				x.shares[x.a].add(dk.Players, placed)
			}
		case placed:
			t.none.add(dk.Players, true)
		default:
			t.noise.add(dk.Players, false)
		}
		seen := map[string]bool{}
		for _, c := range dk.Cards {
			if f := bundle.Feature(d.byID[c.ID]); !seen[f] {
				seen[f] = true
				t.cards[f] += dk.Players
			}
		}
	}
	return t
}

func (t *tally) ranked() []*bundle.Archetype {
	return slices.SortedFunc(maps.Keys(t.roots), func(a, b *bundle.Archetype) int {
		return cmp.Or(t.roots[b].players-t.roots[a].players, t.roots[b].decks-t.roots[a].decks, strings.Compare(a.Name, b.Name))
	})
}

func featureName(f string) string {
	if name, ok := strings.CutPrefix(f, "c:"); ok {
		return name + " (vampire)"
	}
	return strings.TrimPrefix(f, "l:")
}

func writeWindow(out *strings.Builder, d *data, cur, prev *tally, months int) {
	r := d.rel
	c := r.Class
	length := fmt.Sprintf("%d months", months)
	if months == 1 {
		length = "month"
	}
	fmt.Fprintf(out, "Window: the %s to %s, the newest deck of the TWDA: %s, %d decks won at events of %d players in all. "+
		"Previous window: %s, %d decks, %d players.\n",
		length, cur.to.Format(time.DateOnly), cur.span, cur.decks, cur.players, prev.span, prev.decks, prev.players)
	out.WriteString("A deck weighs its event's players: a share is the share of those players whose event a deck of the archetype won.")
	if n := cur.unweighted + prev.unweighted; n > 0 {
		fmt.Fprintf(out, " %d of these decks have no player count and weigh nothing.", n)
	}
	out.WriteString(" These decks won: a share says what wins, not what is played, and the full field is not served.\n")
	fmt.Fprintf(out, "Archetypes: codex-of-the-damned's labels (codex-of-the-damned.org), by its maintainer, pass of %s, curated and revisable, "+
		"a variant counted in its archetype. A deck the labels lack (every deck since that pass, and some before it) is placed in its nearest archetype "+
		"when at %.2f or more from it, the bar vtes_deck uses, which 95 in 100 of an archetype's own decks reach.\n\n", c.Generated, r.Threshold)
}

func writeMeta(out *strings.Builder, d *data, m *metaData, months int, detail bool) {
	r := d.rel
	shown := metaShown[detail]
	cur, prev := d.tally(m, m.span(months, 0)), d.tally(m, m.span(months, 1))
	out.WriteString("# The meta among tournament-winning decks\n\n")
	writeWindow(out, d, cur, prev, months)

	labelled, placed := 0, 0
	for _, s := range cur.roots {
		labelled += s.players - s.placedPlayers
		placed += s.placedPlayers
	}
	out.WriteString("## Shares\n\n")
	fmt.Fprintf(out, "Of the window's players, %.0f%% saw a deck the labels hold win, %.0f%% one placed by its nearest archetype, "+
		"and **%.0f%% one in no archetype**: %.0f%% the labels put in none (%d decks), %.0f%% near no archetype (%d decks: a new archetype, or a pass over the labels).\n\n",
		cur.pct(labelled), cur.pct(placed), cur.pct(cur.noise.players+cur.none.players),
		cur.pct(cur.noise.players), cur.noise.decks, cur.pct(cur.none.players), cur.none.decks)

	ranked := cur.ranked()
	out.WriteString("| # | Archetype | Share | Decks (placed by similarity) | Previous share | Change, in points |\n|---|---|---|---|---|---|\n")
	for i, a := range ranked[:min(len(ranked), shown.archetypes)] {
		s, p := cur.of(a), prev.of(a)
		fmt.Fprintf(out, "| %d | %s | %.1f%% | %d (%d) | %.1f%% | %+.1f |\n", i+1, archetypeName(r, a), cur.pct(s.players), s.decks, s.placed,
			prev.pct(p.players), cur.pct(s.players)-prev.pct(p.players))
	}
	if rest := ranked[min(len(ranked), shown.archetypes):]; len(rest) > 0 {
		players := 0
		for _, a := range rest {
			players += cur.of(a).players
		}
		fmt.Fprintf(out, "| | %d more archetypes (detail lists more) | %.1f%% together | | | |\n", len(rest), cur.pct(players))
	}
	fmt.Fprintf(out, "| | in no archetype | %.1f%% | %d (%d near none) | %.1f%% | %+.1f |\n\n",
		cur.pct(cur.noise.players+cur.none.players), cur.noise.decks+cur.none.decks, cur.none.decks,
		prev.pct(prev.noise.players+prev.none.players), cur.pct(cur.noise.players+cur.none.players)-prev.pct(prev.noise.players+prev.none.players))
	var variants []string
	for _, a := range ranked[:min(len(ranked), shown.archetypes)] {
		for i := range r.Class.Archetypes {
			if v := &r.Class.Archetypes[i]; v.VariantOf == a.Ref && cur.groups[v] != nil {
				variants = append(variants, fmt.Sprintf("%s (%s) %d of %s's %d", v.Name, v.Ref, cur.groups[v].decks, a.Name, cur.of(a).decks))
			}
		}
	}
	if len(variants) > 0 {
		fmt.Fprintf(out, "Their variants' decks in the window: %s.\n\n", strings.Join(variants, "; "))
	}

	writeMoves(out, r, cur, prev, shown.moves)
	writeCardMoves(out, cur, prev, shown.cards)

	fmt.Fprintf(out, "## The core of the leading archetypes\n\nFrom each archetype's decks since %s, labelled or placed, its variants' with it: "+
		"the vampires in a quarter of them or more and the library cards in half or more, the most played first "+
		"(call vtes_meta with the archetype for its whole core and its decks).\n\n", r.Class.Since)
	for _, a := range ranked[:min(len(ranked), shown.cores)] {
		decks := archetypeDecks(d, m, a)
		vampires, library := core(d, decks)
		fmt.Fprintf(out, "- %s, %d decks: vampires %s; library %s.\n", archetypeName(r, a), len(decks),
			coreLine(vampires, len(decks), shown.vampires), coreLine(library, len(decks), shown.library))
	}
	out.WriteString("\n")
}

func writeMoves(out *strings.Builder, r *bundle.Relatives, cur, prev *tally, n int) {
	all := map[*bundle.Archetype]float64{}
	for a := range cur.roots {
		all[a] = cur.pct(cur.of(a).players) - prev.pct(prev.of(a).players)
	}
	for a := range prev.roots {
		all[a] = cur.pct(cur.of(a).players) - prev.pct(prev.of(a).players)
	}
	byChange := slices.SortedFunc(maps.Keys(all), func(a, b *bundle.Archetype) int {
		return cmp.Or(cmp.Compare(all[b], all[a]), strings.Compare(a.Name, b.Name))
	})
	line := func(a *bundle.Archetype) {
		c, p := cur.of(a), prev.of(a)
		fmt.Fprintf(out, "- %s: %.1f%% → %.1f%% (%+.1f), %d → %d decks.\n", archetypeName(r, a), prev.pct(p.players), cur.pct(c.players), all[a], p.decks, c.decks)
	}
	out.WriteString("## What rises and falls\n\nAgainst the previous window, in points of share, a move under one point left out:\n\nRising:\n")
	k := 0
	for _, a := range byChange {
		if all[a] < 1 || k == n {
			break
		}
		line(a)
		k++
	}
	if k == 0 {
		out.WriteString("- none.\n")
	}
	out.WriteString("\nFalling:\n")
	k = 0
	for _, a := range slices.Backward(byChange) {
		if all[a] > -1 || k == n {
			break
		}
		line(a)
		k++
	}
	if k == 0 {
		out.WriteString("- none.\n")
	}
	out.WriteString("\n")
}

func writeCardMoves(out *strings.Builder, cur, prev *tally, n int) {
	change := map[string]float64{}
	for _, t := range []*tally{cur, prev} {
		for f, w := range t.cards {
			if float64(w) >= moveMin*float64(t.players) {
				change[f] = cur.pct(cur.cards[f]) - prev.pct(prev.cards[f])
			}
		}
	}
	byChange := slices.SortedFunc(maps.Keys(change), func(a, b string) int { return cmp.Or(cmp.Compare(change[b], change[a]), strings.Compare(a, b)) })
	fmt.Fprintf(out, "## The cards behind the move\n\nThe share of players whose event a deck playing the card won, against the previous window, "+
		"for the cards at %.0f%% or more in either; a card rises with the archetypes playing it:\n\nRising:\n", 100*moveMin)
	line := func(f string) {
		fmt.Fprintf(out, "- %s: %.1f%% → %.1f%% (%+.1f).\n", featureName(f), prev.pct(prev.cards[f]), cur.pct(cur.cards[f]), change[f])
	}
	k := 0
	for _, f := range byChange {
		if change[f] <= 0 || k == n {
			break
		}
		line(f)
		k++
	}
	if k == 0 {
		out.WriteString("- none.\n")
	}
	out.WriteString("\nFalling:\n")
	k = 0
	for _, f := range slices.Backward(byChange) {
		if change[f] >= 0 || k == n {
			break
		}
		line(f)
		k++
	}
	if k == 0 {
		out.WriteString("- none.\n")
	}
	out.WriteString("\n")
}

// archetypeDecks are the window's decks of a (its variants' with a root's),
// labelled or placed.
func archetypeDecks(d *data, m *metaData, a *bundle.Archetype) []int {
	r := d.rel
	var out []int
	for i := range r.Decks {
		if g := m.group[i]; g >= 0 {
			if b := &r.Class.Archetypes[g]; b == a || r.Root(b) == a {
				out = append(out, i)
			}
		}
	}
	return out
}

type coreCard struct {
	feature       string
	decks, copies int
}

// core counts the decks playing each card, vampires and library apart, the
// most played first.
func core(d *data, decks []int) (vampires, library []coreCard) {
	n, copies := map[string]int{}, map[string]int{}
	for _, i := range decks {
		seen := map[string]bool{}
		for _, c := range d.rel.Decks[i].Cards {
			f := bundle.Feature(d.byID[c.ID])
			copies[f] += c.Count
			if !seen[f] {
				seen[f] = true
				n[f]++
			}
		}
	}
	for _, f := range slices.SortedFunc(maps.Keys(n), func(a, b string) int { return cmp.Or(n[b]-n[a], strings.Compare(a, b)) }) {
		cc := coreCard{f, n[f], copies[f]}
		if strings.HasPrefix(f, "c:") {
			if float64(n[f]) >= coreVampireMin*float64(len(decks)) {
				vampires = append(vampires, cc)
			}
		} else if float64(n[f]) >= coreLibraryMin*float64(len(decks)) {
			library = append(library, cc)
		}
	}
	return vampires, library
}

func coreLine(cards []coreCard, of, n int) string {
	if len(cards) == 0 {
		return "none that often"
	}
	parts := make([]string, 0, n)
	for _, c := range cards[:min(len(cards), n)] {
		parts = append(parts, fmt.Sprintf("%s %.0f%%", strings.TrimSuffix(featureName(c.feature), " (vampire)"), 100*float64(c.decks)/float64(of)))
	}
	if k := len(cards) - n; k > 0 {
		parts = append(parts, fmt.Sprintf("and %d more", k))
	}
	return strings.Join(parts, ", ")
}

func writeArchetype(out *strings.Builder, d *data, m *metaData, a *bundle.Archetype, months int, detail bool) {
	r := d.rel
	shown := metaShown[detail]
	cur, prev := d.tally(m, m.span(months, 0)), d.tally(m, m.span(months, 1))
	root := r.Root(a)
	fmt.Fprintf(out, "# %s in the meta\n\n", archetypeName(r, a))
	var variants []string
	for i := range r.Class.Archetypes {
		if v := &r.Class.Archetypes[i]; v.VariantOf == a.Ref {
			variants = append(variants, fmt.Sprintf("%s (%s)", v.Name, v.Ref))
		}
	}
	if len(variants) > 0 {
		fmt.Fprintf(out, "Its variants, counted with it: %s.\n\n", strings.Join(variants, ", "))
	}
	writeWindow(out, d, cur, prev, months)

	out.WriteString("## Share\n\n")
	c, p := cur.of(a), prev.of(a)
	fmt.Fprintf(out, "In the window: %.1f%% of the players, %d decks (%d placed by similarity)", cur.pct(c.players), c.decks, c.placed)
	if a == root {
		if k := slices.Index(cur.ranked(), a); k >= 0 {
			fmt.Fprintf(out, ", ranked %d of %d archetypes", k+1, len(cur.roots))
		}
	} else {
		fmt.Fprintf(out, ", of %s's %.1f%%", root.Name, cur.pct(cur.of(root).players))
	}
	fmt.Fprintf(out, ". Previous window: %.1f%%, %d decks. Change: %+.1f points.\n\n", prev.pct(p.players), p.decks, cur.pct(c.players)-prev.pct(p.players))
	fmt.Fprintf(out, "Over the whole windows of %d months since the labels' window opened, oldest first:\n", months)
	var history []string
	for back := 0; !m.span(months, back).from.Before(m.since); back++ {
		t := d.tally(m, m.span(months, back))
		history = append(history, fmt.Sprintf("- %s: %.1f%%, %d of %d decks.\n", t.span, t.pct(t.of(a).players), t.of(a).decks, t.decks))
	}
	for _, h := range slices.Backward(history) {
		out.WriteString(h)
	}
	out.WriteString("\n")

	decks := archetypeDecks(d, m, a)
	vampires, library := core(d, decks)
	labelled := 0
	for _, i := range decks {
		if m.labelled[i] {
			labelled++
		}
	}
	fmt.Fprintf(out, "## Its core\n\nFrom its %d decks since %s (%d labelled, %d placed by similarity), each card with the share of those decks playing it "+
		"and the copies they play on average.\n\nVampires in a quarter of them or more:\n", len(decks), r.Class.Since, labelled, len(decks)-labelled)
	coreRows := func(cards []coreCard, n int) {
		if len(cards) == 0 {
			out.WriteString("- none that often.\n")
		}
		for _, c := range cards[:min(len(cards), n)] {
			fmt.Fprintf(out, "- %s: %.0f%%, %.1f copies.\n", strings.TrimSuffix(featureName(c.feature), " (vampire)"), 100*float64(c.decks)/float64(len(decks)),
				float64(c.copies)/float64(c.decks))
		}
		if k := len(cards) - n; k > 0 {
			fmt.Fprintf(out, "- and %d more (detail lists more).\n", k)
		}
	}
	coreRows(vampires, shown.vampires)
	out.WriteString("\nLibrary cards in half of them or more:\n")
	coreRows(library, shown.library)
	out.WriteString("\n## Its decks\n\nCall vtes_deck with an id for the whole list.\n\n")
	deckLine := func(i int) {
		t := r.Decks[i]
		fmt.Fprintf(out, "- %s, %s: won by %s at %s, %s, %s, %d players.", t.ID, cmp.Or(t.Name, "(unnamed)"), t.Player, t.Event, t.Place, t.Date, t.Players)
		g := &r.Class.Archetypes[m.group[i]]
		switch {
		case !m.labelled[i] && g != a:
			fmt.Fprintf(out, " Placed by similarity in %s.", archetypeName(r, g))
		case !m.labelled[i]:
			out.WriteString(" Placed by similarity.")
		case g != a:
			fmt.Fprintf(out, " Labelled %s.", archetypeName(r, g))
		}
		if t.URL != "" {
			fmt.Fprintf(out, " Event report: %s", t.URL)
		}
		out.WriteString("\n")
	}
	var reps, newest []int
	for _, i := range decks {
		if slices.Contains(r.Class.Archetypes[m.group[i]].Representative, r.Decks[i].ID) {
			reps = append(reps, i)
		}
		if cur.holds(r.Decks[i].Date) {
			newest = append(newest, i)
		}
	}
	if len(reps) > 0 {
		out.WriteString("The representative decks the labels pick:\n")
		for _, i := range reps {
			deckLine(i)
		}
	} else {
		out.WriteString("The labels pick no representative deck.\n")
	}
	slices.SortStableFunc(newest, func(x, y int) int {
		return cmp.Or(strings.Compare(r.Decks[y].Date, r.Decks[x].Date), r.Decks[y].Players-r.Decks[x].Players)
	})
	if len(newest) == 0 {
		out.WriteString("\nNo deck of it won in the window.\n\n")
		return
	}
	fmt.Fprintf(out, "\nIts newest decks in the window, of %d:\n", len(newest))
	for _, i := range newest[:min(len(newest), shown.decks)] {
		deckLine(i)
	}
	out.WriteString("\n")
}
