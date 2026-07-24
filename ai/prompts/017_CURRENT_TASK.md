# Current Engineering Task 017: QWSG User CLI & Snapshot Explorer

## Task Metadata

- Task ID: `017`
- Task slug: `qwsg-user-cli-and-snapshot-explorer`
- Status: `approved`
- Date opened: `2026-07-24` UTC
- Human authority: Attila — Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

QWSG User CLI & Snapshot Explorer


## Objective


Transform the current engineering-oriented one-shot CLI into the first production-ready, user-oriented QWSG application that a Linux administrator can build, install, discover, and use without repository-specific or Go-tooling knowledge.

Deliver a coherent command hierarchy, complete contextual help, human-readable Inventory and snapshot presentation, explicit JSON output where machine-readable compatibility is required, Snapshot Explorer commands, build/version information, a reversible installation workflow, and complete user and demonstration documentation.

The implementation shall expose and refine only capabilities already provided by Inventory 1.0, the Collector Framework, Canonical System Inventory v1, and the Digital Twin Foundation. It shall not add monitoring or intelligence.

Project Owner acceptance requires a clean Ubuntu 24.04 environment to build and install QWSG and then execute the documented `qwsg` command sequence successfully. Every user-visible result shall be meaningful, deterministic where the underlying data is deterministic, safe for terminal display, and clear about complete, partial, failed, missing, corrupt, or incompatible data.

Existing machine consumers shall retain an explicit documented way to obtain the established Inventory 1.0 JSON envelope with additive `canonical_inventory` and the established status exit codes.


## Scope


- Inspect the current `cmd/qwsg`, application, inventory, Inventory Store, build, test, and documentation boundaries before design or implementation.
- Define and document a stable user CLI information architecture for:

  - `qwsg help`;
  - `qwsg version`;
  - `qwsg inventory`;
  - `qwsg inventory save`;
  - `qwsg inventory list`;
  - `qwsg inventory info`;
  - `qwsg inventory load`.
- Provide command-specific help for the root command and every inventory subcommand, including supported options, defaults, output modes, exit semantics, examples, privacy boundaries, and failure behavior.
- Make no-argument and invalid-argument behavior deliberate, consistent, tested, and documented.
- Introduce human-readable terminal output for interactive administrator use.
- Retain explicit JSON output for machine use and backward compatibility; define the exact compatibility path and test it before changing any current default.
- Present Inventory summaries without changing Inventory 1.0, canonical inventory, collector statuses, issues, redactions, or exit-code truthfulness.
- Implement Snapshot Explorer listing over the existing file-backed Inventory Store.
- Implement Snapshot Explorer metadata inspection over existing validated snapshot envelopes.
- Allow explicit loading of the latest or a named snapshot through the documented CLI.
- Reuse Inventory Store validation for stored content; do not duplicate or bypass its integrity, compatibility, permission, path, file-type, retention, privacy, or corruption checks.
- Design list and info output so stored observations are never presented as current live state, desired state, comparison results, health verdicts, or monitoring data.
- Show useful snapshot metadata available from the existing persistence envelope, including snapshot name or selector, observation/completion time, snapshot identity in its privacy-safe documented form, schema/format identity, status, integrity-validation result, and other already persisted non-secret metadata that materially aids browsing.
- Preserve non-root, read-only collection and explicit one-shot persistence.
- Define consistent stdout/stderr separation and stable exit-code behavior for success, partial Inventory, usage error, missing snapshot, invalid/corrupt store, permission failure, and runtime failure.
- Ensure terminal rendering escapes or safely represents untrusted stored or observed strings and cannot emit uncontrolled terminal control sequences.
- Add version/build information suitable for installed binaries, including the QWSG version and reproducible build metadata that can be supplied by the build without requiring a Git checkout at runtime.
- Keep development builds deterministic and useful when optional build metadata is unavailable.
- Extend the Makefile with documented build and install targets.
- Support a configurable installation prefix and destination directory suitable for normal installation and isolated packaging/tests.
- Install the `qwsg` binary with an appropriate executable mode and without creating services, users, groups, state directories, configuration, scheduled jobs, or network listeners.
- Define safe install defaults, preconditions, exact installed paths, overwrite behavior, verification, and bounded uninstall or manual removal guidance.
- Verify the Project Owner's exact acceptance flow, including `make build`, `make install`, `qwsg version`, `qwsg help`, and every required inventory command.
- Where system installation requires privilege, use `sudo` only for the documented installation step, after verifying the exact target path and rollback; QWSG runtime commands shall not require root.
- Add unit, integration, CLI golden/fixture, install-layout, compatibility, and regression tests using synthetic privacy-reviewed data and isolated temporary stores.
- Use only the Go standard library and existing repository tooling.
- Update canonical architecture, development, installation, administration/user, demonstration, README, history, and delivery documentation necessary to make the CLI an official supported user boundary.
- Provide English engineering and user documentation and a functionally equivalent Hungarian user guide for the owner-facing product workflow.
- Preserve Engineering Framework 1.0, Inventory 1.0, Collector Framework, Canonical System Inventory v1, and Digital Twin Foundation compatibility.
- Record the complete implementation, validation evidence, rollback procedure, known limitations, Git delivery, and Project Owner acceptance in the Task 017 history and delivery report.


