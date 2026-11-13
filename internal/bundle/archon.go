package bundle

import (
	"regexp"
	"strings"
)

// The pages archon.vekn.net serves the documents at.
var archonPages = map[string]string{
	DocRulebook:   "https://archon.vekn.net/help/rules",
	DocTournament: "https://archon.vekn.net/help/tournament-rules",
	DocJudges:     "https://archon.vekn.net/help/judges-guide",
}

var (
	reTag          = regexp.MustCompile(`<[^>]*>`)
	reNotWord      = regexp.MustCompile(`[^a-z0-9_\s-]`)
	reSpaces       = regexp.MustCompile(`\s+`)
	reDashes       = regexp.MustCompile(`-+`)
	rePinnedID     = regexp.MustCompile(`\s*\{#[^}]+\}\s*$`)
	reStrong       = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	reCalloutTitle = regexp.MustCompile(`^[A-Z][A-Z\s&:,\-]+$`)
	reBoldOnly     = regexp.MustCompile(`^\*\*([^*]+)\*\*\s*$`)
)

// archonSlug is the id archon.vekn.net gives a heading (frontend/src/lib/markdown.ts).
func archonSlug(text string) string {
	s := strings.ToLower(text)
	s = reTag.ReplaceAllString(s, "")
	s = reNotWord.ReplaceAllString(s, "")
	s = reSpaces.ReplaceAllString(s, "-")
	return strings.TrimSpace(reDashes.ReplaceAllString(s, "-"))
}

// escaped is a text as the renderer writes it in HTML, which bold terms and
// boxes are slugged from: "During X" gives quotduring-x.
var escaped = strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", `"`, "&quot;", "'", "&#39;").Replace

// plain is a markdown heading's text as the renderer reads it.
func plain(s string) string {
	s = reMdLink.ReplaceAllString(s, "$1")
	return strings.NewReplacer("**", "", "\\", "", "`", "").Replace(s)
}

// element is what opens a line, as the renderer gives it an id.
type element struct {
	kind  int // 0 none, a heading, a box, a bold term
	title string
	id    string
}

const (
	noElement = iota
	headingElement
	boxElement
	termElement
)

