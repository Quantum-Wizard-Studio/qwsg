# Current Engineering Task 016: Inventory Persistence & Digital Twin Foundation

## Task Metadata

- Task ID: `016`
- Task slug: `inventory-persistence-digital-twin-foundation`
- Status: `complete`
- Date opened: `2026-07-24` UTC
- Human authority: Attila — Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Inventory Persistence & Digital Twin Foundation


## Objective

Transform the current one-shot Canonical System Inventory into QWSG's first persistent Inventory Store and Digital Twin foundation.

Design and implement a minimal, versioned, deterministic, privacy-safe on-disk storage layer that can persist a completed validated inventory snapshot and load previously stored snapshots without changing the meaning or compatibility of Inventory 1.0 or the canonical inventory model.

The persisted Digital Twin shall represent the canonical last known observed state of a managed server. It remains evidence, not a health verdict, desired state, monitoring stream, or authorization to mutate the host.


## Scope

- Define the Inventory persistence architecture and the boundary between collection, validation, serialization, storage, and loading.
- Define a canonical, versioned on-disk storage layout for privacy-safe Inventory snapshots and store metadata.
- Define the Digital Twin persistence envelope, including format identity, schema/version compatibility, snapshot identity, subject identity, creation time, integrity metadata, and the embedded validated Inventory 1.0 document.
- Implement a minimal file-backed Inventory Store using the Go standard library only.
- Persist only fully assembled snapshots that pass the existing Inventory and Canonical System Inventory validators.
- Write snapshots atomically using a same-directory temporary file, restrictive permissions, synchronization where supported, and a no-partial-success contract.
- Load and validate previously stored snapshots without silently repairing, reinterpreting, or overwriting incompatible or corrupt data.
- Provide deterministic snapshot naming and deterministic listing/order behavior.
- Define and implement a bounded retention policy appropriate to a manually invoked, non-daemon foundation. Retention must never silently remove the last known valid snapshot and must be testable without wall-clock races.
- Provide integrity verification using standard-library cryptographic hashing and explicit metadata; integrity failure must be reported as data corruption, not as an empty or healthy state.
- Add an explicit, operator-invoked CLI path for persistence and retrieval if required by the architecture. Existing `qwsg inventory` stdout behavior and exit-code semantics must remain compatible.
- Keep collectors independent of persistence. Storage integration must occur only after collection, redaction, assembly, and validation.
- Add unit, integration, adversarial, compatibility, atomic-write, corruption, permission, and deterministic-output tests using temporary fixtures.
- Update canonical architecture, developer, and English/Hungarian user documentation.
- Produce Task 016 engineering history and a delivery report with exact validation and rollback evidence.


## Out of Scope

- No monitoring daemon, resident process, service unit, timer, scheduler, polling loop, or automatic background execution.
- No health scoring, policy evaluation, change detection, configuration-drift analysis, alerting, notifications, incident generation, or remediation.
- No web interface, REST API, network listener, remote upload, cloud synchronization, or external presentation layer.
- No database, embedded database engine, SQL layer, or new external dependency.
- No implementation of snapshot comparison; the format may support a future comparison layer without implementing that behavior.
- No general-purpose backup system or repository backup redesign.
- No secret collection, raw evidence persistence, credential storage, or weakening of redaction and privacy rules.
- No root requirement, privilege escalation, host mutation, or writes outside the explicitly selected QWSG state directory and test fixtures.
- No collector redesign or unrelated collector modification. Collector changes are allowed only when strictly necessary to preserve an existing validated contract, and every such change requires explicit justification and regression coverage.
- No incompatible removal, rename, or semantic change to Inventory 1.0 fields, `canonical_inventory`, collector descriptors, collector results, existing CLI output, or exit codes.
- No broad engineering-framework development unless a verified blocker prevents this product task.
- No Task 017 creation, preparation, or implementation.


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

