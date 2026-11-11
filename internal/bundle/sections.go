package bundle

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// The documents the sections come from, by the name a section's ID begins with.
const (
	DocRulebook   = "rulebook"
	DocAdvanced   = "advanced"
	DocTournament = "tournament"
	DocJudges     = "judges"
)

// DocTitles names each document as a reader knows it.
var DocTitles = map[string]string{
	DocRulebook:   "the rulebook",
	DocAdvanced:   "the advanced rules",
	DocTournament: "the tournament rules",
	DocJudges:     "the judges' guide",
}

// Section is a heading of a rules document and the text up to the next one.
type Section struct {
	ID    string // "<doc>:<number or slug>", stable while the heading is
	Doc   string
	Title string
	// Path holds the titles above it, outermost first.
	Path  []string
	Level int
	Text  string
	URL   string
	// Notes are the advanced rules' footnotes the text calls, which cite its sources.
	Notes []string
}

const (
	rulebookURL = "https://www.vekn.net/rulebook"
	advancedURL = "https://vtes-biased.github.io/vtes-advanced-rules/"
)

var (
	reHeading   = regexp.MustCompile(`^(#{1,6})\s+(.*?)\s*$`)
	reMdLink    = regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)
	reBoldTitle = regexp.MustCompile(`^\*\*([^*]+?)(:)?\*\*(.*)$`)
	reNumber    = regexp.MustCompile(`^(\d+(?:\.\d+)*[a-z]?)\.?\s+`)
	reFootnote  = regexp.MustCompile(`^\[\^([^\]]+)\]:\s*(.*)$`)
	reNoteCall  = regexp.MustCompile(`\[\^([^\]]+)\]`)
	reSeeAlso   = regexp.MustCompile(`\s*\(See [^)]*\)+\s*$`)
	reBlanks    = regexp.MustCompile(`\n{3,}`)
)

// slug is the anchor GitHub, Jekyll and vekn.net give a heading.
func slug(title string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(title) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '-':
			b.WriteRune(r)
		case r == ' ':
			b.WriteByte('-')
		}
	}
	return b.String()
}

// mdTitle strips a markdown heading to its words, and its link if it has one.
func mdTitle(s string) (string, string) {
	link := ""
	s = reMdLink.ReplaceAllStringFunc(s, func(m string) string {
		sub := reMdLink.FindStringSubmatch(m)
		if link == "" {
			link = sub[2]
		}
		return sub[1]
	})
	s = strings.ReplaceAll(s, "**", "")
	return strings.TrimSpace(reSeeAlso.ReplaceAllString(s, "")), link
}

type sectionBuilder struct {
	doc      string
	sections []Section
	stack    []Section // the open headings, outermost first
	ids      map[string]int
	text     strings.Builder
}

func (sb *sectionBuilder) flush() {
	if len(sb.stack) > 0 {
		s := &sb.sections[len(sb.sections)-1]
		s.Text = strings.TrimSpace(reBlanks.ReplaceAllString(sb.text.String(), "\n\n"))
	}
	sb.text.Reset()
}

// open starts a section; key is its number or else its slug.
func (sb *sectionBuilder) open(level int, title, key, url string) {
	sb.flush()
	for len(sb.stack) > 0 && sb.stack[len(sb.stack)-1].Level >= level {
		sb.stack = sb.stack[:len(sb.stack)-1]
	}
	// a repeated title is told apart by its parent's, so that an id moves
	// only when its own heading does
	id := sb.doc + ":" + key
	if sb.ids[id] > 0 && len(sb.stack) > 0 {
		id += "-in-" + slug(reNumber.ReplaceAllString(sb.stack[len(sb.stack)-1].Title, ""))
	}
	if sb.ids[id]++; sb.ids[id] > 1 {
		id += "-" + strconv.Itoa(sb.ids[id])
	}
	var path []string
	for _, p := range sb.stack {
		path = append(path, p.Title)
	}
	s := Section{ID: id, Doc: sb.doc, Title: title, Path: path, Level: level, URL: url}
	sb.stack = append(sb.stack, s)
	sb.sections = append(sb.sections, s)
}

func key(title string) string {
	if m := reNumber.FindStringSubmatch(title); m != nil {
		return m[1]
	}
	return slug(title)
}

