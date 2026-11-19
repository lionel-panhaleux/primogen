package bundle

import (
	"cmp"
	"context"
	"crypto/sha3"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"go.yaml.in/yaml/v3"
)

// Upstream says where Build reads.
type Upstream struct {
	Cards       string // vtes.json
	Sets        string // expansions.json, the names and release dates the cards' printings refer to by code
	TWDA        string // twda.json
	RulingsHead string // answers the vtes-rulings commit to read, as a bare SHA
	RulingsFile string // a file at that commit, formatted with the SHA and the file name
	Archive     string // the archive's bulk file
	// primogen's own knowledge/: a local directory, or else a head and files
	// as for the rulings.
	KnowledgeDir  string
	KnowledgeHead string
	KnowledgeFile string
	// archon-vibe, whose help pages are the rulebook, the tournament rules and
	// the judges' guide: the last commit touching a path, and files at one
	ArchonCommits string // the last commit touching a path, formatted with the path
	ArchonFile    string
}

var KRCG = Upstream{
	Cards:         "https://static.krcg.org/data/vtes.json",
	Sets:          "https://static.krcg.org/data/v5/expansions.json",
	TWDA:          "https://static.krcg.org/data/twda.json",
	RulingsHead:   "https://api.github.com/repos/vtes-biased/vtes-rulings/commits/main",
	RulingsFile:   "https://raw.githubusercontent.com/vtes-biased/vtes-rulings/%s/rulings/%s",
	Archive:       "https://usenet.krcg.org/archive.json.gz",
	KnowledgeHead: "https://api.github.com/repos/lionel-panhaleux/primogen/commits/main",
	KnowledgeFile: "https://raw.githubusercontent.com/lionel-panhaleux/primogen/%s/knowledge/%s",
	ArchonCommits: "https://api.github.com/repos/vtes-biased/archon-vibe/commits?per_page=1&path=%s",
	ArchonFile:    "https://raw.githubusercontent.com/vtes-biased/archon-vibe/%s/frontend/src/lib/help-content/%s",
}

// The names of Bundle.Sources, which Probe answers in the same terms.
const (
	SourceCards     = "cards of static.krcg.org"
	SourceTWDA      = "TWDA of static.krcg.org"
	SourceRulings   = "vtes-rulings"
	SourceKnowledge = "knowledge"
	SourceArchive   = "rulings archive to"
	SourceArchon    = "rules of archon-vibe"
)

const untaggedShown = 50

