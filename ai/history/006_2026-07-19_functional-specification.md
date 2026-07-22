# Task History 006: QWSG Core Alpha Functional Specification

## Task Metadata

- Task ID: `006`
- Title: `QWSG Core Alpha Functional Specification`
- Date opened: `2026-07-19` UTC
- Date completed: `2026-07-20` UTC
- Status: `complete with disclosed release gates and permission limitation`
- Responsible agent: `Aikó/Codex`
- Human authority: `Attila (Project Owner)`
- Owner communication language: `Hungarian`
- Active prompt: `ai/prompts/006_CURRENT_TASK.md`
- Dependency: completed Task 005 Product & System Blueprint

## Objective and completion boundary

Create one authoritative, internally consistent, implementation-neutral Functional Specification for QWSG Core Alpha. The deliverable defines observable behavior in enough detail to begin implementation and acceptance-test design without choosing internal architecture. Code, infrastructure, security hardening, performance work, cloud services, and component implementation remained excluded.

## Required reading completed

Read and applied the Project Philosophy, Constitution, Agent Rules, Architecture Governance, System Map, Engineering Standards, Job Standard, Delivery Policy, Documentation Policy, Security Policy, Release Policy, Roadmap, Product Definition, Product & System Blueprint, repository `AGENTS.md`, the `qwsg-job` skill, active Prompt 006, and Task Histories 001–005.

## Starting state

The verified root was `<repository-root>`, branch `main`, at HEAD `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`. `bin/job --check` passed with exactly one active prompt and matching History 006. The working tree intentionally contained uncommitted authorized results and prompt rotations from Tasks 003–006; those changes were preserved. A pre-existing Task 007 snapshot was also present and remained outside scope.

Required directories were writable with setgid and default owner/group-write ACLs. Prompt 006 was mode `0660`. History 006 pre-existed as mode `0600`; permission normalization was not authorized and was not performed.

## Snapshot location and verification

Snapshot: `ai/backups/20260720T175721Z_task006_functional_specification/`.

It contains `START_STATE.md`, exact pre-task Git status and log, a binary-capable pre-task Git patch, a target manifest, preserved copies of every existing Task 006 documentation target, SHA-256 checksums, and a root-guarded bounded restore procedure. The restore script removes only the Task 006-created Functional Specification and restores only the eight captured existing targets.

## Risk assessment

- Security, stability, and infrastructure: low; documentation only, with no runtime or host mutation.
- Product ambiguity: medium; mitigated through normative requirement IDs, explicit defaults, acceptance criteria, and release gates.
- Unauthorized architecture selection: medium; mitigated by specifying observable behavior while deferring languages, frameworks, topology, protocols, schemas, package layout, storage technology, and secrets backend.
- Open owner decisions: medium; unresolved release choices are preserved as explicit gates rather than silently inferred.
- Rollback and data loss: low; all existing targets were captured and restoration is bounded.
- Localization: low; machine tokens remain stable while every user-facing surface is required to be localizable.
- Permissions: disclosed limitation; History 006 remains mode `0600` because permission changes require separate authority.

## Planned and performed work

1. Validated the active task and recorded repository, Git, file, ownership, mode, and ACL state.
2. Read all mandatory governance, product, blueprint, and historical sources.
3. Created and integrity-verified the rollback snapshot before modifying task targets.
4. Created `docs/FUNCTIONAL_SPECIFICATION.md` with normative Core Alpha behavior and traceable acceptance criteria.
5. Defined actors, authority, operating profiles, capabilities, configuration, defaults, required checks, state and incident semantics, maintenance, alert delivery, reports, CLI, Installer lifecycle, Console behavior, persistence, diagnostics, failure isolation, security, privacy, localization, and exclusions.
6. Preserved eight unresolved product and release choices as explicit release gates that do not block contract implementation but do block a supported release.
7. Added concise repository cross-references, changelog and milestone entries, completed Prompt 006, and completed this delivery record.

## Files changed

Created:

- `docs/FUNCTIONAL_SPECIFICATION.md`
- `ai/backups/20260720T175721Z_task006_functional_specification/`

Updated:

