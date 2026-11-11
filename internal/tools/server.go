// Package tools is primogen's MCP surface: the server, its instructions and
// its tools, the same for the hosted endpoint, stdio and `primogen query`.
package tools

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

const instructions = `Primogen grounds answers about Vampire: The Eternal Struggle (VTES).
Never state a card's text or a ruling from memory: call vtes_cards for every card the answer relies on, and quote what it returns.
Each item carries its authority and its link; cite the link. When the result says no ruling covers a point, say the point is unsettled rather than filling the gap.`

func NewServer(b *bundle.Bundle) *mcp.Server {
	version := "dev"
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" {
		version = info.Main.Version
	}
	s := mcp.NewServer(&mcp.Implementation{Name: "primogen", Version: version}, &mcp.ServerOptions{Instructions: instructions})
	addCards(s, b)
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

// dataLine dates what an answer was read from.
func dataLine(b *bundle.Bundle) string {
	parts := make([]string, 0, len(b.Sources))
	for _, s := range b.Sources {
		parts = append(parts, s.Name+" "+s.Version)
	}
	return fmt.Sprintf("Data: %s; bundle built %s.", strings.Join(parts, "; "), b.Built.Format(time.DateOnly))
}
