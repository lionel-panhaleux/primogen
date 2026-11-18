# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **The dossier reads the deck by module** (M:deck): a card-to-module tag file in `knowledge/`, derived from `modules.md`, every claim checked against the card texts. `vtes_deck` groups the library by module, with each module's density and expected copies per hand, and names the cards that belong to no module. Done when every card in the core of each `modules.md` module is tagged, a reference deck's modules come back as the catalog describes them, and the status page lists TWDA cards with no tag. Docs: provenance.md, product.md, operations.md.

- **Advice surveys the competitors for a slot** (M:deck): `vtes_card_stats` ships with play rates, co-occurrence, and the cards tournament decks play in a card's place, overall or within an archetype (moved into M:deck, owner, 2026-11-15). Done when, for the Toreador vote entry in the regression set (`calibration/corpus/2026-08-26-toreador-vote-push.md`), Awe's competitors include Scalpel Tongue and Closed Session, the alternatives the owner's grade named; or else the line comes back to the owner with why the data doesn't show them. Docs: product.md, roadmap.md.

- **An owner-graded deck session** (M:deck): a `/calibrate` session of deck analyses and advice on primogen's tools, including the two deck entries in the regression set (the Toreador vote deck, the Sword of Nuln deck). Done when the entries are recorded and the owner's grade passes them or names the lines that follow. Docs: roadmap.md (M:deck deleted when it passes).

<!-- completed-since-upkeep: 0 -->
