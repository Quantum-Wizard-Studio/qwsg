# Current Engineering Task 034: Canonical Current Operator State

## Task Metadata

- Task ID: `034`
- Task slug: `canonical-current-operator-state`
- Status: `complete`
- Date opened: `2026-08-08` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Canonical Current Operator State


## Objective

Establish Canonical Current Operator State 1.0 as the smallest durable, process-boundary handoff between a successful current canonical observation and every later local operator interface.

Task 034 shall make an explicitly publishing `qwsg check` observation available to a separately started bare `qwsg` process without moving engineering interpretation into the Console. It shall persist a validated Canonical Operator Overview inside a versioned integrity-protected Current State envelope, publish it atomically after successful eligible execution and typed projection, load and validate it through the application OverviewProvider, and ask the Canonical Operator Presentation Model to requalify time-dependent freshness before presentation.

The stored choice is a validated canonical Operator Overview rather than raw `command.StageResult.Value` data or a second engineering evidence model. This preserves the typed semantic result already owned and validated by `presentationmodel`, avoids serializing interface-valued stage payloads whose concrete type identity would be lost, and gives future Console/API/Dashboard consumers one presentation-independent source. The Current State envelope owns storage identity, publication provenance, observation and freshness bounds, integrity, and replacement metadata; it does not own condition, attention, Guardian, Alert, completeness, or recommendation interpretation.

Task 034 ends when the Current State contract/store, explicit publication policy, typed projection adapter, separate-process Console consumption, missing/stale/corrupt/incompatible behavior, tests, permanent documentation, rollback evidence, lifecycle completion, archive, and canonical idle state are complete. It is not a general persistence platform.


## Scope

Define Canonical Current Operator State 1.0, expected in a focused package such as `internal/operatorstate`, with:

- fixed schema name, schema version, model version, state ID, payload digest algorithm, and maximum encoded size;
- one validated `presentationmodel.Overview` payload with stable Overview ID;
- publication provenance containing canonical Command Definition ID, Command Execution ID, profile, selected source, covered stages, publication reason, and application version, using stable bounded identifiers rather than arbitrary errors or host values;
- UTC observation time, publication time, and exclusive freshness deadline with exact boundary semantics;
- explicit coverage declaring which canonical domains were actually observed, so successful partial execution is never represented as a full health, policy, Alert, Runtime, or Guardian observation;
- deterministic canonical JSON serialization, strict decoding with unknown-field and trailing-data rejection, identity recomputation, payload digest verification, version rejection, timestamp ordering, collection/resource bounds, and terminal/privacy-safe diagnostics;
- validated load results that distinguish available, missing, stale, corrupt, incompatible, unsafe-path, permission, and IO failure without leaking paths, content, or raw errors into beginner views;
- immutable returned values and no ambient cache or background refresh.

The Current State storage adapter shall:

- resolve one documented per-user local state location through an injected resolver; production defaults shall follow `QWSG_STATE_DIR` when explicitly set, otherwise `XDG_STATE_HOME/qwsg`, otherwise the conventional absolute user state directory under the verified home directory;
- treat path resolution as an application/storage concern, never an engineering-policy input;
- require an absolute clean path, reject symlink path components and non-private ownership/modes, create only the exact state directory with mode `0700`, and publish one current-state file with mode `0600`;
- write a same-directory uniquely named temporary file with no-follow/exclusive creation, bounded bytes, full write, file sync, close, atomic rename, directory sync, and cleanup on every pre-rename failure;
- never overwrite the last valid state until the complete replacement is durable, never expose a partial file, and deterministically recover or fail closed after interrupted publication;
- retain only the single current envelope. It shall create no history, queue, database, migration framework, retention engine, lock service, or remote backend.

Publication policy shall preserve Command Architecture compatibility:

- the existing `check` profile remains exactly `live Inventory -> Snapshot`; its stage order, output, exit codes, explicit store behavior, and advanced composition semantics shall not be silently redefined;
- a successful normalized `check` execution with complete Inventory and Snapshot typed stage outputs is eligible to publish a current observation, but its coverage is explicitly `inventory_snapshot`; it is not a qualified Health, Policy, Alert, Runtime, Runtime Service, or Guardian verdict;
- successful Command Execution alone is insufficient. Publication requires exact typed extraction and validation of every stage promised by the eligible coverage, correlation to the Definition/Plan/Execution identities, and successful `presentationmodel.Project` output;
- failed, incomplete, cancelled, mismatched, corrupt, or unsupported execution shall not replace the last valid Current State;
- `status` remains a read-only live Inventory profile and shall not silently publish; `health` and `report` retain existing store-based semantics and shall not become implicit publishers in this task;
- Task 034 may add an explicit `--no-publish-current` escape hatch and an explicit state-directory option if needed for deterministic administration and tests, but defaults for bare `qwsg check` and bare `qwsg` must resolve to the same documented local location;
- terminal rendering success is not a publication precondition, and publication failure must be reported distinctly without falsely reporting a durable current state. Exact exit behavior shall be documented and compatibility-tested before implementation completion.

Add the narrow typed projection seam at the application/presentationmodel boundary:

- extract `inventory.Snapshot` values only from validated Inventory and Snapshot `command.StageResult` entries for the eligible `check` plan;
- reject interface-decoded maps, duplicate/missing stages, contract/version mismatch, identity mismatch, incomplete stages, or changed ordering;
- construct the existing typed `presentationmodel.Input` using only supported observations and explicit observation/freshness times;
- where the current Presentation Model lacks an Inventory observation input, add the smallest typed `InventoryObservation`/coverage input necessary to distinguish fresh validated Inventory evidence from absent evidence without deriving Health from Inventory;
- project an honest Overview whose condition may remain `unknown` when no Health evidence exists, while freshness is current and coverage/completeness truthfully shows the limited observation;
- add a presentationmodel-owned requalification operation that accepts a validated stored Overview plus its original freshness deadline and a supplied UTC evaluation time, changes only model-owned time-dependent freshness/completeness/condition/recommendations according to one documented table, recomputes identity, and never upgrades or invents evidence;
- keep all condition, attention, completeness, freshness, Guardian, Alert, and recommendation rules inside `internal/presentationmodel`.

Consumption shall occur through the application OverviewProvider path:

- bare `qwsg`, both interactive and non-interactive, attempts one bounded load before rendering and never starts collection automatically;
- a valid fresh state is requalified at the current injected time and supplied to `internal/operatorconsole` as a validated Overview;
- a valid stale state remains visible as stale/partial according to Presentation Model rules, preserving source time and provenance rather than becoming missing;
- missing state produces the existing validated unavailable/not-observed Overview and `run_fresh_check` recommendation;
- corrupt, incompatible, unsafe, or unreadable state fails closed with a bounded localized diagnostic and must not be treated as healthy, current, or absent without distinction;
- interactive explicit refresh uses the same eligible one-shot `check` execution, typed projection, publication, and accepted Overview result; it performs exactly one provider call, no polling/retry/background work, and retains the previous valid Overview if execution or publication fails;
- the Console continues to consume only validated `presentationmodel.Overview` values and shall not import Current State storage, Command, Pipeline, Inventory, or other engine packages.

The task shall include focused unit, integration, separate-process CLI, corruption, compatibility, atomicity, crash-window, permissions, privacy, resource, determinism, freshness-boundary, publication-policy, typed-extraction, Console regression, and rollback tests using temporary injected state roots, clocks, providers, collectors, and fixtures. No test may require the real home directory, live host collection, ambient state, real service, privilege, network, or sleep.


## Out of Scope

Task 034 shall not implement or redesign:

- Scheduler state, schedules, next-run state, job persistence, monitoring loops, automatic refresh, polling, daemon work, watchdog, heartbeat, service discovery, systemd/init installation, supervision, or restart policy;
- Alert history, incident persistence, acknowledgement/suppression persistence, Notification queues, delivery attempts, provider state, Runtime history, Runtime Service history, audit history, reports, trends, or long-term observations;
- a general database, object store, repository abstraction, write-ahead log, event log, multi-record history, retention policy, backup service, migration framework, replication, remote storage, cloud synchronization, multi-user state, fleet state, locking service, or distributed coordination;
- new Health, Rule, Policy, Alert, Notification, Runtime, Runtime Service, Configuration, Scheduler, Report, comparison, drift, or inventory-collection decisions;
- redefining `check` as Health or Report, changing existing Command stage dependencies, adding a second Pipeline, changing canonical profile output schemas, or treating Inventory success as healthy server status;
- persisting raw `command.Execution`, arbitrary `StageResult.Value`, unvalidated interface values, raw collector output, secrets, credentials, environment contents, provider payloads, terminal output, keystrokes, or errors;
- Console condition/recommendation logic, interface-specific evidence interpretation, duplicate Operator Overview semantics, REST API, Web Dashboard, remote access, remediation, shell execution, AI, packaging, deployment, installation, release, staging, commit, push, fetch, branch, or tag operations.

Task 034 is one local current-record persistence/recovery slice. Broader product persistence and operational continuity remain separately authorized Version 1.0 work.


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

- verify the exact QWSG root, repository markers, Framework 1.x configuration, canonical remote, `main` branch, HEAD, ahead/behind, complete Git status, empty index, ownership, modes, and ACLs;
- require canonical idle with Task 033 as the unique latest completed archive/history pair, no active prompt, and no Task 034 prompt/history/archive collision;
- run `bin/job --check`, `ai/scripts/next-task.sh --check`, `bin/job --check-test-tasks`, `ai/scripts/framework-check.sh`, Builder tests, Framework tests, configured engineering tests, Go tests, vet, format, and Git diff checks;
- preserve every pre-existing Owner-owned modified and untracked file exactly, including Tasks 025–033, QWCS work, Builder inputs, backups, code, tests, and documentation;
- reproduce by contract inspection and deterministic tests that `check` resolves only Inventory and Snapshot, its typed values remain process-local without an explicit store, bare Console starts from a missing Overview, and the current refresh provider runs `status` and supplies only `CommandObservation`;
- confirm there is no existing Current Operator State contract/store, no equivalent durable current Overview, and no architecture-owned freshness requalification seam;
- confirm Task 032 already owns the necessary operator semantics and Task 033 already renders validated Overview values, so no duplicate presentation layer is required;
- inspect current private atomic Inventory store techniques for reusable safety patterns without coupling the new single-record store to snapshot retention or copying a general repository abstraction;
- record the Release Minimalism decision: process-boundary current state is mandatory because otherwise a successful observation is inaccessible to the next ordinary-user process; general persistence remains deferred;
- validate this Builder source as readable non-symlink UTF-8 data with no NUL, unfinished fence, placeholder, embedded approval protocol value, unresolved choice, competing lifecycle, or ambiguous Builder field mapping.

### Separately authorized Builder installation

- repeat canonical idle, Task 033 baseline, Task 034 destination absence, source hash/content, Framework, repository/Git, permissions, ACL, Builder interface, and lifecycle checks immediately before installation;
- create and verify a bounded external Builder-installation snapshot of exact lifecycle targets, their absence, source/hash evidence, repository identity, Task 033 baseline, complete Git state, ownership, modes, ACLs, and restore instructions;
- map only the owner-authored Builder fields into a mode-0700 temporary input directory; keep generated metadata, fixed Required Reading, approval prose, and approval protocol data separate;
- run canonical read-only Builder input validation with separately supplied owner approval protocol data, then install exactly one Task 034 prompt/history pair only after separate explicit Owner authorization;
- require Task 034 as the sole active approved task, Task 033 as latest completed baseline, empty index, preserved unrelated status, and all Builder/lifecycle/Framework/repository checks passing;
- stop at the installed approved state and do not begin implementation during installation.

### Implementation starting state

