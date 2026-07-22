# Current Engineering Task 008: Core Alpha Architecture and Slice 1 Definition

## Task Metadata

- Task ID: `008`
- Task slug: `core-alpha-architecture`
- Status: `complete with disclosed architecture gates and permission limitation`
- Date opened: `2026-07-20` UTC
- Human authority: `Project Owner`
- Owner or lead-developer communication language: `Hungarian`
- Task type: `Documentation and architecture design`
- Requires snapshot: `Yes`
- Requires rollback validation: `Yes`
- Implementation changes allowed: `No`
- Destructive actions allowed: `No`

## Title

QWSG Core Alpha Architecture and Core Alpha Slice 1 Definition

## Owner Decisions and Ratification

The Project Owner authorizes this architecture task and ratifies the current Product Definition, Product System Blueprint, Functional Specification, and Task 007 audit outputs as the design baseline for this task, subject to the following limits:

1. Statements already marked as proposals, alternatives, unresolved decisions, assumptions, or release gates remain unresolved unless this task records an explicit owner-approved decision already present in the repository.
2. A downstream document must not silently convert a proposal into a mandatory product requirement.
3. Where documents conflict, the authority order defined by the Constitution and documentation policy must be followed.
4. The first implementation milestone shall be named `Core Alpha Slice 1`.
5. The intended Slice 1 direction is `Read-only Server Discovery and System Inventory`.
6. This task may recommend decisions, but it must not invent business policy, security trust decisions, retention periods, supported platform commitments, commercial boundaries, or cryptographic trust rules on behalf of the owner.
7. Any unresolved decision that blocks safe implementation must be recorded as a clearly named architecture gate.

## Objective

Create the authoritative, implementation-ready Core Alpha architecture for QWSG and define a narrow first vertical product slice called `Core Alpha Slice 1: Read-only Server Discovery and System Inventory`.

The result must establish enough verified architecture, security, permission, data, component, interface, error-handling, testing, and delivery boundaries to permit a later implementation task to begin safely.

This task must transform the existing product-level requirements into a coherent technical design without writing production source code.

## Scope

This task is authorized to:

- inspect the entire repository and all current engineering records;
- verify the completed state of Tasks 006 and 007;
- review all audit findings, release gates, traceability records, and readiness conclusions;
- distinguish ratified requirements from proposals, assumptions, recommendations, and unresolved owner decisions;
- design the Core Alpha system architecture;
- define component boundaries and responsibilities;
- define the Agent, Console, shared contracts, storage, transport, and trust boundaries at architecture level;
- define the minimum security and permission model required before implementation;
- define the canonical internal data model for discovered system inventory;
- define collection provenance, timestamps, freshness, partial-result, and unknown-value semantics;
- define read-only discovery constraints and prohibited mutation behavior;
- define command execution boundaries and privilege behavior;
- define supported and unsupported execution contexts for Slice 1 without making unsupported platform promises;
- define error categories, degraded operation, cancellation, timeout, and retry principles;
- define logging, auditability, privacy, redaction, and secret-handling requirements;
- define interfaces between discovery collectors, normalization, validation, storage or transient state, and presentation;
- define a test strategy and acceptance mapping for Slice 1;
- define implementation sequencing and dependency order;
- create architecture decision records where justified;
- create a decision and gate register for unresolved owner choices;
- update only documentation and lifecycle records required by this task.

Authorized documentation locations include:

- `docs/architecture/`
- `docs/development/`
- `docs/security/`
- `ai/projects/QWSG/`
- `ai/history/`
- `ai/core/05_SYSTEM_MAP.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/13_ROADMAP.md`
- `ai/projects/QWSG.md`
- `README.md`
- `CHANGELOG.md`
- the active Task 008 prompt
- the Task 008 snapshot directory

New architecture documents may be created when they have a clear, non-duplicative authority role.

## Out of Scope

Do not:

- implement Agent, Console, installer, service, daemon, API, database, UI, collectors, or product tests;
- add production code in Python, Go, Rust, PHP, JavaScript, TypeScript, Bash, or any other language;
- choose a final implementation language merely because it is installed on the current server;
- install packages, compilers, runtimes, databases, services, or dependencies;
- change operating-system, web-server, PHP, database, firewall, SSH, HestiaCP, Webmin, DNS, TLS, or production settings;
- execute privileged discovery against the host beyond safe repository and environment inspection needed for documentation;
- perform destructive rollback;
- modify secrets, credentials, ownership, ACLs, or permissions without explicit separate approval;
- commit or push;
- resolve commercial, legal, pricing, licensing, retention, telemetry, e-mail transport, update-signing, or support-policy gates without explicit owner evidence;
- design the whole final product in exhaustive detail;
- allow Core Alpha Slice 1 to grow into general monitoring, remediation, configuration management, package updates, alerting, remote control, or automated repair;
- silently rewrite previously approved product intent.

