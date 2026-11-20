# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- `vtes_deck` sets a draft, and a list near no archetype, beside its
  nearest archetype's staples: the library cards in three quarters of its
  decks that the list lacks, defenses first, the archetype named with its
  distance. Done when: `primogen query -draft deck` on the 2026-11-20
  werewolf draft lists Bait and Switch among its nearest archetype's missing
  staples, and a fresh regression answer to that entry offers it. Docs:
  product.md (the `vtes_deck` row, "Grading a deck").

- The deploy runs from CI: `mise run release` on a GitHub Actions
  workflow, the server-setup deploy key decrypting `secrets.sops.yaml`
  (added to `.sops.yaml`, `sops updatekeys`) and reaching frankfurt. Done
  when: a release tagged from a workflow run deploys and `mcp.krcg.org`
  reports its tag. Docs: operations.md (Secrets, Releases and deploys).

<!-- completed-since-upkeep: 17 -->
