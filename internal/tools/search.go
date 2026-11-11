package tools

import (
	"cmp"
	"math"
	"slices"
	"strings"
	"unicode"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

var stopwords = map[string]bool{}

func init() {
	for _, w := range strings.Fields(`a an and are as at be but by can could did do does for from had has have he her his
		how i if in into is it its may me might my no not of on or our she should so than that the their them then there
		these they this those to too was we were what when where which while who why will with would you your yes
		after before during only also any each other one same does doesn just get got still even play played plays
		card cards use used using`) {
		stopwords[w] = true
	}
}

// terms are the words a text is searched by: folded, stopwords out, each cut
// to a crude stem so that "blocks" finds "blocked".
func terms(s string) []string {
	var out []string
	for _, w := range strings.FieldsFunc(bundle.Fold(s), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
	}) {
		w = strings.Trim(w, "'")
		w = strings.TrimSuffix(w, "'s")
		if len(w) < 2 || stopwords[w] {
			continue
		}
		out = append(out, stem(w))
	}
	return out
}

func stem(w string) string {
	if len(w) <= 4 {
		return w
	}
	for _, suffix := range []string{"ing", "ed", "es", "s"} {
		if base, ok := strings.CutSuffix(w, suffix); ok && len(base) >= 4 {
			return base
		}
	}
	return w
}

// bm25 ranks documents for a weighted query.
type bm25 struct {
	docs  []map[string]float64 // term frequencies, a title's terms counted thrice
	lens  []float64
	df    map[string]int
	avgLn float64
}

func newBM25(texts []string, titles []string) *bm25 {
	ix := &bm25{df: map[string]int{}}
	total := 0.0
	for i, text := range texts {
		tf := map[string]float64{}
		for _, t := range terms(text) {
			tf[t]++
		}
		for _, t := range terms(titles[i]) {
			tf[t] += 3
		}
		n := 0.0
		for t, f := range tf {
			ix.df[t]++
			n += f
		}
		ix.docs = append(ix.docs, tf)
		ix.lens = append(ix.lens, n)
		total += n
	}
	ix.avgLn = total / float64(max(1, len(texts)))
	return ix
}

type scored struct {
	i     int
	score float64
}

// rank answers the documents matching the query, best first; keep, when
// given, says which documents may be answered.
func (ix *bm25) rank(query map[string]float64, keep func(int) bool) []scored {
	const k1, b = 1.2, 0.75
	n := float64(len(ix.docs))
	var out []scored
	for i, tf := range ix.docs {
		if keep != nil && !keep(i) {
			continue
		}
		s := 0.0
		for t, w := range query {
			f := tf[t]
			if f == 0 {
				continue
			}
			idf := math.Log(1 + (n-float64(ix.df[t])+0.5)/(float64(ix.df[t])+0.5))
			s += w * idf * f * (k1 + 1) / (f + k1*(1-b+b*ix.lens[i]/ix.avgLn))
		}
		if s > 0 {
			out = append(out, scored{i, s})
		}
	}
	slices.SortStableFunc(out, func(a, b scored) int { return cmp.Compare(b.score, a.score) })
	return out
}

func query(s string, weight float64, into map[string]float64) map[string]float64 {
	if into == nil {
		into = map[string]float64{}
	}
	for _, t := range terms(s) {
		into[t] = max(into[t], weight)
	}
	return into
}

// excerpt keeps the paragraphs of a text that share most terms with the
// query, in their order, up to about limit bytes, marking each cut with […].
func excerpt(text string, q map[string]float64, limit int) (string, bool) {
	if len(text) <= limit {
		return text, false
	}
	paras := strings.Split(text, "\n\n")
	type para struct {
		i     int
		score float64
	}
	ranked := make([]para, len(paras))
	for i, p := range paras {
		s := 0.0
		seen := map[string]bool{}
		for _, t := range terms(p) {
			if w, ok := q[t]; ok && !seen[t] {
				s += w
				seen[t] = true
			}
		}
		ranked[i] = para{i, s}
	}
	slices.SortStableFunc(ranked, func(a, b para) int { return cmp.Compare(b.score, a.score) })
	keep := map[int]bool{}
	size := 0
	for _, p := range ranked {
		if size > 0 && (p.score == 0 || size+len(paras[p.i]) > limit) {
			continue
		}
		keep[p.i] = true
		size += len(paras[p.i])
		if size >= limit {
			break
		}
	}
	var out []string
	cut := false
	for i, p := range paras {
		if !keep[i] {
			cut = true
			continue
		}
		if cut && len(out) > 0 {
			out = append(out, "[…]")
		}
		cut = false
		if len(p) > limit {
			p = p[:strings.LastIndex(p[:limit], " ")+1] + "[…]"
		}
		out = append(out, p)
	}
	return strings.Join(out, "\n\n"), true
}