Later implementation, verification, packaging, and hardening tasks are explicitly deferred.

## Required Reading

Read and reconcile at minimum:

- `AGENTS.md`
- `README.md`
- `CHANGELOG.md`
- `VERSION`
- `ai/README.md`
- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/02_PROJECT_STRUCTURE.md`
- `ai/core/03_AGENTS.md`
- `ai/core/04_ARCHITECTURE.md`
- `ai/core/05_SYSTEM_MAP.md`
- `ai/core/06_ENGINEERING_STANDARDS.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/09_DELIVERY_POLICY.md`
- `ai/core/10_DOCUMENTATION_POLICY.md`
- `ai/core/11_SECURITY_POLICY.md`
- `ai/core/12_RELEASE_POLICY.md`
- `ai/core/13_ROADMAP.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/projects/QWSG.md`
- `ai/projects/QWSG/QWSG_MASTER_PLAN.md`
- `docs/PRODUCT_DEFINITION.md`
- `docs/PRODUCT_SYSTEM_BLUEPRINT.md`
- `docs/FUNCTIONAL_SPECIFICATION.md`
- `docs/development/REQUIREMENTS_TRACEABILITY_MATRIX.md`
- `docs/development/CORE_ALPHA_READINESS.md`
- `ai/audits/2026-07-20_QWSG_REPOSITORY_DEEP_AUDIT.md`
- `ai/audits/2026-07-20_QUANTUM_CREATOR_CONFORMANCE.md`
- `ai/history/006_2026-07-19_functional-specification.md`
- `ai/history/007_2026-07-20_repository-deep-audit.md`
- the complete active Task 008 prompt

If a required record has a different verified path, locate it and record the variance. Do not silently substitute unrelated material.

## Starting State Verification

Before modifying any file:

1. Confirm the repository root.
2. Run `bin/job --check`, `bin/job --path`, and `bin/job --history`.
3. Confirm that exactly one active prompt exists and that its metadata resolves to Task 008.
4. Confirm that Tasks 006 and 007 are complete and that their expected deliverables exist.
5. Record:
   - current UTC time;
   - branch and commit;
   - complete Git status;
   - existing modified and untracked files;
   - relevant file ownership, permissions, and ACLs;
   - available toolchain versions;
   - current repository structure;
   - existing architecture and decision records.
6. Detect placeholders, unresolved lifecycle states, duplicate authoritative documents, broken internal links, inconsistent task references, and unexpected permission limitations.
7. If the active prompt is stale in the current agent session, stop and report that a new session or explicit reload is required.
8. If a contradiction prevents safe architecture work, create a bounded preflight failure report and stop without modifying product documentation.
9. Do not require a clean working tree, but preserve and clearly distinguish pre-existing changes from Task 008 changes.

## Snapshot Requirements

Before architecture edits, create a timestamped snapshot under:

`ai/backups/<UTC_TIMESTAMP>_task008_core_alpha_architecture/`

The snapshot must include at minimum:

- `START_STATE.md`
- `git-status-before.txt`
- `git-diff-before.patch`
- `permissions-before.txt`
- `affected-files.txt`
- `manifest.txt`
- `SHA256SUMS`
- `restore.sh`
- preserved copies of every existing file that Task 008 may modify

Requirements:

- use UTC timestamp format consistent with the repository;
- record newly created files separately from pre-existing files;
- validate checksums after snapshot creation;
- statically validate the restore script;
- ensure rollback is limited to Task 008 outputs and does not erase unrelated working-tree changes;
- do not execute the restore script during normal completion;
- preserve the snapshot as audit evidence after successful completion;
- disclose any file that cannot be preserved because of permissions.

## Risk Assessment

Record and rate at least:

- security risk;
- privilege and permission risk;
- data-model risk;
- privacy and secret-exposure risk;
- architecture lock-in risk;
- platform compatibility risk;
- documentation-authority risk;
- scope-expansion risk;
- rollback risk;
- implementation-readiness risk.

Expected overall task risk is low for host stability and medium for future architectural impact.

Use conservative decisions. Prefer reversible architecture boundaries, explicit interfaces, least privilege, safe defaults, local-first operation, and truthful partial results.

## Planned Work

### Phase 1 — Preflight and Evidence Baseline

- Validate the active task and lifecycle records.
- Create and verify the Task 008 snapshot.
- Read all required records.
- Build an evidence list of:
  - ratified facts;
  - mandatory requirements;
  - proposals;
  - assumptions;
  - unresolved release gates;
  - contradictions;
  - audit recommendations.

### Phase 2 — Requirement Authority and Ratification Boundary

Create a concise requirement-authority analysis that:

- prevents proposal-to-requirement promotion without approval;
- identifies which requirements are safe architecture inputs;
- records unresolved owner decisions;
- maps architecture outputs back to FR and AC identifiers;
- identifies any requirement that cannot yet be designed safely.

Do not renumber the existing 125 FR identifiers or 19 AC identifiers.

### Phase 3 — Core Alpha Architecture

Define an implementation-ready architecture covering at minimum:

1. System context and trust boundaries.
2. Core Alpha component map.
3. Agent responsibilities and non-responsibilities.
4. Console responsibilities and non-responsibilities.
5. Shared contract and schema ownership.
6. Discovery collector interface.
7. Normalization and validation pipeline.
8. Inventory domain model.
9. Provenance and freshness semantics.
10. Unknown, unavailable, unsupported, permission-denied, timeout, and partial-result states.
11. Read-only guarantee and mutation prohibition.
12. Privilege model and least-privilege rules.
13. Command execution policy.
14. Filesystem and process inspection boundaries.
15. Local storage versus transient-state boundaries.
16. Transport boundary between Agent and Console, while respecting unresolved security gates.
17. Logging, redaction, audit trail, and privacy rules.
18. Error taxonomy and recovery behavior.
19. Configuration boundary.
20. Versioning and schema-compatibility approach.
21. Test architecture.
22. Packaging and update boundaries without resolving authenticity gates.
23. Observability of the QWSG components themselves.
24. Extension strategy that does not compromise Slice 1 simplicity.

Where a final technology choice is not yet justified, define the contract and decision criteria instead of choosing by preference.

### Phase 4 — Core Alpha Slice 1

Define the first narrow vertical slice:

`Core Alpha Slice 1: Read-only Server Discovery and System Inventory`

The Slice 1 definition must include:

- user-visible purpose;
- exact included capabilities;
- explicit non-goals;
- minimum collection categories;
- permission behavior;
- supported execution assumptions;
- structured result contract;
- result provenance;
- handling of unavailable data;
- security and privacy constraints;
- Console or presentation boundary;
- testable acceptance criteria;
- traceability to existing FR and AC identifiers;
- implementation stages;
- completion definition;
- release blockers.

The minimum collection categories should be justified and may include:

- operating-system identity and version;
- kernel and architecture;
- host identity with privacy-safe handling;
- CPU summary;
- memory summary;
- filesystem and storage summary;
- network-interface summary with sensitive fields controlled;
- running-service summary where safely accessible;
- installed runtime and server-component versions where safely detectable;
- permission and capability report for the collector itself.

Do not include remediation, modification, continuous monitoring, alerting, or remote command execution.

### Phase 5 — Security and Permission Architecture

Produce a dedicated, explicit security boundary for Slice 1:

- default non-root operation;
- privileged data must be optional and separately identifiable;
- no privilege escalation initiated by the product without future explicit design and approval;
- command allowlisting or equivalent bounded execution;
- no shell-string construction from untrusted input;
- timeouts and output-size limits;
- environment sanitization;
- path and symlink safety;
- secret and credential redaction;
- no collection whose necessity is not documented;
- distinction between inventory facts and sensitive operational data;
- truthful reporting when access is denied.

### Phase 6 — Decision Records and Gates

Create architecture decision records only for decisions supported by current owner authority and evidence.

Create a separate architecture gate register for unresolved matters, including at least the previously disclosed release gates where relevant:

- supported platform matrix;
- e-mail transport;
- retention policy;
- Console security model;
- update authenticity and signing;
- any other gate found by Tasks 006 or 007.

Each gate must state:

- why it remains open;
- what it blocks;
- available options;
- recommended decision criteria;
- whether Slice 1 implementation can proceed before resolution.

### Phase 7 — Roadmap and Implementation Handoff

Define the smallest safe implementation sequence following Task 008.

Recommend, but do not create or activate, later tasks such as:

- Task 009 — Core Alpha Slice 1 Implementation;
- Task 010 — Slice 1 Verification and Hardening.

The implementation handoff must identify:

- prerequisites;
- permitted source areas;
- required test layers;
- prohibited shortcuts;
- exact architecture documents that will govern implementation.

### Phase 8 — Verification and Lifecycle Completion

- validate internal links;
- validate unique FR and AC references;
- confirm no requirement identifiers were lost or duplicated;
- check for placeholders and accidental draft markers;
- check documentation authority and naming consistency;
- validate Markdown formatting;
- run `git diff --check`;
- verify snapshot checksums and restore syntax;
- run all safe existing repository checks;
- update lifecycle records;
- mark Task 008 complete only if every completion criterion passes.

## Rollback Plan

The Task 008 restore script must:

1. restore every pre-existing file modified by Task 008 from its preserved copy;
2. remove only files newly created by Task 008;
3. preserve unrelated changes that existed before Task 008;
4. avoid `git reset --hard`, broad cleanup commands, recursive deletion outside the bounded task file set, or destructive repository-wide operations;
5. verify restored file checksums where possible;
6. report permission failures without attempting unsafe escalation.

The final delivery must disclose the exact restore command.

## Deliverables

Create or update an authoritative, non-duplicative set of documents. At minimum, produce:

1. `docs/architecture/CORE_ALPHA_ARCHITECTURE.md`
2. `docs/architecture/CORE_ALPHA_SLICE_1.md`
3. `docs/architecture/CORE_ALPHA_DATA_MODEL.md`
4. `docs/security/CORE_ALPHA_SECURITY_MODEL.md`
5. `docs/development/CORE_ALPHA_IMPLEMENTATION_PLAN.md`
6. `docs/development/ARCHITECTURE_GATES.md`
7. `docs/development/REQUIREMENTS_ARCHITECTURE_MAPPING.md`
8. architecture decision records under a suitable existing or newly established ADR location, only where decisions are genuinely ratified
9. `ai/history/008_2026-07-20_core-alpha-architecture.md`

Update as necessary:

- `README.md`
- `CHANGELOG.md`
- `ai/core/05_SYSTEM_MAP.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/13_ROADMAP.md`
- `ai/projects/QWSG.md`
- `ai/projects/QWSG/QWSG_MASTER_PLAN.md`
- the active Task 008 prompt

If repository conventions require different exact filenames or paths, use the conventionally correct location and document the variance. Avoid creating two documents with the same authority.

## Verification

At minimum verify:

- `bin/job --check` succeeds before and after work;
- the active prompt resolves to Task 008;
- snapshot checksums pass;
- the restore script passes shell syntax validation;
- every deliverable exists and is readable;
- no `[REQUIRES HUMAN EDITING]`, `[PENDING]`, accidental template placeholder, or unsupported completion marker remains in Task 008 outputs;
- no production source code was added;
- no package or service was installed;
- no host configuration was changed;
- every Slice 1 capability traces to existing requirements or is explicitly identified as a proposed architecture support mechanism;
- all 125 FR and 19 AC identifiers remain unique in their authoritative source;
- architecture mapping references valid identifiers;
- proposals remain distinguishable from mandatory requirements;
- open gates are not silently presented as resolved;
- the read-only and least-privilege guarantees are explicit and testable;
- error, partial-result, unsupported, and permission-denied semantics are defined;
- internal Markdown links resolve;
- `git diff --check` succeeds;
- unrelated pre-existing working-tree changes remain untouched;
- permissions are preserved unless explicit approval exists;
- no commit or push occurred.

## Documentation Updates

Update documentation only where necessary to:

- establish the new architecture documents as authoritative;
- link product requirements to architecture;
- record Task 008 completion;
- place Core Alpha Slice 1 correctly in the roadmap;
- disclose unresolved gates;
- prepare a safe implementation handoff.

Do not rewrite stable philosophy or product-definition text merely for style. Do not duplicate authority across multiple files.

## Completion Criteria

Task 008 is complete only when:

1. the Core Alpha architecture is coherent and implementation-ready;
2. Core Alpha Slice 1 has a narrow, testable, read-only boundary;
3. component, data, security, permission, error, and interface models are defined;
4. architecture outputs trace to existing functional requirements and acceptance criteria;
5. proposals and unresolved gates remain visibly separated from ratified requirements;
6. the architecture does not depend on unapproved host mutation, privilege escalation, cloud service, or technology choice;
7. a later implementation task can determine exactly what may be built and what remains forbidden;
8. all verification checks pass;
9. the snapshot and bounded rollback are valid;
10. lifecycle and history records are complete;
11. no implementation, commit, push, package installation, or infrastructure change occurred;
12. the final report states whether Task 009 may safely begin and lists every condition that still blocks it.

If these conditions are not met, leave the task incomplete and report the exact blockers.

## Final Delivery Report

Report in Hungarian:

- starting-state summary;
- files created and modified;
- architecture decisions made;
- proposals not ratified;
- open architecture and release gates;
- Slice 1 exact boundary;
- requirement and acceptance mapping status;
- security and permission conclusions;
- verification commands and results;
- permission limitations;
- pre-existing working-tree changes preserved;
- exact rollback command;
- whether Task 009 is recommended;
- confirmation that no commit or push occurred.

## Owner Approval Requirements

This prompt is explicitly approved by Attila, Project Owner, for Task 008.

The task may begin within the exact scope above.

Any scope expansion, implementation work, destructive action, privilege escalation, package installation, infrastructure modification, secret access, unsupported product-policy decision, commit, or push requires separate explicit owner approval.
