# Active Prompt and Task Rotation Workflow

## Purpose

This document defines the permanent lifecycle, storage, numbering, rotation, and traceability rules for Quantum Wizard Server Guardian engineering prompts and task histories.

## Status

Active governance refined by Engineering Update E002 and Task 013. `ai/scripts/task-builder.sh` implements official approved task generation; `ai/scripts/next-task.sh` preserves the bounded unapproved-draft workflow.

## Directory model

- `ai/prompts/` contains zero or one active Markdown task file named `NNN_CURRENT_TASK.md`. Zero is the canonical idle state after a completed task is archived without an authorized successor.
- `ai/archive_prompts/` permanently stores prior prompts as `NNN_YYYY-MM-DD_task-slug.md`.
- `ai/history/` stores one independent history file per task as `NNN_YYYY-MM-DD_task-slug.md`.
- `ai/scripts/task-builder.sh` is the official structured, approved task-generation workflow.
- `ai/scripts/next-task.sh` remains the guarded compatibility workflow for preparing an unapproved draft.
- `ai/test_tasks/` stores explicitly diverted experimental or aborted task
  evidence under independent identifiers such as `001_TEST_TASK`. It is not a
  production prompt, archive, or history directory.
- `ai/framework/VERSION` identifies the reusable framework release.
- `ai/config/engineering-project.conf` declares validated project identity and
  lifecycle configuration without becoming executable shell input.

Production task numbers are sequential and zero-padded: `001`, `002`, `003`,
and so on. Completed or normally archived numbers are never reused. The sole
exception is an incomplete active number explicitly released through the
Project Owner-authorized aborted-test diversion protocol; the preserved test
record proves the release and retains the original identity. Test-task IDs use
an independent `NNN_TEST_TASK` sequence and never influence production
numbering. Prompt archives and task histories may share filenames because they
occupy different directories. All lifecycle and diversion records must remain
committed. Generated binary backup archives may remain ignored where
appropriate; Markdown audit records must remain trackable.

## One active prompt

At most one active Markdown prompt may exist in `ai/prompts/`. Governance documentation does not belong there. A prompt defines one bounded task; unrelated work and multiple milestones do not belong in one file. A completed prompt may be archived without creating a successor. In that idle state the highest-numbered archive and its matching history are the complete lifecycle baseline. No next task may start until a separately authorized prompt/history pair is installed transactionally.

## Required structure

Each Framework 1.1.0 active prompt follows `ai/core/08_JOB_TEMPLATE.md` and
contains Task ID, title, task slug, Objective, Scope, Out of Scope, Authority
Envelope, Required Reading, Starting State Verification, Snapshot Requirements,
Risk Assessment, Planned Work, Rollback Plan, Deliverables, Verification,
Documentation Updates, Completion Criteria, and Owner Approval Requirements.
Historical prompts remain valid without retroactive Authority Envelope edits.

Task Metadata declares one semantic status: `draft`, `approved`, `active`, `complete`, `superseded`, or `archived without execution`. It also records human authority and the preferred communication language when verified. The language preference is a configurable policy concept; no configuration file is defined yet.

## Lifecycle

The authoritative lifecycle and transaction semantics are defined in `11_ENGINEERING_LIFECYCLE.md`; this document owns prompt storage and naming details.

1. Validate the current lifecycle. After verified completion, either archive the completed prompt and enter the idle state, or—only with separate owner authority—run `ai/scripts/task-builder.sh` from the exact project root and provide every structured owner field. Use `--input-dir` for deterministic automation or no arguments for interactive multi-line input.
2. Explicitly approve the complete structured definition. The builder generates an approved prompt/history pair with no editing markers. If a separate review cycle is needed, use `next-task.sh`; its draft must remain unexecuted until owner approval.
3. Read `00_PROJECT_PHILOSOPHY.md`, `01_CONSTITUTION.md`, `03_AGENTS.md`, and `08_JOB_TEMPLATE.md`; verify the environment and snapshot before changes.
4. Update the task's independent pending history file during work and at delivery.
5. Set the prompt's semantic status accurately. Archiving does not by itself mean completion or execution.
6. Before the next task begins, rotate the prior prompt into `ai/archive_prompts/` and create the next active prompt and pending history record without overwriting files.

