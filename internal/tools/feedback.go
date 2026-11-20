package tools

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Report is a user's report on an answer, filed as a public issue.
type Report struct {
	Question string `json:"question"`
	Wrong    string `json:"wrong"`
	Expected string `json:"expected"`
	Public   bool   `json:"public"`
}

// The size cap of each field, in characters.
const reportField = 2000

var feedbackSchema = &jsonschema.Schema{
	Type: "object",
	Properties: map[string]*jsonschema.Schema{
		"question": {Type: "string", Description: "the user's question, as asked"},
		"wrong":    {Type: "string", Description: "what was wrong in the answer, in the user's words where possible"},
		"expected": {Type: "string", Description: `what the user expected instead, or "not known"`},
		"public":   {Type: "boolean", Description: "true once the user, told that the report becomes a public GitHub issue, agreed to file it"},
	},
	Required: []string{"question", "wrong", "expected"},
}

// Filer files a report and answers the issue's link; from is the client's
// address when the server knows it. Its error is the model's to read.
type Filer func(ctx context.Context, r Report, from string) (string, error)

// Filed is the result's data, which a local install reads back.
type Filed struct {
	Issue string `json:"issue"`
}

func addFeedback(s *mcp.Server, l *Live) {
	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_feedback",
		Title: "Report a wrong primogen answer",
		Description: "Call when the user says an answer grounded in primogen's Vampire: The Eternal Struggle (VTES) tools was wrong, and agrees to report it. " +
			"Files the report (the question, what was wrong, what was expected) as a public issue on primogen's GitHub repository, which its maintainer reads, and returns the issue's link. " +
			"Before calling, tell the user the report will be public and show them its text; leave out names, emails and any other personal detail.",
		InputSchema: feedbackSchema,
		Annotations: &mcp.ToolAnnotations{DestructiveHint: new(false), OpenWorldHint: new(true)},
	}, func(ctx context.Context, req *mcp.CallToolRequest, in Report) (*mcp.CallToolResult, *Filed, error) {
		if !in.Public {
			return nil, nil, fmt.Errorf("not filed: tell the user the report becomes a public GitHub issue, show them its text without personal details, and call again with public true once they agree")
		}
		for name, v := range map[string]string{"question": in.Question, "wrong": in.Wrong, "expected": in.Expected} {
			if strings.TrimSpace(v) == "" {
				return nil, nil, fmt.Errorf("not filed: %s is empty", name)
			}
			if n := utf8.RuneCountInString(v); n > reportField {
				return nil, nil, fmt.Errorf("not filed: %s holds %d characters, at most %d: shorten it", name, n, reportField)
			}
		}
		if l.File == nil {
			return nil, nil, fmt.Errorf("not filed: this server files no reports; the user can open an issue at https://github.com/lionel-panhaleux/primogen/issues")
		}
		from := ""
		if req.Extra != nil && req.Extra.Header != nil {
			from = req.Extra.Header.Get("X-Real-IP")
		}
		link, err := l.File(ctx, in, from)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: "Filed as " + link + ". Give the user this link: they can follow the issue there."}}}, &Filed{Issue: link}, nil
	})
}
