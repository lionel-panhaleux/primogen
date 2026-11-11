# Sources: the public VTES data stack

Who owns which fact, and how it reaches primogen. Read from each service's
repository and public pages on 2026-11-10 and 2026-11-11; a claim about a
service is checked against that service before it is relied on.

## The sources

| Source | Owns | Refreshes | Access |
|---|---|---|---|
| `static.krcg.org/data/v4` (`vtes.json`, `twda.json`) | Cards with their rulings; the tournament winning deck archive | Every 6 hours, and at once when a ruling is approved | Public files |
| `vtes-biased/vtes-rulings` (`rulings.yaml`, `groups.yaml`, `references.yaml`) | The rulings database: the durable source of truth for rulings, the card groups they apply to, the URL of each reference | Pushed by the rulings site on each approval | Public repository |
| `rulings.krcg.org` | Curation of the rulings database by the rulemongers; proposals under discussion | Live | Pages public; a public read API (below) |
| `usenet.krcg.org` (`vtes-biased/newsgroup-archive`) | The rulings archive: 14,417 threads in which a rules director took part, 1994 to today (the newsgroup, the V:EKN forum, BoardGameGeek) | Nightly forum sweep at 04:17 UTC | Pages and a search index; every thread in one file, `usenet.krcg.org/archive.json.gz` (below) |
| `vtes-biased/vtes-advanced-rules` | The 2024 rulebook text and the advanced rules | By hand; primogen copies them with `/sync` ([provenance.md](provenance.md)) | Public repository |
| `api.archon.vekn.net` | Tournaments, published decklists, members and ratings, community links | Export rebuilt within 15 minutes of a change | App token (below) |
| `engine.krcg.org` | Rules enforcement for the cards it covers; bots | On each engine release | Public API |
| `api.krcg.org` | Card search and decklist conversion | Live | Public API |
| `gist.github.com/lionel-panhaleux/a4a9cad2e492af7c45c50a1cea7e6cf6` | The history of card and rules changes: each errata and rules change, dated, linking its announcement | By hand, by the owner; primogen copies it with `/sync` ([provenance.md](provenance.md)) | Public gist |
| `vekn.net` (`/tournament-rules`, `/judges-guide`, `/code-of-ethics`) | The official tournament rules, judges' guide and code of ethics | When the VEKN publishes a revision | Public pages |

## The rulings chain

A rulemonger approves a proposal on `rulings.krcg.org`; the site writes
`vtes-rulings`; the copy packaged with `krcg` feeds the card data at
`static.krcg.org`, and lags: on 2026-11-11 its references still pointed at
Google Groups where `references.yaml` points at the archive. Primogen reads
`vtes-rulings` itself.

Each ruling ends with reference ids (`[LSJ 20040518]`), which
`references.yaml` resolves to a URL. A newsgroup-era reference points into
the archive (`usenet.krcg.org/t/<thread>/#m<n>`, the anchor positional); a
forum reference points at `vekn.net/forum/<category>/<topic>-<slug>#<post>`,
which the archive holds as `usenet.krcg.org/t/vekn-<topic>/#<post>` (the
`newsgroup-archive` README). Rules Team rulings on `blackchantry.com` and
rulebook references are not in the archive.

A ruling has no id of its own in the YAML. The rulings site's id is the first
5 bytes of SHAKE128 over the ruling's text (`[REMINDER]` stripped, each card
token `{100788|Freak Drive}` cut to `{100788}`), in base32; for a group
ruling, the group's text, whatever prefix or override a card shows. The card
page carries every ruling that applies, group rulings included, at
`rulings.krcg.org/index.html?uid=<card id>#r-<id>`; the id changes when the
wording is edited. A ruling ending in `[REMINDER]` restates card text or a
core rule and may carry no reference.

## The rulings site's read API

Public by the owner's decision (2026-11-11), JSON, no login; checked
2026-11-11: `GET /api/card/<id>` (the card, its rulings with their ids, its
groups, the rulings naming it), `/api/search?query=<text>` (cards, groups and
rulings matching), `/api/group` and `/api/group/<id>`, `/api/reference` (every
reference id and its URL).

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
