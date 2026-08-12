# Current Engineering Task 032: Canonical Operator Presentation Model

## Task Metadata

- Task ID: `032`
- Task slug: `canonical-operator-presentation-model`
- Status: `complete`
- Date opened: `2026-08-07` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Canonical Operator Presentation Model


## Objective


Establish the Canonical Operator Presentation Model as the single deterministic, presentation-independent read model that translates existing validated QWSG engineering and operational contracts into an immediately understandable operator overview.

Task 032 shall define a versioned Operator Presentation Model 1.0 that answers, without requiring knowledge of internal engines or canonical identifiers:

- whether the observed server is healthy;
- whether attention is needed and at what severity;
- what materially changed;
- whether active, acknowledged, suppressed, recovered, or expired Alert facts exist;
- whether the Guardian Runtime Service is running, stopping, stopped, failed, or not observed;
- how fresh and complete the underlying evidence is;
- what bounded, non-remediating action is recommended, if any.

The model shall consume only explicit validated canonical outputs and explicit observation context. It shall preserve their authority, uncertainty, provenance, and identities while projecting stable operator concepts, fixed localization tokens, bounded counts, ordered attention items, change summaries, service status, freshness, and recommended next-step tokens. It shall not collect, execute, schedule, monitor, persist, render an interface, or reinterpret engineering truth.

Task 032 ends when the shared read model, deterministic projection engine, strict contracts and validation, tests, architecture documentation, rollback evidence, and canonical lifecycle closure are complete. Interactive Terminal Console, bare `qwsg` integration, REST API, Web Dashboard, persistence, monitoring, providers, and daemon installation remain later tasks.



## Scope


Task 032 shall define and implement one presentation-independent package, expected at `internal/presentationmodel`, with Canonical Operator Presentation Model 1.0 contracts for:

- an explicit Projection Input containing observation time, optional validated Command Execution/typed canonical stage outputs, optional validated Runtime Result, optional validated Runtime Service State or terminal Service Result, and explicit freshness policy;
- a top-level Operator Overview with stable schema/model version, identity, observed-at time, overall condition, attention state, evidence freshness/completeness, Guardian state, bounded summary counts, ordered change summary, ordered attention items, and ordered recommendations;
- exact overall-condition values such as `healthy`, `degraded`, `critical`, `unknown`, and `unavailable`, with documented precedence and no numeric or probabilistic score;
- exact attention values such as `none`, `review`, `urgent`, and `unknown`, derived only from canonical Alert, Health, Rule, Policy, Runtime, Service, and completeness facts through a closed documented mapping;
- Guardian status that distinguishes `running`, `starting`, `stopping`, `stopped`, `failed`, and `not_observed` without probing processes, PID files, systemd, or the host;
- freshness and completeness states that distinguish current, stale, partial, unsupported, missing, invalid, and not observed evidence without presenting absence as health;
- bounded change summaries derived from validated Comparison/Drift facts already present in canonical execution outputs, without performing comparison or drift classification;
- bounded Alert summaries and attention items derived from immutable Alert Records and lifecycle facts already present in validated Runtime outputs, without creating, correlating, suppressing, acknowledging, expiring, recovering, or persisting Alerts;
- fixed recommendation tokens limited to safe read-only next steps such as inspect attention, review changes, run a fresh check, inspect failed operation, verify Guardian operation, or no action; recommendations shall not authorize or describe remediation commands;
- stable source references that retain canonical contract/version/record identities for advanced drill-down while keeping identifiers out of the beginner summary by default;
- canonical ordering, strict validation, canonical JSON, UTC normalization, stable content identity, bounded strings/collections/counts, privacy limits, fixed reason/title/action tokens, schema evolution, and compatibility behavior;
- localization-ready display tokens and parameter values only; no English or Hungarian prose shall be embedded as canonical meaning.

The projection shall use a closed precedence table. Canonical source ownership remains unchanged: Comparison owns factual changes; Drift owns semantic classification; Health owns engineering condition; Rule owns matching; Policy owns governance interpretation; Report owns engineering presentation artifacts; Alert owns alert existence and lifecycle; Runtime owns one-cycle outcome; Runtime Service owns process lifecycle. Task 032 may combine those facts for operator comprehension but shall never recompute or contradict them.