- `README.md`
- `CHANGELOG.md`
- `ai/core/05_SYSTEM_MAP.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/13_ROADMAP.md`
- `ai/projects/QWSG.md`
- `ai/prompts/006_CURRENT_TASK.md`
- `ai/history/006_2026-07-19_functional-specification.md`

No application, executable, dependency, service, system configuration, infrastructure, Installer, Agent, Console, or cloud implementation was created or changed.

## Decisions

- Core Alpha is a coherent single-server, reporting-only protection slice; automatic remediation remains excluded.
- The normative severity model is `OK`, `WARNING`, `CRITICAL`, `EMERGENCY`, and orthogonal `UNKNOWN`, with explicit persistence, hysteresis, freshness, incident, maintenance, acknowledgement, and recovery semantics.
- Concrete default thresholds are specified where evidence semantics are universal enough; backup policy requires target-specific age and optional size configuration.
- The logical CLI command is `qwsg`, with stable human and versioned JSON behaviors and documented exit-code classes.
- Lifecycle workflows bind consent to a previewed plan and preserve operator data by default on removal.
- Exact platform versions, Console sequencing and security model, e-mail transport, retention, implementation technologies, update authenticity, and business decisions remain release gates.
- Architecture must implement the contract and may not silently weaken or reinterpret it.

## Verification evidence

- Active-task validation, required-document presence, and previous-history checks: passed before modification.
- Snapshot checksum verification and restore-script syntax: passed.
- Requirement-ID uniqueness, required section, placeholder, state-term, release-gate, traceability, and acceptance-criterion checks: passed.
- Parent Product Definition and Blueprint consistency review: passed; open decisions remain disclosed rather than asserted as approved facts.
- Documentation cross-reference and relative-path checks: passed.
- Markdown whitespace and scoped Git-diff checks: passed.
- Secret-pattern review: no credential or private-key material introduced.
- Excluded implementation and infrastructure paths: unchanged by Task 006.
- Final ownership, mode, ACL, active prompt count, and rollback validity checks: passed with the pre-existing History 006 mode limitation disclosed.

## Problems and disclosed limitations

The repository was intentionally dirty from earlier authorized, uncommitted task work. Task 006 preserved that state and used a separate bounded snapshot. A Task 007 snapshot existed before Task 006 execution and was not inspected as task authority or modified.

History 006 was created before execution with mode `0600`, so group collaborators cannot read it through normal group permissions. Task 006 did not authorize permission changes; its content is complete, but mode normalization remains an owner decision.

Eight choices remain release gates. They do not prevent implementation against the functional contracts, but a build cannot be represented as a supported production release until the gates are resolved and applicable acceptance criteria pass.

## Rollback procedure

From exactly `<repository-root>`, run `ai/backups/20260720T175721Z_task006_functional_specification/restore.sh`, review its bounded effects, and type `ROLLBACK-QWSG-TASK-006`. It removes only `docs/FUNCTIONAL_SPECIFICATION.md` and restores only README, CHANGELOG, System Map, Engineering History, QWSG project record, Prompt 006, and History 006. It preserves previous-task work, unrelated changes, application paths, planning sources, and all other snapshots.

## Git record

Starting and final HEAD remain `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`; Task 006 did not authorize a commit or push.

## Open questions and release gates

1. Initial supported Ubuntu/Debian versions and CPU architectures.
2. Console inclusion or later sequencing for the first release.
3. Required local or SMTP e-mail transport.
4. Default retention periods and storage budgets.
5. Console authentication, recovery, and exposure model.
6. Configuration syntax, secrets backend, storage technology, and topology.
7. Distribution integrity and update authenticity.
8. Licensing, editions, telemetry, commercial, and hosted-service choices.

## Recommended next task

Resolve the owner-controlled release-gate decisions needed before architecture, then authorize a separate Core System Architecture task that allocates responsibilities and trust boundaries while tracing every design decision to this Functional Specification. Do not begin implementation as part of Task 006.

## Delivery result

**Complete with disclosed release gates and permission limitation.** The authoritative Functional Specification exists, contains no unresolved template placeholders, defines testable Core Alpha behavior, preserves implementation neutrality, and is rollback-capable.
