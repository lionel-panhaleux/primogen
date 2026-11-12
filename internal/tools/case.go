package tools

import (
	"cmp"
	"context"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"
	"unicode"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

type caseInput struct {
	Question string   `json:"question"`
	Cards    []string `json:"cards,omitempty"`
	Detail   bool     `json:"detail,omitempty"`
}

var caseSchema = &jsonschema.Schema{
	Type: "object",
	Properties: map[string]*jsonschema.Schema{
		"question": {Type: "string", Description: "the rules question, as the player asked it"},
		"cards": {
			Type: "array", Items: &jsonschema.Schema{Type: "string"},
			Description: "every card the question involves, in English or translated, exact or approximate; at most 10. Given none, the cards the question names are found; given some, the others it names are only suggested",
		},
		"detail": {Type: "boolean", Description: "every ruling's original post, longer, and more sections, as far as the result's size allows"},
	},
	Required: []string{"question"},
}

// casePack is what one question gathers, before it is written.
type casePack struct {
	d        *data
	question string
	q        map[string]float64
	cards    []*bundle.Card
	notes    []string // how the names resolved
	detail   bool
	// posts holds the rulings whose original posts the pack may quote, best first.
	posts []*bundle.Ruling
}

func addCase(s *mcp.Server, l *Live) {
	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_rules_case",
		Title: "VTES rules question",
		Description: "Call first for any Vampire: The Eternal Struggle (VTES) rules question: never answer one from memory. " +
			"Give the question and the cards involved; returns in one call the evidence that settles it: each card's official text and " +
			"a breakdown of its effects and timing, every ruling with the rules director's original post, the rulings naming these cards " +
			"together, the governing sections of the rulebook and the advanced rules, what no source covers, and how to answer.",
		InputSchema: caseSchema,
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: new(false)},
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in caseInput) (*mcp.CallToolResult, any, error) {
		if strings.TrimSpace(in.Question) == "" || len(in.Cards) > maxNames || slices.ContainsFunc(in.Cards, func(n string) bool { return bundle.Fold(n) == "" }) {
			return nil, nil, fmt.Errorf(`give the question and at most %d non-empty card names, e.g. {"question": "Can I play Freak Drive after The Kiss of Ra?", "cards": ["Freak Drive", "The Kiss of Ra"]}`, maxNames)
		}
		p := newCasePack(l.p.Load(), in)
		return text(p.write()), nil, nil
	})
}

func newCasePack(d *data, in caseInput) *casePack {
	p := &casePack{d: d, detail: in.Detail}
	var dropped []string
	add := func(cards []*bundle.Card) {
		for _, c := range cards {
			switch {
			case slices.Contains(p.cards, c):
			case len(p.cards) < maxNames:
				p.cards = append(p.cards, c)
			default:
				dropped = append(dropped, c.Name)
			}
		}
	}
	for _, name := range in.Cards {
		cards, note := d.ix.resolve(name, "vtes_rules_case")
		if note != "" {
			p.notes = append(p.notes, note)
		}
		add(cards)
	}
	cards, maybe := d.findCards(in.Question)
	var found []string
	for _, c := range cards {
		if !slices.Contains(p.cards, c) {
			found = append(found, c.Name)
			if len(in.Cards) == 0 {
				add([]*bundle.Card{c})
			}
		}
	}
	switch {
	case len(found) > 0 && len(in.Cards) == 0:
		p.notes = append(p.notes, fmt.Sprintf("Found in the question: %s. Pass the cards yourself in cards if another was meant.", strings.Join(found, "; ")))
	case len(found) > 0:
		p.notes = append(p.notes, fmt.Sprintf("The question also names %s, not in cards: call again with them in cards if they are meant.", strings.Join(found, "; ")))
	}
	maybe = slices.DeleteFunc(maybe, func(n string) bool {
		return slices.ContainsFunc(p.cards, func(c *bundle.Card) bool { return c.Name == n })
	})
	if len(maybe) > 0 {
		p.notes = append(p.notes, fmt.Sprintf("The question may also name %s: pass them in cards if meant.", strings.Join(maybe, "; ")))
	}
	if len(dropped) > 0 {
		p.notes = append(p.notes, fmt.Sprintf("Left out past the %d cards a call takes: %s.", maxNames, strings.Join(dropped, "; ")))
	}
	p.question = in.Question
	p.q = query(in.Question, 1, nil)
	for _, c := range p.cards {
		query(strings.Join(c.Types, " ")+" "+reDisciplines.ReplaceAllString(c.Text, "$2"), 0.15, p.q)
	}
	p.choosePosts()
	return p
}

