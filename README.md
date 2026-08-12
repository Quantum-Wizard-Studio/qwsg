# Quantum Wizard Server Guardian

QWSG 1.0 is a local Linux Server Guardian. The supported platform is
Ubuntu 24.04 LTS, systemd 255+, linux-amd64, running the Guardian as an ordinary
non-root user. Start with the [release Quick Start](docs/release/QUICK_START.md);
support, operations, troubleshooting, security, upgrade/rollback/uninstall and
limitations are documented under `docs/release/`.

`qwsg observe` is the simple full operator evaluation. Its first run creates a private baseline and truthfully remains unknown; a later run composes the existing Inventory→Snapshot→Compare→Drift→Health→Rule→Policy→Report pipeline and publishes one private Current Operator State record. A separately started bare `qwsg` displays that qualified evidence. `qwsg check` remains the compatible Inventory→Snapshot profile.

Large evaluations remain bounded: QWSG ranks attention globally, correlates
validated Rule/Policy views, preserves the most severe evidence, and explicitly
reports how many additional concerns were correlated or omitted. Observation
failures distinguish Pipeline evaluation, operator projection, and state
publication without exposing host-sensitive error details.

Bare `qwsg` provides the local read-only Operator Console: interactive terminals receive keyboard navigation and explicit refresh, while non-interactive use receives a concise nonblocking view. See the [English](docs/user/INTERACTIVE_OPERATOR_CONSOLE.en.md) and [Hungarian](docs/user/INTERACTIVE_OPERATOR_CONSOLE.hu.md) guides.

`qwsg guardian run` now composes the existing Runtime Service as one
foreground, continuously operating local Guardian. The supported systemd user
unit supervises that unprivileged process, while the Console reports only
validated current lifecycle evidence. See the [operational architecture](docs/architecture/OPERATIONAL_GUARDIAN_SERVICE.md)
and [installation guide](docs/installation/INSTALL.md).

## Purpose

Quantum Wizard Server Guardian (QWSG) is a Professional Linux Server
Engineering Toolkit: an independent, modular system for trustworthy Linux
evidence, change understanding, and controlled protection.

## Status

Version `1.0.0` provides a user-installable local Linux
Inventory CLI, Snapshot Explorer, canonical Snapshot Comparison Engine,
Canonical Drift Engine, deterministic Canonical Health Engine, Canonical Rule
Engine, Canonical Policy Engine, Canonical Report Engine, and Canonical
Professional Alert Engine and provider-neutral Notification Delivery contracts.
Task 024 adds the presentation-independent Canonical Command Architecture,
deterministic analysis-pipeline orchestration, simple command profiles,
advanced composition, Command Definition 1.0, Command Execution 1.0, and a CLI
reference adapter. Future Interactive Terminal, Dashboard, and REST API
interfaces must consume the same contracts and cannot implement orchestration
or engineering logic.
Task 025 inserts deterministic Policy Profile evaluation between Rule and
Policy-backed Report generation without adding operational actions.
Task 026 establishes the versioned Canonical Configuration Contract, exact
source precedence, field provenance, immutable Effective Configuration, and a
scheduler-ready Schedule Definition without implementing scheduling or file
activation.
Task 027 establishes the deterministic Canonical Professional Scheduler, its
versioned state and execution records, and an explicitly invoked one-cycle
local adapter over the existing Command and Pipeline contracts. It adds no
daemon, service, monitoring, alert, or notification behavior.
Task 028 establishes the pure deterministic Canonical Professional Alert Engine,
versioned Alert Records, lifecycle state, acknowledgement, bounded suppression,
maintenance, deduplication, expiration, correlation, and recovery decisions. It
decides when an alert exists but provides no persistence or notification
delivery.
Task 029 establishes deterministic provider-neutral delivery planning, bounded
retry and queue-state proposals, canonical attempt/status/acknowledgement/evidence
records, and an explicitly invoked one-cycle adapter over injected providers.
It includes no production transport, persistence, daemon, monitoring, or Alert
decision behavior.
The local Terminal Console and continuously supervised Guardian now ship.
Concrete notification transports, an API, and a Web UI are explicit post-1.0
families and are not claimed by this release.

## Scope

Architecture, implementation, tests, tools, and operating guidance belong here as they are approved. Secrets, credentials, host-specific data, generated artifacts, and undocumented infrastructure changes do not belong in the repository.

