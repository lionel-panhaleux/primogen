package tools

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"math"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

const maxDeckInput = 20000

type deckInput struct {
	Deck   string `json:"deck"`
	Detail bool   `json:"detail,omitempty"`
}

var deckSchema = &jsonschema.Schema{
	Type: "object",
	Properties: map[string]*jsonschema.Schema{
		"deck": {Type: "string", Description: `the deck: a decklist as text, one card per line with its count ("3x Conditioning", ` +
			`"Anson x2", crypt lines with their capacity and disciplines accepted), a TWDA deck id ("2k6munduko"), or a deck link from VDB (https://vdb.im/decks/...), VTES Decks (https://vtesdecks.com/deck/...) or Amaranth (https://amaranth.vtes.co.nz/#deck/...)`},
		"detail": {Type: "boolean", Description: "the deck author's whole comment, where the default cuts it, and a larger result, keeping texts a large list's default drops"},
	},
	Required: []string{"deck"},
}

func addDeck(s *mcp.Server, l *Live) {
	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_deck",
		Title: "VTES deck",
		Description: "Call first to analyze or improve a Vampire: The Eternal Struggle (VTES) deck: never read a decklist from memory " +
			"of its cards. Takes a decklist as text, a TWDA deck id or a deck link from VDB, VTES Decks or Amaranth; validates it (sizes, crypt groups, banned cards, " +
			"legality today) and returns the computed facts: the crypt table with capacities, disciplines, titles and texts, " +
			"which crypt cards have what each library card requires and at which level, the library by type with every card's official text, " +
			"the library by module with each module's density and expected copies per hand, the costs, the nearest archetypes and tournament decks with the cards that differ, and the procedure a deck analysis and advice follow.",
		InputSchema: deckSchema,
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: new(false)},
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in deckInput) (*mcp.CallToolResult, any, error) {
		text := strings.TrimSpace(in.Deck)
		if text == "" || len(text) > maxDeckInput {
			return nil, nil, fmt.Errorf(`give one deck of at most %d characters: a decklist as text, a TWDA id or a deck link, e.g. {"deck": "2k6munduko"}`, maxDeckInput)
		}
		d := l.p.Load()
		dk, err := d.readDeck(ctx, l, text)
		if err != nil {
			return nil, nil, err
		}
		today := time.Now().UTC().Format(time.DateOnly)
		fixes := dk.fixes(today)
		if dk.twda == nil && len(fixes) > 0 {
			return nil, nil, fmt.Errorf("this list is not tournament legal today (%s); fix it and call vtes_deck again:\n- %s",
				today, strings.Join(fixes, "\n- "))
		}
		limit := budget
		if in.Detail {
			limit = detailBudget
		}
		var out strings.Builder
		for fit := full; ; fit++ {
			out.Reset()
			if dk.write(&out, d, today, fixes, in.Detail, fit); out.Len()+len(DataLine(d.b)) <= limit {
				break
			}
			if fit == bare {
				return nil, nil, fmt.Errorf("this list of %d distinct cards does not fit a result of %d characters even without its texts: give a list of fewer distinct cards", len(dk.entries), limit)
			}
		}
		out.WriteString(DataLine(d.b))
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: out.String()}}}, nil, nil
	})
}

type deckEntry struct {
	card    *bundle.Card
	count   int
	comment string
}

type deck struct {
	name, author, description string
	source                    string
	twda                      *bundle.Deck
	entries                   []deckEntry
	notes                     []string // how the list was read
}

var reSchemeless = regexp.MustCompile(`(?i)^(www\.)?(vdb\.im|vtesdecks\.com|amaranth\.vtes\.co\.nz)/`)

func (d *data) readDeck(ctx context.Context, l *Live, text string) (*deck, error) {
	if reSchemeless.MatchString(text) {
		text = "https://" + text
	}
	if u, err := url.Parse(text); err == nil && u.Host != "" && (u.Scheme == "https" || u.Scheme == "http") {
		switch strings.TrimPrefix(strings.ToLower(u.Host), "www.") {
		case "vdb.im":
			return d.readVDB(ctx, l, u)
		case "vtesdecks.com":
			return d.readVTESDecks(ctx, l, u)
		case "amaranth.vtes.co.nz":
			return d.readAmaranth(ctx, l, u)
		}
		return nil, fmt.Errorf("primogen reads deck links from VDB (vdb.im), VTES Decks (vtesdecks.com) and Amaranth (amaranth.vtes.co.nz), not %s: give the list as text, one card per line with its count", u.Host)
	}
	if !strings.ContainsAny(text, " \t\n") {
		if t := d.decks[text]; t != nil {
			return d.fromTWDA(t), nil
		}
		return nil, fmt.Errorf("no deck of the TWDA has the id %q: give a TWDA id (e.g. 2k6munduko), a deck link, or the list as text, one card per line with its count", text)
	}
	return d.readText(text)
}

func (d *data) fromTWDA(t *bundle.Deck) *deck {
	dk := &deck{name: t.Name, author: cmp.Or(t.Author, t.Player), twda: t}
	for _, c := range t.Cards {
		dk.entries = append(dk.entries, deckEntry{d.byID[c.ID], c.Count, c.Comment})
	}
	return dk
}

// fetch reads a deckbuilder's API answer into v; site and what name the
// deckbuilder and what is fetched in the errors.
func (l *Live) fetch(ctx context.Context, site, host, path, what string, v any) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cmp.Or(l.Sites, "https:/")+"/"+host+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "primogen")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s did not answer for %s: %v; give the list as text instead", site, what, err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return fmt.Errorf("%s has no public %s (%s): check the link, or give the list as text", site, what, resp.Status)
	default:
		return fmt.Errorf("%s did not answer for %s (%s); give the list as text instead", site, what, resp.Status)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("%s's answer for %s: %w; give the list as text instead", site, what, err)
	}
	return nil
}

// stored is a deck as a deckbuilder stores it, its cards by card id.
type stored struct {
	Name, Author, Description string
	Cards                     map[string]int
}

