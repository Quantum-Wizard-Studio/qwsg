# Current Engineering Task 090: Local Evidence Interruption Integrity & Recovery

## Task Metadata

- Task ID: `090`
- Task slug: `local-evidence-interruption-integrity-recovery`
- Status: `complete`
- Date opened: `2026-09-21` UTC
- Human authority: Project Owner explicitly authorizes Task 090 C3 implementation, tests, documentation, lifecycle closure, commit and push in this session.
- Owner or lead-developer communication language: Hungarian

## Title

Local Evidence Interruption Integrity & Recovery


## Objective

Close only frozen Community 1.0 C3 if all interruption integrity and recovery requirements are implemented and proven. Complete in one session; do not start Task 091.


## Scope

Local inventory/evidence persistence, safe lock recovery, interrupted write and retention handling, bounded evidence/checkpoint reads, directly necessary durability ordering, focused tests and minimal documentation. Canonical scope: docs/PRODUCT_1_0_SCOPE_FREEZE.md.


## Out of Scope

C7 update/rollback semantics; C8 general readiness closure; C9 release/version/signing; Pro entitlement; privilege configuration; production mutation; databases, generalized WAL, remote storage, telemetry, new services and unrelated persistence.


## Authority Envelope

**Task targets and boundaries:** Smallest deterministic C3 implementation and nine-case acceptance matrix; no general persistence redesign.
**Permitted external actions:** Fetch, reviewed commit and fast-forward push to canonical origin/main.
**Owner-reserved decisions:** Scope expansion, release, deployment and other gate changes.
**Task-specific STOP conditions:** Material baseline difference; unsafe ambiguity must preserve evidence and fail closed. Leave C3 PARTIAL if full closure is not proven.


## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/core/17_EXECUTION_MODEL.md`
- `ai/core/18_BOUNDED_DIAGNOSTIC_RUNNER.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

Verify main and fetched origin/main at d8b8e598aebb25cbd126227a890b5d346faf9f0a, clean worktree, divergence 0/0, framework/lifecycle valid and latest completed task 089. Verified idle; install this Owner-authorized task pair through Builder.


## Snapshot Requirements

Before mutation capture protected external-to-Git tracked baseline archive, SHA-256 manifest, readability/member verification and restore notes; retain until Owner acceptance.


## Risk Assessment

Prevent concurrent writers, loss of prior evidence, false commit success, unbounded recovery and privacy/permission regressions. Preserve ambiguous artifacts and refuse explicitly. Backward-compatible evidence meaning; no historical bulk rewrite.


## Planned Work

Map actual committed/lock/temp/retirement/sync/checkpoint transaction states before source edits. Add deterministic acceptance at real boundaries; implement minimal recovery; focused tests then stable-source mandatory validation once. Update only C3 if justified, archive idle, commit/push and verify. At half budget freeze architecture; last quarter blockers/validation/closure only.


## Rollback Plan

Verify protected snapshot hash and members. Extract into new private directory, compare exact changed paths, restore only reviewed bounded targets after drift/authority checks. Preserve new artifacts until approved disposition. No broad reset, clean or published history rewrite. Revalidate tests, lifecycle and Git.


## Deliverables

Minimal implementation and tests; accurate persistence contract; individually reported Cases 1–9; C3 evidence and gate decision; finalized prompt/history and clean pushed idle repository.


## Verification

Cases 1 normal canonical persistence; 2 A preserved before B commit; 3 B recognized after durable commit despite cleanup interruption; 4 stale lock recovery; 5 live writer exclusion; 6 interrupted retention; 7 bounded reads at limit+1; 8 malformed/incomplete refusal; 9 idempotent repeated recovery. Run gofmt/vet/full Go tests, focused race, inventory/privacy/retention/resource regressions, engineering/build contract, framework/v2/diagnostic/lifecycle checks. Verify rollback, permissions, privacy, exact gate totals and clean pushed Git.


## Documentation Updates

Minimal directly affected persistence/recovery contract, canonical C3 gate only if proven, engineering milestone, Task 090 prompt/history. C7 and C8 unchanged.


## Completion Criteria

PASS only with all nine cases and mandatory validation passing, C3 accurate, archive idle, worktree clean, commit pushed, HEAD=origin/main and divergence 0/0. Otherwise truthful PARTIAL/BLOCKED; never sacrifice integrity. Expected C3-only totals Community 6 PASS/2 PARTIAL/1 MISSING; Pro total 7 PASS/4 PARTIAL/3 MISSING. Final report covers contract, individual cases, validation, integrity/security, repository and important deferred findings. Stop without successor.


## Owner Approval Requirements

Approved by Project Owner explicitly authorizes Task 090 C3 implementation, tests, documentation, lifecycle closure, commit and push in this session. through the Engineering Task Builder on 2026-09-21 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.

## Completion evidence

Task 090 implemented and proved all nine C3 acceptance cases and mandatory local
validation. C3 alone is PASS; C7/C8/C9 and all other gates are unchanged. Detailed
evidence is in `ai/history/090_2026-09-21_local-evidence-interruption-integrity-recovery.md`.
Archive idle without any successor. Reviewed commit/push and final Git identity
are recorded in the Owner handoff.
