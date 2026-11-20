# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **The rulings archive holds every V:EKN forum topic Ankha answered** [M:rulings] (work in `newsgroup-archive`): the forum sweep fetches the topics on Ankha's forum profile that the forum's search omits (5,609 of 6,947 posts found by search on 2026-11-11), so `usenet.krcg.org/archive.json.gz` holds the threads no ruling cites yet. Done when `archive.json.gz` includes the missing Rules Questions threads from Ankha's profile and `sources.md` records the closed gap. Docs: sources.md "The archive's limits".
- **The rulings database holds the rulings calibration and the advanced rules already rely on** [M:rulings] (work in `vtes-rulings` / `rulings.krcg.org`, then re-pinned here): the sitting rules directors' rulings that `calibration/corpus/` cites directly from the archive (Foreshadowing Destruction with Conditioning in `vekn-78044`, Trap, Conquer the Beast and Terror Frenzy in `6iew9bqYFVY` and `p-Xf1uFNdPE`, Political Flux in `z_kvG6T-DTQ` and `w6jOOcMIUbM`), the 9 archive posts `knowledge/rules/advanced-rules.md` cites that `references.yaml` lacks, and the `zero-vtes` ledger's card rulings missing from `rulings.yaml` are curated in `vtes-rulings`, `testdata/upstream` is refreshed, and the corpus entries' `settled-by` name their ruling labels. Done when `vtes_cards` returns those rulings from the pinned bundle, no archive post cited by `advanced-rules.md` is absent from `references.yaml`, and `go test ./...` passes. Docs: product.md "Evaluation".
- **Primogen lists the sitting rules directors' posts no ruling cites, for the rulemongers** [M:rulings]: from the bundle and the rulings archive, an audit lists by card (most played in the TWDA window first) and by card pair the sitting rules directors' posts naming them (in their own words or in the question they answer, outside a decklist) that no ruling on those cards cites in `vtes-rulings`, each with its link, date, quoted question and answer, existing rulings on the card, and whether `card-changes.md` records a change to the card after the post's date. Done when the command runs over the bundle and archive, finds the uncited director posts in the pinned test archive, and writes the report the rulemongers curate from. Docs: operations.md "Commands", sources.md "The rulings chain".

<!-- completed-since-upkeep: 24 -->
