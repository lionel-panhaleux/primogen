package tools

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

// A library card is rarely played by fewer than rareShare of the window's
// decks; one legal for fewer months than newMonths before the window's
// newest deck is new.
const (
	rareShare = 0.01
	newMonths = 18
)

// rarelyPlayed says whether a card k of n decks play is rarely played.
func rarelyPlayed(k, n int) bool {
	return float64(k) < rareShare*float64(n)
}

// window is how many of the window's decks play each card, by feature (a
// vampire's groups merged), each window deck's share of rarely played copies
// in its library, by id, itself left out of the counts, and the date from
// which a card is new.
func (d *data) window() (map[string]int, map[string]float64, string) {
	d.windowOnce.Do(func() {
		newest := ""
		d.windowPlay = map[string]int{}
		for _, t := range d.rel.Decks {
			newest = max(newest, t.Date)
			for _, f := range d.features(t.Cards) {
				d.windowPlay[f]++
			}
		}
		at, _ := time.Parse(time.DateOnly, newest) // the TWDA's dates are ISO
		d.newSince = at.AddDate(0, -newMonths, 0).Format(time.DateOnly)
		d.windowRare = map[string]float64{}
		n := len(d.rel.Decks)
		for _, t := range d.rel.Decks {
			if share, ok := d.rareCopies(t.Cards, n-1, func(f string) int { return d.windowPlay[f] - 1 }); ok {
				d.windowRare[t.ID] = share
			}
		}
	})
	return d.windowPlay, d.windowRare, d.newSince
}

func (d *data) features(cards []bundle.DeckCard) []string {
	var fs []string
	for _, c := range cards {
		if f := bundle.Feature(d.byID[c.ID]); !slices.Contains(fs, f) {
			fs = append(fs, f)
		}
	}
	return fs
}

// rareCopies is the share, in percent, of a list's library copies in cards
// rarely played among n decks, new cards left out of both; false when no
// library card is counted.
func (d *data) rareCopies(cards []bundle.DeckCard, n int, plays func(string) int) (float64, bool) {
	counted, rare := 0, 0
	for _, e := range cards {
		c := d.byID[e.ID]
		if c.IsCrypt() || c.Legal >= d.newSince {
			continue
		}
		counted += e.Count
		if rarelyPlayed(plays(bundle.Feature(c)), n) {
			rare += e.Count
		}
	}
	return pct(rare, counted), counted > 0
}

