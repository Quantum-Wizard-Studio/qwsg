# Current Engineering Task 033: Interactive Operator Console

## Task Metadata

- Task ID: `033`
- Task slug: `interactive-operator-console`
- Status: `complete`
- Date opened: `2026-08-07` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Interactive Operator Console


## Objective

Establish the Interactive Operator Console as the first immediately understandable, local operator-facing interface over the existing Canonical Operator Presentation Model.

Task 033 shall make bare `qwsg` useful to an ordinary Linux user within seconds while preserving every advanced explicit CLI command and machine-readable workflow. In an interactive terminal, bare `qwsg` shall open a keyboard-driven, terminal-safe Console whose home view answers only from a validated `presentationmodel.Overview`: server condition, attention, changes, Alert summary, Guardian state, evidence freshness/completeness, and recommended next step. In a non-interactive context, bare `qwsg` shall emit a deterministic concise non-full-screen overview derived from the same model.

The Console shall be presentation only. It shall accept Overview values through a narrow injected read-only provider and shall never import, call, or reinterpret Inventory, Comparison, Drift, Health, Rule, Policy, Configuration, Scheduler, Alert, Notification, Runtime, or Runtime Service logic. A separate application adapter may perform an explicit operator-requested refresh through the existing Canonical Command/Pipeline boundary and ask `internal/presentationmodel` to project the already-produced canonical outputs. No engineering decision may move into the Console or adapter.

Task 033 ends when Console Model 1.0, deterministic view composition, keyboard interaction, explicit refresh, localization catalogs, terminal capability/fallback behavior, bare-command selection, comprehensive injected tests, permanent architecture documentation, rollback evidence, and canonical lifecycle closure are complete. It does not create monitoring, persistence, restart recovery, a resident daemon, REST API, Web Dashboard, or production notification provider.


## Scope

Task 033 shall define and implement one presentation-only Console package, expected at `internal/operatorconsole`, with versioned Console Model 1.0 contracts for:

- an injected `OverviewProvider` whose only product data result is one validated `presentationmodel.Overview` or a bounded fixed failure classification;
- immutable Console Session State containing current screen, selection, viewport, terminal dimensions/capabilities, locale, refresh state, last accepted Overview identity, and bounded diagnostic token;
- normalized input actions for up, down, select, back, refresh, help, quit, resize, and unsupported input;
- deterministic screen models for Home, Attention, Changes, Guardian, Evidence/Source Details, and Help;
- stable view element, navigation, label, severity, freshness, completeness, recommendation, and diagnostic tokens with typed parameters;
- bounded selection, paging/viewport, resize, narrow-terminal fallback, monochrome behavior, and no-color semantics;
- deterministic transition rules, including empty collections, invalid selection after refresh, stale/partial/unavailable Overview, provider failure, cancellation, EOF, resize, and quit;
- validation, canonical ordering, strict model versions, terminal-safe text, resource bounds, and privacy limits.

The Console shall consume only `presentationmodel.Overview`. It may select, order, page, and render fields already present in that model, but it shall not change condition, attention, Guardian, freshness, completeness, change, Alert, or recommendation meaning. It shall never use canonical source IDs as the beginner Home summary; exact existing references may appear only in an explicitly selected advanced detail view.

The Console shall provide a line-oriented keyboard interaction that is fully testable with injected input/output and requires no raw-terminal dependency. Single-key commands followed by Enter are acceptable for navigation. ANSI cursor/screen control may be used only behind an explicit terminal-capability adapter; unsupported, dumb, constrained, or non-interactive terminals shall receive a deterministic plain-text fallback. No color may be the sole carrier of meaning.

The selected information architecture is:

1. Home: condition, attention, Guardian, freshness/completeness, summary counts, and first recommended next step;
2. Attention: bounded ordered attention items from the Overview;
3. Changes: bounded category-level change summaries from the Overview;
4. Guardian: model-owned Guardian and operational evidence state only;
5. Evidence/Source Details: copyable model source references and timestamps for advanced users;
6. Help: stable keys and navigation descriptions from localization tokens.

