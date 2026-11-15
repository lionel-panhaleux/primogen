# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **`vtes_deck` returns a decklist's computed facts** (M:deck): the bundle reads the TWDA (`static.krcg.org/data/v5/twda.json`). `vtes_deck` takes a decklist as text, a TWDA id or a VDB link, resolves each card name the same way `vtes_cards` does, and validates the list (crypt and library sizes, crypt groups, banned cards, legality at today's date). It returns the crypt table (capacity, inferior and superior disciplines, titles, group), which crypt vampires can play each library card and at what level, the library by type, the costs, and the procedures of `deck-analysis.md` and `deck-advice.md` as the answer contract. Done when a TWDA deck by id, the same list as text, and a VDB link each return the dossier from `primogen query` and from the hosted server, an unknown card or an illegal list gets an error that says what to fix, and every check of operations.md passes. Docs: product.md, sources.md, operations.md, README, AGENTS.md "Right now".

- **The dossier names the deck's relatives** (M:deck): codex's `classify.py` (`.claude/skills/twda/scripts/`), ported to Go, runs over `knowledge/classification.json`. `vtes_deck` returns the nearest archetypes, the nearest tournament decks with their authors' comments, and the cards that differ from those decks; a deck near no archetype says so. Done when the decks of `classification.json` classify back to their own group (leaving each deck out of its group in turn, at a rate the test states), a list from outside the TWDA gets its relatives, and the status page lists recent TWDA decks that match no archetype. Docs: product.md, provenance.md, operations.md, roadmap.md ("Later": the meta keeps the clustering and the shares; the card stats move to M:deck).

- **The dossier reads the deck by module** (M:deck): a card-to-module tag file in `knowledge/`, derived from `modules.md`, every claim checked against the card texts. `vtes_deck` groups the library by module, with each module's density and expected copies per hand, and names the cards that belong to no module. Done when every card in the core of each `modules.md` module is tagged, a reference deck's modules come back as the catalog describes them, and the status page lists TWDA cards with no tag. Docs: provenance.md, product.md, operations.md.

- **Advice surveys the competitors for a slot** (M:deck): `vtes_card_stats` ships with play rates, co-occurrence, and the cards tournament decks play in a card's place, overall or within an archetype (moved into M:deck, owner, 2026-11-15). Done when, for the Toreador vote entry in the regression set (`calibration/corpus/2026-08-26-toreador-vote-push.md`), Awe's competitors include Scalpel Tongue and Closed Session, the alternatives the owner's grade named; or else the line comes back to the owner with why the data doesn't show them. Docs: product.md, roadmap.md.

- **An owner-graded deck session** (M:deck): a `/calibrate` session of deck analyses and advice on primogen's tools, including the two deck entries in the regression set (the Toreador vote deck, the Sword of Nuln deck). Done when the entries are recorded and the owner's grade passes them or names the lines that follow. Docs: roadmap.md (M:deck deleted when it passes).

<!-- completed-since-upkeep: 18 -->