Start with [`ai/core/00_PROJECT_PHILOSOPHY.md`](ai/core/00_PROJECT_PHILOSOPHY.md) and [`ai/core/01_CONSTITUTION.md`](ai/core/01_CONSTITUTION.md). This document will evolve throughout development.

The owner-review Product Definition is maintained in [`docs/PRODUCT_DEFINITION.md`](docs/PRODUCT_DEFINITION.md). Established constraints in that document apply now; strategic proposals remain subject to explicit owner approval.

The authoritative product-level relationship between QWSG's identity, boundaries, Agent, Installer, Console, MVP, future direction, and deferred architecture decisions is maintained in [`docs/PRODUCT_SYSTEM_BLUEPRINT.md`](docs/PRODUCT_SYSTEM_BLUEPRINT.md).

The canonical long-term Product Architecture is maintained in
[`docs/PRODUCT_ARCHITECTURE.md`](docs/PRODUCT_ARCHITECTURE.md). It defines one
deterministic engineering core shared by the complete local Community Guardian,
future central/remote Professional services, and operations-focused Provider
services, together with workspace, terminal, Web Dashboard, licensing, privacy,
deployment, automation, AI, ecosystem, and roadmap principles. Described future
capabilities are architecture, not claims of current implementation. Community
use and redistribution are governed by the proprietary source-available QWS
Community / Free License Version 1.0, not an OSI open-source license.

The authoritative observable behavior and acceptance boundary for QWSG Core Alpha is maintained in [`docs/FUNCTIONAL_SPECIFICATION.md`](docs/FUNCTIONAL_SPECIFICATION.md).

The authoritative Core Alpha technical design and `Core Alpha Slice 1: Read-only Server Discovery and System Inventory` are maintained in [`docs/architecture/CORE_ALPHA_ARCHITECTURE.md`](docs/architecture/CORE_ALPHA_ARCHITECTURE.md) and [`docs/architecture/CORE_ALPHA_SLICE_1.md`](docs/architecture/CORE_ALPHA_SLICE_1.md).

Build and test with `make build` and `make test`, then run
`build/qwsg help`. Installation guidance is in
[`docs/installation/INSTALL.md`](docs/installation/INSTALL.md). The supported
CLI and Snapshot Explorer guides are available in
[English](docs/user/CLI_AND_SNAPSHOT_EXPLORER.en.md) and
[Hungarian](docs/user/CLI_AND_SNAPSHOT_EXPLORER.hu.md), with a complete
[demonstration walkthrough](docs/user/DEMONSTRATION.md).
[English snapshot comparison](docs/user/SNAPSHOT_COMPARISON.en.md) and
[Hungarian snapshot comparison](docs/user/SNAPSHOT_COMPARISON.hu.md) document
the factual Change Record workflow.

System installation deliberately separates privilege: run `make build` as the
normal user, then use `sudo make install` only to copy the verified artifact.
The privileged step does not require Go in root's `PATH`.

Canonical System Inventory v1 now provides the authoritative internal Linux host model through the Collector Registry while preserving the Inventory 1.0 compatibility envelope. Its explicitly invoked file-backed [Inventory Persistence and Digital Twin foundation](docs/architecture/INVENTORY_PERSISTENCE_AND_DIGITAL_TWIN.md) can save and reload validated snapshots without monitoring or background execution. See the [developer guide](docs/development/CANONICAL_SYSTEM_INVENTORY.md); user guidance is available in [English](docs/user/CANONICAL_SYSTEM_INVENTORY.en.md) and [Hungarian](docs/user/CANONICAL_SYSTEM_INVENTORY.hu.md).

The [Canonical Drift Engine](docs/architecture/CANONICAL_DRIFT_ENGINE.md) is
the deterministic, offline semantic layer above the Snapshot Comparison Engine.
It emits one versioned Drift Record per Change Record and makes no health, risk,
or policy judgement.

The [Canonical Health Engine](docs/architecture/CANONICAL_HEALTH_ENGINE.md)
evaluates validated Drift Records through a deterministic Health taxonomy and
emits one versioned Health Record per Drift Record. It performs no monitoring,
alerting, policy, compliance, reporting, remediation, networking, or AI work.

The [Canonical Rule Engine](docs/architecture/CANONICAL_RULE_ENGINE.md)
evaluates bounded Rule Definition 1.0 data against validated Health Records and
emits explainable, versioned Rule Evaluation Records. It performs no Policy,
alerting, reporting, remediation, command execution, networking, or AI work.