## Out of Scope


- Modifying collector implementations, collector registration, collector contracts, collector timeouts, collector budgets, or collector evidence acquisition.
- Changing Inventory 1.0 or Canonical System Inventory schemas, semantics, status aggregation, privacy/redaction rules, or compatibility guarantees.
- Changing the Digital Twin persisted format, integrity model, atomicity model, retention semantics, locking, permissions, or storage safety rules unless a CLI-only defect cannot be corrected without a separately reviewed scope change.
- Automatically creating, choosing, discovering, migrating, repairing, pruning, or deleting an operator's Inventory Store.
- Comparing snapshots or implementing diffs, drift analysis, trends, timelines, baselines, desired state, policy evaluation, health scoring, recommendations, or intelligence.
- Continuous monitoring, polling, watching, background refresh, scheduling, timers, daemon mode, systemd units, cron jobs, services, or long-running processes.
- Alerts, notifications, e-mail, webhooks, reports derived from comparison, or remediation.
- REST, HTTP, RPC, socket, network API, Web UI, Console, dashboard, or remote access.
- Database, external storage engine, cloud upload, remote synchronization, or telemetry.
- External Go modules, third-party CLI frameworks, terminal libraries, package managers, or runtime dependencies.
- Debian, RPM, Snap, Flatpak, Homebrew, container, or other distribution packaging.
- Installer scripts that execute downloaded content, modify shell profiles, manage users/groups, elevate the QWSG runtime, or alter host services.
- Automatic `sudo`, privilege escalation inside QWSG, setuid/setgid binaries, Linux capabilities, or root-only collection.
- Localization infrastructure beyond the English and Hungarian documentation and a CLI design that remains localization-ready.
- Guaranteeing unsupported platforms; the acceptance platform for this task is Ubuntu 24.04, while broader Linux support must be described only to the extent verified.
- Rewriting the Inventory Store or CLI with an unrelated architecture.
- Modifying the Engineering Framework, Task Builder, lifecycle tooling, Git policy, historical task records, or prior delivery evidence except for narrowly required current indexes.
- Weakening snapshot, rollback, approval, validation, targeted-staging, history, privacy, or security requirements.
- Deleting or modifying pre-existing untracked files or backup directories.
- Creating Task 018 or beginning monitoring, intelligence, comparison, daemon, API, Console, or packaging milestones.


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
- Verify `ai/framework/VERSION` and `ai/config/engineering-project.conf` declare framework version `1.0.0`.
- Verify branch `main`, canonical HTTPS remote `origin`, exact HEAD, upstream relationship, tags, complete Git status, and ahead/behind state before modifying any target.
- Verify Task 016 is complete and archived, Task 017 is the sole approved active prompt with exactly one matching history record, and no Task 018 exists.
- Run `bin/job --check`, `ai/scripts/next-task.sh --check`, and `bin/job --check-test-tasks`.
- Record and preserve every pre-existing untracked path; do not stage, delete, move, clean, archive, or modify it.
- Verify the current project version, Go version, Ubuntu/kernel environment, Make implementation, install tools, user identity, groups, umask, relevant ACLs, and write permissions.
- Inspect the current Makefile targets and confirm the existing development binary path and version source.
- Inspect the complete current CLI grammar, help/version behavior, stdout/stderr behavior, and exit codes.
- Verify the current root help advertises `inventory`, `version`, and `help`.
- Verify `qwsg inventory` currently emits the Inventory 1.0 JSON envelope with additive `canonical_inventory`, runs one-shot, performs no persistence by default, and returns `0` for complete, `2` for partial but usable, and `1` for failed.
- Verify the existing explicit persistence commands and their exact current syntax:

  - `qwsg inventory save --store DIR [--retention N]`;
  - `qwsg inventory load --store DIR [--retention N] [--snapshot NAME]`.
