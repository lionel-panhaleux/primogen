# Roadmap

Rewritten, never appended. The work itself is on [`../BOARD.md`](../BOARD.md).

## Who it is for

VTES players and judges, through whatever LLM assistant they use: an answer
they can trust on a rules question, a deck read the way a tournament player
reads it, a meta picture with numbers. We would know by the owner's grading
of real questions ([product.md](product.md) "Evaluation").

## The phase: live, with users

Announced and served at `mcp.krcg.org`. `vtes_feedback` turns what users
report into the `feedback` issues `/intake` reads; real usage drives what
moves next.

## Ranking

1. A line another line, another repository or a person waits on.
2. The current milestone before the next.
3. What a user would meet, before internal cleanup.

Ties break to the smaller line.

## Milestones

At most three open, each named by what someone can do afterwards. A board
line serving one carries its token; done is deleted. Opening one is the
owner's decision.

- **`[M:rulings]` What the rules directors settled on cards is in the rulings database**: the rulings primogen's calibration, the advanced rules, the `zero-vtes` ledger and the archive audit (Ankha's tenure from 2017–present, plus cards played in the TWDA window since 2021-07 and co-named card pairs) surface are curated by the rulemongers in `vtes-rulings`, and `vtes_cards` and `vtes_rules_case` serve them as official rulings.

## Later, in this order unless the owner reorders

- The legality check on the engine (`vtes_engine_check`), which needs a
  position-check endpoint in zero-vtes.
- Bot-based deck testing, once the engine covers the pool and the bot is
  strong enough.