- Begin from `main` and verify the configured canonical HTTPS remote, exact HEAD, tags, ahead/behind state, and complete Git status.
- Verify the Reusable Engineering Framework version and run `ai/scripts/framework-check.sh`, `bin/job --check`, `ai/scripts/next-task.sh --check`, and `bin/job --check-test-tasks`.
- Confirm Task 015 is complete and archived, Task 016 is the only active task after builder installation, and Task 017 does not exist.
- Record all pre-existing tracked modifications and untracked files; preserve unrelated backup directories and task-maker source files.
- Inspect `ai/core/12_INVENTORY_ARCHITECTURE.md`, `docs/architecture/CANONICAL_SYSTEM_INVENTORY_V1.md`, `docs/development/CANONICAL_SYSTEM_INVENTORY.md`, both canonical inventory user guides, `internal/inventory`, `internal/app`, `internal/collector`, `cmd/qwsg`, and their tests.
- Verify the current command remains one-shot: `qwsg inventory` emits a validated Inventory 1.0 JSON envelope with additive `canonical_inventory`, uses exit code `0` for complete, `2` for partial, and `1` for failed, and performs no persistence.
- Record current package boundaries, test counts, file permissions, Inventory schema constants, canonical profile identifiers, and privacy/redaction behavior.
- Run the complete baseline test, vet, format, framework, lifecycle, builder, and diversion suites before implementation. Stop on any unexplained failure or material mismatch.


## Snapshot Requirements

- Before modifying repository files, create a unique UTC-dated rollback snapshot outside the repository under `/tmp`.
- Include every intended source, test, documentation, task-history, and delivery-report target; record verified absence for every proposed new path.
- Capture branch, HEAD, remote configuration, ahead/behind state, complete Git status, relevant file modes, baseline validation output, and hashes of all affected inputs.
- Include a manifest, SHA-256 checksums, retention decision, and exact bounded restore instructions.
- Verify every copied file hash and restore target before implementation.
- Do not include secrets, credentials, raw host inventory, unrelated backup payloads, build artifacts, or untracked owner content.
- Retain the snapshot through Project Owner acceptance and use only exact-path restoration. Broad reset, checkout, restore, clean, wildcard deletion, or worktree overlay is prohibited.


## Risk Assessment

- Persistence changes the system from stateless observation to durable handling of potentially host-identifying operational data. Persist only post-redaction validated documents with restrictive permissions.
- Interrupted writes or full filesystems could leave partial state. Use same-directory temporary files, bounded writes, explicit flush/close handling, atomic rename, cleanup, and failure-injection tests.
- Corrupt or tampered snapshots could be mistaken for current truth. Verify envelope schema, semantic validation, integrity metadata, and subject identity before use; fail closed without replacing the last valid snapshot.
- Schema evolution could make stored snapshots unreadable or silently lossy. Define supported major/minor ranges and explicit rejection or migration boundaries; do not silently migrate.
- Non-deterministic JSON, filenames, or ordering could undermine hashes and repeatability. Define canonical persisted bytes and stable ordering before hashing.
- Retention could delete valuable evidence or grow without bound. Define a conservative bounded policy, protect the last valid snapshot, and test selection and deletion in isolated fixtures.
- Symlinks, path traversal, races, permissive modes, or unsafe directory ownership could redirect writes or expose data. Validate roots and path components, reject unsafe file types, and test permission and symlink cases.
- CLI integration could break existing Inventory 1.0 consumers or exit codes. Preserve the existing command unchanged by default and add compatibility tests.
- Collector coupling could weaken isolation. Keep storage outside collectors and reject any design that makes collectors persist or load data.
- Tests that use live host data could leak identifiers. Use synthetic privacy-reviewed fixtures and temporary directories only.


## Planned Work

### Phase 1 — Baseline and design

- Complete starting-state verification and the rollback snapshot.
- Inventory existing schema, canonical model, collection, CLI, privacy, and validation boundaries.
- Resolve the minimal store API, state-directory ownership, persisted envelope, file format, filename rules, integrity model, compatibility behavior, and conservative retention policy.
- Record architecture decisions and explicitly preserve the collector/persistence boundary.

### Phase 2 — Persistence contracts

