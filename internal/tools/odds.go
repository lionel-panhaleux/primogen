package tools

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

const (
	handSize      = 7
	draftLibrary  = 90 // a draft's library, unless the caller gives its size
	maxRecipes    = 5
	maxGroups     = 6
	maxGroupNames = 12
)

// seenWithin are the counts of cards seen, the opening hand included, at which
// a recipe's odds are given past the opening hand.
var seenWithin = []int{14, 21, 28}

// binom is the binomial coefficient as a float: the largest a library needs,
// C(90, 45), is about 1e26, which a float64 holds to about 16 digits.
func binom(n, k int) float64 {
	if k < 0 || k > n {
		return 0
	}
	k = min(k, n-k)
	r := 1.0
	for i := 1; i <= k; i++ {
		r = r * float64(n-k+i) / float64(i)
	}
	return r
}

// anyInHand is the chance n cards drawn from a library of size hold at least
// one of its k copies.
func anyInHand(k, size, n int) float64 {
	return 1 - binom(size-k, n)/binom(size, n)
}

// pool is the copies of the cards sharing one set of groups: the cards are
// interchangeable, so the draw is hypergeometric over pools.
type pool struct {
	groups uint // a bit per group the cards fill
	copies int
}

// recipeOdds is the exact chance that n cards drawn from a library of size
// hold a distinct copy for each group, a copy filling one group only. A card
// named in several groups (or a group repeated) needs as many copies.
func recipeOdds(pools []pool, nGroups, size, n int) float64 {
	caps := make([]int, len(pools))
	other := size
	for i, p := range pools {
		caps[i] = poolCap(p)
		other -= p.copies
	}
	nStates := states(pools)
	// dist[t][s]: the ways to draw t cards from the pools so far, s their
	// copies drawn (capped), in mixed radix
	dist := make([][]float64, n+1)
	for t := range dist {
		dist[t] = make([]float64, nStates)
	}
	dist[0][0] = 1
	radix := 1
	for i, p := range pools {
		next := make([][]float64, n+1)
		for t := range next {
			next[t] = make([]float64, nStates)
		}
		for t, row := range dist {
			for s, w := range row {
				if w == 0 {
					continue
				}
				for x := 0; x <= p.copies && t+x <= n; x++ {
					next[t+x][s+radix*min(x, caps[i])] += w * binom(p.copies, x)
				}
			}
		}
		dist, radix = next, radix*(caps[i]+1)
	}
	total := 0.0
	drawn := make([]int, len(pools))
	for t, row := range dist {
		rest := binom(other, n-t)
		if rest == 0 {
			continue
		}
		for s, w := range row {
			if w == 0 {
				continue
			}
			for i, v := 0, s; i < len(pools); i++ {
				drawn[i], v = v%(caps[i]+1), v/(caps[i]+1)
			}
			if fills(pools, drawn, nGroups) {
				total += w * rest
			}
		}
	}
	return total / binom(size, n)
}

// maxStates bounds the draws recipeOdds tells apart.
const maxStates = 1 << 16

// poolCap is the copies of a pool worth telling apart: more than the groups
// it fills never help.
func poolCap(p pool) int { return min(p.copies, popcount(p.groups)) }

func states(pools []pool) int {
	n := 1
	for _, p := range pools {
		if n *= poolCap(p) + 1; n > maxStates {
			return n
		}
	}
	return n
}

// fills is Hall's condition: every set of groups has at least as many copies
// drawn among the pools able to fill one of them.
func fills(pools []pool, drawn []int, nGroups int) bool {
	for set := uint(1); set < 1<<nGroups; set++ {
		have := 0
		for i, p := range pools {
			if p.groups&set != 0 {
				have += drawn[i]
			}
		}
		if have < popcount(set) {
			return false
		}
	}
	return true
}

func popcount(x uint) int {
	n := 0
	for ; x != 0; x &= x - 1 {
		n++
	}
	return n
}

// recipe is a caller's recipe, its names resolved: each group's cards.
type recipe [][]*bundle.Card

// readRecipes resolves the recipes' names as vtes_cards does; it answers the
// errors to fix.
func (d *data) readRecipes(in [][][]string, dk *deck) ([]recipe, []string) {
	var out []recipe
	var errs []string
	if len(in) > maxRecipes {
		return nil, []string{fmt.Sprintf("%d recipes: give at most %d", len(in), maxRecipes)}
	}
	for _, groups := range in {
		if len(groups) == 0 || len(groups) > maxGroups {
			errs = append(errs, fmt.Sprintf("a recipe of %d groups: give 1 to %d groups, each a list of card names one of which is needed", len(groups), maxGroups))
			continue
		}
		var r recipe
		for _, names := range groups {
			if len(names) == 0 || len(names) > maxGroupNames {
				errs = append(errs, fmt.Sprintf("a group of %d names: give 1 to %d card names, one of which is needed", len(names), maxGroupNames))
				continue
			}
			var group []*bundle.Card
			for _, name := range names {
				cards, note := d.ix.resolve(name, "vtes_deck")
				if cards == nil {
					errs = append(errs, note)
					continue
				}
				if note != "" {
					dk.notes = append(dk.notes, note)
				}
				c := cards[0]
				switch {
				case c.IsCrypt():
					errs = append(errs, fmt.Sprintf("%s is a crypt card: a recipe holds library cards, the hand's", c.Name))
				case !slices.Contains(group, c):
					group = append(group, c)
				}
			}
			r = append(r, group)
		}
		out = append(out, r)
	}
	return out, errs
}