// archonElements reads a document as archon.vekn.net renders it: headings and
// bold-only lines (made headings) and boxed rules with a capitalized title
// first, then every other bold term, in order, an id when it is still free.
func archonElements(lines []string) []element {
	elements := make([]element, len(lines))
	ids := map[string]bool{}
	consumed := map[int]int{} // line → the bold spans a box title took
	inBox := false
	for i, line := range lines {
		t := strings.TrimSpace(line)
		boxed := strings.HasPrefix(t, ">")
		opens := boxed && !inBox
		inBox = boxed
		if m := reHeading.FindStringSubmatch(t); m != nil {
			title := plain(rePinnedID.ReplaceAllString(m[2], ""))
			elements[i] = element{headingElement, title, archonSlug(title)}
			ids[elements[i].id] = true
			continue
		}
		if m := reBoldOnly.FindStringSubmatch(t); m != nil {
			elements[i] = element{headingElement, m[1], archonSlug(m[1])}
			ids[elements[i].id] = true
			continue
		}
		if !opens {
			continue
		}
		body := strings.TrimSpace(strings.TrimPrefix(t, ">"))
		m := reStrong.FindStringSubmatchIndex(body)
		if m == nil || m[0] != 0 {
			continue
		}
		title := body[m[2]:m[3]]
		if !reCalloutTitle.MatchString(title) {
			continue
		}
		consumed[i] = 1
		// a subtitle opening the box's next words joins the title
		if rest := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(body[m[1]:]), `\`)); rest == "" {
			for j := i + 1; j < len(lines) && strings.HasPrefix(strings.TrimSpace(lines[j]), ">"); j++ {
				next := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[j]), ">"))
				if next == "" || next == `\` {
					continue
				}
				if s := reStrong.FindStringSubmatchIndex(next); s != nil && s[0] == 0 {
					title += " — " + next[s[2]:s[3]]
					consumed[j] = 1
				}
				break
			}
		}
		elements[i] = element{boxElement, title, archonSlug(escaped(title))}
		ids[elements[i].id] = true
	}
	for i, line := range lines {
		if elements[i].kind == headingElement {
			continue
		}
		body := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), ">"))
		// a list item opening on a bold term opens as a paragraph does
		lead := 0
		if strings.HasPrefix(body, "- ") {
			lead = 2
		}
		for k, m := range reStrong.FindAllStringSubmatchIndex(body, -1) {
			if k < consumed[i] {
				continue
			}
			term := strings.TrimSuffix(strings.TrimSuffix(body[m[2]:m[3]], ":"), `\`)
			if strings.ContainsAny(term, "<[*") {
				continue
			}
			id := archonSlug(escaped(term))
			if ids[id] {
				id = ""
			} else {
				ids[id] = true
			}
			// a title ends on a colon, or stands alone: "**Stealth** represents…" is a word
			titled := strings.HasSuffix(body[m[2]:m[3]], ":") || strings.Trim(body[m[1]:], ` \`) == ""
			if k == 0 && m[0] == lead && titled && elements[i].kind == noElement {
				elements[i] = element{termElement, term, id}
			}
		}
	}
	return elements
}

// parseArchon splits a document of archon.vekn.net at its headings and, given
// terms, at a box or a paragraph opening on a bold title (**Combat Ends:**),
// each linked to the id the page gives it, else to the section before it. A
// heading's ID is its number in a numbered document, else its id. The text's
// links within the page are given the page, to be cited without it.
func parseArchon(md []byte, doc string, terms bool) []Section {
	page := archonPages[doc]
	var lines []string
	for line := range strings.Lines(string(md)) {
		lines = append(lines, strings.TrimRight(line, "\n"))
	}
	elements := archonElements(lines)
	sb := &sectionBuilder{doc: doc, ids: map[string]int{}}
	linked := map[string]bool{}
	link := func(id string) string {
		if id != "" && !linked[id] {
			linked[id] = true
			return page + "#" + id
		}
		for j := len(sb.stack) - 1; j >= 0; j-- {
			if sb.stack[j].URL != "" {
				return sb.stack[j].URL
			}
		}
		return page
	}
	level := 1
	for i, line := range lines {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "![") {
			continue
		}
		e := elements[i]
		switch {
		case e.kind == headingElement && reHeading.MatchString(t):
			level = len(reHeading.FindStringSubmatch(t)[1])
			// the title of a document that is not chaptered heads no section
			if level == 1 && doc != DocRulebook {
				continue
			}
			// an infraction's Definition, Examples, Penalty… stay in the
			// numbered section they belong to
			if doc != DocRulebook && reNumber.FindString(e.title) == "" && len(sb.stack) > 0 &&
				reNumber.FindString(sb.stack[len(sb.stack)-1].Title) != "" && level > sb.stack[len(sb.stack)-1].Level {
				sb.text.WriteString("**" + e.title + "**\n")
				level = sb.stack[len(sb.stack)-1].Level
				continue
			}
			k := e.id
			if doc != DocRulebook {
				k = key(e.title)
			}
			sb.open(level, e.title, k, link(e.id))
			sb.text.WriteString(line + "\n")
			continue
		case terms && e.kind != noElement:
			title := strings.TrimSuffix(strings.TrimSpace(e.title), ":")
			sb.open(level+1, title, archonSlug(title), link(e.id))
		}
		sb.text.WriteString(line + "\n")
	}
	sb.flush()
	for i := range sb.sections {
		s := &sb.sections[i]
		// the heading line opens the text; the title says it
		if first, rest, ok := strings.Cut(s.Text, "\n"); ok && reHeading.MatchString(strings.TrimSpace(first)) {
			s.Text = strings.TrimSpace(rest)
		} else if reHeading.MatchString(strings.TrimSpace(s.Text)) {
			s.Text = ""
		}
		s.Text = strings.ReplaceAll(s.Text, "](#", "]("+page+"#")
	}
	return sb.sections
}
