# Operations

Interviewed with the owner, 2026-11-10: same conventions as zero-vtes. The
commands are added here by the lines that create them.

## Environments

- **This machine**: sessions in the one checkout on `main`.
- **frankfurt** (the VPS shared with the engine and Archon beta): serves
  `https://mcp.krcg.org`. `primogen serve` listens on `127.0.0.1:8050` as
  the `primogen` systemd unit (a dynamic user, its bundle in
  `/var/lib/primogen`, logs under `journalctl -t primogen`), behind
  `server-setup`'s `nginx_site` and certbot, which passes the client's
  address as `X-Real-IP`, the address `vtes_feedback`'s limit counts. The
  limit's counts are held in memory: a restart or a deploy resets them.

## Secrets

`secrets.sops.yaml`, encrypted with sops to the SSH keys `.sops.yaml` lists
(the owner's, and server-setup's `deploy` key, which the release workflow
holds), holds `github_app_key`, the private key of the
GitHub App `app-feedback-issues` (Issues read and write, no webhook),
installed on several of the owner's repositories: each token the server
asks for is scoped to this repository's issues; and `archon_client_id` and
`archon_client_secret`, the credentials of the Archon OAuth client (`api:read`
and `profile:read`, callback `https://mcp.krcg.org/oauth/callback`). `mise run deploy` decrypts them and puts them
on frankfurt as `/etc/primogen/github-app.pem` and `/etc/primogen/archon-secret` (root, 600, never diffed),
handed to the dynamic user through systemd's `LoadCredential`; the GitHub App's
client ID, not a secret, is in `deploy/deploy.py`. `sops edit
secrets.sops.yaml` edits it; a new key is generated on the App's settings
page, the old one deleted there after the deploy.

## Data refresh

The server refreshes itself; there is no cron job and no operator.

- A scheduler inside `primogen serve` checks the cheap sources at start and
  hourly: the `Last-Modified` of the card data and of the TWDA, the head commits of
  `vtes-rulings` and of primogen's `main` (its `knowledge/`), and the last
  commit touching `archon-vibe`'s documents (its rulebook, tournament rules
  and judges' guide), so that a push elsewhere in the app moves nothing.
  When one moved, it rebuilds. The heavy source, the archive's bulk file, is fetched
  at start and again after 06:00 UTC each night; it rebuilds when its newest post
  moved. A failed fetch keeps the archive held, and the cheap sources still
  refresh ([sources.md](sources.md) says how each source refreshes). The
  archive is held whole in memory for `vtes_rulings_archive` and the rules
  directors' posts of `vtes_rules_case` (about 750 MB
  resident at its peak, measured on a local query, and briefly twice that
  while a nightly fetch replaces it); the bundle keeps only the posts the
  rulings and the advanced rules cite.
- A refresh builds a new bundle, checks it (no fewer cards than the live one,
  and no fewer TWDA decks than 99% of its, since a deck may leave the TWDA),
  saves it
  beside the live file and renames it over, then swaps it in between two tool
  calls. The saved file is the published bundle, the server's
  `/bundle.json.gz`.
- A failed build leaves the previous bundle serving. It shows on the status
  page, the server's `/status`, and in the data date of every answer, which
  stops advancing.
- With `-archon-client` and `-archon-secret`, `primogen serve` fetches
  `api.archon.vekn.net/v1/export` with a `client_credentials` (`api:read`) app
  token at start and every 15 minutes, verifies its `header` and `eof` row
  count, and swaps it in memory on `Live`, never into the published bundle. A
  failed fetch leaves the previous export serving and shows on `/status`.
- A local install downloads the published bundle on first run and again when
  a newer one exists, checked at most daily, and the archive's bulk file the
  same way, from `usenet.krcg.org`, the first time a tool needs it
  (`vtes_rulings_archive`, or `vtes_rules_case` given a single card or two
  cards no ruling names together). Loading it takes about 5 seconds (measured on a local
  `primogen query`, against a tenth of a second for a call that needs no
  archive); `primogen stdio` loads it once. Offline, it serves what it has
  and says how old it is.
- Each refresh reports what needs a person, on the status page and in the
  journal: the build's warnings (below), and every link of the catalog
  (`knowledge/llms.txt`) that does not answer a GET: no answer within 30
  seconds, a 404, a 410 or a server error (an endpoint that refuses a GET or
  wants a token answers). An upstream rules diff waiting for sync and an
  archive whose newest post is old join them once primogen can compute
  them.

## Commands

Go, jq, node, sops and uv come from mise: `mise install` in the checkout reads
`mise.toml`.

