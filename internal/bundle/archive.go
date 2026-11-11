package bundle

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	reArchive = regexp.MustCompile(`^https://usenet\.krcg\.org/t/([^/]+)/(?:#(.*))?$`)
	reForum   = regexp.MustCompile(`^https://www\.vekn\.net/forum/(?:[^#?]*/)?(\d+)-[^/#?]*(?:\?[^#]*)?(?:#(\d+))?$`)
)

// Archive is the rulings archive, the heaviest source: every thread, and the
// date of the newest post.
type Archive struct {
	Threads  []Thread
	byID     map[string]*Thread
	LastPost string
}

type Thread struct {
	ID       string
	Title    string
	Messages []Message
}

type Message struct {
	ID     string // the forum's or BoardGameGeek's post number, if any
	Author string
	Date   string // YYYY-MM-DD
	Body   string
}

// Post is a message of the archive as the bundle keeps it: one a ruling or the
// advanced rules cite.
type Post struct {
	Link   string
	Thread string // its title
	Author string
	Date   string
	Body   string
}

// find answers the thread and message a link of the archive points at.
func (a *Archive) find(u string) (*Thread, int, bool) {
	m := reArchive.FindStringSubmatch(u)
	if m == nil {
		return nil, 0, false
	}
	t := a.byID[m[1]]
	if t == nil {
		return nil, 0, false
	}
	if n, found := strings.CutPrefix(m[2], "m"); found {
		i, err := strconv.Atoi(n)
		return t, i, err == nil && i < len(t.Messages)
	}
	if m[2] == "" {
		return t, 0, true
	}
	for i, msg := range t.Messages {
		if msg.ID == m[2] {
			return t, i, true
		}
	}
	return nil, 0, false
}

// link is the archive's copy of the post a reference URL cites, if it holds it.
// Forum topic 123 is thread vekn-123, its posts anchored by their forum number.
func (a *Archive) link(u string) string {
	if reArchive.MatchString(u) {
		if _, _, ok := a.find(u); ok {
			return u
		}
		return ""
	}
	if m := reForum.FindStringSubmatch(u); m != nil {
		link := "https://usenet.krcg.org/t/vekn-" + m[1] + "/"
		if m[2] != "" {
			link += "#" + m[2]
		}
		if _, _, ok := a.find(link); ok {
			return link
		}
	}
	return ""
}

// Post answers the message a link of the archive points at.
func (a *Archive) Post(u string) (Post, bool) {
	t, i, ok := a.find(u)
	if !ok {
		return Post{}, false
	}
	m := t.Messages[i]
	return Post{Link: u, Thread: t.Title, Author: m.Author, Date: m.Date, Body: m.Body}, true
}

// FetchArchive reads the bulk file.
func FetchArchive(ctx context.Context, up Upstream) (*Archive, error) {
	resp, err := open(ctx, http.MethodGet, up.Archive, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	a, err := ReadArchive(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", up.Archive, err)
	}
	return a, nil
}

// ReadArchive decodes a gzipped bulk file, thread by thread.
func ReadArchive(r io.Reader) (*Archive, error) {
	z, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(z)
	if _, err := dec.Token(); err != nil {
		return nil, err
	}
	a := &Archive{byID: map[string]*Thread{}}
	var lastPost string
	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch key {
		case "LastPost":
			err = dec.Decode(&lastPost)
		case "Threads":
			if _, err = dec.Token(); err != nil {
				return nil, err
			}
			for dec.More() {
				var t struct {
					ThreadId string
					Title    string
					Messages []struct{ Id, Author, Date, Body string }
				}
				if err := dec.Decode(&t); err != nil {
					return nil, err
				}
				th := Thread{ID: t.ThreadId, Title: t.Title, Messages: make([]Message, len(t.Messages))}
				for i, m := range t.Messages {
					th.Messages[i] = Message{ID: m.Id, Author: m.Author, Date: archiveDate(m.Date), Body: m.Body}
				}
				a.Threads = append(a.Threads, th)
			}
			_, err = dec.Token()
		default:
			err = dec.Decode(&json.RawMessage{})
		}
		if err != nil {
			return nil, err
		}
	}
	if len(a.Threads) == 0 || len(lastPost) < len(time.DateOnly) {
		return nil, fmt.Errorf("no threads or no LastPost")
	}
	for i := range a.Threads {
		a.byID[a.Threads[i].ID] = &a.Threads[i]
	}
	a.LastPost = lastPost[:len(time.DateOnly)]
	return a, nil
}

// archiveDate reads the archive's two date shapes ("Jul 6, 1997, 9:00:00 AM",
// the forum's without seconds), a narrow no-break space before AM or PM.
func archiveDate(s string) string {
	s = strings.ReplaceAll(s, " ", " ")
	for _, layout := range []string{"Jan 2, 2006, 3:04:05 PM", "Jan 2, 2006, 3:04 PM"} {
		if t, err := time.Parse(layout, s); err == nil {
			return t.Format(time.DateOnly)
		}
	}
	return ""
}
