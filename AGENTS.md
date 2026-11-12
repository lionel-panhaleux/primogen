# AGENTS.md

Primogen is a VTES (Vampire: The Eternal Struggle) councilor any LLM can use:
rules answers, deck analysis and advice, meta reading. It is an MCP server
with no LLM on the server side, hosted at `mcp.krcg.org` and runnable locally
from the same Go binary. `CLAUDE.md` is a symlink to this file.

## Three lifespans

Every artifact has exactly one. What fits none is noise: delete it.

- **Code** (and `knowledge/`, which is served to users): permanent, the source
  of truth for *how* and for what the councilor says.
- **[`docs/`](docs/index.md)**: standing, the source of truth for *what is*
  and *what was decided*. Kept true by editing, never by appending.
- **Task context** (plans, findings, reasoning): dies with its task. A line
  in progress may keep its plan in `board/<slug>.md`, deleted with the line.

No other tracker: no plan document, no TODO file, no TODO comment. Discovered
work goes through `/intake` or is done now. No memory outside the repository
(`autoMemoryEnabled` is off): a fact worth keeping goes in `docs/`.

## Start here

1. [`docs/index.md`](docs/index.md): the map. Read the pages your task
   touches before the code. [`docs/principles.md`](docs/principles.md) is the
   standard intake and review check against.
2. [`BOARD.md`](BOARD.md): what must change, in priority order. Continue a
   line before opening one. [`docs/roadmap.md`](docs/roadmap.md) holds the
   phase, the ranking rules and the open milestones.

Context lives in the docs, asks live on the board.

## The loops

- **`/intake`**: nothing reaches the board unchallenged. Intake proposes, the
  owner accepts.
- **`/ship`**: one line in, one change out: code or knowledge, the docs pages
  named at intake, the board line deleted; then the `reviewer` agent, fresh
  context, before the commit is final.
- **`/sync`**: the synced knowledge files checked against their upstreams
  (`vtes-advanced-rules`, the owner's gist of card changes), every diff read
  before it lands.
- **`/calibrate`**: an owner-graded session recorded in `calibration/`, or a
  regression re-run after a revision.
- **`/upkeep`**: every 20 completed lines or monthly: docs review, board
  eviction, the retrospective that alone changes this harness.

Work in another repository (`vtes-advanced-rules`, `newsgroup-archive`,
`rulings-website`, `zero-vtes`) follows that repository's own harness. The
board here holds only the line that says what must exist upstream.

## Models

Review and decisions on the strongest model (Fable 5.2: the `reviewer` agent
is pinned to it). Implementation on Opus 5.5. A mechanical-edit agent on the
cheapest model only once it is earned. The owner changes this, not a session.

## Talking to the owner

Decisions, not narration. Interrupt only for a choice of principle, an
irreversible or outward-facing action, a scope or phase change, or a review
deadlocked after two rounds. A request for a decision stands alone: numbered
questions, the recommendation first, every excerpt needed to answer, and it
ends the message.

## Commits

Trunk-based: straight to `main`, no feature branch, no worktree. Describe the
change itself. A commit that met friction (a correction by the owner, a
second review round, a surprise, a failing check) says so in a `Friction:`
trailer; upkeep reads them.

## Right now

`vtes_cards` and the four rules tools (`vtes_rules_case`,
`vtes_rules_search`, `vtes_rules_read`, `vtes_rulings_archive`) hosted at
`https://mcp.krcg.org/mcp` and refreshing themselves (not announced yet), and
from a local install (`primogen stdio`) downloading the published bundle. The
first owner-graded rules session on them is in `calibration/corpus/`
(2026-11-12). From a checkout: `primogen build`, then
`primogen query` or `primogen stdio` with `-bundle data/bundle.json.gz`.
The board holds the next lines.
