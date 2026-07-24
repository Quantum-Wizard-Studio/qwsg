# Current Engineering Task 020: Canonical Drift Engine

## Task Metadata

- Task ID: `020`
- Task slug: `canonical-drift-engine`
- Status: `complete`
- Date opened: `2026-07-24` UTC
- Human authority: Project Owner through the Task 020 implementation instruction
- Owner or lead-developer communication language: Hungarian

## Title

Canonical Drift Engine


## Objective

Implement the permanent Canonical Drift Engine as the deterministic semantic interpretation layer above the Snapshot Comparison Engine, classifying every canonical Change Record into one versioned canonical Drift Record without health, risk, policy, or remediation judgement.


## Scope

Design and implement the Canonical Drift Architecture, public versioned Drift Engine and Drift Record contracts, extensible canonical taxonomy, deterministic classification pipeline, metadata and versioning conventions, documentation, tests, and required lifecycle records. Integrate only where necessary to establish the Compare to Drift pipeline while preserving existing comparison and CLI behavior.


## Out of Scope

Do not implement Health evaluation, risk scoring, policy or compliance engines, alerts, email, scheduler, daemon, monitoring, AI integration, automatic remediation, system modifications, or unrelated product features. Do not reinterpret inventory or snapshots directly; consume only canonical Change Records and produce only canonical Drift Records. Do not break existing public contracts or CLI behavior.


## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

Verify canonical idle state after completed Task 019; repository root, framework version, main branch, canonical origin, HEAD and origin relationship, complete tracked and untracked status, lifecycle consistency, relevant Comparison Engine contracts, permissions, ownership, ACLs, Go toolchain, and existing validations. Treat the documented local Task 019 commit ahead of origin and pre-existing untracked owner artifacts as expected and preserve them. Stop on any other material difference.


## Snapshot Requirements

Before any Task 020 target modification, create a timestamped private snapshot under /tmp containing a complete Git bundle, the pre-change lifecycle state, and a bounded archive of every existing implementation, test, documentation, and governance target. Record absent new paths explicitly, SHA-256 checksums, file list, permissions, ownership, ACL evidence, and verify bundle and archive readability. Retain both lifecycle and implementation snapshots through owner acceptance.


## Risk Assessment

Primary risks are semantic misclassification, unstable IDs or ordering, contract ambiguity, accidental coupling to future Health or Policy concerns, backward-incompatible comparison changes, metadata privacy leakage, nondeterministic map or time behavior, and incomplete rollback coverage. Mitigate with explicit versioned enums and schemas, pure deterministic classification, stable canonical serialization inputs, conservative taxonomy precedence, validation, golden/repeatability tests, compatibility tests, bounded metadata, documentation, and exact snapshots.


## Planned Work

Inspect canonical Inventory and Comparison contracts and their tests. Define taxonomy boundaries, classification precedence, scope semantics, deterministic confidence representation, identifiers, metadata constraints, version negotiation and validation. Implement the smallest isolated internal drift package consuming comparison.ChangeRecord values only and producing validated Drift Records in stable order. Add integration only as needed to demonstrate Compare to Drift composition without changing established compare output. Add exhaustive unit, determinism, validation, compatibility, privacy, and pipeline tests. Update permanent architecture, system map, roadmap/history, and canonical Drift documentation. Verify and complete lifecycle records, then archive Task 020 to canonical idle state after all gates pass.


## Rollback Plan

Rollback is bounded to Task 020 paths. Verify snapshot hashes and archive listings first; compare current targets and obtain confirmation before any destructive restoration. Restore only pre-existing Task 020 targets from the implementation archive, remove only new paths whose pre-task absence is recorded, and restore lifecycle files from the lifecycle archive if required. Re-run framework, lifecycle, Go, formatting, vet, race, documentation, permission, ACL, Git, and repository consistency checks. Never use broad reset, checkout, clean, wildcard deletion, or extraction over the live worktree.


## Deliverables

A public versioned Drift Record contract; canonical taxonomy including Identity, Software, Hardware, Platform, Filesystem, Storage, Network, Service, Configuration, Security, Capability, and Environment Drift; deterministic offline AI-independent Drift Engine; stable validation and classification behavior; Compare to Drift pipeline tests; permanent canonical architecture and lifecycle documentation; compatibility strategy; completed Task 020 history and idle lifecycle closure.


## Verification

Run all configured Engineering Framework validations and lifecycle validators; Builder input and installation checks; all Go tests; gofmt check; go vet; race tests; deterministic repeatability and stable JSON tests; drift taxonomy coverage and invalid-contract rejection tests; comparison backward-compatibility tests; privacy and offline-boundary review; documentation link and terminology consistency checks; git diff checks; snapshot checksum and restore-readiness checks; ownership, permission, ACL, scope, and complete Git-state review. Every command and result must be truthfully recorded.


## Documentation Updates

Create docs/architecture/CANONICAL_DRIFT_ENGINE.md as the permanent reference covering Drift Architecture, taxonomy, lifecycle, Compare to Drift pipeline, future Health, Rule and Policy integration boundaries, and compatibility strategy. Update docs/architecture/SNAPSHOT_COMPARISON_ENGINE.md, ai/core/04_ARCHITECTURE.md, ai/core/05_SYSTEM_MAP.md, ai/core/07_ENGINEERING_HISTORY.md, ai/core/13_ROADMAP.md, README.md, the active prompt, and the independent Task 020 history as required by actual implementation.


## Completion Criteria

Complete only when every canonical Change Record accepted by the engine yields exactly one validated canonical Drift Record; all required taxonomy categories and public versioned contracts exist; classification, IDs, ordering, confidence, metadata, and serialization are deterministic and reproducible; privacy, offline, AI-independent, compatibility, and separation boundaries are demonstrated; all mandated validations pass; documentation and lifecycle evidence are complete; rollback remains usable; Task 020 is committed with targeted staging and the repository returns to valid canonical idle state. Otherwise report blocked or complete with explicit disclosed limitations without claiming success.


## Owner Approval Requirements

Approved by Project Owner through the Task 020 implementation instruction through the Engineering Task Builder on 2026-07-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