- start only through an explicit canonical `job` invocation after Builder installation; read the entire active prompt, history, skill, and every Required Reading item as data;
- require Task 034 as the sole active approved task and Task 033 as the unique completed baseline;
- reverify the exact `check` plan and typed `StageResult` values, Pipeline validation, presentationmodel validators/recommendations/freshness rules, OverviewProvider, Console startup/session paths, CLI output/exit contracts, state path inputs, and atomic filesystem primitives available in the Go standard library;
- verify target absence for `internal/operatorstate` and `docs/architecture/CANONICAL_CURRENT_OPERATOR_STATE.md`, and identify every overlapping Owner-owned target before editing;
- prove tests can inject state root, clock, collector, Pipeline/provider, process environment, failures, and filesystem operations without live host, real home, privilege, network, or sleep;
- create and verify the proportional implementation snapshot before modifying any target;
- stop on any material need to redefine Command profiles, add engineering decisions outside presentationmodel, persist arbitrary interface payloads, weaken privacy/path safety, or expand into general persistence.

The valid preparation/install baseline is canonical idle after Task 033. The valid post-install state is one active approved Task 034. The valid completion state is canonical idle with Task 034 as the unique latest completed archive/history pair and no successor.


## Snapshot Requirements

Task preparation modifies only `current-task-job.txt`; no Builder-installation or implementation snapshot is created during preparation. The previous Task 033 source is recoverable from its archived prompt/history, and the pre-edit source hash and status shall be recorded.

Before separately authorized Builder installation, create one unique external mode-0700 snapshot under `/tmp` covering exact Task 034 prompt/history/archive destinations and absence, `current-task-job.txt`, source hash, repository identity, Task 033 idle baseline, complete Git status, empty staged list, ownership, modes, ACLs, checksum manifest, payload readability, collision guards, retention, and bounded restore instructions. Verify every checksum and absence record before installation.

Before implementation, create one unique external rollback snapshot for every existing directly affected Command, Pipeline, presentationmodel, Console/application, test, Makefile if changed, README, English/Hungarian user guide, Product Architecture, Functional Specification, Roadmap, System Map, architecture document, prompt, history, and archive target. Record verified absence for `internal/operatorstate`, its tests, the Current State architecture document, and any new user document. Preserve exact working-tree bytes rather than HEAD content. Record Git identity/state, permissions, ownership, ACLs, target inventory, hashes, collision guards, restore preconditions, and exact removal identities for created paths.

Snapshots shall exclude broad repository archives, build caches, real user state, live host evidence, terminal input, secrets, credentials, networks, processes, services, and unrelated content. Retain both installation and implementation snapshots through validation and Owner acceptance.


## Risk Assessment

- **False health claim — high:** Inventory/Snapshot success is not Health. Mitigate with explicit coverage, unchanged `check` profile, typed validation, and Presentation Model-owned unknown/partial semantics.
- **Lost type identity — high:** JSON-decoding `StageResult.Value` yields untyped maps. Extract only in-process concrete typed results and persist the validated Overview, never arbitrary execution payloads.
- **Stale state shown as current — high:** store an exact freshness deadline and require presentationmodel-owned requalification using an injected current time with boundary tests.
- **Corruption or partial replacement — high:** strict envelope validation, digest/identity recomputation, bounded IO, same-directory sync/rename/directory-sync publication, last-valid preservation, and fault-injection tests.
- **Unsafe local path or disclosure — high:** private per-user directory/file modes, owner and symlink checks, injected roots, bounded diagnostics, no raw paths/content/errors in Console output, and no secrets in state.
- **Profile compatibility regression — medium:** preserve the existing `check` Definition/Plan/output; publication is an application side effect only after validated success and is explicitly documented/tested.
- **Publication falsely coupled to rendering — medium:** publish after canonical projection and before presentation; broken stdout shall not corrupt state, while publication failure is separately observable.
- **Task 033 reopening — medium:** change only provider/application seams required to load/publish Overview; keep `internal/operatorconsole` Overview-only and audit imports.
- **Over-generalization — medium:** one current envelope, no history/retention/database/queue/monitor; audit source and documentation for excluded infrastructure.
- **Owner-owned dirty targets — medium:** exact working-tree snapshot, focused patches, empty index, target/status diff review, and no broad Git recovery.
- **Platform durability variance — medium:** use documented POSIX/Linux atomic rename and fsync behavior, injectable filesystem failure seams, and explicitly fail closed on unsupported operations.

Overall risk is medium-high because the task creates the first durable product-state boundary. Scope remains narrow and read-only at consumption, with no service, monitoring, network, privilege, or remediation authority.


## Planned Work

