# Current Engineering Task 031: Canonical Runtime Service

## Task Metadata

- Task ID: `031`
- Task slug: `canonical-runtime-service`
- Status: `complete`
- Date opened: `2026-08-07` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Canonical Runtime Service


## Objective


Establish the Canonical Runtime Service as the single continuously operating local process-lifecycle coordinator above the existing Canonical Runtime Engine.

Task 031 shall repeatedly invoke `runtime.Coordinator.Run` through a narrow Runtime Runner interface at explicit deterministic fixed-rate boundaries. It shall provide graceful startup, cancellation-aware operation, bounded cycle deadlines, SIGINT/SIGTERM handling, graceful shutdown, canonical service lifecycle/evidence/result contracts, and exact in-memory handoff of Runtime-proposed state between sequential cycles.

The Runtime Service owns only recurrence and service-process lifecycle. The Runtime Engine remains the sole owner of one complete execution cycle and continues to coordinate Scheduler, Command/Pipeline evidence, Alert, and Notification. Task 031 shall neither copy Runtime orchestration nor call any downstream canonical component directly.

Task 031 ends when Runtime Service Model 1.0, the deterministic service loop, a replaceable signal adapter, comprehensive no-sleep tests, permanent architecture documentation, rollback evidence, and canonical lifecycle closure are complete. It does not create an installed or production-supported daemon.



## Scope


Task 031 shall define and implement versioned Canonical Runtime Service Model 1.0 contracts for:

- Service Definition containing stable service identity, positive execution interval, positive per-cycle timeout, startup mode, exact resource limits, and schema/model versions;
- Service State with exact `created`, `starting`, `running`, `stopping`, `stopped`, and `failed` lifecycle states, one active Runtime cycle at most, next nominal cycle time, sequence, counters, and last Runtime Result reference;
- Service Input containing one validated Service Definition and one validated seed `runtime.Input`; the seed contains already resolved Effective Configuration, Alert controls/state, Notification policy/queue, and Runtime state and is not configuration activation;
- Service Event and Evidence records for startup, cycle scheduled, cycle started, cycle completed, missed service interval, shutdown requested, shutdown completed, and terminal failure;
- Service Result containing bounded aggregate counts, final Service State, last Runtime Result identity/outcome, terminal reason token, and no unbounded history or raw Runtime payload duplication;
- strict validation, canonical ordering, stable content identity, UTC normalization, JSON encoding/strict decoding, versioning, privacy limits, failure tokens, and resource bounds;
- injected Clock, Waiter/Timer, Runtime Runner, and Evidence Sink boundaries so deterministic tests require no real sleep, signal, daemon, host collection, provider, network, or service manager.

The selected service lifecycle is:

1. validate Service Definition, seed Runtime Input, dependencies, limits, and initial `created` Service State;
2. transition `created -> starting`, emit startup evidence synchronously, and fail closed if canonical evidence cannot be accepted;
3. establish the first nominal boundary at the explicit service start time; startup mode is `immediate`, so the first Runtime cycle is eligible once startup validation finishes;
4. transition `starting -> running` and execute cycles sequentially with no overlap;
5. for each nominal boundary, derive a deterministic Runtime Execution Context from service identity, cycle sequence, nominal boundary, and per-cycle timeout, then invoke the existing Runtime Runner exactly once with a bounded child context;
6. validate every returned `runtime.Result`; treat a valid `completed`, `partial`, `failed`, `cancelled`, or `timed_out` Runtime outcome as Runtime-owned truth rather than reclassifying component semantics;
7. mechanically carry forward only exact Runtime-proposed `NextState`, `FinalAlertState`, and `FinalNotificationQueue` into the next `runtime.Input`; keep Effective Configuration, Alert controls, Notification policy, and other seed values unchanged;
8. preserve Scheduler state ownership inside the existing Runtime/Scheduler dependencies and never persist or reconstruct it in the Service;
9. advance nominal boundaries by the configured interval from the prior nominal boundary, never from wall-clock completion time;
10. when one or more nominal boundaries elapsed while a cycle was running, emit bounded missed-interval evidence and advance directly to the first future boundary; do not catch up, overlap, burst, or invoke Runtime more than once for a boundary;
11. on caller cancellation, SIGINT, or SIGTERM, transition `running -> stopping`, cancel the active Runtime context, wait for the Runtime call to honor its existing context contract, emit final evidence, and transition to `stopped`;
12. on an invalid Runtime contract, dependency contract violation, impossible/non-monotonic clock observation, evidence-sink refusal, identity failure, or resource-limit violation, cancel active work, transition to `failed`, and stop without another cycle;
13. return one bounded canonical Service Result after stopped or failed termination.

