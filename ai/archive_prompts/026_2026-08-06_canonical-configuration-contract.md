# Current Engineering Task 026: Canonical Configuration Contract

## Task Metadata

- Task ID: `026`
- Task slug: `canonical-configuration-contract`
- Status: `complete — Canonical Configuration Contract and all required gates passed`
- Date opened: `2026-08-06` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

Canonical Configuration Contract


## Objective


Establish the permanent Canonical Configuration Contract of Quantum Wizard Server Guardian.

The contract shall become the single deterministic, versioned, immutable, and explainable source of effective runtime configuration for the existing canonical pipeline and future operational components.

Task 026 shall define configuration semantics before Professional Scheduler implementation begins. Future Task 027 shall consume canonical scheduling configuration and shall not define configuration precedence, identity, validation, or interpretation independently.



## Scope


Task 026 shall design and implement the Canonical Configuration Contract as a bounded, presentation-independent engineering-core component.

The task scope includes:

- defining Canonical Configuration Model 1.0;
- defining Configuration Source Record 1.0 and Effective Configuration 1.0;
- defining stable configuration, profile, section, check, target, schedule, rule-profile, policy-profile, and source identities where applicable;
- defining exact schema, contract, profile, and content-version semantics;
- defining deterministic source precedence consistent with the Functional Specification: command-specific temporary override, explicitly activated local override, primary local configuration, then documented built-in default;
- preserving field-level provenance so every effective value identifies its source, source version, precedence, and resolution result;
- defining deterministic merge behavior, explicit replacement behavior, duplicate handling, and conflict detection without using input order as an undocumented tie-breaker;
- defining immutable configuration for instance identity, locale, time zone, enabled checks and targets, canonical Rule and Policy profile selection, snapshot retention, schedules, execution timeouts, bounded concurrency, bounded retry policy, report policy, and future-compatible typed sections;
- defining a scheduler-ready Schedule Definition 1.0 with stable identity, enabled state, time-zone basis, trigger kind, normalized interval or calendar fields, priority, misfire policy, overlap policy, execution timeout, retry reference, and applicability boundaries;
- defining schedule validation without executing schedules, reading the clock, starting timers, or calculating operational due work;
- defining canonical duration, size, percentage, count, time-zone, calendar, identifier, and enum normalization where used by the model;
- defining explicit unknown-key, unsupported-version, unsupported-extension, missing-value, invalid-reference, duplicate-identity, contradictory-value, overflow, and resource-limit behavior;
- defining typed secret references and redacted metadata without implementing secret storage, secret resolution, or accepting secret material in canonical configuration;
- defining deterministic defaulting with every built-in default declared, versioned, inspectable, and represented in provenance;
- defining canonical validation and immutable resolution from Configuration Source Records to one Effective Configuration;
- defining deterministic ordering, stable content identity, and byte-stable canonical JSON serialization where applicable;
- preserving offline operation, privacy, redaction, bounded-resource behavior, localization boundaries, and AI independence;
- integrating Effective Configuration into the existing canonical pipeline only where required to replace hard-coded configurable inputs while preserving the single `internal/pipeline` execution path;
- preserving existing CLI behavior and public contracts through compatible defaults or narrowly additive versioned extensions;
- adding focused unit, contract, precedence, provenance, schedule-definition, conflict, compatibility, determinism, serialization, integration, privacy, bounded-resource, and regression tests;
- creating permanent Canonical Configuration Contract architecture documentation;
- updating only directly affected permanent architecture, system-map, roadmap, engineering-history, README, and lifecycle documentation.

The Configuration component shall remain a pure model, validation, and resolution boundary. It may normalize and resolve supplied immutable configuration sources. It shall not collect host evidence, execute the engineering pipeline, schedule work, activate files, mutate runtime state, or perform operational actions.



## Out of Scope


Task 026 shall not implement:

- Professional Scheduler or schedule execution;
- due-work calculation, timers, polling loops, cron daemon behavior, missed-run execution, locking, worker pools, or background processing;
- monitoring daemon or service lifecycle;
- Alert Engine, incidents, notifications, email, webhooks, or delivery channels;
- configuration file discovery, filesystem watching, hot reload, activation transaction, configuration editor, or Configuration UI;
- a secrets backend, credential storage, secret resolution, or secret material;
- Dashboard, Terminal UI, REST API, remote management, fleet configuration, or network calls;
- licensing, edition enforcement, automation actions, remediation, remote execution, or host mutation;
- package installation, service installation, privilege changes, deployment, or release;
- AI or Machine Learning;
- changes to Inventory, Compare, Drift, Health, Rule, Policy, Report, or Command semantics.

Task 026 shall not select a user-facing configuration file syntax or storage backend. Canonical JSON is an interchange and identity representation, not authorization for file activation or persistence behavior.



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


Before modifying any Task 026 implementation or documentation target:

