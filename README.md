# Quantum Wizard Server Guardian

## Purpose

Quantum Wizard Server Guardian (QWSG) is a Professional Linux Server
Engineering Toolkit: an independent, modular system for trustworthy Linux
evidence, change understanding, and controlled protection.

## Status

Version `0.0.1-prealpha` now provides a user-installable one-shot Linux
Inventory CLI, Snapshot Explorer, canonical Snapshot Comparison Engine,
Canonical Drift Engine, deterministic Canonical Health Engine, Canonical Rule
Engine, and Canonical Report Engine contracts.
Task 024 adds the presentation-independent Canonical Command Architecture,
deterministic analysis-pipeline orchestration, simple command profiles,
advanced composition, Command Definition 1.0, Command Execution 1.0, and a CLI
reference adapter. Future Interactive Terminal, Dashboard, and REST API
interfaces must consume the same contracts and cannot implement orchestration
or engineering logic.
It remains a pre-alpha release and does not provide monitoring, daemon mode, services,
alerts, an API, or a Web UI.

## Scope

Architecture, implementation, tests, tools, and operating guidance belong here as they are approved. Secrets, credentials, host-specific data, generated artifacts, and undocumented infrastructure changes do not belong in the repository.

Start with [`ai/core/00_PROJECT_PHILOSOPHY.md`](ai/core/00_PROJECT_PHILOSOPHY.md) and [`ai/core/01_CONSTITUTION.md`](ai/core/01_CONSTITUTION.md). This document will evolve throughout development.

The owner-review Product Definition is maintained in [`docs/PRODUCT_DEFINITION.md`](docs/PRODUCT_DEFINITION.md). Established constraints in that document apply now; strategic proposals remain subject to explicit owner approval.

The authoritative product-level relationship between QWSG's identity, boundaries, Agent, Installer, Console, MVP, future direction, and deferred architecture decisions is maintained in [`docs/PRODUCT_SYSTEM_BLUEPRINT.md`](docs/PRODUCT_SYSTEM_BLUEPRINT.md).

The canonical long-term Product Architecture is maintained in
[`docs/PRODUCT_ARCHITECTURE.md`](docs/PRODUCT_ARCHITECTURE.md). It defines one
deterministic engineering core shared by the complete Community toolkit, the
automation-focused Professional Edition, and the operations-focused Provider
Edition, together with workspace, terminal, Web Dashboard, licensing, privacy,
deployment, automation, AI, ecosystem, and roadmap principles. Described future
capabilities are architecture, not claims of current implementation.

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

The [Canonical Report Engine](docs/architecture/CANONICAL_REPORT_ENGINE.md)
transforms validated Rule Evaluation Records into presentation-neutral,
traceable Canonical Reports 1.0 with deterministic text rendering. It does not
re-evaluate upstream evidence or implement Policy, monitoring, alerts,
Dashboard, export, delivery, remediation, networking, or AI.

The [Canonical Command Architecture](docs/architecture/CANONICAL_COMMAND_ARCHITECTURE.md)
defines one command and execution model above the complete canonical pipeline.
`internal/pipeline` is the only orchestration layer; CLI and all future
presentations are replaceable adapters over the same deterministic behavior.

Engineering tasks follow [`ai/core/11_ENGINEERING_LIFECYCLE.md`](ai/core/11_ENGINEERING_LIFECYCLE.md). The official `ai/scripts/task-builder.sh` workflow generates an approved prompt/history pair from structured owner input after a completed task; `ai/scripts/next-task.sh` remains available when a separate unapproved draft/review cycle is required. Explicitly owner-authorized incomplete-task diversion uses `ai/scripts/divert-task-to-test.sh` to preserve failed evidence under the independent `ai/test_tasks/` namespace without weakening production completion gates. See [`docs/architecture/ENGINEERING_TASK_BUILDER.md`](docs/architecture/ENGINEERING_TASK_BUILDER.md).

The versioned [Reusable Engineering Framework](docs/architecture/REUSABLE_ENGINEERING_FRAMEWORK.md)
validates project identity, canonical Git state, lifecycle paths, required
reading, and project-specific validation commands through
`ai/scripts/framework-check.sh`. QWSG is its reference implementation.
