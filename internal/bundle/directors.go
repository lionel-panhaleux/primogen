package bundle

import "slices"

// director is a seat of the rules directors' succession. Where its start is
// unconfirmed, from is the first day it may have begun and sure the first day
// it certainly had; where its end is, until is the last day certain and
// possibly the last day it may still have been held.
type director struct {
	name     string
	spelled  []string // as the archive's authors spell them
	from     string
	sure     string
	until    string
	possibly string
}

var directors = []director{
	{name: "Thomas R. Wylie", spelled: []string{"Thomas R Wylie", "Tom Wylie"}, from: "1994-12-15", until: "1996-07-29"},
	{name: "Shawn F. Carnes", spelled: []string{"Shawn F. Carnes"}, from: "1996-07-29", until: "1996-10-18"},
	{name: "Jon Wilkie", spelled: []string{"Jon Wilkie"}, from: "1996-10-18", until: "1997-02-24"},
	{name: "L. Scott Johnson", spelled: []string{"L. Scott Johnson", "LSJ", "LSJ (VtES Rep)", "L. Scott Johnson (Rulemonger)"}, from: "1997-02-24", until: "2011-07-06"},
	{name: "Pascal Bertrand", spelled: []string{"Pascal Bertrand"}, from: "2011-07-06", until: "2016-12-04", possibly: "2017-01-11"},
	{name: "Vincent Ripoll (Ankha)", spelled: []string{"Ankha"}, from: "2016-12-04", sure: "2017-01-11"},
}

// Seat says where a post's author stood when they wrote it.
type Seat int

const (
	Player   Seat = iota // never a rules director
	Sitting              // the sitting rules director
	Possibly             // within a handover whose date is unconfirmed
	Other                // a rules director at another time
)

// Standing is a post's author's seat then, and how to say it.
type Standing struct {
	Seat  Seat
	Label string
}

// when dates a day, or the span of days it fell in.
func when(first, last string) string {
	if last == "" || last == first {
		return "on " + first
	}
	return "between " + first + " and " + last
}

// Role says where a post's author stood on the date they wrote it.
func Role(author, date string) Standing {
	for _, d := range directors {
		if !slices.Contains(d.spelled, author) {
			continue
		}
		sure, end := max(d.from, d.sure), max(d.until, d.possibly)
		switch {
		case date == "":
			return Standing{Other, d.name + ", a rules director (the post is undated)"}
		case date >= sure && (d.until == "" || date < d.until):
			return Standing{Sitting, d.name + ", the sitting rules director"}
		case date >= d.from && date < end:
			return Standing{Possibly, d.name + ", possibly the sitting rules director: the handover's date is unconfirmed"}
		case date < d.from:
			return Standing{Other, d.name + ", not yet the rules director (seated " + when(d.from, d.sure) + ")"}
		}
		return Standing{Other, d.name + ", a former rules director (succeeded " + when(d.until, d.possibly) + ")"}
	}
	return Standing{Player, author + ", a player, not a rules director"}
}
