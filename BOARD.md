# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- `vtes_meta` (M:meta): archetype shares over a window weighted by event
  size, what rises and falls against the previous window, an archetype's
  core cards and representative decks; the decks since the labels placed by
  their nearest archetype at `vtes_deck`'s bar, those near none counted and
  said. Done when: `primogen query vtes_meta` on the current bundle gives
  the latest window's shares with the unplaced share stated. Docs:
  product.md (the `vtes_meta` row, the shipped-tools sentence), AGENTS.md
  "Right now".
- Clustering (M:meta): the window's decks no archetype holds, grouped into
  proposed archetypes with their core cards and decks, each named or
  rejected by the owner in `classification.json`. Done when: the build
  proposes the clusters, the owner has ruled on each, and `vtes_meta` shows
  the unplaced share that remains. Docs: provenance.md
  (`classification.json`), product.md.
- `vtes_feedback`: a user's report on an answer (the question, what was
  wrong, what was expected) filed by a GitHub App as a public issue labelled
  `feedback` on `lionel-panhaleux/primogen`, the user told before filing
  that it is public and the model told to leave out personal details; local
  installs send it through the hosted endpoint; a size cap and a rate
  limit. Done when: a Claude client's call to the hosted server and a call
  from `primogen stdio` each open a labelled issue, and the open `feedback`
  issues reach `/intake` (read by it, or by a feedback-triage skill feeding
  it: chosen at ship with the owner). Docs: product.md (the tool's row,
  what becomes public), operations.md (the App's key in sops, the limits),
  the skill changed or added.

<!-- completed-since-upkeep: 14 -->
