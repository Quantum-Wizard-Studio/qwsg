# Current Engineering Task 093: Community C9 Release Candidate Preparation

## Task Metadata

- Task ID: `093`
- Task slug: `community-c9-release-candidate-preparation`
- Status: `active`
- Date opened: `2026-09-25` UTC
- Human authority: Project Owner explicitly authorizes Task 093 in this session: release candidate preparation, lifecycle, source commit and push; no signing, publication or VPS mutation.
- Owner or lead-developer communication language: Hungarian

## Title

Community C9 Release Candidate Preparation


## Objective

Prepare the deterministic post-1.3.1 Community release candidate and exact Dell1 offline signing handoff for C9. C9 remains MISSING.


## Scope

Inspect version policy and Tasks 082–092; choose smallest correct distinct semantic version; minimally update version/build identity, release notes and candidate metadata; build reproducibly; prepare authenticated 1.3.1 bootstrap/migration authority and deterministic signing input; validate; lifecycle/archive/history; commit and push source preparation.


## Out of Scope

No production signing, publication, tagging, VPS mutation, real install/upgrade/rollback acceptance, Pro gate changes, new product features, POST-1.0 work or Task 094. Historical 1.3.1 is immutable.


## Authority Envelope

- **Task targets and boundaries:** Task 093 release preparation only, with immutable historical 1.3.1 and frozen product scope.
- **Permitted external actions:** Fetch origin and reviewed source commit/push on canonical main.
- **Owner-reserved decisions:** Offline production signing, publication, real-system acceptance and material product/version policy ambiguity.
- **Task-specific STOP conditions:** Material baseline mismatch; ambiguous version authority; unproven security-critical bootstrap authority. Never weaken release verification.


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

Verify main, expected HEAD 5b2d4acf36cf844988851e4b551bc7a9a3a7aa1c equal fetched origin/main, divergence 0 0, clean worktree, valid framework/idle lifecycle with latest complete Task 092.


## Snapshot Requirements

Before mutation create protected /tmp/qwsg-task093-snapshot baseline.tar from tracked HEAD, archive and per-member SHA256, readability and bounded restore instructions. Retain through Owner acceptance and release rollback dependency.


## Risk Assessment

High: signing ambiguous bytes or false bootstrap authority; mitigate exact source/artifact/index binding and fail-closed inspection/tests. Medium: stale provenance or nonreproducibility; freeze source commit and controlled build inputs. No runtime-host mutation.


## Planned Work

Verify baseline and snapshot; install lifecycle; inspect authority and determine version; inspect released 1.3.1 bootstrap; minimally prepare source; validate and commit source boundary if needed; build/rebuild candidate and verify identity/integrity; generate deterministic unsigned authority and signing bytes; freeze inputs; mandatory validation; handoff; archive and commit/push; verify clean synchronized idle; STOP.


## Rollback Plan

Verify snapshot archive and member checksums; extract only into a new empty private review directory. Compare exact Task 093 changed paths before bounded restoration. Never extract over live checkout or broad reset/clean. Preserve evidence and published history. Destructive restoration requires explicit authority; rerun framework/lifecycle/diff checks.


## Deliverables

Distinct candidate version and immutable source commit; reproducible Linux amd64 archive, manifest/checksum; staging unsigned authenticated release/migration metadata; deterministic signing input; exact Dell1 handoff including filename, size, SHA256, public authority, signature output and immutable values.


## Verification

All 18 Owner acceptance cases: baseline/lifecycle; evidence-based version; 1.3.1 immutability; candidate version/build/rebuild/manifest/checksum; consistent metadata; authenticated supported bootstrap with unknown paths refused; deterministic signing input binding security fields; no signing/publication/VPS; unchanged C9; clean lifecycle and synchronized HEAD/origin. Focused release/update tests and mandatory engineering/release validation once after stabilization.


## Documentation Updates

English candidate release notes/changelog, version rationale, bootstrap and offline handoff, independent Task 093 history and concise milestone; archive completed prompt without successor. Explicit release candidate/signed/published/accepted distinctions.


## Completion Criteria

PASS only if applicable Owner acceptance matrix passes; otherwise PARTIAL/BLOCKED with exact security or authority blocker. Community stays 8 PASS / 0 PARTIAL / 1 MISSING; Pro stays 9 PASS / 2 PARTIAL / 3 MISSING. Stop before production signing/publication/real acceptance; no Task 094.


## Owner Approval Requirements

Approved by Project Owner explicitly authorizes Task 093 in this session: release candidate preparation, lifecycle, source commit and push; no signing, publication or VPS mutation. through the Engineering Task Builder on 2026-09-25 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
