# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **`vtes_events` and Archon decklists** [M:members]: for a logged-in VEKN
  member, `vtes_events` returns tournaments (by code, uid, country, date
  window, rank or format) with their standings and published decklists,
  rankings and ratings by category and country, a member's ratings, known
  wins, tournament history and published decks (by VEKN ID), and leagues;
  community links from `/v1/export` are served to everyone; `vtes_deck`
  accepts an Archon deck uid or tournament deck for a logged-in member;
  without login (or from a local install) member queries state that Archon
  login at `mcp.krcg.org/mcp` is required. Done when: tests on a pinned
  Archon export cover tournament lookup and filter, rankings, a member's
  wins and decks, community links with and without login, an Archon deck in
  `vtes_deck`, and the unauthenticated refusal. Docs: product.md ("Tools",
  "Connecting"), operations.md ("Commands"), AGENTS.md "Right now",
  README.md.
- **Full-field meta in `vtes_meta`** [M:members]: for a logged-in VEKN
  member, `vtes_meta` places the window's published Archon decks by
  similarity alongside the winners and reports play share against win share
  (and conversion) over the events that published decklists beyond the
  winner. Done when: `vtes_meta` with a member token reports play share
  beside win share on the pinned Archon export and without a token continues
  to serve the TWDA winners' shares while noting that full-field shares need
  Archon login. Docs: product.md (`vtes_meta` row), AGENTS.md "Right now".

<!-- completed-since-upkeep: 21 -->