// writeRecord says how winners have played this deck's cards: a list's share
// of rarely played copies beside its reference, and each card no window deck
// plays with its record in the whole TWDA.
func (dk *deck) writeRecord(out *strings.Builder, d *data, ref *reference, today string, fit int) {
	out.WriteString("## Track record\n\n")
	r := d.rel
	if r == nil {
		out.WriteString("This data bundle holds no archetype labels: no track record until it is rebuilt.\n\n")
		return
	}
	plays, rareOf, since := d.window()
	n := len(r.Decks)
	own := func(f string) int { return plays[f] }
	leftOut := ""
	if t := dk.twda; t != nil && slices.ContainsFunc(r.Decks, func(w *bundle.Deck) bool { return w.ID == t.ID }) {
		n--
		mine := d.features(t.Cards)
		own = func(f string) int {
			if slices.Contains(mine, f) {
				return plays[f] - 1
			}
			return plays[f]
		}
		leftOut = " (this deck left out of the counts)"
	}
	fewest := 0
	for rarelyPlayed(fewest+1, n) {
		fewest++
	}
	fmt.Fprintf(out, "Among the %d TWDA decks since %s%s. A card is rarely played when fewer than %.0f%% of them play it (%d decks or fewer), "+
		"and new when legal on or after %s, %d months before the newest of them: a new card had little time to be played, so it is listed apart and not counted.\n\n",
		len(r.Decks), r.Class.Since, leftOut, 100*rareShare, fewest, since, newMonths)

	var list []bundle.DeckCard
	var fresh []string
	library := false
	for _, e := range dk.entries {
		list = append(list, bundle.DeckCard{ID: e.card.ID, Count: e.count})
		if !e.card.IsCrypt() {
			library = true
			if e.card.Legal >= since {
				fresh = append(fresh, fmt.Sprintf("%s (legal since %s)", e.card.Name, e.card.Legal))
			}
		}
	}
	switch share, ok := d.rareCopies(list, n, own); {
	case !library:
		out.WriteString("- No library card given: no share of rarely played copies.\n")
	case dk.draft:
		out.WriteString("- A deck idea's share of rarely played copies is not computed: its counts are not a list's.\n")
	case !ok:
		out.WriteString("- Every library card is new: no share of rarely played copies.\n")
	default:
		var s []float64
		for _, t := range ref.decks {
			if v, ok := rareOf[t.ID]; ok {
				s = append(s, v)
			}
		}
		slices.Sort(s)
		fmt.Fprintf(out, "- Library copies in rarely played cards: %.0f%% of those not new. Beside the reference decks of the section above, each against the others, in %%: %s (min, 10%%, 25%%, median, 75%%, 90%%, max); %s.\n",
			share, spread(s), rank(s, share))
	}
	if len(fresh) > 0 {
		fmt.Fprintf(out, "- New, not counted: %s.\n", strings.Join(fresh, ", "))
	}

	type unplayed struct {
		card       *bundle.Card
		ever       int
		last, ofID string
	}
	byFeature := map[string]*unplayed{}
	for _, e := range dk.entries {
		if f := bundle.Feature(e.card); own(f) == 0 && byFeature[f] == nil {
			byFeature[f] = &unplayed{card: e.card}
		}
	}
	if len(byFeature) == 0 {
		out.WriteString("- Every card of this deck is played by at least one of these decks.\n\n")
		return
	}
	for i := range d.b.Decks {
		t := &d.b.Decks[i]
		for _, f := range d.features(t.Cards) {
			if u := byFeature[f]; u != nil {
				u.ever++
				if t.Date > u.last {
					u.last, u.ofID = t.Date, t.ID
				}
			}
		}
	}
	var none []unplayed
	for _, u := range byFeature {
		none = append(none, *u)
	}
	slices.SortFunc(none, func(a, b unplayed) int {
		return cmp.Or(cmp.Compare(a.ever, b.ever), strings.Compare(a.card.Legal, b.card.Legal), strings.Compare(a.card.Name, b.card.Name))
	})
	if fit > full {
		var names []string
		for _, u := range none[:min(len(none), fewShown)] {
			names = append(names, u.card.Name)
		}
		if k := len(none) - fewShown; k > 0 {
			names = append(names, fmt.Sprintf("and %d more", k))
		}
		fmt.Fprintf(out, "- Played by none of these decks, %d cards (their record in the whole TWDA, and past %d their names, left out to fit this result; detail gives a larger one): %s.\n\n",
			len(none), fewShown, strings.Join(names, ", "))
		return
	}
	inputs := ""
	if none[len(none)-1].ever > 0 {
		inputs = " (the last deck's id is an input of vtes_deck)"
	}
	fmt.Fprintf(out, "- Played by none of these decks; each searched in the whole TWDA, %d decks%s%s:\n",
		len(d.b.Decks), map[bool]string{true: ", this deck included"}[dk.twda != nil], inputs)
	for _, u := range none {
		legal := "no legality date in the card data"
		if u.card.Legal != "" {
			legal = "legal since " + u.card.Legal
			at, err1 := time.Parse(time.DateOnly, u.card.Legal)
			now, err2 := time.Parse(time.DateOnly, today)
			if err1 == nil && err2 == nil {
				if y := int(now.Sub(at).Hours() / 24 / 365.25); y >= 1 {
					legal += fmt.Sprintf(", %d year%s", y, map[bool]string{true: "s"}[y > 1])
				}
			}
		}
		if u.ever == 0 {
			fmt.Fprintf(out, "  - %s: %s; in no TWDA deck.\n", u.card.Name, legal)
		} else {
			fmt.Fprintf(out, "  - %s: %s; in %d TWDA deck%s, the last of %s (%s).\n", u.card.Name, legal, u.ever, map[bool]string{true: "s"}[u.ever > 1], u.last, u.ofID)
		}
	}
	out.WriteString("\n")
}
