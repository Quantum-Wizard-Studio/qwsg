# Current Engineering Task 018: Snapshot Comparison Engine

## Task Metadata

- Task ID: `018`
- Task slug: `snapshot-comparison-engine`
- Status: `complete`
- Date opened: `2026-07-24` UTC
- Human authority: Attila — Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Snapshot Comparison Engine


## Objective


Implement the first deterministic intelligence layer of Quantum Wizard Server Guardian: a canonical Snapshot Comparison Engine that analyzes system evolution between two validated Inventory snapshots without comparing raw JSON documents.

Establish one versioned Change Record model and one comparison result envelope as the only supported source of system-change information for future Configuration Drift, Health, Alert, Reporting, CLI, e-mail, API, and Web UI consumers.

Compare every Canonical System Inventory layer independently, preserve typed fact semantics, and expose identical canonical results through machine-readable JSON and a concise human presentation derived exclusively from Change Records.

The engine detects facts only. It shall never decide whether a change is good, bad, expected, unexpected, healthy, unhealthy, important, actionable, or alert-worthy.

Project Owner acceptance requires deterministic latest-versus-previous and explicit snapshot-versus-snapshot comparison, zero Added/Removed/Modified changes for semantically identical snapshots, stable Change Records for changed synthetic snapshots, unchanged Inventory compatibility, and complete administrator documentation.


## Scope


- Inspect the completed Task 017 CLI, Canonical System Inventory, Inventory Store, persistence envelope, privacy, validation, output, installation, test, and documentation boundaries before design.
- Define the Comparison Engine as a dedicated internal architectural layer between validated Inventory Store data and every future interpretation or presentation subsystem.
- Add a focused internal comparison package using only the Go standard library and existing QWSG packages.
- Accept only two fully validated Inventory snapshots with valid additive canonical inventory.
- Reject missing canonical inventory, unsupported schemas or profiles, incompatible comparison contracts, corrupt store data, subject mismatch where comparison would be misleading, and unsafe input before producing a result.
- Compare semantic Canonical System Inventory structures rather than serialized JSON bytes, map ordering, formatting, field order, or persistence-envelope metadata.
- Compare every canonical layer independently.
- Define a canonical, versioned Comparison Result envelope containing, at minimum:

  - comparison schema name and version;
  - engine version;
  - deterministic comparison identifier;
  - `from` and `to` snapshot selectors and identities;
  - privacy-safe subject identity;
  - source Inventory schema/profile information;
  - deterministic comparison timestamp;
  - ordered Change Records;
  - Added, Removed, Modified, and Unchanged counts;
  - explicit comparison metadata.
- Define a canonical Change Record containing, at minimum:

  - stable record identifier;
  - canonical layer identifier;
  - object identifier;
  - canonical object path;
  - change type;
  - typed previous value;
  - typed current value;
  - deterministic comparison timestamp;
  - comparison and provenance metadata.
- Support exactly these initial change types:

  - `added`;
  - `removed`;
  - `modified`;
  - `unchanged`.