// findCards answers the cards a text names, in its order: an English name,
// then, in the words no name took, the words opening or ending one printed
// name and no other, judged by their words that are no stopword; neither a
// game term (a discipline, a clan, a title) nor a phrase the rulebook uses, a
// single word only capitalized, a name in the possessive read without its 's.
// It answers apart, by name, the cards it may name: a single word opening a
// sentence, a capitalized phrase of the rulebook, a single word of a name
// beside another capitalized word, words opening or ending a few names, words
// of a name none capitalized.
func (d *data) findCards(text string) (found []*bundle.Card, maybe []string) {
	words := splitWords(text)
	taken := make([]bool, len(words))
	type hit struct {
		at    int
		cards []*bundle.Card // found
		names []string       // suggested
	}
	var hits []hit
	ends := func(w string, marks string) bool { return strings.ContainsAny(w[len(w)-1:], marks) }
	capitalized := func(w string) bool {
		w = cleanWord(w)
		return w != "" && unicode.IsUpper([]rune(w)[0]) && !stopwords[bundle.Fold(w)]
	}
	for pass, index := range []map[string][]*bundle.Card{d.ix.english, d.ix.part} {
		for i := 0; i < len(words); i++ {
			matched := 0
			for n := min(8, len(words)-i); n >= 1 && matched == 0; n-- {
				if slices.Contains(taken[i:i+n], true) {
					continue
				}
				parts := make([]string, n)
				for k := range n {
					parts[k] = cleanWord(words[i+k])
				}
				name := bundle.Fold(strings.Join(parts, " "))
				base := strings.TrimSuffix(name, "'s")
				cards := index[name]
				if cards == nil {
					cards = index[base]
				}
				if cards == nil || slices.Contains(parts, "") || d.rules.terms[name] || d.rules.terms[base] {
					continue
				}
				core := parts
				capital := unicode.IsUpper([]rune(parts[0])[0])
				names, within := []string{cards[0].Name}, []string(nil)
				if pass == 1 {
					core = slices.DeleteFunc(slices.Clone(parts), func(w string) bool { return stopwords[bundle.Fold(w)] })
					if w := bundle.Fold(core[0]); len(core) == 1 && (d.rules.terms[w] || d.rules.terms[strings.TrimSuffix(w, "'s")]) {
						continue
					}
					capital = slices.ContainsFunc(core, capitalized)
					names = nil
					for _, c := range cards {
						if !slices.Contains(names, c.PrintedName) {
							names = append(names, c.PrintedName)
						}
					}
					within = names
				}
				single := len(core) == 1
				opens := i == 0 || ends(words[i-1], ".!?:")
				// a single word beside another capitalized one is part of a name given whole
				inName := pass == 1 && single && (i > 0 && !ends(words[i-1], ".!?:,") && capitalized(words[i-1]) ||
					i+n < len(words) && !ends(words[i+n-1], ".!?:,") && capitalized(words[i+n]))
				prose := d.rules.inRulebook(name, within) || d.rules.inRulebook(base, within)
				matched = n
				switch {
				case single && !capital:
					matched = 0
				case pass == 1 && (prose || len(names) > maxSuggested):
					// prose, or naming none: no shorter part of these words is tried
				case single && opens, inName, len(names) > 1, prose, pass == 1 && !capital:
					if capital || pass == 1 {
						hits = append(hits, hit{at: i, names: names})
					}
				default:
					hits = append(hits, hit{at: i, cards: cards})
				}
			}
			for k := range matched {
				taken[i+k] = true
			}
			i += max(0, matched-1)
		}
	}
	slices.SortStableFunc(hits, func(a, b hit) int { return cmp.Compare(a.at, b.at) })
	for _, h := range hits {
		for _, c := range h.cards {
			if !slices.Contains(found, c) {
				found = append(found, c)
			}
		}
		for _, n := range h.names {
			if !slices.Contains(maybe, n) {
				maybe = append(maybe, n)
			}
		}
	}
	return found, maybe
}

