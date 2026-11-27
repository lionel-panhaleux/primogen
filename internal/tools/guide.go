package tools

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

type guideInput struct {
	Topic    string `json:"topic"`
	Section  string `json:"section,omitempty"`
	Language string `json:"language,omitempty"`
	Filter   string `json:"filter,omitempty"`
}

var guideTopics = []string{
	"rules-answer",
	"deck-analysis",
	"deck-advice",
	"card-comparison",
	"meta-reading",
	"theory",
	"modules",
	"terms",
	"examples",
}

var guideSchema = &jsonschema.Schema{
	Type: "object",
	Properties: map[string]*jsonschema.Schema{
		"topic": {
			Type:        "string",
			Description: "what to read: rules-answer, deck-analysis, deck-advice, card-comparison, meta-reading, theory, modules, terms, or examples",
		},
		"section": {
			Type:        "string",
			Description: `optional: a "##" heading within theory or modules (e.g. "Combat defenses"), a category within terms (e.g. "combat"), or an example name within examples`,
		},
		"language": {
			Type:        "string",
			Description: "for terms: language code (EN, FR, ES, PT-BR, IT, JP, LATIN); EN when omitted",
		},
		"filter": {
			Type:        "string",
			Description: "for terms or examples: substring to narrow the entries (matches English or translated terms, or example name, title and text)",
		},
	},
	Required: []string{"topic"},
}

func addGuide(s *mcp.Server, l *Live) {
	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_guide",
		Title: "VTES guide, theory, modules, game terms and examples",
		Description: "Read primogen's curated Vampire: The Eternal Struggle (VTES) guide instead of inventing task procedures, module vocabulary, game-term translations or worked examples from memory. " +
			"Returns, by topic: task procedures (rules-answer, deck-analysis, deck-advice, card-comparison, meta-reading), VTES strategy theory (theory), " +
			"the module catalog and combat defenses (modules), game terms and their translations in EN, FR, ES, PT-BR, IT, JP or LATIN (terms; translations are unchecked consistency references, not official authorities), " +
			"or owner-graded served examples (examples).",
		InputSchema: guideSchema,
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: new(false)},
	}, func(_ context.Context, _ *mcp.CallToolRequest, in guideInput) (*mcp.CallToolResult, any, error) {
		d := l.p.Load()
		b := d.b
		topic := normGuideKey(in.Topic)
		var body, footer string
		var err error
		footer = "This is primogen's curated guide, not an official V:EKN rules source: ground any card text, ruling or rule in vtes_cards or vtes_rules_case, never from memory.\n\n"

		switch topic {
		case "rules-answer":
			body, err = guideProcedure(topic, b.RulesAnswer)
		case "deck-analysis":
			body, err = guideProcedure(topic, b.DeckAnalysis)
		case "deck-advice":
			body, err = guideProcedure(topic, b.DeckAdvice)
		case "card-comparison":
			body, err = guideProcedure(topic, b.CardComparison)
		case "meta-reading":
			body, err = guideProcedure(topic, b.MetaReading)
		case "theory":
			body, err = guideDocument(topic, b.Theory, "Authority: curated guide authored here (not an official V:EKN rules source).", in.Section)
		case "modules":
			body, err = guideDocument(topic, b.ModuleCatalog, `Authority: curated guide authored here (module catalog canonized by the owner; "Combat defenses" checked against card texts, not yet canonized by the owner; not an official V:EKN rules source).`, in.Section)
		case "terms":
			body, err = guideTerms(b.Terms, in.Language, in.Section, in.Filter)
			footer = "Use these terms for consistency across answers; translations are unchecked and not an official authority. Ground any card text, ruling or rule in vtes_cards or vtes_rules_case, never from memory.\n\n"
		case "examples":
			body, err = guideExamples(b, in.Section, in.Filter)
		default:
			return nil, nil, fmt.Errorf("unknown topic %q: choose one of %s", in.Topic, strings.Join(guideTopics, ", "))
		}
		if err != nil {
			return nil, nil, err
		}
		dataLine := DataLine(b)
		body = strings.TrimSpace(body) + "\n\n"
		const cutNotice = "\n\n(cut here at the 40,000-character bound: narrow with section or filter)\n\n"
		maxBody := budget - len(footer) - len(dataLine)
		if len(body) > maxBody {
			keep := max(maxBody-len(cutNotice), 0)
			if nl := strings.LastIndex(body[:keep], "\n"); nl > 0 {
				keep = nl
			}
			body = strings.ToValidUTF8(strings.TrimRight(body[:keep], "\n"), "") + cutNotice
		}
		return text(body + footer + dataLine), nil, nil
	})
}

