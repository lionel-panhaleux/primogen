package main

import (
	"context"
	"strings"
	"testing"
)

func TestQueryCards(t *testing.T) {
	for _, tc := range []struct {
		name string
		want []string
		not  []string
	}{
		{"Freak Drive", []string{
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
			"ANK 20190425: https://usenet.krcg.org/t/vekn-77581/#94652",
		}, nil},
		{"Pulsion monstrueuse", []string{"## Freak Drive", "#r-KHOYNPJM"}, nil},
		{"Freak Drvie", []string{`No card is named "Freak Drvie"; this is Freak Drive`, "## Freak Drive"}, nil},
		{"Blood Doll Supreme", []string{`No card is named "Blood Doll Supreme". Closest names: Blood Doll;`}, []string{"## "}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var out, errOut strings.Builder
			err := run(context.Background(), []string{"query", "-bundle", "testdata/bundle.json.gz", "cards", tc.name}, &out, &errOut)
			if err != nil {
				t.Fatal(err, errOut.String())
			}
			for _, w := range tc.want {
				if !strings.Contains(out.String(), w) {
					t.Errorf("missing %q in:\n%s", w, out.String())
				}
			}
			for _, n := range tc.not {
				if strings.Contains(out.String(), n) {
					t.Errorf("unexpected %q in:\n%s", n, out.String())
				}
			}
		})
	}
}
