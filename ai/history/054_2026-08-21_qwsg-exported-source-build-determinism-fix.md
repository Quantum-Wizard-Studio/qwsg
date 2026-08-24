# Task History 054: Approved Task

## Task metadata

- Task ID: `054`
- Task slug: `qwsg-exported-source-build-determinism-fix`
- Status: `complete`
- Date generated: `2026-08-21` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/054_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded. The Project Owner started Task 054 through the canonical `job` workflow on 2026-08-21 UTC. The narrow implementation was integrated into `main` as commit `ef513dde187e4119b6aa04a3439a879056f6cc69` (`fix: make exported-source builds deterministic`) and directly verified on Forgejo. On 2026-08-24 UTC the Owner accepted the integrated implementation and authorized canonical lifecycle completion and archival without creating a successor task.

## Starting state

`HEAD`, tracking `origin/main` and direct remote `main` were
`3e7d2f9d543078b49d1afa522dc6bf3baba1c949`, ahead/behind `0/0`, branch
`main`, and index empty. The expected lifecycle-only working changes from Task
053 closure/Task 054 installation and the excluded Owner-owned untracked path
were the only local state. `VERSION` was `1.1.0-rc.3`; no RC.4 artifact or tag
existed. RC.1/RC.2/RC.3 and historical findings were preserved.

The genuine failure was reproduced once in private mode-0700 export
`/tmp/qwsg-task054-export-failure.7QAy5h` with `.git` absent and `GOFLAGS`
unset: ordinary `make build` stopped with `error obtaining VCS status`.

## Snapshot

Private mode-0700 snapshot:
`/tmp/qwsg-task054-implementation.9N9Xio`. It contains a readable tracked-HEAD
archive, Task 053 closure/RC.3 ledger, Task 054 prompt/history and Builder
source/input, build/release targets and absence records, tool/Git/mode/ACL and
protected preservation evidence, plus bounded restore instructions. RC.3
candidate bytes and Owner content are excluded.

## Work performed

- Added explicit `-buildvcs=false` to the canonical Makefile build invocation.
- Added `scripts/test-build-contract.sh` and Makefile wiring for genuine
  no-`.git`, no-`GOFLAGS` twin exports, truthful defaults, explicit identity,
  byte identity and absence of ambient Go `vcs.*` settings.
- Advanced source identity and installation instructions to `1.1.0-rc.4`,
  added private RC.4 notes and extended release plumbing while preserving RC.3
  protocol/ledger and `QWSG-053-F001` evidence.
- Integrated the exact reviewed 12-path endpoint allowlist through a clean
  fast-forward push to `main` at
  `ef513dde187e4119b6aa04a3439a879056f6cc69`.
- No runtime source was modified and no candidate artifact was constructed.

## Verification

PASS results:

- ordinary checkout build with explicit `-buildvcs=false` and RC.4 identity;
- two independent genuine no-`.git` exports with `GOFLAGS` unset, truthful
  `unknown` commit/date, byte-identical binaries and no Go `vcs.*` settings;
- explicit full commit/date injection, exact `qwsg version` output and
  byte-identity between checkout and export with identical controlled inputs;
- complete Go tests, repository-wide race tests, vet, formatting and focused
  build/runtime/assessment/setup packages;
- RC.4 validate-only release plumbing, full-commit/epoch/collision policy and
  immutable RC.3 protocol/ledger/finding assertions;
- Framework 21, Builder 38, lifecycle 28, diversion 36, active job, test-task,
  shell syntax and whitespace checks.

The sandboxed static systemd probe could not bind its private socket or connect
to the system bus, returned success, and is retained only as a host-local
limitation rather than product or external acceptance evidence. The release
builder itself was not changed and no RC.4 candidate bytes were constructed;
its deterministic policy is validated without crossing that explicit boundary.

Post-integration verification passed for the canonical checkout build, two
genuine no-`.git` exported-source builds, truthful default and exact explicit
provenance, controlled byte identity, release plumbing, full Go and race tests,
vet, formatting, shell and whitespace checks, Framework 21, Builder 38,
lifecycle 28, diversion 36, active-job/test-task validation, security/exclusion
and protected-evidence preservation. No RC.4 candidate construction or external
clean-host acceptance occurred.

## Rollback

Restore only the literal Task 054 targets from the verified snapshot after
Owner authorization. Remove only new files recorded absent. Never use broad
reset/clean, alter the private RC.3 roots, or touch Owner/historical content.

## Completion state

`complete` — the exported-source build determinism correction is integrated in
`main`; source identity is QWSG `1.1.0-rc.4`. No RC.4 candidate artifact has
been constructed and no RC.4 external clean-host acceptance has occurred.