- Verify `inventory list` and `inventory info` are not yet user commands.
- Verify `internal/inventorystore` already owns validated save, deterministic list, exact/latest load, integrity checks, version checks, permission/path checks, locking, and retention.
- Verify collectors do not import or invoke persistence.
- Verify the repository currently uses no external Go dependency and record `go.mod` and `go.sum` state.
- Verify current English and Hungarian canonical inventory guidance, persistence architecture, README claims, and missing installation/demo documentation.
- Run the complete baseline Go, race where supported, vet, format, Engineering Framework, lifecycle, diversion, Task Builder, build, CLI, Inventory, persistence, privacy, and Git checks.
- Capture the exact baseline behavior for every command that Task 017 may change, including representative stdout, stderr, and exit status, using sanitized output or synthetic fixtures.
- Stop before implementation on any unexplained failure, lifecycle mismatch, unexpected active task, repository identity difference, unsafe install target, dependency change, or material difference from this baseline.


## Snapshot Requirements


- Before modifying any repository or installation target, create a unique UTC-dated rollback snapshot outside the repository under the configured `/tmp` snapshot location.
- Include every intended existing source, test, Makefile, documentation, task history, audit, and index target.
- Record verified absence for every proposed new repository path and every proposed installed path.
- Capture branch, HEAD, tags, remotes, upstream and ahead/behind state, complete Git status, relevant untracked paths, user/groups, umask, ACLs, file ownership, modes, and line-ending/UTF-8 state.
- Capture current CLI command behavior, help/version text, representative sanitized output, exit codes, build metadata, binary hash, and Make target behavior.
- Capture the exact pre-install state of the intended prefix and binary path, including absence or a byte-for-byte backup of any existing `qwsg` binary, ownership, mode, and symlink status.
- Capture hashes and modes of `go.mod`, `go.sum` when present, Makefile, CLI sources/tests, application boundaries, Inventory Store sources/tests, affected documents, Task 017 prompt/history, and configured validators.
- Include baseline validation output, a manifest, SHA-256 checksums, a retention decision, and exact bounded restore instructions.
- Verify every copied file and recorded-absence target before implementation.
- Keep repository rollback and system-install rollback separately bounded and documented.
- Do not capture credentials, environment secrets, raw live Inventory output, host identifiers, operator stores, unrelated backup payloads, or build caches.
- Retain the verified snapshot through Project Owner acceptance.
- Do not begin code, documentation, Makefile, or installation-target modification until snapshot integrity and rollback instructions pass review.


## Risk Assessment


