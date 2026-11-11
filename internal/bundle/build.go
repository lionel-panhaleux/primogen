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
}

var KRCG = Upstream{
	Cards:         "https://static.krcg.org/data/v4/vtes.json",
	RulingsHead:   "https://api.github.com/repos/vtes-biased/vtes-rulings/commits/main",
	RulingsFile:   "https://raw.githubusercontent.com/vtes-biased/vtes-rulings/%s/rulings/%s",
	Archive:       "https://usenet.krcg.org/archive.json.gz",
	KnowledgeHead: "https://api.github.com/repos/lionel-panhaleux/primogen/commits/main",
	KnowledgeFile: "https://raw.githubusercontent.com/lionel-panhaleux/primogen/%s/knowledge/%s",
}

// The names of Bundle.Sources, which Probe answers in the same terms.
const (
	SourceCards     = "cards of static.krcg.org"
	SourceRulings   = "vtes-rulings"
	SourceKnowledge = "knowledge"
	SourceArchive   = "rulings archive to"
)

// Build reads every upstream but the archive, fetched apart by FetchArchive.
// The warnings name what needs a person.
func Build(ctx context.Context, up Upstream, archive *Archive) (*Bundle, []string, error) {
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
	knowledgeVersion, history, err := readKnowledge(ctx, up, "card-changes.md")
	if err != nil {
		return nil, nil, err
	}

	b := &Bundle{
		Built: time.Now().UTC(),
		Sources: []Source{
			{SourceCards, cardsVersion},
			{SourceRulings, db.commit[:7]},
			{SourceKnowledge, knowledgeVersion},
			{SourceArchive, archive.LastPost},
		},
		Cards: make([]Card, len(raw)),
	}
	for i, k := range raw {
		b.Cards[i] = k.card()
	}
	var warnings []string
	warn := func(format string, a ...any) { warnings = append(warnings, fmt.Sprintf(format, a...)) }
	b.attachRulings(db, archive.index, warn)
	b.attachChanges(history, warn)
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
	knowledge := "local"
	if up.KnowledgeDir == "" {
		if knowledge, err = headCommit(ctx, up.KnowledgeHead); err != nil {
			return nil, err
		}
		knowledge = knowledge[:7]
	}
	return []Source{{SourceCards, cards}, {SourceRulings, rulings[:7]}, {SourceKnowledge, knowledge}}, nil
}

// readKnowledge reads a file of knowledge/ and says which version it read.
func readKnowledge(ctx context.Context, up Upstream, name string) (string, []byte, error) {
	if up.KnowledgeDir != "" {
		body, err := os.ReadFile(filepath.Join(up.KnowledgeDir, name))
		return "local", body, err
	}
	commit, err := headCommit(ctx, up.KnowledgeHead)
	if err != nil {
		return "", nil, err
	}
	body, err := fetch(ctx, fmt.Sprintf(up.KnowledgeFile, commit, name), "")
	return commit[:7], body, err
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
func (b *Bundle) attachRulings(db *rulingsDB, archive archiveIndex, warn func(string, ...any)) {
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
