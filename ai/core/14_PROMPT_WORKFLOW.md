# Active Prompt and Task Rotation Workflow

## Purpose

This document defines the permanent lifecycle, storage, numbering, rotation, and traceability rules for Quantum Wizard Server Guardian engineering prompts and task histories.

## Status

Active governance refined by Engineering Update E002. `ai/scripts/next-task.sh` implements the bounded prompt rotation described here.

## Directory model

- `ai/prompts/` contains exactly one active Markdown task file named `NNN_CURRENT_TASK.md`.
- `ai/archive_prompts/` permanently stores prior prompts as `NNN_YYYY-MM-DD_task-slug.md`.
- `ai/history/` stores one independent history file per task as `NNN_YYYY-MM-DD_task-slug.md`.
- `ai/scripts/next-task.sh` performs guarded rotation and initialization.

Task numbers are sequential, zero-padded, and never reused: `001`, `002`, `003`, and so on. Prompt archives and task histories may share filenames because they occupy different directories. Both are engineering records and must remain committed to Git. Generated binary backup archives may remain ignored where appropriate; Markdown audit records must remain trackable.

## One active prompt

Exactly one active Markdown prompt may exist in `ai/prompts/`. Governance documentation does not belong there. A prompt defines one bounded task; unrelated work and multiple milestones do not belong in one file. A completed or replaced prompt remains the sole active record only until the next authorized rotation; `next-task.sh` archives it before creating the next active prompt. No next task may start before that rotation succeeds.

## Required structure

Each active prompt follows `ai/core/08_JOB_TEMPLATE.md` and contains Task ID, title, task slug, Objective, Scope, Out of Scope, Required Reading, Starting State Verification, Snapshot Requirements, Risk Assessment, Planned Work, Rollback Plan, Deliverables, Verification, Documentation Updates, Completion Criteria, and Owner Approval Requirements.

Task Metadata declares one semantic status: `draft`, `approved`, `active`, `complete`, `superseded`, or `archived without execution`. It also records human authority and the preferred communication language when verified. The language preference is a configurable policy concept; no configuration file is defined yet.

## Lifecycle

1. Initialize or rotate with `ai/scripts/next-task.sh` from the exact project root.
2. Edit every `[REQUIRES HUMAN EDITING]` field and obtain explicit human approval before implementation.
3. Read `00_PROJECT_PHILOSOPHY.md`, `01_CONSTITUTION.md`, `03_AGENTS.md`, and `08_JOB_TEMPLATE.md`; verify the environment and snapshot before changes.
4. Update the task's independent pending history file during work and at delivery.
5. Set the prompt's semantic status accurately. Archiving does not by itself mean completion or execution.
6. Before the next task begins, rotate the prior prompt into `ai/archive_prompts/` and create the next active prompt and pending history record without overwriting files.

The general Engineering History is a milestone index, not an infinitely growing task log. Detailed work belongs only in independent task history files. Existing valid historical records keep their original names and meaning for backward compatibility; they must not be rewritten to imply execution.

Creating, archiving, or reviewing a prompt does not authorize execution. Prompts contain instructions and acceptance criteria, not secrets, credentials, unverified environment claims, application output, or completed architecture decisions. This workflow will evolve through approved engineering-governance updates.
