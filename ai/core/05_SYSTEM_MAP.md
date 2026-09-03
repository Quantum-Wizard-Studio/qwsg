# System Map

- `internal/operatorstate`: private, atomic, integrity-validated single Current Operator State record; no presentation semantics or history.

- `internal/operatorconsole`: read-only Console Model 1.0, localized views including bounded-attention reduction disclosure, navigation, and injected session loop over `presentationmodel.Overview`.
- `cmd/qwsg`: terminal selection, `observe` baseline/full-evaluation composition, typed Pipeline/projection/publication diagnostics, typed Current State publication, and explicit one-shot refresh adapter.

```text
qwsg observe -> existing live canonical Pipeline -> typed Operator Overview
             -> Current Operator State -> new process -> qwsg Console
```

## Purpose

This document will map QWSG components, external boundaries, trust relationships, and operational dependencies.

## Status

A supported Version 1.0 RC local CLI, Terminal Console, Snapshot Explorer and
systemd-supervised Guardian front the implemented canonical boundaries. The product map remains
`docs/PRODUCT_SYSTEM_BLUEPRINT.md`, observable behavior remains
`docs/FUNCTIONAL_SPECIFICATION.md`, and technical allocation is documented
under `docs/architecture/`. The canonical long-term ecosystem, edition, user
experience, deployment, licensing, privacy, automation, and AI boundaries are
defined in `docs/PRODUCT_ARCHITECTURE.md`.

All editions share one deterministic engineering flow:

Every presentation enters that flow through the Canonical Command Architecture:

```text
CLI / local Interactive Terminal / future Dashboard / future REST API
                              |
                    Command Definition 1.0
                              |
                    Canonical Command Plan
                              |
                    Pipeline Orchestration
                              |
 Inventory -> Snapshot -> Compare -> Drift -> Health -> Rule -> Policy -> Report
                              |
                    Command Execution 1.0
                              |
                 Replaceable presentation
```

Only `internal/pipeline` orchestrates engines. `internal/command` owns the
presentation-neutral public request, plan, profile, parameter, projection, and
execution contracts. `internal/presentation` renders completed executions and
cannot call the pipeline or engines. `cmd/qwsg` is the first adapter, not an
alternative source of command behavior.

Canonical configuration enters orchestration through one validated boundary:

```text
Configuration Sources -> Configuration Resolver -> Effective Configuration
                                                     |
                                                     +-> canonical pipeline
                                                     `-> Scheduler Engine
                                                              |
                                                   one-cycle local adapter
                                                              |
                                                   Command -> Pipeline
```

Configuration owns source precedence, defaults, validation, stable identity,
field provenance, and schedule-definition semantics. It performs no activation
or execution. Scheduler consumes Effective Configuration and cannot
invent an independent configuration path.

Scheduler owns due-time evaluation, bounded execution planning, durable local
state, and canonical Scheduler records. Its explicit local cycle resolves only
existing Command profiles and invokes only the Pipeline Orchestrator. It does
not own recurring process lifecycle, monitoring, Policy interpretation, alerts,
notifications, remediation, or presentation.

Alert consumes validated canonical outputs after their owning engines have
finished. Its pure decision engine combines those immutable records with
explicit prior Alert state, evaluation time, acknowledgement facts, and bounded
suppression/maintenance windows. It owns canonical Alert identity, severity,
category, lifecycle, deduplication, expiration, correlation, and recovery
decisions. It does not persist state, deliver notifications, run a daemon,
monitor hosts, remediate, or become a Command/Pipeline stage.

Notification consumes validated immutable Alert Records only. Its pure planner
combines them with explicit Notification-owned routes, endpoint/provider
references, previous Queue State, and explicit time. It emits deterministic
Delivery Plans and proposed Queue State. An explicitly invoked one-cycle
adapter calls replaceable injected providers and records outcomes. Notification
does not evaluate upstream engines, create Alerts, persist queues, run a daemon,
or implement a concrete transport.

```text
Health / Rule / Policy / Scheduler Event / Canonical Report
                              |
                     pure Alert Engine
                              |
              Alert Records + proposed Alert State
                              |
         pure Notification Delivery planner
                              |
       Delivery Plan + proposed Queue State
                              |
          explicit one-cycle Provider adapter
```

```text
Collectors -> Canonical Inventory -> Snapshot Store -> Comparison Engine -> Drift Engine -> Health Engine -> Rule Engine -> Policy Engine -> Report Engine
                                                                                                                       |
                                                                                                                       v
                                                                                                             future operations
                                              |
                 +----------------------------+---------------------------+
                 |                            |                           |
        Community local tools       Professional automation      Provider operations
