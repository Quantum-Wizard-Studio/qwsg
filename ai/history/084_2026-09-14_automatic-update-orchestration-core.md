# Task History 084: Automatic Update Orchestration Core

## Task metadata

- Task ID: `084`
- Task slug: `automatic-update-orchestration-core`
- Status: `complete`
- Date: `2026-09-14` UTC
- Agent: Codex
- Human authority: Project Owner explicit Task 084 APPROVE; Task 083 accepted.
- Preferred owner communication language: Hungarian, verified project configuration.
- Related prompt: `ai/archive_prompts/084_2026-09-14_automatic-update-orchestration-core.md`
- Dependencies: accepted Tasks 082 and 083.

## Objective, scope and exclusions

Implement the explicitly callable automatic update orchestration core over the
common authenticated engine, with policy/capability checks, structured evidence,
mutation boundary, rollback and conflict protection. Keep one session and close,
commit, push and verify the repository. No scheduler, timer, Guardian trigger,
maintenance/reboot orchestration, fleet/remote management, licensing backend,
release publication, v1.3.2, notification redesign or unrelated refactor.
Task 085 must not be started. Historical v1.3.1 remains immutable.

## Required reading and starting state

Read the project-local qwsg-job skill, AGENTS.md and all prompt Required Reading;
backup, engineering, delivery and security policies; Task 083 archived prompt and
history; architecture and Task 082 authenticated migration contract; common
update, authority, installed classification, policy, helper and fixture code.

Framework identity and idle lifecycle checks passed. Main, HEAD and fetched
origin/main were `abd7b1174ebd774dab5fa3d67bc537d0253125f8`; divergence 0/0,
worktree clean. Task 083 complete and archived; no Task 084 existed. Canonical
HTTPS origin matched validated configuration. Go: `go1.26.5 linux/amd64`.
Source modes and ownership inspected; new sources/docs retain inherited
non-executable modes. Existing module caches reused. No new dependencies,
services, infrastructure, production installation or credentials changed.

## Snapshot, rollback plan and risk

Before lifecycle or source mutation, protected external-to-Git local snapshot
`baseline.tar` captured tracked HEAD with 576 regular members. Archive SHA-256:

`8587c2adf4fe33855368478b3ffdaf5ad99b6a6771d3f6f397d5d9dc7d54e109`

Archive readability, archive hash and every member's SHA-256 were verified at
creation and again after implementation. Retain payload and manifest until
Owner acceptance. No payload, private path or host inventory is committed.
Restore only after validating both hash layers: extract into a new empty private
directory, compare the exact Task 084 changed-path allowlist to baseline, review
individual restoration/removal, preserve history and obtain authority before
destructive restoration. Never extract over live checkout or broadly reset/clean.
Rerun affected tests and lifecycle checks after an authorized restore.

High-impact integrity/rollback risk is mitigated by shared authority/package
primitives, unchanged helper security gates, explicit recovery outcomes and real
isolated transaction tests. Conflict guard scope is the canonical installation
update directory and process. Service coordination is deliberately conservative:
the production adapter requires verified inactive Guardian and does not start or
stop services. No production automatic capability source is introduced.

## Planned work and work performed

The Builder transaction installed exactly Task 084 using Owner-approved fields.
Inspection and snapshot preceded implementation. The smallest change preserves
the common engine and adds a coordinator plus trusted host adapter:

- `internal/automaticupdate`: synchronous lifecycle, immutable Task 083 capability
  and canonical policy evaluation, Task 082 authorization and exact-source
  understood migration eligibility, watermark, common staging/package checks,
  preflight, helper transaction, canonical validation and local commit. Ordered
  stages and bounded result fields expose decisions, versions, failure category,
  mutation knowledge and recovery outcome. No trigger or background goroutine.
- `internal/update/transaction.go`: common apply now reports entry to the first
  destination-changing operation and internal rollback outcome. Rollback errors
  are retained; restored managed files are checked for hash/mode or absence.
  Existing transaction metadata remains readable; private user data is excluded.