// Build reads every upstream but the archive, fetched apart by FetchArchive.
// The warnings name what needs a person.
func Build(ctx context.Context, up Upstream, archive *Archive) (*Bundle, []string, error) {
	var warnings []string
	warn := func(format string, a ...any) { warnings = append(warnings, fmt.Sprintf(format, a...)) }
	var raw []krcgCard
	cardsVersion, err := fetchJSON(ctx, up.Cards, &raw)
	if err != nil {
		return nil, nil, err
	}
	if len(raw) == 0 {
		return nil, nil, fmt.Errorf("%s: no cards", up.Cards)
	}
	var sets []krcgSet
	if body, err := fetch(ctx, up.Sets, ""); err != nil {
		return nil, nil, err
	} else if err := json.Unmarshal(body, &sets); err != nil {
		return nil, nil, fmt.Errorf("%s: %w", up.Sets, err)
	}
	db, err := fetchRulings(ctx, up)
	if err != nil {
		return nil, nil, err
	}
	knowledgeVersion, read, err := knowledge(ctx, up)
	if err != nil {
		return nil, nil, err
	}
	var history, advanced, procedure, analysis, advice, terms, shorthand, catalog, classification, tags, modules, theory []byte
	for name, body := range map[string]*[]byte{
		"card-changes.md": &history, "rules/advanced-rules.md": &advanced,
		"guide/rules-answer.md": &procedure, "guide/deck-analysis.md": &analysis, "guide/deck-advice.md": &advice,
		"game-terms.json": &terms, "shorthand.json": &shorthand, "llms.txt": &catalog, "classification.json": &classification,
		"guide/module-tags.json": &tags, "guide/modules.md": &modules, "guide/theory.md": &theory,
	} {
		if *body, err = read(name); err != nil {
			return nil, nil, err
		}
	}
	gameTerms, err := readGameTerms(terms)
	if err != nil {
		return nil, nil, err
	}
	labels, err := readClassification(classification)
	if err != nil {
		return nil, nil, err
	}
	var short struct{ Shorthand []string }
	if err := json.Unmarshal(shorthand, &short); err != nil {
		return nil, nil, fmt.Errorf("shorthand.json: %w", err)
	}
	archonCommit, err := lastCommit(ctx, up, archonDocuments)
	if err != nil {
		return nil, nil, err
	}
	if renderer, err := lastCommit(ctx, up, archonRendererPath); err != nil {
		return nil, nil, err
	} else if !strings.HasPrefix(renderer, archonRenderer) {
		warn("archon's renderer (%s) changed at %s, after the %s whose ids archon.go reproduces: "+
			"the links of the rulebook, the tournament rules and the judges' guide may miss (testdata/upstream/refresh.sh, then TestArchonLinks)",
			archonRendererPath, renderer[:7], archonRenderer)
	}
	var rulebook, tournament, judges []byte
	for name, body := range map[string]*[]byte{"vtes-rules.md": &rulebook, "tournament-rules.md": &tournament, "judges-guide.md": &judges} {
		if *body, err = fetch(ctx, fmt.Sprintf(up.ArchonFile, archonCommit, name), ""); err != nil {
			return nil, nil, err
		}
	}

	b := &Bundle{
		Built: time.Now().UTC(),
		Sources: []Source{
			{SourceCards, cardsVersion},
			{SourceRulings, db.commit[:7]},
			{SourceKnowledge, knowledgeVersion},
			{SourceArchive, archive.LastPost},
			{SourceArchon, archonCommit[:7]},
		},
		Cards: make([]Card, len(raw)),
		Sections: slices.Concat(parseArchon(rulebook, DocRulebook, true), parseAdvanced(advanced),
			parseArchon(tournament, DocTournament, false), parseArchon(judges, DocJudges, false)),
		Posts:          map[string]Post{},
		RulesAnswer:    string(procedure),
		DeckAnalysis:   string(analysis),
		DeckAdvice:     string(advice),
		GameTerms:      gameTerms,
		Shorthand:      short.Shorthand,
		Catalog:        string(catalog),
		Classification: labels,
	}
	setByCode := map[string]krcgSet{}
	for _, s := range sets {
		setByCode[s.Code] = s
	}
	for i, k := range raw {
		if b.Cards[i], err = k.card(setByCode); err != nil {
			return nil, nil, fmt.Errorf("%s: %w", up.Cards, err)
		}
	}
	if b.Modules, err = readModules(tags, modules, theory, b.Cards, warn); err != nil {
		return nil, nil, err
	}
	decks, twdaVersion, err := fetchDecks(ctx, up.TWDA, b.Cards, warn)
	if err != nil {
		return nil, nil, err
	}
	b.Decks, b.Sources = decks, append(b.Sources, Source{SourceTWDA, twdaVersion})
	if unmatched := NewRelatives(b).Unmatched(); len(unmatched) > 0 {
		warn("%d TWDA decks since %s, not in the archetype labels of %s, match no archetype (a new one, or a pass over the labels): %s",
			len(unmatched), labels.Since, labels.Generated, strings.Join(unmatched, "; "))
	}
	if untagged := b.untagged(); len(untagged) > 0 {
		shown := untagged[:min(len(untagged), untaggedShown)]
		warn("%d library cards the TWDA decks since %s play have no module tag (knowledge/guide/module-tags.json); the %d most played, with their decks: %s",
			len(untagged), labels.Since, len(shown), strings.Join(shown, "; "))
	}
	b.attachRulings(db, archive, warn)
	b.attachChanges(history, warn)
	b.keepPosts(archive, warn)
	b.linkRulebook(warn)
	return b, warnings, nil
}

// Probe answers the versions of every source but the archive, as Build would
// record them, at the cost of a few small requests.
func Probe(ctx context.Context, up Upstream) ([]Source, error) {
	resp, err := open(ctx, http.MethodHead, up.Cards, "")
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	cards, err := lastModified(resp)
	if err != nil {
		return nil, err
	}
	if resp, err = open(ctx, http.MethodHead, up.TWDA, ""); err != nil {
		return nil, err
	}
	resp.Body.Close()
	twda, err := lastModified(resp)
	if err != nil {
		return nil, err
	}
	rulings, err := headCommit(ctx, up.RulingsHead)
	if err != nil {
		return nil, err
	}
	version, _, err := knowledge(ctx, up)
	if err != nil {
		return nil, err
	}
	archon, err := lastCommit(ctx, up, archonDocuments)
	if err != nil {
		return nil, err
	}
	return []Source{
		{SourceCards, cards}, {SourceTWDA, twda}, {SourceRulings, rulings[:7]}, {SourceKnowledge, version}, {SourceArchon, archon[:7]},
	}, nil
}

