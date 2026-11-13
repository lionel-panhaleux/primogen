# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **The advanced rules' rulebook links point to archon.vekn.net**: every `[RBK …]` citation in `vtes-advanced-rules` (162, 56 distinct, all to `vekn.net/rulebook` today) links the id `archon.vekn.net/help/rules` gives the section it cited, as the rulings database's references do, done there through its own harness, then `/sync`. Done when `knowledge/rules/advanced-rules.md` holds no `vekn.net/rulebook` link. Docs: sources.md (the vekn.net row).
- **`mcp.krcg.org` lists the KRCG stack for developers and their agents**: the root and `/llms.txt` serve one catalog, authored in `knowledge/`, of the public services (the MCP endpoint and the published bundle, `static.krcg.org`, `api.krcg.org`, the rulings site's read API, the archive's bulk file, the engine, Archon's API and its token rule, the upstream repositories), each linking its own live specification; the status page moves to `/status`. Done when `curl https://mcp.krcg.org/` and `/llms.txt` return the catalog, every URL in it answers at each refresh (a dead one is a warning on the status page), and a fresh agent given only `https://mcp.krcg.org` finds the card data and the MCP endpoint. Docs: product.md "Connecting" (authored, linking each live specification, links checked at every refresh, replaces "generated from the live specifications"), operations.md, provenance.md, roadmap.md (the `llms.txt` item leaves "Later").

<!-- completed-since-upkeep: 15 -->
