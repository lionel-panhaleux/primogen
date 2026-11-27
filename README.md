# Primogen

A [VTES](https://www.vekn.net) (Vampire: The Eternal Struggle) councilor for
any LLM assistant: rules answers that quote what settles them, deck analysis
and advice, a meta picture with numbers.

It is an MCP server with no LLM on the server side: it returns verbatim card
texts, rulings and their sources, and your own assistant writes the answer.

The hosted server is `https://mcp.krcg.org/mcp`. Public tools need no account
and no sign-in; signing in with [Archon](https://archon.vekn.net) unlocks
members-only tournament standings, rankings, player records, published
Archon decklists, and full-field meta.

## Tools

- `vtes_cards`: a card's official text, history and rulings.
- `vtes_rules_case`: the evidence for a rules question in one call: the cards, their rulings with the original posts, the rules directors' posts naming the cards together where no ruling does, the governing rules sections, what no source covers.
- `vtes_rules_search`, `vtes_rules_read`: the rulebook, the advanced rules, the tournament rules and the judges' guide, searched by keyword or read by section.
- `vtes_rulings_archive`: full-text search of the rulings archive, each post's author marked when they were the sitting rules director.
- `vtes_deck`: a decklist (as text, a TWDA id, a link from VDB, VTES Decks or Amaranth, or, with Archon login, a published Archon deck uid or tournament) validated and read: the crypt, which crypt cards have what each card requires and at which level, the library by type with every card's text and by module with each module's density, copies per hand and chance of an opening hand holding one, the odds that the cards a plan needs meet in a hand, its answers to the field's combat defenses, the costs, its budgets, blood and pool sources, concentration and staples beside its reference decks, its track record among winning decks, the nearest archetypes and tournament decks; or a deck idea's cards, a count optional, its sizes not checked.
- `vtes_card_stats`: a card among the tournament-winning decks, overall or within an archetype: how many play it and with how many copies, what is played with it, and its competitors for the slot with their texts.
- `vtes_meta`: the archetypes' shares of the tournament-winning decks of the latest window, weighted by event size, the share in no archetype stated, what rises and falls and the cards behind it, and, with Archon login, full-field play share against win share and conversion over the events that published decklists beyond the winner; an archetype's share over time, core cards and decks.
- `vtes_events`: community links (without login); with Archon login, tournaments, standings, published decklists, rating rankings, a member's ratings, known wins, tournament history and decks (by VEKN ID), and leagues.
- `vtes_guide`: task procedures, VTES strategy theory, the module vocabulary and combat defenses, game-term translations (labeled unchecked, consistency reference), and owner-graded worked examples.
- `vtes_feedback`: a report on a wrong answer (the question, what was wrong, what was expected), filed with the user's agreement as a public issue on this repository, labelled `feedback`.

## Sources

- [KRCG](https://static.krcg.org) for cards data and the Tournament Winning Deck Archive (TWDA)
- [VDB](https://vdb.im), [VTES Decks](https://vtesdecks.com) and [Amaranth](https://amaranth.vtes.co.nz) for the decks shared from them
- The [rulings database](https://rulings.krcg.org)
- The [usenet archive](https://usenet.krcg.org)
- The [rulebook](https://archon.vekn.net/help/rules), the [tournament rules](https://archon.vekn.net/help/tournament-rules) and the [judges' guide](https://archon.vekn.net/help/judges-guide) of the V:EKN, from [Archon](https://archon.vekn.net)
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