### Phase 1 — Confirm contracts and publish policy

- complete implementation starting checks and snapshot;
- freeze the existing `check` plan/output compatibility and document that complete execution is not a health verdict;
- define eligible coverage, typed extraction/correlation rules, publication timing, failure/exit behavior, and default/injected state location;
- specify Current Operator State 1.0 and freshness requalification tables before code.

### Phase 2 — Current State contract and private atomic store

- implement deterministic envelope normalization, canonical serialization, IDs/digests, strict validation/decoding, bounds, and typed errors;
- implement private path resolution, load, same-directory atomic replacement, sync, cleanup, and crash/fault behavior;
- add exhaustive contract/store tests for deterministic bytes, tampering, versions, paths, modes, ownership where testable, symlinks, truncation, oversize, rename/sync failures, and last-valid preservation.

### Phase 3 — Typed projection and time semantics

- add the smallest presentationmodel Inventory/coverage observation needed for honest `check` projection;
- implement application extraction of validated in-process Inventory/Snapshot results and correlation to Definition/Plan/Execution;
- add presentationmodel-owned stored-Overview freshness requalification that can only preserve or degrade state;
- test exact fresh-until boundary, stale transitions, missing evidence, incomplete coverage, IDs, recommendations, and impossibility of upgrading stale/partial evidence.

### Phase 4 — Publish and consume across processes

- publish after eligible successful `check` projection without changing the canonical plan or rendering contract;
- load through the OverviewProvider before bare interactive/non-interactive rendering;
- make explicit Console refresh use the same one-shot check/project/publish path;
- add subprocess tests proving writer termination followed by independent reader success, plus missing/stale/corrupt/incompatible/private-path behavior and preservation of previous valid state after failed publication.

### Phase 5 — Verify, document, and close

- run focused/full/race/vet/format, Framework, Builder, lifecycle, diversion, CLI regression, import/source, privacy, permissions/ACL, Git diff, snapshot, and rollback validations;
- document representation choice, profile semantics, state path, atomicity, failure behavior, privacy, freshness, Console boundary, and deferred persistence;
- finalize Task 034 history, mark complete only after every mandatory gate, archive without a successor, and verify canonical idle.


## Rollback Plan

Rollback is exact, file-bounded, identity-checked, and requires Owner confirmation before overwriting material post-snapshot work.

Record the failing/current diff and validation evidence; stop active test processes; verify snapshot manifests and hashes; confirm repository root, lifecycle, exact target inventory, and that no later Owner edits exist. Restore only pre-existing targets from their exact working-tree snapshot payloads with recorded modes and ACLs where authorized. Remove only Task 034-created paths whose pre-change absence and current Task 034 identity are both proven. Never use wildcard deletion, recursive repository cleanup, broad reset, clean, checkout, restore, or replacement from HEAD.

Test/product state created under temporary injected roots may be removed only by exact resolved path after validating the root belongs to the Task 034 test. Real user Current State must never be deleted by engineering rollback; if a manual acceptance record was created, preserve it or obtain explicit Owner authority for its exact removal.

After rollback, verify checksums, presence/absence, ownership, modes, ACLs, complete Git status, empty index, Framework/lifecycle/diverted-task state, original `check` output/plan, original bare Console missing behavior, full Go tests/race/vet/format, engineering tests, and Git diff checks. Retain snapshot, failed-work diff, restore log, and validation report for Owner review.


## Deliverables

- Canonical Current Operator State 1.0 envelope, validation, canonical serialization, identity, integrity, provenance, coverage, and bounded error contracts;
- private single-record atomic store with deterministic path resolution, validated loading, crash-safe replacement, and fail-closed recovery;
- typed `check` Inventory/Snapshot extraction and projection seam without persisting `any` values;
- minimal presentationmodel Inventory/coverage input and freshness requalification owned by the Operator Presentation Model;
- compatible `check` publication and explicit Console refresh publication paths;
- bare interactive and non-interactive Current State loading through OverviewProvider;
- exact missing, fresh, boundary, stale, corrupt, incompatible, unsafe, unreadable, failed-publication, and last-valid behaviors;
- deterministic unit, fault-injection, integration, subprocess, CLI regression, boundary/import, privacy, resource, permission, atomicity, rollback, and lifecycle tests;
- `docs/architecture/CANONICAL_CURRENT_OPERATOR_STATE.md`;
- directly affected README, English/Hungarian user guidance, Product Architecture, Functional Specification, Command/Operator Model/Console architecture, Architecture, System Map, Roadmap, Engineering History, prompt/history/archive records;
- verified snapshots, completion evidence, archive, and canonical idle closure.