Missing optional inputs shall produce explicit `unknown`, `unavailable`, `not_observed`, or partial states according to the contract. The model shall not claim a server is healthy merely because no Alert Record, Runtime Result, Service observation, or stored history was supplied.

The existing Command Architecture remains the advanced public composition boundary. Task 032 may add a narrow typed projection seam or validator to expose already-produced canonical stage values safely, but it shall not change Command grammar, profiles, pipeline selection, engine order, or execution semantics. Existing `command.Execution.View` remains stage metadata and shall not be stretched into a competing operator model.

The existing Canonical Report Engine remains the owner of deterministic engineering reports. Task 032 shall consume or reference Report/Policy Report facts where useful and shall not recreate report sections, rule explanations, or policy evaluation.

No terminal renderer is required. A minimal diagnostic text fixture may exist only in tests if necessary to prove token completeness; production output of the task is the canonical read model and canonical JSON contract, ready for later replaceable interfaces.



## Out of Scope


Task 032 shall not implement:

- an Interactive Terminal Console, TUI framework, full-screen mode, navigation, keyboard handling, color theme, screen layout, terminal detection, or bare `qwsg` behavior;
- a new CLI command, changed default CLI behavior, REST endpoint, HTTP listener, Web Dashboard, desktop interface, public API server, or remote interface;
- collection, Inventory assembly, Snapshot storage, comparison, Drift classification, Health evaluation, Rule matching, Policy interpretation, Report generation, Alert decisions, Notification planning/delivery, Scheduler evaluation, Runtime execution, or Runtime Service recurrence;
- a second Command Definition, parser, profile registry, plan, Pipeline, stage order, view filter/group/sort language, report model, alert model, or service-health engine;
- process discovery, PID inspection, systemd queries, watchdog, health probes, heartbeat transport, monitoring loop, polling, timer, goroutine, background worker, daemon, installation, enablement, supervision, or automatic restart;
- durable Operator Overview, Runtime, Service, Alert, Notification, evidence, history, audit, incident, trend, graph, or recommendation persistence;
- restart recovery, checkpointing, replay, migration, transactional cross-engine state, database, queue broker, file activation, workspace adoption, retention cleanup, backup, or restore product behavior;
- Alert acknowledgement or suppression commands, incident workflow, maintenance editing, configuration editing/activation/reload, secrets, credentials, recipient management, provider selection, or concrete Email/Webhook/other transport;
- remediation, repair, shell commands, host mutation, automatic action, remote execution, AI, machine learning, probabilistic scoring, generated advice, licensing, fleet management, packaging, deployment, release, or support declaration;
- dependency installation, infrastructure mutation, staging, commit, push, fetch, branch, or tag operations.

Task 032 may document how later CLI, Console, REST API, and Dashboard adapters consume the model. It shall not implement those adapters or select a UI toolkit. Persistence and restart recovery remain mandatory Version 1.0 operational gates, but they are not prerequisites for defining and testing the shared operator semantics from explicit inputs.



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


### Task-creation state

- verify the exact QWSG root, Framework 1.x configuration, repository markers, canonical remote URL, primary branch, HEAD, local/remote relationship, complete Git status, empty staged path list, ownership, permissions, and ACLs;
- run `bin/job --check`, `ai/scripts/next-task.sh --check`, `bin/job --check-test-tasks`, and `ai/scripts/framework-check.sh`; require canonical idle with Task 031 as the unique latest completed archive/history pair;
- verify `ai/prompts/` is empty and Task 032 prompt, history, archive, `internal/presentationmodel`, tests, and `docs/architecture/CANONICAL_OPERATOR_PRESENTATION_MODEL.md` are absent;
- preserve every pre-existing unstaged and untracked Owner-owned change, including completed Tasks 025–031, QWCS documents, Builder sources, backups, and this prepared source; require the index to remain empty;
- inspect actual Command, Pipeline, Presentation, Report, Alert, Runtime, and Runtime Service public contracts and tests; confirm the existing Command View contains stage metadata only and no existing contract answers the complete operator overview;
- prove that the selected model can consume validated existing outputs through narrow typed boundaries without changing their semantics, and that future CLI, Console, API, and Dashboard can reuse it without depending on terminal behavior;
- record the Release Minimalism decision: a Console now would duplicate operator interpretation; persistence/recovery, monitoring, providers, and installation are necessary later operational gates but do not define shared user meaning; the presentation model is the smallest prerequisite that prevents interface-specific rework;
- validate `current-task-job.txt` as readable non-symlink UTF-8 data with LF-normalizable line endings, no NUL, unfinished fence, placeholder, embedded approval protocol value, ambiguous field mapping, competing lifecycle, or unresolved authority decision.