// readGameTerms answers the English terms of game-terms.json, every key of
// its categories.
func readGameTerms(raw []byte) ([]string, error) {
	var categories map[string]map[string]json.RawMessage
	if err := json.Unmarshal(raw, &categories); err != nil {
		return nil, fmt.Errorf("game-terms.json: %w", err)
	}
	var out []string
	for name, terms := range categories {
		if name != "_meta" {
			out = append(out, slices.Collect(maps.Keys(terms))...)
		}
	}
	slices.Sort(out)
	return out, nil
}

// knowledge says which version of knowledge/ a build reads, and reads its files
// at that version.
func knowledge(ctx context.Context, up Upstream) (string, func(string) ([]byte, error), error) {
	if up.KnowledgeDir != "" {
		return "local", func(name string) ([]byte, error) { return os.ReadFile(filepath.Join(up.KnowledgeDir, name)) }, nil
	}
	commit, err := headCommit(ctx, up.KnowledgeHead)
	if err != nil {
		return "", nil, err
	}
	return commit[:7], func(name string) ([]byte, error) { return fetch(ctx, fmt.Sprintf(up.KnowledgeFile, commit, name), "") }, nil
}

var reCommit = regexp.MustCompile(`^[0-9a-f]{40}$`)

const (
	archonDocuments    = "frontend/src/lib/help-content"
	archonRendererPath = "frontend/src/lib/markdown.ts"
	// archonRenderer is the commit of archon's renderer whose ids archon.go reproduces
	archonRenderer = "9248307"
)

// lastCommit asks GitHub's API for the last commit touching a path of
// archon-vibe, so that a push elsewhere in the app moves nothing.
func lastCommit(ctx context.Context, up Upstream, path string) (string, error) {
	u := fmt.Sprintf(up.ArchonCommits, path)
	body, err := fetch(ctx, u, "")
	if err != nil {
		return "", err
	}
	var commits []struct{ SHA string }
	if err := json.Unmarshal(body, &commits); err != nil {
		return "", fmt.Errorf("%s: %w", u, err)
	}
	if len(commits) == 0 || !reCommit.MatchString(commits[0].SHA) {
		return "", fmt.Errorf("%s: no commit", u)
	}
	return commits[0].SHA, nil
}

// headCommit asks GitHub's API for a branch's commit, as a bare SHA.
func headCommit(ctx context.Context, u string) (string, error) {
	head, err := fetch(ctx, u, "application/vnd.github.sha")
	if err != nil {
		return "", err
	}
	commit := strings.TrimSpace(string(head))
	if !reCommit.MatchString(commit) {
		return "", fmt.Errorf("%s: not a commit: %q", u, commit)
	}
	return commit, nil
}

type rulingsDB struct {
	commit     string
	rulings    map[string][]rulingEntry     // "<card id or group id>|<name>" → entries
	groups     map[string]map[string]string // group key → member key → prefix
	references map[string]string            // label → URL
}

// A rulingEntry is a plain string, or a {text, overrides} map giving some
// group members their own wording, which keeps the group's references.
type rulingEntry struct {
	Text      string
	Overrides map[string]string
}

func (e *rulingEntry) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind == yaml.ScalarNode {
		return n.Decode(&e.Text)
	}
	type plain rulingEntry
	return n.Decode((*plain)(e))
}

func fetchRulings(ctx context.Context, up Upstream) (*rulingsDB, error) {
	commit, err := headCommit(ctx, up.RulingsHead)
	if err != nil {
		return nil, err
	}
	db := &rulingsDB{commit: commit}
	for name, v := range map[string]any{"rulings.yaml": &db.rulings, "groups.yaml": &db.groups, "references.yaml": &db.references} {
		u := fmt.Sprintf(up.RulingsFile, db.commit, name)
		body, err := fetch(ctx, u, "")
		if err != nil {
			return nil, err
		}
		if err := yaml.Unmarshal(body, v); err != nil {
			return nil, fmt.Errorf("%s: %w", u, err)
		}
	}
	if len(db.rulings) == 0 {
		return nil, fmt.Errorf("vtes-rulings %s: no rulings", commit)
	}
	return db, nil
}

