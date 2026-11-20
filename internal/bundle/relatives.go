package bundle

import (
	"cmp"
	"encoding/json"
	"fmt"
	"maps"
	"math"
	"slices"
	"strings"
)

// Classification is the owner's archetype labels of the TWDA decks, from
// knowledge/classification.json.
type Classification struct {
	Generated   string
	Revised     string // the last day the owner ruled on the build's proposals, if after Generated
	Since       string // the first date of the labelled window
	CryptWeight float64
	MinDF       int // the decks a card must be in to count
	Archetypes  []Archetype
	Noise       []string // decks of the window the owner put in no archetype
}

type Archetype struct {
	Ref       string
	Name      string
	VariantOf string // the ref of the archetype this one is a variant of
	Decks     []string
	// Representative are the decks the owner picked to show the archetype.
	Representative []string
}

// Pass dates the labels.
func (c *Classification) Pass() string {
	if c.Revised == "" {
		return c.Generated
	}
	return c.Generated + ", revised on " + c.Revised + " by the same maintainer from primogen's build's proposals"
}

func readClassification(raw []byte) (Classification, error) {
	var f struct {
		Generated string
		Revised   string
		Params    struct {
			Since       string  `json:"since"`
			CryptWeight float64 `json:"crypt_weight"`
			MinDF       int     `json:"min_df"`
		}
		Groups []struct {
			Ref       string
			Name      string
			VariantOf *string `json:"variant_of"`
			Decks     []struct {
				ID             string
				Representative bool
			}
		}
		Noise []struct{ ID string }
	}
	if err := json.Unmarshal(raw, &f); err != nil {
		return Classification{}, fmt.Errorf("classification.json: %w", err)
	}
	c := Classification{Generated: f.Generated, Revised: f.Revised, Since: f.Params.Since, CryptWeight: f.Params.CryptWeight, MinDF: f.Params.MinDF}
	if c.Since == "" || c.CryptWeight <= 0 || c.CryptWeight >= 1 || c.MinDF < 1 || len(f.Groups) == 0 {
		return Classification{}, fmt.Errorf("classification.json: no groups, or params without since, crypt_weight in (0, 1) and min_df")
	}
	for _, g := range f.Groups {
		a := Archetype{Ref: g.Ref, Name: g.Name}
		if g.VariantOf != nil {
			a.VariantOf = *g.VariantOf
		}
		for _, d := range g.Decks {
			a.Decks = append(a.Decks, d.ID)
			if d.Representative {
				a.Representative = append(a.Representative, d.ID)
			}
		}
		c.Archetypes = append(c.Archetypes, a)
	}
	for _, a := range c.Archetypes {
		// a variant of a variant, or of no archetype, has no root
		if a.VariantOf != "" && !slices.ContainsFunc(c.Archetypes, func(p Archetype) bool { return p.Ref == a.VariantOf && p.VariantOf == "" }) {
			return Classification{}, fmt.Errorf("classification.json: %s is a variant of %q, no archetype that is itself no variant", a.Ref, a.VariantOf)
		}
	}
	for _, d := range f.Noise {
		c.Noise = append(c.Noise, d.ID)
	}
	return c, nil
}

// Relatives places decks among the TWDA decks of the labelled window and the
// owner's archetypes.
type Relatives struct {
	Class *Classification
	// Threshold is the similarity 95 in 100 of the labelled decks reach to
	// their own archetype, each left out of it.
	Threshold float64
	// GroupRate and RootRate are the labelled decks that classify back to their
	// own archetype, each left out of it, and to its root (variants merged).
	GroupRate, RootRate float64

	column  map[string]int
	idf     []float64
	crypt   []bool
	byID    map[int]*Card
	Decks   []*Deck // the window's
	vecs    []vector
	sums    [][]float64 // per archetype, of its decks' vectors
	sizes   []int       // per archetype, its decks the TWDA holds
	norms   []float64   // of sums
	groupOf map[string]int
	noise   map[string]bool
}

// vector is a deck's normalized weights, sparse.
type vector struct {
	cols []int
	vals []float64
}

// Feature is a card's column key: a vampire's printed name across its groups
// (re-printed in another group, it plays the same deck), advanced apart.
func Feature(c *Card) string {
	if c.IsCrypt() {
		if c.Advanced {
			return "c:" + c.PrintedName + " (ADV)"
		}
		return "c:" + c.PrintedName
	}
	return "l:" + c.Name
}