The core Service shall accept cancellation through `context.Context`. A separate local signal adapter may use the Go standard library only to map SIGINT and SIGTERM into context cancellation. The signal adapter shall register and release handlers exactly once per explicit invocation, own no loop beyond the Service call, install no service, and contain no platform policy.

Runtime calls shall remain strictly sequential. The Service shall not launch detached Runtime goroutines, background worker pools, retry Runtime failures, or perform automatic restart. If the injected Runtime Runner violates its cancellation contract, the Service shall not claim graceful bounded shutdown; tests and documentation shall expose this dependency.

Service evidence shall be emitted synchronously through an injected sink and shall use fixed bounded reason tokens and canonical references only. The Service shall retain bounded counters and the last Runtime Result reference, not an ever-growing in-memory event log. A sink may be an in-memory test double; durable logging or persistence is not authorized.

In-memory handoff exists only to preserve deterministic continuity while one Service process remains alive. Restart continuity is explicitly not provided by Task 031 and shall not be implied by Service State or documentation.



## Out of Scope


Task 031 shall not implement:

- any Runtime Engine behavior, Runtime component ordering, Scheduler invocation, Command/Pipeline projection, Alert evaluation, Notification planning, provider invocation, or canonical business decision;
- Scheduler due-time, misfire, retry, overlap, priority, concurrency, locking, persistence, Command resolution, Pipeline execution, or Scheduler event semantics;
- Alert identity, lifecycle, acknowledgement, suppression, correlation, expiration, recovery, severity, category, persistence, or incident semantics;
- Notification routing, provider selection, delivery requests, retry, queue semantics, acknowledgement, failure classification, provider transport, or delivery persistence;
- durable Service, Runtime, Alert, Notification, incident, history, evidence, log, audit, or configuration persistence;
- recovery after process crash, restart continuity, journal replay, durable checkpoints, transactional state, queue broker, database, or cross-engine rollback;
- systemd unit files, systemd installation or enablement, init integration, service supervision, watchdog protocol, socket activation, automatic restart, health probe, PID file, privilege/user management, filesystem layout, log rotation, package installation, deployment, or release;
- continuous monitoring semantics, product-health evaluation, channel-health Alert generation, self-remediation, automatic repair, or host mutation;
- configuration discovery, parsing, merge, activation, reload, file watching, secret resolution, credential handling, environment-variable policy, or provider activation;
- CLI service commands, REST API, Dashboard, Console, Terminal UI, public listener, inbound callback, remote control, remote execution, clustering, fleet management, licensing, billing, AI, or machine learning;
- concrete Email, Webhook, Slack, Discord, Telegram, SMS, or other external transport;
- dependency installation, staging, commit, push, branch, tag, packaging, system installation, deployment, or production-support declaration.

Operating-system installation, durable state, crash recovery, configuration activation, production providers, monitoring/product health, diagnostics, and support hardening remain separately authorized future tasks. Task 031 may document those dependencies but shall not implement substitutes.



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


### Task preparation

- verify the exact QWSG root, Framework 1.x configuration, repository marker, canonical remote URL, primary branch, HEAD, local/remote relationship, complete Git status, empty staged path list, ownership, permissions, and ACLs;
- run `bin/job --check`, `ai/scripts/next-task.sh --check`, `bin/job --check-test-tasks`, and `ai/scripts/framework-check.sh`; require canonical idle with Task 030 as the unique latest completed archive/history pair;
- verify `ai/prompts/` is empty and Task 031 prompt, history, archive, `internal/runtimeservice`, tests, and `docs/architecture/CANONICAL_RUNTIME_SERVICE.md` are absent;
- preserve every pre-existing unstaged and untracked Owner-owned change from Tasks 025–030, QWCS work, Builder sources, and historical backups; require the index to remain empty;
- inspect actual Runtime public contracts and tests and confirm `Coordinator.Run(context.Context, runtime.Input) (runtime.Result, error)` is the sole one-cycle execution boundary required by the Service;
- inspect Product Architecture, Functional Specification, Roadmap, System Map, Runtime architecture, and Task 030 evidence; distinguish the required Version 1.0 continuous local Agent capability from later installation, persistence, recovery, and support gates;
- confirm the Service can carry exact Runtime-proposed states in memory without inventing Runtime semantics and that lack of restart continuity is an explicit limitation, not hidden persistence;
- confirm no systemd integration, package installation, production provider, API, monitoring engine, configuration activation, or durable store is a prerequisite for implementing and testing the canonical service-process contract;
- validate `current-task-job.txt` as readable non-symlink UTF-8 data with no NUL, unfinished fence, placeholder, approval token, ambiguous section, competing lifecycle, or material unresolved decision.

