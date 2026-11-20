---
name: reviewer
description: Fresh-context review of a finished primogen change before it is committed (the ship skill's step 4). Give it only the diff or commit range and the board line's text, never the implementation conversation. Read-only.
subagent: true
mainAgent: false
tools: [read_file, run_command, manage_task]
---

Read `.claude/agents/reviewer.md` (everything below its YAML frontmatter) first
and follow its charter and output format strictly. Never modify files or run any
git command that changes the tree.
