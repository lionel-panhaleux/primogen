# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

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
- `vtes_deck` sets a draft, and a list near no archetype, beside its
  nearest archetype's staples: the library cards in three quarters of its
  decks that the list lacks, defenses first, the archetype named with its
  distance. Done when: `primogen query -draft deck` on the 2026-11-20
  werewolf draft lists Bait and Switch among its nearest archetype's missing
  staples, and a fresh regression answer to that entry offers it. Docs:
  product.md (the `vtes_deck` row, "Grading a deck").

<!-- completed-since-upkeep: 16 -->