No general persistence, history, monitoring, Scheduler, Alert/Notification/Runtime history, database, API, Dashboard, service, remediation, remote, packaging, deployment, release, stage, commit, or push is a deliverable.


## Verification

Builder/lifecycle validation shall prove exact Task 034 metadata and slug, Task 033 baseline, mandatory sections, lossless Builder mapping, separate approval protocol data, no placeholder/embedded approval, correct pre-install idle/post-install sole-active/post-completion idle states, and no successor.

Implementation validation shall include:

- focused `internal/operatorstate`, `internal/presentationmodel`, provider/application, Console boundary, and `cmd/qwsg` tests;
- `make build`, full `go test ./...`, repository-wide race tests with configured writable caches, vet, all Go formatting, Framework 1.x, engineering tests, Builder tests, lifecycle checks, diverted-task audit, and Git diff checks;
- exact tests proving `check` still resolves only Inventory/Snapshot and existing help, human/JSON output, profiles, advanced composition, store selection, exit behavior, and Command/Pipeline identity remain compatible;
- tests proving only validated complete `check` coverage publishes; technical completion with missing/wrong/untyped/duplicate/tampered stages cannot publish or replace state;
- tests proving Inventory success never becomes healthy, and absent Health/Policy/Alert/Runtime/Service evidence remains explicit;
- canonical serialization golden tests, repeated-byte determinism, stable IDs/digests, strict unknown/trailing rejection, tamper detection, version incompatibility, timestamp/coverage/provenance correlation, and maximum-size/count/token limits;
- filesystem tests for clean absolute paths, default/injected resolution, private `0700/0600` modes, symlink/non-directory/wrong-mode/permission rejection, exclusive temporary creation, complete write, fsync/close/rename/directory-sync ordering, cleanup, and collision handling;
- failure injection before and after every publication boundary proving no partial state is loadable and the last valid state survives every failed pre-rename replacement;
- exact freshness tests for before, at, and after the exclusive deadline; supplied-clock determinism; stale state visibility; missing distinction; no stale-to-current upgrade; and model-owned recommendation/condition recalculation;
- separate-process acceptance tests: process A runs eligible fresh `check` against deterministic injected collection/state root and exits; process B starts bare `qwsg`, loads the state, and displays current limited coverage without unavailable/missing or a false healthy verdict;
- subprocess variations for missing, stale, truncated, digest-corrupt, identity-corrupt, unsupported-schema/model, unsafe-path, unreadable, and oversize state with deterministic bounded output/exit behavior;
- interactive explicit-refresh tests proving one call, one eligible execution, one publication, zero startup collection, zero polling/retry/background work, retention of last Overview/state on failure, and cancellation propagation;
- import/source audits proving Console remains presentationmodel-only and contains no store/Command/Pipeline/Inventory/engine/persistence logic; store contains no presentation semantics; presentationmodel owns all semantic requalification;
- privacy tests proving state excludes secrets, raw environment, raw host output, credentials, arbitrary errors, terminal content, and unbounded paths; Console diagnostics expose only localized bounded classifications;
- resource tests for file size, read/write bytes, tokens, provenance, sources, temporary files, attempts, and no accumulating history;
- documentation consistency and Release Minimalism review against Product Architecture, Functional Specification, Command Architecture, Operator Presentation Model, Console, Runtime/Service, Roadmap, and System Map;
- exact changed-target, ownership, mode, ACL, index, complete status, unrelated Owner-content preservation, snapshot checksum, restore feasibility, and confirmation that nothing was installed, monitored, transmitted, staged, committed, pushed, packaged, deployed, or released.

