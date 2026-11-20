---
date: 2026-11-20
mode: advice
answered-by: primogen 2c096e2
status: graded
regression: yes
settled-by: owner's grade; rulebook, Paths of Enlightenment; card text of Qawiyya el-Ghaduba, Soulgrinder, Immortal Grapple, Death Seeker; Qawiyya Caine (G29) 4 decks, 1.3% of the 12 months to 2026-10-25 (vtes_meta, TWDA of 2026-11-19)
---

## Question

I'm thinking of running a Qawwiya Caine rush deck for the next tournament, how does that sit in the current meta?

## Answer given

I read "Qawwiya Caine" as **Qawiyya el-Ghaduba** (G5, Salubri antitribu, capacity 9). Her text: *"Sabbat: Qawiyya can burn 2 blood to unlock during your minion phase. Red List. +2 strength."* Primogen's labels have an archetype for this deck: **Qawiyya Caine** (G29).

**Its record**
- 12 months: 1.3% share, 4 decks, ranked 21st of 69 archetypes.
- Last 6 months: 2.2% share, 9th place, up from 0.
- It is entirely new this window. The Caine cards it uses (Abuse of Power, Augury of Doom, Rage of Apedemak, Soulgrinder, Twisted Bloodhound) became legal on 2025-10-26.

The four wins, all between May and August 2026:
- Bologna, 21 players
- Polish NC, 40 players (the archetype's representative deck)
- Helsinki, 27 players
- Helsinki, 10 players

Two of the four were won by the same player. That is a promising start but a small sample, and it is concentrated in a few hands.

**What the list does**, from the Polish NC list (13350):
- Crypt: 5× Qawiyya, then Andre LeRoux and cheap support (Anarch Convert, New Blood, Walker Grimes, Donald Cargill). The crypt groups are Any and G5.
- Freak Drive ×12 alongside Qawiyya's own unlock gives multiple combat actions per turn. Freak Drive is in 74% of opening hands.
- Immortal Grapple ×7 and Taste of Vitae ×7 form the POT grapple kit: 23.6% of the library, in 86% of opening hands.
- Death Seeker ×3, Unflinching Persistence ×5 and Rolling with the Punches ×3 protect Qawiyya in combat.
- Ashur Tablets ×6 provide the pool and the recursion.

**Against the field's defenses** (winning decks since 2021):
- **Combat ends** is in 55.6% of them: Earth Meld, Form of Mist, Majesty. The current leaders carry it heavily. Gangrel Resistance plays Earth Meld in 94% of its decks and Form of Mist in 88%. The Unnamed plays Form of Mist in 91%, and Living Museum plays both in 100%.
  - The deck answers with **Immortal Grapple**: *"Strikes that are not hand strikes cannot be used this round (by either combatant)."*
  - It also has **Death Seeker**: *"Cancel a combat card played by the opposing minion as it is played."* Death Seeker is limited to Qawiyya (Salubri antitribu), and it costs 1 blood.
  - So the top three archetypes are exactly what this deck is built to beat. How often it actually wins those fights is not in the data.
- **Prevention** is in 31.4% of winners. Its only answers are Death Seeker and Soulgrinder (superior: *"damage from this strike cannot be prevented by cards requiring Fortitude"*).
  - The data shows the Caine path on Soulgrinder and the other Caine cards, but no crypt card in the list carries the path. From the data I can't tell whether that limits who can play them. All four winning lists play them anyway.

**Verdict**: it's a new, legitimately successful rush deck that rose into the top 10 over the last 6 months. Its kit targets the combat-ends defenses that the three leading archetypes rely on. The evidence is four wins, so treat it as a promising contender, not a proven top tier.

## Owner grade

All three are very good answers. One slight wrinkle on (B):

> The data shows the Caine path on Soulgrinder and the other Caine cards, but no crypt card in the list carries the path. From the data I can't tell whether that limits who can play them.

I'm a bit suprised at that. Primogen should know that any sabbat vampire without path can elect a path for one blood when it is influenced out. That is what the star vampire will do in this deck. Maybe the LLM should have asked a follow-up question to primogen and did not?

## Correction

Qawiyya el-Ghaduba is Sabbat and has no printed Path, so she can follow the Path of Caine. The rulebook, Paths of Enlightenment (https://archon.vekn.net/help/rules#paths-of-enlightenment): "During your influence phase, each of your Sabbat vampires not following a Path can burn 1 blood to follow a Sabbat Path from the above list." The deck's Caine cards are therefore played by its star vampire, which is why the lists play them.

Two causes. The answering model did not query a gap it could have closed: `vtes_rules_search` "path Sabbat vampire influenced burn 1 blood" returns that section first. And `vtes_deck` primed the gap: its requirement check counts only a printed Path ("Crypt cards with the path Caine it requires: none"; "Soulgrinder ×2 … needs [POT] or [THA] and path Caine: no crypt card here has it"), and it does not count a Sabbat vampire with no printed Path as able to follow one.

## Generalization

A tool fix, not guide wording (the guide already says to search before
calling a point unsettled): the BOARD.md line "`vtes_deck` counts a Sabbat
vampire with no printed Path as able to follow a Sabbat Path", approved by
the owner on 2026-11-20.