Task 033 shall add localization-ready rendering with built-in English and Hungarian test catalogs. User-visible text shall live in catalogs keyed by stable Console and Operator Model tokens; commands, keys, schema values, and canonical identifiers remain language-independent. Unknown required tokens fail closed in tests and render a bounded safe fallback in production without exposing raw errors.

Refresh shall be explicit operator intent. The Console calls only its injected provider. The production local provider, owned by the `cmd/qwsg` application adapter rather than the Console, may:

- execute one existing bounded canonical live-analysis Command through the existing Command Definition, Plan, and Pipeline contracts;
- extract only typed stage outputs already present in the validated Command Execution through a narrow presentationmodel-owned or presentationmodel-validated projection seam;
- call `presentationmodel.Project` with explicit observation time and freshness policy;
- return the validated Overview to the Console.

The provider shall not define an engine order, Command profile, Health mapping, Alert rule, Runtime behavior, retry, polling interval, cache, monitoring loop, or persistent state. It shall perform no refresh until the operator explicitly requests one. Provider failure shall produce a bounded Console diagnostic while retaining the last validated Overview if one exists.

Bare-command behavior shall be:

- interactive terminal input and output: launch the Console without automatically collecting or persisting host evidence; show an initial validated unavailable/not-observed Overview and make explicit refresh discoverable;
- non-interactive input or output: emit one concise deterministic plain-text unavailable/not-observed Overview unless an explicitly supplied validated Overview source is available; never start a full-screen or input-waiting session;
- explicit `qwsg help`, `version`, `inventory`, `compare`, `status`, `check`, `changes`, `health`, `report`, and `analyze` behavior: remain unchanged and retain existing exit, JSON, and advanced composition contracts;
- explicit `--help`/help behavior: remain help, not Console launch.

Task 033 may add an explicit `console` command as the deterministic way to request Console behavior and test terminal selection. Bare `qwsg` and `qwsg console` shall resolve to the same Console application contract; non-interactive safety remains mandatory.

No hidden environment default may change engineering semantics. Locale, color/capability, and terminal dimensions are presentation inputs only. Refresh source, store, or other engineering selection must use existing explicit Command selections or documented existing environment adapters and must become explicit data before Command normalization.


## Out of Scope

Task 033 shall not implement:

- Inventory collection logic, Snapshot persistence logic, Comparison, Drift, Health, Rule, Policy, Report, Configuration, Scheduler, Alert, Notification, Runtime, or Runtime Service decisions;
- a second Command Definition, profile registry, parser, planner, Pipeline, engine sequence, typed evidence model, Operator Overview, condition/attention mapping, Alert summary, freshness/completeness policy, Guardian status, or recommendation taxonomy;
- automatic refresh, polling, periodic scheduling, background monitoring, watchdog, heartbeat, goroutine worker, daemon, resident process, automatic startup, systemd/init integration, service installation, supervision, or process discovery;
- durable Console state, Overview cache, history, Alert/Notification state, Runtime/Service state, configuration, evidence, audit, database, checkpoint, replay, migration, restart recovery, or retention behavior;
- REST API, HTTP listener, WebSocket, Web Dashboard, browser Console, desktop GUI, public application server, remote access, remote control, or fleet interface;
- concrete Email, SMTP, Webhook, Slack, Discord, Telegram, SMS, push, desktop, or other notification provider;
- Alert acknowledgement or suppression mutation, maintenance editing, configuration editing/activation/reload, scheduler control, service control, install/update/remove, remediation, repair, shell execution, host mutation, remote execution, AI, machine learning, licensing, packaging, deployment, or release;
- authentication, authorization, sessions, TLS, network security, multi-user Console, or privileged helper;
- raw-terminal dependency installation, third-party TUI framework, uncontrolled ANSI output, alternate-screen requirement, mouse handling, animation, graphs, themes, plugin system, or optional visual polish;
- dependency installation, infrastructure mutation, staging, commit, push, fetch, branch, or tag operations.

