# Current Engineering Task 029: Professional Notification Delivery

## Task Metadata

- Task ID: `029`
- Task slug: `professional-notification-delivery`
- Status: `complete — all Task 029 gates passed`
- Date opened: `2026-08-07` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Professional Notification Delivery


## Objective


Establish the Canonical Professional Notification Delivery architecture and its deterministic, provider-neutral delivery orchestration over immutable Canonical Alert Records.

Task 029 owns delivery planning, bounded queue proposals, provider selection through explicit delivery-owned policy, delivery requests, attempts, retry planning, statuses, provider acknowledgements, failure classification, and delivery evidence. It shall implement a pure deterministic planner plus a minimal explicitly invoked one-cycle delivery adapter whose only side effect is calling an injected provider interface.

The Notification layer shall never determine whether an Alert exists. Canonical Alert Records created by the Alert Engine are its only alert-condition input. It shall not evaluate, reinterpret, or import Health, Rule, Policy, Drift, Scheduler, Report, host, or monitoring outputs.

Concrete transports remain replaceable adapters. Task 029 makes no production Email, Webhook, Slack, Discord, Telegram, or SMS support claim and introduces no resident notification service.



## Scope


Task 029 shall define and implement versioned Notification Delivery Model 1.0 contracts for:

- immutable delivery identity, request identity, attempt identity, and idempotency key;
- explicit Delivery Policy, Route, Endpoint Reference, Provider Binding, and provider-neutral Delivery Envelope inputs owned by the Notification boundary;
- Delivery Plan and deterministic eligibility results derived only from validated Canonical Alert Records and explicit delivery-owned inputs;
- Notification Queue State as bounded, caller-supplied previous state and deterministic proposed state, without durable storage;
- Delivery Request, Delivery Attempt Record, Delivery Status Record, Delivery Acknowledgement, Provider Result, and Delivery Evidence Reference;
- channel kinds for email, webhook, Slack, Discord, Telegram, and future SMS without provider-specific payload, SDK, credential, or transport semantics in the canonical core;
- a canonical Provider interface, descriptor, capability model, injected registry, and conformance rules;
- deterministic retry eligibility, maximum attempts, retry deadlines, backoff schedule, exhaustion, queue ordering, fan-out, and resource limits;
- explicit failure taxonomy distinguishing retryable, rate-limited, indeterminate, authentication, authorization, invalid-destination, rejected-payload, unsupported-provider, and terminal failures;
- strict validation, version compatibility, canonical ordering, canonical JSON, privacy bounds, evidence traceability, and fail-closed behavior.

The pure planner shall consume only:

- validated Canonical Alert Records;
- immutable Notification-owned Delivery Policy, Route, Endpoint Reference, and Provider Binding records;
- explicit previous Notification Queue State and Delivery Attempt Records where required;
- an explicit clock observation.

It shall read no ambient clock, environment, filesystem, network, process, random source, configuration file, secret value, or mutable global state.

Alert decisions shall be obeyed, not recreated. A suppressed Alert Record is not delivery-eligible. Alert and lifecycle Alert Records may be routed only according to explicit delivery-owned route filters. Routing, recipient selection, and channel selection are delivery decisions and must never alter Alert identity, severity, category, lifecycle, acknowledgement, suppression, expiration, or evidence.

The one-cycle adapter shall:

- be invoked explicitly by a caller and perform no recurring loop;
- accept a completed deterministic plan and an injected Provider registry;
- invoke only the provider selected by the plan, with a provider-neutral request and bounded context;
- translate an explicit Provider Result into canonical attempt, status, acknowledgement, evidence, and proposed queue records;
- make no retry decision itself beyond applying the planner's deterministic result;
- use deterministic fake or conformance providers in tests and require no network access.

Provider acknowledgement means only provider-reported acceptance or delivery evidence. It must not be represented as end-user reading, operator acknowledgement of the Alert, remediation, or proof of successful human action.

