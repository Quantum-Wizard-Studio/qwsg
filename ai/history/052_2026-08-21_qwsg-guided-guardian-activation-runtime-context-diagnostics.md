# Task History 052: Approved Task

## Task metadata

- Task ID: `052`
- Task slug: `qwsg-guided-guardian-activation-runtime-context-diagnostics`
- Status: `active — source integration authorized; lifecycle closure not authorized`
- Date generated: `2026-08-21` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/052_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded. The Project Owner started Task 052 through the canonical `job` workflow on 2026-08-21 UTC.

After local implementation and verification, the Project Owner authorized one
exact 29-path source-integration commit and clean fast-forward push. This gate
does not authorize Task 052 lifecycle closure, RC.3 construction, external
acceptance, credentials, tags, releases, publication, or Task 053.

## Starting state

Verified at `2026-08-21T09:46:14Z`: repository root
`/home/qws/web/qwsg.quantumwizard.hu/qwsg`, ordinary user `attila`, branch
`main`, canonical HTTPS `origin`, HEAD and existing `origin/main` both
`6d3f79accd4d52b94c960eefa93e2f51fbc9a48c`, ahead/behind `0/0`, and empty
index. The expected authorized Task 051 closure/Task 052 lifecycle changes are
unstaged. The only unrelated path is the pre-existing untracked Owner-owned
`docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`; it was not read, copied,
hashed, modified, moved, staged, ignored, or packaged.

Task 051 is archived truthfully as `complete with disclosed limitations — NOT
READY FOR QWSG 1.1.0 RELEASE`; `QWSG-051-F001` remains OPEN/BLOCKING at guided
Guardian activation. RC.2 source and archive identities remain respectively
`6d3f79accd4d52b94c960eefa93e2f51fbc9a48c` and
`73d045cbc5577d3e9921a44760ba316d2094cf13fafe82f873be9f3600547315`.
RC.1/Task 049, v1.0.0 and LICENSE preservation hashes matched the Task 051
closure baseline. No RC.3 tag, release note or repository artifact existed.

The source audit reproduced the blocker architecture: assessment validates the
effective UID's canonical runtime directory and supplies its trusted
`XDG_RUNTIME_DIR` to fixed probes, while `userservice.New` constructs the same
bounded runner without trusted runtime context. Guided activation therefore
cannot reliably reach the manager that readiness already proved reachable.

## Snapshot

Private mode-0700 snapshot:
`/tmp/qwsg-task052-implementation.ECFPGI`. It contains a verified readable
tracked-HEAD archive, the authorized unstaged lifecycle/evidence state, Task
052 Builder source, target absence records, tool/baseline identities, and
bounded restore instructions. Owner content is represented only by excluded
stat metadata.

## Work performed

Implemented `internal/userruntime` as the single trusted boundary for
effective-UID `/run/user/<uid>` derivation, lstat directory/symlink/type,
ownership and private-mode validation. It returns stable valid/missing/unsafe
outcomes and a context whose only projection is canonical
`XDG_RUNTIME_DIR=/run/user/<decimal-uid>`. The package also owns the shared
fixed `systemctl --user is-system-running` classifier for reachable,
transient, unavailable, timeout, output-limit, failed-probe and unrecognized
outcomes without reading stderr.

Refactored assessment to consume the shared validator/classifier while
preserving every Task 047/050 classification and evidence token. Hardened
`runner.Bounded` so each fixed operation accepts at most one canonical trusted
XDG entry; duplicate, noncanonical, HOME, DBus and other environment entries
fail before execution, and ambient environment remains replaced by fixed
PATH/C locale.

Refactored `userservice.Controller` to derive the current context itself and
attach it only to three fixed absolute systemctl operation IDs. Activation now
proves user-manager reachability before the unchanged ordered
`daemon-reload` then `enable --now qwsg-guardian.service` mutation. Typed
privacy-safe errors preserve the runtime-context, manager, daemon-reload or
enable/start stage and distinguish missing/unsafe context, unreachable/
transient/unrecognized/failed manager probe, timeout, output limit and fixed
operation failure. No raw stderr, environment, path, username or host data is
returned.

