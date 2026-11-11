# Calibration

The owner-graded corpus and its regression set ([product.md](../docs/product.md)
"Evaluation"). The `/calibrate` skill fills it; the lessons it distills go
into the guide, `knowledge/guide/`.

## Corpus entries: `corpus/YYYY-MM-DD-<slug>.md`

One file per graded exchange:

```markdown
---
date: 2026-08-25
mode: rules | analysis | advice
answered-by: primogen <short commit> | vtes-skill <short commit>
model: <only when not the default model of the session>
status: graded | candidate
regression: yes | no
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
- An entry is a record: once written, only its front matter changes. The
  `vtes-skill` entries name that repository's files (`LESSONS.md`,
  `references/`); their lessons are now in `knowledge/guide/`.
