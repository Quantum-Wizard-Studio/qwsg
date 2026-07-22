# QWSG Core Alpha Architecture

## Authority and scope

This is the authoritative technical architecture for Core Alpha and the governing architecture for `Core Alpha Slice 1: Read-only Server Discovery and System Inventory`. It is subordinate to the Constitution and the ratified baseline named by Task 008. The Functional Specification remains authoritative for observable behavior. Product Definition proposals remain proposals unless Task 008 explicitly ratified them; this document does not settle commercial, platform-support, retention, Console trust, transport, or cryptographic policy.

No runtime, language, framework, database, protocol, package format, or cloud service is selected here. Slice 1 implementation may select a proportional stack only through a later authorized task and must preserve these contracts.

## Architecture principles

- Local-first, single-host operation; no vendor cloud dependency or telemetry.
- Default non-root collection and explicit truthful loss of visibility.
- Read-only collectors cannot remediate, configure, install, restart, signal, or write monitored resources.
- Evidence is separated from interpretation, and unknown is never healthy.
- Partial success is a first-class result; one collector cannot invalidate unrelated results.
- Contracts, schemas, configuration, and stored envelopes are versioned before implementation.
- Agent operation is independent of Console availability.
- Privileged lifecycle work, future remediation, and discovery are separate authority domains.

## System context and trust boundaries

The protected Linux host is an untrusted evidence source. The local operator invokes the Agent-facing CLI. The Agent coordinates bounded collectors, normalization, validation, inventory assembly, optional local persistence, and presentation. A future Console is a consumer of Agent-owned contracts, never a privileged command channel. External Console transport and authentication remain gate `AG-004`.

Trust boundaries are:

1. operator input to CLI or configuration;
2. coordinator to collector invocation;
3. collector to operating-system files, process metadata, and allowlisted commands;
4. untrusted raw evidence to normalization and schema validation;
5. Agent state to local storage;
6. Agent contract to Console or other presentation;
7. future Installer/remediation authority, which is outside Slice 1.

## Component map

| Component | Owns | Must not own in Slice 1 |
| --- | --- | --- |
| Agent entry point | command parsing, request identity, cancellation, result and exit contract | privilege escalation, arbitrary commands |
| Discovery coordinator | collector registry, deadlines, bounded concurrency, deterministic aggregation | parsing collector-specific evidence |
| Collectors | one declared inventory category, capability probe, bounded raw evidence | storage, presentation, cross-collector policy, mutation |
| Normalizer | typed canonical values, units, identifiers, redaction labels | host access or command execution |
| Validator | envelope/schema/invariant enforcement | repairing malformed results |
| Inventory assembler | snapshot status, category results, provenance, completeness | interpreting unknown as healthy |
| State adapter | atomic write/read of versioned inventory envelopes if enabled | selecting retention duration |
| Presentation adapter | stable machine output and localizable human rendering | changing canonical state |
| Console adapter | future read contract and authorization boundary | direct collector or shell access |
| Installer | future privileged lifecycle boundary | routine discovery execution |

## Discovery flow

`request -> validate -> plan collectors -> run bounded probes -> normalize -> redact -> validate -> assemble -> optionally persist atomically -> render`

Cancellation stops pending work and requests termination of active probes. Completed categories remain reportable with snapshot status `cancelled` or `partial`. Every subprocess has a monotonic deadline, output limit, fixed executable and argument grammar, sanitized environment, closed stdin, and captured categorized failure.

## Collector contract

Each collector declares a stable ID, contract version, category, supported evidence sources, privilege class, sensitivity classification, timeout ceiling, output ceiling, and platform capability predicate. Input is an immutable request containing request ID, deadline, locale-independent options, and cancellation signal. Output is exactly one category result defined by the data model.

Collectors return facts only. They do not emit user-facing prose, write persistent state, call other collectors, construct shell strings, follow unsafe symlinks, or accept executable names and arbitrary paths from untrusted input. Platform adapters may vary evidence acquisition while producing the same contract.

## Agent and Console boundaries

The Agent owns authoritative local inventory and schema validation. The Console, if later approved, reads a versioned redacted projection and presents it; it cannot reinterpret missing data, invoke arbitrary commands, or gain privileges beyond a separately authenticated Agent interface. Slice 1 requires CLI/JSON presentation and permits a local in-process view. Network transport, sessions, recovery, and exposure are blocked by `AG-004`.

## Storage and transient state

A single discovery run may operate fully in memory. If the implementation task enables persistence, only complete validated envelopes or explicitly partial validated envelopes may be atomically committed to a QWSG-owned location. The previous valid envelope remains readable until replacement succeeds. Storage technology and retention duration remain gates; Slice 1 needs at most the latest inventory plus bounded diagnostic metadata. Raw command output is transient and is not retained by default.

## Errors and degraded operation

Canonical category outcomes are `available`, `unavailable`, `unsupported`, `permission_denied`, `timeout`, `error`, and `cancelled`. Snapshot outcomes are `complete`, `partial`, `failed`, and `cancelled`. Each failure has a stable code, safe summary, retryability flag, and provenance. Retries are absent by default for local deterministic probes; an implementation may retry once only for a documented transient condition within the original deadline. Failure in one category does not stop independent categories.

## Configuration and versioning

Slice 1 configuration is limited to enabling categories, privacy-safe field policy, finite time/output limits within architecture maxima, and optional latest-envelope persistence. Unknown keys fail validation. Configuration syntax is a later implementation decision. Contract versions use `major.minor`: incompatible removal or semantic change increments major; additive optional fields increment minor. Readers reject unsupported major versions and preserve unknown additive fields only where explicitly safe.

## Logging and observability

Structured operational events contain UTC timestamp, monotonic duration where relevant, severity, component, request ID, collector ID, outcome, safe error code, and redaction count. They exclude raw environment dumps, command output, secrets, full host identifiers, and sensitive network fields. Audit events identify operator-requested runs and policy changes; ordinary automatic read-only collection is operational history, not an authorization event. Scheduler, persistence, collector freshness, and schema failures become future product-health subjects.

## Packaging, update, and extension boundaries

Package format, service topology, runtime, dependency set, and update authenticity are not selected. Implementation must be reproducible, locked, testable without root, and unable to claim supported release status before gates close. Extensions register through the collector contract; they cannot bypass coordinator limits, broaden privilege, add network transmission, or mutate the host.

## Testing architecture

Required layers are contract/schema tests, collector unit tests using fixtures, platform-adapter tests, coordinator failure-isolation tests, security/adversarial tests, storage fault tests if persistence is enabled, CLI/JSON golden tests, and declared-platform integration tests. Root and mutation assertions are negative acceptance tests. Deterministic fixtures cover every outcome and sensitive-data redaction.

## Governing documents

- [Slice 1 definition](CORE_ALPHA_SLICE_1.md)
- [Inventory data model](CORE_ALPHA_DATA_MODEL.md)
- [Security model](../security/CORE_ALPHA_SECURITY_MODEL.md)
- [Implementation plan](../development/CORE_ALPHA_IMPLEMENTATION_PLAN.md)
- [Architecture gates](../development/ARCHITECTURE_GATES.md)
- [Requirements mapping](../development/REQUIREMENTS_ARCHITECTURE_MAPPING.md)