- Define stable object identity from canonical layer/resource/fact contracts, never from array position, display text, map iteration, raw JSON offsets, or host-specific filesystem traversal.
- Define canonical object paths with an unambiguous escaped representation suitable for stable machine use.
- Treat a resource present only in `to` as Added and a resource present only in `from` as Removed.
- Compare facts belonging to resources present in both snapshots by stable fact name and typed canonical value.
- Treat facts present only in `to` as Added, facts present only in `from` as Removed, equal facts as Unchanged, and unequal typed facts as Modified.
- Compare canonical status, issue, redaction, relationship, label, and metadata fields only where they are part of the approved semantic comparison contract; document every included and excluded field.
- Exclude observation timestamps, durations, request IDs, persistence hashes, serialization details, and other expected per-observation metadata from change detection unless they are explicitly comparison-significant.
- Preserve typed values without lossy string conversion in canonical JSON.
- Define deterministic ordering across layers, objects, paths, change types, and records.
- Derive the comparison timestamp deterministically from the `to` snapshot completion timestamp; do not use the current wall clock in canonical output.
- Derive comparison and record identifiers deterministically from versioned canonical inputs with standard-library cryptographic hashing.
- Ensure repeated comparison of the same validated snapshot pair produces byte-identical canonical JSON.
- Represent Unchanged records in the canonical model so human or future consumers may include them; define “zero detected changes” as zero Added, Removed, and Modified records even when Unchanged records are present.
- Add the `qwsg compare` command family.
- Support default latest-versus-previous comparison in the explicitly selected Inventory Store.
- Support explicit `--from SNAPSHOT --to SNAPSHOT` comparison.
- Require `--from` and `--to` together; reject ambiguous half-specified selections.
- Reuse Task 017 explicit store selection through `--store` or `QWSG_STORE`; do not discover or create an implicit store.
- Support `--retention N` consistently with the existing immutable Inventory Store configuration.
- Support `--format json` and `--format human`.
- Keep machine-readable JSON as the compare command's canonical default.
- Generate human output exclusively from the canonical Comparison Result and Change Records.
- Group human output by Added, Removed, Modified, and Unchanged with deterministic ordering and concise administrator-oriented descriptions.
- Use canonical kinds, paths, and facts to produce descriptions such as kernel changes, filesystem additions, memory changes, network resource removals, or application configuration changes only when those facts actually exist in the canonical input.
- Escape terminal control characters and retain Task 017 output safety.
- Clearly identify the compared snapshots, comparison timestamp, change counts, and no-change result.
- Keep human presentation factual and avoid health judgement, recommendation, priority, severity, scoring, alerting, remediation, or policy language.
- Define consistent stdout/stderr and exit-code behavior for success with changes, success without changes, partial-but-valid source snapshots, usage error, insufficient snapshot count, missing selector, incompatible subjects, corrupt store, and comparison failure.
- Preserve all existing `qwsg inventory`, Snapshot Explorer, JSON compatibility, version, help, build, and install behavior.
- Add contextual help and examples for the compare command.
- Add unit, property-oriented, deterministic golden, integration, CLI, adversarial, compatibility, privacy, and regression tests using synthetic privacy-reviewed fixtures and isolated temporary stores.
- Test values of all supported canonical types, missing/additional layers/resources/facts, ordering permutations, control characters, Unicode, null/empty values, large bounded structures, partial snapshots, and invalid input.
- Document the Comparison Engine architecture, Change Record schema, semantic include/exclude rules, CLI use, output contracts, privacy interpretation, limitations, and future consumer boundary.
- Provide English engineering and user documentation and a functionally equivalent Hungarian user guide.
- Record complete implementation, validation, rollback, Git delivery, and Project Owner acceptance evidence in Task 018 history and an English Engineering Delivery Report.


## Out of Scope


- Raw JSON text, byte, whitespace, object-key-order, array-position, or persistence-envelope diffing.
- Configuration Drift classification.
- Desired-state comparison, baseline policy, allowlists, maintenance windows, expected-change policy, or compliance evaluation.
- Health Engine, health scoring, severity, priority, risk rating, impact analysis, or root-cause analysis.
- Alert Engine, alerts, notifications, e-mail, webhooks, incident generation, or escalation.
- Recommendations, explanations generated by AI, natural-language analysis, remediation plans, or automated corrective action.
- Scheduler, timer, polling loop, watch mode, continuous comparison, daemon, service, background process, or automatic execution after save.
- Web UI, Console, REST API, RPC, socket, network listener, remote access, upload, synchronization, or telemetry.
- Database, search index, external storage engine, or new external dependency.
- Collector implementation, registration, contract, timeout, budget, command, or evidence-acquisition changes.
- Inventory 1.0 or Canonical System Inventory schema, validation, aggregation, privacy/redaction, or exit-code changes.
- Inventory Store format, integrity, retention, atomicity, locking, permission, path, or migration changes.
- Comparing different subjects as though they represented one evolving system.
- Cross-host fleet comparison, aggregation, trend analysis, time series, or reporting dashboards.
- Persisting comparison results, caching, indexing, signing, authenticating, encrypting, or transmitting them.
- Automatic snapshot selection beyond latest-versus-previous in an explicitly selected store.
- Modifying host state, requiring root, privilege escalation, or writing outside explicit test fixtures and existing build/documentation targets.
- Package-manager artifacts or installation redesign.
- Weakening Engineering Framework, lifecycle, approval, snapshot, rollback, targeted-staging, Git, privacy, or validation rules.
- Deleting, repairing, migrating, rewriting, pruning, or otherwise mutating operator snapshot stores.
- Modifying pre-existing untracked backups or task-maker inputs.
- Starting Task 019 or implementing any future consumer of Change Records.


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


