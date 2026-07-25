# Engineering History

## Purpose

This document indexes major engineering milestones and points to detailed chronological delivery records.

## Status

The first milestone is the `0.0.1-prealpha` documented project bootstrap in `ai/history/000_project_bootstrap.md`.

## Milestones

- `2026-07-18`: Created and verified the recoverable QWSG project foundation.
- `2026-07-18`: The project ACL defaults were corrected to `default:user::rwx` and `default:group::rwx`; a live creation probe confirmed new directories and files inherit owner and group write permission.
- `2026-07-18`: Established `08_JOB_TEMPLATE.md` as the definitive, backward-compatible engineering task standard; adopted English engineering artifacts, Hungarian owner communication, and localization-ready user-facing design. See `ai/history/001_engineering_standard_update.md`.
- `2026-07-18`: Engineering Update E001 refined the official task structure and prompt workflow, and replaced the fixed owner-language rule with a configurable preferred-language policy. See `ai/history/E001_engineering_workflow_refinement.md`.
- `2026-07-18`: Engineering Update E002 introduced guarded sequential prompt rotation, a single active prompt, permanent prompt archives, and independent per-task history. See `ai/history/002_2026-07-18_task-workflow-automation.md`.
- `2026-07-18`: Task 003 consolidated established product constraints and clearly labeled strategic proposals into an owner-review Product Definition without beginning Product Architecture. See `ai/history/003_2026-07-18_product-definition.md`.
- `2026-07-19`: Task 005 consolidated the Product Definition and original comprehensive plan into the authoritative Product & System Blueprint, defining product boundaries, the Agent/Installer/Console model, MVP, future capability groups, and deferred architecture decisions without implementing the system. See `ai/history/005_2026-07-19_product-architecture.md`.
- `2026-07-20`: Task 006 established the authoritative Core Alpha Functional Specification, including testable monitoring, state, alert, CLI, lifecycle, failure-isolation, release-gate, and acceptance behavior without selecting architecture or implementing the product. See `ai/history/006_2026-07-19_functional-specification.md`.
- `2026-07-20`: Task 007 completed an evidence-based repository, documentation-authority, workflow, backup, Quantum Creator conformance, traceability, and Core Alpha readiness audit; it confirmed that product implementation is absent and recommended a bounded architecture milestone before Core Alpha Slice 1. See `ai/history/007_2026-07-20_repository-deep-audit.md`.
- `2026-07-20`: Task 008 established the authoritative Core Alpha architecture and defined Core Alpha Slice 1 as non-root read-only server discovery and system inventory, with data/security contracts, gates, traceability, ADRs, and a bounded Task 009 handoff. See `ai/history/008_2026-07-20_core-alpha-architecture.md`.
- `2026-07-20`: Task 009 implemented the internal Go-based Core Alpha Slice 1 read-only inventory CLI, bounded collectors, versioned JSON, tests, documentation, and rollback evidence. See `ai/history/009_2026-07-20_core-alpha-slice-1-implementation.md`.
- `2026-07-21`: Task 011 established the platform-wide Inventory Architecture: a canonical digital-twin object model, common collector contract, JSON and versioning rules, resource-efficiency and localization contracts, and strict consumer boundaries. See `ai/history/011_2026-07-21_core-inventory-architecture.md`.
- `2026-07-21`: Task 012 implemented the internal Collector Framework with validated descriptors, capabilities and dependencies, duplicate-safe registration, deterministic dependency-aware execution, availability checks, bounded contexts, cancellation, structured results, and collector failure isolation while preserving the existing Inventory output contract. See `ai/history/012_2026-07-21_core-collector-framework.md`.
- `2026-07-21`: Task 013 implemented the official Engineering Task Builder with structured interactive and deterministic multi-line input, automatic metadata and approval generation, validated transactional lifecycle installation, bounded rollback, architecture documentation, and dedicated tests. See `ai/history/013_2026-07-21_engineering-task-builder.md`.
- `2026-07-21`: Task 014 implemented Canonical System Inventory v1: nine bounded Linux collectors, canonical layers/resources/facts and collector results, deterministic validated aggregation, privacy-safe identity/redaction, Registry output-limit enforcement, and additive Inventory 1.0 compatibility. See `ai/history/014_2026-07-21_canonical-system-inventory.md`.
- `2026-07-23`: A Project Owner-authorized aborted-test diversion preserved the
  incomplete Forgejo Task 015 attempt as `001_TEST_TASK`, introduced
  transactional failure containment and independent test-task auditing, and
  released production ID 015 without claiming completion. See
  `ai/audits/2026-07-23_ABORTED_TEST_TASK_DIVERSION.md`.