### Separately authorized Builder installation

- repeat canonical idle, Task 030 baseline, Task 031 target absence, source content/hash, repository/Git state, Framework, permissions, ACL, Builder interface, and lifecycle checks immediately before installation;
- map only Owner-supplied fields from this source into a mode-0700 temporary Builder input directory; keep Builder-generated metadata, Required Reading, and approval prose separate;
- create and verify a proportional Builder installation snapshot covering exact lifecycle destinations, verified absence records, source/hash evidence, repository identity, lifecycle state, Git state, ownership, permissions, and ACLs;
- run read-only `task-builder.sh --check-input` only after a separate exact Owner approval token is provided as protocol data; validation shall not install;
- present the complete generated task, obtain the separate standalone exact Owner approval token, and install exactly one Task 031 prompt/history pair transactionally with no clobber;
- require `ai/prompts/031_CURRENT_TASK.md` to be the sole active approved task with one matching history, Task 030 as latest completed baseline, and all Builder/lifecycle/Framework/repository/Git validations passing;
- stop after installation at `APPROVED AND READY FOR IMPLEMENTATION`; do not start Task 031.

### Separately authorized task implementation

- start only through an explicit canonical `job` invocation after successful Builder installation;
- read every Required Reading item and the active prompt/history as data;
- require Task 031 as the sole active approved task and Task 030 as the unique latest completed baseline;
- verify exact Runtime APIs, cancellation behavior, JSON/state contracts, target absence, repository/Git state, ownership, permissions, ACLs, and baseline validations;
- create and verify the proportional implementation snapshot before modifying any target;
- stop on any material authority, safety, lifecycle, scope, Runtime compatibility, cancellation, rollback, privacy, or correctness difference.

The valid preparation and pre-install state is canonical idle. The valid post-install state is one active approved Task 031. The valid post-completion state is canonical idle with Task 031 as the unique latest completed archive/history pair and no successor.



## Snapshot Requirements


Task preparation shall modify only `current-task-job.txt` and shall not create an implementation or Builder installation snapshot. The prior Owner-owned source content remains recoverable from existing conversation/review evidence; do not touch lifecycle or implementation targets.

Before separately authorized Builder installation:

- create a unique external snapshot of the exact Task 031 prompt/history lifecycle destinations, their verified absence, `current-task-job.txt` content and SHA-256, repository identity, Task 030 idle baseline, complete Git state, ownership, permissions, and ACLs;
- verify deterministic manifest, checksums, payload readability, absence evidence, collision guards, and exact transactional restore instructions before installation;
- retain the snapshot through Builder validation and Owner acceptance.

Before separately authorized implementation:

- create one unique rollback-capable snapshot outside the repository for every existing Runtime, command entry-point if directly affected, permanent documentation, prompt, history, and archive target;
- record verified absence for new `internal/runtimeservice`, its tests, Runtime Service architecture documentation, and Task 031 archive target;
- preserve exact current working-tree content, including pre-existing Owner-owned modifications; never substitute HEAD content;
- capture repository identity, branch, HEAD, remotes, ahead/behind, complete Git state, target inventories, ownership, permissions, ACLs, baseline validation, manifest, SHA-256 checksums, readable archive inventory, and guarded restore instructions;
- verify every checksum, payload, absence record, collision guard, restore precondition, and proportional target before implementation;
- retain the snapshot through completion and Owner acceptance.

Snapshot scope shall exclude broad repository archives, live process state, signal state, Scheduler stores, provider payloads, secrets, credentials, external responses, build caches, and unrelated data.



## Risk Assessment


Primary risks and mandatory mitigations:

- The Service could become a second Runtime Engine. Depend only on a narrow Runtime Runner and never import or call Scheduler, Command, Pipeline, Health, Rule, Policy, Report, Alert, or Notification entry points.
- A repeated loop could duplicate alerts after each cycle. Carry forward exact `FinalAlertState`, `FinalNotificationQueue`, and `NextState` from each validated Runtime Result without interpretation; document that restart continuity is not available without later persistence.
- Fixed-delay timing could drift with cycle duration. Anchor every nominal boundary to explicit service start plus sequence multiplied by interval, and use an injected waiter.
- Fixed-rate catch-up could create bursts or overlap. Permit one active Runtime call only, skip elapsed nominal boundaries deterministically, emit one bounded skipped-count record, and advance to the first future boundary.
- Cancellation could start another cycle or lose a terminal result. Check context before waiting and invocation, cancel the active bounded Runtime context, wait for Runtime's existing cancellation contract, and never schedule again after stopping begins.
- Signal handling could leak registrations or become OS installation. Isolate SIGINT/SIGTERM mapping in one standard-library adapter, release it on return, and create no unit, package, process supervisor, or startup policy.
- A Runtime implementation could ignore cancellation. Do not detach or abandon it, do not falsely claim bounded shutdown, and expose the Runtime Runner cancellation contract as a verification and support prerequisite.
- Runtime operational outcomes could be reinterpreted as service decisions. Validate and record Runtime outcome/reference exactly; continue on valid terminal Runtime results according to interval policy, and fail only on invalid contracts or Service-owned failures.
- Continuous evidence could grow without bound. Emit immutable records synchronously to an injected sink and retain only bounded counters plus the last Runtime reference; durable storage remains outside scope.
- Evidence sink failure could hide service behavior. Fail closed with a bounded reason token and stop; never embed raw sink errors.
- Ambient clock changes could make behavior nondeterministic. Require explicit UTC observations, monotonic non-regression within the service contract, fixed nominal arithmetic, and tests for forward jumps, exact boundaries, and invalid backward observations.
- Cycle identity collisions could duplicate Runtime evidence. Derive identity deterministically from Service identity, sequence, nominal boundary, and schema version and test uniqueness/canonical reproduction.
- A seed change could become configuration reload. Hold all non-state seed values immutable for the Service invocation; configuration activation/reload requires later authority.
- Continuous tests could hang or sleep. Use fake Runner, Clock, Waiter, Sink, and signal channel with explicit bounded test contexts; prohibit real-time sleeps and real OS signal delivery in unit tests.
- Service records could expose Runtime data, configuration, destinations, or raw errors. Store only canonical IDs, outcomes, counters, timestamps, and bounded fixed tokens; never embed Runtime Result payloads, Alert Records, Queue entries, destination/secret references, provider payloads, report prose, host paths, or raw errors.
- Resource arithmetic could overflow or run unbounded. Validate interval/timeout ranges, sequence/counter maxima, duration arithmetic, event size, and skipped-boundary counts before use; fail closed on overflow.
- Rollback could overwrite prior Owner work. Snapshot exact working-tree targets and allow only collision-aware bounded restoration of proven Task 031 changes.

No risk requires an Owner scope change. Persistent restart continuity, daemon installation/supervision, operational recovery, configuration activation, production transport, product-health monitoring, and support hardening remain explicit future prerequisites for a supported Version 1.0 release, not prerequisites for this bounded canonical Service layer.



## Planned Work


### Phase 1 — Task preparation

1. Verify canonical idle, Task 030 completion, repository identity/Git state, target absence, and authoritative Version 1.0 product requirements without installing Task 031.
2. Inspect exact Runtime Model 1.0 APIs and prove that a narrow Runner is sufficient for recurrence.
3. Select fixed-rate, sequential, skip-missed-boundaries timing with injected clock/waiter and no overlap or catch-up burst.
4. Define Service Definition, State, Input, Event, Evidence, Result, lifecycle, timing, identity, cancellation, signal, failure, privacy, and resource contracts.
5. Define exact in-memory Runtime-state handoff and its non-durable restart limitation without introducing persistence or Runtime semantics.
6. Simulate Owner, Builder, installation, implementation, validation, completion, archive, and canonical-idle transitions; encode every preventable interruption.
7. Perform the Version 1.0 Release Minimalism check and confirm continuous local Agent operation is required while OS installation and persistence remain later gates.
8. Prepare and self-validate `current-task-job.txt` as data only; do not add approval data, install, implement, stage, commit, or push.
9. Stop with Builder-ready preparation.

### Phase 2 — Separately authorized Builder installation

