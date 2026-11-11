# Calibration

The owner-graded corpus and its regression set ([product.md](../docs/product.md)
"Evaluation"), and `codex-cases.md`, the codex strategist's graded cases. The
`/calibrate` skill fills the corpus; the lessons it distills go into the
guide, `knowledge/guide/`.

## Corpus entries: `corpus/YYYY-MM-DD-<slug>.md`

One file per graded exchange:

```markdown
---
date: 2026-08-25
mode: rules | analysis | advice
answered-by: primogen <short commit> | vtes-skill <short commit>
model: <only when not the default model of the session>
status: graded | candidate | retired
regression: yes | no
settled-by: <what settles it: rulebook or advanced rules section titles, ruling references, or "owner's grade">
---

## Question

Verbatim, as asked.

## Answer given

Verbatim, as produced.

## Owner grade

The verdict in the owner's words.

## Correction

What the right answer was, and why. Omitted when the grade is clean.

## Generalization

Where the lesson went: the guide file it corrected, or "case-specific:
none".
```

- `answered-by` ties the answer to a revision; drift is meaningless without
  it. The entries answered by `vtes-skill` are the baseline primogen must
  beat ([roadmap.md](../docs/roadmap.md) "M:rules").
- `status: graded` needs the owner's verdict. Feedback from anyone else
  enters as `status: candidate`, attributed, and drives nothing until the
  owner grades it.
- `regression: yes` puts the question in the regression set. Proposed for
  any substantive correction; the owner confirms.
- `settled-by` names what settles the question as graded: section titles
  rather than numbers, which move, ruling references, card names for card
  text. `owner's grade` covers only what judgement settles; a fact a
  correction rests on names its source beside it. When a `/sync` or an
  evidence test shows one of them changed, the entry goes back to the owner,
  who re-grades it or retires it. Retiring sets `status: retired` and
  `regression: no`: the entry stays as a record. A served example drawn from
  an entry the owner re-grades or retires is removed.
- A rules entry's `settled-by` is what its evidence test reads: items
  separated by `; `, each `rulebook, <title>`, `advanced rules, <title>`
  (without its number), `tournament rules, <title>`, `judges' guide,
  <title>`, a ruling reference (`ANK 20211003`), `card text of <name>, <name>`
  (the cards the test passes), or `owner's grade`; a closing parenthesis
  explains an item and is not read.
- No entry is served or shown to an answering model with its answer: only
  its question. The one exception is an excerpt the owner promotes to a
  served example (principle 8, `/calibrate` "Graded session" step 7).
- An entry is a record: once written, only its front matter changes. The
  `vtes-skill` entries name that repository's files (`LESSONS.md`,
  `references/`); their lessons are now in `knowledge/guide/`.
