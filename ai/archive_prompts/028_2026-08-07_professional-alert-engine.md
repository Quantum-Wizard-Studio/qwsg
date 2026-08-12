# Current Engineering Task 028: Professional Alert Engine

## Task Metadata

- Task ID: `028`
- Task slug: `professional-alert-engine`
- Status: `complete — all Task 028 gates passed`
- Date opened: `2026-08-07` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Professional Alert Engine


## Objective


Establish the permanent Canonical Professional Alert Engine of Quantum Wizard Server Guardian as a pure deterministic decision boundary.

Task 028 shall implement versioned Alert contracts and a side-effect-free engine that determines when a canonical alert exists from explicit validated canonical evidence, prior Alert lifecycle facts, explicit clock observations, and explicit bounded control inputs. It shall produce immutable Canonical Alert Records and proposed lifecycle state without delivering, persisting, displaying, transmitting, or acting on alerts.

The engine shall define stable alert identity, severity, category, source, event and observation timestamps, lifecycle state, acknowledgement, suppression, deduplication, expiration, correlation, recovery, and evidence-reference semantics. Equivalent valid inputs, including the same explicit time observation, shall produce byte-identical canonical outputs.

Task 028 ends when the pure Alert Engine, its public contracts, focused integration, architecture documentation, tests, rollback evidence, and lifecycle records are complete. Notification delivery and every operational adapter remain later tasks.



## Scope


Task 028 shall design and implement the Canonical Professional Alert Engine as a bounded pure decision layer over existing canonical outputs.

The selected boundary is engine-only. No persistence adapter, delivery adapter, command adapter, runtime loop, daemon, service, monitoring process, or presentation adapter is authorized.

The task scope includes:

- defining Canonical Alert Model 1.0;
- defining Alert Evaluation Input 1.0 and Alert Evaluation Result 1.0;
- defining immutable Canonical Alert Record 1.0;
- defining the minimal previous Alert State 1.0 required for deterministic lifecycle continuity, supplied to and returned by the pure engine as data only;
- defining immutable Alert Acknowledgement Record 1.0 and bounded Alert Suppression Window 1.0 as explicit control inputs, without implementing their creation UI, persistence, configuration activation, authorization transport, or administration workflow;
- consuming only validated canonical Health Results, Rule Results, Policy Results, Scheduler Events, Effective Configuration, Canonical Reports, previous Alert state, acknowledgements, suppression windows, and an explicit clock observation;
- permitting a bounded subset of those source families in one evaluation while rejecting unsupported, contradictory, ambiguous, duplicate, stale, tampered, or incompatible input combinations;
- defining exact source-adapter ownership so downstream evidence does not duplicate upstream alerts: Policy-backed Rule evidence supersedes the corresponding direct Rule candidate, Rule evidence supersedes its referenced direct Health candidate, and Report items do not recreate candidates already represented by their canonical Policy or Rule sources;
- allowing Canonical Reports only as validated presentation-contract evidence or report-level completeness/failure evidence, never as authority to reinterpret Health, Rule, or Policy semantics;
- allowing Effective Configuration only as validated identity, provenance, applicability, and bounded policy context already present in version 1.0, never as a place to invent absent alert, maintenance, acknowledgement, or notification configuration fields;
- consuming Scheduler Events as scheduler-operational evidence while preserving Scheduler identity, event kind, timestamp, and source references without re-evaluating schedule due-time or execution behavior;
- defining Alert severity taxonomy 1.0 independently from Health status and Policy outcome, with explicit deterministic mappings and no probabilistic scoring;
- defining at least `informational`, `warning`, `critical`, and `emergency` alert severities, plus an explicit indeterminate decision where evidence cannot justify an alert severity;
- defining Alert category taxonomy 1.0 with stable machine identifiers for engineering condition, rule match, policy governance, scheduler operation, report completeness, evidence loss, and Alert-engine input failure where safely representable without recursion;
- defining exact candidate rules for confirmed unhealthy or unknown evidence, Rule match and evaluation failure, Policy escalated/conflict/indeterminate outcomes, applicable Scheduler failure or interruption events, and Canonical Report incompleteness;
- preserving `accepted`, `observe`, `suppressed`, `escalated`, `indeterminate`, `not_applicable`, and `conflict` as Policy facts; in particular, Policy `suppressed` shall not delete evidence and shall not automatically become a time-bounded operational suppression window;
- defining alert identity at three levels: stable condition key, lifecycle/correlation identity, and immutable record identity;
- deriving identities from canonical source identity, subject/scope, category, condition generation, lifecycle event, and canonical content, never from map order, random values, process identity, locale, or ambient time;
- defining source taxonomy and source references containing exact schema, version, record identity, evidence identity, and privacy-bounded scope;
- defining event time, observation time, evaluation time, acknowledgement time, suppression start/end, expiration time, and recovery time as distinct explicit UTC fields with normative meanings;
- using only supplied timestamps; the engine shall not read the wall clock or monotonic clock;
- defining lifecycle states and events for at least candidate, active, acknowledged, suppressed, expired, resolved/recovered, and indeterminate outcomes without mutating the underlying Health, Rule, Policy, Scheduler, Configuration, or Report evidence;
- defining acknowledgement as operator-awareness evidence that never changes measured condition, severity, recovery, or escalation and never grants remediation authority;
- binding acknowledgement to a stable alert lifecycle/correlation identity, actor identity, time, optional localization-ready note token or bounded note reference, and source authority evidence;
- rejecting acknowledgement of a nonexistent, resolved, expired, mismatched, or future alert when the contract cannot justify it;
- defining suppression as a bounded, explicit, scope-matched decision with stable identity, start, end, actor/authority reference, reason token, severity applicability, and emergency-suppression flag;
- forbidding open-ended suppression in Alert Model 1.0;
- continuing alert condition evaluation during suppression, recording the matching suppression reason, and producing suppressed Alert Records rather than deleting or falsifying the condition;
- requiring an explicit emergency-suppression choice for any suppression window that includes emergency alerts;
- defining maintenance as a suppression category only; Task 028 shall evaluate supplied bounded maintenance windows but shall not create, store, activate, discover, schedule, or administer them;
- defining maintenance-end evaluation so one current-status Alert Record may be proposed for a still-active condition and suppressed historical transitions are not replayed;
- defining deduplication from stable condition and lifecycle identities so equivalent unchanged evidence does not create repeated entry alerts;
- defining deterministic escalation, de-escalation, evidence-update, acknowledgement, suppression-entry, suppression-exit, maintenance-end status, reminder eligibility, expiration, full recovery, and recurrence behavior;
- defining reminder eligibility only as an Alert decision and record; no reminder scheduling or delivery is authorized;
- defining exact default fixed Model 1.0 reminder behavior only where no missing configuration semantics are required, including a bounded emergency reminder interval evaluated from explicit prior Alert state and explicit time;
- defining expiration separately from recovery: expiration ends an Alert record's actionability because its evidence lifetime elapsed, while recovery requires canonical evidence that the condition ended;
- ensuring stale or missing evidence never fabricates recovery and instead yields the defined evidence-loss, expired, or indeterminate outcome;
- defining recurrence after full recovery as a new lifecycle/correlation identity while preserving a stable condition key for history correlation;
- defining correlation through exact canonical subject/scope and evidence relationships, not similarity, prose, heuristics, host queries, or AI;
- defining deterministic ordering, aggregation, collision handling, and maximum input/output cardinalities;
- producing proposed next Alert State only; callers remain responsible for any later persistence under a separately authorized task;
- defining byte-stable canonical JSON, strict decoding where exposed, content-derived SHA-256 identities, complete version information, and validation that recomputes identities and ordering;
- defining fail-closed behavior for unsupported schemas, versions, taxonomies, mappings, lifecycle transitions, time relationships, references, and bounds;
- preserving privacy by copying only canonical identifiers, localization-ready tokens, bounded status facts, and evidence references rather than raw host evidence, arbitrary report prose, secrets, credentials, or private configuration values;
- remaining offline, local, deterministic, side-effect-free, presentation-independent, delivery-independent, remediation-independent, and AI-independent;
- integrating only through direct typed consumption of existing public engine contracts; Alert shall not be inserted into the Canonical Command/Pipeline stage order unless architecture review proves a narrowly additive typed adapter is necessary and does not change existing Command 1.0 semantics;
- preserving all existing CLI, Command, Pipeline, Scheduler, Configuration, Policy, Report, and manual-engine behavior;
- adding focused unit, contract, lifecycle, mapping, timestamp, acknowledgement, suppression, maintenance, deduplication, expiration, correlation, recurrence, determinism, canonical-serialization, bounds, privacy, compatibility, source-boundary, and regression tests;
- creating permanent Canonical Alert Engine architecture documentation;
- updating only directly affected permanent architecture, system map, roadmap, engineering history, README, and lifecycle documentation.

