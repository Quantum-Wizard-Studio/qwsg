# Current Engineering Task 037: Operational Guardian Service

## Task Metadata

- Task ID: `037`
- Task slug: `operational-guardian-service`
- Status: `complete`
- Date opened: `2026-08-10` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Operational Guardian Service


## Objective


Turn the existing Canonical Runtime Service into a genuinely operating, continuously supervised local Linux Guardian and publish current validated lifecycle evidence through the existing Current Operator State, Operator Presentation Model, and Console chain.

Implement the smallest operational composition around the existing Runtime Engine and Runtime Service: one foreground `qwsg guardian run` process, one supported systemd user unit, canonical configuration consumption, a private bounded recovery checkpoint, atomic operator-state publication, and truthful lifecycle/failure/freshness projection. The service must execute recurring canonical Runtime cycles without copying Runtime, Scheduler, Pipeline, Alert, Notification, Health, Rule, Policy, or Report logic.

After completion, a separately started bare `qwsg` must truthfully show Guardian `starting`, `running`, `degraded`, `stopping`, `stopped`, `unavailable`, or `not observed` from current validated evidence. A one-shot observation, PID, installed unit, prior process, or Console availability must never establish a running claim. A stale former-running record must become unavailable, not remain running.

Task 037 ends when the supported Ubuntu 24.04 systemd service can be installed as a bounded artifact, started, observed across recurring cycles and processes, gracefully stopped, restarted, recovered after a simulated host/service restart, failed safely, resource-tested, documented, and archived. It does not perform Version 1.0 release publication or general release hardening.



## Scope


