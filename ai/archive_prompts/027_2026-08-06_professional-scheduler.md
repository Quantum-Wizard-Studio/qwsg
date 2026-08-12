# Current Engineering Task 027: Professional Scheduler

## Task Metadata

- Task ID: `027`
- Task slug: `professional-scheduler`
- Status: `complete`
- Date opened: `2026-08-06` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

Professional Scheduler


## Objective


Establish the permanent Canonical Professional Scheduler of Quantum Wizard Server Guardian.

Task 027 shall implement a deterministic Scheduler Engine and a minimal explicitly invoked local execution adapter that evaluate validated Schedule Definition 1.0 records from Effective Configuration 1.0, produce canonical scheduling records, and invoke existing Canonical Command and Pipeline orchestration without duplicating their logic.

The Scheduler shall provide a real, testable, restart-safe local scheduling capability while remaining independent of daemon, service-management, monitoring, alerting, notification, remote execution, and presentation responsibilities.



## Scope


Task 027 shall design and implement the Canonical Scheduler as a bounded Professional automation component over the existing immutable engineering core.

The selected implementation boundary is option B: the Canonical Scheduler Engine plus a minimal local execution adapter. The adapter shall execute one explicitly requested scheduling cycle; it shall not create a background loop, daemon, installed service, or implicit startup behavior.

The task scope includes:

- defining Canonical Scheduler Model 1.0;
- defining Scheduler Evaluation Record 1.0;
- defining Scheduler State Record 1.0;
- defining Scheduler Event Record 1.0;
- defining Scheduler Execution Request 1.0;
- defining Scheduler Execution Result Record 1.0;
- defining stable schedule, evaluation, due occurrence, execution request, attempt, state, event, lock-owner, and result identities;
- consuming only validated Effective Configuration 1.0 and Schedule Definition 1.0 records;
- consuming configured concurrency, retry policy, timeout, priority, overlap, misfire, time-zone, daylight-saving, Check applicability, and Command profile values without redefining configuration semantics;
- implementing deterministic enabled and disabled schedule handling;
- implementing deterministic interval and calendar due-time evaluation;
- implementing time-zone-aware next-run calculation through an explicit validated time-zone resolver;
- defining exact daylight-saving evaluation for the existing Schedule Definition 1.0 policies, including ambiguous and nonexistent local times, without changing the configuration contract;
- defining interval schedule anchoring in Scheduler State rather than adding hidden Schedule Definition fields;
- defining a half-open evaluation window and exact behavior for first evaluation, routine cycles, restart, long gaps, backward wall-clock movement, forward discontinuity, and time-zone data failure;
- using supplied wall-clock observations for persisted timestamps and monotonic elapsed observations for interval, timeout, and clock-discontinuity reasoning where available;
- implementing deterministic priority ordering using priority, scheduled occurrence, and stable schedule identity;
- implementing configured global concurrency bounds and explicit delayed, skipped, queued, or indeterminate capacity outcomes;
- implementing overlap behavior with bounded single-run protection and no unbounded backlog;
- implementing finite retry eligibility and deterministic retry planning from the referenced Retry Policy without retry randomness or unlimited attempts;
- implementing exact missed-run behavior for `skip`, `run_once`, and `indeterminate`;
- implementing deterministic next-run calculation without executing future work;
- generating immutable Scheduler Execution Requests that reference the exact Effective Configuration, Schedule Definition, due occurrence, Command profile, Check scope, attempt, trigger, and initiating scheduler evaluation;
- treating an empty Schedule Check scope as execution of the complete referenced Command profile;
- treating a non-empty Schedule Check scope as explicitly inapplicable when Command Definition 1.0 cannot represent that scope, producing no execution rather than silently ignoring, broadening, or inventing Command semantics;
- creating canonical Command Definitions through the existing Command profile resolver and submitting them only to the existing Pipeline Orchestrator;
- recording the resulting Command Execution identity, completeness, stage contracts, Policy Evaluation references and outcomes where present, and failure information without re-evaluating or reinterpreting engineering evidence;
- implementing a minimal local Scheduler Cycle adapter that accepts explicit dependencies: validated Effective Configuration, scheduler-state location, explicit validated Command selection context including required Inventory Store selectors, clock observation source, time-zone resolver, lock provider, Command resolver, and Pipeline executor;
- implementing versioned local scheduler-state persistence with deterministic serialization, integrity verification, restrictive permissions, atomic replacement, corruption isolation, and bounded retention where applicable;
- implementing process-safe local locking and per-schedule single-run protection with explicit lock acquisition, contention, release, stale or invalid lock, and recovery records;
- ensuring interrupted in-progress requests are marked incomplete after restart and are reevaluated under configured misfire and retry policy rather than treated as successful;
- preserving manual Canonical Command and Pipeline execution when Scheduler state, locking, time-zone resolution, or persistence fails;
- defining scheduler self-observability through canonical state, event, and result records without implementing monitoring, Health evaluation, alerts, or notification delivery;
- defining exact schema, contract, taxonomy, engine, state, event, result, and compatibility versions;
- defining strict validation for malformed, incomplete, duplicate, contradictory, unsupported, stale, tampered, oversized, or inapplicable Scheduler records;
- defining deterministic ordering, stable content identities, complete provenance, and byte-stable canonical JSON where applicable;
- preserving offline operation, least privilege, privacy, redaction, bounded-resource behavior, localization boundaries, and AI independence;
- preserving existing CLI behavior and public contracts unless a narrowly additive, versioned, and tested local Scheduler-cycle command adapter is strictly required;
- adding focused unit, contract, time-zone, daylight-saving, interval, calendar, clock-discontinuity, priority, concurrency, overlap, lock, retry, misfire, restart, persistence, integrity, determinism, integration, Policy-traceability, compatibility, privacy, bounded-resource, and regression tests;
- creating permanent Canonical Scheduler architecture documentation;
- updating only directly affected permanent architecture, system-map, roadmap, engineering-history, README, and lifecycle documentation.

