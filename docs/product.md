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

`vtes_cards` and the four rules tools have shipped and keep their names; the
others are provisional until they ship.

| Tool | What it returns |
|---|---|
| `vtes_cards` | Cards resolved from names (exact, an English name before another card's translated one; translated; or one or two letters off a single card): full text, requirements and costs, legality, every ruling with its `rulings.krcg.org` link and each cited post's link in the rulings archive, the card's entries in the history of changes (an entry whose name several cards share, or an older spelling, says so). An unknown name returns the closest names and no card. The detailed mode adds the printings and the translated texts. |
| `vtes_rules_case` | For a question and its cards, given, or else found in the question (an English name that is no game term (the game terms of `knowledge/`, titles and traits among them, like the trait Seraph, a clan, a title printed on a vampire, or a player's shorthand, like NRA) and no phrase of the rulebook, else the words opening or ending one card's name and no other, e.g. "Magnum" for .44 Magnum, weighed by their words that are no stopword, a single one capitalized, no game term, and no phrase the rulebook uses outside that name, nor such a word when they are a single one ("the Death of my vampire"); a name in the possessive read without its 's; only suggested: a single word opening a sentence or beside another capitalized one, a capitalized phrase of the rulebook, words opening or ending a few names, words of a name none capitalized, or another card when some are given): each card's verbatim text and requirements, a breakdown computed from the text (type, effects by level, timing and limits), every ruling with its link (the closest to the question only, when all would not fit, and how many more `vtes_cards` gives), the original post of the rulings closest to the question with its author's role then; the rulings on other cards naming two of its cards; the sections joined to the cards (naming them, citing the posts their rulings cite, or linked by them) and the sections closest to the question's words, each with the footnotes it calls; what no source covers (no ruling for a card, no advanced rules section naming it, no ruling naming two of the cards together); and the procedure of `knowledge/guide/rules-answer.md` as the answer contract. |
| `vtes_rules_search`, `vtes_rules_read` | Keyword search over the rulebook, the advanced rules, the tournament rules and the judges' guide by section, each labelled with its authority; a section read whole by id, with the sections under it and its footnotes. |
| `vtes_rulings_archive` | Full-text search of the rulings archive over what each author wrote (not what they quoted): posts with author, date, whether the author was the sitting rules director then ([sources.md](sources.md) "The rules directors"), and the citable link. |
| `vtes_deck` | A decklist (text, TWDA id, VDB link) parsed and validated, then a dossier: crypt table, which vampires can play each library card and at which level, library by type and module, hand expectation per module, costs, nearest archetypes, nearest tournament decks with their authors' comments, differences from those relatives. |
| `vtes_meta` | Archetype shares over a window weighted by event size, what rises and falls against the previous window, an archetype's core cards and representative decks. With Archon login: play share against win share over the full field. |
| `vtes_card_stats` | Play rates, co-occurrence, competitors for the same slot, overall or within an archetype. |
| `vtes_events` | With Archon login only: tournaments, standings, rankings, ratings. |
| `vtes_engine_check` | A described position and a proposed play, answered by the engine at `engine.krcg.org` with the legal options; says plainly when a card is outside the engine's pool. |
| `vtes_guide` | The procedure for a task (rules answer, deck analysis, advice, card comparison, meta reading), the theory of the game, the module vocabulary, the game terms in a language with their label (translations unchecked, not an authority). |

Every tool that excerpts has a concise default and a detailed mode; every
tool says how to narrow when it truncates and dates the data behind its
answer. `primogen query` calls
the same tools from a shell. A tool that ships adds its one line to the README's
"Tools".

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
  file the first time `vtes_rulings_archive` is called. They have
  every tool except the Archon-backed ones.
- **`mcp.krcg.org/llms.txt`** is the front door of the whole KRCG stack for
  developers and their agents: a catalog of the public services, their
  specifications, data formats and recipes, generated from the live
  specifications.

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
is marked read-only and closed-world, the marks by which ChatGPT, Gemini and
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
  show retrieval missing. Since the graded session of 2026-11-12: a rules
  director's posts no ruling cites, waiting on `vtes_rules_case` reaching
  the archive (the board), and two judges' guide sections a judge's
  question shares no word with (Game Rule Violation, Rolling Back the Game
  State), the retrieval miss that decision waits on.
- **Owner-graded sessions** on Claude, per milestone, with a regression set
  re-run after each revision (`calibration/`).
- **Other models** are spot-checked by hand, not graded.
