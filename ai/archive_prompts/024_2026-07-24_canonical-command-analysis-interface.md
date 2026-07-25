# Current Engineering Task 024: Canonical Command and Analysis Interface

## Task Metadata

- Task ID: `024`
- Task slug: `canonical-command-analysis-interface`
- Status: `complete`
- Date opened: `2026-07-24` UTC
- Human authority: Attila — Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Canonical Command and Analysis Interface


## Objective

Establish and implement the permanent canonical command architecture and analysis interface of Quantum Wizard Server Guardian.

This task marks the transition from isolated internal engineering engines to one stable, deterministic, presentation-independent public interaction layer. Beginners must be able to execute complete workflows through memorable simple commands, while experienced Linux administrators must be able to compose deterministic workflows through a structured advanced command language.

Both usage models must resolve to the same canonical command definitions, execution model, analysis pipeline, and engineering engines. No interface-specific duplication of engineering logic may exist. The result must become the permanent foundation consumed by future CLI, Interactive Terminal, Dashboard, and REST API interfaces.


## Scope

- Design and implement one canonical, presentation-independent command architecture shared by all current and future QWSG interfaces.
- Establish two equally important usage models:
  - a Simple Command Model with memorable one- or two-word commands and deterministic predefined command profiles, including concepts equivalent to `qwsg`, `qwsg status`, `qwsg check`, `qwsg changes`, `qwsg health`, and `qwsg report`;
  - an Advanced Command Model with a deterministic, composable, versionable, and extensible command language.
- Support advanced command concepts for source selection, snapshot selection, pipeline selection, engine inclusion and exclusion, filtering, grouping, sorting, output formatting, and presentation selection.
- Define and implement the Canonical Analysis Pipeline:
  Inventory → Snapshot → Compare → Drift → Health → Rule → Report.
- Execute only the stages required by the selected command profile while preserving explicit deterministic dependencies and stage order.
- Define the Command Profile model, Command Definition 1.0, Command Execution model, Pipeline Orchestration model, parameter model, command taxonomy, and versioned public contracts.
- Provide deterministic human-readable terminal presentation derived from canonical command execution results without duplicating analysis logic.
- Integrate the existing canonical Inventory, Snapshot, Compare, Drift, Health, Rule, and Report engines through orchestration rather than reimplementing their responsibilities.
- Add architecture documentation, engineering tests, lifecycle records, and only the required canonical repository documentation updates.
- Preserve compatibility with existing public contracts and established QWSG behavior unless an explicitly documented, tested compatibility transition is required.


## Out of Scope

Task 024 must not implement:

- an Interactive Terminal UI or other TUI;
- a Web Dashboard;
- a REST API;
- a monitoring daemon;
- a scheduler;
- a Policy Engine;
- an Alert Engine;
- an Automation Engine;
- AI integration or machine learning;
- remote execution;
- remediation or host mutation;
- duplicated analysis or engineering logic for individual interfaces.

Future interface integration may be specified through stable contracts and documented boundaries only. No excluded future component, deployment, release, infrastructure change, dependency installation, commit, push, or unrelated repository cleanup is authorized.


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

Before implementation:

1. Run `bin/job --check` and verify Task 024 is the only active prompt, has the expected ID and slug, is approved, has exactly one matching history record, and contains no unresolved human-editing fields.
2. Read every Required Reading document, this prompt, Task 023 history, the Canonical Report architecture, and the existing Inventory, Snapshot, Compare, Drift, Health, Rule, Report, CLI, public-contract, architecture, system-map, and roadmap sources relevant to orchestration.
3. Run `ai/scripts/framework-check.sh` and validate project identity, framework version, repository root, configured canonical remote and branch, lifecycle configuration, and configured validations without sourcing configuration.
4. Record UTC time, effective user, working directory, branch, HEAD, remotes, ahead/behind relationship, complete Git status, relevant tags, Go version, file state, modes, ownership, ACLs, and dependency state.
5. Confirm Task 023 is complete and its uncommitted or untracked deliverables and all owner-owned paths are preserved as the immediate implementation baseline.
6. Verify the existing engine contracts and real pipeline behavior rather than assuming adapter signatures, completeness semantics, ordering, or compatibility.
7. Confirm no conflicting command-architecture implementation or documentation target already exists.
8. Run the required baseline framework, lifecycle, diversion, Builder, Go, race where supported, vet, format, installed CLI, and Git checks.

Stop on any material difference, unexpected divergence, lifecycle inconsistency, destination collision, unresolved field, validation failure, or inability to distinguish Task 024 targets from prior or owner-owned work.


## Snapshot Requirements

Before modifying any Task 024 implementation or documentation target, create a private UTC-timestamped rollback-capable snapshot outside the repository, under the configured `/tmp` snapshot location.