The Scheduler Engine shall remain deterministic and side-effect-free. It shall transform validated configuration, scheduler state, and explicit clock observations into scheduling decisions, next-run data, execution requests, and proposed state transitions.

Only the minimal local adapter may acquire locks, persist Scheduler records, and invoke an already-constructed Command/Pipeline dependency. The adapter shall perform one bounded cycle per explicit call and shall not own a recurring loop.



## Out of Scope


Task 027 shall not implement:

- a daemon, monitoring loop, polling service, resident background process, systemd unit, init integration, automatic startup, watchdog, or service manager;
- package installation, service installation, privilege changes, deployment, release, update, repair, removal, or host lifecycle management;
- a user-facing configuration file syntax, configuration discovery, activation, editor, hot reload, or Configuration UI;
- changes to Schedule Definition 1.0 or competing configuration precedence, defaults, identity, validation, or provenance;
- a second Command model, Command parser, Pipeline, engine sequence, report generator, or presentation path;
- Inventory collection, Compare, Drift, Health, Rule, Policy, Report, or Command evaluation logic inside Scheduler;
- Alert Engine, incidents, maintenance windows, suppression decisions, notifications, email, webhooks, delivery channels, or notification retry;
- interpreting a Policy Outcome as notification, remediation, or execution authority;
- automatic remediation, host mutation, remote execution, fleet scheduling, multi-host coordination, or network calls;
- Dashboard, Terminal UI, Web UI, REST API, public management API, or remote-control surface;
- licensing, entitlement enforcement, telemetry, AI, or Machine Learning;
- arbitrary command execution, shell execution, plugin execution, or user-supplied executable paths;
- a general-purpose database or shared product-state platform;
- production support claims for daemon operation, operating-system service lifecycle, or unattended installation.

Task 027 shall not silently select a scheduler-state directory, Inventory Store path, deployment identity, or service account. The minimal adapter receives explicit validated local dependencies. Product installation and runtime-path defaults remain separate lifecycle and packaging work.



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


Task-creation preflight, performed before Builder installation, shall verify:

- the repository is in canonical idle lifecycle state;
- Task 026 is the latest completed task;
- no Task 027 prompt, history, or archive destination exists;
- the structured Builder input is complete, explicitly approved, and passes `task-builder.sh --check-input`;
- Builder installation targets are absent and rollback coverage for installation is verified.