Task 033 is the local read-only presentation slice only. Persistence/restart recovery, monitoring/product health, production notification transport, installation/supervision, REST API, Web Dashboard, and supported release evidence remain separately authorized Version 1.0 gates.


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
- run `bin/job --check`, `ai/scripts/next-task.sh --check`, `bin/job --check-test-tasks`, and `ai/scripts/framework-check.sh`; require canonical idle with Task 032 as the unique latest completed archive/history pair;
- verify `ai/prompts/` is empty and Task 033 prompt, history, archive, `internal/operatorconsole`, tests, and `docs/architecture/INTERACTIVE_OPERATOR_CONSOLE.md` are absent;
- preserve every pre-existing unstaged and untracked Owner-owned change from Tasks 025–032, QWCS work, Builder sources, backups, and `current-task-job.txt`; require the index to remain empty;
- inspect the actual Operator Presentation Model, Command/Pipeline, current presentation package, CLI process adapter, Product Architecture, Functional Specification, Roadmap, and System Map;
- confirm the Operator Model already owns the complete beginner status semantics and that no existing Console or equivalent application interface exists;
- confirm the Console can be implemented with standard-library injected terminal boundaries and no unavailable dependency, real terminal, live service, provider, persistence, network, privilege, or infrastructure requirement;
- confirm the application refresh adapter can obtain typed canonical execution outputs and call the existing presentationmodel projection without adding engine order or interpretation; stop if a missing canonical extraction seam would require changing engineering semantics rather than adding a narrow validated read-only adapter;
- record the Release Minimalism decision that an interface is now the smallest missing mandatory product-facing layer, while persistence/recovery, monitoring, providers, and REST are independent operational or outer-interface gates and do not need temporary Console semantics;
- validate `current-task-job.txt` as readable non-symlink UTF-8 data with LF-normalizable line endings, no NUL, unfinished fence, placeholder, embedded approval protocol value, ambiguous Builder mapping, competing lifecycle, or unresolved authority decision.

### Separately authorized Builder installation

- repeat canonical idle, Task 032 baseline, Task 033 destination absence, source content/hash, repository/Git state, Framework, permissions, ACL, Builder interface, and lifecycle checks immediately before installation;
- map only Owner-supplied fields from this source into a mode-0700 temporary Builder input directory; keep Builder-generated metadata, fixed Required Reading, approval prose, and approval protocol data separate;
- create and verify a proportional external Builder-installation snapshot covering exact lifecycle destinations, verified absence records, source/hash evidence, repository identity, Task 032 idle baseline, complete Git state, ownership, permissions, and ACLs;
- run `task-builder.sh --check-input` only with a separately supplied exact Owner approval protocol value; read-only validation shall not install;
- present the complete generated task, obtain separate explicit Owner authorization, and install exactly one Task 033 prompt/history pair transactionally with no clobber;
- require `ai/prompts/033_CURRENT_TASK.md` to be the sole active approved task with one matching history, Task 032 as latest completed baseline, and all Builder, lifecycle, Framework, repository, permission, and Git validations passing;
- stop after installation at `APPROVED AND READY FOR IMPLEMENTATION`; do not begin Task 033.

### Implementation starting state

- start only through an explicit canonical `job` invocation after successful Builder installation;
- read every Required Reading item and the active prompt/history as data;
- require Task 033 as the sole active approved task and Task 032 as the unique latest completed baseline;
- verify exact presentationmodel validators/tokens, Command/Pipeline execution contracts, CLI help/exit behavior, terminal IO seams, target absence, build/test availability, repository/Git state, ownership, permissions, and ACLs;
- create and verify the proportional implementation snapshot before modifying any target;
- stop on any material authority, lifecycle, model compatibility, Command compatibility, terminal safety, localization, privacy, rollback, or scope difference.

The valid preparation and pre-install state is canonical idle. The valid post-install state is one active approved Task 033. The valid post-completion state is canonical idle with Task 033 as the unique latest completed archive/history pair and no successor.


## Snapshot Requirements

