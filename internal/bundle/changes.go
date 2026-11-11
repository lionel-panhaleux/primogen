package bundle

import (
	"cmp"
	"regexp"
	"slices"
	"strings"
	"unicode"
)

var (
	reDate         = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	reAnnouncement = regexp.MustCompile(`^\[(.+)\]\((.+)\)$`)
	reNameList     = regexp.MustCompile(`,\s*(?:and\s+)?|\s+and\s+`)
)

// attachChanges files each "- Card: change" line of the history under the
// cards it names, or as possible under the cards closest to a name it misspells.
func (b *Bundle) attachChanges(history []byte, warn func(string, ...any)) {
	byName := map[string][]*Card{}
	byWord := map[string][]*Card{}
	for i := range b.Cards {
		c := &b.Cards[i]
		for _, n := range c.Names() {
			k := Fold(n)
			if !slices.Contains(byName[k], c) {
				byName[k] = append(byName[k], c)
			}
			for _, w := range words(k) {
				if !slices.Contains(byWord[w], c) {
					byWord[w] = append(byWord[w], c)
				}
			}
		}
	}
	// exact also says whether a name matched several cards (a vampire's groups).
	exact := func(name string) ([]*Card, bool) {
		if base, ok := strings.CutSuffix(name, " ADV"); ok && byName[Fold(name)] == nil {
			name = base + " (ADV)"
		}
		if m := byName[Fold(name)]; m != nil {
			return m, len(m) > 1
		}
		parts := reNameList.Split(name, -1)
		if len(parts) < 2 {
			return nil, false
		}
		var all []*Card
		shared := false
		for _, part := range parts {
			m := byName[Fold(part)]
			if m == nil {
				return nil, false
			}
			all, shared = append(all, m...), shared || len(m) > 1
		}
		return all, shared
	}
	near := func(name string) []*Card {
		k := Fold(name)
		var found []*Card
		for n, cards := range byName {
			if Distance(k, n) <= max(1, len(k)/6) {
				found = append(found, cards...)
			}
		}
		if found == nil { // the rarest word of the name, if rare enough
			for _, w := range words(k) {
				if m := byWord[w]; len(m) > 0 && len(m) <= 5 && (found == nil || len(m) < len(found)) {
					found = m
				}
			}
		}
		found = slices.SortedFunc(slices.Values(found), func(a, b *Card) int { return cmp.Compare(a.ID, b.ID) })
		return slices.Compact(found)
	}

	var change Change
	for line := range strings.Lines(string(history)) {
		line = strings.TrimSpace(line)
		if date, ok := strings.CutPrefix(line, "# "); ok && reDate.MatchString(date) {
			change = Change{Date: date}
			continue
		}
		if m := reAnnouncement.FindStringSubmatch(line); m != nil && change.Date != "" && change.URL == "" {
			change.Announcement, change.URL = m[1], m[2]
			continue
		}
		item, ok := strings.CutPrefix(line, "- ")
		if !ok || change.Date == "" {
			continue
		}
		// A name can hold a colon itself (Powerbase: Madrid): the longest
		// prefix naming cards is the name.
		c := change
		c.Text = item
		var matched []*Card
		for i := range item {
			sep := 0
			if item[i] == ':' {
				sep = 1
			} else if strings.HasPrefix(item[i:], " - ") {
				sep = 3
			}
			if sep == 0 {
				continue
			}
			if c.Named == "" {
				c.Named, c.Text = item[:i], item[i+sep:]
			}
			if m, shared := exact(item[:i]); m != nil {
				matched, c.Shared, c.Named, c.Text = m, shared, item[:i], item[i+sep:]
			}
		}
		c.Text = strings.TrimSpace(c.Text)
		if matched == nil && c.Named != "" {
			matched, c.Possible = near(c.Named), true
		}
		// A change older than a card's first printing cannot concern it.
		matched = slices.DeleteFunc(slices.Clone(matched), func(card *Card) bool { return change.Date < card.Legal })
		c.Shared = c.Shared && len(matched) > 1
		if len(matched) == 0 {
			warn("card-changes.md %s: no card named in %q", change.Date, line)
			continue
		}
		if c.Possible {
			var names []string
			for _, m := range matched {
				names = append(names, m.Name)
			}
			warn("card-changes.md %s: %q names no card exactly; filed as possible under %s", change.Date, c.Named, strings.Join(names, "; "))
		}
		for _, card := range matched {
			card.Changes = append(card.Changes, c)
		}
	}
}

func words(folded string) []string {
	return strings.FieldsFunc(folded, func(r rune) bool { return !unicode.IsLetter(r) })
}