After Builder installation and before modifying any Task 027 implementation or documentation target:

- verify the QWSG repository identity and configured project root;
- verify branch `main`, HEAD, configured remote, canonical remote URL, and local-to-remote relationship;
- verify Task 027 is the sole active approved task and its matching history exists;
- verify Task 026 is the latest completed canonical implementation baseline;
- verify Canonical Policy, Configuration, Schedule Definition 1.0, Command, Pipeline, and Policy-backed Report implementations and documentation are present;
- verify repository, Framework 1.0, lifecycle, Go, vet, formatting, and test-task validation pass;
- verify Framework 1.0 remains authoritative and QWCS principles are guidance only;
- record the complete working-tree state, including all pre-existing unstaged and untracked owner content;
- verify intended Task 027 targets do not overwrite, absorb, stage, remove, or reinterpret unrelated Task 025, Task 026, QWCS, backup, or task-source work;
- identify every intended new and modified target before snapshot creation;
- verify no unexplained lifecycle collision, destination collision, repository divergence, dependency change, unsupported platform assumption, or material architecture mismatch exists;
- report failed preconditions and stop only when they affect authority, safety, rollback, scope, or implementation correctness.

The valid post-install state is one active approved Task 027. The pre-install requirement for canonical idle must not be applied after Builder installation.



## Snapshot Requirements


Before modifying any Task 027 implementation or documentation target:

- create a rollback-capable snapshot outside the repository for every existing target Task 027 may modify;
- record verified absence for every new implementation, test, architecture, history, and archive target;
- capture repository identity, branch, HEAD, remote relationship, lifecycle state, complete Git status, target inventory, permissions, and relevant validation baseline;
- include exact pre-task content for any tracked target that enters scope after initial planning and prove its baseline from recorded clean Git evidence;
- generate a deterministic manifest and SHA-256 checksum list;
- verify every checksum and the readability and complete listing of every archive payload;
- write an exact bounded restore procedure for modified and newly created targets;
- define collision checks that prevent rollback from overwriting later owner work;
- retain the snapshot through Task 027 completion and Owner acceptance;
- keep payloads, private absolute paths, host data, process data, lock contents, Scheduler state, and unpublished source copies outside Git in accordance with the Engineering Backup Policy.

Snapshot scope shall remain proportional to the actual Scheduler target set. Runtime state-store fixtures belong in isolated temporary test directories. Broad repository archives or destructive recovery shortcuts are prohibited.



## Risk Assessment


Primary risks:

- redefining Schedule Definition, Configuration, Policy, Command, Pipeline, or Report semantics inside Scheduler;
- conflating pure schedule evaluation with state persistence, command execution, daemon lifecycle, monitoring, or notifications;
- calendar ambiguity across time zones, daylight-saving transitions, locale changes, leap years, and wall-clock discontinuities;
- interval ambiguity because Schedule Definition 1.0 intentionally stores duration but not an activation anchor;
- duplicate, lost, or fabricated executions after restart, lock contention, state corruption, persistence failure, or clock movement;
- uncontrolled overlap, retry storms, unbounded backlog, starvation, or nondeterministic capacity ordering;
- treating incomplete or failed Command Execution as success;
- treating Policy Outcomes as alert, remediation, or authorization instructions;
- executing arbitrary commands instead of validated Canonical Command profiles;
- silently choosing runtime paths, service identities, daemon behavior, or platform lifecycle;
- mutating existing CLI behavior or preventing manual operation when Scheduler fails;
- leaking host paths, lock tokens, configuration values, Policy evidence, or sensitive execution details;
- tests depending on the real wall clock, local time zone, sleep timing, process races, or external services;
- excessive snapshot or validation requirements becoming unrelated procedural blockers.

Risk mitigation:

- keep the Scheduler Engine pure with explicit configuration, state, and clock inputs;
- place interval anchor, last-evaluation, active-run, and pending-replacement facts in versioned Scheduler State rather than Configuration;
- specify time windows, daylight-saving behavior, clock-discontinuity outcomes, and restart transitions exactly;
- inject deterministic clock, time-zone, lock, persistence, Command-resolution, and Pipeline-execution dependencies;
- use content-derived occurrence, request, attempt, event, state, and result identities for idempotency;
- enforce validated concurrency, finite retries, process-safe locking, per-schedule overlap policy, and at most one coalesced replacement;
- persist state atomically before and after execution and record incomplete transitions explicitly;
- reject corrupt or unsupported state without resetting it and keep manual Command execution available;
- record Policy results only as immutable downstream evidence references;
- execute only Canonical Command Definitions through `internal/pipeline`;
- require explicit local state and workspace dependencies and implement no service lifecycle;
- use virtual-clock, synthetic time-zone, temporary-store, contention, crash-recovery, and failure-injection tests;
- apply full validation only to task-relevant code and repository gates required by Framework 1.0.



## Planned Work


Task 027 shall:

1. inspect Tasks 025 and 026, Schedule Definition 1.0, Command profiles, Pipeline orchestration, Policy-backed Report integration, Product Architecture, Functional Specification, System Map, and Roadmap;
2. define Scheduler ownership and the exact separation among Schedule Definition, evaluation, execution planning, request creation, Command execution, state persistence, daemon lifecycle, monitoring, alerting, and notification;
3. define Scheduler Model, Evaluation, State, Event, Execution Request, and Execution Result contracts with exact versions, identities, provenance, bounds, ordering, and compatibility behavior;
4. define exact interval anchor, calendar evaluation, evaluation window, next-run, time-zone, daylight-saving, restart, missed-run, backward-clock, forward-discontinuity, and unavailable-time-zone semantics;
5. implement the pure deterministic Scheduler Engine over Effective Configuration, Scheduler State, and explicit clock observations;
6. implement deterministic priority, configured concurrency, overlap, one-replacement backlog, single-run protection, finite retry eligibility, and retry planning;
7. implement canonical serialization and validation for Scheduler inputs and outputs;
8. implement versioned local Scheduler state persistence with integrity checking, restrictive permissions, atomic replacement, explicit corruption behavior, and bounded records;
9. implement process-safe locking with contention and recovery evidence without service or process management;
10. implement the minimal explicitly invoked local Scheduler Cycle adapter that acquires the lock, loads state, evaluates due work, persists reservations, resolves existing Command profiles, calls the existing Pipeline Orchestrator, and persists results and final state;
11. capture Command Execution and Policy Evaluation identities and outcomes without evaluating Policy or implementing Alert behavior;
12. prove manual Command and Pipeline operation remains available when Scheduler evaluation, locking, state, time-zone, or persistence fails;
13. add focused unit, contract, determinism, temporal, daylight-saving, discontinuity, priority, capacity, overlap, lock, retry, misfire, restart, corruption, persistence, integration, traceability, privacy, resource, compatibility, and regression tests using controlled dependencies;
14. create `docs/architecture/CANONICAL_SCHEDULER.md` and update only directly affected canonical documentation;
15. run every required build, test, race, static-analysis, formatting, framework, lifecycle, documentation, source-boundary, rollback, permission, and Git-state validation;
16. record exact implementation, decisions, validation evidence, failed attempts, limitations, rollback, and Git state in Task 027 history;
17. mark Task 027 complete only after all gates pass, archive it without creating a successor, and return the repository to canonical idle without staging, committing, or pushing.



## Rollback Plan


If implementation or validation fails:

- stop further mutation and preserve failure, state, lock, and validation evidence without recording secret or private payloads;
- verify the Task 027 snapshot manifest, checksums, payload readability, target inventory, absence records, and restore instructions;
- compare each affected target with the snapshot and current working tree;
- refuse rollback over later or unrelated owner work;
- release only locks proven to belong to the failed isolated Task 027 test or adapter invocation; never delete an ambiguous lock;
- restore only verified pre-existing Task 027 targets from the bounded snapshot;
- remove only verified Task 027-created targets whose pre-task absence was recorded and which contain no later owner work;
- preserve corrupt or failed runtime-state fixtures as isolated test evidence until validation completes; do not reset real Scheduler state silently;
- preserve truthful Task 027 lifecycle and history evidence rather than rewriting completed or failed facts;
- rerun framework, lifecycle, repository, Go, race, vet, formatting, documentation, focused Scheduler, permission, and snapshot-integrity validation after restoration;
- report resulting repository consistency and every unresolved condition.

