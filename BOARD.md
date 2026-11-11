# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **The hosted server, refreshing itself** [M:cards]: Streamable HTTP, the refresh scheduler and the bundle swap of operations.md "Data refresh", the published bundle, the deploy on `server-setup`. Done when an MCP client connected to `https://mcp.krcg.org/mcp` gets a card, a local install downloads the bundle, and the data date advances after an upstream change with nobody touching the server. Docs: operations.md.
- **What each client does with the server** [M:cards]: Claude, ChatGPT, Gemini and one local client, hands-on: whether the server's instructions reach the model, how prompts appear, what a tool result looks like to the user. Done when product.md "Connecting" carries a table of the findings and the server's instructions and results are adjusted to the weakest client. Docs: product.md.
- **A README a newcomer can follow** [M:cards]: what primogen is, the hosted URL, a configuration snippet per checked client, the local install. Done when each checked client connects by following it. Docs: none (the README is the page).
- **The history of changes names every card as the card data does** (work in the owner's gist `a4a9cad2e492af7c45c50a1cea7e6cf6`): each card line names its card by the card data's name (the lines `primogen build` reports, e.g. "Monacle of Clarity", "Tereza Rotas"; the 1999-04-01 promo line written as prose, not an item), and each `groups.google.com` announcement link replaced by its `usenet.krcg.org` copy; then `/sync`, and the build's near-name fallback deleted. Done when `primogen build` prints no `card-changes.md` line and no served announcement link points at `groups.google.com`. Docs: provenance.md (the copied commit), operations.md "Commands" (the build's warnings).
- **The archive's sweep does not miss or stop silently** [M:rules] (work in `newsgroup-archive`): the `--refresh` run scheduled weekly over recent topics; the `--expect` check fed every forum topic the rulings database cites; a monthly dated heartbeat commit so the schedule is not disabled. Done when the three run in its workflow. Docs: sources.md "The archive's limits".
- **A rules answer's evidence in one call** [M:rules]: the section index over rulebook, advanced rules, tournament rules and judges' guide; `vtes_rules_case`, `vtes_rules_search`, `vtes_rules_read`, `vtes_rulings_archive`; the evidence tests seeded from the calibration corpus (each rules entry's `settled-by` traced to today's sources), the codex cases, the `rules-qa` items as questions and the engine's rulings ledger. Done when the evidence tests pass and the owner has graded a rules session. Docs: product.md "Tools" and "Evaluation".

<!-- completed-since-upkeep: 5 -->
