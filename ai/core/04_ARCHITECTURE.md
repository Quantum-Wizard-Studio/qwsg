# Architecture Governance

## Version 1.0 local release boundary

QWSG 1.0 ships one local deterministic core, CLI/Terminal Console, private
bounded file state, and an ordinary-user systemd-supervised Guardian. Console
remains presentation-only and Guardian composes Runtime Service. Network
interfaces, providers, fleet, remediation and AI are post-1.0.

Current local presentation flows through typed projection, one Canonical Current Operator State envelope, OverviewProvider, Presentation Model freshness requalification, and the replaceable interface.

Canonical Operator Evaluation is an application composition, not a new engine: `observe` selects the existing live eight-stage Command/Pipeline, bootstraps an absent baseline through compatible `check`, projects exact typed outputs, and publishes Current Operator State. Existing engine ownership and `check` semantics remain unchanged.

The local presentation chain is `Canonical Engineering and Operational Data -> Canonical Operator Presentation Model -> Interactive Operator Console`; interfaces do not reinterpret engineering evidence.

Operator Presentation Model 1.1 globally ranks and bounds valid attention,
correlates exact Rule-to-Policy provenance, and publishes explicit reduction
counts. Application diagnostics distinguish Pipeline, projection, and Current
State publication boundaries. These harden the product composition without
changing canonical engine semantics.

## Purpose

This document will record approved system boundaries, components, interfaces, data flows, and architectural decisions for Quantum Wizard Server Guardian.

## Status

The canonical ecosystem and edition architecture is defined in
`docs/PRODUCT_ARCHITECTURE.md`. It governs the shared deterministic engineering
core; complete Community toolkit; automation-focused Professional Edition;
operations-focused Provider Edition; product identity; workspace and interface
philosophy; licensing, privacy, deployment, automation, AI, ecosystem, and
roadmap boundaries. Capability architecture documents refine that product
architecture and must not introduce an edition-specific correctness tier or a
mandatory cloud or AI dependency.

The platform-wide Inventory Architecture is defined by `12_INVENTORY_ARCHITECTURE.md` and is the canonical system-description contract for future collectors and consumers. The Core Alpha architecture and `Core Alpha Slice 1: Read-only Server Discovery and System Inventory` architecture were approved under Task 008; their narrower `1.x` inventory envelope is a bounded implementation profile subject to the platform-wide compatibility and migration rules. The authoritative Core Alpha package begins at `docs/architecture/CORE_ALPHA_ARCHITECTURE.md` and links its data model, security model, gate register, requirements mapping, implementation plan, and accepted Slice 1 ADRs.

The architecture establishes a digital-twin object model, extensible canonical layers and relationships, a common collector contract, canonical JSON, schema evolution rules, resource limits, localization boundaries, and consumer separation. The internal implementation realizes that contract through `internal/collector`: validated descriptors, explicit capabilities and dependencies, duplicate-safe registration, availability checks, deterministic dependency-aware execution, bounded per-collector contexts, cancellation, panic isolation, and structured results. `internal/app` obtains collector contributions only through this Registry and preserves the legacy `1.0` Inventory projection until a separately authorized canonical-model migration.

Core Alpha additionally establishes a non-root read-only collector boundary, Agent-owned local truth, bounded command execution, privacy controls, and an implementation/test handoff. Runtime, packaging, supported platforms, full storage, Console security, e-mail, retention, update authenticity, and business policy remain explicit gates.

Verified design decisions and their rationale belong here. Speculation, credentials, host-specific configuration, and unapproved implementation commitments do not. The architecture will evolve during development through documented decisions.

The Engineering Task Builder architecture is defined in `docs/architecture/ENGINEERING_TASK_BUILDER.md`. It is an engineering-governance component, separate from the QWSG runtime, that converts owner-authored structured input into deterministic approved prompt/history pairs through a validated rollback-capable transaction.

The versioned Reusable Engineering Framework is defined in
`docs/architecture/REUSABLE_ENGINEERING_FRAMEWORK.md`. It owns reusable task,
approval, lifecycle, configuration-validation, Git-safety, snapshot, rollback,
history, and delivery boundaries. QWSG remains the reference implementation;
product architecture and runtime behavior remain separate.

The permanent future-state architecture for the Quantum Wizard Creator System
(QWCS), a reusable multi-project AI Engineering Operating System and intended
successor to Framework 1.x, is defined in
`docs/architecture/QWCS_ENGINEERING_OPERATING_SYSTEM.md`. It separates
Engineering Principles, Constitution, Rule Registry, Task Compiler, Governance,
Validation, Lifecycle, Override, Evidence, and Compatibility responsibilities.
It is architecture only and does not modify current Framework 1.x rules or
authorize implementation or migration.

