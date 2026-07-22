# Core Alpha Slice 1 Implementation Plan

## Recommendation

Task 009 implemented the internal slice using repository-pinned Go `1.26`, one self-contained CLI binary, standard-library-only dependencies, and exit codes `0` complete, `1` fatal, and `2` partial. No supported release is claimed. Task 010 should independently verify and harden the slice.

## Prerequisites

- Owner-authorized Task 009 with implementation permission, snapshot, rollback, and exact source scope.
- Runtime/package decision record using the criteria in `AG-006`; locked dependencies and reproducible commands.
- Stable JSON Schema or equivalent generated schema for the data model.
- Stable partial-result CLI exit behavior under `AG-011`.
- Test fixtures that contain no production host data.
- Explicit declaration of whether latest-envelope persistence is included; if yes, bounded path/ownership and atomic-write design.

## Permitted source areas recommendation

Task 009 should be limited to `agent/`, `modules/`, `tests/`, a narrowly justified shared-contract area, and required build/package metadata. `installer/` and `console/` remain excluded. No service unit, privileged helper, host configuration, dependency installation, production deployment, or network listener should be included unless separately authorized.

## Dependency order

1. Build/test scaffolding and dependency locks.
2. Contract types/schema, invariants, error catalog, sensitivity model, and redaction.
3. Fixture-only collector interface and coordinator with deadlines/cancellation/output bounds.
4. CLI request and deterministic JSON/human presentation.
5. Non-root collectors: capability report, OS/kernel/architecture, host identity, CPU, memory, storage, network, services, components.
6. Optional latest-envelope persistence behind a contract and fault tests.
7. Platform adapters and controlled integration tests.
8. Security/adversarial/no-mutation test suite and documentation.

## Required test layers

- Unit tests for parsing, normalization, redaction, validation, and errors.
- Contract tests for every collector and schema version.
- Coordinator tests for timeout, cancellation, malformed evidence, output exhaustion, and partial success.
- Golden tests for locale-independent JSON and localizable human rendering.
- Security tests listed in the security model.
- Storage crash/fault tests if persistence is included.
- Integration tests on every platform before it enters the support matrix.
- Acceptance mapping tests tied to Slice 1 criteria and existing FR/AC identifiers.

## Prohibited shortcuts

No root-as-default, shell strings, arbitrary command/path configuration, raw command-output persistence, invented fallback values, `UNKNOWN` to healthy conversion, unbounded reads/processes/retries, shared mutable global collector state, direct Console access, outbound telemetry, package installation during tests, production-host fixtures, silent schema reset, or requirement promotion.

## Task 010 verification and hardening

Task 010 should review architecture conformance, fuzz parsers and schema boundaries, test restrictive permissions and containers/VPS contexts, measure resource bounds, verify clean installation/build reproducibility, audit dependencies, run no-mutation probes, validate documentation, and decide whether any platform can be entered into `AG-001`. It must leave unsupported contexts explicit.

## Governing documents

Implementation is governed by `CORE_ALPHA_ARCHITECTURE.md`, `CORE_ALPHA_SLICE_1.md`, `CORE_ALPHA_DATA_MODEL.md`, `CORE_ALPHA_SECURITY_MODEL.md`, `ARCHITECTURE_GATES.md`, and `REQUIREMENTS_ARCHITECTURE_MAPPING.md`, plus the Constitution and Functional Specification.
