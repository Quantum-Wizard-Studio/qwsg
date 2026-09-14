# Current Engineering Task 083: Pro Update Policy & Capability Foundation

## Task Metadata

- Task ID: `083`
- Task slug: `pro-update-policy-capability-foundation`
- Status: `complete`
- Date opened: `2026-09-14` UTC
- Human authority: Project Owner explicit Task 083 APPROVE in session; Task 082 accepted
- Owner or lead-developer communication language: Hungarian

## Title

Pro Update Policy & Capability Foundation


## Objective

Establish deterministic fail-closed product capabilities update.manual and update.automatic and canonical manual/automatic update policy. Community remains complete and operator-controlled. Pro test authority may validate automatic policy but MUST NOT execute updates.


## Scope

Canonical capability and update-policy packages; existing configuration/status integration; focused tests including Task 082 security regression; concise architecture documentation, history, snapshot metadata and lifecycle closure. Reuse the single authenticated update engine.


## Out of Scope

Task 084; automatic/unattended execution, Guardian installation, scheduler, maintenance windows, reboot, new rollback orchestration, licensing server, payments/accounts, network entitlements, fleet/remote control, dashboard, release publication or v1.3.2; unrelated refactoring. Published v1.3.1 is immutable.


## Authority Envelope

**Task targets and boundaries:** Only Task 083 foundation and its configuration, visibility, tests and documentation.
**Permitted external actions:** Fetch, reviewed commit, dry-run and fast-forward push to canonical origin/main, post-push fetch/verification. No deployment or publication.
**Owner-reserved decisions:** Commercial entitlement mechanism, Task 084 and releases.
**Task-specific STOP conditions:** None beyond standard; stop expanding scope if it cannot fit one session.


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

Verify main at Task 082 closure 38ee760c267aef826c8fde97843853ad206d72ca; fetch, HEAD == origin/main, divergence 0/0, clean worktree and idle lifecycle; inspect architecture and configuration before mutation. Actual baseline verified unchanged.


## Snapshot Requirements

Before mutation preserve protected /tmp/qwsg-task083-snapshot/baseline.tar of tracked baseline with archive/member SHA-256 manifest and readability check. Retain until Owner acceptance. No payload in Git.


## Risk Assessment

Medium: accidental capability grant or security bypass; mitigate immutable validated capability result, explicit authority seam, pure policy validation and negative tests. Low configuration risk: preserve existing schema, manual default and legacy notify semantics, existing private file handling. No runtime deployment.


## Planned Work

Inspect baseline/architecture; minimal capability/policy design; implement; focused security/config tests; regression and required gates; concise docs; archive Task 083, commit, push, fetch and final clean idle verification. Fit one session; defer orchestration.


## Rollback Plan

Verify archive and member SHA-256, extract only into new empty private directory, compare task changed-path allowlist against baseline. Individually review bounded restoration, never extract over live tree or broad reset/clean. Re-run relevant tests and lifecycle validators. Preserve history and obtain authority before destructive restoration.


## Deliverables

Capability model; manual/automatic authority and policy; compatible Community default; deterministic Pro test fixture; operator visibility; security/regression tests; concise docs; history evidence and canonical closure.


## Verification

gofmt, go vet, focused capability/policy/configuration tests, relevant race tests, Task 082 authenticated authority regression and frozen-client gate; full local Go and configured framework/lifecycle gates as required. Verify snapshot, scope/privacy/modes, immutable historical release, reviewed staging, dry-run push, HEAD == origin/main, 0/0, clean worktree and idle lifecycle.


## Documentation Updates

Concise capability/update policy contract covering Community vs Pro, default manual, fail closed, independent Task 082 security gates, future production entitlement source and deferred Task 084. Task history and completed prompt archival.


## Completion Criteria

PASS only when all Task 083 deliverables and required gates pass; automatic without capability refused, Pro fixture automatic policy accepted with zero execution, malformed/unknown/conflicting/schema authority fails closed, Task 082 cannot be bypassed, old config compatible; committed/pushed clean idle main equal origin/main. Report implementation commit and tests. Task 084 must NOT be started or prepared. Stop after closure.


## Owner Approval Requirements

Approved by Project Owner explicit Task 083 APPROVE in session; Task 082 accepted through the Engineering Task Builder on 2026-09-14 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
