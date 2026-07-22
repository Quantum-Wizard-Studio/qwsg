# Inventory Architecture

## Purpose and authority

This document defines the canonical Inventory Architecture for Quantum Wizard Server Guardian (QWSG). Inventory is a versioned digital description of an observed system: a digital twin of evidence-based reality, not a simulation, monitor, log stream, metric series, diagnosis, or health verdict.

This specification is the single source of truth for platform inventory concepts. Collectors produce it; policy, reporting, REST, user-interface, automation, and AI components consume it. Those consumers MUST NOT query the host to supplement inventory, create competing host models, or redefine canonical CPU, memory, storage, network, service, or other inventory objects.

This document is subordinate to the Project Constitution and approved product requirements. The narrower Core Alpha Slice 1 model remains a valid `1.x` implementation profile only where it conforms to this specification. If an older document conflicts, this document governs the platform model; a migration MUST be explicit rather than silently reinterpreting stored data.

Normative terms `MUST`, `MUST NOT`, `SHOULD`, and `MAY` express requirement strength.

## Core principles

- Inventory describes what was observed, how, when, and with what confidence; it does not declare that state healthy or desired.
- Unknown, unavailable, unsupported, denied, timed out, failed, cancelled, stale, and redacted are explicit states. None means healthy, absent, zero, or false.
- Each fact carries provenance, observation time, quality, and sensitivity sufficient for a consumer to assess its usability.
- The model is platform-neutral. Platform adapters change evidence acquisition, not canonical object meaning.
- A complete inventory is assembled from independently valid contributions. Partial success is first class and never presented as complete.
- Collection is read-only, stateless by default, bounded, measurable, interruptible, timeout-aware, and resource-efficient.
- Internal identifiers, enum tokens, schema keys, and machine-readable errors are locale-independent English.
- Data acquisition, normalization, evaluation, presentation, and authorized mutation are separate responsibilities.

## Architectural context

```text
Observed system
      |
      v
bounded collectors --> normalization/redaction --> schema validation
                                                   |
                                                   v
                                            InventorySnapshot
                                              /    |    \
                                             v     v     v
                                         policy  reports  API/UI/AI
                                             |
                                             v
                                  separately authorized action
```

The preferred execution lifecycle is:

```text
systemd timer -> oneshot service -> collect -> validate -> evaluate -> report -> exit
```

A daemon MAY be introduced by a separately approved design, but the Core and Inventory contracts MUST NOT depend on one. Scheduling is outside collector responsibility.

## Canonical object model

### Aggregate hierarchy

```text
InventorySnapshot
├── producer and schema identity
├── subject identity and snapshot timing
├── collection summary, issues, and redactions
└── layers[]
    └── InventoryLayer
        └── resources[]
            └── InventoryResource
                ├── identity, kind, lifecycle, labels
                ├── facts{}
                │   └── InventoryFact
                └── relationships[] -> InventoryResource identity
```

Collector execution evidence is represented separately by `CollectorResult` objects referenced by the snapshot and by contributed layers/resources. This prevents operational collector health from becoming a property of the observed host.

### InventorySnapshot

Every serialized inventory MUST contain:

| Field | Meaning |
| --- | --- |
| `schema_name` | Constant canonical identifier, `qwsg.inventory`. |
| `schema_version` | Semantic data-contract version in `major.minor` form. |
| `profile` | Optional named compatible subset, such as `core-alpha-slice-1`. |
| `snapshot_id` | Unique opaque identifier; never a secret or raw hardware ID. |
| `request_id` | Correlation identifier for the collection request. |
| `subject_id` | Privacy-safe stable identity of the observed system. |
| `observed_at` | UTC timestamp at the start of the observation window. |
| `completed_at` | UTC timestamp when assembly and validation completed. |
| `fresh_until` | Optional bounded validity deadline; never a health assertion. |
| `duration_ms` | Monotonic elapsed collection duration in integer milliseconds. |
| `status` | `complete`, `partial`, `failed`, or `cancelled`. |
| `producer` | QWSG component identity, version, and supported contract range. |
| `collector_results` | Deterministically ordered collector execution results. |
| `layers` | Deterministically ordered canonical inventory layers. |
| `issues` | Snapshot-level structured issues only. |
| `redactions` | Structured counts/reasons; never removed values. |
| `metadata` | Bounded extension metadata under registered namespaces. |

