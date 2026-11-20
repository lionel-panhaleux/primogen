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
The announcement waits for M:meta, so it covers the three questions of
[product.md](product.md) (owner, 2026-11-20); `vtes_feedback` turns what
users report into the issues `/intake` reads.

## Ranking

1. A line another line, another repository or a person waits on.
2. The current milestone before the next.
3. What a user would meet, before internal cleanup.

Ties break to the smaller line.

## Milestones

At most three open, each named by what someone can do afterwards. A board
line serving one carries its token; done is deleted. Opening one is the
owner's decision.

- **M:meta**: a player gets the current meta read with numbers. Done when
  an owner-graded session of meta questions passes. `vtes_meta` reads the
  archetype labels, the decks since the labels placed by their nearest
  archetype, and the build proposes new archetypes from the decks still
  unplaced, curated by the owner (owner, 2026-11-20).

## Later, in this order unless the owner reorders

- The announcement (the phase above).
- A pass over the rulings database, with primogen's tools finding what the
  rules directors settled and no ruling holds, the curated rulings kept by
  the rulemongers (owner, 2026-11-12). Until then the tools work around
  what it misses.
- The legality check on the engine (`vtes_engine_check`), which needs a
  position-check endpoint in zero-vtes.
- Members: Archon login, tournaments, full-field meta, rankings, community
  links.
- Bot-based deck testing, once the engine covers the pool and the bot is
  strong enough.
