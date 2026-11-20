# Provenance of the knowledge files

Where each file in `knowledge/` comes from and how it changes.
[principles.md](principles.md) "Knowledge files" says who may edit what.

## Synced

Copied byte for byte from their upstream, never edited here. `/sync`
compares them with upstream and records the commit below.

| File | Upstream | Path | Copied at |
|---|---|---|---|
| `rules/advanced-rules.md` | `vtes-biased/vtes-advanced-rules` | `docs/advanced-rules.md` | `f8aa01d`, 2026-11-18 |
| `card-changes.md` | the owner's gist `a4a9cad2e492af7c45c50a1cea7e6cf6` | `VTES History of changes.md` | `f1be31f`, 2026-11-11 |

## Authored here

- `guide/`: what the model reads to do a task. `theory.md` and the
  procedures (`rules-answer.md`, `deck-analysis.md`, `deck-advice.md`,
  `card-comparison.md`, `meta-reading.md`) come from the `vtes` skill, the
  codex-of-the-damned strategist and their owner-graded lessons; the
  strategist's graded cases stay in `calibration/codex-cases.md`, unserved.
  `rules-answer.md` closes every `vtes_rules_case` pack as its answer
  contract, `deck-analysis.md` and `deck-advice.md` every `vtes_deck`
  dossier, `deck-advice.md` every `vtes_card_stats` result; `theory.md`, `modules.md`, `card-comparison.md` and
  `meta-reading.md` are served by no tool yet.
  `examples/` holds the served examples (principles.md, principle 8): each
  an excerpt of one of primogen's own answers the owner graded clean, named
  after its task, naming its corpus entry and the dates of the answer and
  of its data, its entry in the regression set. `card-comparison-trap-conquer-the-beast.md`,
  from `calibration/corpus/2026-11-20-trap-or-conquer-the-beast.md`, promoted
  on 2026-11-20. Served by no tool until `vtes_guide` ships.
  `modules.md` is the module catalog the owner canonized on 2026-08-25,
  checked against the card texts on 2026-11-19, but for its "Combat
  defenses", written on 2026-11-19 after codex-of-the-damned's
  combat-modules article, each card read from its text, not canonized by
  the owner. `module-tags.json` gives each of its modules, by its name and
  section there, its cards: the core its entry names first, the support it
  adds (refills, enablers, fetch targets), each checked against its text in
  the card data, and gives the reason a module has no card of its own; and
  each combat defense its cards and its answers, grouped by how they answer
  it, each with the level it needs; and its `economy`, the blood and pool
  sources by kind (pool, blood, steal, blood to pool, pool to blood, each
  defined there, the served wording), each
  card read from its text, the cards a text finds gaining pool or blood
  that are no source excluded with the reason, the doubtful ones settled
  by the owner on 2026-11-20. `vtes_deck` reads a library by module
  with it, serving the modules' names, not the catalog's prose; its
  combat posture, serving the defenses' names and their answers' `how`;
  and its sources by kind beside its reference.
- `classification.json`: the owner's archetype labels for the
  tournament-winning decks since 2021-07 (the 2026-07 pass), forked from
  codex-of-the-damned @ `18be59e`. `vtes_deck` places a deck among its
  archetypes ([product.md](product.md)); the build names the decks since
  that the labels lack and no archetype matches. A variant is of an archetype
  that is itself no variant; a file breaking that fails the build.
- `game-terms.json`: the game terms in English (the 2024 rulebook's, plus
  the two 2025 terms and Seraph, a trait printed on Black Hand vampires)
  and their translations in six languages, from
  `vtes-skill` @ `65b5c3e`. The translations are unchecked, the exception
  [principles.md](principles.md) "Knowledge files" records; the file says so
  itself.
- `llms.txt`: the catalog of the KRCG stack the hosted server serves at
  its root ([product.md](product.md) "Connecting"), each line checked
  against its service's own specification when written.
- `shorthand.json`: players' shorthand for words of the game that the game
  terms do not list ("NRA", which the advanced rules print in §3.5's title;
  "hunter"), which `vtes_rules_case` takes for no card's name. Kept apart
  from the game terms, it grows when a question finds a card by such a word.

## Left behind

What `vtes-skill` @ `65b5c3e` carried and primogen does not:

- The rulings database: the data bundle reads `vtes-rulings` itself
  ([sources.md](sources.md)). The rulings canon, the advanced rules
  maintainers' cache of rulings' original wording: the archive serves the
  originals.
- Condensed copies of the tournament rules, the judges' guides (current and
  2004), the code of ethics and the rulebook: paraphrases, where primogen
  serves verbatim text. Primogen reads the rulebook, the tournament rules
  and the judges' guide from `archon-vibe` at each refresh
  ([sources.md](sources.md)); the code of ethics and the 2004 judges' guide
  are not read.
- Its notes on gaps in the official documents: unverified.
- The meta tables by year: computed facts, which the meta tools compute
  from the archive (principle 7).
- The community tier (blog index, KGS Cards posts): primogen serves no
  external blogs ([product.md](product.md)).
- The usenet mailboxes, kept as the `usenet-archives` release of
  `vtes-skill`: primogen reads the rulings archive at `usenet.krcg.org`.
- Its Python scripts and the skill's router: primogen's tools replace them.
