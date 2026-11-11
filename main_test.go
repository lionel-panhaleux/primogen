package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

// testBundle is built once, from the real upstreams pinned in testdata/upstream.
var (
	testBundle   string
	testWarnings []string
)

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "primogen-test")
	if err == nil {
		err = buildTestBundle(filepath.Join(dir, "bundle.json.gz"))
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}

func buildTestBundle(path string) error {
	srv := httptest.NewServer(http.FileServer(http.Dir("testdata/upstream")))
	defer srv.Close()
	b, warnings, err := bundle.Build(context.Background(), bundle.Upstream{
		Cards:       srv.URL + "/vtes.json",
		RulingsHead: srv.URL + "/commit",
		RulingsFile: srv.URL + "/vtes-rulings/%s/%s",
		Archive:     srv.URL + "/archive.json.gz",
	}, "knowledge")
	if err != nil {
		return err
	}
	testBundle, testWarnings = path, warnings
	return b.Save(path)
}

func TestQueryCards(t *testing.T) {
	for _, tc := range []struct {
		args      []string // after query -bundle
		want, not []string
		err       string
	}{
		{args: []string{"cards", "Freak Drive"}, want: []string{
			"## Freak Drive",
			"> [FOR] Only usable if the action was blocked. Unlock this vampire.",
			"Requires [for].",
			"Cost: 1 blood.",
			// the rulings site's own ids for the card's four rulings, two of them group rulings
			"https://rulings.krcg.org/index.html?uid=100788#r-KHOYNPJM",
			"https://rulings.krcg.org/index.html?uid=100788#r-ROKWE7BA",
			"https://rulings.krcg.org/index.html?uid=100788#r-NKLCIFEB",
			"https://rulings.krcg.org/index.html?uid=100788#r-KFKPOACQ",
			"LSJ 19981028: https://usenet.krcg.org/t/JZHY1bTCAa0/#m1",
			"ANK 20190425: https://usenet.krcg.org/t/vekn-77581/#94652 (archived from https://www.vekn.net/",
			"Changes: no line of the history of card changes names this card.",
			"Data: cards of static.krcg.org ",
		}, not: []string{"Printings:"}},
		{args: []string{"cards", "Lord of Serenity"}, want: []string{
			"- 2018-07-19, RTR 20180719 (https://www.blackchantry.com/2018/07/18/rules-team-rulings-rtr-19-07-2018/): Gaining blood is now optional.\n",
		}},
		{args: []string{"-detail", "cards", "Freak Drive"}, want: []string{"Printings: Jyhad 1994-08-16", "Translation (fr), Pulsion monstrueuse:"}},
		{args: []string{"cards", "Pulsion monstrueuse"}, want: []string{"## Freak Drive", "#r-KHOYNPJM"}},
		{args: []string{"cards", "Freak Drvie"}, want: []string{`No card is named "Freak Drvie"; this is Freak Drive`, "## Freak Drive"}},
		{args: []string{"cards", "Blood Doll Supreme"}, want: []string{`No card is named "Blood Doll Supreme". Closest names: Blood Doll;`}, not: []string{"## "}},
		// Misdirection is "Distraction" in French
		{args: []string{"cards", "Distraction"}, want: []string{"## Distraction"}, not: []string{"## Misdirection"}},
		{args: []string{"cards", "Theo Bell"}, want: []string{"## Theo Bell (G2)\n", "## Theo Bell (G2 ADV)\n", "## Theo Bell (G6)\n"}},
		{args: []string{"cards", "Count Germaine (G4)"}, want: []string{`[the history names "Count Germaine", which several cards share`}},
		// the 1998 entry naming Kalinda predates the G6 card
		{args: []string{"cards", "Kalinda (G6)"}, want: []string{"Changes: no line of the history of card changes names this card."}},
		{args: []string{"cards", "Constanza Vinti"}, want: []string{`[the history names "Constanza Vinzi", no card exactly: possibly this one]`}},
		{args: []string{"cards", "Anson"}, want: []string{"## Anson (G1)", "Rulings: the rulings database has none for this card."}},
		{args: []string{"cards", " "}, err: "give 1 to 10 non-empty card names"},
	} {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			var stdout, stderr strings.Builder
			err := run(t.Context(), append([]string{"query", "-bundle", testBundle}, tc.args...), &stdout, &stderr)
			if tc.err != "" {
				if err == nil || !strings.Contains(err.Error(), tc.err) {
					t.Fatalf("error %v, want %q", err, tc.err)
				}
				return
			}
			if err != nil {
				t.Fatal(err, stderr.String())
			}
			out := stdout.String()
			for _, w := range tc.want {
				if !strings.Contains(out, w) {
					t.Errorf("missing %q in:\n%s", w, out)
				}
			}
			for _, n := range tc.not {
				if strings.Contains(out, n) {
					t.Errorf("unexpected %q in:\n%s", n, out)
				}
			}
		})
	}
}

func TestBuildWarnings(t *testing.T) {
	for _, w := range []string{
		`"Monacle of Clarity" names no card exactly; filed as possible under Monocle of Clarity`,
		`no card named in "- The promo vampires`,
	} {
		if !slices.ContainsFunc(testWarnings, func(s string) bool { return strings.Contains(s, w) }) {
			t.Errorf("no warning %q in:\n%s", w, strings.Join(testWarnings, "\n"))
		}
	}
}