```

Comparison owns facts about change, Drift owns semantic change classification,
and Health owns deterministic engineering-condition evaluation. Each Health
Record preserves its Drift and Change references. Rule owns deterministic
matching of predefined conditions and preserves Health evidence references.
Policy owns deterministic governance interpretation without changing Rule
evidence. Report owns deterministic presentation contracts and preserves Policy,
Rule Evaluation, and Health evidence references. Health, Rule, Policy, and
Report are pure offline libraries and introduce no scheduler, monitoring,
alert, remediation, process, delivery, network, or AI boundary.

Community, Professional, and Provider are orchestration and operating-context
layers above the same versioned core. No dashboard, control plane, notification
service, license service, ecosystem service, or AI adapter may become an
alternative source of engineering truth.

For Slice 1, a local operator invokes the non-root Agent boundary. The discovery coordinator runs bounded read-only collectors, followed by normalization, redaction, validation, inventory assembly, optional latest-envelope persistence, and CLI/JSON presentation. A future Console consumes an Agent-owned redacted contract and cannot access collectors, shell execution, or privilege directly. Installer, remediation, network Console, e-mail, and update boundaries remain outside Slice 1.

The Canonical System Inventory is now the single internal host-information boundary. Host, OS, kernel, CPU, memory, storage, filesystem, network, and virtualization collectors register through the Collector Registry. The coordinator assembles both the authoritative canonical representation and its Inventory 1.0 compatibility projection from the same structured Results. The explicitly invoked Inventory Store can persist and revalidate that synchronized envelope after collection; it has no collector, scheduler, monitoring, or network responsibility. Future Health, Rule, Alert, Policy, Automation, Console, API, and reporting components consume validated canonical inventory and do not query Linux directly.

System evolution follows the permanent boundary:
`Inventory -> Snapshot Store -> Comparison Engine -> Drift Engine -> Health Engine -> Rule Engine -> Policy Engine -> Report Engine`. Comparison
emits deterministic factual Change Records. Drift emits deterministic semantic
Drift Records without judging them. Health evaluates engineering condition,
Rule matches fixed conditions, Policy interprets their governance treatment,
and Report presents their canonical evidence.
No downstream module may independently diff Inventory snapshots, reclassify
Change Records, re-evaluate Health, Rules, or Policy, or rebuild engineering summaries.

The local administrator invokes `qwsg`. Command Execution 1.0 is the canonical
machine boundary; JSON serializes it and the separate terminal renderer exposes
safe summaries. Snapshot
Explorer list/info/load operations consume only validated Inventory Store data.
The Makefile installs one binary and creates no runtime service or state.

Verified components and connections belong here. Assumed services, secret endpoints, live credentials, and invented dependencies do not. The map will evolve during development as facts are verified.

`internal/runtime` coordinates one explicit bounded cycle over Scheduler,
validated Pipeline traces, Alert, and Notification. It provides no service
loop, persistence, or monitoring behavior.

`internal/runtimeservice` sits directly above Runtime and owns only fixed-rate
recurrence plus graceful local process lifecycle. Runtime calls remain
sequential. `internal/guardian` supplies its production foreground composition,
private exact-state checkpoint and publisher. The systemd user manager
supervises that process; it never owns QWSG engineering decisions.

`systemd user service -> qwsg guardian run -> Runtime Service -> Runtime ->
Scheduler -> Command/Pipeline -> Alert -> Notification`.

Validated Runtime/Service evidence flows through `internal/guardian` into
Current Operator State, then the read-only Console. Explicit `observe` shares
the Guardian operation lock and cannot race these writers.

Validated Command, change, Health, Report, Alert, Runtime, and Runtime Service
observations may flow into `internal/presentationmodel`. It emits one canonical
operator Overview for replaceable interfaces without calling engines or
probing operational state:

`Canonical Engineering and Operational Data -> Operator Presentation Model -> Replaceable Interface`.

Public release awareness is an independent outbound-only side flow:

```text
approved HTTPS source -> Release Source adapter -> strict release-index parse
                      -> Ed25519 authenticity -> installed-aware evaluation
                      -> private Update Awareness State
                                           +-> network-free update status
                                           `-> transition -> existing Notification transport

explicit operator action -> existing acquire/package/migration/update/rollback transaction
```

Update Discovery owns no Health, readiness, local observation, SMTP,
privileged installation or rollback behavior. Guardian may schedule it at low
frequency with bounded failure isolation; release-source failure cannot block
the primary local monitoring flow. The canonical contract is
`docs/architecture/UPDATE_DISCOVERY_AND_RELEASE_AWARENESS.md`.

Task 076 implements the flow through installed-aware evaluation. Task 077 adds
the independent strict atomic awareness state plus explicit check and
network-free status. Task 078 fixes the single production endpoint and bundled
public trust anchor, adds deterministic offline-signing/publication plumbing,
and stops before infrastructure mutation or publication. Task 079 connects a
separate Guardian-owned 24-hour due loop to the existing awareness manager
after local startup publication. Persisted last-attempt time prevents duplicate
restart retrieval; one bounded check runs at a time and its failure cannot
enter Runtime, Health, Alert, Notification, acquisition, or installation
flows. Task 080 adds a post-authentication Guardian-only branch through the
configured Community SMTP transport, with successful-delivery identity stored
atomically in awareness state and deduplicated across restarts. Failed or
disabled delivery never enters Guardian health and cannot retry before the next
release-check interval. The implemented contracts are
`docs/architecture/RELEASE_INDEX_AND_SOURCE_CONTRACT.md`,
`docs/architecture/UPDATE_AWARENESS_STATE.md`, and
`docs/release/RELEASE_INDEX_PUBLICATION.md`.

Local installed identity is a separate read-only input:

```text
canonical package layout + installed RELEASE.json + binary identity
    -> Installed Package Classifier
       +-> guided-install package decision
       `-> existing updater -> declared migration -> transaction/rollback
```

The classifier owns no installation mutation, remote release trust,
configuration/readiness interpretation, Guardian state or rollback state. Its
contract is `docs/architecture/INSTALLED_PACKAGE_CLASSIFICATION.md`.
