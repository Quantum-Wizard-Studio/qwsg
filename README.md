# Quantum Wizard Server Guardian

## Purpose

Quantum Wizard Server Guardian (QWSG) is an independent, modular Linux server protection project. This repository currently provides its safe engineering and governance foundation.

## Status

Version `0.0.1-prealpha`: project bootstrap only. No application, installer, service, dependency, or operational automation has been implemented.

## Scope

Architecture, implementation, tests, tools, and operating guidance belong here as they are approved. Secrets, credentials, host-specific data, generated artifacts, and undocumented infrastructure changes do not belong in the repository.

Start with [`ai/core/00_PROJECT_PHILOSOPHY.md`](ai/core/00_PROJECT_PHILOSOPHY.md) and [`ai/core/01_CONSTITUTION.md`](ai/core/01_CONSTITUTION.md). This document will evolve throughout development.

The owner-review Product Definition is maintained in [`docs/PRODUCT_DEFINITION.md`](docs/PRODUCT_DEFINITION.md). Established constraints in that document apply now; strategic proposals remain subject to explicit owner approval.

The authoritative product-level relationship between QWSG's identity, boundaries, Agent, Installer, Console, MVP, future direction, and deferred architecture decisions is maintained in [`docs/PRODUCT_SYSTEM_BLUEPRINT.md`](docs/PRODUCT_SYSTEM_BLUEPRINT.md).

The authoritative observable behavior and acceptance boundary for QWSG Core Alpha is maintained in [`docs/FUNCTIONAL_SPECIFICATION.md`](docs/FUNCTIONAL_SPECIFICATION.md).

The authoritative Core Alpha technical design and `Core Alpha Slice 1: Read-only Server Discovery and System Inventory` are maintained in [`docs/architecture/CORE_ALPHA_ARCHITECTURE.md`](docs/architecture/CORE_ALPHA_ARCHITECTURE.md) and [`docs/architecture/CORE_ALPHA_SLICE_1.md`](docs/architecture/CORE_ALPHA_SLICE_1.md).

Slice 1 now has an internal, unsupported Go CLI implementation. Build and test it with `make build` and `make test`, then run `build/qwsg inventory`. See [`docs/development/SLICE_1_DEVELOPMENT.md`](docs/development/SLICE_1_DEVELOPMENT.md). It is not a supported public release.

Canonical System Inventory v1 now provides the authoritative internal Linux host model through the Collector Registry while preserving the Inventory 1.0 compatibility envelope. Its explicitly invoked file-backed [Inventory Persistence and Digital Twin foundation](docs/architecture/INVENTORY_PERSISTENCE_AND_DIGITAL_TWIN.md) can save and reload validated snapshots without monitoring or background execution. See the [developer guide](docs/development/CANONICAL_SYSTEM_INVENTORY.md); user guidance is available in [English](docs/user/CANONICAL_SYSTEM_INVENTORY.en.md) and [Hungarian](docs/user/CANONICAL_SYSTEM_INVENTORY.hu.md).

Engineering tasks follow [`ai/core/11_ENGINEERING_LIFECYCLE.md`](ai/core/11_ENGINEERING_LIFECYCLE.md). The official `ai/scripts/task-builder.sh` workflow generates an approved prompt/history pair from structured owner input after a completed task; `ai/scripts/next-task.sh` remains available when a separate unapproved draft/review cycle is required. Explicitly owner-authorized incomplete-task diversion uses `ai/scripts/divert-task-to-test.sh` to preserve failed evidence under the independent `ai/test_tasks/` namespace without weakening production completion gates. See [`docs/architecture/ENGINEERING_TASK_BUILDER.md`](docs/architecture/ENGINEERING_TASK_BUILDER.md).

The versioned [Reusable Engineering Framework](docs/architecture/REUSABLE_ENGINEERING_FRAMEWORK.md)
validates project identity, canonical Git state, lifecycle paths, required
reading, and project-specific validation commands through
`ai/scripts/framework-check.sh`. QWSG is its reference implementation.