- `cmd/qwsg`: explicit code-only composition, no automatic CLI dispatch.
  `privileged-apply-report` shares every check with manual privileged apply and
  returns a bounded receipt. Missing/inconsistent receipt fails conservatively.
  Sudo is noninteractive. Recovery waits for the finite helper to finish rather
  than race a killed parent with its still-running privileged child.
- Process atomic guard and nonblocking private file lock reject re-entry and
  conflicting automatic calls. Locks remain held through validation/recovery.
  This is not distributed locking or manual-operation scheduling.
- Pre-mutation failures return failure without rollback. Apply's own successful
  restore is consumed without duplication. Later validation/commit errors use
  existing privileged rollback and source validation. Failed restoration is
  `rollback_failure`, never ordinary success/failure. Caller cancellation does
  not abandon rollback. No long observation window or crash-recovery daemon.

## Verification evidence

- Focused automatic orchestration tests: PASS. Cases include Community/zero
  capability, manual/unknown/conflicting policy, unsigned/tampered metadata,
  unsupported migration, source/target/platform/artifact/provenance mismatch,
  invalid package, watermark, preflight and helper refusal; installed file bytes
  and modes remain unchanged before mutation.
- Real common-engine fixture transaction tests: PASS for successful trace,
  failure during apply, post-validation failure, commit failure, successful
  rollback, explicit degraded rollback failure and cancellation after apply.
  Missing helper receipts and failed internal-rollback receipts remain degraded.
- Re-entry/concurrency and independent-descriptor flock contention tests: PASS;
  symlink/unsafe lock refusal and lock release tested.
- Extended frozen-old-client gate: PASS for automatic success, post-apply rollback
  and actual root-helper authority rejection, alongside all original Task 082
  Community/manual scenarios. Real staging, signature/provenance checks, apply,
  installed/config validation and rollback run only in isolated fixtures.
- `make fmt-check vet test`: formatting and full vet PASS; full Go suite PASS
  after the environment-specific TLS rerun below. Final preparation-evidence
  correction received focused tests and another format/vet/full Go run.
- Affected `go test -race`: PASS for automaticupdate, update, updateauthority,
  productcapability, updatepolicy, configuration and cmd/qwsg, including the
  frozen-old-client integration gate.
- Framework v2: 15 assertions PASS; bounded diagnostic runner: 11 PASS.
- Active lifecycle and diverted-task audit: PASS.
- Snapshot/member hash revalidation, reviewed scope/modes and diff whitespace: PASS.
- `make engineering-test`: checkout/export provenance PASS; 25 framework,
  36 diversion, 29 lifecycle and 49 task-builder assertions PASS.
- Final idle lifecycle and framework identity gates: PASS.

### Diagnosed attempts

`ENVIRONMENTAL ISSUE`: sandbox denied writing `.git/FETCH_HEAD`; the approved
scoped fetch succeeded and confirmed the baseline. Targeted index writes use
approved sandbox elevation because `.git` is read-only in the workspace sandbox.

`ENVIRONMENTAL ISSUE`: existing releasediscovery TLS test could not bind its
local fixture socket in sandbox. The unchanged full `make test` outside sandbox
passed. No test was skipped, weakened or modified to hide that failure.

## Documentation updates

Added `AUTOMATIC_UPDATE_ORCHESTRATION.md`, updated the architecture and milestone
indexes and replaced stale future-Task-084 wording in the capability/policy
contract. This history and completed archived prompt contain lifecycle evidence.

## Completion state and delivery

All mandatory implementation and verification gates passed. Task 084 is
complete; its prompt is archived without a successor and idle validators pass.
The reviewed Task 084 path set contains only the common transaction, automatic
core/adapter, tests and narrow documentation/lifecycle records. No source changed
after final validation. Integration uses targeted staging, commit, explicit
dry-run/fast-forward push and post-push fetch/integrity verification. Final identities
are reported in the Owner handoff rather than embedded self-referentially here.
Scheduler/Guardian triggering, active-service coordination, maintenance/reboot,
production entitlement and crash/long-observation orchestration remain future
scope. Task 085 has NOT been started or prepared. No release is published. No unresolved implementation defect is known.
