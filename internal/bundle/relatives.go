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

func readClassification(raw []byte) (Classification, error) {
	var f struct {
		Generated string
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
	c := Classification{Generated: f.Generated, Since: f.Params.Since, CryptWeight: f.Params.CryptWeight, MinDF: f.Params.MinDF}
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

// feature is a card's column key: a vampire's printed name across its groups
// (re-printed in another group, it plays the same deck), advanced apart.
func feature(c *Card) string {
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
			out[feature(c)] += dc.Count
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

// Unmatched are the decks of the window the owner has not labelled, archetype
// or noise, that reach no archetype's threshold.
func (r *Relatives) Unmatched() []string {
	var out []string
	for i, d := range r.Decks {
		if _, ok := r.groupOf[d.ID]; ok || r.noise[d.ID] {
			continue
		}
		best := r.Archetypes(r.vecs[i], d.ID, 1)[0]
		if best.Similarity < r.Threshold {
			out = append(out, fmt.Sprintf("%s %q (%s, closest %s at %.2f)", d.ID, d.Name, d.Date, r.Class.Archetypes[best.Index].Name, best.Similarity))
		}
	}
	return out
}