func (d *data) fromStored(site string, u *url.URL, s stored) (*deck, error) {
	source := *u
	source.Fragment = ""
	dk := &deck{name: s.Name, author: s.Author, description: s.Description, source: source.String()}
	var unknown []string
	for _, k := range slices.Sorted(maps.Keys(s.Cards)) {
		id, _ := strconv.Atoi(k)
		c := d.byID[id]
		if c == nil {
			unknown = append(unknown, k)
			continue
		}
		if s.Cards[k] > 0 {
			dk.entries = append(dk.entries, deckEntry{card: c, count: s.Cards[k]})
		}
	}
	if len(unknown) > 0 {
		return nil, fmt.Errorf("the %s deck names cards the card data does not hold (ids %s): give the list as text", site, strings.Join(unknown, ", "))
	}
	if len(dk.entries) == 0 {
		return nil, fmt.Errorf("the %s deck %s holds no card", site, u)
	}
	return dk, nil
}

func (d *data) readVDB(ctx context.Context, l *Live, u *url.URL) (*deck, error) {
	id := strings.TrimPrefix(u.Path, "/decks/")
	// VDB names a TWDA deck of a nine-character id with a trailing 0
	twda := d.decks[id]
	if len(id) == 10 && twda == nil {
		twda = d.decks[strings.TrimSuffix(id, "0")]
	}
	var s stored
	switch {
	case u.Path == "/decks/deck":
		s.Name, s.Author, s.Description = u.Query().Get("name"), u.Query().Get("author"), u.Query().Get("description")
		s.Cards = map[string]int{}
		for pair := range strings.SplitSeq(u.Fragment, ";") {
			k, v, _ := strings.Cut(pair, "=")
			n, err := strconv.Atoi(v)
			if k == "" || err != nil {
				return nil, fmt.Errorf("the VDB link's cards, after #, are not id=count pairs separated by ';': %q", u.Fragment)
			}
			s.Cards[k] += n
		}
	case id == "" || strings.Contains(id, "/"):
		return nil, fmt.Errorf("not a VDB deck link: %s; give https://vdb.im/decks/<id> or the link VDB's \"Copy URL\" makes", u)
	case twda != nil:
		dk := d.fromTWDA(twda)
		dk.source = u.String()
		return dk, nil
	default:
		if err := l.fetch(ctx, "VDB", "vdb.im", "/api/deck/"+url.PathEscape(id), fmt.Sprintf("deck %q", id), &s); err != nil {
			return nil, err
		}
	}
	return d.fromStored("VDB", u, s)
}

// readVTESDecks reads a deck of vtesdecks.com, a TWDA deck (its id
// "tournament-" and the TWDA's) from the bundle.
func (d *data) readVTESDecks(ctx context.Context, l *Live, u *url.URL) (*deck, error) {
	id, ok := strings.CutPrefix(u.Path, "/deck/")
	if !ok || id == "" || strings.Contains(id, "/") {
		return nil, fmt.Errorf("not a VTES Decks deck link: %s; give https://vtesdecks.com/deck/<id>", u)
	}
	if twda, ok := strings.CutPrefix(id, "tournament-"); ok && d.decks[twda] != nil {
		dk := d.fromTWDA(d.decks[twda])
		dk.source = u.String()
		return dk, nil
	}
	var answer struct {
		Name, Author, Description string
		Crypt, Library            []struct{ ID, Number int }
	}
	if err := l.fetch(ctx, "VTES Decks", "api.vtesdecks.com", "/1.0/decks/"+url.PathEscape(id), fmt.Sprintf("deck %q", id), &answer); err != nil {
		return nil, err
	}
	s := stored{Name: answer.Name, Author: answer.Author, Description: answer.Description, Cards: map[string]int{}}
	for _, c := range slices.Concat(answer.Crypt, answer.Library) {
		s.Cards[strconv.Itoa(c.ID)] += c.Number
	}
	return d.fromStored("VTES Decks", u, s)
}

// Amaranth names an advanced vampire "Theo Bell (ADV2)", the card data "Theo
// Bell (G2 ADV)".
var reAmaranthAdvanced = regexp.MustCompile(`^(.+) \(ADV(\d)\)$`)

// readAmaranth reads a deck of Amaranth, whose cards are its own ids, named by
// its catalog: a name is the card's own, never a misspelling or another card's.
func (d *data) readAmaranth(ctx context.Context, l *Live, u *url.URL) (*deck, error) {
	// the link Amaranth shares routes in its fragment: /#deck/<id>
	id, ok := strings.CutPrefix(cmp.Or(u.Fragment, strings.TrimPrefix(u.Path, "/")), "deck/")
	if !ok || id == "" || strings.Contains(id, "/") {
		return nil, fmt.Errorf("not an Amaranth deck link: %s; give the link Amaranth shares, https://amaranth.vtes.co.nz/#deck/<id>", u)
	}
	var answer struct {
		Success bool
		Error   struct{ Message string }
		Result  struct {
			Title, Author, Description string
			Cards                      map[string]int
		}
	}
	if err := l.fetch(ctx, "Amaranth", "amaranth.vtes.co.nz", "/api/deck?id="+url.QueryEscape(id), fmt.Sprintf("deck %q", id), &answer); err != nil {
		return nil, err
	}
	if !answer.Success {
		return nil, fmt.Errorf("Amaranth has no public deck %q (%s): check the link, or give the list as text", id, answer.Error.Message)
	}
	names, err := l.amaranthNames(ctx, slices.Collect(maps.Keys(answer.Result.Cards)))
	if err != nil {
		return nil, err
	}
	dk := &deck{name: answer.Result.Title, author: answer.Result.Author, description: answer.Result.Description,
		source: "https://amaranth.vtes.co.nz/#deck/" + id}
	var lines []nameLine
	var errs []string
	for _, k := range slices.Sorted(maps.Keys(answer.Result.Cards)) {
		name, ok := names[k]
		switch {
		case !ok:
			errs = append(errs, fmt.Sprintf("Amaranth's card %s is not in its catalog", k))
		case answer.Result.Cards[k] > 0:
			var cards []*bundle.Card
			if m := reAmaranthAdvanced.FindStringSubmatch(name); m != nil {
				cards = d.ix.lookup(bundle.Fold(m[1] + " (G" + m[2] + " ADV)"))
			} else {
				cards = slices.DeleteFunc(slices.Clone(d.ix.lookup(bundle.Fold(name))), func(c *bundle.Card) bool {
					return bundle.Fold(c.PrintedName) != bundle.Fold(name) || strings.Contains(c.Name, "ADV")
				})
			}
			if len(cards) == 0 {
				errs = append(errs, fmt.Sprintf("Amaranth's card %s, %q, is not in the card data", k, name))
				continue
			}
			lines = append(lines, nameLine{name: name, count: answer.Result.Cards[k], cards: cards})
		}
	}
	errs = append(errs, d.resolveLines(dk, lines, true)...)
	if len(errs) > 0 {
		return nil, fmt.Errorf("some cards of the Amaranth deck %s cannot be read; give the list as text instead:\n- %s", id, strings.Join(errs, "\n- "))
	}
	if len(dk.entries) == 0 {
		return nil, fmt.Errorf("the Amaranth deck %s holds no card", u)
	}
	return dk, nil
}