// maxSuggested is the most names words opening or ending them suggest: words
// fitting more name none.
const maxSuggested = 4

// splitWords cuts a text into words, keeping the punctuation a card's name
// holds; cleanWord trims what may end a word but no name.
func splitWords(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && !strings.ContainsRune("'’-.:,!?&", r)
	})
}

func cleanWord(w string) string { return strings.Trim(w, ".,:!?'’") }

// nameWords cuts a name as findCards cuts a text.
func nameWords(name string) []string {
	words := splitWords(name)
	for i, w := range words {
		words[i] = cleanWord(w)
	}
	return words
}

// choosePosts picks the rulings whose original posts the pack quotes: all in
// detail, else the ones naming another card of the question and the ones
// closest to its words.
func (p *casePack) choosePosts() {
	type cand struct {
		r     *bundle.Ruling
		score float64
	}
	var cands []cand
	for _, c := range p.cards {
		for i := range c.Rulings {
			cands = append(cands, cand{&c.Rulings[i], p.closeness(c, &c.Rulings[i])})
		}
	}
	slices.SortStableFunc(cands, func(a, b cand) int { return cmp.Compare(b.score, a.score) })
	for i, c := range cands {
		if p.detail || (i < 6 && c.score > 0) {
			p.posts = append(p.posts, c.r)
		}
	}
}

// closeness says how near a card's ruling is to the question: its words, and
// far more, another card of the question it names.
func (p *casePack) closeness(c *bundle.Card, r *bundle.Ruling) float64 {
	s := 0.0
	for _, t := range terms(reCardToken.ReplaceAllString(r.Text, "$1")) {
		s += p.q[t]
	}
	for _, other := range p.cards {
		if other != c && strings.Contains(r.Text, fmt.Sprintf("{%d|", other.ID)) {
			s += 100
		}
	}
	return s
}

// closest answers a card with only its n rulings closest to the question, in
// their order, and how many it left out; n 0 keeps them all.
func (p *casePack) closest(c *bundle.Card, n int) (*bundle.Card, int) {
	if n == 0 || len(c.Rulings) <= n {
		return c, 0
	}
	order := make([]int, len(c.Rulings))
	for i := range order {
		order[i] = i
	}
	slices.SortStableFunc(order, func(a, b int) int {
		return cmp.Compare(p.closeness(c, &c.Rulings[b]), p.closeness(c, &c.Rulings[a]))
	})
	keep := order[:n]
	slices.Sort(keep)
	short := *c
	short.Rulings = nil
	for _, i := range keep {
		short.Rulings = append(short.Rulings, c.Rulings[i])
	}
	return &short, len(c.Rulings) - n
}