Task 014 implements Canonical System Inventory v1 as the authoritative read-only host model. The existing Collector Registry is the sole acquisition boundary; `internal/app` produces the validated canonical layers/resources/facts representation and the legacy Inventory 1.0 projection from the same Results. Implementation details and compatibility rules are defined in `docs/architecture/CANONICAL_SYSTEM_INVENTORY_V1.md`.

Task 016 adds the first explicitly invoked file-backed Inventory Store after
canonical validation. It retains the complete compatibility/canonical envelope
as a versioned Digital Twin observation with deterministic naming, integrity
verification, restrictive permissions, atomic installation, validated loading,
and bounded retention. It remains outside collectors and introduces no daemon,
scheduler, comparison, health, alert, API, database, or network boundary. The
profile is defined in
`docs/architecture/INVENTORY_PERSISTENCE_AND_DIGITAL_TWIN.md`.

Task 017 establishes `cmd/qwsg` as the first supported user application
boundary. It provides contextual help, controlled build/version data,
terminal-safe human summaries, explicit JSON compatibility, and Snapshot
Explorer list/info/load commands. Every explorer result passes through the
existing Inventory Store validation boundary. Build and installation remain
one binary with no service, daemon, scheduler, network listener, database, or
external dependency.

Task 018 establishes `internal/comparison` as the exclusive factual
system-evolution boundary. Task 020 establishes `internal/drift` as the first
semantic layer above it: one versioned Drift Record per canonical Change
Record, without health, risk, policy, or remediation judgement. Task 021
establishes `internal/health` as the deterministic engineering-condition layer:
one versioned Health Record per validated Drift Record, with explicit
unknown, unsupported, and insufficient-evidence semantics. Task 022 establishes
`internal/rule` as the pure deterministic matching layer over versioned Health
Records, producing canonical Rule Evaluation Records with explicit match,
non-match, insufficient, unsupported, invalid, disabled, and error outcomes.
Task 023 establishes `internal/report` as the presentation-contract layer over
validated Rule Evaluation Records, producing traceable Canonical Reports 1.0
without re-evaluating upstream evidence. Task 025 establishes `internal/policy`
as the sole deterministic governance interpreter over immutable Rule Evaluation
Records and adds a Policy-backed Report contract while retaining the original
Rule-backed Report API. Future Automation, alerts, dashboards, exports,
notifications, and interfaces consume canonical Drift, Health, Rule, Policy,
and Report contracts and must not compare Inventory snapshots,
reclassify Change Records, create competing Health semantics, or reimplement
Rule matching or engineering summaries. The contracts are defined in
`docs/architecture/SNAPSHOT_COMPARISON_ENGINE.md` and
`docs/architecture/CHANGE_RECORD_SCHEMA.md`; the permanent semantic boundary is
defined in `docs/architecture/CANONICAL_DRIFT_ENGINE.md`, and the Health
evaluation boundary is defined in
`docs/architecture/CANONICAL_HEALTH_ENGINE.md`, and the permanent Rule boundary
is defined in `docs/architecture/CANONICAL_RULE_ENGINE.md`. The permanent
presentation boundary is defined in
`docs/architecture/CANONICAL_REPORT_ENGINE.md`; the governance boundary is
defined in `docs/architecture/CANONICAL_POLICY_ENGINE.md`.

Task 024 establishes `internal/command` as the permanent
presentation-independent Command Definition 1.0 and Command Execution boundary,
`internal/pipeline` as the only canonical engine-orchestration layer, and
`internal/presentation` as a replaceable consumer of completed executions.
Simple profiles and advanced composition resolve to the same plans and the same
Inventory → Snapshot → Compare → Drift → Health → Rule → Policy → Report engines. The
CLI is the first reference adapter only. Future Interactive Terminal, Dashboard,
and REST API adapters must consume the same contracts and may not implement
orchestration or engineering logic. The permanent boundary is defined in
`docs/architecture/CANONICAL_COMMAND_ARCHITECTURE.md`.

Task 026 establishes `internal/configuration` as the sole canonical
configuration-normalization and resolution boundary. Versioned Configuration
Source Records resolve by explicit precedence into immutable Effective
Configuration with field provenance, stable identity, strict validation, and
Schedule Definition 1.0. `internal/pipeline` consumes the effective result;
future Scheduler code may not recreate configuration semantics. No file
activation, secret backend, Scheduler, daemon, or operational action is added.
The permanent boundary is defined in
`docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md`.

Task 027 establishes `internal/scheduler` as the canonical deterministic
schedule-evaluation boundary. The pure engine consumes Effective Configuration,
Schedule Definition 1.0, Scheduler State, and explicit clock observations. Its
one-cycle local adapter alone owns bounded locking, state persistence, Command
profile resolution, and submission to the existing Pipeline Orchestrator. It
does not add a daemon, service lifecycle, monitoring, alerts, notifications, or
an alternative execution path. The permanent boundary is defined in
`docs/architecture/CANONICAL_SCHEDULER.md`.