- CLI compatibility risk is high because scripts may depend on current JSON output and exit codes. Preserve an explicit JSON path, test exact schemas and status codes, and document any default-output decision.
- Usability and truthfulness risk is high because a friendly summary can hide partial collection, redactions, stale observations, or corrupt stores. Always expose status and issues clearly and never translate absence into health.
- Terminal-injection risk is medium to high because observed or stored strings may contain control bytes. Sanitize all human-readable rendering and test adversarial fixture values.
- Persistence safety risk is high because browsing could accidentally bypass validation or mutate operator state. Reuse the Inventory Store read APIs and keep list/info/load read-only except for explicit save.
- Privacy risk is high because snapshots contain host evidence. Display only already approved privacy-safe fields by default, avoid raw host identifiers in fixtures/reports, and preserve restrictive store permissions.
- Install risk is high because a privileged install can overwrite a system binary. Verify exact targets, reject unsafe file types, use configurable `PREFIX`/`DESTDIR`, snapshot any pre-existing binary, and provide exact restoration.
- Privilege risk is high if build or runtime begins to require root. Build and run as an ordinary user; constrain `sudo` to a reviewed install command only.
- Build metadata risk is medium because non-reproducible timestamps or repository assumptions can make binaries inconsistent. Define injected fields and deterministic fallbacks; do not require `.git` at runtime.
- Help/output drift risk is medium because documentation and command behavior can diverge. Generate or test stable examples and verify every documented command.
- Store compatibility risk is high if list/info interpret persisted envelopes differently from load. Use a single validated parsing boundary and adversarial compatibility fixtures.
- Error-contract risk is medium because ambiguous exit codes can mislead automation. Define and test usage, operational, partial, and failure semantics without relabeling Inventory status.
- Localization risk is medium because terminal strings may become an accidental permanent English-only API. Keep rendering boundaries separable and avoid embedding user-visible strings throughout domain logic.
- Scope-expansion risk is high because an explorer suggests comparison and monitoring. Enforce explicit negative tests and import/process reviews for all prohibited capabilities.
- External-dependency risk is low but mandatory: reject `go.mod`/`go.sum` dependency additions and third-party CLI frameworks.
- Rollback risk is medium because both repository files and an installed binary may change. Prove both restore paths independently before the acceptance installation.


## Planned Work


### Phase 1 — Baseline, snapshot, and CLI contract

- Complete starting-state verification and create the verified rollback snapshot.
- Inventory current commands, output, error, exit, build, install, documentation, Inventory, and Inventory Store contracts.
- Define the user CLI grammar, global/subcommand option rules, output-mode rules, stdout/stderr boundary, exit-code table, and compatibility policy before implementation.
- Decide and document the safest default for `qwsg inventory` based on the existing compatibility promise; require an explicit tested JSON mode if human-readable output becomes the default.
- Define Snapshot Explorer list/info fields using only existing validated store metadata and privacy-safe Inventory summary data.

### Phase 2 — Command architecture and help

- Refactor the CLI into focused, testable command parsing and rendering boundaries without introducing a third-party framework.
- Implement root and contextual help, consistent usage errors, option validation, examples, and discoverable command grouping.
- Ensure `qwsg help`, `qwsg help inventory`, and subcommand help cover all supported operations.
- Preserve one-shot execution and avoid global mutable runtime state.

### Phase 3 — Human-readable Inventory presentation

- Add a deterministic, safe terminal renderer for live and loaded Inventory summaries.
- Clearly show observation time, status, partial/failure conditions, redactions, relevant issues, and canonical categories without implying health.
- Preserve full JSON output through the approved compatibility path and verify byte-semantic schema compatibility.
- Keep output formatting separate from Inventory domain structures and collection.

### Phase 4 — Snapshot Explorer

- Implement `inventory list` using the existing deterministic Inventory Store listing boundary.
- Implement `inventory info` for an explicitly selected snapshot, with a clearly documented latest/default rule only if ambiguity is eliminated.
- Implement or refine `inventory load` for latest and named snapshots with human-readable and JSON output modes.
- Make list, info, and load fail closed on unsafe, corrupt, unsupported, integrity-invalid, or permission-invalid stores.
- Do not mutate store contents during list, info, or load.

### Phase 5 — Version, build, and installation

- Add user-meaningful `qwsg version` output with version and controlled build information.
- Add or refine Makefile `build` and `install` targets with explicit `PREFIX`, `DESTDIR`, binary path, permissions, overwrite rules, and help.
- Test install layout without privilege in isolated temporary destinations.
- Perform the final development-server installation demonstration only after exact target preflight and system-install rollback are verified.
- Confirm the installed binary does not depend on the source checkout or manual Go invocation.

