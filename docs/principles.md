# Principles

Chosen with the owner, 2026-11-10. Intake checks new work against this page;
review checks every change against it. Overturning a principle is valid work;
violating one silently is not.

## Design

1. **The server does the grounding.** One call returns the complete, verbatim
   evidence for a question; the model is never asked to remember a card or a
   ruling. Why: retrieval tools that leave application to the model still
   hallucinate on 17 to 33% of queries
   ([Magesh et al.](https://arxiv.org/abs/2405.20362)).
2. **The method lives in the tools' shape, not in prose.** Why: a prompt-only
   instruction to read evidence moved accuracy by about zero, an environment
   that required the read gained 3 to 9 points
   ([arXiv 2608.02011](https://arxiv.org/abs/2608.02011)).
3. **Structured lookup and keyword search, no embeddings.** Rulings are keyed
   by card, so finding them is a join; rules text is searched by section. Why:
   exact names and terms are what keyword search is best at, and the corpus
   is small ([Anthropic, Contextual Retrieval](https://www.anthropic.com/news/contextual-retrieval)).
   Embeddings are added only if the evidence tests show misses.
4. **Applicability is computed.** What a ruling applies to, what a card said
   at a date, what is banned: the server filters, the model reads.
5. **Authority is labelled on every item** (below).
6. **Absence is stated.** "No ruling names these two cards together" is a
   line in the result, so a model can say a question is unsettled.
7. **Computed facts before prose.** For a deck, the server computes what can
   be computed; the model interprets.
8. **Worked examples over rule lists.** A served example is an excerpt of
   one of primogen's own answers the owner graded clean, dated with its data,
   and re-checked by every regression re-run. Older graded material is
   evaluation data, never served: it ages, and a model anchors on it.
9. **Few task-shaped tools, bounded payloads, actionable errors, stable ids
   and links** ([Anthropic, Writing tools for agents](https://www.anthropic.com/engineering/writing-tools-for-agents)).
10. **Advice lays out options.** The competitors for a slot with the evidence
    and the trade-off, the player's plan asked for when unclear, never a bare
    endorsement of the card the player named. Why: answers that lean toward
    the user's view increase over-reliance
    ([Cheng, Ibrahim et al.](https://arxiv.org/abs/2609.14849)).
11. **User ease over implementation ease.** One endpoint, one binary, no
    tokens for local use; the complexity is ours.

## Authority of rules sources

- **Official**: the rulebook; the rulings database (curated by the
  rulemongers, each ruling citing a rules director's post); card text; the
  V:EKN's tournament rules and judges' guide.
- **Curated, revisable, to be taken with a grain of salt**: the advanced
  rules, a synthesis by the owner and other rulemongers that no rules
  director has verified. Within them, a reading no official source settles is
  marked as a reading. The rules director is not asked to confirm them.
- **Archive posts**: by a sitting rules director, or by a player.
- **The history of card and rules changes**: the owner's dated digest of the
  official announcements. Each entry links its announcement, which settles
  it.

## Knowledge files

- `knowledge/` is product: it is what users are served. A change to it is
  reviewed like code, and every card or rules claim in it is checked against
  the sources, never written from memory. One exception: the game terms'
  translations, kept unchecked as a consistency reference and labelled so in
  the file (owner, 2026-11-11: a fixed reference beats guessing a term in
  every answer).
- **Synced files are never hand-edited.** The advanced rules are authored in
  `vtes-advanced-rules`, the history of card and rules changes in the
  owner's gist: fix upstream, then `/sync`. The rulebook, the tournament
  rules and the judges' guide are read live from `archon-vibe`, like the
  card data.
  [provenance.md](provenance.md) says which file is which.
- **Authored here**: the guide (the theory, the procedure for each task,
  the module vocabulary and card-to-module tags, the served examples), the
  archetype labels, the game terms and their translations.

## Code

- **Go**, for the server and all data tooling. One static binary.
- **Locality first.** Keep together what changes together; files short enough
  to read in one pass. Extract a module only when its interface is much
  narrower than what it hides. Similar code that changes for unrelated
  reasons stays repeated.
- **Simplicity means avoiding hazards**: dependence at a distance, behaviour
  invisible where it lives, traps for a reader without today's context.
- **One source of truth.** The docs may name parts of the code; the code
  never points back. A comment is justified only by a trap invisible where
  the code is read, two lines at most. No narration, no TODOs.
- **Look it up.** Before using a library, a protocol or a tool, read its
  current documentation ([operations.md](operations.md) "Looking things
  up"). Never from memory.

## Testing

- **Few tests, at the edges**: a tool call and its result, the command line,
  the refresh end to end, the failure modes that matter. Never internals.
- **Each test traces to a claim in the docs**; a test nobody can connect to
  one is removed. The evidence tests trace to
  [product.md](product.md) "Evaluation".
- **No mocks.** Tests run against a pinned copy of real data.
- **A test that caught a real bug is protected**: deleting it takes a stated
  reason.
- Tests written to drive development may be dropped before landing. Property
  tests only for parsing.