- Start from repository root `/home/qws/web/qwsg.quantumwizard.hu/qwsg`.
- Validate Engineering Framework 1.0 project identity and configured validations with `ai/scripts/framework-check.sh --run-validations`.
- Verify branch `main`, canonical HTTPS remote `origin`, exact HEAD, tags, upstream relationship, ahead/behind state, and complete Git status.
- Verify Task 017 is complete and archived, Task 018 is the sole approved active prompt after builder installation, exactly one matching Task 018 history exists, and Task 019 does not exist.
- Run `bin/job --check`, `ai/scripts/next-task.sh --check`, and `bin/job --check-test-tasks`.
- Record and preserve every pre-existing untracked path without staging, deleting, moving, cleaning, archiving, or modifying it.
- Verify Ubuntu version, kernel, Go, GNU Make, installed and repository QWSG versions, user identity, groups, umask, relevant ACLs, file ownership, modes, and writable boundaries.
- Inspect `ai/core/12_INVENTORY_ARCHITECTURE.md`, the canonical inventory architecture and developer guides, persistence architecture, Task 017 CLI/user/install documentation, Task 017 history and delivery report, `internal/inventory`, `internal/inventorystore`, `internal/app`, and `cmd/qwsg`.
- Record current Inventory 1.0 and canonical schema/profile constants, value types, resource identity rules, ordering, status semantics, privacy/redaction rules, persisted format identity, retention, and store validation behavior.
- Verify the Inventory Store exposes deterministic listing and validated exact/latest loading, and identify the smallest safe way to select latest and previous without weakening store encapsulation.
- Verify `qwsg inventory`, save, list, info, and load behavior, JSON compatibility, human rendering, contextual help, exit codes, and installed-binary operation remain as accepted under Task 017.
- Verify no comparison package, `qwsg compare` command, Task 018 artifact, or external Go dependency exists.
- Capture `go.mod`/`go.sum` state and collector, Inventory, persistence, CLI, Makefile, documentation, and test hashes before implementation.
- Run all Go package tests, `go test -race ./...` where supported, vet, format, engineering-framework, lifecycle, diversion, Task Builder, installed CLI, persistence, privacy, and Git checks.
- Capture sanitized baseline outputs and exit statuses for existing commands and the current unknown-command result for `qwsg compare`.
- Stop before implementation on any unexplained failure, lifecycle mismatch, unexpected task, repository identity difference, dependency drift, installed/repository binary mismatch that affects acceptance, or material difference from this baseline.


## Snapshot Requirements


