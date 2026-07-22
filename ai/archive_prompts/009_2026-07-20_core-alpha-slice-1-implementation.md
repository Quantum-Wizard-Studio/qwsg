# Current Engineering Task 009: Core Alpha Slice 1 Implementation

## Task Metadata

- Task ID: `009`
- Task slug: `core-alpha-slice-1-implementation`
- Status: `complete`
- Date opened: `2026-07-20` UTC
- Human authority: `Project Owner`
- Owner or lead-developer communication language: `Hungarian`
- Task type: `Implementation`
- Requires snapshot: `Yes`
- Requires rollback validation: `Yes`
- Production implementation changes allowed: `Yes, only within the exact Slice 1 scope`
- Destructive actions allowed: `No`
- Commit allowed: `No`
- Push allowed: `No`

## Title

QWSG Core Alpha Slice 1 Implementation — Read-only Server Discovery and System Inventory

## Owner Decisions and Ratification

The Project Owner approves the implementation of `Core Alpha Slice 1` under the architecture completed in Task 008.

The following implementation baseline is ratified for this task:

1. Runtime: `Go`.
2. Toolchain: a repository-pinned, explicitly documented supported Go toolchain.
3. Packaging: one self-contained Linux CLI binary named `qwsg`.
4. No installer, package-manager integration, daemon, systemd unit, background service, automatic update, network listener, or Console integration is authorized.
5. Execution model: local, one-shot, non-root by default, read-only server discovery.
6. Exit codes:
   - `0` — complete and valid inventory;
   - `1` — fatal failure, no valid inventory produced;
   - `2` — partial but valid inventory produced.
7. Partial results must remain structurally valid and must identify every unavailable, unsupported, permission-denied, failed, or timed-out collection result.
8. No privilege escalation may be initiated.
9. No host configuration, service state, package state, filesystem content, account state, or network configuration may be changed.
10. Task 008 architecture, security model, data model, ADRs, implementation plan, architecture gates, and requirement mapping are authoritative for this implementation.

## Objective

Implement the first working QWSG product slice as a safe, deterministic, testable Linux command-line application that performs local, one-shot, read-only server discovery and emits a structured system inventory.

The implementation must:

- follow the Task 008 architecture;
- run without root privileges by default;
- collect only the approved inventory categories;
- preserve truthful partial results;
- apply bounded command execution and timeout rules;
- emit machine-readable structured output;
- expose clear CLI behavior and exit codes;
- include sufficient automated tests to verify the security and functional boundaries;
- avoid all host mutation.

This task creates the first production source code in the repository, but it does not create a supported public release.

## Scope

This task is authorized to:

- create the Go module and project structure required by the approved architecture;
- add the `qwsg` CLI entry point;
- implement the approved Slice 1 discovery collectors;
- implement normalization, validation, result aggregation, provenance, issue reporting, and status calculation;
- implement structured JSON output;
- implement CLI help and version output;
- implement timeout, cancellation, output-size, environment, path, and subprocess safety controls;
- implement complete, partial, and fatal result semantics;
- implement deterministic collector interfaces;
- implement unit tests;
- implement bounded integration tests that do not mutate the host;
- add test fixtures and safe fake command runners;
- add build scripts or Makefile targets if justified and non-duplicative;
- add static-analysis configuration if it is local, documented, and does not require unsafe host changes;
- document local development, build, test, and execution procedures;
- update lifecycle, system map, roadmap, history, changelog, and project documentation as needed;
- create the Task 009 snapshot and bounded restore script;
- build a local development binary in a repository build-output directory;
- execute the resulting CLI against the current host only in read-only mode and only after automated tests pass;
- capture a sanitized sample inventory that does not disclose secrets or unnecessarily sensitive host data.

Authorized source and documentation areas may include:

- `cmd/qwsg/`
- `internal/`
- `pkg/` only if the architecture explicitly requires a public package boundary
- `testdata/`
- `scripts/`
- `build/` or another repository-approved ignored output directory
- `go.mod`
- `go.sum`
- `Makefile`
- `.gitignore`
- `README.md`
- `CHANGELOG.md`
- `docs/`
- `ai/core/05_SYSTEM_MAP.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/13_ROADMAP.md`
- `ai/projects/QWSG.md`
- `ai/projects/QWSG/QWSG_MASTER_PLAN.md`
- `ai/history/`
- the active Task 009 prompt
- the Task 009 snapshot directory