`observed_at <= completed_at`; when present, `completed_at <= fresh_until`. A `complete` snapshot requires all requested required capabilities to have usable results. A `partial` snapshot contains at least one usable requested contribution and at least one non-usable requested contribution. `failed` means no requested contribution is usable or the envelope is invalid. Cancellation is explicit and completed contributions MAY remain usable.

### InventoryLayer

An `InventoryLayer` groups canonical resources by architectural concern. It contains `layer_id`, `contract_version`, `status`, timing, contributing collector IDs, resources, issues, redactions, and bounded metadata. Layer status uses `available`, `partial`, `unavailable`, `unsupported`, `permission_denied`, `timeout`, `error`, or `cancelled`.

The minimum registered layers are:

| Layer ID | Canonical responsibility | Illustrative resource kinds |
| --- | --- | --- |
| `host` | Subject identity and physical/virtual context | `host`, `virtualization_context` |
| `hardware` | Physical or visible compute devices and capacity | `cpu`, `memory`, `device` |
| `operating_system` | OS, kernel, distribution, boot/runtime environment | `operating_system`, `kernel` |
| `runtime` | Language, container, and execution runtimes | `runtime`, `container_runtime` |
| `storage` | Filesystems, mounts, block/storage abstractions | `filesystem`, `mount`, `volume` |
| `network` | Interfaces, addresses, routes, and bounded local topology | `interface`, `address`, `route` |
| `services` | Service-manager-visible service definitions and state | `service`, `socket_unit`, `timer_unit` |
| `applications` | Allowlisted installed or deployed applications/components | `application`, `component` |
| `users` | Privacy-controlled local principals and groups | `user`, `group` |
| `security` | Observed security capabilities and configuration facts | `security_control`, `certificate_ref` |
| `policies` | Policy assignments and evaluation references, not verdict invention | `policy_assignment`, `policy_result_ref` |
| `metadata` | Inventory-wide annotations not owned by another layer | `annotation`, `tag_set` |

Layers MAY be absent when not requested or inapplicable. A required but unobtainable layer MUST be present with a truthful non-available status. New layers use the extension rules below and MUST NOT duplicate an existing layer's responsibility.

### InventoryResource

Every resource contains:

- `resource_id`: stable, opaque identity within the subject and resource kind;
- `kind`: registered locale-independent resource-kind token;
- `layer_id`: owning canonical layer;
- optional localizable `display_name_key` plus safe interpolation parameters;
- `lifecycle_state`: observed existence/state token where meaningful, not a health verdict;
- `facts`: a map of registered fact names to `InventoryFact` values;
- `relationships`: typed references to other resources;
- `labels`: bounded non-secret machine labels under registered namespaces;
- `observed_at`, `collector_id`, and `metadata`.

Identity MUST survive harmless presentation changes but MUST NOT expose raw serial numbers, MAC addresses, credentials, or other sensitive source identifiers. If a stable privacy-safe identity cannot be produced, the resource MUST declare an observation-scoped identity and consumers MUST NOT correlate it across snapshots.

### InventoryFact

Every fact contains:

| Field | Rule |
| --- | --- |
| `value` | Typed JSON value; absent for `unknown` or `redacted`. |
| `value_type` | Registered type such as `string`, `integer`, `number`, `boolean`, `timestamp`, `duration`, `bytes`, `percentage`, `enum`, `object`, or `array`. |
| `unit` | Canonical unit when applicable; bytes are integer bytes, durations declare their unit. |
| `quality` | `observed`, `reported`, `derived`, `estimated`, `unknown`, or `redacted`. |
| `sensitivity` | `public`, `operational`, `host_identifying`, `network_sensitive`, `personal`, or `secret_prohibited`. |
| `observed_at` | UTC timestamp for the fact's evidence. |
| `provenance` | Collector, logical source type/label, and transformation identifier. |
| `reason_code` | Required for unknown, redacted, or estimated values. |

`null`, zero, an empty string, and `false` are real values and MUST NOT encode unknown. A `secret_prohibited` fact MUST be rejected before inventory assembly; it MUST NOT be retained merely with a sensitivity label.

