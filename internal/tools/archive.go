package tools

import (
	"cmp"
	"context"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

// archiveSearch holds the archive and each message's own lines, folded: what
// its author wrote, not what they quoted.
type archiveSearch struct {
	a    *bundle.Archive
	msgs []archiveMsg
	// directors indexes the messages by the rules director sitting when they
	// were written.
	directors []int
}

type archiveMsg struct {
	thread *bundle.Thread
	i      int
	own    string // folded
	title  string // folded
	// question is what a rules director's reply answers, folded, its
	// decklist out; "" for anyone else's post and a thread's first.
	question string
}

func newArchiveSearch(a *bundle.Archive) *archiveSearch {
	s := &archiveSearch{a: a}
	for t := range a.Threads {
		th := &a.Threads[t]
		title := bundle.Fold(th.Title)
		for i, m := range th.Messages {
			if m.Body == "" {
				continue
			}
			seat, question := bundle.Role(m.Author, m.Date).Seat, ""
			if seat == bundle.Sitting || seat == bundle.Possibly {
				question = bundle.Fold(prose(questionOf(th, i)))
			}
			if seat == bundle.Sitting {
				s.directors = append(s.directors, len(s.msgs))
			}
			s.msgs = append(s.msgs, archiveMsg{th, i, bundle.Fold(ownLines(m.Body)), title, question})
		}
	}
	return s
}

// named is where a rules director's message names a card outside a
// decklist, folded: its own words, the question it answers; "" where not.
type named struct{ own, question string }

// naming answers the rules directors' messages whose own words, or the
// question they answer, name a card outside a decklist.
func (s *archiveSearch) naming(names []string) map[int]named {
	found := map[int]named{}
	if len(names) == 0 {
		return found
	}
	names1 := func(text string) bool {
		return slices.ContainsFunc(names, func(n string) bool { return countWord(text, n) > 0 })
	}
	for _, k := range s.directors {
		m := s.msgs[k]
		var n named
		if names1(m.own) {
			if own := bundle.Fold(prose(ownLines(m.thread.Messages[m.i].Body))); names1(own) {
				n.own = own
			}
		}
		if names1(m.question) {
			n.question = m.question
		}
		if n != (named{}) {
			found[k] = n
		}
	}
	return found
}

// prose drops a text's decklist lines.
func prose(text string) string {
	var b strings.Builder
	for line := range strings.Lines(text) {
		if len(line) >= 80 || !reDeckLine.MatchString(line) {
			b.WriteString(line)
		}
	}
	return b.String()
}

// quotesNone says whether a reply quotes nothing of what it answers.
func quotesNone(th *bundle.Thread, i int) bool {
	return i > 0 && strings.TrimSpace(unquote(th.Messages[i].Body)) == ""
}

// questionOf answers what a reply answers: the lines it quotes, unquoted, else
// the own words of the post before it in its thread (the archive keeps no
// parent: the post before may answer another).
func questionOf(th *bundle.Thread, i int) string {
	if q := unquote(th.Messages[i].Body); strings.TrimSpace(q) != "" {
		return q
	}
	if i == 0 {
		return ""
	}
	return ownLines(th.Messages[i-1].Body)
}

func isQuote(line string) bool { return strings.HasPrefix(strings.TrimSpace(line), ">") }

// unquote answers a text's quoted lines, their marks stripped.
func unquote(text string) string {
	var b strings.Builder
	for line := range strings.Lines(text) {
		if isQuote(line) {
			b.WriteString(strings.TrimLeft(strings.TrimSpace(line), "> ") + "\n")
		}
	}
	return b.String()
}

// answering answers a reply as evidence, its question quoted first, up to
// about qLimit bytes of it, and its own words up to about ownLimit (0: whole).
// A reply quotes nothing when before is set. A reply quoting is kept as written from
// its quoted block holding most needles through its own words after it: one
// reply may answer several questions in turn. A reply quoting nothing follows
// the post before it, which before says where from.
func answering(th *bundle.Thread, i int, needles []string, qLimit, ownLimit int) (body, before string) {
	lines := strings.Split(strings.ReplaceAll(th.Messages[i].Body, "\r\n", "\n"), "\n")
	best, most := -1, -1
	for j := range lines {
		if !isQuote(lines[j]) || (j > 0 && isQuote(lines[j-1])) {
			continue
		}
		end := j
		for end < len(lines) && isQuote(lines[end]) {
			end++
		}
		block, n := bundle.Fold(unquote(strings.Join(lines[j:end], "\n"))), 0
		for _, nd := range needles {
			if countWord(block, nd) > 0 {
				n++
			}
		}
		if n > most {
			best, most = j, n
		}
	}
	var q, own string
	if best >= 0 {
		end := best
		for end < len(lines) && isQuote(lines[end]) {
			end++
		}
		q = unquote(strings.Join(lines[best:end], "\n"))
		var after []string
		for _, l := range lines[end:] {
			if isQuote(l) && strings.TrimSpace(strings.Join(after, "")) != "" {
				break
			}
			after = append(after, l)
		}
		own = strings.Join(after, "\n")
	} else if i > 0 {
		prev := th.Messages[i-1]
		q, own = ownLines(prev.Body), th.Messages[i].Body
		before = fmt.Sprintf("the post before it in its thread, by %s, https://usenet.krcg.org/t/%s/#m%d, which it does not quote", prev.Author, th.ID, i-1)
	}
	if qLimit > 0 {
		q = around(q, needles, qLimit)
	}
	if ownLimit > 0 && len(own) > ownLimit {
		own = around(own, needles, ownLimit)
	}
	own = strings.TrimSpace(own)
	if q = strings.TrimSpace(q); q == "" {
		return own, before
	}
	return "> " + strings.ReplaceAll(q, "\n", "\n> ") + "\n\n" + own, before
}

// reDeckLine is a decklist's line, "3x Conditioning" or "Conditioning x3":
// a card listed, not discussed.
var reDeckLine = regexp.MustCompile(`(?i)^\s*(?:\d+\s*[x×]?|[x×]\s*\d+)\s+\S|\s[x×]\s*\d+\s*$`)

func ownLines(body string) string {
	var b strings.Builder
	for line := range strings.Lines(body) {
		if !strings.HasPrefix(strings.TrimSpace(line), ">") {
			b.WriteString(line)
		}
	}
	return b.String()
}

// SetArchive gives the archive tool the archive to search, as the hosted
// server fetches it.
func (l *Live) SetArchive(a *bundle.Archive) {
	if cur := l.archive.Load(); cur != nil && cur.a == a {
		return
	}
	l.archive.Store(newArchiveSearch(a))
}

// archiveFor answers the archive to search, loading it on first use where the
// Live was given a loader (a local install).
func (l *Live) archiveFor(ctx context.Context) (*archiveSearch, error) {
	if s := l.archive.Load(); s != nil {
		return s, nil
	}
	if l.LoadArchive == nil {
		return nil, fmt.Errorf("the rulings archive is still loading on the server: call again in a minute")
	}
	l.loading.Lock()
	defer l.loading.Unlock()
	if s := l.archive.Load(); s != nil {
		return s, nil
	}
	a, err := l.LoadArchive(ctx)
	if err != nil {
		return nil, fmt.Errorf("the rulings archive could not be loaded: %w", err)
	}
	l.SetArchive(a)
	return l.archive.Load(), nil
}

type archiveInput struct {
	Query     string `json:"query"`
	Author    string `json:"author,omitempty"`
	Directors bool   `json:"directors,omitempty"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	Detail    bool   `json:"detail,omitempty"`
}

var (
	rePhrase = regexp.MustCompile(`"([^"]+)"`)
	reDay    = regexp.MustCompile(`^\d{4}(-\d{2}(-\d{2})?)?$`)
)