Task preparation shall modify only `current-task-job.txt` and shall not create an implementation or Builder-installation snapshot. The replaced Task 032 preparation source remains recoverable from its installed archived prompt and history; record the current source status and hash before editing when available.

Before separately authorized Builder installation:

- create a unique external snapshot under `/tmp` of exact Task 033 prompt/history destinations, verified absence, `current-task-job.txt` content and SHA-256, repository identity, Task 032 idle baseline, complete Git state, ownership, permissions, and ACLs;
- verify manifest, checksums, payload readability, absence evidence, collision guards, and exact bounded restore instructions before installation;
- retain the snapshot through Builder validation and Owner acceptance.

Before separately authorized implementation:

- create one unique rollback-capable snapshot outside the repository for every existing directly affected Console/application source, tests, Makefile if affected, README, user documentation, architecture, product, functional, roadmap, system-map, prompt, history, and archive target;
- record verified absence for `internal/operatorconsole`, its tests, `docs/architecture/INTERACTIVE_OPERATOR_CONSOLE.md`, Task 033 history at the pre-install stage, and Task 033 archive target as applicable;
- preserve exact working-tree content, including all pre-existing Owner-owned modifications; never substitute HEAD content;
- capture repository identity, branch, HEAD, remotes, ahead/behind, complete Git state, staged paths, target inventory, ownership, permissions, ACLs, baseline validation, manifest, SHA-256 checksums, readable payload inventory, and guarded restore instructions;
- verify every checksum, payload, absence record, collision guard, restore precondition, and proportional target before implementation;
- retain the snapshot through completion and Owner acceptance.

Snapshot scope shall exclude broad repository archives, build caches, live terminal state, raw keystrokes beyond bounded test fixtures, live host evidence, running processes, systemd state, product stores, provider payloads, secrets, credentials, network responses, and unrelated data.


## Risk Assessment

Primary risks and mandatory mitigations:

- The Console could reinterpret engineering evidence. Accept only validated `presentationmodel.Overview` and render its exact states, counts, tokens, references, and order; prohibit all canonical engine imports from `internal/operatorconsole`.
- Refresh could turn the Console into an orchestration engine. Use an injected OverviewProvider; keep the production one-shot adapter outside the Console and require existing Command/Pipeline plus presentationmodel projection boundaries.
- Bare `qwsg` could unexpectedly block scripts. Detect interactive capability through an injected process/terminal adapter; non-interactive bare invocation emits concise output and never waits for input or enters full-screen mode.
- Bare launch could perform unauthorized collection. Initial Console state uses a validated unavailable/not-observed Overview; live analysis occurs only after the operator explicitly requests refresh.
- A terminal dependency could block execution readiness. Use standard-library line-oriented IO and injected capabilities; no raw-terminal library or third-party framework is required.
- ANSI sequences could corrupt dumb terminals or logs. Gate screen control by explicit capability and provide plain-text fallback; escape control characters and prohibit arbitrary upstream text.
- Keyboard behavior could become nondeterministic. Normalize a closed action set and test every state transition, EOF, cancellation, resize, invalid input, and selection boundary.
- Refresh failure could discard trusted evidence. Retain the last validated Overview, show a fixed bounded diagnostic token, and never merge partial provider payloads.
- Localization could hardcode user text or change machine semantics. Resolve stable tokens through complete catalogs; keep command keys and canonical values stable and test English plus Hungarian expansion.
- Advanced source references could overwhelm beginners or expose sensitive values. Keep IDs out of Home and summary screens; show only existing bounded model references in an explicit detail view.
- Narrow terminals and accessibility could rely on color or layout. Provide labels, monochrome output, deterministic wrapping/truncation, and plain-text fallback with complete navigation equivalents.
- Console navigation might imply mutation authority. Task 033 includes read-only navigation and explicit refresh only; no acknowledgement, suppression, configuration, service control, or remediation action exists.
- The application adapter could invent freshness or analysis defaults. Pass explicit observation time and bounded documented freshness input to presentationmodel and use existing Command selection semantics; no hidden engine or policy default is added.
- Existing CLI behavior could regress. Preserve every explicit command, help topic, JSON contract, exit behavior, environment selection, and advanced composition test; isolate bare-command selection.
- Output or session state could grow without bound. Bound screens, viewport, terminal dimensions, selections, tokens, catalog values, rendered lines/columns, input length, diagnostics, sources, and refresh attempts per interaction.
- Tests could require an ambient terminal or live collection. Use fake provider, terminal, input, output, clock/observation values, capability, and resize sources; integration tests use injected canonical fixtures and no live host.
- Repository targets overlap Owner-owned dirty files. Snapshot exact working-tree versions, use focused patches, review every target diff, keep the index empty, and prove unrelated status preservation.
- Scope could expand into monitoring, persistence, REST, providers, installation, or polish. Enforce import/source audits and explicit exclusions; stop on any required new authority.