func normGuideKey(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(bundle.Fold(s), " ", "-"), "_", "-")
}

func guideProcedure(topic, content string) (string, error) {
	if strings.TrimSpace(content) == "" {
		return "", fmt.Errorf("this data bundle does not carry %s, built before it included it: rebuild the bundle", topic)
	}
	return "Authority: curated guide authored here (not an official V:EKN rules source).\n\n" + strings.TrimSpace(content), nil
}

type mdSection struct {
	heading string
	text    string
}

func splitMarkdownSections(doc string) []mdSection {
	var sections []mdSection
	var cur *mdSection
	var buf strings.Builder
	flush := func() {
		if cur != nil {
			cur.text = strings.TrimSpace(buf.String())
			sections = append(sections, *cur)
		}
		buf.Reset()
	}
	for line := range strings.Lines(doc) {
		if h, ok := strings.CutPrefix(line, "## "); ok {
			flush()
			cur = &mdSection{heading: strings.TrimSpace(h)}
		}
		if cur != nil {
			buf.WriteString(line)
		}
	}
	flush()
	return sections
}

func guideDocument(topic, content, authority, section string) (string, error) {
	if strings.TrimSpace(content) == "" {
		return "", fmt.Errorf("this data bundle does not carry %s, built before it included it: rebuild the bundle", topic)
	}
	if strings.TrimSpace(section) == "" {
		return authority + "\n\n" + strings.TrimSpace(content), nil
	}
	sections := splitMarkdownSections(content)
	q := bundle.Fold(section)
	var exact, partial []mdSection
	headings := make([]string, len(sections))
	for i, s := range sections {
		headings[i] = s.heading
		h := bundle.Fold(s.heading)
		if h == q {
			exact = append(exact, s)
		} else if strings.Contains(h, q) {
			partial = append(partial, s)
		}
	}
	matched := exact
	if len(matched) == 0 {
		matched = partial
	}
	if len(matched) == 0 {
		return "", fmt.Errorf("no section of %s matches %q; sections: %s", topic, section, strings.Join(headings, "; "))
	}
	var out strings.Builder
	out.WriteString(authority + "\n\n")
	for i, s := range matched {
		if i > 0 {
			out.WriteString("\n\n")
		}
		out.WriteString(s.text)
	}
	return out.String(), nil
}

func glossaryLanguages(g bundle.Glossary) []string {
	seen := map[string]bool{}
	for k := range g.Sources {
		seen[k] = true
	}
	for _, cat := range g.Categories {
		for _, t := range cat.Terms {
			for k := range t.Translations {
				seen[k] = true
			}
		}
	}
	delete(seen, "EN")
	return append([]string{"EN"}, slices.Sorted(maps.Keys(seen))...)
}

