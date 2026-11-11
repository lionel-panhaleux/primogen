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
	RulingsHead string // answers the vtes-rulings commit to read, as a bare SHA
	RulingsFile string // a file at that commit, formatted with the SHA and the file name
	Archive     string // the archive's bulk file
	// primogen's own knowledge/: a local directory, or else a head and files
	// as for the rulings.
	KnowledgeDir  string
	KnowledgeHead string
	KnowledgeFile string
	// vekn.net's pages
	TournamentRules string
	JudgesGuide     string
	Rulebook        string
}

var KRCG = Upstream{
	Cards:           "https://static.krcg.org/data/v4/vtes.json",
	RulingsHead:     "https://api.github.com/repos/vtes-biased/vtes-rulings/commits/main",
	RulingsFile:     "https://raw.githubusercontent.com/vtes-biased/vtes-rulings/%s/rulings/%s",
	Archive:         "https://usenet.krcg.org/archive.json.gz",
	KnowledgeHead:   "https://api.github.com/repos/lionel-panhaleux/primogen/commits/main",
	KnowledgeFile:   "https://raw.githubusercontent.com/lionel-panhaleux/primogen/%s/knowledge/%s",
	TournamentRules: "https://www.vekn.net/tournament-rules",
	JudgesGuide:     "https://www.vekn.net/judges-guide",
	Rulebook:        rulebookURL,
}

// The names of Bundle.Sources, which Probe answers in the same terms.
const (
	SourceCards      = "cards of static.krcg.org"
	SourceRulings    = "vtes-rulings"
	SourceKnowledge  = "knowledge"
	SourceArchive    = "rulings archive to"
	SourceTournament = "tournament rules of vekn.net"
	SourceJudges     = "judges' guide of vekn.net"
)

// Build reads every upstream but the archive, fetched apart by FetchArchive.
// vekn.net's pages failing, the previous bundle's copy of them stands in,
// when there is one. The warnings name what needs a person.
func Build(ctx context.Context, up Upstream, archive *Archive, previous *Bundle) (*Bundle, []string, error) {
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
	db, err := fetchRulings(ctx, up)
	if err != nil {
		return nil, nil, err
	}
	knowledgeVersion, read, err := knowledge(ctx, up)
	if err != nil {
		return nil, nil, err
	}
	var history, rulebook, advanced, procedure, terms []byte
	for name, body := range map[string]*[]byte{
		"card-changes.md": &history, "rules/rulebook.md": &rulebook, "rules/advanced-rules.md": &advanced,
		"guide/rules-answer.md": &procedure, "game-terms.json": &terms,
	} {
		if *body, err = read(name); err != nil {
			return nil, nil, err
		}
	}
	gameTerms, err := readGameTerms(terms)
	if err != nil {
		return nil, nil, err
	}
	tournament, tournamentVersion, err := fetchVEKNOr(ctx, up.TournamentRules, DocTournament, SourceTournament, previous, warn)
	if err != nil {
		return nil, nil, err
	}
	judges, judgesVersion, err := fetchVEKNOr(ctx, up.JudgesGuide, DocJudges, SourceJudges, previous, warn)
	if err != nil {
		return nil, nil, err
	}
	anchors := map[string]bool{}
	for _, u := range db.references {
		if a, ok := strings.CutPrefix(u, rulebookURL+"#"); ok {
			anchors[strings.ToLower(a)] = true
		}
	}
	if page, err := fetchAnchors(ctx, up.Rulebook); err != nil {
		warn("the rulebook's page: %v; its sections link only the anchors the rulings cite", err)
	} else {
		maps.Copy(anchors, page)
	}

	b := &Bundle{
		Built: time.Now().UTC(),
		Sources: []Source{
			{SourceCards, cardsVersion},
			{SourceRulings, db.commit[:7]},
			{SourceKnowledge, knowledgeVersion},
			{SourceArchive, archive.LastPost},
			{SourceTournament, tournamentVersion},
			{SourceJudges, judgesVersion},
		},
		Cards:       make([]Card, len(raw)),
		Sections:    slices.Concat(parseRulebook(rulebook), parseAdvanced(advanced), tournament, judges),
		Posts:       map[string]Post{},
		RulesAnswer: string(procedure),
		GameTerms:   gameTerms,
	}
	for i, k := range raw {
		b.Cards[i] = k.card()
	}
	linkRulebook(b.Sections, anchors)
	b.attachRulings(db, archive, warn)
	b.attachChanges(history, warn)
	b.keepPosts(archive, warn)
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
	rulings, err := headCommit(ctx, up.RulingsHead)
	if err != nil {
		return nil, err
	}
	version, _, err := knowledge(ctx, up)
	if err != nil {
		return nil, err
	}
	probed := []Source{{SourceCards, cards}, {SourceRulings, rulings[:7]}, {SourceKnowledge, version}}
	// a vekn.net page that fails is left out: it moves nothing
	for _, p := range []struct{ name, url, doc string }{{SourceTournament, up.TournamentRules, DocTournament}, {SourceJudges, up.JudgesGuide, DocJudges}} {
		if _, v, err := fetchVEKN(ctx, p.url, p.doc); err == nil {
			probed = append(probed, Source{p.name, v})
		}
	}
	return probed, nil
}

