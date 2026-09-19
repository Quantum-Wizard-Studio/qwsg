# Current Engineering Task 087: Guardian Handoff & Rollback Recovery Hardening

## Task Metadata

- Task ID: `087`
- Task slug: `guardian-handoff-rollback-recovery-hardening`
- Status: `complete`
- Date opened: `2026-09-19` UTC
- Human authority: Project Owner explicit authorization in this session for task creation, implementation, validation, commit, push and closure.
- Owner or lead-developer communication language: Hungarian

## Title

Guardian Handoff & Rollback Recovery Hardening


## Objective

Make Guardian recovery explicitly owned, preserve pre-handoff service intent, and durably distinguish package, rollback and verified recovery outcomes. Complete this bounded task in one session.


## Scope

Task 085 trigger/transient handoff, Task 084 result composition, shared Guardian recovery and manual rollback evidence; focused tests, required validation and minimal lifecycle/documentation.


## Out of Scope

No release, version bump, production acceptance/mutation, sudo provisioning, privilege redesign, retention, inventory, checkpoints, installer redesign, entitlement backend, native packages or unrelated refactors.


## Authority Envelope

**Task targets and boundaries:** Guardian handoff/recovery and naturally shared manual rollback evidence only.
**Permitted external actions:** Canonical Git fetch, dry-run push, fast-forward push and verification; isolated local tests.
**Owner-reserved decisions:** Out-of-scope architecture/security, production configuration, unrelated services, global sudo policy, destructive recovery and credentials.
**Task-specific STOP conditions:** Cannot safely complete within these boundaries.


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

Verified main HEAD and freshly fetched origin/main both 9cd2cec207dd7f78b8fab58f7a5cf16a9118d67b; divergence 0/0; clean worktree; valid idle lifecycle, latest completed Task 086.


## Snapshot Requirements

Protected local baseline archive before lifecycle mutation; 594 regular members with SHA-256 manifest, readable tar, archive SHA-256 0afe2448ff9f3c79bb4026cbeeeb08ff0747b3ff26b16ef68a82255ef6ec57fd. Retain until Owner acceptance.


## Risk Assessment

High impact: recovery is not authorization. Fail closed on unknown package/service state. Preserve Tasks 082–086, helper and rollback authority. No unconditional worker-exit restart.


## Planned Work

Map ownership before editing; implement single explicit recovery owner; early focused tests; concurrency/security regression, full/race/framework/lifecycle validation; truthful evidence, commit/push and idle closure.


## Rollback Plan

Verify archive SHA-256 and per-member manifest; extract only into new empty private directory; compare exact task changed paths before authorized bounded restoration. Never extract over live worktree or broad reset/clean. Rerun affected tests and lifecycle/Git checks.


## Deliverables

Explicit service intent and verified recovery, distinct durable package/rollback/recovery outcomes, focused running/stopped/pre-mutation/post-mutation/rollback/restart/abnormal/concurrency/Community tests, closed task and synchronized clean main.


## Verification

Formatting; focused scenarios; Task 086 common exclusion/reentry; Tasks 082–085 and helper security; go test ./...; go vet ./...; affected race tests; make engineering-test; framework-v2/diagnostic/lifecycle validators; snapshot verification; targeted staged/privacy review; final HEAD == origin/main, 0/0 clean idle.


## Documentation Updates

Task history, archived prompt, concise existing orchestration contract and chronological milestone only.


## Completion Criteria

PASS only after implementation and all mandatory validation succeed, terminal recovery evidence fails closed, authority unchanged, commit/push and normal idle closure verified. No release.


## Owner Approval Requirements

Approved by Project Owner explicit authorization in this session for task creation, implementation, validation, commit, push and closure. through the Engineering Task Builder on 2026-09-19 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
