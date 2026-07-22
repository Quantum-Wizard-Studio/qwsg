# Core Alpha Slice 1: Read-only Server Discovery and System Inventory

## Purpose

Slice 1 gives a local operator one truthful, structured inventory of a Linux server without changing it. It proves the Agent's collector, normalization, schema, partial-result, security, CLI, and test boundaries. It is an internal implementation milestone, not the complete Core Alpha product or a supported production release.

## Included capabilities

- One-shot local discovery invoked through the future `qwsg` CLI.
- Machine-readable, versioned JSON and localizable human presentation.
- Operating-system identity/version; kernel and machine architecture.
- Privacy-safe host identity: display label and pseudonymous local instance identifier; no default raw hardware serials.
- CPU topology/capacity summary and memory/swap summary.
- Mounted filesystem/storage summary without file-content traversal.
- Network-interface summary with addresses and hardware identifiers redacted or omitted by default.
- Running-service summary only through safely accessible service-manager metadata; no service control.
- Installed runtime/server-component versions only from allowlisted, bounded detection mechanisms.
- Collector self-capability and permission report.
- Per-field provenance, observation/completion timestamps, freshness deadline, sensitivity, and quality.
- Honest `unsupported`, `unavailable`, `permission_denied`, `timeout`, `error`, `cancelled`, and partial results.
- Optional atomic persistence of the latest validated envelope only if the implementation task explicitly includes it.

## Non-goals

No continuous monitoring, scheduler, thresholds, health severity, incidents, alerts, e-mail, daily reports, remediation, configuration management, package installation, service control, backup inspection, HTTP/TLS probing, remote command execution, network Console, fleet management, telemetry, update, or lifecycle mutation is included.

## Read-only and permission behavior

The normal process runs non-root. It never invokes `sudo`, `su`, polkit, setuid helpers, capabilities elevation, or a privileged daemon. A category needing unavailable privilege returns `permission_denied`, names the missing capability generically, and continues. No implementation may recommend broad root execution as the normal remedy. Any future privileged companion requires a separate architecture and owner authorization.

## Execution assumptions

Slice 1 targets Linux with procfs/sysfs-like evidence and may provide adapters for distributions later declared supported. Exact Ubuntu/Debian versions and CPU architectures remain gate `AG-001`. Running in containers, restricted VPS environments, chroots, namespaces, or non-systemd hosts is supported only to the extent detected capabilities truthfully allow; absence is not failure of unrelated categories.

## Result contract

The output is the `InventorySnapshot` envelope in [CORE_ALPHA_DATA_MODEL.md](CORE_ALPHA_DATA_MODEL.md). Each category is independently valid and ordered by stable category ID. Human rendering derives only from the envelope. Exit behavior for Slice 1 is: `0` complete; `3` request-level permission denial; `4` requested capability unsupported/unavailable; `5` internal validation/storage failure. A partial multi-category result must not return a misleading complete success; the implementation plan must choose and document a stable partial-result CLI policy before coding.

## Security and privacy

Collection is purpose-limited to inventory. Secrets, credential material, environment values, command lines with arguments, file contents, application configuration, listening-connection payloads, user data, and arbitrary process data are prohibited. Network discovery is local metadata only; no outbound probes occur. Hostname, addresses, interface hardware identifiers, mount sources, service names, and component paths carry sensitivity labels and default redaction rules.

## Presentation boundary

Slice 1 requires a CLI human view and versioned JSON. User text is externalizable and machine tokens are locale independent. Console presentation is deferred; any future Console consumes the redacted envelope and cannot invoke collectors directly.

## Acceptance criteria

1. A non-root fixture-backed run returns a schema-valid envelope containing every enabled category or an explicit categorized absence.
2. No test observes file, service, package, network, process, permission, ownership, timestamp, or configuration mutation.
3. Each collector enforces a deadline and output limit; one timeout or malformed result leaves unrelated results available.
4. Permission denial, unsupported environment, missing command, cancellation, and partial completion are distinguishable.
5. Seeded secrets and terminal/markup control payloads do not appear in JSON, human output, or logs.
6. Timestamps are UTC with offsets, durations use monotonic time, and provenance identifies source type without leaking unsafe source content.
7. Unknown major schema versions are rejected without overwrite; additive compatible fields do not corrupt known data.
8. JSON output is deterministic for deterministic fixtures and human strings are not stored as canonical values.
9. If persistence is included, interrupted or failed writes preserve the previous valid envelope.
10. Declared-platform integration tests pass before any platform is described as supported.

## Traceability

Direct inputs: `FR-AUTH-001`, `FR-AUTH-003`, `FR-CAP-001` through `FR-CAP-005`, `FR-CHECK-001` through `FR-CHECK-007`, `FR-CLI-001` through `FR-CLI-005`, `FR-DATA-002`, `FR-DATA-004`, `FR-NFR-001` through `FR-NFR-006`, `FR-REL-002`, and `FR-DIAG-001`. Architecture support mechanisms that are not new product promises are the inventory category set, envelope, collector interface, and optional latest-snapshot persistence.

Relevant acceptance inheritance: `AC-AGENT-002`, `AC-AGENT-004`, `AC-UX-002`, `AC-UX-004`, `AC-DATA-001`, `AC-DOC-001`, and `AC-REL-001`. Slice 1 does not claim full satisfaction of any broader Core Alpha acceptance criterion.

## Implementation stages and completion

1. Select the proportional runtime/package approach under Task 009 and establish locked build/test scaffolding.
2. Implement contracts, validation, errors, redaction, and fixture harness.
3. Implement coordinator and non-root collectors category by category.
4. Add CLI/JSON presentation and optional latest-envelope adapter if authorized.
5. Run security, fault, platform, permission, and no-mutation acceptance tests.

Slice 1 is complete only when all ten criteria pass, documentation names actual supported contexts, no gate is represented as closed without approval, and no excluded behavior exists. Public/production release remains blocked by all applicable gates in `ARCHITECTURE_GATES.md`.