func addArchive(s *mcp.Server, l *Live) {
	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_rulings_archive",
		Title: "VTES rulings archive",
		Description: "Never state a ruling from memory: quote what settles it. " +
			"Full-text search of the Vampire: The Eternal Struggle (VTES) rulings archive (usenet.krcg.org): the newsgroup " +
			"from 1994, the V:EKN forum and BoardGameGeek, every thread a rules director took part in. Returns posts with their author, " +
			"date, whether the author was the sitting rules director then, and the citable link. For a point no ruling settles; " +
			"a post by a player is not a ruling. A rules director's reply is found by its own words, and also by its thread's title " +
			"and the question it answers, shown with it. vtes_rules_case already quotes the posts the rulings cite, " +
			"and the rules directors' posts on a card asked about alone or on two cards no ruling names together.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"query":     {Type: "string", Description: `whole words every post must hold in what its author wrote, not what it quotes (a rules director's reply: or in its thread's title or the question it answers); "a quoted phrase" as written; card names work`},
				"author":    {Type: "string", Description: "part of the author's name, e.g. LSJ, Ankha"},
				"directors": {Type: "boolean", Description: "only posts by the rules director sitting when they were written"},
				"from":      {Type: "string", Description: "the earliest date, YYYY, YYYY-MM or YYYY-MM-DD"},
				"to":        {Type: "string", Description: "the latest date, YYYY, YYYY-MM or YYYY-MM-DD"},
				"detail":    {Type: "boolean", Description: "more posts, each whole as far as the result's size allows"},
			},
			Required: []string{"query"},
		},
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: new(false)},
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in archiveInput) (*mcp.CallToolResult, any, error) {
		var needles []string
		rest := rePhrase.ReplaceAllStringFunc(in.Query, func(m string) string {
			if p := bundle.Fold(strings.Trim(m, `"`)); p != "" {
				needles = append(needles, p)
			}
			return " "
		})
		for _, w := range strings.Fields(bundle.Fold(rest)) {
			if w = strings.Trim(w, ".,;:!?()"); len(w) > 1 && !stopwords[w] {
				needles = append(needles, w)
			}
		}
		if len(needles) == 0 || (in.From != "" && !reDay.MatchString(in.From)) || (in.To != "" && !reDay.MatchString(in.To)) {
			return nil, nil, fmt.Errorf(`give words to search and dates as YYYY, YYYY-MM or YYYY-MM-DD, e.g. {"query": "Freak Drive ended", "directors": true}; got %+v`, in)
		}
		as, err := l.archiveFor(ctx)
		if err != nil {
			return nil, nil, err
		}
		author := bundle.Fold(in.Author)
		type hit struct {
			m     *archiveMsg
			msg   *bundle.Message
			role  bundle.Standing
			score int
			// answer: a director's reply found by its thread's title or the
			// question it answers, not by its own words alone
			answer bool
		}
		var hits []hit
		for k := range as.msgs {
			m := &as.msgs[k]
			msg := &m.thread.Messages[m.i]
			if (in.From != "" && msg.Date < in.From) || (in.To != "" && msg.Date[:min(len(msg.Date), len(in.To))] > in.To) {
				continue
			}
			if author != "" && !strings.Contains(bundle.Fold(msg.Author), author) {
				continue
			}
			score, answer := 0, false
			for _, n := range needles {
				c, t := countWord(m.own, n), 0
				if countWord(m.title, n) > 0 {
					t = 2
				}
				if c == 0 {
					if m.question == "" {
						score = -1
						break
					}
					c, answer = countWord(m.question, n), true
					if c == 0 && t == 0 {
						score = -1
						break
					}
				}
				score += min(c, 5) + t
			}
			if score < 0 {
				continue
			}
			role := bundle.Role(msg.Author, msg.Date)
			if in.Directors && role.Seat != bundle.Sitting && role.Seat != bundle.Possibly {
				continue
			}
			hits = append(hits, hit{m, msg, role, score, answer})
		}
		rank := func(s bundle.Seat) int {
			switch s {
			case bundle.Sitting:
				return 0
			case bundle.Possibly:
				return 1
			}
			return 2
		}
		slices.SortStableFunc(hits, func(a, b hit) int {
			return cmp.Or(cmp.Compare(rank(a.role.Seat), rank(b.role.Seat)), compareBool(a.answer, b.answer),
				cmp.Compare(b.score, a.score), strings.Compare(b.msg.Date, a.msg.Date))
		})
		n, answers, limit, total := 8, 3, 600, budget
		if in.Detail {
			n, answers, limit, total = 25, 8, 0, detailBudget
		}
		var picked []hit
		own, replies := 0, 0
		for _, h := range hits {
			if h.answer {
				replies++
				if replies <= answers {
					picked = append(picked, h)
				}
			} else if own++; own <= n {
				picked = append(picked, h)
			}
		}
		var out strings.Builder
		switch {
		case len(hits) == 0:
			fmt.Fprintf(&out, "No post of the archive holds all of %q as words of its own, and no rules director's reply answers a question that does. Try fewer words, or other spellings.\n\n", needles)
		case replies == 0:
			fmt.Fprintf(&out, "%d posts hold %q as words of their own; %d below, by the sitting rules director first, then possibly sitting, then everyone else, each by how often it holds them.\n\n",
				own, needles, len(picked))
		default:
			fmt.Fprintf(&out, "%d posts hold %q as words of their own, and %d rules directors' replies hold them only with their thread's title or the question they answer, quoted before the reply; "+
				"%d below, by the sitting rules director first, then possibly sitting, then everyone else, a post holding them before a reply found by its question, each by how often it holds them.\n\n",
				own, needles, replies, len(picked))
		}
		shown := 0
		for _, h := range picked {
			p := bundle.Post{
				Link:   fmt.Sprintf("https://usenet.krcg.org/t/%s/#m%d", h.m.thread.ID, h.m.i),
				Thread: h.m.thread.Title, Author: h.msg.Author, Date: h.msg.Date, Body: h.msg.Body,
			}
			// a reply's question is its evidence: twice the room, a third of it quoted
			heading, room := "Post", limit
			body := func(limit int) string { return around(h.msg.Body, needles, limit) }
			if h.answer || h.m.question != "" && quotesNone(h.m.thread, h.m.i) {
				heading, room = "Reply, after the question it quotes", 2*limit
				if _, before := answering(h.m.thread, h.m.i, needles, 0, 0); before != "" {
					heading = "Reply, after " + before
				}
				body = func(limit int) string {
					b, _ := answering(h.m.thread, h.m.i, needles, limit*2/3, limit)
					return b
				}
				p.Body = body(0)
			}
			if !in.Detail {
				p.Body = body(limit)
			}
			left := total - out.Len() - 1200 // and the closing lines
			if left < 1000 {
				break
			}
			var post strings.Builder
			writePost(&post, p, heading, "", room)
			text := post.String()
			if len(text) > left {
				post.Reset()
				p.Body = body(left)
				writePost(&post, p, heading, "", 0)
				// cut as written: each line's prefix counts too
				text = post.String()
				text = text[:strings.LastIndex(text[:min(len(text), left)], "\n")+1] + "(cut here: the post is whole at its link)\n"
			}
			out.WriteString(text + "\n")
			shown++
		}
		if len(hits) > shown {
			fmt.Fprintf(&out, "%d more posts: add words, an author, dates, or directors only.\n\n", len(hits)-shown)
		}
		out.WriteString("A post weighs as its author: the sitting rules director's settles a point no ruling covers, a player's settles nothing. " +
			"Quote it with its link and its author's standing, never a ruling from memory.\n\n")
		fmt.Fprintf(&out, "Archive: posts to %s.\n", as.a.LastPost)
		return text(out.String()), nil, nil
	})
}

