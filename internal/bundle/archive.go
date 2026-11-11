package bundle

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

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

// link is the archive's copy of the post a reference URL cites, if it holds it.
// Forum topic 123 is thread vekn-123, its posts anchored by their forum number.
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

// Archive is what a bundle reads of the rulings archive, the heaviest source:
// which posts it holds, and the date of its newest.
type Archive struct {
	index    archiveIndex
	LastPost string
}

// FetchArchive streams the bulk file thread by thread, keeping only the ids.
func FetchArchive(ctx context.Context, up Upstream) (*Archive, error) {
	u := up.Archive
	resp, err := open(ctx, http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fail := func(err error) (*Archive, error) { return nil, fmt.Errorf("%s: %w", u, err) }
	z, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fail(err)
	}
	dec := json.NewDecoder(z)
	if _, err := dec.Token(); err != nil {
		return fail(err)
	}
	index := archiveIndex{}
	var lastPost string
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
	if len(index) == 0 || len(lastPost) < len("2006-01-02") {
		return fail(fmt.Errorf("no threads or no LastPost"))
	}
	return &Archive{index, lastPost[:len("2006-01-02")]}, nil
}
