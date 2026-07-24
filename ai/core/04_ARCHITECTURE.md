# Architecture Governance

## Purpose

This document will record approved system boundaries, components, interfaces, data flows, and architectural decisions for Quantum Wizard Server Guardian.

## Status

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

Task 018 establishes `internal/comparison` as the exclusive system-evolution
boundary. Future drift, health, alert, reporting, and interface modules consume
canonical Change Records and must not compare Inventory snapshots directly.
The contract is defined in
`docs/architecture/SNAPSHOT_COMPARISON_ENGINE.md` and
`docs/architecture/CHANGE_RECORD_SCHEMA.md`.