- Before modifying repository targets, create a unique UTC-dated rollback snapshot outside the repository under the configured `/tmp` snapshot location.
- Include every intended existing comparison, Inventory, store, CLI, test, Makefile, documentation, prompt, history, audit, and canonical-index target.
- Record verified absence for each proposed new comparison package, schema document, user guide, test fixture, and delivery report path.
- Capture branch, HEAD, tags, remotes, ahead/behind state, complete Git status, relevant untracked paths, environment/tool versions, installed binary state, file ownership, modes, ACLs, UTF-8/LF state, and dependency state.
- Capture baseline command help, version, JSON compatibility, human output, store listing/loading, exit statuses, and installed-binary hash without preserving raw live host Inventory.
- Hash all affected inputs, especially canonical model, validator, store, CLI, tests, `go.mod`, `go.sum` when present, architecture, Task 018 prompt/history, and configured validators.
- Include a manifest, SHA-256 checksums, recorded-absence list, retention decision, and exact bounded restore instructions.
- Verify every copied file checksum and restore/absence target before implementation.
- Do not capture secrets, credentials, raw host identifiers, operator Inventory Stores, private runtime evidence, unrelated backup payloads, build caches, or generated binaries.
- Retain the verified snapshot through Project Owner acceptance.
- Do not begin implementation or documentation changes until snapshot integrity and rollback instructions pass.


## Risk Assessment


- Canonical-authority risk is high: divergent comparison implementations would produce conflicting system-change truth. Enforce one dedicated package and prohibit direct future snapshot diffing.
- Determinism risk is high: wall-clock timestamps, map iteration, raw JSON ordering, unstable paths, or random IDs could change output. Use canonical traversal, sorted keys, deterministic timestamps, versioned hashes, and golden repeatability tests.
- Semantic-noise risk is high: observation times, durations, request IDs, provenance timing, and persistence metadata naturally differ. Define and test explicit comparison-significant fields.
- Identity risk is high: array position or mutable display values could misclassify Modified as Added/Removed. Use canonical resource and fact identities and fail on duplicate/ambiguous identities.
- Type-safety risk is high: stringifying values could equate or confuse numbers, strings, booleans, nulls, arrays, and objects. Preserve typed canonical values and compare normalized semantic representations.
- Privacy risk is high: previous/current values can expose evidence. Compare only already validated redacted canonical data, preserve sensitivity boundaries, and keep fixtures synthetic.
- Subject-confusion risk is high: comparing different hosts could be reported as evolution. Require compatible privacy-safe subject identity and fail explicitly otherwise.
- Partial-data risk is high: unavailable collectors can make absence look like removal. Preserve source status/provenance, define incomplete-evidence metadata, and never interpret change meaning.
- Store-safety risk is high: selection or comparison could mutate stored evidence. Use read-only list/load APIs and verify store hashes/modes remain unchanged.
- Compatibility risk is high: CLI integration could break accepted Inventory or Snapshot Explorer behavior. Add exact regression and JSON schema tests.
- Human-presentation risk is medium: friendly descriptions can imply causality or health. Generate only factual text from Change Records and state limitations.
- Output-volume risk is medium: Unchanged records and large inventories can be verbose. Keep processing bounded, deterministic, and documented without silently dropping canonical results.
- Stable-ID collision risk is low but material: use SHA-256 with version/domain separation and test identifier inputs.
- External-dependency and scope-expansion risk is mandatory: reject new modules and every prohibited future subsystem.
- Rollback risk is medium: new canonical contracts may attract future consumers. Complete rollback before Task 019 and document that no downstream consumer exists yet.


## Planned Work


### Phase 1 — Baseline, snapshot, and semantic contract

- Complete starting-state verification and the verified rollback snapshot.
- Inventory canonical identities, typed values, ordering, status, privacy, store selection, and CLI contracts.
- Define comparison schema/version, Change Record fields, canonical paths, stable IDs, deterministic timestamp, metadata, counts, ordering, included semantic fields, and excluded volatile fields.
- Record explicit behavior for partial snapshots, subject mismatch, identical inputs, empty layers, missing objects/facts, and unsupported types.

### Phase 2 — Canonical comparison model

- Implement versioned Comparison Result, Snapshot Reference, Change Record, Change Type, typed value, and metadata structures.
- Implement deterministic identifier and canonical serialization inputs using standard-library hashing and explicit domain/version separation.
- Implement validation that rejects malformed results, duplicate IDs/paths, invalid ordering, inconsistent counts, illegal type/value combinations, and incompatible sources.

