# Task History 010: Automated Task Lifecycle and Safe Next-Task Preparation

## Task metadata

- Task ID: `010`
- Task slug: `automated-task-lifecycle`
- Status: `complete`
- Date opened: `2026-07-21` UTC
- Human authority: Project Owner
- Responsible agent: Aikó/Codex

## Starting state

Repository root `<repository-root>`, branch `main`, HEAD `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`. Exactly one valid Task 010 prompt and matching history existed after the owner removed the earlier duplicate. The worktree already contained extensive uncommitted Tasks 003–009 documentation, workflow, source, test, and backup work; it was preserved as the Task 010 baseline.

The prior `next-task.sh` enforced root/markers, one prompt, safe names/slugs, sequential IDs, destination collision checks, same-directory temporary files, no-clobber moves, and partial rollback. Gaps were interactive-only operation; no completed-history gate or prompt/history metadata cross-validation; no read-only check mode; history was not removed after late transaction failure; no post-rotation validation; no separate generated-draft validation; and no automated failure-injection suite. `bin/job` validated executable prompts but rejected generated placeholders, requiring a distinct prepared-state mode.

ShellCheck was unavailable and was not installed. Task 010 prompt/history modes were `0600` with ACL mask `---`; this pre-existing collaboration limitation was preserved and disclosed.

## Snapshot

`ai/backups/20260720T235829Z_task010_automated_task_lifecycle/` captures the full Git baseline, permissions, preserved target files, manifest, SHA-256 hashes, and guarded bounded restore. Checksums and restore syntax passed.

## Work performed

Created the authoritative `ai/core/11_ENGINEERING_LIFECYCLE.md` with No Task Without History, authority states, gates, transactional preparation, validation modes, and rollback rules. Added concise integrations to the Constitution, Agent Rules, Job Template, Prompt Workflow, and README.

Hardened `ai/scripts/next-task.sh` with read-only `--check`, non-interactive `--prepare --slug`, compatible interactive mode, prompt/history Task ID and slug cross-validation, finalized-history and evidence gates, sequential/collision validation, stale-temporary detection, complete rollback after archive/prompt/history stages, failure injection, prepared-state post-validation, and explicit `READY FOR OWNER REVIEW` reporting.

Added `bin/job --prepared-check` while retaining strict executable-task behavior for existing modes. Generated prompt/history scaffolds contain matching machine-readable metadata and explicit draft, unapproved, not-started states.

Added isolated lifecycle tests covering interactive and non-interactive success, read-only checking, malformed/missing/duplicate/mismatched records, pending history, every destination conflict, three injected transaction failures, full restoration, no orphan temporary files, and matching generated metadata.

## Files changed

Modified `ai/scripts/next-task.sh`, `bin/job`, `README.md`, `ai/core/01_CONSTITUTION.md`, `ai/core/03_AGENTS.md`, `ai/core/08_JOB_TEMPLATE.md`, `ai/core/14_PROMPT_WORKFLOW.md`, this history, and the Task 010 prompt. Created `ai/core/11_ENGINEERING_LIFECYCLE.md`, `ai/tests/test-next-task.sh`, and the Task 010 snapshot. Final preparation archives the Task 010 prompt and creates the Task 011 prompt/history pair.

## Decisions

Executable validation and prepared-draft validation are separate to prevent circular rejection without weakening authority. Rotation requires both prompt and history to declare completion and the history to contain the required evidence headings with no pending markers. Preparation never invents an objective, approves, or executes the next task. Test-only failures use the bounded `QWSG_TEST_FAIL_AFTER` values `archive`, `prompt`, and `history`.

## Verification

Passed Bash syntax validation, 23 isolated lifecycle assertions, current lifecycle validation, executable `bin/job` validation, Git whitespace checks, documentation-link validation, Task 009 snapshot integrity and restore syntax, Task 010 snapshot integrity and restore syntax, Go formatting, `go vet ./...`, `go test ./...`, `go test -race ./...`, QWSG binary build, and sanitized JSON fixture parsing. ShellCheck was unavailable; no package was installed. HEAD remained unchanged.

## Problems encountered

The first test run exposed a `set -e` interaction in the clean temporary-file check; it was corrected to an explicit conditional and the complete suite then passed. The earlier duplicate Task 010 history blocked execution until the owner resolved it before this run.

## Rollback

From the repository root run `ai/backups/20260720T235829Z_task010_automated_task_lifecycle/restore.sh`. It restores only captured Task 010 files, removes only enumerated Task 010-created and final-preparation artifacts, and preserves unrelated pre-existing changes.

## Open questions

The historical `0600`/ACL mask limitation remains under the existing governance permission gate. Reusable cross-project extraction remains deferred. No commit or push occurred.

## Recommended next task

Task 011: `core-alpha-platform-hardening`, prepared for owner review but not approved or started.

## Completion state

`complete`
