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
| `usenet.krcg.org` (`vtes-biased/newsgroup-archive`) | The rulings archive: 14,418 threads, 1994 to today (the newsgroup, the V:EKN forum, BoardGameGeek), each one a rules director took part in, or that a ruling or the history of changes cites (the 1998-08-01 DCI floor rules) | Nightly forum sweep at 04:17 UTC | Pages and a search index; every thread in one file, `usenet.krcg.org/archive.json.gz` (below) |
| `vtes-biased/vtes-advanced-rules` | The advanced rules | By hand; primogen copies them with `/sync` ([provenance.md](provenance.md)) | Public repository |
| `vtes-biased/archon-vibe` (`frontend/src/lib/help-content/`) | The official rulebook, tournament rules, judges' guide and code of ethics, as markdown, served at `archon.vekn.net/help/<page>`, each heading's id made by archon's renderer (`frontend/src/lib/markdown.ts`) | On each push to `main`, deployed with Archon; primogen probes the last commit touching the documents, and warns when the renderer moved past the one it reproduces | Public repository; pages rendered in the browser, which scroll to an id once drawn |
| `api.archon.vekn.net` | Tournaments, published decklists, members and ratings, community links | Export rebuilt within 15 minutes of a change | App token (below) |
| `engine.krcg.org` | Rules enforcement for the cards it covers; bots | On each engine release | Public API |
| `api.krcg.org` | Card search and decklist conversion | Live | Public API |
| `gist.github.com/lionel-panhaleux/a4a9cad2e492af7c45c50a1cea7e6cf6` | The history of card and rules changes: each errata and rules change, dated (one section, 1995-10, by its month), linking its announcement (a newsgroup or forum post at its `usenet.krcg.org` copy) | By hand, by the owner; primogen copies it with `/sync` ([provenance.md](provenance.md)) | Public gist |
| `vekn.net` (`/rulebook`, `/tournament-rules`, `/judges-guide`) | Other copies of the same documents, worded apart in places; Archon's are the official, maintained ones (owner, 2026-11-12); the rulebook's anchors are what the `RBK` labels of the rulings database and the advanced rules are named after | When the VEKN publishes a revision | Public pages; primogen reads none |

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
`newsgroup-archive` README). Rules Team rulings on `blackchantry.com` are
not in the archive. A rulebook reference (`RBK <label>`, a closed set of 140
the rulings site offers and never adds to) links an id archon's renderer
gives `archon.vekn.net/help/rules` (a heading, a box or a bold term), the
label still the vekn.net anchor or page it linked until November 2026;
primogen joins it to the section that id opens, and its build warns of a
cited one no section opens. The advanced rules' `[RBK <label>]` citations
link the URL `references.yaml` gives the label. An id is its element's words
slugged: an element archon rewords breaks the links to it.

A ruling has no id of its own in the YAML. The rulings site's id is the first
5 bytes of SHAKE128 over the ruling's text (`[REMINDER]` stripped, each card
token `{100788|Freak Drive}` cut to `{100788}`), in base32; for a group
ruling, the group's text, whatever prefix or override a card shows. The card
page carries every ruling that applies, group rulings included, at
`rulings.krcg.org/index.html?uid=<card id>#r-<id>`; the id changes when the
wording is edited. A ruling ending in `[REMINDER]` restates card text or a
core rule and may carry no reference.

## The rules directors

The succession, as `krcg`'s `RULING_AUTHORS` dates it (`krcg/rulings.py`),
each start checked against its announcement in the archive where one was
found (2026-11-11 and 2026-11-12). Ankha's is a span: krcg's 2016-12-04 is the
day the V:EKN announced he would take over, not the day he did. Wizards of
the Coast's net representatives answered the newsgroup until the V:EKN took
the seat.

| Label | Director | From | Announced |
|---|---|---|---|
| `TOM` | Thomas R. Wylie | 1994-12-15 | none held: he posted for Wizards of the Coast from 1994-04 |
| `SFC` | Shawn F. Carnes | 1996-07-29 | ["I guess I'm the net NetRep!"](https://usenet.krcg.org/t/WzCRgg6ZYZ0/#m0) |
| `JON` | Jon Wilkie | 1996-10-18 | ["OFFICIAL: New NetRep!"](https://usenet.krcg.org/t/p43vjjRTpds/#m0), by Carnes |
| `LSJ` | L. Scott Johnson | 1997-02-24 | ["New Net.Rep"](https://usenet.krcg.org/t/w3et8cc8unk/#m0); his last ruling 2011-06-11, on BoardGameGeek |
| `PIB` | Pascal Bertrand | 2011-07-06 | [the V:EKN's press release](https://usenet.krcg.org/t/vekn-5958/#m0) |
| `ANK` | Vincent Ripoll (Ankha) | between 2016-12-04 and 2017-01-11 | [the V:EKN's new Inner Circle](https://usenet.krcg.org/t/vekn-75402/#m0), 2016-12-04: Bertrand stays Rules Director "until he finishes the new qualification system", then "Vincent Ripoll will take over"; the first ruling the database cites by Ankha is of 2017-01-11 |

The archive spells them several ways (`LSJ`, `L. Scott Johnson
(Rulemonger)`, `Tom Wylie`: the rulings site's `USENET_AUTHORS`). Primogen
calls a post's author the sitting director only within these dates; between
the Inner Circle's announcement and Ankha's first cited ruling, either may
have been sitting, and a post of theirs there says so.

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

- A reply posted after the director's last post in a topic is seen only by
  the weekly refresh, which reads again the topics with a post in the last 60
  days: a reply to a topic quiet for longer is missed.
- The forum's search under-reports (5,609 of the 6,947 posts on Ankha's
  profile), so a topic no ruling cites can be missing. Every forum topic the
  rulings database cites by its forum link is checked each night and fetched
  if absent; one the forum no longer serves fails the sweep.
- GitHub disables a scheduled workflow after 60 days without repository
  activity, a term its docs leave undefined. A monthly empty commit, dated in
  its message, is the activity.
- Nothing published records when a sweep last ran: a sweep that finds nothing
  commits and deploys nothing, and `LastPost` is the newest post's date. A
  failed sweep goes red in the repository's Actions tab, but from outside a
  stopped sweep and a quiet forum look the same.

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