Task 028 establishes `internal/alert` as the pure canonical decision boundary
that determines when an Alert lifecycle event exists. It consumes only
validated Health, Rule, Policy, Scheduler Event, Effective Configuration, and
Canonical Report contracts plus explicit prior Alert state, time,
acknowledgements, and bounded suppression windows. It emits immutable Alert
Records and proposed state without adding persistence, delivery, a Command or
Pipeline stage, daemon, monitoring, interfaces, remediation, network access, or
AI. The permanent boundary is defined in
`docs/architecture/CANONICAL_ALERT_ENGINE.md`.

Task 029 establishes `internal/notification` as the provider-neutral delivery
boundary downstream of immutable Canonical Alert Records. Its pure planner owns
deterministic routing, fan-out, queue proposals, retry eligibility, deadlines,
and idempotency. Its explicitly invoked one-cycle adapter calls only injected
providers and records attempts, statuses, provider acknowledgements, and
privacy-bounded evidence. It adds no concrete transport, persistence, daemon,
monitoring, upstream evaluation, Alert creation, interface, remediation, or AI.
The permanent boundary is defined in
`docs/architecture/CANONICAL_NOTIFICATION_DELIVERY.md`.

## Canonical Runtime Engine

`internal/runtime` is the one-cycle coordinator. It invokes Scheduler once,
validates Scheduler-owned execution traces, projects exact typed Pipeline
outputs to Alert, and passes only resulting Alert Records to Notification. It
owns no scheduling, business evaluation, alert decision, delivery policy,
persistence, service loop, interface, or mutation. The permanent contract is
`docs/architecture/CANONICAL_RUNTIME_ENGINE.md`.

## Canonical Runtime Service

`internal/runtimeservice` owns deterministic recurrence and local process
lifecycle above `internal/runtime`. It invokes one Runtime Runner sequentially,
forwards exact proposed states in memory, emits bounded synchronous evidence,
and maps SIGINT/SIGTERM to cancellation through a separate adapter. It owns no
Runtime or downstream engine logic, persistence, system installation,
supervision, monitoring, interface, or mutation. The canonical contract is
`docs/architecture/CANONICAL_RUNTIME_SERVICE.md`.

## Operational Guardian Service

`internal/guardian` is the narrow local operational adapter around Runtime
Service. It supplies canonical dependency wiring, a private bounded restart
checkpoint, single-writer locking, Current Operator State publication, and a
generation-correlated demotion-only exit boundary. The foreground
`qwsg guardian run` process is supervised by the supported systemd user unit;
systemd owns process management while Runtime Service retains recurrence. The
adapter owns no engineering, scheduling, Alert, Notification, or remediation
decision. Its contract is
`docs/architecture/OPERATIONAL_GUARDIAN_SERVICE.md`.

## Canonical Operator Presentation Model

`internal/presentationmodel` is the presentation-independent read model below
replaceable interfaces. It consumes validated canonical records, preserves
their ownership and uncertainty, and produces one bounded operator Overview.
It owns no engineering evaluation, interface, monitoring, persistence,
process observation, or remediation. The permanent contract is
`docs/architecture/CANONICAL_OPERATOR_PRESENTATION_MODEL.md`.

## Update Discovery and Release Awareness

Task 074 defines the QWSG 1.3 read-only public release-awareness boundary in
`docs/architecture/UPDATE_DISCOVERY_AND_RELEASE_AWARENESS.md`. It extends the
existing QWSG 1.2.0 Forgejo discovery, strict versioning, package verification,
migration and transactional rollback path through a source-neutral versioned
manifest, explicit authenticity model, private awareness state, isolated
Guardian schedule and persistent notification transitions. It does not create
a second updater, treat discovery as Health/readiness, or authorize automatic
installation, telemetry, registration, central access or Pro implementation.

Task 076 implements the first remote-metadata layer as
`internal/releasediscovery`, with the permanent contract in
`docs/architecture/RELEASE_INDEX_AND_SOURCE_CONTRACT.md`. It strictly parses
`qwsg.release-index/1`, verifies deterministic Ed25519 signatures against
explicit trust anchors, retrieves through a bounded source-neutral HTTPS
interface, and performs pure evaluation through Task 075 installed identity
and the existing local migration registry. It activates no production source
or key and owns no awareness state, Guardian schedule, notification, artifact
acquisition, installation, rollback, publication, telemetry or Pro behavior.

## Installed Package Classification

Task 075 implements the local installed-identity prerequisite in
`internal/installation`, documented by
`docs/architecture/INSTALLED_PACKAGE_CLASSIFICATION.md`. Complete safe package
layout, strict installed `qwsg.release/1` provenance and exact binary identity
agreement establish a verified supported installation. Binary presence,
version output, configuration and runtime state cannot do so. A supplied
candidate becomes a supported upgrade only through the existing declared
migration registry; legacy, unknown, incomplete and inconsistent evidence
fails closed before mutation. Guided installation and native update
orchestration share this classifier. Remote release authenticity remains an
independent Update Discovery concern.
