// Package tools is primogen's MCP surface: the server, its instructions and
// its tools, the same for the hosted endpoint, stdio and `primogen query`.
package tools

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"sync/atomic"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

const instructions = `Primogen grounds answers about Vampire: The Eternal Struggle (VTES).
Never state a card's text or a ruling from memory: call vtes_cards for every card the answer relies on, and quote what it returns.
Each item carries its authority and its link; cite the link. When the result says no ruling covers a point, say the point is unsettled rather than filling the gap.`

// Live is the data the tools serve. A call reads one snapshot, so a swap lands
// between two calls.
type Live struct{ p atomic.Pointer[data] }

type data struct {
	b  *bundle.Bundle
	ix *cardIndex
}

func NewLive(b *bundle.Bundle) *Live {
	l := &Live{}
	l.Swap(b)
	return l
}

func (l *Live) Swap(b *bundle.Bundle) { l.p.Store(&data{b, newCardIndex(b.Cards)}) }

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
