---
date: 2026-08-25
mode: rules
answered-by: vtes-skill eb6dcd4
status: graded
regression: no
settled-by: rulebook, Who May Attempt to Block; advanced rules, Declaring and declining; ANK 20211003; card text of Deflection, Direct Intervention, Govern the Unaligned
---

## Question

I bleed with govern, prey plays Deflection, grand-prey DIs, can my prey block ?

## Answer given

**No — your prey cannot block anymore.** The bleed resolves against them (unless they have
other reactions left).

The chain, with authorities:

1. **Playing Deflection declined the block.** Deflection's own text: *"Only usable if a minion
   is bleeding you, **after blocks are declined**."* Playing a card so worded declines the block
   for its controller (advanced-rules §3.3: "Playing a card worded 'after blocks are declined'
   declines the block for its controller" — [ANK 20211003]). So the moment your prey played
   Deflection, they had passed on blocking.

2. **A pass is final unless the target changes.** Rulebook block-sequencing rules: *"if a
   Methuselah passes, that Methuselah cannot declare any block attempt until the end of the
   action **unless the target of the action changes**"* — and the declined state re-opens only
   *"if the target of the action is changed."*

3. **The canceled Deflection never changed the target.** Direct Intervention cancels the card
   *as it is played*: it has no effect and its cost is not paid. The target of your Govern bleed
   never moved off your prey — so no new block window ever opened for them. (Your prey stays the
   target; your grand-prey, who was never the target, has no block standing either.)

Two edges worth knowing in the same situation:

- **The same vampire cannot try again**: a canceled card *has still been played*, and "the same
  reaction or modifier cannot be played again by the same minion" (DI ruling, [ANK 20190104]
  [LSJ 19980212]). A **different** ready vampire of your prey could still play a second
  Deflection — reactions remain legal until resolution.
- Because the cost was never paid, the reacting vampire **doesn't lock** (relevant for inferior
  `dom` Deflection).

So: prey declined by playing Deflection → DI canceled it before it could change the target → no
target change, no new block opportunity → the +2 Govern bleed proceeds against the prey.

## Owner grade

Great answer. A bit too long but that's externally fixable.

## Generalization

Case-specific — none. The one note (length) is externally fixable per the owner: a
harness/verbosity concern, not skill knowledge. No substantive correction → `regression: no`
by the default rule.