var (
	reBracket         = regexp.MustCompile(`\[([^\]\n]+)\]`)
	reReferenceShaped = regexp.MustCompile(`\[[^\]\n]*\s\d{6,}[\w-]*\]`)
	reReminder        = regexp.MustCompile(`(?i)\s*\[REMINDER\]\s*$`)
	reCard            = regexp.MustCompile(`\{(\d+)\|[^}]*\}`)
	reBareCard        = regexp.MustCompile(`\{[^}|]*\}`)
	reCardMark        = regexp.MustCompile(`<([^<>\n]+)>`) // krcg 5 marks the cards a text names
)

// attachRulings files each ruling under the card it names, and a group's
// rulings under each member, behind the member's prefix or in its own wording.
func (b *Bundle) attachRulings(db *rulingsDB, archive *Archive, warn func(string, ...any)) {
	refs := map[string]Reference{}
	for label, u := range db.references {
		ref := Reference{Label: label, URL: u, Archive: archive.link(u)}
		if ref.Archive == "" && (reArchive.MatchString(u) || reForum.MatchString(u)) {
			ref.Unarchived = true
			warn("reference %s: the archive does not hold %s", label, u)
		}
		refs[label] = ref
	}
	byID := map[string]*Card{}
	for i := range b.Cards {
		byID[strconv.Itoa(b.Cards[i].ID)] = &b.Cards[i]
	}

	for _, key := range slices.Sorted(maps.Keys(db.rulings)) {
		target, group, _ := strings.Cut(key, "|")
		members := map[string]string{} // card id → prefix
		if strings.HasPrefix(target, "G") {
			for member, prefix := range db.groups[key] {
				id, _, _ := strings.Cut(member, "|")
				members[id] = prefix
			}
			if len(members) == 0 {
				warn("%s: no member in groups.yaml", key)
			}
		} else {
			members[target], group = "", ""
		}
		for _, e := range db.rulings[key] {
			reminder := reReminder.MatchString(e.Text)
			text := strings.TrimSpace(reReminder.ReplaceAllString(e.Text, ""))
			if reBareCard.MatchString(text) || strings.Contains(text, "_") || duplicateReference(text) {
				warn("%s: a ruling not in the rulings site's canonical form, its link may miss: %q", key, text)
			}
			ruling := Ruling{Group: group, Reminder: reminder}
			for _, m := range reBracket.FindAllStringSubmatch(text, -1) {
				if ref, ok := refs[m[1]]; ok {
					ruling.References = append(ruling.References, ref)
				} else if reReferenceShaped.MatchString(m[0]) {
					warn("%s: %s is absent from references.yaml", key, m[0])
				}
			}
			if len(ruling.References) == 0 && !reminder {
				warn("%s: a ruling citing no reference: %q", key, text)
			}
			overrides := map[string]string{}
			for k, v := range e.Overrides {
				id, _, _ := strings.Cut(k, "|")
				overrides[id] = v
			}
			for id, prefix := range members {
				card := byID[id]
				if card == nil {
					warn("%s: card %s is absent from the card data", key, id)
					continue
				}
				r := ruling
				r.Text = cmp.Or(overrides[id], strings.TrimSpace(prefix+" "+text))
				r.Link = rulingLink(card.ID, text)
				card.Rulings = append(card.Rulings, r)
			}
		}
	}
}

var reArchiveLink = regexp.MustCompile(`https://usenet\.krcg\.org/t/[^/\s)]+/#[\w-]+`)

// linkRulebook joins each ruling's rulebook reference, an id of
// archon.vekn.net's rules page, to the section that id opens.
func (b *Bundle) linkRulebook(warn func(string, ...any)) {
	opens := map[string]string{}
	for _, s := range b.Sections {
		// a section the page gives no id of its own links the one it is in, which comes first
		if s.Doc == DocRulebook && opens[s.URL] == "" {
			opens[s.URL] = s.ID
		}
	}
	missed := map[string]bool{}
	for i := range b.Cards {
		for j := range b.Cards[i].Rulings {
			refs := b.Cards[i].Rulings[j].References
			for k := range refs {
				if !strings.HasPrefix(refs[k].Label, "RBK ") {
					continue
				}
				if refs[k].Section = opens[refs[k].URL]; refs[k].Section == "" && !missed[refs[k].Label] {
					missed[refs[k].Label] = true
					warn("reference %s: no section of the rulebook opens at %s", refs[k].Label, refs[k].URL)
				}
			}
		}
	}
}

