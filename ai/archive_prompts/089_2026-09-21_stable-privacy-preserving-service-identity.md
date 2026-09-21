# Current Engineering Task 089: Stable Privacy-Preserving Service Identity

## Task Metadata

- Task ID: `089`
- Task slug: `stable-privacy-preserving-service-identity`
- Status: `complete`
- Date opened: `2026-09-21` UTC
- Human authority: Project Owner, explicit Task 089 authorization in this session including implementation, lifecycle completion, commit and push
- Owner or lead-developer communication language: Hungarian

## Title

Stable Privacy-Preserving Service Identity


## Objective

Close only frozen Community C2 with stable deterministic privacy-preserving service identity, meaningful changes and explicit historical compatibility. Complete in this session; no Task 090.


## Scope

Service collector identity, directly required canonical inventory/comparison/drift compatibility, focused acceptance tests, minimum privacy/architecture documentation, C2 register and aggregate counts only if proven, Task 089 lifecycle, commit and push.


## Out of Scope

No general monitoring, service-health rules, expanded service states, secret/key lifecycle, external identity service, database, telemetry, dependencies, historical bulk migration, C3/C7/C8/C9/P1/P3/P4/P5 implementation, version bump, release, production mutation or next task.


## Authority Envelope

**Task targets and boundaries:** The Owner explicitly authorizes the smallest deterministic solution for C2, including existing privacy primitive reuse, canonical ordering and historical representation compatibility.
**Permitted external actions:** Fetch and task-scoped clean fast-forward push to canonical origin/main; no production mutation.
**Owner-reserved decisions:** Scope expansion, secret/key lifecycle, release/publication and other gates.
**Task-specific STOP conditions:** Material baseline mismatch; new secret lifecycle authority needed; standard security/privacy/rollback boundaries.


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

Verify expected HEAD 975789e60e9a87a224665693964d987b52f02a10 equals freshly fetched origin/main, main, clean worktree, idle lifecycle, latest complete Task 088. Read canonical docs/PRODUCT_1_0_SCOPE_FREEZE.md, collector, inventory, privacy primitive, comparison and tests before coding.


## Snapshot Requirements

Before mutations capture tracked baseline in protected external-to-Git /tmp snapshot with per-member SHA-256 manifest, archive checksum and readability verification; retain through Owner acceptance.


## Risk Assessment

Deterministic identifiers are pseudonymous, not anonymous: preserve existing privacy semantics and disclose dictionary guessing. Old ordinal identities must never become stable identities by reinterpretation. Keep historical evidence readable and immutable; explicitly refuse unsupported exact comparison or otherwise expose uncertainty.


## Planned Work

Verify baseline and snapshot; install approved lifecycle pair; inspect focused contracts; reuse existing deterministic privacy primitive for unit identity; sort protected IDs; define explicit historical boundary; test all seven cases early; stable-source mandatory validation; minimal documentation and C2 decision; archive idle, targeted commit/push and final verification.


## Rollback Plan

Verify snapshot archive and manifest; extract only into new private empty directory. Review exact Task 089 changed paths and intervening drift before bounded restoration from baseline. New lifecycle/test paths absent at baseline require explicit destructive removal authority. No broad reset/clean or history rewrite. Preserve evidence and rerun focused tests/framework/lifecycle/Git checks.


## Deliverables

Stable protected service identity and canonical ordering; historical compatibility behavior; seven deterministic acceptance cases; minimum privacy/architecture documentation; honest C2 status/totals; Task 089 history/archive.


## Verification

Repeated alpha/beta/gamma observation; reordered gamma/alpha/beta unchanged; beta disappearance and appearance; same-count beta to delta replacement; no raw names in canonical inventory/report; old ordinal/new stable boundary explicit. Formatting, vet, focused collector/inventory/comparison/drift/privacy tests, relevant Community integration, focused race if relevant, required full local engineering/build contract after source stabilizes, framework/lifecycle and snapshot integrity. Final clean main, commit/push, HEAD equals origin/main, divergence 0 0.


## Documentation Updates

Only minimum stable protected identity/privacy semantics, historical comparison boundary, C2 evidence/status and aggregate counts, engineering milestone and Task 089 prompt/history. Do not change other gate statuses.


## Completion Criteria

PASS only when all seven cases and mandatory validations pass, rollback verified, documentation truthful, lifecycle archived idle, reviewed commit pushed, clean worktree and HEAD == origin/main with divergence 0 0. Otherwise report PARTIAL/BLOCKED and leave C2 PARTIAL. Stop without another task.


## Owner Approval Requirements

Approved by Project Owner, explicit Task 089 authorization in this session including implementation, lifecycle completion, commit and push through the Engineering Task Builder on 2026-09-21 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.

## Completion record

Task 089 implementation and all seven acceptance cases passed. Mandatory final
Go, focused race, framework and engineering/build-contract checks passed; the
history records exact scope and evidence. C2 alone advances to PASS. Archive
this completed prompt without a successor and validate idle before reviewed
commit/push and final synchronization verification. No Task 090 is authorized.
