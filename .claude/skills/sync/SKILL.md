---
name: sync
description: The upstream loop of primogen. Check the upstreams of the synced files in knowledge/ (vtes-advanced-rules for the rulebook and the advanced rules, the owner's gist for the history of card changes), adjudicate the diff, copy, and record the new commit in docs/provenance.md. Use for "sync", "check upstream", or when the rules or the history of changes moved.
---

# /sync

The diff is ingress, not a notification: every upstream change is read
before it lands. `docs/provenance.md` "Synced" lists each synced file, its
upstream, its path there and the commit it was copied at; that commit is the
baseline.

1. Clone each upstream into `$TMPDIR`:
   - `git clone --recurse-submodules https://github.com/vtes-biased/vtes-advanced-rules.git`
   - `git clone https://gist.github.com/a4a9cad2e492af7c45c50a1cea7e6cf6.git`
2. For each synced file, `cmp` the upstream file with the one in
   `knowledge/`. A difference that the upstream history between the recorded
   commit and `HEAD` does not explain is a hand edit here: stop and tell the
   owner. All identical: report "clean", done.
3. Read each diff (`git diff <recorded commit> HEAD -- "<path>"`, quoted:
   the gist's path has spaces; inside the submodule for the rulebook). Check
   a rules change against `docs/principles.md` "Authority of rules sources"
   and the other synced files. Check each added or changed entry of the
   history of changes against the announcement it links; the card's current
   text cross-checks only its most recent entry, since later errata change
   older ones by design. A change that contradicts its source goes to the
   owner, unapplied: never half a sync.
4. Copy the changed files over their `knowledge/` counterparts.
5. In `docs/provenance.md`, the new commits (the submodule's for the
   rulebook) and today's date.
6. If the change moves a fact the guide relies on, a source a calibration
   entry's `settled-by` names, or a fact an entry without `settled-by` relies
   on, say so to the owner: that entry goes back to them to re-grade or
   retire, and a regression re-run may be due (`/calibrate`).
7. One commit for the sync, saying what changed upstream and what was
   decided. A `knowledge/` change is reviewed (`/ship` step 4).
