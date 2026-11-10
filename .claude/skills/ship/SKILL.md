---
name: ship
description: The execution loop of primogen. Take a line of BOARD.md, do the work, and land it as one change (code or knowledge, the docs pages named on the line, the board line deleted), reviewed by the fresh-context reviewer agent before the commit. Use to start or continue any board line, and as the default way a substantive change lands here.
---

# /ship

Arguments name the line; otherwise take the top line of `BOARD.md`.

## 1. Claim

Read `BOARD.md`, `docs/index.md` and the pages the line names. Check
`git status` for another session's work in progress: a change you did not
make is someone's work, never debris. Say which line you take. If the line
has a `board/<slug>.md`, read it and keep working notes there.

## 2. Do the work

Hold to `docs/principles.md`. Before using a library, a protocol or a tool,
read its current documentation (`docs/operations.md` "Looking things up").
Before writing any card or rules claim into `knowledge/`, a test or a tool's
text, check it against the sources; never from memory.

Scope grows in place: what exploration turns up and shares the change is done
in it. Only a genuinely separate topic goes back through `/intake`. Work in
another repository follows that repository's harness.

## 3. Land one change

1. **Code or knowledge**, with the checks of `docs/operations.md` green.
2. **Docs**: every page named on the line updated, or its absence justified
   in a sentence. A pivot is an edit: the docs never keep a deprecated fact.
3. **Board**: the line deleted, its `board/<slug>.md` too.

Before review, delete every comment that repeats the docs, another comment or
the code.

## 4. Review

Start the `reviewer` agent with the diff (or the commit range), the board
line's text and nothing else: it has not seen this conversation and reads the
docs and the code itself. A docs-only change is not reviewed; a change to
`knowledge/` is.

Blocking and advisory findings are both fixed in the same change. A finding
you think wrong is argued in the next round; after two rounds of
disagreement, stop and give the owner the disagreement at its turning point.
"Looks good" is a normal verdict.

## 5. Commit and report

Commit to `main`, describing the change itself. Add a `Friction:` trailer if
there was any: an owner correction, a second review round, a surprise, a
failing check. Bump `<!-- completed-since-upkeep: N -->` in `BOARD.md`; at
20, or a month after the last pass, run `/upkeep`. Tell the owner the line,
what landed, the pages that moved and the commit.
