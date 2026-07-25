# Canonical Health Engine

## Permanent architectural boundary

The Health Engine is QWSG's deterministic engineering-condition evaluation
layer:

```text
Inventory -> Snapshot Store -> Comparison Engine -> Drift Engine -> Health Engine
                  what existed       what changed     what kind       condition
                                                                    |
                                   +--------------------------------+----------+
                                   |                   |                       |
                           future Rule Engine  future Policy Engine  future Report Engine
```

Inventory describes the system. Snapshots preserve observations. Comparison
produces factual Change Records. Drift classifies the kind of change. Health
evaluates the engineering condition represented by validated Drift evidence.
Health never collects, snapshots, compares, or reclassifies evidence.

`internal/health` is a pure, offline, AI-independent Go package. It accepts only
a validated `drift.Result` 1.0 envelope and returns only a validated Health
Result 1.0 envelope. It has no wall clock, randomness, locale, environment,
filesystem, process, scheduler, daemon, network, monitoring, alerting,
notification, policy, compliance, reporting, remediation, or AI dependency.

## Health evaluation model

Evaluation is a total deterministic transformation over supported canonical
Drift Records:

1. validate the complete Drift 1.0 contract;
2. preserve each Drift ID, Change ID, category, privacy-safe scope, and evidence
   timestamp;
3. apply the fixed Health Taxonomy 1.0 matrix;
4. derive a Health ID from versioned semantic fields;
5. sort Health Records by layer, object ID, path, and Drift ID;
6. aggregate status and evidence state using fixed precedence;
7. validate the complete Health result before release.

Every accepted Drift Record produces exactly one Health Record. Health does not
read previous or current comparison values and does not copy arbitrary Drift
metadata. The only record metadata is the bounded source classification token
needed to validate derivation. Equivalent canonical input produces
byte-identical canonical JSON and does not mutate the input.

Health evaluates the supplied evidence only. It does not claim that a one-shot
record is fresh, continuously observed, policy-compliant, safe, or remediated.

## Canonical Health Record 1.0

The public schema, engine, and taxonomy begin at version `1.0`.

| Field | Contract |
| --- | --- |
| `id` | Stable SHA-256-derived Health ID. |
| `drift_id` | Canonical Drift Record that is the sole evaluation input. |
| `change_id` | Originating canonical Change Record reference. |
| `category` | Preserved Drift category; Health never recategorizes drift. |
| `status` | Health Taxonomy 1.0 status. |
| `evidence_state` | `sufficient`, `insufficient`, or `unsupported`. |
| `confidence_basis_points` | Deterministic evidence sufficiency, never probability or risk. |
| `reason` | Stable machine-readable evaluation rule. |
| `scope` | Preserved privacy-safe layer, object ID, and path. |
| `evidence_timestamp` | Source comparison timestamp normalized to UTC. |
| `metadata` | Bounded source classification token. |
| `versions` | Health schema, engine, taxonomy, Drift schema, and Drift taxonomy versions. |

The envelope exposes schema, engine, and taxonomy versions, `overall_status`,
aggregate `evidence_state`, ordered records, and the exact input-contract and
pipeline identifiers. Canonical serialization validates before encoding.

Confidence is not probabilistic confidence. Supported fixed rules use `10000`
because their required evidence is sufficient for that rule. Unsupported
categories use `0`. An empty envelope has no fabricated record and explicitly
reports `unknown` with `insufficient` evidence.

## Health taxonomy 1.0

### Status model

| Status | Normative meaning |
| --- | --- |
| `healthy` | Canonical evidence states that the represented value is unchanged. |
| `informational` | Canonical evidence states that a value or presence was added. |
| `advisory` | Canonical evidence states that a supported value was modified. |
| `warning` | Canonical evidence states that a supported presence was removed. |
| `critical` | Canonical security evidence states that a security presence was removed. |
| `unknown` | No canonical Drift Record exists from which to evaluate condition. |
| `unsupported` | Valid Drift evidence uses a category not evaluated by Health 1.0. |

These statuses describe deterministic engineering condition. They are not alert
severity, business risk, compliance, policy, priority, likelihood, diagnosis,
or remediation instructions.

### Evidence model

`status` and `evidence_state` are separate:

- `healthy` with `sufficient` evidence is not the same as absent evidence;
- `unknown` is paired with `insufficient` evidence and cannot be treated as
  healthy;
- `unsupported` is paired with `unsupported` evidence and cannot be treated as
  unknown or healthy;
- sufficient supported findings never hide an unsupported finding in the
  aggregate evidence state.

### Evaluation matrix

Health first preserves the Drift category, then applies:

| Drift evidence | Health status | Evidence | Reason |
| --- | --- | --- | --- |
| category `extension` | `unsupported` | `unsupported` | `unsupported_drift_category` |
| `state_unchanged` | `healthy` | `sufficient` | `canonical_state_unchanged` |
| `presence_added` | `informational` | `sufficient` | `canonical_presence_added` |
| `value_modified` | `advisory` | `sufficient` | `canonical_value_modified` |
| `presence_removed` in `security` | `critical` | `sufficient` | `canonical_security_presence_removed` |
| other `presence_removed` | `warning` | `sufficient` | `canonical_presence_removed` |

The security-removal rule is a structural Health taxonomy rule, not a policy
profile. Changing it requires versioned taxonomy evolution.

The aggregate status precedence is `critical`, `unknown`, `unsupported`,
`warning`, `advisory`, `informational`, `healthy`. Aggregate evidence
precedence is `insufficient`, `unsupported`, `sufficient`. Empty input is the
only Health 1.0 path to `unknown`; records never manufacture evidence.

Health categories use the canonical Drift taxonomy directly: identity,
software, hardware, platform, filesystem, storage, network, service,
configuration, security, capability, environment, and extension. A future
upstream category such as a dedicated memory category remains `extension` and
therefore `unsupported` until Drift and Health contracts explicitly add it.
Health must not infer a category from paths or values.

## Drift to Health pipeline

The only supported pipeline is:

```text
comparison.ChangeRecord
        |
        v
drift.Classify -> validated drift.Result 1.0
        |
        v
health.Evaluate -> validated health.Result 1.0
```

Health rejects an unsupported envelope version, invalid ID, duplicate Drift or
Change identity, invalid derivation, unordered evidence, or other failed Drift
validation. It does not repair, reinterpret, or partially accept invalid input.
The one-Change/one-Drift/one-Health provenance chain remains explicit.

## Future Rule Engine integration

A future Rule Engine may match versioned Health fields, combine multiple Health
Records, and produce separately versioned rule outcomes. It must not change a
Health Record, infer hidden Health status, compare snapshots, or reclassify
drift. Rule definitions, profiles, expression languages, and execution are not
implemented by Health.

## Future Policy Engine integration

A future Policy Engine may select rule sets, define organization-specific
acceptance, and resolve policy conflicts over immutable Health and Rule
outcomes. Policy must remain distinguishable from the universal Health
taxonomy. No Policy or Compliance Engine is implemented here.

## Future Report Engine integration

A future Report Engine may render localized terminal, file, Web, notification,
or export views from canonical Health Records. Presentation must preserve
Health ID, Drift ID, Change ID, status, evidence state, reason, and versions.
Health performs no rendering, localization, notification, or remote delivery.

## Compatibility strategy

Consumers must reject unsupported Health schema, engine, taxonomy, Drift
schema, or Drift taxonomy versions rather than guess.

Additive optional envelope metadata may use a schema minor version only when
older consumers can safely ignore it. New statuses, evidence states, reason
semantics, categories, or precedence rules require explicit taxonomy
compatibility analysis. Removing or renaming fields, changing field types,
identity inputs, ordering, aggregation, confidence meaning, fixed evaluation
rules, or one-Drift/one-Health cardinality requires a major version.

New Drift categories remain truthfully `unsupported` until a Health taxonomy
version defines their evaluation. This conservative boundary permits upstream
evolution without silently assigning false Health meaning.

## Explicit exclusions

The Health Engine implements no monitoring, scheduler, daemon, Alert Engine,
email notification, automatic remediation, script execution, Rule Engine,
Policy Engine, Compliance Engine, Report Engine, network or remote
communication, AI integration, or operational automation. Those capabilities
require separate contracts, threat models, lifecycle authority, and tasks.
