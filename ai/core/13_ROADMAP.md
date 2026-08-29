# Roadmap

- Task 035: Canonical Operator Evaluation — additive `observe` profile composes the complete existing live pipeline, truthfully bootstraps missing baseline evidence, and publishes a qualified Overview for a later Console process.
- Task 036: Operator Projection Hardening — bounded severity-first attention, validated Rule/Policy correlation, explicit overflow disclosure, and differentiated observation diagnostics make large valid evaluations publishable.

- Task 034: Canonical Current Operator State — single-record process continuity implemented; general persistence, monitoring, and operational history remain separate Version 1.0 gates.

- Task 033: Interactive Operator Console — local read-only interface implemented; persistence/recovery, monitoring, providers, installation/supervision, REST API, and Dashboard remain independent Version 1.0 gates.

## Purpose and authority

This document structures future QWSG outcomes and dependencies. The canonical
product direction, edition boundaries, and architecture gates are defined in
`docs/PRODUCT_ARCHITECTURE.md`. This roadmap orders capability streams; it does
not authorize a task, promise a date, select commercial packaging, or claim
support.

## Current state

The Foundation Phase is complete:

- Tasks 003–008 established the Product Definition, Product & System Blueprint,
  Functional Specification, repository audit, and Core Alpha architecture.
- Tasks 009, 011, 012, and 014 established the one-shot Inventory slice,
  platform-wide Inventory Architecture, Collector Framework, and Canonical
  System Inventory v1.
- Tasks 013 and 015 established the governed Engineering Task Builder and
  versioned Reusable Engineering Framework.
- Tasks 016–018 established validated Digital Twin snapshot persistence, the
  supported user CLI and Snapshot Explorer, and the deterministic Snapshot
  Comparison Engine with canonical Change Records.
- Task 019 establishes the complete QWSG Product Architecture and edition
  strategy. It introduces no runtime or behavioral change.
- Task 020 establishes the deterministic Canonical Drift Engine and versioned
  Drift Records as the semantic boundary above Change Records.
- Task 021 establishes the deterministic Canonical Health Engine, Health Record
  1.0, explicit evidence semantics, and the evaluation boundary above Drift.
- Task 022 establishes the deterministic Canonical Rule Engine, Rule Definition
  1.0, Rule Evaluation Record 1.0, and matching boundary above Health.
- Task 023 establishes the deterministic Canonical Report Engine, Canonical
  Report 1.0, source traceability, and presentation boundary above Rule.
- Task 024 establishes the presentation-independent Canonical Command
  Architecture, deterministic pipeline orchestration, simple profiles, advanced
  composition, Command Definition 1.0, Command Execution 1.0, and the CLI
  reference adapter shared by all future interfaces.
- Task 025 establishes the deterministic Canonical Policy Engine, Policy
  Profile 1.0, Policy Evaluation Record 1.0, precedence, inheritance,
  conflict behavior, Policy-backed Report integration, and the permanent
  governance boundary above Rule.
- Task 026 establishes the Canonical Configuration Contract, deterministic
  source precedence, field provenance, Effective Configuration 1.0, and
  scheduler-ready Schedule Definition 1.0 without operational scheduling.
- Task 027 establishes the Canonical Professional Scheduler, deterministic
  evaluation and state contracts, bounded retry/overlap/concurrency behavior,
  restart-safe persistence, and an explicitly invoked one-cycle local adapter
  over the existing Command and Pipeline contracts.
- Task 028 establishes the pure Canonical Professional Alert Engine, immutable
  Alert Records, explicit lifecycle state, acknowledgement, bounded suppression
  and maintenance, deduplication, expiration, correlation, recovery, and
  deterministic source precedence without persistence or delivery.
- Task 029 establishes provider-neutral Canonical Professional Notification
  Delivery over immutable Alert Records: deterministic routing, queue and retry
  proposals, idempotency, delivery audit records, and an explicit one-cycle
  injected-provider adapter without production transports, durable persistence,
  daemon operation, monitoring, or upstream re-evaluation.

The prepared final repository identity is `1.0.0`. The implemented product is a user-installable
local Inventory, Snapshot Explorer, Comparison, full operator evaluation,
Current Operator State, Terminal Console, and continuously supervised Guardian,
composing the canonical Drift, Health, Rule, Policy, Report, Command,
Configuration, Scheduler, Alert, Notification, Runtime and Runtime Service
boundaries. Production notification transports, Dashboard/API, license enforcement,
remote agents, fleet/provider operations, remediation and AI are not implemented
and are post-1.0 rather than hidden release gates.