Broad `git reset`, `git checkout`, `git restore`, `git clean`, wildcard deletion, repository-wide extraction, ambiguous lock deletion, and silent state reset are prohibited.



## Deliverables


- Canonical Professional Scheduler implementation;
- Canonical Scheduler Model 1.0;
- Scheduler Evaluation Record 1.0;
- Scheduler State Record 1.0;
- Scheduler Event Record 1.0;
- Scheduler Execution Request 1.0;
- Scheduler Execution Result Record 1.0;
- deterministic interval and calendar evaluation;
- exact time-zone and daylight-saving behavior;
- next-run and missed-run calculation;
- deterministic priority, concurrency, overlap, and bounded backlog behavior;
- finite retry eligibility and retry planning;
- restart, interruption, and clock-discontinuity semantics;
- stable identities, provenance, canonical serialization, and strict validation;
- versioned local scheduler-state persistence and integrity handling;
- process-safe local locking and single-run protection;
- minimal explicitly invoked local Scheduler Cycle adapter;
- Canonical Command and Pipeline integration with Policy-result traceability;
- permanent Scheduler architecture documentation;
- focused engineering tests;
- updated directly affected canonical documentation and Task 027 history.



## Verification


Successful completion requires:

- identical Effective Configuration, Scheduler State, clock observations, and time-zone data produce byte-identical evaluation, request, event, state-transition, result, identity, and next-run outputs;
- Scheduler input enumeration, map ordering, locale, or presentation cannot change scheduling decisions;
- disabled schedules never create execution requests and remain visible in evaluation records;
- interval schedules use explicit Scheduler State anchors and monotonic elapsed evidence where available;
- calendar schedules handle month, weekday, leap-year, time-zone, ambiguous, and nonexistent local-time fixtures according to documented Schedule Definition 1.0 evaluation semantics;
- first evaluation, normal cycles, long gaps, restart, interrupted execution, backward clock, forward discontinuity, and unavailable time-zone fixtures produce explicit deterministic outcomes;
- `skip`, `run_once`, and `indeterminate` missed-run behavior is exact and bounded;
- overlap and lock contention cannot create duplicate execution, and any replacement backlog is coalesced to at most one request per schedule;
- global concurrency and priority ordering are deterministic and bounded by Effective Configuration;
- retries are finite, observable, idempotent by attempt identity, and generated only for eligible incomplete or failed execution results;
- Scheduler Execution Requests reference exact configuration, schedule, occurrence, evaluation, Command profile, Check scope, and attempt identities;
- the minimal adapter resolves only existing Canonical Command profiles and invokes only the existing Pipeline Orchestrator;
- Command Execution, Policy Evaluation, and Policy-backed Report identities remain unchanged and traceable through Scheduler results;
- Scheduler never evaluates Inventory, Compare, Drift, Health, Rule, Policy, Report, or Command semantics;
- state writes are atomic, integrity-checked, versioned, permission-restricted, and reject corruption or unsupported versions without silent reset;
- process-safe lock acquisition, contention, release, invalid-lock, failure, and recovery behavior is tested without sleep-based timing;
- empty Check scope schedules execute the complete referenced Command profile, while non-empty scopes unsupported by Command Definition 1.0 produce a deterministic inapplicable record and no broadened execution;
- Scheduler failure does not corrupt canonical records or prevent direct manual Command/Pipeline execution;
- serialization and diagnostics preserve privacy and contain no secret material, private host evidence, arbitrary command payload, or unsafe path disclosure;
- source audits prove no daemon loop, service manager, notification, email, webhook, remediation, remote execution, arbitrary shell, network call, AI, or host mutation was introduced;
- focused Scheduler package, state-store, lock, adapter, and integration tests pass;
- existing Configuration, Command, Pipeline, Policy, Report, CLI, and repository regression tests pass;
- `make build` passes;
- `make test` passes;
- `GOCACHE=/tmp/qwsg-go-cache GOMODCACHE=/tmp/qwsg-go-modcache go test -race ./...` passes;
- `make vet` passes;
- `make fmt-check` passes;
- `ai/scripts/framework-check.sh --run-validations` passes;
- `make engineering-test` passes;
- `bin/job --check`, `ai/scripts/next-task.sh --check`, and `bin/job --check-test-tasks` pass at applicable active, completed, archived, and idle phases;
- `git diff --check`, `git diff --cached --check`, documentation-reference audits, and source-boundary audits pass;
- snapshot checksums, archive readability, target inventory, absence records, permissions, and guarded restore procedure pass;
- final Git evidence identifies exact changed, staged, and untracked paths and proves unrelated owner content remained untouched.