1. Revalidate source hash/content, canonical idle, Task 030 baseline, targets, Builder interface, Framework, repository/Git state, ownership/ACLs, and rollback assumptions.
2. Map each reviewed substantive section to its exact Builder field without executing task prose or copying builder-owned metadata.
3. Create/verify the Builder installation snapshot and validate structured input read-only after receiving the exact separate approval protocol token.
4. Present the complete generated task and obtain the exact separate standalone Owner approval token.
5. Install the Task 031 prompt/history pair transactionally and verify sole-active approved state.
6. Stop after installation; do not execute Task 031.

### Phase 3 — Separately authorized task implementation

1. Start through `job`, read all authority, verify baseline and targets, and create/verify the implementation snapshot.
2. Implement Runtime Service Model 1.0 contracts, validation, canonical JSON/strict decoding, identities, limits, and privacy boundaries in `internal/runtimeservice`.
3. Implement one sequential fixed-rate Service loop over the injected Runtime Runner, with immediate first cycle, no overlap, deterministic missed-boundary skipping, and exact in-memory proposed-state handoff.
4. Implement context cancellation, bounded per-cycle deadlines, graceful lifecycle transitions, and the narrow SIGINT/SIGTERM adapter without system installation.
5. Implement synchronous bounded evidence emission and final aggregate Service Result without unbounded retained history.
6. Add comprehensive fake-clock/fake-waiter/fake-runner/fake-sink tests with no real sleep, signal delivery, host collection, external provider, network, filesystem state, or daemon.
7. Create `docs/architecture/CANONICAL_RUNTIME_SERVICE.md` and update only directly affected permanent documentation.
8. Run focused and repository-wide validation, architecture/import/privacy/resource audits, exact diff/permission review, and snapshot/rollback integrity checks.
9. Finalize Task 031 history with exact evidence and disclosed non-durable limitations.
10. Mark complete and archive only after every gate passes; create no successor and return to canonical idle.



## Rollback Plan


Preparation rollback is limited to `current-task-job.txt`. Preserve any later Owner edits and replace it only with explicit Owner direction. No lifecycle or implementation rollback applies because preparation installs and implements nothing.

Builder-installation rollback shall use only the verified transactional Builder snapshot. Remove only a proven Builder-created Task 031 prompt/history pair and restore exact Task 030 canonical idle. Stop if ownership, checksum, destination identity, or collision safety cannot be proven.

Implementation rollback shall:

- stop the Service test/process invocation and preserve truthful partial evidence;
- verify snapshot manifests, checksums, archive readability, payloads, permissions, ACLs, absence records, collision guards, and restore instructions;
- compare every affected target against snapshot and current working tree and refuse to overwrite later or unrelated Owner work;
- restore only verified pre-existing Runtime and documentation targets;
- remove only verified Task 031-created files whose pre-task absence and lack of later Owner edits remain proven;
- preserve lifecycle/history truth rather than rewriting failed or completed facts;
- rerun focused Service/Runtime tests, build, full tests, race, vet, format, Framework, Builder, lifecycle, diverted-test, repository, architecture, privacy, import, permission, ACL, Git-diff, and snapshot-integrity checks;
- report exact restored state, staged/unstaged paths, and unresolved conditions.

Broad `git reset`, `git checkout`, `git restore`, `git clean`, wildcard deletion, repository-wide extraction, process killing by broad pattern, or removal of Owner-owned untracked content is prohibited.



## Deliverables


- Canonical Runtime Service and Runtime Service Model 1.0;
- Service Definition, Service State, Service Input, Service Event/Evidence, Service Result, exact schemas/versions/identities, lifecycle, outcome, reason, and resource taxonomies;
- one continuously running but explicitly invoked local Service loop that repeatedly calls the existing Runtime Runner and terminates only by cancellation, signal, or terminal Service-owned failure;
- deterministic fixed-rate interval arithmetic, immediate startup cycle, single active cycle, missed-boundary skip behavior, bounded cycle timeout, and no overlap/catch-up burst;
- exact in-memory handoff of Runtime-proposed state with no persistence or reinterpretation;
- graceful startup, shutdown, cancellation propagation, and standard-library SIGINT/SIGTERM adapter;
- synchronous privacy-bounded Service evidence and bounded aggregate final result;
- strict validation, canonical ordering/JSON/decoding, deterministic identities, overflow/resource guards, and compatibility rules;
- focused unit, contract, lifecycle, timing, cancellation, signal, failure, determinism, privacy, resource, and Runtime-integration tests without real sleeps;
- `docs/architecture/CANONICAL_RUNTIME_SERVICE.md`;
- directly affected README, Product Architecture, Functional Specification, Roadmap, System Map, Runtime architecture, project architecture/history, prompt/history/archive updates;
- complete Task 031 history, verified rollback evidence, archived completed prompt, and canonical idle closure.

