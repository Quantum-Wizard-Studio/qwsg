# Current Engineering Task 084: Automatic Update Orchestration Core

## Task Metadata

- Task ID: `084`
- Task slug: `automatic-update-orchestration-core`
- Status: `complete`
- Date opened: `2026-09-14` UTC
- Human authority: Project Owner explicit Task 084 APPROVE in session; Task 083 accepted
- Owner or lead-developer communication language: Hungarian

## Title

Automatic Update Orchestration Core


## Objective

Implement deterministic explicitly callable automatic update orchestration using Task 082 authenticated update primitives and Task 083 capability/policy. Fit one session and complete committed pushed clean lifecycle closure.


## Scope

Automatic orchestration core, structured evidence, shared update transaction mutation and rollback evidence, safe helper integration, deterministic tests, narrow documentation and Task 084 lifecycle.


## Out of Scope

Scheduler, timers, Guardian trigger, maintenance windows, reboot, remote/fleet management, licensing backend, payments, dashboard, release publication, v1.3.2, notification redesign, unrelated refactors and Task 085. Historical v1.3.1 immutable.


## Authority Envelope

**Task targets and boundaries:** Only Task 084 orchestration core and necessary common update primitives, tests and documentation.
**Permitted external actions:** Fetch, targeted commit, dry-run and fast-forward push to canonical origin/main, post-push verification. No production installation or release.
**Owner-reserved decisions:** Scheduling, Guardian integration, licensing and release decisions.
**Task-specific STOP conditions:** None beyond standard; mandatory single-session budget and no scope expansion.


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

Verified main at abd7b1174ebd774dab5fa3d67bc537d0253125f8, fetched origin/main identical, divergence 0/0, worktree clean, idle after complete archived Task 083. Framework valid. Inspect existing authority, policy, transaction, helper and rollback before implementation.


## Snapshot Requirements

Protected /tmp/qwsg-task084-snapshot/baseline.tar captured tracked baseline before lifecycle mutation, with archive SHA-256 and per-member manifest verified. Retain until Owner acceptance; no payload in Git.


## Risk Assessment

High impact update integrity and rollback risk: reuse authenticated authority and common transaction; preserve helper reauthentication; expose mutation and rollback failure; deterministic fixture tests without live installation changes. Guard conflicting invocation.


## Planned Work

Inspect baseline; snapshot; design on existing primitives; implement core; focused deterministic/security/race tests and required local gates; narrow documentation; record evidence, archive Task 084, commit, push, fetch and final integrity verification.


## Rollback Plan

Verify archive and member SHA-256; extract only into a new empty private directory. Review exact changed-path allowlist for individual restoration. Never extract over live worktree or reset/clean broadly; preserve history, obtain authority before destructive restoration, rerun tests and lifecycle.


## Deliverables

Explicit state/result model with source/target, capability/policy decisions, mutation, rollback and failure category; gated authenticated staging/preflight/apply/validation coordination; deterministic rollback/conflict handling; required tests and documentation; completed history/archive and clean pushed main.


## Verification

gofmt, go vet, focused unit tests, affected race tests, Task 082/083 regressions, update transaction and rollback tests, full Go/configured framework/lifecycle gates. Cover capability absent, manual/unknown policy, signed compatible success, signature/migration/source/target/platform/artifact/provenance failures, invalid package, preflight failure, post-mutation failure, rollback success/degraded failure, reentry, helper rejection and Community manual regression. No real release/live installation.


## Documentation Updates

Narrow lifecycle, mutation boundary, rollback semantics, policy/capability and security contract; automatic authorization differs from automatic trigger; scheduler/Guardian deferred. Task 084 history and archived completed prompt.


## Completion Criteria

PASS only with all mandatory Task 084 acceptance checks, no scheduler/trigger/release, closed lifecycle, committed pushed clean main equal origin/main with 0/0 divergence. Report commit, identities, tests, architecture, boundaries, rollback and conflict semantics. Task 085 NOT started; stop for Owner direction.


## Owner Approval Requirements

Approved by Project Owner explicit Task 084 APPROVE in session; Task 083 accepted through the Engineering Task Builder on 2026-09-14 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