- Define versioned store metadata and persisted Digital Twin structures without duplicating the canonical host model.
- Define UTF-8 JSON serialization, stable ordering, timestamps, identifiers, permissions, size limits, and integrity verification.
- Define errors for missing, incompatible, corrupt, unsafe, permission-denied, and partial-write states.
- Define load behavior that never converts missing or invalid evidence into a healthy state.

### Phase 3 — File-backed Inventory Store

- Implement the store in a focused internal package using only the Go standard library.
- Implement validated save, atomic installation, deterministic listing, validated load, integrity checking, and bounded retention.
- Protect the last known valid snapshot and ensure failed writes do not damage existing state.
- Keep storage paths explicit and testable; do not embed environment-specific production paths in core logic.

### Phase 4 — Explicit application integration

- Add the minimum explicit operator-invoked integration required to save and load snapshots.
- Preserve the existing `qwsg inventory` output, one-shot behavior, privacy rules, status meaning, and exit codes.
- Do not start a daemon, scheduler, listener, background goroutine lifecycle, or automatic collection loop.

### Phase 5 — Verification and hardening

- Add deterministic fixtures and tests for round trips, compatibility, atomicity, corruption, truncation, hash mismatch, permissions, unsafe paths, symlinks, retention, concurrent/conflicting access where relevant, partial inventories, and privacy.
- Run race-sensitive tests where supported and prove no external dependency or database was added.
- Run all existing collector, inventory, application, CLI, engineering-framework, lifecycle, diversion, builder, vet, format, and Git checks.

### Phase 6 — Documentation and delivery

- Update canonical architecture and developer documentation.
- Update English and Hungarian user documentation for explicit persistence behavior, paths, privacy, retention, recovery, and limitations.
- Record exact changed files, compatibility evidence, validation results, known limitations, rollback procedure, and Project Owner acceptance requirements.
- Stage only reviewed Task 016 paths and do not create Task 017.


## Rollback Plan

- Stop immediately on a mandatory validation failure, unexpected lifecycle change, unsafe storage behavior, data-loss risk, privacy regression, or compatibility failure.
- Before restoration, capture the exact failure, Git status, affected test fixture paths, and hashes without collecting live host inventory or secrets.
- Remove only Task 016-created temporary test artifacts and new repository paths whose verified pre-task state was absent.
- Restore only Task 016-modified files and executable modes from the verified snapshot using explicit paths.
- Do not delete or rewrite any operator-created inventory store unless it was created in an isolated Task 016 test fixture. Production-like store data requires separate owner authorization before destructive recovery.
- Rerun baseline framework, lifecycle, Go test, vet, format, CLI compatibility, inventory privacy, and collector regression checks after restoration.
- Confirm Task 016 lifecycle records remain truthful and unrelated untracked files remain untouched.
- Never use broad Git reset, restore, checkout, clean, wildcard deletion, force-push, or history rewriting as rollback.


## Deliverables

- Canonical Inventory persistence and Digital Twin storage architecture.
- A documented, versioned persisted snapshot and store-metadata format.
- A minimal file-backed Inventory Store implemented with the Go standard library.
- Safe save, load, list, integrity-validation, and bounded-retention behavior.
- Explicit one-shot application or CLI integration that preserves existing `qwsg inventory` compatibility.
- Unit, integration, adversarial, corruption, atomicity, permission, privacy, compatibility, and regression tests.
- Updated canonical architecture, developer documentation, and English/Hungarian user documentation.
- Task 016 history, verified snapshot and rollback evidence, validation record, and delivery report.


## Verification

