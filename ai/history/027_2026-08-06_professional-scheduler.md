# Task History 027: Professional Scheduler

## Task metadata

- Task ID: `027`
- Task slug: `professional-scheduler`
- Status: `complete`
- Date generated: `2026-08-06` UTC
- Human authority: Project Owner
- Related prompt: `ai/prompts/027_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder installed Task 027 as the sole active approved
task over the completed Task 026 baseline. Implementation was then started by
the canonical `job` workflow. No successor task was created.

## Starting state

- Repository identity and configured root matched QWSG.
- Active branch was `main`; local HEAD and `origin/main` were both
  `0a8a5c7a81f9597993c44bc932e76b89edb0b87a`.
- Task 027 was the sole active approved prompt and Task 026 was the latest
  completed canonical implementation baseline.
- Framework, lifecycle, task, Go test, vet, and formatting baselines passed.
- The working tree already contained unstaged and untracked Owner work from
  Tasks 025 and 026, QWCS architecture, Builder sources, and a historical
  backup. That content was recorded and preserved; it was neither staged nor
  treated as Task 027 output.

## Snapshot

The mandatory rollback-capable implementation snapshot is:

`/tmp/qwsg-task027-implementation-HzGxHfXX`

It contains the exact pre-task target payload, absence records for new Task 027
targets, repository/Git/lifecycle/validation evidence, permissions, a bounded
restore procedure, a deterministic manifest, and SHA-256 checksums. Archive
listing, readability, and checksum verification passed before implementation.

## Work performed

- Added `internal/scheduler` with Scheduler Model, Evaluation, State, Event,
  Execution Request, and Execution Result 1.0 contracts.
- Implemented pure deterministic interval and calendar evaluation using the
  half-open `(last observation, current observation]` window, explicit interval
  anchors, injected clock observations, IANA time-zone resolution, DST
  selection, deterministic next-run calculation, and explicit clock-failure
  outcomes.
- Implemented enabled/disabled and Command-scope applicability behavior,
  missed-run policies, descending priority, stable tie-breaking, configured
  concurrency, overlap protection, and a one-occurrence bounded replacement.
- Implemented finite deterministic retry planning, process-session restart
  recovery, interrupted results, and prevention of already completed occurrence
  retries.
- Added stable content identities, strict validation, deterministic canonical
  JSON, bounded results and occurrences, and Policy/Command/Pipeline provenance.
- Added an explicit one-cycle adapter. It receives every runtime dependency,
  resolves only existing Canonical Command profiles, and invokes only the
  existing Pipeline executor. It owns no recurring loop or service lifecycle.
- Added a versioned file store with strict decoding, SHA-256 integrity,
  restrictive permissions, synchronized temporary writes, atomic replacement,
  and corrupt-state rejection.
- Added non-blocking process-safe local file locking with validated owner
  identity, contention behavior, and explicit release.
- Added focused deterministic, interval, misfire, clock-discontinuity,
  priority, concurrency, time-zone, DST, overlap, retry, restart, persistence,
  integrity, permission, locking, adapter, Policy-traceability, serialization,
  privacy, and Pipeline compatibility tests.
- Created `docs/architecture/CANONICAL_SCHEDULER.md` and updated only the
  directly affected permanent architecture, map, roadmap, README, Policy,
  Configuration, Command, Report, lifecycle, and history references.

## Engineering decisions

- Task 027 uses the approved option B boundary: pure Scheduler Engine plus one
  explicitly invoked local cycle. A daemon or installed service remains out of
  scope.
- The first interval evaluation establishes state and produces no immediate
  work. This avoids inventing an activation time in Schedule Definition 1.0.
- Empty Schedule Check scope means the complete Command profile. Non-empty
  scope is `inapplicable` because Command Definition 1.0 cannot represent it;
  the Scheduler never broadens the request silently.
- Reservations are persisted before Pipeline invocation. A later process
  session turns any surviving active reservation into an interrupted result and
  applies only the configured finite retry policy.
- Policy outcomes captured from Command Execution remain traceability facts,
  never scheduling or action authority.

## Failed attempts and corrections

Initial focused restart and adapter tests exposed a mutation bug in evaluation
identity calculation: a shallow copy shared nested slice storage and cleared
returned request references. Identity normalization now operates on a fresh
deep value. The same review found that restart retry planning needed to enforce
the configured maximum-attempt contract and that multiple candidates for a
forbid-overlap schedule needed reservation-aware selection. Both were corrected
and covered by focused tests before repository-wide validation.

## Verification

All completion gates passed:

- focused Scheduler tests and `make test`;
- `make build` and repository-wide `go test -race ./...`;
- `make vet` and `make fmt-check`;
- `ai/scripts/framework-check.sh --run-validations`;
- `make engineering-test` (Framework, diversion, lifecycle, and Builder test
  suites);
- active `bin/job --check`, lifecycle, and test-task validation;
- `git diff --check` and `git diff --cached --check`;
- Scheduler source-boundary and canonical Command/Pipeline dependency audits;
- snapshot archive listing, readability, and every `SHA256SUMS` checksum.

Final evidence showed no staged changes. The pre-existing Owner work remained
unstaged/untracked alongside Task 027 changes; no commit or push was performed.

## Rollback

Use only the guarded `RESTORE.md` procedure in the snapshot. It restores exact
pre-existing Task 027 targets and removes only targets whose verified absence
was recorded, after collision checks. It does not reset, clean, or overwrite
unrelated Owner work. The snapshot remains outside Git through Owner
acceptance.

## Completion state

`complete — ready for canonical archive`