// writeOdds gives each recipe's odds, over a library of size cards.
func writeOdds(out *strings.Builder, recipes []recipe, library, crypt []deckEntry, size, nCrypt int, sizeNote string) {
	out.WriteString("## Odds of a hand\n\n")
	fmt.Fprintf(out, "Exact odds over a library of %d cards.%s Only the opening hand is a random draw of fixed size: "+
		"past it, how many cards a deck has seen depends on what it plays, discards and searches, so %d, %d and %d are cards seen, "+
		"the opening %d included, not turns. A copy fills one group only.\n\n",
		size, sizeNote, seenWithin[0], seenWithin[1], seenWithin[2], handSize)
	held := map[*bundle.Card]int{}
	for _, e := range library {
		held[e.card] = e.count
	}
	for i, r := range recipes {
		var parts, absent, unchecked []string
		bits := map[*bundle.Card]uint{}
		for g, group := range r {
			var names []string
			for _, c := range group {
				bits[c] |= 1 << g
				if held[c] > 0 {
					names = append(names, copies(deckEntry{card: c, count: held[c]}))
				} else if !slices.Contains(absent, c.Name) {
					absent = append(absent, c.Name)
				}
				if m := reTextRequires.FindString(c.Text); m != "" && held[c] > 0 && !slices.Contains(unchecked, c.Name) {
					unchecked = append(unchecked, c.Name)
				}
			}
			parts = append(parts, "("+cmp.Or(strings.Join(names, ", "), "none in this deck")+")")
		}
		byGroups := map[uint]int{}
		for c, b := range bits {
			byGroups[b] += held[c]
		}
		var pools []pool
		for b, k := range byGroups {
			if k > 0 {
				pools = append(pools, pool{b, k})
			}
		}
		slices.SortFunc(pools, func(a, b pool) int { return cmp.Compare(a.groups, b.groups) })
		fmt.Fprintf(out, "%d. One of each: %s.", i+1, strings.Join(parts, " + "))
		if len(absent) > 0 {
			fmt.Fprintf(out, " Not in this deck: %s.", strings.Join(absent, ", "))
		}
		if states(pools) > maxStates {
			out.WriteString("\n   Too many cards shared between its groups to compute: give groups sharing fewer cards.\n")
			continue
		}
		fmt.Fprintf(out, "\n   In an opening hand: %.1f%%; seen within", 100*recipeOdds(pools, len(r), size, handSize))
		for j, n := range seenWithin {
			if j > 0 {
				out.WriteString(",")
			}
			fmt.Fprintf(out, " %d cards %.1f%%", n, 100*recipeOdds(pools, len(r), size, n))
		}
		out.WriteString(".\n")
		out.WriteString("   " + ableToPlay(r, held, crypt, nCrypt))
		if len(unchecked) > 0 {
			fmt.Fprintf(out, " A requirement only the text states, unchecked here: %s.", strings.Join(unchecked, ", "))
		}
		out.WriteString("\n")
	}
	out.WriteString("\n")
}

// ableToPlay names the crypt cards having what a card the deck holds of each
// group requires, with their levels in those cards' disciplines.
func ableToPlay(r recipe, held map[*bundle.Card]int, crypt []deckEntry, nCrypt int) string {
	if nCrypt == 0 {
		return "No crypt given: who can play every piece is unchecked."
	}
	var disciplines []string
	for _, group := range r {
		for _, c := range group {
			if held[c] > 0 {
				for _, d := range c.Disciplines {
					if !slices.Contains(disciplines, strings.ToLower(d)) {
						disciplines = append(disciplines, strings.ToLower(d))
					}
				}
			}
		}
	}
	var able []string
	n := 0
	for _, e := range crypt {
		if slices.ContainsFunc(r, func(group []*bundle.Card) bool {
			return !slices.ContainsFunc(group, func(c *bundle.Card) bool { _, ok := meets(c, e.card); return held[c] > 0 && ok })
		}) {
			continue
		}
		name := copies(e)
		var levels []string
		for _, d := range cryptDisciplines(e.card) {
			if slices.Contains(disciplines, strings.ToLower(d)) {
				levels = append(levels, d)
			}
		}
		if len(levels) > 0 {
			name += " [" + strings.Join(levels, " ") + "]"
		}
		able = append(able, name)
		n += e.count
	}
	head := "Crypt cards having what a card of each group held here requires (the requirement, not the level an effect needs: their levels shown, the texts above say which level gives what)"
	if n == 0 {
		return head + ": none."
	}
	return fmt.Sprintf("%s: %d of %d: %s.", head, n, nCrypt, strings.Join(able, "; "))
}