The Alert Engine decides whether an alert exists and what its canonical lifecycle record is. It performs no action based on that decision.



## Out of Scope


Task 028 shall not implement:

- email delivery or e-mail transport;
- SMS;
- Discord;
- Telegram;
- Slack;
- webhooks;
- push notifications;
- desktop notifications;
- any notification channel, notification router, recipient model, template renderer, delivery queue, delivery retry, delivery audit, channel-health incident, SMTP selection, or transport recovery;
- REST API, HTTP listener, application API, SDK, public integration endpoint, or remote-control surface;
- Dashboard, Web UI, Terminal UI, interactive incident view, notification view, or presentation adapter;
- CLI alert commands, incident commands, acknowledgement commands, maintenance commands, or user-facing configuration syntax unless a separately reviewed compatibility-only change is strictly necessary to keep existing builds valid; no new runtime behavior is authorized;
- monitoring daemon, polling loop, resident agent, background worker, timer service, service manager, systemd unit, automatic startup, watchdog, or daemon lifecycle;
- alert persistence, incident database, notification store, delivery store, maintenance store, file activation, hot reload, state migration, retention cleanup, or a general-purpose product-state platform;
- host remediation, automatic repair, action execution, shell execution, process execution, plugin execution, privilege change, or infrastructure mutation;
- AI reasoning, probabilistic classification, machine learning, semantic similarity, generated recommendations, or external AI integration;
- host collection, Snapshot creation, Snapshot comparison, Drift classification, Health evaluation, Rule matching, Policy interpretation, Report generation, Configuration resolution, Scheduler due-time calculation, Command planning, Pipeline orchestration, or presentation rendering;
- redefining existing Health status, Rule outcome, Policy outcome, Scheduler decision/event, Configuration precedence, Report completeness, Command, or Pipeline semantics;
- treating Policy `escalated` as delivery authority, Policy `suppressed` as deletion, acknowledgement as health change, suppression as recovery, expiration as recovery, or an alert as remediation authority;
- creating or activating configuration fields absent from Effective Configuration 1.0;
- unbounded or open-ended suppression, unbounded reminders, unbounded lifecycle history, recursive Alert-engine health alerts, or implicit defaults that depend on deployment environment;
- remote agents, fleet correlation, multi-host coordination, tenancy, licensing, telemetry, cloud services, network calls, release, deployment, installation, update, repair, or removal.

Task 028 shall not install dependencies, alter infrastructure, stage, commit, push, fetch, create branches or tags, or begin any later delivery task.



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


Task-preparation preflight, completed before Owner review, shall establish:

- the repository is the verified QWSG root and Framework 1.x is authoritative;
- `bin/job --check` reports canonical idle;
- Task 027 is the highest completed and archived production task with one matching complete history record;
- no Task 028 prompt, history, or archive destination exists;
- Tasks 021, 025, 026, and 027 and the Canonical Report, Command, Product Architecture, Functional Specification, Roadmap, and System Map sources establish the required upstream contracts;
- the Scheduler is intentionally one-cycle and no daemon or notification delivery prerequisite is required for a pure Alert decision engine;
- Effective Configuration 1.0 does not yet contain alert, acknowledgement, maintenance, or notification fields, so Task 028 shall use explicit typed bounded Alert control inputs rather than silently extending Configuration;
- no mandatory architectural prerequisite remains for the selected engine-only boundary;
- existing unstaged and untracked Owner work is preserved and is not Task 028 preparation output except for the explicitly authorized `current-task-job.txt` replacement;
- the Builder source contains no placeholders, unresolved mandatory sections, approval token, installation instruction, or execution claim.

Before Builder installation, a later authorized installation phase shall verify:

- canonical idle still holds and Task 027 remains the latest completed baseline;
- no Task 028 lifecycle target exists;
- the Builder interface, versioned project configuration, input contract, target paths, and validation commands are unchanged from preparation assumptions;
- the reviewed source hash and generated deterministic field hashes are unchanged;
- structured Builder input is complete and passes `task-builder.sh --check-input` only after the Owner supplies the exact separate approval token;
- installation snapshot and exact rollback coverage exist before any lifecycle mutation.