- verify the QWSG repository identity and configured project root;
- verify branch `main`, HEAD, configured remote, remote URL, and local-to-remote relationship;
- verify the canonical lifecycle is valid and Task 026 is the sole active approved task;
- verify Task 025 is the latest completed task and canonical implementation baseline;
- verify all Task 025 Policy deliverables required by the pipeline are present and repository validation passes;
- verify the Framework 1.0 project configuration and validation commands remain authoritative;
- record the complete working-tree state, including pre-existing unstaged and untracked owner content;
- verify intended Task 026 targets do not overwrite or absorb unrelated owner work;
- verify no unexplained lifecycle collision, destination collision, repository divergence, dependency change, or material architecture mismatch exists;
- report every failed precondition and stop only when it affects authority, safety, rollback, scope, or the correctness of the planned implementation.

The QWCS architecture and Migration Blueprint guide engineering progress and minimal intervention but do not replace or modify the current Framework 1.0 lifecycle during Task 026.



## Snapshot Requirements


Before modifying any Task 026 implementation or documentation target:

- create a rollback-capable snapshot outside the repository for every existing target that Task 026 may modify;
- record verified absence for every new target;
- capture repository identity, branch, HEAD, remote relationship, lifecycle state, complete Git status, target inventory, and relevant validation baseline;
- generate a deterministic manifest and SHA-256 checksum list;
- verify every checksum and the readability of every snapshot payload;
- write an exact bounded restore procedure covering modified and newly created targets;
- define collision checks that prevent rollback from overwriting later owner work;
- retain the snapshot through Task 026 completion and Owner acceptance;
- keep payloads, private paths, host data, and unpublished source copies outside Git in accordance with the Engineering Backup Policy.

The snapshot boundary shall be proportional to the actual Task 026 target set. A broad repository rewrite or destructive recovery shortcut is prohibited.



## Risk Assessment


Primary risks:

- creating configuration semantics inside Scheduler or presentation code instead of the canonical contract;
- introducing hidden defaults or non-explainable precedence;
- allowing equal-precedence input order to resolve conflicts;
- coupling canonical configuration to one file format, storage backend, CLI, or future interface;
- placing mutable activation, filesystem, clock, or operational behavior inside the configuration boundary;
- accepting secrets rather than protected secret references;
- changing current pipeline, CLI, Rule, Policy, or Report behavior;
- creating schedule syntax that is ambiguous across time zones, daylight-saving transitions, locales, or unsupported calendar values;
- permitting unbounded configuration, inheritance, reference graphs, strings, collections, retries, concurrency, or durations;
- making Task 027 reinterpret configuration instead of consuming Effective Configuration.

Risk mitigation:

- keep configuration normalization, validation, and resolution pure and deterministic;
- define exact precedence, conflict, provenance, default, and unsupported-version behavior;
- reject ambiguous or contradictory equal-precedence values explicitly;
- use stable typed schedule fields and explicit time-zone and misfire semantics;
- prohibit schedule execution and all operational side effects;
- represent only secret references and enforce redacted serialization;
- impose documented resource bounds and validate all references and identities;
- preserve existing behavior through versioned compatible defaults and regression tests;
- keep `internal/pipeline` as the only canonical engine orchestrator;
- require Task 027 to consume the Task 026 contracts without redefining them.



## Planned Work


Task 026 shall:

1. inspect the implemented Task 025 baseline, current configuration-related requirements, and all hard-coded configurable pipeline inputs;
2. define the configuration boundary, taxonomy, identities, versions, sources, precedence, provenance, defaults, conflicts, references, extensions, and resource limits;
3. define Canonical Configuration Model 1.0, Configuration Source Record 1.0, Effective Configuration 1.0, and Schedule Definition 1.0;
4. implement immutable normalization, validation, resolution, canonical identity, and canonical serialization;
5. implement explicit behavior for malformed, incomplete, contradictory, unsupported, unknown, inapplicable, redacted, and missing configuration;
6. provide a deterministic built-in default source that preserves current supported CLI and pipeline behavior;
7. integrate Effective Configuration through the existing pipeline dependency boundary only where required for configurable Rule profiles, Policy profiles, retention, and future scheduler input;
8. prove that no second pipeline, Scheduler, daemon, activation, secret backend, network, process, or host-mutation path exists;
9. add focused tests for contracts, precedence, field provenance, defaults, conflicts, references, schedules, time-zone semantics, compatibility, determinism, canonical JSON, privacy, resource limits, integration, and regression;
10. create `docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md` and update only directly affected canonical documentation;
11. run every required implementation, framework, lifecycle, formatting, race, static-analysis, documentation, rollback, and repository validation;
12. record exact implementation, evidence, unresolved issues, rollback instructions, and Git state in Task 026 history;
13. complete and archive Task 026 transactionally, leaving the repository in canonical idle lifecycle state without staging, committing, or pushing unless separately authorized.



## Rollback Plan


If implementation or validation fails:

- stop further mutation and preserve the failure evidence;
- verify the Task 026 snapshot manifest, checksums, payload readability, and target inventory;
- compare every affected target with the snapshot and current working tree;
- refuse rollback over later or unrelated owner changes;
- restore only verified pre-existing Task 026 targets from the bounded snapshot;
- remove only verified Task 026-created targets whose pre-task absence was recorded and which contain no later owner work;
- preserve truthful Task 026 lifecycle and history evidence rather than rewriting completed or failed facts;
- rerun framework, lifecycle, repository, Go, race, vet, formatting, documentation, and focused configuration validations after restoration;
- report resulting repository consistency and any unresolved condition.