// parseRulebook splits the rulebook at its headings and at a paragraph, a box
// or a list item opening on a bold title (**Combat Ends:**). Each is given the
// anchor vekn.net would slug it to, without the step number of "1. UNLOCK
// PHASE", which linkRulebook checks; a chapter, its page.
func parseRulebook(md []byte) []Section {
	sb := &sectionBuilder{doc: DocRulebook, ids: map[string]int{}}
	level := 1
	for line := range strings.Lines(string(md)) {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "![") {
			continue
		}
		if m := reHeading.FindStringSubmatch(trimmed); m != nil {
			title, link := mdTitle(m[2])
			level = len(m[1])
			anchor := slug(reNumber.ReplaceAllString(title, ""))
			url := rulebookURL + "#" + anchor
			if level == 1 && link != "" {
				anchor, url = link[strings.LastIndex(link, "/")+1:], link
			}
			sb.open(level, title, anchor, url)
			continue
		}
		unboxed := strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(trimmed, ">"), "- "))
		if m := reBoldTitle.FindStringSubmatch(unboxed); m != nil && len(sb.stack) > 0 && (m[2] != "" || m[3] == "" || m[3] == `\`) && m[1] != "ADVANCED RULES" {
			title := strings.TrimSpace(m[1])
			anchor := slug(reNumber.ReplaceAllString(title, ""))
			sb.open(level+1, title, anchor, rulebookURL+"#"+anchor)
		}
		sb.text.WriteString(line)
	}
	sb.flush()
	return sb.sections
}

// bare drops an anchor's small words, which vekn.net words otherwise.
func bare(anchor string) string {
	return strings.Join(slices.DeleteFunc(strings.Split(anchor, "-"), func(w string) bool {
		return slices.Contains([]string{"of", "in", "an", "a", "the", "and"}, w)
	}), "-")
}

// linkRulebook keeps a rulebook section's anchor where vekn.net's page has
// it, corrects one the page words otherwise ("IMPORTANT TERMS IN THE GAME" is
// #important-terms-of-the-game), and else links the nearest section above it
// that has one. The id stays the heading's, whatever the page holds.
func linkRulebook(sections []Section, anchors map[string]bool) {
	known := map[string]string{} // anchor without its small words → anchor
	for a := range anchors {
		known[bare(a)] = a
	}
	for i := range sections {
		s := &sections[i]
		a, ok := strings.CutPrefix(s.URL, rulebookURL+"#")
		if s.Doc != DocRulebook || !ok || anchors[a] {
			continue
		}
		if fixed, found := known[bare(a)]; found {
			s.URL = rulebookURL + "#" + fixed
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if sections[j].Doc == DocRulebook && sections[j].Level < s.Level {
				s.URL = sections[j].URL
				break
			}
		}
	}
}

// fetchAnchors answers the ids of a vekn.net page's content, lowercased.
func fetchAnchors(ctx context.Context, u string) (map[string]bool, error) {
	page, err := fetch(ctx, u, "")
	if err != nil {
		return nil, err
	}
	root, err := html.Parse(bytes.NewReader(page))
	if err != nil {
		return nil, err
	}
	content := find(root, "content")
	if content == nil {
		return nil, fmt.Errorf("%s: no content", u)
	}
	ids := map[string]bool{}
	for n := range content.Descendants() {
		if id := attr(n, "id"); id != "" {
			ids[strings.ToLower(id)] = true
		}
	}
	return ids, nil
}

func find(root *html.Node, id string) *html.Node {
	for n := range root.Descendants() {
		if n.Type == html.ElementNode && attr(n, "id") == id {
			return n
		}
	}
	return nil
}

// parseAdvanced splits the advanced rules at their headings, each section
// keeping the footnotes it calls.
func parseAdvanced(md []byte) []Section {
	body, notes := splitFootnotes(md)
	sb := &sectionBuilder{doc: DocAdvanced, ids: map[string]int{}}
	inFront := bytes.HasPrefix(body, []byte("---\n"))
	first := true
	for line := range strings.Lines(string(body)) {
		trimmed := strings.TrimSpace(line)
		if inFront {
			if trimmed == "---" && !first {
				inFront = false
			}
			first = false
			continue
		}
		if m := reHeading.FindStringSubmatch(trimmed); m != nil {
			// the document's own title heads no section
			if len(m[1]) > 1 {
				title, _ := mdTitle(m[2])
				sb.open(len(m[1]), title, key(title), advancedURL+"#"+slug(title))
			}
			continue
		}
		if trimmed == "---" {
			continue
		}
		sb.text.WriteString(line)
	}
	sb.flush()
	for i := range sb.sections {
		s := &sb.sections[i]
		for _, m := range reNoteCall.FindAllStringSubmatch(s.Text, -1) {
			if n, ok := notes[m[1]]; ok && !slices.Contains(s.Notes, n) {
				s.Notes = append(s.Notes, n)
			}
		}
	}
	return slices.DeleteFunc(sb.sections, func(s Section) bool { return s.Title == "References" && s.Text == "" })
}

// splitFootnotes takes the footnote definitions out of a markdown text, each
// joined to one line.
func splitFootnotes(md []byte) ([]byte, map[string]string) {
	notes := map[string]string{}
	var body bytes.Buffer
	current := ""
	for line := range strings.Lines(string(md)) {
		if m := reFootnote.FindStringSubmatch(strings.TrimRight(line, "\n")); m != nil {
			current = m[1]
			notes[current] = strings.TrimSpace(m[2])
			continue
		}
		if current != "" && (strings.HasPrefix(line, "    ") || strings.TrimSpace(line) == "") {
			if t := strings.TrimSpace(line); t != "" {
				notes[current] += " " + t
			}
			continue
		}
		current = ""
		body.WriteString(line)
	}
	for k, n := range notes {
		notes[k] = "[^" + k + "]: " + n
	}
	return body.Bytes(), notes
}

// parseVEKN reads a page of vekn.net, split at its headings. A heading linking
// to an article's own page sets the link of the sections under it.
func parseVEKN(page []byte, doc, pageURL string) ([]Section, error) {
	root, err := html.Parse(bytes.NewReader(page))
	if err != nil {
		return nil, err
	}
	sb := &sectionBuilder{doc: doc, ids: map[string]int{}}
	pu, err := url.Parse(pageURL)
	if err != nil {
		return nil, err
	}
	base := pu.Scheme + "://" + pu.Host
	n := find(root, "content")
	if n == nil {
		return nil, fmt.Errorf("%s: no content", pageURL)
	}
	article := pageURL
	{
		for c := range n.Descendants() {
			if c.Type != html.ElementNode || skipped(c, n) {
				continue
			}
			switch c.Data {
			case "h1", "h2", "h3", "h4", "h5", "h6":
				title := strings.Join(strings.Fields(textOf(c)), " ")
				if title == "" {
					continue
				}
				for a := range c.Descendants() {
					if href := attr(a, "href"); a.Data == "a" && strings.HasPrefix(href, "/") {
						article = base + href
					}
				}
				link := article
				if id := attr(c, "id"); id != "" {
					link = pageURL + "#" + id
				}
				// vekn.net's tournament rules are all h2: a number says the depth
				level := int(c.Data[1] - '0')
				if m := reNumber.FindStringSubmatch(title); m != nil {
					level = strings.Count(m[1], ".") + 1
				}
				sb.open(level, title, key(title), link)
			case "p", "li", "tr":
				if len(sb.stack) == 0 || inBlock(c, n) {
					continue
				}
				var t string
				if c.Data == "tr" {
					var cells []string
					for td := range c.ChildNodes() {
						if td.Type == html.ElementNode {
							cells = append(cells, strings.Join(strings.Fields(textOf(td)), " "))
						}
					}
					t = "| " + strings.Join(cells, " | ") + " |"
				} else {
					t = strings.Join(strings.Fields(textOf(c)), " ")
				}
				if t == "" {
					continue
				}
				if c.Data == "li" {
					t = "- " + t
				}
				sb.text.WriteString(t + "\n\n")
			}
		}
	}
	sb.flush()
	sections := slices.DeleteFunc(sb.sections, func(s Section) bool { return s.Title == "Table of Contents" })
	if len(sections) == 0 {
		return nil, fmt.Errorf("%s: no section found", pageURL)
	}
	return sections, nil
}

func attr(n *html.Node, name string) string {
	for _, a := range n.Attr {
		if a.Key == name {
			return a.Val
		}
	}
	return ""
}

func textOf(n *html.Node) string {
	var b strings.Builder
	for c := range n.Descendants() {
		if c.Type == html.TextNode {
			b.WriteString(c.Data)
		} else if c.Type == html.ElementNode && c.Data == "br" {
			b.WriteByte(' ')
		}
	}
	return b.String()
}

// skipped marks the page's furniture inside the content: the print buttons.
func skipped(n, article *html.Node) bool {
	for a := range n.Ancestors() {
		if a == article {
			return false
		}
		if strings.Contains(attr(a, "class"), "icons") {
			return true
		}
	}
	return false
}

// inBlock says whether a paragraph sits in another block already read whole.
func inBlock(n, article *html.Node) bool {
	for a := range n.Ancestors() {
		if a == article {
			return false
		}
		switch a.Data {
		case "p", "li", "tr", "h1", "h2", "h3", "h4", "h5", "h6":
			return true
		}
	}
	return false
}

var reVEKNDate = regexp.MustCompile(`(?i)\b(?:version updated|effective)\s+[^.|\n]{4,30}?\d{4}`)

// fetchVEKN reads one of vekn.net's rules pages, dated by the date it states
// and the hash of its text, the server stamping every page with the time asked.
func fetchVEKN(ctx context.Context, u, doc string) ([]Section, string, error) {
	page, err := fetch(ctx, u, "")
	if err != nil {
		return nil, "", err
	}
	sections, err := parseVEKN(page, doc, u)
	if err != nil {
		return nil, "", err
	}
	h := sha256.New()
	date := ""
	for _, s := range sections {
		fmt.Fprintf(h, "%s\n%s\n", s.Title, s.Text)
		if date == "" {
			date = reVEKNDate.FindString(s.Text)
		}
	}
	version := fmt.Sprintf("%x", h.Sum(nil))[:7]
	if date != "" {
		version = strings.Join(strings.Fields(date), " ") + ", " + version
	}
	return sections, version, nil
}
