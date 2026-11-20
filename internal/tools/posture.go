package tools

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

// writePosture gives each combat defense's share of the archetype labels'
// window and the deck's answers to it.
func writePosture(out *strings.Builder, b *bundle.Bundle, library, crypt []deckEntry, nCrypt int) {
	if len(b.Defenses) == 0 {
		return
	}
	window := 0
	carried := make([]int, len(b.Defenses))
	byCard := map[int]int{} // card → the window's decks playing it
	for _, d := range b.Decks {
		if d.Date < b.Classification.Since {
			continue
		}
		window++
		held := map[int]bool{}
		for _, c := range d.Cards {
			held[c.ID] = true
			byCard[c.ID]++
		}
		for i, def := range b.Defenses {
			if slices.ContainsFunc(def.Cards, func(t bundle.Tag) bool { return slices.ContainsFunc(t.IDs, func(id int) bool { return held[id] }) }) {
				carried[i]++
			}
		}
	}
	fmt.Fprintf(out, "## Combat posture against the field\n\n"+
		"The combat defenses of primogen's catalog (knowledge/guide/modules.md, after codex-of-the-damned's combat-modules article, curated and revisable), "+
		"the most carried first: how many of the %d TWDA decks since %s carry a card giving it (what winners play, not strength), the cards most often, "+
		"and this deck's cards answering it, how, and what an answer requires (capitals: superior) with the crypt cards having it.\n\n",
		window, b.Classification.Since)
	held := map[int]deckEntry{}
	for _, e := range library {
		held[e.card.ID] = e
	}
	order := make([]int, len(b.Defenses))
	for i := range order {
		order[i] = i
	}
	slices.SortStableFunc(order, func(x, y int) int { return cmp.Compare(carried[y], carried[x]) })
	for _, i := range order {
		def := b.Defenses[i]
		fmt.Fprintf(out, "- **%s**: %d of %d decks (%.1f%%)", def.Name, carried[i], window, 100*float64(carried[i])/float64(window))
		type often struct {
			name  string
			decks int
		}
		var top []often
		for _, t := range def.Cards {
			n := 0
			for _, id := range t.IDs {
				n += byCard[id]
			}
			if n > 0 {
				top = append(top, often{t.Name, n})
			}
		}
		slices.SortStableFunc(top, func(x, y often) int { return cmp.Compare(y.decks, x.decks) })
		var names []string
		for _, o := range top[:min(3, len(top))] {
			names = append(names, fmt.Sprintf("%s (%d)", o.name, o.decks))
		}
		if len(names) > 0 {
			fmt.Fprintf(out, ", most often %s", strings.Join(names, ", "))
		}
		var answers []string
		for _, as := range def.Answers {
			for _, a := range as.Cards {
				for _, id := range a.IDs {
					if e, ok := held[id]; ok {
						answers = append(answers, copies(e)+", "+as.How+answerNeeds(e.card, a.Level, crypt, nCrypt))
					}
				}
			}
		}
		if len(answers) == 0 {
			out.WriteString(". Not answered here.\n")
		} else {
			fmt.Fprintf(out, ". Answered here by %s.\n", strings.Join(answers, "; "))
		}
	}
	out.WriteString("\n")
}

// answerNeeds says what an answer requires (its level for this effect, its
// clan and path) and the crypt cards having it; a requirement only its text
// states is quoted, unchecked.
func answerNeeds(c *bundle.Card, level [][]string, crypt []deckEntry, nCrypt int) string {
	var needs []string
	if len(level) > 0 {
		alts := make([]string, len(level))
		for i, set := range level {
			alts[i] = "[" + strings.Join(set, "][") + "]"
		}
		needs = append(needs, strings.Join(alts, " or "))
	}
	if len(c.Clans) > 0 {
		needs = append(needs, "clan "+strings.Join(c.Clans, " or "))
	}
	if c.Path != "" {
		needs = append(needs, "path "+c.Path)
	}
	var line string
	switch {
	case len(needs) == 0:
	case nCrypt == 0:
		line = ", needs " + strings.Join(needs, " and ") + ", unchecked without a crypt"
	default:
		n, elected := 0, 0
		for _, e := range crypt {
			has := cryptDisciplines(e.card)
			if clanAndPath(c, e.card) && (len(level) == 0 || slices.ContainsFunc(level, func(set []string) bool {
				return !slices.ContainsFunc(set, func(d string) bool {
					return !slices.ContainsFunc(has, func(h string) bool {
						return strings.EqualFold(h, d) && (d != strings.ToUpper(d) || h == strings.ToUpper(h))
					})
				})
			})) {
				n += e.count
				if electsPath(c, e.card) {
					elected += e.count
				}
			}
		}
		switch {
		case n == 0:
			line = ", needs " + strings.Join(needs, " and ") + ": no crypt card here has it"
		case elected > 0:
			line = fmt.Sprintf(", needs %s: %d of %d crypt cards, %d of them electing the Path: %s", strings.Join(needs, " and "), n, nCrypt, elected, electionRule)
		default:
			line = fmt.Sprintf(", needs %s: %d of %d crypt cards", strings.Join(needs, " and "), n, nCrypt)
		}
	}
	if m := reTextRequires.FindString(c.Text); m != "" {
		line += fmt.Sprintf(" (its text also requires, unchecked: %q)", m)
	}
	return line
}