### Phase 3 — Comparison Engine

- Implement deterministic traversal by canonical layer, resource, and fact identity.
- Compare semantic typed values independently of JSON representation and map order.
- Emit Added, Removed, Modified, and Unchanged records with deterministic order and metadata.
- Preserve truthful partial/incomplete-evidence context without producing health, drift, or alert classifications.
- Ensure repeated and reverse-direction comparisons behave according to documented contracts.

### Phase 4 — Store selection and CLI integration

- Add `qwsg compare` using the existing explicit store selection and validated list/load boundaries.
- Implement latest-versus-previous selection and explicit paired `--from`/`--to`.
- Keep JSON canonical by default and implement human output strictly as a projection of Change Records.
- Add contextual help, safe terminal rendering, stdout/stderr separation, and stable exit behavior.
- Preserve every existing Inventory and Snapshot Explorer command unchanged.

### Phase 5 — Verification and hardening

- Add deterministic fixtures for unchanged, added, removed, modified, reordered, partial, incompatible-subject, invalid-schema, duplicate-identity, Unicode, control-character, typed-value, and bounded-large cases.
- Add byte-identical repeat tests, input-immutability tests, reverse-comparison tests, JSON golden tests, human renderer tests, CLI integration tests, and store non-mutation proofs.
- Run race, vet, format, privacy, no-dependency, no-collector-change, framework, lifecycle, builder, diversion, installed CLI, and Git gates.

### Phase 6 — Documentation and delivery

- Add canonical Comparison Engine architecture and Change Record schema documentation.
- Update system architecture and map to make the engine the sole supported system-evolution source.
- Update English/Hungarian user guides and CLI demonstration with default and explicit comparisons.
- Record exact changed paths, decisions, schemas, outputs, tests, limitations, rollback, Git evidence, and Project Owner acceptance requirements.
- Stage and deliver only Task 018 paths; do not create Task 019.


## Rollback Plan


- Stop immediately on lifecycle drift, nondeterministic output, raw JSON diffing, identity ambiguity, typed-value loss, subject confusion, privacy regression, store mutation, Inventory compatibility regression, prohibited interpretation, external dependency, or mandatory validation failure.
- Capture the exact failure, fixture identity, hashes, command, sanitized output, exit status, Git state, and affected paths without storing raw live Inventory or secrets.
- Restore only Task 018-modified files and modes from the verified snapshot using explicit paths.
- Remove only Task 018-created repository paths whose verified pre-task absence and current Task 018 identity/hash are recorded.
- Never remove, rewrite, repair, migrate, prune, or restore an operator Inventory Store.
- Remove only isolated synthetic Task 018 comparison fixtures and temporary stores created under exact verified paths.
- Rerun baseline framework, lifecycle, Go, race, vet, format, Inventory, store, CLI, privacy, install, documentation, and Git checks after restoration.
- Verify no comparison package, command, schema, generated artifact, or Task 019 remains after a full rollback.
- Preserve Task 018 prompt/history truthfully even if implementation is rolled back or blocked.
- Never use broad reset, restore, checkout, clean, wildcard deletion, recursive repository overlay, force-push, rebase, or history rewriting.
- If exact safe restoration cannot be proven, stop and request Project Owner direction.


## Deliverables


- Canonical Snapshot Comparison Engine architecture.
- Versioned Comparison Result and Change Record contracts.
- Dedicated internal comparison package with validation and deterministic engine behavior.
- Stable comparison and Change Record identifiers.
- Deterministic semantic comparison of every canonical Inventory layer.
- Added, Removed, Modified, and Unchanged Change Records with typed previous/current values.
- Latest-versus-previous and explicit snapshot-pair selection.
- `qwsg compare` contextual help and CLI integration.
- Canonical machine-readable JSON output.
- Terminal-safe human comparison report generated only from Change Records.
- Unit, deterministic golden, property-oriented, integration, CLI, adversarial, privacy, compatibility, non-mutation, and regression tests.
- Canonical architecture and Change Record schema documentation.
- Updated English/Hungarian user and demonstration documentation.
- Task 018 history, verified rollback snapshot, Engineering Delivery Report, validation evidence, Git evidence, and Project Owner acceptance record.