- `2026-07-24`: Task 015 delivered and verified Reusable Engineering Framework
  `1.0.0`: validated project identity/configuration, canonical Git policy,
  direct-argv validation commands, framework-aware lifecycle entry points, and
  isolated portability/regression coverage. The Project Owner accepted the
  delivery and Task 015 was completed and archived. See
  `ai/audits/2026-07-24_REUSABLE_ENGINEERING_FRAMEWORK.md`.
- `2026-07-24`: Task 016 implemented and verified the first file-backed
  Inventory Store and Digital Twin persistence foundation with versioned
  envelopes, atomic save/load, integrity validation, private storage, bounded
  retention, Inventory 1.0 compatibility, explicit one-shot CLI integration,
  and no monitoring or background execution. Final lifecycle completion remains
  subject to Project Owner acceptance. See
  `ai/audits/2026-07-24_INVENTORY_PERSISTENCE_DIGITAL_TWIN.md`.
- `2026-07-24`: Task 017 implemented and verified the first user-installable
  QWSG CLI and Snapshot Explorer with contextual help, human and JSON output,
  validated snapshot list/info/load commands, controlled build information,
  prefix-aware installation, Ubuntu 24.04 demonstration, and English/Hungarian
  user guidance. The Project Owner accepted the completed delivery. See
  `ai/audits/2026-07-24_QWSG_USER_CLI_SNAPSHOT_EXPLORER.md`.
- `2026-07-24`: Task 018 implemented and verified the canonical Snapshot
  Comparison Engine, deterministic Change Record 1.0 contract, latest/previous
  and explicit CLI comparison, and record-derived JSON/human presentation.
  The Project Owner accepted the installed and verified delivery. See
  `ai/history/018_2026-07-24_snapshot-comparison-engine.md`.
- `2026-07-24`: Task 019 established the canonical QWSG Product Architecture:
  one deterministic core shared by the complete Community engineering toolkit,
  the automation-focused Professional Edition, and the operations-focused
  Provider Edition, with durable workspace, interface, licensing, privacy,
  deployment, automation, AI-separation, ecosystem, and roadmap boundaries.
  See `ai/history/019_2026-07-24_product-architecture.md`.
- `2026-07-24`: Task 020 implemented the Canonical Drift Engine, versioned
  Drift Record 1.0 contract, extensible deterministic taxonomy, and permanent
  Compare-to-Drift boundary without Health or Policy judgement. See
  `ai/history/020_2026-07-24_canonical-drift-engine.md`.
- `2026-07-24`: Task 021 implemented the Canonical Health Engine, versioned
  Health Record 1.0 contract, explicit deterministic status and evidence
  taxonomy, stable Drift-to-Health evaluation, and future Rule, Policy, and
  Report boundaries without monitoring, alerting, remediation, or AI. See
  `ai/history/021_2026-07-24_canonical-health-engine.md`.

Completed milestones, dates, outcomes, and links belong here. Detailed task evidence belongs in independent files under `ai/history/`; this index must not become a continuously growing general task log. Future claims, raw logs, credentials, and rewritten history do not. The index will evolve through concise milestone entries.