Broad `git reset`, `git checkout`, `git restore`, `git clean`, wildcard deletion, repository-wide extraction, and silent removal of legacy or owner data are prohibited.



## Deliverables


- Canonical Configuration Contract implementation;
- Canonical Configuration Model 1.0;
- Configuration Source Record 1.0;
- Effective Configuration 1.0;
- Schedule Definition 1.0 suitable for consumption by Task 027;
- deterministic configuration source taxonomy and precedence model;
- field-level provenance and resolution results;
- conflict-detection and conflict-resolution rules;
- versioned built-in defaults preserving current supported behavior;
- typed secret-reference and redaction contract;
- deterministic canonical identity and serialization;
- pipeline configuration integration without a second execution path;
- permanent architecture documentation;
- focused engineering tests;
- updated directly affected canonical documentation and Task 026 history.



## Verification


Successful completion requires:

- Canonical Configuration normalization and resolution are deterministic and immutable;
- equal semantic inputs produce equal Effective Configuration identity and byte-stable canonical JSON regardless of source enumeration or map order;
- every effective value has complete source and default provenance;
- precedence is exact, explainable, and consistent with `FR-CFG-002`;
- equal-precedence contradictions fail explicitly and are never resolved by incidental input order;
- malformed, incomplete, duplicate, unsupported, unknown, contradictory, oversized, cyclic, and invalid-reference inputs produce deterministic bounded failures;
- Schedule Definition validation covers interval and calendar boundaries, time-zone basis, daylight-saving policy representation, misfire policy, overlap policy, timeouts, retries, priority, and disabled or inapplicable states without executing a schedule;
- serialization and diagnostics contain no secret material and preserve redaction guarantees;
- default Effective Configuration preserves existing supported CLI commands, pipeline stage order, canonical engine behavior, and public output contracts;
- Rule and Policy engines continue to own their existing semantics;
- `internal/pipeline` remains the only canonical engine orchestrator;
- source audits prove no Scheduler, daemon, timer, polling, activation, file-watching, notification, network, process-execution, remediation, AI, or host-mutation behavior was introduced;
- focused configuration package tests pass;
- integration and regression tests pass;
- `make build` passes;
- `make test` passes;
- `go test -race ./...` passes;
- `make vet` passes;
- `make fmt-check` passes;
- `ai/scripts/framework-check.sh --run-validations` passes;
- `make engineering-test` passes;
- `bin/job --check`, `ai/scripts/next-task.sh --check`, and `bin/job --check-test-tasks` pass at their applicable lifecycle phases;
- `git diff --check` and documentation-reference audits pass;
- snapshot checksums, payload readability, target inventory, absence records, and guarded restore procedure pass;
- final Git evidence identifies exact changed, staged, and untracked paths and confirms unrelated owner content remained untouched.



## Documentation Updates


Create:

- `docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md`.

Update only where directly affected:

- `README.md`;
- `ai/core/04_ARCHITECTURE.md`;
- `ai/core/05_SYSTEM_MAP.md`;
- `ai/core/07_ENGINEERING_HISTORY.md`;
- `ai/core/13_ROADMAP.md`;
- `docs/PRODUCT_ARCHITECTURE.md` only if a clarification is strictly required without changing approved product direction;
- `docs/FUNCTIONAL_SPECIFICATION.md` only if traceability requires a non-semantic reference correction;
- directly affected canonical Command, Pipeline, Policy, Rule, Report, configuration, or development mapping documentation;
- the Builder-generated Task 026 prompt and matching dated history record through the canonical lifecycle.

Do not rewrite unrelated architecture, historical tasks, framework documents, QWCS architecture, or the QWCS Migration Blueprint.



## Completion Criteria


Task 026 is complete only when:

- the Canonical Configuration Contract is the sole authoritative source of effective configuration semantics for QWSG;
- Canonical Configuration Model 1.0, Configuration Source Record 1.0, Effective Configuration 1.0, and Schedule Definition 1.0 are implemented, versioned, validated, immutable, deterministic, and documented;
- every effective value is traceable to an exact source or declared built-in default;
- source precedence, conflicts, references, extensions, unsupported versions, redaction, and resource limits have exact tested behavior;
- current supported CLI and canonical pipeline behavior remain compatible;
- future Task 027 can consume canonical schedule definitions, timeouts, retry limits, concurrency, priorities, and policies without defining configuration semantics;
- Professional Scheduler, daemon behavior, schedule execution, Alert Engine, configuration activation, secret storage, network behavior, and host mutation remain unimplemented;
- all mandatory engineering validations pass;
- rollback evidence is complete and verified;
- documentation and Task 026 history truthfully describe the delivered boundary;
- no unrelated owner content is modified, staged, removed, or absorbed;
- no files are staged, committed, or pushed unless separately authorized;
- Task 026 is completed and archived and the repository returns to canonical idle lifecycle state.



## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-06 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
