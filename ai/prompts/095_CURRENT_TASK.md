# Current Engineering Task 095: Community 1.4.0 Real-System Clean Install Acceptance

## Task Metadata

- Task ID: `095`
- Task slug: `community-1-4-0-real-system-clean-install-acceptance`
- Status: `active`
- Date opened: `2026-09-30` UTC
- Human authority: Project Owner explicitly authorizes Task 095 in the current session under the supplied eight-phase specification and thirty-item acceptance matrix.
- Owner or lead-developer communication language: Hungarian

## Title

Community 1.4.0 Real-System Clean Install Acceptance


## Objective

Prove official production-signed Community 1.4.0 public acquisition, authentication, clean installation and real supported OVH Ubuntu operation in one bounded session.


## Scope

Task 095 lifecycle/evidence; canonical release read verification; designated OVH playground preflight and approved bounded clean installation/operation only. All requirements of the Owner-supplied Task 095 specification apply.


## Out of Scope

No historical upgrade, migration bootstrap, rollback acceptance, production mutation, publication, rebuilding/resigning, Task 096, Pro or unrelated refactoring. C9 remains MISSING.


## Authority Envelope

- **Task targets and boundaries:** Clean-install acceptance only on the designated OVH playground; immutable 1.4.0 and task evidence.
- **Permitted external actions:** Canonical public release reads, existing authorized SSH read-only inventory, task-scoped Git integration. Target reset/install only after exact plan and explicit Owner approval at the external destructive/credential boundary.
- **Owner-reserved decisions:** Target identification when unavailable, exact destructive reset/install approval, credential-bound actions, replacement releases and scope expansion.
- **Task-specific STOP conditions:** Release authority failure before target access; ambiguous host; unsafe baseline; credential boundary; product defect requiring release replacement.


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

Verified main HEAD/origin/main c4dfe92c010195a7fb43162498815ec4748d64b4, fetched origin, divergence 0/0, clean worktree, idle after complete 094.


## Snapshot Requirements

Protected /tmp/qwsg-task095 baseline.tar from git archive HEAD; SHA256 recorded and tar readability verified before lifecycle mutation. Retain through Owner acceptance and rollback dependency expiry. Prior target evidence/snapshot required before approved reset.


## Risk Assessment

High external destructive/credential risk: inventory first, exact bounded plan, Owner checkpoint; no unrelated mutation. High trust risk: exact public identity and production signature before installation.


## Planned Work

Execute Owner phases 1-8 and acceptance items 1-30 sequentially. Inventory only after authority PASS. Use documented install and stability contracts. Respect 50/25/10 percent budget convergence; deterministic partial evidence if unable to finish.


## Rollback Plan

Local archive is restored only into a separate directory after digest/readability verification; compare and restore exact task-owned paths with review, never extract over live worktree. Target recovery plan must be derived from inventory and approved before reset. No historical rollback acceptance.


## Deliverables

Sanitized acceptance/history record, exact public release verification, target preflight/reset plan, authenticated official install and bounded real-system/stability evidence if authorized, truthful PASS/PARTIAL/BLOCKED and Git closure.


## Verification

All thirty Owner acceptance items mandatory for PASS. Version 1.4.0 source 996fb90d0f53d4ae6088cd884a654d9d0a32fff9; artifact 3680796 bytes SHA256 26cbf8d76fa27c652e11ee49f4570e05b3be6b969309e3cfd486977a27a6224d; signed index 1153 bytes SHA256 fbb39b52e240ff3691090a48a1e7915e516da6fb2f6a1faa9ed5fc069b633c41; production pinned signature; unchanged historical 1.3.1. Existing collector/evidence/Guardian/scheduler/awareness/current/resource/stability/notification contracts only.


## Documentation Updates

Task prompt/history, sanitized clean-install acceptance evidence and chronological milestone; C9 remains MISSING and Community 8 PASS/0 PARTIAL/1 MISSING.


## Completion Criteria

PASS/COMPLETE/CLOSED only with every required check, push, HEAD==origin/main, divergence 0/0, clean worktree and idle lifecycle. Otherwise truthful PARTIAL/BLOCKED with deterministic continuation; do not falsely archive incomplete task as complete. Stop without successor work.


## Owner Approval Requirements

Approved by Project Owner explicitly authorizes Task 095 in the current session under the supplied eight-phase specification and thirty-item acceptance matrix. through the Engineering Task Builder on 2026-09-30 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