// keepPosts copies from the archive every post a ruling or the advanced rules
// cite, which the evidence for a question quotes.
func (b *Bundle) keepPosts(archive *Archive, warn func(string, ...any)) {
	keep := func(link, citer string) {
		if _, ok := b.Posts[link]; ok {
			return
		}
		if p, ok := archive.Post(link); ok {
			b.Posts[link] = p
		} else if citer != "" {
			warn("%s: the archive does not hold %s", citer, link)
		}
	}
	for _, c := range b.Cards {
		for _, r := range c.Rulings {
			for _, ref := range r.References {
				if ref.Archive != "" {
					keep(ref.Archive, "")
				}
			}
		}
	}
	for _, s := range b.Sections {
		for _, n := range s.Notes {
			for _, link := range reArchiveLink.FindAllString(n, -1) {
				keep(link, "advanced-rules.md "+s.ID)
			}
		}
	}
}

func duplicateReference(text string) bool {
	seen := map[string]bool{}
	for _, m := range reReferenceShaped.FindAllString(text, -1) {
		if seen[m] {
			return true
		}
		seen[m] = true
	}
	return false
}

// rulingLink must hash as the rulings site does (vtesrulings.utils.hash_text),
// which first normalizes a text attachRulings only checks is canonical.
func rulingLink(card int, text string) string {
	sum := sha3.SumSHAKE128([]byte(reCard.ReplaceAllString(text, "{$1}")), 5)
	return fmt.Sprintf("https://rulings.krcg.org/index.html?uid=%d#r-%s", card, base32.StdEncoding.EncodeToString(sum))
}

// krcgCard is a card of krcg 5's model (krcg/models.py): a crypt card or a
// library card, by its kind.
type krcgCard struct {
	ID           int    `json:"id"`
	Kind         string `json:"kind"`
	PrintedName  string `json:"printed_name"`
	Unicity      string `json:"unicity_suffix"`
	Suffix       string `json:"suffix"`
	NameVariants []struct {
		Name string `json:"name"`
	} `json:"name_variants"`
	URL   string   `json:"url"`
	Types []string `json:"types"`
	Text  string   `json:"text"`
	I18n  map[string]struct {
		Name string `json:"name"`
		Text string `json:"text"`
	} `json:"i18n"`
	Prints []struct {
		Set struct {
			Code string `json:"code"`
		} `json:"set"`
		Occurrences []struct {
			Type   string `json:"type"`
			Bundle string `json:"bundle"`
			Date   string `json:"date"`
		} `json:"occurrences"`
	} `json:"prints"`
	Legal  string `json:"legal"`
	Banned string `json:"banned"`

	Advanced    bool     `json:"advanced"`
	Clan        string   `json:"clan"`
	Path        string   `json:"path"`
	Capacity    int      `json:"capacity"`
	Group       string   `json:"group"`
	Disciplines []string `json:"disciplines"`
	Title       string   `json:"title"`

	Cost *struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"` // a number, or "X"
	} `json:"cost"`
	BurnOption            bool     `json:"burn_option"`
	ClanRequirement       []string `json:"clan_requirement"`
	PathRequirement       []string `json:"path_requirement"`
	DisciplineRequirement *struct {
		Type        string   `json:"type"` // Mono, Combo, or Choice
		Disciplines []string `json:"disciplines"`
	} `json:"discipline_requirement"`
}

type krcgSet struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	Bundles     map[string]struct {
		Name        string `json:"name"`
		ReleaseDate string `json:"release_date"`
	} `json:"bundles"`
}

