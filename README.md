# Primogen

A [VTES](https://www.vekn.net) (Vampire: The Eternal Struggle) councilor for
any LLM assistant: rules answers that quote what settles them, deck analysis
and advice, a meta picture with numbers.

It is an MCP server with no LLM on the server side: it returns verbatim card
texts, rulings and their sources, and your own assistant writes the answer.

The hosted server is `https://mcp.krcg.org/mcp`. It needs no account and no
sign-in.

## Tools

- `vtes_cards`: a card's official text, history and rulings.
- `vtes_rules_case`: the evidence for a rules question in one call: the cards, their rulings with the original posts, the governing rules sections, what no source covers.
- `vtes_rules_search`, `vtes_rules_read`: the rulebook, the advanced rules, the tournament rules and the judges' guide, searched by keyword or read by section.
- `vtes_rulings_archive`: full-text search of the rulings archive, each post's author marked when they were the sitting rules director.

## Sources

- [KRCG](https://static.krcg.org) for cards data
- The [rulings database](https://rulings.krcg.org)
- The [usenet archive](https://usenet.krcg.org)
- The [rulebook](https://www.vekn.net/rulebook), the [tournament rules](https://www.vekn.net/tournament-rules) and the [judges' guide](https://www.vekn.net/judges-guide) of the V:EKN
- The [advanced rules](https://vtes-biased.github.io/vtes-advanced-rules/)

## Connect your assistant

Use the publicly available MCP `https://mcp.krcg.org/mcp`.
Check your LLM documentation on how to add *Connectors* or *MCP* tools.

- Claude: `Customize > Connectors > Add > Add custom connector`
- Claude Code: `claude mcp add --transport http primogen https://mcp.krcg.org/mcp`
- Gemini: `Settings > Connected Apps > Custom apps`

## Run it locally

With [Go](https://go.dev/dl/) 1.27 or later:

```sh
go install github.com/lionel-panhaleux/primogen@latest
primogen query cards "Freak Drive"
```

You can then add it as a local MCP, for example with Claude Code:

```sh
claude mcp add primogen -- primogen stdio
```

## Contributing

This software was agentic-coded. Check [`AGENTS.md`](AGENTS.md).

## License

Licensed under the [MIT License](LICENSE).

Portions of the materials are the copyrights and trademarks of Paradox Interactive AB,
and are used with permission. All rights reserved.
For more information please visit [white-wolf.com](http://www.white-wolf.com).

![Dark Pack](dark-pack.png)
