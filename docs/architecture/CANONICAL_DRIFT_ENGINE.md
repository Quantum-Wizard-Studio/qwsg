# Canonical Drift Engine

## Permanent architectural boundary

The Drift Engine is QWSG's first semantic interpretation layer. It answers only
what kind of change occurred:

```text
Inventory -> Snapshot Store -> Comparison Engine -> Drift Engine
                                                        |
                         +------------------------------+----------------+
                         |                              |                |
                  future Health                   future Rules     future Policy
```

Inventory defines what exists. Snapshots preserve what existed. Comparison
produces factual Change Records. Drift classifies those records. The canonical
Health Engine evaluates the engineering condition represented by validated
Drift Records. Rules, Policy,
Reports, interfaces, and Automation consume Drift and Health contracts; they
must not compare snapshots or recreate drift classification.

`internal/drift` is a pure, offline, AI-independent Go package. It consumes only
validated `comparison.ChangeRecord` values and produces only canonical Drift
Records. It performs no collection, storage, host access, networking,
scheduling, logging, policy evaluation, scoring, diagnosis, recommendation, or
mutation. The existing `qwsg compare` interface and Change Record 1.0 contract
remain unchanged.

## Drift Record 1.0

The `qwsg.drift` schema, engine, and taxonomy begin at version `1.0`. A record
contains:

| Field | Contract |
| --- | --- |
| `id` | Stable SHA-256-derived Drift ID. |
| `change_id` | The sole canonical Change Record that caused this record. |
| `category` | One taxonomy category below. |
| `scope` | Canonical `layer`, privacy-safe `object_id`, and escaped `path`. |
| `classification` | `presence_added`, `presence_removed`, `value_modified`, or `state_unchanged`. |
| `confidence_basis_points` | Deterministic integer from 0 through 10000; never a probability or risk score. |
| `timestamp` | The source comparison timestamp, normalized to UTC. |
| `metadata` | Bounded classifier rule and source change-type tokens only. |
| `versions` | Drift schema, engine, taxonomy, and Change schema versions. |

The envelope repeats the public schema and engine versions, contains an ordered
`records` array, and declares its exact input contract. Every accepted Change
Record yields exactly one Drift Record, including `unchanged` comparison
evidence. Drift records never copy previous/current values or arbitrary source
metadata, reducing duplication and preventing downstream value disclosure.

Confidence describes only how specifically the taxonomy rule matched. Exact
registered-layer rules use 10000, the registered storage subtype rule uses
9500, and the forward-compatible extension rule uses 5000. It does not express
health, likelihood, severity, importance, trust, or operational risk.

## Taxonomy 1.0

| Category | Canonical responsibility |
| --- | --- |
| Identity Drift (`identity`) | Subject, account, and physical/virtual identity context. |
| Software Drift (`software`) | Installed software, applications, and package state. |
| Hardware Drift (`hardware`) | Compute devices and visible physical capacity. |
| Platform Drift (`platform`) | Operating system, kernel, distribution, and runtime platform. |
| Filesystem Drift (`filesystem`) | Filesystem and mount semantics within storage. |
| Storage Drift (`storage`) | Block devices, volumes, and other storage semantics. |
| Network Drift (`network`) | Interfaces, addresses, routes, and local topology facts. |
| Service Drift (`service`) | Service-manager-visible definitions and state. |
| Configuration Drift (`configuration`) | Configuration resources and values. |
| Security Drift (`security`) | Security capabilities and configuration facts. |
| Capability Drift (`capability`) | Declared system or collector capability state. |
| Environment Drift (`environment`) | Environment and otherwise unowned metadata state. |
| Extension Drift (`extension`) | A valid but not-yet-registered canonical layer. |

Classification precedence is deterministic: exact canonical layer mapping
first; the `storage` layer then distinguishes filesystem/mount subtypes from
other storage by canonical object-kind, fact-name, and privacy-safe object ID;
an unknown layer becomes `extension`. No heuristic examines raw values. The
extension category prevents a new Inventory layer from being falsely assigned
an existing meaning and allows a future taxonomy minor release to register it.

## Lifecycle and determinism

1. Comparison validates snapshots and emits ordered canonical Change Records.
2. Drift validates each record's required identity, path, timestamp, values,
   change-type semantics, and uniqueness.
3. A fixed precedence table selects category, confidence, and rule.
4. Change type maps one-to-one to drift classification.
5. The engine derives the Drift ID only from canonical semantic fields and
   versioned classifier output.
6. Records are sorted by layer, object, path, and Change ID.
7. The complete output is validated before release.

There is no wall clock, randomness, locale, environment, map iteration, network,
AI, or mutable global state in this pipeline. Repeating classification over the
same records produces byte-identical JSON and does not mutate input.

## Compatibility strategy

Consumers must reject unsupported schema, engine, or taxonomy versions rather
than guess. Additive categories may be introduced in a taxonomy minor version
because category is an extensible string token; unknown canonical layers remain
truthfully classified as `extension` until registered. Optional envelope or
metadata additions may use a schema minor version when older consumers can
ignore them safely.

Removing or renaming fields or categories, changing field types, classification
meaning, identity inputs, ordering, confidence semantics, existing precedence,
or the one-change/one-drift invariant requires a new major version. Existing
Change Record 1.0 behavior is not modified by Drift evolution.

## Future consumers

Health consumes validated Drift Results and emits one versioned Health Record
per Drift Record under the contract defined in
`CANONICAL_HEALTH_ENGINE.md`. Rules may match public Drift and Health fields
without reaching back to snapshots. Policy may select or combine Rules and
Health outcomes without changing Drift classification. Reports and Automation
must preserve the originating Drift ID and Change ID. None of these consumers
may write back a category, bypass Comparison, or turn confidence into a risk
score.