- Reconfirm the Task 030 Runtime Engine and Task 031 Runtime Service as the only owners of one-cycle orchestration and fixed-rate recurrence. Build a thin operational adapter that supplies their dependencies and never invokes downstream engines outside Runtime/Scheduler's existing ownership.
- Add a foreground, non-forking `qwsg guardian run` application entry point. It shall run exactly one Canonical Runtime Service invocation until SIGINT, SIGTERM, or terminal failure; it shall create no PID file, self-daemonization, detached goroutine, polling supervisor, or automatic restart loop.
- Keep existing canonical command profiles and bare `qwsg` behavior compatible. Service start/stop/enable/disable/restart remain systemd operations and are not Console actions. Add only bounded read-only Guardian configuration validation/status diagnostics if needed for safe operation; do not create a general lifecycle CLI.
- Use a systemd user service as the supported Version 1.0 supervisor. Ship one auditable unit under the existing installation layout and bounded Make targets or documented copy steps that install but do not silently enable it. The process runs as the invoking ordinary user. Privileged responsibility is limited to installing the executable/unit and, when boot-before-login is required, explicitly enabling that user's lingering; runtime collection and state publication remain unprivileged.
- Define the unit as `Type=simple` with the QWSG process in the foreground, `KillSignal=SIGTERM`, bounded `TimeoutStopSec`, `Restart=on-failure`, fixed `RestartSec`, and `StartLimitIntervalSec`/`StartLimitBurst`. Use `UMask=0077`, `NoNewPrivileges=true`, a private temporary directory, a read-only system view except declared private QWSG state paths, and explicit task/memory/CPU limits compatible with read-only collectors. Do not use systemd watchdog, socket activation, dynamic privilege, root runtime, or a competing health decision.
- Add one narrow systemd exit-report adapter invoked by `ExecStopPost`. It may accept only allowlisted systemd termination categories plus a QWSG launch-generation identity, correlate them to the exact active Guardian checkpoint, and demote a prior active claim to stopped, degraded, or unavailable. It must never promote a service to running, parse free-form `systemctl`/journal text, expose raw exit data, or operate on another generation. This closes the crash/invalid-start window that freshness alone cannot close immediately.
- Use systemd enablement as service enablement. Document `systemctl --user start`, `stop`, `restart`, `enable`, `disable`, `status`, and `journalctl --user-unit`; distinguish installed, enabled, active, and canonically observed states. Enabling at boot requires a verified user manager and explicit lingering; lack of lingering is unavailable boot operation, never an implied guarantee.
- Consume Canonical Configuration Source Record 1.0 through strict bounded decoding and resolve it together with the existing built-in source into Effective Configuration 1.0. The operational adapter may supply one documented built-in operational source containing the minimum enabled local `observe` schedule/check binding needed for a useful default Guardian, but it must pass through the existing Configuration normalization, precedence, provenance, identity, reference, and validation rules. It must not add configuration precedence, hidden overrides, Rule/Policy semantics, or an alternate configuration engine.
- Keep Runtime Service cadence and per-cycle timeout explicit through its existing validated Service Definition. Define stable Version 1.0 operational defaults in the adapter and documentation, derive no timing from ambient heuristics, require cycle timeout to be shorter than the recurrence interval and systemd stop timeout, and allow only bounded validated overrides through the Guardian command/configuration boundary. Scheduler remains the owner of schedule due-time, misfire, overlap, retry, and persisted scheduling state.
- Reuse `scheduler.FileStore` and `scheduler.FileLocker` in a private state subtree. Use the existing Inventory Store and Current Operator State roots. Resolve all roots from explicit command inputs first and the existing XDG/HOME policy second; require clean absolute paths, reject unsafe symlinks, preserve `0700` directories and `0600` records, and never mix privileged and unprivileged state.
- Add the minimum private atomic Guardian recovery checkpoint required for exact process restart continuity. It may contain only schema/model identity, service/configuration identity, last completed cycle reference and time, and the exact validated Runtime-proposed `NextState`, `FinalAlertState`, and `FinalNotificationQueue`. It must be bounded, integrity-protected, strict-decoded, owner-only, atomically replaced, and written only after a validated Runtime result. It is not history, an event log, an incident database, or general persistence.
- On first start, initialize canonical empty Runtime/Alert/Notification states. On restart, validate and resume only the last completely checkpointed proposed states. A missing checkpoint is a first start. A corrupt, incompatible, permission-invalid, identity-mismatched, or configuration-mismatched checkpoint fails closed with a bounded diagnostic and must not be overwritten. Scheduler retains its independent durable state and recovery semantics.
- Before configuration activation can leave a historical running claim current, record a bounded launch generation under the operation lock. If configuration or dependency validation then fails, the process or correlated `ExecStopPost` adapter must atomically publish degraded/unavailable operational evidence while preserving the last qualified engineering facts. If that publication itself fails, the old record remains intact and its fixed freshness deadline makes it unavailable; logs and systemd status remain the only immediate evidence, never a false new state.
- Treat an interrupted cycle as incomplete. Never checkpoint an unvalidated or partial write. Preserve the previous valid checkpoint and Scheduler evidence; on restart let the existing Scheduler configuration/misfire/retry/overlap contracts decide due work. Do not synthesize successful completion, recovery, Alert transitions, or Notification delivery.
- Extend the Runtime Service's existing synchronous evidence boundary only as much as necessary to expose each exact validated proposed Service State with its corresponding Event/Evidence. Do not reconstruct lifecycle state from log text, event position, PID state, systemctl prose, or timing. Preserve existing deterministic Service behavior and compatibility where possible; if the public evidence contract must version, apply one coherent minimal version change with legacy read compatibility where safe.
- Publish Current Operator State atomically on starting/running cycle evidence, every completed Runtime cycle, stopping, stopped, and terminal failed/degraded evidence. Publication shall consume validated Service State/Result, Runtime Result, Alert Records, and exact typed Command/Pipeline outputs already produced by Scheduler traces; it shall not rerun or reinterpret an engine. Refactor the existing Task 035 typed projection helper when needed instead of duplicating its eight-stage validation/correlation rules.
- Extend Current Operator State coverage/provenance minimally for Guardian operation. A pre-first-cycle state may truthfully contain service lifecycle evidence with unknown server condition. A completed cycle publication shall preserve the complete qualified engineering evaluation and add current Runtime/Service evidence. Publication failures are terminal visible service failures and preserve the last valid state.
- Keep exactly one operational publisher through a private nonblocking instance lock. A second Guardian fails safely. Explicit one-shot `observe` must not race a running Guardian's Inventory/Scheduler/Current State writes: either acquire the same bounded operation lock or refuse with a privacy-safe `guardian_active` diagnostic. It must not silently overwrite current Guardian evidence with `not_observed`.
- Extend the existing Presentation Model only where its current taxonomy cannot express the required facts. `running` requires current validated Service State. `degraded` means the process is currently running but the last validated Runtime cycle is partial, failed, cancelled, timed out, publication-invalid, or otherwise not successfully complete. `starting`, `stopping`, and `stopped` come only from matching Service lifecycle states. `unavailable` means previously observed operational evidence can no longer support a current lifecycle claim, including stale former-running evidence. `not_observed` remains absence of any valid lifecycle evidence. A terminal Service failure remains visible and must not become healthy.
- Requalify lifecycle freshness in the Presentation Model, not the Console. At and after the exclusive lifecycle freshness deadline, a former `running`, `starting`, or `stopping` claim becomes `unavailable`, completeness becomes partial, condition cannot be healthy, and the recommendation becomes inspection/service verification. Stopped evidence may remain stopped only while its stated evidence is current; stale stopped evidence also becomes unavailable.
- Apply the smallest coherent Presentation Model and Current Operator State version change if the new closed Guardian tokens/coverage cannot be safely read by older strict validators. New readers must retain explicit compatibility with valid 1.0/1.1 stored records where safe; old readers must fail incompatibly rather than misstate a new lifecycle token.
- Keep Console rendering read-only and localization-owned. Add English/Hungarian labels and recommendations only for new canonical tokens. The Console loads Current Operator State and never calls systemd, probes a process, manages the service, reads journals, or decides whether Guardian is healthy.
- Emit bounded privacy-safe service diagnostics and journal lines containing stable lifecycle/failure tokens, version/configuration/service identities only where non-sensitive, cycle sequence/outcome, and timestamps. Never log raw configuration, host evidence, paths, usernames, environment, secret references, destinations, provider payloads, or wrapped Go errors.
- Notification operation is optional for Task 037. The default Guardian uses a valid empty delivery policy and provider registry while Alert Records remain locally visible. If an explicitly configured provider dependency is absent or fails, Runtime's existing partial/failed outcome must degrade Guardian evidence; the service must not invent delivery success. No concrete provider is required.
- Verify bounded operation with fixed service limits of `MemoryMax=128M`, `TasksMax=32`, and `CPUQuota=25%`; acceptance additionally records steady idle RSS, file descriptors, goroutines, and CPU over at least five recurrence boundaries and proves no monotonic growth. Adjusting these stated bounds requires evidence and Owner approval because it changes the approved operating contract.
- Preserve all unrelated Owner-owned working-tree content. Engineering artifacts remain English; operator content is localization-ready. Do not stage, commit, push, fetch, branch, tag, package, publish, or deploy beyond the explicitly isolated Task 037 service acceptance.