After Builder installation and before any Task 028 implementation or documentation target changes, the Builder agent shall:

- verify QWSG repository identity, configured root, branch `main`, HEAD, configured remote, canonical remote URL, and local-to-remote relationship;
- verify Task 028 is the sole active approved prompt with exactly one matching history record and Task 027 is the latest completed implementation baseline;
- verify the installed prompt and history match Task ID, slug, title, authority, approval, and lifecycle state;
- verify all required upstream implementations, contracts, versions, public types, tests, and documents are present and consistent;
- run Framework, lifecycle, test-task, Go build, test, race, vet, and formatting baselines using configured writable cache paths;
- record complete Git status and distinguish all pre-existing Owner work from Task 028 targets;
- identify every intended Task 028 target, including new files and every existing documentation or integration file, before snapshot creation;
- verify no destination collision, unexplained repository divergence, dependency change, unsupported platform assumption, permission change, lifecycle inconsistency, or architecture conflict affects correctness or rollback;
- stop on material differences affecting authority, safety, scope, compatibility, rollback, or implementation correctness.

The valid pre-install state is canonical idle. The valid post-install state is one active approved Task 028. These checks must not be confused or applied in the wrong lifecycle phase.



## Snapshot Requirements


Task preparation shall not create an implementation snapshot because no implementation or lifecycle installation is authorized. The Builder source itself is Owner-owned preparation material and remains outside the active-task lifecycle.

Before a separately authorized Builder installation:

- create a proportional rollback snapshot covering the exact lifecycle destinations the Builder may create and the source/hash evidence required to prove the reviewed input;
- verify recorded absence of Task 028 prompt, history, and archive targets;
- rely on the Builder's transactional no-clobber installation and automatic restoration, then validate canonical lifecycle state.

Before Task 028 implementation changes:

- create a unique timestamped rollback-capable snapshot outside the repository for every existing implementation, test, integration, and documentation target;
- record verified absence for every new Alert package, test, architecture, prompt archive, and other new target;
- capture repository identity, branch, HEAD, remotes, lifecycle state, complete Git status, exact target inventory, ownership, permissions, ACLs, and baseline validation evidence;
- preserve exact pre-task content, including pre-existing unstaged Owner changes, without using HEAD as a substitute for working-tree truth;
- include a deterministic manifest, SHA-256 checksums, archive inventory, and exact guarded restore procedure;
- verify every checksum, payload, listing, absence record, and restore precondition before implementation;
- define collision checks that refuse to overwrite later Owner work;
- retain the snapshot through completion and Owner acceptance;
- keep private source payloads and host-specific paths outside Git.

Snapshot scope must be proportional to the actual Alert target set. Broad repository archives, runtime host-state capture, unrelated backups, and destructive recovery shortcuts are prohibited.



## Risk Assessment


Primary risks:

- conflating Alert decision with notification delivery, incident persistence, monitoring, daemon lifecycle, presentation, or remediation;
- inventing alert-policy, maintenance, or notification fields inside Effective Configuration 1.0;
- treating Policy `escalated` as an automatic action or Policy `suppressed` as operational deletion;
- emitting duplicate alerts from Health, Rule, Policy, and Report representations of the same evidence;
- mapping Health status or Policy outcomes to Alert severity without an explicit fixed versioned rule;
- using ambient time and making reminder, expiration, suppression, or recovery results nondeterministic;
- allowing acknowledgement or suppression to falsify source health, block legitimate escalation/recovery facts, or imply remediation authority;
- confusing expiration, evidence loss, and recovery;
- reopening a recovered lifecycle instead of creating a recurrence, or creating a new lifecycle for unchanged evidence;
- generating recursive Alert-engine health alerts from malformed Alert input;
- copying raw host evidence, report prose, configuration secrets, actor notes, or sensitive paths into Alert Records;
- unbounded lifecycle state, correlation groups, acknowledgements, suppression windows, evidence references, or output records;
- altering Command/Pipeline stage semantics or existing CLI behavior to integrate Alert prematurely;
- tests depending on real time, map order, external services, environment state, or network availability;
- rollback overwriting the pre-existing unstaged Task 025–027 and Owner-owned work.

Mitigations:

- keep `internal/alert` pure and restrict imports to required canonical model packages and the standard library;
- use explicit typed evaluation inputs, previous state, controls, and clock observation; read no environment, filesystem, network, or process state;
- publish exact source precedence and de-duplication rules and reject ambiguous mixed sources;
- preserve every upstream outcome unchanged and define Alert mapping as its own versioned taxonomy;
- distinguish condition key, lifecycle identity, and immutable record identity;
- model acknowledgement, suppression, maintenance, expiration, recovery, and recurrence as validated lifecycle facts;
- require bounded time windows and explicit emergency suppression;
- emit privacy-bounded identifiers, tokens, and references only;
- fail closed on unsupported versions, invalid time order, tampering, ambiguity, and resource excess;
- use virtual explicit timestamps and deterministic fixtures for all temporal tests;
- keep Command/Pipeline integration absent unless a typed additive need is proven during architecture review;
- snapshot exact working-tree targets and use collision-aware bounded restoration.

No risk requires an Owner scope change. Delivery-channel selection, runtime persistence, daemon architecture, operational authorization, and end-user workflows remain consciously deferred.



## Planned Work


### Phase 1 — Task preparation

1. Verify canonical idle and Task 027 completion without installing or activating Task 028.
2. Review every authoritative source and confirm the Alert Engine is the next correct Professional decision capability.
3. Confirm Health, Rule, Policy, Configuration, Scheduler, Report, Command, Product Architecture, Functional Specification, Roadmap, and System Map provide sufficient upstream contracts.
4. Select the engine-only boundary and document why daemon and notification delivery are not prerequisites.
5. Resolve the Effective Configuration gap by specifying explicit Alert-owned control inputs without changing Configuration 1.0.
6. Simulate Owner, Builder, installation, execution, validation, completion, archive, and canonical-idle transitions and encode every preventable stop condition in this source.
7. Prepare and self-review `current-task-job.txt` as data only, with no placeholders, approval token, installation, implementation, lifecycle mutation, staging, commit, or push.
8. Stop at Builder-ready preparation and report whether any engineering-scope Owner input remains.

### Phase 2 — Separately authorized Builder installation

1. Revalidate lifecycle, repository, Builder interface, source hash, target absence, and installation rollback assumptions.
2. Transform the reviewed source into deterministic Builder fields without executing its prose.
3. Present the complete generated task for Owner review.
4. Obtain the exact separate Builder approval token from the Owner; never infer or manufacture it from this task source.
5. Run Builder input validation and install the Task 028 prompt/history pair transactionally.
6. Validate the sole active task, matching history, approval state, permissions, and rollback result.
7. Stop with `APPROVED AND READY FOR IMPLEMENTATION`; do not execute Task 028 during installation.

### Phase 3 — Separately authorized task execution

1. Start only through the canonical `job` workflow and read all governing and task-specific sources.
2. Verify the exact starting state and create the proportional verified implementation snapshot.
3. Inspect public canonical types and finalize Alert source, taxonomy, identity, lifecycle, control, temporal, privacy, version, and compatibility contracts before code changes.
4. Define source ownership, candidate generation, ambiguity rejection, severity/category mappings, and evidence traceability.
5. Define condition key, lifecycle identity, record identity, correlation, deduplication, escalation, de-escalation, recovery, recurrence, reminder, expiration, and indeterminate behavior.
6. Define acknowledgement and bounded suppression/maintenance input contracts and exact lifecycle effects without persistence or administration.
7. Implement the pure deterministic engine and strict validation/canonical serialization in a new bounded `internal/alert` package.
8. Add typed adapters only for the selected existing canonical outputs and keep every upstream engine's meaning immutable.
9. Add focused tests for every taxonomy, mapping, lifecycle transition, time boundary, ambiguity, limit, determinism, privacy, compatibility, and exclusion.
10. Create `docs/architecture/CANONICAL_ALERT_ENGINE.md` and update only directly affected permanent documents.
11. Run all mandatory focused and repository-wide verification, inspect exact diffs and permissions, and validate rollback evidence.
12. Record implementation, decisions, failures/corrections, exact evidence, limitations, rollback, and Git state in Task 028 history.
13. Complete and archive Task 028 only after every gate passes; create no successor and return to canonical idle.



## Rollback Plan


Preparation rollback is limited to `current-task-job.txt`: preserve any later Owner edits, compare the file with the preparation baseline, and replace it only with explicit Owner direction. No repository-wide rollback is permitted.

Builder-installation rollback shall use only the Builder's verified transactional restoration and exact lifecycle-target snapshot. It shall remove only a Task 028 prompt/history pair proven to have been installed by that failed transaction and restore the prior canonical idle state. If exact restoration cannot be proven, stop and request Owner direction.