- Run `ai/scripts/framework-check.sh --run-validations`.
- Run `make test`, `make vet`, `make fmt-check`, and `make engineering-test`.
- Run `bin/job --check`, `ai/scripts/next-task.sh --check`, and `bin/job --check-test-tasks`.
- Run all new Inventory Store unit and integration tests, including deterministic round-trip and persisted-byte golden tests.
- Verify save rejects an invalid Inventory envelope and load rejects unsupported versions, malformed JSON, duplicate keys where the decoder boundary can detect them, truncation, integrity mismatch, unsafe file types, and path traversal.
- Verify an injected failure before atomic rename leaves the previous valid snapshot readable and creates no partial canonical file.
- Verify restrictive directory/file permissions and safe handling of symlinks and conflicting paths.
- Verify retention is bounded, deterministic, protects the last valid snapshot, and affects only isolated test fixtures.
- Verify partial but valid Inventory 1.0 snapshots can be persisted and retain truthful status, errors, redactions, and canonical data.
- Verify existing `qwsg inventory` JSON shape, `canonical_inventory`, deterministic ordering, privacy/redaction behavior, and exit codes remain compatible.
- Verify collectors do not import or invoke persistence and no unrelated collector behavior changes.
- Verify no network listener, daemon, scheduler, timer, background service, database, or external dependency was added.
- Run `go test -race ./...` when supported by the environment and report any environmental limitation truthfully.
- Run shell syntax, UTF-8, LF, file-mode, secret/private-host, generated-artifact, and `git diff --check` reviews.
- Compare final Git status with the recorded baseline; stage only explicit reviewed Task 016 files.
- Verify Task 017 does not exist and Task 016 implementation is not claimed complete until every mandatory gate passes.


## Documentation Updates

- Update `ai/core/12_INVENTORY_ARCHITECTURE.md` only where the canonical persistence contract requires clarification; do not weaken its collector, privacy, or consumer boundaries.
- Add or update canonical architecture documentation for the Inventory Store, persisted Digital Twin envelope, directory layout, atomicity, integrity, compatibility, and retention.
- Update `docs/development/CANONICAL_SYSTEM_INVENTORY.md` with store APIs, explicit invocation, fixtures, compatibility, and troubleshooting guidance.
- Update `docs/user/CANONICAL_SYSTEM_INVENTORY.en.md` and `docs/user/CANONICAL_SYSTEM_INVENTORY.hu.md` with equivalent persistence, privacy, retention, recovery, and limitation information.
- Update README, project structure, system map, engineering history, or other canonical indexes only where required to keep references accurate.
- Record architecture decisions, configuration/path assumptions, validation evidence, known limitations, and rollback instructions in the Task 016 history and delivery report.
- Clearly label monitoring, comparison, drift analysis, health evaluation, alerts, scheduling, APIs, and web presentation as future capabilities not delivered by Task 016.


## Completion Criteria

- A versioned file-backed Inventory Store exists and persists only validated, privacy-safe Inventory snapshots.
- Stored snapshots load successfully through the same Inventory and canonical-model validation boundaries and retain their exact semantic status, provenance, issues, and redactions.
- The persisted Digital Twin envelope, directory layout, naming, integrity, compatibility, atomicity, permissions, and retention contracts are canonical and documented.
- Failed or interrupted writes cannot replace or corrupt the last known valid snapshot.
- Unsupported, corrupt, unsafe, or integrity-invalid stored data fails closed with explicit errors and is never presented as empty, current, or healthy.
- Existing Inventory 1.0 fields, additive `canonical_inventory`, collector contracts, `qwsg inventory` output, one-shot execution, and exit codes remain compatible.
- Collectors remain independent and contain no persistence responsibility.
- Non-root, read-only observation, privacy minimization, redaction-before-persistence, bounded resource use, and deterministic behavior remain enforced.
- No daemon, scheduler, background service, monitoring, comparison engine, drift analysis, health scoring, alerting, notification, API, web interface, database, network upload, or external dependency is introduced.
- All new and existing mandatory tests and validators pass with no unexplained waiver.
- English and Hungarian user documentation and canonical engineering documentation are consistent.
- Snapshot, rollback, history, delivery, Git, and validation evidence are complete.
- Only Task 016-scoped files are staged or delivered; unrelated untracked files remain untouched.
- Task 017 does not exist.
- The Project Owner has reviewed the delivery report and explicitly accepted the implementation before Task 016 is marked complete or archived.


## Owner Approval Requirements

Approved by Attila — Project Owner through the Engineering Task Builder on 2026-07-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