### Separately authorized Builder installation

- repeat canonical idle, Task 031 baseline, Task 032 destination absence, source content/hash, repository/Git state, Framework, permissions, ACL, Builder interface, and lifecycle checks immediately before installation;
- map only Owner-supplied fields from this source into a mode-0700 temporary Builder input directory; keep Builder-generated metadata, fixed Required Reading, approval prose, and approval protocol data separate;
- create and verify a proportional external Builder-installation snapshot covering exact lifecycle destinations, verified absence records, source/hash evidence, repository identity, Task 031 idle baseline, Git state, ownership, permissions, and ACLs;
- run `task-builder.sh --check-input` only with a separately supplied exact Owner approval protocol value; read-only validation shall not install;
- present the complete generated task for review, obtain separate explicit Owner authorization, and install exactly one Task 032 prompt/history pair transactionally with no clobber;
- require `ai/prompts/032_CURRENT_TASK.md` to be the sole active approved task with one matching history, Task 031 as the latest completed baseline, and all Builder, lifecycle, Framework, repository, permission, and Git validations passing;
- stop after installation at `APPROVED AND READY FOR IMPLEMENTATION`; do not begin Task 032.

### Implementation starting state

- start only through an explicit canonical `job` invocation after successful Builder installation;
- read every Required Reading item and the active prompt/history as data;
- require Task 032 as the sole active approved task and Task 031 as the unique latest completed baseline;
- verify exact public validators, schemas, identities, typed output access, package dependency direction, target absence, build/test availability, repository/Git state, ownership, permissions, and ACLs;
- create and verify the proportional implementation snapshot before modifying any target;
- stop on any material authority, lifecycle, scope, compatibility, source-contract, privacy, localization, rollback, or correctness difference.

The valid preparation and pre-install state is canonical idle. The valid post-install state is one active approved Task 032. The valid post-completion state is canonical idle with Task 032 as the unique latest completed archive/history pair and no successor.



## Snapshot Requirements


Task preparation modifies only `current-task-job.txt` and shall not create an implementation or Builder-installation snapshot. Preserve the previous source through the workspace history available to the Project Owner and record its pre-edit status and hash when available.

Before separately authorized Builder installation:

- create a unique external snapshot under `/tmp` of the exact Task 032 prompt/history lifecycle destinations, their verified absence, `current-task-job.txt` content and SHA-256, repository identity, Task 031 idle baseline, complete Git state, ownership, permissions, and ACLs;
- verify manifest, checksums, payload readability, absence evidence, collision guards, and exact bounded restore instructions before installation;
- retain the snapshot through Builder validation and Owner acceptance.

Before separately authorized implementation:

- create one unique rollback-capable snapshot outside the repository for every existing directly affected source, test, architecture, product, roadmap, system-map, README, prompt, history, and archive target;
- record verified absence for `internal/presentationmodel`, its test files, `docs/architecture/CANONICAL_OPERATOR_PRESENTATION_MODEL.md`, and the Task 032 archive destination;
- preserve exact working-tree content, including pre-existing Owner-owned modifications; never substitute HEAD content for working-tree state;
- record repository identity, branch, HEAD, remotes, ahead/behind, complete Git state, staged path list, target inventory, ownership, permissions, ACLs, baseline validations, manifest, SHA-256 checksums, readable payload inventory, and guarded restore instructions;
- verify every checksum, payload, absence record, collision guard, restore precondition, and proportional target before implementation;
- retain the snapshot through completion and Owner acceptance.

Snapshot scope shall exclude broad repository archives, build caches, live host evidence, running process state, systemd state, persistent product stores, provider payloads, secrets, credentials, network responses, and unrelated data.



## Risk Assessment


Primary risks and mandatory mitigations:

- The model could become a second engineering engine. Validate and map only owning-engine outputs through a closed table; never compare, classify Drift, evaluate Health/Rules/Policy, create Alerts, or infer service process state.
- Existing Command View or Report semantics could be duplicated. Keep Command View as stage metadata, keep Report as engineering-report authority, and define only the cross-domain operator overview missing from both.
- A Console-first implementation could lock meaning to terminal widgets. Produce only canonical contracts, tokens, ordering, and JSON; add no screen, color, navigation, or terminal framework.
- Missing evidence could be presented as healthy. Use explicit unknown/unavailable/not-observed states and make completeness/freshness part of every overall conclusion.
- Runtime Service `running` is observable only while a process is active, while terminal results describe stopped/failed completion. Accept explicit Service State or Result observations and never claim live status from age, PID guesses, or absence.
- Freshness could use hidden wall time. Require explicit observation time and explicit bounded freshness policy; normalize UTC and test exact boundaries and clock-invalid input.
- Recommendations could become remediation or new policy. Limit output to a closed set of read-only navigation/check tokens whose precedence follows existing severity, failure, freshness, and completeness facts.
- Alert summaries could recreate incident lifecycle. Consume immutable Alert Records and existing lifecycle values only; no acknowledgement, suppression, correlation, recovery, or expiry decisions occur here.
- Typed values inside Command Stage Results could be unsafe or ambiguous. Add only narrow contract-specific extraction/validation with exact schema checks and fail closed on type/schema/identity mismatch.
- Multiple sources could conflict or represent different observations. Require explicit correlation and observation context, reject incompatible bundles, retain source references, and expose partial/unknown instead of guessing.
- Canonical identifiers could overwhelm amateur users or leak data. Keep references available for drill-down but separate from concise summaries; enforce privacy-bounded tokens and exclude raw host values, paths, errors, configuration, destinations, provider payloads, and secrets.
- Localization could be hardcoded into the model. Store stable tokens and typed parameters only; language-specific prose belongs to later replaceable adapters.
- Collections or source payloads could grow without bound. Define strict input/output cardinality, string, count, and reference limits; summaries remain bounded and deterministic.
- New imports could reverse dependency direction. The presentation model may depend on public canonical contracts; canonical engines, Runtime, Runtime Service, Command, and Report shall not depend on it.
- Repository tests may be disproportionate or unavailable. Use pure fixtures with no live host, clock, service, network, provider, persistence, privilege, or terminal dependency; retain full existing repository validation as regression evidence.
- Owner-owned dirty files overlap likely documentation targets. Snapshot exact working-tree versions, edit only authorized targets, inspect each diff, keep the index empty, and stop on material overlapping changes that cannot be preserved.
- Scope could expand into Version 1.0 operational infrastructure. Enforce source/import audits and explicit exclusions for persistence, monitoring, provider, Console, CLI integration, installation, and release.

Overall implementation risk is medium because the task defines public user meaning across several canonical domains. Mutation and operational risk remain low because the implementation is pure, local, deterministic, read-only, and side-effect free.



## Planned Work


### Phase 1 — Verify and freeze ownership

- complete Starting State Verification and implementation snapshot requirements;
- inventory exact canonical values available from Command/Pipeline, Comparison, Drift, Health, Rule, Policy, Report, Alert, Runtime, and Runtime Service;
- publish an ownership and precedence matrix showing which source exclusively answers each operator question and how missing/partial/stale facts propagate;
- confirm no existing model already satisfies the complete overview and record why extending Command View or Report would violate their current boundary.

### Phase 2 — Define Operator Presentation Model 1.0

- define explicit input, overview, condition, attention, Guardian, freshness, completeness, summary, change, attention-item, recommendation, and source-reference contracts;
- define closed taxonomies, ordering, identity, UTC, strict decoding, JSON, compatibility, privacy, and resource rules;
- define exact mapping and precedence tables, including contradictions, partial results, unavailable sources, stale evidence, Service failure, Runtime failure, Alert severities/lifecycles, and no-evidence cases;
- define localization token registries and prove that canonical semantics contain no interface prose.

### Phase 3 — Implement the pure projection

- implement strict input validation and narrow typed canonical source extraction;
- implement deterministic projection without ambient time, filesystem, environment, process, service manager, network, persistence, goroutine, or engine invocation;
- validate the completed overview before returning it and provide canonical JSON/strict decode helpers if consistent with repository contract patterns;
- make invalid, incompatible, uncorrelated, over-limit, or tampered inputs fail closed with bounded fixed errors/tokens.

