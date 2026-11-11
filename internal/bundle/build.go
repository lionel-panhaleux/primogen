package bundle

import (
	"compress/gzip"
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
	"unicode"

	"go.yaml.in/yaml/v3"
)

const (
	cardsURL    = "https://static.krcg.org/data/v4/vtes.json"
	rulingsRepo = "vtes-biased/vtes-rulings"
	archiveURL  = "https://usenet.krcg.org/archive.json.gz"
)

// Build reads every upstream and the knowledge directory. Warnings name what
// needs a person: a reference that resolves nowhere, a change naming no card.
func Build(ctx context.Context, knowledge string) (b *Bundle, warnings []string, err error) {
	warn := func(format string, a ...any) { warnings = append(warnings, fmt.Sprintf(format, a...)) }
	b = &Bundle{Built: time.Now().UTC()}

	var raw []krcgCard
	cardsDate, err := fetchJSON(ctx, cardsURL, &raw)
	if err != nil {
		return nil, nil, err
	}
	b.Sources = append(b.Sources, Source{"cards of static.krcg.org", cardsDate})

	sha, err := fetch(ctx, "https://api.github.com/repos/"+rulingsRepo+"/commits/main", "application/vnd.github.sha")
	if err != nil {
		return nil, nil, err
	}
	var rulings map[string][]any
	var groups map[string]map[string]string
	var references map[string]string
	for name, v := range map[string]any{"rulings.yaml": &rulings, "groups.yaml": &groups, "references.yaml": &references} {
		body, err := fetch(ctx, "https://raw.githubusercontent.com/"+rulingsRepo+"/"+string(sha)+"/rulings/"+name, "")
		if err != nil {
			return nil, nil, err
		}
		if err := yaml.Unmarshal(body, v); err != nil {
			return nil, nil, fmt.Errorf("%s: %w", name, err)
		}
	}
	b.Sources = append(b.Sources, Source{"vtes-rulings", string(sha[:7])})

	archive, lastPost, err := fetchArchive(ctx)
	if err != nil {
		return nil, nil, err
	}
	b.Sources = append(b.Sources, Source{"rulings archive to", lastPost[:10]})

	refs := map[string]Reference{}
	for label, u := range references {
		ref := Reference{Label: label, URL: u, Archive: archive.link(u)}
		if ref.Archive == "" && (reArchive.MatchString(u) || reForum.MatchString(u)) {
			ref.Unarchived = true
			warn("reference %s: the archive does not hold %s", label, u)
		}
		refs[label] = ref
	}
	// text is what the site hashes, display what this card shows: a group
	// ruling's text behind the card's prefix, or the card's own override.
	ruling := func(card int, text, display string, reminder bool, group string) Ruling {
		r := Ruling{Text: display, Group: group, Reminder: reminder, Link: rulingLink(card, text)}
		for _, m := range reBracket.FindAllStringSubmatch(text, -1) {
			if ref, ok := refs[m[1]]; ok {
				r.References = append(r.References, ref)
			} else if reReferenceShaped.MatchString(m[0]) {
				warn("ruling of %d cites %s, absent from references.yaml", card, m[0])
			}
		}
		if len(r.References) == 0 && !reminder {
			warn("ruling of %d cites no reference: %q", card, text)
		}
		return r
	}

	byID := map[int]*Card{}
	for _, k := range raw {
		b.Cards = append(b.Cards, k.card())
	}
	for i := range b.Cards {
		byID[b.Cards[i].ID] = &b.Cards[i]
	}
	keys := make([]string, 0, len(rulings))
	for k := range rulings {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, key := range keys {
		target, groupName, _ := strings.Cut(key, "|")
		members := map[string]string{target: ""}
		if strings.HasPrefix(target, "G") {
			members = map[string]string{}
			for member, prefix := range groups[key] {
				id, _, _ := strings.Cut(member, "|")
				members[id] = prefix
			}
		} else {
			groupName = ""
		}
		for _, entry := range rulings[key] {
			text, overrides := rulingEntry(entry)
			reminder := reReminder.MatchString(text)
			text = strings.TrimSpace(reReminder.ReplaceAllString(text, ""))
			if reBareCard.MatchString(text) || strings.Contains(text, "_") || hasDuplicateReference(text) {
				warn("%s: a ruling not in the site's canonical form, its link may miss: %q", key, text)
			}
			for id, prefix := range members {
				n, _ := strconv.Atoi(id)
				card := byID[n]
				if card == nil {
					warn("%s names card %s, absent from the card data", key, id)
					continue
				}
				display := overrides[id]
				if display == "" {
					display = strings.TrimSpace(prefix + " " + text)
				}
				card.Rulings = append(card.Rulings, ruling(n, text, display, reminder, groupName))
			}
		}
	}

	if err := attachChanges(b, filepath.Join(knowledge, "card-changes.md"), warn); err != nil {
		return nil, nil, err
	}
	return b, warnings, nil
}

var (
	reBracket         = regexp.MustCompile(`\[([^\]\n]+)\]`)
	reReferenceShaped = regexp.MustCompile(`\[[^\]\n]*\s\d{6,}[\w-]*\]`)
	reReminder        = regexp.MustCompile(`(?i)\s*\[REMINDER\]\s*$`)
	reCard            = regexp.MustCompile(`\{(\d+)\|[^}]*\}`)
	reBareCard        = regexp.MustCompile(`\{[^}|]*\}`)
)

func hasDuplicateReference(text string) bool {
	seen := map[string]bool{}
	for _, m := range reBracket.FindAllString(text, -1) {
		if seen[m] && reReferenceShaped.MatchString(m) {
			return true
		}
		seen[m] = true
	}
	return false
}

// rulingLink must hash as the rulings site does (vtesrulings.utils.hash_text),
// which first normalizes a text Build only checks is canonical.
func rulingLink(card int, text string) string {
	sum := sha3.SumSHAKE128([]byte(reCard.ReplaceAllString(text, "{$1}")), 5)
	return fmt.Sprintf("https://rulings.krcg.org/index.html?uid=%d#r-%s", card, base32.StdEncoding.EncodeToString(sum))
}

// rulingEntry reads a plain string or a {text, overrides} map, overrides being
// one card's own wording of a group ruling that keeps the group's references.
func rulingEntry(entry any) (string, map[string]string) {
	m, ok := entry.(map[string]any)
	if !ok {
		s, _ := entry.(string)
		return s, nil
	}
	overrides := map[string]string{}
	if o, ok := m["overrides"].(map[string]any); ok {
		for k, v := range o {
			id, _, _ := strings.Cut(k, "|")
			overrides[id], _ = v.(string)
		}
	}
	s, _ := m["text"].(string)
	return s, overrides
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
		Types: k.Types, Clans: k.Clans, Disciplines: k.Disciplines,
		Combo: k.Combo, Multidisc: k.Multidisc,
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

// archiveIndex holds, per thread, its message count and the forum's post numbers.
type archiveIndex map[string]archiveThread

type archiveThread struct {
	count int
	ids   map[string]bool
}

var (
	reArchive = regexp.MustCompile(`^https://usenet\.krcg\.org/t/([^/]+)/(?:#(.*))?$`)
	reForum   = regexp.MustCompile(`^https://www\.vekn\.net/forum/(?:[^#?]*/)?(\d+)-[^/#?]*(?:\?[^#]*)?(?:#(\d+))?$`)
)

// link is the archive's copy of the post a reference URL cites, if the archive
// holds it. A forum topic is archived as vekn-<topic>, its posts keeping their
// forum number as anchor (the newsgroup-archive README).
func (a archiveIndex) link(u string) string {
	if m := reArchive.FindStringSubmatch(u); m != nil {
		t, ok := a[m[1]]
		if !ok {
			return ""
		}
		if n, found := strings.CutPrefix(m[2], "m"); found {
			if i, err := strconv.Atoi(n); err != nil || i >= t.count {
				return ""
			}
		} else if m[2] != "" && !t.ids[m[2]] {
			return ""
		}
		return u
	}
	if m := reForum.FindStringSubmatch(u); m != nil {
		t, ok := a["vekn-"+m[1]]
		if !ok || (m[2] != "" && !t.ids[m[2]]) {
			return ""
		}
		link := "https://usenet.krcg.org/t/vekn-" + m[1] + "/"
		if m[2] != "" {
			link += "#" + m[2]
		}
		return link
	}
	return ""
}

// fetchArchive streams the bulk file thread by thread: whole, it is over 50 MB
// compressed, and only the ids are kept.
func fetchArchive(ctx context.Context) (archiveIndex, string, error) {
	resp, err := open(ctx, archiveURL, "")
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	z, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", archiveURL, err)
	}
	dec := json.NewDecoder(z)
	index := archiveIndex{}
	var lastPost string
	fail := func(err error) (archiveIndex, string, error) {
		return nil, "", fmt.Errorf("%s: %w", archiveURL, err)
	}
	if _, err := dec.Token(); err != nil {
		return fail(err)
	}
	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			return fail(err)
		}
		switch key {
		case "LastPost":
			err = dec.Decode(&lastPost)
		case "Threads":
			if _, err = dec.Token(); err != nil {
				return fail(err)
			}
			for dec.More() {
				var t struct {
					ThreadId string
					Messages []struct{ Id string }
				}
				if err := dec.Decode(&t); err != nil {
					return fail(err)
				}
				th := archiveThread{count: len(t.Messages), ids: map[string]bool{}}
				for _, m := range t.Messages {
					if m.Id != "" {
						th.ids[m.Id] = true
					}
				}
				index[t.ThreadId] = th
			}
			_, err = dec.Token()
		default:
			err = dec.Decode(&json.RawMessage{})
		}
		if err != nil {
			return fail(err)
		}
	}
	if len(index) == 0 || len(lastPost) < 10 {
		return fail(fmt.Errorf("no threads or no LastPost"))
	}
	return index, lastPost, nil
}