- `primogen build` fetches the card data, the TWDA, `vtes-rulings` and `archon-vibe`
  at their latest commits and the archive's bulk file, reads `knowledge/` in
  the current
  directory (`-knowledge` to change it), and writes the bundle (by default
  `data/bundle.json.gz`, `-bundle` to change it, and to query it). It
  prints what needs a person: a reference the archive does not hold, a ruling
  not in the rulings site's canonical form or citing no reference, a line of
  the history of changes naming no card exactly or under no dated heading,
  left out of the bundle: the fix is in the gist, then `/sync`; a TWDA deck
  naming a card the card data lacks, left out until the card data has it;
  each archetype it proposes from the decks of the archetype labels' window
  in none ([provenance.md](provenance.md) `classification.json`), to name
  or reject there, and the decks the owner has not labelled that match no
  archetype nor proposal (a new one, or a pass over the labels); a name of
  `knowledge/guide/module-tags.json` no card has, left out, a module there
  that `modules.md` does not name under its section, or a module of
  `modules.md` with no tags and no reason under `untagged`, a reason there
  for no such module, a module needing more pieces than it tags, a defense there that `modules.md`'s "Combat defenses" does not name or one it names without tags, an answer's level naming a discipline its card does not require, or
  `theory.md` no longer reading the density classes `vtes_deck` applies; how many library cards the TWDA
  decks of that window play and no module tags, the 50 most played named
  with how many decks play them (a pass over the tags); the library cards
  of that window whose text gains pool or blood (a pattern in the code:
  "gain N pool", "from the blood bank", "steal N blood" and their kin) and
  that the tags' `economy` neither lists nor excludes with a reason, a
  key of its `sources` that is none of its `kinds`; any TWDA `Place` in
  the archetype labels' window that maps to no country and is not Online; a post the
  advanced rules cite that the archive does not hold, or a ruling's rulebook
  reference no section of the rulebook opens: the fix is upstream, unless
  archon gives that id to a paragraph `archon.go` opens no section at; and
  archon's renderer changed past the commit `archon.go` reproduces: re-pin
  the fixtures (`testdata/upstream/refresh.sh`), then `TestArchonLinks`
  says whether the links still land.
- `primogen audit [-bundle PATH] [-archive PATH|URL] [-since YYYY-MM-DD] [-twda] [-out PATH] [CARD...]`
  reads the bundle and the rulings archive (a local path or URL, cached like
  `query` when a URL) and writes a Markdown report for the rulemongers listing,
  by card (most played in the TWDA window first) and by card pair, the sitting
  rules directors' posts naming them (in their own words or in the question
  they answer, outside a decklist) that no ruling on those cards cites in
  `vtes-rulings`, each with its link, date, author standing, quoted question
  and answer, existing rulings on the card, and any entry of `card-changes.md`
  dated after the post, skipping non-rules forum categories (event reports/TWD,
  deck clinic, logistics, off-topic) and TWD or tournament-report threads
  (`-since` for a date lower bound, `-twda` for cards played in the TWDA window
  only, `-out` to write to a file, and optional card names to narrow the audit).
- `primogen query TOOL ARG...` (`-detail` for the detailed mode) calls a
  tool through the MCP server and prints what the model reads: `cards
  NAME...`, `case QUESTION [CARD...]`, `search WORDS...`, `read ID...`,
  `archive WORDS...` (`-directors` for the sitting rules directors' posts
  only), `deck LIST|ID|LINK` (the list's lines as one or
  several arguments, or `-` to read it from stdin; a link of VDB, VTES Decks,
  Amaranth or Archon; `-draft` for a deck idea's cards, `-library-size` for a draft's
  library size, `-together` for hand-odds recipes as JSON), `stats CARD
  [ARCHETYPE]`, `meta [ARCHETYPE]` (`-months` for the window's length,
  `-country` and `-continent` for geographic filters),
  `events [TOURNAMENT]` (`-member`, `-rankings`, `-category`, `-league`,
  `-links`, `-country`, `-since`, `-until`, `-rank`, `-format`, `-state`),
  `guide TOPIC [SECTION...]` (`-lang` for game-term translations, `-filter` to
  narrow terms or examples),
  `feedback QUESTION WRONG EXPECTED` (files a public issue, through the
  hosted server, or the one `-from` names).
- `primogen stdio` serves the tools over stdio for a local MCP client.
  Like `query`, it files `vtes_feedback`'s reports through the hosted
  server's `/mcp`, at the host `-from` names.
