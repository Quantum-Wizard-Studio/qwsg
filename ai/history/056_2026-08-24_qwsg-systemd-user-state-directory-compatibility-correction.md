# Task History 056: Approved Task

## Task metadata

- Task ID: `056`
- Task slug: `qwsg-systemd-user-state-directory-compatibility-correction`
- Status: `complete with disclosed limitations — external RC.5 acceptance pending`
- Date generated: `2026-08-24` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/archive_prompts/056_2026-08-24_qwsg-systemd-user-state-directory-compatibility-correction.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded; implementation has not started.

## Starting state

On 2026-08-24 UTC, `HEAD`, `origin/main`, and direct Forgejo `main` were
`029e01bc2e85f524f72e4598135cc29daf336dca`, ahead/behind `0/0`, on `main`
with an empty index and clean tracked tree. Task 056 was the sole approved
active task; its Builder-created prompt/history and the excluded Owner-owned
blueprint were the only untracked paths. Source identity was `1.1.0-rc.4`.
Task 054 and Task 055 source commits were ancestors. Task 055 remained complete
with disclosed limitations and NOT READY; QWSG-055-F001 remained OPEN/BLOCKING.

## Snapshot

Private mode-0700 implementation snapshot:
`/tmp/qwsg-task056-implementation.jVXTTl`. It contains a readable exact
tracked-HEAD archive, Task 056 lifecycle before-images, affected target copies,
RC.5 notes absence evidence, Builder identities, protected-history hashes, and
literal bounded rollback instructions. Owner content, credentials, candidate
bytes, caches and private host identity are excluded.

## Work performed

- Confirmed setup writes configuration before its activation boundary and the
  activation callback then performs trusted runtime-context validation,
  user-manager reachability, daemon-reload and enable/start.
- Confirmed Guardian/runtime storage uses the canonical `localStateRoot`
  contract and `operatorstate.EnsurePrivateRoot` already rejects symlinks,
  unsafe components, wrong types, ownership and modes without repair.
- Added canonical state-root preparation immediately before activation. Unsafe
  preparation now stops before every systemd operation with deterministic,
  privacy-safe stage-specific guidance.
- Removed the primitive's post-creation chmod so a raced or restrictive-mode
  result is rejected rather than repaired; mode-0700 creation is followed only
  by strict validation.
- Added clean-state ordering/idempotence and fail-closed symlink, unsafe
  component, wrong-type, unsafe-mode and ownership regression coverage.
- Retained the packaged systemd unit and all hardening directives unchanged;
  pre-creation is sufficient for systemd to encounter a real state directory.
- Advanced source/release-plumbing identity to `1.1.0-rc.5` and replaced the
  obsolete RC.4 initial-ledger assertion with immutable terminal-failure
  chronology checks. No candidate was constructed.

## Verification

- Focused `cmd/qwsg`, `internal/operatorstate`, and `internal/userservice` tests
  pass. They prove preparation-before-activation, safe mode-0700 ownership and
  non-symlink state, idempotence, state/service-root alignment, and fail-closed
  symlink, component, type, mode and ownership behavior without mutation.
- Existing userservice fixed operations, trusted Task 052 runtime environment,
  timeout/output bounds and deterministic stage errors pass unchanged.
- Full `go test ./...`, repository-wide `go test -race ./...`, `go vet ./...`,
  formatting and canonical checkout build pass with source identity
  `QWSG 1.1.0-rc.5`.
- Two genuine no-`.git` ordinary exports, explicit-identity export and checkout
  build-contract comparisons pass with no GOFLAGS workaround. RC.5 release
  validate-only and release-plumbing checks pass; no archive or sidecar was
  constructed.
- Assessment/readiness, Guardian, setupflow, notification, SMTP and Runtime
  Service regressions pass. Controlled package-fixture install/uninstall passes
  and preserves user configuration/state. Shell syntax and static systemd-unit
  verification pass; the unit is unchanged.
- Framework, Builder, lifecycle, diversion, active-job and test-task checks
  pass. Protected RC.1/RC.2/failed RC.3/failed RC.4, findings, v1.0.0 and LICENSE
  evidence remain unchanged. Owner content remained excluded.
- Local validation does not prove external systemd behavior. QWSG-055-F001
  remains historical OPEN/BLOCKING pending a separately authorized RC.5
  candidate and fresh Checkpoint 01 clean-host acceptance.

## Rollback

Use only the literal Task 056 targets and absence records in
`/tmp/qwsg-task056-implementation.jVXTTl` after Owner authorization. The exact
tracked baseline archive is readable, preserved before-files match the starting
state, and the RC.5 notes path was absent. Never reset/clean broadly, touch
Owner content or historical evidence, or remove unrelated files. After bounded
restore, rerun focused/full/security/governance/protected-hash validation.

## Completion state

`complete with disclosed limitations — external RC.5 acceptance pending`.
Task 056 implementation is complete, and the state-directory compatibility
correction is integrated into `main` as
`fc52295f6cc5b61078b32535230bdc704168d13c` (`fix: prepare Guardian state
before systemd activation`). Local and regression validation passed. The
packaged systemd user unit remains unchanged, and the source version is QWSG
`1.1.0-rc.5`.

No RC.5 candidate artifact has been constructed, and no RC.5 external
clean-host acceptance has occurred. QWSG-055-F001 remains historical `OPEN,
BLOCKING` until a separately authorized RC.5 candidate passes fresh external
clean-host acceptance from Checkpoint 01 on a fully reinstalled disposable
VPS. Task 056 completion does not claim external correction or release
readiness.