## Verification


- Run `ai/scripts/framework-check.sh --run-validations`.
- Run `make test`, `make vet`, `make fmt-check`, and `make engineering-test`.
- Run `bin/job --check`, `ai/scripts/next-task.sh --check`, and `bin/job --check-test-tasks`.
- Run `go test -race ./...` when supported and report environmental limitations truthfully.
- Run every new comparison model, validator, engine, deterministic golden, property-oriented, renderer, CLI, store-integration, compatibility, and adversarial test.
- Verify two semantically identical snapshots produce zero Added, Removed, and Modified records, while any emitted Unchanged records are deterministic.
- Verify changed synthetic snapshots produce the exact expected Change Records, counts, typed values, stable IDs, paths, metadata, and order.
- Compare the same pair repeatedly and across ordering permutations; require byte-identical canonical JSON.
- Verify reverse comparison swaps Added/Removed and previous/current semantics consistently.
- Verify comparison timestamp and identifiers contain no wall-clock or random input.
- Verify JSON formatting, key ordering, and record ordering are stable under the canonical serializer.
- Verify all canonical layers are traversed independently and no duplicate or unknown layer/resource/fact identity is accepted.
- Verify strings, booleans, integers, finite numbers, null/absent values, lists, maps, Unicode, escaped paths, and terminal controls retain correct typed semantics.
- Verify volatile observation metadata does not create false changes.
- Verify different subjects, missing canonical data, unsupported schemas/profiles, invalid source snapshots, corrupt stores, insufficient snapshot count, half-specified selectors, unsafe paths, and retention mismatch fail closed.
- Verify partial source snapshots remain truthfully labeled and are never converted into health, drift, or alert conclusions.
- Verify latest-versus-previous selects the two newest deterministic store entries in the correct direction.
- Verify explicit `--from`/`--to` selects exactly the named validated snapshots.
- Verify comparison leaves both input objects and every store file, hash, mode, and timestamp unchanged.
- Verify human output is derived only from canonical Change Records, grouped by Added/Removed/Modified/Unchanged, terminal-safe, deterministic, concise, and free of judgement, recommendation, severity, score, alert, or remediation.
- Verify JSON is the default for `qwsg compare` and `--format json`; verify `--format human`.
- Verify contextual help, examples, stdout/stderr separation, and exit codes for all compare modes.
- Verify every existing `qwsg inventory`, save, list, info, load, help, version, build, install, JSON, human, status, and exit-code regression remains compatible.
- Verify the installed binary supports compare after the normal-user build and artifact-only privileged install workflow.
- Verify `go.mod` and `go.sum` add no external dependency.
- Verify collectors, Inventory schemas/validators, persisted store format, retention, atomicity, integrity, permissions, privacy/redaction, and installed-system configuration are unchanged.
- Verify no Configuration Drift, Health Engine, Alert Engine, e-mail, scheduler, daemon, Web UI, REST API, database, AI analysis, recommendation, network listener, telemetry, or host mutation was added.
- Run UTF-8, LF, file-mode, ACL where relevant, secret/private-host-data, generated-artifact, documentation-link, and `git diff --check` reviews.
- Prove the exact-path rollback in an isolated or non-destructive way.
- Compare final Git status with baseline, stage only explicit Task 018 paths, review the staged diff, commit, run HTTPS dry-run push, push normally, and verify local/remote synchronization.
- Verify Task 018 remains the only active task during implementation and Task 019 does not exist.


## Documentation Updates