// attachChanges files each "- Card: change" line of the history under the
// cards it names; the history's paragraphs are rules changes and name none.
// A name matching no card exactly is filed as possible under the cards it is
// close to.
func attachChanges(b *Bundle, path string, warn func(string, ...any)) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	byName := map[string][]*Card{}
	byWord := map[string][]*Card{}
	for i := range b.Cards {
		c := &b.Cards[i]
		for _, n := range slices.Concat([]string{c.Name, c.PrintedName}, c.Aliases) {
			k := Fold(n)
			if !slices.Contains(byName[k], c) {
				byName[k] = append(byName[k], c)
			}
			for _, w := range words(k) {
				if !slices.Contains(byWord[w], c) {
					byWord[w] = append(byWord[w], c)
				}
			}
		}
	}
	// exact also says whether one name matched several cards (a vampire's groups).
	exact := func(name string) ([]*Card, bool) {
		if m := byName[Fold(name)]; m != nil {
			return m, len(m) > 1
		}
		if base, ok := strings.CutSuffix(name, " ADV"); ok {
			m := byName[Fold(base+" (ADV)")]
			return m, len(m) > 1
		}
		var all []*Card
		shared := false
		parts := reNameList.Split(name, -1)
		for _, part := range parts {
			m := byName[Fold(part)]
			if m == nil || len(parts) < 2 {
				return nil, false
			}
			all, shared = append(all, m...), shared || len(m) > 1
		}
		return all, shared
	}
	near := func(name string) []*Card {
		var found []*Card
		k := Fold(name)
		for n, cards := range byName {
			if Distance(k, n) <= max(1, len(k)/6) {
				found = append(found, cards...)
			}
		}
		if found == nil {
			for _, w := range words(k) {
				if m := byWord[w]; len(m) > 0 && len(m) <= 5 && (found == nil || len(m) < len(found)) {
					found = m
				}
			}
			found = slices.Clone(found)
		}
		slices.SortFunc(found, func(a, b *Card) int { return a.ID - b.ID })
		return slices.Compact(found)
	}

	var change Change
	for line := range strings.Lines(string(data)) {
		line = strings.TrimSpace(line)
		if date, ok := strings.CutPrefix(line, "# "); ok && reDate.MatchString(date) {
			change = Change{Date: date}
			continue
		}
		if m := reAnnouncement.FindStringSubmatch(line); m != nil && change.Date != "" && change.URL == "" {
			change.Announcement, change.URL = m[1], m[2]
			continue
		}
		item, ok := strings.CutPrefix(line, "- ")
		if !ok || change.Date == "" {
			continue
		}
		// A name can hold a colon itself (Powerbase: Madrid): the longest
		// prefix naming cards is the name.
		c := change
		var matched []*Card
		for i := range item {
			sep := 0
			if item[i] == ':' {
				sep = 1
			} else if strings.HasPrefix(item[i:], " - ") {
				sep = 3
			}
			if sep == 0 {
				continue
			}
			if c.Named == "" {
				c.Named, c.Text = item[:i], item[i+sep:]
			}
			if m, shared := exact(item[:i]); m != nil {
				matched, c.Shared, c.Named, c.Text = m, shared, item[:i], item[i+sep:]
			}
		}
		c.Text = strings.TrimSpace(c.Text)
		if c.Named == "" {
			c.Text = item
		}
		if matched == nil && c.Named != "" {
			matched, c.Possible = near(c.Named), true
		}
		// A change older than a card's first printing cannot concern it.
		matched = slices.DeleteFunc(slices.Clone(matched), func(card *Card) bool { return change.Date < card.Legal })
		c.Shared = c.Shared && len(matched) > 1
		if len(matched) == 0 {
			warn("card-changes.md %s: no card named in %q", change.Date, line)
			continue
		}
		if c.Possible {
			var names []string
			for _, m := range matched {
				names = append(names, m.Name)
			}
			warn("card-changes.md %s: %q names no card exactly; filed as possible under %s", change.Date, c.Named, strings.Join(names, "; "))
		}
		for _, card := range matched {
			card.Changes = append(card.Changes, c)
		}
	}
	return nil
}

func words(folded string) []string {
	return strings.FieldsFunc(folded, func(r rune) bool { return !unicode.IsLetter(r) })
}

var (
	reDate         = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	reAnnouncement = regexp.MustCompile(`^\[(.+)\]\((.+)\)$`)
	reNameList     = regexp.MustCompile(`,\s*(?:and\s+)?|\s+and\s+`)
)

var client = &http.Client{Timeout: 10 * time.Minute}

func open(ctx context.Context, u, accept string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
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
	resp, err := open(ctx, u, accept)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// fetchJSON decodes the body and returns the server's date for it.
func fetchJSON(ctx context.Context, u string, v any) (string, error) {
	resp, err := open(ctx, u, "")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return "", fmt.Errorf("%s: %w", u, err)
	}
	date, err := http.ParseTime(resp.Header.Get("Last-Modified"))
	if err != nil {
		return "", fmt.Errorf("%s: no Last-Modified to date the data: %w", u, err)
	}
	return date.UTC().Format(time.DateOnly), nil
}