// amaranthNames answers Amaranth's card names by its ids, its catalog fetched
// on first use and again, at most hourly, when a deck names an id it lacks.
func (l *Live) amaranthNames(ctx context.Context, ids []string) (map[string]string, error) {
	l.amaranth.Lock()
	defer l.amaranth.Unlock()
	if l.amaranth.names != nil && (time.Since(l.amaranth.fetched) < time.Hour ||
		!slices.ContainsFunc(ids, func(id string) bool { _, ok := l.amaranth.names[id]; return !ok })) {
		return l.amaranth.names, nil
	}
	var catalog struct {
		Result []struct {
			ID   int
			Name string
		}
	}
	if err := l.fetch(ctx, "Amaranth", "amaranth.vtes.co.nz", "/api/cards", "card catalog", &catalog); err != nil {
		if l.amaranth.names != nil {
			return l.amaranth.names, nil
		}
		return nil, err
	}
	names := map[string]string{}
	for _, c := range catalog.Result {
		names[strconv.Itoa(c.ID)] = c.Name
	}
	l.amaranth.names, l.amaranth.fetched = names, time.Now()
	return names, nil
}

var (
	// "3x Card", "3 x Card", "x3 Card", "3\tCard", "3 Card"; "Card x3"
	reCountFirst = regexp.MustCompile(`^[-*•\s]*(?:[xX]\s*)?(\d{1,2})(?:\s*[xX×])?\s+(.+)$`)
	reCountLast  = regexp.MustCompile(`^[-*•\s]*(.+?)\s*[xX×]\s*(\d{1,2})$`)
	reLineNote   = regexp.MustCompile(`\s+(?:--+|//+)\s*(.*)$`)
	reColumns    = regexp.MustCompile(`\t+|\s{2,}`)
	// a crypt line's group column: "Tremere:5", "G5"
	reGroupColumn = regexp.MustCompile(`(?i)(?::|\bg)(\d)\s*$`)
	// a section's heading, or a rule of dashes
	reHeader      = regexp.MustCompile(`(?i)^[\W_]*$|^((crypt|library|master|action|political|ally|allies|retainer|equipment|modifier|reaction|combat|event|conviction|power|combo|imbued|vampire)s?[\s/&,]*)+(\(.*\)|:?\s*\d*\s*(cards?)?)$`)
	reNameField   = regexp.MustCompile(`(?i)^\s*(?:deck\s*)?name\s*:\s*(.+)$`)
	reAuthorField = regexp.MustCompile(`(?i)^\s*(?:author|created by|deck by)\s*:?\s*(.+)$`)
)

// readText reads a decklist line by line.
func (d *data) readText(text string) (*deck, error) {
	dk := &deck{}
	var lines []nameLine
	for raw := range strings.Lines(text) {
		raw = strings.TrimSpace(raw)
		if raw == "" {
			continue
		}
		if len(lines) == 0 {
			if m := reNameField.FindStringSubmatch(raw); m != nil {
				dk.name = m[1]
				continue
			}
			if m := reAuthorField.FindStringSubmatch(raw); m != nil {
				dk.author = m[1]
				continue
			}
		}
		comment := ""
		if m := reLineNote.FindStringSubmatchIndex(raw); m != nil {
			comment, raw = raw[m[2]:m[3]], raw[:m[0]]
		}
		var count int
		var rest string
		if m := reCountFirst.FindStringSubmatch(raw); m != nil {
			count, rest = atoi(m[1]), m[2]
		} else if m := reCountLast.FindStringSubmatch(raw); m != nil {
			count, rest = atoi(m[2]), m[1]
		} else {
			switch {
			case d.ix.lookup(bundle.Fold(raw)) != nil:
				dk.notes = append(dk.notes, fmt.Sprintf("%q names a card but gives no count: left out", raw))
			case len(lines) > 0 && !reHeader.MatchString(raw):
				dk.notes = append(dk.notes, fmt.Sprintf("%q reads as no card with its count: left out", raw))
			}
			continue
		}
		name, group := d.cryptColumns(rest)
		lines = append(lines, nameLine{name: name, group: group, count: count, comment: comment})
	}
	if errs := d.resolveLines(dk, lines, false); len(errs) > 0 {
		return nil, errors.New("some lines name no card; correct them and call vtes_deck again:\n- " + strings.Join(errs, "\n- "))
	}
	if len(dk.entries) == 0 {
		return nil, errors.New(`no line of the list reads as a card with its count: give one card per line, e.g. "3x Conditioning" or "2 Anson"`)
	}
	return dk, nil
}

// nameLine is a card of a list by its name, with the group its line gives,
// and the cards the name may be.
type nameLine struct {
	name, group string
	count       int
	comment     string
	cards       []*bundle.Card
}

