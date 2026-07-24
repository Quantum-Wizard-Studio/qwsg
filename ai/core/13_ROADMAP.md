# Roadmap

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

The repository remains `0.0.1-prealpha`. The implemented product is a
user-installable, one-shot local Inventory, Snapshot Explorer, and Comparison
workflow. Monitoring, Drift, Health, alerts, daemon mode, scheduler, Web
Dashboard, licensing, remote agents, fleet management, Provider operations,
and AI adapters are not implemented.

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
- versioned schemas, deterministic output, privacy-safe identity, resource
  bounds, and failure isolation.

### Next architectural outcomes

1. **Drift contract** — distinguish factual change from policy-relevant drift
   without bypassing Change Records.
2. **Health contract** — deterministic findings, severity, completeness,
   aggregation, evidence links, and unknown/unsupported behavior.
3. **Rule and policy contract** — versioned declarative evaluation with
   provenance, validation, profiles, and conflict behavior.
4. **Report contract** — canonical structured report data from which terminal,
   file, Web, and notification views derive.
5. **Configuration contract** — layered, validated, explainable configuration
   and secret references without hidden defaults.
6. **Core hardening and release gates** — platform matrix, performance and
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

1. scheduler and daemon architecture, including locking, recovery, missed runs,
   bounded retries, self-observability, and lifecycle;
2. alert lifecycle, maintenance, suppression, recovery, delivery audit, email,
   multiple recipients, and pluggable channels;
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

## Roadmap governance

Approved goals, dependencies, risks, and gates belong here. Detailed task plans
belong in active prompts; evidence belongs in task histories; product philosophy
belongs in `docs/PRODUCT_ARCHITECTURE.md`. Fixed promises, speculative dates,
pricing, credentials, and implementation without an approved task do not belong
in the roadmap.