### Phase 6 — Tests and security hardening

- Add table-driven parser and help tests for every command, option combination, invalid input, and exit path.
- Add golden or structured assertions for human-readable and JSON output.
- Add synthetic Snapshot Explorer tests for empty, one, many, latest, named, partial, corrupt, incompatible, hash-mismatch, permission-denied, symlink, traversal, and adversarial terminal-string cases.
- Add isolated install tests for default/custom prefix, `DESTDIR`, modes, existing target, missing directory, and rollback.
- Run compatibility, privacy, race, format, vet, Engineering Framework, lifecycle, diversion, builder, and Git policy gates.

### Phase 7 — Documentation, acceptance demonstration, and delivery

- Update canonical architecture and developer contracts for the CLI boundary without changing Inventory or persistence semantics.
- Add installation instructions, an English user guide, a functionally equivalent Hungarian user guide, and a first demonstration walkthrough.
- Run the complete clean Ubuntu 24.04 acceptance flow and the development-server command demonstration exactly as documented.
- Record exact environment, commands, outputs in sanitized form, exit codes, changed paths, validation evidence, known limitations, and rollback.
- Stage and deliver only explicit Task 017 paths; do not create Task 018.


## Rollback Plan


- Stop immediately on lifecycle drift, compatibility regression, collector change, Inventory schema change, store mutation during read operations, terminal-injection failure, privacy exposure, unsafe install behavior, unexplained validation failure, or prohibited capability/dependency introduction.
- Before restoration, capture the exact failure, affected command, sanitized stdout/stderr, exit status, Git status, target modes/hashes, and isolated fixture state without recording secrets or live host Inventory.
- Restore only Task 017-modified repository files and modes from the verified snapshot using explicit paths.
- Remove only Task 017-created repository paths whose pre-task absence was recorded and whose identity/hash still matches the task artifact.
- For an installed binary, restore the exact pre-install file, ownership, mode, and path from the separate installation snapshot; if the target was absent, remove only the exact installed binary after verifying its Task 017 hash.
- Never remove an operator Inventory Store or snapshots as rollback.
- Remove only isolated Task 017 test stores and install staging roots created under verified temporary paths.
- Rerun baseline build, Go, race where supported, vet, format, CLI compatibility, Inventory, store, privacy, Engineering Framework, lifecycle, diversion, builder, install-layout, and Git checks after restoration.
- Confirm Task 017 lifecycle records remain truthful and all unrelated untracked paths remain untouched.
- Never use broad Git reset, restore, checkout, clean, wildcard deletion, unbounded recursive removal, force-push, history rewriting, or blind system-path overwrite.
- If safe restoration of a pre-existing installed binary or lifecycle state cannot be proven, stop and request Project Owner direction instead of guessing.


## Deliverables


- A production-ready, user-oriented `qwsg` command hierarchy.
- Complete root and contextual CLI help.
- Human-readable Inventory presentation with explicit truthful status and issue reporting.
- An explicit documented JSON compatibility mode preserving Inventory 1.0 plus `canonical_inventory`.
- Snapshot Explorer commands for deterministic listing, metadata inspection, and validated loading of latest or named snapshots.
- Consistent stdout/stderr and documented exit-code behavior.
- Version and build information suitable for installed binaries.
- Makefile build and install workflow with configurable prefix and isolated destination support.
- A verified user-installable `qwsg` binary for the Ubuntu 24.04 acceptance environment.
- Unit, integration, golden/fixture, adversarial, security, compatibility, install-layout, and regression tests.
- Installation documentation.
- An English user guide and functionally equivalent Hungarian user guide.
- A first demonstration walkthrough containing the exact Project Owner acceptance commands.
- Updated canonical architecture/developer documentation, README and indexes as required.
- Task 017 history, verified rollback snapshot, delivery audit/report, validation evidence, Git evidence, and Project Owner acceptance record.


## Verification


