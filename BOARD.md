# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **The hosted server, live** [M:cards]: `primogen serve` and its deploy are built (operations.md); what remains is the owner's DNS record `mcp.krcg.org` → frankfurt, the first `mise run deploy`, the `primogen` row in `server-setup`'s OPERATIONS.md "Who deploys what", and the check. Done when an MCP client connected to `https://mcp.krcg.org/mcp` gets a card, a local install downloads the bundle, and the data date advances after an upstream change with nobody touching the server. Docs: operations.md ("not deployed yet" goes), AGENTS.md "Right now".
- **What each client does with the server** [M:cards]: Claude, ChatGPT, Gemini and one local client, hands-on: whether the server's instructions reach the model, how prompts appear, what a tool result looks like to the user. Done when product.md "Connecting" carries a table of the findings and the server's instructions and results are adjusted to the weakest client. Docs: product.md.
- **A README a newcomer can follow** [M:cards]: what primogen is, the hosted URL, a configuration snippet per checked client, the local install. Done when each checked client connects by following it. Docs: none (the README is the page).
- **The history of changes links no announcement at `groups.google.com`** (work in `newsgroup-archive`, then the owner's gist): the card names and the other links are done; what remains is the 1998-08-01 floor rules, Google thread `E7fFdAhjpdg`, a single post by the DCI's Robert Goudie that the archive's rule (a rules director took part) left out. The archive takes the thread, as it took the topics a ruling cites; then the gist's DCI 19980801 link moves to its `usenet.krcg.org` copy, and `/sync`. Done when no served announcement link points at `groups.google.com`. Docs: sources.md (the gist row's exception).
- **The archive's sweep does not miss or stop silently** [M:rules] (work in `newsgroup-archive`): the `--refresh` run scheduled weekly over recent topics; the `--expect` check fed every forum topic the rulings database cites; a monthly dated heartbeat commit so the schedule is not disabled. Done when the three run in its workflow. Docs: sources.md "The archive's limits".
- **A rules answer's evidence in one call** [M:rules]: the section index over rulebook, advanced rules, tournament rules and judges' guide; `vtes_rules_case`, `vtes_rules_search`, `vtes_rules_read`, `vtes_rulings_archive`; the evidence tests seeded from the calibration corpus (each rules entry's `settled-by` traced to today's sources), the codex cases, the `rules-qa` items as questions and the engine's rulings ledger. Done when the evidence tests pass and the owner has graded a rules session. Docs: product.md "Tools" and "Evaluation".

<!-- completed-since-upkeep: 5 -->
