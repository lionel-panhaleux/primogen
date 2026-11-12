---
date: 2026-11-12
mode: rules
answered-by: primogen 3d72039
status: graded
regression: no
settled-by: card text of Govern the Unaligned, Enhanced Senses; judges' guide, Procedural Error — Game Rule Violation; judges' guide, Rolling Back the Game State
---

## Question

Player A announces Gover Sup and their predator plays Enhanced Senses to block. Then the moment they start combat they realise the acting vampire did not have superior dominate so couldn’t have played Govern this way in the first place. As a judge for the table, what do I do ?

## Answer given

Call the head judge. It's a Game Rule Violation by player A, with a Warning rather than a Caution, because it can no longer be fixed with a simple roll-back. The head judge then decides between a full roll-back to before the Govern and a forward fix. This one has few, clean steps, so the roll-back is the natural choice.

**1. The infraction: Procedural Error — Game Rule Violation** (judges' guide, official, https://archon.vekn.net/help/judges-guide#23-procedural-error-game-rule-violation). Its examples include "A player attempts to perform an action that their minion is not allowed to perform." Superior Govern the Unaligned needs [DOM], and the vampire didn't have it.

**2. The penalty: a Warning.** The guide gives "**Caution**, if the violation can be fixed with a simple roll-back, or with a correct declaration of the played card's effect. **Warning**, if the violation cannot be fixed with a simple roll-back." A simple roll-back "reverses only the most recently completed game action" and "must not involve randomization of any kind" (1.4.1). Here, after the error, A drew Govern's replacement, the predator played Enhanced Senses, drew its replacement, and blocked. So it's past a simple roll-back.

**The "correct declaration" route doesn't fit,** even if the vampire has inferior Dominate. Inferior Govern is "Ⓓ Bleed with +2 bleed": a directed action against the prey, which the predator could not have blocked. Re-declaring it would make the predator's block, and the Enhanced Senses, illegal after the fact. That's my reading of the rulebook's blocking rules. No ruling names this case.

**3. The remedy: the least disruptive fix the guide's hierarchy allows** ("Simple Roll-back", then "Partial Fix", then "Structured Forward Fix").
- **A full roll-back needs the head judge:** "Only the Head Judge may approve a roll-back" (https://archon.vekn.net/help/judges-guide#141-rolling-back-the-game-state). It "should be attempted only in situations with very few subsequent choices or hidden effects", and here there are only three steps after the error. Reversed "in precise reverse order", it gives:
  - Combat is cancelled and the blocker unlocked.
  - Enhanced Senses goes back to the predator's hand.
  - The Govern goes back to A's hand, its 1 blood is restored, and the acting vampire is unlocked.
  - Each drawn replacement goes back to its library. The guide says that "When a card that must be returned to a hidden area is unknown to one or more players (e.g., a card drawn into a hand), the judge selects a random card from among all legal candidates in that hidden area."
  - A then takes a legal action.
- **If the head judge finds that too disruptive,** the guide says "apply a forward correction that removes any advantage gained: unlock minions improperly locked, restore costs, return resources… If the correction would be too disruptive, the judge may leave the game state as-is." At the least, the illegal superior effect must not happen: no +1 stealth and no 3 blood. How exactly the combat is handled is left to the judge. The guide doesn't settle this case.
- **The predator gets no Failure to Maintain Game State.** That penalty "applies only when the judge determines that the player actually noticed the error", and the predator spoke up as soon as they noticed (https://archon.vekn.net/help/judges-guide#24-procedural-error-failure-to-maintain-game-state).

**The cards:**
- **Govern the Unaligned:** "Requires [dom]. Cost: 1 blood. [dom] Ⓓ Bleed with +2 bleed. [DOM] +1 stealth action. Add 3 blood to a younger vampire in your uncontrolled region."
- **Enhanced Senses:** "[aus] +1 intercept. [AUS] +2 intercept."

**What's ruled and what's judgment:** the infraction and the Warning come straight from the guide. The choice between roll-back and forward fix is the head judge's call. No ruling and no section of the guide covers this exact case.

## Owner grade

Good answer, great.

Let’s stop here for now. For all answers I suppose they could be shorter but that’s for the llm and its user to figure out I suppose, no?

## Generalization

case-specific: none. The answer reached the judges' guide's Game Rule Violation and its
roll-back section through `vtes_rules_search`: the pack `vtes_rules_case` returned for the
question held Failure to Maintain Game State and Slow Play instead.