### Phase 4 — Prove reuse and boundaries

- add focused golden/table tests for each mapping, precedence, missing/partial/stale case, ordering, identity, localization token, privacy limit, resource bound, and strict validation rule;
- add contract tests showing equivalent explicit inputs produce byte-identical outputs and that a CLI, Console, REST, or Dashboard adapter can consume the same overview without engine access;
- add source/import audits proving no reverse dependencies, presentation framework, operational probe, persistence, monitoring, provider, or remediation behavior exists;
- retain existing CLI and advanced Command behavior unchanged through regression tests.

### Phase 5 — Document and close

- create the permanent architecture document and update only directly affected architecture, product, roadmap, system-map, README, prompt, and history records;
- run every required verification and inspect exact diffs, permissions, ACLs, snapshot evidence, and rollback validity;
- complete and archive Task 032 without installing or beginning a successor; restore canonical idle.

Decision points are limited to verified source compatibility. If an existing public validator cannot safely expose a required already-owned fact, add the smallest read-only seam in its owning package without changing semantics. Any need for new engineering judgment, durable state, process observation, interface behavior, or remediation is a hard scope stop requiring Owner authority.



## Rollback Plan


Rollback is file-bounded and preserves all pre-existing Owner-owned work.

Before rollback, stop active task work, record the current Task 032 diff and validation failure, verify the snapshot manifest/checksums, confirm the exact repository root and target list, and obtain Project Owner confirmation before any destructive overwrite of material post-snapshot work.

Restore only snapshotted pre-existing files from their exact snapshot payloads, preserving recorded modes and ownership where authorized. Remove only Task 032 paths whose verified absence was recorded and only after confirming they still contain Task 032 artifacts; never use wildcards, recursive repository cleanup, `git reset`, `git clean`, broad checkout/restore, or deletion of unrelated untracked files.

For lifecycle rollback during implementation, restore the exact Task 032 prompt/history working copies and remove an archive only if the snapshot proves it was absent and identity checks prove it was created by Task 032. Builder installation rollback remains the Builder transaction's own no-clobber restoration procedure and its external snapshot.

After rollback, verify checksums, target absence/presence, permissions, ownership, ACLs, complete Git status, empty index, Framework validation, lifecycle identity, `bin/job --check`, `ai/scripts/next-task.sh --check`, and relevant baseline Go tests. Retain the failed-work diff, snapshot, manifest, checksums, and rollback report for Owner review.

Rollback does not remove or alter pre-existing Task 025–031, QWCS, Builder, backup, documentation, source, test, or other Owner-owned changes outside the exact Task 032 target manifest.



## Deliverables


- Canonical Operator Presentation Model 1.0 in one presentation-independent package;
- explicit Projection Input and Operator Overview contracts;
- exact condition, attention, Guardian state, freshness, completeness, change-summary, Alert-summary, recommendation, localization-token, source-reference, failure, version, and resource taxonomies;
- documented ownership, mapping, precedence, conflict, missing-evidence, partial-evidence, stale-evidence, and not-observed semantics;
- a pure deterministic projection that consumes validated canonical outputs and returns one bounded overview without invoking engines or operational systems;
- canonical JSON, strict validation/decoding, deterministic ordering and identity, UTC behavior, privacy rules, and compatibility strategy;
- narrow typed read-only extraction/validation seams only where existing canonical outputs require them;
- comprehensive unit, golden, contract, determinism, boundary, privacy, localization, resource, and regression tests;
- `docs/architecture/CANONICAL_OPERATOR_PRESENTATION_MODEL.md`;
- directly affected Product Architecture, Functional Specification, Roadmap, System Map, project Architecture/Engineering History, README, prompt/history/archive updates;
- complete Task 032 history, verified rollback evidence, archived completed prompt, and canonical idle closure.

No Console, TUI, CLI integration, default-command change, API, Dashboard, renderer, monitoring loop, process probe, persistence, restart recovery, concrete provider, configuration activation, daemon installation, remediation, package, deployment, or release artifact is a Task 032 deliverable.



## Verification


Builder and lifecycle verification shall prove:

- exact Task 032 ID/slug/title, authority, language, mandatory sections, unique Task 031 baseline, no destination collisions, and no placeholders or unresolved content;
- explicit separation of task creation, Builder installation, implementation starting state, implementation work, completion, archive, and canonical idle;
- no embedded or inferred approval protocol value in `current-task-job.txt`;
- lossless mapping of every owner-authored Builder field and successful canonical read-only Builder input validation using only separately supplied approval protocol data;
- correct pre-install idle, post-install sole-active approved, and post-completion idle states with exact prompt/history/archive identity and no successor.

Implementation verification shall include:

- focused `internal/presentationmodel` tests and all directly affected canonical package regression tests;
- `make build`, full `go test ./...`, repository-wide `go test -race ./...` with writable configured caches, `go vet ./...`, and complete Go formatting checks;
- Framework 1.x configured validations, `make engineering-test`, Builder tests, lifecycle checks, diverted-test audit, active-task validation, and final idle validation;
- golden or equivalent tests for every public Model 1.0 contract and byte-identical canonical JSON/identity for equivalent explicit inputs;
- complete mapping-table tests for healthy/degraded/critical/unknown/unavailable, none/review/urgent/unknown attention, every Guardian state, freshness boundary, completeness state, Runtime outcome, Service lifecycle/result, Alert severity/lifecycle/event, and relevant canonical report/policy/health/change facts;
- precedence and contradiction tests proving critical owned facts cannot be masked by healthy facts, missing evidence cannot become healthy, and incompatible or uncorrelated inputs fail closed;
- change-summary tests proving only validated Comparison/Drift outputs are summarized and no comparison or classification is performed;
- Alert tests proving only validated Alert-owned facts are summarized and no lifecycle decision or persistence is performed;
- recommendation tests proving every recommendation is a fixed bounded read-only token, no remediation authority is implied, and no free-form generated advice exists;
- service-status tests proving `running` requires an explicit valid running Service State observation and is never inferred from a terminal result, timestamp age, process, PID, or systemd;
- freshness tests using explicit times only, including exact boundary, stale, future/invalid, partial, missing, and unsupported evidence;
- strict typed-extraction tests rejecting wrong Go type, schema, version, identity, stage correlation, duplicate source, malformed value, and excessive input;
- privacy tests excluding raw errors, report prose, host values and paths, configuration bodies, secret/destination/provider references, credentials, environment data, and unbounded metadata;
- localization tests proving canonical output contains registered stable tokens/parameters and no interface-language prose;
- resource tests for maximum sources, changes, alerts, attention items, recommendations, references, string lengths, counts, and overflow behavior;
- import/source audits proving canonical engines, Runtime, Runtime Service, Command, and Report do not depend on the presentation model and the new package contains no collector, executor, scheduler action, process probe, systemd, filesystem persistence, network, UI toolkit, terminal control, provider, remediation, remote, AI, package, installation, or deployment behavior;
- reuse fixtures demonstrating one identical overview can be consumed by hypothetical terminal, JSON API, and dashboard adapters without changing its semantics; no production adapter is implemented;
- regression tests proving current bare `qwsg` help, explicit CLI commands, structured JSON, terminal renderer, profiles, advanced composition, and Pipeline behavior remain unchanged;
- documentation consistency review against Product Architecture, Functional Specification, Roadmap, System Map, Command, Report, Alert, Runtime, and Runtime Service boundaries;
- exact changed-target audit, ownership, permissions, ACLs, staged/unstaged paths, `git diff --check`, `git diff --cached --check`, and preservation of unrelated Owner-owned content;
- snapshot checksum, payload, readability, absence evidence, collision guards, bounded restore verification, and confirmation that rollback remains usable;
- confirmation that nothing was installed, activated, monitored, persisted as product state, staged, committed, pushed, packaged, deployed, or released.

Verification requires no live host collection, real clock sleep, real signal, running daemon, systemd/service manager, persistent product store, provider, credential, network, privileged operation, remote system, terminal emulator, or infrastructure mutation.



## Documentation Updates


Expected direct documentation targets are:

- `docs/architecture/CANONICAL_OPERATOR_PRESENTATION_MODEL.md`;
- `docs/architecture/CANONICAL_COMMAND_ARCHITECTURE.md` for the unchanged Command boundary and the new downstream operator-model consumer;
- `docs/architecture/CANONICAL_REPORT_ENGINE.md` for the unchanged Report authority and downstream summary boundary;
- `docs/PRODUCT_ARCHITECTURE.md` for the shared operator-model layer before replaceable interfaces;
- `docs/FUNCTIONAL_SPECIFICATION.md` for beginner-visible status semantics without claiming an interface or operational persistence;
- `ai/core/04_ARCHITECTURE.md`;
- `ai/core/05_SYSTEM_MAP.md`;
- `ai/core/07_ENGINEERING_HISTORY.md`;
- `ai/core/13_ROADMAP.md`;
- `README.md`;
- `ai/prompts/032_CURRENT_TASK.md` during active implementation;
- `ai/history/032_2026-08-07_canonical-operator-presentation-model.md`;
- `ai/archive_prompts/032_2026-08-07_canonical-operator-presentation-model.md` at successful closure.

Runtime, Runtime Service, Alert, Notification, Scheduler, Configuration, CLI user guides, installation, monitoring, provider, and security documents shall change only if implementation proves a direct boundary clarification is necessary. Every actual update and every justified omission shall be recorded in Task 032 history.

Documentation shall state the permanent flow:

`Canonical Engineering and Operational Data -> Canonical Operator Presentation Model -> Replaceable Interface`.

It shall also state that the model is a read-only projection, not a monitor or persistent current-state database; that service status is only as current as its explicit observation; that missing inputs remain visible; that advanced users retain Command Definition/Execution and structured canonical contracts; and that Console, CLI integration, API, Dashboard, persistence/recovery, monitoring, production delivery, installation, and release remain separately governed.



## Completion Criteria


Task 032 is complete only when:

- one Canonical Operator Presentation Model 1.0 exists and is the sole shared cross-domain operator-overview contract;
- it deterministically answers health, attention, change, Alert, Guardian, freshness/completeness, and recommended-next-step questions from explicit validated canonical inputs;
- amateur-facing summaries require no understanding of Drift Records, Rule Evaluation, Policy internals, Runtime contracts, or canonical IDs, while stable source references preserve advanced drill-down and machine composition;
- exact ownership and precedence rules prevent re-evaluation or contradiction of Comparison, Drift, Health, Rule, Policy, Report, Alert, Runtime, Runtime Service, Command, and Pipeline semantics;
- missing, stale, partial, unsupported, unavailable, invalid, and not-observed evidence remain explicit and can never be silently reported as healthy;
- recommendations are closed, localization-ready, read-only next-step tokens and never grant remediation authority;
- equivalent valid explicit inputs produce byte-identical bounded output, and strict validation rejects tampering, incompatible versions, uncorrelated bundles, privacy violations, and resource-limit breaches;
- future CLI, Interactive Terminal Console, REST API, and Web Dashboard can consume the same model without engine access or interface-specific reinterpretation;
- current CLI, Command profiles, advanced composition, Pipeline, terminal/JSON output, Report, Alert, Runtime, Runtime Service, and manual workflows remain compatible;
- no Console, TUI, CLI integration, API, Dashboard, monitor, process probe, persistence, restart recovery, provider, configuration activation, installation, remediation, AI, packaging, deployment, or release behavior exists;
- the Release Minimalism record confirms this model is the smallest mandatory layer preventing duplicated future interface semantics and does not absorb later Version 1.0 operational gates;
- all focused, build, full test, race, vet, format, Framework, Builder, lifecycle, diversion, determinism, mapping, boundary, privacy, localization, resource, regression, documentation, permission/ACL, Git-diff, snapshot, and rollback validations pass;
- rollback remains proportional, collision-aware, verified, and preserves unrelated Owner-owned work;
- no dependency installation, staging, commit, push, branch, tag, package, deployment, release, or infrastructure mutation occurred;
- the completed Task 032 prompt/history are archived without a successor and `bin/job --check` confirms canonical idle with Task 032 as the unique latest completed baseline.

A valid result is `complete`, `complete with disclosed limitations`, or `blocked`. Completion may not be claimed while any mandatory model contract, ownership mapping, precedence rule, uncertainty behavior, source validation, determinism, interface-independence, privacy/localization/resource rule, test, documentation, rollback, or lifecycle gate remains unresolved.



## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-07 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