Expected implementation targets are a narrow `internal/guardian` operational adapter and private checkpoint/lock code, the minimal `internal/runtimeservice` state-observation seam, shared application projection/publication code, `cmd/qwsg`, `internal/presentationmodel`, `internal/operatorstate`, `internal/operatorconsole`, one systemd user unit, Make/install documentation, tests, and directly affected architecture/product/lifecycle documents. Existing canonical engine logic is not a target.



## Out of Scope


Task 037 shall not implement or redesign:

- Runtime, Scheduler, Command/Pipeline, Inventory, Snapshot, Comparison, Drift, Health, Rule, Policy, Report, Alert, or Notification business semantics;
- a second recurrence loop, daemon framework, supervisor, watchdog protocol, PID file, socket activation, systemd decision engine, service-status polling loop, or process discovery engine;
- a root-running Guardian, privileged collectors, capability grants, sudo inside the service, arbitrary shell execution, remediation, automatic repair, or infrastructure management;
- general configuration activation, editor, hot reload, file watcher, secret backend, credential retrieval, or a service-specific precedence model;
- durable observation history, incident database, audit database, Runtime event history, journal ingestion, general persistence platform, distributed transaction, queue broker, or replay engine;
- concrete SMTP, Telegram, Discord, webhook, Slack, SMS, or other provider transport; provider installation; durable provider delivery continuity beyond the minimal exact checkpoint;
- Console service controls, general lifecycle CLI, Web Dashboard, REST API, public listener, remote access, remote agents, fleet management, clustering, cloud control plane, billing, licensing, AI, or telemetry;
- distro packaging, signed repositories, automatic updater, broad installer redesign, supported-platform expansion, release artifact publication, Version 1.0 declaration, or Task 038 Release Candidate work;
- destructive manipulation of unrelated user/system services, real user state, external configuration, logs, packages, accounts, groups, or dependencies;
- staging, commit, push, fetch, branch, tag, release, or deployment operations.

Task 038 remains the earliest possible Version 1.0 Release Hardening/Release Candidate task. It must not begin during Task 037.



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


### Task-preparation state