Guided setup now renders the typed failure as a stage-specific explanation,
states that configuration was preserved, and supplies a bounded
`install --check`/`readiness` plus resumable `setup` action. It supplies no
guessed systemctl remediation. The success path still waits for fresh
integrity-checked Guardian evidence and independently invokes unchanged
readiness.

Added focused context, manager-state, runner-environment, exact-operation,
early-stop, success/failure, timeout/output-bound, deterministic diagnostic and
guided configuration-preservation fixtures. Updated Smart Install/Setup
architecture, installation, setup, troubleshooting, security/privacy and
known-limitations documentation.

Advanced only source metadata and collision-aware release validation to
`1.1.0-rc.3`, added private RC.3 notes, and updated the canonical installation
filename examples. No RC.3 archive, sidecar, tag, release, acceptance record or
publication was created. RC.1/RC.2 release notes, protocols and evidence were
not changed by Task 052.

## Verification

Pre-change PASS: build, full tests, repository-wide race tests, vet, formatting,
RC.2 release-check, shell syntax, static user-unit verification, Framework 21,
Builder 38, lifecycle 28, diversion 36, active-job/test-task and Git whitespace
checks. The static unit verifier emitted its expected sandbox inability to bind
a private socket/system bus but returned PASS for the unit itself.

Post-change PASS: focused tests for `internal/userruntime`, `internal/runner`,
`internal/assessment`, `internal/userservice` and `cmd/qwsg`; build; full tests;
repository-wide race tests; vet; formatting; RC.3 validate-only release and
collision plumbing; shell syntax; static user-unit verification; Framework 21;
Builder 38; lifecycle 28; diversion 36; active job; test-task audit; Git
whitespace. The development-host `build/qwsg install --check --format json`
remained read-only, reported RC.3 source-build identity and preserved the
existing assessment schema/classifications; its user-manager result was
environment-limited host-local evidence, not external acceptance.

Security/diff audit found no shell, sudo, loginctl, arbitrary executable/argv,
ambient XDG/DBus/HOME, privilege/system/lingering/package/network/remediation,
raw host output, credential or external-host action. The Git index remained
empty. An isolated snapshot extraction reproduced the exact baseline
controller and proved the two Task 052-created paths absent, validating bounded
rollback without touching the worktree.

Preservation PASS: Task 051 terminal prompt/history and RC.2 acceptance record
match the pre-Task-052 snapshot; RC.1 acceptance/protocol hashes, Task 049
prompt/history hashes, RC.2 protocol hash, v1.0.0 dereferenced target and
LICENSE hash remain exact. Owner-owned path stat metadata is unchanged and its
content was never read/copied/hashed. No RC.3 artifact or tag exists.

Source-integration pre-commit PASS: the snapshot and baseline revalidated; the
literal file-level staged set was exactly the 29 Owner-authorized paths under
`--no-renames`; no tracked change remained unstaged; only the excluded
Owner-owned file remained untracked. Complete staged diff, modes, whitespace,
binary, secret/private-key/access-key, payload/path and Owner-content audits
passed. Build, focused setup/runtime tests, full tests, repository-wide race,
vet, formatting, RC.3 release-check, shell/static unit, Framework 21, Builder
38, lifecycle 28, diversion 36, job and test-task gates all passed.

## Rollback

Verify later-edit absence, restore only literal tracked targets from the
snapshot, restore the preserved lifecycle files, and remove only Task
052-created paths with recorded pre-task absence. Never use broad Git recovery
or touch Owner content, RC.1/RC.2, external state, v1.0.0 or LICENSE.

## Completion state

`active — source integration gate authorized; lifecycle closure remains separately gated`

Task 052 is not lifecycle-complete and `QWSG-051-F001` is not externally
verified/corrected. The next authorized operation must be an exact path-based
staging/review/commit/push gate for the Task 051 closure, Task 052 lifecycle,
product correction, tests and RC.3 source metadata. After integration, Task
052 closure requires another explicit Owner gate. A future clean-host task
must construct RC.3 twice from that exact clean commit and restart at
Checkpoint 01; historical host-independent evidence may be cited only as
chronology while every artifact-identity and product-dependent external gate
is re-executed. Task 053 is not created here.