// write writes the pack within its bound: the cards, their rulings (the
// closest to the question only, when all would not fit), the rulings naming
// them together, what no source covers and the procedure always; then the
// sections and the original posts, best first, while they fit.
func (p *casePack) write() string {
	d := p.d
	total, postLimit := budget, 700
	if p.detail {
		total, postLimit = detailBudget, 4000
	}
	posts := map[string]string{} // by ruling link
	for _, r := range p.posts {
		var b strings.Builder
		for _, ref := range r.References {
			if post, ok := d.b.Posts[ref.Archive]; ok {
				writePost(&b, post, "Original post", "   ", postLimit)
			}
		}
		if b.Len() > 0 {
			posts[r.Link] = b.String()
		}
	}
	head, sections, tail := p.sections()
	kept := map[string]bool{}
	rulings := 0 // the rulings a card keeps, 0 for all
	render := func(withSections []string) string {
		var out strings.Builder
		for _, n := range p.notes {
			fmt.Fprintf(&out, "%s\n\n", n)
		}
		if len(p.cards) == 0 {
			out.WriteString("No card given or found in the question: the rules sections below answer the words alone. " +
				"If a card is involved, call again with it in cards.\n\n")
		} else {
			out.WriteString("# The cards\n\n")
		}
		for _, full := range p.cards {
			c, more := p.closest(full, rulings)
			writeCard(&out, c, false, true, len(full.Rulings), func(r *bundle.Ruling) {
				if kept[r.Link] {
					out.WriteString(posts[r.Link])
				}
			})
			if more > 0 {
				fmt.Fprintf(&out, "%d more rulings on %s, less close to the question: vtes_cards gives them all.\n\n", more, c.Name)
			}
			if b := breakdown(c); len(b) > 0 {
				out.WriteString("Breakdown (computed from the text above, which governs):\n")
				for _, line := range b {
					fmt.Fprintf(&out, "- %s\n", line)
				}
				out.WriteString("\n")
			}
		}
		p.writeTogether(&out)
		if len(withSections) > 0 {
			out.WriteString(head)
			for _, s := range withSections {
				out.WriteString(s)
			}
			out.WriteString(tail)
		}
		p.writeAbsent(&out)
		out.WriteString(strings.TrimSpace(d.b.RulesAnswer))
		out.WriteString("\n\nAnswer from this evidence, quoting it and citing its links, never from memory. " +
			"Where nothing here settles a point, say it is unsettled rather than filling the gap.\n\n")
		return out.String()
	}
	const trailer = 800 // what is left out, and the data line
	for _, n := range []int{0, 25, 15, 10, 6, 3, 1} {
		if rulings = n; len(render(nil))+trailer <= total {
			break
		}
	}
	// the sections first, up to most of what is left, then the posts, then
	// the sections still out
	left := total - len(render(nil)) - len(head) - len(tail) - trailer
	var chosen []string
	in := 0
	for _, s := range sections {
		if len(s) > left*3/4-in {
			break
		}
		chosen = append(chosen, s)
		in += len(s)
	}
	left -= in
	shown := map[string]bool{}
	for _, full := range p.cards {
		c, _ := p.closest(full, rulings)
		for _, r := range c.Rulings {
			shown[r.Link] = true
		}
	}
	postsOut := 0
	for _, r := range p.posts {
		if t, ok := posts[r.Link]; ok && shown[r.Link] {
			if len(t) <= left {
				kept[r.Link] = true
				left -= len(t)
			} else {
				postsOut++
			}
		}
	}
	for _, s := range sections[len(chosen):] {
		if len(s) > left {
			break
		}
		chosen = append(chosen, s)
		left -= len(s)
	}
	out := render(chosen)
	if sectionsOut := len(sections) - len(chosen); sectionsOut > 0 || postsOut > 0 {
		cut := fmt.Sprintf("Left out to keep this result within its size: %d sections and %d rulings' original posts; ", sectionsOut, postsOut)
		if p.detail {
			cut += "ask about fewer cards, or read the sections with vtes_rules_read.\n\n"
		} else {
			cut += "the detail mode, fewer cards, or vtes_rules_read give more.\n\n"
		}
		return out + cut + DataLine(d.b)
	}
	return out + DataLine(d.b)
}

// writePost quotes a post under its author's standing: the lines it quotes,
// shortened, then its own words, up to limit bytes (0: whole).
func writePost(out *strings.Builder, post bundle.Post, heading, indent string, limit int) {
	fmt.Fprintf(out, "%s%s: %s, %s, %q, %s\n", indent, heading, bundle.Role(post.Author, post.Date).Label, post.Date, post.Thread, post.Link)
	var quoted, own []string
	for line := range strings.Lines(post.Body) {
		line = strings.TrimRight(line, " \n\r")
		if strings.HasPrefix(strings.TrimSpace(line), ">") && len(own) == 0 {
			quoted = append(quoted, line)
			continue
		}
		if line == "" && len(own) == 0 {
			continue
		}
		own = append(own, line)
	}
	write := func(lines []string, budget int) {
		size := 0
		for i, line := range lines {
			if line == "" && i > 0 && lines[i-1] == "" {
				continue
			}
			if budget > 0 && size+len(line) > budget {
				fmt.Fprintf(out, "%s| %s\n", indent, clip(line, budget-size))
				return
			}
			size += len(line)
			fmt.Fprintf(out, "%s| %s\n", indent, line)
		}
	}
	write(quoted, limit/3)
	write(own, limit)
}

