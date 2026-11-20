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
The tools cover the three questions of [product.md](product.md), each
through an owner-graded session; the announcement is the owner's to make.
`vtes_feedback` turns what users report into the issues `/intake` reads.

## Ranking

1. A line another line, another repository or a person waits on.
2. The current milestone before the next.
3. What a user would meet, before internal cleanup.

Ties break to the smaller line.

## Milestones

At most three open, each named by what someone can do afterwards. A board
line serving one carries its token; done is deleted. Opening one is the
owner's decision.

- **M:members**: a logged-in VEKN member gets Archon tournaments, standings,
  rankings, member records (by VEKN ID), published decklists beyond the TWDA,
  and full-field play share against win share; an unlogged user keeps every
  public tool and community links (owner, 2026-11-20). Done when an
  owner-graded session of member questions passes.

## Later, in this order unless the owner reorders

- The announcement (the phase above).
- A pass over the rulings database, with primogen's tools finding what the
  rules directors settled and no ruling holds, the curated rulings kept by
  the rulemongers (owner, 2026-11-12). Until then the tools work around
  what it misses.
- The legality check on the engine (`vtes_engine_check`), which needs a
  position-check endpoint in zero-vtes.
- Bot-based deck testing, once the engine covers the pool and the bot is
  strong enough.