Use the exact paths and component boundaries established by Task 008 where they differ from this illustrative list.

## Out of Scope

Do not:

- implement a Console, web UI, desktop UI, REST API, GraphQL API, socket API, remote API, or network listener;
- implement continuous monitoring, scheduling, polling, alerting, e-mail, notification, telemetry, analytics, or remote reporting;
- implement remediation, automatic repair, configuration management, package updates, service restart, process termination, file editing, permission changes, firewall changes, user changes, or any other host mutation;
- implement privilege escalation, `sudo` prompts, setuid behavior, root-helper binaries, capabilities, or privileged agents;
- implement an installer, `.deb`, `.rpm`, Snap, Flatpak, Docker image, systemd unit, init script, package repository, or update mechanism;
- resolve open Task 008 architecture gates beyond the owner decisions explicitly recorded above;
- implement retention, persistent database storage, secrets backend, e-mail transport, Console authentication, licensing, billing, distribution, or update signing;
- collect secrets, private keys, tokens, password hashes, environment secrets, full command histories, private file contents, or application data;
- add shell command construction from untrusted values;
- call arbitrary commands supplied by users;
- depend on network availability;
- send data off the host;
- require a clean working tree;
- overwrite unrelated pre-existing changes;
- commit or push;
- install system packages or globally modify the development environment without separate explicit owner approval.

## Required Reading

Before implementation, read and reconcile at minimum:

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
- `docs/architecture/CORE_ALPHA_ARCHITECTURE.md`
- `docs/architecture/CORE_ALPHA_SLICE_1.md`
- `docs/architecture/CORE_ALPHA_DATA_MODEL.md`
- `docs/security/CORE_ALPHA_SECURITY_MODEL.md`
- `docs/development/CORE_ALPHA_IMPLEMENTATION_PLAN.md`
- `docs/development/ARCHITECTURE_GATES.md`
- `docs/development/REQUIREMENTS_ARCHITECTURE_MAPPING.md`
- all accepted Task 008 ADRs
- `docs/development/REQUIREMENTS_TRACEABILITY_MATRIX.md`
- `docs/development/CORE_ALPHA_READINESS.md`
- `ai/audits/2026-07-20_QWSG_REPOSITORY_DEEP_AUDIT.md`
- `ai/audits/2026-07-20_QUANTUM_CREATOR_CONFORMANCE.md`
- `ai/history/008_2026-07-20_core-alpha-architecture.md`
- the complete active Task 009 prompt

If an authoritative document exists at a different verified path, use the verified path and record the variance.

## Starting State Verification

Before modifying any file:

1. Confirm the repository root.
2. Run:
   - `bin/job --check`
   - `bin/job --path`
   - `bin/job --history`
3. Confirm that exactly one active prompt exists and resolves to Task 009.
4. Confirm Task 008 is complete and all required architecture deliverables exist.
5. Record:
   - current UTC time;
   - branch and HEAD commit;
   - full Git status;
   - pre-existing modified and untracked files;
   - repository structure;
   - existing source-code state;
   - file ownership, permissions, and ACLs for affected paths;
   - available Go installation and version;
   - whether network access would be needed for module resolution;
   - existing build, lint, and test tools.
6. Validate the owner-ratified decisions:
   - Go runtime;
   - self-contained CLI binary;
   - exit codes 0, 1, and 2.
7. Check that no unresolved architecture gate blocks Slice 1 implementation.
8. Detect placeholders, lifecycle conflicts, stale prompts, broken required links, or permission limitations.
9. If the Go toolchain is missing, unsupported, or requires package installation, stop and report the exact blocker. Do not install anything without approval.
10. If module dependencies require network access, prefer the standard library and architecture-compatible local design. Do not silently add external dependencies.
11. If the active prompt changed after the current agent session started, stop and request a new session or explicit reload.
12. Preserve all unrelated pre-existing working-tree changes.

## Snapshot Requirements

Before implementation, create:

`ai/backups/<UTC_TIMESTAMP>_task009_core_alpha_slice_1_implementation/`

The snapshot must include at minimum:

- `START_STATE.md`
- `git-status-before.txt`
- `git-diff-before.patch`
- `permissions-before.txt`
- `affected-files.txt`
- `new-files.txt`
- `manifest.txt`
- `SHA256SUMS`
- `restore.sh`
- preserved copies of every existing file that may be modified

Requirements:

- use the repository-standard UTC timestamp format;
- distinguish pre-existing files from files created by Task 009;
- verify all snapshot checksums;
- statically validate the restore script;
- ensure rollback restores only Task 009 changes;
- preserve unrelated pre-existing worktree changes;
- avoid broad deletion, hard reset, or repository-wide cleanup;
- do not execute rollback during normal completion;
- disclose files that cannot be preserved because of permissions.

## Risk Assessment

Record and rate at minimum:

- host-mutation risk;
- command-execution risk;
- privilege risk;
- secret-exposure risk;
- privacy risk;
- data-correctness risk;
- partial-result ambiguity risk;
- Linux compatibility risk;
- runtime and dependency risk;
- output-schema compatibility risk;
- test-coverage risk;
- rollback risk;
- scope-expansion risk.

Expected host-stability risk must remain low.

Any implementation choice that increases privilege, mutation, network exposure, or dependency risk requires a stop and separate owner approval.

## Planned Work

### Phase 1 — Preflight and Implementation Baseline

- Validate Task 009 and Task 008 outputs.
- Create and verify the snapshot.
- Confirm the Go toolchain.
- Record the exact Go version selected for repository support.
- Establish or validate the source layout from the architecture.
- Confirm the standard-library-first dependency policy.
- Confirm ignored build-output paths.
- Confirm no implementation blocker remains.

### Phase 2 — Go Module and CLI Skeleton

Create the minimum project skeleton required by the architecture.

The CLI must support at least:

- `qwsg inventory`
- `qwsg version`
- `qwsg help` or equivalent standard help behavior

The inventory command must support machine-readable JSON output.

Do not add speculative commands.

The CLI must:

- reject unknown commands and unsupported options clearly;
- write structured inventory to standard output;
- write human-facing errors and diagnostics to standard error;
- use exit code `0`, `1`, or `2` according to the ratified rules;
- avoid interactive prompts;
- avoid requiring root;
- avoid network access;
- avoid host mutation.

### Phase 3 — Core Domain Model

Implement the Task 008 inventory model, including:

- schema or document version;
- tool version;
- collection timestamp;
- collection duration where approved;
- host or target identity with privacy-safe treatment;
- collection status;
- per-section status;
- provenance;
- issues;
- unknown and unavailable values;
- permission-denied state;
- unsupported state;
- timeout state;
- partial result state.

Do not silently omit failed sections.

Ensure output remains structurally valid when any optional collector fails.

### Phase 4 — Safe Execution Abstraction

Implement a bounded command-runner abstraction where commands are required.

Requirements:

- fixed executable and argument lists;
- no shell invocation unless Task 008 explicitly requires it;
- no string-concatenated shell commands;
- context-based timeout and cancellation;
- bounded captured output;
- controlled environment;
- explicit working directory where needed;
- deterministic error classification;
- no privilege escalation;
- testable fake runner;
- no user-controlled executable path;
- no unbounded subprocess tree.

Prefer direct reads from safe virtual filesystems or standard APIs when the architecture designates them as safer than command execution.

### Phase 5 — Approved Discovery Collectors

Implement only the collectors approved by `CORE_ALPHA_SLICE_1.md`.

The intended categories are:

1. operating-system identity and version;
2. kernel version;
3. machine architecture;
4. privacy-safe host identity;
5. CPU summary;
6. memory summary;
7. filesystem and storage summary;
8. network-interface summary with sensitive fields controlled;
9. running-service or service-manager summary only where safely accessible and architecture-approved;
10. installed runtime and server-component version summary only where safely detectable;
11. collector permission and capability report.

Every collector must:

- be read-only;
- have a defined timeout;
- return structured provenance;
- distinguish unsupported, unavailable, permission denied, timeout, parse error, and success;
- avoid secret collection;
- avoid arbitrary file reads;
- avoid mutation;
- have unit tests;
- degrade safely.

If Task 008 defines a narrower collector set, obey Task 008.

