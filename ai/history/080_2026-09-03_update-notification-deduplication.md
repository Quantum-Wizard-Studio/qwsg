# Task History 080: Approved Task

## Task metadata

- Task ID: `080`
- Task slug: `update-notification-deduplication`
- Status: `active — Owner-authorized OOM remediation and Task 080 acceptance in progress`
- Date generated: `2026-09-03` UTC
- Human authority: Project Owner — explicitly authorized Task 080 and supplied the exact engineering scope with APPROVE on 2026-09-03 UTC.
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/080_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured Owner input. Exact `APPROVE` was received, active lifecycle validation passed, and implementation began under Framework 2.0 Standard Execution Authority. Task 081 remains prohibited.

## Starting state

- Verified repository root, Framework 2.0.0, branch `main`, canonical HTTPS origin, and fresh fetched `HEAD == origin/main == 1b3fb747e91a3f01256c22250e31fe3f99e242fd`; divergence `0/0`; tracked/index worktree clean; lifecycle canonical idle after complete Task 079 before Builder installation.
- `VERSION` is `1.2.0`; the exact Task 079 implementation and closure identities match the Owner baseline.
- Repository, lifecycle, user, UTC time, mode, ownership and ACL evidence was captured. User-service inspection from the restricted sandbox could not connect to the user bus and was classified `ENVIRONMENTAL ISSUE`; no product correction was made from that observation.
- Read the required governance/configuration documents, active prompt, Task 046/068 notification records, Task 074–079 release records, release/awareness/Guardian/SMTP/configuration code, tests, and EN/HU documentation before design and product changes.

## Snapshot

- Pre-intake snapshot `/tmp/qwsg-task080-intake.kxiFwc`: mode `0700`; complete mode-`0600` Git bundle SHA-256 `c1c7800022979817f3c60fccd6f03505a8542ba1d20fdd90e3c960c9b3fdee83`; readable lifecycle archive SHA-256 `8f2c7e0267dc102ae5eaaf4ea23d84b0fff6a9861f1656971cec694ce2393402`; bundle completeness, archive listing and permissions passed before Builder mutation.
- Execution snapshot `/tmp/qwsg-task080-execution.uZEkhN`: mode `0700`, all evidence mode `0600`; complete Git bundle SHA-256 `c1c7800022979817f3c60fccd6f03505a8542ba1d20fdd90e3c960c9b3fdee83`; target archive SHA-256 `771861ca20c276770ded493c7a50bbcf66edb0fbc768f06064d03b74a61921e6`; active prompt/history, ACL, scope and bounded rollback evidence are covered by the verified `SHA256SUMS`. Bundle verification, archive readability, every checksum and mode passed before implementation targets changed.
- Both snapshots exclude credentials, runtime state and private keys and are retained through Owner acceptance.
- OOM remediation snapshot `/tmp/qwsg-task080-oom-remediation.Ts9Eyg`: mode `0700`; complete baseline Git bundle, scheduler source/unit before-images, installed binary/unit, exact private production scheduler directory, config SHA-256-only evidence and pre-mutation service properties. Core checksums, bundle completeness and private modes passed. Rollback restores the exact saved binary/unit/scheduler directory after stopping only the Guardian, then starts and verifies that unit; unrelated state/configuration is untouched.

## Work performed