The snapshot must include:

- the exact starting Git status, branch, HEAD, remotes, ahead/behind evidence, and relevant tags;
- a verified complete Git bundle;
- a bounded archive of every pre-existing planned Task 024 target, including untracked dependencies that Task 024 may modify;
- explicit absence records for every planned new path;
- SHA-256 manifests for captured files, archive, bundle, metadata, and absence records;
- modes, ownership, ACL evidence where supported, and the exact target list;
- a written exact bounded restore procedure that preserves later owner work and requires confirmation before destructive replacement.

Verify checksums, archive listings, bundle integrity, absence records, privacy, readability, and restore boundaries before implementation. Retain the snapshot through owner acceptance. Never use broad reset, checkout, restore, clean, wildcard deletion, or repository-wide archive extraction as rollback.


## Risk Assessment

- High architectural-compatibility risk: a permanent public command grammar can constrain all future interfaces. Mitigate with explicit Command Definition 1.0 contracts, deterministic semantics, extension points, compatibility rules, golden tests, and clear versioning.
- High semantic-divergence risk: simple and advanced commands could produce different behavior. Mitigate by resolving both forms to the same canonical command profile, execution plan, pipeline, and engine adapters, with equivalence tests.
- High responsibility-duplication risk: orchestration could reproduce Inventory, Compare, Drift, Health, Rule, or Report logic. Mitigate with strict adapter boundaries, dependency audits, and end-to-end traceability tests.
- High pipeline-correctness risk: skipped, reordered, or unnecessary stages could produce invalid results. Mitigate with explicit stage dependency validation, profile planning, deterministic ordering, and stage-selection tests.
- Medium public-contract risk: unstable grammar, parameters, output, IDs, or serialization could break scripts and future interfaces. Mitigate with versioned contracts, strict validation, compatibility fixtures, canonical serialization, and unsupported-version behavior.
- Medium CLI-regression risk: existing commands may change unexpectedly. Mitigate with baseline inspection, focused regression tests, installed-binary tests, and documented migration only where explicitly required.
- Medium safety and privacy risk: sources, snapshots, filters, metadata, or terminal output could expose sensitive values or control characters. Mitigate with bounded inputs, safe rendering, redaction preservation, escaping, and privacy tests.
- Medium resource-exhaustion risk: composable filtering, grouping, sorting, or pipelines may become unbounded. Mitigate with explicit limits, deterministic failure, and adversarial tests.
- Medium localization risk: terminal presentation could hardcode assumptions that block future localization. Mitigate by separating canonical data from presentation and keeping user-visible content localization-ready.
- Medium rollback risk in the existing dirty worktree. Mitigate with exact target inventory, verified bounded snapshot, no broad cleanup, and preservation of all unrelated owner-owned content.
- Low dependency and operational risk: no new dependency, daemon, scheduler, remote call, privilege change, host mutation, deployment, commit, or push is authorized.


## Planned Work

### Phase 1 — Baseline and contract audit

Validate governance, lifecycle, repository identity, Git state, existing CLI behavior, and every canonical engine contract. Record the exact orchestration boundaries and create the verified bounded snapshot before target changes.

### Phase 2 — Canonical command architecture

Define the permanent presentation-independent command taxonomy and grammar. Specify simple commands as stable predefined profiles and advanced commands as structured composition over the same model. Prevent an uncontrolled collection of unrelated switches.

### Phase 3 — Versioned command contracts

Define Command Definition 1.0, Command Profile, parameter, selection, filtering, grouping, sorting, formatting, presentation, execution-request, execution-plan, result, diagnostic, and compatibility contracts with strict validation, deterministic identity and ordering, bounded values, canonical serialization where applicable, and explicit unsupported-version behavior.

### Phase 4 — Canonical analysis pipeline

Implement deterministic planning and orchestration for Inventory → Snapshot → Compare → Drift → Health → Rule → Report. Execute only required stages, enforce dependencies and ordering, propagate canonical identities and completeness, and call existing engines without reproducing their logic.

### Phase 5 — Simple and advanced command models

Implement the simple profile resolution for concepts equivalent to `qwsg`, `status`, `check`, `changes`, `health`, and `report`. Implement advanced structured parameter composition. Prove that equivalent simple and advanced definitions resolve to and execute the same canonical plan.

### Phase 6 — Execution and presentation

Implement the Command Execution model and deterministic, human-readable terminal presentation exclusively from canonical execution results. Keep analysis data presentation-neutral and preserve a clean boundary for future Interactive Terminal, Dashboard, and REST API adapters.

### Phase 7 — Tests and compatibility

