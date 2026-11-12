package bundle

import (
	"bytes"
	"regexp"
	"slices"
	"strconv"
	"strings"
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

const advancedURL = "https://vtes-biased.github.io/vtes-advanced-rules/"

var (
	reHeading  = regexp.MustCompile(`^(#{1,6})\s+(.*?)\s*$`)
	reMdLink   = regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)
	reNumber   = regexp.MustCompile(`^(\d+(?:\.\d+)*[a-z]?)\.?\s+`)
	reFootnote = regexp.MustCompile(`^\[\^([^\]]+)\]:\s*(.*)$`)
	reNoteCall = regexp.MustCompile(`\[\^([^\]]+)\]`)
	reBlanks   = regexp.MustCompile(`\n{3,}`)
)

// slug is the anchor GitHub and Jekyll give a heading.
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
	return strings.TrimSpace(s), link
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