Overall implementation risk is medium because bare-command behavior and terminal interaction are public UX contracts. Engineering-truth, mutation, privilege, network, and persistence risk remain low because the Console is read-only and receives only the canonical Operator Model.


## Planned Work

### Phase 1 — Verify interface and data boundaries

- complete implementation Starting State Verification and snapshot requirements;
- inventory exact Overview fields/tokens, existing presentation helpers, CLI routing/help behavior, Command/Pipeline typed outputs, and test seams;
- publish the dependency boundary `Canonical Data -> Operator Model -> OverviewProvider -> Console -> Terminal` and prove the Console imports only presentationmodel plus standard library;
- define the one-shot application refresh adapter and confirm it adds no canonical decision.

### Phase 2 — Define Console Model 1.0

- define Session State, screens, normalized actions, terminal capabilities, viewport, locale/catalog, provider result, diagnostic, and transition contracts;
- define Home, Attention, Changes, Guardian, Evidence/Details, and Help information architecture;
- define exact keyboard map, resize/fallback behavior, localization tokens, privacy/resource bounds, and compatibility policy;
- document interactive versus non-interactive bare-command selection and exit behavior.

### Phase 3 — Implement pure state and rendering

- implement deterministic state transitions and selection/viewport normalization;
- implement catalog-driven screen composition from validated Overview only;
- implement terminal-safe line-oriented renderer and capability-gated redraw adapter;
- implement injected input/output/provider loop with cancellation, EOF, refresh, failure retention, and quit behavior;
- introduce no ambient terminal, live service, persistence, network, or engine dependency.

### Phase 4 — Integrate the local application boundary

- add the smallest presentationmodel-owned or validated typed Command Execution extraction seam only if existing public contracts require it;
- implement the explicit one-shot refresh provider through existing Command/Pipeline APIs and presentationmodel projection;
- route interactive bare `qwsg` and explicit `qwsg console` to the Console; route non-interactive bare invocation to concise plain-text rendering without blocking;
- preserve all explicit current commands, help, JSON, profiles, advanced composition, and exit semantics.

### Phase 5 — Verify, document, and close

- add focused model, transition, rendering, localization, terminal-safety, refresh, failure, cancellation, capability, resource, privacy, CLI-routing, and regression tests using injected fixtures;
- update permanent architecture, product/functional, roadmap/system-map, README, CLI user documentation, prompt/history/archive, and chronological history records only where directly affected;
- run every mandatory verification, inspect exact diffs/permissions/ACLs, reverify snapshot/rollback, archive Task 033 without a successor, and restore canonical idle.

If a production refresh cannot be built without defining new engineering semantics or adding monitoring/persistence, implement only the provider interface and deterministic unavailable Console, record the limitation, and stop before claiming full completion. Such a material difference requires Owner review rather than hidden scope expansion.


## Rollback Plan

Rollback is exact, file-bounded, collision-aware, and preserves all pre-existing Owner-owned content.

Before rollback, stop Console/task processes, record the current Task 033 diff and validation evidence, verify snapshot manifests/checksums, confirm the exact repository root and target inventory, and obtain Project Owner confirmation before overwriting material post-snapshot work.