- Run `ai/scripts/framework-check.sh --run-validations`.
- Run `make test`, `make vet`, `make fmt-check`, and `make engineering-test`.
- Run `bin/job --check`, `ai/scripts/next-task.sh --check`, and `bin/job --check-test-tasks`.
- Run `go test -race ./...` when supported and report any environmental limitation truthfully.
- Run all new CLI parser, help, renderer, Snapshot Explorer, store-fixture, compatibility, security, and installation tests.
- Verify `go.mod` and `go.sum` add no external dependency and all imports use the standard library or existing project packages.
- Verify collectors are byte-for-byte unchanged and do not import CLI, renderer, install, or persistence responsibilities.
- Verify Inventory 1.0 and canonical schema constants, validation, status aggregation, privacy/redaction, and collector behavior are unchanged.
- Verify the documented JSON compatibility command preserves the complete Inventory 1.0 envelope, additive `canonical_inventory`, deterministic ordering where guaranteed, and `0`/`2`/`1` Inventory status exits.
- Verify human-readable live and stored Inventory output identifies complete, partial, and failed states truthfully and does not claim health, comparison, recency beyond recorded timestamps, or desired state.
- Verify output safely handles ANSI escape sequences, control bytes, invalid display data reachable through fixtures, long values, Unicode, empty values, and terminal width without hiding status.
- Verify `qwsg help`, `qwsg version`, `qwsg inventory`, `qwsg inventory save`, `qwsg inventory list`, `qwsg inventory info`, and `qwsg inventory load` are discoverable and match documentation.
- Verify every command-specific help path and invalid argument combination has the documented stdout/stderr and exit behavior.
- Verify list ordering is deterministic and list/info/load all use existing store validation.
- Verify list/info/load reject corrupt, truncated, duplicate-key, integrity-invalid, unsupported-version, unsafe-permission, symlink, traversal, conflicting, and unexpected store content without modification.
- Verify Snapshot Explorer does not delete, rewrite, repair, migrate, prune, compare, or automatically choose an undeclared store.
- Verify `make build` produces the documented binary without requiring repository-specific knowledge beyond the checked-out source.
- Verify `make install` supports isolated `DESTDIR` and configurable `PREFIX`, installs only documented paths with correct modes, and does not create a service, user, group, state directory, configuration, scheduler, or network listener.
- Verify the installed binary runs without the source tree and reports controlled version/build data.
- On a clean Ubuntu 24.04 acceptance environment, execute and record:

  - `make build`;
  - `make install`;
  - `qwsg version`;
  - `qwsg help`;
  - `qwsg inventory`;
  - `qwsg inventory save`;
  - `qwsg inventory list`;
  - `qwsg inventory info`;
  - `qwsg inventory load`.
- On the development server, execute the same documented application sequence after exact installation-target preflight; use `sudo make install` only when required by the selected prefix and explicitly record the bounded restoration.
- Verify a normal Linux administrator can follow the documentation without editing source code, directly invoking Go tooling, or requiring root to run QWSG.
- Verify build, install, runtime, help, and snapshot demonstrations produce meaningful human-readable output and the documented JSON path remains usable.
- Verify no daemon, scheduler, background loop, monitoring, comparison, drift, health, alert, notification, REST API, Web UI, database, external dependency, telemetry, or network listener was added.
- Run shell syntax, UTF-8, LF, file-mode, ACL where relevant, secret/private-host-data, generated-artifact, documentation-link, and `git diff --check` reviews.
- Prove repository and installed-binary rollback using isolated or non-destructive verification before acceptance.
- Compare final Git status with the recorded baseline and stage only explicit reviewed Task 017 paths.
- Perform the canonical targeted-staging, staged-diff, commit, HTTPS dry-run push, push, and post-push synchronization checks required by `ai/core/16_GIT_POLICY.md`.
- Verify Task 017 remains the only active task during implementation and Task 018 does not exist.


## Documentation Updates