Add focused unit, contract, golden, negative, determinism, limits, privacy, safe-terminal, compatibility, regression, and integration tests. Add real pipeline tests across existing engines and prove excluded capabilities and duplicated logic are absent.

### Phase 8 — Permanent documentation

Document the Canonical Command Architecture, grammar, taxonomy, profiles, analysis pipeline, orchestration, parameter model, public contracts, deterministic execution, compatibility strategy, terminal presentation, and future Interactive Terminal, Dashboard, and REST API integration boundaries. Update only necessary canonical indexes and Task 024 lifecycle records.

### Phase 9 — Final verification and delivery

Run every required focused and repository-wide validation, review exact diffs and Git state, verify snapshot and rollback evidence, complete the English engineering history and delivery record, and provide the Hungarian owner handoff. Do not begin any future task.


## Rollback Plan

Rollback is bounded to the exact Task 024 targets recorded before implementation.

1. Stop and preserve truthful failure evidence.
2. Verify repository identity, Task ID, current Git state, exact target list, snapshot manifest, hashes, bundle, archive, and absence records.
3. Review the restore set and obtain explicit confirmation before destructive replacement.
4. Restore only pre-existing Task 024 targets from exact snapshot counterparts, preserving modes, ownership, and ACLs where authorized.
5. Remove only Task 024-created paths recorded as absent after confirming they contain no later owner work.
6. Preserve truthful lifecycle evidence; restore prompt/history only if governance files are damaged and exact restoration is separately authorized.
7. Never use broad reset, checkout, restore, clean, wildcard deletion, or archive extraction over the live worktree.
8. Re-run framework, lifecycle, diversion, Builder, focused, repository, race, vet, format, CLI, documentation, privacy, permissions, ACL, Git diff, and status checks.
9. Record exact rollback operations, restored and retained paths, results, and unresolved issues in Task 024 history.

If safe exact restoration cannot be proven, stop and request Project Owner direction.


## Deliverables

- Canonical Command Architecture implementation.
- Canonical Analysis Pipeline implementation.
- Simple Command Model with deterministic predefined command profiles.
- Advanced Command Model with deterministic composable structured parameters.
- Command Profile model.
- Command Definition 1.0.
- Command Execution model.
- Pipeline Orchestration model for Inventory → Snapshot → Compare → Drift → Health → Rule → Report.
- Versioned presentation-independent public contracts, strict validation, canonical serialization where applicable, and compatibility strategy.
- Deterministic human-readable terminal presentation derived from canonical execution results.
- Shared adapters to existing canonical engineering engines with no duplicated engineering logic.
- Engineering unit, contract, golden, negative, compatibility, regression, and real pipeline integration tests.
- Permanent architecture documentation covering:
  - Canonical Command Architecture;
  - Command Grammar;
  - Command Taxonomy;
  - Command Profiles;
  - Canonical Analysis Pipeline;
  - Pipeline Orchestration;
  - Parameter Model;
  - Public Contracts;
  - deterministic Command Execution;
  - terminal presentation;
  - compatibility and extension strategy;
  - future Interactive Terminal integration;
  - future Dashboard integration;
  - future REST API integration;
  - security, privacy, localization, boundedness, and no-side-effect boundaries.
- Necessary canonical architecture, system-map, roadmap, README, Engineering History, and Task 024 lifecycle updates.
- Complete Task 024 history with starting state, snapshot, decisions, changes, exact validation evidence, rollback, Git state, limitations, and result.


## Verification

Verification must include:

- `bin/job --check`;
- `ai/scripts/framework-check.sh` and all configured validations;
- `ai/scripts/next-task.sh --check`;
- `bin/job --check-test-tasks`;
- all Engineering Framework, lifecycle, diversion, and Task Builder assertion suites;
- focused command, profile, grammar, parameter, planning, orchestration, execution, presentation, and compatibility tests;
- all repository Go tests;
- race-enabled Go tests where supported;
- `go vet`, formatting, build, installed CLI, and Git diff checks;
- golden or equivalent Command Definition 1.0 and public-contract tests;
- deterministic repeated parsing, normalization, planning, execution, output, identity, ordering, and canonical serialization;
- equivalence tests proving simple and advanced forms execute the same canonical plan and engines;
- profile tests for concepts equivalent to `qwsg`, `status`, `check`, `changes`, `health`, and `report`;
- source, snapshot, pipeline, engine inclusion/exclusion, filter, grouping, sorting, output-format, and presentation-selection tests;
- tests proving only required stages execute, dependencies cannot be violated, invalid combinations fail explicitly, and no stage executes twice;
- real integration coverage through Inventory → Snapshot → Compare → Drift → Health → Rule → Report where the selected profile requires the full pipeline;
- traceability and completeness tests across every executed stage and final result;
- missing, empty, duplicate, contradictory, invalid, unsupported-command, unsupported-parameter, unsupported-profile, and unsupported-version handling;
- compatibility and regression tests for existing commands and public contracts;
- privacy, redaction, bounded-input, resource-limit, terminal/control-character safety, stable error, and localization-readiness checks;
- tests proving presentation derives only from canonical execution results and never re-analyzes evidence;
- dependency and behavior audits proving no Interactive Terminal, Dashboard, REST API, monitoring daemon, scheduler, Policy, Alert, Automation, AI, machine learning, remote execution, remediation, host mutation, or hidden side effect was introduced;
- exact target mode, ownership, and ACL review where supported;
- snapshot checksum, archive, bundle, absence-record, and bounded rollback validation;
- `git diff --check`, exact staged/unstaged review, and `git diff --cached --check` if staged;
- confirmation that unrelated owner content remains untouched and no unauthorized dependency, infrastructure, release, deployment, staging, commit, push, or future capability occurred.