// resolveLines adds the lines' cards to the deck, resolved as vtes_cards
// resolves them; a vampire's name shared by several groups takes the group of
// the rest of the crypt. It answers what names no card. A deckbuilder's link
// gives each line's cards, which no one can then choose among.
func (d *data) resolveLines(dk *deck, lines []nameLine, link bool) []string {
	var rs []nameLine
	var errs []string
	for _, l := range lines {
		if link {
			rs = append(rs, l)
			continue
		}
		cards, note := d.ix.resolve(l.name, "vtes_deck")
		if l.group != "" {
			if grouped := d.ix.lookup(bundle.Fold(l.name + " (G" + l.group + ")")); grouped != nil {
				cards, note = grouped, ""
			}
		}
		switch {
		case cards == nil:
			errs = append(errs, note)
			continue
		case note != "":
			dk.notes = append(dk.notes, note)
		}
		l.cards = cards
		rs = append(rs, l)
	}
	groups := map[string]bool{}
	for _, r := range rs {
		if len(r.cards) == 1 && r.cards[0].Group != "" {
			groups[r.cards[0].Group] = true
		}
	}
	for _, r := range rs {
		c := r.cards[0]
		if len(r.cards) > 1 {
			fits := slices.DeleteFunc(slices.Clone(r.cards), func(c *bundle.Card) bool { return !groups[c.Group] })
			if len(fits) > 1 {
				fits = slices.DeleteFunc(fits, func(c *bundle.Card) bool { return strings.Contains(c.Name, "ADV") })
			}
			if len(fits) != 1 {
				var names []string
				for _, c := range r.cards {
					names = append(names, c.Name)
				}
				if link {
					errs = append(errs, fmt.Sprintf("%q is several cards (%s), and the rest of the crypt does not say which", r.name, strings.Join(names, "; ")))
				} else {
					errs = append(errs, fmt.Sprintf("%q is several cards: %s. Give the one meant, e.g. %q.", r.name, strings.Join(names, "; "), names[0]))
				}
				continue
			}
			c = fits[0]
			dk.notes = append(dk.notes, fmt.Sprintf("%q read as %s, of the group of the rest of the crypt (several vampires bear the name).", r.name, c.Name))
		}
		if i := slices.IndexFunc(dk.entries, func(e deckEntry) bool { return e.card == c }); i >= 0 {
			dk.entries[i].count += r.count
		} else {
			dk.entries = append(dk.entries, deckEntry{c, r.count, r.comment})
		}
	}
	return errs
}