- Add canonical Comparison Engine architecture describing the dedicated layer, inputs, outputs, trust boundaries, determinism, privacy, and future-consumer prohibition on direct snapshot comparison.
- Add a versioned Change Record and Comparison Result schema/contract reference with field semantics, typed values, identifiers, paths, ordering, counts, compatibility, and examples.
- Update `ai/core/04_ARCHITECTURE.md` and `ai/core/05_SYSTEM_MAP.md` so the Comparison Engine is the sole supported source of system evolution.
- Update `ai/core/12_INVENTORY_ARCHITECTURE.md` only to add the comparison consumer boundary; do not change Inventory semantics.
- Update the persistence architecture only to document read-only comparison consumption; do not change store format or behavior.
- Update the canonical Inventory developer guide with comparison package boundaries, fixtures, compatibility rules, and prohibited direct diffs.
- Update CLI help, README, English/Hungarian user guides, and the demonstration walkthrough with latest-versus-previous, explicit pair, JSON, human, no-change, and limitation examples.
- Update project structure, engineering history, or other indexes only where needed for accurate references.
- Record starting state, snapshot, every changed path, architecture decision, schema/version, deterministic evidence, compatibility results, acceptance demonstration, limitation, rollback, Git delivery, and Project Owner acceptance in Task 018 history.
- Create an English Task 018 Engineering Delivery Report under the canonical audit/report location.
- Clearly label Configuration Drift, Health, Alert, e-mail, scheduler, daemon, Web UI, API, database, AI analysis, recommendations, persistence of comparisons, and Task 019 as not delivered.


## Completion Criteria


- A dedicated canonical Snapshot Comparison Engine exists and is the only supported comparison implementation.
- The engine accepts only validated compatible canonical Inventory snapshots and never compares raw JSON documents.
- Every canonical layer is compared independently using stable semantic resource/fact identities and typed values.
- Versioned Comparison Result and Change Record contracts are complete, validated, documented, and deterministic.
- Every record contains stable ID, layer, object ID, canonical path, Added/Removed/Modified/Unchanged type, typed previous/current value, deterministic comparison timestamp, and metadata.
- Repeated comparison of the same pair produces byte-identical canonical JSON, stable identifiers, stable counts, and stable ordering.
- Semantically identical snapshots produce zero Added, Removed, and Modified changes.
- Changed fixtures produce exact deterministic Change Records; reverse comparisons have consistent inverse semantics.
- Expected per-observation metadata does not create false system changes.
- Partial evidence remains truthful and no comparison result implies health, drift, expectation, severity, recommendation, or alert state.
- `qwsg compare` supports latest-versus-previous and exact `--from`/`--to` selectors in an explicitly chosen validated store.
- JSON is the canonical default and human output is derived exclusively from Change Records.
- Human output is grouped, concise, deterministic, terminal-safe, privacy-aware, and judgement-free.
- Inputs and stores remain byte-, mode-, and timestamp-unchanged by comparison.
- All Task 017 Inventory, Snapshot Explorer, help, version, build, install, JSON, human, and exit-code contracts remain compatible.
- Inventory schemas, collectors, store format and safety, privacy/redaction, non-root operation, one-shot execution, and no-external-dependency guarantees remain unchanged.
- No Configuration Drift, Health Engine, Alert Engine, notification, scheduler, daemon, Web UI, API, database, AI analysis, recommendation, comparison persistence, network activity, privilege escalation, or host mutation is introduced.
- All new and existing mandatory tests and validators pass without unexplained waiver.
- English/Hungarian user documentation, architecture, schema reference, CLI help, examples, and delivery evidence are consistent.
- Snapshot, rollback, history, delivery, security, determinism, compatibility, Git, and Project Owner acceptance evidence are complete.
- Only Task 018-scoped paths are staged and delivered; pre-existing untracked content remains untouched.
- Task 018 is the sole active task during implementation and Task 019 does not exist.
- The Project Owner explicitly accepts the final delivery before Task 018 is marked complete or archived.


## Owner Approval Requirements

Approved by Attila — Project Owner through the Engineering Task Builder on 2026-07-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