Restore only listed pre-existing targets from their exact working-tree snapshot payloads, preserving recorded modes and ownership where authorized. Remove only Task 033-created paths whose absence was recorded and only after verifying their identity and that they contain no later Owner work. Never use wildcards, recursive repository cleanup, broad `git reset`, `git clean`, checkout, restore, or deletion of unrelated untracked content.

If CLI routing was changed, restore only the exact snapshotted `cmd/qwsg` files and tests. If a presentationmodel extraction seam was added, restore only exact snapshotted presentationmodel files. For lifecycle rollback, restore the exact Task 033 prompt/history working copies and remove the Task 033 archive only when the snapshot proves absence and identity checks prove Task 033 created it.

After rollback, verify checksums, target absence/presence, permissions, ownership, ACLs, complete Git status, empty index, Framework and lifecycle identity, diverted-task audit, bare/explicit CLI baseline behavior, full Go tests, race where relevant, vet, format, and Git diff checks. Retain the snapshot, failed-work diff, manifest, checksums, and rollback report for Owner review.

Rollback shall not remove or alter pre-existing Tasks 025–032, QWCS, Builder, backup, source, documentation, test, or other Owner-owned changes outside the exact Task 033 manifest.


## Deliverables

- Interactive Operator Console and Console Model 1.0;
- OverviewProvider, Session State, screen, action, terminal-capability, viewport, locale/catalog, diagnostic, and transition contracts;
- deterministic Home, Attention, Changes, Guardian, Evidence/Details, and Help views derived only from validated Operator Overview;
- line-oriented keyboard navigation, explicit refresh, selection, paging, back/help/quit, resize normalization, terminal-safe redraw, and plain-text fallback;
- complete English and Hungarian token catalogs with stable language-independent commands and model values;
- one-shot application refresh adapter using only existing Command/Pipeline and presentationmodel boundaries;
- interactive and non-interactive bare `qwsg` behavior plus explicit `qwsg console`, without regression to existing commands;
- strict terminal safety, privacy, localization, accessibility, resource, deterministic transition, and compatibility behavior;
- focused unit, golden, state-machine, provider, failure, terminal, localization, CLI, boundary, privacy, resource, and regression tests without live host/terminal/service/network dependencies;
- `docs/architecture/INTERACTIVE_OPERATOR_CONSOLE.md`;
- directly affected README, Product Architecture, Functional Specification, Roadmap, System Map, Command/Operator Model architecture, CLI user guidance, project Architecture/Engineering History, prompt/history/archive updates;
- complete Task 033 history, verified rollback evidence, archived completed prompt, and canonical idle closure.

No monitoring, automatic refresh, persistence, restart recovery, provider transport, Runtime/Service logic, REST API, Dashboard, remote capability, mutation, remediation, AI, package, deployment, or release artifact is a Task 033 deliverable.


## Verification

Builder and lifecycle verification shall prove:

- exact Task 033 ID/slug/title, authority, language, mandatory sections, unique Task 032 baseline, no destination collision, and no placeholder or unresolved content;
- explicit separation of task creation, Builder installation, implementation starting state, implementation work, completion validation, archive, and canonical idle;
- no embedded or inferred approval protocol value in `current-task-job.txt`;
- lossless mapping of every owner-authored Builder field and successful canonical read-only Builder input validation using only separately supplied approval protocol data;
- correct pre-install idle, post-install sole-active approved, and post-completion idle states with exact prompt/history/archive identity and no successor.

Implementation verification shall include:

- focused `internal/operatorconsole`, directly affected presentationmodel/presentation, and `cmd/qwsg` tests;
- `make build`, full `go test ./...`, repository-wide `go test -race ./...` with writable configured caches, `go vet ./...`, and complete Go formatting checks;
- Framework 1.x configured validations, `make engineering-test`, Builder tests, lifecycle checks, diverted-task audit, active-task validation, and final idle validation;
- golden tests for every screen in English and Hungarian, monochrome, narrow-terminal, dumb-terminal, interactive redraw, and non-interactive plain-text modes;
- state-machine table tests for every screen/action pair, up/down/select/back/refresh/help/quit, selection boundaries, empty lists, paging, viewport resize, unsupported key, excessive input, EOF, cancellation, and provider failure;
- tests proving Home answers condition, attention, change, Alert, Guardian, freshness/completeness, and recommendation questions only from exact Overview fields;
- tests proving critical/urgent, degraded/review, unknown, unavailable, stale, partial, unsupported, failed/stopped/not-observed Guardian, no-change, no-alert, and no-action states remain visibly distinct without color;
- refresh call-count/order tests proving zero calls before explicit refresh, one provider call per accepted refresh action, no polling/retry/background work, exact acceptance of only validated Overview, retention of last good Overview on failure, and cancellation propagation;
- application-provider tests proving one explicit bounded existing Command/Pipeline execution, typed validated projection through presentationmodel, no duplicated stage order or engine calls, and no direct Console dependency on any canonical engine;
- source/import audits proving `internal/operatorconsole` imports only presentationmodel and standard library and contains no Inventory, Comparison, Drift, Health, Rule, Policy, Configuration, Scheduler, Alert, Notification, Runtime, Runtime Service, Pipeline, collector, persistence, monitoring, provider transport, process, systemd, network, remediation, remote, or AI logic;
- tests proving interactive terminal launch, explicit `console`, non-interactive nonblocking fallback, `--help`, help topics, explicit commands, JSON, profiles, advanced composition, stdout/stderr separation, broken output, cancellation, and exit codes;
- terminal-safety tests rejecting/control-escaping ANSI/control injection from tokens, IDs, locale values, diagnostics, dimensions, and provider failures; ANSI control appears only from the trusted capability adapter;
- localization tests proving complete English/Hungarian catalogs, stable keys/commands/schema values, missing-token behavior, string expansion, Unicode width-safe fallback, and absence of user-facing prose outside catalogs;
- accessibility tests proving no color-only meaning, explicit text labels, deterministic focus/selection markers, plain-text equivalents, and bounded readable output at constrained sizes;
- privacy tests proving Home and ordinary views exclude canonical IDs, raw host values, paths, errors, configuration, destinations, provider payloads, credentials, environment values, and secrets; detail view exposes only existing presentationmodel SourceReference fields;
- resource tests for terminal rows/columns, rendered line/byte counts, input length, action count, selections, pages, catalog values, attention/change/source collections, diagnostics, and fail-closed overflow behavior;
- determinism tests proving equivalent Overview/session/capability/locale inputs produce byte-identical screen models and rendered output;
- documentation consistency review against Product Architecture, Functional Specification, Roadmap, System Map, Command Architecture, Operator Presentation Model, Runtime, and Runtime Service boundaries;
- exact changed-target audit, ownership, permissions, ACLs, staged/unstaged paths, `git diff --check`, `git diff --cached --check`, and preservation of unrelated Owner-owned content;
- snapshot checksum, payload, readability, absence evidence, collision guards, bounded restore verification, and confirmation rollback remains usable;
- confirmation that nothing was installed as a service, monitored, automatically refreshed, persisted as product state, transmitted, staged, committed, pushed, packaged, deployed, or released.

Verification requires no live host collection, real terminal/raw mode, ambient keyboard input, real clock sleep, real signal, running daemon, systemd/service manager, persistent product store, external provider, credential, network, privilege, remote system, or infrastructure mutation.


## Documentation Updates

Expected direct documentation targets are:

- `docs/architecture/INTERACTIVE_OPERATOR_CONSOLE.md`;
- `docs/architecture/CANONICAL_OPERATOR_PRESENTATION_MODEL.md` for the Console consumer and unchanged model ownership;
- `docs/architecture/CANONICAL_COMMAND_ARCHITECTURE.md` for the one-shot application provider and unchanged advanced command boundary;
- `docs/PRODUCT_ARCHITECTURE.md` for the local Console and bare-command behavior;
- `docs/FUNCTIONAL_SPECIFICATION.md` for observable read-only Console behavior without operational support overclaim;
- `ai/core/04_ARCHITECTURE.md`;
- `ai/core/05_SYSTEM_MAP.md`;
- `ai/core/07_ENGINEERING_HISTORY.md`;
- `ai/core/13_ROADMAP.md`;
- `README.md`;
- directly affected English and Hungarian CLI/Console user guidance;
- `ai/prompts/033_CURRENT_TASK.md` during active implementation;
- `ai/history/033_2026-08-07_interactive-operator-console.md`;
- `ai/archive_prompts/033_2026-08-07_interactive-operator-console.md` at successful closure.

