package tools

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

// AuditOptions narrows an uncited-posts audit for the rulemongers.
type AuditOptions struct {
	Since    string   // YYYY-MM-DD lower bound on post dates
	TWDAOnly bool     // only cards played in the TWDA window
	Cards    []string // optional card names to restrict the audit to
}

// Audit lists by card (most played in the TWDA window first) and by card pair
// the sitting rules directors' posts naming them that no ruling on those cards
// cites in vtes-rulings.
func Audit(b *bundle.Bundle, a *bundle.Archive, opt AuditOptions) (string, error) {
	if opt.Since != "" && !reDay.MatchString(opt.Since) {
		return "", fmt.Errorf("invalid -since %q: use YYYY, YYYY-MM or YYYY-MM-DD", opt.Since)
	}
	d := NewLive(b).p.Load()
	as := newArchiveSearch(a)

	type entry struct {
		feat    string
		card    *bundle.Card
		cards   []*bundle.Card
		spelled []string
		twda    int
		cited   map[string]bool
		rulings []bundle.Ruling
		changes []bundle.Change
	}
	entryName := func(e *entry) string {
		if e.card.IsCrypt() {
			if e.card.Advanced {
				return e.card.PrintedName + " (ADV)"
			}
			return e.card.PrintedName
		}
		return e.card.Name
	}

	byFeat := map[string]*entry{}
	var entries []*entry
	addCard := func(c *bundle.Card) *entry {
		feat := bundle.Feature(c)
		e := byFeat[feat]
		if e == nil {
			e = &entry{feat: feat, card: c, cited: map[string]bool{}}
			byFeat[feat] = e
			entries = append(entries, e)
		}
		if slices.Contains(e.cards, c) {
			return e
		}
		e.cards = append(e.cards, c)
		for _, sp := range spellings(c) {
			if !slices.Contains(e.spelled, sp) {
				e.spelled = append(e.spelled, sp)
			}
		}
		for _, r := range c.Rulings {
			for _, ref := range r.References {
				if ref.Archive != "" {
					e.cited[ref.Archive] = true
				}
			}
			if !slices.ContainsFunc(e.rulings, func(x bundle.Ruling) bool { return x.Link == r.Link && x.Text == r.Text }) {
				e.rulings = append(e.rulings, r)
			}
		}
		for _, ch := range c.Changes {
			if !slices.ContainsFunc(e.changes, func(x bundle.Change) bool { return x.Date == ch.Date && x.Text == ch.Text }) {
				e.changes = append(e.changes, ch)
			}
		}
		return e
	}

	if len(opt.Cards) > 0 {
		for _, name := range opt.Cards {
			cards, note := d.ix.resolve(name, "audit")
			if cards == nil {
				return "", fmt.Errorf("%s", note)
			}
			for _, c := range cards {
				addCard(c)
			}
		}
	} else {
		for i := range b.Cards {
			addCard(&b.Cards[i])
		}
	}

	// Count TWDA window decks per card feature.
	windowDecks, windowSince := 0, ""
	if d.rel != nil {
		windowDecks, windowSince = len(d.rel.Decks), d.rel.Class.Since
		for _, deck := range d.rel.Decks {
			seen := map[string]bool{}
			for _, dc := range deck.Cards {
				c := d.byID[dc.ID]
				if c == nil {
					continue
				}
				feat := bundle.Feature(c)
				if !seen[feat] {
					seen[feat] = true
					if e := byFeat[feat]; e != nil {
						e.twda++
					}
				}
			}
		}
	}

	if opt.TWDAOnly {
		entries = slices.DeleteFunc(entries, func(e *entry) bool { return e.twda == 0 })
	}

	// Build an inverted word index over sitting directors' posts (within opt.Since)
	// so scanning all cards against the full archive takes seconds.
	byWord := map[string][]int{}
	directorCount := 0
	for _, k := range as.directors {
		m := as.msgs[k]
		if opt.Since != "" && m.thread.Messages[m.i].Date < opt.Since {
			continue
		}
		directorCount++
		seen := map[string]bool{}
		for _, text := range []string{m.own, m.question} {
			for _, w := range foldedWords(text) {
				if !seen[w] {
					seen[w] = true
					byWord[w] = append(byWord[w], k)
				}
			}
		}
	}

	isCited := func(e *entry, k int) bool {
		m := as.msgs[k]
		base := "https://usenet.krcg.org/t/" + m.thread.ID + "/"
		id := m.thread.Messages[m.i].ID
		return e.cited[base+"#m"+strconv.Itoa(m.i)] || id != "" && e.cited[base+"#"+id] || m.i == 0 && e.cited[base]
	}

	type cardHit struct {
		entry *entry
		named named
	}
	cardPosts := map[*entry][]int{}
	postCards := map[int][]cardHit{}

	for _, e := range entries {
		// Gather candidate director posts using the rarest word of each spelling.
		cands := map[int]bool{}
		for _, sp := range e.spelled {
			ws := foldedWords(sp)
			if len(ws) == 0 {
				continue
			}
			rarest := byWord[ws[0]]
			for _, w := range ws[1:] {
				if list := byWord[w]; len(list) < len(rarest) {
					rarest = list
				}
			}
			for _, k := range rarest {
				cands[k] = true
			}
		}
		names1 := func(text string) bool {
			return slices.ContainsFunc(e.spelled, func(n string) bool { return countWord(text, n) > 0 })
		}
		for k := range cands {
			if isCited(e, k) {
				continue
			}
			m := as.msgs[k]
			var n named
			if names1(m.own) {
				if own := bundle.Fold(prose(ownLines(m.thread.Messages[m.i].Body))); names1(own) {
					n.own = own
				}
			}
			if names1(m.question) {
				n.question = m.question
			}
			if n == (named{}) {
				continue
			}
			cardPosts[e] = append(cardPosts[e], k)
			postCards[k] = append(postCards[k], cardHit{entry: e, named: n})
		}
	}

	// Build uncited card pairs from posts naming two or more cards close together.
	type pairKey [2]*entry
	pairPosts := map[pairKey][]int{}
	spelledMap := map[*bundle.Card][]string{}
	for _, e := range entries {
		spelledMap[e.card] = e.spelled
	}
	for k, hits := range postCards {
		if len(hits) < 2 {
			continue
		}
		slices.SortFunc(hits, func(a, b cardHit) int { return cmp.Compare(a.entry.card.ID, b.entry.card.ID) })
		for i := range hits {
			for j := i + 1; j < len(hits); j++ {
				a, b := hits[i], hits[j]
				if a.entry.card.IsCrypt() && b.entry.card.IsCrypt() && a.entry.card.PrintedName == b.entry.card.PrintedName {
					continue
				}
				together := func(x, y string) bool {
					return x != "" && y != "" && closeTogether(x, a.entry.spelled, b.entry.spelled)
				}
				own := together(a.named.own, b.named.own)
				if !own && !(together(a.named.question, b.named.question) && shownNames(as, k, [2]*bundle.Card{a.entry.card, b.entry.card}, spelledMap)) {
					continue
				}
				pk := pairKey{a.entry, b.entry}
				pairPosts[pk] = append(pairPosts[pk], k)
			}
		}
	}

	// Sort cards: most played in the TWDA window first, then uncited post count, then name.
	cardOrder := slices.Collect(maps.Keys(cardPosts))
	slices.SortFunc(cardOrder, func(a, b *entry) int {
		return cmp.Or(
			cmp.Compare(b.twda, a.twda),
			cmp.Compare(len(cardPosts[b]), len(cardPosts[a])),
			strings.Compare(entryName(a), entryName(b)),
		)
	})

	// Sort pairs: highest combined TWDA play first, then minimum of the pair, then post count, then names.
	pairOrder := slices.Collect(maps.Keys(pairPosts))
	slices.SortFunc(pairOrder, func(a, b pairKey) int {
		return cmp.Or(
			cmp.Compare(b[0].twda+b[1].twda, a[0].twda+a[1].twda),
			cmp.Compare(min(b[0].twda, b[1].twda), min(a[0].twda, a[1].twda)),
			cmp.Compare(len(pairPosts[b]), len(pairPosts[a])),
			strings.Compare(entryName(a[0]), entryName(b[0])),
			strings.Compare(entryName(a[1]), entryName(b[1])),
		)
	})

	sortPosts := func(ks []int) {
		slices.SortFunc(ks, func(a, b int) int {
			ma, mb := as.msgs[a], as.msgs[b]
			return cmp.Or(
				strings.Compare(mb.thread.Messages[mb.i].Date, ma.thread.Messages[ma.i].Date),
				strings.Compare(ma.thread.ID, mb.thread.ID),
				cmp.Compare(ma.i, mb.i),
			)
		})
	}

	var out strings.Builder
	out.WriteString("# Uncited rules directors' posts\n\n")
	fmt.Fprintf(&out, "Scope: %d cards, %d sitting rules directors' posts (archive to %s", len(entries), directorCount, a.LastPost)
	if windowDecks > 0 {
		fmt.Fprintf(&out, "; TWDA window: %d decks since %s", windowDecks, windowSince)
	}
	if opt.Since != "" {
		fmt.Fprintf(&out, "; posts since %s", opt.Since)
	}
	if opt.TWDAOnly {
		out.WriteString("; TWDA-played cards only")
	}
	out.WriteString(").\n")
	fmt.Fprintf(&out, "Found %d cards and %d card pairs with uncited sitting rules directors' posts (%d distinct posts).\n\n",
		len(cardOrder), len(pairOrder), len(postCards))

	writeRulings := func(e *entry) {
		if len(e.rulings) == 0 {
			fmt.Fprintf(&out, "Existing rulings on %s: none.\n", entryName(e))
			return
		}
		fmt.Fprintf(&out, "Existing rulings on %s:\n", entryName(e))
		for _, r := range e.rulings {
			fmt.Fprintf(&out, "- %s (%s)\n", reCardToken.ReplaceAllString(r.Text, "$1"), r.Link)
		}
	}

	writeChangesAfter := func(e *entry, postDate string) {
		var after []bundle.Change
		for _, ch := range e.changes {
			if postDate != "" && ch.Date > postDate[:min(len(postDate), len(ch.Date))] {
				after = append(after, ch)
			}
		}
		if len(after) == 0 {
			fmt.Fprintf(&out, "Card changes after %s (%s): none.\n", postDate, entryName(e))
			return
		}
		fmt.Fprintf(&out, "Card changes after %s (%s):\n", postDate, entryName(e))
		for _, ch := range after {
			if ch.URL != "" {
				fmt.Fprintf(&out, "- %s: %s (%s)\n", ch.Date, ch.Text, ch.URL)
			} else {
				fmt.Fprintf(&out, "- %s: %s\n", ch.Date, ch.Text)
			}
		}
	}

	writeAuditPost := func(k int, targets []*entry) {
		m := as.msgs[k]
		msg := m.thread.Messages[m.i]
		link := fmt.Sprintf("https://usenet.krcg.org/t/%s/#m%d", m.thread.ID, m.i)
		var needles []string
		for _, e := range targets {
			needles = append(needles, e.spelled...)
		}
		body, before := auditExcerpt(m.thread, m.i, needles, 600, 1000)
		fmt.Fprintf(&out, "#### %s · %s · %q\n", msg.Date, bundle.Role(msg.Author, msg.Date).Label, m.thread.Title)
		fmt.Fprintf(&out, "Link: %s\n", link)
		if before != "" && m.i > 0 && m.thread.Messages[m.i-1].Body != "" {
			fmt.Fprintf(&out, "Question context: %s\n", before)
		}
		for _, e := range targets {
			writeChangesAfter(e, msg.Date)
		}
		out.WriteString("\n")
		for line := range strings.Lines(body) {
			out.WriteString("| " + strings.TrimRight(line, " \r\n") + "\n")
		}
		out.WriteString("\n")
	}

	out.WriteString("## By card (most played in the TWDA window first)\n\n")
	if len(cardOrder) == 0 {
		out.WriteString("None.\n\n")
	}
	for _, e := range cardOrder {
		ks := cardPosts[e]
		sortPosts(ks)
		fmt.Fprintf(&out, "### %s (%d TWDA decks, %d uncited posts)\n\n", entryName(e), e.twda, len(ks))
		writeRulings(e)
		out.WriteString("\n")
		for _, k := range ks {
			writeAuditPost(k, []*entry{e})
		}
	}

	out.WriteString("## By card pair (most played in the TWDA window first)\n\n")
	if len(pairOrder) == 0 {
		out.WriteString("None.\n\n")
	}
	for _, pk := range pairOrder {
		ks := pairPosts[pk]
		sortPosts(ks)
		fmt.Fprintf(&out, "### %s + %s (%d + %d TWDA decks, %d uncited posts)\n\n",
			entryName(pk[0]), entryName(pk[1]), pk[0].twda, pk[1].twda, len(ks))
		writeRulings(pk[0])
		writeRulings(pk[1])
		out.WriteString("\n")
		for _, k := range ks {
			writeAuditPost(k, []*entry{pk[0], pk[1]})
		}
	}

	out.WriteString(DataLine(b))
	return out.String(), nil
}