No installed daemon, systemd artifact, persistent store, production transport, monitoring engine, interface, remediation, remote capability, AI, package, deployment, or release artifact is a Task 031 deliverable.



## Verification


Builder and lifecycle verification shall prove:

- exact Task 031 ID/slug/title, authority, language, mandatory sections, unique Task 030 baseline, and absence of placeholders or unresolved content;
- explicit separation of preparation, Builder installation, and implementation;
- no embedded, inferred, or reused Owner approval token in `current-task-job.txt`;
- correct pre-install idle, post-install sole-active approved, and post-completion idle states;
- structured Builder input maps losslessly and passes canonical read-only validation only after separate Owner approval protocol input;
- prompt/history/archive identity, permissions, statuses, chronology, rollback, and no-successor closure are exact.

Implementation verification shall include:

- focused `internal/runtimeservice` tests and existing `internal/runtime` regression tests;
- `make build`, full `go test ./...`, repository-wide `go test -race ./...` with writable configured caches, `go vet ./...`, and complete Go formatting;
- Framework 1.x configured validations, `make engineering-test`, Builder assertions, lifecycle checks, diverted-test audit, active-task validation, and idle-closure validation;
- golden or equivalent tests for every public Runtime Service Model 1.0 contract and byte-identical canonical JSON/identity for equivalent explicit inputs and injected observations;
- startup transition tests for `created -> starting -> running`, immediate first cycle, startup cancellation, invalid seed/dependency, and evidence refusal;
- timing tests for exact fixed-rate boundaries, long/short cycles, forward clock jumps, backward/invalid time, exact-boundary completion, duration/counter overflow, missed one/multiple intervals, and no drift from completion time;
- call-count/order tests proving exactly one Runtime call per executed nominal boundary, one active call maximum, no overlap, no catch-up burst, no Service retry, and no direct downstream engine call;
- state-handoff tests proving only exact Runtime `NextState`, `FinalAlertState`, and `FinalNotificationQueue` are forwarded and all immutable seed fields remain unchanged;
- Runtime outcome tests for `completed`, `partial`, `failed`, `cancelled`, and `timed_out` valid results, plus invalid/tampered Runtime results and Runner errors;
- cancellation tests before startup, while waiting, immediately before a cycle, during a cycle, after partial Runtime completion, and during shutdown without ambient sleep;
- signal-adapter tests using injected/fake signal channels for SIGINT, SIGTERM, registration cleanup, duplicate signal, pre-cancelled context, and unsupported signal isolation; unit tests shall not deliver real OS signals;
- tests proving a Runtime Runner that honors context yields graceful bounded shutdown and documentation/source assertions proving the Service makes no claim when that dependency violates its contract;
- evidence tests for sequence/identity, fixed tokens, synchronous order, sink refusal, bounded counters, missed-count compression, and absence of unbounded retained event history;
- privacy tests rejecting raw errors, Runtime payloads, Alert Records, Queue entries, destination/secret references, provider payloads, report prose, host paths, credentials, configuration bodies, signal implementation details, and unbounded metadata in Service records;
- resource tests for interval, timeout, maximum cycle sequence/counters, evidence sizes, duration arithmetic, and fail-closed limit behavior;
- import/source audits proving `internal/runtimeservice` depends on `internal/runtime` only among canonical engine packages and contains no Scheduler, Command, Pipeline, Health, Rule, Policy, Report, Alert, Notification, provider, persistence, monitoring, API, Dashboard, remediation, remote, AI, network, package, or installation boundary;
- regression tests for Runtime, Scheduler, Command/Pipeline, Alert, Notification, Configuration, CLI, Framework, and manual one-shot workflows;
- documentation terminology and cross-document consistency review, including explicit distinction among service recurrence, Scheduler due-time, Runtime execution, and systemd/OS supervision;
- exact changed targets, ownership, permissions, ACLs, staged/unstaged paths, `git diff --check`, `git diff --cached --check`, and preservation of all unrelated Owner-owned content;
- snapshot checksum, payload, readability, absence evidence, collision guards, and bounded rollback verification;
- confirmation that nothing was installed, started as a real resident daemon, persisted, staged, committed, pushed, packaged, deployed, or released.

