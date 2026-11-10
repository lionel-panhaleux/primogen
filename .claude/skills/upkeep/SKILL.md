---
name: upkeep
description: The maintenance loop of primogen, every 20 completed board lines (the counter at the bottom of BOARD.md) or monthly, whichever comes first. Reviews the docs, evicts board lines and milestones, and holds the retrospective that alone changes the harness. Use when the counter reaches 20, when the owner asks for a maintenance pass, or when docs and code feel out of step.
---

# /upkeep

Three passes, in this order.

## 1. Docs review

Read every page of `docs/` against what it describes:

- **Contradictions** between pages.
- **Stale claims**: sample the load-bearing ones against the code; names,
  paths, commands and procedures rot fastest.
- **Domain claims** (`docs/sources.md`) have only their sources to check
  against: does the service still behave as stated? Is an interviewed fact
  old enough to ask again?
- **Orphans**: a page `docs/index.md` does not reach, or one nothing has
  cited or touched in 20 lines, is reviewed for deletion.
- **Size**: a page that became a summary of a summary folds into its parent.

Fix what is unambiguous. What needs a decision goes through `/intake`.

## 2. Eviction

Walk `BOARD.md` oldest first (`git blame` knows each line's age). Every line
gets a verdict, none skipped twice: **do**, **watch** (it waits on a date or
an event: tell the owner, delete the line), **promote** (no done-condition:
its content moves to a docs page), **drop**. Then the milestones of
`docs/roadmap.md`: one nobody served in a long while is a subject or a drop.
The owner decides drops and reordering.

Then ask: are the ranking rules still true? Is the phase still the one we are
in? Is the board getting shorter?

## 3. Retrospective

The only way the harness changes. Read the `Friction:` trailers since the
last pass (`git log --grep '^Friction:'`) and the reviewer's advisory
findings. The same friction seen twice becomes a proposed rule, with its
evidence and one line of reason; a correction the owner had to repeat becomes
a proposed hook. Remove rules that lapsed as readily as you add new ones.
Nothing joins the harness because it seems useful. Proposals go to the owner.

Then reset `<!-- completed-since-upkeep: 0 -->` in `BOARD.md`.
