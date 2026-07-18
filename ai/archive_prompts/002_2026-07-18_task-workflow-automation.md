# Current Engineering Task 002: Task Workflow Automation

## Task Metadata

- Task ID: `002`
- Task slug: `task-workflow-automation`
- Status: `complete — awaiting next authorized rotation`
- Engineering update: `E002`
- Date opened: `2026-07-18` UTC
- Human authority: current project owner
- Owner communication language: Hungarian for this owner
- Legacy numbering note: governance record E001 and archived Prompt 001 predate the unified sequential workflow; E002 uses task number `002` as the next available prompt number.

## Title

Task Workflow Automation

## Objective

Create a safe, dependency-free workflow that maintains exactly one active prompt, archives prior prompts without losing semantic status, and creates an independent pending history record for every new task.

## Scope

Prompt migration, relevant engineering-governance documents, `ai/scripts/next-task.sh`, isolated tests, the E002 snapshot, and the independent task history record.

## Out of Scope

QWSG application functionality, dependency installation, server or operating-system changes, Product Architecture execution, and any later engineering task.

## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- Relevant structure, standards, documentation, history, and prior prompt-workflow records

## Starting State Verification

Verified clean `main` at `fa156697198a2230c938fb6780e9b8e42c860f57`, `attila:qwdev` ownership, setgid directories, owner/group write ACL inheritance, one governance README, and one unexecuted Product Architecture draft under `ai/prompts/`.

## Snapshot Requirements

Use `ai/backups/20260718T215938Z_E002_task_workflow_automation/`; validate its baseline commit, permissions record, root guard, explicit confirmation, and exact rollback targets.

## Risk Assessment

- High: prompt or history overwrite — mitigate with complete preflight validation and no-clobber operations.
- High: false execution history — preserve explicit semantic status and migration reasoning.
- Medium: partial rotation — create and validate temporary outputs before moving the active prompt; report failures without deleting history.
- Medium: unsafe input — normalize the slug to lowercase ASCII letters, digits, and single hyphens; reject an empty result.
- Low: permission drift — verify inherited owner/group write, setgid, group, and non-world-writable modes.

## Planned Work

Create the snapshot; migrate governance and draft prompt content by role; update core policies; implement bounded initialization and rotation; test in isolated fixtures; update this prompt and its independent history; commit and verify a clean tree.

## Rollback Plan

Run the guarded E002 snapshot restore procedure only from the exact project root after explicit confirmation. It restores baseline core and prompt files and removes only named E002-created files and empty directories.

## Deliverables

Required directories and policies, one active prompt, the preserved archived draft, `next-task.sh`, isolated test evidence, and `002_2026-07-18_task-workflow-automation.md`.

## Verification

Validate shell syntax, initialization, normal rotation, incrementing, malformed and multiple active files, duplicate archive refusal, unsafe slug handling, outside-root refusal, no-overwrite behavior, permissions, ACLs, semantic migration, exact one-active count, Git diff, and excluded work.

## Documentation Updates

Update only relevant core workflow documents, Engineering History, the migrated archive record, this active prompt, and the independent E002 history.

## Completion Criteria

All required workflow behavior and isolated tests pass; migration remains traceable; Product Architecture remains unexecuted; rollback is valid; the E002 commit hash is recorded in the matching history; and the Git worktree is clean after its audit commit. This file remains the sole active-directory prompt until the next authorized use of `next-task.sh` archives it.

## Owner Approval Requirements

This task is explicitly authorized. Any product work, Product Architecture execution, project-wide permission change, dependency installation, or server change requires separate owner approval.