// NewRelatives is nil for a bundle without archetype labels, built before it carried them.
func NewRelatives(b *Bundle) *Relatives {
	c := &b.Classification
	if len(c.Archetypes) == 0 {
		return nil
	}
	r := &Relatives{Class: c, column: map[string]int{}, byID: map[int]*Card{}, groupOf: map[string]int{}, noise: map[string]bool{}}
	for i := range b.Cards {
		r.byID[b.Cards[i].ID] = &b.Cards[i]
	}
	for i := range b.Decks {
		if b.Decks[i].Date >= c.Since {
			r.Decks = append(r.Decks, &b.Decks[i])
		}
	}
	df := map[string]int{}
	for _, d := range r.Decks {
		for k := range r.counts(d.Cards) {
			df[k]++
		}
	}
	keys := slices.Sorted(maps.Keys(df))
	n := float64(len(r.Decks))
	for _, k := range keys {
		if df[k] >= c.MinDF {
			r.column[k] = len(r.idf)
			r.idf = append(r.idf, math.Log((1+n)/(1+float64(df[k])))+1)
			r.crypt = append(r.crypt, strings.HasPrefix(k, "c:"))
		}
	}
	for _, d := range r.Decks {
		r.vecs = append(r.vecs, r.Vector(d.Cards))
	}
	for i, a := range c.Archetypes {
		for _, id := range a.Decks {
			r.groupOf[id] = i
		}
	}
	for _, id := range c.Noise {
		r.noise[id] = true
	}
	r.sums, r.sizes = make([][]float64, len(c.Archetypes)), make([]int, len(c.Archetypes))
	for i := range r.sums {
		r.sums[i] = make([]float64, len(r.idf))
	}
	for i, d := range r.Decks {
		if g, ok := r.groupOf[d.ID]; ok {
			r.sizes[g]++
			for j, col := range r.vecs[i].cols {
				r.sums[g][col] += r.vecs[i].vals[j]
			}
		}
	}
	r.norms = make([]float64, len(r.sums))
	for i, s := range r.sums {
		for _, v := range s {
			r.norms[i] += v * v
		}
		r.norms[i] = math.Sqrt(r.norms[i])
	}
	r.validate()
	return r
}

func (r *Relatives) counts(cards []DeckCard) map[string]int {
	out := map[string]int{}
	for _, dc := range cards {
		if c := r.byID[dc.ID]; c != nil && dc.Count > 0 {
			out[Feature(c)] += dc.Count
		}
	}
	return out
}

