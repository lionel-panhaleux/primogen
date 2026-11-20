# Product

Decided with the owner, 2026-11-10. This page states the intended product;
[`../AGENTS.md`](../AGENTS.md) "Right now" says how much of it exists.

## Who it is for

VTES players and judges using any LLM assistant (Claude, ChatGPT, Gemini,
open-source models). Three questions set the bar:

- "Can I play a Freak Drive after I played superior Kiss of Ra?"
- "Here is my decklist, what could I improve?"
- "What's the current meta?"

The quality bar is the owner's grading, and the target is clearly better than
the `vtes` skill and the codex strategist it replaces. Their failures:
confident wrong statements about cards and rulings, wrong answers on subtle
interactions, deck analysis that misses how the deck works as a whole,
generic or wrong swaps.

## What it is

An MCP server with no LLM on the server side. It returns evidence and
computed facts; the user's model writes the answer. One Go binary serves the
hosted endpoint and local installs.

Primogen supersedes the `vtes-skill` repository, whose knowledge, sync and
calibration moved here ([provenance.md](provenance.md)). It owns the TWDA analytics (clustering, archetype
classification, best cards), ported to Go from codex-of-the-damned, which
then consumes primogen's output.

## Tools

`vtes_cards`, the four rules tools, `vtes_deck`, `vtes_card_stats` and `vtes_meta` have shipped and
keep their names; the others are provisional until they ship.