// cryptColumns cuts a crypt line's columns after the name (capacity,
// disciplines, title, clan and group) and answers the name and the group.
func (d *data) cryptColumns(rest string) (string, string) {
	group := ""
	if m := reGroupColumn.FindStringSubmatch(rest); m != nil {
		group = m[1]
	}
	if cols := reColumns.Split(strings.TrimSpace(rest), -1); len(cols) > 1 {
		return cols[0], group
	}
	if d.ix.lookup(bundle.Fold(rest)) != nil {
		return rest, ""
	}
	// single spaces: the longest opening words naming a card, a capacity after them
	words := strings.Fields(rest)
	for i := len(words) - 1; i > 0; i-- {
		if _, err := strconv.Atoi(words[i]); err == nil && d.ix.lookup(bundle.Fold(strings.Join(words[:i], " "))) != nil {
			return strings.Join(words[:i], " "), group
		}
	}
	return rest, ""
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func (dk *deck) split() (crypt, library []deckEntry, nCrypt, nLibrary int) {
	for _, e := range dk.entries {
		if e.card.IsCrypt() {
			crypt, nCrypt = append(crypt, e), nCrypt+e.count
		} else {
			library, nLibrary = append(library, e), nLibrary+e.count
		}
	}
	slices.SortStableFunc(crypt, func(a, b deckEntry) int {
		return cmp.Or(cmp.Compare(b.count, a.count), cmp.Compare(b.card.Capacity, a.card.Capacity), strings.Compare(a.card.Name, b.card.Name))
	})
	return crypt, library, nCrypt, nLibrary
}

// fixes names what keeps the list from tournament play today, per the
// rulebook's deck construction rules and the card data's dates.
func (dk *deck) fixes(today string) []string {
	crypt, _, nCrypt, nLibrary := dk.split()
	var out []string
	if nCrypt < 12 {
		out = append(out, fmt.Sprintf("the crypt has %d cards: at least 12 (add %d)", nCrypt, 12-nCrypt))
	}
	if nLibrary < 60 {
		out = append(out, fmt.Sprintf("the library has %d cards: 60 to 90 (add %d)", nLibrary, 60-nLibrary))
	} else if nLibrary > 90 {
		out = append(out, fmt.Sprintf("the library has %d cards: 60 to 90 (cut %d)", nLibrary, nLibrary-90))
	}
	var groups []int
	for _, e := range crypt {
		if g, err := strconv.Atoi(strings.TrimPrefix(e.card.Group, "G")); err == nil && !slices.Contains(groups, g) {
			groups = append(groups, g)
		}
	}
	if slices.Sort(groups); len(groups) > 0 && groups[len(groups)-1]-groups[0] > 1 {
		out = append(out, fmt.Sprintf("the crypt mixes groups %s: a crypt is of one group or two consecutive ones (a vampire of group Any fits either)", joinInts(groups)))
	}
	for _, e := range dk.entries {
		switch {
		case e.card.Banned != "" && e.card.Banned <= today:
			out = append(out, fmt.Sprintf("%s is banned since %s: take it out", e.card.Name, e.card.Banned))
		case e.card.Legal > today:
			out = append(out, fmt.Sprintf("%s is tournament legal only from %s", e.card.Name, e.card.Legal))
		}
	}
	return out
}

func joinInts(ns []int) string {
	s := make([]string, len(ns))
	for i, n := range ns {
		s[i] = strconv.Itoa(n)
	}
	return strings.Join(s, ", ")
}

const commentCut = 1500

// How much of the dossier is written, to fit the result's bound.
const (
	full      = iota
	lean      // the relatives' events, comments and differences left out
	noLibrary // the library's texts left out
	bare      // every text and the crypt's rosters left out
)

func (dk *deck) write(out *strings.Builder, d *data, today string, fixes []string, detail bool, fit int) {
	b := d.b
	crypt, library, nCrypt, nLibrary := dk.split()
	fmt.Fprintf(out, "# Deck: %s\n", cmp.Or(dk.name, "(unnamed)"))
	if t := dk.twda; t != nil {
		fmt.Fprintf(out, "TWDA deck %s: won by %s at %s, %s, %s, %d players (%s).", t.ID, t.Player, t.Event, t.Place, t.Date, t.Players, t.Format)
		if t.URL != "" {
			fmt.Fprintf(out, " Event report: %s", t.URL)
		}
		out.WriteString("\n")
		if t.Author != "" && t.Author != t.Player {
			fmt.Fprintf(out, "Deck by %s.\n", t.Author)
		}
	} else if dk.author != "" {
		fmt.Fprintf(out, "By %s.\n", dk.author)
	}
	if dk.source != "" {
		fmt.Fprintf(out, "Read from %s\n", dk.source)
	}
	comment := dk.description
	if dk.twda != nil {
		comment = dk.twda.Comment
	}
	if strings.TrimSpace(comment) != "" {
		out.WriteString("\n")
		cut := commentCut
		if detail {
			cut = math.MaxInt
		}
		writeComment(out, comment, cut, "detail gives it whole")
	}
	if len(dk.notes) > 0 {
		out.WriteString("\nHow the list was read:\n")
		for _, n := range dk.notes {
			fmt.Fprintf(out, "- %s\n", n)
		}
	}
	if len(fixes) == 0 {
		fmt.Fprintf(out, "\nTournament legal today (%s): yes.\n", today)
	} else {
		fmt.Fprintf(out, "\nTournament legal today (%s): no, as archived:\n", today)
		for _, f := range fixes {
			fmt.Fprintf(out, "- %s\n", f)
		}
	}

	if fit >= noLibrary {
		fmt.Fprintf(out, "\nLeft out to fit this result: the texts of the %s. Call vtes_cards with the names, 10 at a time.\n",
			map[int]string{noLibrary: "library's cards", bare: "crypt's and the library's cards, and each discipline's vampires"}[fit])
	}
	writeCrypt(out, crypt, nCrypt, fit)
	writeLibrary(out, library, crypt, nLibrary, nCrypt, fit)
	writeModules(out, b.Modules, library, crypt, nLibrary, fit)
	dk.writeRelatives(out, d, detail, fit)

	out.WriteString("## How to answer\n\nThese facts are computed from the card data; the texts above are official. " +
		"Rulings are not in this dossier: call vtes_cards for a card's rulings and history, vtes_rules_case for a rules point. " +
		"The modules say which cards belong together, not how this deck plays them. " +
		"Primogen does not serve theory.md and modules.md, which the procedures name: what you say of them beyond these facts is your own reading, " +
		"to be said so.\n\n")
	out.WriteString(strings.TrimSpace(b.DeckAnalysis))
	out.WriteString("\n\n")
	out.WriteString(strings.TrimSpace(b.DeckAdvice))
	out.WriteString("\n\nAnswer from these facts and texts, never from memory of the cards. " +
		"Where neither a text nor a ruling settles a point, say it is unsettled rather than filling the gap.\n\n")
}

func writeCrypt(out *strings.Builder, crypt []deckEntry, n, fit int) {
	capacity, lo, hi := 0, 0, 0
	var groups []string
	for i, e := range crypt {
		c := e.card.Capacity
		capacity += c * e.count
		if i == 0 || c < lo {
			lo = c
		}
		hi = max(hi, c)
		if !slices.Contains(groups, e.card.Group) {
			groups = append(groups, e.card.Group)
		}
	}
	slices.Sort(groups)
	fmt.Fprintf(out, "\n## Crypt (%d cards, %d distinct", n, len(crypt))
	if n > 0 {
		fmt.Fprintf(out, "; capacity %d to %d, average %.1f; groups %s", lo, hi, float64(capacity)/float64(n), strings.Join(groups, ", "))
	}
	out.WriteString(")\n\n| Copies | Name | Card | Capacity | Group | Clan | Disciplines (capitals: superior) | Title |\n|---|---|---|---|---|---|---|---|\n")
	for _, e := range crypt {
		c := e.card
		fmt.Fprintf(out, "| %d | %s | %d | %d | %s | %s | %s | %s |\n", e.count, c.Name, c.ID, c.Capacity, c.Group,
			strings.Join(c.Clans, ", "), strings.Join(cryptDisciplines(c), " "), c.Title)
	}
	type roster struct {
		copies   [2]int
		superior []string
		inferior []string
	}
	disciplines := map[string]*roster{}
	for _, e := range crypt {
		for _, d := range cryptDisciplines(e.card) {
			k := strings.ToLower(d)
			if disciplines[k] == nil {
				disciplines[k] = &roster{}
			}
			r := disciplines[k]
			if d == strings.ToUpper(d) {
				r.copies[0] += e.count
				r.superior = append(r.superior, copies(e))
			} else {
				r.copies[1] += e.count
				r.inferior = append(r.inferior, copies(e))
			}
		}
	}
	keys := slices.Collect(maps.Keys(disciplines))
	slices.SortFunc(keys, func(a, b string) int {
		ra, rb := disciplines[a].copies, disciplines[b].copies
		return cmp.Or(cmp.Compare(rb[0]+rb[1], ra[0]+ra[1]), cmp.Compare(rb[0], ra[0]), strings.Compare(a, b))
	})
	out.WriteString("\nEach discipline, by crypt cards:\n")
	for _, k := range keys {
		r := disciplines[k]
		fmt.Fprintf(out, "- [%s] %d: superior %d", k, r.copies[0]+r.copies[1], r.copies[0])
		if fit < bare && len(r.superior) > 0 {
			fmt.Fprintf(out, " (%s)", strings.Join(r.superior, ", "))
		}
		fmt.Fprintf(out, "; inferior %d", r.copies[1])
		if fit < bare && len(r.inferior) > 0 {
			fmt.Fprintf(out, " (%s)", strings.Join(r.inferior, ", "))
		}
		out.WriteString("\n")
	}
	if fit < bare {
		out.WriteString("\nTexts (official):\n")
	}
	for _, e := range crypt {
		if fit < bare {
			fmt.Fprintf(out, "- **%s**:\n", e.card.Name)
			quote(out, e.card.Text)
		}
		if e.comment != "" {
			fmt.Fprintf(out, "  The author, on %s: %s\n", e.card.Name, e.comment)
		}
	}
	out.WriteString("\n")
}

// cryptDisciplines are a crypt card's disciplines, Flight included, which
// the card data gives only in the text.
func cryptDisciplines(c *bundle.Card) []string {
	if strings.Contains(c.Text, "[FLIGHT]") {
		return append(slices.Clone(c.Disciplines), "FLIGHT")
	}
	return c.Disciplines
}

var typeOrder = []string{"Master", "Conviction", "Action", "Political Action", "Ally", "Retainer", "Equipment", "Power",
	"Action Modifier", "Reaction", "Combat", "Event"}

func writeLibrary(out *strings.Builder, library, crypt []deckEntry, n, nCrypt, fit int) {
	fmt.Fprintf(out, "## Library (%d cards, %d distinct)\n\n", n, len(library))
	costs := map[string][2]int{} // kind → amount, cards
	var variable []string
	for _, e := range library {
		for _, cost := range [][2]string{{e.card.PoolCost, "pool"}, {e.card.BloodCost, "blood"}, {e.card.ConvictionCost, "conviction"}} {
			if cost[0] == "" {
				continue
			}
			if v, err := strconv.Atoi(cost[0]); err == nil {
				c := costs[cost[1]]
				costs[cost[1]] = [2]int{c[0] + v*e.count, c[1] + e.count}
			} else {
				variable = append(variable, fmt.Sprintf("%s (%s %s)", e.card.Name, cost[0], cost[1]))
			}
		}
	}
	var parts []string
	for _, kind := range []string{"pool", "blood", "conviction"} {
		if c, ok := costs[kind]; ok {
			parts = append(parts, fmt.Sprintf("%d %s over %d cards", c[0], kind, c[1]))
		}
	}
	if len(variable) > 0 {
		parts = append(parts, "variable: "+strings.Join(variable, ", "))
	}
	if len(parts) > 0 {
		fmt.Fprintf(out, "Costs, every copy counted: %s.\n\n", strings.Join(parts, "; "))
	}

	byType := map[string][]deckEntry{}
	for _, e := range library {
		t := strings.Join(e.card.Types, "/")
		byType[t] = append(byType[t], e)
	}
	types := slices.Collect(maps.Keys(byType))
	rank := func(t string) int {
		first, _, _ := strings.Cut(t, "/")
		if i := slices.Index(typeOrder, first); i >= 0 {
			return i
		}
		return len(typeOrder)
	}
	slices.SortFunc(types, func(a, b string) int { return cmp.Or(cmp.Compare(rank(a), rank(b)), strings.Compare(a, b)) })
	for _, t := range types {
		entries := byType[t]
		total := 0
		for _, e := range entries {
			total += e.count
		}
		slices.SortStableFunc(entries, func(a, b deckEntry) int {
			return cmp.Or(cmp.Compare(b.count, a.count), strings.Compare(a.card.Name, b.card.Name))
		})
		fmt.Fprintf(out, "### %s (%d)\n\n", t, total)
		for _, e := range entries {
			fmt.Fprintf(out, "- %d× **%s** (card %d)", e.count, e.card.Name, e.card.ID)
			if reqs := requirements(e.card); len(reqs) > 0 {
				fmt.Fprintf(out, ". %s", strings.Join(reqs, " "))
			}
			out.WriteString("\n")
			who := playable(e.card, crypt, nCrypt)
			if who != "" {
				fmt.Fprintf(out, "  %s\n", who)
			}
			if fit < noLibrary {
				quote(out, e.card.Text)
			} else if m := reTextRequires.FindString(e.card.Text); m != "" && who == "" {
				fmt.Fprintf(out, "  Its text requires: %q\n", m)
			}
			if e.comment != "" {
				fmt.Fprintf(out, "  The author: %s\n", e.comment)
			}
		}
		out.WriteString("\n")
	}
}

// writeModules groups the library by the catalog's modules.
func writeModules(out *strings.Builder, modules []bundle.Module, library, crypt []deckEntry, n, fit int) {
	out.WriteString("## Library by module\n\n")
	if len(modules) == 0 {
		out.WriteString("This data bundle holds no module tags: no modules until it is rebuilt.\n\n")
		return
	}
	fmt.Fprintf(out, "The modules of primogen's catalog (knowledge/guide/modules.md, curated by the owner) of which this deck holds at least one core card "+
		"(a star engine naming several: two of its cards), a card counted in each module it belongs to. Density: the module's copies over the library's %d cards, in theory.md's classes "+
		"(%s); per hand: the copies a hand of 7 holds on average, density × 7.\n\n", n, bundle.DensityClasses)
	if fit > full {
		out.WriteString("What each module lacks of its core is left out to fit this result.\n\n")
	}
	held := map[int]deckEntry{}
	for _, e := range slices.Concat(library, crypt) {
		held[e.card.ID] = e
	}
	type shown struct {
		line   string
		copies int
	}
	var lines []shown
	var single []string // the modules with one library card or none
	tagged := map[int]bool{}
	for _, m := range modules {
		var parts [2][]string // core, support
		var inCrypt, absent []string
		var ids []int
		copiesHere, coreHeld, pieces := 0, 0, 0
		for i, tags := range [][]bundle.Tag{m.Core, m.Support} {
			for _, t := range tags {
				found := false
				for _, id := range t.IDs {
					e, ok := held[id]
					if !ok {
						continue
					}
					found = true
					if slices.Contains(e.card.Types, "Vampire") || slices.Contains(e.card.Types, "Imbued") {
						inCrypt = append(inCrypt, copies(e))
						continue
					}
					ids = append(ids, id)
					copiesHere += e.count
					parts[i] = append(parts[i], copies(e))
				}
				switch {
				case found:
					pieces++
					if i == 0 {
						coreHeld++
					}
				case i == 0 && fit == full:
					absent = append(absent, t.Name)
				}
			}
		}
		if coreHeld == 0 || pieces < m.Pieces {
			continue
		}
		for _, id := range ids {
			tagged[id] = true
		}
		section, _, _ := strings.Cut(m.Section, " (")
		if copiesHere <= 1 {
			single = append(single, fmt.Sprintf("**%s** (%s)", m.Name, strings.Join(slices.Concat(parts[0], parts[1], inCrypt), ", ")))
			continue
		}
		density := float64(copiesHere) / float64(n)
		var line strings.Builder
		fmt.Fprintf(&line, "- **%s** (%s): %d copies, %.1f%%, %s, %.1f per hand", m.Name, section, copiesHere, 100*density, densityClass(density), 7*density)
		for _, p := range []struct {
			head  string
			cards []string
		}{{"core", parts[0]}, {"support", parts[1]}, {"in the crypt", inCrypt}, {"not here, of its core", absent}} {
			if len(p.cards) > 0 {
				fmt.Fprintf(&line, "; %s: %s", p.head, strings.Join(p.cards, ", "))
			}
		}
		line.WriteString("\n")
		lines = append(lines, shown{line.String(), copiesHere})
	}
	slices.SortStableFunc(lines, func(a, b shown) int { return cmp.Compare(b.copies, a.copies) })
	if len(lines)+len(single) == 0 {
		out.WriteString("No module of the catalog has a core card in this deck.\n")
	}
	for _, l := range lines {
		out.WriteString(l.line)
	}
	if len(single) > 0 {
		fmt.Fprintf(out, "- With one library copy or none: %s\n", strings.Join(single, "; "))
	}
	var none []string
	count := 0
	for _, e := range slices.SortedStableFunc(slices.Values(library), func(a, b deckEntry) int {
		return cmp.Or(cmp.Compare(b.count, a.count), strings.Compare(a.card.Name, b.card.Name))
	}) {
		if !tagged[e.card.ID] {
			none = append(none, copies(e))
			count += e.count
		}
	}
	if len(none) > 0 {
		fmt.Fprintf(out, "\nIn no module shown (%d copies, %.1f%% of the library): %s\n", count, 100*float64(count)/float64(n), strings.Join(none, ", "))
	}
	out.WriteString("\n")
}

// densityClass is theory.md's class of a module's density.
func densityClass(d float64) string {
	switch {
	case d > 0.25:
		return "spam"
	case d >= 0.14:
		return "strong"
	case d >= 0.07:
		return "standard"
	}
	return "tactical"
}

var reTextRequires = regexp.MustCompile(`Requires [^.]*\.`)

// playable says which crypt cards have what a library card requires as the
// card data holds it (disciplines, clans, path), at which level; empty when
// it holds none. A requirement only the text states is quoted, not checked.
func playable(c *bundle.Card, crypt []deckEntry, nCrypt int) string {
	if len(c.Disciplines) == 0 && len(c.Clans) == 0 && c.Path == "" {
		return ""
	}
	var needs []string
	if len(c.Disciplines) > 0 {
		ds := make([]string, len(c.Disciplines))
		for i, d := range c.Disciplines {
			ds[i] = "[" + d + "]"
		}
		needs = append(needs, strings.Join(ds, " or "))
	}
	if len(c.Clans) > 0 {
		needs = append(needs, "clan "+strings.Join(c.Clans, " or "))
	}
	if c.Path != "" {
		needs = append(needs, "path "+c.Path)
	}
	head := "Crypt cards with the " + strings.Join(needs, " and ") + " it requires"
	if len(c.Disciplines) > 1 {
		head += " (each effect's symbols in the text say which)"
	}
	var able []string
	n, superior := 0, 0
	for _, e := range crypt {
		v := e.card
		if len(c.Clans) > 0 && !slices.ContainsFunc(v.Clans, func(cl string) bool { return slices.Contains(c.Clans, cl) }) {
			continue
		}
		if c.Path != "" && (v.Path == "" || !strings.Contains(c.Path, v.Path)) {
			continue
		}
		var levels []string
		for _, req := range c.Disciplines {
			for _, d := range cryptDisciplines(v) {
				if strings.EqualFold(d, req) {
					levels = append(levels, d)
				}
			}
		}
		if len(c.Disciplines) > 0 && len(levels) == 0 {
			continue
		}
		name := copies(e)
		if len(levels) > 0 {
			name += " [" + strings.Join(levels, " ") + "]"
		}
		if len(levels) == 1 && levels[0] == strings.ToUpper(levels[0]) {
			superior += e.count
		}
		able = append(able, name)
		n += e.count
	}
	var line string
	switch {
	case n == 0:
		line = head + ": none."
	case len(c.Disciplines) == 1 && len(c.Clans) == 0 && c.Path == "":
		line = fmt.Sprintf("%s: %d of %d, %d at superior (the crypt's [%s] line names them).", head, n, nCrypt, superior, strings.ToLower(c.Disciplines[0]))
	case len(c.Disciplines) == 0 && n == nCrypt:
		line = fmt.Sprintf("%s: all %d.", head, nCrypt)
	default:
		line = fmt.Sprintf("%s: %d of %d: %s.", head, n, nCrypt, strings.Join(able, "; "))
	}
	if m := reTextRequires.FindString(c.Text); m != "" {
		line += fmt.Sprintf(" Its text also requires, unchecked here: %q", m)
	}
	return line
}

func copies(e deckEntry) string {
	if e.count > 1 {
		return fmt.Sprintf("%s ×%d", e.card.Name, e.count)
	}
	return e.card.Name
}

// quote writes a card's text verbatim, as vtes_cards does.
func quote(out *strings.Builder, text string) {
	for line := range strings.Lines(strings.TrimSpace(text)) {
		fmt.Fprintf(out, "  > %s", line)
	}
	out.WriteString("\n")
}

const (
	relativesShown      = 3
	relativeCommentCut  = 600
	relativeCommentWide = 1500
	differencesShown    = 12 // of each kind, the most copies first
)

// writeRelatives places the deck among the archetypes of the labels and the
// TWDA decks of their window.
func (dk *deck) writeRelatives(out *strings.Builder, d *data, detail bool, fit int) {
	out.WriteString("## Relatives\n\n")
	r := d.rel
	if r == nil {
		out.WriteString("This data bundle holds no archetype labels: no relatives until it is rebuilt.\n\n")
		return
	}
	c := r.Class
	var cards []bundle.DeckCard
	for _, e := range dk.entries {
		cards = append(cards, bundle.DeckCard{ID: e.card.ID, Count: e.count})
	}
	id := ""
	if dk.twda != nil {
		id = dk.twda.ID
	}
	fmt.Fprintf(out, "Similarity from card counts against the %d TWDA decks since %s, rare cards weighing more, the crypt %.0f%% and the library %.0f%%: "+
		"1 for the same list, 0 for no card in common. Archetypes: codex-of-the-damned's labels (codex-of-the-damned.org), by its maintainer, "+
		"pass of %s, curated and revisable; each of their decks left out in turn, %.0f%% are placed back in their archetype, %.0f%% in its family (variants merged).\n\n",
		len(r.Decks), c.Since, 100*c.CryptWeight, 100*(1-c.CryptWeight), c.Generated, 100*r.GroupRate, 100*r.RootRate)
	v := r.Vector(cards)
	if v.Empty() {
		fmt.Fprintf(out, "No card of this deck is played by %d or more of those decks: it has no relatives there.\n\n", c.MinDF)
		return
	}
	if a := r.Label(id); a != nil {
		fmt.Fprintf(out, "This deck is labelled %s.\n\n", archetypeName(r, a))
	}
	matches := r.Archetypes(v, id, relativesShown)
	fmt.Fprintf(out, "Nearest archetypes (95 in 100 of an archetype's own decks are at %.2f or more from it):\n", r.Threshold)
	for _, m := range matches {
		a := &c.Archetypes[m.Index]
		fmt.Fprintf(out, "- %s: %.2f; %d decks", archetypeName(r, a), m.Similarity, r.Size(m.Index))
		if len(a.Representative) > 0 {
			fmt.Fprintf(out, ", its representative decks %s", strings.Join(a.Representative, ", "))
		}
		out.WriteString(".\n")
	}
	if matches[0].Similarity < r.Threshold {
		fmt.Fprintf(out, "\nThis deck is near no archetype: even the closest is under %.2f. Read it on its own mechanics; "+
			"the decks below are its nearest lists, not its archetype.\n", r.Threshold)
	}
	out.WriteString("\nNearest tournament decks (call vtes_deck with an id for the whole list):\n\n")
	if fit > full {
		out.WriteString("Their events, their authors' comments and the cards that differ from this deck are left out to fit this result.\n")
	}
	here := map[int]int{}
	for _, e := range dk.entries {
		here[e.card.ID] += e.count
	}
	for _, m := range r.Nearest(v, id, relativesShown) {
		t := r.Decks[m.Index]
		if fit > full {
			fmt.Fprintf(out, "- %s, %s: %.2f", t.ID, cmp.Or(t.Name, "(unnamed)"), m.Similarity)
			if a := r.Label(t.ID); a != nil {
				fmt.Fprintf(out, "; labelled %s", archetypeName(r, a))
			}
			out.WriteString(".\n")
			continue
		}
		fmt.Fprintf(out, "### %s, %s: %.2f\n\nWon by %s at %s, %s, %s, %d players.", t.ID, cmp.Or(t.Name, "(unnamed)"), m.Similarity, t.Player, t.Event, t.Place, t.Date, t.Players)
		if t.URL != "" {
			fmt.Fprintf(out, " Event report: %s", t.URL)
		}
		if a := r.Label(t.ID); a != nil {
			fmt.Fprintf(out, " Labelled %s.", archetypeName(r, a))
		}
		out.WriteString("\n")
		cut, whole := relativeCommentCut, "vtes_deck "+t.ID+" gives it"
		if detail {
			cut = relativeCommentWide
		}
		writeComment(out, t.Comment, cut, whole)
		out.WriteString(d.differences(here, t))
		out.WriteString("\n")
	}
	if fit > full {
		out.WriteString("\n")
	}
}

// writeComment quotes an author's comment, cut at a word past cut runes,
// saying where to read it whole.
func writeComment(out *strings.Builder, comment string, cut int, whole string) {
	if comment = strings.TrimSpace(comment); comment == "" {
		return
	}
	if r := []rune(comment); len(r) > cut {
		kept := string(r[:cut])
		comment = kept[:max(strings.LastIndexAny(kept, " \n"), 0)] + fmt.Sprintf(" […] (cut at %d of %d characters: %s)", cut, len(r), whole)
	}
	out.WriteString("The author's comment:\n")
	for line := range strings.Lines(comment) {
		fmt.Fprintf(out, "> %s", line)
	}
	out.WriteString("\n")
}

func archetypeName(r *bundle.Relatives, a *bundle.Archetype) string {
	if root := r.Root(a); root != a {
		return fmt.Sprintf("**%s** (%s), a variant of **%s**", a.Name, a.Ref, root.Name)
	}
	return fmt.Sprintf("**%s** (%s)", a.Name, a.Ref)
}

// differences names the cards a relative plays otherwise than the deck.
func (d *data) differences(here map[int]int, t *bundle.Deck) string {
	there := map[int]int{}
	for _, c := range t.Cards {
		there[c.ID] += c.Count
	}
	var onlyHere, onlyThere, counts []string
	ids := slices.Collect(maps.Keys(here))
	for id := range there {
		if _, ok := here[id]; !ok {
			ids = append(ids, id)
		}
	}
	slices.SortFunc(ids, func(a, b int) int {
		return cmp.Or(cmp.Compare(max(here[b], there[b]), max(here[a], there[a])), strings.Compare(d.byID[a].Name, d.byID[b].Name))
	})
	for _, id := range ids {
		name := d.byID[id].Name
		switch h, th := here[id], there[id]; {
		case th == 0:
			onlyHere = append(onlyHere, fmt.Sprintf("%d× %s", h, name))
		case h == 0:
			onlyThere = append(onlyThere, fmt.Sprintf("%d× %s", th, name))
		case h != th:
			counts = append(counts, fmt.Sprintf("%s, %d here and %d there", name, h, th))
		}
	}
	if len(onlyHere)+len(onlyThere)+len(counts) == 0 {
		return "The same list as this deck.\n"
	}
	var out strings.Builder
	for _, part := range []struct {
		head  string
		cards []string
	}{{"Only in this deck", onlyHere}, {"Only in " + t.ID, onlyThere}, {"Other counts", counts}} {
		if n := len(part.cards); n > differencesShown {
			part.cards = append(part.cards[:differencesShown], fmt.Sprintf("and %d more", n-differencesShown))
		}
		if len(part.cards) > 0 {
			fmt.Fprintf(&out, "%s: %s.\n", part.head, strings.Join(part.cards, "; "))
		}
	}
	return out.String()
}
