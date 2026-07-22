# Current Engineering Task 012: Core Collector Framework

## Task Metadata

- Task ID: `012`
- Task slug: `core-collector-framework`
- Status: `complete`
- Date opened: `2026-07-21` UTC
- Human authority: `Project Owner`
- Owner or lead-developer communication language: `Hungarian`

## Title

Design and implement the official Collector Framework for the Quantum Wizard Server Guardian platform.

## Objective

Design and establish the canonical Collector Framework used by the QWSG Core.

The framework defines how collectors are identified, registered, executed, validated and integrated into the Inventory Architecture introduced in Task 011.

Every collector shall follow a common contract, contribute to the canonical Inventory model, operate independently and remain resource-efficient.

The Collector Framework becomes the only supported mechanism for future Inventory collection.

## Scope

- Design the canonical Collector interface.
- Implement the Collector Registry.
- Define the collector lifecycle.
- Implement the capability model.
- Implement availability and dependency checks.
- Implement the standard collector result object.
- Implement error isolation.
- Support timeout and context cancellation.
- Ensure deterministic execution order.
- Migrate existing Inventory collector(s).
- Update architecture documentation.
- Add unit tests.

## Out of Scope

- New Inventory domains
- Policy Engine
- Configuration Engine
- REST API
- Dashboard
- Web interface
- Daemon mode
- Automatic remediation
- Dynamic external plugins
- Remote collectors

Task 012 establishes the internal Collector Framework only.

## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`

## Starting State Verification

- Verify Task 011 is complete.
- Verify Inventory Architecture consistency.
- Execute `bin/job --check`.
- Verify working tree status.
- Create verified snapshot.
- Record baseline test status.

## Snapshot Requirements

Create a complete snapshot before any modification, including source tree, AI documentation, collector implementation, tests and configuration. Record the snapshot location and rollback procedure in Task History.

## Risk Assessment

Primary risks:
- incompatible collector interfaces
- duplicated Inventory logic
- collector coupling
- unstable execution order
- Inventory regression
- excessive resource consumption

The framework shall remain fully compatible with Task 011.

## Planned Work

1. Analyse current collectors.
2. Design the canonical Collector interface.
3. Design and implement the Collector Registry.
4. Implement registration.
5. Implement capability model.
6. Implement availability and dependency checks.
7. Implement standard collector result.
8. Implement deterministic execution.
9. Implement timeout and context handling.
10. Migrate existing collectors.
11. Update documentation.
12. Execute complete verification.

## Rollback Plan

If verification fails, restore the verified snapshot, validate hashes, rerun baseline tests, document the rollback and leave the repository in the verified pre-task state.

## Deliverables

- Collector Framework
- Collector Registry
- Collector interface
- Collector result model
- Capability model
- Updated Inventory integration
- Updated architecture documentation
- Updated Engineering History
- Tests

## Verification

Verify:
- `bin/job --check`
- Registry operation
- Duplicate registration rejection
- Deterministic execution
- Timeout handling
- Dependency checks
- Existing Inventory compatibility
- Canonical Inventory output
- Go tests
- `git diff --check`

## Documentation Updates

Update:
- `ai/core/04_ARCHITECTURE.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/13_ROADMAP.md`

Create additional architecture documentation if required.

## Completion Criteria

Task 012 is complete only when:
- Collector Framework exists.
- Registry implemented.
- Collector contract implemented.
- Existing collectors migrated.
- Inventory compatibility preserved.
- Documentation updated.
- Verification successful.
- Snapshot and rollback documented.
- No unresolved architectural inconsistencies remain.

## Engineering Principle

> Collectors observe and describe. They do not evaluate, decide, repair or modify the system. Their sole responsibility is to produce accurate, deterministic and canonical Inventory data.

## Owner Approval Requirements

Approved by the Project Owner.

This task has been reviewed and explicitly approved for implementation.

Implementation may begin in accordance with the Engineering Lifecycle.

Further scope changes require a new explicit Project Owner approval.
