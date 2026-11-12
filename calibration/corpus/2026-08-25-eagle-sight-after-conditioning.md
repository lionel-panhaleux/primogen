---
date: 2026-08-25
mode: rules
answered-by: vtes-skill eb6dcd4
model: opus
status: graded
regression: yes
settled-by: rulebook, Stealth and Intercept Resolution; rulebook, Who May Attempt to Block; card text of Eagle's Sight, Conditioning; owner's grade (the table's repair: the Conditioning back to hand if possible)
---

## Question

I bleed my prey. They announce no block, I play Conditioning, my predator plays Eagle Sight to
block me. Is this allowed ?

## Answer given

(Answered by an Opus subagent given only the question and the instruction to use the vtes
skill.)

**Yes — that block is legal.** (With one caveat, below.)

### Why

**1. The block window was still open when your predator acted.**

The 2024 rulebook's *Detailed Course of an Action* (`references/rules/rulebook.md`) puts the
action in state **A. There is no current block attempt** until *every* Methuselah has passed:

> A.3 — *if a Methuselah passes, that Methuselah cannot declare any block attempt until the end
> of the action unless the target of the action changes.*
> A.4 — *once every Methuselah has passed, switch to "C. Blocks have been declined by all
> Methuselahs".*

Your prey's "no block" binds **your prey only**. Your predator hadn't passed — and in fact
couldn't have acted earlier: a Methuselah who is not the target of the action must wait until
the targeted Methuselah declares whether he blocks or declines
(`references/rules/advanced-rules.md` §3.3.2, per the rulebook Sequencing sidebar). So the
prey's declaration is precisely the moment your predator's window opens.

**2. Playing Conditioning doesn't close it.**

Conditioning (KRCG id 100401) reads, verbatim: *"Only usable during a bleed action. [dom] +2
bleed (limited). [DOM] +3 bleed (limited)."* No stealth, no "cannot be blocked." Playing an
action modifier right after the prey declines is explicitly blessed by the rulebook's own
worked example (Sarah plays Bonding after Alexis declines to block), and by the Sequencing
sidebar: *"if at any point any Methuselah uses a card or effect, the acting Methuselah again
gets the impulse back"* — the impulse goes round again, it does not shut the window.

**3. Eagle's Sight [AUS] lifts exactly the restriction that would otherwise bar your
predator.**

A bleed is a *directed* action, and under *Who May Attempt to Block*: *"only the Methuselahs
who are targeted may use their ready unlocked minions to attempt to block."* Normally your
predator simply cannot block your bleed. Eagle's Sight (id 100598) verbatim: *"[aus] +1
intercept. [AUS] This vampire attempts to block the current action, ignoring the normal prey,
predator or target restrictions for blocking actions."* That superior effect is the whole
point of the card.

### The caveat

If your predator had **already explicitly passed** before you played Conditioning (i.e.
everyone had passed and the action was in state C), then no — rulebook A.3 bars them for the
rest of the action, and the Eagle's Sight ruling is explicit that it does not save them:

> "[AUS] Only affects the rule restricting who gets to block an action. The vampire must still
> meet all other requirements to block. In particular, Blood Bond, Day Operation, Seduction and
> **prior decision not to block** are not circumvented." — [RTR 19950413] [RTR 20020501]

Once passed, only a **change of target** reopens blocks (rulebook C.2; e.g. a Deflection
bouncing the bleed). Conditioning is not a target change.

In the normal table flow you describe — prey says "no block", you immediately play
Conditioning, predator responds — the predator has not passed, and the block stands.

### Checks that must also hold

- **Superior Auspex is required.** The [aus] inferior gives only +1 intercept and does *not*
  let the predator block. Note also that the superior effect does **not** carry the +1
  intercept — it's a separate effect line.
- **Intercept vs. stealth.** A bleed has 0 inherent stealth and Conditioning adds none, so with
  no other stealth the attempt succeeds on 0 intercept. If you had added stealth, they'd need
  separate intercept cards.
- **The blocking vampire must be ready and unlocked** (or have used a wake effect that permits
  a block attempt).
- **One attempt only.** Eagle's Sight is spent by that single block attempt; if the action
  later continues as if unblocked (Form of Mist and friends), another copy is needed. —
  [LSJ 20030227]

### Practical note

If the block lands, your bleed is blocked, the Conditioning is wasted, and your acting vampire
is in combat with a fresh predator minion. This is exactly why Conditioning-after-"no block" is
a known trap into an Auspex predator: the safe line is to commit your bleed modifiers *before*
asking, or to hold them when a predator with superior Auspex is untapped.

**Authorities relied on:** `references/rules/rulebook.md` — *Who May Attempt to Block*,
*Detailed Course of an Action* (states A / C), *Sequencing* sidebar; card text for Conditioning
(100401) and Eagle's Sight (100598) via the KRCG v4 snapshot; Eagle's Sight rulings
[RTR 19950413] [RTR 20020501] and [LSJ 20030227]; `references/rules/advanced-rules.md` §3.3.1,
§3.3.2, §3.7.1.3 as locators.

## Owner grade

Ok this was a tricky one, you're right at the beginning but miss an important fact here. In
terms of sequencing the prey foregoing the block is a pass, so the impulse indeed goes around:
the window to play Eagle Sight is _before_ the acting methuselah gets his opportunity to play
conditioning.

So they played too fast and should have checked whether there was any Eagle Sight, but their
conditioning shouldn't be wasted either on the predator's block. If possible it should be
rolled back to hand.

## Correction

Two-part miss:

1. **Impulse order.** The prey's "no block" is a pass, which hands the impulse onward around
   the table — the predator's window to play Eagle's Sight comes *before* the impulse returns
   to the acting Methuselah for Conditioning. (Rulebook: "Once all Methuselahs have declined to
   block, the impulse goes back to the acting Methuselah (then the other Methuselahs following
   the order previously given) to play more cards and effects before resolving the action.")
   The answer treated "prey declines → I play Conditioning → predator responds" as clean
   sequencing when it is in fact out-of-order play.
2. **The remedy.** Both sides played too fast: the acting Methuselah should have checked for an
   Eagle's Sight before Conditioning; the predator's late block stands, but the Conditioning
   should not be wasted on it — if possible it is rolled back to hand.

## Generalization

Folded into `LESSONS.md` § Rules — sequencing (owner-approved 2026-08-26): walk the impulse,
and answer the table repair too.
