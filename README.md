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

## Sources

- [KRCG](https://static.krcg.org) for cards data
- The [rulings database](https://rulings.krcg.org)
- The [usenet archive](https://usenet.krcg.org)

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
