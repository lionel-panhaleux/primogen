---
date: 2026-11-20
mode: analysis
answered-by: primogen 0165806
model: Gemini 4 Neon (Medium)
status: graded
regression: no
settled-by: owner's grade; vtes_meta regional shares (South America and Europe, online events excluded) of the 12 months to 2026-10-25 (TWDA of 2026-11-20 06:56 UTC, archon.vekn.net export 2026-11-20 16:03 UTC)
---

## Question

what are the differences between south american and european metas nowadays?

## Answer given

*(First pass missed `continent`/`country` because the client's cached `vtes_meta` schema pre-dated `0165806`; after the owner pointed out regional filters had shipped, `vtes_meta` was called with `continent: "South America"` and `continent: "Europe"` and produced the answer below.)*

### Scope and data limits
- **Window**: 12 months from **`2025-10-26` to `2026-10-25`** vs. prior 12 months (`2024-10-26` to `2025-10-25`), **online events excluded**, weighted by event size. Data as of **2026-11-20** (TWDA `06:56 UTC`, Archon export `16:03 UTC`).
- **TWDA (what wins, not what is played)**:
  - **South America**: **111 winning decks**, **1,798 players** (4 decks have no player count; prior window: 112 decks, 1,973 players). **12.6%** (16 decks) in no archetype.
  - **Europe**: **179 winning decks**, **4,370 players** (prior window: 158 decks, 3,955 players). **8.0%** (16 decks) in no archetype.
- **Full-field Archon (play share vs. win share)**:
  - **Europe** has a usable sample: **38 of 391** finished events (1,472 players, 263 published decks; **13.3%** near no archetype).
  - **South America** does not: only **4 of 249** finished events (57 players, 37 published decks; **32.4%** near no archetype). Regional comparisons below therefore rest on **TWDA winning shares**, with European Archon full-field figures noted where helpful.

---

### 1. Shared core, different trajectories

| Archetype | South America TWDA (change) | Europe TWDA (change) |
|---|---|---|
| **Gangrel Resistance** (C11) | **10.5%** (-7.1 pts, 8 decks) | **11.8%** (+1.5 pts, 22 decks) |
| **Pop Captivation** (C27) | **6.7%** (+6.7 pts, 7 decks) | **6.1%** (+6.1 pts, 9 decks) |
| **Ravnos Bonds** (C2) | **4.7%** (-0.4 pts, 4 decks) | **5.3%** (+3.9 pts, 11 decks) |

- **Gangrel Resistance** still leads both regions, but it is **retreating sharply in South America** (17.6% → 10.5%) while **holding and slightly growing in Europe** (10.3% → 11.8%; also #1 in Europe's Archon full-field sample at 6.8% play share and 27.8% conversion).
- **Pop Captivation** broke out to virtually identical shares in both regions.
- **Ravnos Bonds** is top-5 in both, steady in South America and rising in Europe.

---

### 2. Where Europe diverges from South America (TWDA win share)

Europe is more top-heavy (its top 6 archetypes take **41.7%** of player-weighted wins vs. **33.7%** in South America; only **8.0%** in no archetype vs. **12.6%** in SA) and leans heavily into several archetypes that barely register or are absent in South America:

- **Lasombra Pop** (C51): **7.6%** in Europe (#2, +7.1 pts, 5 decks; 60.0% conversion in EU Archon full-field) vs. **0.0%** in South America.
- **The Unnamed** (C14): **6.9%** in Europe (#3, +4.8 pts, 7 decks; 50.0% conversion in EU Archon) vs. **3.5%** in South America (#7, -0.1 pts, 5 decks) — surging in Europe, flat in SA.
- **Juliet's Dream** (C45): **4.0%** in Europe (#6, +2.3 pts) vs. **1.8%** in South America (#17, +1.8 pts).
- **Platinum Revelation** (C22): **3.6%** in Europe (#7, +0.9 pts; #2 most-played in EU Archon at 4.6% play share, 25.0% conversion) vs. **0.0%** in South America (down from 2.4%).
- **Hecata Corpse** (G16): **2.9%** in Europe (#8, 4.6% play share in EU Archon) vs. **1.4%** in South America (#20).
- **Europe-only or near-only winners in the top 15**: **Qawiyya Caine** (2.2%), **emerald-legion** (2.2%), **Setite Bleed** (2.1%) — all 0.0% in South America — plus **Enkidu Multirush** (1.9% EU vs. 0.8% SA).

---

### 3. Where South America diverges from Europe (TWDA win share)

South America is flatter, has a larger share of winners outside established archetypes (**12.6%** vs. **8.0%**), and keeps winning with several archetypes that have faded in Europe:

- **Endless Dance** (C4): **4.1%** in South America (#4, **+2.0 pts**, 6 decks) vs. **0.9%** in Europe (#30, **-3.4 pts**, 2 decks) — moving in opposite directions (though Endless Dance still sees 3.4% play share in EU Archon full-field).
- **Govern Resistance** (C32): **3.9%** in South America (#5, **+3.4 pts**) vs. **1.4%** in Europe (#22, -0.6 pts).
- **Haqim Royalty** (C1): **3.8%** in South America (#6, -1.3 pts) vs. **2.0%** in Europe (#13, -1.2 pts).
- **Lutz Politics** (C15): **2.1%** in South America (#13, +0.7 pts) vs. **0.3%** in Europe (#43, -1.6 pts).
- **Higher mid-tier shares in South America**: **Illegal Brawl** (2.8% vs. 1.7%), **Salubri Powerbleed** (2.7% vs. 1.3%), **Oblivion Bleed** (2.7% vs. 1.8%), **Kalinda** (2.5% vs. 1.3%), plus **Ghede Typhonic Beast** (2.1%) and **Inside Dirt** (1.7%), both 0.0% in Europe's TWDA.

---

### 4. Card-level shifts behind the differences

- **Shared wave**: Both regions saw **vote defense, bounce, and the Pop Captivation / Sabbat package** surge — *Cold War* (+14.6 pts EU / +9.1 pts SA), *Delaying Tactics* (+12.7 / +9.9), *Deflection* (+9.4 / +11.3), *Govern the Unaligned* (+10.0 / +13.0), *Direct Intervention* (+8.0 / +8.8), *Obedience* (+6.8 / +9.3), and *Unthinkable Humiliation* / *Unholy Sacrament* / *Privileged Position* / *Magnetic Authority*.
- **Europe's specific shift**: Strong rise in **Lasombra Pop and Ravnos Bonds** cards (*Ambrosio Luis Monçada* and *Francisco Domingo de Polonia* +7.1 pts each, *Information Highway* +9.0, *Private Audience* +7.8, *Narrow Minds* +7.9, *Memory Rift* +7.0), paired with a sharp drop in **combat/rush and Auspex** staples (*On the Qui Vive* -17.2 pts, *Fame* -12.7, *Taste of Vitae* -8.7, *Bowl of Convergence* -8.2, *Deep Song* -8.1, *Tension in the Ranks* -7.4, *Psyche!* -7.3).
- **South America's specific shift**: Much steeper drop in **Gangrel / Anarch** infrastructure (*Backways* -12.8 pts, *Ecoterrorists* -11.8, *Anarch Railroad* -11.2, *Bait and Switch* -10.2, *Form of Mist* -9.0, *Rewilding* -8.6), while **pool/blood economy and Baron/bleed tools** climbed (*Giant's Blood* +14.8 pts, *Heart of Nizchetus* +12.7, *Villein* +12.1 to 74.5%, *Freak Drive* +10.1, *Papillon* +9.8 — whereas *Papillon* fell -6.2 pts in Europe).

## Owner grade

good answer. you can count everything as calibrated now (including meta and full-meta, plus members and events info from authenticated version)

## Generalization

Case-specific: none.
