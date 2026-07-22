# Task History 013: Engineering Task Builder

## Task metadata

- Task ID: `013`
- Task slug: `engineering-task-builder`
- Status: `complete`
- Date opened: `2026-07-21` UTC
- Responsible agent: Aikó/Codex
- Human authority: Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/013_CURRENT_TASK.md`
- Dependency: completed Task 012 and the existing Engineering Lifecycle

## Starting state

Task 012 is complete in both its archived prompt and independent history. `bin/job --check` and `ai/scripts/next-task.sh --check` accepted the approved Task 013 prompt/history identity. The inherited worktree is intentionally dirty with prior tracked and untracked engineering work; no inherited change is treated as Task 013 work. Lifecycle scripts are owned by `attila:nogroup`, mode `0770`, with equivalent access ACLs. The active prompt and history were mode `0600` before this update.

Baseline verification passed with a workspace-safe Go cache:

- `GOCACHE=/tmp/qwsg-task013-go-cache GOPATH=/tmp/qwsg-task013-gopath go test ./...`: passed for all five Go packages.
- `bash ai/tests/test-next-task.sh`: `PASS: 23 lifecycle assertions`.
- A default-cache `go test ./...` attempt could not create `<user-cache>/go-build` because that path is read-only in the execution sandbox; this is an environmental cache-path constraint, not a product failure.

## Snapshot

- Location: `ai/backups/20260721T_task013_engineering_task_builder/`
- Archive: `repository-before-task013.tar.gz`
- Scope: complete repository worktree excluding `.git`, prior `ai/backups`, and generated `build` output; includes lifecycle scripts, engineering documentation, prompt templates and records, history handling, validation logic, tests, and implementation source.
- Integrity: `SHA256SUMS` was verified with `sha256sum -c`; archive readability was verified with `tar -tzf`.
- Retention: retain through owner acceptance and any subsequent task that depends on this builder.

## Risk assessment

| Risk | Rating | Mitigation |
|---|---|---|
| Invalid or inconsistent metadata | High | Generate ID/date/status centrally and validate prompt/history identity before commit. |
| Unapproved task installation | High | Require an exact explicit owner approval token and emit an approved record only after it is present. |
| Partial lifecycle mutation | High | Use same-directory temporary files, bounded moves, failure injection, and automatic rollback. |
| Input execution or injection | High | Read owner content strictly as data; never source or evaluate it. |
| Multi-line truncation | Medium | Use explicit terminators interactively and one-file-per-field deterministic input. |
| Existing lifecycle regression | High | Preserve `bin/job`, `next-task.sh`, and their existing tests; add dedicated builder tests. |

## Planned work

Implement a dependency-free shell builder with interactive and deterministic input-directory modes, centralized metadata and document rendering, explicit approval, pre-install validation, transactional archive/install/rollback, lifecycle integration, documentation, and unit-style shell tests. Preserve all existing lifecycle interfaces and avoid application-runtime changes.

## Rollback

Preconditions: stop lifecycle tooling, inspect the current worktree, verify the exact archive and `SHA256SUMS`, and obtain owner confirmation before overwriting any Task 013 target.

1. Verify `ai/backups/20260721T_task013_engineering_task_builder/SHA256SUMS` from the repository root.
2. List the archive and compare only Task 013-created or modified targets with the live worktree.
3. Restore only the exact affected lifecycle scripts, tests, documentation, prompt, and history paths from the archive; remove only Task 013-created paths after explicit confirmation.
4. Re-run `bin/job --check`, lifecycle tests, Go tests with a writable cache, and `git diff --check`.

Broad Git reset, wildcard deletion, and replacement of inherited dirty-worktree changes are forbidden.

## Work performed

Implemented `ai/scripts/task-builder.sh` as the official repository-local task generator. It supports interactive structured collection with explicit multi-line terminators, deterministic one-file-per-field `--input-dir` generation, and read-only `--check-input` validation. All owner content is read as data and never sourced or evaluated.

The builder validates a completed predecessor, all required non-empty fields, single-line metadata, lowercase kebab-case slug, unresolved editing markers, the exact `APPROVE` token, sequential ID range, and destination collisions. It generates the UTC date, Task ID, filenames, approved status, mandatory required reading, Owner Approval text, and matching approved-but-unstarted history metadata.

Generation occurs in same-directory temporary files. Required sections, prompt/history identity, and approval state are checked before mutation. The transaction archives the completed prompt, installs the next prompt, installs its history, and runs both `bin/job --check` and lifecycle consistency validation. Bounded automatic rollback tracks every move, removes only builder-installed targets, restores the original active prompt, and cleans temporary files.

Added `ai/tests/test-task-builder.sh` with deterministic input, interactive multi-line input, invalid/missing input, missing approval, marker rejection, incomplete predecessor, post-archive/post-prompt/post-history injected failures, lifecycle compatibility, and byte-identical generation coverage. Existing `next-task.sh` behavior remains supported for a separate unapproved draft/review workflow.

## Verification

All final gates passed:

- `bash -n ai/scripts/task-builder.sh ai/tests/test-task-builder.sh ai/scripts/next-task.sh bin/job`: passed.
- `bash ai/tests/test-task-builder.sh`: `PASS: 33 task-builder assertions`.
- `bash ai/tests/test-next-task.sh`: `PASS: 23 lifecycle assertions`.
- Builder rollback injection after `archive`, `prompt`, and `history`: original prompt restored; generated prompt/history/archive and temporary files absent.
- Determinism: two independent fixtures with identical structured input produced byte-identical prompt and history files using `cmp`.
- Interactive multi-line generation: two objective lines were preserved and lifecycle validation passed.
- `bin/job --check`: passed for Task 013.
- `ai/scripts/next-task.sh --check`: passed for Task 013.
- `GOCACHE=/tmp/qwsg-task013-go-cache GOPATH=/tmp/qwsg-task013-gopath go test ./...`: passed for all five Go packages.
- `make fmt-check`: passed.
- `git diff --check`: clean.
- Task targets are owned by `attila:nogroup`; scripts/tests are mode `0770`, documentation is mode `0660`, and inherited ACLs preserve owner/group write access.
- No generated task can contain an unresolved `[REQUIRES HUMAN EDITING]` marker; literal marker occurrences are limited to rejection logic and its negative test.
- No GUI, web interface, remote management, AI task generation, planning, roadmap management, task execution, or automatic implementation was added.

## Documentation updates

- Added `docs/architecture/ENGINEERING_TASK_BUILDER.md` with boundaries, input contract, ownership, deterministic pipeline, transaction, rollback, security, and compatibility decisions.
- Updated `ai/core/04_ARCHITECTURE.md` to register the builder architecture.
- Updated `ai/core/08_JOB_TEMPLATE.md`, `ai/core/11_ENGINEERING_LIFECYCLE.md`, and `ai/core/14_PROMPT_WORKFLOW.md` with the official builder workflow and retained draft compatibility path.
- Updated `README.md` with task-generation entry points.
- Updated `ai/core/07_ENGINEERING_HISTORY.md` and `ai/core/13_ROADMAP.md` with the completed Task 013 outcome.
- Updated the active prompt and this independent history to their truthful completion state.

## Unresolved issues

None within Task 013 scope. The repository remains intentionally dirty with inherited work and Task 013 changes; no commit or push was performed.

## Completion state

`complete`

Delivery result: `complete`.

The official Engineering Task Builder, structured and multi-line input modes, automatic metadata and approval generation, lifecycle validation, transactional installation and rollback, documentation, and tests all exist and passed their required verification. The next task was not generated, approved, or started.
