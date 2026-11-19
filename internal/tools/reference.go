package tools

import (
	"cmp"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

// The reference: an archetype or the nearest decks only from this many decks.
const referenceMin = 10

// Staples are in this share of the reference's decks; few play a card in
// this share or less.
const (
	stapleShare = 0.75
	fewShare    = 0.10
	fewShown    = 15
)

// actionTypes are the library types needing a minion's action, in the order shown.
var actionTypes = []string{"Action", "Political Action", "Equipment", "Ally", "Retainer"}

// reTrifle finds a trifle in a master card's opening line ("Master: unique trifle.").
var reTrifle = regexp.MustCompile(`(?i)\btrifle\b`)

// measures are a list's counts, its library's.
type measures struct {
	lib, action, masters, trifles, blood, pool, distinct, top10, noModule int
	byType                                                                map[string]int
	copies                                                                map[int]int    // library card id → copies
	gains                                                                 map[string]int // economy kind → copies
}

// figure is a measure as the section shows it, with what the list's own
// value adds.
type figure struct {
	name   string
	get    func(m *measures) int
	detail func(m *measures) string
}

// budgets are also what the outlier line counts.
var budgets = []figure{
	{"Cards needing a minion's action (action, political action, equipment, ally, retainer)", func(m *measures) int { return m.action }, func(m *measures) string {
		var types []string
		for _, t := range actionTypes {
			if k := m.byType[t]; k > 0 {
				types = append(types, fmt.Sprintf("; %s %d", strings.ToLower(t), k))
			}
		}
		return strings.Join(types, "")
	}},
	{"Master cards, trifles included", func(m *measures) int { return m.masters }, func(m *measures) string { return fmt.Sprintf("; trifles %d", m.trifles) }},
	{"Blood the library costs, every copy, X costs left out", func(m *measures) int { return m.blood }, nil},
	{"Pool the library costs, every copy, X costs left out", func(m *measures) int { return m.pool }, nil},
}

var concentration = []figure{
	{"Distinct library cards", func(m *measures) int { return m.distinct }, nil},
	{"Copies in the ten largest playsets", func(m *measures) int { return m.top10 }, nil},
	{"Copies in no module shown", func(m *measures) int { return m.noModule }, nil},
}

func measure(cards map[int]int, byID map[int]*bundle.Card, modules []bundle.Module, economy []bundle.Gain) *measures {
	m := &measures{byType: map[string]int{}, copies: map[int]int{}, gains: map[string]int{}}
	for _, g := range economy {
		for _, t := range g.Cards {
			for _, id := range t.IDs {
				m.gains[g.Kind] += cards[id]
			}
		}
	}
	tagged := map[int]bool{}
	for _, mod := range modules {
		if !moduleHeld(mod, func(id int) bool { return cards[id] > 0 }) {
			continue
		}
		for _, t := range slices.Concat(mod.Core, mod.Support) {
			for _, id := range t.IDs {
				tagged[id] = true
			}
		}
	}
	var counts []int
	for id, n := range cards {
		c := byID[id]
		if c.IsCrypt() {
			continue
		}
		m.lib += n
		m.copies[id] += n
		counts = append(counts, n)
		if i := slices.IndexFunc(actionTypes, func(t string) bool { return slices.Contains(c.Types, t) }); i >= 0 {
			m.action += n
			m.byType[actionTypes[i]] += n
		}
		if slices.Contains(c.Types, "Master") {
			m.masters += n
			if first, _, _ := strings.Cut(c.Text, "\n"); reTrifle.MatchString(first) {
				m.trifles += n
			}
		}
		m.blood += n * atoi(c.BloodCost)
		m.pool += n * atoi(c.PoolCost)
		if !tagged[id] {
			m.noModule += n
		}
	}
	m.distinct = len(counts)
	slices.Sort(counts)
	slices.Reverse(counts)
	for _, n := range counts[:min(10, len(counts))] {
		m.top10 += n
	}
	return m
}

// windowMeasures are the measures of the archetype labels' window's decks, by id.
func (d *data) windowMeasures() map[string]*measures {
	d.measuresOnce.Do(func() {
		d.measures = map[string]*measures{}
		for _, t := range d.rel.Decks {
			cards := map[int]int{}
			for _, c := range t.Cards {
				cards[c.ID] += c.Count
			}
			d.measures[t.ID] = measure(cards, d.byID, d.b.Modules, d.b.Economy)
		}
	})
	return d.measures
}

// reference are the decks a list is set beside, and what they are.
type reference struct {
	what  string
	decks []*bundle.Deck
}

// referenceOf is the nearest archetype with its variants, when the list is
// within the bar and it holds enough decks; else the nearest decks within
// the bar, when enough are; else the whole window, loose. id is left out.
func (d *data) referenceOf(cards []bundle.DeckCard, id string) reference {
	r := d.rel
	window := fmt.Sprintf("the %d TWDA decks since %s", len(r.Decks), r.Class.Since)
	v := r.Vector(cards)
	if !v.Empty() {
		best := r.Archetypes(v, id, 1)[0]
		if best.Similarity >= r.Threshold {
			a := &r.Class.Archetypes[best.Index]
			var decks []*bundle.Deck
			for _, t := range r.Decks {
				if l := r.Label(t.ID); l != nil && t.ID != id && (l == a || r.Root(l) == a) {
					decks = append(decks, t)
				}
			}
			if len(decks) >= referenceMin {
				what := "its decks"
				if slices.ContainsFunc(r.Class.Archetypes, func(o bundle.Archetype) bool { return o.VariantOf == a.Ref }) {
					what = "its decks and its variants'"
				}
				return reference{what: fmt.Sprintf("its nearest archetype, %s, %s among %s: %d decks", archetypeName(r, a), what, window, len(decks)), decks: decks}
			}
		}
		var decks []*bundle.Deck
		for _, m := range r.Nearest(v, id, len(r.Decks)) {
			if m.Similarity < r.Threshold {
				break
			}
			decks = append(decks, r.Decks[m.Index])
		}
		if len(decks) >= referenceMin {
			return reference{what: fmt.Sprintf("its %d nearest decks among %s, those at %.2f or more from it (no archetype of %d decks or more is that near)",
				len(decks), window, r.Threshold, referenceMin), decks: decks}
		}
	}
	var decks []*bundle.Deck
	for _, t := range r.Decks {
		if t.ID != id {
			decks = append(decks, t)
		}
	}
	return reference{what: fmt.Sprintf("all %s: no archetype of %d decks or more is at %.2f or more from it, nor are %d of its nearest decks. "+
		"This reference is loose, a mix of archetypes: a figure far from it says the list is unlike most winners, not what it lacks",
		window, referenceMin, r.Threshold, referenceMin), decks: decks}
}

// writeReference sets the list beside its reference decks, and returns
// them, nil when it sets none.
func (dk *deck) writeReference(out *strings.Builder, d *data, fit int) *reference {
	out.WriteString("## Beside its reference\n\n")
	_, _, _, nLibrary := dk.split()
	switch {
	case d.rel == nil:
		out.WriteString("This data bundle holds no archetype labels: no reference until it is rebuilt.\n\n")
		return nil
	case dk.draft:
		out.WriteString("A deck idea is not set beside winning lists: its sizes are not a list's.\n\n")
		return nil
	case nLibrary == 0:
		out.WriteString("No library card given: nothing to set beside a reference.\n\n")
		return nil
	}
	cards := map[int]int{}
	var list []bundle.DeckCard
	for _, e := range dk.entries {
		cards[e.card.ID] += e.count
		list = append(list, bundle.DeckCard{ID: e.card.ID, Count: e.count})
	}
	id := ""
	if dk.twda != nil {
		id = dk.twda.ID
	}
	ref := d.referenceOf(list, id)
	all := d.windowMeasures()
	var ms []*measures
	for _, t := range ref.decks {
		ms = append(ms, all[t.ID])
	}
	here := measure(cards, d.byID, d.b.Modules, d.b.Economy)
	n := len(ms)

	fmt.Fprintf(out, "The reference: %s. These decks won a tournament: their figures say what winners play, not what is strong. "+
		"Each figure is this list's, then the reference decks' as min, 10%%, 25%%, median, 75%%, 90%% and max, then how many of them are above and below it by count.\n\n", ref.what)
	fmt.Fprintf(out, "- Library: %d cards; the reference's %s.\n", here.lib, spread(values(ms, func(m *measures) float64 { return float64(m.lib) })))

	out.WriteString("\n### Budgets\n\n")
	lines := func(figures []figure) {
		for _, f := range figures {
			v, detail := f.get(here), ""
			if f.detail != nil {
				detail = f.detail(here)
			}
			fmt.Fprintf(out, "- %s: %d, %.0f per 100 library cards%s. Reference: %s; per 100 library cards: %s; %s.\n", f.name, v, pct(v, here.lib), detail,
				spread(values(ms, func(m *measures) float64 { return float64(f.get(m)) })),
				spread(values(ms, func(m *measures) float64 { return pct(f.get(m), m.lib) })),
				rank(values(ms, func(m *measures) float64 { return float64(f.get(m)) }), float64(v)))
		}
	}
	lines(budgets)
	outside := 0
	for i, m := range ms {
		for _, b := range budgets {
			others := sortedWithout(ms, i, func(m *measures) float64 { return float64(b.get(m)) })
			if v := float64(b.get(m)); v < others[0] || v > others[len(others)-1] {
				outside++
				break
			}
		}
	}
	fmt.Fprintf(out, "\nHow common an outlier is: of the %d reference decks, each set against the others, %d (%s%%) fall outside the others' full range on at least one of these %d budgets "+
		"(at most two a budget can, so the larger the reference, the rarer). Where that is common, a single figure far from the reference is a question and several pointing the same way are a finding.\n",
		n, outside, num(math.Round(10*pct(outside, n))/10), len(budgets))

	writeGains(out, d, lines)

	out.WriteString("\n### Concentration\n\n")
	lines(concentration)

	plays := map[int]int{}
	for _, m := range ms {
		for id := range m.copies {
			plays[id]++
		}
	}
	if fit > full {
		out.WriteString("\nThe staples this list lacks and its cards few of the reference play are counted, not named, to fit this result.\n")
	}
	writeStaples(out, d, here, ms, plays, fit == full)
	writeFew(out, d, here, ms, plays, fit == full)
	out.WriteString("\n")
	return &ref
}

// writeGains counts the list's blood and pool sources by kind, naming its cards.
func writeGains(out *strings.Builder, d *data, lines func([]figure)) {
	out.WriteString("\n### Blood and pool sources\n\n")
	if len(d.b.Economy) == 0 {
		out.WriteString("This data bundle holds no blood and pool sources: none until it is rebuilt.\n")
		return
	}
	out.WriteString("The copies of the library cards giving each kind of source, a card counted in each of its kinds, by the curated and revisable tags of `module-tags.json`. " +
		"Copies, not yield: most gains are conditional (a strike that lands, a vote that passes, the Edge held), and how much each gives is in its text.\n\n")
	var figures []figure
	for _, g := range d.b.Economy {
		figures = append(figures, figure{fmt.Sprintf("%s%s (%s)", strings.ToUpper(g.Kind[:1]), g.Kind[1:], g.What),
			func(m *measures) int { return m.gains[g.Kind] },
			func(m *measures) string {
				var names []string
				for _, t := range g.Cards {
					k := 0
					for _, id := range t.IDs {
						k += m.copies[id]
					}
					if k > 0 {
						names = append(names, fmt.Sprintf("%s ×%d", t.Name, k))
					}
				}
				if len(names) == 0 {
					return ""
				}
				return "; " + strings.Join(names, ", ")
			}})
	}
	lines(figures)
}

// writeStaples names the reference's staples this list lacks or plays below
// the reference's lower quartile, with the same count for each reference deck
// against the others.
func writeStaples(out *strings.Builder, d *data, here *measures, ms []*measures, plays map[int]int, named bool) {
	n := len(ms)
	out.WriteString("\n### Staples\n\n")
	copiesOf := map[int][]float64{} // sorted, every reference deck, zeros included
	for id, k := range plays {
		if float64(k) >= stapleShare*float64(n-1) {
			s := make([]float64, n)
			for i, m := range ms {
				s[i] = float64(m.copies[id])
			}
			slices.Sort(s)
			copiesOf[id] = s
		}
	}
	type miss struct {
		card        *bundle.Card
		have, plays int
		low, median float64
	}
	var misses []miss
	staples := 0
	for id, s := range copiesOf {
		if float64(plays[id]) < stapleShare*float64(n) {
			continue
		}
		staples++
		low, med := quantile(s, 0.25), quantile(s, 0.5)
		if have := here.copies[id]; have == 0 || float64(have) < low {
			misses = append(misses, miss{d.byID[id], have, plays[id], low, med})
		}
	}
	var own []float64
	for _, m := range ms {
		k := 0
		for id, s := range copiesOf {
			c := m.copies[id]
			p := plays[id]
			if c > 0 {
				p--
			}
			if float64(p) < stapleShare*float64(n-1) {
				continue
			}
			others := without(s, float64(c))
			if c == 0 || float64(c) < quantile(others, 0.25) {
				k++
			}
		}
		own = append(own, float64(k))
	}
	if staples == 0 {
		out.WriteString("No library card is in three quarters of the reference decks or more: the reference has no staples.\n")
		return
	}
	slices.Sort(own)
	fmt.Fprintf(out, "The reference's staples are the library cards in three quarters of its decks or more: %d. "+
		"This list lacks %d of them or plays fewer copies than the reference's lower quartile; each reference deck against the others: %s.\n",
		staples, len(misses), fiveNumbers(own))
	if !named {
		return
	}
	slices.SortFunc(misses, func(a, b miss) int {
		return cmp.Or(cmp.Compare(b.plays, a.plays), strings.Compare(a.card.Name, b.card.Name))
	})
	for _, m := range misses {
		fmt.Fprintf(out, "- %s: in %d of %d reference decks, median %s copies, lower quartile %s; here %d.\n",
			m.card.Name, m.plays, n, num(m.median), num(m.low), m.have)
	}
}

// writeFew names this list's library cards a tenth or fewer of the reference
// decks play, with the same count for each reference deck against the others.
func writeFew(out *strings.Builder, d *data, here *measures, ms []*measures, plays map[int]int, named bool) {
	n := len(ms)
	out.WriteString("\n### Cards few of the reference play\n\n")
	var few []int
	copies := 0
	for id, c := range here.copies {
		if float64(plays[id]) <= fewShare*float64(n) {
			few = append(few, id)
			copies += c
		}
	}
	var own []float64
	for _, m := range ms {
		k := 0
		for id := range m.copies {
			if float64(plays[id]-1) <= fewShare*float64(n-1) {
				k++
			}
		}
		own = append(own, float64(k))
	}
	slices.Sort(own)
	fmt.Fprintf(out, "This list holds %d library cards (%d copies) a tenth or fewer of the reference decks play; each reference deck against the others: %s.\n",
		len(few), copies, fiveNumbers(own))
	if !named {
		return
	}
	slices.SortFunc(few, func(a, b int) int {
		return cmp.Or(cmp.Compare(here.copies[b], here.copies[a]), strings.Compare(d.byID[a].Name, d.byID[b].Name))
	})
	var names []string
	for _, id := range few[:min(len(few), fewShown)] {
		names = append(names, fmt.Sprintf("%d× %s (in %d)", here.copies[id], d.byID[id].Name, plays[id]))
	}
	if k := len(few) - fewShown; k > 0 {
		names = append(names, fmt.Sprintf("and %d more", k))
	}
	if len(names) > 0 {
		fmt.Fprintf(out, "%s.\n", strings.Join(names, "; "))
	}
}

func pct(a, b int) float64 {
	return 100 * float64(a) / float64(b)
}

func values(ms []*measures, get func(*measures) float64) []float64 {
	s := make([]float64, len(ms))
	for i, m := range ms {
		s[i] = get(m)
	}
	slices.Sort(s)
	return s
}

func sortedWithout(ms []*measures, i int, get func(*measures) float64) []float64 {
	return without(values(ms, get), get(ms[i]))
}

// without is a sorted slice less one occurrence of v.
func without(s []float64, v float64) []float64 {
	i, _ := slices.BinarySearch(s, v)
	return slices.Concat(s[:i], s[i+1:])
}

// quantile interpolates linearly between the closest ranks of a sorted slice.
func quantile(s []float64, p float64) float64 {
	x := p * float64(len(s)-1)
	lo := int(x)
	if lo+1 >= len(s) {
		return s[len(s)-1]
	}
	return s[lo] + (x-float64(lo))*(s[lo+1]-s[lo])
}

// spread is a sorted slice's min, deciles, quartiles, median and max.
func spread(s []float64) string {
	var parts []string
	for _, p := range []float64{0, 0.1, 0.25, 0.5, 0.75, 0.9, 1} {
		parts = append(parts, fmt.Sprintf("%.0f", quantile(s, p)))
	}
	return strings.Join(parts, ", ")
}

func fiveNumbers(s []float64) string {
	var parts []string
	for _, p := range []float64{0.1, 0.25, 0.5, 0.75, 0.9} {
		parts = append(parts, fmt.Sprintf("%.0f", quantile(s, p)))
	}
	return strings.Join(parts, ", ") + " (10%, 25%, median, 75%, 90%)"
}

func rank(s []float64, v float64) string {
	above, below := 0, 0
	for _, x := range s {
		switch {
		case x > v:
			above++
		case x < v:
			below++
		}
	}
	return fmt.Sprintf("%d of the %d above it, %d below", above, len(s), below)
}

func num(f float64) string {
	if f == math.Trunc(f) {
		return fmt.Sprintf("%.0f", f)
	}
	return fmt.Sprintf("%.1f", f)
}
