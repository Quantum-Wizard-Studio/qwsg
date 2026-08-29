# Changelog

## [1.2.0-rc.7] - 2026-08-29

- Bound Go-managed Guardian memory below the unchanged 128 MiB systemd cgroup ceiling with measured headroom for subprocess, non-heap, and file-cache accounting.
- Release the completed scheduler execution graph after canonical publication so consecutive loaded-host inventory working sets do not overlap.
- Add repeated-cycle memory diagnostics and preserve deterministic RC.2 migration and rollback through the RC.7 unit replacement.

## [1.2.0-rc.6] - 2026-08-28

- Synchronize guided completion with the newly activated Guardian's first fresh canonical evidence using the setup path's shared bounded waiter.
- Reject preserved stale evidence by lifecycle identity; retain bounded timeout, cancellation, diagnostic exit 4 and optional-notification Partial exit 0 semantics.
- Add deterministic installer completion, delayed evidence, preserved-state, timeout, localization and RC.2-to-RC.6 migration regressions.

## [1.2.0-rc.5] - 2026-08-28

- Add a declarative, fail-closed compatibility path for deterministic updates from the real installed `1.2.0-rc.2` predecessor.
- Revalidate migration compatibility in the privileged transaction boundary and require installed-configuration preflight before Guardian or package mutation.
- Add a credential-free RC.2 fixture proving configuration, credential, state and Guardian-unit preservation plus exact package rollback to RC.2.

## [1.2.0-rc.4] - 2026-08-28

- Added localized administrator notifications for QWSG-managed installation,
  update, rollback, configuration and guided Guardian activation changes.
- Reused Community SMTP configuration and protected credentials while keeping
  lifecycle operation and notification-delivery results independent.
- Preserved deterministic canonical release packaging from RC.3.

## [1.2.0-rc.2] - 2026-08-27

- Added a QWSG-owned guided terminal installer with explicit preflight, plan,
  package, configuration, optional notification, update-policy, activation,
  readiness, and completion phases.
- Added truthful state-derived progress, a terminal dashboard and line-mode
  fallback, plus English, Hungarian, and German catalogs with English fallback.
- Added strict Ubuntu 24.04 LTS amd64 platform gating and reusable installer
  engine contracts for future concise and unattended adapters.
- Added explicit manual/notify update-policy configuration. Unattended
  privileged automatic updating remains disabled.

## [1.2.0-rc.1] - 2026-08-26

- Added canonical public Release discovery and strict SemVer update ordering.
- Added private bounded acquisition plus archive, sidecar, manifest and release-provenance verification.
- Added rollback-capable fixed-destination package transactions preserving configuration, credentials, state and Guardian service intent.
- Added deterministic migration planning and explicit native rollback metadata/workflow.

All notable changes to Quantum Wizard Server Guardian will be recorded here. This changelog currently tracks pre-alpha engineering milestones and will evolve during development; transient logs, secrets, and unverified work do not belong here.

## [Unreleased]

## [1.1.0] - 2026-08-26

### Fixed

- Recovered Guardian processes now replace persisted Scheduler state owned by
  a superseded effective configuration, allowing fresh canonical evidence to
  converge after systemd automatic recovery while retaining same-generation
  restart and failure evidence.

- Added a versioned resumable Smart Setup plan, explicit fixed user-service
  activation, and operator-first README/installation packaging guidance.

### Added

- Versioned read-only Smart Install dependency/capability registry and bounded
  supported-host probes.
- `qwsg install --check` archive bootstrap assessment and composite `qwsg
  readiness` human/JSON reporting.
- Common five-state assessment classifications shared with Community email
  preflight, without automatic remediation or provisioning.

- Optional one-recipient Community SMTP notifications with a private credential
  store, deterministic readiness preflight, bounded retry, and privacy-safe
  Alert messages.

- Task 045 per-user Setup and Configuration foundation with strict XDG
  discovery, atomic private persistence, deterministic CLI management,
  effective Guardian scheduling, invalid-configuration preflight, and future
  secret/notification boundaries without delivery or entitlement.

- Authoritative QWSG Core Alpha Functional Specification covering actors, checks, configuration, state and incident semantics, alerting, reporting, CLI and lifecycle behavior, failure handling, release gates, and testable acceptance criteria.
- Evidence-based repository deep audit, Quantum Creator conformance review, requirements traceability matrix, and Core Alpha readiness assessment.
- Authoritative Core Alpha architecture, inventory data and security models, architecture gates, requirements mapping, and implementation handoff for the read-only Slice 1 milestone.
- Internal Go implementation of the non-root, one-shot, read-only Slice 1 inventory CLI, including bounded collectors, privacy-safe JSON, exit semantics, tests, and developer documentation.
- Canonical System Inventory v1 with host, operating-system, kernel, CPU, memory, storage, filesystem, network, and virtualization collectors; deterministic canonical layers/resources/facts; privacy-safe identifiers; Registry output-limit enforcement; and additive Inventory 1.0 compatibility.

## [1.0.0] - 2026-08-11

### Added

- Final QWSG 1.0 linux-amd64 release identity over the accepted RC.3 product
  baseline and Task 043 clean-host, physical reboot, recurrence, restart and
  uninstall evidence.
- QWS Community / Free License Version 1.0, preserving the complete accepted
  local Community Guardian while reserving future central and managed services.

### Changed

- Reconciled the accepted Task 025–043 source, architecture, packaging and
  lifecycle evidence into an explicit canonical Git source classification.
- Generalized deterministic release assembly to accept final `1.0.0` while
  retaining strict `1.0.0-rc.N` validation.

## [1.0.0-rc.3] - 2026-08-11

### Fixed

- A first observation now safely creates a missing per-user QWSG state
  hierarchy while retaining ownership, exact-mode and symlink protections.
- Valid partial Inventory and Snapshot evidence can establish and publish an
  honest first-use baseline without requiring optional Go tooling.
- First-use bootstrap and Current State publication failures now use bounded,
  privacy-safe diagnostics instead of misleading unreadable-state evidence.

## [1.0.0-rc.2] - 2026-08-11

### Fixed

- Large canonical Policy Reports use one bounded aggregate Alert evidence
  reference while retaining full traceability in the Report envelope.
- Interactive Console refresh reloads Current Operator State without starting
  a competing observation, and privacy-safe Runtime failure causes reach the
  operator view.
- Repeated Attention meaning is correlated into a bounded useful operator view.
- Guardian liveness claims are bounded by exclusive lifecycle freshness after
  graceful or unexpected termination.

## [1.0.0-rc.1] - 2026-08-10

### Added

- Complete local Inventory-to-Report evaluation, bounded Operator Overview,
  durable Current Operator State, local Console, and continuously supervised
  non-root Guardian service.
- Deterministic linux-amd64 Release Candidate archive, SHA-256 manifest,
  fixed-prefix installer/uninstaller, and operator release documentation.

### Fixed

- Interactive Console redraw replaces the prior screen after the single
  initial Overview instead of appending duplicate full-screen output.

## [0.0.1-prealpha] - 2026-07-18

### Added

- Recoverable initial project snapshot and guarded rollback procedure.
- Engineering constitution, philosophy, policies, agent rules, and project record.
- Initial repository and documentation directory structure.
- Temporary proprietary notice and baseline repository hygiene rules.
