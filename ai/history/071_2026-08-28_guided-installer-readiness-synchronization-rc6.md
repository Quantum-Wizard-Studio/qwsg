# Task History 071: Approved Task

## Task metadata

- Task ID: `071`
- Task slug: `guided-installer-readiness-synchronization-rc6`
- Status: `active — implementation and verification in progress`
- Date generated: `2026-08-28` UTC
- Human authority: Project Owner: Attila
- Preferred owner communication language: English
- Related prompt: `ai/prompts/071_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this first canonical Task 071 pair from the complete Owner-approved, explicitly renumbered remediation specification. Exact `APPROVE`, protected input validation, Framework and lifecycle consistency passed; execution began under Framework 2.0 authority.

## Starting state

- Verified ordinary user, exact QWSG root, Framework 2.0.0, canonical HTTPS origin, `main`, and fresh `HEAD == origin/main == 641be36ae3098042331858ec2cc7b8c5108967b6` with `0 0` divergence.
- Before Builder installation the full worktree/index was clean and lifecycle canonically idle after Task 070; no prior Task 071 prompt/archive/history existed. After installation only the expected Task 071 prompt/history were untracked.
- Verified RC.5 artifact SHA-256 `ca3123f313c2fb95138fceb0fc2af17c5a8295b28e7b13e60f5d7fd57ed19faf`, source `3f2f7822473d3d97c38c171f006e8b263250321c`, Task 070 READY classification and absence of final `v1.2.0*` tag.
- Recorded Owner manual RC.5 evidence without contacting OVH/Contabo: clean install/reboot/uninstall preservation passed; preserved-state reinstall installed/activated successfully but returned exit 4 at 87% before new canonical evidence, then mandatory readiness passed with only optional notification Partial and no corruption.

## Snapshot

- Protected snapshot `/tmp/qwsg-task071-execution.3zJbjO`, directory mode `0700`, files `0600`, retained through Owner acceptance/release rollback-window closure.
- Captured complete tracked HEAD plus exact Builder-created Task 071 prompt/history before-images with README, bounded rollback and deterministic SHA-256 manifest. Checksums/readability passed; isolated protected extraction reproduced RC.5 and prompt/history byte comparisons passed.

## Work performed

- Read the active prompt, required governance/configuration and backup policy plus relevant installer/setup/readiness/Guardian and Tasks 065/068/069/070 evidence.
- Confirmed root cause: guided install immediately assessed readiness after activation, while setup had a separate bounded poll. A transient missing/not-running-fresh canonical finding caused exit 4 before completion; optional notification Partial was not causal.
- Extracted one cancellable bounded waiter shared by guided setup and guided install. It captures the pre-activation evidence ID and accepts only a different integrity-checked running/current canonical state, preventing preserved stale evidence from passing.
- Guided install renders readiness at 87%, waits up to validated cycle timeout plus the existing five-second margin, then assesses all domains. Timeout/cancellation preserves exit 4 and no summary; optional-only Partial reaches completion/100%/summary/exit 0.
- Added deterministic tests for clean immediate/delayed evidence, preserved-state stale rejection/new evidence, bounded timeout/cancellation, mandatory NotReady, optional Partial, 97%/100%/summary, EN/HU/DE guidance and secret-free diagnostics.
- Advanced the declared fail-closed RC.2 migration target and private release identity to RC.6; updated release plumbing, migration/notification regressions and relevant installer/operations/security/troubleshooting/release/acceptance documentation while preserving RC.5 evidence.

## Verification

- Protected Builder input `0700`/`0600`; `task-builder.sh --check-input`: VALID. Builder installation and active lifecycle: PASS.
- Snapshot integrity, archive readability and isolated restore rehearsal: PASS.
- Focused installer/migration/notification tests, full `go test ./...`, `go vet`, formatting and build-contract checks: PASS.
- Full `go test -race ./...`: PASS across every package.
- Framework v2 (15), bounded runner (11), lifecycle (29), Builder (49), active Task 071 and diverted-test-task audit: PASS. One initial verification command referenced a nonexistent legacy lifecycle test filename; this was a test-invocation error and the canonical `test-next-task.sh` suite subsequently passed.
- Release plumbing and cross-umask/source-mode reproducibility regression: PASS.

## Rollback

Verify `/tmp/qwsg-task071-execution.3zJbjO/SHA256SUMS` and follow its bounded `ROLLBACK.md`; restore only explicit reviewed Task 071 paths from isolated extraction and rerun focused/full/Framework/lifecycle checks.

## Completion state

`active — implementation and verification in progress`