The [Canonical Policy Engine](docs/architecture/CANONICAL_POLICY_ENGINE.md)
interprets immutable Rule Evaluation Records through deterministic, versioned
Policy Profiles. It emits traceable Policy Evaluation Records with explicit
precedence, defaults, inheritance, and conflicts, but performs no Scheduler,
Alert, notification, automation, remediation, networking, host mutation, or AI
work.

The [Canonical Report Engine](docs/architecture/CANONICAL_REPORT_ENGINE.md)
retains its Rule-backed Report 1.0 API and adds a Policy-backed Report contract
for the canonical pipeline. Both are presentation-neutral and traceable. It
does not re-evaluate upstream evidence or implement Policy evaluation, monitoring, alerts,
Dashboard, export, delivery, remediation, networking, or AI.

The [Canonical Command Architecture](docs/architecture/CANONICAL_COMMAND_ARCHITECTURE.md)
defines one command and execution model above the complete canonical pipeline.
`internal/pipeline` is the only orchestration layer; CLI and all future
presentations are replaceable adapters over the same deterministic behavior.

The [Canonical Configuration Contract](docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md)
resolves versioned built-in, primary, activated-override, and temporary sources
into one immutable Effective Configuration with explicit field provenance.
The canonical pipeline and Scheduler consume that result, including the same
schedule, timeout, retry, concurrency, and applicability semantics.

The [Canonical Professional Scheduler](docs/architecture/CANONICAL_SCHEDULER.md)
evaluates Schedule Definition 1.0 using explicit clock observations and durable
state. Its pure engine emits deterministic decisions and execution requests;
the minimal adapter performs one explicitly requested local cycle through the
existing Command resolver and Pipeline Orchestrator.

The [Canonical Professional Alert Engine](docs/architecture/CANONICAL_ALERT_ENGINE.md)
consumes validated canonical outputs, prior Alert state, explicit time, and
bounded acknowledgement/suppression facts. It emits immutable Canonical Alert
Records and proposed state with deterministic lifecycle, source-deduplication,
expiration, correlation, and recovery semantics. It performs no persistence,
delivery, monitoring, daemon, interface, remediation, network, or AI work.

The [Canonical Professional Notification Delivery](docs/architecture/CANONICAL_NOTIFICATION_DELIVERY.md)
consumes only validated Canonical Alert Records. Its pure planner and explicit
one-cycle adapter provide provider-neutral routing, bounded retries, queue
proposals, attempts, statuses, acknowledgements, and evidence without concrete
transport, durable storage, daemon, upstream re-evaluation, or Alert creation.

Engineering tasks follow [`ai/core/11_ENGINEERING_LIFECYCLE.md`](ai/core/11_ENGINEERING_LIFECYCLE.md). The official `ai/scripts/task-builder.sh` workflow generates an approved prompt/history pair from structured owner input after a completed task; `ai/scripts/next-task.sh` remains available when a separate unapproved draft/review cycle is required. Explicitly owner-authorized incomplete-task diversion uses `ai/scripts/divert-task-to-test.sh` to preserve failed evidence under the independent `ai/test_tasks/` namespace without weakening production completion gates. See [`docs/architecture/ENGINEERING_TASK_BUILDER.md`](docs/architecture/ENGINEERING_TASK_BUILDER.md).

The versioned [Reusable Engineering Framework](docs/architecture/REUSABLE_ENGINEERING_FRAMEWORK.md)
validates project identity, canonical Git state, lifecycle paths, required
reading, and project-specific validation commands through
`ai/scripts/framework-check.sh`. QWSG is its reference implementation.

The [Canonical Runtime Engine](docs/architecture/CANONICAL_RUNTIME_ENGINE.md)
coordinates one explicit bounded Scheduler → Alert → Notification cycle. It is
not a daemon or monitoring service.

The [Canonical Runtime Service](docs/architecture/CANONICAL_RUNTIME_SERVICE.md)
adds explicitly invoked fixed-rate recurrence, graceful cancellation, and
process-local state handoff above that one-cycle boundary. The narrow
[Operational Guardian Service](docs/architecture/OPERATIONAL_GUARDIAN_SERVICE.md)
now supplies installation, systemd user supervision, private restart handoff,
and lifecycle publication without moving those responsibilities into Runtime.

The [Canonical Operator Presentation Model](docs/architecture/CANONICAL_OPERATOR_PRESENTATION_MODEL.md)
projects validated canonical engineering and operational records into one
bounded, localization-ready operator overview. It adds no interface,
monitoring, persistence, process probe, or remediation behavior.
