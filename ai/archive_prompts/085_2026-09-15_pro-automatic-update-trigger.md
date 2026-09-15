# Current Engineering Task 085: Pro Automatic Update Trigger and Guardian Integration

## Task Metadata

- Task ID: `085`
- Task slug: `pro-automatic-update-trigger`
- Status: `complete`
- Date opened: `2026-09-15` UTC
- Human authority: Project Owner explicit Task 085 APPROVE; Task 084 accepted.
- Owner or lead-developer communication language: Hungarian

## Title

Pro Automatic Update Trigger and Guardian Integration


## Objective

Integrate a deterministic recurring automatic update trigger into the existing Guardian lifecycle, reusing Tasks 082, 083 and 084. Deliver in one session.


## Scope

Guardian release-check scheduler, capability/policy gating, authenticated candidate decision, safe handoff to Task 084, evidence, focused security and integration tests, concise documentation and lifecycle closure; commit and push main.


## Out of Scope

No advanced scheduling, maintenance/blackout/timezone rules, reboot, fleet/remote management, licensing backend, release publication, v1.3.2 or Task 086. No new daemon or independent updater service unless existing architectural constraints make a bounded handoff unavoidable. Historical v1.3.1 immutable.


## Authority Envelope

**Task targets and boundaries:** Task 085 scope above; one common update engine.
**Permitted external actions:** Canonical Git fetch, dry-run push, fast-forward push and verification only; isolated fixtures, no live installation mutation.
**Owner-reserved decisions:** Releases, deployment, scope expansion and Task 086.
**Task-specific STOP conditions:** None beyond standard security/rollback/authority boundaries.


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

Verify main at ab66c7323999e9748ca53a952f0e80b3f2daef7e, HEAD equal fetched origin/main, 0/0 and clean. Lifecycle idle after archived Task 084; Task 085 absent; record historical v1.3.1 object. Inspect Guardian and scheduler before design.


## Snapshot Requirements

Before lifecycle or source mutation, create protected /tmp baseline archive, SHA-256 and per-member manifest; verify readability and hashes. Retain until Owner acceptance.


## Risk Assessment

High-impact mutation and service lifecycle risk: preserve inactive Guardian requirement, signed authority, canonical capability/policy gates, transaction lock, rollback evidence. Never update production during acceptance.


## Planned Work

Verify baseline; inspect existing Guardian scheduler and service primitives; select narrow deterministic trigger/handoff; implement gating, authenticated candidate handling, Task 084 reuse and evidence; focused tests and regressions; minimal documentation; close/archive; commit/push/fetch and final verification.


## Rollback Plan

Validate archive and member hashes, extract into new empty private directory, compare exact changed paths, review bounded restoration with destructive approval where required. Never extract over live checkout, reset or clean broadly. Rerun affected tests and lifecycle checks.


## Deliverables

Guardian recurring trigger; Task 083 gating; Task 082 candidate authority; safe Task 084 handoff; duplicate/re-entry protection; result persistence; deterministic tests/black-box evidence; concise documentation/history; committed pushed clean idle repository.


## Verification

gofmt, go vet, focused and affected race tests, Guardian scheduler and Tasks 082/083/084 regressions, handoff tests and canonical framework/lifecycle gates. Cover Community default/manual, Pro manual, authorized current no-op, supported future candidate exactly once, authentication and migration refusal, repeats, active transaction, handoff failure, rollback success/failure and Guardian stability. Isolated black-box Pro running Guardian to Task 084 terminal result and Community same candidate no apply. No real release.


## Documentation Updates

Concise architecture, Community/Pro boundary, safe handoff, repeated execution, failures/rollback, preserved security and deferred advanced scheduling; Task 085 history, milestone and completed archived prompt.


## Completion Criteria

PASS only after all deliverables and tests pass, Task 082/083/084 boundaries preserved, Community manual unchanged, explicit rollback failure, lifecycle closed and archived, commit/push/fetch done, clean main equals origin/main at 0/0; no release or Task 086. Final report includes all requested architecture, behavior, evidence and Git fields.


## Owner Approval Requirements

Approved by Project Owner explicit Task 085 APPROVE; Task 084 accepted. through the Engineering Task Builder on 2026-09-15 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