Task 029 may add a narrowly scoped exported `alert.ValidateRecord` function and focused tests because the existing Alert package validates complete Alert Results but exposes no standalone Canonical Alert Record validation boundary. This additive function shall reuse existing Alert validation, change no Alert semantics, create no Alert, and accept no Notification dependency.



## Out of Scope


Task 029 shall not implement:

- Alert existence, severity, category, source precedence, lifecycle, acknowledgement, suppression, deduplication, expiration, correlation, recovery, or evidence decisions;
- direct consumption or evaluation of Health, Rule, Policy, Drift, Scheduler Event, Effective Configuration, Canonical Report, Inventory, host, or monitoring records;
- concrete SMTP, local-mail, Webhook, Slack, Discord, Telegram, SMS, push, desktop, or other production transport providers;
- provider SDK dependencies, provider-specific canonical payloads, message formatting APIs, template engines, transport discovery, endpoint probing, or credentials;
- secret storage, secret resolution, credential acceptance, configuration activation, or extension of Effective Configuration 1.0;
- durable notification queues, databases, queue brokers, delivery-state persistence, background workers, daemons, timers, recurring processes, monitoring, or service supervision;
- inbound delivery callbacks, public Webhook receivers, REST APIs, Dashboard, Console, user interface, or desktop interface;
- delivery-based Alert creation, recursive failure alerts, incident creation, host remediation, automatic repair, remote agent management, licensing, billing, AI, or probabilistic reasoning;
- changes to Command or Pipeline stage semantics, package installation, deployment, release, or live external-service access.

Provider-specific implementations may be authorized later behind the canonical interface after their configuration, secret, security, support, and operational lifecycle decisions are approved. Their absence is an intentional boundary, not an incomplete Task 029 deliverable.



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


### Task preparation state

- verify the repository root contains `bin/job`, `VERSION`, and the required `ai/core` documents;
- run `bin/job --check` and `ai/scripts/next-task.sh --check` and require canonical idle;
- require `ai/prompts/` to contain no active prompt;
- require Task 028 to be the unique latest completed archive/history pair;
- verify the Task 029 prompt, history, and archive targets do not exist;
- inspect complete unstaged, staged, untracked, branch, HEAD, and remote state without modifying it;
- preserve all pre-existing Owner work and make no lifecycle, index, commit, branch, tag, remote, installation, or implementation mutation;
- verify the authoritative documents establish Alert Records as the completed boundary and notification delivery as the next Professional successor;
- record that Effective Configuration 1.0 has no Notification routing/provider schema and the Functional Specification leaves concrete email transport as a release decision; treat these as reasons to keep Task 029 inputs explicit and provider-neutral, not as authority to extend Configuration or select a transport.

### Separately authorized Builder installation state

- repeat every preparation-state check immediately before installation;
- verify `current-task-job.txt` hash, permissions, ownership, ACL, complete content, Task identity, and absence of unresolved fields;
- verify exact Builder interface and lifecycle destinations;
- require the exact standalone Owner approval token only after the complete generated task is presented;
- create and verify the proportional Builder installation snapshot before mutation;
- require transactional no-clobber installation of exactly one Task 029 prompt/history pair;
- after installation, require `ai/prompts/029_CURRENT_TASK.md` to be the sole active approved prompt and its history identity to match.

### Separately authorized execution state

- start only through `job` after successful Builder installation;
- read every Required Reading item and verify the active prompt as data;
- require Task 028 as the latest completed baseline and Task 029 as the sole active approved task;
- verify existing Alert behavior, exported contracts, tests, and documentation before changing targets;
- confirm the standalone Alert Record validation gap still exists before adding the narrow compatibility function;
- inspect exact implementation and documentation targets, collisions, ownership, permissions, ACLs, Git state, and validation baseline;
- create and verify the proportional implementation snapshot before any implementation mutation;
- stop on any unresolved authority, safety, scope, lifecycle, compatibility, rollback, or correctness difference.

