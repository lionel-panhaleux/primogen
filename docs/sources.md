# Sources: the public VTES data stack

Who owns which fact, and how it reaches primogen. Read from each service's
repository and public pages on 2026-11-10; a claim about a service is checked
against that service before it is relied on.

## The sources

| Source | Owns | Refreshes | Access |
|---|---|---|---|
| `static.krcg.org/data/v4` (`vtes.json`, `twda.json`) | Cards with their rulings; the tournament winning deck archive | Every 6 hours, and at once when a ruling is approved | Public files |
| `vtes-biased/vtes-rulings` (`rulings.yaml`, `groups.yaml`, `references.yaml`) | The rulings database: the durable source of truth for rulings, the card groups they apply to, the URL of each reference | Pushed by the rulings site on each approval | Public repository |
| `rulings.krcg.org` | Curation of the rulings database by the rulemongers; proposals under discussion | Live | Pages public; a read API under `/api` (undocumented, public access unchecked) |
| `usenet.krcg.org` (`vtes-biased/newsgroup-archive`) | The rulings archive: 14,417 threads in which a rules director took part, 1994 to today (the newsgroup, the V:EKN forum, BoardGameGeek) | Nightly forum sweep at 04:17 UTC | Pages and a search index; every thread in one file, `usenet.krcg.org/archive.json.gz` (below) |
| `vtes-biased/vtes-advanced-rules` | The 2024 rulebook text and the advanced rules | By hand | Public repository |
| `api.archon.vekn.net` | Tournaments, published decklists, members and ratings, community links | Export rebuilt within 15 minutes of a change | App token (below) |
| `engine.krcg.org` | Rules enforcement for the cards it covers; bots | On each engine release | Public API |
| `api.krcg.org` | Card search and decklist conversion | Live | Public API |

## The rulings chain

A rulemonger approves a proposal on `rulings.krcg.org`; the site writes
`vtes-rulings`; that repository feeds the card data at `static.krcg.org`.
Each ruling ends with a reference id (`[LSJ 20040518]`), which
`references.yaml` resolves to an exact message in the archive
(`usenet.krcg.org/t/<thread>/#m<n>`).

A ruling has no id of its own in the YAML. The rulings site identifies it by
a stable hash of its text and links to it as
`rulings.krcg.org/index.html?uid=<card id>#r-<hash>`; the id changes when the
wording is edited. A ruling ending in `[REMINDER]` restates card text or a
core rule and may carry no reference.

## The archive's bulk file

`usenet.krcg.org/archive.json.gz` (about 57 MB) is rebuilt with the site on
every push and every sweep that finds something. Unzipped, it is one JSON
object: `LastPost`, the date of the newest post the archive holds (ISO, no
time zone), and `Threads`, every thread file as the repository holds it, in
the format its README describes.

GitHub Pages stamps its `ETag` and `Last-Modified` at each deploy, not from
the content, so a `200` does not mean a thread changed: a pass that must know
compares the unzipped content.

## The archive's limits

- A plain nightly sweep does not see a reply posted after the director's last
  post in a topic; only its `--refresh` run does.
- The forum's search under-reports (5,609 of the 6,947 posts on Ankha's
  profile), so a topic can be missing.
- GitHub disables a scheduled workflow after 60 days without a commit, which
  would stop the sweep.
- Nothing published records when a sweep last ran: a sweep that finds nothing
  commits and deploys nothing, and `LastPost` is the newest post's date. A
  stopped sweep and a quiet forum look the same from outside.

## Archon

- The public API needs an app token (client credentials; the app's owner
  holds the DEV role). `/v1/export` is the whole corpus as one file.
- Members' names and ratings are visible to VEKN members only. Primogen
  serves Archon data only to a user logged in through Archon OAuth
  (owner, 2026-11-10: a GDPR matter), except community links
  ([product.md](product.md)).
- Archon's OAuth wants pre-registered clients with exact callback URLs.
- A member is identified by `/oauth/userinfo` with the `profile:read` grant.

## The engine

`engine.krcg.org` covers part of the card pool (`/cards` lists it) and hosts
one human per game. It has no endpoint to check a described position and no
bot-only tables.
