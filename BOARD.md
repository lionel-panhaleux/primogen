# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **A deck link from any major deckbuilder** (M:deck): `vtes_deck` reads vtesdecks.com and Amaranth links as it reads VDB's; a link it cannot read is refused with the list-as-text fallback. Done when `primogen query deck` on a vtesdecks link and an Amaranth link both return the dossier. Docs: product.md, sources.md.
- **A deck idea gets the advice procedure** (M:deck): `vtes_deck` accepts a draft (a few cards, no counts or sizes enforced) and returns its cards, modules, nearest archetypes and the advice procedure; the server instructions send a deck idea there. Done when `primogen query deck` on the Sword of Nuln idea's cards returns them, and a re-run of that regression entry answers with the procedure in hand. Docs: product.md.
- **A deck's combat posture against the field** (M:deck): `vtes_deck` reports which prominent combat defenses (S:CE, grapple, dodge, prevention…) the deck answers and which it lacks, each with its share of TWDA winners in the window, from a combat section of `modules.md` inspired by codex-of-the-damned's combat-modules article. Done when the Sword of Nuln draft and the Banu list both show their S:CE answers or the lack of one, with the shares. Docs: product.md, modules.md.
- **Grade a deck's economy, efficiency and synergy** (M:deck): research by a Fable agent into measures the server can compute and a player can trust: pool and blood gained against spent (bloat per turn and the net difference), action density against the turns a deck gets, the odds that a module or a kill lands from a hand (copies, draws, the cards it needs together), how well the modules feed each other, and the level of an idea against its archetype's winners. Each measure is checked against the graded deck entries (the Sword of Nuln idea, the Banu list). Done when docs record the measures chosen, each with its computation and the graded entries it separates, approved by the owner, and the lines implementing them have passed intake. Docs: product.md (the `vtes_deck` row), theory.md if a measure is theory.

<!-- completed-since-upkeep: 5 -->