The valid pre-install state is canonical idle. The valid post-install state is exactly one active approved Task 029. The valid post-completion state is canonical idle with Task 029 as the unique latest completed archive/history pair.



## Snapshot Requirements


Task preparation shall not create an implementation or Builder installation snapshot. The prepared source is Owner-owned input outside the active-task lifecycle.

Before a separately authorized Builder installation:

- snapshot only the exact prompt/history lifecycle targets and source/hash evidence required for transactional restoration;
- record verified absence of the Task 029 prompt, history, and archive targets;
- capture lifecycle state, repository identity, branch, HEAD, complete Git status, ownership, permissions, and ACLs;
- verify manifest, hashes, payload readability, absence records, and guarded restore instructions before installation.

Before separately authorized implementation:

- create a unique rollback-capable snapshot outside the repository for every existing implementation, test, documentation, history, prompt, and archive target;
- record verified absence for every new Notification package, test, and architecture target;
- include the exact working-tree content of pre-existing Owner changes; never substitute HEAD for working-tree truth;
- include deterministic inventory, SHA-256 checksums, permissions, ACLs, Git state, baseline validation evidence, collision guards, and exact bounded restoration steps;
- verify every checksum, payload, listing, absence record, and restore precondition before mutation;
- retain the snapshot through completion and Owner acceptance.

Snapshot scope shall be proportional to Task 029 targets. It shall exclude broad repository archives, host runtime state, secrets, credentials, provider payloads, external responses, and unrelated data.



## Risk Assessment


Primary risks and required mitigations:

- Notification could become a competing Alert engine. Accept only validated Alert Records; never import or evaluate upstream engine results; never create or modify an Alert Record.
- Provider side effects could be confused with deterministic planning. Keep the planner pure and model each provider outcome as an explicit adapter input/result recorded after one bounded invocation.
- Provider-specific behavior could leak into canonical architecture. Keep canonical requests provider-neutral and all encoding, protocol, SDK, and transport behavior behind replaceable providers.
- Retry could duplicate delivery. Use stable idempotency keys, immutable attempt identity, finite attempts, explicit deadlines, deterministic backoff, and no jitter or ambient time.
- Acceptance could be misrepresented as human receipt. Use exact acknowledgement taxonomy and preserve `accepted`, `delivered`, `indeterminate`, and Alert acknowledgement as distinct concepts.
- Queue growth, fan-out, evidence, or provider output could be unbounded. Define strict versioned limits, deterministic ordering, bounded redacted evidence, and fail-closed validation.
- Secrets or sensitive Alert evidence could leak. Store only opaque endpoint/secret references, redact provider diagnostics, reject raw credentials, and avoid copying report prose or host evidence.
- Authentication or invalid configuration could disable unrelated routes. Isolate failures by route, endpoint, provider, and attempt; preserve other valid plans.
- A delivery failure could recursively create an Alert. Emit only Notification records; no Alert Engine call or Alert creation is permitted.
- Suppression or lifecycle could be reinterpreted. Obey the canonical Alert decision/event fields exactly and record delivery eligibility separately without changing them.
- Queue persistence or a daemon could be smuggled into the adapter. Keep previous/proposed queue state caller-supplied, one-cycle, and non-durable.
- Configuration semantics could be invented. Use explicit Notification-owned contracts and defer Configuration 1.0 integration until separately authorized.
- Tests could depend on real services, time, sleeps, randomness, or environment. Use explicit virtual timestamps and deterministic conformance providers only.
- Rollback could overwrite pre-existing Owner work. Snapshot exact working-tree targets and restore only collision-free verified Task 029 changes.

No risk requires an Owner scope change. Concrete provider transport, secrets, durable queues, daemon lifecycle, configuration integration, inbound callbacks, and production support decisions remain intentionally deferred.



## Planned Work


### Phase 1 — Task preparation

