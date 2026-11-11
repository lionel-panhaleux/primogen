# Primogen

A [VTES](https://www.vekn.net) (Vampire: The Eternal Struggle) councilor for
any LLM assistant: rules answers that quote what settles them, deck analysis
and advice, a meta picture with numbers.

It is an MCP server with no LLM on the server side: it returns verbatim card
texts, rulings and their sources, and your own assistant writes the answer.
Today it offers one tool, `vtes_cards`: a card's official text, costs,
legality, every ruling with its link and the post it cites, and the card's
history of changes. The data refreshes itself from the
[KRCG](https://static.krcg.org) card data and the
[rulings database](https://rulings.krcg.org).

The hosted server is `https://mcp.krcg.org/mcp`. It needs no account and no
sign-in.

## Connect your assistant

### Claude (web and desktop)

1. **Customize > Connectors**, then **Add > Add custom connector**.
2. Name `Primogen`, server URL `https://mcp.krcg.org/mcp`, **Continue**.
3. Authentication: **No sign-in**, then **Add**.

In a chat, the **+** button, **Connectors**, turns it on. The Free plan
allows one custom connector. On Team and Enterprise plans an owner adds it
under **Organization settings > Connectors**.

### Claude Code

```sh
claude mcp add --transport http primogen https://mcp.krcg.org/mcp
```

Add `--scope user` to have it in every project.

### Gemini app

On [gemini.google.com](https://gemini.google.com) (a personal account, 18 or
over, in the US, with Keep Activity on):

1. **Settings > Connected Apps**, then under "Custom apps", **Add a custom
   app**.
2. Server URL `https://mcp.krcg.org/mcp`, **Next**, and follow the steps.

In a chat, type `@` and pick the app so that Gemini uses it.

### Other MCP clients

The server speaks Streamable HTTP at `https://mcp.krcg.org/mcp`, without
authentication.

## Run it locally

With [Go](https://go.dev/dl/) 1.27 or later:

```sh
go install github.com/lionel-panhaleux/primogen@latest
primogen query cards "Freak Drive"
```

`go install` puts `primogen` in `$(go env GOPATH)/bin`, which must be on
your `PATH`. The first run downloads the published data (about a megabyte)
and checks for a newer one at most once a day. `primogen stdio` serves the
tools to a local MCP client, for instance Claude Code:

```sh
claude mcp add primogen -- primogen stdio
```

## Contributing

[`docs/index.md`](docs/index.md) maps what primogen is and what was
decided; [`BOARD.md`](BOARD.md) holds the work.

## License

Licensed under the [MIT License](LICENSE).

Portions of the materials are the copyrights and trademarks of Paradox Interactive AB,
and are used with permission. All rights reserved.
For more information please visit [white-wolf.com](http://www.white-wolf.com).

![Dark Pack](dark-pack.png)
