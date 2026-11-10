# Operations

Interviewed with the owner, 2026-11-10: same conventions as zero-vtes. The
commands are added here by the lines that create them.

## Environments

- **This machine**: sessions in the one checkout on `main`.
- **frankfurt** (the VPS shared with the engine and Archon beta): serves
  `https://mcp.krcg.org`, deployed with pyinfra on `server-setup`, a systemd
  unit behind nginx and certbot, secrets in sops.

## Data refresh

The server refreshes itself; there is no cron job and no operator.

- A scheduler inside `primogen serve` checks the cheap sources hourly and the
  heavy ones nightly ([sources.md](sources.md) says how each one refreshes).
- A refresh builds a new bundle beside the live one, checks it (card and deck
  counts did not fall, the evidence tests pass), swaps it in between two
  requests, and publishes the public bundle for download.
- A failed build leaves the previous bundle serving. It shows on the status
  page and in the data date of every answer, which stops advancing.
- Archon data goes into a separate database that is never published.
- A local install downloads the published bundle on first run and again when
  a newer one exists, checked at most daily. Offline, it serves what it has
  and says how old it is.
- Each refresh reports what needs a person: decks that matched no archetype,
  cards with no module tag, evidence tests that started failing, an upstream
  rules diff waiting for sync, an archive whose last sweep is old.

## Releases and deploys

Trunk-based. A release is a tag; the owner tags and deploys. Knowledge files
travel inside the data bundle, so a knowledge change reaches users at the
next refresh without a new binary.

## Looking things up

- MCP: the specification at `modelcontextprotocol.io/specification` (the
  revision in use is named in the code's dependency on the Go SDK,
  `github.com/modelcontextprotocol/go-sdk`).
- Libraries: the Context7 server, then the official pages.
- Claude Code: `code.claude.com/docs`.
- The KRCG services: [sources.md](sources.md), then each service's own
  documentation (`api.krcg.org`, `api.archon.vekn.net/docs`,
  `engine.krcg.org/docs`).