1. Verify canonical idle and Task 028 completion without installing or activating Task 029.
2. Review every authoritative input and confirm Notification Delivery is the direct successor to the pure Alert Engine.
3. Confirm no mandatory prerequisite remains for a provider-neutral deterministic planner and injected one-cycle adapter.
4. Define exact Alert-only consumption, delivery-owned routing, queue, retry, provider, status, acknowledgement, evidence, privacy, and failure boundaries.
5. Resolve the standalone Alert Record validation gap as a narrow additive compatibility seam, not a separate prerequisite or Alert semantic change.
6. Simulate Owner, Builder, installation, execution, validation, completion, archive, and canonical-idle transitions and encode all preventable interruption checks.
7. Prepare and self-validate `current-task-job.txt` as data only, without approval token, installation, implementation, staging, commit, or push.
8. Stop at Builder-ready preparation and report any remaining engineering risk and Owner-input requirement.

### Phase 2 — Separately authorized Builder installation

1. Revalidate canonical idle, Task 028 baseline, target absence, repository state, Builder interface, source hash, and rollback assumptions.
2. Transform the reviewed source into deterministic Builder fields without executing its prose.
3. Present the complete generated task to the Owner.
4. Obtain the exact separate standalone Builder approval token; never infer or manufacture it.
5. Run Builder input validation and install the Task 029 prompt/history pair transactionally.
6. Verify Task 029 is the sole active approved task with matching history, permissions, identity, and rollback evidence.
7. Stop at `APPROVED AND READY FOR IMPLEMENTATION`; do not execute Task 029 during installation.

### Phase 3 — Separately authorized task execution

1. Start only through the canonical `job` workflow and verify all authority, lifecycle, baseline, and snapshot gates.
2. Finalize Notification Delivery Model 1.0 identities, taxonomies, versions, bounds, privacy rules, canonical ordering, and compatibility behavior.
3. Add only the narrowly additive standalone Canonical Alert Record validator in `internal/alert`, with no semantic or dependency change.
4. Implement the pure deterministic Notification planner in a new bounded `internal/notification` package.
5. Implement deterministic route matching, fan-out, queue ordering, idempotency, attempt limits, deadlines, retry schedules, exhaustion, and failure isolation.
6. Define the provider-neutral request/result interface, injected registry, capability validation, and deterministic provider conformance harness.
7. Implement the explicitly invoked one-cycle adapter and canonical attempt, status, acknowledgement, evidence, and proposed queue records.
8. Add exhaustive focused tests using explicit time and fake providers, with no live transport or network dependency.
9. Create `docs/architecture/CANONICAL_NOTIFICATION_DELIVERY.md` and update only directly affected permanent documentation.
10. Run every focused and repository-wide required verification, inspect exact diffs and permissions, and validate rollback evidence.
11. Record decisions, implementation, failures/corrections, evidence, limitations, rollback, and Git state in Task 029 history.
12. Complete and archive Task 029 only after all gates pass; create no successor and return to canonical idle.



## Rollback Plan


Preparation rollback is limited to `current-task-job.txt`. Preserve later Owner edits and replace the prepared source only with explicit Owner direction. No repository-wide rollback is authorized.

Builder-installation rollback shall use only the verified transactional Builder snapshot and restore exactly the failed Task 029 lifecycle targets to canonical idle. If content ownership or collision safety cannot be proven, stop and request Owner direction.

Implementation rollback shall:

- stop further mutation and preserve truthful failure and validation evidence;
- verify snapshot manifests, checksums, payloads, permissions, ACLs, absence records, collision guards, and restore instructions;
- compare every affected target with both snapshot and current working tree;
- restore only verified pre-existing Task 029 targets;
- remove only verified Task 029-created files whose prior absence and lack of later Owner edits are proven;
- preserve lifecycle/history truth and all unrelated Owner work;
- rerun lifecycle, Framework, engineering, Go build/test/race/vet/format, focused Notification, source-boundary, documentation, permission, ACL, Git-diff, and snapshot-integrity checks;
- report the exact restored state and every unresolved condition.

