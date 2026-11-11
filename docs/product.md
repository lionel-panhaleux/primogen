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

`vtes_cards` has shipped and keeps its name; the others are provisional until
they ship.

| Tool | What it returns |
|---|---|
| `vtes_cards` | Cards resolved from names (exact, an English name before another card's translated one; translated; or one or two letters off a single card): full text, requirements and costs, legality, every ruling with its `rulings.krcg.org` link and each cited post's link in the rulings archive, the card's entries in the history of changes (an entry whose name several cards share, or an older spelling, says so). An unknown name returns the closest names and no card. The detailed mode adds the printings and the translated texts. |
| `vtes_rules_case` | For the cards involved and a question: an evidence pack. Each card's verbatim text, every ruling that applies with the rules director's original post, the governing rulebook and advanced-rules sections, a structured breakdown of each card (type, cost, inferior and superior effects, timing restrictions), what is explicitly absent, and the answer contract. |
| `vtes_rules_search`, `vtes_rules_read` | Keyword search over rulebook, advanced rules, tournament rules and judges' guide by section; a section read whole by id. |
| `vtes_rulings_archive` | Full-text search of the rulings archive: posts with author, date, whether the author was the sitting rules director then, and the citable link. |
| `vtes_deck` | A decklist (text, TWDA id, VDB link) parsed and validated, then a dossier: crypt table, which vampires can play each library card and at which level, library by type and module, hand expectation per module, costs, nearest archetypes, nearest tournament decks with their authors' comments, differences from those relatives. |
| `vtes_meta` | Archetype shares over a window weighted by event size, what rises and falls against the previous window, an archetype's core cards and representative decks. With Archon login: play share against win share over the full field. |
| `vtes_card_stats` | Play rates, co-occurrence, competitors for the same slot, overall or within an archetype. |
| `vtes_events` | With Archon login only: tournaments, standings, rankings, ratings. |
| `vtes_engine_check` | A described position and a proposed play, answered by the engine at `engine.krcg.org` with the legal options; says plainly when a card is outside the engine's pool. |
| `vtes_guide` | The procedure for a task (rules answer, deck analysis, advice, card comparison, meta reading), the theory of the game, the module vocabulary, the game terms in a language with their label (translations unchecked, not an authority). |

Every tool has a concise default and a detailed mode, says how to narrow when
it truncates, and dates the data behind its answer. `primogen query` calls
the same tools from a shell.

## Connecting

- **One endpoint**, `mcp.krcg.org/mcp`. It offers the Archon login when the
  user adds the server and never requires it. Without login every public tool
  works, and the user is told once that tournaments, full-field meta and
  rankings need the login.
- **MCP only.** The method reaches the model through the server's
  instructions (the grounding rules first), through each tool's description
  and every tool result (the contract and the next step), and through MCP
  prompts where the client offers them. Nothing is shipped per client beyond
  a configuration snippet in the README.
- **Local installs** run the same binary over stdio. They download the
  published data bundle and build nothing. They have every tool except the
  Archon-backed ones.
- **`mcp.krcg.org/llms.txt`** is the front door of the whole KRCG stack for
  developers and their agents: a catalog of the public services, their
  specifications, data formats and recipes, generated from the live
  specifications.

### What each client does with a remote server

As of 2026-11-11. Claude Code and the Gemini app were tried by hand; the
rest is read from each vendor's documentation and issue tracker, linked in
each cell.

| Client | Adding the server | Instructions | Prompts | Results and tools | Checked |
|---|---|---|---|---|---|
| Claude (web, desktop, mobile) | A custom connector, "No sign-in" accepted, one on the Free plan ([docs](https://claude.com/docs/connectors/custom/add-unlisted.md)). Since 2026-11-03 authless connectors fail to connect ([claude-ai-mcp#1214](https://github.com/anthropics/claude-ai-mcp/issues/1214), open) | Not passed to the model ([#93](https://github.com/anthropics/claude-ai-mcp/issues/93), open) | Documented, reported missing from the menu since 2026-05 ([#333](https://github.com/anthropics/claude-ai-mcp/issues/333)) | About 150,000 characters at most ([docs](https://claude.com/docs/connectors/building/index.md)); the tools deferred, listed to the model by name only ([#1105](https://github.com/anthropics/claude-ai-mcp/issues/1105)); the list refreshed hours after a change ([#1044](https://github.com/anthropics/claude-ai-mcp/issues/1044)) | Docs and issues |
| ChatGPT | Plugins, "Add custom MCP server", no authentication accepted ([docs](https://developers.openai.com/api/docs/guides/custom-mcp-server)); which plans qualify is unconfirmed | Read since 2026-05-26 ([changelog](https://developers.openai.com/plugins/changelog)); the first 512 characters must stand alone ([docs](https://developers.openai.com/apps-sdk/build/mcp-server)) | Not documented | The model reads `content` and `structuredContent` ([reference](https://developers.openai.com/apps-sdk/reference)); a tool not marked read-only asks for confirmation ([docs](https://developers.openai.com/api/docs/guides/custom-mcp-server)); a change needs a manual Refresh and a new conversation ([docs](https://developers.openai.com/apps-sdk/deploy/connect-chatgpt)) | Docs |
| Gemini app | Settings, Connected apps, Custom apps: in the US, 18 or over, a personal account, Keep Activity on, added from the web ([help](https://support.google.com/gemini/answer/17209137)); a server without sign-in accepted (by hand), invoked as `@Primogen` | Added to the system prompt ([docs](https://developers.google.com/connected-apps/mcp/server-optimization)) | Not documented | `structuredContent` not supported; a tool marked destructive or open-world asks for confirmation; the schema and instructions snapshotted at registration (same page) | By hand, 2026-11-11, before `vtes_cards` carried the grounding rule: asked whether Freak Drive can follow The Kiss of Ra, it answered right and quoted both cards exactly, but cited no link, left out the two rulings that settle it, and added an unsourced timing rule. Whether it called the tool, and how a call looks, unchecked |
| Claude Code (local) | `claude mcp add`, HTTP or stdio ([docs](https://code.claude.com/docs/en/mcp)) | Passed, cut at 4,096 characters | Slash commands | 25,000 tokens at most by default, a text over 50,000 characters saved to a file; the tools deferred | By hand, hosted and stdio: the instructions quoted back, `vtes_cards` found and called unprompted |

Primogen is built for the weakest of them. Claude's connectors drop the
instructions and defer the tools, so a model may first meet primogen through
a tool's name alone: the grounding rule opens each tool's description and is
repeated after the cards of every result. Results are text only. Every tool
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

- **Evidence tests**, with no LLM: for each graded case, the pack returned
  for the question must contain the rule or ruling that settles it.
- **Owner-graded sessions** on Claude, per milestone, with a regression set
  re-run after each revision (`calibration/`).
- **Other models** are spot-checked by hand, not graded.
