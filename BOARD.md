# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **The rulings database holds the rulings calibration and the advanced rules already rely on** [M:rulings] (work in `vtes-rulings` / `rulings.krcg.org`, then re-pinned here): the sitting rules directors' rulings that `calibration/corpus/` cites directly from the archive (Foreshadowing Destruction with Conditioning in `vekn-78044`, Trap, Conquer the Beast and Terror Frenzy in `6iew9bqYFVY` and `p-Xf1uFNdPE`, Political Flux in `z_kvG6T-DTQ` and `w6jOOcMIUbM`), the 9 archive posts `knowledge/rules/advanced-rules.md` cites that `references.yaml` lacks, and the `zero-vtes` ledger's card rulings missing from `rulings.yaml` are curated in `vtes-rulings`, `testdata/upstream` is refreshed, and the corpus entries' `settled-by` name their ruling labels. Done when `vtes_cards` returns those rulings from the pinned bundle, no archive post cited by `advanced-rules.md` is absent from `references.yaml`, and `go test ./...` passes. Docs: product.md "Evaluation".

<!-- completed-since-upkeep: 1 -->
