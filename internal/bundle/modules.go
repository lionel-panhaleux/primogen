package bundle

import (
	"cmp"
	"encoding/json"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"
)

// Module is a module of knowledge/guide/modules.md with its cards, from
// knowledge/guide/module-tags.json.
type Module struct {
	Section string
	Name    string
	Core    []Tag
	Support []Tag
	// Pieces is how many of its cards, one at least core, a deck holds for it to show.
	Pieces int
}

// Tag is a card name of the tags and the cards it stands for.
type Tag struct {
	Name string
	IDs  []int
}

// Defense is a combat defense of modules.md's "Combat defenses", with the
// cards giving it and those answering it.
type Defense struct {
	Name    string
	Cards   []Tag
	Answers []Answers
}

// Answers are the cards answering a defense one way.
type Answers struct {
	How   string
	Cards []Answer
}

// Answer is a card answering a defense, at the level it needs:
// alternatives, each a set of disciplines all needed (capitals: superior).
type Answer struct {
	Tag
	Level [][]string
}

var (
	reCatalogEntry = regexp.MustCompile(`(?ms)^## ([^\n]+)$|^- \*\*(.+?)\*\*`)
	reGroup        = regexp.MustCompile(` \(G[^)]*\)$`)
	reLevelSymbol  = regexp.MustCompile(`\[(\w+)\]`)
)

// defensesSection is the section of modules.md whose entries are combat
// defenses, not modules.
const defensesSection = "Combat defenses"

// DensityClasses are theory.md's classes of a module's density, which
// vtes_deck applies.
const DensityClasses = "spam above 25%, strong 14 to 25%, standard 7 to 14%, tactical below 7%"

// readModules resolves the tags' card names, a vampire's name without its
// group standing for each of its groups; a name no card has is left out.
func readModules(raw, catalog, theory []byte, cards []Card, warn func(string, ...any)) ([]Module, []Defense, error) {
	var f struct {
		Modules []struct {
			Section, Module string
			Core, Support   []string
			Pieces          int
		}
		Untagged map[string]string // a catalog module with no card of its own, and why
		Defenses []struct {
			Defense string
			Cards   []string
			Answers []struct {
				How   string
				Cards []string
			}
		}
	}
	if err := json.Unmarshal(raw, &f); err != nil {
		return nil, nil, fmt.Errorf("guide/module-tags.json: %w", err)
	}
	if !strings.Contains(strings.Join(strings.Fields(string(theory)), " "), DensityClasses) {
		warn("guide/theory.md no longer reads %q, the density classes vtes_deck applies: change them in the code too", DensityClasses)
	}
	named := map[string]string{} // module → its section
	defenses := map[string]bool{}
	section := ""
	for _, m := range reCatalogEntry.FindAllSubmatch(catalog, -1) {
		name := strings.Join(strings.Fields(string(m[2])), " ")
		switch {
		case m[1] != nil:
			section = string(m[1])
		case section == defensesSection:
			defenses[name] = true
		default:
			named[name] = section
		}
	}
	ids := map[string][]int{}
	disciplines := map[string][]string{}
	for _, c := range cards {
		disciplines[c.Name] = c.Disciplines
		ids[c.Name] = append(ids[c.Name], c.ID)
		if base := reGroup.ReplaceAllString(c.Name, ""); base != c.Name {
			ids[base] = append(ids[base], c.ID)
		}
	}
	resolve := func(names []string, of string) []Tag {
		var found []Tag
		for _, n := range names {
			if len(ids[n]) == 0 {
				warn("guide/module-tags.json: no card is named %q (%s), left out", n, of)
				continue
			}
			found = append(found, Tag{n, ids[n]})
		}
		return found
	}
	var out []Module
	for _, m := range f.Modules {
		if s, ok := named[m.Module]; !ok || s != m.Section {
			warn("guide/module-tags.json: %q of %q is no module of that section of guide/modules.md", m.Module, m.Section)
		}
		delete(named, m.Module)
		mod := Module{m.Section, m.Module, resolve(m.Core, m.Module), resolve(m.Support, m.Module), max(m.Pieces, 1)}
		if len(mod.Core)+len(mod.Support) < mod.Pieces {
			warn("guide/module-tags.json: %q needs %d pieces and tags fewer cards: it never shows", m.Module, mod.Pieces)
		}
		out = append(out, mod)
	}
	for n := range f.Untagged {
		if _, ok := named[n]; !ok {
			warn("guide/module-tags.json: %q, under untagged, is no untagged module of guide/modules.md", n)
		}
	}
	for _, n := range slices.Sorted(maps.Keys(named)) {
		if f.Untagged[n] == "" {
			warn("guide/modules.md: the module %q has no tags in guide/module-tags.json, nor a reason under untagged", n)
		}
	}
	var defs []Defense
	for _, d := range f.Defenses {
		if !defenses[d.Defense] {
			warn("guide/module-tags.json: the defense %q is no entry of guide/modules.md's %q", d.Defense, defensesSection)
		}
		delete(defenses, d.Defense)
		def := Defense{Name: d.Defense, Cards: resolve(d.Cards, d.Defense)}
		for _, a := range d.Answers {
			ans := Answers{How: a.How}
			for _, c := range a.Cards {
				name, level, _ := strings.Cut(c, " [")
				t := resolve([]string{name}, d.Defense)
				if len(t) == 0 {
					continue
				}
				var alts [][]string
				if level != "" {
					for alt := range strings.SplitSeq("["+level, " or ") {
						var set []string
						for _, m := range reLevelSymbol.FindAllStringSubmatch(alt, -1) {
							if !slices.ContainsFunc(disciplines[name], func(d string) bool { return strings.EqualFold(d, m[1]) }) {
								warn("guide/module-tags.json: %q answering %q names [%s], which it does not require", name, d.Defense, m[1])
							}
							set = append(set, m[1])
						}
						alts = append(alts, set)
					}
				}
				ans.Cards = append(ans.Cards, Answer{t[0], alts})
			}
			def.Answers = append(def.Answers, ans)
		}
		defs = append(defs, def)
	}
	for _, n := range slices.Sorted(maps.Keys(defenses)) {
		warn("guide/modules.md: the defense %q has no tags in guide/module-tags.json", n)
	}
	return out, defs, nil
}

// untagged names the library cards that the TWDA decks of the archetype
// labels' window play and no module tags, the most played first.
func (b *Bundle) untagged() []string {
	tagged := map[int]bool{}
	for _, m := range b.Modules {
		for _, t := range slices.Concat(m.Core, m.Support) {
			for _, id := range t.IDs {
				tagged[id] = true
			}
		}
	}
	byID := map[int]*Card{}
	for i := range b.Cards {
		byID[b.Cards[i].ID] = &b.Cards[i]
	}
	decks := map[int]int{}
	for _, d := range b.Decks {
		if d.Date < b.Classification.Since {
			continue
		}
		for _, c := range d.Cards {
			if t := byID[c.ID].Types; !tagged[c.ID] && !slices.Contains(t, "Vampire") && !slices.Contains(t, "Imbued") {
				decks[c.ID]++
			}
		}
	}
	order := slices.SortedFunc(maps.Keys(decks), func(x, y int) int {
		return cmp.Or(cmp.Compare(decks[y], decks[x]), strings.Compare(byID[x].Name, byID[y].Name))
	})
	out := make([]string, len(order))
	for i, id := range order {
		out[i] = fmt.Sprintf("%s (%d)", byID[id].Name, decks[id])
	}
	return out
}