- Extended `qwsg.update-awareness/1` with a backward-compatible optional successful notification record. Its SHA-256 identity is derived only from authenticated source, stable channel, release version, artifact digest and trusted signing key ID; timestamps and cache validators cannot create duplicate release identities.
- Added atomic lock-protected post-success recording to the existing private awareness store. Failure before SMTP acceptance writes nothing; persistence uncertainty after acceptance is surfaced rather than falsely claiming deduplication.
- Added `internal/updatenotification` for strict eligibility, localized EN/HU rendering, configured transport dispatch, restart-safe persistent suppression and in-process non-overlap. It contains no discovery, artifact or installation operation.
- Composed notification only after the Guardian's existing authenticated Manager check. Explicit `update.policy=notify` plus enabled Community email is required. Existing SMTP configuration, credential loading and `accepted`/`delivered` result semantics are reused; manual check/status are not wired to notification.
- Added focused deterministic eligibility, current/older/untrusted, first/same/restart/different, success/failure/retry, disabled-policy, persistence-error, private-state and concurrency coverage. Updated architecture, system map, awareness state, operations, security/privacy, README, changelog and EN/HU operator guidance.
- After the production-health blocker report, the Owner explicitly expanded Task 080 for bounded Guardian OOM diagnosis/remediation. Structural production evidence proved the release-check and notifier were not the cause: the persisted Scheduler state contained 984 results, averaged 22,815 encoded bytes per result with 375 policy IDs, and had grown to 22,451,945 bytes plus abandoned temporary writes. Repeated load/clone/hash/save JSON allocations combined with the normal Guardian working set reached `MemoryMax=128M`; restart recovery appended another result and sustained the loop.
- Corrected the underlying Scheduler persistence defect without increasing resource limits: retained results are capped at 64, encoded state is capped at 8 MiB, and load checks size before allocation. Added deterministic retention and oversized-before-decode regressions. `GOMEMLIMIT=64MiB`, `MemoryMax=128M`, `TasksMax=32`, release scheduling and Task 080 behavior remain unchanged.

## Verification

Builder input, metadata, prompt/history identity, approval state, lifecycle installation, both snapshots and initial Framework/Git checks: PASS. Initial focused Go test using the default cache was blocked by a read-only sandbox cache and classified `ENVIRONMENTAL ISSUE`; the identical run with private `/tmp` Go caches passed `internal/updateawareness`, `internal/updatenotification`, `internal/guardian`, and `cmd/qwsg`.

- `make fmt-check vet test`: PASS with unrestricted module/loopback access after the restricted run was correctly classified as an environmental cache/network denial.
- `go test -race ./...`: PASS for every package, including the new notification concurrency test.
- `make engineering-test`: PASS; checkout/export build provenance, Framework 25, diversion 36, lifecycle 29, and Builder 49 assertions passed after exact-path staging.
- Release reproducibility and release-authority reproducibility: PASS. Protected final artifact SHA-256 remains `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`; annotated `v1.2.0` tag object remains `ac395b568b8e1f83c0ef85c9aa02f98c15402af0` and peels to `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`. `make release-check` refused at its intentional clean-output precondition because the immutable published artifact exists; classified `EXPECTED BEHAVIOR`.
- Isolated official production-index check at `/tmp/qwsg-task080-production-acceptance.9IXy9N`: authenticated current `1.2.0` PASS; awareness mode `0600`, SHA-256 `8639f36ae7798f4f63cf059ca4c28c9b6ba5fb98b61263a8e6c0aee33f3deb03`; no `last_notification`, artifact, installation, SMTP or configuration state was created. Unrestricted `strace -f -e trace=network` around `update status` showed no network syscalls and status remained current.
- Production Guardian health gate: FAIL. Read-only systemd evidence at 22:24–22:25 UTC showed repeated `oom-kill` exits, `MemoryPeak=MemoryMax=134217728`, restart counter `1839`, and transition from `active/running` to `activating/auto-restart` with `Result=oom-kill`. The production service was not restarted, reconfigured or otherwise mutated by Task 080. This is a material pre-existing production baseline variance outside update-notification scope; completion, commit, push and lifecycle closure stopped rather than misreporting health.

## Rollback

Verify snapshot checksums and archive readability, extract only to an isolated directory, compare current hashes, then restore only reviewed Task 080 before-images and remove only paths recorded absent before Task 080. Preserve unrelated Owner work. Never extract over the live tree or use broad reset, checkout, clean, wildcard deletion, or published-history rewriting. After publication, use a new exact-path revert commit and rerun focused/full/race/framework/lifecycle/privacy/Git checks.

## Completion state

`active — Owner-authorized root-cause correction is implemented locally; production activation, stability acceptance, final integration and closure remain.`
