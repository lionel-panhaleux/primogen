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
}

type archiveMsg struct {
	thread *bundle.Thread
	i      int
	own    string // folded
	title  string // folded
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
			s.msgs = append(s.msgs, archiveMsg{th, i, bundle.Fold(ownLines(m.Body)), title})
		}
	}
	return s
}

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
			"a post by a player is not a ruling. vtes_rules_case already quotes the posts the rulings cite.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"query":     {Type: "string", Description: `whole words every post must hold in what its author wrote, not what it quotes; "a quoted phrase" as written; card names work`},
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
			score := 0
			for _, n := range needles {
				c := countWord(m.own, n)
				if c == 0 {
					score = -1
					break
				}
				score += min(c, 5)
				if countWord(m.title, n) > 0 {
					score += 2
				}
			}
			if score < 0 {
				continue
			}
			role := bundle.Role(msg.Author, msg.Date)
			if in.Directors && role.Seat != bundle.Sitting && role.Seat != bundle.Possibly {
				continue
			}
			hits = append(hits, hit{m, msg, role, score})
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
			return cmp.Or(cmp.Compare(rank(a.role.Seat), rank(b.role.Seat)), cmp.Compare(b.score, a.score), strings.Compare(b.msg.Date, a.msg.Date))
		})
		n, limit, total := 8, 600, budget
		if in.Detail {
			n, limit, total = 25, 0, detailBudget
		}
		var out strings.Builder
		if len(hits) == 0 {
			fmt.Fprintf(&out, "No post of the archive holds all of %q as words of its own. Try fewer words, or other spellings.\n\n", needles)
		} else {
			fmt.Fprintf(&out, "%d posts hold %q as words of their own; %d below, by the sitting rules director first, then possibly sitting, then everyone else, each by how often it holds them.\n\n",
				len(hits), needles, min(n, len(hits)))
		}
		shown := 0
		for _, h := range hits[:min(n, len(hits))] {
			p := bundle.Post{
				Link:   fmt.Sprintf("https://usenet.krcg.org/t/%s/#m%d", h.m.thread.ID, h.m.i),
				Thread: h.m.thread.Title, Author: h.msg.Author, Date: h.msg.Date, Body: h.msg.Body,
			}
			if !in.Detail {
				p.Body = around(h.msg.Body, needles, limit)
			}
			left := total - out.Len() - 1200 // and the closing lines
			if left < 1000 {
				break
			}
			var post strings.Builder
			writePost(&post, p, "Post", "", limit)
			text := post.String()
			if len(text) > left {
				post.Reset()
				p.Body = around(h.msg.Body, needles, left)
				writePost(&post, p, "Post", "", 0)
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

// countWord counts a folded needle in a folded text as whole words.
func countWord(text, needle string) int {
	n := 0
	for i := 0; ; {
		j := strings.Index(text[i:], needle)
		if j < 0 {
			return n
		}
		start, end := i+j, i+j+len(needle)
		if (start == 0 || !isWordByte(text[start-1])) && (end == len(text) || !isWordByte(text[end])) {
			n++
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