Broad `git reset`, `git checkout`, `git restore`, `git clean`, wildcard deletion, repository-wide extraction, ambiguous deletion, and removal of Owner-owned untracked content are prohibited.



## Deliverables


- Canonical Professional Notification Delivery architecture;
- Notification Delivery Model 1.0;
- Delivery Policy, Route, Endpoint Reference, Provider Binding, and provider-neutral Delivery Envelope 1.0;
- deterministic Delivery Plan and eligibility result 1.0;
- bounded previous/proposed Notification Queue State 1.0 without persistence;
- Delivery Request, Delivery Attempt Record, Delivery Status Record, Delivery Acknowledgement, Provider Result, and Delivery Evidence Reference 1.0;
- canonical provider interface, descriptor, capability model, injected registry, and conformance harness;
- deterministic route matching, fan-out, queue ordering, idempotency, retry limit, retry deadline, backoff, exhaustion, and failure-isolation semantics;
- exact delivery channel, status, acknowledgement, and failure taxonomies;
- strict validation, versioning, canonical JSON, privacy controls, resource bounds, and compatibility behavior;
- pure planner and explicitly invoked bounded one-cycle adapter;
- narrowly additive `alert.ValidateRecord` compatibility boundary;
- focused unit, contract, integration, temporal, provider-conformance, determinism, privacy, resource-bound, and regression tests;
- `docs/architecture/CANONICAL_NOTIFICATION_DELIVERY.md` and directly affected canonical documentation updates;
- complete Task 029 history, verified rollback evidence, completed prompt archive, and canonical idle state after execution.

Concrete production providers, real network delivery, secrets, durable queue storage, background processing, and Notification configuration integration are not Task 029 deliverables.



## Verification


Builder and lifecycle verification shall prove:

- correct Task ID, slug, title, authority, language, canonical sections, and absence of placeholders or unresolved mandatory content;
- explicit separation of preparation, Builder installation, and execution;
- absence of embedded or inferred approval token;
- correct pre-install idle, post-install sole-active, and post-completion idle states;
- deterministic Builder input passes canonical validation after separate Owner approval;
- prompt/history identity, approval state, permissions, archive identity, and latest-completed ordering remain exact.

Implementation verification shall include:

- focused `internal/alert` validation-boundary and `internal/notification` tests;
- repository-wide build, Go tests, race tests with configured writable caches, vet, and complete formatting checks;
- Framework 1.x configured validations, engineering suites, lifecycle, active-task, idle-closure, Builder, and test-task validation;
- golden or equivalent contract tests for every public Notification 1.0 record;
- byte-identical canonical JSON and identity across equivalent input ordering;
- tests for Alert, lifecycle, and suppressed Alert decisions proving Notification obeys but never recreates or changes them;
- tests proving only Canonical Alert Records are accepted as alert-condition inputs and no Health, Rule, Policy, Drift, Scheduler, Report, Configuration, Inventory, or host package is imported or evaluated;
- tests for no route, single route, deterministic multi-route fan-out, duplicate routes, endpoint isolation, provider mismatch, unknown provider, and bounded registry behavior;
- tests for queue order, stable idempotency, duplicate-plan replay, attempt identity, finite retries, exact deadline boundaries, rate limits, exhaustion, terminal failures, and indeterminate outcomes;
- tests distinguishing provider acceptance, provider-reported delivery, human receipt, and Alert acknowledgement;
- tests for malformed, unsupported-version, contradictory, tampered, oversized, stale, future-dated, duplicate, secret-bearing, and privacy-unsafe input;
- tests proving no ambient time, environment, filesystem, process, randomness, mutable global state, sleep, live service, or network is used by planner tests;
- provider conformance tests using deterministic fake providers only;
- tests proving one-cycle adapter bounds and absence of daemon, loop, durable persistence, automatic retry execution, callbacks, API, UI, remediation, AI, or recursive Alert creation;
- import and source audits proving provider-specific transport logic and dependencies are absent from canonical architecture;
- existing Alert, Scheduler, Configuration, Policy, Report, Command, Pipeline, CLI, and manual workflow regression compatibility;
- architecture terminology, functional traceability, documentation, privacy, and secret-disclosure audits;
- exact target, ownership, permissions, ACLs, unstaged/staged state, `git diff --check`, and `git diff --cached --check` review;
- verified snapshot checksums, payloads, absence records, collision guards, and bounded restore procedure;
- confirmation that nothing was installed beyond the approved task lifecycle, staged, committed, pushed, deployed, or released.

