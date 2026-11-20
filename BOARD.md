# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **Full-field meta in `vtes_meta`** [M:members]: for a logged-in VEKN
  member, `vtes_meta` places the window's published Archon decks by
  similarity alongside the winners and reports play share against win share
  (and conversion) over the events that published decklists beyond the
  winner. Done when: `vtes_meta` with a member token reports play share
  beside win share on the pinned Archon export and without a token continues
  to serve the TWDA winners' shares while noting that full-field shares need
  Archon login. Docs: product.md (`vtes_meta` row), AGENTS.md "Right now".

<!-- completed-since-upkeep: 22 -->
