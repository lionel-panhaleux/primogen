---
name: calibrate
description: The calibration loop of primogen. Run an owner-graded session (the owner asks, primogen's tools ground the answer, the owner grades, the exchange is recorded) or a regression re-run after a revision. Use for "calibration session", "grade primogen", "run the regression set".
---

# /calibrate

The owner is the only grader. The entry format and regression membership are
in `calibration/README.md`; follow them exactly.

## Graded session

1. Note `git rev-parse --short HEAD`: it goes in every entry's `answered-by`.
2. Take the owner's question; infer the mode (rules, analysis, advice).
3. Answer as a user's assistant would: with primogen's tools only (the MCP
   server, or `primogen query`), no repository file, no other skill, no extra
   hedging because it is a test. The corpus grades what users get.
4. Take the grade and the correction. Write the entry in
   `calibration/corpus/` verbatim before distilling anything, with its
   `settled-by`.
5. Distill: a correction that generalizes becomes proposed wording for the
   guide file it corrects (`knowledge/guide/`: the task's procedure, the
   theory, the modules), never a separate list of lessons; the owner
   approves it before it is written. A correction that contradicts a synced
   file is an upstream fix and a `/sync`. Record in the entry's
   Generalization where it went.
6. Propose regression membership for a substantive correction; the owner
   decides.
7. An answer graded clean that shows the bar may be proposed as a served
   example (principle 8): an excerpt in `knowledge/guide/examples/`, naming
   its corpus entry, the date of the answer and of its data, with its line
   in `docs/provenance.md` and the `vtes_guide` row of `docs/product.md`.
   Its entry joins the regression set. The owner decides.
8. One commit for the session, entries and guide edits together.

## Regression re-run

After a change to the tools, to `knowledge/` or to the server's
instructions.

1. Collect every entry with `regression: yes`.
2. For each, a fresh subagent whose prompt holds only the question and the
   instruction to answer with primogen's tools: never the recorded answer,
   grade or correction. If a served example is drawn from that entry, the
   re-run serves a copy of `knowledge/` without that example.
3. Compare each fresh answer with the graded answer and its correction:
   holds, improved, or drifted (the corrected mistake is back, or a new one).
4. Give the owner the drift report, naming the served examples whose entry
   drifted. Never self-grade; entries, examples and the guide stay untouched
   until the owner rules.

Escalate instead of deciding alone when a lesson would change
`docs/principles.md` or a decision recorded in `docs/`.