Record exact commands, meaningful output summaries, assertion counts, limitations, and failures truthfully. Any mandatory failure blocks completion.


## Documentation Updates

Create or update only documentation needed to make the command and analysis interface permanent:

- a dedicated Canonical Command Architecture document;
- Command Grammar and Command Taxonomy;
- Simple and Advanced Command Models;
- Command Profiles;
- Command Definition 1.0 and all versioned public contracts;
- Canonical Analysis Pipeline and stage dependency model;
- Pipeline Orchestration and stage-selection behavior;
- Parameter Model for sources, snapshots, pipelines, engine selection, filters, grouping, sorting, formatting, and presentation;
- deterministic Command Execution model, results, diagnostics, identity, ordering, limits, and serialization;
- human-readable terminal presentation and localization boundary;
- compatibility, deprecation, versioning, and extension strategy;
- security, privacy, boundedness, offline, AI-independent, and no-side-effect guarantees;
- future Interactive Terminal integration;
- future Dashboard integration;
- future REST API integration;
- canonical architecture, system map, roadmap, README, and concise Engineering History references;
- Task 024 history and required repository indexes.

All engineering artifacts must be English. Documentation must distinguish implemented command and orchestration contracts from future Interactive Terminal, Dashboard, REST API, monitoring, scheduling, Policy, Alert, Automation, AI, remote-execution, and remediation components. It must state that every future interface consumes the same canonical command definitions, execution plans, pipeline, results, and engines instead of recreating engineering logic.


## Completion Criteria

Task 024 is complete only when:

- one permanent presentation-independent Canonical Command Architecture exists;
- the Simple Command Model and Advanced Command Model both resolve to the same canonical Command Definition 1.0, profiles, execution plans, pipeline, and engineering engines;
- memorable simple commands provide complete deterministic workflows without requiring internal architecture knowledge;
- advanced structured parameters compose deterministic workflows for source, snapshot, pipeline, engine selection, filtering, grouping, sorting, output formatting, and presentation;
- the canonical Inventory → Snapshot → Compare → Drift → Health → Rule → Report pipeline executes only required stages with validated dependencies, stable order, explicit traceability, and no duplicated analysis logic;
- the Command Profile, Command Definition 1.0, Command Execution, Pipeline Orchestration, parameter, result, diagnostic, and presentation contracts are versioned, strictly validated, deterministic, bounded, documented, and tested;
- human-readable terminal output derives only from canonical execution results and remains localization-ready;
- compatibility, versioning, deprecation, and extension behavior supports future interfaces without destabilizing existing contracts;
- future Interactive Terminal, Dashboard, and REST API interfaces are documented as adapters over the same canonical command architecture and results;
- all required architecture and integration boundaries are documented without implementing excluded future components;
- no Interactive Terminal, Dashboard, REST API, monitoring daemon, scheduler, Policy, Alert, Automation, AI, machine learning, remote execution, remediation, host mutation, duplicated engine, or unauthorized side effect exists;
- all focused, repository, race, vet, format, build, installed CLI, framework, Builder, lifecycle, diversion, determinism, compatibility, privacy, localization, permission, ACL, rollback, and Git checks pass;
- existing public contracts and CLI behavior do not regress except through an explicitly documented and tested compatible transition;
- Task 024 history is complete and truthful, unrelated owner content remains untouched, and no excluded work occurred.

The valid delivery result is `complete`, `complete with disclosed limitations`, or `blocked` under the Engineering Framework. Completion does not authorize Task 025 or any future interface, monitoring, Policy, Alert, Automation, AI, remote-execution, or remediation work.


## Owner Approval Requirements

Approved by Attila — Project Owner through the Engineering Task Builder on 2026-07-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