### Relationships

A relationship is a directed typed edge with `relationship_type`, `source_resource_id`, `target_resource_id`, optional observed facts, provenance, and timing. Initial canonical types are:

- `contains` / `contained_by`;
- `hosts` / `hosted_on`;
- `runs` / `runs_on`;
- `uses` / `used_by`;
- `depends_on` / `required_by`;
- `mounted_at` / `mount_of`;
- `connected_to`;
- `member_of`;
- `secured_by`;
- `governed_by`.

Edges MUST reference resources in the same snapshot or use an explicitly typed external reference. Relationships MUST NOT imply health, causality, ownership authority, or authorization unless the registered relationship semantics explicitly define it. Adding an inverse edge is optional; if emitted, both directions MUST be consistent.

## Collector contract

### Descriptor

Every collector MUST expose a machine-readable descriptor before execution:

| Required field | Contract |
| --- | --- |
| `collector_name` | Stable registered collector ID; not localized prose. |
| `version` | Collector implementation version. |
| `contract_version` | Collector contract version implemented. |
| `inventory_compatibility` | Supported Inventory schema range and produced layer/resource versions. |
| `capability` | Stable capability ID and precise evidence responsibility. |
| `supported_platforms` | Explicit predicates or declared adapters; never an untested broad claim. |
| `privilege_class` | Expected authority, ordinary user by default. |
| `timeout_ms` | Finite default and enforced maximum. |
| `output_limit_bytes` | Finite maximum raw and normalized contribution sizes. |
| `resource_budget` | CPU-time, memory, disk-I/O, process, and concurrency ceilings where measurable. |
| `sensitivity_classes` | Maximum classes the collector may encounter and emit. |

### Request

Collector input is immutable and contains a request ID, subject context, requested capability, absolute deadline, cancellation signal, privacy policy, bounded locale-independent options, and resource budget. It MUST NOT contain arbitrary executable names, shell fragments, uncontrolled filesystem roots, or presentation locale as a data-shaping instruction.

### CollectorResult

Every invocation returns exactly one structured result containing all minimum task fields:

| Required field | Meaning |
| --- | --- |
| `collector_name` | Descriptor identity. |
| `version` | Executed collector version. |
| `capability` | Capability attempted. |
| `supported_platforms` | Descriptor reference or resolved compatibility evidence. |
| `execution_time_ms` | Monotonic elapsed execution time. |
| `timestamp` | UTC observation timestamp. |
| `health_status` | Collector execution status, never host health. |
| `warnings` | Structured warning codes and localizable message keys. |
| `errors` | Structured error codes, class, retryability, safe details, and timestamps. |
| `collected_data` | Zero or more schema-valid layer/resource/fact contributions. |
| `metadata` | Bounded namespaced operational metadata. |

Collector `health_status` is one of `available`, `partial`, `unavailable`, `unsupported`, `permission_denied`, `timeout`, `error`, or `cancelled`. It describes the execution and usability of its contribution only.

Collectors MUST return facts, not user-facing prose or policy conclusions. They MUST NOT persist inventory, call other collectors, construct shell command strings, mutate the observed system, broaden privilege, perform undeclared network access, or bypass coordinator budgets. Raw evidence is transient by default. Platform-specific collectors and adapters MUST normalize into the same registered model.

### Measurability and interruption

The coordinator records wall and monotonic duration, outcome, bytes read/emitted, subprocess count, and timeout/cancellation behavior where the platform can measure them safely. Cancellation MUST propagate to pending and active work. Subprocesses MUST use fixed executables and bounded arguments, sanitized environments, closed standard input, finite output capture, and termination within the request deadline.

## Validation and assembly

The processing boundary is `acquire -> parse -> normalize -> classify sensitivity -> redact/reject -> validate contribution -> assemble -> validate snapshot`. Invalid contributions are never repaired by guessing. One failed collector does not invalidate independent valid contributions.

Validation MUST enforce:

- supported schema and contract versions;
- unique and bounded identifiers;
- registered layer, kind, fact, relationship, state, unit, and error tokens;
- deterministic ordering for arrays with canonical keys;
- valid timestamp ordering and non-negative resource measurements;
- referential integrity of relationships;
- declared compatibility between collector output and snapshot schema;
- absence of prohibited secrets and unredacted restricted values;
- agreement between contribution states and aggregate status.

## Serialization

### Canonical JSON

JSON is the mandatory canonical interchange representation. Canonical output MUST:

- be UTF-8 encoded and conform to RFC 8259;
- use the exact lowercase `snake_case` field names defined by the schema;
- encode timestamps as RFC 3339 UTC strings with `Z` and sufficient precision to preserve observations;
- encode integral byte counts and durations as JSON integers within the documented safe range;
- reject duplicate object keys, non-finite numbers, and invalid Unicode;
- omit only fields declared optional; never omit required status or provenance;
- order arrays by their registered stable keys and emit object members in schema order for deterministic QWSG output;
- include `schema_name` and `schema_version` in every top-level document.

Whitespace and object-member ordering are not semantic for consumers. When hashing, signing, or byte comparison is later required, a separately approved canonicalization profile MUST be declared; ordinary JSON serialization MUST NOT be assumed cryptographically canonical.

YAML, MessagePack, Protocol Buffers, and future encodings MAY be adapters over the same logical model. They MUST round-trip all supported types and explicit states without semantic loss and MUST NOT introduce an alternative object model. JSON remains the required interoperability baseline.

## Versioning and compatibility

Inventory schema versions use `major.minor`:

- increment `minor` for backward-compatible additive optional fields, enum values explicitly declared extensible, or new registered resource kinds/layers that old consumers may safely ignore;
- increment `major` for removal, rename, type change, changed meaning, newly required fields, identity changes, or altered status/relationship semantics;
- corrections that change no machine-observable contract MAY be published without a version change and MUST be documented.

Collectors explicitly declare an Inventory version range and the contract version of every contribution. Producers MUST NOT assemble outputs outside that intersection. Consumers:

- MUST reject unsupported major versions safely;
- MAY accept a newer minor only if they implement unknown-field and extensible-enum handling required by that major;
- MUST preserve the original validated document when forwarding unknown data, or disclose that the projection is lossy;
- MUST NOT overwrite a last-known valid snapshot with an incompatible or invalid document.

Migrations are explicit, directional, testable, non-destructive, and retain the source until verification. Downgrade is not assumed. A migration records source/target versions, tool version, time, result, and safe issues.

The Core Alpha Slice 1 envelope (`schema_version: 1.0`, categories and items) is a bounded legacy profile. Until migrated, adapters MAY project its categories into canonical layers and items into resources if every mapping is deterministic and lossless or disclosed as lossy. New platform modules MUST target this architecture, not extend the legacy category shape independently.

## Extension model

Extensions use reverse-domain or QWSG-owned namespaces for metadata, labels, facts, resource kinds, and capabilities. An extension MUST publish ownership, version, schema, sensitivity rules, resource budget, compatibility range, and collision-free identifiers. It MUST NOT:

- redefine or shadow canonical semantics;
- change required fields through metadata;
- treat unknown extensions as authorization to ignore validation;
- introduce user-facing hardcoded prose;
- bypass collection, privacy, privilege, or resource contracts.

A broadly useful extension SHOULD be promoted through an approved schema revision. Consumers ignore unknown optional extensions only when the enclosing contract marks them ignorable; otherwise they fail with an explicit compatibility error.

## Resource-efficiency contract

Collection MUST minimize CPU use, resident memory, disk activity, subprocess count, and process lifetime. Specifically:

- collection is one-shot and stateless by default;
- collectors operate under finite deadlines, read/output limits, concurrency limits, and memory budgets;
- streaming or bounded parsing is preferred to retaining raw evidence and duplicate object graphs;
- arbitrary recursive filesystem traversal and broad process/package enumeration are prohibited unless separately required and bounded;
- continuous polling is prohibited unless explicitly approved for a capability that cannot meet its requirements otherwise;
- temporary data is minimized and cleaned safely; durable writes occur only in an authorized state adapter;
- exceeding a budget stops the collector with `resource_limit` evidence and preserves unrelated results.

Every collector is measurable, interruptible, and timeout-aware. A module that cannot state and test reasonable bounds is not compatible with the Inventory Architecture.