Tests shall use injected or synthetic clock observations, deterministic time-zone fixtures, isolated temporary state directories, controlled lock providers, and fake Command/Pipeline dependencies where possible. Validation shall not depend on wall-clock waiting, an installed daemon, systemd, network access, email, external services, or privileged execution.



## Documentation Updates


Create:

- `docs/architecture/CANONICAL_SCHEDULER.md`.

Update only where directly affected:

- `README.md`;
- `ai/core/04_ARCHITECTURE.md`;
- `ai/core/05_SYSTEM_MAP.md`;
- `ai/core/07_ENGINEERING_HISTORY.md`;
- `ai/core/13_ROADMAP.md`;
- `docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md` for the Task 027 consumer boundary only;
- `docs/architecture/CANONICAL_COMMAND_ARCHITECTURE.md` for the Scheduler adapter boundary only;
- `docs/architecture/CANONICAL_POLICY_ENGINE.md` and `docs/architecture/CANONICAL_REPORT_ENGINE.md` only for immutable result-consumer traceability if directly affected;
- `docs/PRODUCT_ARCHITECTURE.md` or `docs/FUNCTIONAL_SPECIFICATION.md` only for non-semantic traceability corrections strictly required by implementation evidence;
- directly affected development mapping or user documentation only if a supported user-visible adapter is added;
- the Builder-generated Task 027 prompt and matching dated history record through the canonical lifecycle.

Do not rewrite unrelated architecture, historical tasks, Framework 1.0 documents, QWCS architecture, or the QWCS Migration Blueprint.



## Completion Criteria


Task 027 is complete only when:

- the Canonical Scheduler is the sole QWSG owner of schedule evaluation, due occurrences, next-run calculation, scheduler state, execution request generation, retry planning, overlap decisions, and scheduling records;
- Scheduler Model, Evaluation, State, Event, Execution Request, and Execution Result 1.0 contracts are implemented, versioned, immutable, deterministic, bounded, validated, and documented;
- the Scheduler consumes Effective Configuration 1.0 and Schedule Definition 1.0 without redefining configuration semantics;
- interval, calendar, time-zone, daylight-saving, missed-run, restart, clock-discontinuity, priority, concurrency, overlap, lock, retry, and next-run behavior has exact tested semantics;
- versioned local state persistence survives restart fixtures, preserves incomplete execution truth, detects corruption, and never silently resets;
- process-safe locking and per-schedule single-run protection prevent duplicate or uncontrolled overlap;
- the explicitly invoked local adapter can perform one real bounded scheduling cycle through existing Command and Pipeline contracts and record the resulting Command and Policy evidence;
- manual canonical Command and Pipeline execution remains available when Scheduler components fail;
- Scheduler does not duplicate Inventory, Compare, Drift, Health, Rule, Policy, Report, Command, Pipeline, Configuration, or presentation logic;
- daemon/service lifecycle, monitoring, Alert Engine, notifications, configuration activation, arbitrary execution, remote operation, network behavior, AI, and host mutation remain unimplemented;
- current supported CLI and canonical engine behavior remain compatible;
- all mandatory engineering validations pass;
- rollback evidence is complete and verified;
- documentation and Task 027 history truthfully describe the delivered and excluded boundaries;
- no unrelated owner content is modified, staged, removed, or absorbed;
- no files are staged, committed, or pushed unless separately authorized;
- Task 027 is completed and archived and the repository returns to canonical idle lifecycle state.



## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-06 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
