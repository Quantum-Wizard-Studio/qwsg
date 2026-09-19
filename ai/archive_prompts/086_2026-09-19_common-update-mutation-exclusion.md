# Current Engineering Task 086: Common Update Mutation Exclusion

## Task Metadata

- Task ID: `086`
- Task slug: `common-update-mutation-exclusion`
- Status: `complete`
- Date opened: `2026-09-19` UTC
- Human authority: Project Owner explicit authorization in the Task 086 request, including creation, execution, commit, push and closure.
- Owner or lead-developer communication language: Hungarian

## Title

Common Update Mutation Exclusion


## Objective

For one installed QWSG state, mutually conflicting manual authenticated update, automatic authenticated update and manual rollback transactions share one nonblocking process-safe exclusion boundary. Complete within one session.


## Scope

Inspect mutation paths and existing locks; reuse existing locking behavior through a small common mechanism; integrate manual update, automatic orchestration and manual rollback; focused deterministic contention/security/race tests; required validation, minimal documentation, snapshot, lifecycle closure, commit and push.


## Out of Scope

No release or version bump, production deployment/acceptance, Guardian handoff/restart redesign, rollback receipt redesign, sudo provisioning, privilege readiness, backup lifecycle, inventory, checkpoint, installer UX, native packages, entitlement backend or unrelated refactors.


## Authority Envelope

**Task targets and boundaries:** Common update mutation exclusion and focused tests/documentation only.
**Permitted external actions:** Canonical Git fetch, dry-run push, fast-forward push and verification. Isolated local fixtures only.
**Owner-reserved decisions:** Architectural/security changes outside scope, production configuration, destructive recovery, private credentials and publication.
**Task-specific STOP conditions:** Solution cannot safely remain in the stated scope.


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

Verified main HEAD and freshly fetched origin/main both 8dea8fcc649e89c995c3120052af6b564adb6ff2; divergence 0/0; clean worktree; valid idle lifecycle with latest completed archived Task 085. Framework configuration validates canonical identity.


## Snapshot Requirements

Protected local baseline archive created before lifecycle mutation; 589 regular tracked members, archive SHA-256 acf29560500dbec72aa2f068fcebe8fd2851515efce77d4786f09417e7bf718f; archive readable and every member digest verified. Retain until Owner acceptance.


## Risk Assessment

High-impact concurrency change: lock ownership is coordination only, never authorization. Preserve authenticated release/migration, capability/policy, helper and rollback gates. Acquire before mutable state reads and service/package changes; hold through recovery and record finalization. Nonblocking contention; no nested acquisition.


## Planned Work

Map current entry points, helper and first mutation; extract existing secure nonblocking flock into common ownership API; keep automatic process reentry protection; wire manual entry points; focused tests early, security regressions, full and race validation; concise evidence and idle closure.


## Rollback Plan

Verify baseline archive and per-member digests; extract into new empty private directory only; compare and restore exact reviewed task paths with authority for destructive recovery. Never extract over live checkout, reset or clean broadly. Recheck tests, Git and lifecycle after bounded restoration.


## Deliverables

Common exclusion, explicit ownership contract, deterministic distinguishable contention, focused concurrency/read-only/failure-release tests, preserved security gates, required evidence, archived completed task and clean synchronized main.


## Verification

gofmt; focused update/manual/automatic contention tests in both directions, reentry, separate processes, failure release and read-only checks; security regression suite for Tasks 082-085/helper/rollback; go test ./...; go vet ./...; affected concurrency race tests; make engineering-test and configured framework/lifecycle checks; snapshot and staged diff/privacy review; final HEAD equal origin/main, 0/0 clean idle.


## Documentation Updates

Task history and archived prompt, concise mutation ownership contract in update orchestration documentation, chronological engineering milestone.


## Completion Criteria

PASS only when all scoped paths share exclusion, required tests and security gates pass, snapshot remains verified, task closed/archived, implementation committed/pushed, main equals origin/main with clean worktree and 0/0 divergence. No release. Report implementation/closure hashes and any relevant deferred findings.


## Owner Approval Requirements

Approved by Project Owner explicit authorization in the Task 086 request, including creation, execution, commit, push and closure. through the Engineering Task Builder on 2026-09-19 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