Version 1.0 release-gate classification at Task 038:

- **MUST:** satisfied by Tasks 038–043, including the Owner-run clean-host physical reboot and uninstall evidence. Task 044 owns final source, license, identity and reproducibility reconciliation.
- **SHOULD:** concrete notification transport and durable delivery continuity, unless the Project Owner makes off-console notification a Version 1.0 release requirement.
- **LATER:** REST API, Web Dashboard, fleet/provider operations, AI, remediation, licensing, and broader integrations.

## Roadmap rules

1. The shared deterministic core is implemented once and used by every edition.
2. Community remains a complete professional local engineering toolkit.
3. Professional work extends automation and management, never correctness.
4. Provider work extends tenant-aware operations, never core truth.
5. Core dependencies precede interface, automation, and commercial layers.
6. Every capability progresses through product obligation, threat/privacy
   analysis, canonical contract, bounded implementation, verification and
   migration, and an explicit support decision.
7. Unknown dates and unresolved business decisions remain unknown rather than
   becoming implied commitments.
8. A roadmap item becomes executable only through the canonical task lifecycle
   and explicit Project Owner authority.

## Stream 1 — Engineering Core

### Established foundation

- Collector Framework and Canonical System Inventory;
- validated Snapshot Store and Digital Twin envelopes;
- exclusive Snapshot Comparison Engine and canonical Change Records;
- canonical Drift taxonomy, engine, and Drift Records;
- canonical Health taxonomy, engine, and Health Records;
- canonical Rule definitions, deterministic operator model, engine, and Rule
  Evaluation Records;
- canonical Policy Profiles, deterministic governance outcomes, precedence,
  conflict semantics, engine, and Policy Evaluation Records;
- canonical Report taxonomy, engine, Canonical Report 1.0, rendering model, and
  source traceability;
- canonical command definitions, profiles, deterministic pipeline orchestration,
  execution results, parameter projection, and replaceable presentation;
- canonical configuration sources, deterministic precedence and conflicts,
  effective values with provenance, typed secret references, and scheduler-ready
  schedule definitions;
- canonical deterministic schedule evaluation, Scheduler state, events,
  execution requests/results, bounded local persistence and locking, and an
  explicit one-cycle Command/Pipeline adapter;
- canonical deterministic Alert decisions, Alert Records and state,
  acknowledgement, bounded suppression/maintenance, deduplication, expiration,
  correlation, recovery, and exact upstream source precedence;
- versioned schemas, deterministic output, privacy-safe identity, resource
  bounds, and failure isolation.

### Next architectural outcomes

1. **Core hardening and release gates** — platform matrix, performance and
   resource envelopes, migrations, packaging, security review, and support
   evidence.

No scheduler, dashboard, or Provider layer should invent these semantics.

## Stream 2 — Community Experience

Community work turns the core into a complete daily Linux engineering toolkit:

1. canonical private workspace adoption and migration;
2. coherent CLI vocabulary and configuration workflows;
3. full local Drift, Health, policy, report, export, and diagnostic workflows
   as their core contracts become available;
4. future keyboard-first Terminal UI with complete non-interactive equivalents;
5. localization, accessibility, terminal compatibility, and low-resource
   behavior;
6. native packaging, installation, update, removal, backup, restore, and
   supported-platform guidance;
7. documentation and demonstrations suitable for a supported Community
   release.

Community correctness, security, privacy, schema access, and manual operation
must not depend on a license, account, cloud service, or AI provider.

## Stream 3 — Automation and Professional

Professional outcomes start only after the corresponding Community workflow and
core contract are correct:

The one-cycle Canonical Runtime Engine is complete. It is the bounded
coordination foundation for later operational hosting and deliberately adds no
resident process. The next automation step must separately authorize daemon/
service lifecycle, durable cross-component state, and operational recovery.

The Operational Guardian Service now composes Runtime Service as one
unprivileged foreground process under systemd user supervision, with canonical
configuration, bounded exact-state restart handoff, single-writer protection,
truthful lifecycle freshness, and cross-process Console evidence. Release
hardening remains the sole Version 1.0 MUST gate; production transports and
broader automation remain optional or later.

