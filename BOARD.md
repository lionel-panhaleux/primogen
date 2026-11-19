# Board

What must change. The goal is zero. Completion is deletion: no closed state,
no archive. Context lives in [`docs/`](docs/index.md), never here. Order is
priority, by the ranking rules of [`docs/roadmap.md`](docs/roadmap.md). At
most 12 lines: adding a thirteenth means removing one. A line waiting on a
date or an outside event is not a line.

---

- **`vtes_deck` gives a list's track record** (M:deck): the share of the library in cards under 1% of the window's decks play beside the reference's quartiles, those legal under 18 months apart and uncounted; each card no window deck plays, with the TWDA decks ever playing it, the last one's date and the date it became legal; for a list or a draft. Done when: the Sword of Nuln draft shows the Sword legal since 2003 and in no TWDA deck; the Keepers of Tradition Ventrue precon's list stands above the window's 90th percentile and the Fifth Edition Malkavian precon's within its reference's range (both from the card data's set information); the Banu list says every card is played; `calibration/corpus/2026-11-19-tremere-sword-of-nuln.md` is re-run and the fresh answer states the Sword is in no TWDA deck since 2003 and calls the idea non-competitive. Docs: product.md (the `vtes_deck` row, "Grading a deck"), `deck-advice.md` step 7.
- **`vtes_deck` gives the odds of a hand** (M:deck): for each module the chance an opening hand holds a copy; with `together` (recipes, each a list of groups of card names), the exact chance an opening hand holds one of each group and of having seen them within 14, 21 and 28 cards, and the crypt cards able to play every piece; a draft's library taken as 90 cards unless the caller gives its size, said; the opening hand said to be the only random one. Done when: on the Banu list, Hunger of Marduk with (Blur, Lightning Reflexes, Quickness, Side Strike, Pursuit) returns 22.0% and TWDA deck 12494 36.9%; `deck-advice.md` tells the model to name the recipe from the card texts. Docs: product.md (the `vtes_deck` row, "Grading a deck").
- **The blood and pool sources of a deck** (M:deck): `module-tags.json` gains an `economy` block, each card gaining its controller pool or blood with its kind, every entry checked against the card's text and the doubtful ones settled by the owner in the session; the build warns of a card the window plays whose text gains its controller pool or a minion blood ("gain N pool", "move N blood from the blood bank", "steal N blood" and their kin) and is neither tagged nor excluded with a reason; `vtes_deck` prints each kind's copies beside the reference, copies not yield. Done when: the build passes with no such warning, and the Banu list prints its sources by kind beside Haqim Royalty's quartiles, its 7 Hunger of Marduk counted as a steal. Docs: product.md (the `vtes_deck` row, "Grading a deck"), operations.md (the build's warning), provenance.md (what `module-tags.json` holds).

<!-- completed-since-upkeep: 11 -->
