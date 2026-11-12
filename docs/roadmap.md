# Roadmap

Rewritten, never appended. The work itself is on [`../BOARD.md`](../BOARD.md).

## Who it is for

VTES players and judges, through whatever LLM assistant they use: an answer
they can trust on a rules question, a deck read the way a tournament player
reads it, a meta picture with numbers. We would know by the owner's grading
of real questions ([product.md](product.md) "Evaluation").

## The phase: building, before any user

Served at `mcp.krcg.org`, not announced. Ends when the owner announces it
to the community; real users change the ranking and may change the harness.

## Ranking

1. A line another line, another repository or a person waits on.
2. The current milestone before the next.
3. What a user would meet, before internal cleanup.

Ties break to the smaller line.

## Milestones

At most three open, each named by what someone can do afterwards. A board
line serving one carries its token; done is deleted. Opening one is the
owner's decision.

- **M:cards**: anyone can add `mcp.krcg.org` to their assistant and get exact
  card text and rulings. Done when a card and its rulings, each with its
  link, come back in at least Claude, ChatGPT and one local client, from the
  hosted server and from a local install, and the README says how.
- **M:rules**: a player gets a rules answer that quotes what settles it. Done
  when the evidence tests pass on the seed cases and an owner-graded rules
  session beats the `vtes` skill on the regression set.
- **M:deck**: a player gets their deck explained and improved. Done when an
  owner-graded session of deck analyses and advice passes.

## Later, in this order unless the owner reorders

- A pass over the rulings database, with primogen's tools finding what the
  rules directors settled and no ruling holds, the curated rulings kept by
  the rulemongers (owner, 2026-11-12). Until then the tools work around
  what it misses.
- The legality check on the engine (`vtes_engine_check`), which needs a
  position-check endpoint in zero-vtes.
- The meta: the Go port of the TWDA analytics, `vtes_meta`,
  `vtes_card_stats`; codex-of-the-damned reads primogen's output.
- Members: Archon login, tournaments, full-field meta, rankings, community
  links.
- `llms.txt`, the front door of the KRCG stack.
- Bot-based deck testing, once the engine covers the pool and the bot is
  strong enough.
