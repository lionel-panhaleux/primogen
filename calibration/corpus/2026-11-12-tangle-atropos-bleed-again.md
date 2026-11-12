---
date: 2026-11-12
mode: rules
answered-by: primogen 3d72039
status: graded
regression: no
settled-by: card text of Tangle Atropos' Hand; LSJ 20070709; ANK 20210124; rulebook, Who can bleed; advanced rules, Reached resolution
---

## Question

If I use Tangle Atropos on a bleed, can I bleed again ?

## Answer given

Yes. Tangle Atropos' Hand cancels the bleed and unlocks your vampire, and a canceled action hasn't reached resolution. The same vampire can bleed again this turn, with one exception: if the bleed had already been blocked and then continued (for example by superior Form of Mist) before you played Tangle, it can't.

**What settles it** (official, from the rulings database):
- "The action has not reached resolution. The same vampire can undertake the same action again, but not if it was blocked then continued before (eg. [PRO] Form of Mist). [LSJ 20070709] [LSJ 20090220] [ANK 20210124]", https://rulings.krcg.org/index.html?uid=101935#r-YHEMKHX6.
  - L. Scott Johnson, the sitting rules director, on 2007-07-09: the minion can act again "Only if the action that was canceled hadn't reached resolution" (https://usenet.krcg.org/t/qKrJPLXBFFw/#m2).
  - He also said a vampire isn't "untainted" for that purpose after a block and a continued action (2009-02-20, https://usenet.krcg.org/t/JMgZa9jdrc4/#m1).
- "The action was still taken (eg. counts against Enkil Cog or [OBF] Veil the Legions). [ANK 20210124]", https://rulings.krcg.org/index.html?uid=101935#r-MWBZCDOL. So the canceled bleed still counts for cards that count actions. It just doesn't stop a second bleed.

**The card text:** "Only usable if a minion attempts to block. [pot] Cancel the action and unlock this vampire. (The action is not blocked. Do not lock the blocking minion.) [tem] As [pot] above, and if an action card was played, take it back in your hand (discard down afterward). [TEM] As [tem] above, and this vampire gets +1 stealth during their next action this turn."

That fits the second bleed well: at [tem] you get the bleed card back, and at [TEM] the next bleed has +1 stealth.

**The rulebook's limit doesn't bar it.** "A minion cannot perform more than one bleed action each turn, even if they unlock" (official, https://archon.vekn.net/help/rules#who-can-bleed). The ruling above settles that a canceled bleed doesn't use it up. The advanced rules (curated, not verified by a rules director) agree: "A canceled action has not reached resolution. It is not performed, its cost is not paid, and the same minion may attempt the same action again. It may not if the action was blocked and then continued before the cancellation." They also say "No further action modifier or reaction can be played" once the action ends this way (https://vtes-biased.github.io/vtes-advanced-rules/#351-reached-resolution).

## Owner grade

Good answer on tangle.

## Generalization

case-specific: none. Called with the question alone, `vtes_rules_case` did not find the card
from "Tangle Atropos", the first words of its name; the answer called it again with the
card. Proposed to the owner through intake.