// writeTogether lists the rulings that name two cards of the question, wherever
// they are filed, and in detail the rulings on other cards naming one.
func (p *casePack) writeTogether(out *strings.Builder) {
	if len(p.cards) == 0 {
		return
	}
	ix := p.d.rules
	inCase := map[int]bool{}
	for _, c := range p.cards {
		inCase[c.ID] = true
	}
	seen := map[string]bool{}
	var together, others []string
	for _, c := range p.cards {
		for _, nr := range ix.rulingsNaming[c.ID] {
			if inCase[nr.card.ID] || seen[nr.r.Text] {
				continue
			}
			seen[nr.r.Text] = true
			named := 0
			for _, m := range reCard.FindAllStringSubmatch(nr.r.Text, -1) {
				for _, o := range p.cards {
					if fmt.Sprint(o.ID) == m[1] {
						named++
					}
				}
			}
			line := fmt.Sprintf("- on %s: %s\n  Ruling: %s", nr.card.Name, reCardToken.ReplaceAllString(nr.r.Text, "$1"), nr.r.Link)
			if named >= 2 {
				together = append(together, line)
			} else {
				others = append(others, line)
			}
		}
	}
	if len(together) > 0 {
		out.WriteString("# Rulings on other cards naming two of these cards\n\n")
		out.WriteString(strings.Join(together, "\n") + "\n\n")
	}
	if len(others) > 0 && p.detail {
		out.WriteString("# Rulings on other cards naming one of these cards\n\n")
		out.WriteString(strings.Join(others[:min(len(others), 60)], "\n") + "\n\n")
		if len(others) > 60 {
			fmt.Fprintf(out, "%d more: vtes_cards gives each card's rulings.\n\n", len(others)-60)
		}
	} else if len(others) > 0 {
		fmt.Fprintf(out, "%d rulings on other cards name one of these cards: the detail mode lists them.\n\n", len(others))
	}
}

// sections renders the sections joined to the question's cards, then the
// sections closest to its words. A section joins a card by naming it, by
// citing a post one of its rulings cites, or as the rulebook section a ruling
// links to; the joined ones are ranked by their links and the words.
func (p *casePack) sections() (head string, sections []string, tail string) {
	ix := p.d.rules
	joined, ranked, limit := 6, 8, 1300
	if p.detail {
		joined, ranked, limit = 10, 12, 3000
	}
	links := map[int]int{}
	for _, c := range p.cards {
		for _, i := range ix.sectionsNaming[c.ID] {
			links[i]++
		}
		for _, r := range c.Rulings {
			for _, ref := range r.References {
				for _, i := range ix.sectionsCiting[ref.Label] {
					links[i]++
				}
				if i, ok := ix.byID[ref.Section]; ok {
					links[i]++
				}
			}
		}
	}
	question := " " + strings.Join(terms(p.question), " ") + " "
	results := ix.search.rank(p.q, nil)
	scores := map[int]float64{}
	for _, r := range results {
		scores[r.i] = r.score
	}
	if len(results) > 0 {
		// a section the question names by its title
		for i, s := range ix.sections {
			if t := terms(reNumberPrefix.ReplaceAllString(s.Title, "")); len(t) >= 2 && strings.Contains(question, " "+strings.Join(t, " ")+" ") {
				scores[i] += results[0].score
			}
		}
		slices.SortStableFunc(results, func(a, b scored) int { return cmp.Compare(scores[b.i], scores[a.i]) })
	}
	top := 1.0
	if len(results) > 0 {
		top = results[0].score
	}
	weight := func(i int) float64 { return (scores[i] + 0.25*top) * float64(1+min(links[i], 3)) }
	join := slices.Collect(maps.Keys(links))
	slices.SortFunc(join, func(a, b int) int { return cmp.Or(cmp.Compare(weight(b), weight(a)), cmp.Compare(a, b)) })
	chosen := slices.Clone(join[:min(joined, len(join))])
	for _, r := range results {
		if len(chosen) >= joined+ranked {
			break
		}
		if !slices.Contains(chosen, r.i) {
			chosen = append(chosen, r.i)
		}
	}
	if len(chosen) == 0 {
		return "", nil, ""
	}
	head = "# The rules sections\n\n"
	if len(join) > 0 {
		head += fmt.Sprintf("%d sections name these cards, cite the posts their rulings cite, or are linked by their rulings: the %d closest to the question first, then the sections closest to the question's words.\n\n",
			len(join), min(joined, len(join)))
	}
	for _, i := range chosen {
		var b strings.Builder
		writeSection(&b, ix.sections[i], p.q, limit)
		sections = append(sections, b.String())
	}
	var rest []string
	for _, i := range join {
		if !slices.Contains(chosen, i) {
			rest = append(rest, ix.sections[i].ID)
		}
	}
	if len(rest) > 0 {
		tail = fmt.Sprintf("Other sections joined to these cards: %s (vtes_rules_read).\n\n", strings.Join(rest[:min(len(rest), 20)], ", "))
	}
	return head, sections, tail
}