// auditExcerpt returns the quoted question and answer around needles: answering
// when a quoted block holds a needle, the reply quotes nothing, or the reply's
// own lines hold no needle; otherwise the reply's body windowed around needles.
func auditExcerpt(th *bundle.Thread, i int, needles []string, qLimit, ownLimit int) (body, before string) {
	raw := th.Messages[i].Body
	hasNeedle := func(folded string) bool {
		return slices.ContainsFunc(needles, func(nd string) bool { return countWord(folded, nd) > 0 })
	}
	if hasNeedle(bundle.Fold(unquote(raw))) || quotesNone(th, i) || !hasNeedle(bundle.Fold(ownLines(raw))) {
		full, bef := answering(th, i, needles, 0, 0)
		if full != "" {
			if q := strings.TrimSpace(unquote(full)); q != "" {
				own := strings.TrimSpace(ownLines(full))
				q = windowText(q, needles, qLimit)
				own = windowText(own, needles, ownLimit)
				return "> " + strings.ReplaceAll(q, "\n", "\n> ") + "\n\n" + own, bef
			}
			return windowText(full, needles, ownLimit), bef
		}
	}
	return windowText(raw, needles, ownLimit), ""
}

// windowText keeps the region around needles up to limit bytes, stepping down
// to lines when the matching paragraph itself exceeds limit.
func windowText(text string, needles []string, limit int) string {
	text = strings.TrimSpace(strings.ReplaceAll(text, "\r\n", "\n"))
	if limit <= 0 || len(text) <= limit {
		return text
	}
	paras := strings.Split(text, "\n\n")
	best, most := 0, 0
	for i, p := range paras {
		f, found := bundle.Fold(p), 0
		for _, n := range needles {
			if countWord(f, n) > 0 {
				found++
			}
		}
		if found > most {
			best, most = i, found
		}
	}
	if len(paras[best]) > limit {
		lines := strings.Split(paras[best], "\n")
		hit := 0
		for i, l := range lines {
			f := bundle.Fold(l)
			if slices.ContainsFunc(needles, func(n string) bool { return countWord(f, n) > 0 }) {
				hit = i
				break
			}
		}
		start := max(0, hit-2)
		var out []string
		if best > 0 || start > 0 {
			out = append(out, "[…]")
		}
		size := 0
		for _, l := range lines[start:] {
			if size+len(l) > limit {
				if c := clip(l, limit-size); c != "[…]" || len(out) == 0 || out[len(out)-1] != "[…]" {
					out = append(out, c)
				}
				break
			}
			out = append(out, l)
			size += len(l)
		}
		return strings.Join(out, "\n")
	}
	start := best
	if best > 0 && len(paras[best-1]) <= limit/2 {
		start = best - 1
	}
	var out []string
	if start > 0 {
		out = append(out, "[…]")
	}
	size := 0
	for _, p := range paras[start:] {
		if size+len(p) > limit {
			if c := clip(p, limit-size); c != "[…]" || len(out) == 0 || out[len(out)-1] != "[…]" {
				out = append(out, c)
			}
			break
		}
		out = append(out, p)
		size += len(p)
	}
	return strings.Join(out, "\n\n")
}

// foldedWords splits a folded string into word runs matching isWordByte.
func foldedWords(s string) []string {
	var ws []string
	start := -1
	for i := 0; i < len(s); i++ {
		if isWordByte(s[i]) {
			if start < 0 {
				start = i
			}
		} else if start >= 0 {
			ws = append(ws, s[start:i])
			start = -1
		}
	}
	if start >= 0 {
		ws = append(ws, s[start:])
	}
	return ws
}
