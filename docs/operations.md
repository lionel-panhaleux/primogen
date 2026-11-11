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
  rules diff waiting for sync, an archive whose newest post is old.

## Commands

Go and jq come from mise: `mise install` in the checkout reads `mise.toml`.

- `primogen build` fetches the card data, `vtes-rulings` at its latest
  commit and the archive's bulk file, reads `knowledge/` in the current
  directory (`-knowledge` to change it), and writes the bundle (by default
  `bundle.json.gz` in the user cache directory, `-bundle` to change it). It
  prints what needs a person: a reference the archive does not hold, a ruling
  not in the rulings site's canonical form or citing no reference, a line of
  the history of changes naming no card exactly (filed as possible under the
  closest cards, or under none).
- `primogen query cards NAME...` (`-detail` for the detailed mode) calls
  `vtes_cards` through the MCP server and prints what the model reads.
- `primogen stdio` serves the tools over stdio for a local MCP client.

## Checks

`gofmt -l .` prints nothing, `go vet ./...` and `go test ./...` pass, in
about a second. The tests build a bundle from `testdata/upstream`, the real
upstreams pinned and served by a local file server (the whole card data,
`vtes-rulings` at one commit, the archive cut to its ids), with
`knowledge/card-changes.md` as it stands, then query it as a user does.
`testdata/upstream/refresh.sh` pins them again; the expectations are then
checked against the new data, and a `/sync` changing an entry a test names
updates that test in the same change.

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
