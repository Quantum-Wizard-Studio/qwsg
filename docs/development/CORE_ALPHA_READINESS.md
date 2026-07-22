# QWSG Core Alpha Readiness Review

## Decision

QWSG is **not ready for direct Core Alpha implementation**, but it is ready for a bounded Core Alpha Architecture milestone. The Functional Specification provides a strong behavioral contract; the missing architecture and data/security/test contracts are prerequisites to safe coding.

The narrow initial slice proposed by Task 007 is a sound engineering sequence if the owner names it “Core Alpha Slice 1” and does not represent it as complete Core Alpha functionality.

## Slice readiness

| Item | Classification | Evidence and condition |
| --- | --- | --- |
| System discovery | NOT_READY | Capability behavior exists, but support matrix, inventory schema, platform adapters, and privilege boundary are missing. |
| Disk usage monitoring | READY_WITH_CONDITIONS | Threshold/evidence/state requirements exist; define filesystem scope, normalized observation schema, collector contract, and fixtures. |
| Inode monitoring | READY_WITH_CONDITIONS | Same requirements as disk; define independent subject identity and incident behavior. |
| State transitions | READY_WITH_CONDITIONS | Functional semantics are detailed; select persistence model, transactions, monotonic/wall-clock handling, schema versioning, and restart algorithm. |
| Warning events | READY_WITH_CONDITIONS | Entry/persistence/alert rules are specified; define event schema and delivery boundary. |
| Critical events | READY_WITH_CONDITIONS | Escalation semantics are specified; define durable ordering and failure handling. |
| Recovery events | READY_WITH_CONDITIONS | Full recovery vs de-escalation is explicit; define restart-safe incident closure transaction. |
| Structured local logging | NOT_READY | Categories are required, but record schema, format, sink, rotation, retention, permissions, and injection-safe encoding are undefined. |
| Configuration validation | READY_WITH_CONDITIONS | Validation/precedence/activation behavior is specified; choose syntax, schema mechanism, storage, permissions, and secret references. |
| CLI status command | READY_WITH_CONDITIONS | Required human/JSON behavior exists; define versioned JSON schema, permission model, and service boundary. |
| CLI check command | READY_WITH_CONDITIONS | On-demand parity and exit classes exist; define target lookup, concurrency, transport, and partial-result representation. |
| Dry-run behavior | NOT_READY | Lifecycle preview and plan binding exist, but a general dry-run term/flag, machine contract, side-effect boundary, and exit semantics are not normative. |
| Unit tests | READY_WITH_CONDITIONS | Acceptance cases are detailed; choose runtime/test framework and deterministic clock/filesystem/process fixtures. |
| Isolated integration tests | READY_WITH_CONDITIONS | Failure and lifecycle cases are specified; define sandbox boundaries, fake systemd/mail/storage, and no-root CI execution. |

## Missing decisions and contracts

### Product decisions

- Ratify downstream use of Product Definition proposals.
- Confirm the first slice is an internal milestone, not the complete Core Alpha release.
- Console shipment, e-mail transport, retention, licensing, and hosted-service decisions may remain release gates for Slice 1.

### Architecture decisions

- Runtime/language and dependency/package strategy.
- Agent process, scheduler, CLI, module, and privileged Installer boundaries.
- Filesystem paths, users/groups, ownership, permissions, and service lifecycle.
- Error/failure isolation, concurrency, timeouts, restart, and upgrade topology.

### Data contracts

- Stable IDs; observation, state, incident, event, audit, log, and JSON schemas.
- Schema versions, transactions, durability, retention, cleanup, export, migration, corruption isolation, and clock semantics.

### Security decisions

- Threat model and least-privilege boundary.
- Configuration and secret-reference permissions/redaction.
- Input/output encoding, terminal safety, update integrity, dependency provenance, and audit integrity.
- Console security may be deferred if Console is excluded from Slice 1.

### Test contracts

- Test framework and repository layout.
- Deterministic clock, filesystem, mount, capacity, restart, full-storage, and corruption fixtures.
- Contract tests for CLI JSON and exit codes.
- Non-root isolated integration environment and coverage/release gates.

### Hidden dependencies

- Reading mount/inode state reliably across filesystem types and namespaces.
- Stable CPU/container capability discovery even if CPU monitoring follows later.
- A durable state store before meaningful recovery/restart behavior can be claimed.
- Scheduler and locking behavior behind on-demand `check` parity.
- Permissions that let read-only CLI users inspect state without exposing secrets.

## Unnecessary scope for Slice 1

Console, cloud/fleet functions, automatic remediation, database/mail/log/hardware checks, package installation, e-mail delivery, full update/uninstall, and public-release support are unnecessary for the first implementation slice. They remain valid later requirements and must not be deleted or silently weakened.

## Required entry gate

Implementation may begin only after an approved architecture record defines at minimum:

1. product authority inputs and Slice 1 boundary;
2. runtime/package/test stack;
3. privilege, process, module, and filesystem boundaries;
4. observation/state/incident/config/log/CLI contracts;
5. restart, failure, concurrency, and migration rules;
6. deterministic test strategy and acceptance mapping;
7. snapshot, rollback, and version-control baseline.

## Recommendation

Authorize a documentation-only `core-alpha-architecture` task next. After owner approval of its decisions, implement Slice 1 as a small vertical path: validated configuration -> disk/inode observation -> retained transition -> structured local event/log -> `qwsg status`/`qwsg check` -> deterministic unit and isolated integration tests.
