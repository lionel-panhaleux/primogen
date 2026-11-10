---
name: intake
description: The ingress loop of primogen. Challenge an incoming request, idea or discovery before it reaches BOARD.md. Use whenever new work arrives (an owner ask, a bug, a graded session's correction, something found mid-task) and is not being done right now, and whenever tempted to write a TODO, a plan document or a follow-up note.
---

# /intake

No work reaches the board unchallenged. Intake proposes; the owner accepts.
Refusing a line is far cheaper than evicting one.

Read `docs/index.md`, `docs/roadmap.md` and `BOARD.md` first. For each item,
answer four questions, briefly:

1. **Conflict.** Does it contradict `docs/principles.md` or a decision on
   another page, or duplicate a line? A contradiction is surfaced as a
   decision to overturn, naming the page and the sentence; a duplicate
   sharpens the existing line.
2. **Completability.** Can "done" be stated, checkable where possible (a
   command's output, a test, a graded session)? If not it is documentation:
   put it on the right docs page. Waiting on a date or an outside event is
   not a line either: say so to the owner.
3. **Scope.** Split only on a real boundary, never to put part of it off. If
   it is small and nothing is in flight, do it now instead of adding a line.
   Work that belongs to another repository is done there, through that
   repository's own intake; the line here says only what must exist upstream.
4. **Docs.** Which `docs/` pages will change, or "none" with the reason.

Then propose to the owner, in a message that stands alone and ends there: the
ask, what the four questions found, the line verbatim (the ask, its milestone
token if it serves one, `Done when:`, `Docs:`), and its position by the
ranking rules of `docs/roadmap.md`. The board holds at most 12 lines: a
thirteenth names the line it replaces. A new milestone or a change of phase
is the owner's decision. On acceptance, insert the line at that position.

Bulky context for a line about to be shipped goes in `board/<slug>.md`,
deleted with the line.
