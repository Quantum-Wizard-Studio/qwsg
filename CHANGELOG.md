# Changelog

All notable changes to Quantum Wizard Server Guardian will be recorded here. This changelog currently tracks pre-alpha engineering milestones and will evolve during development; transient logs, secrets, and unverified work do not belong here.

## [Unreleased]

### Added

- Authoritative QWSG Core Alpha Functional Specification covering actors, checks, configuration, state and incident semantics, alerting, reporting, CLI and lifecycle behavior, failure handling, release gates, and testable acceptance criteria.
- Evidence-based repository deep audit, Quantum Creator conformance review, requirements traceability matrix, and Core Alpha readiness assessment.
- Authoritative Core Alpha architecture, inventory data and security models, architecture gates, requirements mapping, and implementation handoff for the read-only Slice 1 milestone.
- Internal Go implementation of the non-root, one-shot, read-only Slice 1 inventory CLI, including bounded collectors, privacy-safe JSON, exit semantics, tests, and developer documentation.
- Canonical System Inventory v1 with host, operating-system, kernel, CPU, memory, storage, filesystem, network, and virtualization collectors; deterministic canonical layers/resources/facts; privacy-safe identifiers; Registry output-limit enforcement; and additive Inventory 1.0 compatibility.

## [0.0.1-prealpha] - 2026-07-18

### Added

- Recoverable initial project snapshot and guarded rollback procedure.
- Engineering constitution, philosophy, policies, agent rules, and project record.
- Initial repository and documentation directory structure.
- Temporary proprietary notice and baseline repository hygiene rules.
