# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **Primogen reads the krcg 5 card data** (M:cards): the bundle reads `static.krcg.org/data/vtes.json` in the krcg 5 model (`krcg/models.py`) in place of `data/v4`, a snapshot frozen on 2026-06-28 (krcg-static `62cab1c`) that lacks the 63 cards since (Backroom Deal, Aaron Fletcher, Silent Investor); the pinned `testdata/upstream/vtes.json` moves with it. Done when those cards come back from `vtes_cards` with their text and rulings, every check of operations.md passes, and the data line no longer dates the frozen file by a `Last-Modified` its deploys move. Docs: sources.md (the first row), operations.md, provenance.md if it names the file.

<!-- completed-since-upkeep: 17 -->