// writeAbsent states what no source covers, so that the answer can say so.
func (p *casePack) writeAbsent(out *strings.Builder) {
	var lines []string
	ix := p.d.rules
	for _, c := range p.cards {
		if len(c.Rulings) == 0 {
			lines = append(lines, fmt.Sprintf("No ruling of the rulings database is filed under %s.", c.Name))
		}
		if len(ix.sectionsNaming[c.ID]) == 0 {
			lines = append(lines, fmt.Sprintf("No section of the advanced rules names %s.", c.Name))
		}
	}
	for i, a := range p.cards {
		for _, b := range p.cards[i+1:] {
			if !p.namedTogether(a, b) {
				lines = append(lines, fmt.Sprintf("No ruling names %s and %s together.", a.Name, b.Name))
			}
		}
	}
	if len(lines) == 0 {
		return
	}
	out.WriteString("# What no source here covers\n\n")
	for _, l := range lines {
		fmt.Fprintf(out, "- %s\n", l)
	}
	out.WriteString("\n")
}

func (p *casePack) namedTogether(a, b *bundle.Card) bool {
	names := func(r *bundle.Ruling, c *bundle.Card) bool {
		return strings.Contains(r.Text, fmt.Sprintf("{%d|", c.ID))
	}
	for i := range a.Rulings {
		if names(&a.Rulings[i], b) {
			return true
		}
	}
	for i := range b.Rulings {
		if names(&b.Rulings[i], a) {
			return true
		}
	}
	for _, nr := range p.d.rules.rulingsNaming[a.ID] {
		if names(nr.r, b) {
			return true
		}
	}
	return false
}

var (
	reNumberPrefix = regexp.MustCompile(`^\d+(?:\.\d+)*[a-z]?\.?\s+`)
	reDisciplines  = regexp.MustCompile(`^((?:\[[a-zA-Z]+\]\s*)+)(.*)$`)
	reTiming       = regexp.MustCompile(`(?i)\b(only usable|usable|only one|once each|do not replace|cannot be played|as the action is announced|before range|after (?:combat|resolution|action resolution|blocks? (?:are|is) declined)|during (?:an?|your|the|each)|out-of-turn|reaction|unique|limited)\b`)
	reSentence     = regexp.MustCompile(`[^.]+[.)]*`)
)

// breakdown reads a library card's text into its effects by level and the
// sentences limiting when it is played.
func breakdown(c *bundle.Card) []string {
	if slices.Contains(c.Types, "Vampire") || slices.Contains(c.Types, "Imbued") {
		return nil
	}
	var lines []string
	var timing []string
	for line := range strings.Lines(c.Text) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if m := reDisciplines.FindStringSubmatch(line); m != nil {
			marks := strings.Fields(strings.NewReplacer("][", "] [").Replace(m[1]))
			level := "With"
			switch {
			case allCase(marks, unicode.IsLower):
				level = "Inferior"
			case allCase(marks, unicode.IsUpper):
				level = "Superior"
			}
			lines = append(lines, fmt.Sprintf("%s %s: %s", level, strings.Join(marks, " "), m[2]))
		} else {
			lines = append(lines, "Any level: "+line)
		}
		for _, s := range reSentence.FindAllString(line, -1) {
			s = strings.TrimSpace(reDisciplines.ReplaceAllString(strings.TrimSpace(s), "$2"))
			if reTiming.MatchString(s) && !slices.Contains(timing, s) {
				timing = append(timing, s)
			}
		}
	}
	if len(timing) > 0 {
		lines = append(lines, "Timing and limits: "+strings.Join(timing, " / "))
	} else {
		lines = append(lines, "Timing and limits: the text states none; the card type's rules apply.")
	}
	return slices.Concat([]string{"Type: " + strings.Join(c.Types, " / ")}, lines)
}

func allCase(marks []string, is func(rune) bool) bool {
	for _, m := range marks {
		for _, r := range strings.Trim(m, "[]") {
			if unicode.IsLetter(r) && !is(r) {
				return false
			}
		}
	}
	return true
}
