# Canonical Configuration Contract

## Status and authority

Task 026 establishes this document as the permanent architecture of QWSG
configuration semantics. It defines Canonical Configuration Model 1.0,
Configuration Source Record 1.0, Effective Configuration 1.0, and Schedule
Definition 1.0.

The contract is an engineering-core boundary. It does not select a user-facing
file format or storage backend and does not implement activation, file
discovery, a configuration editor, secret storage, scheduling, background
execution, networking, or host mutation.

## Purpose

Every configurable QWSG consumer needs one exact answer to three questions:

1. What is the effective value?
2. Which source supplied it?
3. Why did that source win?

`internal/configuration` is the sole canonical owner of those answers. Future
Scheduler, Alert, Dashboard, Terminal UI, REST API, and remote-management code
must consume Effective Configuration rather than create precedence, default,
identity, or validation behavior of its own.

## Boundary

```text
built-in defaults ----+
primary local --------+--> normalize --> validate --> resolve
activated override ---+                                  |
temporary override ---+                                  v
                                              Effective Configuration 1.0
                                                          |
                          +-------------------------------+----------------+
                          |                                                |
               canonical pipeline                             future Scheduler
```

Configuration accepts supplied immutable data. It does not locate, read,
watch, write, activate, or persist configuration files. An eventual activation
component must validate a candidate through this contract and remains a
separate transaction boundary.

## Contracts and versions

The canonical contracts are:

- `qwsg.configuration-source/1.0`: Configuration Source Record 1.0;
- `qwsg.effective-configuration/1.0`: Effective Configuration 1.0;
- Canonical Configuration Model `1.0`;
- Schedule Definition `1.0`.

Unsupported contract or model versions fail. Source content has an independent
declared version. Every normalized source and Effective Configuration has a
SHA-256 identity derived from its complete canonical content with its identity
field omitted.

Canonical JSON is used for stable interchange and identity. Task 045 selects a
strict per-user Source Record file and activation transaction around this
unchanged semantic contract; see `CONFIGURATION_ACTIVATION.md`.

## Source taxonomy and precedence

The exact lowest-to-highest precedence is:

1. `built_in_default` (`0`);
2. `primary_local` (`100`);
3. `activated_local_override` (`200`);
4. `command_temporary_override` (`300`).

This implements `FR-CFG-002`. A higher-precedence explicit field replaces the
complete lower-precedence field. Collections use explicit replacement; there
is no implicit list append or map overlay.

Sources at equal precedence are sorted by stable source ID only for processing
and evidence. Equal values merge their source provenance. Different values are
an explicit conflict and resolution fails. Input enumeration order is never a
tie-breaker.

## Field provenance

Every Effective Configuration field has a provenance record containing:

- canonical field name;
- sorted source IDs and source versions;
- source kind and numeric precedence;
- resolution result: `selected`, `overridden`, or `equal_values_merged`.

Built-in defaults are ordinary versioned source values and therefore never
hidden. Missing required fields fail resolution.

## Canonical Configuration Model 1.0

The model contains typed configuration for:

- instance identity, locale, and time-zone basis;
- checks and targets with stable IDs and references;
- Canonical Rule Definitions and Policy Profiles;
- snapshot retention;
- Schedule Definitions;
- global execution timeout and bounded concurrency;
- bounded Retry Policies;
- report policy;
- typed secret references;
- optional non-required extension sections.

All collections are deterministically sorted with unique identities. Strings,
collections, metadata, durations, counts, retries, concurrency, retention,
priorities, and graph references are bounded. Unknown JSON keys fail in strict
decoding. Cross-references must resolve after precedence is applied.

Rule semantics remain owned by the Rule Engine. The Rule package exposes
configuration-time definition validation without evaluating Health evidence.
Policy semantics remain owned by the Policy Engine and Policy Profiles pass its
canonical validation.

## Schedule Definition 1.0

Schedule Definition describes scheduler input without executing time-based
behavior. It includes:

- stable ID and contract version;
- enabled state;
- explicit time-zone basis;
- interval or calendar trigger;
- explicit daylight-saving policy;
- bounded priority;
- `skip`, `run_once`, or `indeterminate` misfire policy;
- `forbid` or `allow` overlap policy;
- bounded execution timeout;
- Retry Policy reference;
- sorted Check references;
- Command profile applicability.