func (k krcgCard) card(sets map[string]krcgSet) (Card, error) {
	c := Card{
		ID: k.ID, Name: k.PrintedName, PrintedName: k.PrintedName, Image: k.URL, Types: k.Types,
		Text: reCardMark.ReplaceAllString(k.Text, "$1"), Legal: k.Legal, Banned: k.Banned, Advanced: k.Advanced,
	}
	if k.Suffix != "" {
		c.Name += " (" + k.Suffix + ")"
	}
	switch k.Kind {
	case "Crypt":
		if k.Clan != "" {
			c.Clans = []string{k.Clan}
		}
		c.Disciplines, c.Capacity, c.Group, c.Title, c.Path = k.Disciplines, k.Capacity, k.Group, k.Title, k.Path
	case "Library":
		c.Clans, c.Path, c.BurnOption = k.ClanRequirement, strings.Join(k.PathRequirement, ", "), k.BurnOption
		if d := k.DisciplineRequirement; d != nil {
			c.Combo, c.Multidisc = d.Type == "Combo", d.Type == "Choice"
			for _, disc := range d.Disciplines {
				if disc == "fli" { // as the card text writes it
					disc = "FLIGHT"
				}
				c.Disciplines = append(c.Disciplines, disc)
			}
		}
		if k.Cost != nil {
			value := strings.Trim(string(k.Cost.Value), `"`)
			switch k.Cost.Type {
			case "Blood":
				c.BloodCost = value
			case "Pool":
				c.PoolCost = value
			case "Conviction":
				c.ConvictionCost = value
			default:
				return c, fmt.Errorf("card %d: cost type %q", k.ID, k.Cost.Type)
			}
		}
	default:
		return c, fmt.Errorf("card %d: kind %q, not krcg 5's model", k.ID, k.Kind)
	}
	var translated []string
	for _, lang := range slices.Sorted(maps.Keys(k.I18n)) {
		t := k.I18n[lang]
		c.Translations = append(c.Translations, Translation{lang, t.Name, reCardMark.ReplaceAllString(t.Text, "$1")})
		if t.Name != k.PrintedName {
			translated = append(translated, t.Name)
		}
	}
	names := []string{k.PrintedName}
	if k.Unicity != "" {
		names = append(names, k.PrintedName+" ("+k.Unicity+")")
	}
	for _, v := range k.NameVariants {
		names = append(names, v.Name)
	}
	add := func(a string) {
		if a != c.Name && a != c.PrintedName && !slices.Contains(c.Aliases, a) {
			c.Aliases = append(c.Aliases, a)
		}
	}
	for _, n := range names {
		// krcg gives some spellings only with the suffix: "Lucinde (G3)"
		bare := n
		for _, suffix := range []string{k.Suffix, k.Unicity} {
			if suffix != "" {
				bare = strings.TrimSuffix(bare, " ("+suffix+")")
			}
		}
		// the variants include the translated names, which stay behind every English one
		if n == "" || slices.Contains(translated, bare) {
			continue
		}
		add(n)
		add(bare)
		if k.Advanced {
			add(bare + " (ADV)")
		}
	}
	for _, p := range k.Prints {
		set, ok := sets[p.Set.Code]
		if !ok {
			return c, fmt.Errorf("card %d: set %q is not in the sets", k.ID, p.Set.Code)
		}
		printing := Printing{Set: set.Name, Date: set.ReleaseDate}
		// a printing only in singles or in precons released apart is dated by them
		var dates, bundles []string
		for _, o := range p.Occurrences {
			date := o.Date
			if o.Type == "Precon" {
				date = set.Bundles[o.Bundle].ReleaseDate
				bundles = append(bundles, set.Bundles[o.Bundle].Name)
			}
			dates = append(dates, date)
		}
		if len(dates) > 0 && !slices.Contains(dates, "") {
			printing.Date = slices.Min(dates)
			if len(bundles) == len(dates) {
				printing.Set += " (" + strings.Join(bundles, ", ") + ")"
			}
		}
		c.Sets = append(c.Sets, printing)
	}
	slices.SortStableFunc(c.Sets, func(a, b Printing) int { return strings.Compare(a.Date, b.Date) })
	return c, nil
}

var client = &http.Client{Timeout: 10 * time.Minute}

func open(ctx context.Context, method, u, accept string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "primogen")
	if accept != "" {
		req.Header.Set("Accept", accept)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("%s: %s", u, resp.Status)
	}
	return resp, nil
}

func fetch(ctx context.Context, u, accept string) ([]byte, error) {
	resp, err := open(ctx, http.MethodGet, u, accept)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// fetchJSON decodes the body and returns the server's date for it.
func fetchJSON(ctx context.Context, u string, v any) (string, error) {
	resp, err := open(ctx, http.MethodGet, u, "")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return "", fmt.Errorf("%s: %w", u, err)
	}
	return lastModified(resp)
}

func lastModified(resp *http.Response) (string, error) {
	date, err := http.ParseTime(resp.Header.Get("Last-Modified"))
	if err != nil {
		return "", fmt.Errorf("%s: no Last-Modified to date the data: %w", resp.Request.URL, err)
	}
	return date.UTC().Format("2006-01-02 15:04 UTC"), nil
}
