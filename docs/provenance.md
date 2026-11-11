# Provenance of the knowledge files

Where each file in `knowledge/` comes from and how it changes.
[principles.md](principles.md) "Knowledge files" says who may edit what.

## Synced

Copied byte for byte from their upstream, never edited here. `/sync`
compares them with upstream and records the commit below.

| File | Upstream | Path | Copied at |
|---|---|---|---|
| `rules/rulebook.md` | `vtes-biased/vtes-advanced-rules` | `rulebook2024/content.md` (submodule) | `rulebook2024` @ `0644d3d`, 2026-11-11 |
| `rules/advanced-rules.md` | `vtes-biased/vtes-advanced-rules` | `docs/advanced-rules.md` | `5c844a5`, 2026-11-11 |
| `card-changes.md` | the owner's gist `a4a9cad2e492af7c45c50a1cea7e6cf6` | `VTES History of changes.md` | `b8beddd`, 2026-11-11 |

## Authored here

- `guide/`: what the model reads to do a task. `theory.md` and the
  procedures (`rules-answer.md`, `deck-analysis.md`, `deck-advice.md`,
  `card-comparison.md`, `meta-reading.md`) come from the `vtes` skill, the
  codex-of-the-damned strategist and their owner-graded lessons; the
  strategist's graded cases stay in `calibration/codex-cases.md`, unserved.
  `modules.md` is the module catalog the owner canonized on 2026-08-25.
- `classification.json`: the owner's archetype labels for the
  tournament-winning decks since 2021-07 (the 2026-07 pass), forked from
  codex-of-the-damned @ `18be59e`. The input of the archetype tools.
- `game-terms.json`: the game terms in English (the 2024 rulebook's, plus
  the two 2025 terms) and their translations in six languages, from
  `vtes-skill` @ `65b5c3e`. The translations are unchecked, the exception
  [principles.md](principles.md) "Knowledge files" records; the file says so
  itself.

## Left behind

What `vtes-skill` @ `65b5c3e` carried and primogen does not:

- The rulings database: the data bundle reads `vtes-rulings` itself
  ([sources.md](sources.md)). The rulings canon, the advanced rules
  maintainers' cache of rulings' original wording: the archive serves the
  originals.
- Condensed copies of the tournament rules, the judges' guides (current and
  2004), the code of ethics and the rulebook: paraphrases, where primogen
  serves verbatim text. The official documents are on `vekn.net`
  ([sources.md](sources.md)).
- Its notes on gaps in the official documents: unverified.
- The meta tables by year: computed facts, which the meta tools compute
  from the archive (principle 7).
- The community tier (blog index, KGS Cards posts): primogen serves no
  external blogs ([product.md](product.md)).
- The usenet mailboxes, kept as the `usenet-archives` release of
  `vtes-skill`: primogen reads the rulings archive at `usenet.krcg.org`.
- Its Python scripts and the skill's router: primogen's tools replace them.
