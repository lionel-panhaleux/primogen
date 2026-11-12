---
layout: rules
title: VTES Advanced Rules
permalink: /
---

# VTES Advanced Rules

> **UNOFFICIAL DRAFT**
> This document has not been approved by the Rules Director. It's an unofficial community effort,
> **contributions are welcome**.

## About this document

This document supplements the [VTES rulebook](https://www.vekn.net/rulebook). It collects the general principles, corner
cases and non-obvious interactions established by thirty years of rulings, organized by game mechanic.

It is written for judges and advanced players: it assumes full fluency with the base rules and terminology, and favors
precision over accessibility.

A statement marked *(reading)* settles a point no official ruling settles: its footnote names where the reading is
recorded, and an official ruling would supersede it.

## Driving principles

Four principles drive most rulings:

**Card text governs.** Where a card contradicts the rules, the card takes precedence [[RBK
important-terms-of-the-game]](https://www.vekn.net/rulebook#important-terms-of-the-game). However odd the result,
presume the text applies naturally, as written, in the expected manner ([§1.1](#11-card-text-and-the-rules)).

**You cannot do what you cannot do.** A play barred in whole or in relevant part is not available at all, even just to
cycle the card: no adding a value you do not need ([§1.6.5](#165-playing-a-card-that-will-do-nothing)), no maneuver from
a strike you cannot use ([§4.10](#410-weapons-and-equipment-in-combat)), no announcement without a legal target
([§3.1.3](#313-legal-targets)), no play with a cost you cannot afford ([§1.7](#17-costs-and-payment)).

**There is no stack.** Effects are instantaneous and resolve completely before the next opportunity opens; nothing is
played "in response" ([§2.1.3](#213-effects-resolve-atomically)).

**Effects that apply, apply.** Unless it has become impossible, a scheduled effect applies: its source leaving play,
other effects landing before or after, or the action failing change nothing
([§2.5.3](#253-effects-outliving-their-source), [§3.5.3](#353-ended-and-failed)).

---

## 1. Cards

### 1.1 Card text and the rules

**Base rules.** Where a card contradicts the rules, the card takes precedence. [[RBK
important-terms-of-the-game]](https://www.vekn.net/rulebook#important-terms-of-the-game)

#### 1.1.1 Mandatory and optional effects

A printed statement of fact applies automatically and its controller has no choice in the matter, e.g. {Writ of
Acceptance}, {Enchanted Marionette}.[^1-1-1] Only a clause the card marks as optional — "may", "can", "optional" — is
optional. {Changeling Skin Mask} carries both: the bearer has [OBF] whether you want it or not, and burning the card for
intercept is your choice.

An ability is mandatory unless the card marks it optional. {Renegade Garou} gets a mandatory additional strike and an
optional maneuver on the same line; {Lorrie Dunsirn}'s press and additional strike are both mandatory.[^1-1-2]

A cost printed inside an effect is mandatory. {Vast Wealth} equips with the equipment found and you pay for it even if
the payment ousts you; {Serpent's Numbing Kiss} at superior burns 1 blood.[^1-1-3] A printed cost reduction likewise
cannot be declined, e.g. {Frondator}.

An opportunity is not a compulsion. A card that lets a minion strike with a weapon does not force the strike, e.g.
{Bomb}, {Bundi}; a wake card lets its vampire attempt a block but never requires it, e.g. {Wake with Evening's
Freshness}.[^1-1-4]

A choice inside an effect is made when the effect resolves, not when the card is played, e.g. the "up to 2 blood" on
{Third Tradition: Progeny}.[^1-1-5] Who takes or declines an optional effect on a card in play follows the card's
wording, not control of the card: an effect granted to "this vampire" is used by the vampire — hence by its current
controller — even on a master card another Methuselah played and still controls, e.g. {Perfectionist}, {Corporal
Reservoir}.

#### 1.1.2 Applying a mandatory effect

A mandatory effect is applied at its timing point, before its controller passes the opportunity to play effects, e.g.
burning pool to {Anarch Revolt}.[^1-1-6] A player who has passed that opportunity does not get it back because the
effect was overlooked: {The Coven} still moves to the predator and its controller cannot lock it retroactively.

A mandatory "during each X phase" effect on a card entering play during an X phase is due in that same phase, e.g.
{Brujah Debate} played in a master phase, {Dragonbound} put in play in a discard phase *(reading)*.[^1-1-11]

Where a mandatory choice has no printed default and every player missed it, imposing one of the choices after the fact
is not a remedy, e.g. {Leandro}.

#### 1.1.3 "Cannot"

A "cannot" on a card is absolute. Where it names the cards it forbids, those cards cannot be played at all, not even for
no effect to cycle them: the {Blood Fury} template ("damage from this strike cannot be prevented by cards requiring
Fortitude [for]") stops the opponent from playing [for] prevention entirely.[^1-1-7] "Cannot use any additional strikes"
({Rigor Mortis} and its template) is read the same way, not as a barred outcome: it bars playing an additional-strike
card even for a side benefit, e.g. {Acrobatics} for its dodge.

Where the "cannot" names an outcome instead, the play remains legal and only the outcome fails. {Charnas the Imp} is
immune to damage from its employer, but the employer may still use effects that would normally damage it.[^1-1-8] A
vampire under {Legacy of Caine} who cannot gain blood still takes the hunt action; the blood he steals goes to the blood
bank. Blood moved onto a vampire who cannot gain blood is a gain all the same, e.g. with {Heidelberg Castle, Germany}
onto an {Ex Nihilo} vampire.[^1-1-15]

#### 1.1.4 Reading card text

"This minion" is the acting minion, "that minion" the target, e.g. {Big Game}.[^1-1-10] Such referents are re-read as
the effect resolves: two {Major Boon} cannot be burned on the same bleed, because after the first "you" are no longer
the Methuselah burning pool. "You" and "your" are the card's controller, whoever brought the effect about: {Open War}'s
4 pool goes to its controller, not to the Methuselah who moved the fourth counter *(reading)*.[^1-1-12]

A higher level carries the lower level's text only where it prints "as above"; otherwise it stands alone.[^1-1-14] So
{Read the Winds} at [ANI][AUS] locks its vampire for the block *(reading)*.[^1-1-13]

### 1.2 Wording templates: periodicity, duration, triggers

#### 1.2.1 How often an effect may be used

"Once each turn" means once during each player's turn. "Once during each of your turns" is the narrower form, e.g.
{Maila}. Turn, master phase, unlock phase, action, bleed and referendum all read the same way, and so does "during X, …
can …" or "can … during X": one use per X with each copy, e.g. {Vessel} once per unlock phase, {Eyes of the Wild}'s burn
once in its action.[^1-2-1]

A limit keyed to an event is spent per occurrence, and the event can recur in one turn: {Andre LeRoux} may be used on
each successful bleed. An ability worded "after X" is limited by that event even with no other limit: once after each X
with each copy, e.g. {Richard de Worde} *(reading)*.[^1-2-17] An ability whose use costs a master phase action is
likewise spent per action, so extra master phase actions buy extra uses, e.g. {Nahir}. An ability framed by a phase is
spent for that phase: {Nonu Dis} may be used once per master phase however many master cards are played in it.[^1-2-2]

Absence of a periodicity clause means no limit — {Maris Streck} may use her ability as often as she can pay for it. A
trigger condition is not a limit: {Slake the Thirst} may be played several times on one blood gain.[^1-2-3]

The limit attaches to the card, not to the effect, so a second copy carries its own use whatever the period, e.g.
{Sawed-Off Shotgun}. See [§1.15](#115-cumulative-and-stacking-effects) for multiple copies generally.[^1-2-4]

The limit travels with the card, not the holder: {Owain Evans, The Wanderer}'s use is spent for that unlock phase even
if control of him changes during it; a weapon's per-combat allowance tracks the weapon; a spent referendum ability is
not restored by unlocking. A minion who takes an action over gets no fresh use of a per-action ability already used on
it.[^1-2-5] Per [§1.15](#115-cumulative-and-stacking-effects) a once-per-action limit on *playing* a card binds the
minion instead.

An effect marked "(limited)" is capped by the rulebook, not the card, and is not optional: once a minion has used a
limited additional strike this round, he cannot play the card for its other clause, e.g. {Dust Up} [cel].[^1-2-6] Each
cap binds what the rulebook names: a limited bleed increase binds the bleed action — a substitute acting minion cannot
use one if one has already been used on the action — while a limited additional strike binds the minion per round of
combat, each combatant separately.[^1-2-15] A canceled use spends no limit, nor does one granting nothing, e.g.
{Lightning Reflexes} [CEL] with X = 0 *(reading)*.[^1-2-18] A hunting ground used on a vampire at full capacity grants
him blood, so it is still his one hunting ground that turn, e.g. {Zoo Hunting Ground} *(reading)*.[^1-2-20]

A referendum ability is usable once per referendum; which states it works from is
[§3.7.6.4](#3764-locked-and-torpid-vampires)'s.[^1-2-7]

"Only once in a game" survives the card's return from the ash heap, so a replay effect cannot bring it back, e.g.
{Filchware's Pawn Shop}.[^1-2-9] "Only one X can be played each turn" or "… per game" names no player, so it counts
every Methuselah's plays, a canceled one included, e.g. {Ashur Tablets}, {Fortschritt Library}; each event card is
likewise played once a game across the table *(reading)*.[^1-2-19] "Since your last turn" means between the end of your
previous turn and now: the previous turn does not qualify, the current one does.

#### 1.2.2 Trigger and condition wording

"Reduced to X" names an event, not a state. {Anathema} does not fire on a vampire that entered combat already at zero
blood, fires on any loss in combat however caused, and does not fire on diablerie.[^1-2-12]

A condition referring to the card's own effect is not satisfied by another effect producing the same event: {Escaped
Mental Patient} is not burned when another card grants him the hand strike.[^1-2-13] An "either A or B" condition needs
one branch completed in full; partial progress on both does not combine, e.g. {First Tradition: The Masquerade}.

"When the action/bleed would be successful" is a timing window, not merely a condition
(see [§3.4.3](#343-the-would-be-successful-window) for sequencing).[^1-2-14] "Would X" with a stated substitute marks a
replacement effect; [§2.4.1](#241-ordering-within-a-window) sequences it.[^1-2-16]

### 1.3 Card types and multi-type cards

#### 1.3.1 Locquipments

An equipment card printed "represents a location and does not count as equipment while in play" is an equipment in hand,
library and ash heap; both equipment and location as it is played; and only a location once in play, e.g. {Palatial
Estate}.[^1-3-1]

Cost effects reach it under either wording, since the card is checked in hand and is both types at the moment it is
played. An effect reducing the cost of locations applies, e.g. {Therbold Realty}; so does one reducing the cost of
equipment cards, e.g. {The British Museum, London}.[^1-3-2] See [§1.7](#17-costs-and-payment) for cost reduction.

A search for an equipment card finds a locquipment in the library, e.g. {Magic of the Smith}; a search for a location
does not, e.g. {Danylo}. An effect that moves "the equipment" after an equip action cannot move it, e.g. {Reg
Driscoll}.[^1-3-3]

A vampire who cannot have or use equipment may still equip with a locquipment, e.g. {Enkidu, The Noah}. {Beast, The
Leatherface of Detroit} also cannot play action cards, so a card must put it on him, e.g. {Vast Wealth}.[^1-3-4]

#### 1.3.2 Multi-type cards

A card with two types is of the type it is played as, never both.[^1-3-6] {Adana de Sforza} reduces the cost of a
[COMBAT]/[MODIFIER] card only when it is played as a combat card; {Conrad Adoula} reads a [REACTION]/[MODIFIER] card
only when it is played as a reaction. A card offering a choice of disciplines is of the discipline it was played at,
e.g. {Horrock}. In hand, library or ash heap a card printing several Disciplines counts as requiring every one of them,
e.g. {Alexa Draper} may discard any card listing [dom] for her vote.[^1-3-8]

An effect that retrieves or counts cards by type reads the type the card was played as: {Henry Taylor} recovers
{Resist Earth's Grasp} only if it was played as a combat card.

The printed type is unchanged by an unusual play route. A retainer employed through {Pack Alpha} is still an action
card; see [§1.7.3](#173-cost-reductions) for cost reductions.[^1-3-7]

Whether a card keeps its type after being used is settled by its own text, not by a general rule. {Shackles of Enkidu}
is still an equipment once it has been moved onto the opposing minion. {Enhanced Coagulant} prints replacement text on
use and becomes typeless, so it is no longer a weapon.

### 1.4 Representation and placeholders

**Base rules.** Several cards put a card into play to stand in for something else — a Master: Discipline card as a
1-capacity Sabbat vampire ({Shock Troops}), a library card as a ghoul ally ({Absimiliard's Army}), an ash-heap minion as
a wraith ally ({Khazar's Diary (Endless Night)}).

The representing card contributes nothing of its own. Its name, card type, cost, clan, disciplines, text and unique
status are all inert (see [§1.13.1](#1131-what-contests-what) for contests); the card is exactly what the effect that
put it there describes, and nothing more.[^1-4-1] An {Agent of Power} put into play by {Shock Troops} is a non-unique
clanless 1-capacity Sabbat vampire with no discipline, and is not a Master: Discipline card while it is in play.

The placeholder sits in play face up; a card merely stored on another card is out of play and sits face down instead,
e.g. {Père Lachaise, France}.[^1-4-2]

Once the card leaves play it is read as printed again; what it represented is gone.

Standing in for something else does not launder a play restriction. An ability that lets a minion count as having a
required discipline substitutes only for that discipline; clan and other printed restrictions still apply, e.g. {Reality
Mirror}.[^1-4-5] See [§1.6](#16-requirements) for requirements and requirement-faking.

### 1.5 Abilities and card plays

**Base rules.** Using an ability printed on a card already in play is not playing a card, and is not taking an action.

#### 1.5.1 Ability use is not a card play

Costs and restrictions that attach to playing a card do not reach ability use. The [obe] version of {Safe Passage}
raises the cost of reaction cards, not of reaction powers. {Secure Haven}'s surcharge applies to master cards played at
the minion, not to using the effect of a master already in play, e.g. {KRCG News Radio}.[^1-5-1]

An ability marked [REACTION] (a marker only imbued powers bear) does not count as playing a reaction card, but it can be
used only when reactions can be used — unlocked, or with a wake effect, e.g. {Champion}. Card text overrides the
unlocked half: {Vigilance} is only usable by a locked imbued, so it needs no wake and only the reaction window itself —
during an action — applies.[^1-5-2] An effect that merely lets a locked minion block is not a wake, e.g.
{No Secrets From the Magaji}.

#### 1.5.2 Locked bearers and torpor

Locking bars acting and blocking; it does not bar ability use. An ability stays usable while its bearer is locked unless
its own text requires the bearer to be unlocked, e.g. {Montano}, {Toby}. Where the ability costs a lock, only that part
needs the bearer unlocked: {Courier} lets you look at the card while locked and needs the lock only to discard
it.[^1-5-4]

Torpor bars no ability either: abilities remain usable there, including lock-to-use ones, e.g. {Montano}, {Mariel, Lady
Thunder}.[^1-5-5] Action modifiers are likewise not barred: they carry no unlocked requirement; what a vampire in torpor
can and cannot play is [§5.3.2](#532-what-a-torpid-vampire-can-still-do)'s.

#### 1.5.3 Ability use is not an action

A minion barred from acting may still use its abilities. A recruited ally cannot act the turn it enters play, but it can
use its abilities that turn, including ones that lock it, e.g. {Draeven Softfoot}, {War Ghoul}. The same holds for a
minion influenced out: its ability is usable during that influence phase.[^1-5-6] See
[§3.7.4](#374-employ-retainer-and-recruit-ally) for the recruit action and [§5.2](#52-locking-and-unlocking) for the
lock.

A vampire treated as an ally keeps abilities that let him play discipline-requiring cards, and uses them on an ally's
terms.[^1-5-14]

#### 1.5.4 Reuse

An ability may be used any number of times during a single action unless its own text limits it, as a "during" clause
does, e.g. {Corpse Minion}.[^1-5-7] Three limits come from the wording. If use locks the bearer or its source, it is
once per unlock, e.g. {Forest of Shadows}, {Hukros}. If it is keyed to a phase action, it is once per window available,
e.g. {Josef von Bauren}. Any other condition in the text is rechecked at each use: {Osric Vladislav} may repeat only
while he is acting and stealth is needed.[^1-5-8]

A once-each-turn ability is spent only if it actually applies. {Nergal}'s cost reduction is not used when the action
cannot be paid for even with it, or when the card is canceled as played and no cost is paid.[^1-5-9]

#### 1.5.5 Timing, participation, prevention

An ability that modifies an action has action-modifier timing: usable at any point during the action, including after
block attempts conclude, e.g. {Pentex(TM) Loves You!}, {Spiridonas}.[^1-5-10]

An ability whose text is framed by combat requires its bearer to be a combatant, not merely that a combat is occurring,
e.g. {Watenda}, {White Lily}.[^1-5-11]

An ability granted by a card in play is unavailable while an effect prevents the bearer from using that card: {Drawing
Out the Beast} stops the vampire using {Drum of Xipe Totec}, so neither the [CEL] nor the maneuver it grants is
available.[^1-5-12]

### 1.6 Requirements

#### 1.6.1 When a requirement is checked

A requirement is checked when the card is played and on each later play: {Undead Persistence} cannot be used twice in
one combat.[^1-6-1] During an action it is checked continuously; a minion who stops meeting it before the action is
blocked or resolves *fizzles* it at once: the action is successful, but it has no effect and its cost is not paid, and a
block attempt under way does not resolve.[^1-6-29] One met at the time of the block stands if the acting vampire is
later replaced.

A card already in play stays in play when its requirement lapses.[^1-6-2] A trait gained later applies to cards already
in play if the card checks for it, e.g. {Orun}'s capacity bonus[^1-6-3], not if it's only checked when the card is
played, eg. the discipline level for a retainer like {Raven Spy}.

#### 1.6.2 What counts as meeting a requirement

What "Only usable by ..." restricts depends on whether using the card is the same act as playing it. On a card that is
played and resolves — action, action modifier, reaction, combat card — it bars the play: {Regeneration} cannot be played
by an untorpid vampire. On a card that goes into play it gates only the conferred ability: a capacity-5 vampire can
still equip {Seal of Veddartha} and gains its [dom] and [for] levels, but cannot use its bleed action.[^1-6-4]

Meeting a requirement is not possessing the trait: {Talaq, the Immortal} plays [qui] cards but has no [qui] and gains
nothing from cards keyed to having it.[^1-6-5] A power the card prints as "not a Discipline" sets the card's level, but
nothing that counts, requires or taxes Disciplines sees it, e.g. [str] on {Masca} for {Dabbler}.[^1-6-26]

A card printing a clan and its antitribu as alternatives needs only one of the two, e.g. {Defender of the Haven}; a lone
clan requirement is not met by the antitribu clan.[^1-6-6]

#### 1.6.3 Cards entering play abnormally

Only "put ... in play" bypasses requirements and cost, e.g. {Summon History}, {Compel the Spirit} [NEC].[^1-6-7]
**Equip**, **recruit** and **employ** invoke the normal machinery, so requirements and cost apply as printed, e.g.
{Concealed Weapon}, {Piper}, {Magic of the Smith}. Such a card has still been played; see
[§1.8](#18-playing-and-canceling-a-card) for what it loses.

Where requirements are bypassed, a required discipline is used at inferior and a cost of X is zero.[^1-6-8] Card text
splits the two checks at will: {Bindusara, Historian of the Kindred} puts in play but prints that cost is paid; {Horrid
Reality} equips paying no cost, and burns the weapon if requirements are unmet.[^1-6-9]

Bypassing requirements does not bypass prohibitions: {Heidelberg Castle, Germany} cannot move equipment onto {Enkidu,
The Noah}. A card placed where it cannot legally sit is burned and no cost is paid for it, e.g. a second vehicle found
by {Vast Wealth};[^1-6-27] whether an effect it imposed on placement survives is
[§2.5.3](#253-effects-outliving-their-source)'s. On-entry clauses firing however the card entered are
[§5.5](#55-allies)'s.[^1-6-10]

A clause firing on equipping does not fire when the card is merely put in play: {Helicopter} comes locked when equipped,
unlocked when {Alastor} places it.[^1-6-12] An ability modifying equipping or recruiting does apply to non-normal ones,
e.g. {Little Tailor of Prague} with {Piper}. One keyed to the action applies only where there is an action: {Zhenga}
triggers on announcing a recruit or employ action, so {Piper}'s actionless recruit is outside her ability, while a
non-standard recruit that is an action is not.[^1-6-13]

The requirements and cost of a fetched card are not those of the fetching action: {CrimethInc.} cannot be used off {The
Summoning} for an ally requiring an Anarch.[^1-6-14]

#### 1.6.4 Vampires that meet requirements they do not have

{Mata Hari} and her kind are treated as meeting the requirement for all effects of the card played, and only them,
including its duration effects — never for what the card does from being in play, e.g. {Lady Scarlett Churchill}, chosen
as the locked Toreador of a {Toreador Grand Ball}, still unlocks as normal.[^1-6-15] {Vlad Tepes}, who "uses and can
play" such cards, also meets it for what the card does from being in play, e.g. {Abbot}'s intercept.[^1-6-28]

The ability operates only on a card played normally, not for a {Piper} recruit or a {Concealed Weapon} equip.[^1-6-16]
It operates only where the card prints the requirement, never to match a trait the card merely acts on, e.g.
{Sacrifice}.[^1-6-17] It does not operate while the vampire is uncontrolled or in the crypt.[^1-6-18] Controlling such a
vampire does not shut off burn options: a card in hand that only she could meet the requirement to play may still be
discarded via its burn option.[^1-6-25]

Title requirements are the special case: sect membership, the choice of city, and {Vidal Jarbeaux}'s printed
once-per-game caps are [§5.8.3](#583-how-requirements-read-titles)'s.

An ability granting a discipline for playing cards meets one, not both, of a two-discipline requirement, e.g. {Infernal
Familiar}, {Grey Thorne}.[^1-6-20]

#### 1.6.5 Playing a card that will do nothing

Futility is no bar by default: a card may be played, and an action taken, knowing it will accomplish nothing, e.g.
{Absorb the Mind} against a bloodless vampire, {Draba} on a minion already at 0 stealth, {Fragment of the Book of Nod}
with an empty library.[^1-6-21]

Three families bar the play. A card whose text names the object it acts on cannot be played when no such object exists,
e.g. {Shattering Blow} with no equipment, {Storage Annex} as the only card in hand, an empty crypt to draw
from.[^1-6-22] Immediate damage prevention cannot be played with no damage to prevent.[^1-6-23] Stealth or intercept
cannot be gained when not needed, e.g. {Bonding}. Reducing an opposing value already at its floor stays legal.

A card adding the unneeded value is barred even if it does more, e.g. {Bonding} at [DOM], whose +1 bleed the inferior
form still offers.[^1-6-30] These bars are on the play only. A card granting a standing effect may be played before the
need exists: {Repulsion} [OBE] may be played though the stealth is not currently needed, and once on the vampire its +1
stealth applies on every action regardless.[^1-6-24]

### 1.7 Costs and payment

#### 1.7.1 When a cost is paid

An action's cost is paid at resolution, not at announcement; any other card burns its cost as it is played.[^1-7-1]

An additional cost is paid at the same time as the cost it adds to. Where that cost is never reached — the action is
blocked or canceled, the referendum fails — nothing is paid.[^1-7-3] A charge incurred per combat is incurred anew if a
fresh combat begins: {Blithe Acceptance} burns its blood (or itself) again under {Psyche!}.

A cost charged for failing to block burns when the action begins to resolve, and reactions may still be played, and paid
for, before that burn.[^1-7-4] Costs and triggered burns charged on the block itself are
[§3.3.3](#333-the-cost-of-blocking)'s; a cost to *attempt* is per attempt and never retroactive, e.g. {Tenebrous Form}.

The cost of playing a card and the gain that card yields happen in one step, so no oust intervenes — even where another
card imposed the cost, e.g. {Secure Haven}'s surcharge on a {Minion Tap}. Where paying does oust a player, resolve the
oust before the effects that produced it.[^1-7-5]

#### 1.7.2 Cost arithmetic

Multiplication and division are applied first, then addition and subtraction.[^1-7-6] The modified cost is the card's
cost for every other effect that reads it, and a reduction tied to a minion follows the card onto and off that
minion.[^1-7-7] Modifications apply before affordability is checked, but paying from another source leaves the cost
itself unchanged.[^1-7-8] Where a reduction and another cost effect could both fire, the controller chooses the
order.[^1-7-9]

You must be able to pay the printed cost, even when the card returns more than it costs, e.g. {The Eternals of Sirius},
{Villein}; a reduction that brings the cost within reach makes the play legal.[^1-7-10]

Some cards have a cost of X, chosen as the card is played — for an action card, as the action is announced, though it is
paid at resolution. X may be zero, and "X" in the card's own text means the cost paid for that card. A card put into
play rather than played has X of zero, whatever the fetching card announced, e.g. {Summon History}.[^1-7-11]

#### 1.7.3 Cost reductions

A "lock to reduce the cost" location may be locked at any point from announcement until just before the cost is paid.
When the reduction is what makes the card affordable at all, it must be locked at announcement, e.g. {Sunset Strip,
Hollywood}. A location locked for an action card then canceled as it is played stays unlocked, e.g. {The Line}.[^1-7-12]
So does one locked for a card canceled with its cost still paid, though that card's other announced reductions apply
*(reading)*.[^1-7-30]

A reduction reaches what its wording keys it to: one on a card's cost applies however the card reaches play, including a
retainer employed by {Pack Alpha}; one keyed to an action does not apply where no action is taken, e.g.
{Charisma}.[^1-7-13] A reduction on "cards you play" covers minion cards, which the Methuselah also plays; wording
naming pool does not reach a cost paid in vampire blood, and the converse.[^1-7-14] Reductions reaching an equipment
that is a location are [§1.3](#13-card-types-and-multi-type-cards)'s.[^1-7-15]

#### 1.7.4 Costs and "burn blood" effects

A "burn X blood" clause inside a card's effect is not a cost. Cost reductions do not touch it, and the card cannot be
played if the minion cannot afford the burn, e.g. {Preternatural Evasion}, {Shadow Boxing}.[^1-7-16] An ability sparing
a vampire the cost of a class of cards therefore does not spare a burn-blood effect printed on one, e.g. {Dragos}
against {Chiropteran Marauder}.[^1-7-17]

An effect that adds to a card's cost is a cost: other reductions apply to it, and it reaches every card in the class it
names, including cards played at end of round, e.g. {Terror Frenzy}.[^1-7-18]

#### 1.7.5 Paying what you can

Where a card names an amount to take from a target, take what is there, e.g. {Villein}, {Theft of Vitae}.[^1-7-19] Where
the amount is a cost the payer cannot meet, the payment is still made as far as it goes and the effect it was to buy
does not happen, e.g. {Sword of Nuln}.[^1-7-20] An action whose cost can no longer be paid in full at resolution
fizzles: unblocked, it is still successful, so {Freak Drive} [for] may follow, but it has no effect.[^1-7-31]

An ally has life, not blood, and cannot pay a blood cost at all. Where the payer chooses the source, choosing an empty
one does not discharge the obligation, e.g. {Smiling Jack, The Anarch}.[^1-7-22]

#### 1.7.6 Whose cost, and paid in what

The fetching card's cost is paid even if its search finds nothing.[^1-7-23] Where a card names who pays, that minion
pays, acting or not, e.g. {Alastor}.[^1-7-24] Where a cost may be paid in blood or pool, or a choice of costs is
offered, the player playing the card chooses.[^1-7-25]

Using an ability of a card in play is not playing a card, so an added cost on card plays does not reach it.[^1-7-27] A
canceled card's cost is [§1.8](#18-playing-and-canceling-a-card)'s.[^1-7-28]

The "keep … by repaying their pool cost" template ({Kindred Segregation}, {Peace Treaty}) repays the pool cost only:
blood costs are not repaid, and a card with only a blood cost is kept for 0 pool. The controller may instead let the
card burn, even where the amount is zero.[^1-7-29]

### 1.8 Playing and canceling a card

#### 1.8.1 The "as played" window

A card played from hand in the normal fashion can be canceled as it is played. The cancellation must be immediate, and
cards are not replaced during the window, so a library-search master cannot fetch a canceller, e.g. {The Barrens},
{Dreams of the Sphinx}.[^1-8-1] A wake effect is the one other card playable in the window; a reaction that unlocks to
attempt a block is not.[^1-8-2] A cancellation from an ability in play is applied before anyone may play a further
cancellation, e.g. {Andrew Stuart} pre-empting {Direct Intervention}. A cancellation that is itself a card play opens
its own window and can be canceled there first — {Sudden Reversal} on the {Direct Intervention}.[^1-8-3] A reflex is
played by the minion the canceled card is played on, locked or not, e.g. {Predator's Transformation}
*(reading)*.[^1-8-21]

The window opens on an out-of-turn master, on a combat card played at the end of a round or "when combat would end", and
on a modifier played after the action succeeds, e.g. {Voter Captivation}.[^1-8-4] A card is cancelable as what it was
played as: played with [obe] it does not require [aus], e.g. {Hide the Mind}, {Iron Heart}.[^1-8-5]

#### 1.8.2 Played, but not in the normal fashion

A card brought into play by another card's effect, "put … in play" included, has still been played, but not in the
normal fashion. It has no as-played window, so it cannot be canceled as it is played. Cancellation reaches only normal
plays, minion cards and master cards alike — from hand, or "as if from your hand" where a card says so, e.g. {Persistent
Echo}, {The Erciyes Fragments} — not an ally recruited by {Piper} or a weapon equipped by {Disguised Weapon}. Nor has a
minion played it: the card that brought it in did, so an effect keyed to a minion playing a card does not see it, e.g.
{Louhi}'s "played by a vampire" against {Piper}'s ally.[^1-8-6] Requirements and cost still apply, unless the card is
"put … in play" ([§1.6.3](#163-cards-entering-play-abnormally)).

An effect that lets a minion *use the ability* of a card plays no card: no window opens, and it does not consume his own
once-per-action limit on that modifier, e.g. {Inscription}, {Shadow Court Satyr}. A cancellation in the used card's own
text still applies, e.g. {Target Vitals}.[^1-8-7]

When a card counts as played fixes when the window is. A political action card played from hand by {Charming Lobby}
counts as played only at successful resolution and so never becomes cancelable; one retained by {Echo of Harmonies} is
played normally and is.[^1-8-8]

#### 1.8.3 A canceled card has still been played

Cancellation stops the effect, not the play. The same reaction or modifier cannot be played again by that minion, and a
card limited to one per round cannot be replayed, e.g. {Immortal Grapple}.[^1-8-9] Effects that count or retrieve cards
played count the canceled one, e.g. {Dabbler}, {Marthe Dizier}, {Perfectionist}.[^1-8-10]

The cost is still paid unless the canceling card says otherwise. Most say "its cost is not paid", e.g. {Direct
Intervention}; those that do not, do not, e.g. {Santaleous} — including a blood cost imposed by {Terror
Frenzy}.[^1-8-11] Where only part of an effect is canceled, bids made are still paid and counters burned stay
burned.[^1-8-12]

A canceled limited effect does not trigger its limit, so another bleed or additional strike remains
available.[^1-8-14]

A canceled play spends a use that attached to the play itself: the canceled card was still played, so {Vidal Jarbeaux}
has spent his once-per-game requirement. A use that attached to an element the cancellation removed is not spent: no
cost was paid, so {Nergal}'s reduction never applied and his once-per-turn ability is not used.[^1-8-20]

#### 1.8.4 What a cancellation reaches

Cancellation propagates downward only: canceling an effect cancels what that effect provides, never the effect that
provided it. Canceling a strike cancels the rest of the strike card's effect, and the striking minion chooses another
strike.[^1-8-15] Canceling a maneuver does not cancel the strike that provided it — and that strike cannot be changed,
e.g. {Aid from Bats} — but does cancel a press the maneuver provides, e.g. the optional press on {Masque of Judas}'s
maneuver.[^1-8-16] Canceling an aim leaves the strike to resolve on its default target, e.g. {Target Retainer}.[^1-8-17]

### 1.9 Replacement

**Base rules.** A card played from hand is replaced immediately unless card text says otherwise, and the draw comes
after the "as played" window closes [[RBK playing-a-card]](https://www.vekn.net/rulebook#playing-a-card) [[RBK
drawing-cards]](https://www.vekn.net/rulebook#drawing-cards). With an empty library you do not draw and play continues.

#### 1.9.1 When the replacement is drawn

The action card is replaced before any action modifier is played on that action.[^1-9-1]

#### 1.9.2 Delayed replacement

A card printing "do not replace until after this action" is not replaced when the action is blocked; it is replaced when
the action ends, after all combats.[^1-9-4]

A delayed replacement is drawn at the first moment its condition is met, ahead of any other effect triggering then:
before cards unlock ({Port Authority}), before other effects keyed to the same torpor ({The New Inquisition} ahead of
{Fame}), before a queued combat begins ({Lucky Blow} ahead of {Psyche!}).[^1-9-5] The delay runs to its condition even
if the card has since been burned. A delay imposed by another card in play ends with that card: when it leaves play, the
held replacements are drawn at once, e.g. {The Meddling of Semsith} *(reading)*.[^1-9-12]

"Do not replace" effects stack with a card's own clause and the longest delay wins.[^1-9-6] Where two effects
each offer a different replacement treatment, the controller chooses, e.g. {Visit from the Capuchin} against {Steely
Tenacity}.

An effect tied to the moment a card is played resolves at that moment, and a delayed replacement is simply not there
yet. That can reduce the effect — {Troglodytia} after {Wash} sees the hand without the replacement — or negate it
entirely: {Agaitas, The Scholar of Antiquities}' ability applies only as the card is played, so any delay prevents it.
An effect keyed to the replacement itself waits for it, e.g. {Learjet}.[^1-9-7]

#### 1.9.3 Abnormal entry and cancellation

A "do not replace until after this action" clause governs only a normal play
([§1.8.2](#182-played-but-not-in-the-normal-fashion)). A card put into play in a special way is replaced immediately,
even mid-action, e.g. {Baseball Bat} played via {Concealed Weapon}; played normally the clause still governs, even when
the card reaches play by another route during the resulting combat, e.g. {Gift of Bellona}.[^1-9-8] A card played
normally is replaced before its action resolves, so an ally recruited in the normal fashion can be given cards drawn as
its own replacement; one entering play by other means cannot, e.g. {Corrupt Construction}. A card played "as if from
your hand" off another card is never replaced, so its "do not replace" clause costs no hand size, e.g. {The Erciyes
Fragments}.[^1-9-13]

A canceled card is replaced normally, its "do not replace until" or alternate-replacement clause canceled with it, e.g.
{Steely Tenacity}.[^1-9-9]

#### 1.9.4 Empty crypt, hand size

A card whose effect draws from or moves a card out of a crypt cannot be played when that crypt is empty — your own for
{Illusions of the Kindred}, the blocker's for {Bear-Baiting}.[^1-9-10] This is [§1.6](#16-requirements)'s missing-object
family, not futility.

Cards not yet replaced count against hand size, and an effect that shows you a hand does not show the card drawn to
replace a discard from it. See [§6.8](#68-hand-draw-and-discard) for both.[^1-9-11]

### 1.10 Burn, removal from the game, and shuffling back

#### 1.10.1 Burn is not removal from the game

Burning a card and removing it from the game are separate events. A "when burned" trigger does not fire when its object
is removed instead: {Soul Gem of Etrius} does nothing when its bearer is removed by {Golconda: Inner Peace}.[^1-10-1]

The clauses apply in printed order: where a card is burned and then removed, the burn happened and every when-burned
trigger fires — an {Absimiliard's Army} ghoul is burned first, so {Tension in the Ranks} sees it. Cards controlled by an
ousted Methuselah are removed, not burned, so nothing triggers.

Where a card prints removal as its own disposition, an actual burn pre-empts it and the card goes to the ash heap
instead, e.g. an illusionary vampire burned in combat, or a {Grasp the Ghostly} equipment burned before its pathos
counters run out.[^1-10-2]

A card removed as part of resolving its own effect never reaches the ash heap, so ash-heap retrieval cannot find it,
e.g. {Summon Soul} at [NEC].

Cards removed from the game are public.

#### 1.10.2 "Burned" means burned from play

An effect keyed to a card being burned means burned from play, not any route to the ash heap. A card discarded from hand
has not been burned, and neither has an ally card whose recruit action was blocked; {Resurrection} recovers
neither.[^1-10-3]

The same test governs a shuffle-back clause: a card that shuffles into its owner's library when burned is not shuffled
back when it is discarded.

Such a card is burned first and shuffled back second, so other when-burned effects still apply — {Set's Curse} still
puts its ally into play when the minion it burns is {Kherebutu}.

#### 1.10.3 Effects conditioned on a burn

A card whose effect is conditioned on a burn may be played when that burn will not happen. The conditioned part simply
does not fire. {Set's Curse} may be played on {Ilomba}, which redirects its own burn, but then nothing is put into play.
{Sacrificial Lamb} gains no blood and moves no equipment if the target does not burn.[^1-10-4]

A card that burns itself at a stated moment does not burn if that moment is never reached: {Bomb} used as a strike does
not burn if combat ends before the strike resolves, and a weapon out of play when a burn would apply escapes it.

A card that burns itself for want of counters is a counter-count question, not a burn question; see
[§1.12](#112-cards-and-counters-on-other-cards) for it.[^1-10-5]

#### 1.10.4 Crypt and library are separate

Crypt cards cannot be put into the library and library cards cannot be put into the crypt. Continuous effects stop
tracking a card shuffled back into the crypt, and any effect standing on that card is voided. Where one effect both
discards and shuffles the discards back, resolve all discards and draws first, then shuffle.[^1-10-6]

An effect does not need the minion it names to remain in play. A contract can still be used after its target — or the
contract card itself — is burned, e.g. {The Black Throne}; a card keyed to what happened in the combat may be played
after the minion it measures was burned, e.g. {Taste of Vitae}; and a prohibition fixed on a card survives even that
card's own trip through the ash heap — {The Eternal Mask} still cannot be burned after its vampire burned and returned.
See [§2.5](#25-duration-and-persistence) for persistence.[^1-10-7]

### 1.11 The ash heap

The ash heap has no order. A player removing or taking cards from an ash heap chooses them freely, e.g. {Trochomancy},
and the choice is visible to every player even when the card is then placed face down, e.g. {Maabara}.[^1-11-7]

#### 1.11.1 Retrieval: what counts as played

A retrieval effect worded on "a card played" reaches only cards played the normal way. A card brought into play by
another card's effect is not retrievable by such an effect, e.g. {Concealed Weapon}, {Charming Lobby}. That card has
still been played, though not in the normal fashion; the two tests are separate. A card canceled as it is played counts
as played and can be retrieved. The action card of the action itself counts among the cards the acting minion played
during that action, e.g. {The Art of Memory}.[^1-11-1]

Burning a card is not playing it. A card burned from hand for a vote was not played (it is burned, but neither played
nor burned from play;), and a card in play burned by its own text was not played to call the referendum that burned it —
{Echo of Harmonies} cannot retrieve a {National Guard Support} burned that way.[^1-11-2]

#### 1.11.2 The ash heap is read when the effect resolves

A card that goes to the ash heap when played and is referred to later has no effect if it has been removed from the ash
heap before that moment, e.g. {Melange} when the vampire blocks. A retrieval likewise finds nothing if its target has
left the ash heap before the retrieval resolves, e.g. {Echo of Harmonies} after {Delaying Tactics}.[^1-11-3]

A card diverted before it reaches the ash heap was never there. {Delaying Tactics} returns the political action card to
its owner's hand before it is burned, so a card played from {The Erciyes Fragments} returns to hand rather than being
removed from the game. A retrieval naming a card in the ash heap cannot be played at all if that card was moved
elsewhere, e.g. {Compel the Spirit}.[^1-11-4]

#### 1.11.3 What a card is in the ash heap

A card in the ash heap is its own printed card, not what it represented in play. A library card that became a minion in
play is not a minion of that type there: {Jake Washington} is no ally in the ash heap, and an action card that created a
vampire leaves no vampire there. A crypt card in the ash heap is a vampire whatever it represented in play, so a burned
{Spell of Life} mummy can be retrieved by {Possession} but not by {Compel the Spirit} — and an ally that was represented
by a crypt card cannot be retrieved as an ally. An effect keyed to a card "burned from play" does not reach a card that
was never in play: a discard, or an ally card sent to the ash heap by a blocked recruit action, does not qualify, e.g.
{Resurrection}.[^1-11-5]

#### 1.11.4 Whose ash heap

A burned or discarded card goes to its owner's ash heap. A minion controlled by another Methuselah passes through its
owner's ash heap when burned, breaking all control effects, e.g. {The Capuchin}; a stolen equipment burned in play
returns to its owner's ash heap and is not removed from the game unless an effect says so, e.g. {Grasp the Ghostly}. A
retrieval worded "your ash heap" means the ash heap you own — it does not matter who first controlled the card, e.g.
{Père Lachaise, France}. Effects may reach any player's ash heap, and an action targeting one is always undirected
([[RBK important-terms-of-the-game]](https://www.vekn.net/rulebook#important-terms-of-the-game)).[^1-11-6]

### 1.12 Cards and counters on other cards

**Base rules.** Burning a card or removing it from the game burns any counters and cards on it. Moving a minion between
regions does not: cards and counters follow the minion and are out of play while it sits in torpor or in the
uncontrolled region.

#### 1.12.1 The host leaving play

A card on a host is burned when the host leaves play, whatever sends it away: burned, removed from the game, returned to
a hand, library or crypt, or set aside with no printed provision for its cards.[^1-12-1] Cards on the initial target of
a card effect burn like any other, e.g. {Memory's Fading Glimpse} (host to the bottom of the crypt), {Raw Recruit} (host
set aside out of play). Two families of departures keep the cards instead, out of play but intact: torpor, the
uncontrolled region and contest, where they burn only if the host is yielded or burned; and a set-aside whose own text
carries them along, e.g. {Descent into Darkness} ("this vampire and any other cards and counters on them").[^1-12-8] A
vampire that burns and returns comes back with nothing on it and is a new minion for every purpose.

A hosted card printing its own "burn this card if this vampire is about to leave the ready region" clause resolves
before the move. It burns even under an effect that would otherwise keep cards on the vampire, e.g. {Righteous Aura}
under {Banishment}.[^1-12-2]

Effects the hosted card already produced are not undone when it burns with its host, e.g. {The Eternal Mask}. Counters
it placed elsewhere keep working, e.g. {Tegyrius, Vizier}'s allegiance counters. See
[§2.5](#25-duration-and-persistence) for persistence.[^1-12-3]

#### 1.12.2 Status of a hosted card

A card put on another card is not in play, even where the host's text does not say so, e.g. {Shadow Court
Satyr}.[^1-12-4]

A minion card on a minion is controlled by that minion's controller, so it stays when the Methuselah who played it is
ousted, e.g. {Anathema}.[^1-12-5] It follows its host on a change of control, and the new controller may use it. When it
would move to an ash heap or library it goes to the owner's, e.g. {Storage Annex} stolen.

A copy entering play into a contest resolves its enter-play effects before the contest begins: a second {Mokolé Blood}
puts its cards on itself first; its standing effects do not apply. See [§1.13.2](#1132-entering-contest) for entering
contest.[^1-12-6]

#### 1.12.3 Counters on hosts

A card that enters play with zero counters and prints "burn this card when it has no counters" burns immediately, e.g.
{Secret Horde} put into play with X=0.[^1-12-7]

A counter moved onto a master card becomes the counter type that card uses, e.g. {Goth Band}.

A threshold of counters "equal to the card's cost" reads the number, not the currency: {Dominique} burns a location that
cost X blood with X vandal counters, the same as one that cost X pool.

### 1.13 Contests

**Base rules.** Contested cards are turned face down and out of play for the contest's duration; a yielded card is
burned. A contested *title* leaves its vampires in play and acting normally, treated as having no title.

#### 1.13.1 What contests what

A contest needs two cards that are both unique and share a card name, whatever the card type. The unique mummies made by
{Spell of Life} contest each other and a unique vampire of that name, but never the non-unique {Aabbt Kindred}.[^1-13-1]

A card whose text says it does not contest is treated as a different card for everything that reads names, not merely
spared the contest. Its duplicate title still votes, and an effect triggered by a second copy of that name does not
fire: an illusionary {Jimmy Dunn} from {Illusions of the Kindred} does not make the real one burn himself. {Jimmy
Dunn}'s own "cannot be contested" works the same way, so a second copy may be influenced out.[^1-13-2]

You can never choose to contest a card with yourself. A play or choice that would do so is illegal: a second {Visit from
the Capuchin} cannot be played, and {Chain of Command} cannot bring out two copies of one unique vampire. Where an
effect forces the contest instead, the incoming copy is burned, e.g. two {Parmenides} influenced in one turn.[^1-13-3] A
title claimed by two of one's own vampires is [§5.8.1](#581-printed-titles-granted-titles-contests)'s.

#### 1.13.2 Entering contest

Entry costs, cards placed on the card, and "as this enters play" abilities all resolve before it enters contest, and
only before. {Dr. Solomon Grey} burns his pool first; a second {Mokolé Blood} collects its cards first; {Anarch Convert}
may remove himself from the game, averting the contest entirely.[^1-13-4]

The copy already in play leaves play when the contest starts, so its "if this leaves play" clause fires: {Sonja Blue} is
removed from the game, handing the contest to the incoming copy.[^1-13-5]

#### 1.13.3 While contested

A contested card is out of play and nothing printed on it applies. A contested location gives no hand size, whether
printed on it ({Elder Library}) or counted by another card ({Guillaume Giovanni}); {Secure Haven} is not burned when its
minion goes to torpor; {Byzar}'s burn-replacement does not save him from being yielded.[^1-13-6] The cards on a
contested minion are out of play with him, so none of their text applies; if he is yielded they burn with him, none of
their burn riders firing, e.g. {Soul Gem of Etrius} *(reading)*.[^1-13-12]

A weapon taken into contest mid-combat is unusable by either side, and a strike already chosen with it has no effect,
e.g. {Disguised Weapon}.

An effect elsewhere that reads the contested card is suspended, not ended, and resumes when the contest resolves, e.g.
{Betrayer} naming a vampire who becomes contested. A contested card carried by a minion that an effect moves out of play
drops out of the contest and re-enters it when the minion returns, e.g. {Descent into Darkness}.[^1-13-7]

#### 1.13.4 After the contest

What a card returning from contest remembers is
[§6.4.1](#641-set-aside-out-of-play--the-card-remembers-everything)'s.

A fresh copy of the same name is a different card, and a choice fixed on a copy that was burned or yielded does not move
to it. {The Rack} does not pick up the new copy, nor may its controller re-choose merely because the chosen vampire
became ineligible.[^1-13-10]

A card that leaves play through a contest is not burned by any minion, so nothing keyed to a minion burning it
applies.[^1-13-11]

### 1.14 Set-aside and announced cards

#### 1.14.1 When the card is named

An action that moves a named card out of your ash heap names that card when the action is declared, not at
resolution.[^1-14-1] This holds whether the action comes from a card, e.g. {The Sargon Fragment}, or from a minion's own
ability, e.g. {Pochtli}.

A search of your library declares nothing: the card is chosen at resolution, e.g. {Magic of the Smith}.[^1-14-2] So is a
choice belonging to an effect that only triggers later, e.g. {Ashur Tablets}.

The card providing the current action is not itself available to that action. While it is resolving it is in limbo —
neither in hand nor in the ash heap — so it cannot be named among the cards it retrieves, e.g. {Siphon}, {Sudario
Refraction}.[^1-14-3]

#### 1.14.2 Cards named from hand

A card named from hand must be in hand when the naming card is played, before the replacement for that card is
drawn.[^1-14-4] It need not be shown: you may set it apart face down, e.g. {Disguised Weapon}, {Charming Lobby}. The
same applies where the naming card takes a count rather than a name, e.g. {Gift of Proteus}.

Naming does not commit the card. If the action is canceled or fails, the named card stays in hand and is not discarded,
e.g. {Jack of Both Sides}.

A card playable "as if from your hand" is not in hand. An effect requiring a card in hand cannot reach it: {Charming
Lobby} cannot call a referendum off a political action card held by {Echo of Harmonies}.[^1-14-5] Likewise, an effect
that returns the political action card played to call a referendum finds nothing when the referendum came from a card in
play, e.g. {Delaying Tactics}.

#### 1.14.3 Cards out of play

A card set aside face down under an effect is out of play. It is in no zone that discard and burn effects reach, so it
cannot be discarded or burned to satisfy another effect, e.g. a card on {The Erciyes Fragments} against {Tension in the
Ranks}.[^1-14-6] See [§1.12](#112-cards-and-counters-on-other-cards) for cards placed on cards in play.

An equipment that is temporarily out of play is not burned by a scheduled burn that applies while it is out.[^1-14-7]
{Kerrie}'s ranged strike turns it face down until the end of the action ({Dagger} likewise), so {Baal's Bloody Talons}'
end-of-round burn misses it; {Cleave}'s burn lands at the end of the action itself, where the acting Methuselah orders
it against the weapon's return.

### 1.15 Cumulative and stacking effects

Multiple copies of the same card in play each exert their own effect and impose their own cost. Two {Anarch Revolt} burn
2 pool in the unlock phase; two {First Tradition: The Masquerade} burn 4 pool.[^1-15-1] Burn-pool and burn-blood
requirements from any number of sources are cumulative. Where a card offers an alternative to its cost, taking that
alternative once satisfies every copy — skipping a turn counts against all copies of {First Tradition: The Masquerade}.

Numeric bonuses from multiple copies add, e.g. {Mass Reality} damage, {Orun} capacity.[^1-15-2] An absolute value does
not add: a card granting "X" rather than "+X" sets its value, so multiple {Legacy of Pander} give a non-titled Pander 1
vote in total, and {Torn Signpost} sets strength to 2 or 3 — which can even reduce it, e.g. a 3-strength {Rock Cat}
using it at [pot]. "+X" effects then add to the set value. {Orun} counts its own copies by its own text: one vote per
three. A vote bonus that is not a title adds to the votes of a title gained later, e.g. {Xeper, Sultan of Lepers}.

Each copy triggers and resolves on its own. Multiple copies of {Camarilla Exemplary} may each name the same vampire, and
multiple {Courier} each look at and discard a card in succession.[^1-15-3] Copies attached to one minion are never
merged: two {Kindred Society Games} on a vampire cost 2 blood to unlock and move separately, and stealing a minion under
two {Temptation} burns the counters on only one of them.

An effect conditioned on having *any* counters is flat and does not scale with the count. A minion with graft counters
from {Dr. Morrow, The Skindoctor} gets -1 stealth, not -X stealth; {Kahina the Sorceress} inflicts 1 damage however many
corruption counters the minion carries.[^1-15-4] The same holds for a trigger tied to an event rather than to a
magnitude: a vampire with several {Orun} burns one on a successful bleed for more than 2, and each {Scorn of Adonis}
costs a given Methuselah at most 1 pool however many "no" votes that Methuselah casts.

Effects from *different* cards that grant the same discipline do not combine into the superior level. An ally allowed to
play [pot] cards by both {Leech} and {Putrescent Servitude} still plays only [pot], never [POT].[^1-15-5]

A once-per-action limit on playing a card binds the minion, not the action. Different minions may each play their own
copy on the same action, e.g. {Cloak the Gathering}, {Suppressing Fire}.[^1-15-6]

Two copies of a card that imposes a mandatory directed action can deadlock the bearer: a minion under two {Lunatic
Eruption} is stuck and cannot act.[^1-15-7] See [§3.9](#39-mandatory-actions).

---

## 2. Timing and Sequencing

### 2.1 Effect windows and impulse

**Base rules.** The impulse is the opportunity to play the next card or effect. The acting Methuselah plays first and
gets the impulse back whenever anyone else uses a card or effect; the window closes only when every Methuselah passes in
turn.

#### 2.1.1 Where the impulse exists

An ability worded "when its controller has the opportunity to play effects" is usable at every such point, including
inside an action.[^2-1-1] An ability framed by a phase, e.g. "during your minion phase", is likewise usable at any
impulse in that phase, inside another minion's action or a combat included, e.g. {Rutor's Hand}, {Qawiyya el-Ghaduba}
*(reading)*.[^2-1-10] Mandatory effects are applied before their controller passes the impulse.

Permission to play one card outside its usual window opens no window for anything else: a cancel card usable when there
is no action does not thereby allow a wake effect, e.g. {Rewind Time}.

#### 2.1.2 Who may play what

Only the acting minion plays action modifiers; only other Methuselahs' ready unlocked minions play reaction cards. The
window runs until resolution, so a modifier may still be played after all block attempts are concluded.[^2-1-2]

The template "Only usable by a ready vampire other than the acting minion" widens which minion plays the modifier, not
which player. The card is still an action modifier, only the acting minion's controller may play it, and the vampire
must be one that Methuselah controls, e.g. {Blanket of Night}, {Siren's Lure}, {Hidden Lurker}.[^2-1-9]

A reaction card widens its user only as far as its own text goes; the "not involved in the current combat" permission is
[§4.1](#41-combat-sequence-and-rounds)'s.[^2-1-3] A [REACTION] ability — the marker borne by imbued powers, e.g. {Hide},
{Surge} — is not a reaction card play but still needs a reaction window.

#### 2.1.3 Effects resolve atomically

All effects are instantaneous except actions and certain combat effects. Nothing may be played "in response" to an
effect, and no window opens between the sub-steps of one resolving effect, e.g. {Heidelberg Castle, Germany} cannot be
locked in response.[^2-1-4] Only a card that prints the exception opens one, and a "for each X" clause resolves once,
as a whole, over the X in play when it resolves.[^2-1-8]

So a vampire chosen to go to torpor by {Baltimore Purge} cannot be unlocked between the choice and the torpor by another
effect; a vampire changing controller because of {Temptation} cannot be locked between the unlock and the change of
control; and life cannot be added in the middle of a single damage resolution step by {Vagabond Mystic}.

Between two actions, including during another Methuselah's turn, every Methuselah has the impulse, and any effect not
bound to an action window is usable there — a card barred "during an action" like {Heidelberg Castle, Germany} only
there, an unbarred one like {The Louvre, Paris} or {Dreams of the Sphinx} at any impulse. A combat generated by an
action modifier is still part of that action, so no between-actions window opens between two such combats ({Siren's
Lure}).[^2-1-5]

#### 2.1.4 Step windows

A card that changes how a step resolves must be played before that step resolves. {Spirit Claws} affects the current
strike only if played before strike resolution; {Rötschreck} is played after a strike is declared and before resolution,
in either strike order. A card played on a diablerie may come before or after the Discipline card is taken, but must
precede the blood hunt.[^2-1-6]

A lock-to-reduce cost reduction may be applied at any point up to payment. A card in play that must be burned on an
action is burned before that action resolves ({Melange}).

The "as the action/bleed would be successful" window is the last one before resolution; redirection must precede
it.[^2-1-7]

### 2.2 "As the action is announced" effects

**Base rules.** A card worded "only usable as the action is announced" must be played before any regular action modifier
or reaction card [[RBK
summary-of-the-course-of-an-action]](https://www.vekn.net/rulebook#summary-of-the-course-of-an-action).

The window opens once the action card has been played and closes as soon as the first regular action modifier or
reaction card is played, by any player. No as-announced card may follow it.[^2-2-1]

The timing is a property of the card's wording, not of its card type. Action modifiers, reaction cards and cards already
in play all carry it, e.g. {Car Bomb} [REACTION] and the location {Yoruba Shrine}. The acting Methuselah has the first
opportunity to play in this window, as everywhere else in the action.[^2-2-2]

The action card is played, and replaced, before the window opens. A card drawn as that replacement may therefore be
played as the action is announced. See [§1.9](#19-replacement) for replacement timing.[^2-2-3]

Because the action card has already been played, an effect used in the window that forbids playing cards does not reach
it. {Concoction of Vitality} may be burned as the action is announced even though the action was announced with a card
requiring a discipline, e.g. {Govern the Unaligned}.[^2-2-4]

The deadline at the action's other end is [§3.4.3](#343-the-would-be-successful-window)'s.

[§1.6](#16-requirements)'s bar on gaining stealth you do not need reaches this window: an as-announced card comes before
any block attempt, so the stealth is never needed yet. Stealth granted here therefore prints its own override —
{Predator's Transformation}: "+1 stealth, even if stealth is not yet needed".

### 2.3 After resolution and after combat

**Base rules.** An action resolves before it ends. Cards keyed "after resolution" or "after combat" are played in the
gap between the two, while the action is still in progress.

#### 2.3.1 The post-resolution window

A card played after resolution is played during the action, not after it.[^2-3-1] Durations running to the end of the
action are still running: a minion taken control of until the end of the action returns to its controller only once
these effects are done, e.g. {Spirit Marionette}. An oust does not close the window either — with {Last Stand} in play,
post-resolution effects are played before the turn ends. Being played during the action, such a card can be canceled as
it is played, e.g. {Ophidian Gaze} on {Voter Captivation}.

Once a card keyed to resolution has been played, only cards with that same timing may follow. Cards keyed to an earlier
point — {Deflection}, {Archon Investigation}, {Conditioning} — must precede it.[^2-3-2] A printed after-success effect
on a permanent works the same way and does not close the window: cards playable after resolution may still be played
after {Perfectionist} gains its blood. The window opens at resolution; the "would be successful" window is earlier still
and cannot follow it.[^2-3-3]

Modifiers and reactions usable at the end of or after an action survive a failed action and survive an intervening
combat, e.g. {Champion}, {Car Bomb}, {Ensconced}.[^2-3-4] An action that *ends* closes the window immediately and
nothing may follow, e.g. {Change of Target}.

{Momentary Delay} puts the action back in progress rather than adding to it, continuing it as if unblocked: other "after
resolution" cards may be played before it, not after.[^2-3-5]

#### 2.3.2 Ordering

The acting player chooses the order of effects played after resolution and after combat, and may play further cards of
that timing between them.[^2-3-6] Cards keyed "after combat" order freely against one another, e.g. {Provision of the
Silsila} and {Monster}.[^2-3-7]

Damage taken after resolution, e.g. {Force of Will}, and blood gained after resolution are both inside the window; cards
usable after resolution may be played before or after them.[^2-3-8] Blood gained there is a gain like any other and
triggers gain-keyed cards, e.g. {Slake the Thirst}.

A reaction card that unlocks its vampire may be played after resolution to enable another reaction card, e.g. {Wake with
Evening's Freshness} into {Fast Reaction} after combat.[^2-3-9]

#### 2.3.3 After combat

Cards keyed "after combat", "after block resolution" or "after action resolution" are played once every queued combat of
the action has been handled, e.g. {Cats' Guidance}, {Freak Drive}, {Hay Ride}.[^2-3-10] A card usable after block
resolution cannot be played at all if the action was not blocked.[^2-3-14] The combat a block causes is the resolution
of the blocked action, and where the action resolves before combat begins, "after resolution" effects wait for that
combat too; the same holds for a success-conditioned effect from a card played earlier in the action, e.g. {The Art of
Memory}.[^2-3-11] With a further combat queued, none of these windows opens between the combats — all wait for the last
one. The impulse between two combats belongs to other effects: a reaction keyed to the coming combat can end the action
there ({Obedience}), and a per-combat effect drops and re-applies ({Raptor}'s hand-size penalty is redone for the new
combat); cards barred during the action stay barred.[^2-3-13]

#### 2.3.4 What resolution includes

An effect embedded in the action's own resolution happens before the window opens. A blood hunt referendum called by the
action's text is part of resolution, so "after resolution" cards follow it and {Heidelberg Castle, Germany} — which
cannot be used during an action — cannot be used before it.[^2-3-12] The diablerie blood hunt is different: part of the
action but not of its resolution; the window before it is [§3.7.8.1](#3781-the-window-before-the-blood-hunt)'s.

### 2.4 Simultaneous effects and ordering

**Base rules.** When two or more effects apply at the same moment, the acting Methuselah plays and orders first, then
the impulse passes ([[RBK sequencing]](https://www.vekn.net/rulebook#sequencing)).

#### 2.4.1 Ordering within a window

The acting Methuselah decides the order of the effects available in a window, and may play further effects in it before
passing.[^2-4-1] Nothing else forces an order: cards playable at the end of the round or "when the combat would end" may
be played before or after one another,[^2-4-2] and a limited bleed modifier or limited additional strike may be played
before or after an unlimited one.[^2-4-3]

Order is fixed only where wording fixes it. A replacement effect — "would X … instead" — must be played while the event
is still replaceable, so it cannot follow an effect keyed to the event merely happening: {Telepathic Tracking} at [AUS]
("combat would end. Instead, start a new round") cannot be played after {Psyche!} ("combat is about to end"); played
first, combat continues and {Psyche!} waits for the end of the new round.[^2-4-4] An effect that reads a value applies
after every effect that modifies that value, e.g. {Protected Resources} after all bleed modifiers.[^2-4-5] An "after X"
effect of a card in play resolves as soon as X happens, inside whatever caused it: {Fame}'s 3 pool is burned mid-combat,
not in the after-combat window *(reading)*.[^2-4-25]

Which window a card belongs to is [§2.3](#23-after-resolution-and-after-combat) and
[§4.9](#49-end-of-round-end-of-combat-and-new-combats); ordering runs the same in each. An action that ENDS closes its
window at once, leaving nothing to order.

Across Methuselahs, whoever uses an effect first has priority and whoever uses one last has the final say. In the "when
a vampire should go into torpor" window the acting minion goes first, and going first can lock the other side
out.[^2-4-9]

#### 2.4.2 Order is a real choice

Effects resolve one at a time, so a later one can find nothing left to do: a second {Anathema} is burned
unresolved.[^2-4-10] Ordering also dodges a cost or a trigger — apply a cost reduction before {Jenna Cross}'s surcharge,
or burn the last non-Camarilla vampire before {Judgment: Camarilla Segregation} bills you.[^2-4-11] Separate triggers
from one event resolve one at a time in that order, so the ousts they cause are not simultaneous: two minions burned at
once under {Tension in the Ranks} cost their controllers 1 pool each, in the order the acting Methuselah sets.[^2-4-26]

Where one effect could attach to any of several simultaneous applications, its controller chooses which, e.g. {Slake the
Thirst}.[^2-4-12] The clauses of a single card resolve in printed order.[^2-4-13]

#### 2.4.3 Delayed effects triggered by a block

When a block triggers a delayed effect, no other effect may be used before it resolves, e.g. {Millicent Smith, Puritan
Vampire Hunter}, {Unleash Hell's Fury}. Effects triggering at that same moment are still applied and ordered by the
acting Methuselah; modifiers may follow.[^2-4-18]

#### 2.4.4 Unlock phase

Held-back replacements and control reversions resolve before unlocking and cannot be ordered among unlock
effects.[^2-4-20] Within the phase the acting Methuselah orders his own effects freely, but must finish the mandatory
ones before another Methuselah gets the impulse.[^2-4-21]

#### 2.4.5 Effects that are not ordered

Some effects are simultaneous, not ordered. Two {Theft of Vitae} strikes net the gain against the loss; no blood wears
off.[^2-4-22] {Deep Song}'s lock and enter-combat are one moment; both are lost if the action ends first.[^2-4-23]
Damage is prevented in sequence — acting player first — but all of it is applied at once.[^2-4-24]

### 2.5 Duration and persistence

#### 2.5.1 Effects with a stated duration

An effect granted for the duration of an action runs until the action ends. A combat the action causes is inside it: the
acting minion stays the same and effects applied during the action still apply there. They also survive an action
continuing as if unblocked.[^2-5-1]

An allowance given for one action ends with the action even when the card granting it stays in play: {Eyes of the
Beast}'s maneuver serves only that action though the card sits on the vampire until the discard phase. Within the action
it is not spent by being available — {Marciana Giovanni, Investigator}'s beneficiary may burn the blood for intercept at
any point. A prohibition triggered by attempting to block runs from the attempt to the end of the action, whether or not
the block succeeds, e.g. {Valerius Maior, Hell's Fool (ADV)}.[^2-5-2]

"Until your next X phase" ends as that phase begins, and an effect created during your own turn before that phase
reaches this turn's, e.g. {Fotini Katsikaris}'s hand size, {Mirror Walk}'s "do not replace until your next discard
phase" played in your minion phase *(reading)*.[^2-5-15]

#### 2.5.2 Fixed at play versus read continuously

A value a card writes is fixed when the card is played and is not re-read. {Derange} takes the clan and sect from the
acting vampire at play; that vampire changing clan afterwards does not move it. Counters added later do not change an X
already set, e.g. {Leadership Vacuum}. Removing the card restores the underlying value.[^2-5-3]

An effect whose text reads a trait on an ongoing basis re-reads it, including mid-action. A relational referent is
re-read at each use — {Victim of Habit} keeps its named card when the prey is ousted and then works against the new
prey.[^2-5-4] "A bleed against your prey" still names the Methuselah that bleed ousted, e.g. {Enkil Cog} played after it
*(reading)*.[^2-5-16]

A requirement is checked when the card is played and the effect continues after the vampire stops meeting it, e.g.
{Concordance}, {Cry Wolf}.[^2-5-5]

An effect granted to a class of minions applies to the members in play when it resolves: {Annabelle Triabell}'s
turn-long +1 bleed does not reach a Toreador entering play later that turn. A standing grant is not consumed by use: the
Toreador chosen for {Toreador Grand Ball} stays unblockable on non-bleed actions even after attempting a bleed.[^2-5-7]

#### 2.5.3 Effects outliving their source

An effect already applied stands when its source leaves play. {Tegyrius, Vizier}'s allegiance counters keep working
after he leaves; {Camarilla Vitae Slave} grants its discipline until the next master phase even if the retainer is
burned or stolen.[^2-5-8]

An effect equally outlives the player or minion that produced it: {Shatter the Gate}'s counters stay active after the
acting Methuselah is ousted, and minions that equipped while {NRA PAC} was in play unlock at end of turn whether or not
it still is. A card in play keeps working while its bearer is in torpor, e.g. {Nahir}'s hand size.[^2-5-9]

A card reaches only what happens after it is played. {Frenzy} does not undo equipment used earlier in the round;
{Kraken's Kiss} does not count damage inflicted before it; gaining or losing {Sire's Index Finger} does not reach a
frenzy card already played.[^2-5-10]

Whether a persistent effect ends with its source is decided by the card's printed duration clause. {Rowan Ring} and
{Wooden Stake} tie "does not unlock as normal" to the victim remaining in torpor, so it holds after the weapon is burned
or moved. {Rötschreck} states it flatly in its own in-play text, so it ends when that card is burned or
removed.[^2-5-11]

An effect naming a card as its object keeps tracking it across a change of bearer or controller: {Horrid Reality} burns
the weapon at end of combat wherever it now sits, and {Imposing Phantasm} returns the blood even if the opposing minion
changed controller.[^2-5-12]

#### 2.5.4 Leaving play

A minion that goes to the ash heap loses every continuous effect on it and is lost track of by cards and effects, e.g.
{The Capuchin}. Cards shuffled back into the crypt are lost track of the same way. A minion that changes type loses the
continuous effects attached to its old form, e.g. {Demdemeh}'s retainers.[^2-5-13]

The uncontrolled region suspends rather than erases: cards, counters, titles and continuous effects — even those worded
for the rest of the game — resume on return, e.g. {Banishment}; suspended capacity modifiers are
[§5.1.1](#511-capacity)'s.[^2-5-14]

---

## 3. Actions and Politics

### 3.1 Announcement and targets

#### 3.1.1 What is fixed at announcement

Every term the card calls for is named when the action is declared: the target, the counters spent, the card burned, the
affected Methuselah.[^3-1-1] A lock-to-reduce cost reduction is not such a term: it may be applied at any point from
announcement until just before the cost is paid.

A card named as part of the announcement must be in hand at that moment, face down if you wish, and stays in hand if the
action is canceled or fails; a card retrieved from your ash heap is named at declaration, while a library search chooses
only at resolution. [§1.14](#114-set-aside-and-announced-cards) owns all of these.

#### 3.1.2 Declaration or resolution

Some effects choose their object at declaration, others at resolution, and both can sit on the same card: {Break the
Bonds} fixes the uncontrolled vampire at [pre] on declaration, the locked minion at [obf] on resolution.[^3-1-6] Which
one applies decides what a target-protection or target-changing effect can reach.

An effect worded on the action being successful chooses its object then, e.g. {Keystone Kine}; so does a search that
need not name what it fetches, e.g. {Magic of the Smith}.[^3-1-7]

#### 3.1.3 Legal targets

An action cannot be announced at all if the card requires an object and no legal one exists.[^3-1-8] But the object need
only exist — it need not be able to yield the full effect. A vampire with no blood is a legal choice for a
blood-stealing action; at resolution you take what is there, e.g. {Siphon}, {Villein}.[^3-1-9]

If some of the intended objects cannot be targeted, the announcement stands; partial application is
[§3.4](#34-resolution-success-and-effect)'s.[^3-1-10]

An effect granting an action or a block does not lift the normal prey, predator and target restrictions.[^3-1-11]

#### 3.1.4 Choosing is targeting

A directed action that chooses, selects or names a minion targets it: protection effects bar it and cost increases
apply, e.g. {Spirit Marionette} against {Secure Haven}, and a master card that chooses a vampire pays its extra pool,
e.g. {The Rack}. The same test decides whether a frenzy card is played "on" a minion; one boosting its own player's side
is played on the opposing minion, so a {Blade of Enoch} bearer may play {Dam the Heart's River}. An action targeting a
retainer, equipment or other card on a minion does not target the minion, and neither does a card elsewhere that merely
named him.[^3-1-12]

An effect that resolves outside an action targets nothing: an unlock-phase choice reaches a minion under {Secure Haven},
e.g. {Baltimore Purge}.[^3-1-13]

#### 3.1.5 Directed actions

An action directed at a card controlled by another Methuselah is directed at that Methuselah; it is undirected if you
control the target.[^3-1-14] So an action against an equipment, retainer or location is directed at the Methuselah
controlling it, not at the minion carrying it, e.g. {Conceal}, {Haunt}.[^3-1-15]

{Abbot} grants intercept during actions directed at its controller, so it applies to those; {Detect Authority} requires
an action directed at a minion or location, so it does not. An ability keyed to directed actions never fires against
actions by its own controller's minions.[^3-1-16]

A Methuselah's hand, library, crypt and uncontrolled region count as his for this test: an action that targets them, or
cards in them, is directed at him whether or not its card is marked directed, e.g. {Petra Resonance}, directed at every
other Methuselah.[^3-1-20]

Acting on another Methuselah's ash heap is not a directed action — the ash heap is not a card in play, e.g. {Daemonic
Possession}. Neither is a referendum, so {Secure Haven} gives no protection against a vampire being chosen by one, e.g.
{Banishment}.[^3-1-17]

An action with several targets is directed at every Methuselah controlling one of them, e.g. {Sowing
Dissension}.[^3-1-18] A reactive effect keyed to being targeted applies if any one target of the action would enable it,
e.g. {Chanjelin Ward}; where such an action's effect then names a single minion, that minion must be controlled by one
of the target Methuselahs, e.g. {Weigh the Heart}.[^3-1-19]

### 3.2 Stealth and intercept

**Base rules.** Stealth may be added only while the action is being blocked by a minion with enough intercept to stop
it; intercept only by a blocking minion whose intercept is below the acting minion's stealth. [[RBK
stealth-and-intercept]](https://www.vekn.net/rulebook#stealth-and-intercept)

#### 3.2.1 Adding what is not needed

A card that adds stealth or intercept cannot be played when the value is not needed, e.g. {Bonding}, {Mask of a Thousand
Faces} at [OBF]. The restriction reaches an ability that adds it as well, e.g. {Osric Vladislav}.[^3-2-1] See
[§1.6](#16-requirements) for the rule; this is its main instance. An ability whose stealth depends on a random outcome
may be used before stealth is needed; a stealth result is then not gained, e.g. {Gathii}.[^3-2-12]

Need is a threshold, not a cap. Once some intercept is needed the minion may take more than it needs, e.g. {Angelica,
The Canonicus}. Printed text waives the restriction where it says so — "even if intercept is not yet needed", e.g.
{Phased Motion Detector}, {Under Siege}.[^3-2-2]

#### 3.2.2 Reducing the opposing value

Reducing is always legal, including against a minion already at 0 or negative, where it simply has no effect, e.g.
{Draba}, {Night Terrors}. Such a card needs no block attempt under way. Reducing stealth "to 0" is a -X stealth modifier
on the action, X fixed at the time the card is played; stealth added afterwards is added to the reduced value.[^3-2-3]

#### 3.2.3 Duration and re-evaluation

Modifications last for the duration of the action [[RBK
stealth-and-intercept]](https://www.vekn.net/rulebook#stealth-and-intercept). An option to add, granted during the
action, likewise stays open until resolution, e.g. {Marciana Giovanni, Investigator}.[^3-2-5]

Having intercept is not permission to block. {Marciana Giovanni, Investigator} may lock to grant herself the option, and
may do so from torpor, but still needs a wake to attempt the block.[^3-2-6]

A bonus conditioned on the action's target is satisfied by any one qualifying target, e.g. {Talley, The Hound};
"directed at their controller" covers actions directed at cards that controller owns, e.g. {Abbot}. A bonus conditioned
on a trait is re-read if the trait changes mid-action, e.g. {Ministry}.[^3-2-7] Intercept "against" a class of minion
reads the acting minion: {Carlton Van Wyk}'s +1 intercept against vampires does not apply to an ally's action
*(reading)*.[^3-2-11]

#### 3.2.4 Who plays it, when, and how often

The window does not close when one side passes. Intercept may still be added after the acting Methuselah declines to add
stealth, e.g. {Guardian Vigil}, and playing a modifier that changes neither value leaves it open, e.g. {Inspire
Greatness}. Sequencing puts the targeted Methuselah's declaration first. Failing to block for want of intercept is still
declining to block.[^3-2-9]

How often a stealth or intercept effect may be used is fixed by its own wording: once per action, e.g. {Phased Motion
Detector}; repeatable, e.g. {Matteus, Flesh Sculptor}; once per unlock where locking is the cost, e.g. {Starshell
Grenade Launcher}. See [§1.2](#12-wording-templates-periodicity-duration-triggers) for the wording templates.[^3-2-10]

### 3.3 Blocking

#### 3.3.1 Who may block

An effect that lets a minion ignore the prey, predator or target restriction reaches only that check. Every other
condition on blocking still applies, including a prior decision not to block, e.g. {Eagle's Sight} [AUS],
{Anneke}.[^3-3-1] It covers one block attempt; if the action later continues as if unblocked, another is needed.[^3-3-2]
Blocking "as an ally" is the same shape. A block attempt that becomes illegal while under way fails, and that minion
cannot attempt again this action: {Daring the Dawn} [FOR] played during a vampire's attempt fails it, while an ally's
attempt stands *(reading)*.[^3-3-24]

The unlock-to-block template ({Guard Duty}, {Second Tradition: Domain}):

- Unusable when that vampire could not attempt to block the action anyway ({Daring the Dawn}).[^3-3-3]
- Unlocking to block does not lift the prey, predator or target restriction.[^3-3-4]
- Usable by a vampire already attempting to block, including a locked one that woke earlier.[^3-3-5]
- If the unlock leaves the action with no suitable target ({Ambush}), the action does not end
  ([§3.4](#34-resolution-success-and-effect)).

#### 3.3.2 Declaring and declining

Sequencing governs the block window: a Methuselah who is not the target must wait until the targeted Methuselah declares
whether he attempts or declines to block. Using a card in play is bound the same as playing one from hand, e.g. {Hide
the Heart} [VAL], {Draba}.[^3-3-7]

Playing a card worded "after blocks are declined" declines the block for its controller, and a block attempt that
already failed against stealth counts as declining for those cards, e.g. {Archon Investigation}, {Deflection}.[^3-3-8]
Waking or unlocking to block compels nothing: the vampire need not attempt to block nor play further reaction
cards.[^3-3-9]

A declared attempt cannot be withdrawn — the minion is still attempting to block after a stealth modifier lands
({Faceless Night} [OBF]).[^3-3-10] An effect imposing a blood cost to attempt is the exception: a minion already
attempting may resign instead of paying ({Tenebrous Form} [OBT]).[^3-3-11]

Declining belongs to the Methuselah, not to the minion. Being made to enter combat with the acting minion is not
blocking ({Brujah Frenzy}); substituting another minion into the block-induced combat leaves the block standing, so the
blocker still locks.[^3-3-14]

#### 3.3.3 The cost of blocking

"Must burn 1 blood to attempt to block" is a cost, so a minion who cannot pay cannot attempt, e.g. {Dominion},
{Tenebrous Form} ([§5.5](#55-allies) for allies).[^3-3-15] "Vampires blocking this vampire burn 1 blood" is an automatic
effect: an empty vampire still blocks and burns what it can, e.g. {Archon}, {Dónal O'Connor}. That burn happens
immediately on the successful block, whether or not combat follows, and before any card or effect may be used: {Truth of
Blood}'s discard is completed before the blocker can answer, {Obedience} included.[^3-3-16]

A toll on attempting to block is owed only by an attempt. When the card that would force the attempt is canceled as
played, no attempt happens and the toll is not paid, e.g. {Parijat, the Dark Oracle} and a canceled {Second Tradition:
Domain}; the canceled card's own cost is still paid ([§1.8.3](#183-a-canceled-card-has-still-been-played)). A card
forcing an attempt whose toll cannot be paid cannot be played: under Parijat, Second Tradition cannot be played from an
empty library.[^3-3-26]

A restriction triggered by attempting to block runs for the duration the card prints, and a failed attempt does not cut
it short ({Blessing of Chaos} [dem]). A cost worded as a state — "while attempting to block" — stops once the attempt
has failed, so it never reaches the action card's cost ({Libertas}).[^3-3-17]

#### 3.3.4 After the block succeeds

Locking the blocker and entering combat are the two simultaneous consequences of block resolution, not of the attempt
succeeding. Cards timed "before block resolution" fill that window ({Change of Target}, {Angel of Berlin}). An effect
that ends the action or cancels the block there prevents both consequences ({Clan Loyalty}); canceling only the combat
does not.[^3-3-18] A blocker leaving play before block resolution leaves the action blocked with no combat, e.g. when
the pool {Aelswith, The Irresistible} makes its Methuselah burn ousts him *(reading)*.[^3-3-25]

Effects triggered by the block fire once the block would be successful; ending the action afterwards does not undo them
({Aching Beauty}, {Unleash Hell's Fury}). An effect that makes the action resolve as if unblocked instead means no block
was successful, so effects conditioned on a successful block never fire ({Lesser Boon}).[^3-3-19] An effect keyed on a
block does nothing when none occurred ({Spiritual Protector}).[^3-3-20]

Delayed effects triggered by the block are [§2.4.3](#243-delayed-effects-triggered-by-a-block)'s. An effect granted for
blocking is scoped to the block-induced combat ([§3.6.3](#363-effects-scoped-to-one-block)). Methuselahs who had already
declined get no fresh opportunity when the action continues as if unblocked.[^3-3-23]

### 3.4 Resolution, success and effect

**Base rules.** An action that survives all block attempts is successful; its cost is then paid and its effects take
place. [[RBK resolve-the-action]](https://www.vekn.net/rulebook#resolve-the-action)

#### 3.4.1 A successful action need not do anything

An unblocked action is successful even when its effect cannot happen, and its cost is then not paid, unless what fails
is the cost itself ([§1.7.5](#175-paying-what-you-can)).[^3-4-1] The usual causes are a target that no longer qualifies
at resolution, e.g. {Ambush} when the target unlocks, and a target removed or protected by another effect, e.g. {Warsaw
Station}. A stolen target keeps the action pointed at it; the action ends with no effect only if that target has become
illegal, e.g. {Temptation}.

An action provided by an equipment the acting minion no longer bears at resolution ends with no effect, though still
successful if unblocked, e.g. {Jar of Skin Eaters}.[^3-4-16]

The target ceasing to qualify does not end the action: block attempts and any resulting combat happen as normal, e.g.
{Coterie Tactics}.[^3-4-3]

Riders conditioned on the action being successful fire on such an action. {Forced March} and {Instantaneous
Transformation} unlock; {Perfectionist} gains the blood; {Hrothulf} burns the Edge.[^3-4-4]

Where a card names the effect rather than the success, the effect must actually occur. {Abactor} calls its blood hunt
only after "successful resolution", meaning unblocked *and* blood gained (for a hunt at full capacity,
[§3.7.2.1](#3721-a-successful-hunt-and-a-successful-hunt-action-are-different-conditions)).[^3-4-5] Whether the hunt or
bleed itself succeeded is [§3.7.1.2](#3712-reduction-and-pool-burned)'s and [§3.7.2](#372-hunt)'s.

An effect that cannot apply to part of what it names applies to the rest, e.g. {Edged Illusion} against minions another
effect protects, or {Enticement} when the Edge is gone before resolution.[^3-4-7]

An action card voided as it is played by a standing "no effect" bar produces no action at all — this is not a
cancellation, but the card counts as played, the vampire does not lock, and may act again, the same action included,
e.g. under {Veil of Darkness}.[^3-4-8]

#### 3.4.2 What is read at resolution

The action's terms are named at announcement; resolution supplies the rest. A choice the wording embeds in the effect
rather than in the terms is made at resolution, e.g. the minion taking {The Platinum Protocol}'s corruption counter; an
option is exercised then — moving blood onto the progeny or not ({Third Tradition: Progeny}); a state is read then, e.g.
the disciplines copied by {Dual Form}.[^3-4-9] A library search provided by an action happens only once the action
succeeds, e.g. {Magic of the Smith}; a failed action's additional cost is unpaid.[^3-4-10]

An effect keyed to the action *reaching* resolution fires even where the action does nothing, or ends before combat,
e.g. {Daring the Dawn}.[^3-4-11]

An action that puts its acting minion into combat resolves first; "after action resolution" effects wait for the combat,
e.g. {Strix}.[^3-4-12]

#### 3.4.3 The "would be successful" window

Once all blocks are declined, a card usable "when the action/bleed would be successful" may be played, e.g. {Spying
Mission}, {Andre LeRoux}, {Dis Pater}. After one is played the action is about to resolve: only other cards with that
same timing may follow.[^3-4-13] Anything that changes or defeats the action — {Deflection}, {Archon Investigation},
{Telepathic Counter} — must be played before that point.[^3-4-14] The post-resolution windows sequence the same way.

A bleed reduced to zero is not successful ([§3.7.1.2](#3712-reduction-and-pool-burned)), so a "would be successful"
card cannot then be used on it.[^3-4-15]

### 3.5 Action repetition (NRA) and canceled actions

**Base rules.** A minion cannot perform an action with the same action card, from hand or in play, more than once each
turn, even if it unlocks; a minion cannot bleed more than once each turn. [[RBK
action-card-or-card-in-play]](https://www.vekn.net/rulebook#action-card-or-card-in-play) [[RBK
bleed]](https://www.vekn.net/rulebook#bleed)

#### 3.5.1 Reached resolution

The No Repeated Action rule bites once the action reaches resolution. An action that fails, or that is ended by a card
or an ability, has reached resolution: the same minion cannot attempt it again this turn, e.g. {Change of Target},
{Champion}, {Krassimir}.[^3-5-1] Being blocked and failing still counts as the attempt. An action continued after a
reaction made it fail still fails.

A canceled action has not reached resolution. It is not performed, its cost is not paid, and the same minion may attempt
the same action again.[^3-5-2] It may not if the action was blocked and then continued before the cancellation.[^3-5-3]
Canceled as its card is played, the acting minion does not lock; canceled later, he stays locked unless the canceling
card unlocks him: {Tangle Atropos' Hand} does, {The Kiss of Ra} does not.[^3-5-26]

An action canceled after its card was played was nonetheless taken, so it counts against effects that count actions,
e.g. {Enkil Cog}, {Veil the Legions}. One canceled as its card is played was not: the next action is still the first for
{Channel 10}.[^3-5-4]

Cancellation undoes the cost of the action: an {Enrage}d vampire canceled mid-action does not burn 2 blood.[^3-5-5]
Failure does not — {Mobile HQ, Operation Antigen} stays locked when the action it paid for fails. A minion locked by a
card in play to take an action canceled as it is played stays locked, and may act again or let the effect be
lost.[^3-5-6]

#### 3.5.2 "The same action"

{Change of Target}, {Obedience} and {Red Herring} forbid repeating "the same action". Two actions are compared as
follows.

- Rulebook actions (bleed, hunt, equip, …) are never the same as an action from a card, played or in play.[^3-5-7]
- Two rulebook actions differ if their targets differ. Equipping from a minion targets the equipment taken: if any one
  of them is the same, the action is the same.[^3-5-8]
- Actions granted by two different cards in play are not the same, even if the cards share a name.[^3-5-9]
- Substituting the acting minion only in combat does not change who took it, e.g. {Malleable Visage}.[^3-5-10] Handing
  the action itself to another minion does: the new acting minion has taken it, e.g. {Mask of a Thousand Faces}
  *(reading)*.[^3-5-29]
- A minion that leaves play and returns is a new minion, bound by nothing its previous incarnation did.[^3-5-11]
- Putting a card into play, or being made to recruit or employ out of turn, is not performing that card's action, so the
  normal action is still available, e.g. {Piper}, {Muricia's Call}, {Clandestine Contract}.[^3-5-12]

#### 3.5.3 Ended and failed

An action that **ends** ends immediately. No further action modifier or reaction can be played and no other effect
intervenes, e.g. {Mirror Walk}, {Tangle Atropos' Hand}.[^3-5-13] Where the effect starts a combat instead, combat begins
immediately with the same window closed, e.g. {Brujah Frenzy}.[^3-5-14]

An action that **fails** also ends, but modifiers and reactions usable at the end of, or after, an action can still be
played, e.g. {Detect Authority}, {Mistaken Identity}.[^3-5-15] A failed action was not blocked, so cards requiring a
block cannot be played, e.g. {Freak Drive}, {Mirror Image}.[^3-5-16] Conversely, a card that makes the action fail
cannot be played while a block is being attempted or once one has succeeded, e.g. {Faerie Wards} [myt], {Hide the Heart}
[val].[^3-5-27]

Effects scheduled for resolution or later are lost if the action ends first: the lock-and-enter-combat of {Deep Song},
the after-resolution combat of {Siren's Lure}.[^3-5-17] Likewise, an action whose acting minion is not ready when it
would resolve simply ends — [§3.7.5.1](#3751-the-referendum-is-part-of-resolution) owns the referendum case.

Damage a card inflicts on its own acting vampire is keyed to the action reaching resolution
([§3.4](#34-resolution-success-and-effect)); it does not apply if the action was canceled, e.g. {Daring the Dawn},
{Force of Will}.[^3-5-18]

#### 3.5.4 Canceled blocks and canceled combats

A canceled block attempt is neither successful nor unsuccessful. It is simply canceled.[^3-5-20] Nor is it a decline:
the minion may attempt again, e.g. after {Dawn Operation} [for] *(reading)*.[^3-5-28]

Canceling the combat that a successful block created leaves the action blocked. The blocking minion stays locked, {Mask
of a Thousand Faces} cannot take the action over, and reaction cards pertaining to that combat cannot be played
afterwards, e.g. {Venenation}, {Bear-Baiting}.[^3-5-21] An effect that continues the action as if unblocked still works,
e.g. {Crypt's Sons}.

#### 3.5.5 Scope of cancellation

A canceled card is still played, so a cancellation does not give back the play.

Canceling a card that carried a choice returns the choice, and the same choice may be made again — a minion whose strike
is canceled must still choose a strike and may play the same card.[^3-5-24]

Cancel effects reach every action card, including employ retainer, recruit ally and political actions.[^3-5-25]

Multiple copies of one action modifier played by different minions on the same action are a stacking question, not a
repetition one.

### 3.6 Continuing an action as if unblocked

**Base rules.** A successful block locks the blocker and sends both minions into combat; the action card is set aside
out of play until the action resolves. Once a Methuselah declines to make further block attempts, that decision is
final. [[RBK summary-of-the-course-of-an-action]](https://www.vekn.net/rulebook#summary-of-the-course-of-an-action)

#### 3.6.1 What continuation restores

Continuation moves the action card from the ash heap, where it went when the action was blocked, back to limbo. If the
action card is not in the ash heap, the action cannot be continued.[^3-6-1]

All action modifiers and reaction cards played earlier in the action remain in effect, including those that pertained to
the successful block.[^3-6-2]

Methuselahs who had already declined to block get no fresh opportunity. A Methuselah who had not yet passed may still
attempt to block, and a minion who blocked may block again.[^3-6-3]

An effect that penalises a minion for not blocking ignores blocks made before it was played, e.g. {Forced Awakening},
{WMRH Talk Radio}.[^3-6-4]

A card whose effect is one block attempt is spent by that attempt; another copy is needed to attempt the second block,
e.g. {Eagle's Sight}, {Falcon's Eye}. Effects that apply for the duration of the action are not spent this
way.[^3-6-5]

#### 3.6.2 When the action cannot be continued

Only a blocked action can continue as if unblocked: not after a combat that resulted from the action succeeding, e.g.
{Bum's Rush}.[^3-6-6]

Continuation is an "after combat ends" effect: if the combat continues or a new combat begins, it is lost
([§4.9.2](#492-what-survives-a-continued-combat)).[^3-6-7]

A combat queued during the action is still part of that action ([§4.9.3](#493-queuing-a-new-combat)), so the action can
be continued from its last combat, e.g. {Form of Mist} played in the combat {Psyche!} queued. It cannot be continued
while a further combat is queued.[^3-6-8]

No combat need have occurred: a canceled combat leaves the action blocked
([§3.5](#35-action-repetition-nra-and-canceled-actions)) and it can still be continued, e.g. {Crypt's Sons}, {Momentary
Delay}.[^3-6-9]

#### 3.6.3 Effects scoped to one block

An effect a reaction card provides on a successful block applies only in the combat that block induced. It does not
carry over to a follow-up combat, e.g. one queued by {Psyche!}. This holds whatever the effect is — a maneuver, damage
prevention, a strike or a strike restriction, e.g. {Guard Dogs}, {Precognition}, {Night Terrors}.[^3-6-10]

If the action continues and a minion blocks again, that effect is provided afresh.[^3-6-11]

### 3.7 The action types in detail

The rulebook actions, one subsection each. Rules that are not specific to an action type live in
[§3.1](#31-announcement-and-targets)-[§3.6](#36-continuing-an-action-as-if-unblocked) and
[§3.8](#38-actions-provided-by-cards)-[§3.10](#310-changing-the-acting-minion).

#### 3.7.1 Bleed

**Base rules.** A bleed's default amount is 1 and its default target is your prey; if the action succeeds with an amount
of 1 or more the bleed is successful, the target burns that much pool, and the acting minion's controller takes the
Edge. [[RBK bleed]](https://www.vekn.net/rulebook#bleed) The "(limited)" rule caps increases only: one limited action
modifier may increase the bleed amount per action, while modifiers printed as not counting against the limit, and all
reductions, are unrestricted.

##### 3.7.1.1 Bleed amount

The limit counts an increase only while that increase is actually applying.[^3-7-1-2] A conditional limited modifier
played when its condition is not met does not count against the limit and lingers, applying if the condition is later
met — {Foreshadowing Destruction} at [DOM] against a target with 10 or more pool. A bonus lost during the action stops
counting and a further modifier may be played; redirection to another target is the usual cause.

"Only usable during a bleed action" bars the play in every other action, e.g. {Confusion}. Where a card in play instead
grants a bonus "during a bleed action", the bonus reaches only the minion performing that bleed, e.g. {Haqim's Law:
Retribution}, {Anu Diptinatpa}.[^3-7-1-4]

A vampire's "current bleed" read outside a bleed action is the amount he would bleed his prey with, e.g. {Justicar
Retribution}.[^3-7-1-5]

##### 3.7.1.2 Reduction and pool burned

A bleed reduced to zero is not successful. An effect keyed to a successful bleed does not fire either, e.g. {Protected
Resources} is not burned by a bleed that resolves at 0.[^3-7-1-6]

An effect capping the pool a target burns does not change the bleed amount; it applies after every modifier, e.g.
{Protected Resources}.[^3-7-1-7]

##### 3.7.1.3 Redirection

Redirection changes the target of the bleed, so the new target is the target for every purpose: a reaction keyed to a
bleed against you is usable by whoever the bleed ends up against, not only by the Methuselah it was declared against,
e.g. {Elder Intervention}.[^3-7-1-8] Block attempts reopen. [[RBK
summary-of-the-course-of-an-action]](https://www.vekn.net/rulebook#summary-of-the-course-of-an-action)

The new Methuselah must be a legal target for that bleed. A restriction on whom a minion may bleed also bars redirecting
that minion's bleed to the protected Methuselah, e.g. {Minor Boon}.[^3-7-1-9]

Redirection affects the acting vampire, so an effect stopping reaction cards from affecting him stops {Deflection}, e.g.
{Perfect Clarity}.[^3-7-1-10]

Mandatory bleeds are [§3.9](#39-mandatory-actions)'s.

#### 3.7.2 Hunt

**Base rules.** The rulebook hunt action is undirected, at +1 stealth, and gains the acting vampire blood from the blood
bank equal to their hunt amount, by default 1 [[RBK hunt]](https://www.vekn.net/rulebook#hunt). Cards also grant hunt
actions of their own.

##### 3.7.2.1 A successful hunt and a successful hunt action are different conditions

A hunt action that resolves unblocked is successful even when no blood is gained. The hunt is not: with zero blood
gained the hunt is not successful, so effects keyed to a vampire "successfully hunting" do not fire, e.g. {Hunger Moon}.
A card that lets a vampire escape an effect by hunting asks only for the hunt action: a blocked hunt, or one gaining no
blood, still spares him, e.g. {Thirst}.[^3-7-2-1] A hunt at full capacity is a successful hunt, e.g. for {Harvest Rites}
*(reading)*.[^3-7-2-11]

Each card names one of the two: {Triole's Revenge} burns a Ventrue who hunts at full capacity: it names the hunt action,
so it fires whatever the hunt gains, even nothing. It does not fire if the action was blocked.[^3-7-2-2]

See [§3.4](#34-resolution-success-and-effect) for the general rule that an unblocked action is successful even when its
effect cannot occur.

##### 3.7.2.2 Card-provided hunt actions

A hunt action granted by a card is a hunt action for every purpose. It takes hunt bonuses and increases to hunt value,
e.g. {Kyoko Shinsegawa}. A restriction or permission worded around hunting reaches it: {Pariah}, who cannot take
undirected actions other than hunting, may take a card-provided undirected hunt action such as {Vulture's Buffet}. Where
a vampire has more than one granted hunt action, his controller chooses among them.[^3-7-2-3]

An effect worded against the normal hunt reaches only the rulebook hunt action. A card removing "the normal +1 stealth
when hunting" does not touch the stealth printed on a card-provided hunt action, e.g. {Igo the Hungry} against {Loki's
Gift}.[^3-7-2-4]

Where the hunt action takes blood from a target, blood added by a hunt bonus comes from that target, not from the blood
bank.[^3-7-2-5] A bonus that grants blood from the blood bank is a separate gain, not hunt blood, and comes from the
bank as printed, e.g. {Festivo dello Estinto}.

##### 3.7.2.3 Hunts that steal from a vampire

A card granting a hunt that steals from a named vampire licenses that targeting for its own hunt action only. Its bearer
taking a different special hunt action cannot name a vampire, e.g. a {Legacy of Caine} bearer taking
{Abactor}.[^3-7-2-6]

Blood the hunting vampire cannot gain goes to the blood bank rather than staying with the target.[^3-7-2-8] A bonus that
is not the hunt amount does not scale with hunt modifiers: {Loki's Gift} [PRO] burns 1 blood from any vampire whatever
the hunt value.[^3-7-2-9]

##### 3.7.2.4 Full capacity and mandatory hunts

A vampire at full capacity may still hunt; he gains the blood, and the excess returns to the blood bank as normal. An
effect that moves blood the hunt gained takes it before that return, e.g. {Rabbat, The Sewer Goddess}. A card offering
an alternative satisfies a mandatory hunt, e.g. {Undying Thirst} permits diablerie instead.[^3-7-2-10] See
[§3.9](#39-mandatory-actions) for mandatory actions.

#### 3.7.3 Equip actions

**Base rules.** Equipping from hand is an undirected +1 stealth action and pays the card's cost; equipping from another
minion you control is free, and the equipment taken is announced with the action. [[RBK
equip]](https://www.vekn.net/rulebook#equip)

##### 3.7.3.1 What counts as an equip action

A card that has a minion attach an equipment provides an equip action, whatever verb it prints, e.g. {Magic of the
Smith}, {Bloodstone}.[^3-7-3-1] Effects that modify equip actions therefore reach it, e.g. {The British Museum, London}.
One card may provide an equip action or an employ action depending on what it fetches, e.g. {Jack of Both Sides}.

Being an equip action does not settle requirements and cost. Only the wording "put … in play" bypasses them: {Magic of
the Smith} prints that requirements and cost apply.

An equip granted by another card's effect is still an equip, so an ability reading "when this minion equips" fires, e.g.
{Topaz} with {Concealed Weapon} or {Pier 13, Port of Baltimore}.[^3-7-3-2] An equipment's own on-equip clause fires the
same way, e.g. {Dagger}'s second copy. An ability keyed to an equip *action* is not satisfied by an equip that is not an
action, and {Pier 13, Port of Baltimore} prints that it is not an action.

Being put into play is not equipping. A clause conditioned on the card being equipped does not fire when an effect
merely places it, e.g. {Incriminating Videotape} chooses no minion when put on a vampire by {Alastor}, but does when
equipped via {Magic of the Smith}.[^3-7-3-3]

##### 3.7.3.2 Equipment already in play

Moving equipment between minions is neither playing nor equipping it, so a restriction worded against playing or
equipping does not reach the move, e.g. {Heidelberg Castle, Germany}; a card may be moved onto a vampire in
torpor.[^3-7-3-4] A card already played or equipped can still enter play through an effect that is not an equip, e.g.
{Kiss of Lachesis}. Another of your minions may instead take it with an equip action from its holder,[^3-7-3-12] a
holder in torpor included *(reading)*;[^3-7-3-10] that rulebook action is not a card play, so no requirement is checked
— a clan equipment may be taken by a minion who does not meet its requirements.[^3-7-3-9] Its text naming the bearer by
clan or sect reads the current bearer, so such a taker gets none of it, e.g. a non-Ravnos holding {Treasured Samadji}
*(reading)*.[^3-7-3-11]

A locquipment's immunity to equipment movers and equipment prohibitions is
[§1.3](#13-card-types-and-multi-type-cards)'s.

One equip action may take several equipment from another friendly minion. It is an action to equip with each of them, so
each one's own bonus for the action to equip it applies, e.g. {Unlicensed Taxicab}'s +1 stealth, and a card keyed to an
equip action reaches the take, e.g. {Gift of Bellona}, {Hag's Wrinkles}.[^3-7-3-6]

Whether two equip-from-minion actions are the same for the No Repeated Action rule is
[§3.5](#35-action-repetition-nra-and-canceled-actions)'s.

##### 3.7.3.3 Minions who cannot equip

A vampire who cannot have equipment cannot attempt an equip action at all, e.g. {Enkidu, The Noah}, {Beast, The
Leatherface of Detroit}, and cannot strike to steal one.[^3-7-3-8] The bar does not reach locquipments: Enkidu may still
equip with one.

#### 3.7.4 Employ retainer and recruit ally

##### 3.7.4.1 The action

The cost of an employ retainer or recruit ally action is the cost printed on the retainer or ally, and its requirements
must be met to take the action.[^3-7-4-1]

A recruit, employ or equip action of a card from hand requiring X is an action requiring X, e.g. {CrimethInc.} after
recruiting from hand an ally requiring an Anarch, {Keegan}, {"Mother" Anja Giovanni} *(reading)*.[^3-7-4-9] A card
another card fetches or plays is not ([§1.6.3](#163-cards-entering-play-abnormally)).

A minion whose text says it cannot have equipment or retainers can neither take the action nor steal one with a strike,
e.g. {Beast, The Leatherface of Detroit}.

Employ retainer and recruit ally are ordinary actions for cancellation.

Cost reductions read the card as announced. {Ghouled} makes the ally a ghoul only on resolution, so a reduction for
recruiting a mortal still applies.

##### 3.7.4.2 Entry into play other than by the action

Two templates bring an ally or retainer in without the action: a card under which a minion *recruits* or *employs*
outside an action, e.g. {Piper}, {Pack Alpha}; and a card that *puts* the card *in play*, e.g. {Summon History} —
[§1.6](#16-requirements) governs requirements and cost under each.[^3-7-4-3]

Neither is an action, and effects worded as attaching to an action do not attach:

- Cost reductions naming an action, e.g. {Charisma}, {Erichtho}.[^3-7-4-4]
- Abilities triggered by announcing the action, e.g. {Zhenga} does not work with {Piper}.

Effects keyed to the recruit or employ itself, or to the cost of the card, do attach however the minion arrived, e.g.
{Soul of the Earth}, {Little Tailor of Prague}.[^3-7-4-5]

Cancellation of the fetched card is [§1.8.2](#182-played-but-not-in-the-normal-fashion)'s.

Where a minion recruits or employs, its own disciplines are read as normal to set the card's level; where the card is
merely put in play, use the basic version.

The action card of a normally recruited ally is replaced before the ally enters play. A card brought in by an effect is
not, so an ability triggering on entry has only the hand and ash heap as they stand, e.g. {Corrupt
Construction}.[^3-7-4-8]

Whether the new minion may *act* that turn is [§5.5.4](#554-entering-play-and-acting)'s for allies and
[§5.6.5](#565-entering-play)'s for retainers.

#### 3.7.5 Referendum procedure

##### 3.7.5.1 The referendum is part of resolution

The referendum is part of the action's resolution, not something that follows it. Cards played "after a successful
action" or "after resolution" wait until the referendum is concluded, e.g. {Heidelberg Castle, Germany},
{Sargon}.[^3-7-5-1] Among those cards the order is free, but none of them can forestall an oust the referendum itself
causes.[^3-7-5-2]

The referendum is conducted only if the acting minion is still ready when the action would resolve. If combat leaves him
no longer ready — in torpor or out of play — no referendum happens, e.g. {Yawp Court}. That action was successful, and
its cost was paid before the location was locked.[^3-7-5-3]

Action modifiers and reaction cards may be played during the referendum and after it, and vote-affecting cards before,
during or after votes and ballots are cast. An effect keyed to the tally reaches votes already cast: {Scorn of Adonis}
makes every Methuselah who voted against burn pool, whenever they voted.[^3-7-5-4]

##### 3.7.5.2 Terms

Terms are chosen once the action is successful, before polling. A card usable "during the polling step" cannot be played
before terms are set, e.g. {Business Pressure}.[^3-7-5-5]

What the terms designate is fixed then: {Revolutionary Council}'s chosen Anarchs still count if they lock during the
vote, and pool moving during the referendum does not alter {Parity Shift}'s outcome. A quantity named in the effect
clause instead of the terms is counted when the effect applies — {Domain Challenge} counts locked minions after the
referendum completes.[^3-7-5-6] Where the terms cannot be carried out in full, the calling Methuselah chooses within
them. A cost inside the effect is paid by whoever the terms name ([§1.7.6](#176-whose-cost-and-paid-in-what)).[^3-7-5-7]
The terms must have a legal choice when the action is announced, though they are chosen only on success: {Parity Shift}
cannot be played when no Methuselah has more pool than its caller.[^3-7-5-20] If none remains when the terms are chosen,
the action fizzles and no referendum is held, so an effect keyed to the referendum failing does not apply
*(reading)*.[^3-7-5-21] A referendum may be called with no eligible subject in play, e.g. {Peace Treaty} with no weapons
out.[^3-7-5-8]

##### 3.7.5.3 Polling and the tally

The polling step opens with the cards usable "before votes and ballots are cast", in the normal sequence; no vote or
ballot is cast until every Methuselah has passed in that stage, e.g. {Closed Session}.[^3-7-5-22]

Where the effect lets Methuselahs burn pool or blood for votes, the amounts are decided during that effect's resolution,
and sequencing does not lock them in: there is give and take, each Methuselah free to burn one at a time and wait to see
what the others do, e.g. {Business Pressure}, {Mob Rule}.[^3-7-5-9] The ability does not persist for the rest of the
action.

An ability worded "during a referendum" is usable once per referendum; unlocking mid-vote does not restore the use, e.g.
{Michael Luther}, {Ellison Humboldt}.[^3-7-5-10] An ability keyed to the tally fires when votes are tallied, e.g.
{Astrid Thomas}.[^3-7-5-11]

A cost keyed to "when the votes are tallied" is paid at the tally, before the referendum's effects; a cost keyed only to
the referendum passing is paid after all of them, including any oust and any pool gained. {Treachery} prints both
clauses.[^3-7-5-12] So an effect keyed to a passed referendum still applies when the predator is ousted, e.g. {Donald
Cargill}, but not when its own controller is ousted, e.g. {Lutz von Hohenzollern}.[^3-7-5-13]

##### 3.7.5.4 Automatic pass, cancellation, failure

A referendum that passes automatically has no polling step. No "during a political action", "during a referendum" or
"during the polling step" effect may be played or used — a vampire's ability included, e.g. {Aisha az-Zahra} — and no
card may be burned for a vote during it.[^3-7-5-14] Effects reading the number of votes it passed by have no effect. A
card usable *after* the referendum may still be used, with X = 0, e.g. {Voter Captivation} after a {Cryptic Rider}
referendum.[^3-7-5-15]

Canceling the referendum stops the tally, so cards keyed to the results do nothing, e.g. {Scorn of Adonis} under
{Delaying Tactics}.[^3-7-5-16] A referendum that fails produces none of the calling card's effects, including a
self-inflicted part, e.g. {Aura of Invincibility}.[^3-7-5-17]

The political action card is set aside out of play until the action resolves, so effects reading cards in play or in an
ash heap do not see it, e.g. {Luna Giovanni}.[^3-7-5-18] The blood hunt referendum is not a political action and is not
called by a vampire, so "during a political action" effects and effects keyed to the calling vampire do not apply, e.g.
{Power Structure}.[^3-7-5-19]

#### 3.7.6 Votes and ballots

##### 3.7.6.1 Votes are not ballots

An effect worded on "votes" never reaches ballots or the prisci sub-referendum. That holds whether the effect reduces a
title's vote value, e.g. {Rastacourere}, subtracts votes outright, e.g. {Condemnation: Mute}, or counts the votes a
vampire has, e.g. {Leadership Vacuum}, {Island of Yiaros}.[^3-7-6-1]

An effect that acts on the vampire rather than on its votes reaches both. Forcing a vampire to abstain removes its
ballot from the sub-referendum as well as its votes, and cancels votes and ballots it has already cast, e.g. {Arishat},
{Kateline Nadasdy}.[^3-7-6-2]

What a stripped title leaves behind — printed ballots, the referendum structures it defines — is
[§5.8.2](#582-off-sect-and-off-clan-titles)'s; a printed bonus ballot is cast on the same side as its vampire's votes,
e.g. {Gratiano}.[^3-7-6-3]

##### 3.7.6.2 Which votes count

An effect counting the votes a vampire has counts its title and any unconditional bonus votes from cards or abilities,
e.g. {Eze, The Demon Prince}, {Firebrand}. It does not count votes the vampire could obtain only by paying a cost or
meeting a condition, e.g. {Sundown}, {Aura of Invincibility}.[^3-7-6-4] A vampire given -X votes can go below zero;
the deficit offsets votes he gains later but subtracts nothing at the tally: a vampire at -3 casts 1 vote after
{Bewitching Oration} at [PRE]. A reduction of what a title is worth stops at zero.[^3-7-6-11]

Multiple copies of a vote-granting card do not multiply votes, and an untitled vote bonus adds to a title gained later;
see [§1.15](#115-cumulative-and-stacking-effects) for both.

##### 3.7.6.3 Changing votes, canceling votes, abstaining

Where two effects alter the same vampire's votes, the last one used governs, e.g. {Mustafa, The Heir}. An effect keyed
to the act of casting fires only on the first cast and not again when the votes are changed, e.g. {De Sade}.[^3-7-6-5] A
vote granted to a vampire that has already cast joins its cast on the same side, e.g. {Bewitching Oration}, {Elysium:
The Palace of Versailles} *(reading)*.[^3-7-6-12]

An effect keyed to how a vampire voted reads the disposition at tally. {Madrigal} does nothing if its player ends up
abstaining, forced or not. {Scorn of Adonis} reaches Methuselahs who voted "no" before it was played, and costs each of
them 1 pool however many "no" votes they cast. Votes granted only against the referendum are silenced, not recast, when
the bearer's votes are changed in favor, e.g. {Loyalist}. A vampire that acquires a disqualifying trait after casting
has its votes canceled and abstains, e.g. {Khay'tall, Snake of Eden}.[^3-7-6-6]

An effect that changes votes cannot compel an abstaining vampire to vote; it only restricts its choice if it votes, e.g.
{Kindred Coercion}, {Neferu}. A force-abstain effect may be used on a vampire that has not yet voted, keeping it
abstaining. A vampire that has not yet voted may always choose to abstain, escaping an effect worded on non-abstaining
vampires, e.g. {Astrid Thomas}; canceling that vampire's own votes does not undo votes others already cast with
it.[^3-7-6-7]

Polling-step timing, and referendums that pass automatically, are [§3.7.5](#375-referendum-procedure)'s.

##### 3.7.6.4 Locked and torpid vampires

A card that locks a vampire for voting locks it as it casts — a lock imposed as an effect, not a cost
([§5.2.4](#524-locking-the-locked-unlocking-the-unlocked)), e.g. {Disarming Presence}. A referendum ability printing no
lock cost is usable while locked, once per referendum ([§1.2.1](#121-how-often-an-effect-may-be-used)); an
unconditional one is usable in torpor as well. A vampire in torpor casts no votes, but an effect keyed to its
abstaining still applies, e.g. {Alvaro, The Scion of Angelica}.[^3-7-6-9]

##### 3.7.6.5 The political action card's own vote

A political action card provides its inherent vote whenever it is played, including when another card plays it from
hand, e.g. {Charming Lobby}, {Echo of Harmonies}. Burning a political action card during a referendum for a vote is not
playing it ([§1.11.1](#1111-retrieval-what-counts-as-played)), so a once-per-game play limit is untouched.[^3-7-6-10]

Pool paid, gained or lost as a consequence of a referendum, simultaneous ousts included, is
[§6.5](#65-pool-the-edge-and-ousting)'s; terms that cannot be carried out in full are [§3.7.5.2](#3752-terms)'s.

#### 3.7.7 Leaving torpor and rescuing from torpor

**Base rules.** Both default actions cost 2 blood, and the rescue cost may be paid by the acting vampire, by the rescued
vampire, or split between them. A blocked leave-torpor produces no combat: a vampire blocker may diablerise the acting
vampire instead, otherwise the action simply fails and no cost is paid.[^3-7-7-1]

##### 3.7.7.1 Cards that supply their own action

A card granting a leave-torpor or rescue action supplies its own action, not the rulebook one, so the default 2 blood is
not paid — e.g. {Resume the Coil}, {Rapid Healing}, {Healing Touch}. It is still a leave-torpor (or rescue) action for
anything keyed to one.[^3-7-7-2]

##### 3.7.7.2 A blocked leave-torpor

Leaving torpor produces no combat when blocked, however the action was supplied. Cards conditioned on entering combat
therefore cannot be used, e.g. {Ghoul Escort}.[^3-7-7-3]

The blocker's diablerie opportunity arises at block resolution. An effect that ends the action before block resolution,
or unlocks the acting vampire, denies it, e.g. {Change of Target}, {Mirror Walk} at [THA], {Blood Brother Ambush}.

##### 3.7.7.3 Cost and payer

The rescue cost belongs to the action, not to whoever pays it. A reduction printed on the acting vampire applies even
when the rescued vampire pays, e.g. {Frondator}.[^3-7-7-4] The reduction is mandatory and cannot be declined to keep the
higher cost. Who pays, and how a split is shared, is chosen as the action is announced. A share its payer can no longer
meet at resolution fizzles the action, that payer paying what he can; the other vampire does not pay it instead
*(reading)*.[^3-7-7-7]

The same split works the other way: an effect keyed to "an action costing blood" sees the rescue even when the acting
vampire paid none of it, e.g. {Cavalier}.

##### 3.7.7.4 Torpor and the action window

A vampire in torpor is not ready, but an effect granting extra action opportunities still reaches it. Its only available
action is leave torpor, e.g. {Madness Network}.[^3-7-7-5]

A card effect that moves the vampire out of torpor after a rescue or diablerie action is announced leaves that action
successful if unblocked, but with no effect. One that brings back a vampire sent to torpor during the action ends that
action instead, and it cannot be continued. Both, e.g. {Warsaw Station}.[^3-7-7-6]

#### 3.7.8 Diablerie and the blood hunt

**Base rules.** The steps of diablerie resolve as one unit and cannot be interrupted; effects are played before or
after, never between them [[RBK diablerie]](https://www.vekn.net/rulebook#diablerie). The blood hunt referendum is
automatic, is not an action, cannot be blocked, and takes no action modifiers or reaction cards [[RBK
the-blood-hunt]](https://www.vekn.net/rulebook#the-blood-hunt).

##### 3.7.8.1 The window before the blood hunt

There is a window after the diablerie resolves and before the blood hunt referendum, for the cards keyed to the
diablerie, e.g. {Ritual of the Bitter Rose}, {Slake the Thirst}.[^3-7-8-1] Other after-resolution effects follow the
blood hunt, e.g. {Freak Drive} *(reading)*.[^3-7-8-8] An action modifier keyed to the diablerie may come before or after
the discipline card, but must be played before the blood hunt, e.g. {Draught of the Soul}, {Ritual of the Bitter Rose}.

Where the diablerie happened decides which card types fit the window. A diablerie action takes action modifiers; an
{Amaranth} diablerie happens in combat and takes combat cards. {Slake the Thirst} covers both cases at its two levels.

The blood hunt is part of the action, though independent of it — the window before it is inside the action. A card that
cannot be used during an action therefore never fits before the blood hunt, e.g. {Heidelberg Castle, Germany}.[^3-7-8-2]

What the action itself grants lands before the referendum: votes from {Political Struggle}, and a {Trophy: Diablerie}
retrieved for a Red List victim, which then protects the diablerist in that same referendum.

##### 3.7.8.2 The referendum

The blood hunt referendum is not a political action and is not called by a vampire. Effects keyed to either do not reach
it, e.g. {Charming Lobby}'s follow-up referendum, {Power Structure}.[^3-7-8-3] It is otherwise an ordinary referendum,
so effects keyed to one that passes resolve normally, e.g. {Gangrel Conspiracy} after votes and ballots are cast.

When the referendum follows an {Amaranth}, it happens during combat, so that burn is a burn in combat, e.g. {Hector
Trelane}.[^3-7-8-4]

##### 3.7.8.3 What counts as diablerie

{Amaranth} commits diablerie but is not a diablerie action. {Abactor} is not a diablerie at all and only calls the
referendum. A card requiring a diablerie action is unusable after either, e.g. {Rebirth}; an effect keyed on having
committed diablerie does not see {Abactor}, e.g. {Carlton Van Wyk}.[^3-7-8-5] The referendum either one causes is a real
blood hunt, so blood-hunt effects apply, e.g. {Lay Low}, {Trophy: Diablerie}.

A diablerie the victim survives is unsuccessful: the diablerist gets nothing, blood and equipment stay on the victim,
and no blood hunt is called, e.g. {Reform Body}, {Byzar}.[^3-7-8-6] If the victim is in torpor again, another {Amaranth}
may be played. Bringing the victim out of torpor after the diablerie action is announced does not end the action:
unblocked it resolves successfully with no diablerie, e.g. {Warsaw Station}.

##### 3.7.8.4 Who may commit diablerie

Only a ready vampire may commit diablerie [[RBK diablerie]](https://www.vekn.net/rulebook#diablerie). An ally granted
the use of a card still cannot, e.g. {Shadow Court Satyr} with {Amaranth}.[^3-7-8-7] A vampire barred from diablerie
({Humanitas}, Blood Cursed) cannot be compelled to it, so {Undying Thirst} does nothing to him. Cards compelling
diablerie impose a mandatory action.

A blocker of a leave-torpor action gets the opportunity to diablerize the acting vampire [[RBK
leave-torpor]](https://www.vekn.net/rulebook#leave-torpor); when that opportunity is lost is
[§3.7.7.2](#3772-a-blocked-leave-torpor)'s.

### 3.8 Actions provided by cards

**Base rules.** A minion cannot take the action of a given action card, from hand or in play, more than once each turn,
even if it unlocks [[RBK action-card-or-card-in-play]](https://www.vekn.net/rulebook#action-card-or-card-in-play).

#### 3.8.1 One source, one action

That limit attaches to the card, not to the minion. Each copy in play provides its own action, so a minion that unlocks
in between may take one action per copy, e.g. two {Archon} on the same vampire.[^3-8-1] Actions provided by two
different cards in play are never "the same action", even at identical wording, and no card-provided action is ever the
same as a rulebook action. The same per-copy accounting governs a limited ability granted by a card in play.

One card provides its action once a turn, even where the action would fetch a different card each time, e.g. {Bindusara,
Historian of the Kindred}.[^3-8-2] Where one card provides two different actions they are not cross-restrictive: the
same minion may take each once in the turn. The action taken to put a card into play and the action that card provides
once in play are two different actions, so a minion that unlocks may take both, e.g. {Clandestine Contract}.

#### 3.8.2 What kind of action a provided action is

A provided action has the properties the providing text gives it, and no others.

Where the providing text limits who may take the action, that limit is the action's requirement. {Haqim's Law: Judgment}
provides an action requiring an Independent or Anarch, so {CrimethInc.} can be used after it.[^3-8-3] This is not the
fetched-card case: the requirements of a card brought into play are not the fetching action's.

A provided action of a named type counts as that type for every effect keyed to the type. A card-provided hunt is a
hunt, so {Pariah}, who may take no undirected action but hunting, can take one. A vampire who called a referendum has
taken a political action and can take no other that turn, even if the referendum is canceled.

A card-provided action is not an action card, and neither is a cardless action. Costs and cost reductions worded against
action cards reach neither, e.g. {Ravnos Carnival}'s counters.[^3-8-4] Conversely, an [ACTION] card can provide a
political action without being a [POLITICAL ACTION] card: {Charming Lobby} calls a referendum, but effects naming a
political action card do not see it and it cannot be discarded for a vote.

A provided action that resembles a rulebook action is not that action, and only the cost the card prints is paid: {Go
Anarch} makes the vampire Anarch with no blood paid, unlike the rulebook become-Anarch action [[RBK
become-anarch]](https://www.vekn.net/rulebook#become-anarch).

### 3.9 Mandatory actions

**Base rules.** A mandatory action is performed before any non-mandatory action, and a minion required to take one can
perform no other action. A minion with two different mandatory actions, or with one it cannot take, is stuck and
performs no action.[^3-9-1]

#### 3.9.1 When the obligation attaches

An effect makes an action mandatory only while the condition its text states holds. A vampire whose text compels
diablerie when there are torpid vampires he may diablerize is under no obligation when there are none, and is free to
take any action.[^3-9-2]

Whether another action of the same type discharges the obligation is decided by the mandating card, not by the card used
to satisfy it.

- The card **provides** the action. The obligation is to that action, and another action of the same type does not
  discharge it however it was obtained. {Elen Kamjian} must take the +1 bleed her ability gives her; a bleed from
  {Flurry of Action} does not count.[^3-9-3]
- The card **requires an action of a type**. Any action of that type discharges it, including one taken by playing a
  card. A minion compelled to bleed by {Spirit Marionette} [OBE] may bleed by playing {Computer Hacking}.

Where the mandate names alternatives, either branch discharges it: an empty vampire under {Undying Thirst} satisfies his
mandatory action by hunting or by diablerizing.

#### 3.9.2 Discharge and recurrence

Taking the action discharges the obligation; the action need not succeed. A minion that has performed its mandatory
action and then unlocks is free to take any action — the obligation does not re-attach, e.g. {Cry Wolf}, {Lunatic
Eruption}.[^3-9-4]

A card that requires an action of a type rather than providing one keeps demanding it while the condition holds.
{Phillipe Rigaud} must attempt diablerie again after unlocking as long as an eligible torpid vampire remains.

#### 3.9.3 Stuck

A minion whose mandatory action cannot be taken is stuck: he takes no action at all, not merely no mandatory one. Three
cases recur.[^3-9-5]

- The action was already performed this turn, before the obligation attached. {Elen Kamjian} whose ability turns on
  after she has bled is stuck.
- An effect barred the action. {Change of Target} ends the mandatory action and forbids repeating it, so the acting
  minion is stuck.
- Two copies of the same mandatory-action card sit on one minion, e.g. {Lunatic Eruption}.

#### 3.9.4 Masking

{Mask of a Thousand Faces} can take over an action that is mandatory for the acting minion but not for the masking
vampire. The masker must be capable of the action, not under the same obligation, e.g. a hunt.[^3-9-6] See
[§3.10](#310-changing-the-acting-minion) for masking eligibility.

### 3.10 Changing the acting minion

#### 3.10.1 Substituting the actor in an announced action

An effect that hands an announced action to a different minion leaves the action itself unchanged; it continues from
where it stood. Action modifiers already played stay in effect and every other effect on the action carries over. The
original actor's inherent modifiers do not, e.g. an inherent +1 bleed.[^3-10-1]

The substitute must independently be capable of taking that action, and nothing may have been played or used on the
action that could not have been with him as the acting minion, the original's inherent stealth included. {Force of Will}
requires a locked acting vampire and {Mask of a Thousand Faces} requires an unlocked one, so a {Force of Will} action
cannot be taken over.[^3-10-2] Nor can an action provided by the acting minion's own equipment or ability, even by a
vampire with the same equipment, e.g. a {Bomb} action, {Theo Bell}'s rush.[^3-10-5]

Blood already spent is disregarded in that test. It is not refunded and the substitute does not pay it again.

A mandatory action may be taken over by a minion for whom it was not mandatory, provided he could perform it, e.g. a
hunt.

The substitute may play a modifier the original already played: that limit binds the minion, not the action. He cannot
use a limited effect if a limited effect has already been used on the action.

The replaced minion unlocks. Which of the two took the action for the No Repeated Action rule is
[§3.5.2](#352-the-same-action)'s.

#### 3.10.2 "Is considered the acting minion"

Only text stating that a minion is considered the acting minion rewrites the designation, e.g. {Deep Song} [ANI],
{Nar-Sheptha}. Putting a different minion into the combat does not. A slave locked to enter combat in place of a blocked
clan member is not the acting minion, and {Malleable Visage} leaves the designation with the original.[^3-10-3]

A card that reads "the acting minion" reads the current designation, for its own legality and for its triggers alike.
{Obedience} cannot be played once a slave has taken the acting vampire's place. {FBI Special Affairs Division} does not
trigger when the ally burned is the acting minion. Effects the action gave the vampire that took it stay with him: under
{Deep Song} [ANI] he keeps {Invigorate}'s strength and {Form of Mist}'s continuation *(reading)*.[^3-10-6]

Where two effects write the designation, the later governs, e.g. {Nar-Sheptha} over {Deep Song}. When the later source
leaves play its effect ends immediately and the designation reverts for the rest of the combat, e.g. {Nar-Sheptha}
burned. See [§5.9](#59-traits-and-trait-changes) for the general trait rule.[^3-10-4]

---

## 4. Combat

### 4.1 Combat sequence and rounds

#### 4.1.1 Combat begun outside an action

The acting minion acts first at every step of combat. A combat begun by a card or ability rather than by a block has no
acting minion of its own, so the minion the effect sends into combat takes that role: he declares maneuvers, strikes and
presses first and plays his combat effects first, e.g. {The Guardian}, {Blissful Agony} [VAL]. Where the initiating
player controls neither combatant, the card says which side goes first — {Taunt the Caged Beast} [ANI] gives it to the
prey's vampire.[^4-1-3]

The minion using such an ability is not acting. He plays no action modifiers, and a card whose requirement names an
acting minion cannot be played against him, e.g. {Obedience} against {Marie-Pierre}. Effects that end the combat and
then do something afterwards — including continuing the action as if unblocked — are lost, because there is no action to
return to. Unlock effects still apply.[^4-1-4]

#### 4.1.2 Combat inside an action

A combat that arises during an action is part of that action. The acting minion designation does not change and effects
applied during the action stay in force through the combat. This holds even when the effect that caused the combat also
failed the action, e.g. {Champion}: the action continues until combat ends.[^4-1-5] See
[§3.4](#34-resolution-success-and-effect) for the ordering of after-resolution effects.

A window exists between a successful block and the start of combat. A card barred "during combat" may still be played
there, e.g. {Angel of Berlin}.[^4-1-6]

#### 4.1.3 Who may play cards during combat

A combat card printing permission to be played by a minion "not involved in the current combat" may be played by a
minion of any Methuselah, including one with no minion in the combat, e.g. {Nosferatu Putrescence}, {Bliss}.[^4-1-9]
Using an ability is not playing a card and does not get this latitude.

A reaction card printing "usable even if there is no action" gains no window during combat: it still cannot be played
there.[^4-1-10]

### 4.2 Range

**Base rules.** Range is determined once per round, defaults to close, and a minion needs a card or effect to gain a
maneuver. [[RBK determine-range]](https://www.vekn.net/rulebook#determine-range)

#### 4.2.1 Setting the range

Many cards set the range outright instead of granting a maneuver. Once such an effect resolves, the determine range step
is skipped for that round and no other effect can reset the range.[^4-2-1] Other before-range effects may still be
played afterwards. Some setters fix the *next* round's range instead; the same exclusivity applies to that round, e.g.
{Immortal Grapple} [POT], {Grasp of the Python} [SER].[^4-2-3]

Where two effects would set the same range, the first to resolve has priority and the later one does nothing.[^4-2-2]
Order follows when each effect actually resolves, not when its card was played: a setter that fires on the block
({Squirrel Balance}, then the blocker's {Asanbonsam Ghoul}) resolves before a standing "each round is at long range"
effect that applies at the determine range step ({Neutral Guard}).

A range-setting ability triggered by blocking is checked at the moment of the block and belongs to the minion that
actually blocked, e.g. {Sniper Rifle}. A combat queued by {Psyche!} did not arise from that block, so the ability is not
available in it.[^4-2-6]

#### 4.2.2 The before-range window

Effects usable "before range is determined" must be played before the acting minion decides whether to maneuver. [[RBK
before-range]](https://www.vekn.net/rulebook#before-range) An effect that changes the range in this window takes effect
at once, so range-dependent plays immediately follow the new range: once {Fear of the Void Below} [dai] has made long
the round's default, a close-range card such as {Disarm} cannot be played.[^4-2-4]

Damage inflicted in this window is resolved in it. Both minions may play prevention and non-prevention before-range
effects, and simultaneous before-range damage is prevented and mended together, e.g. {Outside the Hourglass}
[TEM].[^4-2-5]

#### 4.2.3 What the range gates

The range is fixed when the determine range step ends. Effects conditioned on the range apply before strikes are chosen
and cannot be dodged by the strike chosen, e.g. {Vampiric Disease}; a card "only usable at close range" may be played
before or after strikes are chosen, e.g. {Blood to Water}.[^4-2-10]

A card triggered by an incoming effect needs that effect to reach at the current range: {Rötschreck} cannot be played
against aggravated damage that is not effective at the current range.[^4-2-11] The trigger is unmet; futility itself is
no bar.

Range gates strikes and strike resolution, not damage from other sources. Damage that is not strike damage is inflicted
whatever the range, e.g. {Burst of Sunlight}'s damage to the striking vampire. An effect riding on a strike applies at
the range at which that strike resolved, e.g. {Riposte}. Making a strike ranged does not extend its hand-strike or
melee-weapon portion beyond close range, e.g. {Blood of the Cobra} [QUI].[^4-2-12]

At long range a strike may be aimed at a retainer, and a restriction keyed to the opposing minion's own traits does not
bar it, e.g. {Earthshock} against the retainer of a minion with [FLIGHT].[^4-2-13]

### 4.3 Strikes

#### 4.3.1 Choosing a strike

Using the maneuver from a strike card or weapon chooses that minion's initial strike [[RBK
determine-range]](https://www.vekn.net/rulebook#determine-range); a few cards let him instead strike with the weapon he
maneuvered with, e.g. {Blessed Blade}.[^4-3-1] A strike card played in determine range for its maneuver counts as
played; a card merely granting the ability to strike does not, e.g. {.44 Magnum}.[^4-3-2] A cost an opponent imposes on
striking is incurred at that point, and it is a cost of the strike, not of the card played, e.g. {Lucian, the Perfect}
against a {Thrown Gate} played for its maneuver.[^4-3-19] A restriction on playing strike cards reaches strike cards
played in determine range for their maneuver, e.g. {Thoughts Betrayed} [DOM].[^4-3-20]

An aim or other strike modifier is played in the strike step even when the strike was committed during determine range,
and cannot be played once the strike can no longer be chosen, e.g. {Target Vitals} after {Immortal Grapple}. A minion
committed to a strike an effect then bars gets no strike at all.[^4-3-3]

"Strike: make a hand strike" cannot take another strike card as that hand strike; "strike: use a weapon strike" can, and
the result is still a hand strike, e.g. {Bundi}. A nested strike inherits the base strike's properties. A card that is
two strike types needs both available, e.g. {Stutter-Step}.[^4-3-4]

A provided strike carries its card's discipline: {Heroic Might}'s are [pot] strikes. Making a strike ranged leaves an
inner hand or melee portion effective only at close range, e.g. {Blood of the Cobra}.[^4-3-5]

#### 4.3.2 Before strike resolution

"Before resolution" means after both strikes are declared and before they resolve. A modifier that must affect the
current strike is played in that window, ahead of damage prevention, e.g. {Blood Agony}, {Backstab}.[^4-3-7]

A maneuver printed on a strike card or weapon is unavailable to a minion who cannot strike, e.g. against {Lapse} [TEM]
or {Hidden Lurker}; this includes optional weapon maneuvers.[^4-3-8] A maneuver granted by a strike card is usable only
during the round the card was played, e.g. {Rigor Mortis}.[^4-3-21]

Cancellation is [§1.8.4](#184-what-a-cancellation-reaches)'s. A minion whose strike is canceled chooses a new one and
may choose the same again, e.g. {Supernatural Resistance}.[^4-3-9]

#### 4.3.3 Strike resolution

Everything printed on a strike happens at strike resolution, and none of it if the strike does not resolve:
{Contagion}'s steal is dodgeable, {Escaped Mental Patient} does not burn, and burning a committed weapon leaves that
strike with no effect.[^4-3-10]

The strike's effects apply, then damage is resolved. Blood a strike moves is moved before damage is mended, e.g.
{Darkness Within}; blood from burning a minion with a strike arrives after strike resolution and cannot mend that
strike's damage.[^4-3-11]

A damage add-on may be played on any strike and does nothing on one dealing no damage, e.g. {Target Vitals} on a dodge.
Whether the strike deals damage decides, not when — a strike inflicting damage after combat is still damage-dealing,
e.g. {Catatonic Fear} [PRE]. A strike naming the object it acts on cannot be chosen when none exists.[^4-3-12]

An effect worded "after strike resolution" resolves before any further pair of strikes and before the press step, e.g.
{Shoulder Drop}.[^4-3-18]

#### 4.3.4 First strike

A strike done with first strike has its own earlier resolution, so an effect biting when a strike resolves misses it:
{Soul Burn} does not stop a weapon strike made with first strike, and nullifies weapon damage only at
resolution.[^4-3-13] Conversely a modifier excluding "the current strike resolution" reaches the opponent's later normal
resolution — {Scorpion's Touch} with first strike reduces a strength-based strike, as does damage inflicted with first
strike, e.g. {Shambling Hordes}.[^4-3-14]

#### 4.3.5 Additional strikes

Additional strikes, and the cards granting them, come only after the normal pair resolves, and a further additional
strike is gained only once those already granted have been made, e.g. {Quickness} [CEL] after {Blur} [CEL]'s
two.[^4-3-17]

### 4.4 Damage

#### 4.4.1 How much damage

An effect granting "a strength of X" replaces the base strength; modifiers, including inherent ones, then apply to the
new base. A later set overwrites the earlier one, e.g. {Erosion} then {Torn Signpost}. A "+X strength" effect adds to
whichever base is current. A weapon's "current damage" is what it would inflict as a strike by its bearer against a
generic opponent, so strength and other bonuses in play are not part of it. A strike's damage is fixed when strike
resolution begins, e.g. {Shambling Hordes} loses no damage to life burned during that resolution.[^4-4-1]

Additional damage inherits every property of the base damage: type, preventability and source. Aggravated stays
aggravated when {Target Vitals} adds to it, and a gun's added damage is still weapon damage, e.g. {Glaser
Rounds}.[^4-4-2]

{Oubliette} burns blood rather than inflicting damage, so a damage add-on adds nothing to it.[^4-4-3]

#### 4.4.2 Source of damage

The source is read off the card's own wording. Text naming a minion — "this vampire inflicts …" — is minion-inflicted;
text naming the weapon is weapon damage; text naming no one is environmental.[[RBK
damage-resolution]](https://www.vekn.net/rulebook#damage-resolution)

- A strike that lands its damage after combat is still minion-inflicted strike damage: strike modifiers and damage
  reduction reach it, e.g. {Catatonic Fear}, {Outside the Hourglass}, {Riposte}; whether it survives combat continuing
  is [§4.9.2](#492-what-survives-a-continued-combat)'s.[^4-4-4]
- Equipment backfiring on its bearer ({Grenade}), a reaction damaging the striking minion ({Burst of Sunlight}) and a
  printed vampire ability ({Shemti}) are all environmental.[^4-4-5]
- Damage a weapon inflicts outside a strike is the weapon's, neither the bearer's nor environmental, e.g. {Talbot's
  Chainsaw} during the unlock phase. Retainer damage is environmental with the retainer as its source.[^4-4-6]

#### 4.4.3 What sees environmental damage

Environmental damage is not inflicted by any minion, so effects that count, add to or reduce damage a minion inflicts do
not see it, e.g. {Pulled Fangs} counts none of it and {Target Vitals} does not add to {Necrosis}'s press-step damage.
See [§4.5](#45-prevention-and-immunity) for reduction and prevention.[^4-4-7]

A card worded on damage generally does reach it. {Dam the Heart's River} boosts "each strike or damaging effect", so it
boosts the environmental damage as well as the strike; {Dawn Operation} turns environmental and retainer damage
aggravated too. An effect limited to damage "in combat" reaches neither after-combat damage nor anything outside the
combat, e.g. {Domain of Evernight}.[^4-4-8]

What environmental damage survives — dodges, combat ending early — is [§4.6.3](#463-what-survives-a-dodge)'s and
[§4.7.2](#472-combat-ended-outside-strike-resolution)'s.

#### 4.4.4 Aggravated damage

"Treats aggravated damage as normal" changes only how the damage is applied to that minion. The damage is still
aggravated at its source, so {Rötschreck} still works against it; how prevention and immunity read it is
[§4.5](#45-prevention-and-immunity)'s. {Rötschreck} nonetheless requires a strike: it cannot be played against
aggravated damage from a non-strike effect, even one a minion inflicts.[^4-4-11]

A minion burned by damage is burned during damage resolution; damage beyond that is lost, e.g. {Byzar}. An effect keyed
on the vampire's blood reaching 0 does not fire when aggravated damage burns them instead, e.g. {Anathema}.[^4-4-12]

#### 4.4.5 Damage resolution is one block

Damage inflicted at the same moment is prevented and applied together, after all prevention. A single instantaneous
prevention card covers all of it, e.g. {Tunnel Runner}'s theft damage alongside the strike damage. No ability may be
used part-way through, e.g. {Vagabond Mystic}. Damage still applies if the minion inflicting it goes to torpor from the
opposing strike, e.g. {Blood of Acid}.[^4-4-13]

### 4.5 Prevention and immunity

#### 4.5.1 Playing a prevention card

Damage prevention cannot be played when there is no damage to prevent, e.g. {Soak}, {Glancing Blow}.[^4-5-1]

A prevention effect that lingers past a single strike is exempt, and may be played before any damage exists, e.g.
{Apparition}, {Brother's Blood}.[^4-5-2] The restriction reaches the play only
([§1.6.5](#165-playing-a-card-that-will-do-nothing)).[^4-5-3]

#### 4.5.2 How much is prevented

Printed "Prevent X damage" means up to X, e.g. {Soak}, {Nightstick}.[^4-5-5] Where a cost sets X, X may be set higher
than needed, but never negative, e.g. {Hidden Strength}, {Martyr's Resilience}.[^4-5-6]

Unused prevention does not carry over.[^4-5-7] An allotment printed per round is spendable across that round's strikes:
{Armor of Caine's Fury} prevents 2 per round from any strike, or 1 from each of two strikes.[^4-5-8]

Prevention worded against non-aggravated damage cannot prevent aggravated damage, even when the minion treats aggravated
as normal, e.g. {Flesh of Marble} against {Skin of Night}. Where both kinds are inflicted, the preventing minion chooses
which points go unprevented.[^4-5-9]

#### 4.5.3 What can be prevented

Prevention granted for blocking is scoped to the block-induced combat and regranted on a second block
([§3.6.3](#363-effects-scoped-to-one-block)).

A protection effect stopping the opponent's weapon strikes does not stop weapon damage the opponent inflicts on itself,
e.g. {Blood Fury}.[^4-5-11] Neither prevention nor protection applies where the damage would not be effective at the
current range, e.g. {Bollix} at long range.[^4-5-12]

Unpreventable damage still yields to damage reduction, e.g. {Nephandus}.[^4-5-13]

#### 4.5.4 Sequencing

A card modifying a strike must be played before strike resolution, hence before prevention, e.g. {Chiropteran Marauder}.
Mending damage and burning blood to prevent destruction are uninterrupted: no mending part of it, generating blood, then
mending the rest, e.g. {The Coven}.[^4-5-14]

One instantaneous prevention card covers two damage events only if they are simultaneous, e.g. {Tunnel Runner}; where
the second follows strike damage it does not, e.g. {Blood of Acid} — a lingering prevention covers both. Strike damage
is prevented first, acting player first; all damage applies at once afterwards.[^4-5-15]

#### 4.5.5 Immunity

Damage from a source a minion is immune to is inflicted unsuccessfully — no mend, no wound, no destruction ([[RBK
damage-resolution]](https://www.vekn.net/rulebook#damage-resolution)). Immunity holds against unpreventable damage and
outside combat, e.g. {Bloodform}.[^4-5-17] Unlike prevention, immunity to non-aggravated damage does cover aggravated
damage the minion treats as normal, e.g. {Ex Nihilo} with {Skin of Night}.[^4-5-18]

Immunity stops the damage, not the play: the source may still use an effect that would damage an immune minion, e.g.
{Charnas the Imp}. Immunity to a class of source reaches environmental damage from that source, a retainer's damage
included.[^4-5-19]

Prevented damage, like damage a minion is immune to, is inflicted but not successfully inflicted. An effect keyed to
damage successfully inflicted needs damage left unprevented, e.g. {Pulled Fangs}, {Target Vitals}. An effect keyed to
damage inflicted fires anyway, e.g. {Improvised Flamethrower}, {Weighted Walking Stick}, whose "(even if prevented)" is
reminder text. Older cards printed "inflicts" for "successfully inflicts": a ruling reading a card's "inflicts" so
stands until a reprint updates the text, e.g. {Twisting the Knife}, {Blade Clot}; without such a ruling the printed
"inflicts" counts prevented damage, e.g. {Haymaker} *(reading)*. Damage reduced away or dodged is not inflicted at
all.[^4-5-20]

A strike that sends a vampire to torpor leaves that vampire wounded; {Undying Tenacity} defers the torpor to the end of
combat but not the wound, e.g. {Coma}.[^4-5-21]

### 4.6 Dodge

**Base rules.** A dodge is a strike that deals no damage and protects the dodging minion and their possessions from the
effects of the opposing strike; retainers are not protected.

#### 4.6.1 A dodge negates the whole strike

A dodge stops the strike, not merely its damage: {Serpent's Numbing Kiss} places no card and locks no one; {Blissful
Agony} [VAL] starts no new combat.[^4-6-1]

A strike that would move its own card onto the opposing minion leaves that card where it was: a weapon so worded stays
equipped on the bearer, e.g. {Rowan Ring}, {Enhanced Coagulant}.

An effect worded as a step following strike resolution is still part of the strike and is dodged with it, e.g.
{Contagion} [DAI] steals the minion only if the strike resolves.

A secondary effect that disables the opponent's weapons fails too. Dodge the {Blood Fury} template and the dodging
minion's weapon strikes inflict damage normally that round.[^4-6-2]

#### 4.6.2 Dodge against prevention

Preventing the damage leaves the strike resolved, so the rest of the strike's effect still applies unless it requires
damage successfully inflicted ([§4.5.5](#455-immunity)). Dodging means the strike never resolved and none of it
applies. {Fleshcraft} prevented still places the card; {Fleshcraft} dodged does not.[^4-6-3]

#### 4.6.3 What survives a dodge

Only what the dodged strike itself delivers is stopped.

- **Damage from anything but the dodged strike.** A card in play or a combat card inflicting damage on its own schedule
  is unaffected, e.g. {Conscripted Statue} [vis], {Darkling Trickery} [MYT].[^4-6-4]
- **But environmental damage printed as part of the strike** is dodged like the rest of it, e.g. {Necrosis}
  [THN].[^4-6-5] The label does not decide it; the source does.
- **Combat ends**, which the base rules exempt from dodges. An effect ending combat after strike resolution still ends
  it, e.g. {Anesthetic Touch} [obe].[^4-6-6]
- **Effects keyed to the attempt rather than to the damage**, e.g. {Rötschreck} against an announced dodge or combat
  ends.
- **A cost or self-burn the strike card imposes.** {Flash Grenade} burns even though its lock is dodged.[^4-6-7]
- **An ability of the weapon that is not part of the strike.** {Garrote}'s burn-instead-of-torpor is usable when the
  opponent dodged but goes to torpor anyway.

#### 4.6.4 Dodging when it will accomplish nothing

"This strike cannot be dodged" does not stop the opponent choosing a dodge; the dodge simply has no effect against that
strike, e.g. {Scorpion Sting}, {Projectile}.[^4-6-8] That is the permissive default, not an exception to it, and it does
not carry over to damage prevention, which cannot be played with no damage to prevent.

A card that is a dodge and another strike at once needs both to be legal choices. Its dodge half protects, including
against a strike made with first strike, while the other half resolves normally — {Stutter-Step}.[^4-6-9]

### 4.7 Strike: Combat Ends

**Base rules.** A strike: combat ends always resolves first, before a strike done with first strike, and it ends combat
before any other strike or strike resolution effect resolves [[RBK
strike-effects]](https://www.vekn.net/rulebook#strike-effects). A dodge does not stop it.

#### 4.7.1 An unresolved strike does nothing

A strike that never resolves produces none of its effects and pays none of its costs. A weapon or ally whose text burns
it for striking is not burned, e.g. {Bomb}, {Waxen Poetica}.[^4-7-1] The same holds for every other consequence of the
strike: blood is not taken, a card the strike would place is not placed, an unlock it carries does not fire.[^4-7-2] The
strike stays unresolved even if combat then continues through another effect, e.g. {Relentless Reaper}.

A card that is itself a strike: combat ends does resolve when the opponent also plays one: both are combat ends, so both
resolve. Its burn-after-use clause therefore applies, e.g. {Flash Grenade}, {Smoke Grenade}.[^4-7-3] The burn attaches
to the strike, not to the combat: it is skipped only when the card's own strike fails to resolve.

An ability keyed to the strike resolution step is likewise unavailable, e.g. {Gianna di Canneto}'s lock-to-damage
ability. The same is true whenever combat ends during first strike — a combatant sent to torpor by first strike
aggravated damage ends combat before the strike resolution step completes.[^4-7-4]

An effect barring strike: combat ends that arrives after the strike was announced leaves the announced strike with no
effect, e.g. {Dog Pack} put on a minion mid-combat.[^4-7-5]

#### 4.7.2 Combat ended outside strike resolution

An effect that ends combat outside the strike resolution step beats an announced strike: combat ends. {Rötschreck} may
be played after the opponent has announced a strike: dodge or a strike: combat ends, and is effective.[^4-7-6] Combat
ends immediately, no strike resolves — including the announced strike: combat ends — and damage prevention cannot be
played. Self-burning weapons are not burned, exactly as in [§4.7.1](#471-an-unresolved-strike-does-nothing).
Environmental damage, which survives a dodge, does not survive this: a strike: combat ends ends combat before it is
applied.

See [§4.9](#49-end-of-round-end-of-combat-and-new-combats) for what else is lost when combat ends early.

#### 4.7.3 "Combat ends after this strike resolves" is not a strike: combat ends

A card reading *combat ends immediately after this strike resolves* provides an ordinary strike, not a combat ends. Both
strikes resolve normally and simultaneously; the opposing strike inflicts its damage, which can be prevented or mended
as usual, and combat ends afterwards, e.g. {Anesthetic Touch} [obe], {Autonomic Mastery} [DOM].[^4-7-7] A dodge protects
against the opposing strike but does not stop combat from ending, unless the card itself says so — {Autonomic Mastery}
prints "unless it is dodged", {Anesthetic Touch} does not.

### 4.8 Presses

A press provided by a card is usable only during the round in which that card was played; an unused press is lost at
the end of the round.[^4-8-1] This holds regardless of how long the rest of the card lasts: a press riding on a strike,
a maneuver or a range lock expires with the round even where the card's other effect runs for the whole combat, e.g.
{Immortal Grapple} [POT], {Dust to Dust} [thn]. A card whose effect spans several rounds likewise grants one press, for
the current round only, e.g. {Undead Persistence}.[^4-8-2]

A press granted by a minion's own ability is usable only during the press step.[^4-8-3] This applies both to an ability
that gets the bearer a press, e.g. {Aeron}, and to one that gives a press to another minion, e.g. {Don Caravelli}. Such
an ability cannot be used earlier in the round to bank a press.

A press is mandatory unless marked optional: {Talbot's Chainsaw} prints "1 press (mandatory)"; {Chameleon's Colors}
[spi] prints "an optional press". A mandatory press must be used, and where the card restricts its direction — continue
only, or end only — it is used in that direction.[^4-8-4]

An existing press to continue must be handled before another press can be used to continue.[^4-8-5] A second press to
continue cannot be stacked on the first: the opposing minion's only reply is a press to end. So {Trap}'s automatic press
must be answered before a press to continue such as {Boxed In} becomes playable.

Providing a press is not using one. A minion that cannot use presses may still play a card that supplies them, e.g.
{Mukhtar Bey} playing {Trap}: the card provides the press and the minion is not the one pressing.[^4-8-6]

Canceling a maneuver that provides a press cancels the press.[^4-8-7]

Damage inflicted during the press step is environmental damage, e.g. {Drawing Out the Beast} [ANI].[^4-8-8] An effect
scheduled for the press step does not happen if combat ends before it, e.g. {Necrosis} [THN] when a strike ends
combat.[^4-8-9]

### 4.9 End of round, end of combat, and new combats

**Base rules.** The end-of-round step follows the press step and occurs even when combat ends prematurely. [[RBK
end-of-round]](https://www.vekn.net/rulebook#end-of-round)

#### 4.9.1 The end-of-round window

Effects usable at the end of a round wait for presses to be handled, then apply even if the round ended
prematurely.[^4-9-1] This governs triggered effects, not only card plays, e.g. {Ossian}'s life gain, {Masochism}'s
counter.

Ordering within the window is [§2.4.1](#241-ordering-within-a-window)'s: occupants may be played before or after one
another, e.g. {Taste of Vitae}, {Disarm}, {Telepathic Tracking}.[^4-9-2]

Ending combat does not close the window — end-of-round and end-of-combat cards can still be played, e.g. {Elysium: The
Arboretum}.[^4-9-3] Conversely, a card usable at the end of combat may be played before combat ends, e.g. {Amaranth}.

A card played in this window has been played during that round, so cancels and cost modifiers keyed to the round or to
combat cards reach it, e.g. {Death Seeker}, {Terror Frenzy}.[^4-9-4] A cost payable until the end of combat may still be
paid after other end-of-round effects, e.g. {Loving Agony}'s unlock blood.

{Psyche!} is played after presses, when combat is about to end; end-of-round and "at the end of combat" effects may be
played before or after it.[^4-9-5] A "would end" replacement such as {Telepathic Tracking} played before it continues
the combat, and {Psyche!} becomes playable only at the end of the new round.[^4-9-6]

The shape of the step: presses first; then a single shared window — end of round and end of combat are the same window
when no press is taken and combat is about to end — whose occupants order freely, with one fixed point: a "would end"
replacement must be played before the window's "about to end" effects. Effects keyed "after combat" sit outside it and
wait for combat to actually end.

#### 4.9.2 What survives a continued combat

An effect that ends combat and then does something else after combat loses the after-combat part when the combat
continues or a new combat begins, whether the continuation comes from a card or from a minion's ability, e.g. {Psyche!},
{Akram}.[^4-9-7]

Unlock effects are the exception. The unlock half of a "strike: combat ends" resolves when the strike resolves, before
combat ends, so it survives; the rest resolves after combat ends and is lost, e.g. {Mummify}, {Meld with the
Land}.[^4-9-8] {Flash Grenade}'s own lock is not conditioned on the end of combat: it lands even on a vampire whose
simultaneous strike unlocked him — unlocked by {Earth Meld} [PRO] first, then locked by the grenade.[^4-9-23]

Only the after-combat half is lost; what already resolved stands. {Smoke Grenade} still burns. {Flash Grenade} burns but
its lock is lost. {Rötschreck} is still put on the vampire, who does not go to torpor but still does not unlock as
normal.[^4-9-9]

Damage, blood burn and card placement scheduled after combat do not happen, e.g. {Catatonic Fear}, {Riposte}, {Serpent's
Numbing Kiss}.[^4-9-10] An effect keyed to blood reduced to 0 in combat is not triggered by after-combat damage
({Anathema}).[^4-9-11]

An effect timed before range is determined still applies when combat ends at that point, e.g. {Weather
Control}.[^4-9-13]

#### 4.9.3 Queuing a new combat

To queue a combat is to set up a combat that begins later in the same action, e.g. {Hedonism}, {Psyche!}. One queued
during a combat begins when that combat ends. One queued earlier begins once the action has resolved, still as part of
its resolution: after the block's combat, if any, when the action was blocked; after the action's effects when it
succeeded. Every combat of the action occurs during its resolution: an effect used after a combat to start another,
e.g. {Hidden Lurker}, {Fast Reaction}, brings the action back into its resolution for that combat.[^4-9-24] A card
printing its own start for the combat keeps it, e.g. {Siren's Lure} ("After the action resolution").[^4-9-25]

At most one combat may be queued at a time. Every effect starting a new combat in the action uses the same queue,
whatever its wording, so none may be used while a combat is queued, e.g. {Psyche!}, {Hidden Lurker}, {Akram}. A queued
combat does not occur if the action ends first.[^4-9-15] Extending the current combat is not queuing one, e.g.
{Telepathic Tracking}. Nor is replacing the block's combat: a slave may still take a blocked clan member's place while
a combat is queued, e.g. by {Hedonism}.[^4-9-26]

A vampire in torpor cannot enter combat [[RBK leave-torpor]](https://www.vekn.net/rulebook#leave-torpor), and a
combat-starting effect cannot reach a minion in torpor or on his way there, whatever the card's wording: neither {Hidden
Lurker} nor {Blissful Agony} [VAL] may be played on a vampire going to torpor. A queued combat one of whose combatants
is in torpor when it would begin does not happen at all. An effect forbidding further combat bars combat-starting cards,
e.g. {Heaven's Gate}.[^4-9-16]

Per-combat conditions reset. A cost keyed to entering combat is paid again ({Blithe Acceptance}). Permissions tied to
the blocked combat do not carry over ({Sniper Rifle}, {Scry the Hearthstone}). Cards replaced only after combat are
replaced before the new combat begins.[^4-9-18] An effect triggering after combat triggers once per combat when several
occur in one action, e.g. {Amelia, The Blood Red Tears}.[^4-9-19]

#### 4.9.4 After combat

An ability usable after combat has ended comes after cards played when combat "would end" or is "about to end", e.g.
{Marie-Pierre}.[^4-9-20]

Where a new combat follows, effects sequenced between combats apply between them, and effects waiting on the action wait
for the last combat, e.g. {Yawp Court}.[^4-9-21]

A "go to torpor" postponed to the end of combat is postponed again by a new combat, and the postponing card's protection
does not extend into that combat, e.g. {Undead Persistence}.[^4-9-22]

### 4.10 Weapons and equipment in combat

#### 4.10.1 Using a weapon without striking with it

Most weapons print functions besides their strike: an action, a press, a stealth reduction, damage prevention. Using one
never commits the bearer to strike with that weapon, e.g. {Talbot's Chainsaw}'s enter-combat action, {Starshell Grenade
Launcher}'s stealth reduction.[^4-10-1] A weapon's maneuver is the exception ([§4.3.1](#431-choosing-a-strike)). An
additional strike the weapon grants is had even if the weapon's own strike is not used, e.g. {Sword of Judgment}.

An ability keyed to blocking needs the weapon equipped at the time of the block; equipping mid-combat arrives too late,
e.g. {Sniper Rifle}'s range-setting after a {Disguised Weapon}.[^4-10-3]

#### 4.10.2 Once per combat, once per round

The "once each combat" or "once each round" allowance tracks the weapon, so it is not renewed when the weapon changes
hands — {.44 Magnum} gives one maneuver each combat however many bearers it passes through.[^4-10-4] A strike a card
merely grants for the round carries no such limit unless its own text prints one, e.g. {Hunger of Marduk}.

#### 4.10.3 Weapon damage

A weapon's "current damage" is the damage it would inflict as a strike by its bearer against a generic opponent at the
appropriate range.[^4-10-5] Strength and bonuses from the minion or from other cards in play are excluded, and the value
is fixed when the card naming it is announced, e.g. {Machine Blitz}, {Illegal Search and Seizure}.

Reading a weapon's damage value is not using the weapon: no restriction or side effect of the weapon applies, so a
burn-after-use weapon is not burned and a once-per-combat weapon is not spent.[^4-10-6] Such damage is not aggravated
even when the weapon's own strike is.

A weapon that provides a ranged strike is a ranged weapon whatever its printed type, e.g. {Kerrie} under
{Anachronism}.[^4-10-7] Damage from a weapon is weapon damage whatever delivers it: the ranged option on a melee weapon
still does that weapon's damage, and {Talbot's Chainsaw}'s unlock-phase damage is weapon damage, neither inflicted by
the bearer nor environmental.

#### 4.10.4 Weapons that leave or become unusable

A weapon already committed as the strike but unavailable when that strike resolves makes the strike have no effect —
burned by an opponent's effect, or contested because a {Disguised Weapon} equipped a second copy of the same unique
weapon.[^4-10-8] Up to that point the weapon is still usable [[RBK
strike-effects]](https://www.vekn.net/rulebook#strike-effects).

Once a weapon's strike has resolved, its self-burn and bearer damage happen as printed, even if combat continues or a
new combat begins, e.g. {Smoke Grenade}.[^4-10-9]

An effect that prevents a minion from using equipment reaches only what the bearer uses; an equipment effect that
triggers on its own still functions, e.g. {Soul Gem of Etrius} when the bearer is burned.[^4-10-10]

Ammo and other cards improving a weapon before strike resolution can only be played on your own minion's
weapon.[^4-10-11]

---

## 5. Minions and Their States

### 5.1 Vampires: capacity, identity and merging

#### 5.1.1 Capacity

A vampire's capacity can never be reduced below one. A reduction printed as lasting until the end of the game survives
the burning of the card that imposed it, and still applies if the vampire returns to the uncontrolled region.[^5-1-1]

Cards on an uncontrolled vampire are out of play, so a capacity modifier from one of them does not apply there. It
resumes when the vampire is controlled again, and blood held for the raised capacity does not drain off in the
meantime.[^5-1-2] A contested vampire likewise keeps such blood until he returns *(reading)*.[^5-1-12] With the floor of
one, a vampire whose capacity comes entirely from cards on it is a 1-capacity vampire while uncontrolled, e.g. {The
Becoming}.

Blood above capacity drains off as the vampire is moved to the ready region, before anything else happens to them. It
drains before the vampire's own entry text fires, e.g. {Hermana Hambrienta Mayor}'s forced 2-blood burn, and before the
vampire can be merged: {Tariq, The Silent}'s printed capacity reduction applies the moment you control him, so the blood
is lost even if you merge him immediately after.[^5-1-3]

#### 5.1.2 Merging

Merging is not entering play, so effects keyed on a vampire who entered play or entered the ready region in a given
window do not see it, e.g. {Chameleon}, {Legendary Vampire}. It is not moving a vampire from uncontrolled to controlled
either, so a tax on that movement does not apply, e.g. {Masquerade Enforcement}.[^5-1-5]

The advanced card's sect is written onto the vampire as it merges, including when the base version was the one in play,
e.g. {Goratrix} becomes Camarilla. Any anarch token on the vampire is burned, unless the merged sect is itself Anarch,
e.g. {Dancin' Dana}.[^5-1-6] See [§5.9](#59-traits-and-trait-changes) for trait change and its precedence.

The advanced card need not come from your uncontrolled region. A copy you come to control by any means may be merged,
e.g. one stolen with {Graverobbing}.[^5-1-7]

#### 5.1.3 Created vampires and crypt card identity

A library card put into play as a vampire is a vampire card for every purpose from then on, e.g. {Creation Rites}: sent
to the uncontrolled region it stays a crypt card — out of play like any uncontrolled card — and can be influenced out
again. Only in the ash heap is it the library card again, so effects reaching a vampire in the ash heap cannot see
it.[^5-1-8] An effect that would put it into a crypt burns it instead.[^5-1-13]

Such a minion is non-unique unless its card says otherwise, e.g. {Hatchling}. Where the card does make it unique,
uniqueness runs on the card name and crosses minion types: a mummy ally created by {Spell of Life} contests a vampire of
the same name.[^5-1-9]

{Dual Form} copies what its text lists — clan, sect, capacity and Disciplines — so titles and other traits such as
Infernal or Scarce are not copied; there is no generic copy rule beyond the card's own words. The copy is a snapshot
taken at resolution: disciplines held then are copied, including those granted by cards in play, and are kept when the
granting card leaves; disciplines gained later are not. The capacity reduction the card itself imposes applies first, so
the copy has the reduced capacity.[^5-1-10]

A vampire brought into play to replace a burned one is a new vampire for all purposes even when it is the same vampire,
e.g. through {Soul Gem of Etrius}: it may bleed although the burned version already bled that turn.[^5-1-11]

### 5.2 Locking and unlocking

**Base rules.** *Ready* means in the ready region, i.e. not in torpor; *unlocked* means not turned sideways. A vampire
in torpor is not ready but still unlocks each unlock phase — hence "ready, unlocked" in card text.

#### 5.2.1 What a lock prevents

Locking bars acting and blocking, not ability use.[^5-2-1] An action modifier carries no unlocked requirement of its own
— even one played by a minion other than the acting minion, e.g. {Cloak the Gathering}, {Make an Example}.[^5-2-2]
Locked vampires still vote, and casting votes does not lock.[^5-2-3] A locked minion remains a legal target, even of an
effect that locks it, e.g. {Puppet Master}; a lock-to-use cost, by contrast, requires the minion unlocked.[^5-2-4]

#### 5.2.2 The unlock phase and its suppression

Unlocking is the first thing in the phase; every "during your unlock phase" effect follows, in its controller's chosen
order, e.g. {Baleful Doll}.[^5-2-6] A condition on what a Methuselah controls during that phase is checked when the
phase ends: meeting it earlier does not help, e.g. {Anarch Revolt}.[^5-2-7] A "for each" unlock-phase effect instead
resolves once, as a whole, over the objects in play at that moment: objects acquired after it resolves are unaffected
until the Methuselah's next unlock phase, e.g. {Arika}, {Nightmares upon Nightmares}.[^5-2-30]

"Does not unlock as normal" suppresses only that automatic unlock; wakes and unlock effects still reach the minion, e.g.
{Sensory Deprivation}, {Cry Wolf}.[^5-2-8] A card saying minions *cannot unlock* reaches unlock effects too, but still
not wakes, e.g. {The Sleeping Mind}.[^5-2-10]

A stun locks a minion and puts a stun counter on him, e.g. {Mind Numb}, {Hellene Panhard}. A minion with one or more
stun counters "does not unlock as normal" at the beginning of his controller's unlock phase; during that phase, the stun
counters he had at the beginning of the turn are burned.[^5-2-31]

Two burn-at-unlock templates differ. "Choose not to unlock as normal and burn this card" is an independent act, still
available to a minion already prevented from unlocking, e.g. {Putrefaction}. "Burn this card instead of unlocking as
normal" is a substitution: open to an already-unlocked minion, closed to one prevented, e.g. {Fata Amria}.[^5-2-12]

#### 5.2.3 Wakes, unlocking and blocking

A wake opens the reaction window; it does not unlock. A woken locked minion may play reactions and attempt to block, but
cannot use an effect requiring an unlocked minion or costing a lock, e.g. {Familial Bond}.[^5-2-13] Paying a lock leaves
the minion unable to block without a wake or an unlock, e.g. {Starshell Grenade Launcher}.[^5-2-14] An effect merely
letting a locked minion block is not a wake and grants no reaction cards, e.g. {No Secrets From the Magaji}.[^5-2-15]

The unlock-to-block template ({Guard Duty}, {Second Tradition: Domain}) is [§3.3.1](#331-who-may-block)'s.

A block locks the blocker only once it succeeds ([§3.3.4](#334-after-the-block-succeeds)) — a reaction usable on a
successful block thus finds its player locked, e.g. {Hard Case}. {Faceless Night} [OBF] locks failed blockers by its
own text, just before action resolution.[^5-2-19]
A lock from another card is not the block's lock, so an exemption from locking to block does not stop it, e.g. {Mirror
Walk}; locking a minion attempting to block makes the block fail, e.g. {Alexandra}.[^5-2-20] See
[§3.5](#35-action-repetition-nra-and-canceled-actions) for the canceled-combat case.

#### 5.2.4 Locking the locked, unlocking the unlocked

Both are legal, e.g. {Anarch Troublemaker}, {Under Siege}, and locking an already-locked minion can be useful: it
satisfies a lock-or-discard effect that an infernal trait then undoes, e.g. {Nightmares upon Nightmares}.[^5-2-21] No
state changes, so an effect triggered by unlocking does not fire, e.g. {Vampiric Disease}; a mandatory unlock ability is
still applied, e.g. {Eze, The Demon Prince}. Any genuine unlock triggers, including one that already cost
blood.[^5-2-22] A minion returning to play unlocked has not unlocked, e.g. {Banishment}.[^5-2-23] Where a card locks a
minion as part of its effect, nothing need be unlocked and no cost is paid, e.g. {Disarming Presence}.[^5-2-24]

### 5.3 Torpor

**Base rules.** A torpid vampire is controlled but not ready; it unlocks normally, cannot block, play reaction cards or
vote, and acts only to leave torpor [[RBK torpor]](https://www.vekn.net/rulebook#torpor).

#### 5.3.1 Going to torpor

"When a vampire is going into torpor" and "when a vampire should go into torpor" are one window. Both players may play
in it, the acting minion first, and the first card played can strip the other's requirement: once torpor is averted, no
card needing the vampire to be going to torpor may follow — {Undead Persistence} denies {Amaranth}, and the reverse —
including a second copy of the card that averted it.[^5-3-1]
A clause that chooses vampires and sends them to torpor opens no window in between.

A vampire on its way to torpor is still in play and still ready, and may use its abilities and play cards until it
arrives, e.g. {Watenda}, {Revelation of Wrath}.[^5-3-2]

A new combat cannot reach a vampire on his way to torpor. A combat already queued still occurs when the vampire going to
torpor is its queuer rather than a combatant, e.g. {Illusions of the Kindred}.[^5-3-3]

**Delayed torpor.** The {Undying Tenacity} / {Undead Persistence} template pushes torpor to after combat; the vampire
stays wounded and remains burnable.[^5-3-4] Everything else that trigger produced still happens — combat ends, and an
unlock or blood gain in the same clause occurs, e.g. {Ashes to Ashes}. Where torpor is instead bundled with ending
combat and the combat does not end, the vampire unlocks but stays out of torpor, e.g. {Mummify}.

Such a vampire may play cards a vampire going to torpor could not. The protection does not reach a fresh combat
([§4.9.3](#493-queuing-a-new-combat)).[^5-3-5]

#### 5.3.2 What a torpid vampire can still do

Abilities printed on cards in play remain usable in torpor. Printed effects that merely apply do so too, e.g. {Nahir}'s
hand size, as does a trigger keyed to the end of combat, e.g. {Amelia, The Blood Red Tears}.[^5-3-6] Referendum
abilities from torpor are [§3.7.6.4](#3764-locked-and-torpid-vampires)'s.

Torpor bars reaction cards — absent the card's own printed permission — but not combat cards. A combat card requiring an
*unlocked* vampire not involved in the combat may be played from torpor, since a torpid vampire unlocks as normal — e.g.
{Save Face}, {Martyr's Resilience}.[^5-3-8] A requirement naming a *ready* vampire is not met, and a reaction stays
barred even where the vampire reached torpor during the combat that followed his block, e.g. {Cats' Guidance}.[^5-3-9]

A vampire acting from torpor — necessarily the leave-torpor action — plays action modifiers as normal, e.g. {The Kiss
of Ra} on the leave-torpor action itself; a torpid vampire plays no action modifier during another minion's action, e.g.
{Make an Example}.[^5-3-10] Modifiers usable after combat may also be played from torpor, e.g. {Freak
Drive}.[^5-3-11]

#### 5.3.3 Torpid vampires as objects

Torpor removes readiness, not control. A burn option opens when you control no minion meeting the card's requirement. A
torpid vampire who meets it counts as meeting it only for an action modifier, playable on his own leave-torpor action,
so the option stays closed, e.g. {Council of Seraphim}, {The Red Question}; for any other card, which he cannot play, it
opens, e.g. {Emergency Powers}, {Veneficti}, {Evil Eye} *(reading)*.[^5-3-18] Meeting the requirement for one mode of a
card keeps it closed, e.g. {High Orun}.[^5-3-12] Only the requirement is read: whether the card has a legal target does
not open the burn option, e.g. {Consume the Dead}, nor does a copy of the unique card already in play.[^5-3-17]

A card may be put on a torpid vampire unless the operating clause requires readiness: {Fear of Mekhet} cannot be played
on one but can be moved to one.[^5-3-13] A "burn this card after this minion goes to torpor" clause fires on the event,
not the state, so a card entering play on a vampire already in torpor is not burned.[^5-3-14] An ongoing effect whose
subject was chosen while ready keeps applying in torpor, and stops only if that vampire leaves your control, e.g. {The
Rack}.

A change of controller does not change torpor: a stolen or defecting vampire stays in torpor [[RBK
8-glossaries]](https://www.vekn.net/rulebook#8-glossaries).[^5-3-15] Allies have no torpor state: an effect that would
send an ally to torpor burns it instead, and one sending vampires to torpor leaves an opposing ally unaffected.[^5-3-16]
Incapacitation is the imbued analogue, and abilities stay usable there too, e.g. {Abjure}.

### 5.4 Burning minions

Only a controlled minion in play can be burned from play. Uncontrolled vampires and contested cards are out of play and
do not qualify, so effects that retrieve a minion "burned from play" never reach them, e.g. {Blessed
Resilience}.[^5-4-1] What was burned is read off the card, not off the role it filled: a crypt card put into play as a
wraith ally by {Khazar's Diary (Endless Night)} is not an ally burned from play, and a card treated as a vampire only
while blocking is not a minion when it burns, e.g. {Unleash Hell's Fury}. A burned minion goes to its owner's ash heap,
and a pending instruction to move it elsewhere lapses — a vampire put into play by {Chain of Command} and then burned is
not moved to the crypt.

#### 5.4.1 Who burned it

If a minion is burned in combat, its opponent is always considered to have burned it, whatever the cause.[^5-4-2] This
holds when the minion burns itself: an {Escaped Mental Patient} burned at the end of combat is still an ally burned in
combat by its opponent. It also holds when the card's controller is not the combatant — the vampire playing {Blood
Brother Ambush} has not burned a minion, but the opposing minion has.

Outside combat, the acting minion burned it only when the action's own effect did: damage inflicted, or blood or life
taken, by the action. An ally that burns as a result of a directed action is burned by that action, so Trophy cards and
{Predator's Transformation} trigger, e.g. {Brick Laying}, {Cryptic Mission}, {Succubus}.[^5-4-3] A minion burned for any
other reason — a referendum result, a cost paid, its own ability, its last life spent — was not burned by the acting
minion, and cards keyed to "when this vampire burns a minion" do not fire, e.g. {Taking the Skin: Minion}.[^5-4-4]

#### 5.4.2 Effects triggered by the burn

An effect keyed to a minion being burned fires however the burn came about, e.g. {Political Struggle}.[^5-4-5] Blood or
pool awarded for burning a minion with a strike is awarded after strike resolution, so it is not available to heal
damage from that same strike, e.g. {Young Bloods}.

#### 5.4.3 Burns that do not happen

A replacement effect worded "if this minion would be burned, instead ..." means the minion is not burned: effects that
trigger after the burn do not fire, e.g. {Sacrificial Lamb} gains no blood off {Byzar}.[^5-4-6] An instruction or
condition to burn a minion is still satisfied — {Byzar}'s controller may still burn him to burn a {Judgment: Camarilla
Segregation}. Such a replacement takes effect before anything triggering on the burn, e.g. {Soul Gem of Etrius}; a
competing replacement may still be played first, by normal sequencing.

An effect that saves a minion by ending combat cannot reach a burn that occurs as combat ends, e.g. {Heaven's Gate} and
{Left for Dead} cannot save an {Escaped Mental Patient} whose strike resolved. A minion so saved was never burned, so
burn triggers do not fire.

#### 5.4.4 The "would be burned" window

A minion on its way out of play is still in play and may still play cards, e.g. {Revelation of Wrath}.[^5-4-7] A burn
from a blood hunt referendum opens no such window: {Reform Body} and {Abandoning the Flesh} cannot be played against
it.

### 5.5 Allies

#### 5.5.1 Life, not blood

An ally has life, not blood, and can never pay a blood cost. A card or ability whose cost is in blood is unplayable by
an ally, e.g. {Codex of the Edenic Groundskeepers}. An effect that demands blood from the acting or blocking minion
fails for an ally: an ally's action is lost to {Hide the Heart} [val], and an ally cannot block through {Tenebrous Form}
[OBT] or take a directed action against {Étienne Fauberge}. A cost reduced to zero is payable — {The Shard, London} lets
an ally play a blood-cost card that does not otherwise require a vampire.[^5-5-1]

A card whose effect names the opposing vampire has no object in an ally and cannot be played against one, e.g. {Taste of
Vitae}, {Enhanced Coagulant}. Where the blood is only a side effect, the rest still applies: {Shackles of Enkidu} holds
an ally but burns nothing. An effect worded on vampires passes an ally over — {Legacy of Power} sends the vampires to
torpor and leaves the ally where it stands.[^5-5-2]

#### 5.5.2 Life totals

An ally's life is not capped by its starting life; life gained above it stays. "Starting life" is the printed value of
the card that entered play, and later change does not move it: {Ghouled} adds a life but the ally's starting life is
still the mortal's, and an animal retainer moved to the ready region by {Demdemeh} keeps its counters and its original
starting life while losing its abilities and any continuous effect on it. A vampire treated as an ally has no starting
life at all, so {Vagabond Mystic} cannot give it life; its blood is its life, and an "as if he had the Discipline"
ability still works while it is an ally.[^5-5-3]

#### 5.5.3 Acting "as a vampire"

An ally that can play cards as a vampire may play any card its own text admits, not only Discipline cards, e.g. {Shadow
Court Satyr} playing {Taste of Vitae}. It may call a referendum listed on a political action card in hand, e.g. {Herald
of Topheth} with {Charming Lobby}. It cannot commit diablerie, e.g. {Amaranth}. The Discipline grant is whatever the
ally's text says: {Shadow Court Satyr} uses his stored card as a vampire with the basic level of one of its Disciplines,
chosen at each use, so an option requiring two at once is out of reach.[^5-5-4]

The treatment covers only the effect generated by playing the card. An outside effect keying on "a vampire" does not see
the ally: {Veil of Darkness} ignores it and {The Line} cannot reduce its cost. Neither does the played card's own
in-play text, so an ally that plays {Descent into Darkness} never returns. A card that lets a minion act or block "as an
ally" likewise changes only the check it names — {Sonja Blue} blocks as an ally and still plays Discipline
cards.[^5-5-5]

Life pays for the card, so spending the last life burns the ally mid-play. An action modifier played with the last life
fails; a strike played with the last life ends combat before strike resolution; but a cost paid for an action still
resolves the action fully, and the oust order is [§6.5.2](#652-when-the-oust-resolves)'s.[^5-5-6]

#### 5.5.4 Entering play and acting

A recruited ally cannot act the turn it enters play; one brought into play by other means can [[RBK
recruit-ally]](https://www.vekn.net/rulebook#recruit-ally). The keyword decides it: cards that make a minion **recruit**
an ally leave it unable to act that turn, e.g. {The Summoning}, {Piper}; cards that **put** it in play do not, e.g.
{Summon History}, {Khazar's Diary (Endless Night)}.[^5-5-7]

An ally whose text lets it act the turn it is recruited still acts only in your minion phase, so recruiting it on an
opponent's turn gives it no action, e.g. {Nocturn}, {Infernal Servitor}.[^5-5-8]

An effect triggered by the ally entering play applies however it entered. {War Ghoul} burns one of your allies or
retainers on entry, and burns itself if you control none, even when put into play abnormally — it is an on-entry effect,
not a recruit requirement, so [§1.6](#16-requirements)'s bypass does not apply. Abilities keyed on an ally being
"recruited" cover non-standard recruitment routes, e.g. {Sébastien Goulet} with {Piper}.[^5-5-9]

### 5.6 Retainers

#### 5.6.1 The retainer is not the bearer

A retainer that uses a weapon is not its bearer; the employing minion remains the bearer.[^5-6-1] Damage or costs the
weapon puts on the bearer therefore fall on the employer, e.g. {Zip Gun}'s 1 damage during strike resolution. An option
the weapon grants the bearer is not available when the retainer used it: {Garrote} does not let the employer burn the
opposing vampire. A limit on the bearer's use does not limit the retainer — {Ghoul Retainer} may use {Jar of Skin
Eaters} with nothing on it.

A card that restricts a minion's use of equipment does not reach that minion's retainers, e.g. {Spiritual Protector}.

#### 5.6.2 Using a weapon is not a strike

The retainer's use resolves the weapon's whole text, not only damage: {Rowan Ring} sends to torpor, {Flash Grenade} ends
combat, {Weighted Walking Stick} burns its counters.

It is not a strike. It gives the employer no maneuver, no press and no additional strike, e.g. {Meat Hook}. A retainer
that does strike gains no additional strikes from its employer.

A retainer cannot play ammo cards, not even through a {Magazine}. It does benefit from ammo the employer played in an
earlier round, e.g. {Scattershot}.

#### 5.6.3 Damage from a retainer

Damage inflicted by a retainer is environmental damage and the retainer is its source. Immunity to retainers, or to that
type of retainer, stops it.[^5-6-2]

An effect conditioned on the retainer being ready still resolves if the opposing strike burns the retainer that round,
e.g. {Bestial Vengeance}.

#### 5.6.4 Retainers as targets

Cards keyed to actions directed at a minion do not apply to an action targeting its retainer, e.g. {Detect Authority}
[ani].[^5-6-3]

A secondary effect keyed to a minion does nothing when the target is a retainer, e.g. {Shadow Twin} at [OBT].

#### 5.6.5 Entering play

A retainer is usable the turn it is employed; the "cannot act this turn" restriction is on recruited allies
([§5.5.4](#554-entering-play-and-acting)).[^5-6-4]
A retainer whose card has inferior and superior versions is the version matching the Discipline level used to employ it,
and that does not change once it is in play, e.g. {Raven Spy}.[^5-6-5]

### 5.7 Special minion classes and traits

#### 5.7.1 Subtypes

A minion's subtype is what its type line prints, and nothing else. Ghouls are monsters and are not mortal; animals are
neither mortal nor monster.[^5-7-1] The card name is not a subtype: {Camarilla Vitae Slave} is a retainer with no
subtype, and {Gargoyle Slave} is an ally, neither a slave nor a Gargoyle. An effect naming a subtype, e.g. {Abjure}
ending a combat between a monster and a mortal, reads only that line.

#### 5.7.2 Slaves

A slave locking to take a blocked clan member's place cancels the combat, unlocks the acting vampire, and enters a new
combat with the blocker ([[RBK traits]](https://www.vekn.net/rulebook#traits)). The slave becomes the combatant but
never the acting minion ([§3.10](#310-changing-the-acting-minion)).[^5-7-2] It cannot play a card whose use requires it
to be acting, e.g. {Shadow Boxing} at [OBF][POT], and the blocker cannot play a card requiring the opposing minion to be
the acting vampire, e.g. {Obedience}. Effects naming the canceled combat are dead, e.g. {Sniper Rifle} cannot set the
range. The acting vampire keeps what it independently qualifies for after resolution, e.g. {Momentary Delay}.

#### 5.7.3 Infernal

The infernal unlock — the controller burning 1 pool during their unlock phase — is not normal unlocking, so an effect
preventing normal unlocking leaves it available.[^5-7-3] It is not tied to the unlock step: a minion locked by another
effect during that phase can still be unlocked by paying, e.g. {Nightmares upon Nightmares}. An effect keyed to
unlocking *with the infernal trait* does not fire on any other unlock effect, e.g. {Ruins of Charizel}. A card in play
requiring an infernal vampire is satisfied by one in torpor.

#### 5.7.4 Black Hand

Black Hand is a trait, not a title, and is unrelated to sect: a vampire printed Black Hand keeps it through a sect
change.[^5-7-4] A card that confers the trait writes it on a *Sabbat* vampire, so the grant stops applying if he ceases
to be Sabbat, e.g. {Cadet}, {Mustajib}. A vampire who gains the trait by resolving an action has it for cards played
after that resolution, e.g. {Seraph's Second}.

#### 5.7.5 Imbued

Imbued are mortal allies with life rather than blood, and their cost is their starting life ([[RBK
appendix-imbued-rules]](https://www.vekn.net/rulebook#appendix-imbued-rules)). Effects reading a minion's cost read that
starting life, e.g. {Keystone Kine} at [obf]; effects reading a cost to *recruit* read zero, since an Imbued is
influenced, not recruited, e.g. {Kindred Segregation}.[^5-7-5]

An Imbued in the uncontrolled region takes no blood, e.g. {Dreams of the Sphinx}.[^5-7-6] An effect moving a crypt card
to the ready region *with blood* gives an Imbued nothing: it arrives with no life and is incapacitated at once, e.g.
{Soul Gem of Etrius}. Where such an effect compares capacity blind, read the starting life as the capacity.

Incapacitation is neither burning nor torpor, so effects keyed to either do not fire, e.g. {Tension in the Ranks} burns
no pool.[^5-7-7] An incapacitated Imbued is still in play and may be locked to play a card, e.g. {Abjure}. A card that
rewrites the minion's subtype ends the class outright: an Imbued returned as a zombie ally by {Pressing Flesh} has no
creed, gains no conviction and cannot use virtues.

#### 5.7.6 Other traits

Scarce counts the apparent clan while the vampire is controlled, so a clan override changes what the influencing player
pays, e.g. {Clan Impersonation}; that same override does not remove slave status, which is a separate trait.[^5-7-8]
Sterile bars any action to put a vampire into play, including the last step of a multi-action process: the action
placing the counter that completes {Call the Great Beast} is such an action, and {Blood Cult Awareness Network} reads it
the same way. A restriction on cards *requiring* a discipline reaches only the option actually used, so a True Brujah
may play a multi-discipline card for a non-[cel] option.

A vampire follows at most one Path; following a new one replaces the old. During your influence phase, each Sabbat
vampire of yours following none may burn 1 blood to follow a Sabbat Path. A Sabbat Path benefits a vampire only while he
is Sabbat; no Path changes a clan.[^5-7-11] A card printing a Path icon requires a vampire following that Path, as a
clan icon requires that clan, e.g. {Absolute Tyranny} *(reading)*.[^5-7-12]

#### 5.7.7 Blood Brother circles

A vampire with no printed circle is in his own circle of one, and vampires sharing a name do not share a circle and
never contest one. Two uncontrolled copies of {Angelo} or {New Blood} are therefore in different circles, so an effect
reading "the same circle" requires choosing one; a vampire created by {The Embrace} is likewise in a circle of his
own.[^5-7-10]

### 5.8 Titles

#### 5.8.1 Printed titles, granted titles, contests

An Independent vampire's printed "N votes (titled)" is a title of its own and is tied to no sect. It and its votes
survive any sect change, e.g. {Ambrogino Giovanni}, {Xaviar}.[^5-8-1]

A title granted by a library card is tied to that title's sect. While the bearer is off-sect the title is dormant: he
counts as untitled and cannot use the title to make himself titled, e.g. {Fee Stake: Perth} on a vampire who is no
longer Anarch.[^5-8-2]

A card-granted title and a printed title of the same name are one title and contest each other: {Imperator} grants the
same unique Camarilla Imperator title that merged {Karsh} prints.[^5-8-3] Two vampires of one Methuselah claiming the
same unique title contest it, e.g. a {Toreador Justicar} referendum passed for one of his Toreador while he controls
{Madame Guil} *(reading)*.[^5-8-30]

Contesting a title costs 1 blood, paid by the vampire. Effects that raise the pool cost of contesting a card do not
reach it, e.g. {Democritus}.[^5-8-4]

A contested title-providing card is turned face down out of play, and any action it provides is unusable while the
contest lasts, e.g. {Fee Stake: Boston}, {Regent}.[^5-8-5]

Effects keyed to a title are re-evaluated continuously. Capacity granted by a title is lost with the title. A card keyed
to its bearer's title stops working when the title goes and resumes if it returns, e.g. {The Treatment}. An effect
conditioned on the vampire being untitled is suppressed while he holds a title, e.g. {Bloodbath}.[^5-8-6]

Votes granted by an ability are not a title. The vampire remains untitled, so an ability conditioned on his being
untitled keeps working, e.g. {Gerald Windham}, {Xeper, Sultan of Lepers}.[^5-8-7] See
[§1.15](#115-cumulative-and-stacking-effects) for how such votes combine with a title gained later.

A vampire who leaves play and returns remembers titles gained and lost.[^5-8-8] A vampire who gains a title loses the
one he held for good, a printed one included: a prince who becomes {Toreador Justicar} is untitled once that card leaves
play.[^5-8-25]

A representation of a titled vampire uses that title's votes even while another instance of the title is in play; it
does not contest.[^5-8-21]

#### 5.8.2 Off-sect and off-clan titles

A vampire whose sect becomes inappropriate for his title loses its benefit: he gets no vote, is not considered titled,
and yields the title immediately if it is contested or if he receives a new one. The benefit returns when his sect
becomes appropriate again.[^5-8-9] The rule is trait-general — an inappropriate clan change does the same, e.g. {Clan
Impersonation}, {Derange}.[^5-8-10] [§5.9](#59-traits-and-trait-changes) governs which of two competing trait changes
stands, and the "is considered" override that {Writ of Acceptance} applies.

Mass sect-removal events apply the same rule to every holder, e.g. {Fall of the Camarilla}.[^5-8-11] Losing the benefit
strips the title's votes and ballots but not the referendum structures the title defines: under {Fall of the Sabbat} the
priscus subreferendum still happens, and a printed bonus ballot not granted by the title still applies, e.g.
{Gratiano}.[^5-8-12]

An ability that grants its bearer a title cannot be used once he is off-sect: {Horatio Ballard} cannot call the
referendum to become Prince of Chicago while Independent.[^5-8-13]

#### 5.8.3 How requirements read titles

A requirement naming a sect-restricted title is also a requirement for that sect. A card requiring a baron is a card
requiring an Anarch, so {Open War} counts as Anarch-requiring for {Powerbase: Los Angeles}.[^5-8-15]

A title inside a vampire's name is not a title: {The Baron} cannot play cards requiring a baron.[^5-8-16] A permission a
card extends to Anarchs does not make the action one "requiring an Anarch", e.g. {Club Illusion}.[^5-8-17] Nor does a
negative requirement: a card requiring a non-Camarilla vampire is not a card requiring an Anarch, e.g. {Jericho
Founding} for {Garibaldi-Meucci Museum}.[^5-8-26] Prince, baron and archbishop are the titles tied to a city; primogen
and bishop are not.[^5-8-28] "A city title" on a card means one of the former, e.g. {Papillon}, {Mel, a Guerreira
Vibrante} *(reading)*.[^5-8-27]

A title is a vampire trait, so a card that fakes vampire traits reaches title requirements, e.g. {Vidal
Jarbeaux}.[^5-8-22]

A vampire who can meet title requirements ({Vlad Tepes}, {Vidal Jarbeaux}, {Kemintiri}) is also considered a member of
the title's underlying sect; the controller chooses how each requirement is met — any city, even an invented
one.[^5-8-18] {Vidal Jarbeaux}'s own text additionally caps each requirement at once per game: required to be "a titled
vampire" he must name a title, may use each title only once, and may choose the generic "X votes" title only once. A
card canceled as it is played has still spent that use.[^5-8-23]

Such a requirement can be met after the sect has ceased to exist, e.g. a Camarilla or justicar requirement under {Fall
of the Camarilla}.[^5-8-19]

An ability reaching cards that require a named title reaches a requirement that lists the title among others, but not
one the title would only satisfy: no longer Camarilla, {Karl Schrekt} can play {Parity Shift} ("a prince or justicar"),
not {Camarilla's Iron Fist} ("a titled Camarilla vampire").[^5-8-29]

Meeting a title requirement to call a political action grants none of that title's votes in the referendum.[^5-8-20]

Whether the ability reaches cards the Methuselah plays is per card: {Vidal Jarbeaux} meets requirements on master cards
too; merged {Kemintiri} does not enable master card plays.[^5-8-24]

See [§1.6](#16-requirements) for the limits of requirement-faking: it substitutes only for the effects of the card
played, only on a normal play, and only where the card prints the requirement ({Mata Hari}).

### 5.9 Traits and trait changes

#### 5.9.1 Actual change and temporary override

A card that says a vampire *becomes* a sect changes the actual sect, e.g. {Into the Fire}, {Out of the Frying Pan}. A
card that says a minion *is considered* a sect while the card stays in play is a temporary override on an unchanged
underlying sect, e.g. {Writ of Acceptance}, {Field Training}. The override has precedence over the actual sect. A card
that made the actual change and burns itself on a sect change burns when an override lands, e.g. {Go Anarch}.[^5-9-1]

The value an override writes is fixed when the card is played, and the underlying trait resurfaces when the override
leaves play. An effect writing one trait does not write another — {Clan Impersonation} changes clan and leaves sect
alone.[^5-9-2]

#### 5.9.2 Which effect governs

Where two effects write the same trait, the most recent governs. This is not sect-specific: {Clan Impersonation} and
{Derange} yield to whichever was played later, and {Nar-Sheptha} supersedes an earlier {Deep Song} on who is the acting
minion.[^5-9-3] If the governing override leaves play while an earlier one is still live, the earlier override
resurfaces as governing; the underlying trait returns only when no override remains.

An override survives a later change to the *underlying* trait. A vampire with {Writ of Acceptance} who takes the
rulebook action to become anarch is still Camarilla; so is an Assamite carrying a {Tegyrius, Vizier} allegiance
counter.[^5-9-4]

{Fall of the Camarilla} overrides sect but also redirects: a later effect that would set a vampire to Camarilla sets him
Independent instead while the Fall is in play. Such cards remain legal to play. When the Fall leaves play, underlying
sects resurface. {Fall of the Sabbat} is the mirror case.[^5-9-5]

#### 5.9.3 Reading a changed trait

A changed trait is the trait for all purposes, and effects keyed to it re-evaluate as it changes. Requirements are
checked continuously while an action is in progress: if {The Red Question} makes the acting vampire Anarch during an
action that required another sect, the action fizzles — successful but with no effect. {Ministry} grants its extra
intercept only while the acting vampire is Sabbat, gained or lost mid-action; {Teresita, The Godmother} likewise. An
effect applied after resolution reads the trait at that moment — {Warsaw Station} does not unlock an acting vampire who
is no longer Nosferatu when the unlock applies.[^5-9-6]

A vampire who plays cards "as if" he had a sect or title meets those requirements even when the sect has been abolished:
{Vlad Tepes} still plays Camarilla cards under {Fall of the Camarilla}. Conversely, writing a trait onto a vampire does
not make the action one *described* by that trait — {The Red Question} is not an action that "makes this vampire
anarch", so {CrimethInc.} cannot follow it.[^5-9-7]

Allies have no sect. An ally permitted to play cards "as an Anarch vampire" is not Anarch and gains nothing from cards
keyed to the trait, e.g. {Grey Thorne}, {Vivienne Géroux} do not benefit from {An Anarch Manifesto}.[^5-9-8]

---

## 6. Methuselahs and the Game

### 6.1 Owner and controller

**Base rules.** Ownership fixes which library, hand and ash heap a card belongs to. Control fixes who plays it, who
chooses, and what an oust removes. [[RBK
important-terms-of-the-game]](https://www.vekn.net/rulebook#important-terms-of-the-game)

#### 6.1.1 Zones belong to the owner

A card leaves play to its **owner's** ash heap, hand or library, whoever controlled it.[^6-1-1]

- A minion put into play from another Methuselah's ash heap returns to that ash heap when burned, e.g. {Khazar's Diary
  (Endless Night)}.
- A card held face down under a stolen card travels with the card, but still goes to its owner's library or ash heap if
  an effect moves it there, e.g. {Storage Annex}.
- A card drawn from another Methuselah's library returns to its owner's zone even if played or discarded by you, e.g.
  {Agaitas, The Scholar of Antiquities}.
- A vampire burned while under another Methuselah's control goes to his owner's ash heap, breaking every control effect,
  and from there to his owner's uncontrolled region. An effect printed for "you" still pays the Methuselah who
  controlled him when he burned, e.g. {The Capuchin}.

#### 6.1.2 Control decides what an oust removes

An oust removes what the ousted Methuselah **controlled**. Cards he owns but does not control stay in play. [[RBK
important-terms-of-the-game]](https://www.vekn.net/rulebook#important-terms-of-the-game)

Only printed text keeps control elsewhere. {Shackles of Enkidu} goes on the opposing minion but stays with the
controller of the vampire that used it, and changes hands with that vampire. Taking control of a card already in play is
[§6.3](#63-taking-control-of-a-card-in-play)'s.

A vampire in the uncontrolled region belongs to his last permanent controller, who need not be his owner; see
[§6.5.3](#653-cards-of-an-ousted-methuselah) for what an oust then removes.

#### 6.1.3 Control decides who plays and who chooses

The acting minion's controller makes the choices the action calls for, even when another Methuselah steals its outcome:
on a blocked recruit or employ action the acting minion's controller still decides how the ally or retainer enters
play, e.g. {Set's Call}.[^6-1-3]

A minion card is played by the Methuselah as well as the minion, so an effect reducing the cost of "a card you play"
reaches it, e.g. {The Shard, London}. That Methuselah pays.

#### 6.1.4 Cards no one controls

A vampire in the uncontrolled region, or yielded in a contest, is not controlled and his abilities do not apply, e.g.
{Byzar}.[^6-1-4] He is not in play either — in play means controlled, and card effects reach the uncontrolled region
only when their text names it [[RBK targeting-of-cards]](https://www.vekn.net/rulebook#targeting-of-cards) — so an
effect keyed to a vampire burned *from play* does not reach him, e.g. {Blessed Resilience}.

Card text naming cards in play is not limited to cards you control unless it says so, e.g. {Spell of Life} burns copies
controlled by anyone.

### 6.2 Taking control of minions

**Base rules.** An ousted Methuselah's cards are removed from the game; cards he owns but does not control stay in play.
[[RBK 5-ending-the-game]](https://www.vekn.net/rulebook/5-ending-the-game)

#### 6.2.1 When the control change happens

A card that lets you burn counters to steal a minion offers the steal only at its own timing; placing a counter does not
open the opportunity, e.g. {Velvet Tongue}.[^6-2-1] Where the card names a phase, the taking may happen at any point in
it, including after other actions, e.g. {Puppet Master}.

A reversion running until the unlock phase happens at the start of that phase, before unlocking. It cannot be ordered
among unlock effects, and the minion does not unlock that phase, e.g. {Malkavian Dementia}.

A steal does not change the minion's lock state; a card that wants it unlocked prints the unlock, e.g. {Temptation}.
Cards on the minion travel with it.

#### 6.2.2 Stealing mid-action and mid-combat

Stealing a minion that is acting does not end the action. It continues against the same target, and Methuselahs who
already declined to block get no second opportunity; the new controller, if eligible and not yet declined, may still
block with the stolen minion.[^6-2-2][^6-2-3]

If the target no longer qualifies after the steal, the action still resolves and succeeds, with no effect and its cost
unpaid. The same holds when the target empties the action in response — {War Ghoul} locking and burning itself to burn a
location.

Stealing a minion that is attempting to block makes the block attempt fail. Stealing the acting minion instead leaves
the block standing, e.g. {Revelation of Despair}. Stealing a minion in combat ends the combat, but damage still to
resolve resolves before control changes, e.g. {Temptation}.

An effect that stops the control change stops only that: the cost paid to attempt the steal is not refunded, e.g. {The
Diamond Thunderbolt}.

#### 6.2.3 Return to the previous controller

The minion returns in the state it is in, including torpor, and stays in torpor, e.g. {Temptation}.[^6-2-4] Canceling an
action the stolen minion was forced to take does not keep it under the new controller; it still returns, e.g. {Spirit
Marionette}. "After resolution" effects apply before the return.

"Breaking any temporary control effects" hands the minion to its last permanent controller, who then receives the rest
of the effect and any card placed on it — {Lay Low} into that Methuselah's uncontrolled region, {Descent into Darkness}
out of play.

#### 6.2.4 An ousted Methuselah

A minion held on temporary control is removed from the game if its temporary controller is ousted, e.g. {The Ailing
Spirit}.[^6-2-5] It is likewise removed at the moment it should return to a Methuselah who has been ousted, e.g.
{Temptation}.

If that return can no longer occur, control never reverts and the temporary controller keeps the minion indefinitely:
{Parmenides} returns at his owner's next unlock phase, which an ousted owner never has.[^6-2-6] What an oust removes
under permanent control is [§6.5.3](#653-cards-of-an-ousted-methuselah)'s.

#### 6.2.5 Effects after a control change

An effect that charges or credits "the controller" follows the new one: {Betrayer}'s pool loss is taken by whoever
controls the named vampire.[^6-2-7] Who takes or declines an optional effect follows the card's wording, not control of
the card: {The Rack} lets the chosen vampire gain blood, so the new controller of a stolen vampire decides whether it
gains.

A duration or use limit written against "your" turn or unlock reads against the current controller and is not refreshed
by the control change: {Rutor's Hand} grants no extra unlock until a Methuselah begins a turn controlling the bearer,
and a spent per-phase use stays spent. An effect already generated survives the change: {Imposing Phantasm} returns the
blood lost to damage even if the opposing minion has changed controller.

### 6.3 Taking control of a card in play

#### 6.3.1 Requirements are not checked

Taking control of a card already in play is not playing it, so the card's requirements do not apply to the new
controller.[^6-3-1] A vampire may steal a retainer whose discipline requirement it does not meet, e.g. {Far Mastery},
and a Methuselah may take a location whose requirement no minion of his satisfies, e.g. {New Management}. The card stays
in play and works normally.

#### 6.3.2 Where the card is placed

An equipment, retainer or location on a minion whose control changes is placed on a minion the new controller controls;
the new controller chooses which one.[^6-3-2] If the new controller has no minion to take it, it is burned. This is the
rule for locations that are equipment — the {Dartmoor, England} template — as much as for ordinary equipment and
retainers.

Placement happens only on an actual change of controller. A card "taken control of" by the Methuselah who already
controls it stays where it is and cannot be moved to another of his minions.[^6-3-3] Naming your own card is still a
legal play for a card that takes control, e.g. {New Management} on your own location. A card that steals cannot name it:
you cannot steal a card from yourself, so {Saankaláxt} cannot take an equipment from another of your minions.[^6-3-7]

#### 6.3.3 Cards carried along

Control of a minion card on a minion follows control of the minion: take the vampire and you take its equipment and
retainers.[^6-3-4] A card hosted on another card goes with its host, e.g. the face-down card on {Storage Annex}, and the
new controller may use it. A stolen master Discipline card is controlled by the Methuselah of the vampire that took it,
e.g. {Ethan Locke}.

Where a card goes when it later leaves play is set by ownership, not control.

#### 6.3.4 What does not move

A minion written into the card's effect when it was played is fixed and is not re-chosen.[^6-3-5] {Incriminating
Videotape} keeps the minion chosen at play: once the equipment is stolen, that minion cannot block the new bearer and
can block the old one. Only the minion the card sits on is re-designated (6.3.2), never a choice already made by the
effect.

#### 6.3.5 When a forced transfer happens

Some cards hand themselves to another Methuselah during the controller's discard phase. Where the transfer rides on an
optional choice being taken, the card moves as soon as that choice is made, ahead of any mandatory effects still
pending, e.g. {Scourge of the Enochians} when the burn option is used.[^6-3-6] Otherwise the receiving Methuselah takes
the card at his first opportunity to play effects, after the acting Methuselah's other mandatory effects have applied. A
controller who declares the end of his turn cannot then use the card, even where the transfer was overlooked and the
card is still in front of him.

### 6.4 Leaving and re-entering play

**Base rules.** Contested cards are face down and out of play; a vampire in the uncontrolled region is not in play and
is not a legal target. [[RBK contested-cards]](https://www.vekn.net/rulebook#contested-cards)

#### 6.4.1 Set aside out of play — the card remembers everything

A card moved out of play without reaching the ash heap keeps its state. The wording template is {Lay Low} and
{Banishment}: cards and counters stay with the vampire but are out of play while it is.

On return the vampire remembers every effect that had been applied to it, exactly as a contested vampire does, including
gained or lost titles.[^6-4-1] A vampire contested while in torpor returns to torpor if he wins *(reading)*.[^6-4-9]
Cards on it return unlocked, and a contested card on it drops out of the contest until it comes back. A vampire brought
out of the uncontrolled region keeps its blood and its cards.

Leaving play breaks temporary control effects: the vampire goes face down to its permanent controller, and so does the
card that removed it.[^6-4-2] See [§6.2](#62-taking-control-of-minions) for control changes; if the permanent
controller is ousted before the return condition is met, the current controller keeps it.

#### 6.4.2 Through the ash heap — a new card

A card that reaches the ash heap and comes back is a new card: continuous effects applied to it before are lost, e.g.
{Possession}, and a returned ally or retainer is a new ally or retainer for all game purposes, e.g. {Compel the
Spirit}.[^6-4-3] See [§3.5](#35-action-repetition-nra-and-canceled-actions) for the No Repeated Action consequence.

#### 6.4.3 Effects aimed at a card that has left play

An ongoing effect pointed at a card is suspended, not lost, while that card is out of play, and resumes if and when it
returns. This includes effects worded "for the rest of the game". {The Rack} tracks its chosen vampire through a contest
or a {Banishment} and gives blood again on the vampire's return.[^6-4-4]

While the card is out of play the effect does nothing, because the card is not controlled. Torpor is not out of play: an
effect naming a controlled vampire still reaches one in torpor. A pending effect that would land after the card has left
play does not land, e.g. the {Daring the Dawn} damage on a vampire already removed by {Descent into Darkness}.

Effects the card already produced survive it leaving play ([§2.5.3](#253-effects-outliving-their-source)). Conversely a
card in play reaches only events after its arrival: {NRA PAC} does not affect equip actions performed before it entered
play.[^6-4-5] An in-play effect card is removed when the Methuselah who played it is ousted.

#### 6.4.4 Leaving the ready region as a trigger

An effect keyed to a minion leaving the ready region fires whatever the route out — burned, or the controller ousted —
and fires even if the triggering card is burned in the same event, e.g. {The Black Throne}, {Priority Contract}.[^6-4-6]

#### 6.4.5 Entering play

A card returns at the moment its own text names and is not in play before then. {Parmenides} returns as the unlock phase
begins and unlocks as normal; a contested unique won face down turns face up during the next unlock phase.[^6-4-8]

### 6.5 Pool, the Edge and ousting

#### 6.5.1 Gaining and losing pool

Any decrease in pool — burning, paying a cost, a rival taking it — is losing pool; any increase is gaining pool.[^6-5-1]
An effect measured on pool actually lost counts nothing when none was lost, e.g. {Dirty Little Secrets}.

A gain never offsets a loss: where one effect grants then takes pool, a card keying on pool lost sees the loss, not the
net, e.g. {Poison Pill} against {Ancient Influence}.[^6-5-2]

A Methuselah obliged to pay more pool than she holds pays all of it and is ousted; the effect still happens, e.g.
{Thanks for the Donation}. The shortfall is real: an effect saving her must cover the whole amount, not bring her to
zero, e.g. {Life Boon}. A mandatory cost is paid even when paying ousts you.[^6-5-4]

Withdrawal fails on any pool loss, including pool spent as a cost, and a later gain does not repair it [[RBK
5-ending-the-game]](https://www.vekn.net/rulebook/5-ending-the-game).

#### 6.5.2 When the oust resolves

An oust resolves as soon as pool reaches zero, ahead of anything else the same event triggers. A minion spending its
last life to pay for an action still resolves it in full, e.g. {Herald of Topheth}, and the oust precedes the effects
triggered by the minion burning. If a cancel's cost ousts a player, handle the oust before the canceled card.[^6-5-5] A
Methuselah ousted by the pool he pays for votes casts none of them, e.g. {Ferraille}.[^6-5-13]

Referendum-caused ousts, and what effects keyed to the referendum may do around them, are
[§3.7.5](#375-referendum-procedure)'s.

When a card ends the turn on an oust ({Last Stand}), the turn ends after the current action concludes: effects usable
after action resolution are still played, e.g. {Freak Drive}. The next turn then begins immediately, with no discard
phase and no replacement.[^6-5-7]

#### 6.5.3 Cards of an ousted Methuselah

Cards controlled by an ousted Methuselah are removed from the game, not burned, so burn triggers do not fire
([§1.10.1](#1101-burn-is-not-removal-from-the-game)), e.g. {Charnas the Imp}; an ongoing effect such a card imposed is
canceled with it, e.g. {The Meddling of Semsith}. A vampire
in another Methuselah's uncontrolled region is removed if that Methuselah is ousted and stays if only its owner is, e.g.
{Lay Low}; a stolen minion due to return to an ousted one is removed.[^6-5-8] See
[§6.4](#64-leaving-and-re-entering-play) for control and [§6.2](#62-taking-control-of-minions) for theft.

Leaving the ready region because your controller is ousted is still leaving the ready region, so effects keyed to that
fire, e.g. {Priority Contract}. The minions are gone at once, so a later effect needing one has no object: {Revelation
of the Serpent} cannot unlock when the target is ousted by the bleed.[^6-5-9]

#### 6.5.4 Victory points and the Edge

A Methuselah ousted at the same moment as her prey still scores the victory point but takes no pool, and pool comes only
from your own prey, never a grand-prey. Where two effects are available at once you choose the order: the victory point
before or after the 6 pool, or your unlock-phase Edge pool before burning the Edge, e.g. {Sabbat Threat}.[^6-5-10]

A restriction on gaining pool is tested when the pool is gained, not when the action is announced. The Edge follows the
bleed's resolution: pool granted as part of the bleed arrives before it, pool from a separate post-resolution effect
after it. That decides both halves of {The Rising}.[^6-5-11]

The moment an action moves or burns the Edge is its card's. {Sargon}'s Edge gain is worded after the action and applies
after resolution — after the referendum concludes, for a political action; {Hrothulf} burns the Edge at resolution, only
if the action succeeds. A steal of the Edge is not thwarted by the holder burning it after announcement: the stealer
still gets it at resolution, e.g. {Tereza Rostas}. A card that burns your own Edge does nothing if you no longer control
it at resolution, while the rest of its text still applies, e.g. {Enticement}. Legality is checked on announcement: the
Edge cannot be stolen while uncontrolled. "A new Methuselah gets the Edge" means one who did not have it, e.g. {Curse of
Nitocris}.[^6-5-12]

### 6.6 Master phase

**Base rules.** You receive one master phase action per master phase, unused actions are lost, and you choose the order
in which your master phase actions and any other master-phase effects happen [[RBK
master-phase]](https://www.vekn.net/rulebook#master-phase). An out-of-turn master card spends an action from your next
master phase [[RBK master-cards]](https://www.vekn.net/rulebook#master-cards).

#### 6.6.1 Master phase action accounting

Cancellation does not unwind the accounting. A canceled out-of-turn master still costs its player the master phase
action against their next master phase, and that player still cannot play a second out-of-turn master before that
phase.[^6-6-1] A canceled trifle grants no additional master phase action.[^6-6-2] A unique trifle that enters contest
as it is played still grants it, e.g. {Agent of Power}.[^6-6-8]

An additional master phase action handed to a Methuselah by a card effect is not that Methuselah's trifle bonus. A
trifle they play afterwards still grants its one, e.g. {Wash}.[^6-6-3] The one-per-phase cap counts only actions gained
from trifles [[RBK trifle]](https://www.vekn.net/rulebook#trifle).

#### 6.6.2 Out-of-turn masters

One out-of-turn master card between two of your master phases, regardless of how many master phase actions you have
[[RBK master-cards]](https://www.vekn.net/rulebook#master-cards). An ability that grants access to out-of-turn masters,
or that pays for an extra master phase action, does not raise the cap, e.g. {Synesios}.[^6-6-4]

An out-of-turn master played during your own turn under a card's own permission counts against your **next** master
phase, and bars a second out-of-turn master until then, e.g. {Proxy Kissed}.

An out-of-turn master is played from hand in the normal fashion and can be canceled as it is played, even where the
card's text sends it from hand to the ash heap rather than into play, e.g. {Vox Senis}.[^6-6-5]

#### 6.6.3 Master cards on minions

A master card in play is controlled by the Methuselah who played it, even when it sits on another Methuselah's minion
[[RBK master-cards]](https://www.vekn.net/rulebook#master-cards). Who takes or declines a choice its effect offers
follows the card's wording, not control of the card, e.g. {Perfectionist}
([§1.1.1](#111-mandatory-and-optional-effects)).

An upkeep printed as an alternative is a real choice, even when one branch burns the minion: {Ex Nihilo} lets the
vampire's controller burn 1 blood or burn the vampire during the master phase.[^6-6-7]

### 6.7 Influence phase and the uncontrolled region

**Base rules.** Transfers are granted at the start of your influence phase and cannot be saved. The uncontrolled region
is not in play.

#### 6.7.1 The uncontrolled region as a zone

A vampire in the uncontrolled region is not a legal target unless the card names the region, and it cannot play cards,
e.g. {Reform Body}.[^6-7-1]

An effect that adds blood to "a vampire in your uncontrolled region" cannot be used on an imbued there — an imbued is
not a vampire, e.g. {Dreams of the Sphinx}. When the crypt card is chosen blind, from a crypt or from another
Methuselah's uncontrolled region, the effect works with whatever is found, and an imbued's cost counts as its capacity.

Suspended capacity modifiers are [§5.1.1](#511-capacity)'s; created-vampire identity in the uncontrolled region is
[§5.1.3](#513-created-vampires-and-crypt-card-identity)'s; what else the vampire remembers is
[§6.4](#64-leaving-and-re-entering-play)'s.

A vampire moved to the uncontrolled region while under a temporary control effect goes to its **last permanent
controller's** region, which need not be the owner's, e.g. {Banishment}, {Lay Low}; ousts are
[§6.5.3](#653-cards-of-an-ousted-methuselah)'s.[^6-7-3]

#### 6.7.2 Influencing out

A vampire whose blood equals or exceeds its capacity may be moved to the ready region at any moment during its
controller's influence phase, not only at the end.[^6-7-4]

Excess blood drains off as it enters play — before a contest is entered, and before any "as he enters play" effect is
played. A vampire may be influenced out into a contest it will lose.

#### 6.7.3 Transfers and the influence-phase window

The number of transfers is fixed when the phase begins. A vampire influenced out during the phase adds no transfers to
it, e.g. {Ingrid Rossler}.[^6-7-5] Its other abilities are usable for the remainder of that phase, e.g. {Angela
Preston}.

An effect granting an extra transfer may be used at any moment during the phase, not only before the first transfer is
spent, e.g. {Ennoia's Theater}.

"During **your** influence phase" restricts a card to its controller's phase; "during **the** influence phase" makes it
usable in any Methuselah's, e.g. {Gather}.

Transfers and card effects are independent. A Methuselah who takes no transfers, or who has lost them all, may still use
effects that add blood to uncontrolled vampires, e.g. {Powerbase: Montreal}.

#### 6.7.4 Cards that require an uncontrolled vampire

A card whose play clause names a vampire in the uncontrolled region cannot be played with no qualifying vampire there;
the vampire is chosen on announcement, e.g. {Undue Influence}, {Break the Bonds} at [pre].[^6-7-6] The same bars an
ability worded that way, e.g. {Lázár Dobrescu}.

A clause that only describes what the effect does to the region is not a play requirement. {Social Ladder} may be played
with no older vampire in the uncontrolled region; its influence-phase clause then removes the ready vampire and the
blood is lost.

### 6.8 Hand, draw and discard

**Base rules.** Default hand size is seven. Whenever an effect changes your hand size, or adds or removes cards from
your hand, you immediately discard down or draw up to match [[RBK
drawing-cards]](https://www.vekn.net/rulebook#drawing-cards).

#### 6.8.1 Hand size

Cards you play but do not replace count against your hand size, so a "do not replace" clause works as a hand-size
reduction: any further draw forces a discard, e.g. {Hagar Stone}.[^6-8-1] Non-replacements accumulate while the effect
lasts. A hand-size change that comes and goes repeats the adjustment each time, e.g. a fresh combat against {Raptor}
makes the opponent draw up, then discard again.

A hand-size bonus can be taken more than once in a turn if you have more than one master phase action, e.g. {Edward
Neally}.[^6-8-2] It lasts as long as its source is in play: {The Meddling of Semsith} keeps reducing your hand size
after the chosen Methuselah is ousted, and {Nahir}'s bonus survives her going to torpor. A card in play that draws or
raises your hand size may be used during an action, e.g. {The Barrens}, {Dreams of the Sphinx}.

#### 6.8.2 The discard phase

"Until end of turn" durations last until the turn ends, once every Methuselah has passed in the discard phase. Their
expiry is not ordered: discard phase actions and other discard-phase effects come first, e.g. {Dreams of the
Sphinx}.[^6-8-3] Discard-phase effects apply before an ally taken until end of turn goes back; an ally you already lost
control of does not go back at all.

Each discard phase action carries its own use of an ability keyed to one, e.g. {Josef von Bauren}. A Methuselah ousted
during their turn gets no discard phase. A discard imposed by an action card lands after the action resolves and after
any combat it caused.

#### 6.8.3 Draw and discard mechanics

When an effect discards several cards, all discards are made at once and replaced afterwards, e.g. {Ruxandra}.[^6-8-4]
Replacement draws therefore cannot be fed back in as further discards. An effect that repeats a draw-and-discard
resolves one card at a time, e.g. {Infernal Pursuit}. An effect that redirects the discarded cards acts only once every
discard and redraw is done, e.g. {Rachel Brandywine} shuffling them back. A discard you cannot make in full is made as
far as you can.

"Draw, then discard" does not restrict the discard to the cards just drawn; discard any card in hand.[^6-8-5] A trigger
reading "each time you replace a card" sees every replacement draw, whatever took the card out of hand, e.g. {Infernal
Pursuit}. A draw up after a hand-size increase is not a replacement, e.g. for {Visit from the Capuchin}
*(reading)*.[^6-8-7] An effect letting you draw a stored card "if you would draw a card from your library" applies card
by card, and not at all from an empty library, e.g. {Shilmulo Tarot}, {Black Market Cache} *(reading)*.[^6-8-8] A
trigger keyed to replacing a card you *played* does not see extra draws granted by other cards: {Agaitas, The Scholar of
Antiquities} does not redirect a {Learjet} additional draw. Cards returned to an empty library are drawn at once if your
hand is below hand size, e.g. {Waste Management Operation}; that refill precedes any random discard the same card
imposes.

#### 6.8.4 Visibility

Library searches and replacement draws are private. You do not show the cards you look at while searching, e.g. {Vast
Wealth}, and no one sees a card drawn as replacement, e.g. {Vaticination}.[^6-8-6] Where the cards are already public
knowledge, you show which one goes to hand, e.g. {Ashur Tablets}. An effect that reads a hand after the replacement draw
sees it without one when the replacement is delayed, e.g. {Troglodytia} against a master card canceled by {Wash}; an
optional extra draw is decided after the normal replacement has been seen, e.g. {Learjet}.

---

---

## Appendix A. Glossary of card wordings

Idiomatic card wordings, each with the section that governs it. Quotation marks mark printed wording; "e.g." names
exemplar cards.

- **"as if from your hand"** — a normal, cancelable play, but the card is not in hand and effects requiring one there
  cannot reach it. [§1.8.2](#182-played-but-not-in-the-normal-fashion), [§1.14.2](#1142-cards-named-from-hand)
- **"as if he had that required title"** — {Vlad Tepes} and {Kemintiri} meet title requirements and count as members of
  the underlying sect, any city (even invented); only {Vidal Jarbeaux}'s printed text caps each requirement, each
  title, and the generic "X votes" title at once per game. [§5.8.3](#583-how-requirements-read-titles)
- **"as if she were of that required sect and/or clan"** — treated as meeting the requirement for all effects of the
  card played, and only them: never for in-play text, never on an abnormal play, never for a trait the card merely
  acts on (e.g. {Mata Hari}). [§1.6.4](#164-vampires-that-meet-requirements-they-do-not-have),
  [§5.8.3](#583-how-requirements-read-titles)
- **"as long as … / until …"** — the printed duration clause decides survival: tied to a condition ("as long as he or
  she remains in torpor") the effect outlives its card; stated flatly in in-play text it ends when the card leaves
  (e.g. {Rowan Ring}, {Wooden Stake}). [§2.5.3](#253-effects-outliving-their-source)
- **"cannot be prevented by cards requiring …"** — a "cannot" naming the cards it forbids bars playing them entirely,
  even to cycle ("cannot use any additional strikes" is read the same way); dodging the strike defeats the disabling
  (e.g. {Blood Fury}, {Rigor Mortis}). [§1.1.3](#113-cannot), [§4.6.1](#461-a-dodge-negates-the-whole-strike)
- **"do not replace until"** — the delayed card is drawn at the first moment its condition is met, ahead of anything
  else triggering then, even if the card was since burned; canceled with the card, and governing only normal plays.
  [§1.9.2](#192-delayed-replacement), [§1.9.3](#193-abnormal-entry-and-cancellation)
- **"does not unlock as normal"** — suppresses only the automatic unlock; wakes and unlock effects still reach the
  minion. [§5.2.2](#522-the-unlock-phase-and-its-suppression)
- **"even if stealth/intercept is not (yet) needed"** — the printed waiver of the not-needed bar; as-announced stealth
  always carries one. [§3.2.1](#321-adding-what-is-not-needed), [§2.2](#22-as-the-action-is-announced-effects)
- **"gets X" vs "+X"** — an absolute value sets (copies do not stack, and a set can reduce); "+X" adds to whichever base
  is current. [§1.15](#115-cumulative-and-stacking-effects), [§4.4.1](#441-how-much-damage)
- **"inflicts" vs "successfully inflicts"** — "successfully inflicts" counts only damage left unprevented and not
  ignored; "inflicts" also counts damage prevented or ignored through immunity, a printed "(even if prevented)" being
  reminder text, except on older cards a ruling reads as "successfully inflicts". Neither counts damage dodged or
  reduced away. [§4.5.5](#455-immunity)
- **"is considered the acting minion"** — only such text rewrites the designation; entering the combat does not, a later
  writer supersedes an earlier, and the designation reverts when the writer leaves play.
  [§3.10.2](#3102-is-considered-the-acting-minion)
- **"put … in play"** — the only wording that bypasses requirements and cost (required disciplines at inferior, X = 0);
  prohibitions still apply, and the card has still been played, though not in the normal fashion.
  [§1.6.3](#163-cards-entering-play-abnormally), [§1.8.2](#182-played-but-not-in-the-normal-fashion),
  [§3.7.4.2](#3742-entry-into-play-other-than-by-the-action)
- **"queue a combat"** — sets up a combat that begins when the current combat ends, or once the action has resolved,
  still as part of its resolution; one at a time, and every effect starting a new combat in the action shares the queue
  (e.g. {Hedonism}, {Psyche!}). [§4.9.3](#493-queuing-a-new-combat)
- **"represents a location and does not count as equipment while in play"** — equipment in hand, both types as played,
  only a location in play, so equipment-movers and equip bars miss it; on a change of control it is re-placed on a
  minion of the new controller (or burned), like ordinary equipment (e.g. {Dartmoor, England}).
  [§1.3.1](#131-locquipments), [§6.3.2](#632-where-the-card-is-placed)
- **"stun"** — lock the minion and put a stun counter on him; with one or more, he does not unlock as normal at the
  beginning of his controller's unlock phase, which burns the counters he had at the beginning of the turn (e.g.
  {Mind Numb}). [§5.2.2](#522-the-unlock-phase-and-its-suppression)
- **"this vampire can" vs "you may"** — the wording names the decider of an optional effect: "this vampire" → the
  vampire's current controller, "you" → the card's controller. [§1.1.1](#111-mandatory-and-optional-effects)
- **"unlock and attempt to block"** — unusable when the vampire could not attempt anyway; lifts no prey/predator/target
  restriction; usable by one already attempting; the block stands even if the unlock removes the action's target
  (e.g. {Guard Duty}, {Second Tradition: Domain}). [§3.3.1](#331-who-may-block)
- **"will not go to torpor until combat ends"** — pushes torpor to after combat: the vampire stays wounded and
  burnable, the trigger's other effects still happen, and the protection does not reach a fresh combat (e.g.
  {Undying Tenacity}, {Undead Persistence}). [§5.3.1](#531-going-to-torpor)
- **"would X … instead"** — a replacement effect: played only while X is still pending, so it cannot follow an effect
  keyed to X "about to" happen or having happened. [§1.2.2](#122-trigger-and-condition-wording),
  [§2.4.1](#241-ordering-within-a-window)

## References

[^1-1-1]: [[RTR 19980707]](https://usenet.krcg.org/t/t3BWHkOrdyE/#m0) [[LSJ
    19980722]](https://usenet.krcg.org/t/0kbbA-SNchg/#m23) — {Aaron's Feeding Razor}, {The Ancestor's Talisman},
    {Changeling Skin Mask}, {Enchanted Marionette}, {Writ of Acceptance}.

[^1-1-2]: [[PIB 20121028]](https://usenet.krcg.org/t/vekn-39816/#m10)
    [[LSJ 20051016]](https://usenet.krcg.org/t/s_At5syL66k/#m6) [[LSJ
    20051017]](https://usenet.krcg.org/t/s_At5syL66k/#m8) [[LSJ 20051211]](https://usenet.krcg.org/t/TuwXiJ8A9mo/#m1) —
    group "Mandatory additional strike" (G00134), {Lorrie Dunsirn}, {Eze, The Demon Prince}, {Renegade Garou}.

[^1-1-3]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[RTR
    20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[ANK
    20180420-2]](https://usenet.krcg.org/t/vekn-76522/#m5) [[LSJ
    20020304-2]](https://usenet.krcg.org/t/L-8OGYP5xsE/#m6) [[PIB
    20110918]](https://usenet.krcg.org/t/vekn-10458/#m1) — {Repo Man}, {Vast Wealth}, {Serpent's
    Numbing Kiss}, group "Discounts for rescue" (G00075).

[^1-1-4]: [[LSJ 19971215]](https://usenet.krcg.org/t/mfBmRrUKZQ0/#m8) [[TOM
    19951129]](https://usenet.krcg.org/t/jjBzopH-yrQ/#m3) — {Bomb}, {Bundi}, group "Wake" (G00115).

[^1-1-5]: [[PIB 20150418]](https://usenet.krcg.org/t/vekn-70589/#m1) [[ANK
    20210309-3]](https://usenet.krcg.org/t/vekn-79065/#m4) [[RBK
    important-terms-of-the-game]](https://www.vekn.net/rulebook#important-terms-of-the-game) — {Third Tradition:
    Progeny}; group "Master on vampire who can use it" (G00161), {Perfectionist}, {Corporal Reservoir}.

[^1-1-6]: [[LSJ 20100723]](https://usenet.krcg.org/t/0u5KQWiutdg/#m1) [[PIB
    20111002]](https://usenet.krcg.org/t/vekn-8235/#m19) [[LSJ
    20021210-1]](https://usenet.krcg.org/t/kOZf54CTBUU/#m2) — {The Coven}, {Owain Evans, The Wanderer}, {Leandro}.

[^1-1-11]: Owner's reading, no official ruling: [[rules-qa
    33.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-33--pool-masters-wave-1)
    [[rules-qa
    57.5]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-57--legacy-cards-for-twda-decks-staples-steps-13-implemented-readings-2026-10-08)
    — {Brujah Debate}, {Dragonbound}.

[^1-1-7]: [[LSJ 20050628]](https://usenet.krcg.org/t/lN3eieA3xgs/#m4) [[LSJ
    20011214-5]](https://usenet.krcg.org/t/gI44SEC82Yk/#m2) — groups "Prevent discipline based prevention" (G00141) and
    "Prohibits additional strikes" (G00137), {Blood Fury}, {Blood Rage}, {Dead Hand}, {Soul Burn}, {Rigor Mortis}.

[^1-1-8]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[LSJ
    20010809-1]](https://usenet.krcg.org/t/LB6Zg4bEggc/#m5) — {Charnas the Imp}, group "Steal blood as a Hunt action"
    (G00121).

[^1-1-15]: [[ANK
    20221019-2]](https://usenet.krcg.org/t/vekn-80101/#m2)
    — group "Cannot gain blood" (G00163).

[^1-1-10]: [[ANK 20180518]](https://usenet.krcg.org/t/vekn-76623/#m3) [[LSJ
    19990218]](https://usenet.krcg.org/t/B27J0sAwUuw/#m22) — {Big Game}, {Major Boon}.

[^1-1-12]: Owner's reading, no official ruling: [[rules-qa
    32.7]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-32--pool-actions-wave-2)
    [[rules-qa
    49.5]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-49--reign-of-stanislava-25th-anniversary-and-sabbat-paths-scoping-2026-10-05)
    — {Open War}.

[^1-1-14]: [[RBK crypt-cards]](https://www.vekn.net/rulebook#crypt-cards).

[^1-1-13]: Owner's reading, no official ruling: [[rules-qa
    64.9]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-64--legacy-cards-for-twda-decks-rounds-45-the-gangrel-decks-implemented-readings-2026-10-11-owner-2026-10-14-647-648-confirmed-2026-10-20-the-rest-646-restated)
    — {Read the Winds}.

[^1-2-1]: [[ANK 20180719-4]](https://usenet.krcg.org/t/vekn-76833/#m1) [[ANK
    20230122]](https://usenet.krcg.org/t/vekn-79926/#m4) [[ANK
    20190207]](https://usenet.krcg.org/t/vekn-77344/#m22) [[ANK
    20220617]](https://usenet.krcg.org/t/vekn-79835/#m2) [[ANK
    20200729]](https://usenet.krcg.org/t/vekn-78776/#m1)
    [[RBK wording-templates]](https://www.vekn.net/rulebook#wording-templates) — group "Once each turn" (G00014),
    {Maila}, {Qawiyya el-Ghaduba}, {Nonu Dis}, {Carver's Meat Packing and Storage}.

[^1-2-17]: Owner's reading, no official ruling: [[rules-qa
    16.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-16--printed-abilities-after-an-action)
    — {Richard de Worde}; cf. {Andre LeRoux} [[LSJ 20081123]](https://usenet.krcg.org/t/_Pb29pBJ1kU/#m2).

[^1-2-2]: [[LSJ 20081123]](https://usenet.krcg.org/t/_Pb29pBJ1kU/#m2) [[LSJ
    20001102-2]](https://usenet.krcg.org/t/LlPyLJjLdx0/#m5) [[ANK
    20191218]](https://usenet.krcg.org/t/vekn-62700/#m7) [[ANK
    20220617]](https://usenet.krcg.org/t/vekn-79835/#m2) [[LSJ
    20050215]](https://usenet.krcg.org/t/nxlYFsU10Uo/#m1) [[ANK
    20230116]](https://usenet.krcg.org/t/vekn-80266/#m1)
    [[LSJ 20010111]](https://usenet.krcg.org/t/d3WSV1UXBV0/#m5) [[RBK
    wording-templates]](https://www.vekn.net/rulebook#wording-templates) — {Andre LeRoux}, {Courier}, {Nahir}, {Nonu
    Dis}, {NSA Trio}, {Phased Motion Detector}, {Angelica, The Canonicus}.

[^1-2-3]: [[LSJ 20020814]](https://usenet.krcg.org/t/gt8wQhk76lA/#m1) [[ANK
    20180307-2]](https://usenet.krcg.org/t/vekn-76451/#m1)
    [[LSJ 20040617]](https://usenet.krcg.org/t/WxJVsEWWmbc/#m9) [[LSJ
    20100206-2]](https://usenet.krcg.org/t/reXyybyIYX8/#m1) [[LSJ
    20070320-2]](https://usenet.krcg.org/t/QPCnTltI2Rk/#m1) [[ANK
    20200420-2]](https://usenet.krcg.org/t/vekn-58209/#m7) — {Maris Streck},
    {Edith Blount}, {Slake the Thirst}, {Hukros}, {Santaleous}.

[^1-2-4]: [[ANK
    20230316]](https://usenet.krcg.org/t/vekn-80385/#m2)
    [[ANK 20221102-2]](https://usenet.krcg.org/t/vekn-80130/#m1) [[RBK
    wording-templates]](https://www.vekn.net/rulebook#wording-templates) — {Amulet of Temporal Perception}, groups
    "Weapon once per combat" (G00045) and "Weapon once per round" (G00046), {Haqim's Law: Retribution}.

[^1-2-5]: [[LSJ 20040127]](https://usenet.krcg.org/t/2b3SFZyo3ik/#m6) [[ANK
    20221102-2]](https://usenet.krcg.org/t/vekn-80130/#m1) — {Owain Evans, The
    Wanderer}, {Haqim's Law: Retribution}.

[^1-2-6]: [[ANK 20220204]](https://usenet.krcg.org/t/vekn-79634/#m3) [[LSJ
    20100206]](https://usenet.krcg.org/t/cAGrXqpO-YQ/#m1) [[LSJ 20030224]](https://usenet.krcg.org/t/67261v339Ds/#m5)
    [[ANK 20180627-1]](https://usenet.krcg.org/t/vekn-76757/#m1) — {Dust Up},
    {Asguresh}, group "Cancel" (G00058), {Inscription}.

[^1-2-15]: [[ANK
    20190117-1]](https://usenet.krcg.org/t/vekn-77308/#m3) [[RBK
    bleed]](https://www.vekn.net/rulebook#bleed) [[RBK
    additional-strikes]](https://www.vekn.net/rulebook#additional-strikes) — {Mask of a Thousand Faces}.

[^1-2-18]: Owner's reading, no official ruling: [[rules-qa
    29.8]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-29--combat-cards-wave-2) —
    {Lightning Reflexes}.

[^1-2-20]: Owner's reading, no official ruling: [[rules-qa
    7.2]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-7--library-card-contests-hunting-grounds)
    [[rules-qa
    89.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-89--round-20-the-cards-players-asked-for-the-two-player-lists-last-two-readings-graded-claude-2026-10-23-owner-2026-10-23-891-corrected-892-confirmed)
    — {Zoo Hunting Ground}.

[^1-2-7]: [[ANK
    20220705]](https://usenet.krcg.org/t/vekn-79895/#m1)
    [[LSJ 20100705]](https://usenet.krcg.org/t/Dm_Zqyjdx8s/) [[RBK
    wording-templates]](https://www.vekn.net/rulebook#wording-templates) — groups "Unconditional referendum ability"
    (G00039) and "Non-locking referendum ability" (G00040).

[^1-2-9]: [[LSJ 20040518]](https://usenet.krcg.org/t/4emymfUPwAM/#m5) [[ANK
    20210309-2]](https://usenet.krcg.org/t/vekn-79005/#m7) [[LSJ
    20020429]](https://usenet.krcg.org/t/7ZfedGTVt9g/#m1) — {Filchware's Pawn Shop}, group "Vote playable once per game"
    (G00030), group "Since your last turn" (G00010).

[^1-2-19]: Owner's reading, no official ruling: [[rules-qa
    44.10]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-44--legacy-cards-for-tournament-decks-slice-16a)
    [[rules-qa
    71.16]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-71--legacy-cards-for-twda-decks-round-11-the-20252026-decks-entering-the-corpus-readings-no-ruling-settles-claude-2026-10-14-owner-2026-10-14-7117120-71227127-confirmed-7121-corrected-7125s-open-part-ruled)
    [[rules-qa 40.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-40--events) [[RBK
    discard-phase]](https://www.vekn.net/rulebook#discard-phase) — {Ashur Tablets}, {Fortschritt Library}.

[^1-2-12]: [[RTR 19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0) [[LSJ
    20021122-2]](https://usenet.krcg.org/t/1hL0-URtroc/#m5) [[ANK
    20211112]](https://usenet.krcg.org/t/vekn-79475/#m2) — {Anathema}.

[^1-2-13]: [[LSJ 20001127-2]](https://usenet.krcg.org/t/KInac4MQMuA/#m4) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[LSJ 20031011]](https://usenet.krcg.org/t/GZZJjzAeuxU/#m1) —
    {Escaped Mental Patient}, {Flaming Candle}, {First Tradition: The Masquerade}.

[^1-2-14]: [[LSJ 20090205]](https://usenet.krcg.org/t/l5bUtHOejmc/#m2) [[LSJ
    19980105]](https://usenet.krcg.org/t/PzVC-AeFuUQ/#m1) [[LSJ 20081123]](https://usenet.krcg.org/t/_Pb29pBJ1kU/#m2) —
    {Andre LeRoux}, {Archon Investigation}, {Major Boon}.

[^1-2-16]: [[ANK
    20191219]](https://usenet.krcg.org/t/vekn-78241/#m3)
    — {Telepathic Tracking}, {Psyche!}.

[^1-3-1]: [[LSJ 20080114]](https://usenet.krcg.org/t/aMum-K7LEeo/#m1) [[ANK
    20220824]](https://usenet.krcg.org/t/vekn-79986/#m7)
    — group "Locquipments" (G00047), {The British Museum, London}.

[^1-3-2]: [[ANK
    20220824]](https://usenet.krcg.org/t/vekn-79986/#m7)
    [[ANK 20221026]](https://usenet.krcg.org/t/vekn-80117/#m1) [[LSJ
    20080107]](https://usenet.krcg.org/t/XpZ6F53jK-c/#m3) [[LSJ 20080109]](https://usenet.krcg.org/t/XpZ6F53jK-c/#m10)
    [[LSJ 20080114]](https://usenet.krcg.org/t/aMum-K7LEeo/#m1) — {The British Museum, London}, {Therbold Realty},
    {Marie Faucigny}, {Regina Blake}.

[^1-3-3]: [[LSJ 19971001-2]](https://usenet.krcg.org/t/RY_nhdykKP0/#m0) [[LSJ
    20080114]](https://usenet.krcg.org/t/aMum-K7LEeo/#m1) [[ANK
    20221210]](https://usenet.krcg.org/t/vekn-80203/#m2) [[LSJ
    20020211]](https://usenet.krcg.org/t/ubqDaLeG3qo/#m2) [[LSJ 20060221]](https://usenet.krcg.org/t/OC41YQYfJO4/#m1) —
    {Magic of the Smith}, {Danylo}, {Reg Driscoll}.

[^1-3-4]: [[LSJ 20050315]](https://usenet.krcg.org/t/COcJX2hHP-E/#m1) [[LSJ
    20050413-1]](https://usenet.krcg.org/t/1TXRhYopt70/#m25) [[LSJ
    20050413-2]](https://usenet.krcg.org/t/1TXRhYopt70/#m27) [[LSJ 20020211]](https://usenet.krcg.org/t/ubqDaLeG3qo/#m2)
    [[LSJ 20060221]](https://usenet.krcg.org/t/OC41YQYfJO4/#m1) — {Beast, The Leatherface of Detroit}, {Enkidu, The
    Noah}, {Helen Fairchild}, {Lorrie Dunsirn}.

[^1-3-6]: [[PIB 20150924-2]](https://usenet.krcg.org/t/vekn-73293/#m4) [[LSJ
    20031221]](https://usenet.krcg.org/t/OJW-3jpaM04/#m2) — {Adana de Sforza}, {Conrad Adoula}, {Horrock}, {Jane Sims},
    {Nergal}, {Rex, The Necronomist}, {Scout Youngwood}, {Henry Taylor}.

[^1-3-8]: [[LSJ 20011204-3]](https://usenet.krcg.org/t/JhBwcZjDlaQ/#m1) [[PIB
    20130704]](https://usenet.krcg.org/t/vekn-50982/#m2) — group "Multi-discipline"
    (G00086).

[^1-3-7]: [[ANK 20230824]](https://usenet.krcg.org/t/vekn-80782/#m7)
    [[RTR 20070707]](https://usenet.krcg.org/t/vSOt2c1uRzQ/#m0) [[LSJ
    20070927]](https://usenet.krcg.org/t/VaSQ7JL2N2Y/#m1) [[LSJ 20050407]](https://usenet.krcg.org/t/fDl3t2lJ3Pc/#m1)
    [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) [[ANK
    20220818]](https://usenet.krcg.org/t/vekn-79972/#m6)
    — {Pack Alpha}, groups "Lock to reduce cost" (G00130) and "Equip/employ/recruit outside of an action" (G00131),
    {Shackles of Enkidu}, {Enhanced Coagulant}.

[^1-4-1]: [[LSJ 20071001-1]](https://usenet.krcg.org/t/bb9ac-kN5WY/#m1) [[LSJ
    20040623]](https://usenet.krcg.org/t/PlPp5igHOEg/#m1) — {Agent of Power}, {Absimiliard's Army}.

[^1-4-2]: [[LSJ 20040623]](https://usenet.krcg.org/t/PlPp5igHOEg/#m1) [[LSJ
    20071003]](https://usenet.krcg.org/t/duRrP46XygI/#m55) [[LSJ 20040812]](https://usenet.krcg.org/t/n_0DGDsWG0E/#m1) —
    {Absimiliard's Army}, {Père Lachaise, France}.

[^1-4-5]: [[LSJ 20020426]](https://usenet.krcg.org/t/xFGNrTa9CPk/#m1) — {Reality Mirror}.

[^1-5-1]: [[ANK
    20180805]](https://usenet.krcg.org/t/vekn-76897/#m1) [[ANK
    20221120]](https://usenet.krcg.org/t/vekn-80172/#m1)
    — {Safe Passage}, {Secure Haven}.

[^1-5-2]: [[LSJ 20070403]](https://usenet.krcg.org/t/TJ2ktt_1tjk/#m9) [[LSJ
    20070413]](https://usenet.krcg.org/t/umdINigMKqs/#m19) — {Champion}, {Discern}, {Donate}, {Hide}, {Surge},
    {Vigilance}.

[^1-5-4]: [[PIB 20150720]](https://usenet.krcg.org/t/vekn-72088/#m4) [[RTR
    19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[ANK
    20220705]](https://usenet.krcg.org/t/vekn-79895/#m1)
    [[LSJ 20100705]](https://usenet.krcg.org/t/Dm_Zqyjdx8s/) [[LSJ 20001102]](https://usenet.krcg.org/t/LlPyLJjLdx0/#m2)
    [[RBK important-terms-of-the-game]](https://www.vekn.net/rulebook#important-terms-of-the-game) — {Make an Example},
    {Montano}, {Toby}, {Paul Forrest, False Prophet}, {Courier}, groups "Unconditional referendum ability" (G00039) and
    "Non-locking referendum ability" (G00040).

[^1-5-5]: [[LSJ 20011022]](https://usenet.krcg.org/t/KMg4MwD-Jn0/#m1) [[LSJ
    20100902-2]](https://usenet.krcg.org/t/mFpx91METxM/#m1) [[LSJ
    20010813-2]](https://usenet.krcg.org/t/zkKhvgZy9hA/#m2) [[ANK
    20220705]](https://usenet.krcg.org/t/vekn-79895/#m1)
    [[LSJ 20100705]](https://usenet.krcg.org/t/Dm_Zqyjdx8s/) [[RBK torpor]](https://www.vekn.net/rulebook#torpor) —
    group "Ability usable in torpor" (G00027), {Marciana Giovanni, Investigator}, {Pariah}, {Tori Longwood}, {Montano}.

[^1-5-6]: [[ANK
    20180517]](https://usenet.krcg.org/t/vekn-76447/#m33) [[RTR
    20180303]](https://usenet.krcg.org/t/vekn-76447/#m0) [[RBK
    recruit-ally]](https://www.vekn.net/rulebook#recruit-ally) — group "Allies with 'lock this ally to' abilities"
    (G00119), {Abomination}, {Underbridge Stray}, {War Ghoul}, {Paul "Sixofswords29" Moreton}.

[^1-5-14]: [[LSJ 20070301]](https://usenet.krcg.org/t/-CeFWHQ2wXE/#m32) — {The Grandest Trick}.

[^1-5-7]: [[TOM 19960109]](https://usenet.krcg.org/t/gB3xexVAV0s/#m5) [[RTR
    20180511]](https://usenet.krcg.org/t/vekn-76595/#m0) [[LSJ
    20030605]](https://usenet.krcg.org/t/jlVA4lGfpkA/#m4) [[LSJ 20020814]](https://usenet.krcg.org/t/gt8wQhk76lA/#m1)
    [[ANK
    20180307-2]](https://usenet.krcg.org/t/vekn-76451/#m1)
    [[ANK 20200420-2]](https://usenet.krcg.org/t/vekn-58209/#m7) [[LSJ
    20040617]](https://usenet.krcg.org/t/WxJVsEWWmbc/#m9) — {Corpse Minion}, {Count Zaroff}, {General Perfidio Díos},
    {Maris Streck}, {Matteus, Flesh Sculptor}, {Santaleous}, {Edith Blount}.

[^1-5-8]: [[TOM 19960109]](https://usenet.krcg.org/t/gB3xexVAV0s/#m5) [[LSJ
    20070320-2]](https://usenet.krcg.org/t/QPCnTltI2Rk/#m1) [[ANK
    20190725]](https://usenet.krcg.org/t/vekn-77813/#m1) [[ANK
    20191218]](https://usenet.krcg.org/t/vekn-62700/#m7) [[LSJ
    20020814]](https://usenet.krcg.org/t/gt8wQhk76lA/#m1) [[ANK
    20221229]](https://usenet.krcg.org/t/vekn-80231/#m1) —
    {Forest of Shadows}, {Hukros}, {Josef von Bauren}, {Nahir}, {Osric Vladislav}.

[^1-5-9]: [[LSJ 20091021-2]](https://usenet.krcg.org/t/x5oG5J7Egtg/#m7) — {Nergal}.

[^1-5-10]: [[LSJ 19970718]](https://usenet.krcg.org/t/QujjxfQHYzw/#m1) [[SFC
    19960819]](https://usenet.krcg.org/t/G40EE8vCBB8/#m2) — group "Permanents that increase bleed amount during an
    action" (G00117), {Spiridonas}, {Pentex(TM) Loves You!}.

[^1-5-11]: [[ANK 20221021]](https://usenet.krcg.org/t/vekn-80108/#m1) — {Watenda}, {White Lily}.

[^1-5-12]: [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) — {Drum of Xipe Totec}, {Hand of Conrad}.

[^1-6-1]: [[LSJ 20030214]](https://usenet.krcg.org/t/A3U-Dy1yx8Y/#m1) [[ANK
    20171212]](https://usenet.krcg.org/t/vekn-76334/#m14) [[LSJ
    20030202]](https://usenet.krcg.org/t/ox7A8EvaNJo/#m3) [[ANK
    20210124]](https://usenet.krcg.org/t/vekn-79001/#m2) [[RBK
    requirements-for-playing-cards]](https://www.vekn.net/rulebook#requirements-for-playing-cards) — {Undead
    Persistence}, {Mental Maze}, {Maxwell}, {The Red Question}.

[^1-6-29]: [[ANK
    20210124]](https://usenet.krcg.org/t/vekn-79001/#m2) [[ANK
    20220218]](https://usenet.krcg.org/t/vekn-79001/#m13)
    — {The Red Question}, group "Action only usable by vampire in torpor" (G00079), {Warsaw Station}.

[^1-6-2]: [[PIB 20121031]](https://usenet.krcg.org/t/vekn-39816/#m35)
    — group "Require Gehenna" (G00009).

[^1-6-3]: [[LSJ 20051110]](https://usenet.krcg.org/t/fMW3tSF5oUc/#m9) — {Orun}.

[^1-6-4]: [[LSJ 20030607]](https://usenet.krcg.org/t/DG7s60pwv1U/#m2) [[ANK
    20180801]](https://usenet.krcg.org/t/vekn-76839/#m3) [[ANK
    20221208]](https://usenet.krcg.org/t/vekn-80197/#m5)
    [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) — {Seal of Veddartha}, {The Crusader Sword}, {Agate
    Talisman}, {Blade of Enoch}, {Stolen Police Cruiser}, {Exile}, {Soul Gem of Etrius}.

[^1-6-5]: [[LSJ 19980303]](https://usenet.krcg.org/t/Q_kKQiACZqk/#m2) [[LSJ
    20050128-2]](https://usenet.krcg.org/t/HVy8iPUxNbI/#m36) [[ANK
    20221011-4]](https://usenet.krcg.org/t/vekn-79988/#m2) [[RTR
    20070707]](https://usenet.krcg.org/t/vSOt2c1uRzQ/#m0) [[PIB
    20150304]](https://usenet.krcg.org/t/vekn-69627/#m9) — {Talaq, The
    Immortal}, {Kemintiri}, {Vlad Tepes}.

[^1-6-26]: [[LSJ 20081120-1]](https://usenet.krcg.org/t/e2PNDpg-l_c/#m19) — {Vidal Jarbeaux}; the card text of {Masca},
    {Hexe}, {Evil Eye}.

[^1-6-6]: [[ANK
    20190203]](https://usenet.krcg.org/t/vekn-77343/#m5)
    — {Defender of the Haven}, {Derange}, {Ennoia's Theater}.

[^1-6-7]: [[LSJ 20100204]](https://usenet.krcg.org/t/o5Xnzc8G774/#m31) [[LSJ
    20051116-2]](https://usenet.krcg.org/t/5Bovsb8I6R0/#m1) [[LSJ
    20071001-1]](https://usenet.krcg.org/t/bb9ac-kN5WY/#m1) [[LSJ 20100422]](https://usenet.krcg.org/t/tMEimr0yxLA/#m5)
    [[LSJ 20040531]](https://usenet.krcg.org/t/oTb4vsFNi1s/#m2) [[LSJ
    20040518-2]](https://usenet.krcg.org/t/4emymfUPwAM/#m1) [[LSJ 20100303]](https://usenet.krcg.org/t/jmmm0WRUPvs/#m4)
    [[LSJ 20100426]](https://usenet.krcg.org/t/BN3xmoZ0W1A/#m2) — groups "Put card in play ignoring requirements"
    (G00110) and "Equip/employ/recruit outside of an action" (G00131), {Abombwe}, {Agent of Power}, {Absimiliard's
    Army}, {Echo of Harmonies}.

[^1-6-8]: [[LSJ 20071015]](https://usenet.krcg.org/t/Bom6ae7qjbI/#m9) [[LSJ
    20100302-1]](https://usenet.krcg.org/t/jmmm0WRUPvs/#m1) [[PIB
    20111101]](https://usenet.krcg.org/t/vekn-12975/#m2) — {Compel the
    Spirit}, {Pressing Flesh}, {Summon History}.

[^1-6-9]: [[ANK 20200810]](https://usenet.krcg.org/t/vekn-78797/#m2)
    [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[TOM
    19950407]](https://usenet.krcg.org/t/FWVnIu3zLAQ/#m5) [[LSJ 20080702-4]](https://usenet.krcg.org/t/sCHpPQkjeAE/#m3)
    [[ANK
    20210714]](https://usenet.krcg.org/t/vekn-79224/#m3)
    — {Bindusara, Historian of the Kindred}, {Horrid Reality}, {Vast Wealth}, {Topaz}, {Villein}.

[^1-6-27]: [[TOM 19950407]](https://usenet.krcg.org/t/FWVnIu3zLAQ/#m5) — {Vast Wealth}, {Darius Styx}.

[^1-6-10]: [[LSJ 19990602]](https://usenet.krcg.org/t/tnBL-IWbBV4/#m4) [[ANK
    20200320]](https://usenet.krcg.org/t/vekn-78531/#m1)
    [[LSJ 19980206]](https://usenet.krcg.org/t/p_uyqQgE9Ms/#m1) [[LSJ
    20020426]](https://usenet.krcg.org/t/xFGNrTa9CPk/#m1) — {Heidelberg Castle, Germany}, {Shackles of Enkidu}, {Deviki
    Prasanta}, {Lucian}, {Reality Mirror}.

[^1-6-12]: [[LSJ 20090415]](https://usenet.krcg.org/t/oDvJWs7majs/#m4) [[LSJ
    20100119]](https://usenet.krcg.org/t/1eULCGaVcO0/#m1) [[LSJ 20090506]](https://usenet.krcg.org/t/887DQTpntKI/#m5)
    [[LSJ 20060209]](https://usenet.krcg.org/t/ZuOfZorIhhU/#m4) [[LSJ
    20100421]](https://usenet.krcg.org/t/Vp--M79gpqk/#m1) [[LSJ 20080702-2]](https://usenet.krcg.org/t/sCHpPQkjeAE/#m1)
    — {Helicopter}, {Incriminating Videotape}, {Shilmulo Tarot}, {Flaming Candle}, {Dagger}, {Baseball Bat}.

[^1-6-13]: [[PIB
    20150105-2]](https://usenet.krcg.org/t/vekn-68482/#m1)
    [[ANK 20210913]](https://usenet.krcg.org/t/vekn-79322/#m1) [[LSJ
    20090116]](https://usenet.krcg.org/t/RQ3ARP9Kvfk/#m7) [[ANK
    20210928]](https://usenet.krcg.org/t/vekn-79364/#m2) [[LSJ
    20090115-1]](https://usenet.krcg.org/t/RQ3ARP9Kvfk/#m1) [[ANK
    20170309]](https://usenet.krcg.org/t/vekn-75649/#m2) — {Synner-G},
    {Topaz}, {Little Tailor of Prague}, {Sébastien Goulet}, {Kuyén}, {Soul of the Earth}, {Zhenga}.

[^1-6-14]: [[LSJ 20100725]](https://usenet.krcg.org/t/9d1zMZfsfNo/) [[LSJ
    20100303]](https://usenet.krcg.org/t/jmmm0WRUPvs/#m4) [[ANK
    20180817]](https://usenet.krcg.org/t/vekn-76933/#m2) [[ANK
    20220704]](https://usenet.krcg.org/t/vekn-79890/#m1)
    [[LSJ 20100819]](https://usenet.krcg.org/t/x_u1Qtiguzg/#m1) — groups "Equip/employ/recruit outside of an action"
    (G00131) and "Equip/Employ/Recruit action" (G00132), {Charming Lobby}.

[^1-6-15]: [[ANK
    20190228]](https://usenet.krcg.org/t/vekn-77427/#m47)
    [[RTR 20070707]](https://usenet.krcg.org/t/vSOt2c1uRzQ/#m0) [[LSJ
    20070708]](https://usenet.krcg.org/t/vSOt2c1uRzQ/#m11) [[LSJ 20101013]](https://usenet.krcg.org/t/hpFRDAmtSbA/#m4)
    [[ANK 20201028]](https://usenet.krcg.org/t/vekn-78885/#m1) [[LSJ
    20091015]](https://usenet.krcg.org/t/pqa7mYZ6NEM/#m9) [[PIB
    20130128]](https://usenet.krcg.org/t/vekn-43572/#m36)
    [[LSJ 20071109]](https://usenet.krcg.org/t/mXspOwNnPDc/) — group "Impersonating requirements to play cards"
    (G00073), {Mata Hari}, {Vidal Jarbeaux}, {Vlad Tepes}.

[^1-6-28]: [[PIB 20150306]](https://usenet.krcg.org/t/vekn-69627/#m12) — {Vlad
    Tepes}.

[^1-6-16]: [[LSJ 20050119]](https://usenet.krcg.org/t/NKWhBnp7uP4/#m11) — group "Impersonating requirements for cards
    not played normally" (G00072), {Vidal Jarbeaux}.

[^1-6-17]: [[LSJ 20100526]](https://usenet.krcg.org/t/ch0oKrlxX30/#m2) [[ANK
    20220718]](https://usenet.krcg.org/t/vekn-79913/#m3) [[LSJ
    20050721]](https://usenet.krcg.org/t/g39H3dwXqvc/#m20) [[PIB
    20150306]](https://usenet.krcg.org/t/vekn-69627/#m12) — {Sacrifice},
    {Infamous Insurgent}, {Kemintiri}, {Vlad Tepes}.

[^1-6-18]: [[PIB 20150530]](https://usenet.krcg.org/t/vekn-66061/#m7) [[LSJ
    20070708]](https://usenet.krcg.org/t/vSOt2c1uRzQ/#m11) [[LSJ 20071109]](https://usenet.krcg.org/t/mXspOwNnPDc/) —
    {Lisandro Giovanni}, {Vidal Jarbeaux}, {Mata Hari}.

[^1-6-25]: [[LSJ 20050225]](https://usenet.krcg.org/t/OjszbddbvxM/#m2) — {Mata Hari}, {Lisandro Giovanni}, {Petaniqua},
    {Vidal Jarbeaux}, {Winterlich}.

[^1-6-20]: [[LSJ 20090508]](https://usenet.krcg.org/t/B7dz3qoIITQ/#m2) [[LSJ
    20011217]](https://usenet.krcg.org/t/vZ13bh7FEvQ/#m20) [[ANK
    20180731-1]](https://usenet.krcg.org/t/vekn-76877/#m2) [[LSJ
    20031116]](https://usenet.krcg.org/t/RQBw8EnnD5s/#m1) [[ANK
    20230816-2]](https://usenet.krcg.org/t/vekn-80683/#m16) —
    {Inceptor}, {Infernal Familiar}, {Ian Forestal}, {Grey Thorne}, {Inscription}.

[^1-6-21]: [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[LSJ
    20021211]](https://usenet.krcg.org/t/-J07wvmidOA/#m17) [[PIB
    20110725]](https://usenet.krcg.org/t/vekn-6728/#m4) [[LSJ
    20060620]](https://usenet.krcg.org/t/aqascblOZmk/#m3) — {Absorb the Mind}, {Call the Lamprey}, {Draba}, {Night
    Terrors}, {Siphon}, {Kindred Segregation}, {Peace Treaty}, group "Can draw from empty" (G00160).

[^1-6-22]: [[RTR 19980928]](https://usenet.krcg.org/t/Xva4_IRavxM/#m0) [[RTR
    19951110]](https://usenet.krcg.org/t/TXfganI5B2o/#m0) [[ANK
    20200420-3]](https://usenet.krcg.org/t/vekn-78574/#m2) [[LSJ
    20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) [[LSJ 20070901]](https://usenet.krcg.org/t/EouVnC0MuH4/) [[LSJ
    20060623]](https://usenet.krcg.org/t/mfgW0TeoLNM/#m1) [[LSJ 20010810]](https://usenet.krcg.org/t/FUXkrq3B_O8/#m2)
    [[PIB 20121028]](https://usenet.krcg.org/t/vekn-39816/#m10) [[ANK
    20200508-2]](https://usenet.krcg.org/t/vekn-78623/#m2) [[ANK
    20230817]](https://usenet.krcg.org/t/vekn-80780/#m1) — {Shattering Blow}, {Fractured
    Armament}, {Canine Horde}, {Blessing of Durga Syn}, {Brujah Frenzy}, group "Targets card in Ash Heap" (G00122),
    {Storage Annex}, {Liquidation}, {Undue Influence}, {Free States Rant}, {Taste of Vitae}, {T.J.}, {Hide the Heart}.

[^1-6-23]: [[LSJ 20001114]](https://usenet.krcg.org/t/qXSlM7Grq1c/#m1) [[LSJ
    20030121]](https://usenet.krcg.org/t/lED3kZ2UUUo/#m3) [[TOM 19951109]](https://usenet.krcg.org/t/WhJj5K1Fa-0/#m10)
    [[LSJ 20010315]](https://usenet.krcg.org/t/m9CrEOn1veo/#m3) [[LSJ
    20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) — group "Immideate damage prevention" (G00154), {Hidden
    Strength}, {Apparition}, {Brother's Blood}, {Bonding}.

[^1-6-30]: [[TOM 19951109]](https://usenet.krcg.org/t/WhJj5K1Fa-0/#m10) — {Bonding}.

[^1-6-24]: [[LSJ 20011214-2]](https://usenet.krcg.org/t/TinQ8ywzIHU/#m1) — {Repulsion}.

[^1-7-1]: [[LSJ 20091021-2]](https://usenet.krcg.org/t/x5oG5J7Egtg/#m7) [[LSJ
    20080201]](https://usenet.krcg.org/t/y5Uoc7nEulU/#m1) [[ANK
    20200703-2]](https://usenet.krcg.org/t/vekn-78713/#m4)
    — {Nergal}, {Yawp Court}, {Preternatural Evasion}.

[^1-7-3]: [[ANK
    20190625]](https://usenet.krcg.org/t/vekn-77741/#m1)
    [[PIB 20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) [[LSJ
    20090208]](https://usenet.krcg.org/t/07gUFmSeIxU/#m1) [[LSJ 20040730]](https://usenet.krcg.org/t/vCZw1_QnhfE/#m2)
    [[ANK 20230527]](https://usenet.krcg.org/t/vekn-80553/#m1) —
    {The Slow Withering}, {Tryphosa}, {Enrage}, {Aura of Invincibility}, {Blithe Acceptance}.

[^1-7-4]: [[LSJ 19990421]](https://usenet.krcg.org/t/bYwaXWdJX84/#m30) [[LSJ
    20070417]](https://usenet.krcg.org/t/ecDUqbSUsNg/#m1) [[LSJ 20031121-2]](https://usenet.krcg.org/t/1khXmKPU0ws/#m31)
    — {Forced Awakening}, group "Burn blood to attempt to block" (G00088).

[^1-7-5]: [[LSJ 20020620]](https://usenet.krcg.org/t/WoXWzLYaFSY/#m1) [[LSJ
    20080512]](https://usenet.krcg.org/t/z2DGSFph6sM/#m17) [[LSJ 20050607]](https://usenet.krcg.org/t/WLv9R8wA0Ow/#m8)
    [[LSJ 20050608]](https://usenet.krcg.org/t/WLv9R8wA0Ow/#m10) [[RTR
    20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) —
    {Minion Tap}, {Villein}, group "Allies who can play as a vampire" (G00011), {Spying Mission}, {Thanks for the
    Donation}, {Repo Man}.

[^1-7-6]: [[RTR 20070707]](https://usenet.krcg.org/t/vSOt2c1uRzQ/#m0) [[PIB
    20150917]](https://usenet.krcg.org/t/vekn-73156/#m2) — {The
    Ankara Citadel, Turkey}, {Thin-Blooded Seer}.

[^1-7-7]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[PIB
    20150725]](https://usenet.krcg.org/t/vekn-72184/#m1) [[LSJ
    20100429]](https://usenet.krcg.org/t/OUcf-1EMXXA/#m3) [[ANK
    20170716]](https://usenet.krcg.org/t/vekn-75987/#m4) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) — {Black Cat}, {Ivan Krenyenko}, {Cavalier}, {Khobar Towers,
    Al-Khubar}.

[^1-7-8]: [[LSJ 20040701]](https://usenet.krcg.org/t/O7cT51Yl2Uk/#m1) [[LSJ
    20080125]](https://usenet.krcg.org/t/HcCobkGPcFI/#m1) [[ANK
    20181215]](https://usenet.krcg.org/t/vekn-77213/#m39)
    [[RTR 19960112]](https://usenet.krcg.org/t/d3n3StNS7no/#m0) [[LSJ
    20050408]](https://usenet.krcg.org/t/idZy8QCxBT0/#m1) [[RTR
    20111202]](https://usenet.krcg.org/t/vekn-16769/#m0) [[ANK
    20180109]](https://usenet.krcg.org/t/vekn-76360/#m4) — {Concealed Weapon}, {Ravnos
    Cache}, {Ravnos Carnival}.

[^1-7-9]: [[ANK
    20200709]](https://usenet.krcg.org/t/vekn-74710/#m4) —
    {Jenna Cross}.

[^1-7-10]: [[ANK
    20210714-2]](https://usenet.krcg.org/t/vekn-79217/#m6)
    [[ANK
    20210714]](https://usenet.krcg.org/t/vekn-79224/#m3)
    [[ANK
    20230721]](https://usenet.krcg.org/t/vekn-80691/#m15)
    [[LSJ 20080612]](https://usenet.krcg.org/t/THFR8YTXxGI/#m1) — {The Eternals of Sirius}, {Villein}, {Powerbase:
    Tshwane}, {The Shard, London}.

[^1-7-11]: [[LSJ 20030502]](https://usenet.krcg.org/t/qBa9CvBwqNA/#m1) [[LSJ
    19970224]](https://usenet.krcg.org/t/80KRDjVFkyg/#m1) [[LSJ 20071015]](https://usenet.krcg.org/t/Bom6ae7qjbI/#m9)
    [[LSJ 20100302-1]](https://usenet.krcg.org/t/jmmm0WRUPvs/#m1) [[PIB
    20111101]](https://usenet.krcg.org/t/vekn-12975/#m2) [[ANK
    20220704]](https://usenet.krcg.org/t/vekn-79890/#m1)
    [[RBK announce-the-action]](https://www.vekn.net/rulebook#announce-the-action) — group "Cost X" (G00032), {Hidden
    Strength}, {Compel the Spirit}, {Pressing Flesh}, {Summon History}, {Charming Lobby}.

[^1-7-12]: [[ANK
    20230620]](https://usenet.krcg.org/t/vekn-80612/#m1)
    [[ANK 20170605]](https://usenet.krcg.org/t/vekn-75862/#m3) [[LSJ
    20090705]](https://usenet.krcg.org/t/3Ekxk6uQ_wo/#m3) [[LSJ 20030917]](https://usenet.krcg.org/t/i3Ugq5AQWFI/#m1) —
    group "Lock to reduce cost" (G00130), {Powerbase: Tshwane}, {The Line}, {The Shard, London}, {Eccentric
    Billionaire}, {Sunset Strip, Hollywood}.

[^1-7-30]: Owner's reading, no official ruling: [[rules-qa
    101.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-101--reachable-edges-mid-play-ousts-and-costs-the-cards-board-owner-2026-10-26-1011-1013-claude-2026-10-26)
    [[rules-qa
    36.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-36--cost-reductions) — {The
    Shard, London}.

[^1-7-13]: [[RTR 20070707]](https://usenet.krcg.org/t/vSOt2c1uRzQ/#m0) [[ANK
    20230824]](https://usenet.krcg.org/t/vekn-80782/#m7) [[ANK
    20210928]](https://usenet.krcg.org/t/vekn-79364/#m2) [[LSJ
    20090115-1]](https://usenet.krcg.org/t/RQ3ARP9Kvfk/#m1) [[ANK
    20170309]](https://usenet.krcg.org/t/vekn-75649/#m2) [[LSJ
    20090920-2]](https://usenet.krcg.org/t/nbgHtblc8jc/) — {Powerbase: Tshwane}, {Pack Alpha}, {Soul of the Earth},
    {Piper}, {Antonio d'Erlette}, {Sunset Strip, Hollywood}.

[^1-7-14]: [[ANK 20221102]](https://usenet.krcg.org/t/vekn-80129/#m1)
    [[LSJ 20081124-1]](https://usenet.krcg.org/t/Fp4wSGFJ7N4/#m1) [[TOM
    19951208-1]](https://usenet.krcg.org/t/tEHebi9BYfc/#m5) [[ANK
    20181007]](https://usenet.krcg.org/t/vekn-77057/#m1) [[RTR
    19950509]](https://usenet.krcg.org/t/_LKyR7pdMig/#m8) [[ANK
    20230721]](https://usenet.krcg.org/t/vekn-80691/#m15)
    — {Powerbase: Tshwane}, {The Shard, London}, {Walks-With-Might}, {Secure Haven}, {The Line}, {Democritus}.

[^1-7-15]: [[LSJ 20080114]](https://usenet.krcg.org/t/aMum-K7LEeo/#m1) [[LSJ
    20080107]](https://usenet.krcg.org/t/XpZ6F53jK-c/#m3) [[LSJ 20080109]](https://usenet.krcg.org/t/XpZ6F53jK-c/#m10)
    [[ANK 20221026]](https://usenet.krcg.org/t/vekn-80117/#m1) — group
    "Locquipments" (G00047), {Therbold Realty}, {Marie Faucigny}, {Regina Blake}, {The British Museum, London}.

[^1-7-16]: [[ANK
    20210226]](https://usenet.krcg.org/t/vekn-79045/#m21) [[ANK
    20180327]](https://usenet.krcg.org/t/vekn-76480/#m1) — groups "Burn blood for
    effect" (G00065) and "Burn blood for effect (mandatory)" (G00097), {Preternatural Evasion}, {Shadow Boxing}, {Aura
    Absorption}, {Enrage}.

[^1-7-17]: [[LSJ 20090107]](https://usenet.krcg.org/t/qPDLvbgqfCg/#m11) [[LSJ
    20070829-3]](https://usenet.krcg.org/t/sVXZw1J43ik/#m12) [[ANK
    20220809]](https://usenet.krcg.org/t/vekn-79949/#m3) — {Dragos}, {Loki's
    Gift}.

[^1-7-18]: [[LSJ 20070829-1]](https://usenet.krcg.org/t/sVXZw1J43ik/#m1) [[ANK
    20180910-3]](https://usenet.krcg.org/t/vekn-76990/#m2)
    — {Terror Frenzy}.

[^1-7-19]: [[LSJ 20100216]](https://usenet.krcg.org/t/nrXTh1XKJJ8/#m2) [[ANK
    20180512-2]](https://usenet.krcg.org/t/vekn-76595/#m25)
    [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) — {Villein}, {Minion Tap}, {Theft of Vitae}, {Tongue of
    the Serpent}, {Succubus}.

[^1-7-20]: [[LSJ 20090901]](https://usenet.krcg.org/t/9BZPgQH9PFk/#m2) [[LSJ
    20030522-2]](https://usenet.krcg.org/t/iBfBo7CQn4Q/#m1) [[LSJ 20031219]](https://usenet.krcg.org/t/leLLrcPsiBY/#m3)
    [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) — {Sword of Nuln}.

[^1-7-31]: [[ANK
    20210124]](https://usenet.krcg.org/t/vekn-79001/#m2) [[ANK
    20220218]](https://usenet.krcg.org/t/vekn-79001/#m13)
    [[LSJ 20070411]](https://usenet.krcg.org/t/umdINigMKqs/#m1) — group "Requires successful action" (G00077), {Zephyr},
    {Daring the Dawn}, {Forced March}.

[^1-7-22]: [[LSJ 19971212]](https://usenet.krcg.org/t/jQUqciN0Z9E/#m2) [[ANK
    20181006]](https://usenet.krcg.org/t/vekn-77053/#m2) — {Smiling Jack, The Anarch}.

[^1-7-23]: [[LSJ 20100725]](https://usenet.krcg.org/t/9d1zMZfsfNo/) [[LSJ
    20091021]](https://usenet.krcg.org/t/B_MVmqTcUXE/#m2) [[TOM 19951212-1]](https://usenet.krcg.org/t/NJa6LfMSPks/#m6)
    [[LSJ 20080702-4]](https://usenet.krcg.org/t/sCHpPQkjeAE/#m3) [[ANK
    20180509]](https://usenet.krcg.org/t/vekn-76585/#m3) —
    {Muricia's Call}, {Charming Lobby}, {Horrid Reality}, {Topaz}, {Emerald Legionnaire}.

[^1-7-24]: [[LSJ 20040518]](https://usenet.krcg.org/t/4emymfUPwAM/#m5) [[LSJ
    20100315]](https://usenet.krcg.org/t/06C5ufFEaJs/#m2) — {Alastor}, {Cavalier}.

[^1-7-25]: [[LSJ 20060831]](https://usenet.krcg.org/t/wdVbwFZo8Jg/#m12) [[LSJ
    20071015-2]](https://usenet.krcg.org/t/Ei307R78l4A/#m1) [[ANK
    20220503]](https://usenet.krcg.org/t/vekn-39040/#m5) —
    {Lord Aaron Wesley Wilkshire}, {Watenda}, {Ex Nihilo}.

[^1-7-27]: [[ANK
    20180805]](https://usenet.krcg.org/t/vekn-76897/#m1) [[ANK
    20221120]](https://usenet.krcg.org/t/vekn-80172/#m1)
    — {Safe Passage}, {Secure Haven}.

[^1-7-28]: [[PIB
    20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) [[LSJ
    20090601-2]](https://usenet.krcg.org/t/RkEUabeJNdM/#m1) [[LSJ 20040518]](https://usenet.krcg.org/t/4emymfUPwAM/#m5)
    [[LSJ 20040518-2]](https://usenet.krcg.org/t/4emymfUPwAM/#m1) [[LSJ
    20081213-1]](https://usenet.krcg.org/t/MNmJu12AU8I/#m1) [[ANK
    20210124]](https://usenet.krcg.org/t/vekn-79001/#m2) —
    {Target Head}, {Santaleous}, {The Diamond Thunderbolt}, {Enkil Cog}.

[^1-7-29]: [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[ANK
    20190724]](https://usenet.krcg.org/t/vekn-77810/#m2) [[LSJ
    20060409]](https://usenet.krcg.org/t/gsFQXsCGTG4/#m1) [[ANK
    20220528]](https://usenet.krcg.org/t/vekn-76455/#m19) — {Kindred
    Segregation}, {Peace Treaty}, {Keystone Kine}.

[^1-8-1]: [[LSJ 20061207]](https://usenet.krcg.org/t/nqcrJlhg4Ng/#m2) [[LSJ
    20061013]](https://usenet.krcg.org/t/6w8K3yDtBH0/#m16) [[RTR 20040501]](https://usenet.krcg.org/t/7-mp3Ada86I/#m0)
    [[RBK playing-a-card]](https://www.vekn.net/rulebook#playing-a-card) — group "Cancel" (G00058), {The Admonitions},
    {The Barrens}, {Dreams of the Sphinx}, {The Erciyes Fragments}, {Fragment of the Book of Nod}.

[^1-8-2]: [[LSJ 20021011]](https://usenet.krcg.org/t/9WWIzxek9Nc/#m2) [[ANK
    20231028]](https://usenet.krcg.org/t/vekn-80925/#m3) — group "Cancel as a
    reaction" (G00062), {Unleashing the Bestial Soul}, {Deed the Heart's Desire}, {Psalm of the Damned}.

[^1-8-3]: [[LSJ 20100728]](https://usenet.krcg.org/t/mzzPS-cOprI/#m1) [[LSJ
    20090213]](https://usenet.krcg.org/t/hkTPcLPgZk4/#m1) — {Andrew Stuart}.

[^1-8-21]: Owner's reading, no official ruling: [[rules-qa
    64.8]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-64--legacy-cards-for-twda-decks-rounds-45-the-gangrel-decks-implemented-readings-2026-10-11-owner-2026-10-14-647-648-confirmed-2026-10-20-the-rest-646-restated)
    — {Predator's Transformation}; the rulebook sets no unlocked condition [[RBK
    reflex-cards]](https://www.vekn.net/rulebook#reflex-cards).

[^1-8-4]: [[LSJ 20031201]](https://usenet.krcg.org/t/Mi_j7sUsZZw/#m1) [[LSJ
    20070214]](https://usenet.krcg.org/t/o_BltcAJ_So/#m2) [[ANK
    20180910-2]](https://usenet.krcg.org/t/vekn-76990/#m6)
    [[ANK
    20180910-3]](https://usenet.krcg.org/t/vekn-76990/#m2)
    [[ANK
    20180909]](https://usenet.krcg.org/t/vekn-76987/#m2)
    [[PIB
    20120808]](https://usenet.krcg.org/t/vekn-34265/#m17)
    [[LSJ 20080618]](https://usenet.krcg.org/t/9FAkwIsuYZc/#m3) — {Vox Senis}, {Botched Move}, {Death Seeker}, {Ophidian
    Gaze}, group "Cancel with no action" (G00063).

[^1-8-5]: [[PIB 20130704]](https://usenet.krcg.org/t/vekn-50982/#m2) — {Hide the
    Mind}, {Iron Heart}.

[^1-8-6]: [[RTR 20001020]](https://usenet.krcg.org/t/GvxNYsYsWJ4/#m0) [[ANK
    20201229]](https://usenet.krcg.org/t/vekn-78959/#m1) [[LSJ 20060530-2]](https://usenet.krcg.org/t/7ezcZoziFWw/#m3)
    [[RTR 20040501]](https://usenet.krcg.org/t/7-mp3Ada86I/#m0) — group "Cancel an action" (G00061), {Sudden Reversal},
    {Asguresh}, group "Equip/employ/recruit outside of an action" (G00131), {The Summoning}, {Zhenga}, {The Erciyes
    Fragments}.

[^1-8-7]: [[ANK 20180627-1]](https://usenet.krcg.org/t/vekn-76757/#m1) [[LSJ
    20010716]](https://usenet.krcg.org/t/x2EdFtlPs8Q/#m1) [[ANK
    20221011]](https://usenet.krcg.org/t/vekn-80070/#m8)
    — {Inscription}, {Shadow Court Satyr}.

[^1-8-8]: [[ANK
    20220704]](https://usenet.krcg.org/t/vekn-79890/#m1)
    [[LSJ 20090113-3]](https://usenet.krcg.org/t/8FHCL4AQblI/#m4) [[LSJ
    20100426]](https://usenet.krcg.org/t/BN3xmoZ0W1A/#m2) [[LSJ 20091128]](https://usenet.krcg.org/t/-IxzB0bvhKU/#m1)
    [[LSJ 20011205]](https://usenet.krcg.org/t/dkK2L81_cYk/#m24) — {Charming Lobby}, {Echo of Harmonies}.

[^1-8-9]: [[ANK 20190104]](https://usenet.krcg.org/t/vekn-77254/#m11)
    [[LSJ 19980212]](https://usenet.krcg.org/t/fLFLlXZXHqA/#m0) [[ANK
    20200525-2]](https://usenet.krcg.org/t/vekn-78653/#m1) [[RBK
    cancel-a-card]](https://www.vekn.net/rulebook#cancel-a-card) — group "Cancel" (G00058), {Immortal Grapple}.

[^1-8-10]: [[ANK 20220411]](https://usenet.krcg.org/t/vekn-79730/#m1)
    [[LSJ 20000424]](https://usenet.krcg.org/t/6pK4rFvqmmg/#m1) [[LSJ
    20070330]](https://usenet.krcg.org/t/gsI3SYz2g3o/#m1) [[ANK
    20190701]](https://usenet.krcg.org/t/vekn-77763/#m2) — {Dabbler}, {Infernal
    Familiar}, {Perfectionist}, {Marthe Dizier}.

[^1-8-11]: [[LSJ 20090601-2]](https://usenet.krcg.org/t/RkEUabeJNdM/#m1) [[PIB
    20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) [[LSJ
    20071015-2]](https://usenet.krcg.org/t/Ei307R78l4A/#m1) — {Santaleous}, {Watenda}, the Target aim cycle ({Target
    Hand}, {Target Head}, {Target Leg}, {Target Retainer}, {Target Vitals}).

[^1-8-12]: [[LSJ 20040518]](https://usenet.krcg.org/t/4emymfUPwAM/#m5) [[LSJ
    20040518-2]](https://usenet.krcg.org/t/4emymfUPwAM/#m1) [[LSJ
    20070222-2]](https://usenet.krcg.org/t/jBrfK77gayo/#m5) — {The Diamond Thunderbolt}.

[^1-8-14]: [[LSJ 20030224]](https://usenet.krcg.org/t/67261v339Ds/#m5) [[LSJ
    20100206]](https://usenet.krcg.org/t/cAGrXqpO-YQ/#m1) — group "Cancel" (G00058), {Asguresh}.

[^1-8-20]: [[LSJ 20091021-2]](https://usenet.krcg.org/t/x5oG5J7Egtg/#m7) [[ANK
    20200710]](https://usenet.krcg.org/t/vekn-77985/#m3) — {Nergal}, {Vidal
    Jarbeaux}.

[^1-8-15]: [[RTR 20040501]](https://usenet.krcg.org/t/7-mp3Ada86I/#m0) [[LSJ
    20100206]](https://usenet.krcg.org/t/cAGrXqpO-YQ/#m1) [[RBK
    cancel-a-card]](https://www.vekn.net/rulebook#cancel-a-card) — {Primal Instincts}, {Death Seeker}, {Asguresh}.

[^1-8-16]: [[LSJ 20050228-3]](https://usenet.krcg.org/t/UHEZEmX22jA/#m4) [[ANK
    20230111]](https://usenet.krcg.org/t/vekn-80258/#m2)
    [[ANK
    20230114]](https://usenet.krcg.org/t/vekn-80258/#m7)
    — {Rigor Mortis}.

[^1-8-17]: [[ANK 20200311]](https://usenet.krcg.org/t/vekn-78506/#m1) — {Target
    Retainer}.

[^1-9-1]: [[ANK
    20181208]](https://usenet.krcg.org/t/vekn-77210/#m7)
    — group "Modifier as announced" (G00052), {Approximation of Loyalty}, {Predator's Transformation}, {Force of
    Personality}.

[^1-9-4]: [[LSJ 20001120]](https://usenet.krcg.org/t/Br8FPS5mRn4/#m1) — group "Action not replaced" (G00020), {Steely
    Tenacity}.

[^1-9-5]: [[ANK
    20200629-2]](https://usenet.krcg.org/t/vekn-78701/#m2)
    [[LSJ 20091208]](https://usenet.krcg.org/t/ptHbJM9MlVI/#m1) [[RTR
    19951017]](https://usenet.krcg.org/t/ouhNUbHYg50/#m2) [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1)
    [[LSJ 20001122]](https://usenet.krcg.org/t/Br8FPS5mRn4/#m8) [[LSJ
    20080805]](https://usenet.krcg.org/t/SIbzFAwWDKs/#m1) — group "Replace next turn" (G00025), group "Permanent not
    replaced" (G00008), {Port Authority}, {The New Inquisition}, {Psyche!}.

[^1-9-12]: Owner's reading, no official ruling: [[rules-qa
    88.8]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-88--open-contradictions-that-change-play-rules-digest-23-k5k7-k15-k16-k18-k23-k26-k27-k29-readings-graded-claude-2026-10-22-owner-2026-10-22-881-888-889-confirmed-884-answered-58-kept)
    — {The Meddling of Semsith}.

[^1-9-6]: [[LSJ 20070320-1]](https://usenet.krcg.org/t/goUs4JXcqMw/#m1) [[ANK
    20200905]](https://usenet.krcg.org/t/vekn-78836/#m1) — {The
    Meddling of Semsith}, {Hagar Stone}, {Visit from the Capuchin}.

[^1-9-7]: [[PIB 20110718]](https://usenet.krcg.org/t/vekn-6435/#m1) [[LSJ
    20020718]](https://usenet.krcg.org/t/H0ZTfiyxrbg/#m2) [[ANK
    20180512]](https://usenet.krcg.org/t/vekn-76599/#m1) — {Learjet},
    {Agaitas, The Scholar of Antiquities}, {Troglodytia}.

[^1-9-8]: [[LSJ 20080702-2]](https://usenet.krcg.org/t/sCHpPQkjeAE/#m1) [[ANK
    20200517-3]](https://usenet.krcg.org/t/vekn-78606/#m4) [[LSJ
    20070707]](https://usenet.krcg.org/t/ZtRk5z2TcoI/#m1) — {Baseball Bat}, {Corrupt Construction}.

[^1-9-13]: [[LSJ 20081126]](https://usenet.krcg.org/t/EsG0EVCLqX4/#m2) — group "Play stored card as if from hand"
    (G00095).

[^1-9-9]: [[LSJ 20011023]](https://usenet.krcg.org/t/2GOLIrXAF8M/#m1) [[LSJ
    20080630]](https://usenet.krcg.org/t/nvuXBpEaKAA/#m2) [[LSJ 20021022]](https://usenet.krcg.org/t/WIXqYpkKuj8/#m1) —
    group "Cancel" (G00058), {Asguresh}, {Infernal Pursuit}, {The Diamond Thunderbolt}, {React with Conviction}, {Steely
    Tenacity}.

[^1-9-10]: [[RTR 20000501]](https://usenet.krcg.org/t/MKrA0hBXuaU/#m0) [[LSJ
    20070411-2]](https://usenet.krcg.org/t/-B6BlRIT0Rg/#m8) — group "Move or reveal card from crypt" (G00126).

[^1-9-11]: [[ANK
    20180318]](https://usenet.krcg.org/t/vekn-76464/#m1)
    [[ANK 20231229]](https://usenet.krcg.org/t/vekn-81077/#m5) [[LSJ
    20100119]](https://usenet.krcg.org/t/1eULCGaVcO0/#m1) [[PIB
    20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) [[ANK
    20211020]](https://usenet.krcg.org/t/vekn-79416/#m1) — {Dreams of the Sphinx},
    {The Meddling of Semsith}, {Hagar Stone}, {Revelations}, {Sergei Voshkov, The Eye}.

[^1-10-1]: [[LSJ 20000113]](https://usenet.krcg.org/t/-qvB4bbD-rE/#m3) [[LSJ
    20021010]](https://usenet.krcg.org/t/IAamJNEJAug/) — {Soul Gem of Etrius}, {Abandoning the Flesh}; [[LSJ
    20090920]](https://usenet.krcg.org/t/mCPE3I6343Y/#m2) — {Absimiliard's Army}; [[ANK
    20180408]](https://usenet.krcg.org/t/vekn-76500/#m2) — {Charnas the Imp}.

[^1-10-2]: [[PIB
    20110810]](https://usenet.krcg.org/t/vekn-6809/#m14)
    — {Illusions of the Kindred}; [[LSJ 20060209]](https://usenet.krcg.org/t/ZuOfZorIhhU/#m4) — {Grasp the Ghostly};
    [[ANK 20170331]](https://usenet.krcg.org/t/vekn-75683/#m1) — {Summon
    Soul}; [[LSJ 20080314]](https://usenet.krcg.org/t/GqHN0I3c8Sc/#m1) — {Wider View}.

[^1-10-3]: [[LSJ 20011216-1]](https://usenet.krcg.org/t/xwirj771uGM/#m4) — {Resurrection}; [[RTR
    20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[ANK
    20200813]](https://usenet.krcg.org/t/vekn-78798/#m2) — group "Shuffle
    when burned" (G00056); — {Set's Curse}.

[^1-10-4]: [[ANK 20190422]](https://usenet.krcg.org/t/vekn-77571/#m2) — {Set's
    Curse}; [[LSJ 20100423]](https://usenet.krcg.org/t/YnQCu0GeMhc/#m1) — {Sacrificial Lamb}; [[LSJ
    19981006]](https://usenet.krcg.org/t/RU5yM2Ov5Mg/#m0) [[LSJ 20001127-2]](https://usenet.krcg.org/t/KInac4MQMuA/#m4)
    [[LSJ 20010806-1]](https://usenet.krcg.org/t/PuawBcgSIKI/#m5) — {Bomb}.

[^1-10-5]: [[LSJ 20070307]](https://usenet.krcg.org/t/QlMNPzNtXh8/#m1) [[LSJ
    20070308]](https://usenet.krcg.org/t/QlMNPzNtXh8/#m6) — {Riddle Phantastique}, {Secret Horde}; [[LSJ
    20001118]](https://usenet.krcg.org/t/r59jQZbWMi0/#m2) — {Dominique}.

[^1-10-6]: [[LSJ 20041201]](https://usenet.krcg.org/t/DY1OggZnMvs/#m5) — group "From ash heap to deck" (G00002); [[LSJ
    20071014]](https://usenet.krcg.org/t/Bom6ae7qjbI/#m3) — {Kaymakli Fragment}; [[RTR
    20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0) — {Rachel Brandywine}.

[^1-10-7]: [[PIB
    20150512]](https://usenet.krcg.org/t/vekn-71020/#m2) —
    {The Black Throne}, {Provision of the Silsila}; [[LSJ 20031201-2]](https://usenet.krcg.org/t/6lN9MbBeL1E/#m1) —
    {Taste of Vitae}; [[LSJ 20040726]](https://usenet.krcg.org/t/LlqCB6LN64g/#m7) [[ANK
    20191031]](https://usenet.krcg.org/t/vekn-11417/#m10) — {The Eternal
    Mask}.

[^1-11-7]: [[LSJ 20090710]](https://usenet.krcg.org/t/PBi3na64nkw/#m1) [[ANK
    20200523]](https://usenet.krcg.org/t/vekn-78648/#m1)
    — {Trochomancy}, {Maabara}.

[^1-11-1]: [[LSJ 20090430]](https://usenet.krcg.org/t/2STzSnaSgxU/#m4) [[ANK
    20180628]](https://usenet.krcg.org/t/vekn-76758/#m6) [[LSJ
    20041119]](https://usenet.krcg.org/t/HMV_pN5mJhQ/#m3) [[LSJ 19980212]](https://usenet.krcg.org/t/fLFLlXZXHqA/#m0)
    [[ANK 20190701]](https://usenet.krcg.org/t/vekn-77763/#m2) [[LSJ
    20041115]](https://usenet.krcg.org/t/ZCmZl5xUhis/#m2) — {Epikasta Rigatos}, {Marthe Dizier}, {The Art of Memory}.

[^1-11-2]: [[ANK 20210309-2]](https://usenet.krcg.org/t/vekn-79005/#m7)
    [[LSJ 20020911]](https://usenet.krcg.org/t/_lLSO5aevoM/#m1) — group "Vote playable once per game" (G00030), {Echo of
    Harmonies}.

[^1-11-3]: [[ANK
    20220805]](https://usenet.krcg.org/t/vekn-79939/#m3)
    [[LSJ 20050804]](https://usenet.krcg.org/t/Xqa8boGP7C8/#m5) — {Melange}, {Darksight}, {Ghouled}, {Echo of
    Harmonies}.

[^1-11-4]: [[LSJ 20050324]](https://usenet.krcg.org/t/vhSqnoAt_Bs/#m14) [[LSJ
    20050322]](https://usenet.krcg.org/t/vhSqnoAt_Bs/#m2) [[RTR 19980928]](https://usenet.krcg.org/t/Xva4_IRavxM/#m0)
    [[LSJ 19980929]](https://usenet.krcg.org/t/Xva4_IRavxM/#m5) — {Delaying Tactics}, {The Erciyes Fragments}, {Compel
    the Spirit}.

[^1-11-5]: [[LSJ 20020115]](https://usenet.krcg.org/t/wG_tDLgfZso/#m1) [[LSJ
    20071014]](https://usenet.krcg.org/t/Bom6ae7qjbI/#m3) [[LSJ 20030128]](https://usenet.krcg.org/t/d25nONe01WY/#m1)
    [[LSJ 20011216-1]](https://usenet.krcg.org/t/xwirj771uGM/#m4) — group "Action creating vampire" (G00054), {Jake
    Washington}, {Spell of Life}, {Compel the Spirit}, {Resurrection}.

[^1-11-6]: [[ANK
    20200417]](https://usenet.krcg.org/t/vekn-78568/#m16)
    [[LSJ 20060209]](https://usenet.krcg.org/t/ZuOfZorIhhU/#m4) [[LSJ
    20040812]](https://usenet.krcg.org/t/n_0DGDsWG0E/#m1) [[LSJ 20010627]](https://usenet.krcg.org/t/NhNCVCCDyU0/#m5) —
    {The Capuchin}, {Grasp the Ghostly}, {Père Lachaise, France}, {Khazar's Diary (Endless Night)}.

[^1-12-1]: [[LSJ 20021128]](https://usenet.krcg.org/t/6rqhybcysh8/#m1) [[SFC
    19960919]](https://usenet.krcg.org/t/_UxVq0Lrg2U/#m3) [[LSJ 20100302-2]](https://usenet.krcg.org/t/b2cW6X-RQMs/#m3)
    [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[ANK
    20200417]](https://usenet.krcg.org/t/vekn-78568/#m16)
    [[RBK important-terms-of-the-game]](https://www.vekn.net/rulebook#important-terms-of-the-game) — group "Cards
    holding cards" (G00028), {Memory's Fading Glimpse}, {Raw Recruit}, {Betrayer}, {The Capuchin}.

[^1-12-8]: [[LSJ 20080816]](https://usenet.krcg.org/t/FKBTEzLf0_A/#m5) [[RBK
    contested-cards]](https://www.vekn.net/rulebook#contested-cards) — {Descent into Darkness}.

[^1-12-2]: [[ANK 20180511-1]](https://usenet.krcg.org/t/vekn-76594/#m0) — {Righteous
    Aura}.

[^1-12-3]: [[LSJ 20071206]](https://usenet.krcg.org/t/B-5oDAJrwiI/#m3) [[LSJ
    20090725]](https://usenet.krcg.org/t/isG743Sws-8/#m1) [[LSJ 20010616]](https://usenet.krcg.org/t/h7gnHNBFliE/#m12)
    [[LSJ 20030522-1]](https://usenet.krcg.org/t/_krZG-uPtzc/#m19) — {The Eternal Mask}, {Tegyrius, Vizier}.

[^1-12-4]: [[ANK
    20190629]](https://usenet.krcg.org/t/vekn-77740/#m10) — {Shadow
    Court Satyr}.

[^1-12-5]: [[PIB 20121028]](https://usenet.krcg.org/t/vekn-39816/#m10)
    [[LSJ 20001031]](https://usenet.krcg.org/t/y8LNZhRyXO0/#m2) [[PIB
    20121220]](https://usenet.krcg.org/t/vekn-43188/#m8) —
    {Anathema}, {Storage Annex}.

[^1-12-6]: [[ANK 20221028]](https://usenet.krcg.org/t/vekn-80119/#m4) [[LSJ
    20090428-1]](https://usenet.krcg.org/t/887DQTpntKI/#m1) — {Mokolé Blood}, {Shilmulo Tarot}, {Elder Library}.

[^1-12-7]: [[LSJ 20070307]](https://usenet.krcg.org/t/QlMNPzNtXh8/#m1) [[LSJ
    20070308]](https://usenet.krcg.org/t/QlMNPzNtXh8/#m6) [[RTR 19970306]](https://usenet.krcg.org/t/1dlmpgX6t14/#m0)
    [[LSJ 20001118]](https://usenet.krcg.org/t/r59jQZbWMi0/#m2) — {Riddle Phantastique}, {Secret Horde}, {Goth Band},
    {Dominique}.

[^1-13-1]: [[LSJ 20070928-2]](https://usenet.krcg.org/t/duRrP46XygI/#m43) [[LSJ
    20071001-2]](https://usenet.krcg.org/t/XoMeEYJw1ZA/#m10) [[RBK
    contested-cards]](https://www.vekn.net/rulebook#contested-cards) — {Spell of Life}.

[^1-13-2]: [[LSJ 20100213]](https://usenet.krcg.org/t/vXDkYrTmkws/#m2) [[LSJ
    19991110]](https://usenet.krcg.org/t/5hDqAMewLtg/#m2) [[LSJ 20030419]](https://usenet.krcg.org/t/A0mvllC-tgs/#m5)
    [[RTR 20180303]](https://usenet.krcg.org/t/vekn-76447/#m0) —
    {Illusions of the Kindred}, {Jimmy Dunn}.

[^1-13-3]: [[PIB
    20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) [[LSJ
    20040722]](https://usenet.krcg.org/t/z3Tfg9PZVNo/#m5) [[LSJ 20010623-2]](https://usenet.krcg.org/t/65IHHAii7ms/#m6)
    [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[LSJ
    20030918]](https://usenet.krcg.org/t/Lae6k2AG-1c/#m5) [[RBK
    contested-cards]](https://www.vekn.net/rulebook#contested-cards) — {Visit from the Capuchin}, {Chain of Command},
    {Parmenides}, {Soul Gem of Etrius}.

[^1-13-4]: [[LSJ 20080526]](https://usenet.krcg.org/t/zO3l3b76rjQ/#m4) [[LSJ
    20100902]](https://usenet.krcg.org/t/SM2_578Th0U/#m17) [[ANK
    20180626-1]](https://usenet.krcg.org/t/vekn-76751/#m4) [[ANK
    20211214]](https://usenet.krcg.org/t/vekn-79542/#m7) [[ANK
    20221028]](https://usenet.krcg.org/t/vekn-80119/#m4) [[LSJ
    20090428-1]](https://usenet.krcg.org/t/887DQTpntKI/#m1) — {Wormwood}, {Anarch Convert}, {Dr. Solomon Grey}, {Erlik},
    {Mokolé Blood}, {Shilmulo Tarot}.

[^1-13-5]: [[LSJ 20070829-2]](https://usenet.krcg.org/t/drn7wHaGugQ/#m3) — {Sonja Blue}.

[^1-13-6]: [[ANK 20221028]](https://usenet.krcg.org/t/vekn-80119/#m4) [[LSJ
    20080107]](https://usenet.krcg.org/t/XpZ6F53jK-c/#m3) [[RTR 19960124]](https://usenet.krcg.org/t/wF82VdVPlm0/#m0)
    [[LSJ 20100423]](https://usenet.krcg.org/t/YnQCu0GeMhc/#m1) [[LSJ
    19980319]](https://usenet.krcg.org/t/i1Eqqm5Ctv0/#m1) — {Elder Library}, {Guillaume Giovanni}, {Secure Haven},
    {Byzar}, {Disguised Weapon}.

[^1-13-12]: Owner's reading, no official ruling: [[rules-qa
    83.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-83--the-cards-on-a-contested-vampire-z-main-2026-10-18-from-the-owners-playtest-via-z-play-owner-2026-10-19-831-confirmed)
    — {Soul Gem of Etrius}.

[^1-13-7]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[LSJ
    20080816]](https://usenet.krcg.org/t/FKBTEzLf0_A/#m5) — {Betrayer}, {Descent into Darkness}.

[^1-13-10]: [[RTR 19960112]](https://usenet.krcg.org/t/d3n3StNS7no/#m0) [[TOM
    19960122]](https://usenet.krcg.org/t/d3n3StNS7no/#m24) — {The Rack}.

[^1-13-11]: [[LSJ 19990719]](https://usenet.krcg.org/t/kKtW2qo6ACE/#m1) [[RTR
    19991001]](https://usenet.krcg.org/t/RAvWWmYoX3U/#m0) — {Jimmy Dunn}.

[^1-14-1]: [[LSJ 20010627]](https://usenet.krcg.org/t/NhNCVCCDyU0/#m5) [[LSJ
    20050422]](https://usenet.krcg.org/t/nJPfMOuTBtw/#m3) — {The Sargon Fragment}, {Pochtli}, {Sudario Refraction}; same
    wording on {Clio's Kiss}, {Gear Up}, {Pandora's Whisper}, {Scarlet Lore}, {Whispers from the Dead}, {Drozodny},
    {Patrizia Giovanni, Collector of Secrets}.

[^1-14-2]: [[LSJ 20091030]](https://usenet.krcg.org/t/ZKuCyTayYbc/#m1) — {Magic of the Smith}, {Ashur Tablets}.

[^1-14-3]: [[PIB 20110725]](https://usenet.krcg.org/t/vekn-6728/#m4) [[LSJ
    20090722]](https://usenet.krcg.org/t/r-N65rA52uo/#m1) — {Siphon}, {Sudario Refraction}.

[^1-14-4]: [[LSJ 19991207]](https://usenet.krcg.org/t/aQaOTYwC-fg/#m3) [[ANK
    20180925-2]](https://usenet.krcg.org/t/vekn-77029/#m10)
    [[LSJ 20100130]](https://usenet.krcg.org/t/X8Uu7Sk56P4/#m16) [[RTR
    19991001]](https://usenet.krcg.org/t/RAvWWmYoX3U/#m0) [[LSJ 20101210]](https://usenet.krcg.org/t/QTAA0y6cANI/#m11)
    [[ANK
    20220704]](https://usenet.krcg.org/t/vekn-79890/#m1)
    [[PIB 20150428]](https://usenet.krcg.org/t/vekn-70714/#m9) [[LSJ
    20100206-2]](https://usenet.krcg.org/t/reXyybyIYX8/#m1) — {Disguised Weapon}, {Concealed Weapon}, {Charming Lobby},
    {Jack of Both Sides}, {Gift of Proteus}, {Blessing of the Beast}.

[^1-14-5]: [[ANK
    20201029]](https://usenet.krcg.org/t/vekn-78889/#m4) [[LSJ
    20041130]](https://usenet.krcg.org/t/6uTPqRg387A/#m3) [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) —
    {Charming Lobby}, {Delaying Tactics}.

[^1-14-6]: [[LSJ 20050118]](https://usenet.krcg.org/t/W784_3oyz7I/#m2) — {The Erciyes Fragments}.

[^1-14-7]: [[LSJ 20100211-2]](https://usenet.krcg.org/t/pL63VXEPGME/) — {Baal's Bloody Talons}, {Cleave}.

[^1-15-1]: [[LSJ 19980225]](https://usenet.krcg.org/t/62y-5miA8MQ/#m0) [[RTR
    19941222]](https://usenet.krcg.org/t/19Ys4rdQPQw/#m0) [[LSJ 20061018]](https://usenet.krcg.org/t/Prc45fTQd9Y/#m5)
    [[ANK 20200703-4]](https://usenet.krcg.org/t/vekn-78709/#m3) [[RTR
    19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) — {Aching Beauty}, {Anarch Revolt}, {Archon}, {Camarilla
    Exemplary}, {Sabbat Priest}, {First Tradition: The Masquerade}, {Frontal Assault}, {Judgment: Camarilla
    Segregation}, {Betrayer}.

[^1-15-2]: [[LSJ 20020904]](https://usenet.krcg.org/t/Qmi4hFk6QqE/#m2) [[LSJ
    20010819]](https://usenet.krcg.org/t/2dXfHhG98Rs/#m2) [[PIB
    20130723]](https://usenet.krcg.org/t/vekn-52017/#m4) [[PIB
    20130303]](https://usenet.krcg.org/t/vekn-45620/#m4) [[LSJ
    20100813]](https://usenet.krcg.org/t/nOb3cmvA_3U/#m1) [[LSJ 20051204]](https://usenet.krcg.org/t/K-RE-nTJ_Cg/#m4)
    [[LSJ 20060407]](https://usenet.krcg.org/t/YY2aBVshruI/#m10) [[LSJ
    20090623]](https://usenet.krcg.org/t/WqwoI03BW1U/#m1) [[LSJ 20100601]](https://usenet.krcg.org/t/R7YgwD0VlUQ/#m1) —
    {Mass Reality}, {Legacy of Pander}, {Torn Signpost}, {Orun}, {Xeper, Sultan of Lepers}.

[^1-15-3]: [[LSJ 19970707]](https://usenet.krcg.org/t/KWekwiRSa2I/#m1) [[LSJ
    20001102]](https://usenet.krcg.org/t/LlPyLJjLdx0/#m2) [[RTR 19951017]](https://usenet.krcg.org/t/ouhNUbHYg50/#m2)
    [[LSJ 20010227]](https://usenet.krcg.org/t/rdkVAtZFC2I/#m1) — {Camarilla Exemplary}, {Courier}, {Kindred Society
    Games}, {Temptation}.

[^1-15-4]: [[LSJ 20070306]](https://usenet.krcg.org/t/QiFXAaAeJds/#m4) [[LSJ
    20030820]](https://usenet.krcg.org/t/Zjmffk-EbJw/#m1) [[LSJ 20090603]](https://usenet.krcg.org/t/LpFSLRuWONA/#m13)
    [[RTR 19951110]](https://usenet.krcg.org/t/TXfganI5B2o/#m0) — {Dr. Morrow, The Skindoctor}, {Kahina the Sorceress},
    {Orun}, {Scorn of Adonis}.

[^1-15-5]: [[ANK 20180111]](https://usenet.krcg.org/t/vekn-76364/#m3) —
    {Leech}, {Putrescent Servitude}.

[^1-15-6]: [[ANK 20200515]](https://usenet.krcg.org/t/vekn-78634/#m1) — {Blanket of
    Night}, {Cloak the Gathering}, {Inspire Greatness}, {Mask of a Thousand Faces}, {Suppressing Fire}.

[^1-15-7]: [[LSJ 20041103]](https://usenet.krcg.org/t/MiPHVp-NmCA/#m5) — {Lunatic Eruption}.

[^2-1-1]: [[LSJ 20100723]](https://usenet.krcg.org/t/0u5KQWiutdg/#m1) [[ANK
    20200616]](https://usenet.krcg.org/t/vekn-78687/#m1)
    [[PIB
    20120808]](https://usenet.krcg.org/t/vekn-34265/#m17)
    [[RBK targeting-of-cards]](https://www.vekn.net/rulebook#targeting-of-cards) — {Owain Evans, The Wanderer}; group
    "Can draw during action" (G00023); group "Cancel with no action" (G00063).

[^2-1-10]: Owner's reading, no official ruling: [[rules-qa
    26.5]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-26--actions-wave-1) — {Rutor's
    Hand}, {Qawiyya el-Ghaduba}.

[^2-1-2]: [[LSJ 19990425]](https://usenet.krcg.org/t/Mmfn07ib6Yw/#m1) [[SFC
    19960819]](https://usenet.krcg.org/t/G40EE8vCBB8/#m2) [[RBK
    summary-of-the-course-of-an-action]](https://www.vekn.net/rulebook#summary-of-the-course-of-an-action) — {Blanket of
    Night}, {Siren's Lure}, {Veil the Legions}, {Empowering the Puppet King}, {Inspire Greatness}, {Spiridonas}.

[^2-1-9]: [[LSJ 19990425]](https://usenet.krcg.org/t/Mmfn07ib6Yw/#m1) [[ANK
    20200515]](https://usenet.krcg.org/t/vekn-78634/#m1) — {Blanket of Night}, {Cloak the
    Gathering}, {Hidden Lurker}, {Zapaderin}.

[^2-1-3]: [[ANK 20181101]](https://usenet.krcg.org/t/vekn-77132/#m1) [[LSJ
    20070403]](https://usenet.krcg.org/t/TJ2ktt_1tjk/#m9) [[LSJ 20070413]](https://usenet.krcg.org/t/umdINigMKqs/#m19) —
    {Bliss}, {Hide}, {Surge}, {Vigilance}.

[^2-1-4]: [[LSJ 20010110]](https://usenet.krcg.org/t/Vtm465mblX4/#m21) [[LSJ
    20050111]](https://usenet.krcg.org/t/lnW6nMIX-Vw/#m7) [[PIB
    20121028]](https://usenet.krcg.org/t/vekn-39816/#m10) [[LSJ
    20001111]](https://usenet.krcg.org/t/m23Hj3OW2A4/#m1) [[PIB
    20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) —
    {Heidelberg Castle, Germany}, {Baltimore Purge}, {Temptation}, {Vagabond Mystic}, {The Coven}.

[^2-1-8]: [[ANK 20250121]](https://usenet.krcg.org/t/vekn-79080/#m7)
    — {Nightmares upon Nightmares}, {Arika}.

[^2-1-5]: [[ANK
    20210608]](https://usenet.krcg.org/t/vekn-79166/#m13)
    [[ANK
    20170918-2]](https://usenet.krcg.org/t/vekn-76178/#m2)
    — {Heidelberg Castle, Germany}, {The Louvre, Paris}, {Siren's Lure}.

[^2-1-6]: [[RTR 19960112]](https://usenet.krcg.org/t/d3n3StNS7no/#m0) [[LSJ
    19990217]](https://usenet.krcg.org/t/9Bsf2LC1274/#m1) [[RTR 20041202]](https://usenet.krcg.org/t/WUWh7AdooDU/#m5)
    [[LSJ 20100119]](https://usenet.krcg.org/t/1eULCGaVcO0/#m1) [[ANK
    20201228]](https://usenet.krcg.org/t/vekn-78956/#m1) [[RTR
    19991206]](https://usenet.krcg.org/t/N7iEmqgP9WU/#m0) [[ANK
    20230620]](https://usenet.krcg.org/t/vekn-80612/#m1)
    [[ANK 20180531]](https://usenet.krcg.org/t/vekn-76673/#m2) — {Spirit Claws},
    {Rötschreck}, {Slake the Thirst}, {Soul Stealing}, {Powerbase: Tshwane}, {Melange}.

[^2-1-7]: [[LSJ 19980105]](https://usenet.krcg.org/t/PzVC-AeFuUQ/#m1) [[LSJ
    20090205]](https://usenet.krcg.org/t/l5bUtHOejmc/#m2) — {Deflection}, {Redirection}, {Lost in Translation}, {My
    Enemy's Enemy}, {Major Boon}.

[^2-2-1]: [[ANK
    20190425]](https://usenet.krcg.org/t/vekn-77581/#m1)
    [[RBK summary-of-the-course-of-an-action]](https://www.vekn.net/rulebook#summary-of-the-course-of-an-action) — group
    "Modifier as announced" (G00052); {Force of Personality}, {Shroud Mastery}, {Hide}.

[^2-2-2]: [[LSJ 20081105]](https://usenet.krcg.org/t/CYjOJtTBMGU/#m1) [[PIB
    20150820]](https://usenet.krcg.org/t/vekn-72462/#m1) — {Yoruba
    Shrine}, {Car Bomb}.

[^2-2-3]: [[ANK
    20181208]](https://usenet.krcg.org/t/vekn-77210/#m7)
    — group "Modifier as announced" (G00052); {Approximation of Loyalty}, {Force of Personality}.

[^2-2-4]: [[LSJ 20100402]](https://usenet.krcg.org/t/BNPERBLJwRc/) — {Concoction of Vitality}.

[^2-3-1]: [[LSJ 19981028]](https://usenet.krcg.org/t/JZHY1bTCAa0/#m1) [[ANK
    20190425]](https://usenet.krcg.org/t/vekn-77581/#m1)
    — group "Modifier after resolution" (G00006); {Domain of Evernight}, {Forced March}, {Fata Amria}. [[LSJ
    20030219]](https://usenet.krcg.org/t/ugfckv9DAbo/#m5) — {Spirit Marionette}. [[LSJ
    20090731]](https://usenet.krcg.org/t/y6f0s6tUtqs/#m6) — {Last Stand}. [[ANK
    20180909]](https://usenet.krcg.org/t/vekn-76987/#m2)
    — {Ophidian Gaze}.

[^2-3-2]: [[LSJ 20050422]](https://usenet.krcg.org/t/nJPfMOuTBtw/#m3) — {Dis Pater}. [[ANK
    20190113]](https://usenet.krcg.org/t/vekn-77290/#m12) [[ANK
    20190425]](https://usenet.krcg.org/t/vekn-77581/#m1)
    — {Capitalist}, {Cavalier}, {Perfectionist}, {Scrying of Secrets}, {Truth of a Thousand Lies}, {Vigilance}.

[^2-3-3]: [[LSJ 20110502]](https://usenet.krcg.org/t/bgg-648695/#m1) [[LSJ
    19980105]](https://usenet.krcg.org/t/PzVC-AeFuUQ/#m1) — {Spying Mission}.

[^2-3-4]: [[LSJ 20041022]](https://usenet.krcg.org/t/gqhND6kd2wE/#m3) [[RTR
    20180719]](https://www.blackchantry.com/2018/07/18/rules-team-rulings-rtr-19-07-2018/) [[ANK
    20220118]](https://usenet.krcg.org/t/vekn-79389/#m11) [[ANK
    20221205]](https://usenet.krcg.org/t/vekn-80196/#m1) —
    {Champion}, {Car Bomb}, {Ensconced}, {Detect Authority}, {Faerie Wards}, {Final Loosening}, {Hide}, {Hide the
    Heart}, {Malkavian Derangement: Alternate Personality}, {Purification}, {Scobax}. [[LSJ
    20010803-1]](https://usenet.krcg.org/t/5QuGIF5ERUI/#m1) [[ANK
    20200207]](https://usenet.krcg.org/t/vekn-78423/#m4) — {Change of Target},
    group "Action ended by an effect" (G00166).

[^2-3-5]: [[LSJ 19981028]](https://usenet.krcg.org/t/JZHY1bTCAa0/#m1) [[LSJ
    20100206-2]](https://usenet.krcg.org/t/reXyybyIYX8/#m1) — {Momentary Delay}.

[^2-3-6]: [[ANK
    20180206]](https://usenet.krcg.org/t/vekn-76404/#m4) — {The
    Damned}, {Ecstasy}, {Travelers Obey the Tenets}, {Shemti}. [[LSJ
    20100324]](https://usenet.krcg.org/t/PWhoejXDuuA/#m1) [[LSJ 20071116]](https://usenet.krcg.org/t/ZPsShBUBWZI/#m1) —
    {Games of Instinct}, {Save Face}.

[^2-3-7]: [[ANK 20230314]](https://usenet.krcg.org/t/vekn-80378/#m2) —
    {Provision of the Silsila}, {Monster}, {Shadow Boxing}.

[^2-3-8]: [[LSJ 20010715]](https://usenet.krcg.org/t/u91PfJOfOjE/#m1) — {Force of Will}. [[ANK
    20171124]](https://usenet.krcg.org/t/vekn-76304/#m3)
    — {Heidelberg Castle, Germany}. [[LSJ 20100325]](https://usenet.krcg.org/t/PWhoejXDuuA/#m6) — {Slake the Thirst}.
    [[PIB 20150915]](https://usenet.krcg.org/t/vekn-73134/#m3) — {Freak Drive}.

[^2-3-9]: [[LSJ 20091123]](https://usenet.krcg.org/t/Yt7LsvWFwiM/) — {Wake with Evening's Freshness}, {Eyes of Argus},
    {Forced Awakening}, {On the Qui Vive}.

[^2-3-10]: [[LSJ 20030130]](https://usenet.krcg.org/t/TUDO_4FwdyY/#m1) [[ANK
    20181219]](https://usenet.krcg.org/t/vekn-77232/#m2) [[ANK
    20180906]](https://usenet.krcg.org/t/vekn-76981/#m2) — group
    "Modifier after combat" (G00007); [[RTR 19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0) — {Freak Drive}, {Hay
    Ride}; [[LSJ 20040219]](https://usenet.krcg.org/t/zFmoLa6tzWA/#m2) [[ANK
    20221130]](https://usenet.krcg.org/t/vekn-80176/#m3) —
    {Cats' Guidance}.

[^2-3-14]: [[ANK 20250716]](https://usenet.krcg.org/t/vekn-82396/#m16) — {Cats'
    Guidance}.

[^2-3-11]: [[ANK 20200221]](https://usenet.krcg.org/t/vekn-78458/#m2)
    — {Strix}. [[LSJ 20060426]](https://usenet.krcg.org/t/4e6z1_JWIzA/#m1) — {High Aye}. [[LSJ
    20031112]](https://usenet.krcg.org/t/NW4HWUWlzDI/#m13) — {The Art of Memory}.

[^2-3-13]: [[LSJ 19991025]](https://usenet.krcg.org/t/R94tyTGJ6VQ/#m0) [[LSJ
    20030530]](https://usenet.krcg.org/t/SZehI8SwAc4/#m21) [[ANK
    20250710]](https://usenet.krcg.org/t/vekn-82396/#m13) — {Obedience}, {Raptor};
    [[ANK 20250613]](https://usenet.krcg.org/t/vekn-82396/#m0) — {Cats' Guidance},
    {Hedonism}, {Psyche!}.

[^2-3-12]: [[LSJ 20100112]](https://usenet.krcg.org/t/SJu0kgw_2tE/#m1) [[PIB
    20121028]](https://usenet.krcg.org/t/vekn-39816/#m10) — {Abactor}. [[LSJ
    20090722-2]](https://usenet.krcg.org/t/Ry0xU4IuJmQ/#m1) [[LSJ 20030618]](https://usenet.krcg.org/t/AdfUNNicx-Y/#m16)
    — {Heidelberg Castle, Germany}. [[LSJ 20090115-2]](https://usenet.krcg.org/t/rSaiVPMbpvY/#m0) — {Guru}.

[^2-4-1]: [[ANK
    20180206]](https://usenet.krcg.org/t/vekn-76404/#m4) [[LSJ
    20020904-2]](https://usenet.krcg.org/t/ObuKimgcCpI/#m10) [[LSJ 20100211-2]](https://usenet.krcg.org/t/pL63VXEPGME/)
    [[LSJ 20020301]](https://usenet.krcg.org/t/HgmPj9MEdiQ/#m4) [[RBK
    sequencing]](https://www.vekn.net/rulebook#sequencing) — {The Damned}, {Ecstasy}, {Travelers Obey the Tenets},
    {Shemti}, {Dreams of the Sphinx}, {Cleave}, {Darkness Within}.

[^2-4-2]: [[LSJ 20021113]](https://usenet.krcg.org/t/df2P8YHZex8/#m11) [[ANK
    20191219]](https://usenet.krcg.org/t/vekn-78241/#m3)
    [[ANK
    20180910-1]](https://usenet.krcg.org/t/vekn-76990/#m3)
    [[RTR 20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0) — {Disarm}, {Pulled Fangs}, {Street Cred}, {Taste of
    Vitae}, {Relentless Reaper}, {Psyche!}, {Telepathic Tracking}.

[^2-4-3]: [[LSJ 20100218]](https://usenet.krcg.org/t/x-uJChHi76Y/#m2) [[RTR
    20180511-2]](https://usenet.krcg.org/t/vekn-76595/#m30)
    [[LSJ 20001206]](https://usenet.krcg.org/t/kFIO74LxqFQ/#m4) [[ANK
    20211124]](https://usenet.krcg.org/t/vekn-79501/#m2) — {Power of One}, {Quicksilver
    Contemplation}, {Quickness}, {Hell-for-Leather}, {Command of the Beast}, {Divine Image}, {Leverage}.

[^2-4-4]: [[ANK
    20191219]](https://usenet.krcg.org/t/vekn-78241/#m3)
    — {Telepathic Tracking}.

[^2-4-5]: [[TOM 19960413]](https://usenet.krcg.org/t/Gm-NLCP6bF0/#m8) [[LSJ
    20030213]](https://usenet.krcg.org/t/j6cuQ6pFJSA/#m1) — {Protected Resources}, {Merrill Molitor}.

[^2-4-25]: Owner's reading, no official ruling: [[rules-qa
    6.4]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-6--cards-in-play-riders-granted-actions)
    — {Fame}.

[^2-4-9]: [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[LSJ
    20090304-2]](https://usenet.krcg.org/t/PcbRGxbYQUY/#m3) [[PIB
    20120214]](https://usenet.krcg.org/t/vekn-22906/#m14) — {Amaranth},
    {Decapitate}, {Mustafa, The Heir}, {Squirrel Balance}, {Storm Sewers}.

[^2-4-10]: [[LSJ 20021117]](https://usenet.krcg.org/t/4TgUatcTtdk/#m3) [[PIB
    20120225]](https://usenet.krcg.org/t/vekn-24105/#m1) [[LSJ
    20011215]](https://usenet.krcg.org/t/9WJX_WF656A/#m5) [[LSJ 20050805]](https://usenet.krcg.org/t/mJzy9HKytEc/#m2) —
    {Anathema}, {Blissful Agony}, {Dual Form}.

[^2-4-11]: [[ANK
    20200709]](https://usenet.krcg.org/t/vekn-74710/#m4)
    [[ANK
    20181016]](https://usenet.krcg.org/t/vekn-77075/#m9)
    [[ANK 20170625]](https://usenet.krcg.org/t/vekn-75929/#m1) [[ANK
    20210110]](https://usenet.krcg.org/t/vekn-78984/#m10)
    — {Jenna Cross}, {Judgment: Camarilla Segregation}, {Khobar Towers, Al-Khubar}, {First Tradition: The Masquerade}.

[^2-4-26]: [[LSJ 20090208-2]](https://usenet.krcg.org/t/kz7qCu53zvI/#m1) — {Tension in the Ranks}.

[^2-4-12]: [[LSJ 20100203]](https://usenet.krcg.org/t/HQD-UNP7Y0s/#m2) [[ANK
    20200905]](https://usenet.krcg.org/t/vekn-78836/#m1) [[ANK
    20220503]](https://usenet.krcg.org/t/vekn-39040/#m5)
    [[LSJ 20060530-1]](https://usenet.krcg.org/t/7ezcZoziFWw/#m2) [[LSJ
    20090617-2]](https://usenet.krcg.org/t/ClO_zsSSgVU/#m1) — {Slake the Thirst}, {Visit from the Capuchin}, {Drink the
    Blood of Ahriman}, {Zhenga}, {Ilomba}.

[^2-4-13]: [[ANK 20180725]](https://usenet.krcg.org/t/vekn-76858/#m1) [[LSJ
    20050116]](https://usenet.krcg.org/t/yX5rqVaarvs/#m2) [[ANK
    20240724]](https://usenet.krcg.org/t/vekn-81635/#m6) —
    {Feline Saboteur}, {Dual Form}, {Illusions of the Kindred}.

[^2-4-18]: [[ANK 20210627]](https://usenet.krcg.org/t/vekn-79192/#m14)
    [[ANK
    20211207]](https://usenet.krcg.org/t/vekn-79528/#m7)
    [[ANK
    20220116]](https://usenet.krcg.org/t/vekn-79600/#m1)
    [[ANK
    20190524]](https://usenet.krcg.org/t/vekn-77659/#m1)
    [[LSJ 20020126]](https://usenet.krcg.org/t/Fhm2Zi2RZRU/#m1) — group "Triggered by block" (G00105), {Unleash Hell's
    Fury}, {Millicent Smith, Puritan Vampire Hunter}, {Banshee Ironwail}, {Venenation}.

[^2-4-20]: [[ANK
    20200629-2]](https://usenet.krcg.org/t/vekn-78701/#m2)
    [[LSJ 20091208]](https://usenet.krcg.org/t/ptHbJM9MlVI/#m1) [[RTR
    19951017]](https://usenet.krcg.org/t/ouhNUbHYg50/#m2) — group "Replace next turn" (G00025); {Port Authority},
    {Malkavian Dementia}.

[^2-4-21]: [[LSJ 20110121]](https://usenet.krcg.org/t/bgg-609699/#m1) [[ANK
    20200508-1]](https://usenet.krcg.org/t/vekn-78622/#m17) —
    {Fame}, {Scourge of the Enochians}.

[^2-4-22]: [[ANK 20200123]](https://usenet.krcg.org/t/vekn-78373/#m1)
    [[LSJ 20041027]](https://usenet.krcg.org/t/BHeGvhd4yEA/#m2) — {Theft of Vitae}.

[^2-4-23]: [[ANK 20200702]](https://usenet.krcg.org/t/vekn-78711/#m1) [[ANK
    20180627-2]](https://usenet.krcg.org/t/vekn-76756/#m1) — {Deep Song}, {Flames
    of Insurrection}.

[^2-4-24]: [[ANK
    20170427]](https://usenet.krcg.org/t/vekn-75755/#m6) [[LSJ
    20040805]](https://usenet.krcg.org/t/WuER8RUMzTE/#m13) [[LSJ 20090528-1]](https://usenet.krcg.org/t/63JZWOiuAIQ/#m5)
    — {Blood of Acid}, {Blood Shield}.

[^2-5-1]: [[LSJ 20060412]](https://usenet.krcg.org/t/DuZXjDEQ9cE/#m14) [[LSJ
    20070217]](https://usenet.krcg.org/t/HkKuwBe9LRk/#m2) [[LSJ 20010814-2]](https://usenet.krcg.org/t/8MR4bq0Cxj4/#m3)
    [[ANK 20181127]](https://usenet.krcg.org/t/vekn-77176/#m9) —
    {Cleave}, {Ensconced}, {Champion}, {Guardian Vigil}.

[^2-5-2]: [[ANK 20191108]](https://usenet.krcg.org/t/vekn-78081/#m2) [[LSJ
    20100604-2]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m6) [[LSJ 20100611]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m11)
    [[LSJ 20050607-1]](https://usenet.krcg.org/t/vet0HGRxtXQ/#m1) [[LSJ
    20080818]](https://usenet.krcg.org/t/O5HUPRkKY-k/#m1) — {Eyes of the Beast}, {Marciana Giovanni, Investigator},
    {Valerius Maior, Hell's Fool}, {Blessing of Chaos}.

[^2-5-15]: Owner's reading, no official ruling: [[rules-qa
    53.4]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-53--sabbat-v5-the-path-of-cathari-implemented-readings-2026-10-07)
    [[rules-qa
    54.5]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-54--sabbat-v5-the-path-of-caine-implemented-readings-2026-10-07)
    [[rules-qa
    8.4]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-8--cards--effects-replacement-events)
    — {Fotini Katsikaris}, {Fiendish Tongue}, {Mirror Walk}.

[^2-5-3]: [[LSJ 20070707]](https://usenet.krcg.org/t/ZtRk5z2TcoI/#m1) [[RTR
    20201130]](https://www.blackchantry.com/2020/12/25/rtr-30-11-2020/) [[ANK
    20200415]](https://usenet.krcg.org/t/vekn-7516/#m9) [[ANK
    20221208-2]](https://usenet.krcg.org/t/vekn-80200/#m6) [[ANK
    20221024]](https://usenet.krcg.org/t/vekn-80110/#m4) — {Derange},
    {Malkavian Dementia}, {Leadership Vacuum}.

[^2-5-4]: [[ANK 20190416]](https://usenet.krcg.org/t/vekn-77560/#m3) [[ANK
    20180626-2]](https://usenet.krcg.org/t/vekn-76752/#m5)
    [[ANK 20230103]](https://usenet.krcg.org/t/vekn-80237/#m1) —
    {Teresita, The Godmother}, {Tegyrius, Vizier}, {Victim of Habit}.

[^2-5-16]: Owner's reading, no official ruling: [[rules-qa
    61.2]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-61--owner-play-feedback-2026-10-10)
    — {Enkil Cog}.

[^2-5-5]: [[LSJ 20090207]](https://usenet.krcg.org/t/0UGI9W2sSpk/#m1) [[LSJ
    20100212]](https://usenet.krcg.org/t/NIIFHQg_x1c/#m1) [[ANK
    20171212]](https://usenet.krcg.org/t/vekn-76334/#m14) [[LSJ
    20081120-1]](https://usenet.krcg.org/t/e2PNDpg-l_c/#m19) [[LSJ
    20091016]](https://usenet.krcg.org/t/pqa7mYZ6NEM/#m21) [[LSJ 20100226]](https://usenet.krcg.org/t/JnycCGrNQmY/#m3) —
    {Concordance}, {Cry Wolf}, {Mental Maze}, {Vidal Jarbeaux}.

[^2-5-7]: [[LSJ 20030811]](https://usenet.krcg.org/t/Ivy3fm45bb4/#m4) [[LSJ
    19970814]](https://usenet.krcg.org/t/Xd6HOjnqBpw/#m1) — {Annabelle Triabell}, {Toreador Grand Ball}.

[^2-5-8]: [[LSJ 20010616]](https://usenet.krcg.org/t/h7gnHNBFliE/#m12) [[LSJ
    20030522-1]](https://usenet.krcg.org/t/_krZG-uPtzc/#m19) [[LSJ 19970224]](https://usenet.krcg.org/t/80KRDjVFkyg/#m1)
    [[LSJ 20071005-1]](https://usenet.krcg.org/t/Ugcdb0ljZrU/#m1) [[LSJ
    20071206]](https://usenet.krcg.org/t/B-5oDAJrwiI/#m3) [[LSJ 20090725]](https://usenet.krcg.org/t/isG743Sws-8/#m1) —
    {Tegyrius, Vizier}, {Camarilla Vitae Slave}, {Agent of Power}, {The Eternal Mask}.

[^2-5-9]: [[LSJ 20100206-2]](https://usenet.krcg.org/t/reXyybyIYX8/#m1) [[LSJ
    20080619]](https://usenet.krcg.org/t/-JNytF94ST8/) [[LSJ 20021008]](https://usenet.krcg.org/t/Mc3xfym_uw8/#m2) [[ANK
    20180104]](https://usenet.krcg.org/t/vekn-76356/#m4)
    [[ANK 20191218]](https://usenet.krcg.org/t/vekn-62700/#m7) —
    {Shatter the Gate}, {NRA PAC}, {The Meddling of Semsith}, {Illusions of the Kindred}, {Nahir}.

[^2-5-10]: [[TOM 19960326]](https://usenet.krcg.org/t/8DA9p_p5v8s/#m2) [[LSJ
    20010610]](https://usenet.krcg.org/t/KVyVn-Y_UIY/#m2) [[LSJ 20040210]](https://usenet.krcg.org/t/FWzKVXDEJ5k/#m4) —
    {Frenzy}, {Kraken's Kiss}, {Sire's Index Finger}.

[^2-5-11]: [[LSJ 19980831]](https://usenet.krcg.org/t/Za8AS17xXPM/#m1) [[ANK
    20191025]](https://usenet.krcg.org/t/vekn-78050/#m1) — {Rowan Ring},
    {Wooden Stake}, {Rötschreck}.

[^2-5-12]: [[LSJ 20001019]](https://usenet.krcg.org/t/bSH8Q_kDhNQ/#m16) [[ANK
    20220525]](https://usenet.krcg.org/t/vekn-79777/#m5) [[LSJ
    20031106]](https://usenet.krcg.org/t/bFZCLXzzOeM/#m31) — {Horrid Reality}, {Imposing Phantasm}, {Rutor's Hand}.

[^2-5-13]: [[LSJ 20040726]](https://usenet.krcg.org/t/LlqCB6LN64g/#m7) [[ANK
    20200417]](https://usenet.krcg.org/t/vekn-78568/#m16)
    [[LSJ 20071014]](https://usenet.krcg.org/t/Bom6ae7qjbI/#m3) [[LSJ
    20051116-1]](https://usenet.krcg.org/t/MfGC7sJ8vh8/#m1) [[LSJ
    20081213-3]](https://usenet.krcg.org/t/cbZ2jl8-yGQ/#m1) — {Possession}, {The Capuchin}, {Kaymakli Fragment},
    {Demdemeh}.

[^2-5-14]: [[PIB 20140122]](https://usenet.krcg.org/t/vekn-58586/#m3)
    [[TOM 19951209]](https://usenet.krcg.org/t/qP2j6CpBUDI/#m6) [[LSJ
    20010809-3]](https://usenet.krcg.org/t/gLl8F0zcCF0/#m2) [[LSJ
    20010809-2]](https://usenet.krcg.org/t/9ggmJcK2De0/#m10) [[TOM
    19960210]](https://usenet.krcg.org/t/PiOmH08RyVw/#m10) [[RTR 20000501]](https://usenet.krcg.org/t/MKrA0hBXuaU/#m0)
    [[ANK 20210630]](https://usenet.krcg.org/t/vekn-79205/#m3) [[LSJ
    20010211]](https://usenet.krcg.org/t/Fd9tcsKTzjE/#m1) [[LSJ 20051103]](https://usenet.krcg.org/t/BD7KIWBI0Cs/#m1) —
    {Banishment}, {Lay Low}, {Brainwash}, {Dual Form}.

[^3-1-1]: [[RTR 20111202]](https://usenet.krcg.org/t/vekn-16769/#m0) [[RTR
    19980928]](https://usenet.krcg.org/t/Xva4_IRavxM/#m0) — {Ravnos Carnival}, {Travis "Traveler72" Miller}, {Betrayer}.

[^3-1-6]: [[ANK
    20240706]](https://usenet.krcg.org/t/vekn-81563/#m19) [[RTR
    19970425]](https://usenet.krcg.org/t/DhP_l2cX3mQ/#m0) — {Break the Bonds}, {Public Trust}, {Fire Dance}.

[^3-1-7]: [[ANK
    20240706]](https://usenet.krcg.org/t/vekn-81563/#m19) [[LSJ
    20080608]](https://usenet.krcg.org/t/j5ShUmUt-vM/#m1) [[PIB
    20110725]](https://usenet.krcg.org/t/vekn-6728/#m4) — {Keystone Kine}, {The Platinum
    Protocol}, {Magic of the Smith}, {Siphon}.

[^3-1-8]: [[LSJ 20041026]](https://usenet.krcg.org/t/54Zr4RUuVx8/#m2) [[ANK
    20240706]](https://usenet.krcg.org/t/vekn-81563/#m19) — group
    "Justicar title vote without Camarilla" (G00034), {Break the Bonds}.

[^3-1-9]: [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[PIB
    20110725]](https://usenet.krcg.org/t/vekn-6728/#m4) [[LSJ
    20100216]](https://usenet.krcg.org/t/nrXTh1XKJJ8/#m2) [[ANK
    20180512-2]](https://usenet.krcg.org/t/vekn-76595/#m25)
    [[LSJ 20100915]](https://usenet.krcg.org/t/lFLocuKV8NI/#m1) — {Siphon}, {Villein}, {Diversion}, {Donnybrook}, {Drain
    Essence}, {Week of Nightmares}, {Legacy of Caine}, {Kyoko Shinsegawa}.

[^3-1-10]: [[RTR 20080808]](https://usenet.krcg.org/t/1jW6GXrIRSU/#m0) — {Edged Illusion}, {Jaroslav Pascek}.

[^3-1-11]: [[ANK 20180623]](https://usenet.krcg.org/t/vekn-76748/#m4) — {Second Tradition:
    Domain}, {Black Sunrise}, {Sense the Savage Way}.

[^3-1-12]: [[LSJ 20031010]](https://usenet.krcg.org/t/clifT98_Zrk/#m6) [[LSJ
    20090319]](https://usenet.krcg.org/t/G1w6TeymEnQ/#m1) [[ANK
    20180907]](https://usenet.krcg.org/t/vekn-76983/#m7) [[LSJ
    20051113]](https://usenet.krcg.org/t/L-ctaLucuKU/#m1) [[TOM 19951208-1]](https://usenet.krcg.org/t/tEHebi9BYfc/#m5)
    [[LSJ 19980302-1]](https://usenet.krcg.org/t/kijV8VfB56s/#m3) [[ANK
    20220409]](https://usenet.krcg.org/t/vekn-79728/#m5)
    — {Secure Haven}, {Aye}, group "Reflex" (G00060), {Dam the Heart's River}.

[^3-1-13]: [[PIB 20121028]](https://usenet.krcg.org/t/vekn-39816/#m10)
    — {Baltimore Purge}.

[^3-1-14]: [[RBK who-may-attempt-to-block]](https://www.vekn.net/rulebook#who-may-attempt-to-block) [[LSJ
    20090324]](https://usenet.krcg.org/t/Zc_ogoVhsug/#m8) [[ANK
    20190606]](https://usenet.krcg.org/t/vekn-77692/#m2) [[PIB
    20121210]](https://usenet.krcg.org/t/vekn-42498/#m19)
    [[PIB 20121028]](https://usenet.krcg.org/t/vekn-39816/#m10) [[LSJ
    20010328]](https://usenet.krcg.org/t/VMAMljUO4NE/#m1) [[LSJ 20010828]](https://usenet.krcg.org/t/KoP_nqv-feM/#m1)
    [[ANK 20180917]](https://usenet.krcg.org/t/vekn-77011/#m3) — group
    "Directed at a card" (G00036), {Of Noble Blood}, {Renewed Vigor}, {Eurayle Gelasia Mylonas}, {Mylan Horseed},
    {Lunatic Eruption}, {Temptation}, {Impundulu}.

[^3-1-15]: [[LSJ 20090324]](https://usenet.krcg.org/t/Zc_ogoVhsug/#m8) [[RBK
    who-may-attempt-to-block]](https://www.vekn.net/rulebook#who-may-attempt-to-block) [[LSJ
    19970224]](https://usenet.krcg.org/t/80KRDjVFkyg/#m1) — {Conceal}, {Haunt}, {Goblinism}, group "Locquipments"
    (G00047), {Annazir}, {Arcanum Investigator}, {Gremlins}, {Loss}, {Principia Discordia}, {Ethan Locke}.

[^3-1-16]: [[LSJ 20061222]](https://usenet.krcg.org/t/WUKedBpp_h0/) [[LSJ
    20090324]](https://usenet.krcg.org/t/Zc_ogoVhsug/#m8) [[LSJ 20030214-2]](https://usenet.krcg.org/t/nU7PaMBymtY/#m5)
    — {Abbot}, {Detect Authority}, {Ambrosio Luis Monçada, Plenipotentiary}, {Evan Klein}.

[^3-1-20]: [[LSJ 20010925]](https://usenet.krcg.org/t/rZcOI3IZOJw/#m2) [[LSJ
    20100713]](https://usenet.krcg.org/t/sWt7HVuu4Hs/#m3) [[PIB
    20121230]](https://usenet.krcg.org/t/vekn-43471/#m3) —
    group "Action directed at all methuselah" (G00165).

[^3-1-17]: [[LSJ 20010627]](https://usenet.krcg.org/t/NhNCVCCDyU0/#m5) [[ANK
    20221019-3]](https://usenet.krcg.org/t/vekn-21158/#m6) — {Daemonic
    Possession}, {The Eternal Mask}, {Ghost-Eater}, {Khazar's Diary (Endless Night)}, {Pressing Flesh}, {Banishment}.

[^3-1-18]: [[PIB 20150821]](https://usenet.krcg.org/t/vekn-72609/#m7) [[RTR
    20080808]](https://usenet.krcg.org/t/1jW6GXrIRSU/#m0) [[RTR
    20081119]](https://www.vekn.net/card-lists/140-keepers-of-tradition) [[ANK
    20180917]](https://usenet.krcg.org/t/vekn-77011/#m3) [[LSJ
    20080809-1]](https://usenet.krcg.org/t/1jW6GXrIRSU/#m7) — group "Action directed at all methuselah" (G00165),
    {Condemn the Sins of the Father}, {Shepherd's Innocence}, {Wave of Insanity}.

[^3-1-19]: [[LSJ 20080809-2]](https://usenet.krcg.org/t/1jW6GXrIRSU/#m2) [[RTR
    20081119]](https://www.vekn.net/card-lists/140-keepers-of-tradition) [[RTR
    20080808]](https://usenet.krcg.org/t/1jW6GXrIRSU/#m0) [[ANK
    20180515]](https://usenet.krcg.org/t/vekn-76612/#m1) — {Chanjelin
    Ward}, {Talley, The Hound}, {Weigh the Heart}.

[^3-2-1]: [[TOM 19951109]](https://usenet.krcg.org/t/WhJj5K1Fa-0/#m10) [[LSJ
    20060409]](https://usenet.krcg.org/t/gsFQXsCGTG4/#m1) [[ANK
    20200329]](https://usenet.krcg.org/t/vekn-78546/#m1) [[LSJ
    20020814]](https://usenet.krcg.org/t/gt8wQhk76lA/#m1) [[ANK
    20221229]](https://usenet.krcg.org/t/vekn-80231/#m1) —
    {Bonding}, {Mask of a Thousand Faces}, {Familial Bond}, {Osric Vladislav}.

[^3-2-12]: [[ANK 20231031]](https://usenet.krcg.org/t/vekn-80911/#m4) — group "Stealth from random effect" (G00067),
    {Gathii}.

[^3-2-2]: [[RBK wording-templates]](https://www.vekn.net/rulebook#wording-templates) [[RBK
    who-may-attempt-to-block]](https://www.vekn.net/rulebook#who-may-attempt-to-block) [[LSJ
    20010111]](https://usenet.krcg.org/t/d3WSV1UXBV0/#m5) [[LSJ 20100329]](https://usenet.krcg.org/t/1SAloV6P7Xk/#m1)
    [[PIB
    20150216]](https://usenet.krcg.org/t/vekn-57153/#m7)
    — {Angelica, The Canonicus}, {Phased Motion Detector}, {Under Siege}.

[^3-2-3]: [[RTR 20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0) [[LSJ
    20021211]](https://usenet.krcg.org/t/-J07wvmidOA/#m17) [[LSJ 20020612]](https://usenet.krcg.org/t/sC0pTtJJDj4/#m2) —
    {Draba}, {Night Terrors}, groups "Minus stealth" (G00104) and "Minus Intercept" (G00102).

[^3-2-5]: [[RBK stealth-and-intercept]](https://www.vekn.net/rulebook#stealth-and-intercept) [[LSJ
    20100604-2]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m6) [[LSJ 20100611]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m11)
    — {Marciana Giovanni, Investigator}.

[^3-2-6]: [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) [[LSJ
    20010813-2]](https://usenet.krcg.org/t/zkKhvgZy9hA/#m2) — {Marciana Giovanni, Investigator}.

[^3-2-7]: [[RTR 20080808]](https://usenet.krcg.org/t/1jW6GXrIRSU/#m0) [[RTR
    20081119]](https://www.vekn.net/card-lists/140-keepers-of-tradition) [[LSJ
    20061222]](https://usenet.krcg.org/t/WUKedBpp_h0/) [[ANK
    20190416]](https://usenet.krcg.org/t/vekn-77560/#m3) — {Talley, The Hound},
    {Abbot}, {Ministry}, {Protection Racket}, {Teresita, The Godmother}.

[^3-2-11]: Owner's reading, no official ruling: [[rules-qa
    13.2]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-13--printed-vampire-abilities)
    [[rules-qa
    44.6]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-44--legacy-cards-for-tournament-decks-slice-16a)
    — {Carlton Van Wyk}.

[^3-2-9]: [[ANK
    20180926]](https://usenet.krcg.org/t/vekn-77030/#m5)
    [[ANK
    20210926]](https://usenet.krcg.org/t/vekn-79337/#m21)
    [[ANK 20211003]](https://usenet.krcg.org/t/vekn-79372/#m1)
    — {Guardian Vigil}, {Inspire Greatness}, {Deflection}, {My Enemy's Enemy}, {Redirection}.

[^3-2-10]: [[RBK wording-templates]](https://www.vekn.net/rulebook#wording-templates) [[ANK
    20230116]](https://usenet.krcg.org/t/vekn-80266/#m1)
    [[ANK
    20180307-2]](https://usenet.krcg.org/t/vekn-76451/#m1)
    [[LSJ 20011023-2]](https://usenet.krcg.org/t/47-PhTMiAOU/#m1) — {Phased Motion Detector}, {Matteus, Flesh Sculptor},
    {Starshell Grenade Launcher}.

[^3-3-1]: [[RTR 19950413]](https://usenet.krcg.org/t/zB2vyPBnO6g/#m9) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[LSJ 20020112]](https://usenet.krcg.org/t/9yhb-ToEzL8/#m4)
    [[LSJ 20041203]](https://usenet.krcg.org/t/_iogrCkrCmc/#m1) [[LSJ
    20050224-1]](https://usenet.krcg.org/t/enxwleC-ZKw/#m27) — {Eagle's Sight}, {Falcon's Eye}, {Anneke}, {Sonja Blue}.

[^3-3-2]: [[LSJ 20030227]](https://usenet.krcg.org/t/c9KLhUd-Isg/#m6) — {Eagle's Sight}.

[^3-3-24]: Owner's reading, no official ruling: [[rules-qa
    92.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-92--a-block-attempt-under-way-failed-by-a-bar-on-blockers-the-owners-playtest-sc-bot-game-d253f9c8dcf5bcb5-decision-459-readings-graded-claude-2026-10-24-owner-2026-10-24-921)
    — {Daring the Dawn}.

[^3-3-3]: [[LSJ 20010714-2]](https://usenet.krcg.org/t/9U-jt3p3R_Q/#m1) — {Guard Duty}, {Second Tradition: Domain},
    {Black Sunrise}, and the same wording on eight further cards.

[^3-3-4]: [[ANK 20180623]](https://usenet.krcg.org/t/vekn-76748/#m4) — {Guard Duty},
    {Trophy: Domain}, {Under Siege}.

[^3-3-5]: [[ANK 20181122-2]](https://usenet.krcg.org/t/vekn-77176/#m4) —
    {Second Tradition: Domain}, {Guardian Vigil}, {The Mole}.

[^3-3-7]: [[ANK 20200607]](https://usenet.krcg.org/t/vekn-78676/#m1) [[LSJ
    20100507]](https://usenet.krcg.org/t/vJQTPYtp-Eg/#m1) [[PIB
    20150820]](https://usenet.krcg.org/t/vekn-72462/#m1) [[ANK
    20210131-1]](https://usenet.krcg.org/t/vekn-79007/#m2) [[RBK
    sequencing]](https://www.vekn.net/rulebook#sequencing) — {Hide the Heart}, {Draba}, {Netwar}, {Folderol}.

[^3-3-8]: [[LSJ 20000507]](https://usenet.krcg.org/t/6-sMg4Wo7KE/#m3) [[LSJ
    20070203]](https://usenet.krcg.org/t/37EYK3FA30k/#m2) [[ANK
    20211003]](https://usenet.krcg.org/t/vekn-79372/#m1) —
    {Archon Investigation}, {Aksinya Daclau}, {Lost in Translation}, group "Change the target of a bleed" (G00106).

[^3-3-9]: [[TOM 19951129]](https://usenet.krcg.org/t/jjBzopH-yrQ/#m3) — {Wake with Evening's Freshness}, {Eyes of
    Argus}, {Forced Awakening}.

[^3-3-10]: [[LSJ 20081202]](https://usenet.krcg.org/t/GMxeDWiDXP8/#m3) [[LSJ
    19990106]](https://usenet.krcg.org/t/tJR6Hw6CmBQ/#m1) — {Faceless Night}, {Mask of a Thousand Faces}.

[^3-3-11]: [[LSJ 20031121-2]](https://usenet.krcg.org/t/1khXmKPU0ws/#m31) — {Tenebrous Form}.

[^3-3-14]: [[RTR 19950906]](https://usenet.krcg.org/t/P4uSsWk4jj8/#m0) [[ANK
    20200714]](https://usenet.krcg.org/t/vekn-78742/#m1) — {Brujah Frenzy}, {Ohoyo
    Hopoksia (Bastet)}.

[^3-3-15]: [[ANK
    20230226]](https://usenet.krcg.org/t/vekn-31040/#m6)
    [[LSJ 20031121-2]](https://usenet.krcg.org/t/1khXmKPU0ws/#m31) — {Unleash Hell's Fury}, group "Burn blood to attempt
    to block" (G00088).

[^3-3-16]: [[LSJ 19970707]](https://usenet.krcg.org/t/KWekwiRSa2I/#m1) [[TOM
    19951215-1]](https://usenet.krcg.org/t/mLnhwWDRglQ/#m1) [[RTR 19960530]](https://usenet.krcg.org/t/DpvF2Peet9o/#m0)
    — {Sabbat Priest}, {Camarilla Exemplary}, {Archon}, {Dónal O'Connor}.

[^3-3-26]: [[ANK 20260317]](https://usenet.krcg.org/t/vekn-82837/#m7) — {Parijat, the Dark Oracle}; supersedes [[ANK
    20260216]](https://usenet.krcg.org/t/vekn-82835/#m1)'s {Eagle's Sight} example in group "Cancel, cost is still paid"
    (G00155) (owner, 2026-11-11).

[^3-3-17]: [[LSJ 20050607-1]](https://usenet.krcg.org/t/vet0HGRxtXQ/#m1) [[LSJ
    20080818]](https://usenet.krcg.org/t/O5HUPRkKY-k/#m1) [[LSJ 20090810]](https://usenet.krcg.org/t/g3ukXdsh8xo/#m1)
    [[LSJ 20100305]](https://usenet.krcg.org/t/gyDkgm2qXMs/#m1) — {Blessing of Chaos}, {Valerius Maior, Hell's Fool},
    {No Secrets From the Magaji}, {Libertas}.

[^3-3-18]: [[LSJ 19980224]](https://usenet.krcg.org/t/xV22ImgKblY/#m0) [[RTR
    19991206]](https://usenet.krcg.org/t/N7iEmqgP9WU/#m0) [[ANK
    20180321]](https://usenet.krcg.org/t/vekn-76378/#m14) [[LSJ
    20060410]](https://usenet.krcg.org/t/jr8wSeSchsc/#m1) [[ANK
    20171017]](https://usenet.krcg.org/t/vekn-76233/#m4)
    [[ANK 20230305]](https://usenet.krcg.org/t/vekn-63821/#m40)
    [[RBK resolve-the-action]](https://www.vekn.net/rulebook#resolve-the-action) — {Change of Target}, {Clan Loyalty},
    {Angel of Berlin}, {Faceless Night}.

[^3-3-25]: Owner's reading, no official ruling: [[rules-qa
    53.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-53--sabbat-v5-the-path-of-cathari-implemented-readings-2026-10-07)
    — {Aelswith, The Irresistible}.

[^3-3-19]: [[RTR 19991206]](https://usenet.krcg.org/t/N7iEmqgP9WU/#m0) [[LSJ
    20091125]](https://usenet.krcg.org/t/1dYtrdRjRI8/#m1) [[RTR 20080808]](https://usenet.krcg.org/t/1jW6GXrIRSU/#m0)
    [[ANK
    20220116]](https://usenet.krcg.org/t/vekn-79600/#m1)
    — {Aching Beauty}, {Change of Target}, {Lesser Boon}, {Unleash Hell's Fury}.

[^3-3-20]: [[TOM 19951208-2]](https://usenet.krcg.org/t/UqLxnP1C_Wg/#m1) [[RTR
    20080808]](https://usenet.krcg.org/t/1jW6GXrIRSU/#m0) — {Spiritual Protector}, {Lesser Boon}.

[^3-3-23]: [[ANK
    20190116]](https://usenet.krcg.org/t/vekn-77302/#m1) [[RTR
    19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) — {Those Who Endure Judge}.

[^3-4-1]: [[LSJ 20070411]](https://usenet.krcg.org/t/umdINigMKqs/#m1) [[LSJ
    20090514]](https://usenet.krcg.org/t/PKQ6lQUBUOI/#m1) [[ANK
    20220218]](https://usenet.krcg.org/t/vekn-79001/#m13)
    [[ANK 20181107]](https://usenet.krcg.org/t/vekn-77152/#m1) [[LSJ
    20020725]](https://usenet.krcg.org/t/wCPFIH_g5ZE/#m4) [[ANK
    20210124]](https://usenet.krcg.org/t/vekn-79001/#m2) —
    {Ambush}, {Harass}, {Spirit Marionette}, {Warsaw Station}, {Temptation}, and the further cards carrying the same
    three templates.

[^3-4-16]: [[RTR 19960221]](https://usenet.krcg.org/t/UdU535eVm0Y/#m0) [[LSJ
    20020926]](https://usenet.krcg.org/t/C7vgb-0Mcbo/#m1) — {Bomb}, {Camera Phone}, {Chalice of Kinship}, {Guarded
    Rubrics}, {Jar of Skin Eaters}, {Karavalanisha Vrana}.

[^3-4-3]: [[LSJ 20090514]](https://usenet.krcg.org/t/PKQ6lQUBUOI/#m1) — {Coterie Tactics}, {Eternal Vigilance}, {Guard
    Duty}, {Ambush}.

[^3-4-4]: [[LSJ 20070411]](https://usenet.krcg.org/t/umdINigMKqs/#m1) [[LSJ
    20030519]](https://usenet.krcg.org/t/E6Jz8m3iKrA/#m3) [[RTR 19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0)
    [[LSJ 20011222]](https://usenet.krcg.org/t/DlCBJmB2fzY/#m7) [[ANK
    20221028-2]](https://usenet.krcg.org/t/vekn-80122/#m1) — {Forced March},
    {Freak Drive}, {Instantaneous Transformation}, {CrimethInc.}, {Perfectionist}, {Hrothulf}, {Tereza Rostas},
    {Sargon}.

[^3-4-5]: [[LSJ 20090411]](https://usenet.krcg.org/t/2aJf9bN7ZGc/#m6) — {Abactor}.

[^3-4-7]: [[RTR 20080808]](https://usenet.krcg.org/t/1jW6GXrIRSU/#m0) [[LSJ
    20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) — {Edged Illusion}, {Jaroslav Pascek}, {Enticement}.

[^3-4-8]: [[LSJ 20090323]](https://usenet.krcg.org/t/bG4PSmditm4/#m1) [[RBK
    wording-templates]](https://www.vekn.net/rulebook#wording-templates) — {Veil of Darkness}.

[^3-4-9]: [[ANK
    20240706]](https://usenet.krcg.org/t/vekn-81563/#m19) [[LSJ
    20080608]](https://usenet.krcg.org/t/j5ShUmUt-vM/#m1) [[PIB
    20150418]](https://usenet.krcg.org/t/vekn-70589/#m1) [[ANK
    20170226]](https://usenet.krcg.org/t/vekn-75625/#m3) — {The Platinum
    Protocol}, {Third Tradition: Progeny}, {Dual Form}.

[^3-4-10]: [[TOM 19951107]](https://usenet.krcg.org/t/lb3GhBTmgpM/#m56) [[PIB
    20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) — {Magic of
    the Smith}, {Vast Wealth}, {Tryphosa}.

[^3-4-11]: [[ANK 20220331]](https://usenet.krcg.org/t/vekn-79720/#m3)
    [[ANK 20210124]](https://usenet.krcg.org/t/vekn-79001/#m2)
    [[LSJ 20020927]](https://usenet.krcg.org/t/wg3PH7vOs1s/#m1) [[LSJ
    20021115]](https://usenet.krcg.org/t/35vszszGZOA/#m1) [[ANK
    20221011-2]](https://usenet.krcg.org/t/vekn-79984/#m2)
    [[LSJ 20100301]](https://usenet.krcg.org/t/TcUDpHnvW8M/#m6) [[ANK
    20220116]](https://usenet.krcg.org/t/vekn-79600/#m1)
    — {Daring the Dawn}, {Force of Will}, {Unleash Hell's Fury}.

[^3-4-12]: [[ANK 20200221]](https://usenet.krcg.org/t/vekn-78458/#m2)
    — {Strix}.

[^3-4-13]: [[LSJ 20110502]](https://usenet.krcg.org/t/bgg-648695/#m1) [[LSJ
    20081123]](https://usenet.krcg.org/t/_Pb29pBJ1kU/#m2) [[LSJ 20050422]](https://usenet.krcg.org/t/nJPfMOuTBtw/#m3)
    [[LSJ 20090205]](https://usenet.krcg.org/t/l5bUtHOejmc/#m2) — {Spying Mission}, {Andre LeRoux}, {Dis Pater}, {Major
    Boon}.

[^3-4-14]: [[LSJ 19980105]](https://usenet.krcg.org/t/PzVC-AeFuUQ/#m1) [[TOM
    19960303]](https://usenet.krcg.org/t/SEU6ztVpR94/#m2) [[LSJ 20061212]](https://usenet.krcg.org/t/h3onVZ1NqpQ/#m46) —
    group "Change the target of a bleed" (G00106), {My Enemy's Enemy}, {Redirection}, {Telepathic Misdirection},
    {Telepathic Counter}, {Spying Mission}, {Deflection}, {Archon Investigation}, {Lost in Translation}.

[^3-4-15]: [[PIB 20150612]](https://usenet.krcg.org/t/vekn-71660/#m4) — {Andre
    LeRoux}.

[^3-5-1]: [[RTR 20180719]](https://www.blackchantry.com/2018/07/18/rules-team-rulings-rtr-19-07-2018/) [[ANK
    20211015]](https://usenet.krcg.org/t/vekn-79389/#m6) [[LSJ
    20070709]](https://usenet.krcg.org/t/qKrJPLXBFFw/#m2) [[RTR 20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0)
    [[ANK
    20221011-2]](https://usenet.krcg.org/t/vekn-79984/#m2)
    [[LSJ 20011202-2]](https://usenet.krcg.org/t/qnVBhys6loo/#m4) [[LSJ
    20091224]](https://usenet.krcg.org/t/AV34Lb49JlE/#m3) [[LSJ 20070214-2]](https://usenet.krcg.org/t/C_QjWH6VIh4/#m1)
    — {Change of Target}, {Champion}, {Krassimir}, {Obedience}, {Red Herring}, {Secret Passage}, {Black Forest Base},
    and sixteen further cards carrying the identical "has reached resolution" template.

[^3-5-2]: [[LSJ 19980212]](https://usenet.krcg.org/t/fLFLlXZXHqA/#m0) [[RBK
    cancel-a-card]](https://www.vekn.net/rulebook#cancel-a-card) [[LSJ
    20060425]](https://usenet.krcg.org/t/U34QDob4Vco/#m5) [[ANK
    20220411]](https://usenet.krcg.org/t/vekn-79730/#m1) [[LSJ
    20050607]](https://usenet.krcg.org/t/WLv9R8wA0Ow/#m8) [[LSJ 20050608]](https://usenet.krcg.org/t/WLv9R8wA0Ow/#m10) —
    group "Cancel an action" (G00061), {React with Conviction}, {Dabbler}, {Spying Mission}.

[^3-5-3]: [[ANK
    20211015]](https://usenet.krcg.org/t/vekn-79389/#m6) [[ANK
    20210124]](https://usenet.krcg.org/t/vekn-79001/#m2) [[LSJ
    20090220]](https://usenet.krcg.org/t/JMgZa9jdrc4/#m1) [[LSJ 20070709]](https://usenet.krcg.org/t/qKrJPLXBFFw/#m2) —
    {The Kiss of Ra}, {Tangle Atropos' Hand}.

[^3-5-26]: [[ANK
    20211015]](https://usenet.krcg.org/t/vekn-79389/#m6) [[RBK
    cancel-a-card]](https://www.vekn.net/rulebook#cancel-a-card) — {The Kiss of Ra}; the card text of {Tangle Atropos'
    Hand}.

[^3-5-4]: [[ANK
    20210124]](https://usenet.krcg.org/t/vekn-79001/#m2) [[ANK
    20200813-2]](https://usenet.krcg.org/t/vekn-78799/#m1) [[LSJ
    20070526]](https://usenet.krcg.org/t/hm8eaW_dBD0/#m3) [[LSJ 20100505]](https://usenet.krcg.org/t/nAzGpZHw8eA/#m1) —
    {The Kiss of Ra}, {Tangle Atropos' Hand}, {Veil the Legions}, {Waiting Game}, {Channel 10}.

[^3-5-5]: [[LSJ 20090208]](https://usenet.krcg.org/t/07gUFmSeIxU/#m1) [[LSJ
    20070222-1]](https://usenet.krcg.org/t/jBrfK77gayo/#m4) [[ANK
    20221102]](https://usenet.krcg.org/t/vekn-80129/#m1) — {Enrage},
    {React with Conviction}, {Mobile HQ, Operation Antigen}.

[^3-5-6]: [[LSJ 20081213-1]](https://usenet.krcg.org/t/MNmJu12AU8I/#m1) [[ANK
    20210124]](https://usenet.krcg.org/t/vekn-79001/#m2) —
    {Enkil Cog}.

[^3-5-7]: [[LSJ 20060522]](https://usenet.krcg.org/t/2f0wF9CECu8/#m5) [[LSJ
    20060824]](https://usenet.krcg.org/t/zsZTYTbRVRI/#m1) [[LSJ 20080725]](https://usenet.krcg.org/t/jWWCwKnran0/#m1)
    [[LSJ 20090617]](https://usenet.krcg.org/t/OMTF0_ZqUL0/#m2) — {Change of Target}, {Obedience}, {Red Herring}.

[^3-5-8]: [[RTR 19950509]](https://usenet.krcg.org/t/_LKyR7pdMig/#m8) [[LSJ
    20080710]](https://usenet.krcg.org/t/f1NpGhdtk-E/#m1) [[ANK
    20200502]](https://usenet.krcg.org/t/vekn-78616/#m1) — {Change of
    Target}, {Obedience}, {Red Herring}.

[^3-5-9]: [[LSJ 20080725]](https://usenet.krcg.org/t/jWWCwKnran0/#m1) — {Change of Target}, {Obedience}, {Red Herring}.

[^3-5-10]: [[LSJ 20011023]](https://usenet.krcg.org/t/2GOLIrXAF8M/#m1) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) — {Malleable Visage}.

[^3-5-29]: Owner's reading, no official ruling: [[rules-qa
    58.13]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-58--legacy-cards-for-twda-decks-steps-49-implemented-readings-2026-10-09)
    — {Mask of a Thousand Faces}.

[^3-5-11]: [[LSJ 20010627]](https://usenet.krcg.org/t/NhNCVCCDyU0/#m5) [[LSJ
    20100213]](https://usenet.krcg.org/t/vXDkYrTmkws/#m2) — {Compel the Spirit}, {Soul Gem of Etrius}.

[^3-5-12]: [[LSJ 20080815]](https://usenet.krcg.org/t/0_1JBbdwi74/#m1) [[ANK
    20180817]](https://usenet.krcg.org/t/vekn-76933/#m2) [[ANK
    20220910]](https://usenet.krcg.org/t/vekn-80021/#m1) —
    {Piper}, {Muricia's Call}, {Clandestine Contract}.

[^3-5-13]: [[RTR 20180719]](https://www.blackchantry.com/2018/07/18/rules-team-rulings-rtr-19-07-2018/) [[ANK
    20200207]](https://usenet.krcg.org/t/vekn-78423/#m4) [[LSJ
    20010803-1]](https://usenet.krcg.org/t/5QuGIF5ERUI/#m1) [[ANK
    20210627]](https://usenet.krcg.org/t/vekn-79192/#m14) [[ANK
    20211207]](https://usenet.krcg.org/t/vekn-79528/#m7)
    [[ANK
    20220116]](https://usenet.krcg.org/t/vekn-79600/#m1)
    [[TOM 19960303]](https://usenet.krcg.org/t/SEU6ztVpR94/#m2) — group "Action ended by an effect" (G00166), {Tangle
    Atropos' Hand}, {Change of Target}, group "Triggered by block" (G00105), {Banshee Ironwail}, {Telepathic Counter}.

[^3-5-14]: [[LSJ 20010806-2]](https://usenet.krcg.org/t/5QuGIF5ERUI/#m3) [[TOM
    19950829]](https://usenet.krcg.org/t/iVm2VboVp6Q/#m7) — {Blood Brother Ambush}, {Brujah Frenzy}.

[^3-5-15]: [[LSJ 20041022]](https://usenet.krcg.org/t/gqhND6kd2wE/#m3) [[RTR
    20180719]](https://www.blackchantry.com/2018/07/18/rules-team-rulings-rtr-19-07-2018/) [[ANK
    20220118]](https://usenet.krcg.org/t/vekn-79389/#m11) [[ANK
    20221205]](https://usenet.krcg.org/t/vekn-80196/#m1) —
    {Detect Authority}, {Mistaken Identity}, {Scobax}, {Unseen Hibernation}, {Car Bomb}, {Champion}, {Hide the Heart},
    {Malkavian Derangement: Alternate Personality}.

[^3-5-16]: [[LSJ 20070219]](https://usenet.krcg.org/t/HkKuwBe9LRk/#m4) — {Champion}.

[^3-5-27]: [[ANK
    20260316]](https://usenet.krcg.org/t/vekn-82895/#m1) [[LSJ
    20041022]](https://usenet.krcg.org/t/gqhND6kd2wE/#m3) — group "Fail an action (after 'as announced')" (G00156).

[^3-5-17]: [[ANK 20200702]](https://usenet.krcg.org/t/vekn-78711/#m1) [[LSJ
    20090826]](https://usenet.krcg.org/t/KTwa1Hf_gHI/#m1) [[LSJ 20060902]](https://usenet.krcg.org/t/QAKz6Qtr7Ts/#m3)
    [[LSJ 20090325-1]](https://usenet.krcg.org/t/5CW9tD5OfGk/#m6) [[ANK
    20220218]](https://usenet.krcg.org/t/vekn-79001/#m13)
    [[LSJ 20090325-2]](https://usenet.krcg.org/t/5CW9tD5OfGk/#m8) — {Deep Song}, {Siren's Lure}, {Yawp Court}, {Warsaw
    Station}.

[^3-5-18]: [[ANK 20220331]](https://usenet.krcg.org/t/vekn-79720/#m3)
    [[ANK 20210124]](https://usenet.krcg.org/t/vekn-79001/#m2)
    [[LSJ 20020927]](https://usenet.krcg.org/t/wg3PH7vOs1s/#m1) [[LSJ
    20021115]](https://usenet.krcg.org/t/35vszszGZOA/#m1) [[ANK
    20221011-2]](https://usenet.krcg.org/t/vekn-79984/#m2)
    [[ANK 20220401]](https://usenet.krcg.org/t/vekn-79720/#m6)
    [[ANK 20211015]](https://usenet.krcg.org/t/vekn-79389/#m6)
    [[RTR 19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0) — {Daring the Dawn}, {Force of Will}.

[^3-5-20]: [[LSJ 20110419]](https://usenet.krcg.org/t/bgg-643948/#m4) — {Crocodile's Tongue} (ruling removed from the
    database in 2024; original verified at source).

[^3-5-28]: Owner's reading, no official ruling: [[rules-qa
    30.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-30--reactions-and-action-modifiers-wave-2)
    — {Dawn Operation}.

[^3-5-21]: [[LSJ 20090322]](https://usenet.krcg.org/t/tROpvfFdgBI/#m1) [[ANK
    20181003]](https://usenet.krcg.org/t/vekn-77036/#m5)
    [[LSJ 20010803]](https://usenet.krcg.org/t/s1lJEsLMf-8/#m7) [[ANK
    20220127]](https://usenet.krcg.org/t/vekn-79615/#m1)
    [[LSJ 20080611]](https://usenet.krcg.org/t/gvb3uijtpZw/#m1) [[ANK
    20210131]](https://usenet.krcg.org/t/vekn-79008/#m3) —
    {Bear-Baiting}, {Venenation}, {Crypt's Sons}.

[^3-5-24]: [[RBK cancel-a-card]](https://www.vekn.net/rulebook#cancel-a-card) [[LSJ
    20090818]](https://usenet.krcg.org/t/jkKBGVLmFHc/#m1) — {Supernatural Resistance}.

[^3-5-25]: [[ANK 20200813-2]](https://usenet.krcg.org/t/vekn-78799/#m1)
    — group "Cancel an action" (G00061), {React with Conviction}.

[^3-6-1]: [[LSJ 20070808-1]](https://usenet.krcg.org/t/jzMOxZ9oxOs/#m3) — {Form of Mist}, {Mirror Image}, {Shadow Body},
    {Ambulance}, {Foresee}, {Torrent}, and twelve further cards carrying the same wording.

[^3-6-2]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[LSJ
    20080611]](https://usenet.krcg.org/t/gvb3uijtpZw/#m1) [[ANK
    20190116]](https://usenet.krcg.org/t/vekn-77302/#m1) —
    {Form of Mist}, {Crypt's Sons}, {Momentary Delay}, and thirteen further cards carrying the same wording.

[^3-6-3]: [[ANK
    20190116]](https://usenet.krcg.org/t/vekn-77302/#m1) [[ANK
    20191108]](https://usenet.krcg.org/t/vekn-78081/#m2) [[RBK
    summary-of-the-course-of-an-action]](https://www.vekn.net/rulebook#summary-of-the-course-of-an-action) — {Form of
    Mist}, {Mirror Image}, {Wamukota} (and thirteen further cards); {Eyes of the Beast}.

[^3-6-4]: [[ANK
    20230319]](https://usenet.krcg.org/t/vekn-80393/#m2)
    [[LSJ 20070417]](https://usenet.krcg.org/t/ecDUqbSUsNg/#m1) — {Forced Awakening}, {WMRH Talk Radio}.

[^3-6-5]: [[LSJ 20030227]](https://usenet.krcg.org/t/c9KLhUd-Isg/#m6) [[LSJ
    20010814-2]](https://usenet.krcg.org/t/8MR4bq0Cxj4/#m3) [[ANK
    20181127]](https://usenet.krcg.org/t/vekn-77176/#m9) —
    {Eagle's Sight}, {Falcon's Eye}, {Guardian Vigil}.

[^3-6-6]: [[RTR 19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) — {Form of Mist}, {Chameleon's Colors}, {Torrent},
    and nine further cards carrying the same wording.

[^3-6-7]: [[LSJ 19980109]](https://usenet.krcg.org/t/d2U51Uw0Hfg/#m1) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[LSJ 20031123]](https://usenet.krcg.org/t/HvM2cHJ6yIo/#m3) —
    {Form of Mist}, {Fast Reaction}, {Akram}, {Chameleon's Colors}, {Mirror Image}, and seven further cards carrying the
    same wording.

[^3-6-8]: [[LSJ 20021213]](https://usenet.krcg.org/t/Puk_eXStfE8/#m7) [[ANK
    20250613]](https://usenet.krcg.org/t/vekn-82396/#m0) [[ANK
    20250613-2]](https://usenet.krcg.org/t/vekn-82396/#m4) — {Psyche!}, {Form of Mist}.

[^3-6-9]: [[LSJ 20090322]](https://usenet.krcg.org/t/tROpvfFdgBI/#m1) [[LSJ
    20010803]](https://usenet.krcg.org/t/s1lJEsLMf-8/#m7) [[ANK
    20181003]](https://usenet.krcg.org/t/vekn-77036/#m5)
    [[ANK
    20220127]](https://usenet.krcg.org/t/vekn-79615/#m1) —
    {Bear-Baiting}, {Venenation}.

[^3-6-10]: [[LSJ 20010813]](https://usenet.krcg.org/t/8MR4bq0Cxj4/#m1) [[LSJ
    20010819-2]](https://usenet.krcg.org/t/8MR4bq0Cxj4/#m14) [[LSJ
    20010814-2]](https://usenet.krcg.org/t/8MR4bq0Cxj4/#m3) — {Guard Dogs}, {Guardian Vigil}, {Precognition}, {Beast
    Meld}, {Night Terrors}, {One With the Land}, and seven further cards carrying the same wording.

[^3-6-11]: [[LSJ 20010814-2]](https://usenet.krcg.org/t/8MR4bq0Cxj4/#m3) — {Precognition}, {Instinctive Reaction}, {Form
    of the Bat}, and seven further cards carrying the same wording.

[^3-7-1-2]: [[PIB
    20110817]](https://usenet.krcg.org/t/vekn-8405/#m5)
    [[RTR 20191031]](https://www.vekn.net/2-uncategorised/465-vampire-elder-kindred-network-newsletter-october-2019)
    [[ANK
    20211019-2]](https://usenet.krcg.org/t/vekn-79390/#m9)
    — {Foreshadowing Destruction}.

[^3-7-1-4]: [[ANK 20171023]](https://usenet.krcg.org/t/vekn-76252/#m2) [[RTR
    19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[ANK
    20220502]](https://usenet.krcg.org/t/vekn-79757/#m1) —
    {Confusion}, {Power of One}, {Haqim's Law: Retribution}, {Anu Diptinatpa}.

[^3-7-1-5]: [[RTR 19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0) — {Justicar Retribution}.

[^3-7-1-6]: [[PIB 20150612]](https://usenet.krcg.org/t/vekn-71660/#m4) [[LSJ
    20030701-1]](https://usenet.krcg.org/t/8pVaGxWYeyA/#m5) — {Andre LeRoux}, {Protected Resources}.

[^3-7-1-7]: [[TOM 19960413]](https://usenet.krcg.org/t/Gm-NLCP6bF0/#m8) [[LSJ
    20030211]](https://usenet.krcg.org/t/xhpVMShX6qc/#m11) — {Protected Resources}.

[^3-7-1-8]: [[RTR 19960530]](https://usenet.krcg.org/t/DpvF2Peet9o/#m0) — {Elder Intervention}.

[^3-7-1-9]: [[RTR 19950622]](https://usenet.krcg.org/t/86Y38Vps-7E/#m11) [[PIB
    20130711]](https://usenet.krcg.org/t/vekn-51279/#m3)
    — {Deflection}, {Minor Boon}.

[^3-7-1-10]: [[LSJ 20010621]](https://usenet.krcg.org/t/VTHXJOrxlP4/#m1) [[LSJ
    20010830]](https://usenet.krcg.org/t/pWNZjWmtCk0/#m2) — {Perfect Clarity}.

[^3-7-2-1]: [[LSJ 20050720-2]](https://usenet.krcg.org/t/1vbIkgvKDJ0/#m1) [[LSJ
    20050727]](https://usenet.krcg.org/t/QfVIg4fJP38/#m1) — {Hunger Moon}, {Thirst}, {Tainted Vitae}, {Hospital Food},
    {The Anarch Free Press}, and eleven further cards carrying the same ruling.

[^3-7-2-11]: Owner's reading, no official ruling: [[rules-qa
    89.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-89--round-20-the-cards-players-asked-for-the-two-player-lists-last-two-readings-graded-claude-2026-10-23-owner-2026-10-23-891-corrected-892-confirmed)
    — {Harvest Rites}.

[^3-7-2-2]: [[TOM 19951212-3]](https://usenet.krcg.org/t/cYwUsviEhr4/#m2) — {Triole's Revenge}.

[^3-7-2-3]: [[LSJ 20060306]](https://usenet.krcg.org/t/H-W0Wqx3t9w/#m1) [[ANK
    20180202]](https://usenet.krcg.org/t/vekn-76402/#m4) [[RTR
    20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0) [[RTR
    20180511]](https://usenet.krcg.org/t/vekn-76595/#m0) — group "Hunt
    bonus" (G00004), {Pariah}, group "Special Hunt action" (G00064).

[^3-7-2-4]: [[PIB 20150820-2]](https://usenet.krcg.org/t/vekn-72575/#m4) [[PIB
    20120204]](https://usenet.krcg.org/t/vekn-22233/#m2) — {Igo the
    Hungry}, {Strained Vitae Supply}.

[^3-7-2-5]: [[RTR 20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0) [[RTR
    20180511]](https://usenet.krcg.org/t/vekn-76595/#m0) — groups "Hunt
    bonus" (G00004) and "Steal blood as a Hunt action" (G00121); {Festivo dello Estinto}, {Inbase Discotek, Frankfurt}
    (bank-blood bonuses printed on the cards).

[^3-7-2-6]: [[LSJ 20090606]](https://usenet.krcg.org/t/-hy9E1u2t_o/#m4) [[ANK
    20190606-2]](https://usenet.krcg.org/t/vekn-77694/#m1) — {Legacy
    of Caine}, {Week of Nightmares}, {Kyoko Shinsegawa}.

[^3-7-2-8]: [[LSJ 20010809-1]](https://usenet.krcg.org/t/LB6Zg4bEggc/#m5) — {Legacy of Caine}.

[^3-7-2-9]: [[ANK 20220809]](https://usenet.krcg.org/t/vekn-79949/#m3) — {Loki's
    Gift}.

[^3-7-2-10]: [[LSJ 20050331]](https://usenet.krcg.org/t/J0CgOqvbfEk/#m3) [[LSJ
    20081213-2]](https://usenet.krcg.org/t/0Yf2s-qrWpM/#m1) [[RBK hunt]](https://www.vekn.net/rulebook#hunt) — {Rabbat,
    The Sewer Goddess}, {Undying Thirst}.

[^3-7-3-1]: [[TOM 19960130]](https://usenet.krcg.org/t/wF82VdVPlm0/#m13) [[LSJ
    20050313]](https://usenet.krcg.org/t/eSShNquHnXI/#m1) [[ANK
    20230408]](https://usenet.krcg.org/t/vekn-80438/#m2) [[LSJ
    20050315]](https://usenet.krcg.org/t/COcJX2hHP-E/#m1) [[LSJ 20010210]](https://usenet.krcg.org/t/NtMa5w_NVOE/#m21) —
    groups "Equip/Employ/Recruit action" (G00132) and "Retainers and equipment put on different minion" (G00094),
    {Children of Stone}, {Sleight of Hand}, {Lambach}, {Jack of Both Sides}, {Beast, The Leatherface of Detroit},
    {Angelo}, {Lorrie Dunsirn}.

[^3-7-3-2]: [[PIB
    20150105-2]](https://usenet.krcg.org/t/vekn-68482/#m1)
    [[LSJ 20100421]](https://usenet.krcg.org/t/Vp--M79gpqk/#m1) — {Topaz}, {Synner-G}, {Vulture}, {Dagger}.

[^3-7-3-3]: [[LSJ 20090506]](https://usenet.krcg.org/t/887DQTpntKI/#m5) — {Incriminating Videotape}, {Mokolé Blood},
    {Shilmulo Tarot}.

[^3-7-3-4]: [[LSJ 20060209]](https://usenet.krcg.org/t/ZuOfZorIhhU/#m4) [[ANK
    20221103]](https://usenet.krcg.org/t/vekn-80131/#m2)
    [[ANK 20210109]](https://usenet.krcg.org/t/vekn-78983/#m3) — {Flaming
    Candle}, {Fear of Mekhet}.

[^3-7-3-12]: [[RBK equip]](https://www.vekn.net/rulebook#equip) [[ANK
    20221103]](https://usenet.krcg.org/t/vekn-80131/#m2)
    — {Flaming Candle}.

[^3-7-3-10]: Owner's reading, no official ruling: [[rules-qa
    9.2]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-9--equipment) — {Flaming
    Candle}.

[^3-7-3-9]: [[ANK
    20180719-3]](https://usenet.krcg.org/t/vekn-76834/#m1) — group
    "Clan equipment" (G00013).

[^3-7-3-11]: Owner's reading, no official ruling: [[rules-qa
    9.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-9--equipment) — {Treasured
    Samadji}.

[^3-7-3-6]: [[LSJ 20060222]](https://usenet.krcg.org/t/Gv-gf5sAJxM/#m2) [[ANK
    20170326]](https://usenet.krcg.org/t/vekn-75635/#m5) — {Unlicensed Taxicab}, {Gift of Bellona}.

[^3-7-3-8]: [[LSJ 20050315]](https://usenet.krcg.org/t/COcJX2hHP-E/#m1) [[LSJ
    20010210]](https://usenet.krcg.org/t/NtMa5w_NVOE/#m21) [[LSJ 19980206]](https://usenet.krcg.org/t/p_uyqQgE9Ms/#m1)
    [[TOM 19950407]](https://usenet.krcg.org/t/FWVnIu3zLAQ/#m5) — {Beast, The Leatherface of Detroit}, {Enkidu, The
    Noah}, {Howler}, {Lucian}, {Vast Wealth}.

[^3-7-4-1]: [[LSJ 20100303]](https://usenet.krcg.org/t/jmmm0WRUPvs/#m4) [[ANK
    20180817]](https://usenet.krcg.org/t/vekn-76933/#m2) [[LSJ
    20100725]](https://usenet.krcg.org/t/9d1zMZfsfNo/) [[LSJ 20030520-2]](https://usenet.krcg.org/t/GcymCHOJDVY/#m6)
    [[RBK employ-retainer]](https://www.vekn.net/rulebook#employ-retainer) [[RBK
    recruit-ally]](https://www.vekn.net/rulebook#recruit-ally) — groups "Equip/employ/recruit outside of an action"
    (G00131) and "Equip/Employ/Recruit action" (G00132), {Ghouled}.

[^3-7-4-9]: Owner's reading, no official ruling: [[rules-qa
    62.21]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-62--legacy-cards-for-twda-decks-round-2-steps-16-implemented-readings-2026-10-11-owner-2026-10-20-all-confirmed)
    [[rules-qa
    41.12]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-41--vampires-printed-text-slice-12a)
    — {CrimethInc.}, {Keegan}, {"Mother" Anja Giovanni}.

[^3-7-4-3]: [[LSJ 20100204]](https://usenet.krcg.org/t/o5Xnzc8G774/#m31) [[LSJ
    20080803]](https://usenet.krcg.org/t/VgARso4nY7w/#m1) [[RBK
    employ-retainer]](https://www.vekn.net/rulebook#employ-retainer) [[RBK
    recruit-ally]](https://www.vekn.net/rulebook#recruit-ally) — group "Put card in play ignoring requirements"
    (G00110), {Piper}.

[^3-7-4-4]: [[LSJ 20090115-1]](https://usenet.krcg.org/t/RQ3ARP9Kvfk/#m1) [[ANK
    20170309]](https://usenet.krcg.org/t/vekn-75649/#m2) — group
    "Equip/employ/recruit outside of an action" (G00131), {Zhenga}.

[^3-7-4-5]: [[ANK 20210928]](https://usenet.krcg.org/t/vekn-79364/#m2) [[ANK
    20210913]](https://usenet.krcg.org/t/vekn-79322/#m1) [[LSJ
    20090116]](https://usenet.krcg.org/t/RQ3ARP9Kvfk/#m7) [[PIB
    20150105-2]](https://usenet.krcg.org/t/vekn-68482/#m1)
    — {Soul of the Earth}, {Little Tailor of Prague}, {Kuyén}, {Baba Yaga, the Iron Hag}.

[^3-7-4-8]: [[LSJ 20070707]](https://usenet.krcg.org/t/ZtRk5z2TcoI/#m1) — {Corrupt Construction}.

[^3-7-5-1]: [[LSJ 20100112]](https://usenet.krcg.org/t/SJu0kgw_2tE/#m1) [[PIB
    20121028]](https://usenet.krcg.org/t/vekn-39816/#m10) [[LSJ
    20011222]](https://usenet.krcg.org/t/DlCBJmB2fzY/#m7) [[ANK
    20221028-2]](https://usenet.krcg.org/t/vekn-80122/#m1) [[RBK
    politics]](https://www.vekn.net/rulebook#politics) — {Abactor}, {Sargon}.

[^3-7-5-2]: [[RTR 19951110]](https://usenet.krcg.org/t/TXfganI5B2o/#m0) — {Voter Captivation}.

[^3-7-5-3]: [[LSJ 20060902]](https://usenet.krcg.org/t/QAKz6Qtr7Ts/#m3) [[LSJ
    20090325-1]](https://usenet.krcg.org/t/5CW9tD5OfGk/#m6) [[ANK
    20180910-1]](https://usenet.krcg.org/t/vekn-76990/#m3)
    [[LSJ 20060903]](https://usenet.krcg.org/t/QAKz6Qtr7Ts/#m8) [[LSJ
    20080201]](https://usenet.krcg.org/t/y5Uoc7nEulU/#m1) [[ANK
    20220218]](https://usenet.krcg.org/t/vekn-79001/#m13)
    — {Yawp Court}, {Warsaw Station}.

[^3-7-5-4]: [[LSJ 20081202-2]](https://usenet.krcg.org/t/hlK89M1M9rk/#m6) [[LSJ
    19980130]](https://usenet.krcg.org/t/mtjOAd7aaYI/#m3) [[RTR 19951110]](https://usenet.krcg.org/t/TXfganI5B2o/#m0) —
    {Veles' Hunt}, {Bernard, the Scourge}, {Delaying Tactics}, {Telepathic Vote Counting}, {Scorn of Adonis}.

[^3-7-5-5]: [[RTR 20040501]](https://usenet.krcg.org/t/7-mp3Ada86I/#m0) [[RBK
    politics]](https://www.vekn.net/rulebook#politics) — {Business Pressure}.

[^3-7-5-6]: [[LSJ 20081203]](https://usenet.krcg.org/t/gaskAJqA-mE/#m1) [[LSJ
    20041004]](https://usenet.krcg.org/t/UZbyxuVsTJE/#m10) [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) —
    {Revolutionary Council}, {Parity Shift}, {Domain Challenge}.

[^3-7-5-7]: [[LSJ 20010606]](https://usenet.krcg.org/t/2Dj_N6wtifI/#m2) [[LSJ
    20040518]](https://usenet.krcg.org/t/4emymfUPwAM/#m5) — {Parity Shift}, {Alastor}.

[^3-7-5-20]: [[LSJ 20100129]](https://usenet.krcg.org/t/X8Uu7Sk56P4/#m8) [[ANK
    20191228]](https://usenet.krcg.org/t/vekn-78262/#m1) — group "Political
    action with illegal terms" (G00158).

[^3-7-5-21]: Owner's reading, no official ruling: [[rules-qa
    43.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-43--found-by-training-p4) —
    {Parity Shift}, {Cedrick Calhoun}.

[^3-7-5-8]: [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) — {Peace Treaty}.

[^3-7-5-22]: [[RBK politics]](https://www.vekn.net/rulebook#politics) [[RBK
    sequencing]](https://www.vekn.net/rulebook#sequencing) — {Closed Session}.

[^3-7-5-9]: [[RTR 19960530]](https://usenet.krcg.org/t/DpvF2Peet9o/#m0) [[LSJ
    20030602]](https://usenet.krcg.org/t/dy31XE695_M/#m6) — {Business Pressure}, {Mob Rule}.

[^3-7-5-10]: [[LSJ 20041207]](https://usenet.krcg.org/t/6oC7FtG2Ac4/#m3) [[ANK
    20180307-2]](https://usenet.krcg.org/t/vekn-76451/#m1)
    [[RBK wording-templates]](https://www.vekn.net/rulebook#wording-templates) — {Michael Luther}, {Ellison Humboldt}.

[^3-7-5-11]: [[RTR 20001020]](https://usenet.krcg.org/t/GvxNYsYsWJ4/#m0) — {Astrid Thomas}.

[^3-7-5-12]: [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[LSJ
    20090115-2]](https://usenet.krcg.org/t/rSaiVPMbpvY/#m0) — {Treachery}, {Guru}.

[^3-7-5-13]: [[LSJ 20090115-2]](https://usenet.krcg.org/t/rSaiVPMbpvY/#m0) [[ANK
    20220808]](https://usenet.krcg.org/t/vekn-79952/#m1) — {Donald
    Cargill}, {Lutz von Hohenzollern}, {Armin Brenner}.

[^3-7-5-14]: [[LSJ 19980107]](https://usenet.krcg.org/t/aUUs4VCR_Ec/#m1) [[PIB
    20150105]](https://usenet.krcg.org/t/vekn-68465/#m5) — group "Usable during the
    polling step" (G00022), {Charming Lobby}, {Cryptic Rider}, {Distant Friend}, {Quicksilver Contemplation}.

[^3-7-5-15]: [[PIB 20150105]](https://usenet.krcg.org/t/vekn-68465/#m5) [[LSJ
    19980107]](https://usenet.krcg.org/t/aUUs4VCR_Ec/#m1) — {Voter Captivation}, {Cryptic Rider}.

[^3-7-5-16]: [[LSJ 20010209]](https://usenet.krcg.org/t/wYp61ffInqs/#m1) [[ANK
    20201029-2]](https://usenet.krcg.org/t/vekn-78890/#m1) [[LSJ
    20030426]](https://usenet.krcg.org/t/JmGLxQmAF6s/#m1) — {Delaying Tactics}, {Telepathic Vote Counting}.

[^3-7-5-17]: [[LSJ 20040730]](https://usenet.krcg.org/t/vCZw1_QnhfE/#m2) [[ANK
    20220805]](https://usenet.krcg.org/t/vekn-79939/#m3)
    — {Aura of Invincibility}.

[^3-7-5-18]: [[LSJ 20070927]](https://usenet.krcg.org/t/VaSQ7JL2N2Y/#m1) [[LSJ
    20041130]](https://usenet.krcg.org/t/6uTPqRg387A/#m3) [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1)
    [[LSJ 20020911]](https://usenet.krcg.org/t/_lLSO5aevoM/#m1) — {Luna Giovanni}, {Delaying Tactics}, {Echo of
    Harmonies}.

[^3-7-5-19]: [[ANK
    20190606-2]](https://usenet.krcg.org/t/vekn-77694/#m1) [[TOM
    19950921]](https://usenet.krcg.org/t/--zrAV3UcGI/#m2) [[LSJ 20081203-2]](https://usenet.krcg.org/t/hlK89M1M9rk/#m16)
    — {Power Structure}, {Charming Lobby}, {Gangrel Conspiracy}.

[^3-7-6-1]: [[PIB 20110802]](https://usenet.krcg.org/t/vekn-7238/#m4) [[LSJ
    20041202]](https://usenet.krcg.org/t/8MwcrUpCdBc/#m16) [[ANK
    20180307-1]](https://usenet.krcg.org/t/vekn-76452/#m1) [[RTR
    19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[ANK
    20211009]](https://usenet.krcg.org/t/vekn-79388/#m1) [[LSJ
    20041025]](https://usenet.krcg.org/t/sbfkGmojYao/#m4) — {Condemnation: Mute}, {Rastacourere}, {Island of Yiaros},
    {Leadership Vacuum}, {Political Struggle}, group "Reduce votes" (G00021).

[^3-7-6-2]: [[LSJ 19971001]](https://usenet.krcg.org/t/_Yu65rf2qE4/#m0) [[LSJ
    20100311]](https://usenet.krcg.org/t/Ogc7-acbMFs/#m1) — {Arishat}, {Kateline Nadasdy}, {Sundown}.

[^3-7-6-3]: [[LSJ 20040519]](https://usenet.krcg.org/t/fz-EAPmmqZY/#m10) [[RBK
    gaining-votes]](https://www.vekn.net/rulebook#gaining-votes) [[LSJ
    20051113-2]](https://usenet.krcg.org/t/O2cgcyCHBSI/#m3) — {Gratiano}, {Genevieve}.

[^3-7-6-4]: [[ANK
    20211009]](https://usenet.krcg.org/t/vekn-79388/#m1) [[LSJ
    20040518-2]](https://usenet.krcg.org/t/4emymfUPwAM/#m1) — {Leadership Vacuum}, {Political Struggle}.

[^3-7-6-11]: [[ANK 20251125]](https://usenet.krcg.org/t/vekn-82686/#m1) [[ANK
    20251220]](https://usenet.krcg.org/t/vekn-82686/#m13) [[LSJ
    20050221-2]](https://usenet.krcg.org/t/XJqubVDYzxQ/#m2) — group "Vote reduction" (G00146).

[^3-7-6-5]: [[LSJ 20090304-2]](https://usenet.krcg.org/t/PcbRGxbYQUY/#m3) [[ANK
    20221019]](https://usenet.krcg.org/t/vekn-80100/#m1) —
    {Mustafa, The Heir}, {De Sade}.

[^3-7-6-12]: Owner's reading, no official ruling: [[rules-qa
    5.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-5--politics) [[rules-qa
    19.3]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-19--titles-and-votes) —
    {Bewitching Oration}, {Elysium: The Palace of Versailles}.

[^3-7-6-6]: [[LSJ 20020123]](https://usenet.krcg.org/t/F2PELnDgM_g/#m5) [[RTR
    19951110]](https://usenet.krcg.org/t/TXfganI5B2o/#m0) [[LSJ 20031115]](https://usenet.krcg.org/t/quF7butlINo/#m5)
    [[PIB
    20141026-2]](https://usenet.krcg.org/t/vekn-66453/#m20)
    — {Madrigal}, {Scorn of Adonis}, {Loyalist}, {Khay'tall, Snake of Eden}.

[^3-7-6-7]: [[RTR 19991001]](https://usenet.krcg.org/t/RAvWWmYoX3U/#m0) [[LSJ
    20030828]](https://usenet.krcg.org/t/enVCjRhydLo/#m1) [[ANK
    20190731]](https://usenet.krcg.org/t/vekn-77812/#m11) [[RTR
    19941006]](https://usenet.krcg.org/t/HUipqz0LFSw/#m0) [[RTR 19960530]](https://usenet.krcg.org/t/DpvF2Peet9o/#m0) —
    {Kindred Coercion}, {Kindred Manipulation}, {Neferu}, {Astrid Thomas}.

[^3-7-6-9]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[RTR
    19970425]](https://usenet.krcg.org/t/DhP_l2cX3mQ/#m0) [[ANK
    20220705]](https://usenet.krcg.org/t/vekn-79895/#m1)
    [[LSJ 20100705]](https://usenet.krcg.org/t/Dm_Zqyjdx8s/) [[ANK
    20180307-2]](https://usenet.krcg.org/t/vekn-76451/#m1)
    [[LSJ 20010326]](https://usenet.krcg.org/t/0Sy3xNbjYeU/#m2) [[RBK
    wording-templates]](https://www.vekn.net/rulebook#wording-templates) — {Disarming Presence}, {Alvaro, The Scion of
    Angelica}, group "Unconditional referendum ability" (G00039), group "Non-locking referendum ability" (G00040).

[^3-7-6-10]: [[ANK
    20220704]](https://usenet.krcg.org/t/vekn-79890/#m1)
    [[LSJ 20100426]](https://usenet.krcg.org/t/BN3xmoZ0W1A/#m2) [[LSJ
    20091128]](https://usenet.krcg.org/t/-IxzB0bvhKU/#m1) [[ANK
    20210309-2]](https://usenet.krcg.org/t/vekn-79005/#m7) — {Charming
    Lobby}, {Echo of Harmonies}, group "Vote playable once per game" (G00030).

[^3-7-7-1]: [[RBK leave-torpor]](https://www.vekn.net/rulebook#leave-torpor) [[RBK
    rescue-a-vampire-from-torpor]](https://www.vekn.net/rulebook#rescue-a-vampire-from-torpor) — rulebook action
    templates.

[^3-7-7-2]: [[ANK
    20181017]](https://usenet.krcg.org/t/vekn-77086/#m3) [[LSJ
    19980126]](https://usenet.krcg.org/t/C96l_KOk174/#m0) — {Resume the Coil}, {Rapid Healing}, {Healing Touch}, {Recure
    of the Homeland}, {Root of Vitality}, {Sense Vitality}, {Warding the Beast}, {Lord of Serenity}.

[^3-7-7-3]: [[ANK
    20181122-1]](https://usenet.krcg.org/t/vekn-77171/#m13)
    [[LSJ 20050105-2]](https://usenet.krcg.org/t/sAErRDiSXfU/#m5) [[PIB
    20121216]](https://usenet.krcg.org/t/vekn-42974/#m2) [[RTR
    20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) — {Ghoul Escort}, {Change of Target}, {Mirror Walk}, {Blood
    Brother Ambush}.

[^3-7-7-4]: [[LSJ 20020304-2]](https://usenet.krcg.org/t/L-8OGYP5xsE/#m6) [[PIB
    20110918]](https://usenet.krcg.org/t/vekn-10458/#m1) [[LSJ
    20100315]](https://usenet.krcg.org/t/06C5ufFEaJs/#m2) — {Frondator}, {Miriam Benyona}, {Cavalier}.

[^3-7-7-7]: Owner's reading, no official ruling: [[rules-qa
    48.8]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-48--found-by-the-owners-play-2026-10-03)
    — group "Discounts for rescue" (G00075).

[^3-7-7-5]: [[RTR 19950509]](https://usenet.krcg.org/t/_LKyR7pdMig/#m8) — {Madness Network}.

[^3-7-7-6]: [[ANK
    20220218]](https://usenet.krcg.org/t/vekn-79001/#m13)
    [[ANK 20181107]](https://usenet.krcg.org/t/vekn-77152/#m1) [[LSJ
    20090325-2]](https://usenet.krcg.org/t/5CW9tD5OfGk/#m8) — {Warsaw Station}.

[^3-7-8-1]: [[RTR 19991206]](https://usenet.krcg.org/t/N7iEmqgP9WU/#m0) [[ANK
    20201228]](https://usenet.krcg.org/t/vekn-78956/#m1) — {Draught of the Soul}, {Soul Stealing}, {Ritual of the Bitter
    Rose}, {Slake the Thirst}.

[^3-7-8-8]: Owner's reading, no official ruling: [[rules-qa
    88.4]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-88--open-contradictions-that-change-play-rules-digest-23-k5k7-k15-k16-k18-k23-k26-k27-k29-readings-graded-claude-2026-10-22-owner-2026-10-22-881-888-889-confirmed-884-answered-58-kept)
    — {Freak Drive}.

[^3-7-8-2]: [[LSJ 20090722-2]](https://usenet.krcg.org/t/Ry0xU4IuJmQ/#m1) [[LSJ
    20030618]](https://usenet.krcg.org/t/AdfUNNicx-Y/#m16) [[LSJ 20050707]](https://usenet.krcg.org/t/S5UbigI9faM/#m1)
    [[LSJ 20050222]](https://usenet.krcg.org/t/pwPVmNg8hDY/#m8) — {Heidelberg Castle, Germany}, {Political Struggle},
    {Trophy: Diablerie}.

[^3-7-8-3]: [[TOM 19950921]](https://usenet.krcg.org/t/--zrAV3UcGI/#m2) [[ANK
    20190606-2]](https://usenet.krcg.org/t/vekn-77694/#m1) [[LSJ
    20081203-2]](https://usenet.krcg.org/t/hlK89M1M9rk/#m16) — {Charming Lobby}, {Power Structure}, {Gangrel
    Conspiracy}.

[^3-7-8-4]: [[LSJ 20070417]](https://usenet.krcg.org/t/ecDUqbSUsNg/#m1) [[LSJ
    20010814-4]](https://usenet.krcg.org/t/8C05aiy4bdg/#m16) [[LSJ 20091210]](https://usenet.krcg.org/t/6gInH1jtvkc/#m2)
    — {Abandoning the Flesh}, {Ashes to Ashes}, {Reform Body}, {Hector Trelane}.

[^3-7-8-5]: [[LSJ 20090724]](https://usenet.krcg.org/t/1vJy-UKWR7Y/#m2) [[ANK
    20180129-1]](https://usenet.krcg.org/t/vekn-76390/#m2) [[ANK
    20190701]](https://usenet.krcg.org/t/vekn-77763/#m2) [[LSJ
    20091026]](https://usenet.krcg.org/t/rPfGPwFH0_E/#m1) [[LSJ 20100112]](https://usenet.krcg.org/t/SJu0kgw_2tE/#m1) —
    {Abactor}, {Rebirth}.

[^3-7-8-6]: [[LSJ 19970224]](https://usenet.krcg.org/t/80KRDjVFkyg/#m1) [[LSJ
    20011214-4]](https://usenet.krcg.org/t/RSCjaaZXY28/#m1) [[LSJ 20100228]](https://usenet.krcg.org/t/yDeGOj1RBuU/#m1)
    [[ANK 20210424]](https://usenet.krcg.org/t/vekn-79121/#m2) [[ANK
    20220218]](https://usenet.krcg.org/t/vekn-79001/#m13)
    [[ANK 20181107]](https://usenet.krcg.org/t/vekn-77152/#m1) — {Reform
    Body}, {Ashes to Ashes}, {Byzar}, {Warsaw Station}.

[^3-7-8-7]: [[LSJ 20080717]](https://usenet.krcg.org/t/DMsE6V84GWI/#m1) [[RTR
    19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[LSJ 20050324-2]](https://usenet.krcg.org/t/aIzKlsTs51k/#m1)
    [[LSJ 20050514]](https://usenet.krcg.org/t/9803Eu3PvDs/#m4) [[ANK
    20211010]](https://usenet.krcg.org/t/vekn-79335/#m8) [[LSJ
    20081213-2]](https://usenet.krcg.org/t/0Yf2s-qrWpM/#m1) — {Shadow Court Satyr}, {Undying Thirst}, {Phillipe Rigaud}.

[^3-8-1]: [[RTR 20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0) [[LSJ
    20020201]](https://usenet.krcg.org/t/FZ_S1BukETE/#m3) [[LSJ 20060522]](https://usenet.krcg.org/t/2f0wF9CECu8/#m5)
    [[LSJ 20090617]](https://usenet.krcg.org/t/OMTF0_ZqUL0/#m2) [[ANK
    20221102-2]](https://usenet.krcg.org/t/vekn-80130/#m1) — {Archon}, {Templar},
    {Delaying Tactics}, {Red Herring}, {Haqim's Law: Retribution}.

[^3-8-2]: [[LSJ 20100909]](https://usenet.krcg.org/t/9Mn1QueD1I4/#m5) [[ANK
    20220615]](https://usenet.krcg.org/t/vekn-72394/#m7) [[ANK
    20200810]](https://usenet.krcg.org/t/vekn-78797/#m2) [[ANK
    20220910]](https://usenet.krcg.org/t/vekn-80021/#m1) [[RBK
    action-card-or-card-in-play]](https://www.vekn.net/rulebook#action-card-or-card-in-play) — group "Provide multiple
    actions" (G00035), {Annazir}, {Bindusara, Historian of the Kindred}, {Clandestine Contract}.

[^3-8-3]: [[ANK 20180202]](https://usenet.krcg.org/t/vekn-76402/#m4) [[RTR
    19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) — {Haqim's Law: Judgment}, {Pariah}, {Delaying Tactics}.

[^3-8-4]: [[ANK 20180109]](https://usenet.krcg.org/t/vekn-76360/#m4) [[LSJ
    20070224]](https://usenet.krcg.org/t/lJBIPgsqNDg/#m1) [[ANK
    20220704]](https://usenet.krcg.org/t/vekn-79890/#m1)
    [[LSJ 20070927]](https://usenet.krcg.org/t/VaSQ7JL2N2Y/#m1) [[LSJ
    20050407]](https://usenet.krcg.org/t/fDl3t2lJ3Pc/#m1) [[ANK
    20181017]](https://usenet.krcg.org/t/vekn-77086/#m3) — {Ravnos
    Carnival}, {Charming Lobby}, {Go Anarch}.

[^3-9-1]: [[LSJ 20001127]](https://usenet.krcg.org/t/XgeLMemLbj0/#m1) [[LSJ
    20011216-2]](https://usenet.krcg.org/t/eYJFjfISdPY/#m7) [[RBK
    minion-phase]](https://www.vekn.net/rulebook#minion-phase) — {Lunatic Eruption}, {Spirit Marionette}.

[^3-9-2]: [[LSJ 20050514]](https://usenet.krcg.org/t/9803Eu3PvDs/#m4) [[ANK
    20211010]](https://usenet.krcg.org/t/vekn-79335/#m8) — {Phillipe
    Rigaud}.

[^3-9-3]: [[ANK 20211009-2]](https://usenet.krcg.org/t/vekn-79335/#m6)
    [[LSJ 20021121-2]](https://usenet.krcg.org/t/Nc72KRVbd-g/#m1) [[LSJ
    20081213-2]](https://usenet.krcg.org/t/0Yf2s-qrWpM/#m1) — {Elen Kamjian}, {Spirit Marionette}, {Undying Thirst}.

[^3-9-4]: [[LSJ 20090226]](https://usenet.krcg.org/t/f7pLAO9n--U/) [[ANK
    20211009-2]](https://usenet.krcg.org/t/vekn-79335/#m6) [[ANK
    20200227]](https://usenet.krcg.org/t/vekn-78474/#m1) — {Cry Wolf}, {Elen
    Kamjian}, {Lunatic Eruption}.

[^3-9-5]: [[ANK 20211009-2]](https://usenet.krcg.org/t/vekn-79335/#m6)
    [[LSJ 19980112]](https://usenet.krcg.org/t/C7WoO_yDhN0/#m1) [[LSJ
    20041103]](https://usenet.krcg.org/t/MiPHVp-NmCA/#m5) — {Elen Kamjian}, {Change of Target}, {Lunatic Eruption}.

[^3-9-6]: [[LSJ 20001113]](https://usenet.krcg.org/t/m2grSgEaYMw/#m1) — {Mask of a Thousand Faces}.

[^3-10-1]: [[RTR 19951110]](https://usenet.krcg.org/t/TXfganI5B2o/#m0) [[LSJ
    19971201]](https://usenet.krcg.org/t/0I7KUhvhAig/#m1) [[RTR 20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0)
    [[LSJ 20030521]](https://usenet.krcg.org/t/Ude2or4n9nI/#m2) [[ANK
    20210309]](https://usenet.krcg.org/t/vekn-79063/#m5)
    — {Mask of a Thousand Faces}.

[^3-10-2]: [[RTR 19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0) [[RTR
    20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0) [[LSJ 20030520]](https://usenet.krcg.org/t/wltIoG3qv_I/#m3)
    [[RTR 20041202]](https://usenet.krcg.org/t/WUWh7AdooDU/#m5) [[LSJ
    20001113]](https://usenet.krcg.org/t/m2grSgEaYMw/#m1) [[ANK
    20190117-1]](https://usenet.krcg.org/t/vekn-77308/#m3) —
    {Mask of a Thousand Faces}; [[LSJ 20020927]](https://usenet.krcg.org/t/wg3PH7vOs1s/#m1) — {Force of Will}; [[ANK
    20221102-2]](https://usenet.krcg.org/t/vekn-80130/#m1) — {Haqim's Law:
    Retribution}.

[^3-10-5]: [[LSJ 20020926]](https://usenet.krcg.org/t/C7vgb-0Mcbo/#m1) — {Mask of a Thousand Faces}.

[^3-10-3]: [[RBK traits]](https://www.vekn.net/rulebook#traits) [[LSJ
    20011023]](https://usenet.krcg.org/t/2GOLIrXAF8M/#m1) [[RTR 20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) —
    {Malleable Visage}; [[ANK
    20171212]](https://usenet.krcg.org/t/vekn-76334/#m14) —
    {Obedience}; [[ANK 20200114]](https://usenet.krcg.org/t/vekn-78321/#m3) —
    {Shadow Boxing}; [[ANK 20180913-2]](https://usenet.krcg.org/t/vekn-76998/#m8) —
    {FBI Special Affairs Division}; [[ANK
    20230814]](https://usenet.krcg.org/t/vekn-80752/#m12) —
    {Powerbase: Savannah}.

[^3-10-6]: Owner's reading, no official ruling: [[rules-qa
    27.5]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-27--actions-wave-2) — {Deep
    Song}.

[^3-10-4]: [[ANK 20211022]](https://usenet.krcg.org/t/vekn-79422/#m2) — {Deep Song},
    {Nar-Sheptha}.

[^4-1-3]: [[ANK 20221102]](https://usenet.krcg.org/t/vekn-80129/#m1)
    [[LSJ 20011214-3]](https://usenet.krcg.org/t/9WJX_WF656A/#m1) [[LSJ
    20031212]](https://usenet.krcg.org/t/dOATNbVuaqs/#m2) [[LSJ 20070319]](https://usenet.krcg.org/t/9ArBwoYDquw/#m1)
    [[RBK combat]](https://www.vekn.net/rulebook#combat) — {Guardian, The}, {Blissful Agony}, {Taunt the Caged Beast}.

[^4-1-4]: [[ANK 20180929]](https://usenet.krcg.org/t/vekn-77035/#m7)
    [[LSJ 19980109]](https://usenet.krcg.org/t/d2U51Uw0Hfg/#m1) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[ANK
    20180928-2]](https://usenet.krcg.org/t/vekn-77035/#m1) — {Marie-Pierre}.

[^4-1-5]: [[LSJ 20070217]](https://usenet.krcg.org/t/HkKuwBe9LRk/#m2) — {Champion}.

[^4-1-6]: [[LSJ 20060410]](https://usenet.krcg.org/t/jr8wSeSchsc/#m1) [[LSJ
    20020204]](https://usenet.krcg.org/t/DDjGFeolsxg/#m1) — {Angel of Berlin}, {Internal Recursion}.

[^4-1-9]: [[ANK 20181101]](https://usenet.krcg.org/t/vekn-77132/#m1) [[RBK
    combat]](https://www.vekn.net/rulebook#combat) — {Nosferatu Putrescence}, {Bliss}, {Martyr's Resilience}.

[^4-1-10]: [[LSJ 20080618]](https://usenet.krcg.org/t/9FAkwIsuYZc/#m3) — group "Cancel with no action" (G00063).

[^4-2-1]: [[RTR 19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[ANK
    20180720]](https://usenet.krcg.org/t/vekn-76840/#m1) — {Neutral Guard},
    {Squirrel Balance}, {Charge of the Buffalo}, {Omael Kuman}, and sixteen further cards carrying the identical
    wording.

[^4-2-3]: [[RTR 19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[ANK
    20180720]](https://usenet.krcg.org/t/vekn-76840/#m1) [[PIB
    20120214]](https://usenet.krcg.org/t/vekn-22906/#m14) — {Immortal Grapple},
    {Grasp of the Python}, {Lam Into}.

[^4-2-2]: [[PIB 20120214]](https://usenet.krcg.org/t/vekn-22906/#m14) —
    {Asanbonsam Ghoul}, {Neutral Guard}, {Squirrel Balance}, {Gang Tactics}, {Storm Sewers}.

[^4-2-6]: [[LSJ 20021028]](https://usenet.krcg.org/t/g0GGiVIxyis/#m1) [[ANK
    20200115]](https://usenet.krcg.org/t/vekn-78321/#m6) [[LSJ
    20010814]](https://usenet.krcg.org/t/z7uYIO39YCo/#m1) [[RTR 20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) —
    {Sniper Rifle}.

[^4-2-4]: [[LSJ 20100709]](https://usenet.krcg.org/t/j98aqFIFjnE/#m23) [[RBK
    before-range]](https://www.vekn.net/rulebook#before-range) — {Fear of the Void Below}.

[^4-2-5]: [[ANK
    20170930]](https://usenet.krcg.org/t/vekn-76197/#m1)
    — group "Damage before range is determined" (G00164), {Weather Control}.

[^4-2-10]: [[ANK 20210422]](https://usenet.krcg.org/t/vekn-79111/#m5) [[LSJ
    20040830]](https://usenet.krcg.org/t/KHMhiNiSKo4/#m14) [[ANK
    20200703]](https://usenet.krcg.org/t/vekn-78713/#m1)
    [[LSJ 20031008]](https://usenet.krcg.org/t/JUnIGIrb3pw/#m2) — {Vampiric Disease}, {Blood to Water}.

[^4-2-11]: [[RTR 19980928]](https://usenet.krcg.org/t/Xva4_IRavxM/#m0) [[ANK
    20211102]](https://usenet.krcg.org/t/vekn-79447/#m1) — {Rötschreck}.

[^4-2-12]: [[LSJ 19970801]](https://usenet.krcg.org/t/7ZnLKTdPEn4/#m0) [[ANK
    20221212]](https://usenet.krcg.org/t/vekn-80208/#m7) [[ANK
    20181024]](https://usenet.krcg.org/t/vekn-77108/#m5) [[LSJ
    19970915]](https://usenet.krcg.org/t/in3R3mQs5Do/#m0) [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) —
    {Burst of Sunlight}, {Potio Martyrium}, {Riposte}, {Blood of the Cobra}.

[^4-2-13]: [[LSJ 20060509]](https://usenet.krcg.org/t/KBFnBRrOGB4/#m1) — {Earthshock}.

[^4-3-1]: [[ANK 20180704]](https://usenet.krcg.org/t/vekn-76784/#m2) [[LSJ
    20050304]](https://usenet.krcg.org/t/X5JZBbSIj3U/#m2) [[RBK combat]](https://www.vekn.net/rulebook#combat) —
    {Blessed Blade}, {Projectile}.

[^4-3-2]: [[ANK
    20210928-2]](https://usenet.krcg.org/t/vekn-79330/#m14)
    — {Thoughts Betrayed}.

[^4-3-19]: [[LSJ 20090529]](https://usenet.krcg.org/t/LpFSLRuWONA/#m6) — {Jann Berger}, group "Adding costs to a strike
    or strike card" (G00074).

[^4-3-20]: [[ANK
    20210928-2]](https://usenet.krcg.org/t/vekn-79330/#m14)
    — {Thoughts Betrayed}.

[^4-3-3]: [[LSJ 20080409]](https://usenet.krcg.org/t/jrN1Pc_Quk0/#m1) [[TOM
    19951217]](https://usenet.krcg.org/t/Y2_A66iRqMc/#m9) [[LSJ 20100308]](https://usenet.krcg.org/t/-euEN_y8ius/#m1) —
    {Target Vitals}, {Target Head}, {Target Leg}, {Target Hand}, {Target Retainer}, {Immortal Grapple}, {Mind of the
    Wilds}.

[^4-3-4]: [[LSJ 20080702-1]](https://usenet.krcg.org/t/JtLgB7Apqq0/#m1) [[LSJ
    20090114]](https://usenet.krcg.org/t/fZIdIRDDxdo/#m4) [[TOM 19960225]](https://usenet.krcg.org/t/0LLTOfvyVbM/#m11)
    [[LSJ 20071020]](https://usenet.krcg.org/t/ap2IBEgm0gI/#m1) [[LSJ
    20010627]](https://usenet.krcg.org/t/NhNCVCCDyU0/#m5) [[LSJ 20010919]](https://usenet.krcg.org/t/08AhThj0IxI/#m16) —
    {Bundi}, {Lucky Blow}, {Scorpion's Touch}, {Stutter-Step}.

[^4-3-5]: [[LSJ 20081120-2]](https://usenet.krcg.org/t/e2PNDpg-l_c/#m14) [[RTR
    20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[LSJ 20050224-3]](https://usenet.krcg.org/t/tzl6213crbI/#m6)
    — {Heroic Might}, {Blood of the Cobra}, {Projectile}.

[^4-3-7]: [[RTR 19990105]](https://usenet.krcg.org/t/2LT2jSX4Nlk/#m1) [[RTR
    19960112]](https://usenet.krcg.org/t/d3n3StNS7no/#m0) [[LSJ 20000215]](https://usenet.krcg.org/t/8to3EBsk-iY/#m0)
    [[ANK
    20200703]](https://usenet.krcg.org/t/vekn-78713/#m1)
    [[LSJ 20031008]](https://usenet.krcg.org/t/JUnIGIrb3pw/#m2) — group "Improve weapon before resolution" (G00050),
    {Blood Agony}, {Wolf Claws}, {Backstab}, {Blood to Water}.

[^4-3-8]: [[LSJ 20021028]](https://usenet.krcg.org/t/g0GGiVIxyis/#m1) [[LSJ
    20001126]](https://usenet.krcg.org/t/ZK97Kk4WQ00/#m1) [[LSJ 20050224-3]](https://usenet.krcg.org/t/tzl6213crbI/#m6)
    — groups "Weapon with optional maneuver" (G00024) and "Cannot strike" (G00093), {Hidden Lurker}, {Deer Rifle},
    {Lapse}, {Projectile}.

[^4-3-21]: [[TOM 19960521]](https://usenet.krcg.org/t/poYD3n0TKGo/#m4) — group "Optional press" (G00096); as previously
    recorded for member {Dust to Dust}, the ruling covered its optional maneuver as well.

[^4-3-9]: [[RBK cancel-a-card]](https://www.vekn.net/rulebook#cancel-a-card) [[LSJ
    20090818]](https://usenet.krcg.org/t/jkKBGVLmFHc/#m1) [[LSJ 20100206]](https://usenet.krcg.org/t/cAGrXqpO-YQ/#m1)
    [[RTR 20040501]](https://usenet.krcg.org/t/7-mp3Ada86I/#m0) [[LSJ
    20050228-3]](https://usenet.krcg.org/t/UHEZEmX22jA/#m4) [[ANK
    20230111]](https://usenet.krcg.org/t/vekn-80258/#m2)
    — {Supernatural Resistance}, {Death Seeker}, {Primal Instincts}, {Rigor Mortis}.

[^4-3-10]: [[ANK
    20200909]](https://usenet.krcg.org/t/vekn-78845/#m2)
    [[LSJ 20040130]](https://usenet.krcg.org/t/wvuR79dNCDU/#m2) [[LSJ
    20001127-2]](https://usenet.krcg.org/t/KInac4MQMuA/#m4) [[LSJ 20070508]](https://usenet.krcg.org/t/FOLkbrSh0Ns/#m5)
    [[LSJ 20010806-1]](https://usenet.krcg.org/t/PuawBcgSIKI/#m5) [[PIB
    20141026]](https://usenet.krcg.org/t/vekn-66960/#m3) [[LSJ
    20100310]](https://usenet.krcg.org/t/hYN6L3COpqw/#m1) — {Contagion}, {Escaped Mental Patient}, {Flash Grenade},
    {Gianna di Canneto}, {Zip Gun}.

[^4-3-11]: [[LSJ 20051123]](https://usenet.krcg.org/t/Ww3zSY8cVNs/#m2) [[LSJ
    20080212]](https://usenet.krcg.org/t/Yg1nZfgkpGM/#m3) [[RBK
    strike-effects]](https://www.vekn.net/rulebook#strike-effects) — {Darkness Within}, {Young Bloods}.

[^4-3-12]: [[RTR 19960221]](https://usenet.krcg.org/t/UdU535eVm0Y/#m0) [[PIB
    20130319]](https://usenet.krcg.org/t/vekn-46164/#m3) [[ANK
    20170111]](https://usenet.krcg.org/t/vekn-72635/#m10)
    [[LSJ 20071011]](https://usenet.krcg.org/t/GSQqnAEBkgU/#m2) [[RTR
    19980928]](https://usenet.krcg.org/t/Xva4_IRavxM/#m0) — {Target Vitals}, {Target Head}, {Talith}, {Dam the Heart's
    River}, group "Damage after combat ends" (G00091).

[^4-3-18]: [[ANK 20171225]](https://usenet.krcg.org/t/vekn-76349/#m2) [[RBK
    combat]](https://www.vekn.net/rulebook#combat) — {Shoulder Drop}, {Coordinate Attacks}.

[^4-3-13]: [[LSJ 20040928]](https://usenet.krcg.org/t/GBcR6aNjIk4/#m1) [[ANK
    20170519]](https://usenet.krcg.org/t/vekn-75805/#m1) — {Soul
    Burn}, {Blood Fury}.

[^4-3-14]: [[PIB 20110830]](https://usenet.krcg.org/t/vekn-9340/#m1) [[LSJ
    20030307]](https://usenet.krcg.org/t/zhdj4jnSdrM/#m1) [[LSJ 20100119]](https://usenet.krcg.org/t/1eULCGaVcO0/#m1)
    [[ANK 20180913-1]](https://usenet.krcg.org/t/vekn-77003/#m5) [[RBK
    first-strike]](https://www.vekn.net/rulebook#first-strike) — {Scorpion's Touch}, {Shambling Hordes}, {Stutter-Step}.

[^4-3-17]: [[LSJ 20001206]](https://usenet.krcg.org/t/kFIO74LxqFQ/#m4) [[LSJ
    20080210]](https://usenet.krcg.org/t/nL-xqiydvYg/#m1) [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0)
    [[RBK additional-strikes]](https://www.vekn.net/rulebook#additional-strikes) — {Quickness}, {Ghoul Retainer}, group
    "Retainer that strike" (G00029).

[^4-4-1]: [[LSJ 19971211-2]](https://usenet.krcg.org/t/2s5V_AIczYw/#m1) [[PIB
    20130303]](https://usenet.krcg.org/t/vekn-45620/#m4) [[LSJ
    20100813]](https://usenet.krcg.org/t/nOb3cmvA_3U/#m1) [[RTR 19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0)
    [[LSJ 20020821]](https://usenet.krcg.org/t/o_nBNZraMvE/#m1) [[LSJ
    20020904]](https://usenet.krcg.org/t/Qmi4hFk6QqE/#m2) [[LSJ 20030307]](https://usenet.krcg.org/t/zhdj4jnSdrM/#m1)
    [[LSJ 20100119]](https://usenet.krcg.org/t/1eULCGaVcO0/#m1) — {Erosion}, {Torn Signpost}, {Illegal Search and
    Seizure}, {Concealed Weapon}, {Shambling Hordes}.

[^4-4-2]: [[TOM 19960225]](https://usenet.krcg.org/t/0LLTOfvyVbM/#m11) [[RTR
    19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1)
    [[ANK 20210620]](https://usenet.krcg.org/t/vekn-79186/#m2) — {Target
    Vitals}, {Increased Strength}, {Glaser Rounds}, {Merrill Molitor}, {Alejandro Aguirre}, and ~45 further cards
    carrying the same sentence.

[^4-4-3]: [[RTR 19960221]](https://usenet.krcg.org/t/UdU535eVm0Y/#m0) [[LSJ
    20080226]](https://usenet.krcg.org/t/NAfcSQEiLXY/#m1) — {Dam the Heart's River}, {Lucky Blow}, {Rowan Ring},
    {Oubliette}.

[^4-4-4]: [[RTR 19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[LSJ
    19990723]](https://usenet.krcg.org/t/6Lx2jnw7hMA/#m1) [[PIB
    20120426]](https://usenet.krcg.org/t/vekn-27863/#m8) [[ANK
    20170111]](https://usenet.krcg.org/t/vekn-72635/#m10)
    [[LSJ 20071011]](https://usenet.krcg.org/t/GSQqnAEBkgU/#m2) [[PIB
    20130319]](https://usenet.krcg.org/t/vekn-46164/#m3) [[LSJ
    20080630]](https://usenet.krcg.org/t/nvuXBpEaKAA/#m2) — {Catatonic Fear}, {Loving Agony}, {Outside the Hourglass},
    {Riposte}, {Nephandus}.

[^4-4-5]: [[RTR 19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[LSJ
    19970801]](https://usenet.krcg.org/t/7ZnLKTdPEn4/#m0) [[LSJ 20060217]](https://usenet.krcg.org/t/BhecJx5BqtQ/#m19)
    [[RBK damage-resolution]](https://www.vekn.net/rulebook#damage-resolution) — {Grenade}, {Burst of Sunlight},
    {Shemti}, and ~24 further cards carrying the same sentence.

[^4-4-6]: [[LSJ 20090304-1]](https://usenet.krcg.org/t/B7EQiAVOr1Q/#m1) [[LSJ
    20100514]](https://usenet.krcg.org/t/Mz_fcIldfAY/#m2) [[LSJ 20051220]](https://usenet.krcg.org/t/bIcX7F3wzf0/#m1)
    [[ANK 20180612]](https://usenet.krcg.org/t/vekn-76717/#m2) [[ANK
    20211127]](https://usenet.krcg.org/t/vekn-76687/#m9)
    — {Talbot's Chainsaw}, {Kerrie}, group "Retainer doing damage" (G00016).

[^4-4-7]: [[TOM 19950304]](https://usenet.krcg.org/t/DYxpGPWoSaE/#m5) [[LSJ
    20010812]](https://usenet.krcg.org/t/uXGLWUAuByc/#m1) [[ANK
    20201228-2]](https://usenet.krcg.org/t/vekn-78954/#m4) [[ANK
    20200925]](https://usenet.krcg.org/t/vekn-78861/#m1) [[ANK
    20200517]](https://usenet.krcg.org/t/vekn-78638/#m1) [[PIB
    20111017]](https://usenet.krcg.org/t/vekn-12104/#m2)
    [[LSJ 20010626]](https://usenet.krcg.org/t/tkCRUBRp82E/#m1) — {Pulled Fangs}, {Necrosis}, {Nephandus}, {Blood of
    Acid}, {Disarm}.

[^4-4-8]: [[ANK
    20181128]](https://usenet.krcg.org/t/vekn-77187/#m4) [[PIB
    20150520]](https://usenet.krcg.org/t/vekn-68512/#m4)
    [[PIB 20150603-1]](https://usenet.krcg.org/t/vekn-71468/#m1)
    [[LSJ 20020314]](https://usenet.krcg.org/t/1uKDkopJTRo/#m1) [[PIB
    20130319]](https://usenet.krcg.org/t/vekn-46164/#m3) — {Dam the
    Heart's River}, {Dagon's Call}, {Dawn Operation}, {Domain of Evernight}.

[^4-4-11]: [[ANK 20200130-1]](https://usenet.krcg.org/t/vekn-78400/#m1) [[ANK
    20220114]](https://usenet.krcg.org/t/vekn-79432/#m17)
    [[LSJ 20060409]](https://usenet.krcg.org/t/gsFQXsCGTG4/#m1) [[LSJ
    20030213]](https://usenet.krcg.org/t/j6cuQ6pFJSA/#m1) — {Rötschreck}, {Shadow Court Satyr}, {Merrill Molitor}.

[^4-4-12]: [[ANK 20210424]](https://usenet.krcg.org/t/vekn-79121/#m2) [[LSJ
    20021120]](https://usenet.krcg.org/t/4TgUatcTtdk/#m21) [[RBK
    damage-resolution]](https://www.vekn.net/rulebook#damage-resolution) — {Byzar}, {Anathema}.

[^4-4-13]: [[ANK
    20170427]](https://usenet.krcg.org/t/vekn-75755/#m6) [[LSJ
    20040805]](https://usenet.krcg.org/t/WuER8RUMzTE/#m13) [[ANK
    20200517-2]](https://usenet.krcg.org/t/vekn-78629/#m10)
    [[LSJ 20001111]](https://usenet.krcg.org/t/m23Hj3OW2A4/#m1) [[PIB
    20150426]](https://usenet.krcg.org/t/vekn-70713/#m1)
    — {Blood of Acid}, {Tunnel Runner}, {Vagabond Mystic}.

[^4-5-1]: [[LSJ 20001114]](https://usenet.krcg.org/t/qXSlM7Grq1c/#m1) — group "Immideate damage prevention" (G00154),
    {Hidden Strength}.

[^4-5-2]: [[LSJ 20010315]](https://usenet.krcg.org/t/m9CrEOn1veo/#m3) [[LSJ
    20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) — {Apparition}, {Brother's Blood}.

[^4-5-3]: [[LSJ 20011214-2]](https://usenet.krcg.org/t/TinQ8ywzIHU/#m1) — {Repulsion}.

[^4-5-5]: [[RTR 20041202]](https://usenet.krcg.org/t/WUWh7AdooDU/#m5) — {Soak}, {Nightstick}, {Forearm Block}, {Rego
    Motum}.

[^4-5-6]: [[RTR 20041202]](https://usenet.krcg.org/t/WUWh7AdooDU/#m5) [[PIB
    20121202]](https://usenet.krcg.org/t/vekn-42195/#m4)
    [[LSJ 20030121]](https://usenet.krcg.org/t/lED3kZ2UUUo/#m3) [[LSJ
    20001114]](https://usenet.krcg.org/t/qXSlM7Grq1c/#m1) — {Hidden Strength}, {Martyr's Resilience}.

[^4-5-7]: [[LSJ 20001114]](https://usenet.krcg.org/t/qXSlM7Grq1c/#m1) — {Armor of Vitality}, {Rego Motum}, {Undead
    Persistence}, {Skin of Rock}.

[^4-5-8]: [[ANK 20200318]](https://usenet.krcg.org/t/vekn-78525/#m1) [[LSJ
    20081210]](https://usenet.krcg.org/t/pUUGH1nSKpc/#m1) — {Armor of Caine's Fury}, {Apparition}, {Kevlar Vest}.

[^4-5-9]: [[LSJ 20040812-2]](https://usenet.krcg.org/t/E7D1cVmAdqQ/#m5) [[PIB
    20130327]](https://usenet.krcg.org/t/vekn-46279/#m3) [[LSJ
    20021001]](https://usenet.krcg.org/t/FW-bmZpIM68/#m1) — group "Prevent non-aggravated" (G00033), {Flesh of Marble},
    {Resilience}.

[^4-5-11]: [[LSJ 19980216]](https://usenet.krcg.org/t/kyuR_x6pRmo/#m3) — {Blood Fury}, {Blood Rage}, {Soul Burn}.

[^4-5-12]: [[RTR 19980928]](https://usenet.krcg.org/t/Xva4_IRavxM/#m0) [[ANK
    20211102]](https://usenet.krcg.org/t/vekn-79447/#m1) [[ANK
    20230823]](https://usenet.krcg.org/t/vekn-75182/#m10)
    [[LSJ 19980219]](https://usenet.krcg.org/t/FoVObdCLJPM/#m1) — {Rötschreck}, {Bollix}, {Rowan Ring}.

[^4-5-13]: [[LSJ 20100519]](https://usenet.krcg.org/t/criytfhw97o/#m1) [[LSJ
    20090304-1]](https://usenet.krcg.org/t/B7EQiAVOr1Q/#m1) — {Nephandus}, group "Immune to damage" (G00026).

[^4-5-14]: [[RTR 19960112]](https://usenet.krcg.org/t/d3n3StNS7no/#m0) [[PIB
    20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) —
    {Chiropteran Marauder}, {Claws of the Dead}, {The Coven}.

[^4-5-15]: [[ANK
    20200517-2]](https://usenet.krcg.org/t/vekn-78629/#m10)
    [[ANK 20170427]](https://usenet.krcg.org/t/vekn-75755/#m6) [[LSJ
    20040805]](https://usenet.krcg.org/t/WuER8RUMzTE/#m13) [[LSJ 20100205]](https://usenet.krcg.org/t/owQe9egif0U/#m3) —
    {Blood of Acid}, {Tunnel Runner}, {Potio Martyrium}.

[^4-5-17]: [[LSJ 20010627]](https://usenet.krcg.org/t/NhNCVCCDyU0/#m5) [[LSJ
    20010629]](https://usenet.krcg.org/t/MhHKBY7II78/#m1) [[LSJ 20090304-1]](https://usenet.krcg.org/t/B7EQiAVOr1Q/#m1)
    [[RBK damage-resolution]](https://www.vekn.net/rulebook#damage-resolution) — group "Immune to damage" (G00026),
    {Bloodform}, {Ignore the Searing Flames}.

[^4-5-18]: [[LSJ 20040120]](https://usenet.krcg.org/t/5zBRtwg6lpU/#m9) — {Ex Nihilo}.

[^4-5-19]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[ANK
    20180612]](https://usenet.krcg.org/t/vekn-76717/#m2) [[ANK
    20211127]](https://usenet.krcg.org/t/vekn-76687/#m9)
    — {Charnas the Imp}, group "Retainer doing damage" (G00016).

[^4-5-20]: [[RBK damage-resolution]](https://www.vekn.net/rulebook#damage-resolution) [[LSJ
    20020215]](https://usenet.krcg.org/t/FArran2SmSQ/#m1) [[ANK 20171226]](https://usenet.krcg.org/t/vekn-76349/#m8)
    [[ANK 20210324-1]](https://usenet.krcg.org/t/vekn-79087/#m12) [[ANK
    20210324-2]](https://usenet.krcg.org/t/vekn-79087/#m15) [[LSJ 20040802]](https://usenet.krcg.org/t/b6V_OGl-TgE/#m1)
    [[LSJ 20100519]](https://usenet.krcg.org/t/criytfhw97o/#m1) [[LSJ
    20010627]](https://usenet.krcg.org/t/NhNCVCCDyU0/#m5) [[LSJ 20080820]](https://usenet.krcg.org/t/G-TY3oNSBkE/#m1) —
    {Improvised Flamethrower}, {Weighted Walking Stick}, {Shoulder Drop}, {Drink the Blood of Ahriman}, {Pulled Fangs},
    {Target Vitals}, {Twisting the Knife}, {Blade Clot}, {Haymaker}. LSJ's "'Inflict' normally means 'successfully
    inflicts'" (LSJ 20040802, LSJ 20080820) describes the older wording, which the cards he ruled on keep; {Haymaker}
    awaits a Rules Director ruling (owner, 2026-11-12).

[^4-5-21]: [[LSJ 20090622-2]](https://usenet.krcg.org/t/1zt1SZb2TIk/#m1) — {Coma}, {Rowan Ring}, {Rabbat, The Sewer
    Goddess}, {Seren Sukardi}.

[^4-6-1]: [[RTR 20041202]](https://usenet.krcg.org/t/WUWh7AdooDU/#m5) [[LSJ
    19980526]](https://usenet.krcg.org/t/tRTwM9wYaVI/#m0) [[LSJ 19980219]](https://usenet.krcg.org/t/FoVObdCLJPM/#m1)
    [[ANK 20220923]](https://usenet.krcg.org/t/vekn-80044/#m2) [[ANK
    20200909]](https://usenet.krcg.org/t/vekn-78845/#m2)
    [[LSJ 20040130]](https://usenet.krcg.org/t/wvuR79dNCDU/#m2) [[LSJ
    20051020]](https://usenet.krcg.org/t/jqPoW66sEkY/#m1) — {Serpent's Numbing Kiss}, {Morphean Blow}, {Oubliette},
    {Rowan Ring}, {Enhanced Coagulant}, {Contagion}, {Blissful Agony}.

[^4-6-2]: [[LSJ 20040928]](https://usenet.krcg.org/t/GBcR6aNjIk4/#m1) — {Blood Fury}, {Blood Rage}, {Personal Scourge},
    {Soul Burn}.

[^4-6-3]: [[LSJ 19990106-2]](https://usenet.krcg.org/t/6dDhJmBlpXY/#m0) — {Fleshcraft}, {Bonecraft}.

[^4-6-4]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[TOM
    19951216]](https://usenet.krcg.org/t/0G-jqB_fcyE/#m1) [[PIB
    20140324]](https://usenet.krcg.org/t/vekn-60259/#m1)
    [[ANK
    20200422]](https://usenet.krcg.org/t/vekn-78582/#m3)
    [[RBK strike-effects]](https://www.vekn.net/rulebook#strike-effects) — group "Environmental damage" (G00017),
    {Conscripted Statue}, {Darkling Trickery}, {Thoughts Betrayed}.

[^4-6-5]: [[ANK 20200517]](https://usenet.krcg.org/t/vekn-78638/#m1) —
    {Necrosis}.

[^4-6-6]: [[LSJ 20011210-2]](https://usenet.krcg.org/t/mUwzyHLcAx8/#m6) [[LSJ
    20011005]](https://usenet.krcg.org/t/9fyH2X1YGAQ/#m1) [[ANK
    20211102]](https://usenet.krcg.org/t/vekn-79447/#m1) [[RBK
    strike-effects]](https://www.vekn.net/rulebook#strike-effects) — {Anesthetic Touch}, {Rötschreck}.

[^4-6-7]: [[RTR 20041202]](https://usenet.krcg.org/t/WUWh7AdooDU/#m5) [[LSJ
    20060808]](https://usenet.krcg.org/t/ncZn3knH-Uo/#m1) [[ANK
    20210612]](https://usenet.krcg.org/t/vekn-79173/#m11) [[LSJ
    20020416-2]](https://usenet.krcg.org/t/tEC5uN8yqUE/#m4) — {Flash Grenade}, {Garrote}.

[^4-6-8]: [[LSJ 20030902-2]](https://usenet.krcg.org/t/mgZt4f2LyOg/#m6) [[LSJ
    20060808-1]](https://usenet.krcg.org/t/7oK-hKbOs9g/#m2) — group "Cannot be dodged" (G00109), {Shadow Feint}.

[^4-6-9]: [[LSJ 20010919]](https://usenet.krcg.org/t/08AhThj0IxI/#m16) [[ANK
    20180913-1]](https://usenet.krcg.org/t/vekn-77003/#m5) — {Stutter-Step}.

[^4-7-1]: [[LSJ 19981006]](https://usenet.krcg.org/t/RU5yM2Ov5Mg/#m0) [[LSJ
    20001127-2]](https://usenet.krcg.org/t/KInac4MQMuA/#m4) [[LSJ
    20010806-1]](https://usenet.krcg.org/t/PuawBcgSIKI/#m5) [[LSJ 20010627]](https://usenet.krcg.org/t/NhNCVCCDyU0/#m5)
    [[LSJ 20070508]](https://usenet.krcg.org/t/FOLkbrSh0Ns/#m5) — {Bomb}, {Jar of Skin Eaters}, {Waxen Poetica}, {White
    Phosphorus Grenade}, {Dragon's Breath Rounds}, {Elixir of Distillation}, {Escaped Mental Patient}, {Grenade}, {Smoke
    Grenade}, {Molotov Cocktail}.

[^4-7-2]: [[LSJ 20051123]](https://usenet.krcg.org/t/Ww3zSY8cVNs/#m2) [[ANK
    20200203-1]](https://usenet.krcg.org/t/vekn-78416/#m1) [[LSJ
    19980219]](https://usenet.krcg.org/t/FoVObdCLJPM/#m1) [[ANK
    20220923]](https://usenet.krcg.org/t/vekn-80044/#m2) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) — {Darkness Within}, {Molotov Cocktail}, {Rowan Ring},
    {Internal Recursion}.

[^4-7-3]: [[LSJ 20001127-2]](https://usenet.krcg.org/t/KInac4MQMuA/#m4) [[LSJ
    20010806-1]](https://usenet.krcg.org/t/PuawBcgSIKI/#m5) — {Flash Grenade}, {Smoke Grenade}.

[^4-7-4]: [[ANK 20211130]](https://usenet.krcg.org/t/vekn-66960/#m8) [[RTR
    19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[TOM 19951216]](https://usenet.krcg.org/t/0G-jqB_fcyE/#m1)
    [[PIB
    20140324]](https://usenet.krcg.org/t/vekn-60259/#m1)
    [[ANK
    20200422]](https://usenet.krcg.org/t/vekn-78582/#m3)
    [[RBK first-strike]](https://www.vekn.net/rulebook#first-strike) — {Gianna di Canneto}, group "Environmental damage"
    (G00017).

[^4-7-5]: [[ANK
    20190412]](https://usenet.krcg.org/t/vekn-77543/#m6) — {Dog
    Pack}.

[^4-7-6]: [[LSJ 20011005]](https://usenet.krcg.org/t/9fyH2X1YGAQ/#m1) [[ANK
    20211102]](https://usenet.krcg.org/t/vekn-79447/#m1) [[LSJ
    19990217]](https://usenet.krcg.org/t/9Bsf2LC1274/#m1) [[LSJ 20020715]](https://usenet.krcg.org/t/-FBxPz-oms0/#m1)
    [[LSJ 20060803]](https://usenet.krcg.org/t/z8Nd8EFQsKc/#m1) — {Rötschreck}.

[^4-7-7]: [[LSJ 20011210-1]](https://usenet.krcg.org/t/mUwzyHLcAx8/#m22) [[LSJ
    20011210-2]](https://usenet.krcg.org/t/mUwzyHLcAx8/#m6) [[LSJ
    20011210-3]](https://usenet.krcg.org/t/mUwzyHLcAx8/#m12) [[LSJ
    20071112]](https://usenet.krcg.org/t/2fFyhL6YqI0/#m11) — {Anesthetic Touch}, {Autonomic Mastery}.

[^4-8-1]: [[TOM 19960521]](https://usenet.krcg.org/t/poYD3n0TKGo/#m4) — {Immortal Grapple}, {Dust to Dust}, {Chameleon's
    Colors}, and 33 further cards carrying the same wording.

[^4-8-2]: [[ANK 20180110]](https://usenet.krcg.org/t/vekn-76362/#m1) — {Undead Persistence} (ruling removed from the
    database when group "Optional press" (G00096) was created; original verified at source).

[^4-8-3]: [[PIB
    20150121]](https://usenet.krcg.org/t/vekn-68745/#m4) —
    {Aeron}, {Don Caravelli}, and six further crypt cards carrying the same wording.

[^4-8-4]: [[LSJ 20051016]](https://usenet.krcg.org/t/s_At5syL66k/#m6) [[LSJ
    20051017]](https://usenet.krcg.org/t/s_At5syL66k/#m8) [[TOM 19960521]](https://usenet.krcg.org/t/poYD3n0TKGo/#m4) —
    {Talbot's Chainsaw}, {Lorrie Dunsirn}, {Chameleon's Colors}, {Mighty Grapple}.

[^4-8-5]: [[ANK 20200525-1]](https://usenet.krcg.org/t/vekn-78650/#m1) — {Trap}.

[^4-8-6]: [[LSJ 20040521]](https://usenet.krcg.org/t/Giv8nDnYVGo/#m2) — {Trap}.

[^4-8-7]: [[ANK
    20230114]](https://usenet.krcg.org/t/vekn-80258/#m7)
    — {Rigor Mortis}.

[^4-8-8]: [[RTR 19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) — {Drawing Out the Beast}.

[^4-8-9]: [[LSJ 20010312]](https://usenet.krcg.org/t/Mz0Rs9faJ9o/#m1) [[LSJ
    20001027]](https://usenet.krcg.org/t/zUs5JBhRLgE/#m0) — group "Effect during press step" (G00114).

[^4-9-1]: [[RTR 20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0) — {Disarm}, {Pulled Fangs}, {Revelation of Wrath},
    {Street Cred}, {Taste of Vitae}, {Ossian}, {Masochism}.

[^4-9-2]: [[LSJ 20021113]](https://usenet.krcg.org/t/df2P8YHZex8/#m11) [[ANK
    20191219]](https://usenet.krcg.org/t/vekn-78241/#m3)
    [[ANK
    20180910-1]](https://usenet.krcg.org/t/vekn-76990/#m3)
    — {Disarm}, {Pulled Fangs}, {Street Cred}, {Taste of Vitae}, {Telepathic Tracking}, {Relentless Reaper}.

[^4-9-3]: [[RTR 20001020]](https://usenet.krcg.org/t/GvxNYsYsWJ4/#m0) [[LSJ
    20001024]](https://usenet.krcg.org/t/GvxNYsYsWJ4/#m6) [[LSJ 20011214-4]](https://usenet.krcg.org/t/RSCjaaZXY28/#m1)
    — {Elysium: The Arboretum}, {Alpha Glint}, {Garibaldi-Meucci Museum}, {Ashes to Ashes}.

[^4-9-4]: [[ANK
    20180910-2]](https://usenet.krcg.org/t/vekn-76990/#m6)
    [[ANK
    20180910-3]](https://usenet.krcg.org/t/vekn-76990/#m2)
    [[ANK 20180618]](https://usenet.krcg.org/t/vekn-76735/#m1) — {Death Seeker},
    {Terror Frenzy}, {Loving Agony}.

[^4-9-5]: [[RTR 20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0) [[LSJ
    20021113]](https://usenet.krcg.org/t/df2P8YHZex8/#m11) [[ANK
    20191219]](https://usenet.krcg.org/t/vekn-78241/#m3)
    [[ANK
    20180910-1]](https://usenet.krcg.org/t/vekn-76990/#m3)
    — {Psyche!}.

[^4-9-6]: [[ANK
    20191219]](https://usenet.krcg.org/t/vekn-78241/#m3)
    — {Psyche!}, {Telepathic Tracking}.

[^4-9-7]: [[LSJ 19980109]](https://usenet.krcg.org/t/d2U51Uw0Hfg/#m1) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[ANK
    20180928-2]](https://usenet.krcg.org/t/vekn-77035/#m1) — {Psyche!}, {Fast
    Reaction}, {Haven Hunting}, {Akram}, {Jalal Sayad}, {Marie-Pierre}.

[^4-9-8]: [[RTR 19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[LSJ 20030717]](https://usenet.krcg.org/t/IkMD8wfAcqw/#m1) —
    {Mummify}, {Meld with the Land}, {Bond with the Mountain}, {Earth Meld}, {Majesty}, {Loving Agony}.

[^4-9-23]: [[RTR 20041202]](https://usenet.krcg.org/t/WUWh7AdooDU/#m5) [[LSJ
    20060808]](https://usenet.krcg.org/t/ncZn3knH-Uo/#m1) [[LSJ 20040602]](https://usenet.krcg.org/t/bEvYg3Za-fc/#m3)
    [[LSJ 20040601]](https://usenet.krcg.org/t/bEvYg3Za-fc/#m1) [[LSJ
    20071009]](https://usenet.krcg.org/t/vSI2MhyB710/#m9) — {Flash Grenade}, {Earth Meld}, {Rötschreck}, {Loving Agony}.

[^4-9-9]: [[RTR 20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[LSJ
    20040602]](https://usenet.krcg.org/t/bEvYg3Za-fc/#m3) [[LSJ 20071009]](https://usenet.krcg.org/t/vSI2MhyB710/#m9)
    [[LSJ 20021126]](https://usenet.krcg.org/t/9Ui7pesvu5g/#m1) — {Smoke Grenade}, {Flash Grenade}, {Rötschreck},
    {Morphean Blow}, {Legacy of Power}, {Illusions of the Kindred}.

[^4-9-10]: [[RTR 19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[ANK
    20180420-1]](https://usenet.krcg.org/t/vekn-76522/#m1) [[ANK
    20200705]](https://usenet.krcg.org/t/vekn-66086/#m6) — {Catatonic
    Fear}, {Riposte}, {Serpent's Numbing Kiss}, {Mercy for the Weak}, {Unholy Penance}, {Mariel, Lady Thunder},
    {Oubliette}.

[^4-9-11]: [[ANK
    20170111]](https://usenet.krcg.org/t/vekn-72635/#m10)
    [[LSJ 20071011]](https://usenet.krcg.org/t/GSQqnAEBkgU/#m2) [[LSJ
    19980225]](https://usenet.krcg.org/t/62y-5miA8MQ/#m0) — {Dam the Heart's River}, {Anathema}.

[^4-9-13]: [[ANK
    20180104]](https://usenet.krcg.org/t/vekn-76356/#m4)
    [[LSJ 19971110]](https://usenet.krcg.org/t/IlQdghRurtM/#m4) [[ANK
    20190725]](https://usenet.krcg.org/t/vekn-77813/#m1) — {Outside the Hourglass},
    {Weather Control}.

[^4-9-24]: [[RTR 20250612]](https://www.blackchantry.com/2025/06/12/new-sabbat-rules-terms-and-errata/) [[ANK
    20250613]](https://usenet.krcg.org/t/vekn-82396/#m0) — {Hedonism}, {Psyche!},
    {Hidden Lurker}, {Fast Reaction}.

[^4-9-25]: [[ANK 20250614]](https://usenet.krcg.org/t/vekn-82396/#m8) — {Siren's
    Lure}.

[^4-9-15]: [[RTR 20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[LSJ
    19971003]](https://usenet.krcg.org/t/mUZ2704tVsk/#m8) [[LSJ 20021121]](https://usenet.krcg.org/t/UddkVI7G8iA/#m10)
    [[LSJ 20090826]](https://usenet.krcg.org/t/KTwa1Hf_gHI/#m1) — {Psyche!}, {Akram}, {Jalal Sayad}, {Siren's Lure},
    {Illusions of the Kindred}, {Pocket Out of Time}; [[RTR
    20250612]](https://www.blackchantry.com/2025/06/12/new-sabbat-rules-terms-and-errata/) [[ANK
    20250613]](https://usenet.krcg.org/t/vekn-82396/#m0) — {Psyche!}, {Hidden Lurker},
    {Hedonism}.

[^4-9-26]: [[ANK 20250613]](https://usenet.krcg.org/t/vekn-82396/#m0) — {Telepathic
    Tracking}; [[ANK 20251103]](https://usenet.krcg.org/t/vekn-82396/#m18) —
    {Hedonism}.

[^4-9-16]: [[ANK 20220429]](https://usenet.krcg.org/t/vekn-79754/#m1) [[LSJ
    20100604-1]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m7) [[LSJ 20020509]](https://usenet.krcg.org/t/-dqteu2lStc/#m3)
    — {Hidden Lurker}, {Psyche!}, {Blissful Agony}, {Heaven's Gate}.

[^4-9-18]: [[ANK
    20230527]](https://usenet.krcg.org/t/vekn-80553/#m1) [[LSJ
    20010814]](https://usenet.krcg.org/t/z7uYIO39YCo/#m1) [[LSJ 20010813]](https://usenet.krcg.org/t/8MR4bq0Cxj4/#m1)
    [[LSJ 20010819-2]](https://usenet.krcg.org/t/8MR4bq0Cxj4/#m14) [[LSJ
    20030530]](https://usenet.krcg.org/t/SZehI8SwAc4/#m21) [[LSJ 20001122]](https://usenet.krcg.org/t/Br8FPS5mRn4/#m8) —
    {Blithe Acceptance}, {Sniper Rifle}, {Scry the Hearthstone}, {Raptor}, {Psyche!}.

[^4-9-19]: [[ANK 20200420-1]](https://usenet.krcg.org/t/vekn-78576/#m1) —
    {Amelia, The Blood Red Tears}.

[^4-9-20]: [[ANK 20180928-2]](https://usenet.krcg.org/t/vekn-77035/#m1) [[RTR
    19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0) [[LSJ 20040219]](https://usenet.krcg.org/t/zFmoLa6tzWA/#m2)
    [[ANK 20221130]](https://usenet.krcg.org/t/vekn-80176/#m3)
    — {Marie-Pierre}, {Cats' Guidance}.

[^4-9-21]: [[ANK
    20180910-1]](https://usenet.krcg.org/t/vekn-76990/#m3)
    [[LSJ 20060903]](https://usenet.krcg.org/t/QAKz6Qtr7Ts/#m8) [[LSJ
    20080821]](https://usenet.krcg.org/t/hZF1Joxwc2s/#m3) — {Yawp Court}.

[^4-9-22]: [[LSJ 20030214]](https://usenet.krcg.org/t/A3U-Dy1yx8Y/#m1) [[ANK
    20190624]](https://usenet.krcg.org/t/vekn-77737/#m3) [[LSJ
    20020211]](https://usenet.krcg.org/t/ubqDaLeG3qo/#m2) [[LSJ 20021122]](https://usenet.krcg.org/t/LieFYA_gyFo/#m11) —
    {Undead Persistence}, {Ashes to Ashes}.

[^4-10-1]: [[LSJ 19971215]](https://usenet.krcg.org/t/mfBmRrUKZQ0/#m8) [[LSJ
    20030213-2]](https://usenet.krcg.org/t/Jwj7VjhFU5o/#m4) [[LSJ 20050304]](https://usenet.krcg.org/t/X5JZBbSIj3U/#m2)
    — {Talbot's Chainsaw}, {Starshell Grenade Launcher}, {Banshee Ironwail}, {Sword of Judgment}, {Projectile}, {Jar of
    Skin Eaters}.

[^4-10-3]: [[LSJ 20021028]](https://usenet.krcg.org/t/g0GGiVIxyis/#m1) — {Sniper Rifle}.

[^4-10-4]: [[ANK
    20230316]](https://usenet.krcg.org/t/vekn-80385/#m2)
    [[LSJ 19980302-2]](https://usenet.krcg.org/t/9YVFkeiL3Js/#m1) [[ANK
    20220207]](https://usenet.krcg.org/t/vekn-79639/#m1) — groups
    "Weapon once per combat" (G00045) and "Weapon once per round" (G00046), {.44 Magnum}, {Hunger of Marduk}.

[^4-10-5]: [[RTR 19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0) [[LSJ
    20020821]](https://usenet.krcg.org/t/o_nBNZraMvE/#m1) [[LSJ 20020904]](https://usenet.krcg.org/t/Qmi4hFk6QqE/#m2)
    [[LSJ 19970224]](https://usenet.krcg.org/t/80KRDjVFkyg/#m1) — {Machine Blitz}, {Concealed Weapon}, {Illegal Search
    and Seizure}.

[^4-10-6]: [[LSJ 20010806-1]](https://usenet.krcg.org/t/PuawBcgSIKI/#m5) — {Machine Blitz}.

[^4-10-7]: [[LSJ 20060825]](https://usenet.krcg.org/t/2usG7ml8BAw/#m2) [[LSJ
    20051220]](https://usenet.krcg.org/t/bIcX7F3wzf0/#m1) [[LSJ 20090304-1]](https://usenet.krcg.org/t/B7EQiAVOr1Q/#m1)
    [[LSJ 20100514]](https://usenet.krcg.org/t/Mz_fcIldfAY/#m2) [[LSJ
    20091123-2]](https://usenet.krcg.org/t/_IwgQEvViWQ/#m3) — {Anachronism}, {Kerrie}, {Talbot's Chainsaw}, {Nephandus}.

[^4-10-8]: [[PIB 20141026]](https://usenet.krcg.org/t/vekn-66960/#m3) [[LSJ
    19980319]](https://usenet.krcg.org/t/i1Eqqm5Ctv0/#m1) [[RBK
    strike-effects]](https://www.vekn.net/rulebook#strike-effects) — {Gianna di Canneto}, {Disguised Weapon}.

[^4-10-9]: [[LSJ 19981006]](https://usenet.krcg.org/t/RU5yM2Ov5Mg/#m0) [[LSJ
    20001127-2]](https://usenet.krcg.org/t/KInac4MQMuA/#m4) [[LSJ
    20010806-1]](https://usenet.krcg.org/t/PuawBcgSIKI/#m5) [[RTR 20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0)
    [[LSJ 20040602]](https://usenet.krcg.org/t/bEvYg3Za-fc/#m3) — {Grenade}, {Dragon's Breath Rounds}, {Jar of Skin
    Eaters}, {Smoke Grenade}.

[^4-10-10]: [[LSJ 20010729]](https://usenet.krcg.org/t/RkvrP5tplXA/#m1) — {Soul Gem of Etrius}.

[^4-10-11]: [[LSJ 20020425]](https://usenet.krcg.org/t/_Rh8P34zTZo/#m2) — group "Improve weapon before resolution"
    (G00050).

[^5-1-1]: [[RTR 19990712]](https://usenet.krcg.org/t/x5k5Vw_hkI0/#m0) [[LSJ
    20051103]](https://usenet.krcg.org/t/BD7KIWBI0Cs/#m1) — {Dual Form}.

[^5-1-2]: [[PIB 20140122]](https://usenet.krcg.org/t/vekn-58586/#m3)
    [[TOM 19951209]](https://usenet.krcg.org/t/qP2j6CpBUDI/#m6) [[LSJ
    20090625]](https://usenet.krcg.org/t/C3daj-vYqu4/#m3) — {Banishment}, {The Becoming}.

[^5-1-12]: Owner's reading, no official ruling: [[rules-qa
    108.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-108--the-events-deep-sweep-a-reading-graded-the-engine-board-claude-2026-10-30)
    — {Banishment}.

[^5-1-3]: [[LSJ 20051126]](https://usenet.krcg.org/t/dDgHGEG18D0/#m1) [[ANK
    20190522]](https://usenet.krcg.org/t/vekn-77648/#m15) —
    {Hermana Hambrienta Mayor}, {Hermana Hambrienta Menor}, {Tariq, The Silent}.

[^5-1-5]: [[PIB 20150728]](https://usenet.krcg.org/t/vekn-72285/#m2) [[LSJ
    20030519-2]](https://usenet.krcg.org/t/mi8M1lSCLqo/#m3) [[LSJ 20030527]](https://usenet.krcg.org/t/1hViTSXv544/#m4)
    — {Chameleon}, {Legendary Vampire}, {Masquerade Enforcement}.

[^5-1-6]: [[ANK
    20180524]](https://usenet.krcg.org/t/vekn-76595/#m61) [[ANK
    20180523]](https://usenet.krcg.org/t/vekn-76595/#m55) [[LSJ
    20031121-3]](https://usenet.krcg.org/t/JQO1RBmvV_o/#m9) [[RBK
    influence-phase]](https://www.vekn.net/rulebook#influence-phase) — the [MERGED] crypt-card clause, e.g. {Goratrix},
    {Dancin' Dana}, {Tariq, The Silent}.

[^5-1-7]: [[ANK
    20220805-2]](https://usenet.krcg.org/t/vekn-79934/#m4)
    [[RBK influence-phase]](https://www.vekn.net/rulebook#influence-phase) — {Theo Bell}.

[^5-1-8]: [[RTR 19990712]](https://usenet.krcg.org/t/x5k5Vw_hkI0/#m0) [[LSJ
    20020115]](https://usenet.krcg.org/t/wG_tDLgfZso/#m1) — group "Action creating vampire" (G00054), {Creation Rites}.

[^5-1-13]: [[LSJ 20041201-2]](https://usenet.krcg.org/t/DY1OggZnMvs/#m1) — group "Action creating vampire" (G00054).

[^5-1-9]: [[LSJ 20100221]](https://usenet.krcg.org/t/TojSBPeGCFw/) [[LSJ
    20070928-2]](https://usenet.krcg.org/t/duRrP46XygI/#m43) [[LSJ
    20071001-2]](https://usenet.krcg.org/t/XoMeEYJw1ZA/#m10) — {Hatchling}, {Raw Recruit}, {Spell of Life}.

[^5-1-10]: [[LSJ 20050116]](https://usenet.krcg.org/t/yX5rqVaarvs/#m2) [[ANK
    20170226]](https://usenet.krcg.org/t/vekn-75625/#m3) [[LSJ
    20071005-1]](https://usenet.krcg.org/t/Ugcdb0ljZrU/#m1) — {Dual Form}, {Agent of Power}.

[^5-1-11]: [[LSJ 20100213]](https://usenet.krcg.org/t/vXDkYrTmkws/#m2) — {Soul Gem of Etrius}.

[^5-2-1]: [[ANK
    20220705]](https://usenet.krcg.org/t/vekn-79895/#m1)
    [[LSJ 20100705]](https://usenet.krcg.org/t/Dm_Zqyjdx8s/) [[LSJ 20001102]](https://usenet.krcg.org/t/LlPyLJjLdx0/#m2)
    [[TOM 19960214]](https://usenet.krcg.org/t/AP9ub18aVjg/#m6) — {Montano}, {Maris Streck}, {Paul Forrest, False
    Prophet}, {Sundown}, {Toby}, {Courier}, {Alexandra}.

[^5-2-2]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[LSJ
    19980210]](https://usenet.krcg.org/t/v7PxbzgVG0c/#m4) [[PIB
    20150720]](https://usenet.krcg.org/t/vekn-72088/#m4) — {Cloak the Gathering}, {Echo
    of Harmonies}, {Mouthpiece}, {Make an Example}.

[^5-2-3]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0).

[^5-2-4]: [[PIB
    20150412]](https://usenet.krcg.org/t/vekn-70519/#m2) —
    {Puppet Master}.

[^5-2-6]: [[LSJ 20060911]](https://usenet.krcg.org/t/oznDQgIFZ-I/#m3) [[PIB
    20151009]](https://usenet.krcg.org/t/vekn-73577/#m3) [[ANK
    20180702]](https://usenet.krcg.org/t/vekn-76770/#m1) [[LSJ
    20010622]](https://usenet.krcg.org/t/65IHHAii7ms/#m1) [[RBK
    unlock-phase]](https://www.vekn.net/rulebook#unlock-phase) — {Raphael Catarari}, {Baleful Doll}, {Vampiric Disease},
    {Parmenides}.

[^5-2-7]: [[LSJ 20080106]](https://usenet.krcg.org/t/q4s1P6ozsW8/#m1) — {Anarch Revolt}.

[^5-2-30]: [[ANK 20250121]](https://usenet.krcg.org/t/vekn-79080/#m7)
    — {Arika}, {Nightmares upon Nightmares}.

[^5-2-8]: [[PIB 20150111]](https://usenet.krcg.org/t/vekn-68580/#m3) [[LSJ
    20031014]](https://usenet.krcg.org/t/J8eZuZCZJUY/#m1) [[ANK
    20180917]](https://usenet.krcg.org/t/vekn-77011/#m3) [[LSJ
    20010828]](https://usenet.krcg.org/t/KoP_nqv-feM/#m1) — {Sensory Deprivation}, {Cry Wolf}, {Temptation}.

[^5-2-10]: [[PIB 20150703]](https://usenet.krcg.org/t/vekn-71918/#m5) — {The
    Sleeping Mind}.

[^5-2-31]: [[RTR 20250612]](https://www.blackchantry.com/2025/06/12/new-sabbat-rules-terms-and-errata/) — {Mind Numb}.

[^5-2-12]: [[LSJ 20010814-3]](https://usenet.krcg.org/t/4U6VYR9kBTA/#m1) [[LSJ
    20020109]](https://usenet.krcg.org/t/3aWGnESUYi0/#m1) [[RTR 20070707]](https://usenet.krcg.org/t/vSOt2c1uRzQ/#m0) —
    {Fata Amria}, {Putrefaction}.

[^5-2-13]: [[PIB 20121107]](https://usenet.krcg.org/t/vekn-40504/#m3) [[LSJ
    20070403]](https://usenet.krcg.org/t/TJ2ktt_1tjk/#m9) [[LSJ 20070413]](https://usenet.krcg.org/t/umdINigMKqs/#m19) —
    {Familial Bond}, {Champion}, {Discern}, {Donate}.

[^5-2-14]: [[LSJ 20010702]](https://usenet.krcg.org/t/jvIS3SDulqU/#m3) [[LSJ
    20011023-2]](https://usenet.krcg.org/t/47-PhTMiAOU/#m1) [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1)
    — {Starshell Grenade Launcher}, {Marciana Giovanni, Investigator}.

[^5-2-15]: [[LSJ 20070208]](https://usenet.krcg.org/t/ssk96EKyqjQ/#m1) [[LSJ
    20041022-2]](https://usenet.krcg.org/t/UAhJ3vWLyMM/#m2) [[LSJ 20100210]](https://usenet.krcg.org/t/HHJfsSbIEf0/#m11)
    — {No Secrets From the Magaji}, {Madness Network}, {Unleash Hell's Fury}.

[^5-2-19]: [[ANK 20190708]](https://usenet.krcg.org/t/vekn-77776/#m4) [[ANK
    20171017]](https://usenet.krcg.org/t/vekn-76233/#m4)
    [[ANK 20230305]](https://usenet.krcg.org/t/vekn-63821/#m40)
    [[LSJ 20080611]](https://usenet.krcg.org/t/gvb3uijtpZw/#m1) [[ANK
    20210131]](https://usenet.krcg.org/t/vekn-79008/#m3) — {Hard Case},
    {Faceless Night}, {Crypt's Sons}.

[^5-2-20]: [[ANK
    20191204]](https://usenet.krcg.org/t/vekn-78164/#m7)
    [[ANK 20200714]](https://usenet.krcg.org/t/vekn-78742/#m1) [[RTR
    19941006]](https://usenet.krcg.org/t/HUipqz0LFSw/#m0) [[LSJ 19991025]](https://usenet.krcg.org/t/R94tyTGJ6VQ/#m0) —
    {Mirror Walk}, {Ohoyo Hopoksia (Bastet)}, {Alexandra}.

[^5-2-21]: [[PIB
    20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) [[LSJ
    20010508]](https://usenet.krcg.org/t/mvfdiWb-edk/#m4) [[ANK
    20210322-2]](https://usenet.krcg.org/t/vekn-79080/#m6) [[LSJ
    19970707]](https://usenet.krcg.org/t/KWekwiRSa2I/#m1) — {Anarch Troublemaker}, {Brujah Debate}, {Nightmares upon
    Nightmares}, {Glutton}.

[^5-2-22]: [[ANK
    20170227]](https://usenet.krcg.org/t/vekn-75632/#m1) [[LSJ
    20020408]](https://usenet.krcg.org/t/4LXlqwmGQGc/#m1) [[ANK
    20221210-2]](https://usenet.krcg.org/t/vekn-80206/#m1)
    [[LSJ 20051211]](https://usenet.krcg.org/t/TuwXiJ8A9mo/#m1) — {Vampiric Disease}, {Eze, The Demon Prince}.

[^5-2-23]: [[LSJ 20060213]](https://usenet.krcg.org/t/ZZkDUiO9qOw/#m1) [[LSJ
    20060908]](https://usenet.krcg.org/t/CTy2GjM6-Dc/#m1) — {Banishment}, {Descent into Darkness}.

[^5-2-24]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[ANK
    20221102]](https://usenet.krcg.org/t/vekn-80129/#m1) [[ANK
    20180517]](https://usenet.krcg.org/t/vekn-76447/#m33) —
    {Disarming Presence}, {Mobile HQ, Operation Antigen}, {Gianna di Canneto}.

[^5-3-1]: [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[LSJ
    20040608]](https://usenet.krcg.org/t/p1uCHeAO7ak/#m1) [[LSJ 20030214]](https://usenet.krcg.org/t/A3U-Dy1yx8Y/#m1)
    [[LSJ 20020402]](https://usenet.krcg.org/t/eQCXHND3j5o/#m1) [[LSJ
    20050111]](https://usenet.krcg.org/t/lnW6nMIX-Vw/#m7) — {Decapitate}, {Ahriman's Demesne}, {Undead Persistence},
    {Rötschreck}, {Baltimore Purge}.

[^5-3-2]: [[LSJ 20021007]](https://usenet.krcg.org/t/wkyqKMka_F0/#m2) [[ANK
    20201122]](https://usenet.krcg.org/t/vekn-78861/#m6) [[ANK
    20200926]](https://usenet.krcg.org/t/vekn-78861/#m4) [[LSJ
    20100721]](https://usenet.krcg.org/t/Z9QNJ6SPIJM/#m1) [[ANK
    20210612]](https://usenet.krcg.org/t/vekn-79173/#m11) [[LSJ
    20020416-2]](https://usenet.krcg.org/t/tEC5uN8yqUE/#m4) — {Watenda}, {Revelation of Wrath}, {Orgy of Blood},
    {Garrote}.

[^5-3-3]: [[ANK
    20180104]](https://usenet.krcg.org/t/vekn-76356/#m4)
    — {Illusions of the Kindred}.

[^5-3-4]: [[LSJ 20020211]](https://usenet.krcg.org/t/ubqDaLeG3qo/#m2) [[LSJ
    20021122]](https://usenet.krcg.org/t/LieFYA_gyFo/#m11) [[LSJ 20090622-2]](https://usenet.krcg.org/t/1zt1SZb2TIk/#m1)
    [[RTR 20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[LSJ
    20020505]](https://usenet.krcg.org/t/ka9MajuWTAo/#m2) — {Ashes to Ashes}, {Coma}, {Entombment}, {Rowan Ring},
    {Rabbat, The Sewer Goddess}, {Seren Sukardi}, {Mummify}.

[^5-3-5]: [[LSJ 20030214]](https://usenet.krcg.org/t/A3U-Dy1yx8Y/#m1) [[ANK
    20190624]](https://usenet.krcg.org/t/vekn-77737/#m3) — {Undead
    Persistence}, {Undying Tenacity}.

[^5-3-6]: [[LSJ 20011022]](https://usenet.krcg.org/t/KMg4MwD-Jn0/#m1) [[LSJ
    20100902-2]](https://usenet.krcg.org/t/mFpx91METxM/#m1) [[ANK
    20191218]](https://usenet.krcg.org/t/vekn-62700/#m7) [[ANK
    20200420-1]](https://usenet.krcg.org/t/vekn-78576/#m1) [[LSJ
    20010326]](https://usenet.krcg.org/t/0Sy3xNbjYeU/#m2) [[LSJ 20010813-2]](https://usenet.krcg.org/t/zkKhvgZy9hA/#m2)
    — group "Ability usable in torpor" (G00027), {Nahir}, {Amelia, The Blood Red Tears}, {Alvaro, The Scion of
    Angelica}, {Lord Tremere}, {Wamukota}, {Marciana Giovanni, Investigator}.

[^5-3-8]: [[LSJ 20020416]](https://usenet.krcg.org/t/XS0Z0P5qaew/#m4) [[RBK
    torpor]](https://www.vekn.net/rulebook#torpor) — {Save Face}, {Martyr's Resilience}, {Bliss}, {Nosferatu
    Putrescence}.

[^5-3-9]: [[PIB 20150720]](https://usenet.krcg.org/t/vekn-72088/#m4) [[RTR
    19970306]](https://usenet.krcg.org/t/1dlmpgX6t14/#m0) — {Cats' Guidance}, {Make an Example}.

[^5-3-10]: [[LSJ 19970325]](https://usenet.krcg.org/t/R-t138q8688/#m3) [[RTR
    19950509]](https://usenet.krcg.org/t/_LKyR7pdMig/#m8) [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0)
    [[RTR 19970306]](https://usenet.krcg.org/t/1dlmpgX6t14/#m0) [[RBK torpor]](https://www.vekn.net/rulebook#torpor) —
    {The Kiss of Ra}, {Madness Network}, {Blood Brother Ambush}, {Make an Example}.

[^5-3-11]: [[RTR 19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0) [[LSJ
    20030130]](https://usenet.krcg.org/t/TUDO_4FwdyY/#m1) [[ANK
    20180906]](https://usenet.krcg.org/t/vekn-76981/#m2) [[ANK
    20181219]](https://usenet.krcg.org/t/vekn-77232/#m2) — group "Modifier after combat"
    (G00007), {Freak Drive}.

[^5-3-18]: Owner's reading, no official ruling: owner, 2026-11-11 — {Emergency Powers}, {Veneficti}, {Evil Eye},
    {Council of Seraphim}.

[^5-3-12]: [[LSJ 20091203]](https://usenet.krcg.org/t/JBdmMh1udN8/#m5) — {Emergency Powers}, {Legacy}, {Tattoo
    Signal}, {Deep Cover Agent}, {Unleash Hell's Fury}, {Veneficti}, {Barrenness}, {Minor Curse}, {Evil Eye}, {Strix}.
    [[LSJ 20091204]](https://usenet.krcg.org/t/JBdmMh1udN8/#m8) — {Council of Seraphim}, {The Red Question}. [[ANK
    20240610]](https://usenet.krcg.org/t/vekn-81521/#m12) — {High Orun}, {Consume the Dead}. [[RBK
    unlock-phase]](https://www.vekn.net/rulebook#unlock-phase).

[^5-3-17]: [[LSJ 20100203-2]](https://usenet.krcg.org/t/tKnhKlDo8lk/#m3) — group "Unique cards with burn option"
    (G00092).

[^5-3-13]: [[PIB 20150522]](https://usenet.krcg.org/t/vekn-71291/#m2)
    [[LSJ 20010610]](https://usenet.krcg.org/t/KVyVn-Y_UIY/#m2) [[ANK
    20210109]](https://usenet.krcg.org/t/vekn-78983/#m3) — {Blood Doll},
    {Vessel}, {Secure Haven}, {Fear of Mekhet}.

[^5-3-14]: [[RTR 19960124]](https://usenet.krcg.org/t/wF82VdVPlm0/#m0) [[PIB
    20150522]](https://usenet.krcg.org/t/vekn-71291/#m2) [[LSJ
    20010623]](https://usenet.krcg.org/t/GN5MqcCTOo8/#m1) — {Secure Haven}, {The Rack}.

[^5-3-15]: [[LSJ 20010119]](https://usenet.krcg.org/t/XB0IvK7I6PQ/#m2) [[RTR
    19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) — {Puppet Master}, {Temptation}, {Uriah Winter}.

[^5-3-16]: [[LSJ 20060409]](https://usenet.krcg.org/t/gsFQXsCGTG4/#m1) [[RTR
    19960112]](https://usenet.krcg.org/t/d3n3StNS7no/#m0) [[LSJ 20081016]](https://usenet.krcg.org/t/GizIs4K_HiY/#m2) —
    {Shadow Court Satyr}, {Legacy of Power}, {Abjure}.

[^5-4-1]: [[ANK
    20210813]](https://usenet.krcg.org/t/vekn-79279/#m3) [[PIB
    20130128-2]](https://usenet.krcg.org/t/vekn-44232/#m2) [[LSJ
    20100325-2]](https://usenet.krcg.org/t/aC6OOfaulbM/#m1) [[LSJ 20040616]](https://usenet.krcg.org/t/jQkkiC3I8P8/#m1)
    [[ANK 20230317]](https://usenet.krcg.org/t/vekn-76656/#m3)
    [[RBK important-terms-of-the-game]](https://www.vekn.net/rulebook#important-terms-of-the-game) [[RBK
    contested-cards]](https://www.vekn.net/rulebook#contested-cards) — {Blessed Resilience}, {Khazar's Diary (Endless
    Night)}, {Chain of Command}, {Unleash Hell's Fury}.

[^5-4-2]: [[LSJ 20090922]](https://usenet.krcg.org/t/UdvGbJqOeo4/#m15) [[ANK
    20220916]](https://usenet.krcg.org/t/vekn-80030/#m1)
    [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) — {Draught of the Soul}, {Soul Stealing}, {Taking the
    Skin: Minion}, {Blood Brother Ambush}, {Conscripted Statue}, {FBI Special Affairs Division}.

[^5-4-3]: [[ANK
    20220507]](https://usenet.krcg.org/t/vekn-79763/#m2)
    [[LSJ 20080608]](https://usenet.krcg.org/t/j5ShUmUt-vM/#m1) — {Brick Laying}, {Cryptic Mission}, {Succubus},
    {Consignment to Duat}, {Horseshoes}, {Jar the Soul}, {Keystone Kine}, {Smash and Grab}, {Abyssal Hunter}.

[^5-4-4]: [[RTR 19960124]](https://usenet.krcg.org/t/wF82VdVPlm0/#m0) [[LSJ
    20090922]](https://usenet.krcg.org/t/UdvGbJqOeo4/#m15) [[ANK
    20181022]](https://usenet.krcg.org/t/vekn-77103/#m6)
    [[ANK
    20220916]](https://usenet.krcg.org/t/vekn-80030/#m1) —
    {Taking the Skin: Minion}, {Soul Stealing}, {Draught of the Soul}.

[^5-4-5]: [[RTR 19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0) [[LSJ
    19990119]](https://usenet.krcg.org/t/Iez0G178MRM/#m1) [[LSJ 20080212]](https://usenet.krcg.org/t/Yg1nZfgkpGM/#m3) —
    {Anathema}, {Political Struggle}, {Young Bloods}.

[^5-4-6]: [[LSJ 20091217]](https://usenet.krcg.org/t/Xa8dbfnAbv0/#m37) [[LSJ
    20091217-2]](https://usenet.krcg.org/t/Xa8dbfnAbv0/#m32) [[LSJ
    20071003-2]](https://usenet.krcg.org/t/Tr9VQRVDbg0/#m1) [[ANK
    20230816]](https://usenet.krcg.org/t/vekn-80772/#m2) [[LSJ
    20070507]](https://usenet.krcg.org/t/FOLkbrSh0Ns/#m1) [[LSJ 20021205]](https://usenet.krcg.org/t/0MhU0auTJO4/#m5)
    [[PIB 20121028]](https://usenet.krcg.org/t/vekn-39816/#m10) —
    {Byzar}, {Charigger, The Axe}, {Heaven's Gate}, {Escaped Mental Patient}, {FBI Special Affairs Division}.

[^5-4-7]: [[ANK 20201122]](https://usenet.krcg.org/t/vekn-78861/#m6) [[ANK
    20200926]](https://usenet.krcg.org/t/vekn-78861/#m4) [[LSJ
    20100721]](https://usenet.krcg.org/t/Z9QNJ6SPIJM/#m1) [[LSJ 20070417]](https://usenet.krcg.org/t/ecDUqbSUsNg/#m1) —
    {Revelation of Wrath}, {Orgy of Blood}, {Reform Body}, {Abandoning the Flesh}, {Ashes to Ashes}.

[^5-5-1]: [[LSJ 20050216]](https://usenet.krcg.org/t/5_rUFgufFc4/#m2) [[ANK
    20190701]](https://usenet.krcg.org/t/vekn-77763/#m2) [[LSJ
    20030815]](https://usenet.krcg.org/t/sf1U7vnVE-o/#m3) [[ANK
    20230721]](https://usenet.krcg.org/t/vekn-80691/#m15)
    [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) — {Codex of the Edenic Groundskeepers}, {Hide the
    Heart}, {Tenebrous Form}, {Étienne Fauberge}, {The Shard, London}, {Kindred Segregation}.

[^5-5-2]: [[PIB 20121028]](https://usenet.krcg.org/t/vekn-39816/#m10)
    [[ANK
    20220818]](https://usenet.krcg.org/t/vekn-79972/#m6)
    [[ANK 20221102]](https://usenet.krcg.org/t/vekn-80129/#m1) [[RTR
    19971201]](https://usenet.krcg.org/t/zpIp5xOzQT0/#m0) [[RTR 19960112]](https://usenet.krcg.org/t/d3n3StNS7no/#m0) —
    {Taste of Vitae}, {Enhanced Coagulant}, {Shackles of Enkidu}, {Legacy of Power}.

[^5-5-3]: [[ANK 20171109]](https://usenet.krcg.org/t/vekn-76282/#m2) [[LSJ
    20030520-2]](https://usenet.krcg.org/t/GcymCHOJDVY/#m6) [[LSJ
    20051116-1]](https://usenet.krcg.org/t/MfGC7sJ8vh8/#m1) [[LSJ
    20081213-3]](https://usenet.krcg.org/t/cbZ2jl8-yGQ/#m1) [[LSJ
    20011214-1]](https://usenet.krcg.org/t/m2zsNI0McoE/#m1) [[LSJ 20070301]](https://usenet.krcg.org/t/-CeFWHQ2wXE/#m32)
    — {Nephandus}, {Ossian}, {Ghouled}, {Demdemeh}, {The Grandest Trick}.

[^5-5-4]: [[LSJ 20080717]](https://usenet.krcg.org/t/DMsE6V84GWI/#m1) [[RTR
    19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[RTR 20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) —
    {Shadow Court Satyr}, {Charming Lobby}, {Herald of Topheth}.

[^5-5-5]: [[LSJ 20050606]](https://usenet.krcg.org/t/IkZS3ikg7y8/#m1) [[ANK
    20200605]](https://usenet.krcg.org/t/vekn-78671/#m1)
    [[LSJ 20060124]](https://usenet.krcg.org/t/PBEBNFH61ik/#m1) [[LSJ
    20050224-1]](https://usenet.krcg.org/t/enxwleC-ZKw/#m27) [[RBK allies]](https://www.vekn.net/rulebook#allies) —
    {Veil of Darkness}, {The Line}, {Descent into Darkness}, {Sonja Blue}.

[^5-5-6]: [[ANK
    20180928-1]](https://usenet.krcg.org/t/vekn-77034/#m2)
    [[LSJ 20080512]](https://usenet.krcg.org/t/z2DGSFph6sM/#m17) — group "Allies who can play as a vampire" (G00011).

[^5-5-7]: [[RBK recruit-ally]](https://www.vekn.net/rulebook#recruit-ally) [[LSJ
    20100204]](https://usenet.krcg.org/t/o5Xnzc8G774/#m31) [[LSJ 20080426]](https://usenet.krcg.org/t/VIV521VtVfk/#m2) —
    {The Summoning}, {Piper}, group "Put card in play ignoring requirements" (G00110), {Khazar's Diary (Endless Night)}.

[^5-5-8]: [[ANK 20200813-3]](https://usenet.krcg.org/t/vekn-78800/#m2) [[ANK
    20200930]](https://usenet.krcg.org/t/vekn-78800/#m6) — {Nocturn}, {Infernal
    Servitor}.

[^5-5-9]: [[RTR 19961113]](https://usenet.krcg.org/t/VbMEQmJjI_w/#m0) [[ANK
    20170309-2]](https://usenet.krcg.org/t/vekn-75650/#m4) [[ANK
    20210913]](https://usenet.krcg.org/t/vekn-79322/#m1) [[LSJ
    20090116]](https://usenet.krcg.org/t/RQ3ARP9Kvfk/#m7) — groups "Allies with enter play effects" (G00150) and "Allies
    that burn an ally when entering play" (G00151), {Sébastien Goulet}.

[^5-6-1]: [[TOM 19951114]](https://usenet.krcg.org/t/LOEFFpprXKs/#m0) [[LSJ
    20080210]](https://usenet.krcg.org/t/nL-xqiydvYg/#m1) [[ANK
    20171003]](https://usenet.krcg.org/t/vekn-76205/#m1) [[RTR
    19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[RTR 19960112]](https://usenet.krcg.org/t/d3n3StNS7no/#m0) —
    {Ghoul Retainer}, {Jar of Skin Eaters}, {Spiritual Protector}, group "Retainer that strike" (G00029).

[^5-6-2]: [[ANK 20180612]](https://usenet.krcg.org/t/vekn-76717/#m2)
    [[ANK
    20211127]](https://usenet.krcg.org/t/vekn-76687/#m9)
    [[LSJ 20090616]](https://usenet.krcg.org/t/6DEbXHjKGhE/#m1) — group "Retainer doing damage" (G00016), {Bestial
    Vengeance}.

[^5-6-3]: [[LSJ 20090324]](https://usenet.krcg.org/t/Zc_ogoVhsug/#m8) [[LSJ
    20010503]](https://usenet.krcg.org/t/OjxDSCbB6i4/#m1) [[LSJ 19970224]](https://usenet.krcg.org/t/80KRDjVFkyg/#m1) —
    {Detect Authority}, {Shadow Twin}, {Camarilla Vitae Slave}.

[^5-6-4]: [[RBK recruit-ally]](https://www.vekn.net/rulebook#recruit-ally) — rulebook template.

[^5-6-5]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[ANK
    20180612]](https://usenet.krcg.org/t/vekn-76717/#m2) — group
    "Retainer with discipline level" (G00018); {Raven Spy}.

[^5-7-1]: [[ANK 20200203-2]](https://usenet.krcg.org/t/vekn-78415/#m1) [[LSJ
    20060515]](https://usenet.krcg.org/t/bzusnKlnYr8/#m3) [[ANK
    20230308]](https://usenet.krcg.org/t/vekn-80369/#m1)
    [[LSJ 20011210-4]](https://usenet.krcg.org/t/WT1LwCFnU9A/#m1) — group "Animal" (G00019), group "Ghoul" (G00044),
    {Murder of Crows}, {Camarilla Vitae Slave}, {Gargoyle Slave}.

[^5-7-2]: [[ANK 20200114]](https://usenet.krcg.org/t/vekn-78321/#m3) [[ANK
    20200115]](https://usenet.krcg.org/t/vekn-78321/#m6) [[ANK
    20171212]](https://usenet.krcg.org/t/vekn-76334/#m14) —
    {Momentary Delay}, {Shadow Boxing}, {Sniper Rifle}, {Obedience}.

[^5-7-3]: [[LSJ 20050114]](https://usenet.krcg.org/t/JWiZmyC2Y6s/#m3) [[LSJ
    20050228-2]](https://usenet.krcg.org/t/xBb9HQz2KPo/#m1) [[ANK
    20210322-2]](https://usenet.krcg.org/t/vekn-79080/#m6) [[RBK
    traits]](https://www.vekn.net/rulebook#traits) — group "Prevent normal unlock" (G00005), {Ruins of Charizel},
    {Nightmares upon Nightmares}.

[^5-7-4]: [[LSJ 20070322]](https://usenet.krcg.org/t/Ww-4rYJxi4w/#m1) [[ANK
    20180807]](https://usenet.krcg.org/t/vekn-76905/#m1) [[LSJ
    20070702]](https://usenet.krcg.org/t/Y8fFI0VCwfw/#m1) — group "Black Hand" (G00012), {Cadet}, {Mustajib}, {Seraph's
    Second}.

[^5-7-5]: [[ANK 20220528]](https://usenet.krcg.org/t/vekn-76455/#m19)
    [[LSJ 20060409]](https://usenet.krcg.org/t/gsFQXsCGTG4/#m1) [[RBK
    appendix-imbued-rules]](https://www.vekn.net/rulebook#appendix-imbued-rules) — {Keystone Kine}, {Kindred
    Segregation}.

[^5-7-6]: [[LSJ 20070516]](https://usenet.krcg.org/t/zalQ_AHTpAY/#m1) [[LSJ
    20060323]](https://usenet.krcg.org/t/taLKfNyQ3Kc/#m3) [[RBK
    appendix-imbued-rules]](https://www.vekn.net/rulebook#appendix-imbued-rules) — {Dreams of the Sphinx}, {Soul Gem of
    Etrius}, {Illusions of the Kindred}.

[^5-7-7]: [[LSJ 20060409-2]](https://usenet.krcg.org/t/lWdr_Ym-UIg/#m2) [[LSJ
    20081016]](https://usenet.krcg.org/t/GizIs4K_HiY/#m2) [[LSJ 20100211]](https://usenet.krcg.org/t/pL63VXEPGME/#m8)
    [[RBK appendix-imbued-rules]](https://www.vekn.net/rulebook#appendix-imbued-rules) — {Tension in the Ranks},
    {Abjure}, {Pressing Flesh}.

[^5-7-8]: [[LSJ 20011015]](https://usenet.krcg.org/t/ZXW0ScxTsBA/#m24) [[LSJ
    20030511]](https://usenet.krcg.org/t/qSKbZFF7F2A/#m2) [[LSJ 20060815]](https://usenet.krcg.org/t/nxExSHh-QjA/#m3)
    [[LSJ 20011212]](https://usenet.krcg.org/t/3jinbBVvqIE/#m2) [[LSJ
    20060509]](https://usenet.krcg.org/t/KBFnBRrOGB4/#m1) [[LSJ 20050322-2]](https://usenet.krcg.org/t/jiavUD9IgIA/#m5)
    — group "Scarce" (G00043), group "True Brujah" (G00041), {Clan Impersonation}, {Call the Great Beast}, {Earthshock}.

[^5-7-11]: [[RTR 20250612]](https://www.blackchantry.com/2025/06/12/new-sabbat-rules-terms-and-errata/).

[^5-7-12]: Owner's reading, no official ruling: [[rules-qa
    49.7]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-49--reign-of-stanislava-25th-anniversary-and-sabbat-paths-scoping-2026-10-05)
    — {Absolute Tyranny}.

[^5-7-10]: [[LSJ 20020609]](https://usenet.krcg.org/t/Q_Lk0DPHqC8/#m1) [[ANK
    20211113]](https://usenet.krcg.org/t/vekn-79465/#m2)
    [[LSJ 20050927]](https://usenet.krcg.org/t/oLMw6SSSgmA/#m13) [[LSJ
    20091012]](https://usenet.krcg.org/t/UWmuZnBo6FM/#m1) — group "Circle" (G00003), {Angelo}, {New Blood}, {Hermana
    Hambrienta Mayor}.

[^5-8-1]: [[LSJ 20010714]](https://usenet.krcg.org/t/-xrhpMEiMrw/#m10) [[RBK
    6-vampire-sects]](https://www.vekn.net/rulebook#6-vampire-sects) — {Ambrogino Giovanni}, {The Baron}, {Kemintiri},
    {Xaviar}, {Zayyat, The Sandstorm}, and the other crypt cards printing "N votes (titled)".

[^5-8-2]: [[ANK 20190211]](https://usenet.krcg.org/t/vekn-76865/#m15)
    [[RBK 8-glossaries]](https://www.vekn.net/rulebook#8-glossaries) — {Vlad Tepes}.

[^5-8-3]: [[ANK 20190407]](https://usenet.krcg.org/t/vekn-77533/#m2) — {Karsh}.

[^5-8-30]: Owner's reading, no official ruling: [[rules-qa
    3.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-3--p0-open-items) [[RBK
    contested-titles]](https://www.vekn.net/rulebook#contested-titles) — {Toreador Justicar}, {Madame Guil}.

[^5-8-4]: [[RTR 19950509]](https://usenet.krcg.org/t/_LKyR7pdMig/#m8) [[RBK
    contested-titles]](https://www.vekn.net/rulebook#contested-titles) — {Democritus}.

[^5-8-5]: [[LSJ 20070808-2]](https://usenet.krcg.org/t/gaFBeSrs7fU/#m0) [[RBK
    contested-titles]](https://www.vekn.net/rulebook#contested-titles) — group "Title providing action" (G00042).

[^5-8-6]: [[TOM 19960210]](https://usenet.krcg.org/t/PiOmH08RyVw/#m10) [[LSJ
    19980209]](https://usenet.krcg.org/t/8boERT-e5e4/#m0) [[LSJ 19970224]](https://usenet.krcg.org/t/80KRDjVFkyg/#m1) —
    group "Title providing capacity" (G00038); {The Treatment}, {Bloodbath}.

[^5-8-7]: [[LSJ 20100601]](https://usenet.krcg.org/t/R7YgwD0VlUQ/#m1) [[ANK
    20190725]](https://usenet.krcg.org/t/vekn-77813/#m1) — {Xeper, Sultan of Lepers},
    {Gerald Windham}.

[^5-8-8]: [[LSJ 20010809-3]](https://usenet.krcg.org/t/gLl8F0zcCF0/#m2) [[LSJ
    20010809-2]](https://usenet.krcg.org/t/9ggmJcK2De0/#m10) [[TOM
    19960210]](https://usenet.krcg.org/t/PiOmH08RyVw/#m10) [[RTR 20000501]](https://usenet.krcg.org/t/MKrA0hBXuaU/#m0)
    [[LSJ 20060908]](https://usenet.krcg.org/t/CTy2GjM6-Dc/#m1) — {Banishment}, {Descent into Darkness}.

[^5-8-25]: [[RBK 6-vampire-sects]](https://www.vekn.net/rulebook/6-vampire-sects) — {Toreador Justicar}.

[^5-8-21]: [[LSJ 20030419]](https://usenet.krcg.org/t/A0mvllC-tgs/#m5) — {Illusions of the Kindred}.

[^5-8-9]: [[LSJ 20080602]](https://usenet.krcg.org/t/Y7CLzywq1Lk/#m1) [[LSJ
    20060904]](https://usenet.krcg.org/t/L9D4GK0yNv8/#m4) [[RBK
    6-vampire-sects]](https://www.vekn.net/rulebook#6-vampire-sects) — {No Confidence}, {Field Training}, {Go Anarch},
    and the other sect-change cards carrying this ruling.

[^5-8-10]: [[RTR 20201130]](https://www.blackchantry.com/2020/12/25/rtr-30-11-2020/) [[LSJ
    20060904]](https://usenet.krcg.org/t/L9D4GK0yNv8/#m4) — {Clan Impersonation}, {Derange}.

[^5-8-11]: [[LSJ 20060904]](https://usenet.krcg.org/t/L9D4GK0yNv8/#m4) — {Fall of the Camarilla}, {Fall of the Sabbat}.

[^5-8-12]: [[LSJ 20040519]](https://usenet.krcg.org/t/fz-EAPmmqZY/#m10) — {Gratiano}.

[^5-8-13]: [[LSJ 20030202]](https://usenet.krcg.org/t/ox7A8EvaNJo/#m3) — {Horatio Ballard}, {Maxwell}.

[^5-8-15]: [[LSJ 20050128]](https://usenet.krcg.org/t/HVy8iPUxNbI/#m48) [[LSJ
    20080603]](https://usenet.krcg.org/t/9usl4idp-pY/#m1) — group "Require a Baron" (G00037); {CrimethInc.}, {Powerbase:
    Los Angeles}.

[^5-8-16]: [[PIB 20151116]](https://usenet.krcg.org/t/vekn-74317/#m2) — {The Baron},
    {Baron Dieudonne}.

[^5-8-17]: [[ANK
    20220317]](https://usenet.krcg.org/t/vekn-79706/#m2) —
    {Club Illusion}.

[^5-8-26]: [[ANK 20250320]](https://usenet.krcg.org/t/vekn-82172/#m4) —
    {Garibaldi-Meucci Museum}.

[^5-8-28]: [[RBK 6-vampire-sects]](https://www.vekn.net/rulebook/6-vampire-sects) [[RBK
    other-vampire-sects]](https://www.vekn.net/rulebook#other-vampire-sects).

[^5-8-27]: Owner's reading, no official ruling: [[rules-qa
    112.21]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-112--round-27-secrets-of-new-york-and-noites-brasileiras-63-cards-the-cards-board-owner-2026-11-08-red-11214-11219-11221-11247-claude-2026-11-09-the-owner-11214-11219-11221-11247-confirmed)
    — {Papillon}, {Mel, a Guerreira Vibrante}.

[^5-8-22]: [[LSJ 20081120-1]](https://usenet.krcg.org/t/e2PNDpg-l_c/#m19) [[LSJ
    20091016]](https://usenet.krcg.org/t/pqa7mYZ6NEM/#m21) [[LSJ 20100226]](https://usenet.krcg.org/t/JnycCGrNQmY/#m3) —
    {Vidal Jarbeaux}.

[^5-8-18]: [[LSJ 20050128]](https://usenet.krcg.org/t/HVy8iPUxNbI/#m48) [[PIB
    20150306]](https://usenet.krcg.org/t/vekn-69627/#m12) [[PIB
    20150307]](https://usenet.krcg.org/t/vekn-69627/#m20) [[LSJ
    20050526]](https://usenet.krcg.org/t/x98guZkL_CQ/#m1) — {Vidal Jarbeaux}, {Vlad Tepes}, {Kemintiri}, {Mata Hari}.

[^5-8-23]: [[ANK
    20211019]](https://usenet.krcg.org/t/vekn-79278/#m3) [[ANK
    20200710]](https://usenet.krcg.org/t/vekn-77985/#m3) — {Vidal Jarbeaux}.

[^5-8-19]: [[LSJ 20050124]](https://usenet.krcg.org/t/HVy8iPUxNbI/#m13) [[ANK
    20230503]](https://usenet.krcg.org/t/vekn-80486/#m7)
    — group "Impersonating title for non-existent sect" (G00070).

[^5-8-29]: [[ANK 20250410]](https://usenet.krcg.org/t/vekn-82214/#m1) — {Karl Schrekt}.

[^5-8-20]: [[ANK 20231221]](https://usenet.krcg.org/t/vekn-81065/#m4) — group
    "Impersonating a title for political action" (G00069).

[^5-8-24]: [[LSJ 20091015]](https://usenet.krcg.org/t/pqa7mYZ6NEM/#m9) [[PIB
    20130128]](https://usenet.krcg.org/t/vekn-43572/#m36)
    [[LSJ 20050721]](https://usenet.krcg.org/t/g39H3dwXqvc/#m20) — {Vidal Jarbeaux}, {Kemintiri}.

[^5-9-1]: [[LSJ 20100811]](https://usenet.krcg.org/t/kGKWcs3k6vI/#m25) — {Go Anarch}, {Writ of Acceptance}, {Galaric's
    Legacy}, {Into the Fire}, {Out of the Frying Pan}; the card text of {Field Training}.

[^5-9-2]: [[LSJ 20070707]](https://usenet.krcg.org/t/ZtRk5z2TcoI/#m1) [[RTR
    20201130]](https://www.blackchantry.com/2020/12/25/rtr-30-11-2020/) [[ANK
    20200415]](https://usenet.krcg.org/t/vekn-7516/#m9) —
    {Derange}, {Clan Impersonation}.

[^5-9-3]: [[LSJ 20021209]](https://usenet.krcg.org/t/kQpR-Cn9nW0/#m1) — {Clan Impersonation}, {Derange}. [[ANK
    20211022]](https://usenet.krcg.org/t/vekn-79422/#m2) — {Deep Song}, {Nar-Sheptha}. [[ANK
    20230814]](https://usenet.krcg.org/t/vekn-80752/#m12) —
    {Powerbase: Savannah}.

[^5-9-4]: [[LSJ 20100811]](https://usenet.krcg.org/t/kGKWcs3k6vI/#m25) [[ANK
    20190619]](https://usenet.krcg.org/t/vekn-77723/#m1) —
    {Writ of Acceptance}. [[ANK
    20180626-2]](https://usenet.krcg.org/t/vekn-76752/#m5) —
    {Tegyrius, Vizier}.

[^5-9-5]: [[LSJ 20040519]](https://usenet.krcg.org/t/fz-EAPmmqZY/#m10) [[LSJ
    20080510]](https://usenet.krcg.org/t/67g8kq3F_uw/#m6) [[LSJ 20100811]](https://usenet.krcg.org/t/kGKWcs3k6vI/#m25)
    [[ANK
    20190619]](https://usenet.krcg.org/t/vekn-77723/#m1) —
    {Fall of the Camarilla}, {Fall of the Sabbat}.

[^5-9-6]: [[ANK
    20210124]](https://usenet.krcg.org/t/vekn-79001/#m2) — {The
    Red Question}. [[ANK 20190416]](https://usenet.krcg.org/t/vekn-77560/#m3) —
    {Ministry}, {Protection Racket}, {Teresita, The Godmother}. [[LSJ
    20090216]](https://usenet.krcg.org/t/wU4YAwE6wlI/#m2) [[ANK
    20180922]](https://usenet.krcg.org/t/vekn-77023/#m4) — {Warsaw
    Station}.

[^5-9-7]: [[LSJ 20050124]](https://usenet.krcg.org/t/HVy8iPUxNbI/#m13) [[ANK
    20230503]](https://usenet.krcg.org/t/vekn-80486/#m7)
    — {Fall of the Camarilla}, {Fall of the Sabbat}. [[ANK
    20180125]](https://usenet.krcg.org/t/vekn-76385/#m1) — {The Red Question}.

[^5-9-8]: [[RBK allies]](https://www.vekn.net/rulebook#allies) [[ANK
    20230417]](https://usenet.krcg.org/t/vekn-80399/#m4)
    — {Grey Thorne}, {Vivienne Géroux}.

[^6-1-1]: [[RTR 19970425]](https://usenet.krcg.org/t/DhP_l2cX3mQ/#m0) [[LSJ
    20060921]](https://usenet.krcg.org/t/Lj1GDLBmOWQ/#m5) [[PIB
    20130128-2]](https://usenet.krcg.org/t/vekn-44232/#m2) [[LSJ
    20100325-2]](https://usenet.krcg.org/t/aC6OOfaulbM/#m1) [[LSJ 20001031]](https://usenet.krcg.org/t/y8LNZhRyXO0/#m2)
    [[PIB 20121220]](https://usenet.krcg.org/t/vekn-43188/#m8) [[ANK
    20200417]](https://usenet.krcg.org/t/vekn-78568/#m16) —
    {Agaitas, The Scholar of Antiquities}, {Khazar's Diary (Endless Night)}, {Storage Annex}, {The Capuchin}.

[^6-1-3]: [[LSJ 20030701-2]](https://usenet.krcg.org/t/EFBCP5QI3D4/#m3) [[ANK
    20221102]](https://usenet.krcg.org/t/vekn-80129/#m1) — {Echo of
    Harmonies}, {Empowering the Puppet King}, {Set's Call}, {The Shard, London}.

[^6-1-4]: [[LSJ 20100423]](https://usenet.krcg.org/t/YnQCu0GeMhc/#m1) [[ANK
    20210813]](https://usenet.krcg.org/t/vekn-79279/#m3) [[LSJ
    20090428]](https://usenet.krcg.org/t/BmA4xsXoEXc/#m1) — {Byzar}, {Blessed Resilience}, {Spell of Life}.

[^6-2-1]: [[LSJ 20001201]](https://usenet.krcg.org/t/BPHZsY3p20E/#m2) [[ANK
    20200629-2]](https://usenet.krcg.org/t/vekn-78701/#m2)
    [[LSJ 20091208]](https://usenet.krcg.org/t/ptHbJM9MlVI/#m1) [[RTR
    19951017]](https://usenet.krcg.org/t/ouhNUbHYg50/#m2) [[ANK
    20200925]](https://usenet.krcg.org/t/vekn-78861/#m1) [[ANK
    20211105]](https://usenet.krcg.org/t/vekn-79460/#m1) —
    {Puppet Master}, {Malkavian Dementia}, {Velvet Tongue}, {Sarrasine}.

[^6-2-2]: [[LSJ 20020618]](https://usenet.krcg.org/t/DVGe6EsiMZ4/#m7) [[LSJ
    20080326]](https://usenet.krcg.org/t/KUQEznVQlOU/#m3) [[PIB
    20150501]](https://usenet.krcg.org/t/vekn-70780/#m6) [[ANK
    20220127]](https://usenet.krcg.org/t/vekn-79615/#m1) —
    {Temptation}, {Revelation of Despair}.

[^6-2-3]: [[LSJ 20020725]](https://usenet.krcg.org/t/wCPFIH_g5ZE/#m4) [[PIB
    20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) [[LSJ
    20070222-2]](https://usenet.krcg.org/t/jBrfK77gayo/#m5) — {Temptation}, {War Ghoul}, {The Diamond Thunderbolt}; PIB
    20121031's "must then be paid" is superseded on payment by [[ANK
    20210124]](https://usenet.krcg.org/t/vekn-79001/#m2).

[^6-2-4]: [[LSJ 20010119]](https://usenet.krcg.org/t/XB0IvK7I6PQ/#m2) [[RTR
    20020501]](https://usenet.krcg.org/t/M1snoR2msbQ/#m0) [[LSJ 20030219]](https://usenet.krcg.org/t/ugfckv9DAbo/#m5)
    [[LSJ 20040526]](https://usenet.krcg.org/t/M3hDKK2JSWw/#m17) [[RTR
    20000501]](https://usenet.krcg.org/t/MKrA0hBXuaU/#m0) [[ANK
    20210630]](https://usenet.krcg.org/t/vekn-79205/#m3) — {Temptation}, {Puppet
    Master}, {Spirit Marionette}, {Descent into Darkness}, {Lay Low}.

[^6-2-5]: [[LSJ 20081105]](https://usenet.krcg.org/t/CYjOJtTBMGU/#m1) [[ANK
    20181110]](https://usenet.krcg.org/t/vekn-76844/#m8) — group "Temporary
    control" (G00001), {The Ailing Spirit}, {Temptation}, {Puppet Master}, {Spirit Marionette}.

[^6-2-6]: [[LSJ 20021111]](https://usenet.krcg.org/t/Y6_dPn4ONCA/#m2) [[RBK
    5-ending-the-game]](https://www.vekn.net/rulebook/5-ending-the-game) — {Parmenides}.

[^6-2-7]: [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) [[ANK
    20200130-2]](https://usenet.krcg.org/t/vekn-78398/#m2) [[LSJ
    20031106]](https://usenet.krcg.org/t/bFZCLXzzOeM/#m31) [[ANK
    20220525]](https://usenet.krcg.org/t/vekn-79777/#m5) —
    {Betrayer}, {The Rack}, {Rutor's Hand}, {Imposing Phantasm}.

[^6-3-1]: [[TOM 19960226-1]](https://usenet.krcg.org/t/16y-ZD7Xats/#m1) — recorded identically on fifteen cards; {Far
    Mastery}, {New Management}, {Spirit Marionette}.

[^6-3-2]: [[RTR 19960112]](https://usenet.krcg.org/t/d3n3StNS7no/#m0) [[LSJ
    19971002]](https://usenet.krcg.org/t/RY_nhdykKP0/#m1) [[LSJ 20030606]](https://usenet.krcg.org/t/qzKS4OGBgbo/#m4) —
    group "Locquipments" (G00047), {Disputed Territory}, {Dominate Kine}, {Malkavian Time Auction}.

[^6-3-3]: [[LSJ 19971002]](https://usenet.krcg.org/t/RY_nhdykKP0/#m1) [[LSJ
    20080803-2]](https://usenet.krcg.org/t/OknxGrvlaNk/#m2) — {Disputed Territory}, {New Management}.

[^6-3-7]: [[RBK 8-glossaries]](https://www.vekn.net/rulebook#8-glossaries) — glossary "Steal (a card)".

[^6-3-4]: [[RTR 19960708]](https://usenet.krcg.org/t/Rd3Zb61ELuw/#m0) [[LSJ
    20001031]](https://usenet.krcg.org/t/y8LNZhRyXO0/#m2) [[PIB
    20121220]](https://usenet.krcg.org/t/vekn-43188/#m8) [[LSJ
    19990709]](https://usenet.krcg.org/t/F131j-4dabU/#m2) — {Shackles of Enkidu}, {Storage Annex}, {Ethan Locke}.

[^6-3-5]: [[TOM 19960114]](https://usenet.krcg.org/t/cOqrGO0UrSw/#m1) [[LSJ
    20020514]](https://usenet.krcg.org/t/xBDPee5wq40/#m3) — {Incriminating Videotape}.

[^6-3-6]: [[LSJ 20100723]](https://usenet.krcg.org/t/0u5KQWiutdg/#m1) [[PIB
    20111002]](https://usenet.krcg.org/t/vekn-8235/#m19) [[ANK
    20200508-1]](https://usenet.krcg.org/t/vekn-78622/#m17) — {The
    Coven}, {Scourge of the Enochians}.

[^6-4-1]: [[RTR 20000501]](https://usenet.krcg.org/t/MKrA0hBXuaU/#m0) [[LSJ
    20060908]](https://usenet.krcg.org/t/CTy2GjM6-Dc/#m1) [[ANK
    20210630]](https://usenet.krcg.org/t/vekn-79205/#m3) [[LSJ
    20080816]](https://usenet.krcg.org/t/FKBTEzLf0_A/#m5) [[LSJ 20100206-2]](https://usenet.krcg.org/t/reXyybyIYX8/#m1)
    [[TOM 19951209]](https://usenet.krcg.org/t/qP2j6CpBUDI/#m6) [[LSJ
    20010809-2]](https://usenet.krcg.org/t/9ggmJcK2De0/#m10) [[LSJ
    20010809-3]](https://usenet.krcg.org/t/gLl8F0zcCF0/#m2) — {Descent into Darkness}, {Lay Low}, {Thicker than Blood},
    {The Rack}.

[^6-4-9]: Owner's reading, no official ruling: [[rules-qa
    3.1]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-3--p0-open-items) — {Descent
    into Darkness}.

[^6-4-2]: [[LSJ 20040526]](https://usenet.krcg.org/t/M3hDKK2JSWw/#m17) [[LSJ
    20021111]](https://usenet.krcg.org/t/Y6_dPn4ONCA/#m2) — {Descent into Darkness}, {Parmenides}.

[^6-4-3]: [[LSJ 20040726]](https://usenet.krcg.org/t/LlqCB6LN64g/#m7) [[LSJ
    20010627]](https://usenet.krcg.org/t/NhNCVCCDyU0/#m5) — {Possession}, {Compel the Spirit}.

[^6-4-4]: [[TOM 19960210]](https://usenet.krcg.org/t/PiOmH08RyVw/#m10) [[TOM
    19960211]](https://usenet.krcg.org/t/PiOmH08RyVw/#m14) [[LSJ
    20010809-2]](https://usenet.krcg.org/t/9ggmJcK2De0/#m10) [[LSJ
    20010809-3]](https://usenet.krcg.org/t/gLl8F0zcCF0/#m2) [[PIB
    20150522]](https://usenet.krcg.org/t/vekn-71291/#m2) [[LSJ
    20010623]](https://usenet.krcg.org/t/GN5MqcCTOo8/#m1) [[LSJ 20040616]](https://usenet.krcg.org/t/jQkkiC3I8P8/#m1) —
    {The Rack}, {Descent into Darkness}, {Lay Low}.

[^6-4-5]: [[LSJ 20030522-1]](https://usenet.krcg.org/t/_krZG-uPtzc/#m19) [[LSJ
    20061218]](https://usenet.krcg.org/t/DTBI6LkPdZ4/#m1) — {NRA PAC}.

[^6-4-6]: [[PIB
    20150512]](https://usenet.krcg.org/t/vekn-71020/#m2)
    [[ANK
    20180129-2]](https://usenet.krcg.org/t/vekn-76363/#m4)
    — {The Black Throne}, {Priority Contract}, {Provision of the Silsila}.

[^6-4-8]: [[LSJ 20010622]](https://usenet.krcg.org/t/65IHHAii7ms/#m1) [[LSJ
    20070829-2]](https://usenet.krcg.org/t/drn7wHaGugQ/#m3) — {Parmenides}, {Sonja Blue}.

[^6-5-1]: [[ANK
    20180719-2]](https://usenet.krcg.org/t/vekn-76835/#m1)
    [[ANK 20181105]](https://usenet.krcg.org/t/vekn-77140/#m1) [[LSJ
    19970224]](https://usenet.krcg.org/t/80KRDjVFkyg/#m1) [[RBK
    5-ending-the-game]](https://www.vekn.net/rulebook/5-ending-the-game) — {Poison Pill}, {The Rising}, {Dirty Little
    Secrets}.

[^6-5-2]: [[ANK 20180813]](https://usenet.krcg.org/t/vekn-76917/#m2) [[LSJ
    20041025]](https://usenet.krcg.org/t/sbfkGmojYao/#m4) — {Poison Pill}, {Ancient Influence}, {Reins of Power}.

[^6-5-4]: [[RTR 20010711]](https://usenet.krcg.org/t/GsI1UyH54jU/#m0) [[RTR
    19970630]](https://usenet.krcg.org/t/KireUeOYY3c/#m1) [[ANK
    20230817-2]](https://usenet.krcg.org/t/vekn-80779/#m3) [[LSJ
    20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) [[RTR 19941109]](https://usenet.krcg.org/t/_6CXoKTSLnw/#m0) —
    {Thanks for the Donation}, {Life Boon}, {Parity Shift}, {Repo Man}, {Vast Wealth}.

[^6-5-5]: [[LSJ 20080512]](https://usenet.krcg.org/t/z2DGSFph6sM/#m17) [[LSJ
    20050607]](https://usenet.krcg.org/t/WLv9R8wA0Ow/#m8) [[LSJ 20050608]](https://usenet.krcg.org/t/WLv9R8wA0Ow/#m10) —
    {Herald of Topheth}, {Bima}, {Spying Mission}.

[^6-5-13]: [[LSJ 20081105]](https://usenet.krcg.org/t/CYjOJtTBMGU/#m1) — {Ferraille}.

[^6-5-7]: [[LSJ 20090731]](https://usenet.krcg.org/t/y6f0s6tUtqs/#m6) — {Last Stand}.

[^6-5-8]: [[ANK 20180408]](https://usenet.krcg.org/t/vekn-76500/#m2) [[LSJ
    20021008]](https://usenet.krcg.org/t/Mc3xfym_uw8/#m2) [[ANK
    20181110]](https://usenet.krcg.org/t/vekn-76844/#m8) [[ANK
    20200408-2]](https://usenet.krcg.org/t/vekn-78562/#m10) [[ANK
    20210630]](https://usenet.krcg.org/t/vekn-79205/#m3) — {Charnas the Imp}, {The
    Meddling of Semsith}, {Puppet Master}, {Temptation}, {Lay Low}.

[^6-5-9]: [[LSJ 20100206-2]](https://usenet.krcg.org/t/reXyybyIYX8/#m1) [[ANK
    20180129-2]](https://usenet.krcg.org/t/vekn-76363/#m4)
    [[ANK
    20220210]](https://usenet.krcg.org/t/vekn-79644/#m2)
    — {Shatter the Gate}, {Priority Contract}, {The Black Throne}, {Revelation of the Serpent}.

[^6-5-10]: [[LSJ 20000309]](https://usenet.krcg.org/t/ykAzsCzPkvg/#m2) [[LSJ
    20100208]](https://usenet.krcg.org/t/wt2kuUJXoRI/#m9) [[TOM 19951214-1]](https://usenet.krcg.org/t/8Fr6gioYZDI/#m1)
    [[RBK 5-ending-the-game]](https://www.vekn.net/rulebook/5-ending-the-game) — group "Vote damaging multiple players"
    (G00015), {The Rising}, {Sabbat Threat}.

[^6-5-11]: [[LSJ 20100206-2]](https://usenet.krcg.org/t/reXyybyIYX8/#m1) [[LSJ
    20100210-2]](https://usenet.krcg.org/t/wt2kuUJXoRI/#m12) [[LSJ 20100207]](https://usenet.krcg.org/t/wt2kuUJXoRI/#m4)
    — {The Rising}.

[^6-5-12]: [[LSJ 20011222]](https://usenet.krcg.org/t/DlCBJmB2fzY/#m7) [[ANK
    20221028-2]](https://usenet.krcg.org/t/vekn-80122/#m1) [[RTR
    19980623]](https://usenet.krcg.org/t/tSpd9dtTElc/#m0) [[LSJ 20030519]](https://usenet.krcg.org/t/E6Jz8m3iKrA/#m3)
    [[LSJ 20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) [[TOM
    19951214-2]](https://usenet.krcg.org/t/lVTQ1rstpLY/#m1) [[LSJ 19971006]](https://usenet.krcg.org/t/2to9jkow98Q/#m1)
    — {Sargon}, {Tereza Rostas}, {Hrothulf}, {Enticement}, {Curse of Nitocris}.

[^6-6-1]: [[LSJ 20070309-2]](https://usenet.krcg.org/t/c3Kx0-C_5JU/#m1) [[RBK
    master-cards]](https://www.vekn.net/rulebook#master-cards) [[RBK
    master-phase]](https://www.vekn.net/rulebook#master-phase) — {Wash}. The rulebook is explicit that a canceled
    out-of-turn master still counts against the next master phase; the older ruling's parenthetical to the contrary is
    superseded.

[^6-6-2]: [[LSJ 20090626]](https://usenet.krcg.org/t/3Cd-DNfT7yQ/#m4) [[ANK
    20170124]](https://usenet.krcg.org/t/vekn-75556/#m2) [[RBK
    trifle]](https://www.vekn.net/rulebook#trifle) — {Wash}, {Sudden Reversal}.

[^6-6-8]: [[RBK contested-cards]](https://www.vekn.net/rulebook#contested-cards) — group "Unique Trifle" (G00082).

[^6-6-3]: [[ANK 20170124]](https://usenet.krcg.org/t/vekn-75556/#m2) [[LSJ
    20090626]](https://usenet.krcg.org/t/3Cd-DNfT7yQ/#m4) — {Wash}.

[^6-6-4]: [[PIB 20150524]](https://usenet.krcg.org/t/vekn-71300/#m3) [[ANK
    20191004]](https://usenet.krcg.org/t/vekn-77994/#m1) [[RBK
    master-cards]](https://www.vekn.net/rulebook#master-cards) — {Proxy Kissed}, {Synesios}.

[^6-6-5]: [[LSJ 20031201]](https://usenet.krcg.org/t/Mi_j7sUsZZw/#m1) — {Vox Senis}.

[^6-6-7]: [[ANK
    20220503]](https://usenet.krcg.org/t/vekn-39040/#m5) —
    {Ex Nihilo}.

[^6-7-1]: [[LSJ 19970224]](https://usenet.krcg.org/t/80KRDjVFkyg/#m1) [[LSJ
    20070516]](https://usenet.krcg.org/t/zalQ_AHTpAY/#m1) [[RBK
    targeting-of-cards]](https://www.vekn.net/rulebook#targeting-of-cards) [[RBK
    card-rulings]](https://www.vekn.net/rulebook#card-rulings) — {Reform Body}, {Dreams of the Sphinx}.

[^6-7-3]: [[ANK 20200408-2]](https://usenet.krcg.org/t/vekn-78562/#m10) [[RTR
    20000501]](https://usenet.krcg.org/t/MKrA0hBXuaU/#m0) [[ANK
    20210630]](https://usenet.krcg.org/t/vekn-79205/#m3) — {Banishment}, {Lay Low}.

[^6-7-4]: [[TOM 19951209]](https://usenet.krcg.org/t/qP2j6CpBUDI/#m6) [[ANK
    20210630]](https://usenet.krcg.org/t/vekn-79205/#m3) [[LSJ
    20100902]](https://usenet.krcg.org/t/SM2_578Th0U/#m17) [[RTR
    20180303]](https://usenet.krcg.org/t/vekn-76447/#m0) [[RBK
    influence-phase]](https://www.vekn.net/rulebook#influence-phase) — {Lay Low}, {Wormwood}, {Jimmy Dunn}.

[^6-7-5]: [[RTR 20180303]](https://usenet.krcg.org/t/vekn-76447/#m0)
    [[ANK 20240118]](https://usenet.krcg.org/t/vekn-81155/#m1) [[LSJ
    20100527]](https://usenet.krcg.org/t/Tr3X4DsZr-8/#m1) [[RBK
    influence-phase]](https://www.vekn.net/rulebook#influence-phase) — {Ingrid Rossler}, {Angela Preston}, {Paul
    "Sixofswords29" Moreton}, {Ennoia's Theater}, {Gather}, {Leandro}.

[^6-7-6]: [[LSJ 20060623]](https://usenet.krcg.org/t/mfgW0TeoLNM/#m1) [[LSJ
    19990215]](https://usenet.krcg.org/t/_izcAo43T4Q/#m2) [[ANK
    20240706]](https://usenet.krcg.org/t/vekn-81563/#m19) [[LSJ
    20041015]](https://usenet.krcg.org/t/jonKzp3f8wA/#m0) — {Undue Influence}, {Break the Bonds}, {Lázár Dobrescu},
    {Social Ladder}.

[^6-8-1]: [[ANK
    20180318]](https://usenet.krcg.org/t/vekn-76464/#m1)
    [[ANK 20231229]](https://usenet.krcg.org/t/vekn-81077/#m5) [[ANK
    20190606-2]](https://usenet.krcg.org/t/vekn-77694/#m1) [[LSJ
    20030530]](https://usenet.krcg.org/t/SZehI8SwAc4/#m21) — {Dreams of the Sphinx}, {The Meddling of Semsith}, {Hagar
    Stone}, {Sascha Vykos, The Angel of Caine}, {Raptor}.

[^6-8-2]: [[LSJ 20020814]](https://usenet.krcg.org/t/gt8wQhk76lA/#m1) [[LSJ
    20021008]](https://usenet.krcg.org/t/Mc3xfym_uw8/#m2) [[ANK
    20191218]](https://usenet.krcg.org/t/vekn-62700/#m7) [[ANK
    20200616]](https://usenet.krcg.org/t/vekn-78687/#m1)
    — {Edward Neally}, {The Meddling of Semsith}, {Nahir}; group "Can draw during action" (G00023).

[^6-8-3]: [[LSJ 20100812]](https://usenet.krcg.org/t/C4PUHWr_xSQ/#m1) [[LSJ
    20020904-2]](https://usenet.krcg.org/t/ObuKimgcCpI/#m10) [[ANK
    20210313]](https://usenet.krcg.org/t/vekn-79072/#m1)
    [[ANK 20190725]](https://usenet.krcg.org/t/vekn-77813/#m1) [[LSJ
    20090731]](https://usenet.krcg.org/t/y6f0s6tUtqs/#m6) [[LSJ 20060426]](https://usenet.krcg.org/t/4e6z1_JWIzA/#m1) —
    {Dreams of the Sphinx}, {The Art of Love}, {Josef von Bauren}, {Last Stand}, {High Aye}.

[^6-8-4]: [[LSJ 19981005]](https://usenet.krcg.org/t/9vOWIR4P_4o/#m1) [[ANK
    20230404]](https://usenet.krcg.org/t/vekn-80431/#m2) [[LSJ
    20090618]](https://usenet.krcg.org/t/jxTaf6lnYb0/#m1) [[LSJ 20060215]](https://usenet.krcg.org/t/BhecJx5BqtQ/#m10)
    [[ANK 20180725]](https://usenet.krcg.org/t/vekn-76858/#m1) [[RTR
    20030519]](https://usenet.krcg.org/t/NOBWXWrd-vA/#m0) [[ANK
    20211010]](https://usenet.krcg.org/t/vekn-79335/#m8) [[ANK
    20180925-1]](https://usenet.krcg.org/t/vekn-77029/#m6)
    [[LSJ 20100506]](https://usenet.krcg.org/t/H6Jm74N0v7k/#m1) — {Ruxandra}, {Constant Revolution}, {Call the Wild
    Hunt}, {Feline Saboteur}, {Rachel Brandywine}, {Jaggedy Andy}, {Angelica, The Canonicus}, {Infernal Pursuit}.

[^6-8-5]: [[TOM 19950924]](https://usenet.krcg.org/t/yejj-744_zc/#m8) [[ANK
    20180925-1]](https://usenet.krcg.org/t/vekn-77029/#m6)
    [[LSJ 20011202]](https://usenet.krcg.org/t/gyzUP4TCq6M/#m41) [[LSJ
    20080612-2]](https://usenet.krcg.org/t/dVIBPf6EX_8/#m11) [[ANK
    20211201]](https://usenet.krcg.org/t/vekn-79519/#m3)
    — {Infernal Pursuit}, {Agaitas, The Scholar of Antiquities}, {Sudario Refraction}, {Waste Management Operation}.

[^6-8-7]: Owner's reading, no official ruling: [[rules-qa
    8.3]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-8--cards--effects-replacement-events)
    — {Visit from the Capuchin}.

[^6-8-8]: Owner's reading, no official ruling: [[rules-qa
    9.4]](https://github.com/lionel-panhaleux/zero-vtes/blob/main/docs/rules-qa.md#batch-9--equipment) — {Shilmulo
    Tarot}, {Black Market Cache}.

[^6-8-6]: [[PIB 20121028]](https://usenet.krcg.org/t/vekn-39816/#m10)
    [[PIB 20121031]](https://usenet.krcg.org/t/vekn-39816/#m35) [[ANK
    20211020]](https://usenet.krcg.org/t/vekn-79416/#m1) [[LSJ
    20081129]](https://usenet.krcg.org/t/7fMPCYIPrag/#m5) [[ANK
    20180512]](https://usenet.krcg.org/t/vekn-76599/#m1) [[LSJ
    20050224-4]](https://usenet.krcg.org/t/LATlh09TuhA/#m2) — {Vast Wealth}, {Vaticination}, {Sergei Voshkov, The Eye},
    {Ashur Tablets}, {Troglodytia}, {Learjet}.