Runtime, Runtime Service, Alert, Notification, Scheduler, Configuration, installation, monitoring, provider, REST/API, Dashboard, and security documents shall change only if implementation proves a direct boundary clarification is required. Every actual update and justified omission shall be recorded in Task 033 history.

Documentation shall state:

`Canonical Engineering and Operational Data -> Canonical Operator Presentation Model -> Interactive Operator Console`.

It shall also state that Console refresh is explicit and one-shot; the Console never calls engines; non-interactive bare invocation never blocks; source IDs appear only in explicit details; advanced CLI/JSON composition remains available; and persistence/recovery, monitoring, providers, system installation/supervision, REST API, Dashboard, remediation, packaging, deployment, and release remain separately governed.


## Completion Criteria

Task 033 is complete only when:

- one Interactive Operator Console exists and consumes only validated Canonical Operator Presentation Model Overview values;
- interactive bare `qwsg` provides an immediately understandable keyboard-driven local interface and non-interactive bare `qwsg` provides a deterministic nonblocking concise fallback;
- Home clearly exposes condition, attention, changes, Alert summary, Guardian, freshness/completeness, and recommended next step without internal engine vocabulary or canonical IDs;
- Attention, Changes, Guardian, Evidence/Details, and Help views are deterministic, bounded, localized, terminal-safe, and reachable through documented keyboard actions;
- explicit refresh calls only an injected OverviewProvider; the production application provider uses one existing bounded Command/Pipeline execution and presentationmodel projection without defining or duplicating engineering semantics;
- no collection or provider call occurs before explicit refresh, no polling/retry/background loop exists, and provider failure retains the last valid Overview with a bounded diagnostic;
- English and Hungarian catalogs are complete, machine tokens remain stable, plain-text/monochrome/narrow-terminal behavior is accessible, and no color-only meaning exists;
- current explicit CLI commands, help, JSON output, profiles, advanced composition, Command/Pipeline behavior, Operator Model, Runtime, Runtime Service, and manual workflows remain compatible;
- no Runtime, Scheduler, Alert, Notification, Health, Rule, Policy, Configuration, monitoring, persistence, restart recovery, provider transport, REST API, Dashboard, process control, remediation, remote execution, AI, package, deployment, or release behavior exists in the Console;
- Release Minimalism records that the Console is a Version 1.0 requirement because the engineering core and shared operator semantics otherwise remain inaccessible to ordinary users, while the task does not absorb independent operational gates;
- all focused, build, full test, race, vet, format, Framework, Builder, lifecycle, diversion, screen, transition, provider, CLI, terminal, localization, accessibility, privacy, resource, determinism, boundary/import, documentation, permission/ACL, Git-diff, snapshot, and rollback validations pass;
- rollback remains proportional, collision-aware, verified, and preserves unrelated Owner-owned work;
- no dependency installation, automatic live collection, service installation, staging, commit, push, branch, tag, package, deployment, release, or infrastructure mutation occurred;
- the completed Task 033 prompt/history are archived without a successor and `bin/job --check` confirms canonical idle with Task 033 as the unique latest completed baseline.

A valid result is `complete`, `complete with disclosed limitations`, or `blocked`. Completion may not be claimed while any mandatory Console contract, Overview-only boundary, refresh/provider separation, bare-command safety, keyboard/navigation behavior, screen, fallback, localization/accessibility/privacy/resource rule, regression, test, documentation, rollback, or lifecycle gate remains unresolved.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-07 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