Verification requires no live host collection, real clock sleep, real OS signal delivery, provider, credential, network access, systemd/service manager, privileged operation, persistent runtime state, remote system, or infrastructure mutation.



## Documentation Updates


Expected direct documentation targets are:

- `docs/architecture/CANONICAL_RUNTIME_SERVICE.md`;
- `docs/architecture/CANONICAL_RUNTIME_ENGINE.md` for the Service caller boundary and unchanged one-cycle ownership;
- `docs/PRODUCT_ARCHITECTURE.md` for the local continuous Agent foundation and deferred operational-support gates;
- `docs/FUNCTIONAL_SPECIFICATION.md` for explicit service recurrence/start/stop behavior without falsely claiming restart continuity;
- `ai/core/04_ARCHITECTURE.md`;
- `ai/core/05_SYSTEM_MAP.md`;
- `ai/core/07_ENGINEERING_HISTORY.md`;
- `ai/core/13_ROADMAP.md`;
- `README.md`;
- `ai/prompts/031_CURRENT_TASK.md` during active implementation;
- `ai/history/031_2026-08-07_canonical-runtime-service.md`;
- `ai/archive_prompts/031_2026-08-07_canonical-runtime-service.md` at successful closure.

Scheduler, Alert, Notification, Configuration, Command/Pipeline, installation, systemd, monitoring, security, and user documents shall change only if implementation proves a direct boundary correction is necessary; every actual update or justified omission shall be recorded in Task 031 history.

Documentation shall state that Task 031 is an explicitly invoked, continuously running process contract, not an installed/supervised production daemon; state handoff is in-memory only; Scheduler alone retains its existing persistence; graceful shutdown depends on Runtime honoring context; and Version 1.0 still requires separately governed persistence, crash recovery, operational integration, production transport, and support evidence.



## Completion Criteria


Task 031 is complete only when:

- Canonical Runtime Service Model 1.0 and one continuously operating local Service loop exist;
- the Service repeatedly invokes only the existing Runtime Runner at deterministic fixed-rate nominal boundaries and never calls downstream canonical components;
- the first cycle, interval progression, missed-boundary skipping, cycle timeout, no-overlap rule, and no-catch-up behavior are exact, bounded, deterministic, and tested without sleep;
- graceful startup, cancellation, SIGINT/SIGTERM mapping, active-cycle cancellation, graceful shutdown, stopped/failed terminal states, and evidence are exact and comprehensively tested;
- every valid Runtime Result is validated and its proposed states are forwarded exactly without duplication, interpretation, persistence, or false atomicity;
- Runtime business logic and canonical outcomes remain unchanged and existing Runtime/manual workflows pass regression validation;
- Service contracts expose only bounded identifiers, times, outcomes, counters, and fixed tokens and no raw Runtime/host/provider/configuration/secret data;
- restart continuity, durable evidence, persistence, crash recovery, configuration reload, systemd installation/supervision, automatic restart, watchdog, monitoring, product-health decisions, concrete providers, interfaces, remediation, remote execution, AI, packaging, deployment, and release remain absent and explicitly documented;
- the Release Minimalism check records that continuous local Agent execution is required for QWSG Version 1.0, while this task alone does not satisfy all Version 1.0 operational support gates;
- build, full tests, race, vet, format, focused timing/cancellation/signal/resource/privacy tests, Framework, Builder, lifecycle, diverted-test, repository, architecture/import, documentation, permission/ACL, Git-diff, and snapshot/rollback validations all pass;
- rollback remains proportional, collision-aware, verified, and preserves unrelated Owner-owned work;
- no dependency installation, real resident-service start, persistence migration, staging, commit, push, branch, tag, package, deployment, or release occurred;
- the completed Task 031 prompt/history are archived without a successor and `bin/job --check` confirms canonical idle with Task 031 as unique latest completed baseline.

A valid result is `complete`, `complete with disclosed limitations`, or `blocked`. Completion may not be claimed while any mandatory service contract, Runtime boundary, cancellation/signal behavior, deterministic interval rule, state handoff, evidence/privacy/resource limit, test, documentation, rollback, or lifecycle gate remains unresolved.



## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-07 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