// Vector weighs a list; a card no two decks of the window play weighs nothing.
func (r *Relatives) Vector(cards []DeckCard) vector {
	var v vector
	counts := r.counts(cards)
	for _, k := range slices.Sorted(maps.Keys(counts)) {
		if col, ok := r.column[k]; ok {
			v.cols = append(v.cols, col)
			v.vals = append(v.vals, (1+math.Log(float64(counts[k])))*r.idf[col])
		}
	}
	var norms [2]float64 // crypt, library
	for j, col := range v.cols {
		norms[b2i(!r.crypt[col])] += v.vals[j] * v.vals[j]
	}
	weights := [2]float64{r.Class.CryptWeight, 1 - r.Class.CryptWeight}
	total := 0.0
	for j, col := range v.cols {
		k := b2i(!r.crypt[col])
		v.vals[j] *= weights[k] / math.Sqrt(norms[k])
		total += v.vals[j] * v.vals[j]
	}
	for j := range v.vals {
		v.vals[j] /= math.Sqrt(total)
	}
	return v
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (v vector) Empty() bool { return len(v.cols) == 0 }

func (v vector) dot(w vector) float64 {
	s, i, j := 0.0, 0, 0
	for i < len(v.cols) && j < len(w.cols) {
		switch {
		case v.cols[i] < w.cols[j]:
			i++
		case v.cols[i] > w.cols[j]:
			j++
		default:
			s += v.vals[i] * w.vals[j]
			i, j = i+1, j+1
		}
	}
	return s
}

func (v vector) dotDense(d []float64) float64 {
	s := 0.0
	for j, col := range v.cols {
		s += v.vals[j] * d[col]
	}
	return s
}

// Match is an archetype or a deck, and its similarity to the deck placed.
type Match struct {
	Index      int
	Similarity float64
}

// similarities to every archetype, a labelled deck left out of its own.
func (r *Relatives) similarities(v vector, id string) []float64 {
	out := make([]float64, len(r.sums))
	own, labelled := r.groupOf[id]
	for g, s := range r.sums {
		dot := v.dotDense(s)
		norm := r.norms[g]
		if labelled && g == own {
			// the deck's own vector, of norm 1, taken out of the sum
			dot, norm = dot-1, math.Sqrt(max(norm*norm-2*dot+1, 0))
		}
		if norm > 0 {
			out[g] = dot / norm
		}
	}
	return out
}

// Archetypes are the n nearest archetypes; id names a TWDA deck, left out of
// its own archetype, or is empty.
func (r *Relatives) Archetypes(v vector, id string, n int) []Match {
	return top(r.similarities(v, id), n)
}

// Nearest are the n nearest decks of the window, the deck id left out.
func (r *Relatives) Nearest(v vector, id string, n int) []Match {
	sims := make([]float64, len(r.vecs))
	for i, w := range r.vecs {
		if r.Decks[i].ID == id {
			sims[i] = math.Inf(-1)
		} else {
			sims[i] = v.dot(w)
		}
	}
	return top(sims, n)
}

// Around are the n decks of the window nearest the sum of the given decks'
// vectors.
func (r *Relatives) Around(ids map[string]bool, n int) []*Deck {
	sum := make([]float64, len(r.idf))
	for i, d := range r.Decks {
		if ids[d.ID] {
			for j, col := range r.vecs[i].cols {
				sum[col] += r.vecs[i].vals[j]
			}
		}
	}
	sims := make([]float64, len(r.vecs))
	for i, v := range r.vecs {
		sims[i] = v.dotDense(sum)
	}
	var out []*Deck
	for _, m := range top(sims, n) {
		out = append(out, r.Decks[m.Index])
	}
	return out
}

func top(sims []float64, n int) []Match {
	out := make([]Match, len(sims))
	for i, s := range sims {
		out[i] = Match{i, s}
	}
	slices.SortStableFunc(out, func(a, b Match) int { return cmp.Compare(b.Similarity, a.Similarity) })
	return out[:min(n, len(out))]
}

// Size is how many decks of archetype g the TWDA holds.
func (r *Relatives) Size(g int) int { return r.sizes[g] }

// Label is the owner's archetype of a TWDA deck, or nil.
func (r *Relatives) Label(id string) *Archetype {
	if g, ok := r.groupOf[id]; ok {
		return &r.Class.Archetypes[g]
	}
	return nil
}

// Root is the archetype a variant is of, or the archetype itself.
func (r *Relatives) Root(a *Archetype) *Archetype {
	if a.VariantOf == "" {
		return a
	}
	return &r.Class.Archetypes[slices.IndexFunc(r.Class.Archetypes, func(p Archetype) bool { return p.Ref == a.VariantOf })]
}

// validate classifies each labelled deck of an archetype of two or more,
// left out of it, and sets the threshold from their own similarities.
func (r *Relatives) validate() {
	size := make([]int, len(r.sums))
	for i, d := range r.Decks {
		if g, ok := r.groupOf[d.ID]; ok && !r.vecs[i].Empty() {
			size[g]++
		}
	}
	var own []float64
	group, root := 0, 0
	for i, d := range r.Decks {
		g, ok := r.groupOf[d.ID]
		if !ok || size[g] < 2 || r.vecs[i].Empty() {
			continue
		}
		sims := r.similarities(r.vecs[i], d.ID)
		best := top(sims, 1)[0].Index
		own = append(own, sims[g])
		if best == g {
			group++
		}
		if r.Root(&r.Class.Archetypes[best]) == r.Root(&r.Class.Archetypes[g]) {
			root++
		}
	}
	if len(own) == 0 {
		return
	}
	slices.Sort(own)
	r.Threshold = own[len(own)*5/100]
	r.GroupRate, r.RootRate = float64(group)/float64(len(own)), float64(root)/float64(len(own))
}

// Placement is the archetype of the window's deck i: its label, else its
// nearest archetype at the threshold or more. g is -1 for none: the owner
// put it in none (labelled), or it is near none.
func (r *Relatives) Placement(i int) (g int, labelled bool) {
	id := r.Decks[i].ID
	if g, ok := r.groupOf[id]; ok {
		return g, true
	}
	if r.noise[id] {
		return -1, true
	}
	if r.vecs[i].Empty() {
		return -1, false
	}
	if best := r.Archetypes(r.vecs[i], id, 1)[0]; best.Similarity >= r.Threshold {
		return best.Index, false
	}
	return -1, false
}

// Proposal is a group of the window's decks in no archetype of which each,
// left out, reaches the threshold to the rest: it would be placed back in the
// group were the group an archetype.
type Proposal struct {
	Decks             []*Deck
	Vampires, Library []string // in half its decks or more, most particular to them first, at most 5 and 6
	Nearest           Match    // the archetype nearest it
}

const (
	proposalMin      = 3
	proposalVampires = 5
	proposalLibrary  = 6
)

// Proposals groups the window's decks in no archetype, the owner's noise and
// those near none; a group is proposed when one of its decks is not the
// owner's noise. Alone are the decks near none in no proposal.
func (r *Relatives) Proposals() (props []Proposal, alone []string) {
	var rows []int
	for i := range r.Decks {
		if g, _ := r.Placement(i); g < 0 {
			rows = append(rows, i)
		}
	}
	type group struct {
		members []int
		sum     []float64
	}
	groups := make([]*group, len(rows))
	score := make([][]float64, len(rows))
	for a, i := range rows {
		groups[a] = &group{members: []int{i}, sum: make([]float64, len(r.idf))}
		for j, col := range r.vecs[i].cols {
			groups[a].sum[col] = r.vecs[i].vals[j]
		}
		score[a] = make([]float64, len(rows))
		for b, j := range rows[:a] {
			score[a][b] = r.vecs[i].dot(r.vecs[j])
			score[b][a] = score[a][b]
		}
	}
	// the least similarity of a member, left out, to the rest of a and b merged
	merged := func(a, b *group) float64 {
		sum := slices.Clone(a.sum)
		for k, v := range b.sum {
			sum[k] += v
		}
		norm := 0.0
		for _, v := range sum {
			norm += v * v
		}
		least := math.Inf(1)
		for _, i := range append(slices.Clone(a.members), b.members...) {
			dot := r.vecs[i].dotDense(sum)
			least = min(least, (dot-1)/math.Sqrt(max(norm-2*dot+1, 1e-12)))
		}
		return least
	}
	for {
		a, b, best := -1, -1, r.Threshold
		for x := range groups {
			for y := range x {
				if groups[x] != nil && groups[y] != nil && score[x][y] >= best {
					a, b, best = x, y, score[x][y]
				}
			}
		}
		if a < 0 {
			break
		}
		for k, v := range groups[b].sum {
			groups[a].sum[k] += v
		}
		groups[a].members, groups[b] = append(groups[a].members, groups[b].members...), nil
		for x, g := range groups {
			if g != nil && x != a {
				score[a][x] = merged(groups[a], g)
				score[x][a] = score[a][x]
			}
		}
	}
	df := map[string]int{}
	for _, d := range r.Decks {
		for k := range r.counts(d.Cards) {
			df[k]++
		}
	}
	proposed := map[int]bool{}
	for _, g := range groups {
		if g == nil || len(g.members) < proposalMin || !slices.ContainsFunc(g.members, func(i int) bool { return !r.noise[r.Decks[i].ID] }) {
			continue
		}
		var p Proposal
		n := map[string]int{}
		for _, i := range g.members {
			proposed[i] = true
			p.Decks = append(p.Decks, r.Decks[i])
			for k := range r.counts(r.Decks[i].Cards) {
				n[k]++
			}
		}
		lift := func(k string) float64 { return float64(n[k]) / float64(df[k]) }
		var cards []string
		for k, c := range n {
			if 2*c >= len(g.members) {
				cards = append(cards, k)
			}
		}
		slices.SortFunc(cards, func(x, y string) int { return cmp.Or(cmp.Compare(lift(y), lift(x)), cmp.Compare(x, y)) })
		for _, k := range cards {
			if k[0] == 'c' && len(p.Vampires) < proposalVampires {
				p.Vampires = append(p.Vampires, k[2:])
			} else if k[0] == 'l' && len(p.Library) < proposalLibrary {
				p.Library = append(p.Library, k[2:])
			}
		}
		var v vector
		norm := 0.0
		for col, x := range g.sum {
			if x != 0 {
				v.cols, v.vals, norm = append(v.cols, col), append(v.vals, x), norm+x*x
			}
		}
		for j := range v.vals {
			v.vals[j] /= math.Sqrt(norm)
		}
		p.Nearest = r.Archetypes(v, "", 1)[0]
		slices.SortFunc(p.Decks, func(x, y *Deck) int { return cmp.Compare(y.Date, x.Date) })
		props = append(props, p)
	}
	slices.SortStableFunc(props, func(x, y Proposal) int { return cmp.Compare(len(y.Decks), len(x.Decks)) })
	for i, d := range r.Decks {
		if g, labelled := r.Placement(i); g < 0 && !labelled && !proposed[i] {
			best := r.Archetypes(r.vecs[i], d.ID, 1)[0]
			alone = append(alone, fmt.Sprintf("%s %q (%s, closest %s at %.2f)", d.ID, d.Name, d.Date, r.Class.Archetypes[best.Index].Name, best.Similarity))
		}
	}
	return props, alone
}