// fetchVEKNOr reads a vekn.net page, or keeps the previous bundle's copy of it
// when the page fails or holds under half its sections (a changed template).
func fetchVEKNOr(ctx context.Context, u, doc, source string, previous *Bundle, warn func(string, ...any)) ([]Section, string, error) {
	sections, version, err := fetchVEKN(ctx, u, doc)
	if previous == nil {
		return sections, version, err
	}
	var kept []Section
	for _, s := range previous.Sections {
		if s.Doc == doc {
			kept = append(kept, s)
		}
	}
	if err == nil && len(sections) < len(kept)/2 {
		err = fmt.Errorf("%s: %d sections, where the live bundle has %d", u, len(sections), len(kept))
	}
	if err == nil || len(kept) == 0 {
		return sections, version, err
	}
	for _, s := range previous.Sources {
		if s.Name == source {
			version = s.Version
		}
	}
	warn("%v: the live bundle's copy is kept", err)
	return kept, version, nil
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

type krcgCard struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	SortName     string   `json:"_name"`
	PrintedName  string   `json:"printed_name"`
	NameVariants []string `json:"name_variants"`
	Aka          []string `json:"aka"`
	URL          string   `json:"url"`
	Types        []string `json:"types"`
	Clans        []string `json:"clans"`
	Disciplines  []string `json:"disciplines"`
	Combo        bool     `json:"combo"`
	Multidisc    bool     `json:"multidisc"`
	Capacity     int      `json:"capacity"`
	Group        string   `json:"group"`
	Title        string   `json:"title"`
	Path         string   `json:"path"`
	BloodCost    string   `json:"blood_cost"`
	PoolCost     string   `json:"pool_cost"`
	Conviction   string   `json:"conviction_cost"`
	BurnOption   bool     `json:"burn_option"`
	CardText     string   `json:"card_text"`
	I18n         map[string]struct {
		Name     string `json:"name"`
		CardText string `json:"card_text"`
	} `json:"_i18n"`
	Sets map[string][]struct {
		ReleaseDate string `json:"release_date"`
	} `json:"sets"`
	OrderedSets []string `json:"ordered_sets"`
	Legality    string   `json:"legality"`
	Banned      string   `json:"banned"`
}

func (k krcgCard) card() Card {
	c := Card{
		ID: k.ID, Name: k.Name, PrintedName: k.PrintedName, Image: k.URL,
		Types: k.Types, Clans: k.Clans, Disciplines: k.Disciplines, Combo: k.Combo, Multidisc: k.Multidisc,
		Capacity: k.Capacity, Group: k.Group, Title: k.Title, Path: k.Path,
		BloodCost: k.BloodCost, PoolCost: k.PoolCost, ConvictionCost: k.Conviction, BurnOption: k.BurnOption,
		Text: k.CardText, Legal: k.Legality, Banned: k.Banned,
	}
	for _, a := range slices.Concat([]string{k.SortName}, k.NameVariants, k.Aka) {
		if a != "" && a != k.Name && a != k.PrintedName && !slices.Contains(c.Aliases, a) {
			c.Aliases = append(c.Aliases, a)
		}
	}
	for _, lang := range slices.Sorted(maps.Keys(k.I18n)) {
		t := k.I18n[lang]
		c.Translations = append(c.Translations, Translation{lang, t.Name, t.CardText})
	}
	for _, set := range k.OrderedSets {
		p := Printing{Set: set}
		if prints := k.Sets[set]; len(prints) > 0 {
			p.Date = prints[0].ReleaseDate
		}
		c.Sets = append(c.Sets, p)
	}
	return c
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