## Security, privacy, and authority

Collection is non-root and read-only by default. Missing privilege produces truthful partial inventory, not automatic elevation. Future privileged collection requires a separate threat model, minimal capability boundary, owner approval, and the same normalized contract.

Collectors gather only purpose-required evidence. Credentials, secret values, private keys, tokens, payloads, arbitrary user content, full process command lines, and environment dumps are prohibited. Privacy and redaction happen before persistence or external presentation. Logs contain identifiers, safe error codes, counts, and durations—not raw collected values.

Inventory describes observed facts and policy references; it does not authorize repair. Any mutation consumes a validated inventory plus explicit policy and human/machine authority under a separately approved action contract.

## Consumer contracts

- **Policy Engine:** evaluates only validated Inventory objects. It MUST NOT query Linux or call collectors. Desired state and verdicts remain policy objects, separate from observed facts.
- **Reporting Engine:** renders only Inventory and evaluation results. It MUST NOT discover host state or convert unknown into healthy.
- **REST API:** exposes versioned Inventory objects or documented loss-aware projections. It MUST NOT invent alternative host structures.
- **Console:** consumes Inventory through an authorized Agent/API boundary and has no direct collector or shell access.
- **AI assistants:** receive bounded, redacted Inventory as system context. They MUST NOT depend directly on shell commands for canonical host facts and MUST distinguish observation from inference.
- **Automation and repair:** consume immutable snapshot identity and policy decisions; they revalidate freshness and authority before action and never mutate inventory as a substitute for recollection.

## Internationalization

English is canonical for engineering documents, schema keys, stable IDs, enums, error codes, and message keys. Human-visible text is rendered outside collectors from localizable message keys and typed parameters.

English and Hungarian are official user-documentation languages; the data and presentation contracts MUST permit unlimited future locales. No user-facing sentence, label, warning, error, date format, number format, or unit rendering may be hardcoded in a collector or canonical Inventory object. Missing translation falls back according to presentation policy without changing machine semantics.

## Documentation contract

An Inventory feature is incomplete until English and Hungarian user documentation exists for applicable topics and remains structurally translatable. At minimum the documentation set covers:

1. overview and terminology;
2. installation prerequisites and authority boundaries;
3. configuration, privacy, budgets, and collector selection;
4. usage and output interpretation;
5. architecture, schemas, relationships, and extension rules;
6. troubleshooting explicit partial/error states;
7. schema and collector upgrade/migration;
8. removal and retained-data handling;
9. developer notes, contract tests, and compatibility declarations.

Engineering architecture and developer records remain English. User documentation examples MUST use synthetic data and MUST NOT normalize secret exposure.

## Verification requirements

Compatible implementations require automated schema/contract tests, deterministic JSON golden tests, collector fixture tests for every result state, relationship integrity tests, compatibility and migration tests, cancellation/timeout/resource-limit tests, privacy/adversarial tests, partial-assembly tests, and consumer tests proving there is no direct discovery bypass.

Platform support claims require declared matrix evidence. Documentation review verifies both official languages and absence of hardcoded user-visible strings. Resource verification records peak memory where practical, duration, bytes read/emitted, subprocess use, deadline behavior, and bounded failure.

## Future-module compatibility

### Implemented profile

Task 014 implements profile `canonical-system-inventory-v1` for Linux through the existing Collector Registry. The authoritative representation is emitted as the additive `canonical_inventory` member of the Inventory 1.0 envelope so existing consumers retain their category/item projection. Host, operating-system, kernel, CPU, memory, storage, filesystem, network, and virtualization capabilities are implemented; mapping and operational details are defined in `docs/architecture/CANONICAL_SYSTEM_INVENTORY_V1.md`.

The Collector Framework, Platform Hardening, Configuration Engine, Policy Engine, Reporting Engine, REST API, and later Console or AI integrations MUST import or implement these contracts. Each task identifies the Inventory schema range it consumes or produces, adds no competing host model, and treats contract changes as governed schema evolution.

The enduring boundary is:

```text
Collectors collect. Inventory describes. Policies evaluate.
Reports explain. Authorized action changes reality. Recollection verifies it.
```
