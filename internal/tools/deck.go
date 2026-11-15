package tools

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
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

// VDB is the VDB API's route for a stored deck, its id appended.
const VDB = "https://vdb.im/api/deck/"

const maxDeckInput = 20000

type deckInput struct {
	Deck   string `json:"deck"`
	Detail bool   `json:"detail,omitempty"`
}

var deckSchema = &jsonschema.Schema{
	Type: "object",
	Properties: map[string]*jsonschema.Schema{
		"deck": {Type: "string", Description: `the deck: a decklist as text, one card per line with its count ("3x Conditioning", ` +
			`"Anson x2", crypt lines with their capacity and disciplines accepted), a TWDA deck id ("2k6munduko"), or a VDB link (https://vdb.im/decks/...)`},
		"detail": {Type: "boolean", Description: "the deck author's whole comment, where the default cuts it, and a larger result, keeping texts a large list's default drops"},
	},
	Required: []string{"deck"},
}

func addDeck(s *mcp.Server, l *Live) {
	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_deck",
		Title: "VTES deck",
		Description: "Call first to analyze or improve a Vampire: The Eternal Struggle (VTES) deck: never read a decklist from memory " +
			"of its cards. Takes a decklist as text, a TWDA deck id or a VDB link; validates it (sizes, crypt groups, banned cards, " +
			"legality today) and returns the computed facts: the crypt table with capacities, disciplines, titles and texts, " +
			"which crypt cards have what each library card requires and at which level, the library by type with every card's official text, " +
			"the costs, and the procedure a deck analysis and advice follow.",
		InputSchema: deckSchema,
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: new(false)},
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in deckInput) (*mcp.CallToolResult, any, error) {
		text := strings.TrimSpace(in.Deck)
		if text == "" || len(text) > maxDeckInput {
			return nil, nil, fmt.Errorf(`give one deck of at most %d characters: a decklist as text, a TWDA id or a VDB link, e.g. {"deck": "2k6munduko"}`, maxDeckInput)
		}
		d := l.p.Load()
		dk, err := d.readDeck(ctx, cmp.Or(l.VDB, VDB), text)
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
			if dk.write(&out, d.b, today, fixes, in.Detail, fit); out.Len()+len(DataLine(d.b)) <= limit {
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

func (d *data) readDeck(ctx context.Context, vdb, text string) (*deck, error) {
	if u, err := url.Parse(text); err == nil && strings.TrimPrefix(u.Host, "www.") == "vdb.im" {
		return d.readVDB(ctx, vdb, u)
	}
	if !strings.ContainsAny(text, " \t\n") {
		if t := d.decks[text]; t != nil {
			return d.fromTWDA(t), nil
		}
		return nil, fmt.Errorf("no deck of the TWDA has the id %q: give a TWDA id (e.g. 2k6munduko), a VDB link, or the list as text, one card per line with its count", text)
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

func (d *data) readVDB(ctx context.Context, api string, u *url.URL) (*deck, error) {
	id := strings.TrimPrefix(u.Path, "/decks/")
	// VDB names a TWDA deck of a nine-character id with a trailing 0
	twda := d.decks[id]
	if len(id) == 10 && twda == nil {
		twda = d.decks[strings.TrimSuffix(id, "0")]
	}
	var vdb struct {
		Name, Author, Description string
		Cards                     map[string]int
	}
	switch {
	case u.Path == "/decks/deck":
		vdb.Name, vdb.Author, vdb.Description = u.Query().Get("name"), u.Query().Get("author"), u.Query().Get("description")
		vdb.Cards = map[string]int{}
		for pair := range strings.SplitSeq(u.Fragment, ";") {
			k, v, _ := strings.Cut(pair, "=")
			n, err := strconv.Atoi(v)
			if k == "" || err != nil {
				return nil, fmt.Errorf("the VDB link's cards, after #, are not id=count pairs separated by ';': %q", u.Fragment)
			}
			vdb.Cards[k] += n
		}
	case id == "" || strings.Contains(id, "/"):
		return nil, fmt.Errorf("not a VDB deck link: %s; give https://vdb.im/decks/<id> or the link VDB's \"Copy URL\" makes", u)
	case twda != nil:
		dk := d.fromTWDA(twda)
		dk.source = u.String()
		return dk, nil
	default:
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, api+url.PathEscape(id), nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", "primogen")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("VDB did not answer for %s: %v; give the list as text instead", u, err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("VDB has no public deck %q (%s): check the link, or give the list as text", id, resp.Status)
		}
		body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(body, &vdb); err != nil {
			return nil, fmt.Errorf("VDB's answer for %s: %w", u, err)
		}
	}
	source := *u
	source.Fragment = ""
	dk := &deck{name: vdb.Name, author: vdb.Author, description: vdb.Description, source: source.String()}
	var unknown []string
	for _, k := range slices.Sorted(maps.Keys(vdb.Cards)) {
		id, _ := strconv.Atoi(k)
		c := d.byID[id]
		if c == nil {
			unknown = append(unknown, k)
			continue
		}
		dk.entries = append(dk.entries, deckEntry{card: c, count: vdb.Cards[k]})
	}
	if len(unknown) > 0 {
		return nil, fmt.Errorf("the VDB deck names cards the card data does not hold (ids %s): give the list as text", strings.Join(unknown, ", "))
	}
	if len(dk.entries) == 0 {
		return nil, fmt.Errorf("the VDB deck %s holds no card", u)
	}
	return dk, nil
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

// readText reads a decklist line by line, resolving each name as vtes_cards
// does; a vampire's name shared by several groups takes the group of the rest
// of the crypt.
func (d *data) readText(text string) (*deck, error) {
	dk := &deck{}
	type line struct {
		name    string
		cards   []*bundle.Card
		count   int
		comment string
	}
	var lines []line
	var errs []string
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
		cards, note := d.ix.resolve(name, "vtes_deck")
		if group != "" {
			if grouped := d.ix.lookup(bundle.Fold(name + " (G" + group + ")")); grouped != nil {
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
		lines = append(lines, line{name, cards, count, comment})
	}
	groups := map[string]bool{}
	for _, l := range lines {
		if len(l.cards) == 1 && l.cards[0].Group != "" {
			groups[l.cards[0].Group] = true
		}
	}
	for _, l := range lines {
		c := l.cards[0]
		if len(l.cards) > 1 {
			fits := slices.DeleteFunc(slices.Clone(l.cards), func(c *bundle.Card) bool { return !groups[c.Group] })
			if len(fits) > 1 {
				fits = slices.DeleteFunc(fits, func(c *bundle.Card) bool { return strings.Contains(c.Name, "ADV") })
			}
			if len(fits) != 1 {
				var names []string
				for _, c := range l.cards {
					names = append(names, c.Name)
				}
				errs = append(errs, fmt.Sprintf("%q is several cards: %s. Give the one meant, e.g. %q.", l.name, strings.Join(names, "; "), names[0]))
				continue
			}
			c = fits[0]
			dk.notes = append(dk.notes, fmt.Sprintf("%q read as %s, of the group of the rest of the crypt (several vampires bear the name).", l.name, c.Name))
		}
		if i := slices.IndexFunc(dk.entries, func(e deckEntry) bool { return e.card == c }); i >= 0 {
			dk.entries[i].count += l.count
		} else {
			dk.entries = append(dk.entries, deckEntry{c, l.count, l.comment})
		}
	}
	if len(errs) > 0 {
		return nil, errors.New("some lines name no card; correct them and call vtes_deck again:\n- " + strings.Join(errs, "\n- "))
	}
	if len(dk.entries) == 0 {
		return nil, errors.New(`no line of the list reads as a card with its count: give one card per line, e.g. "3x Conditioning" or "2 Anson"`)
	}
	return dk, nil
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

func isCrypt(c *bundle.Card) bool {
	return slices.Contains(c.Types, "Vampire") || slices.Contains(c.Types, "Imbued")
}

func (dk *deck) split() (crypt, library []deckEntry, nCrypt, nLibrary int) {
	for _, e := range dk.entries {
		if isCrypt(e.card) {
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
	noLibrary // the library's texts left out
	bare      // every text and the crypt's rosters left out
)

func (dk *deck) write(out *strings.Builder, b *bundle.Bundle, today string, fixes []string, detail bool, fit int) {
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
	if comment = strings.TrimSpace(comment); comment != "" {
		out.WriteString("\nThe author's comment:\n")
		if r := []rune(comment); !detail && len(r) > commentCut {
			kept := string(r[:commentCut])
			comment = kept[:max(strings.LastIndexAny(kept, " \n"), 0)] + fmt.Sprintf(" […] (cut at %d of %d characters: detail gives it whole)", commentCut, len(r))
		}
		for line := range strings.Lines(comment) {
			fmt.Fprintf(out, "> %s", line)
		}
		out.WriteString("\n")
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

	if fit > full {
		fmt.Fprintf(out, "\nLeft out to fit this result: the texts of the %s. Call vtes_cards with the names, 10 at a time.\n",
			map[int]string{noLibrary: "library's cards", bare: "crypt's and the library's cards, and each discipline's vampires"}[fit])
	}
	writeCrypt(out, crypt, nCrypt, fit)
	writeLibrary(out, library, crypt, nLibrary, nCrypt, fit)

	out.WriteString("## How to answer\n\nThese facts are computed from the card data; the texts above are official. " +
		"Rulings are not in this dossier: call vtes_cards for a card's rulings and history, vtes_rules_case for a rules point. " +
		"This dossier computes no modules, no hand expectation and no relatives (nearest archetypes and tournament decks), " +
		"and primogen does not serve theory.md and modules.md, which the procedures name: what you say of them is your own reading, " +
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
			if fit == full {
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