1. scheduler and daemon architecture, including locking, recovery, missed runs,
   bounded retries, self-observability, and lifecycle;
2. durable Alert/Notification persistence, concrete email and other provider
   transports, channel health, configuration/secret integration, and daemon
   hosting over the established pure Alert and provider-neutral delivery
   contracts;
3. managed history, trends, graphs, and scheduled reports;
4. secured Professional Web Dashboard and application API;
5. multi-server enrollment, remote-agent identity, disconnected operation, and
   centrally managed profiles;
6. signed update orchestration, staged rollout, rollback, and compatibility;
7. deterministic offline-capable licensing, upgrade/downgrade, expiry, and data
   continuity;
8. premium support and privacy-reviewed diagnostic transfer.

Automation acceptance requires evidence that failure leaves local manual
engineering available and does not corrupt or reinterpret core records.

### QWSG 1.3 update-awareness sequence

Task 074 approves architecture only. Implementation remains divided into
separately authorized, reviewable stages: installed-package classification;
versioned public release-index and trust/source adapters; private awareness
state and network-free status; CLI integration over the existing updater;
isolated low-frequency Guardian checks; persistent transition notification and
localization; then outage, privacy, withdrawal, migration, update and rollback
acceptance. Community remains credential-free and operator-controlled.
Manifest signing/key lifecycle, SMTP diagnostic refinement, external-delivery
evidence, central subscriptions and every Pro/fleet capability remain separate
Owner decisions.

## Stream 4 — Provider Operations

Provider outcomes build on verified Professional control-plane contracts:

1. tenant, organization, project, role, and delegated-administration model;
2. independently tested tenant-aware storage, authorization, audit, export,
   retention, deletion, backup, and support access;
3. fleet segmentation, policy rollout, agent lifecycle, and staged updates;
4. provider dashboards and tenant-scoped customer portals;
5. white-label presentation with preserved evidence provenance;
6. integration APIs, webhooks, service-management, and entitlement seams;
7. high availability, disaster recovery, regionality, capacity, and
   control-plane observability;
8. self-hosted and future managed-service deployment and exit contracts.

No Provider implementation begins before a dedicated tenancy threat model and
data architecture are approved.

## Stream 5 — Platform and ecosystem

Cross-cutting platform work may proceed when demanded by an approved consumer:

- stable public CLI, schema, API, and SDK contracts;
- signed collector, rule, report, notification, and integration extensions;
- privacy-reviewed catalog and distribution mechanisms;
- policy and report libraries, training, certification, and support;
- optional managed update, licensing, notification, or control-plane services;
- optional local or remote AI adapters consuming minimized public records.

Every ecosystem service is optional around the local core. AI remains advisory
and cannot become a canonical decision or required dependency.

## Dependency structure

```text
Engineering Core
├── Community Experience
│   └── supported local product
├── Professional Automation
│   └── Provider Operations
└── Platform and Ecosystem adapters
```

Streams may overlap only where their contracts are stable and their authority
boundaries are explicit. Edition priority does not permit an outer stream to
fork or bypass the core.

## Release and support gates

A capability may be described as supported only when:

- its product obligation and edition boundary are explicit;
- security, privacy, privilege, data, and failure models are approved;
- canonical interfaces and compatibility rules exist;
- implementation and migrations are bounded and rollback-capable;
- deterministic, integration, resource, failure, and platform tests pass;
- installation, operation, diagnosis, export, recovery, and removal are
  documented;
- unsupported and degraded contexts remain explicit;
- the Project Owner approves the support decision.

Completed task numbers prove delivery of their bounded scope, not completion of
an entire stream or product edition.

Task 032 establishes the shared Canonical Operator Presentation Model before
interface implementation. It prevents the CLI, future Terminal Console, REST
API, and Web Dashboard from independently interpreting Health, changes,
Alerts, Runtime/Service state, evidence freshness, or recommended next steps.
It does not satisfy persistence, restart recovery, monitoring, provider,
installation, packaging, or release gates.

## Roadmap governance

Approved goals, dependencies, risks, and gates belong here. Detailed task plans
belong in active prompts; evidence belongs in task histories; product philosophy
belongs in `docs/PRODUCT_ARCHITECTURE.md`. Fixed promises, speculative dates,
pricing, credentials, and implementation without an approved task do not belong
in the roadmap.