Implementation rollback shall:

- stop further mutation and preserve truthful failure and validation evidence;
- verify the Task 028 snapshot manifest, checksums, archive listing, payload readability, absence records, permissions, and guarded restore instructions;
- compare each affected target with both the snapshot and current working tree;
- refuse to overwrite later or unrelated Owner work;
- restore only verified pre-existing Task 028 targets from the bounded snapshot;
- remove only verified Task 028-created files whose pre-task absence was recorded and whose current content contains no later Owner work;
- preserve lifecycle and history truth rather than rewriting failed or completed facts;
- rerun Framework, lifecycle, test-task, Go build/test/race/vet/format, focused Alert, documentation, source-boundary, permission, ACL, Git-diff, and snapshot-integrity checks after restoration;
- report exact resulting state and every unresolved condition.

Broad `git reset`, `git checkout`, `git restore`, `git clean`, wildcard deletion, repository-wide extraction, ambiguous file deletion, or removal of owner-owned untracked content is prohibited.



## Deliverables


- pure Canonical Professional Alert Engine implementation;
- Canonical Alert Model 1.0;
- Alert Evaluation Input 1.0;
- Alert Evaluation Result 1.0;
- Canonical Alert Record 1.0;
- previous/proposed Alert State 1.0 contract;
- Alert Acknowledgement Record 1.0;
- bounded Alert Suppression Window 1.0, including maintenance category semantics;
- exact Alert severity, category, source, lifecycle, decision, and event taxonomies;
- deterministic Health, Rule, Policy, Scheduler Event, Effective Configuration, and Canonical Report source-boundary behavior for the selected supported source set;
- stable condition, lifecycle/correlation, and immutable record identities;
- explicit timestamp, acknowledgement, suppression, maintenance-end, deduplication, reminder, expiration, recovery, and recurrence semantics;
- exact evidence references, canonical ordering, validation, versioning, compatibility, resource limits, privacy behavior, and canonical JSON;
- focused unit, contract, integration, temporal, lifecycle, determinism, privacy, source-boundary, and regression tests;
- `docs/architecture/CANONICAL_ALERT_ENGINE.md`;
- directly affected canonical documentation updates;
- complete Task 028 history and verified rollback evidence;
- completed prompt archive and canonical idle state after execution.

No delivery adapter, notification, daemon, persistence adapter, API, dashboard, remediation, or AI artifact is a Task 028 deliverable.



## Verification


Builder and lifecycle verification shall prove:

- the source has the correct Task ID, slug, title, authority, language, canonical sections, and no placeholders or unresolved mandatory content;
- task preparation, Builder installation, and task execution are explicitly separate;
- no approval token is embedded or inferred;
- pre-install canonical idle and post-install sole-active-task states are checked in their correct phases;
- deterministic Builder input passes `task-builder.sh --check-input` after separate Owner approval;
- `bin/job --check`, lifecycle validation, and prompt/history identity validation pass after installation;
- completion and archive return the repository to canonical idle with Task 028 as the unique latest completed archive/history pair.

Implementation verification shall include:

- focused `internal/alert` tests;
- repository-wide build and Go tests;
- repository-wide race tests using configured writable caches;
- vet and complete formatting checks;
- Framework 1.x configured validations and engineering test suites;
- active-task, lifecycle, idle-closure, Builder, and test-task validation;
- golden or equivalent contract tests for every public Alert 1.0 record;
- byte-identical canonical JSON and identity tests across equivalent input ordering;
- tests for every supported source family, taxonomy value, severity mapping, category mapping, and evidence-reference path;
- tests proving direct Health, Rule, Policy, and Report representations cannot create duplicate alerts for the same evidence;
- tests for empty, malformed, duplicate, contradictory, tampered, oversized, stale, future-dated, unsupported-version, unsupported-taxonomy, and ambiguous sources;
- tests for initial entry, unchanged evidence, evidence update, escalation, de-escalation, acknowledgement, suppression entry/exit, emergency suppression, maintenance end, reminder eligibility, expiration, evidence loss, full recovery, and recurrence;
- tests proving acknowledgement does not change severity or recovery, suppression does not delete evidence, expiration does not fabricate recovery, and Policy outcomes are not operational actions;
- tests proving the engine reads no clock, environment, filesystem, network, process, random source, mutable global state, or host evidence;
- tests proving no delivery channel, recipient routing, notification retry, daemon, monitoring loop, persistence, API, Dashboard, remediation, shell, AI, or host mutation exists;
- tests proving existing Command, Pipeline, Scheduler, Configuration, Policy, Report, CLI, and manual workflows remain compatible;
- privacy and secret-disclosure audits;
- deterministic resource-bound and failure-isolation tests;
- architecture terminology and cross-document consistency checks;
- exact target, ownership, permission, ACL, unstaged/staged state, and Git diff review;
- `git diff --check` and `git diff --cached --check`;
- verified snapshot checksums, payloads, absence records, collision guards, and bounded rollback procedure;
- confirmation that nothing was staged, committed, pushed, installed, deployed, or released.