// compareBool orders false before true.
func compareBool(a, b bool) int {
	switch {
	case a == b:
		return 0
	case a:
		return 1
	}
	return -1
}

// countWord counts a folded needle in a folded text as whole words.
func countWord(text, needle string) int { return len(wordsAt(text, needle)) }

// wordsAt answers where a folded needle stands in a folded text as whole words.
func wordsAt(text, needle string) []int {
	var at []int
	for i := 0; ; {
		j := strings.Index(text[i:], needle)
		if j < 0 {
			return at
		}
		start, end := i+j, i+j+len(needle)
		if (start == 0 || !isWordByte(text[start-1])) && (end == len(text) || !isWordByte(text[end])) {
			at = append(at, start)
		}
		i = start + 1
	}
}

func isWordByte(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b >= 0x80
}

// around keeps a post from the paragraph of its own words holding the most
// needles, with the paragraph it quotes just before, up to about limit bytes.
func around(body string, needles []string, limit int) string {
	paras := strings.Split(strings.ReplaceAll(body, "\r\n", "\n"), "\n\n")
	best, most := 0, 0
	for i, p := range paras {
		own := bundle.Fold(ownLines(p))
		found := 0
		for _, n := range needles {
			if countWord(own, n) > 0 {
				found++
			}
		}
		if found > most {
			best, most = i, found
		}
	}
	start := best
	if start > 0 && strings.TrimSpace(ownLines(paras[start-1])) == "" {
		start--
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

// clip cuts a text at the last space within budget, marking the cut; a budget
// too small for a useful start keeps only the mark.
func clip(line string, budget int) string {
	if budget < 80 {
		return "[…]"
	}
	cut := strings.LastIndex(line[:budget], " ")
	if cut < 0 {
		cut = budget
	}
	return strings.ToValidUTF8(line[:cut], "") + " […]"
}