### Phase 6 — Aggregation and Exit Semantics

Implement deterministic aggregation:

- complete valid inventory → status `complete`, exit `0`;
- partial valid inventory → status `partial`, exit `2`;
- fatal failure before a valid inventory exists → status or error appropriate to the architecture, exit `1`.

A failed optional collector must not become fatal unless the architecture explicitly marks it mandatory.

A structurally invalid output must be fatal.

Document the exact aggregation rules in source comments and developer documentation.

### Phase 7 — Output and Redaction

Implement canonical JSON output.

Requirements:

- stable top-level structure;
- deterministic field naming;
- valid JSON for complete and partial results;
- no secrets;
- no private file contents;
- no raw environment dump;
- no unnecessary MAC addresses, public IP addresses, usernames, process arguments, or other sensitive values unless explicitly permitted and redacted by architecture;
- errors and issues must not leak secret values;
- human diagnostics must not corrupt standard-output JSON;
- optional pretty-printing may be added only if architecture-compatible and non-disruptive.

Create a sanitized sample output for documentation or tests.

### Phase 8 — Testing

Implement at minimum:

- unit tests for domain status and exit-code aggregation;
- unit tests for every collector;
- parser tests using fixtures;
- timeout tests;
- permission-denied tests;
- unsupported-platform or missing-file tests;
- partial-result tests;
- fatal-result tests;
- redaction tests;
- command-runner safety tests;
- JSON schema or structural-validation tests;
- CLI argument and exit-code tests;
- deterministic-output tests where applicable.

Tests must not require root, network, package installation, service changes, or host mutation.

Use fake runners and fixtures for risky or environment-dependent behavior.

### Phase 9 — Static Analysis, Build, and Local Read-only Run

Run all available safe checks, including as applicable:

- `gofmt` verification;
- `go vet`;
- `go test ./...`;
- race tests if available and practical;
- build of the `qwsg` binary;
- repository checks;
- `git diff --check`.

After all automated tests pass, execute the built binary locally in read-only mode.

Before preserving any sample output:

- inspect it for secrets and sensitive host data;
- sanitize or discard unsafe output;
- do not commit real host inventory.

A local run may validate behavior, but environment-specific collector gaps must be reported truthfully.

### Phase 10 — Documentation and Handoff

Document:

- supported Go toolchain;
- source layout;
- build command;
- test command;
- CLI usage;
- exit codes;
- JSON output contract;
- complete, partial, and fatal behavior;
- collector limitations;
- security and privacy guarantees;
- unsupported features;
- known implementation limitations;
- open architecture gates;
- exact next hardening work.

Recommend Task 010 for verification and hardening, but do not create or activate it.

### Phase 11 — Lifecycle Completion

- update required history and engineering records;
- verify all generated artifacts;
- verify snapshot and restore integrity;
- preserve file permissions;
- mark Task 009 complete only after every completion criterion passes;
- do not commit or push.

## Rollback Plan

The Task 009 restore script must:

1. restore all pre-existing files modified by Task 009;
2. remove only files newly created by Task 009;
3. remove Task 009 build outputs only from explicitly bounded repository output paths;
4. preserve unrelated pre-existing working-tree changes;
5. preserve later unrelated files if rollback is executed after task completion, unless the script can prove they belong to Task 009;
6. avoid `git reset --hard`, broad `git clean`, recursive deletion outside bounded paths, or host-level cleanup;
7. verify restored checksums where possible;
8. report permission failures without unsafe escalation;
9. never modify host services, packages, accounts, or configuration.

The final delivery must provide the exact restore command.

## Deliverables

At minimum create:

1. Go module files required by the approved implementation.
2. `cmd/qwsg/` CLI entry point or the exact Task 008-approved equivalent.
3. Internal domain, collector, runner, aggregation, and output packages according to the approved architecture.
4. Automated unit and bounded integration tests.
5. Test fixtures and safe fake execution components.
6. A local build procedure producing the `qwsg` binary in an ignored output directory.
7. A sanitized example inventory or test fixture.
8. Developer documentation for building, testing, and running Slice 1.
9. CLI and output-contract documentation.
10. `ai/history/009_2026-07-20_core-alpha-slice-1-implementation.md`.

Update as necessary:

- `README.md`
- `CHANGELOG.md`
- `.gitignore`
- `VERSION` only if release policy and architecture explicitly require it for this milestone
- `ai/core/05_SYSTEM_MAP.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/13_ROADMAP.md`
- `ai/projects/QWSG.md`
- `ai/projects/QWSG/QWSG_MASTER_PLAN.md`
- `docs/development/CORE_ALPHA_IMPLEMENTATION_PLAN.md`
- `docs/development/REQUIREMENTS_ARCHITECTURE_MAPPING.md`
- the active Task 009 prompt

Do not create duplicate authoritative documents.

## Verification

At minimum verify:

- `bin/job --check` succeeds before and after implementation;
- the active prompt resolves to Task 009;
- Task 008 is complete;
- snapshot checksums pass;
- restore script syntax passes;
- the selected Go version is documented;
- `gofmt` produces no pending changes;
- `go vet ./...` succeeds;
- `go test ./...` succeeds;
- every collector has automated tests;
- CLI tests verify exit codes `0`, `1`, and `2`;
- complete output is valid JSON;
- partial output is valid JSON;
- fatal behavior does not emit misleading valid inventory;
- timeout, unavailable, unsupported, permission-denied, and parse-error states are tested;
- no test requires root;
- no test requires network;
- no test mutates host state;
- no secret or sensitive fixture is committed;
- command execution does not use untrusted shell construction;
- no network listener exists;
- no background service exists;
- no installer or package-manager integration exists;
- no privilege-escalation mechanism exists;
- no production dependency was added without justification and documentation;
- the binary builds successfully;
- a local read-only inventory run succeeds or truthfully reports a valid partial result;
- any preserved sample output is sanitized;
- internal links resolve;
- FR and AC mappings remain valid;
- no unresolved gate is represented as resolved;
- no placeholders remain;
- `git diff --check` succeeds;
- unrelated pre-existing changes are untouched;
- HEAD remains unchanged;
- no commit or push occurred.

If the repository contains additional mandated checks, run them too.

## Documentation Updates

Update documentation only as needed to establish:

- the first working implementation;
- the supported development toolchain;
- CLI and exit-code behavior;
- output structure;
- collector boundaries;
- security guarantees;
- known limitations;
- the distinction between Core Alpha implementation and supported release;
- the handoff to Task 010.

Do not claim production readiness, supported release status, broad Linux compatibility, or public distribution readiness without evidence.

## Completion Criteria

Task 009 is complete only when:

1. the `qwsg` CLI builds successfully;
2. the approved Slice 1 collectors are implemented;
3. the application runs locally without root and without host mutation;
4. complete, partial, and fatal outcomes follow the ratified exit-code rules;
5. complete and partial outputs are structurally valid JSON;
6. failed collectors are represented truthfully;
7. security, timeout, output-size, redaction, and command-boundary requirements are enforced;
8. automated tests cover the approved collectors and critical failure modes;
9. all required checks pass;
10. the local read-only run is successful or produces a valid partial result;
11. no network listener, Console, installer, daemon, service, privilege escalation, remediation, or host mutation was added;
12. documentation and lifecycle records are complete;
13. snapshot and bounded rollback are valid;
14. unrelated pre-existing worktree changes remain preserved;
15. no commit or push occurred;
16. the final report clearly states whether Task 010 may safely begin and lists all known limitations and blockers.

If any criterion fails, leave the task incomplete and report the exact blocker.

## Final Delivery Report

Report in Hungarian:

- starting state;
- selected and pinned Go toolchain;
- source structure;
- files created and modified;
- implemented collectors;
- CLI commands and behavior;
- exit-code behavior;
- output contract;
- security and privacy controls;
- tests and exact results;
- build result;
- local read-only execution result;
- sanitization result;
- known limitations;
- open architecture and release gates;
- permission limitations;
- preserved pre-existing changes;
- exact rollback command;
- whether Task 010 is recommended;
- confirmation that no commit or push occurred.

## Owner Approval Requirements

This prompt is explicitly approved by Attila, Project Owner, for Task 009.

The task may begin within the exact scope above.

Any scope expansion, package installation, external dependency requiring network access, host modification, privilege escalation, service creation, network listener, installer, persistent storage, secret access, destructive action, commit, or push requires separate explicit owner approval.
