// Package bundle is primogen's data: what the tools serve, built from the
// upstream sources and saved as one gzipped JSON file.
package bundle

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Bundle struct {
	Built   time.Time
	Sources []Source
	Cards   []Card
}

// Source dates what the bundle read: a commit, a file date, the newest post.
type Source struct {
	Name    string
	Version string
}

type Card struct {
	ID          int
	Name        string
	PrintedName string
	// Aliases holds every other English spelling the card answers to.
	Aliases []string
	Image   string

	Types       []string
	Clans       []string
	Disciplines []string
	// krcg's flags for a library card with several disciplines: one effect
	// combines them (Combo) or effects use different ones (Multidisc).
	Combo          bool
	Multidisc      bool
	Capacity       int
	Group          string
	Title          string
	Path           string
	BloodCost      string
	PoolCost       string
	ConvictionCost string
	BurnOption     bool

	Text         string
	Translations []Translation
	Sets         []Printing
	Legal        string // the date from which it is tournament legal
	Banned       string

	Rulings []Ruling
	Changes []Change
}

// Names are the card's English names.
func (c *Card) Names() []string {
	return slices.Concat([]string{c.Name, c.PrintedName}, c.Aliases)
}

type Translation struct {
	Lang string
	Name string
	Text string
}

type Printing struct {
	Set  string
	Date string
}

type Ruling struct {
	// Text is the database's wording, card tokens included ({100788|Freak Drive}).
	Text  string
	Link  string
	Group string
	// Reminder marks a ruling that restates card text or a core rule.
	Reminder   bool
	References []Reference
}

type Reference struct {
	Label string
	URL   string
	// Archive is the post in usenet.krcg.org, empty when the archive does not hold it.
	Archive string
	// Unarchived marks a newsgroup or forum post the archive should hold and does not.
	Unarchived bool
}

// Change is an entry of the history of card and rules changes.
type Change struct {
	Date         string // a day, or a month
	Announcement string
	URL          string
	Named        string // the name as the history writes it
	Text         string
	// Shared marks a name several cards answer to, the entry filed under each.
	Shared bool
}

func Load(path string) (*Bundle, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	z, err := gzip.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	var b Bundle
	if err := json.NewDecoder(z).Decode(&b); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if len(b.Cards) == 0 {
		return nil, fmt.Errorf("%s: no cards", path)
	}
	return &b, nil
}

// Save writes beside the target and renames, so a reader never sees half a bundle.
func (b *Bundle) Save(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".bundle-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	z := gzip.NewWriter(tmp)
	if err := errors.Join(json.NewEncoder(z).Encode(b), z.Close(), tmp.Chmod(0o644), tmp.Close()); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}

var folder = strings.NewReplacer("’", "'", "‘", "'", "“", `"`, "”", `"`)

// Fold is the form a name is matched in: no accents, no case, straight quotes,
// single spaces.
func Fold(s string) string {
	s, _, _ = transform.String(transform.Chain(norm.NFKD, runes.Remove(runes.In(unicode.Mn))), s)
	return strings.Join(strings.Fields(strings.ToLower(folder.Replace(s))), " ")
}

// Distance is the Levenshtein distance over runes.
func Distance(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	prev, cur := make([]int, len(rb)+1), make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ra); i++ {
		cur[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev, cur = cur, prev
	}
	return prev[len(rb)]
}