| Tool | What it returns |
|---|---|
| `vtes_cards` | Cards resolved from names (exact, an English name before another card's translated one; translated; or one or two letters off a single card): full text, requirements and costs (a card named after a Sabbat Path a vampire follows, as Path of Death and the Soul is, said not to be that Path, with the rulebook section on Paths), legality, every ruling with its `rulings.krcg.org` link and each cited post's link in the rulings archive, the card's entries in the history of changes (an entry whose name several cards share, or an older spelling, says so). An unknown name returns the closest names and no card. The detailed mode adds the printings and the translated texts. |
| `vtes_rules_case` | For a question and its cards, given, or else found in the question (a card's name, or words opening or ending a single card's name, the name in the possessive read without its 's; never a game term of `knowledge/`, a player's shorthand or a phrase of the rulebook; an uncertain match, or another card when some are given, only suggested): each card's verbatim text and requirements and a breakdown computed from it (type, effects by level, timing and limits); every ruling with its link (the closest to the question only when all would not fit, with how many more `vtes_cards` gives), and the original post of the closest with its author's role then; the rulings on other cards naming two of its cards; for a card asked about alone, and for two cards no ruling names together, the posts of the rules director then sitting whose own words name it, or both within a few sentences, outside a decklist, or else whose question does (the lines a reply quotes, kept as written from the quoted block naming them through the answer after it, since one reply may answer several questions; else the post before it, named with its author and link, which the reply may not be answering), the closest to the question and to the words around the names first, at most four of a thread, none a ruling of these cards already cites, each with its author's standing and link and holding as the cards stood at its date (the rulings database misses some of what the rules directors settled); the rules sections joined to the cards and those closest to the question's words, with their footnotes; what no source covers, stated, with how many directors' posts exist beyond those shown or why the archive was not searched; and `knowledge/guide/rules-answer.md` as the answer contract. |
| `vtes_rules_search`, `vtes_rules_read` | Keyword search over the rulebook, the advanced rules, the tournament rules and the judges' guide by section, each labelled with its authority; a section read whole by id, with the sections under it and its footnotes. |
| `vtes_rulings_archive` | Full-text search of the rulings archive over what each author wrote (not what they quoted), and for a reply of the sitting (or possibly sitting) rules director over its thread's title and the question it answers as well (the lines it quotes, else the post before it: "Both, yes." names nothing): posts with author, date, whether the author was the sitting rules director then ([sources.md](sources.md) "The rules directors"), and the citable link; a reply found by its question ranks after the posts holding the words themselves, has three places of its own among those shown (eight in detail), and comes after that question, quoted as `vtes_rules_case` quotes it. |
| `vtes_deck` | A decklist (text, one card per line with its count, crypt lines' columns cut; a TWDA id; or a deck link: from VDB, the cards in it, a stored deck read from VDB, or VDB's own link to a TWDA deck read from the bundle; from VTES Decks, a stored deck read from it, or its link to a TWDA deck read from the bundle; from Amaranth, a stored deck whose cards, Amaranth's own ids, are named by its card catalog, each name read as the card's own (an advanced vampire only where Amaranth marks it, a name several groups share taking the group of the rest of the crypt), a name the card data lacks refused; a link of another site refused, the list as text asked for; or, with `draft`, a deck idea's cards, a line without a count read as one copy, its sizes not checked, a list refused as too small pointed there, a TWDA deck read whole), its names resolved as `vtes_cards` resolves them (a vampire's name several groups share taking the group of the rest of the crypt), validated by the rulebook's deck construction and the cards' bans and legality today: a list illegal today refused with every fix it needs, a draft's other fixes stated, a TWDA deck read as archived with what makes it illegal today stated, an unknown name refused with the closest names. Then a dossier within its bound (the relatives' details, then card texts, cut first; a list too large even then refused): a TWDA deck's event and its author's comments; the crypt with each vampire's text (a draft without one says its requirements go unchecked) and the vampires by discipline and level; which crypt cards have what each library card requires as the card data holds it (for a combination card, every discipline one of its effects names, read line by line from its text, since some of its effects need one alone; else any one of its disciplines; with its level), a requirement only its text states quoted as unchecked; the library by type with each card's text and its requirements as `vtes_cards` gives them; the library by module (the modules of `knowledge/guide/modules.md` of which the deck holds at least one core card, and two of its cards for a star engine naming several, by the tags of `module-tags.json`, a card counted in each of its modules): each module's copies and, but for a draft, its density, its class in `theory.md` (whose bounds the tool's code holds too: the build warns when `theory.md` no longer reads them) and the copies a hand of 7 holds on average; the exact chance an opening hand holds a copy (a draft's library taken as 90 cards unless `library_size`, 60 to 90, gives its size, or its own copies when more; `library_size` left aside for a list, said); its cards (core, support, those in the crypt) and, cut first with the relatives' details, what it lacks of its core; the modules with one library copy or none on one line (a draft's shown each on its own); then the library cards in no module, or no modules from a bundle built before it carried the tags; with `together` (at most 5 recipes, each a list of at most 6 groups of at most 12 library card names, a name resolved as `vtes_cards` does, a crypt card refused), each recipe's exact chance of an opening hand holding a distinct copy of each group and of having seen them within 14, 21 and 28 cards (a card named in two groups needing two copies; groups sharing too many cards to compute refused, said), the pieces not in the deck named and counted for nothing, and the crypt cards having what a card the deck holds of each group requires, with their levels in those cards' disciplines (the requirement checked, not the level an effect needs), the opening hand said to be the only random draw of fixed size; its combat posture (none from a bundle built before it carried the defenses): each combat defense of `modules.md`'s "Combat defenses" (by the tags' `defenses`), the most carried first, with how many TWDA decks of the archetype labels' window carry a card giving it and the three such cards most played, and the deck's cards answering it (or that none does), how, and what an answer requires (the level that effect needs, its clan and path) with how many crypt cards have it (unchecked without a crypt), a requirement only its text states quoted as unchecked; the costs; but for a draft, the list beside its reference decks ("Grading a deck" below): the nearest archetype with its variants, when the list is at or above its relatives' bar (below) and the archetype holds 10 decks or more, else its nearest decks at or above that bar when 10 are, else all the window's decks, said to be loose, a TWDA deck left out; the library's size; the cards needing a minion's action (action, political action, equipment, ally, retainer: a card counted once, in the first of these types it has), by type, the master cards with trifles (a master whose opening line says trifle), the blood and the pool the library costs (X left out), the copies of its blood and pool sources by kind (pool, blood, steal, blood to pool, pool to blood: the cards of `module-tags.json`'s `economy`, the list's named with their copies, a card counted in each of its kinds; copies, not yield, said; none from a bundle built before it carried them), the distinct library cards, the copies in its ten largest playsets and in no module shown, each as a count and per 100 library cards beside the reference decks' min, 10%, 25%, median, 75%, 90% and max, with how many are above and below it; how many reference decks, each set against the others, fall outside their range on one of the four budgets (action cards, masters, blood, pool); the reference's staples (library cards in three quarters of its decks) the list lacks or plays below the reference's lower quartile of copies, and the list's library cards a tenth or fewer of the reference plays, each a count beside the same count for each reference deck against the others, the cards named (cut first with the relatives' details); its track record among the TWDA decks of the archetype labels' window (a TWDA deck read by its id left out of the counts): but for a draft, the share of its library copies in cards fewer than 1% of those decks play, beside the same share for each reference deck against the others; the cards legal for under 18 months before the window's newest deck, listed apart and counted in neither part of that share; and, for a list or a draft, each card none of those decks plays (a vampire's groups merged), with the date it became legal and the years since, how many decks of the whole TWDA play it, and the last one's date and id (the record cut first with the relatives' details, the names past 15 with it); its relatives among the TWDA decks of the archetype labels' window (`knowledge/classification.json`, codex-of-the-damned's labels, curated and revisable), a TWDA deck left out of its own: a labelled deck's label; the three nearest archetypes, each with its deck count and representative decks, a variant with the archetype it varies, or that the deck is near none (below the similarity 95 in 100 labelled decks, each left out, reach to their own archetype); the three nearest decks with their event, its link and their label, and the cards that differ; how often the labelled decks, each left out, are placed back in their own; or no relatives, said, for a deck none of whose cards two decks of the window play, or from a bundle built before it carried the labels; and the procedures of `knowledge/guide/deck-analysis.md` and `deck-advice.md` as the answer contract, with the rulings archive to search before a point is called unsettled. Similarity is codex-of-the-damned's (sublinear TF-IDF over card counts, a vampire's groups merged, crypt and library weighed half each, cosine), with the noise decks kept in the window that weighs the cards, and without its SVD, which placed the labels back less well (95% against 97%). |
| `vtes_meta` | The TWDA decks of a window (the 12 months to the newest TWDA deck unless `months` gives another length, its previous window wholly within the archetype labels' window), each placed in its archetype: its label (`knowledge/classification.json`), or for a deck the labels lack (every deck since their pass, and some before it), its nearest archetype at `vtes_deck`'s bar (the similarity 95 in 100 labelled decks reach to their own), else none. Each deck weighs its event's players (a deck with no player count nothing, said), a variant counted in its root. The window and its previous one with their decks and players, said to be winners (what wins, not what is played); the players' share whose winner was labelled, placed by similarity, or in no archetype (the owner's noise, and near none), stated; the archetypes by share (15, 60 in detail; the rest summed) with their decks, those placed by similarity, the previous share and the change, and the variants' decks; the archetypes rising and falling by a point or more; the cards behind the move (the share of players whose winner played the card, against the previous window, for the cards at 5% or more in either, vampires named so); the leading archetypes' cores (their decks since the labels' window opened, labelled or placed: the vampires in a quarter of them or more, the library cards in half or more). With an archetype (by name or ref, as `vtes_card_stats` resolves it, its variants' decks with a root's; an unknown name pointed at the overview, which names the archetypes with their refs): its share and rank, its share over every whole window of that length since the labels' window opened, its core with each card's average copies, the labels' representative decks and its newest decks in the window, each with its event report's link and, when not the archetype asked, the archetype its label or its placement gives; and `knowledge/guide/meta-reading.md` as the answer contract. Later, with Archon login: play share against win share over the full field. |
| `vtes_card_stats` | One card (resolved as `vtes_cards` resolves it, a vampire counted in every group of its name) among the TWDA decks of the archetype labels' window, or within an archetype (by name or ref, its variants' decks with a root's; a name several archetypes share refused with their refs, an unknown name refused with the five archetypes playing the card most and, apart, up to three whose name holds a word of it, or that none does): its text, its requirements as `vtes_cards` gives them, its link and its ban; how many decks play it and the copies they play; overall, the archetypes playing it most, variants counted in their root, with each one's size; the library cards in a third or more of its decks, those most particular to them first (how many times their rate in the scope), and the vampires most often in their crypts (for a card one deck plays, that deck named instead); for a library card, its competitors: the cards sharing a type with it and its timing (its "Only usable ..." clause, one opening the other, or both without one, the type alone then, said to be broad), counted in the archetype's decks or, overall, in the decks nearest the sum of those playing it (twice as many, 20 at least, by `vtes_deck`'s similarity; not all those playing it need be among them), those played without it first, each with its text and its ban; below three such cards, those of its type with another timing after them; and `deck-advice.md` whole as the answer contract, with the rulings archive to search before a point is called unsettled. A count is what winners play, said so, not strength. |
| `vtes_events` | With Archon login only: tournaments, standings, rankings, ratings. |
| `vtes_engine_check` | A described position and a proposed play, answered by the engine at `engine.krcg.org` with the legal options; says plainly when a card is outside the engine's pool. |
| `vtes_guide` | The procedure for a task (rules answer, deck analysis, advice, card comparison, meta reading), the theory of the game, the module vocabulary, the game terms in a language with their label (translations unchecked, not an authority), the served examples of `knowledge/guide/examples/` (principle 8). |

Every tool that excerpts has a concise default and a detailed mode; every
tool says how to narrow when it truncates and dates the data behind its
answer. `primogen query` calls
the same tools from a shell. A tool that ships adds its one line to the README's
"Tools".

### Grading a deck

Decided with the owner, 2026-11-19, from a prototype over the TWDA window
(1494 decks) and the preconstructed decks the card data's set information
lists (45 from 2000 to 2010, "old"; 22 from 2019 to 2025, "recent").
`vtes_deck` grades no deck: it sets a list beside a reference (the row
above) and the model interprets (principles 7 and 10). Within an
archetype a single outlier is common (in the prototype, about one labelled
winner in five on the four budgets; over the whole window, far rarer), so
the result says how common. The budgets read a list against its relatives
(the Banu list holds 26 action cards, above all 41 Haqim Royalty winners)
but do not grade a level: in the prototype, the old precons' medians sit
inside the winners' central range on all but the masters. Card choice and concentration separate weak
lists (the Banu list's 35 distinct cards and 13 cards few of its
archetype play, against a median of 4, matching the thread's
"unfocused").

Computed since, each figure from the bundle of 2026-11-19:

- **Odds** (the row above): on the Banu list, Hunger of Marduk with an
  additional strike (Blur, Lightning Reflexes, Quickness, Side Strike,
  Pursuit) opens 22.0% of games and is seen within 28 cards in 91.3%; in
  TWDA deck 12494, a winner of its archetype, 36.9% and 97.6%. On the Banu
  list, 11 of 12 crypt cards have what both require (Arjun Shah has no
  Celerity), and the levels shown say the rest: Warmaksan's inferior
  Celerity gets an additional strike from its one Quickness only, Side
  Strike and Pursuit giving one at superior.
- **Rarely played cards** and **track record** (the row above): the
  Keepers of Tradition Ventrue precon holds 43% of its library in rarely
  played cards, against 27% at the 90th percentile of the window, its
  reference, and 4% at the median; the Fifth Edition Malkavian precon 0%, within
  Malk 22's range; the Sword of Nuln idea's engine, legal since 2003, is in
  none of the 4709 TWDA decks. Its real need, an answer to Strike: combat
  ends, is the combat posture's.
- **Blood and pool sources** (the row above): the Banu list's 7 steals are
  its 7 Hunger of Marduk, against a median of 12 in Haqim Royalty, 39 of
  whose 41 decks hold more; only one of those decks holds more than its 5
  blood sources (Fourth Tradition: The Accounting, Taste of Vitae). Which
  cards count, and as what, the owner settled where a text left it open
  (2026-11-20): Vessel and Blood Doll both ways, Camarilla's Iron Fist and
  Anathema pool sources, Week of Nightmares, Restricted Vitae, Open War
  and Wall Street Night excluded.

Rejected: bloat per turn and net pool (most gains are conditional, so a
per-turn yield is a judgement the server would assert), actions per
minion-turn (no turn count a list fixes), module co-occurrence as synergy
(it measures habit, and missed the shared blood budget the owner named),
and any level score (no measure correlates with event size beyond about
0.1). Synergy is read through its constraints: who can play every piece,
the odds the pieces meet, the shared budgets and the concentration. No
measure sees a card's quality against its competitors (`vtes_card_stats`'
slot survey does), nor whether a plan's refills or politics hold at a
table.

## Connecting

- **One endpoint**, `mcp.krcg.org/mcp`. It offers the Archon login when the
  user adds the server and never requires it. Without login every public tool
  works, and the user is told once that tournaments, full-field meta and
  rankings need the login.
- **MCP only.** The method reaches the model through the server's
  instructions (the grounding rules first), through each tool's description
  and every tool result (the contract and the next step), and through MCP
  prompts where the client offers them. Nothing is shipped per client beyond
  its line in the README.
- **The README** is compact and direct: one line per tool, per source and
  per client (where to add the server, or the command), the detail left to
  each client's documentation and to these pages.
- **Local installs** run the same binary over stdio. They download the
  published data bundle and build nothing, and the rulings archive's bulk
  file the first time a tool needs it: `vtes_rulings_archive`, or
  `vtes_rules_case` given a single card or two cards no ruling names
  together. They have
  every tool except the Archon-backed ones.
- **`mcp.krcg.org`** is the front door of the whole KRCG stack for
  developers and their agents: its root and `/llms.txt` serve one catalog
  in the [llms.txt](https://llmstxt.org) format, `knowledge/llms.txt`,
  authored here: each public service of the stack and its source
  repository in a line or two, linking its own live specification, which
  settles the detail. Every link is checked at each
  refresh; one that does not answer is a warning on the status page.

### What each client does with a remote server

As of 2026-11-11. Claude Code, a Claude connector (from Claude Code) and
the Gemini app were tried by hand; the rest is read from each vendor's
documentation and issue tracker, linked in each cell.

| Client | Adding the server | Instructions | Prompts | Results and tools | Checked |
|---|---|---|---|---|---|
| Claude (web, desktop, mobile) | A custom connector, "No sign-in" accepted, one on the Free plan ([docs](https://claude.com/docs/connectors/custom/add-unlisted.md)). Since 2026-11-03 authless connectors fail to connect ([claude-ai-mcp#1214](https://github.com/anthropics/claude-ai-mcp/issues/1214), open) | Not passed to the model ([#93](https://github.com/anthropics/claude-ai-mcp/issues/93), open) | Documented, reported missing from the menu since 2026-05 ([#333](https://github.com/anthropics/claude-ai-mcp/issues/333)) | About 150,000 characters at most ([docs](https://claude.com/docs/connectors/building/index.md)); the tools deferred, listed to the model by name only ([#1105](https://github.com/anthropics/claude-ai-mcp/issues/1105)); the list refreshed hours after a change ([#1044](https://github.com/anthropics/claude-ai-mcp/issues/1044)) | By hand, 2026-11-11, the connector added on claude.ai and used from Claude Code: it connected without sign-in despite #1214, the server's instructions reached the model, and `vtes_cards` returned Freak Drive with every link. The web client's own chat was not tried; the rest is docs and issues |
| ChatGPT | Plugins, "Add custom MCP server", no authentication accepted ([docs](https://developers.openai.com/api/docs/guides/custom-mcp-server)); which plans qualify is unconfirmed | Read since 2026-05-26 ([changelog](https://developers.openai.com/plugins/changelog)); the first 512 characters must stand alone ([docs](https://developers.openai.com/apps-sdk/build/mcp-server)) | Not documented | The model reads `content` and `structuredContent` ([reference](https://developers.openai.com/apps-sdk/reference)); a tool not marked read-only asks for confirmation ([docs](https://developers.openai.com/api/docs/guides/custom-mcp-server)); a change needs a manual Refresh and a new conversation ([docs](https://developers.openai.com/apps-sdk/deploy/connect-chatgpt)) | Docs |
| Gemini app | Settings, Connected apps, Custom apps: in the US, 18 or over, a personal account, Keep Activity on, added from the web ([help](https://support.google.com/gemini/answer/17209137)); a server without sign-in accepted (by hand), invoked as `@Primogen` | Added to the system prompt ([docs](https://developers.google.com/connected-apps/mcp/server-optimization)) | Not documented | `structuredContent` not supported; a tool marked destructive or open-world asks for confirmation; the schema and instructions snapshotted at registration (same page) | By hand, 2026-11-11, before `vtes_cards` carried the grounding rule: asked through `@Primogen` whether Freak Drive can follow The Kiss of Ra, it answered right from memory, without calling the tool (the server's log holds only the registration: `initialize` and `tools/list`), cited no link and left out the two rulings that settle it. Asked for Rötschreck's text and rulings, it called `vtes_cards` once (the response's size in the log matches), quoted the text exactly, paraphrased the eleven rulings without a link, and in two of them turned "the opponent" into "the target vampire" |
| Claude Code (local) | `claude mcp add`, HTTP or stdio ([docs](https://code.claude.com/docs/en/mcp)) | Passed, cut at 4,096 characters | Slash commands | 25,000 tokens at most by default, a text over 50,000 characters saved to a file; the tools deferred | By hand, hosted and stdio: the instructions quoted back, `vtes_cards` found and called unprompted |

Primogen is built for the weakest of them. Claude's connectors drop the
instructions and defer the tools, so a model may first meet primogen through
a tool's name alone: the grounding rule opens each tool's description and is
repeated at the end of every result. Results are text only. The rules
tools bound theirs under 40,000 characters, below the 50,000 past which
Claude Code saves a result to a file; a detailed one, under 80,000, is saved
there and read from it. What a bound leaves out is named, with the call
that gives it. Every tool
is marked read-only and closed-world (`vtes_deck` reading a deck stored on a
deckbuilder is the one call that leaves the server), the marks by which ChatGPT, Gemini and
Claude ([criteria](https://claude.com/docs/connectors/building/review-criteria.md))
decide whether to ask the user before a call. Input
schemas carry no type union: the Go SDK infers `["null", "array"]` for a
list, which the Gemini API answers with a 400
([prometheus-mcp#189](https://github.com/prometheus/prometheus-mcp/pull/189),
[gemini-cli#21094](https://github.com/google-gemini/gemini-cli/issues/21094)).

## What it deliberately does not do

- No LLM, no embeddings and no vector store on the server
  ([principles.md](principles.md) "Design").
- No Archon data for a user who is not a logged-in VEKN member, and none in
  the published bundle. One exception: community links are public links to
  public sites, served to everyone, without the member's name when the user
  is not logged in.
- No external blogs or fan-made material; community pointers are Archon's
  community links.
- No reading of zero-vtes's `rules-qa` as a source: its content is merged
  into the advanced rules upstream, which are the one reference.
- No in-flight proposals from the rulings site: only approved rulings.
- No bot-based deck testing yet. It waits on the engine covering the card
  pool and on a stronger bot; `vtes_deck` is where its results will go.

## Evaluation

- **Evidence tests**, with no LLM: for each graded case, the pack
  `vtes_rules_case` returns for the question must contain the rule or
  ruling that settles it. The cases: the owner-graded rules entries of the
  calibration corpus, by their `settled-by`; the rules facts of the codex
  strategist's cases C and F; each item of zero-vtes's `rules-qa` the
  advanced rules cite, asked as the question, against the sections citing
  it; and each row of zero-vtes's rulings ledger the pinned rulings database
  files the same way: a ruling filed under the card or its group among the
  card's rulings (`vtes_cards`), one filed under another card naming it
  among the rulings naming it (`vtes_rules_case`'s detail). The
  evidence keyword search does not reach is listed as a known miss, with
  why; a miss not listed fails, and so does a listed one found again. The
  three listed first are kept (owner, 2026-11-12): one judges' guide
  section no word of its question shares, and two multi-point `rules-qa`
  items cited for one point; embeddings are weighed only if graded sessions
  show retrieval missing. Since the graded session of 2026-11-12: two
  judges' guide sections a judge's question shares no word with (Game Rule
  Violation, Rolling Back the Game State), the retrieval miss that decision
  waits on.
- **Owner-graded sessions** on Claude, per milestone, with a regression set
  re-run after each revision (`calibration/`).
- **Other models** are spot-checked by hand, not graded.