All verification uses deterministic fixtures and temporary injected roots. It requires no real home state, live host, real clock sleep, terminal, service, systemd, database, network, credential, privilege, or infrastructure mutation.


## Documentation Updates

Create `docs/architecture/CANONICAL_CURRENT_OPERATOR_STATE.md` and directly update:

- `docs/architecture/CANONICAL_COMMAND_ARCHITECTURE.md` for unchanged `check` plan and explicit publication side effect;
- `docs/architecture/CANONICAL_OPERATOR_PRESENTATION_MODEL.md` for Inventory coverage and model-owned freshness requalification;
- `docs/architecture/INTERACTIVE_OPERATOR_CONSOLE.md` for Current State consumption and unchanged Overview-only boundary;
- `docs/PRODUCT_ARCHITECTURE.md` and `docs/FUNCTIONAL_SPECIFICATION.md` for process-boundary current observation behavior without claiming general persistence or health;
- `ai/core/04_ARCHITECTURE.md`, `ai/core/05_SYSTEM_MAP.md`, `ai/core/07_ENGINEERING_HISTORY.md`, and `ai/core/13_ROADMAP.md`;
- `README.md` and directly affected English/Hungarian Console/CLI guidance;
- active prompt, independent Task 034 history, and completed archive.

Documentation shall state:

`Canonical execution -> typed Overview projection -> Canonical Current Operator State -> OverviewProvider -> Canonical Operator Presentation Model freshness requalification -> replaceable interface`.

It shall state that `check` remains Inventory/Snapshot, publication records limited coverage rather than Health, stored data is one validated Overview envelope rather than raw stage values, bare Console never collects automatically, stale/corrupt/incompatible states fail visibly and safely, and all general persistence/monitoring/history/service/API work remains separate. Record every actual update and justified omission in Task 034 history.


## Completion Criteria

Task 034 is complete only when:

- one versioned Canonical Current Operator State contract stores a validated canonical Overview with stable identity, provenance, explicit coverage, observation/publication/freshness times, integrity, canonical bytes, and strict bounds;
- publication is atomic, private, crash-safe, same-directory, sync-backed, last-valid-preserving, symlink-safe, and deterministically rejects corruption/incompatibility;
- existing `check` Command Definition/Plan remains Inventory/Snapshot and compatibility tests pass;
- eligible successful `check` uses exact typed outputs to project and publish honest limited coverage, while failed/incomplete/tampered/untyped executions never publish;
- Inventory success is never labeled healthy and absent Health, Policy, Alert, Runtime, Runtime Service, and Guardian evidence remains explicit;
- a separately started bare `qwsg` loads the last valid state through OverviewProvider, requalifies freshness only through presentationmodel, and renders meaningful fresh or stale status without Console reinterpretation;
- missing state remains unavailable with fresh-check guidance; stale state remains visible; corrupt, incompatible, unsafe, and unreadable state fail closed with bounded distinct diagnostics;
- explicit Console refresh uses the same one-shot check/project/publish path with zero automatic startup collection, polling, retry, persistence history, or background work;
- the Task 033 typed-output/provider gap is corrected only at the application projection/publication/consumption seam and `internal/operatorconsole` remains Overview-only;
- subprocess acceptance proves process A observation, termination, process B consumption, freshness transition, and every required failure mode;
- all focused, full, race, vet, format, Framework, Builder, lifecycle, diversion, compatibility, atomicity, corruption, freshness, privacy, resource, import, documentation, permission/ACL, Git, snapshot, and rollback checks pass;
- Release Minimalism confirms this single current record is required for a truthful usable Version 1.0 Console and no general persistence capability was introduced;
- no real user state was removed, no dependency/service/infrastructure was installed or mutated, and nothing was staged, committed, pushed, packaged, deployed, or released;
- Task 034 prompt/history are complete and archived without a successor, and `bin/job --check` reports canonical idle with Task 034 as the unique latest completed baseline.

A valid result is `complete`, `complete with disclosed limitations`, or `blocked`. Completion may not be claimed while publication eligibility, typed projection, atomic durability, freshness requalification, separate-process consumption, corruption behavior, Console boundary, compatibility, rollback, or lifecycle evidence remains unresolved.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-08 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