An incomplete task never enters `ai/archive_prompts/` as though complete. Under
an explicit Project Owner override, `ai/scripts/divert-task-to-test.sh` may move
its unchanged prompt and history into the next free
`ai/test_tasks/NNN_TEST_TASK/` record. The record contains disposition authority,
reason, original ID and slug, `aborted-test` status, `incomplete` result,
production-number release, preserved-path mapping, and SHA-256 evidence.
Production validators ignore this namespace by default; explicit audit uses
`bin/job --check-test-tasks`.

The general Engineering History is a milestone index, not an infinitely growing task log. Detailed work belongs only in independent task history files. Existing valid historical records keep their original names and meaning for backward compatibility; they must not be rewritten to imply execution.

## Engineering Task Builder

The builder owns Task ID incrementing, UTC creation date, approved status,
prompt/history filenames, required-reading insertion, approval text, and
transaction ordering. Owner input owns the title, authority, communication
language, objective, scope, exclusions, Authority Envelope, starting checks,
snapshot requirements, risks, planned work, rollback, deliverables,
verification, documentation, and completion criteria.

When production is idle after an aborted-test diversion, the builder derives
the next production ID from the latest completed production archive. The
released incomplete ID is therefore reused without reading or incrementing any
test-task identifier. A replacement task must use a clean owner-approved
definition and may reference, but must not inherit, incomplete evidence.

Interactive multi-line fields end with a line containing only `.`. Deterministic mode reads the documented one-file-per-field input directory without sourcing or evaluating any file. The exact `APPROVE` token is mandatory. The builder first renders and validates same-directory temporary documents, then archives the completed prompt and installs the new pair with no-clobber moves. Post-install `bin/job --check` and lifecycle consistency validation are mandatory; any failure performs bounded automatic rollback.

When the versioned framework is present, builder, lifecycle, diversion, and job
entry points require `framework-check.sh` to validate project configuration,
repository identity, canonical branch and remote, required reading, lifecycle
paths, and configured validation argv. Older isolated compatibility fixtures
without `ai/framework/VERSION` retain their bounded legacy validation behavior.

Preparing or reviewing a draft does not authorize execution. Canonical Builder
installation with explicit Owner approval creates an executable task and
authorizes its complete Authority Envelope without a second routine start gate.
Archiving never creates authority for another task. Prompts contain
instructions and acceptance criteria, not secrets, credentials, unverified
environment claims, application output, or completed architecture decisions.
This workflow will evolve through approved engineering-governance updates.

## Project-local job access

`bin/job` provides installation-free, read-only access to an active prompt and validates the canonical idle state. It validates project markers, prompt count and filename, internal Task ID, unresolved human-editing fields, and matching history before producing output. It never executes engineering work or interprets Markdown as shell code.

Use it from the repository without global installation:

```bash
./bin/job
./bin/job --check
./bin/job --check-test-tasks
./bin/job --show
./bin/job --path
./bin/job --history
./bin/job --help
```

The default mode validates, summarizes, and displays an executable prompt.
`--check` validates executable production state or the idle state;
`--check-test-tasks` audits independent diverted records; `--prepared-check`
validates the explicitly unapproved generated draft pair; `--show`, `--path`,
and `--history` require an active prompt. Multiple, malformed, inconsistent,
incomplete, or ambiguously recorded production tasks fail without changing
files. The command does not rotate or generate prompts and does not replace
either lifecycle generator.

The repository skill `.agents/skills/qwsg-job/SKILL.md` tells Codex how to apply the full governed workflow after requests such as `job`, `/job`, `Új feladat`, `aktuális feladat`, or `indítsd a feladatot`. The shell command only validates and displays task data; the Codex skill governs how an authorized task is read and carried out. `/job` is human shorthand and may not be a built-in Codex slash command.

For the current shell session only, a user may optionally run `export PATH="$PWD/bin:$PATH"` from the verified project root and then invoke `job`. This is neither required nor a global shell change. The one-active-prompt rule and all Constitution, authority, snapshot, rollback, verification, and history requirements remain unchanged.
