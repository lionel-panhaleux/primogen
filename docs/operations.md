# Operations

Interviewed with the owner, 2026-11-10: same conventions as zero-vtes. The
commands are added here by the lines that create them.

## Environments

- **This machine**: sessions in the one checkout on `main`.
- **frankfurt** (the VPS shared with the engine and Archon beta): the home
  of `https://mcp.krcg.org`, not deployed yet. `primogen serve` listens on `127.0.0.1:8040` as
  the `primogen` systemd unit (a dynamic user, its bundle in
  `/var/lib/primogen`, logs under `journalctl -t primogen`), behind
  `server-setup`'s `nginx_site` and certbot. No secrets yet; the first one
  goes in sops, as in `server-setup`.

## Data refresh

The server refreshes itself; there is no cron job and no operator.

- A scheduler inside `primogen serve` checks the cheap sources at start and
  hourly: the card data's `Last-Modified`, the head commits of
  `vtes-rulings` and of primogen's `main` (its `knowledge/`). When one moved,
  it rebuilds. The heavy source, the archive's bulk file, is fetched at start
  and again after 06:00 UTC each night; it rebuilds when its newest post
  moved. A failed fetch keeps the archive held, and the cheap sources still
  refresh ([sources.md](sources.md) says how each source refreshes).
- A refresh builds a new bundle, checks it (no fewer cards than the live one;
  decks and the evidence tests join the check when they exist), saves it
  beside the live file and renames it over, then swaps it in between two tool
  calls. The saved file is the published bundle, the server's
  `/bundle.json.gz`.
- A failed build leaves the previous bundle serving. It shows on the status
  page, the server's `/`, and in the data date of every answer, which
  stops advancing.
- Archon data goes into a separate database that is never published.
- A local install downloads the published bundle on first run and again when
  a newer one exists, checked at most daily. Offline, it serves what it has
  and says how old it is.
- Each refresh reports what needs a person, on the status page and in the
  journal: the build's warnings (below). Decks that matched no archetype,
  cards with no module tag, evidence tests that started failing, an upstream
  rules diff waiting for sync and an archive whose newest post is old join
  them once primogen can compute them.

## Commands

Go, jq and uv come from mise: `mise install` in the checkout reads `mise.toml`.

- `primogen build` fetches the card data, `vtes-rulings` at its latest
  commit and the archive's bulk file, reads `knowledge/` in the current
  directory (`-knowledge` to change it), and writes the bundle (by default
  `data/bundle.json.gz`, `-bundle` to change it, and to query it). It
  prints what needs a person: a reference the archive does not hold, a ruling
  not in the rulings site's canonical form or citing no reference, a line of
  the history of changes naming no card exactly or under no dated heading,
  left out of the bundle: the fix is in the gist, then `/sync`.
- `primogen query cards NAME...` (`-detail` for the detailed mode) calls
  `vtes_cards` through the MCP server and prints what the model reads.
- `primogen stdio` serves the tools over stdio for a local MCP client.
- Without `-bundle`, `query` and `stdio` are a local install: they use the
  published bundle in the user cache directory, which only the download
  writes, as "Data refresh" says (`-from` changes where it downloads from,
  for the tests).
- `primogen serve` is the hosted server: the tools over Streamable HTTP at
  `/mcp` (stateless, JSON responses), the published bundle and the status
  page, refreshed as "Data refresh" says, keeping current the file given
  with `-bundle`. It builds one first if there is none, and reads
  `knowledge/` from GitHub unless given `-knowledge`.
- `mise run deploy` (`--dry` only shows) cross-compiles the binary for
  frankfurt and converges it there with pyinfra from `deploy/`, which pins
  `server-setup`'s package. It refuses to run until `mcp.krcg.org` resolves
  to frankfurt.

## Checks

`gofmt -l .` prints nothing, `go vet ./...` and `go test ./...` pass, in
a few seconds. The tests build a bundle from `testdata/upstream`, the real
upstreams pinned and served by a local file server (the whole card data,
`vtes-rulings` at one commit, the archive cut to its ids), with
`knowledge/card-changes.md` as it stands, then query it as a user does. They
also run the hosted server through its refreshes against that server (the
card data changed, a push to `main`, a quiet night, then the card data cut),
as an MCP client and a local install see it.
`testdata/upstream/refresh.sh` pins them again; the expectations are then
checked against the new data, and a `/sync` changing an entry a test names
updates that test in the same change.

## Releases and deploys

Trunk-based. A release is a tag; the owner tags and deploys. Knowledge files
travel inside the data bundle, and the hosted server reads them from `main`
on GitHub: a pushed knowledge change reaches users within the hour, without
a deploy. Pushing `knowledge/` is publishing it, and the deployed binary
reads it: a change to a knowledge file's format is deployed before it is
pushed.

## Looking things up

- MCP: the specification at `modelcontextprotocol.io/specification` (the
  revision in use is named in the code's dependency on the Go SDK,
  `github.com/modelcontextprotocol/go-sdk`).
- Libraries: the Context7 server, then the official pages.
- Claude Code: `code.claude.com/docs`.
- The KRCG services: [sources.md](sources.md), then each service's own
  documentation (`api.krcg.org`, `api.archon.vekn.net/docs`,
  `engine.krcg.org/docs`).