Verification must remain proportional. No external service, live notification transport, real daemon, real host mutation, network access, or sleep-based timing is required or permitted.



## Documentation Updates


Task 028 execution shall create or update only directly affected records, expected to include:

- `docs/architecture/CANONICAL_ALERT_ENGINE.md`;
- `docs/PRODUCT_ARCHITECTURE.md` only where implemented-status wording or permanent Alert boundaries require correction;
- `docs/FUNCTIONAL_SPECIFICATION.md` only where traceability notes must distinguish implemented Alert decision from deferred delivery;
- `docs/architecture/CANONICAL_HEALTH_ENGINE.md`;
- `docs/architecture/CANONICAL_RULE_ENGINE.md`;
- `docs/architecture/CANONICAL_POLICY_ENGINE.md`;
- `docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md`;
- `docs/architecture/CANONICAL_SCHEDULER.md`;
- `docs/architecture/CANONICAL_REPORT_ENGINE.md`;
- `docs/architecture/CANONICAL_COMMAND_ARCHITECTURE.md` only if its future-integration boundary needs a factual additive reference;
- `ai/core/04_ARCHITECTURE.md`;
- `ai/core/05_SYSTEM_MAP.md`;
- `ai/core/07_ENGINEERING_HISTORY.md`;
- `ai/core/13_ROADMAP.md`;
- `README.md`;
- `ai/prompts/028_CURRENT_TASK.md` during active execution;
- `ai/history/028_2026-08-07_professional-alert-engine.md`;
- `ai/archive_prompts/028_2026-08-07_professional-alert-engine.md` at successful idle closure.

Every actual documentation change and every justified omission from this expected list shall be recorded in Task 028 history. Documentation must state that the Alert Engine decides only and that notification delivery, persistence, monitoring daemon, interfaces, remediation, and AI remain unimplemented.



## Completion Criteria


Task 028 is complete only when:

- the pure Canonical Professional Alert Engine exists and determines Alert conditions solely from explicit validated canonical inputs;
- all selected Alert Model 1.0 contracts, taxonomies, mappings, identities, lifecycle semantics, controls, time semantics, evidence references, bounds, and compatibility rules are implemented and documented;
- equivalent valid inputs including the explicit clock observation produce byte-identical ordered results and identities;
- acknowledgement, suppression, maintenance, deduplication, reminder eligibility, expiration, recovery, correlation, and recurrence behavior is exact and covered by tests;
- source ownership prevents competing Health, Rule, Policy, Report, Scheduler, or Alert semantics and prevents duplicate alerts from the same canonical evidence;
- Effective Configuration 1.0 is consumed without silently extending or redefining it;
- existing upstream evidence remains immutable and traceable;
- the engine performs no delivery, notification routing, persistence, daemon, monitoring loop, API, Dashboard, remediation, process execution, network operation, host mutation, AI, or probabilistic work;
- existing CLI, Command, Pipeline, Scheduler, Configuration, Policy, Report, and manual operation remain compatible;
- all focused and repository-wide mandatory verification passes;
- rollback remains complete, bounded, collision-aware, and verified;
- documentation and Task 028 history are complete, accurate, and distinguish implemented behavior from future work;
- no dependency, installation, deployment, release, staging, commit, push, branch, or tag operation occurred;
- the completed prompt is archived without creating a successor and `bin/job --check` confirms canonical idle with Task 028 as the latest completed task.

A valid final result is `complete`, `complete with disclosed limitations`, or `blocked`. Completion may not be claimed while any mandatory contract, boundary, test, documentation, rollback, or lifecycle gate remains unresolved.



## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-07 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
