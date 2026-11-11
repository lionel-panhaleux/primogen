// Package tools is primogen's MCP surface: the server, its instructions and
// its tools, the same for the hosted endpoint, stdio and `primogen query`.
package tools

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

const instructions = `Primogen grounds answers about Vampire: The Eternal Struggle (VTES).
Never state a card's text, a ruling or a rule from memory: call vtes_cards for every card the answer relies on, and vtes_rules_case first for any rules question; quote what they return.
Each item carries its authority and its link; cite the link. When the result says no ruling covers a point, say the point is unsettled rather than filling the gap.`

// The bounds of a result, in characters: under the 50,000 past which Claude
// Code saves a result to a file and, detailed, about 20,000 tokens.
const (
	budget       = 40000
	detailBudget = 80000
)

// Live is the data the tools serve. A call reads one snapshot, so a swap lands
// between two calls.
type Live struct {
	p       atomic.Pointer[data]
	archive atomic.Pointer[archiveSearch]
	// LoadArchive fetches the archive on first use, for a server not handed one.
	LoadArchive func(context.Context) (*bundle.Archive, error)
	loading     sync.Mutex
}

type data struct {
	b     *bundle.Bundle
	ix    *cardIndex
	rules *rulesIndex
}

func NewLive(b *bundle.Bundle) *Live {
	l := &Live{}
	l.Swap(b)
	return l
}

func (l *Live) Swap(b *bundle.Bundle) {
	ix := newCardIndex(b.Cards)
	l.p.Store(&data{b, ix, newRulesIndex(b, ix)})
}

func (l *Live) Bundle() *bundle.Bundle { return l.p.Load().b }

// Version is the binary's module version, "dev" when it has none.
func Version() string {
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" {
		return info.Main.Version
	}
	return "dev"
}

func NewServer(l *Live) *mcp.Server {
	s := mcp.NewServer(&mcp.Implementation{Name: "primogen", Version: Version()}, &mcp.ServerOptions{Instructions: instructions})
	addCards(s, l)
	addCase(s, l)
	addRules(s, l)
	addArchive(s, l)
	return s
}

// Call runs one tool through an in-memory session, as `primogen query` does,
// and returns the text the model would read.
func Call(ctx context.Context, s *mcp.Server, tool string, args map[string]any) (string, error) {
	serverEnd, clientEnd := mcp.NewInMemoryTransports()
	ss, err := s.Connect(ctx, serverEnd, nil)
	if err != nil {
		return "", err
	}
	defer ss.Close()
	cs, err := mcp.NewClient(&mcp.Implementation{Name: "primogen-query"}, nil).Connect(ctx, clientEnd, nil)
	if err != nil {
		return "", err
	}
	defer cs.Close()
	res, err := cs.CallTool(ctx, &mcp.CallToolParams{Name: tool, Arguments: args})
	if err != nil {
		return "", err
	}
	var out strings.Builder
	for _, c := range res.Content {
		if t, ok := c.(*mcp.TextContent); ok {
			out.WriteString(t.Text)
		}
	}
	if res.IsError {
		return "", errors.New(out.String())
	}
	return out.String(), nil
}

// DataLine dates what an answer was read from.
func DataLine(b *bundle.Bundle) string {
	parts := make([]string, 0, len(b.Sources))
	for _, s := range b.Sources {
		parts = append(parts, s.Name+" "+s.Version)
	}
	return fmt.Sprintf("Data: %s; bundle built %s.", strings.Join(parts, "; "), b.Built.Format(time.DateOnly))
}