- Verify repository `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, project markers, Framework 1.x configuration, canonical HTTPS origin, `main`, exact HEAD, `origin/main...main`, complete Git status, empty index, ownership, modes, and ACLs.
- Require canonical idle with Task 036 as the unique latest completed archive/history pair, no active prompt, and no Task 037 prompt/history/archive collision.
- Preserve all pre-existing Owner-owned modified and untracked paths exactly except this authorized `current-task-job.txt` preparation source.
- Read Tasks 030–036, Runtime, Runtime Service, Configuration, Scheduler, Alert, Notification, Presentation Model, Current Operator State, Console, Command, Product Architecture, Functional Specification, Roadmap, System Map, Makefile, installation/service documentation, Builder, Framework, lifecycle, and Git policies.
- Confirm by source/import inspection that Runtime performs one cycle, Runtime Service owns recurrence and signals, Scheduler already owns a private durable state/lock, Current Operator State is one atomic projection handoff, Presentation already accepts Runtime/Service observations, and Console only renders Overview.
- Confirm the genuine missing seams: no installed/supervised process, no production foreground Guardian entry point, no activated canonical operational composition, no durable Runtime/Alert/Notification handoff, no exact running Service State callback/publication, no Guardian-aware Current State coverage, no stale-running demotion, and no service unit/install contract.
- Prove a narrow operational adapter plus a minimal Runtime Service state-observation seam satisfies the requirement. Stop if implementation would require changing any canonical engine decision or creating a competing architecture layer.
- Verify the supported acceptance baseline is Ubuntu 24.04 with systemd 255 or later, GNU Make, the repository Go version, user-manager availability, writable isolated user state/config roots, and authority for a uniquely named temporary user service. Record when the engineering shell is sandboxed even though the underlying system manager is available.
- Validate systemd user-service selection against state ownership: a separate service account cannot read/write the current owner-only State without weakening its contract, while the same ordinary user preserves existing ownership and Console access. Record that privileged installation and optional lingering are separate from unprivileged runtime.
- Validate this source as readable, non-symlink, UTF-8, NUL-free, bounded text with every Builder field, no placeholder, unfinished fence, embedded approval protocol value, competing lifecycle, unresolved product choice, or post-install approval trap.

### Separately authorized Builder installation state

- Immediately reverify canonical idle, Task 036 latest completion, Task 037 destination absence, exact source content/hash, repository identity, Git/index state, Framework, Builder interface, lifecycle, permissions, ACLs, and preservation inventory.
- Create one unique external mode-0700 Builder snapshot covering exact Task 037 prompt/history/archive absence, this source and hash, Task 036 archive/history baseline, repository/Git identity, ownership/modes/ACLs, checksums, collision guards, retention, and exact bounded restore instructions. Verify every manifest and absence record before Builder execution.
- Map each owner-authored Builder field to a separate UTF-8 regular file in a unique mode-0700 temporary input directory. Supply the approval protocol value separately; it must not occur in this source.
- Run `task-builder.sh --check-input` before its transactional installation. Install exactly `ai/prompts/037_CURRENT_TASK.md` and the matching Task 037 history, with no clobber or unrelated lifecycle mutation.
- Require Task 037 as the sole active approved task, Task 036 as latest completed baseline, empty index, unchanged branch/HEAD/ahead-behind relation, preserved Owner content, and passing Builder/Framework/lifecycle/diverted-task/repository validations. Stop without implementation.

### Implementation starting state

- Start only through a later explicit canonical `job` invocation. Read the complete installed prompt, history, skill, governance, Required Reading, Task 030–036 contracts, and service/platform documentation before action.
- Reverify the task-preparation and post-install facts, supported systemd host, service-manager authority, exact service/unit names, no collision, state/config test roots, user identity, permissions, ACLs, processes, and Git state.
- Freeze existing Runtime/Runtime Service/Scheduler/Alert/Notification behavior with focused tests and import/source audits. Specify the operational adapter dependency graph, Service State callback, checkpoint commit point, publisher single-writer rule, lifecycle freshness deadline, Guardian mapping table, restart matrix, unit sandbox, configuration/defaults, and safe diagnostics before modifying targets.
- Prove deterministic tests can inject clocks, waiters, Runtime results, state callbacks, checkpoint/publication failures, locks, filesystem roots, signals, and subprocesses without real home state, host mutation, network, provider credentials, privilege, or sleeps.
- Create and verify one proportional implementation snapshot before code changes. Do not copy caches, live host evidence, journals, credentials, secret material, or unrelated repository content.
- Reserve a unique `qwsg-task037-acceptance` user-unit identity and explicit temporary config/state roots. Refuse implementation-time live service testing if the unit exists, the bus/user manager is unavailable, cleanup/rollback identity is ambiguous, or a real QWSG Guardian is already active.

The valid preparation/install baseline is canonical idle after Task 036. The valid post-install state is exactly one active approved Task 037. The valid completion state is canonical idle with Task 037 as the unique latest completed archive/history pair and no successor.



## Snapshot Requirements


Task preparation creates no implementation or Builder-installation snapshot. The Owner-owned `current-task-job.txt` remains preparation data.

Before separately authorized Builder installation, create a unique `/tmp/qwsg-task037-builder-install-<UTC>-<random>` mode-0700 snapshot. Include exact source bytes/hash, Task 036 archive/history, Task 037 prompt/history/archive absence, repository root/remote/branch/HEAD/ahead-behind, complete status and empty index, ownership/modes/ACLs, Builder/Framework versions, a SHA-256 manifest, absence records, collision guards, retention statement, and exact no-clobber restoration instructions. Verify all entries before installation and retain through Owner acceptance.

Before implementation, create a separate unique `/tmp/qwsg-task037-implementation-<UTC>-<random>` mode-0700 snapshot containing only every existing authorized target: Guardian/application/Runtime Service/Presentation/Current State/Console code and tests, Makefile, unit/install assets, directly affected README/architecture/product/specification/user/core documents, active prompt/history, and verified absence records for new paths/archive. Record repository/Git identity, exact target inventory, modes, ownership, ACLs, SHA-256 checksums, collision guards, bounded restore instructions, and service acceptance pre-state. Verify all manifests before changes.

Before real systemd acceptance, create `/tmp/qwsg-task037-service-acceptance-<UTC>-<random>` mode 0700 containing the exact built binary/unit/config hashes, temporary unit identity, pre-test `systemctl --user` unit state, test-root identity, journal cursor/time boundary, rollback/cleanup commands, and absence/collision evidence. Use only an exact unique unit and exact temporary QWSG roots. Retain acceptance evidence; remove only proven test-created service links/files/state after successful stop/disable/reset-failed and identity checks.

Snapshots exclude broad repository archives, Go caches, real QWSG state/configuration, unrelated systemd units, system journals, secrets, credentials, provider data, network content, and host payload values. Restoration requires Owner confirmation before overwriting any material post-snapshot work and may touch only identity-proven targets.



## Risk Assessment


- **False running claim — critical:** publish only exact validated Service State through the canonical model; demote stale operational evidence to unavailable; never infer from PID/unit/process/Console/one-shot state.
- **Competing orchestration — critical:** Runtime Service remains the sole recurrence owner and Runtime remains the sole cycle coordinator. The operational adapter wires existing dependencies and extracts already-produced typed traces only for projection.
- **Restart state corruption or false transitions — high:** atomically checkpoint only exact validated proposed Runtime/Alert/Notification states after a completed Runtime result; fail closed on corruption/version/config mismatch; never reset silently.
- **Concurrent publisher/store damage — high:** one private instance/operation lock covers Guardian and explicit observe writes; contention returns a bounded diagnostic without mutation.
- **Privilege expansion — high:** runtime is the ordinary user with no new privileges; root is limited to copying public artifacts and optional linger administration. Preserve owner-only state and separate configuration permissions.
- **systemd restart storm — high:** `Restart=on-failure`, five-second delay, three attempts per sixty seconds, no self-restart, and visible final stale/unavailable evidence bound retries.
- **Unbounded or stuck shutdown — high:** Runtime child deadlines are below Service interval and systemd stop timeout; SIGTERM maps to context cancellation; timeout escalation is documented and tested without claiming graceful completion after forced kill.
- **Configuration ambiguity — high:** strict canonical Source decoding plus existing resolution/provenance; one documented operational default source; invalid or unsupported input fails before running/publication and preserves last valid state.
- **Partial cycle shown healthy — high:** Runtime non-completion maps Guardian to degraded, Current State remains partial, and no checkpoint/publication claims a successful evaluation.
- **Abrupt crash has no in-process terminal event — medium:** the generation-correlated systemd exit reporter demotes the claim when it can publish; otherwise the prior record ages to unavailable. QWSG never invents graceful stop evidence it did not observe.
- **Supervisor exit miscorrelation — high:** `ExecStopPost` may only demote the exact launch generation through a closed result mapping; stale, duplicate, malformed, or unrelated exit reports are rejected and can never create running evidence.
- **Notification unavailable — medium:** empty default policy is valid and Alerts remain visible; configured provider failure remains Runtime partial/failed evidence. No delivery claim or hidden retry is added.
- **Version compatibility — medium:** use the smallest coherent version bump, explicit legacy reader tests, strict old-reader failure, and no unrelated migration.
- **Unit sandbox breaks collectors — medium:** validate unit security settings on supported Ubuntu 24.04 with all read-only collectors; relax only the exact incompatible restriction with documented evidence, never by granting root.
- **Systemd acceptance mutates user service state — medium:** unique temporary unit, exact pre-state snapshot, isolated config/state, explicit start/stop/restart/disable/unlink/reset-failed cleanup, and collision refusal.
- **Resource exhaustion — medium:** systemd Memory/Tasks/CPU limits, bounded queues/contracts, sequential cycles, five-boundary measurement, race tests, FD/goroutine growth assertions, and no retained event log.
- **Dirty working tree collision — medium:** exact proportional snapshots, apply-patch changes, complete status review, empty index, and preservation of every unrelated Owner path.
- **Rollback incompleteness — medium:** separate repository and operational snapshots, identity checks, stop-first service restoration, exact targets only, and no wildcard/broad Git recovery.

No risk requires a new business engine, root runtime, provider, network service, or Owner product decision. The systemd user-service choice is the minimal model compatible with current private per-user state.



## Planned Work


### Phase 1 — Task preparation

1. Verify canonical idle after Task 036, repository/Git/Framework/lifecycle state, Owner content, Builder schema, platform/systemd facts, permissions, and absence of Task 037 collisions.
2. Trace Tasks 030–036 and all relevant contracts from Runtime Service through Current State and Console; prove the missing capability is operational composition, supervision, recovery handoff, and lifecycle publication rather than new engine logic.
3. Select the non-root systemd user-service model, foreground command, strict canonical configuration input, private paths, bounded restart/resource policy, and explicit privileged installation/linger separation.
4. Freeze lifecycle mapping, freshness, recovery checkpoint, single-writer, failure, publication, configuration, logging/privacy, and acceptance contracts.
5. Simulate Builder installation, implementation, isolated systemd operation, stop/restart/failure/recovery, validation, rollback, archive, and canonical idle.
6. Prepare and Builder-validate this `current-task-job.txt` as data only, without an embedded approval value, installation, implementation, service mutation, staging, commit, or push.

### Phase 2 — Separately authorized Builder installation

1. Repeat the complete preparation baseline and verify the exact source hash/content and Task 037 destination absence.
2. Create and verify the bounded Builder installation snapshot and field directory; provide approval separately.
3. Run Builder input validation, transactionally install the Task 037 prompt/history pair, and verify prompt/history identity and contents.
4. Prove Task 037 is the sole active approved task and remains executable from the post-install state; run all required Framework, Builder, lifecycle, diverted-task, repository, Git, permission, ACL, and preservation checks.
5. Stop immediately after successful installation without implementation or service action.

### Phase 3 — Separately authorized implementation

1. Start through `job`, reread all authority, reverify systemd/test authority and collisions, freeze the detailed contract, and create/verify the implementation snapshot.
2. Add the minimal Runtime Service exact-State observation seam with deterministic no-sleep compatibility/failure tests. Do not alter recurrence, cycle, signal, or owning-engine semantics.
3. Implement strict canonical configuration loading and the documented operational default source; build the existing Scheduler adapter, Pipeline collector, Runtime Coordinator, empty Notification policy/registry, and Runtime Service input from validated existing contracts.
4. Implement the private single-record Guardian checkpoint and operation lock with bounded strict decoding, integrity, permissions, atomicity, configuration binding, failure-window tests, and exact restart seeding.
5. Implement foreground `guardian run`, the generation-correlated systemd exit reporter, safe diagnostics, signal propagation, bounded logging, state roots, single-instance behavior, and failure exit codes. Refactor shared typed observe projection instead of copying stage validation.
6. Publish exact starting/running/cycle/stopping/stopped/failed lifecycle and completed evaluation evidence through Current Operator State. Apply the minimal coverage/schema compatibility changes and prove last-valid preservation on every failure window.
7. Harden Presentation Guardian mapping and freshness demotion; add localized Console tokens without service probing/control. Preserve server-health uncertainty and all Task 036 attention behavior.
8. Add the systemd user unit, bounded installation targets, configuration example/generation guidance, enable/disable/start/stop/restart/linger/log/rollback documentation, and unit security/resource verification.
9. Run deterministic unit/integration/subprocess tests for first start, five cycles, exact state handoff, large projection compatibility, stop, restart, crash/stale demotion, invalid configuration/checkpoint, lock contention, Runtime/publication/provider failure, privacy, resource bounds, and legacy state compatibility.
10. Create the service acceptance snapshot; use the unique temporary user unit and isolated roots to run real systemd start, recurring cycles, separate-process Console, graceful stop, restart, failure/restart-limit, stale evidence, resource measurement, and exact cleanup/rollback. Do not touch any existing QWSG or unrelated unit.
11. Run full/race/vet/format/build, Framework, Builder, lifecycle, diversion, unit security, installation, Git, snapshot, ownership/mode/ACL, import/source, privacy, rollback, and preservation validations.
12. Update only directly affected documentation and Task 037 history with actual decisions, commands, evidence, limitations, cleanup, and Git state. Mark complete only when real systemd acceptance passes.
13. Archive Task 037 without creating Task 038 and verify canonical idle with Task 037 as the unique latest completed baseline.



## Rollback Plan


Rollback is exact, identity-checked, service-safe, and file-bounded. Stop immediately on uncertain unit identity, target collision, changed post-snapshot content, unavailable supervisor control, or an inability to preserve a later valid state.

For implementation rollback, stop the exact Task 037 acceptance service if and only if its unit identity and process executable/hash match the acceptance snapshot. Disable and unlink only that exact temporary unit, reload the user manager, reset only its failed state, and verify it is inactive/not found. Remove only exact temporary config/state paths whose root, mode, owner, manifest, and Task 037 marker match. Never stop, disable, reload, remove, or reset another unit; never delete real QWSG state, configuration, journal, user, group, linger setting, binary, or service artifact.

After operational cleanup, verify the implementation snapshot manifest, repository identity, target list, current target identities, and absence of later Owner edits. Restore only pre-existing exact payloads with recorded ownership, modes, and ACLs. Remove only Task 037-created paths whose pre-change absence and current Task identity are both proven. Do not use recursive broad deletion, wildcard targets, Git reset/clean/checkout/restore, replacement from HEAD, or changes outside the snapshot.

Builder-install rollback may remove only the exact generated Task 037 prompt/history pair after verifying Builder snapshot absence, content identity, canonical idle baseline, and no implementation. Restore `current-task-job.txt` only with explicit Owner confirmation if later edits exist.

After rollback run focused/full/race/vet/format tests, `systemd-analyze verify`, Framework/Builder/lifecycle/diverted-task checks, `bin/job --check`, complete Git status/index/ahead-behind, ownership/mode/ACL checks, original Current State/Presentation compatibility, snapshot checksum verification, and service absence/inactivity checks. Retain snapshots, manifests, acceptance journal boundary, failure evidence, and rollback report.



## Deliverables


- one foreground `qwsg guardian run` operational composition over the existing Runtime Service and Runtime Engine;
- one supported least-privilege systemd user unit with bounded restart, shutdown, sandbox, and resource policy;
- one generation-correlated, demotion-only systemd exit-report seam for immediate crash/startup-failure truthfulness;
- strict Canonical Configuration Source consumption and one explicit minimal operational default schedule through Effective Configuration;
- one bounded private atomic Guardian recovery checkpoint and single-writer operation lock;
- exact validated Runtime Service State observation and publication seam;
- Guardian-aware Current Operator State coverage and atomic starting/running/cycle/stopping/stopped/failed publication;
- truthful `running`, `degraded`, `starting`, `stopping`, `stopped`, `unavailable`, and `not_observed` presentation with stale-running demotion;
- localized read-only Console rendering and privacy-safe lifecycle diagnostics;
- deterministic unit/integration/subprocess/recovery/failure/resource/compatibility tests and real isolated systemd acceptance evidence;
- bounded service installation, operation, boot/linger, logs, diagnosis, rollback, and removal documentation;
- directly affected architecture, product, specification, roadmap, system-map, README, English/Hungarian user, prompt/history/archive, validation, and rollback evidence.

No new business engine, alternate daemon framework, root runtime, concrete provider, API, Dashboard, remediation, fleet, AI, general persistence/configuration/lifecycle platform, package, release, stage, commit, or push is a deliverable.



## Verification


- Prove by imports, call counts, fakes, and source audit that Runtime Service alone owns recurrence, Runtime alone owns each cycle, Scheduler alone owns due/retry/misfire/locking/state, Pipeline alone invokes engineering engines, Alert alone creates Alert decisions, and Notification alone plans/provider-calls.
- Test Runtime Service's exact-State observation seam for every starting, running-cycle, stopping, stopped, and failed evidence emission; reject invalid/mismatched state; preserve deterministic event/evidence/result identity and existing no-sleep behavior.
- Validate canonical configuration source strictness, built-in plus primary precedence/provenance, default enabled observe schedule, stable identities, bounds, invalid references, unsupported version, malformed/trailing/oversized content, unsafe path, permissions, and configuration-change restart behavior.
- Verify service cadence, cycle timeout, stop timeout, schedule interval, restart limits, startup ordering, process foreground behavior, signals, and systemd unit with `systemd-analyze verify` plus security review. Prove no fork, PID file, self-restart, overlapping Runtime cycle, catch-up burst, or hidden retry.
- Test every allowlisted `ExecStopPost` result and reject unknown, raw, stale-generation, duplicate, and cross-instance reports. Prove the adapter only demotes lifecycle, immediately removes a crashed/invalid-start running claim when publication succeeds, preserves engineering facts, exposes no raw exit detail, and cannot manufacture running.
- Exercise Guardian checkpoint absent/valid/corrupt/incompatible/truncated/tampered/permission/symlink/config-mismatch states and every write failure window. Prove atomic last-valid preservation and exact Runtime/Alert/Notification proposed-state handoff across process restart.
- Interrupt one cycle between Scheduler's durable evaluation and Guardian checkpoint. Prove no completion/checkpoint is invented and restart follows existing Scheduler overlap/misfire/retry facts without false Alert recovery or duplicate successful state.
- Test one operational publisher, second-instance refusal, one-shot `observe` contention, lock release after graceful/failed exit, no stale lock claim, and no Current State overwrite that removes current Guardian evidence.
- Test lifecycle truth table: exact current Service state yields starting/running/stopping/stopped; running plus partial/failed/cancelled/timed-out Runtime yields degraded; terminal service failure is degraded/failed evidence; missing yields not observed; stale previously active/stopped evidence yields unavailable; none may become healthy without complete current engineering evidence.
- Test Current Operator State publication before first cycle, after successful/partial/failed cycle, on stopping/stopped/terminal failure, and on projection/store failure. Validate coverage, provenance, freshness, identities, atomicity, 1.0/1.1 legacy read compatibility, new-version strictness, and last-valid preservation.
- Prove a full Runtime cycle's already-produced eight-stage typed execution projects the same condition/change/Health/Rule/Policy/Report facts as Task 035/036 and adds Runtime/Service/Alert evidence without re-execution or semantic change. Re-run the 367-record/>732-attention regression.
- Verify Console in a separate process performs no collection/systemd call and displays truthful Guardian, condition, attention, changes, Alerts, evidence, recommendation, and overflow disclosure in English and Hungarian. Ensure all new tokens are localized and output contains no raw errors, paths, usernames, environment, config bodies, host identifiers, or secrets.
- Deterministic product subprocess acceptance: process A starts with isolated empty roots and first-cycle baseline; recurring service cycles produce a qualified evaluation and checkpoint; process B bare `qwsg` consumes it; SIGTERM produces stopping/stopped publication; a new service process loads the checkpoint and resumes with correct sequence/state; forced interruption ages to unavailable.
- Real Ubuntu 24.04 systemd acceptance with unique temporary user unit and roots: verify install/link/daemon-reload/start, at least five actual recurring Runtime cycles, Current State changes, separate `qwsg`, `systemctl --user status`, bounded journal tokens, graceful stop, stopped Console, restart, SIGKILL/abnormal exit demotion through the correlated exit reporter, simulated invalid configuration, Runtime failure, bounded automatic restart and start-limit, stale lifecycle demotion, and exact disable/unlink/reset/cleanup. Capture timestamps, unit states, exit statuses, hashes, and sanitized resource evidence.
- Prove boot semantics without changing the host's existing linger setting: verify enabled unit dependency/install state and document that user-manager boot startup requires lingering. Where a disposable acceptance host permits reboot, perform enable/reboot/running verification; otherwise validate with an isolated manager restart and record reboot as an explicit Task 038 platform-release gate, never as already proven support.
- Measure at least five recurrence boundaries under `MemoryMax=128M`, `TasksMax=32`, and `CPUQuota=25%`; assert RSS remains below 96 MiB, file descriptors do not grow by more than two after warm-up, goroutine count does not grow, cycles remain sequential, and no unbounded log/event/state collection exists.
- Run focused tests for Guardian, Runtime Service, Runtime, Scheduler, Presentation Model, Current State, Console, and CLI; `make build`; full `go test ./...`; repository-wide `go test -race ./...` with writable temporary caches; `go vet ./...`; `make fmt-check`; systemd verification; isolated install/rollback; Framework validation/tests; Builder tests; lifecycle tests; diverted-task audit; active/idle job checks; Git diff checks.
- Audit exact changed files/imports and prove no canonical business engine, provider transport, daemon framework, general persistence/configuration engine, service control Console, network listener, remediation, root runtime, dependency, unrelated service, real user state, package, deployment, stage, commit, push, fetch, branch, or tag changed.
- Verify snapshot checksums, rollback feasibility, service cleanup, empty index, unchanged branch/HEAD/ahead-behind, complete Git status, ownership/modes/ACLs, and exact preservation of unrelated Owner content.

Automated tests use injected clocks and explicit temporary roots without sleep, network, credentials, privilege, or host mutation. Only the separately bounded acceptance phase may use real elapsed time and the user systemd manager, under the predeclared unique unit/root identities and rollback guards.



## Documentation Updates


Update only directly affected sections of:

- `docs/architecture/CANONICAL_RUNTIME_SERVICE.md` for operational hosting, exact State observation, restart handoff boundaries, and unchanged recurrence ownership;
- a concise `docs/architecture/OPERATIONAL_GUARDIAN_SERVICE.md` for the adapter, checkpoint, single-writer, lifecycle evidence, systemd, security, recovery, and failure contract;
- `docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md` only for Guardian consumption and the explicit operational source, without changing precedence/semantics;
- `docs/architecture/CANONICAL_CURRENT_OPERATOR_STATE.md` and `CANONICAL_OPERATOR_PRESENTATION_MODEL.md` for Guardian coverage, tokens, freshness, compatibility, and atomic publication;
- `docs/architecture/INTERACTIVE_OPERATOR_CONSOLE.md` and Command Architecture only for read-only consumption/foreground command boundaries;
- `docs/installation/INSTALL.md` for binary/unit installation, normal-user configuration/state, start/stop/restart/enable/disable, lingering, journal, verification, rollback, and safe removal;
- `README.md`, `docs/PRODUCT_ARCHITECTURE.md`, `docs/FUNCTIONAL_SPECIFICATION.md`, and English/Hungarian operator guidance for actual Guardian operation, limits, failures, and truthful status;
- `ai/core/04_ARCHITECTURE.md`, `ai/core/05_SYSTEM_MAP.md`, `ai/core/07_ENGINEERING_HISTORY.md`, and `ai/core/13_ROADMAP.md` for Task 037 completion and the remaining release gate;
- active prompt, independent Task 037 history, completed archive, and actual validation/rollback evidence.

Document exact supported Ubuntu/systemd assumptions, non-root identity, privileged installation/linger separation, configuration defaults/precedence, state locations/modes, restart limits, resource limits, abrupt-crash/stale behavior, optional Notification limitation, service commands, diagnostics, and what Task 038 must still prove. Record every actual update and justified omission in Task 037 history.



## Completion Criteria


Task 037 is complete only when:

- the foreground Guardian composes the existing Runtime Service and Runtime Engine continuously without duplicating any owning-engine decision or recurrence;
- the supported systemd user unit installs as a bounded artifact, runs as the ordinary user, passes security/unit validation, starts/stops/restarts predictably, propagates SIGTERM, obeys bounded restart/resource limits, and leaves no acceptance residue;
- strict canonical configuration produces one explicit useful local observe schedule and invalid configuration fails before a running claim or destructive state replacement;
- Scheduler state remains Scheduler-owned and durable, while the minimal Guardian checkpoint atomically preserves exact Runtime/Alert/Notification proposed handoff across restart without becoming a general persistence system;
- exact validated Service State and Runtime/evaluation evidence publish through Current Operator State, and a separate bare `qwsg` shows truthful condition, attention, changes, Alerts, freshness, recommendation, and Guardian lifecycle;
- running is never inferred from one-shot success, PID/unit/process existence, historical state, systemctl prose, or Console availability;
- current lifecycle maps starting/running/degraded/stopping/stopped accurately; missing evidence is not observed; stale former-running/stopped evidence becomes unavailable and cannot support healthy;
- graceful stop publishes stopped, restart resumes the last valid checkpoint, interrupted cycles remain incomplete, invalid checkpoint/configuration fails closed, Runtime/publication/provider failures remain visible, and no false successful operation survives;
- one-writer locking prevents Guardian/observe races and a second instance fails safely without stale-lock behavior;
- Presentation/Current State compatibility is coherently versioned and legacy 1.0/1.1 records remain safely readable where promised;
- real-scale Task 036 projection behavior remains correct and no false Health/Rule/Policy/Report semantics change was introduced;
- deterministic tests and real isolated systemd acceptance cover starts, at least five recurring cycles, separate-process Console, stop, restart, recovery, failure, stale evidence, privacy, rollback, and documented resource bounds;
- focused/full/race/vet/format/build, systemd/install, Framework, Builder, lifecycle, diversion, Git, snapshot, ownership/mode/ACL, compatibility, privacy, resource, and rollback validations pass;
- installation/operation/recovery/removal and English/Hungarian user documentation are complete and make no unsupported reboot, provider, platform, packaging, or release claim;
- no unrelated Owner content, real QWSG state, external service, dependency, canonical engine semantics, root runtime, provider, API, Dashboard, remediation, package, deployment, stage, commit, push, fetch, branch, or tag changed;
- Task 037 prompt/history are complete and archived without a successor, and `bin/job --check` reports canonical idle with Task 037 as the unique latest completed baseline.

If a disposable reboot test is unavailable, Task 037 may complete with the disclosed limitation that unit enablement and isolated manager restart are proven while physical reboot remains a Task 038 supported-platform release gate. This does not permit claiming reboot-tested Version 1.0 support. Any inability to run real start/cycle/Console/stop/restart/failure systemd acceptance is blocking, not deferrable.

After Task 037 the only remaining Version 1.0 MUST is Task 038 Release Hardening/Release Candidate: supported-platform matrix including reboot evidence, reproducible build/install/upgrade/remove/rollback, configuration and state migration checks, security/privacy/resource review, diagnostics/support documentation, clean-host acceptance, release notes/versioning, and Owner release decision. Concrete notification transport and durable outbound delivery continuity remain SHOULD unless the Owner makes off-console delivery a release requirement. REST API, Dashboard, fleet/provider operations, remediation, licensing, and AI remain post-1.0.

A valid result is `complete`, `complete with the explicitly bounded reboot-test limitation`, or `blocked`. Completion may not be claimed while continuous Runtime cycles, truthful lifecycle freshness, checkpoint recovery, single-writer safety, Current State publication, separate-process Console visibility, real systemd acceptance, failure visibility, resource bounds, rollback, or lifecycle evidence is unresolved.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-10 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
