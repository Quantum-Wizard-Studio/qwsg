# Current Engineering Task 091: Safe Explicit Update, Rollback & Recovery

## Task Metadata

- Task ID: `091`
- Task slug: `safe-explicit-update-rollback-recovery`
- Status: `complete`
- Date opened: `2026-09-22` UTC
- Human authority: Project Owner explicitly authorizes Task 091 C7 implementation, tests, documentation, lifecycle completion, commit and push in this session.
- Owner or lead-developer communication language: Hungarian

## Title

Safe Explicit Update, Rollback & Recovery


## Objective

Close only frozen Community C7 if the complete explicit authenticated update, rollback and Guardian recovery contract is implemented and proven. Failed updates remain failures after recovery.


## Scope

Manual update and rollback CLI, authenticated helper integration, transaction backup/journal durability, shared recovery primitives only where required, focused fault injection, directly affected documentation and C7 gate.


## Out of Scope

No release construction, VERSION change, tags, signing, publication, production/VPS mutation, C8/C9 closure, P1/P3/P4/P5 change, unattended Community updates, automation expansion or Task 092.


## Authority Envelope

1. **Task targets and boundaries:** smallest deterministic C7 fixes in manual update/rollback, shared transaction durability and Guardian recovery; focused tests and minimal documentation.
2. **Permitted external actions:** canonical origin fetch, dry-run push and task-scoped fast-forward push. No production/VPS mutation.
3. **Owner-reserved decisions:** release construction/publication, new product scope and other gate changes.
4. **Task-specific STOP conditions:** materially different baseline from a27935e44fcbe60012cf6869490eefe3d0f71a89 or genuine architecture/security boundary.


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

Verify main, clean worktree, HEAD == origin/main == a27935e44fcbe60012cf6869490eefe3d0f71a89, divergence 0/0 after fetch, valid framework/lifecycle idle after completed 090.


## Snapshot Requirements

Before lifecycle/target mutation, capture tracked baseline in protected local /tmp archive, deterministic member SHA256 manifest, verify archive readability and every member. Retain until Owner acceptance; no payload in Git.


## Risk Assessment

High: false update success, rollback loss, unsafe Guardian start and interruption ambiguity. Mitigate with durable pre-mutation journal, package validation before recovery, separate outcomes, common lock and focused fault injection. Preserve Tasks 082/086/087/090.


## Planned Work

Map current manual transaction first; add boundary tests early; minimally repair remaining gaps; focused iteration; freeze source; run mandatory broad validation once; update only C7 if proven; archive idle, commit, push, verify.


## Rollback Plan

Verify protected snapshot and member hashes; extract to NEW private directory only. Review exact modified targets and intervening drift before bounded restoration. No broad reset/clean, preserve evidence/new files, rerun relevant tests and lifecycle.


## Deliverables

Honest explicit update/rollback outcomes, validated Guardian ordering and durable actionable evidence; twelve-case C7 matrix and final report.


## Verification

Cases 1-12 individually: normal success; failed mutation recovered; post-validation failure recovered; rollback execution failure; rollback validation failure; Guardian failure after rollback; inactive intent; common lock conflict; invalid authority; invalid rollback source; current no-op; distinct durable evidence. Stable source: gofmt, vet, full Go tests, relevant race, update/auth/migration/Guardian regressions, engineering/build contract, framework and lifecycle.


## Documentation Updates

Update matching history throughout, minimal operator/architecture recovery docs and chronological index. Change C7 PARTIAL to PASS only with all acceptance evidence; totals then Community 7/1/1 and inherited Pro 8/3/3. Other gates unchanged.


## Completion Criteria

All C7 cases and mandatory checks pass; truthful report. Archive only when complete, idle, reviewed commit pushed, clean worktree, HEAD == origin/main and divergence 0/0. If genuine blocker keep C7 PARTIAL and report exact missing evidence without manufacturing completion.


## Owner Approval Requirements

Approved by Project Owner explicitly authorizes Task 091 C7 implementation, tests, documentation, lifecycle completion, commit and push in this session. through the Engineering Task Builder on 2026-09-22 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