- Update `README.md` so current product status, build/install entry points, CLI commands, JSON compatibility, Snapshot Explorer, and limitations are accurate.
- Add canonical installation documentation covering Ubuntu 24.04 prerequisites, build, `PREFIX`, `DESTDIR`, system installation, permissions, verification, upgrade/overwrite behavior, rollback, and removal.
- Add an English user CLI and Snapshot Explorer guide.
- Add a functionally equivalent Hungarian user guide.
- Add a first demonstration walkthrough containing the exact clean-build, install, help, version, live Inventory, save, list, info, and load sequence with expected result categories and exit semantics.
- Update `docs/architecture/INVENTORY_PERSISTENCE_AND_DIGITAL_TWIN.md` only to document the user CLI consumer boundary; do not change the store format or safety contract.
- Update `docs/development/CANONICAL_SYSTEM_INVENTORY.md` with CLI architecture, output modes, renderer boundaries, compatibility requirements, tests, and troubleshooting.
- Update canonical user Inventory guides or replace their command sections with authoritative links while preserving both English and Hungarian coverage.
- Update `ai/core/02_PROJECT_STRUCTURE.md`, `ai/core/04_ARCHITECTURE.md`, `ai/core/05_SYSTEM_MAP.md`, `ai/core/07_ENGINEERING_HISTORY.md`, or other canonical indexes only where necessary to reflect the supported CLI and documentation boundaries.
- Record every changed path, design decision, command/output contract, install target, compatibility result, validation result, known limitation, rollback procedure, Git record, and Project Owner acceptance in the Task 017 history.
- Create an English Task 017 Engineering Delivery Report under the canonical audit/report location.
- Keep monitoring, comparison, drift, health, daemon, alerts, API, Web UI, database, packaging, and external dependencies explicitly labeled as not delivered.


## Completion Criteria


- The repository contains one coherent supported `qwsg` CLI with discoverable root and contextual help.
- `qwsg version`, `qwsg help`, `qwsg inventory`, `qwsg inventory save`, `qwsg inventory list`, `qwsg inventory info`, and `qwsg inventory load` work exactly as documented.
- Human-readable output is meaningful, terminal-safe, privacy-aware, and truthful about complete, partial, failed, missing, corrupt, incompatible, and stored observations.
- An explicit documented JSON mode preserves the complete Inventory 1.0 envelope with additive `canonical_inventory` and established Inventory status exit codes.
- Snapshot Explorer lists deterministically, shows useful validated metadata, and loads latest or named snapshots without mutating, repairing, comparing, or reinterpreting store data.
- Existing Inventory Store integrity, atomicity, version, permission, path, locking, retention, and privacy safeguards remain enforced through one canonical validation boundary.
- Inventory 1.0, Canonical System Inventory v1, Collector Framework, Digital Twin Foundation, and Engineering Framework 1.0 remain compatible.
- Collector files and behavior are unchanged.
- The build embeds controlled version/build information and the installed binary needs neither the repository nor manual Go invocation at runtime.
- `make build` and `make install` are documented, tested, prefix-aware, destination-aware, reversible, and install only the expected binary path with the expected mode.
- A clean Ubuntu 24.04 system can build, install, and use QWSG by following the documentation.
- The complete Project Owner acceptance command sequence is demonstrated successfully on the development server with meaningful output and recorded exit statuses.
- QWSG runtime commands require no root; any installation privilege is confined to the reviewed install step.
- No daemon, scheduler, monitoring, comparison, drift analysis, health engine, alerts, notifications, REST API, Web UI, database, network service, telemetry, external dependency, or collector modification is introduced.
- All new and existing mandatory tests and validators pass without unexplained waiver.
- English and Hungarian user documentation, installation guidance, demonstration, README, architecture, and command help are consistent.
- Snapshot, rollback, history, delivery, security, compatibility, installation, Git, and validation evidence are complete and truthful.
- Only Task 017-scoped paths are staged and delivered; pre-existing untracked content remains untouched.
- Task 017 is the sole active production task and Task 018 does not exist.
- The Project Owner has reviewed the final delivery report and explicitly accepted the implementation before Task 017 is marked complete or archived.


## Owner Approval Requirements

Approved by Attila — Project Owner through the Engineering Task Builder on 2026-07-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
