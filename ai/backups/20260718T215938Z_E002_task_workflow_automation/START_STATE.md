# E002 Task Workflow Automation Start State

## Purpose

This snapshot records the verified state before Engineering Update E002 introduces active-prompt rotation and independent task history.

## Status

Captured at `2026-07-18T21:59:38Z` before task targets were modified.

## Baseline

- Root: `/home/qws/web/qwsg.quantumwizard.hu/qwsg`
- Branch: `main`
- HEAD: `fa156697198a2230c938fb6780e9b8e42c860f57`
- Working tree: clean
- Owner and group: `attila:qwdev`
- Project directories: setgid; owner and group write inherited
- Default ACL: `default:user::rwx`, `default:group::rwx`
- `ai/prompts/README.md`: prompt-governance documentation
- `ai/prompts/001_PRODUCT_ARCHITECTURE.md`: unapproved and unexecuted future-task draft
- `ai/archive_prompts/` and `ai/scripts/`: absent

## Planned changes

E002 may migrate the two prompt files according to semantic role; add `ai/core/14_PROMPT_WORKFLOW.md`, `ai/archive_prompts/`, `ai/scripts/next-task.sh`, one `002_CURRENT_TASK.md`, and one independent E002 history file; and update only relevant core structure, standards, history, task, and documentation policies.

No QWSG application functionality, dependency, service, operating-system configuration, or Product Architecture work belongs in this task. This snapshot remains fixed while governance documentation evolves through approved work.