func guideTerms(g bundle.Glossary, language, category, filter string) (string, error) {
	if len(g.Categories) == 0 {
		return "", fmt.Errorf("this data bundle holds no game-term translations, built before it carried them: rebuild the bundle")
	}
	langs := glossaryLanguages(g)
	lang := "EN"
	if strings.TrimSpace(language) != "" {
		want := strings.ToUpper(strings.TrimSpace(language))
		if !slices.Contains(langs, want) {
			return "", fmt.Errorf("unknown language %q: choose one of %s", language, strings.Join(langs, ", "))
		}
		lang = want
	}
	cats := g.Categories
	if strings.TrimSpace(category) != "" {
		want := normGuideKey(category)
		var matched []bundle.TermCategory
		names := make([]string, len(g.Categories))
		for i, c := range g.Categories {
			names[i] = c.Name
			if normGuideKey(c.Name) == want {
				matched = append(matched, c)
			}
		}
		if len(matched) == 0 {
			return "", fmt.Errorf("unknown terms category %q: choose one of %s", category, strings.Join(names, ", "))
		}
		cats = matched
	}
	f := bundle.Fold(filter)
	var out strings.Builder
	fmt.Fprintf(&out, "# VTES game terms (%s)\n\n", lang)
	out.WriteString("Authority: consistency reference, not an official authority. ")
	if lang == "EN" {
		if g.Description != "" {
			out.WriteString(g.Description)
		}
		fmt.Fprintf(&out, " Available languages: %s.\n", strings.Join(langs, ", "))
	} else {
		out.WriteString("Translations have not been checked against the translated rulebooks.")
		if src := g.Sources[lang]; src != "" {
			fmt.Fprintf(&out, " Recorded source for %s: %s (unchecked).", lang, src)
		}
		out.WriteString("\n")
	}
	shown := 0
	for _, cat := range cats {
		var lines []string
		for _, t := range cat.Terms {
			tr := t.Translations[lang]
			if f != "" && !strings.Contains(bundle.Fold(t.English), f) && (lang == "EN" || !strings.Contains(bundle.Fold(tr), f)) {
				continue
			}
			if lang == "EN" {
				lines = append(lines, "- "+t.English)
			} else if tr != "" {
				lines = append(lines, fmt.Sprintf("- %s: %s", t.English, tr))
			} else {
				lines = append(lines, fmt.Sprintf("- %s: (no translation recorded)", t.English))
			}
		}
		if len(lines) == 0 {
			continue
		}
		shown += len(lines)
		fmt.Fprintf(&out, "\n## %s\n\n", cat.Name)
		if cat.Name == "disciplines" && g.Partial != "" {
			fmt.Fprintf(&out, "%s\n\n", g.Partial)
		}
		for _, line := range lines {
			out.WriteString(line + "\n")
		}
	}
	if shown == 0 {
		fmt.Fprintf(&out, "\nNo game terms match filter %q.\n", filter)
	}
	return out.String(), nil
}

func guideExamples(b *bundle.Bundle, name, filter string) (string, error) {
	if len(b.Examples) == 0 {
		if strings.TrimSpace(b.Theory) == "" {
			return "", fmt.Errorf("this data bundle holds no served examples, built before it carried them: rebuild the bundle")
		}
		return "Authority: worked examples (principle 8).\n\nNo served examples are bundled.", nil
	}
	const auth = "Authority: worked example — owner-graded clean excerpt of a primogen answer, dated with its data (counts and shares age with the TWDA; re-checked by regression runs)."
	if strings.TrimSpace(name) != "" {
		want := normGuideKey(strings.TrimSuffix(strings.TrimSpace(name), ".md"))
		var exact, partial []bundle.Example
		names := make([]string, len(b.Examples))
		for i, ex := range b.Examples {
			names[i] = ex.Name
			n := normGuideKey(ex.Name)
			if n == want {
				exact = append(exact, ex)
			} else if strings.Contains(n, want) || strings.Contains(bundle.Fold(ex.Title), bundle.Fold(name)) {
				partial = append(partial, ex)
			}
		}
		matched := exact
		if len(matched) == 0 {
			matched = partial
		}
		if len(matched) == 0 {
			return "", fmt.Errorf("no served example matches %q; available: %s", name, strings.Join(names, ", "))
		}
		if len(matched) > 1 {
			mNames := make([]string, len(matched))
			for i, ex := range matched {
				mNames[i] = ex.Name
			}
			return "", fmt.Errorf("several served examples match %q: %s. Pass the exact name in section", name, strings.Join(mNames, ", "))
		}
		return auth + "\n\n" + strings.TrimSpace(matched[0].Text), nil
	}
	f := bundle.Fold(filter)
	var out strings.Builder
	out.WriteString("Authority: worked examples — owner-graded clean excerpts of primogen answers, dated with their data (counts and shares age with the TWDA; re-checked by regression runs).\n\n")
	out.WriteString("# Served examples\n\n")
	shown := 0
	for _, ex := range b.Examples {
		if f != "" && !strings.Contains(bundle.Fold(ex.Name+" "+ex.Title+" "+ex.Text), f) {
			continue
		}
		shown++
		fmt.Fprintf(&out, "## %s\nName: `%s`\n", ex.Title, ex.Name)
		if header, _, ok := strings.Cut(ex.Text, "\n---\n"); ok {
			if _, rest, ok := strings.Cut(strings.TrimSpace(header), "\n"); ok {
				out.WriteString(strings.TrimSpace(rest) + "\n")
			}
		}
		out.WriteString("\n")
	}
	if shown == 0 {
		fmt.Fprintf(&out, "No served examples match filter %q.\n", filter)
	} else {
		out.WriteString("Call vtes_guide with topic \"examples\" and section set to an example's name to read it whole.\n")
	}
	return out.String(), nil
}
