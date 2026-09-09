# Current Engineering Task 082: Forward-Compatible Authenticated Update Migration Authority

## Task Metadata

- Task ID: `082`
- Task slug: `forward-compatible-authenticated-update-migration-authority`
- Status: `complete`
- Date opened: `2026-09-09` UTC
- Human authority: Project Owner; Task 082 supplied in conversation with explicit APPROVE
- Owner or lead-developer communication language: Hungarian

## Title

Forward-Compatible Authenticated Update Migration Authority


## Objective

Replace future-target-specific local migration authority with authenticated, schema-versioned release compatibility declarations validated against stable capabilities compiled into the installed client. An old client must safely install an unknown future target using the existing update transaction. Preserve fail-closed security and deterministic rollback. This is the common foundation for Community operator updates and later Pro policy orchestration; no automatic installation is authorized.


## Scope

Inspect and modify update, release discovery/index, compatibility, installation classification, awareness, package verification/provenance, release signing/publication tooling, CLI, tests and documentation only as needed. Preserve subsystem boundaries and historical supported migration semantics. Explicit archive updates must enforce the same authority, identity, provenance, integrity and capability gates.


## Out of Scope

Task 083; Pro licensing, entitlements, scheduler, maintenance windows, fleet/remote management, unattended installation, arbitrary remote scripts/commands/binaries as migration authority, unrelated Guardian/UI/notification features, QWSG 1.3.2 publication. Historical v1.3.0/v1.3.1 assets, tags, signatures and index evidence are immutable.


## Authority Envelope

- **Task targets and boundaries:** Task 082 objective and scoped components above; Standard Execution Authority includes implementation, tests, documentation and canonical closure.
- **Permitted external actions:** Canonical repository synchronization and task-scoped Git integration under framework policy; no production deployment/publication.
- **Owner-reserved decisions:** Material authority expansion; releases, production infrastructure and excluded work.
- **Task-specific STOP conditions:** Unsafe or unnecessarily complex architecture requires an explicit narrower/stronger proposal preserving authenticated authority, local executable authority, determinism, fail-closed behavior, rollback and auditability.


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

Before mutation verify branch, HEAD, origin/main, divergence, clean worktree, lifecycle and current task; record baseline and create standard snapshot. Expected Task 081 accepted/complete, release source 45009fe6169bff00842a4c4e9561bf339a5db81e, closure fa570723c4d6b0a1f6758a9c2ea9ddf610716765, v1.3.1 tag. Verify rather than assume.


## Snapshot Requirements

Create and verify a protected pre-task snapshot according to the engineering backup policy before target or lifecycle mutation. Record scope, checksums, readable archive, retention through Owner acceptance, and guarded bounded restore.


## Risk Assessment

High security and compatibility risk: absent, malformed, unauthenticated, unknown, inconsistent, ambiguous or out-of-authority declarations must refuse. Cover schema, mechanism, source, target, platform, provenance, artifact identity/digest, downgrade and conflicting selection. Remote data must never expand compiled executable authority. Preserve configuration, service, state and rollback safety.


## Planned Work

Inspect existing installation/update/migration/rollback/configuration/service/package/release authority. Choose the narrowest coherent stable capability model; extend deterministic canonical signed metadata and offline signing support. Integrate discovery, classification, CLI and explicit archives. Add security, regression, transaction and mandatory old-binary forward-release acceptance. Verify, document and close Task 082 without Task 083.


## Rollback Plan

Use verified pre-task snapshot and exact reviewed changed paths; restore into isolated directory first, never broadly reset or overwrite live worktree. Preserve historical material and unrelated data. Forward migration must use existing deterministic backup/install/post-validation/rollback transaction.


## Deliverables

Architecture and implementation; canonical schema/signing contract; discovery/classification/update/archive integration; deterministic security/unit/integration/release tests; mandatory black-box acceptance; English technical and appropriate Hungarian operator documentation; implementation record, verification evidence, rollback notes and final repository integrity evidence.


## Verification

Use pinned Go toolchain. Run gofmt, vet, go test, race, framework/lifecycle, release/update/discovery/authority and reproducibility checks. Test unknown future target eligible using locally known capability; unknown capability, unauthenticated metadata, source mismatch, artifact/provenance mismatch, ambiguous/conflicting authority refused; historical migration regression; explicit archive cannot bypass authority/integrity; rollback; discovery never installs. Mandatory black-box: build old client before future target exists, no future-specific migration table; authenticated future declaration alone authorizes known local capability through qwsg update and all normal safety gates. Target binary executing migration or injecting a future-specific path is insufficient. Never weaken/skip tests for PASS.


## Documentation Updates

Explain why future hard-coded tables fail; signed authority versus local executable capability, trust boundaries, refusal behavior, operator-controlled Community, unknown-capability operator behavior and safe common foundation for future Pro policy. Update canonical technical and appropriate Hungarian documentation, task implementation history and verification/rollback records.


## Completion Criteria

Complete only when old client needs no future target hard-coding, only authenticated compatibility can authorize compiled capabilities, unknown capabilities refuse, integrity/provenance/rollback remain enforced, Community remains operator-controlled, mandatory black-box and relevant/race tests pass, documentation is updated, lifecycle and worktree reach canonical closure, and no Task 083 work occurs.


## Owner Approval Requirements

Approved by Project Owner; Task 082 supplied in conversation with explicit APPROVE through the Engineering Task Builder on 2026-09-09 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