- Without `-bundle`, `audit`, `query` and `stdio` are a local install: they use the
  published bundle in the user cache directory, which only the download
  writes, as "Data refresh" says (`-from` and `-archive` change where the
  bundle and the archive are downloaded from, `-sites` where the
  deckbuilders' APIs are read, each under its host, for the tests).
- `primogen serve` is the hosted server: the tools over Streamable HTTP at
  `/mcp` (and `/mcp/member`, stateless, JSON responses), the published bundle, the catalog of
  the KRCG stack at `/` and `/llms.txt`, and the status page at `/status`,
  refreshed as "Data refresh" says, keeping current the file given
  with `-bundle`. With `-github-app` (the App's client ID) and `-github-key`
  (its PEM file), it files `vtes_feedback`'s reports; without, the tool
  refuses, pointing to the repository's issues. With `-archon-client` (the
  Archon OAuth client ID) and `-archon-secret` (a file holding its secret),
  it serves the OAuth 2.1 login endpoints bridging to Archon and refreshes
  `api.archon.vekn.net/v1/export` in memory every 15 minutes. It builds a bundle first if there is none, and reads
  `knowledge/` from GitHub unless given `-knowledge`.
- `gh workflow run release` runs `mise run release` on a GitHub Actions
  runner (`.github/workflows/release.yml`), as server-setup's `deploy`
  user: its key, the `production` environment's `DEPLOY_SSH_KEY` (pushed
  by server-setup's `just sync-key`, primogen being in its
  `deploy_targets.py`), reaches frankfurt and decrypts the secrets.
- `mise run release` releases: on a clean `main` not behind `origin`, it
  runs the checks, tags the next `v0.N.0` (none if HEAD has a tag already),
  deploys with pyinfra without asking, pushes `main` and the tag, and fails
  unless `mcp.krcg.org` then reports that tag as its version.
- `mise run deploy` (`--dry` only shows) cross-compiles the binary for
  frankfurt and converges it there with pyinfra from `deploy/`, which pins
  `server-setup`'s package, showing every change and asking first. It
  refuses to run until `mcp.krcg.org` resolves to frankfurt.

## Checks

`gofmt -l .` prints nothing, `go vet ./...` and `go test ./...` pass, in
about half a minute. The tests build a bundle from `testdata/upstream`, the
real upstreams pinned and served by a local file server (the whole card
data with its sets, the TWDA cut to the archetype labels' window and three
decks the tests name, the deckbuilders' APIs, each under its host (one deck as VDB's and VTES Decks's answer it, the graded Banu list as VDB's, an Amaranth deck, the answer for a deck Amaranth lacks, and Amaranth's card catalog), `vtes-rulings` and `archon-vibe`'s three documents at one commit
each, the ids archon.vekn.net gives those documents' headings by archon's
own renderer at that commit, which every link of theirs and of the
advanced rules to them is checked against, the archive cut to
its ids, titles and URLs with the text of the posts the rulings and the advanced
rules cite, of the threads the calibration corpus's `settled-by`
names and of four threads the tests name), with `knowledge/` as it stands, then query it as a user does,
and run the evidence tests ([product.md](product.md) "Evaluation") on it,
with zero-vtes's `rules-qa` and rulings ledger pinned beside them. They
also run the hosted server through its refreshes against that server (the
card data changed, a link of the catalog dead, a push to `main`, a quiet
night, then the card data cut, the catalog's links answered by that server),
as an MCP client and a local install see it.
`testdata/upstream/refresh.sh` pins them again (it runs archon's renderer,
code fetched from GitHub, under node: a session asks the owner first); the expectations are then
checked against the new data, and a `/sync` changing an entry a test names
updates that test in the same change.

## Releases and deploys

Trunk-based. Every deploy is a release, a tag `v0.N.0` the binary reports
as its version; `mise run release` does it all. When the owner says
deploy, a session runs the release workflow, which releases what is
already on `origin/main`; nothing is deployed untagged. Knowledge files
travel inside the data bundle, and the hosted server reads them from `main`
on GitHub: a pushed knowledge change reaches users within the hour, without
a deploy. Pushing `knowledge/` is publishing it, and the deployed binary
reads it: a change to a knowledge file's format is deployed before it is
pushed, by `mise run release` from this machine (which deploys before it
pushes), and a binary reading a new knowledge file after that file is
pushed on its own.

## Looking things up

- MCP: the specification at `modelcontextprotocol.io/specification` (the
  revision in use is named in the code's dependency on the Go SDK,
  `github.com/modelcontextprotocol/go-sdk`).
- Libraries: the Context7 server, then the official pages.
- Claude Code: `code.claude.com/docs`; Antigravity: `antigravity.google/docs`.
- The KRCG services: [sources.md](sources.md), then each service's own
  documentation (`api.krcg.org`, `api.archon.vekn.net/docs`,
  `engine.krcg.org/docs`).