Verification shall require no external service, credentials, network access, real notification, daemon, real clock delay, or host mutation.



## Documentation Updates


Expected direct documentation targets are:

- `docs/architecture/CANONICAL_NOTIFICATION_DELIVERY.md`;
- `docs/architecture/CANONICAL_ALERT_ENGINE.md` for the implemented downstream boundary only;
- `docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md` only to preserve the explicitly deferred integration boundary;
- `docs/PRODUCT_ARCHITECTURE.md` for implemented-status and provider-neutral boundaries;
- `docs/FUNCTIONAL_SPECIFICATION.md` for traceability between implemented orchestration and deferred production transports;
- `ai/core/04_ARCHITECTURE.md`;
- `ai/core/05_SYSTEM_MAP.md`;
- `ai/core/07_ENGINEERING_HISTORY.md`;
- `ai/core/13_ROADMAP.md`;
- `README.md`;
- `ai/prompts/029_CURRENT_TASK.md` during active execution;
- `ai/history/029_2026-08-07_professional-notification-delivery.md`;
- `ai/archive_prompts/029_2026-08-07_professional-notification-delivery.md` at successful idle closure.

Every actual documentation change and justified omission shall be recorded in Task 029 history. Documentation must distinguish deterministic planning, explicit one-cycle provider invocation, provider acknowledgement, production transport support, queue persistence, daemon operation, and Alert decision ownership.



## Completion Criteria


Task 029 is complete only when:

- Notification Delivery Model 1.0 consumes validated Canonical Alert Records as its only alert-condition input and never decides whether an Alert exists;
- pure planning is deterministic for equivalent explicit inputs and produces byte-identical ordered contracts and identities;
- delivery planning, channel abstraction, attempts, statuses, retry scheduling, acknowledgements, failure classification, evidence, queue abstraction, and provider abstraction are implemented and documented;
- retry limits, deadlines, backoff, queue order, fan-out, idempotency, failure isolation, and resource bounds are exact and fully tested;
- the one-cycle adapter invokes only injected providers and records explicit results without daemon, automatic loop, persistence, or hidden time behavior;
- provider-specific logic, credentials, SDKs, and production transports remain outside canonical architecture;
- Alert Records and all upstream evidence remain immutable and traceable, and delivery failures never fabricate Alerts;
- the narrow standalone Alert Record validator changes no Alert semantics and introduces no reverse dependency;
- existing Framework, lifecycle, Alert, Scheduler, Configuration, Policy, Report, Command, Pipeline, CLI, and manual workflows remain compatible;
- all focused and repository-wide mandatory verification passes;
- rollback remains complete, bounded, proportional, collision-aware, and verified;
- documentation and history accurately distinguish implemented architecture from deferred production operations;
- no dependency installation, deployment, release, staging, commit, push, branch, or tag operation occurred;
- the completed prompt is archived without creating a successor and `bin/job --check` confirms canonical idle with Task 029 as latest completed.

A valid final result is `complete`, `complete with disclosed limitations`, or `blocked`. Completion may not be claimed while any mandatory contract, boundary, test, documentation, rollback, or lifecycle gate remains unresolved.



## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-07 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