An interval has one positive bounded nanosecond duration and no calendar
fields. A calendar has valid sorted minute and hour fields, optional month-day,
month, and weekday restrictions, and no interval. Time zones use `UTC` or an
unambiguous IANA-style identifier. Locale cannot change schedule identity or
meaning.

Task 026 validates this data but never reads a clock, calculates due work,
starts a timer, handles a misfire, locks a job, retries execution, or runs a
Command. Those responsibilities belong to Task 027 Scheduler architecture.

## Secret and privacy contract

Canonical configuration can contain only typed secret references: stable ID,
provider ID, and non-secret opaque reference. The contract has no secret-value
field. Strict decoding rejects an attempted value field. Secret retrieval,
storage, credentials, and backend choice remain unimplemented.

Canonical serialization, identities, provenance, and diagnostics contain no
resolved secret material. Metadata and extension values remain bounded and are
subject to publication and redaction policy.

## Extensions and compatibility

Extensions declare stable identity, version, whether they are required, and
bounded string fields. Unknown required behavior is rejected. Optional
extensions may be preserved as inert typed data; no behavior activates merely
because an extension is present.

The built-in source reproduces the existing canonical observation Rule,
Policy Profile, retention, and pipeline behavior. Existing `Orchestrator`
retention, Rule, and Policy fields remain compatibility inputs but are first
projected through the same configuration resolver. Supplying an Effective
Configuration together with compatibility fields is ambiguous and fails.

Current CLI commands, command plans, stage ordering, output contracts, and
engine semantics remain unchanged.

## Pipeline integration

`internal/pipeline` remains the only canonical engine orchestrator. Before any
stage executes, it obtains one validated Effective Configuration. Rule and
Policy stages consume its canonical definitions and profiles; Snapshot Store
access consumes its retention. Future Scheduler consumes its schedules,
timeouts, retry policies, concurrency, priorities, and applicability.

Configuration does not execute the pipeline. Command and presentation layers
do not resolve configuration.

## Determinism and immutability

Normalization deep-copies caller data before sorting or defaulting. Resolution
does not mutate sources. Equal semantic source sets produce the same normalized
sources, Effective Configuration, canonical JSON bytes, identity, ordering,
and provenance regardless of input or map enumeration order.

Malformed identities, unsupported versions, invalid definitions, duplicate
IDs, contradictory equal-precedence values, invalid references, oversized
values, ambiguous schedules, unsupported required extensions, and tampered
effective records fail explicitly.

## Resource limits

The implementation bounds source count, total configured items, string and
metadata sizes, concurrency, retention, retries, priorities, and durations.
Schedule, check, target, Rule, Policy, retry, secret-reference, and extension
collections are finite and validated before consumption.

## Explicit exclusions

This contract implements no Scheduler, due-work computation, timer, polling,
daemon, service, locking, worker pool, alert, incident, notification, e-mail,
webhook, file activation, watch, hot reload, UI, API, secret backend, network
call, process execution, remediation, AI, or host mutation.

## Scheduler consumption

The Canonical Professional Scheduler selects and executes due work only from a validated
Effective Configuration. It must not redefine:

- configuration identity or versioning;
- source precedence or defaults;
- schedule syntax or validation;
- time-zone, daylight-saving, misfire, overlap, retry, timeout, concurrency, or
  priority configuration semantics;
- Check, target, Command profile, Rule, or Policy references.

Task 027 adds operational scheduling behavior around these immutable contracts
while keeping manual canonical pipeline execution available. Its permanent
boundary is `CANONICAL_SCHEDULER.md`.

## Alert consumption

The pure Canonical Alert Engine may consume a validated Effective
Configuration only for exact configuration identity, provenance, and existing
applicability context. Alert State binds to that identity when supplied. Task
028 does not add alert, acknowledgement, maintenance, suppression, recipient,
channel, or notification fields to Configuration Model 1.0. Bounded
acknowledgements and suppression windows are explicit Alert-owned evaluation
inputs until a separately authorized Configuration version defines them.

## Notification Delivery boundary

Task 029 does not extend Effective Configuration 1.0. Delivery Policy, Route,
Endpoint Reference, Provider Binding, opaque destination/secret references,
previous Queue State, and explicit time are Notification-owned inputs. Concrete
configuration integration, secret resolution, provider activation, durable
queue selection, and daemon lifecycle require a separately authorized
Configuration version and security review.

Runtime accepts one already resolved and validated Effective Configuration. It
does not discover, merge, override, activate, or reinterpret configuration.
