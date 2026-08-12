# Canonical Report Engine

## Permanent architectural boundary

The Report Engine is the presentation-contract layer of the canonical
engineering pipeline:

```text
Inventory -> Snapshot -> Compare -> Drift -> Health -> Rule -> Policy -> Report
                                                                          |
                          +----------------+--------------------+----------------+
                          |                |                    |                |
                 future Dashboard   future Export     future Notification  management UI
```

Compare determines what changed. Drift determines what kind of change
occurred. Health evaluates the engineering condition. Rule matches predefined
deterministic conditions. Report transforms the resulting canonical evidence
into a deterministic, presentation-neutral engineering artifact.

`internal/report` is a pure offline Go package. Legacy Report 1.0 accepts only a
validated `rule.Result` 1.0 and produces only a Canonical Report 1.0. It does
not collect Inventory, persist Snapshots, compare, classify Drift, evaluate
Health or Rules, execute Policy, calculate compliance or risk, monitor,
schedule, run a daemon, alert, notify, email, render HTML or PDF, build a
Dashboard, remediate, execute processes/scripts/plugins, communicate remotely,
use cloud services, mutate the host, use AI or machine learning, or perform
probabilistic evaluation. Task 025 adds `qwsg.policy-report/1.0`, which accepts
only validated Policy Evaluation Results and becomes the canonical pipeline's
Report contract. It preserves Policy and Rule source identities without
re-evaluating either. The Rule-backed API remains available unchanged for
compatibility.

## Canonical Report 1.0

The public schema is `qwsg.report` version `1.0`.

| Field | Contract |
| --- | --- |
| `id` | Stable SHA-256-derived Report identity. |
| `schema_name`, `schema_version` | Exact public Report contract. |
| `type` | Report taxonomy type. |
| `title_token` | Machine-readable localization token, not generated prose. |
| `completeness` | `complete` or `incomplete`; empty evidence is never silently complete. |
| `summary` | Exact counts for each preserved Rule outcome. |
| `sections` | Canonically ordered, outcome-specific structured sections. |
| `sources` | Ordered source identity and contract references. |
| `metadata` | Fixed pipeline and rendering-model tokens. |
| `versions` | Report, Rule Evaluation, Rule, and Health semantic versions. |

Every section has a stable ID, Rule outcome, title token, and non-empty ordered
item list. Every item has a stable ID and preserves the Rule ID, optional
Health Record ID, Rule outcome, evaluation status, match result, deterministic
confidence, explanation token, Health evidence references, and exact Rule
Evaluation source reference. Report never changes those meanings.

## Report taxonomy 1.0

The smallest complete initial taxonomy contains:

| Type | Meaning |
| --- | --- |
| `engineering_summary` | Structured presentation of canonical Rule Evaluation Records. |

The supported source taxonomy contains only `rule_evaluation`. Direct
Inventory, Snapshot, Compare, Drift, or Health adapters are unsupported in
Report 1.0 because they would bypass or duplicate established engineering
semantics. New report or source types require explicit versioned contracts and
compatibility review.

The additive `qwsg.policy-report/1.0` source taxonomy contains only
`policy_evaluation`. Its summary and sections preserve the fixed Policy outcome
order `escalated`, `conflict`, `observe`, `indeterminate`, `suppressed`,
`accepted`, and `not_applicable`. Every item retains the Policy Evaluation,
Rule Evaluation, Rule, applied Profile, applied statement, explanation, and
evidence references. Direct Rule input is invalid for this Policy-backed
contract, just as Policy input is invalid for legacy `qwsg.report/1.0`.

Sections follow this fixed outcome order:

```text
matched
not_matched
insufficient_evidence
unsupported_rule
invalid_rule
evaluation_error
disabled_rule
```

Outcomes are never merged. In particular, a technical evaluation error is not
a normal non-match, and missing evidence is not a match decision.

## Deterministic generation pipeline

1. Validate the entire Rule Evaluation Result 1.0.
2. Enforce the 4096-item resource bound.
3. Copy only the public, privacy-bounded presentation fields.
4. Create one Report item and one source reference for every evaluation.
5. Count each outcome without scoring, weighting, or interpretation.
6. Group items by the fixed taxonomy and sort them by Rule ID, Health Record
   ID, and stable item ID.
7. Sort source references by Evaluation ID.
8. Derive stable item, section, and Report IDs.
9. Validate identity, ordering, counts, completeness, source cardinality, and
   traceability before release.

Equivalent valid input produces byte-identical canonical JSON. The pipeline
uses no wall clock, randomness, locale, environment, filesystem, network, or
mutable global state and does not mutate its input.

## Rendering model

Canonical Report data is presentation-neutral. Machine-readable title and
explanation tokens allow future localized views without changing engineering
evidence. `RenderText` is the only initial view: a deterministic minimal text
projection derived solely from an already validated Report. It escapes control
characters, performs no template evaluation, and does not read upstream data.

Text rendering is not an Export Engine, notification system, Dashboard, HTML
renderer, PDF renderer, or delivery mechanism. Future renderers consume the
Canonical Report contract and cannot reconstruct or reinterpret Rule logic.

## Source traceability model

Every Report item has exactly one `rule_evaluation` source containing:

- the canonical Rule Evaluation ID;
- contract name `qwsg.rule-evaluation`;
- contract version `1.0`.

The Report-level source list is a sorted one-to-one index of those references.
Completed Rule matches retain their canonical Health Record ID and evidence
reference. Terminal outcomes retain the explicit absence of Health evidence.
Validation rejects missing, duplicate, altered, unsupported, or unreferenced
sources.

Traceability is therefore:

```text
Report ID
  -> Section ID
    -> Item ID
      -> Rule Evaluation ID
        -> Rule ID
        -> Health Record ID / evidence reference (when evaluated)
```

Report does not copy raw Inventory facts, previous/current Compare values,
arbitrary Rule metadata, or private upstream payloads.

## Complete pipeline relationship

Inventory is the acquisition authority and Snapshot is its durable canonical
state. Compare emits Change Records. Drift classifies those changes. Health
evaluates their engineering condition. Rule determines whether predefined
conditions match. Report presents only the canonical Rule Evaluation evidence.

Each layer consumes the previous layer's public versioned contract:

```text
Inventory 1.0
 -> Snapshot
 -> Change Record 1.0
 -> Drift Record 1.0
 -> Health Record 1.0
 -> Rule Evaluation Record 1.0
 -> Canonical Report 1.0
```

Future components must consume Canonical Reports rather than rebuilding
engineering summaries independently.

## Completeness, invalid, and unsupported behavior

A valid non-empty Rule Result produces a complete Report even when individual
Rule outcomes are insufficient, unsupported, invalid, failed, or disabled;
those states are faithfully presented. A valid empty source produces an
explicitly incomplete empty Report. An invalid or unsupported Rule source
contract produces an API error and no misleading Report. An unsupported Report
contract fails validation and cannot be serialized or rendered.

## Privacy and security

Report can only see privacy-bounded Rule Evaluation fields. It preserves
evidence identifiers but not raw source values. Metadata is a fixed two-token
map; arbitrary upstream metadata and descriptions are excluded. Rendering
escapes control characters. There is no dynamic expression, executable
content, arbitrary template, privilege, persistence, communication, side
effect, or host access.

## Versioning and compatibility strategy

Consumers must reject unsupported Report schema, engine, taxonomy, Rule
Evaluation, Rule contract, Health schema, or taxonomy versions rather than
guess.

Additive optional fields may use a minor schema version only when old consumers
can safely ignore them. A new report type or source adapter requires explicit
taxonomy support and tests. Changes to field meaning, completeness, taxonomy,
traceability cardinality, outcome order, identity inputs, bounds, canonical
serialization, or rendering safety require a major version. Version-specific
adapters may normalize older canonical contracts; they must never infer new
upstream semantics.

## Future Dashboard integration

A Dashboard may localize, filter, and visualize validated Canonical Reports.
It must preserve Report, section, item, source, Rule, and Health identities and
must not evaluate Rules or rebuild engineering summaries.

## Future Export integration

An Export Engine may transform Canonical Reports into HTML, PDF, CSV, or other
formats under a separate contract. Export formatting and delivery remain
outside Report 1.0 and cannot alter canonical content.

## Policy integration

Policy remains a separate governance layer. It consumes canonical Rule
Evaluation Records, not presentation text. Task 025 implements the separately
versioned Policy-backed Report adapter over validated Policy Evaluation
Records. Report never executes Policy or treats a match as authorization.

Dashboard, Export, Notification, and management interfaces shall consume
Canonical Reports instead of implementing independent engineering summary or
Policy interpretation logic.

Scheduler Execution Results may preserve Report stage contract identity and
Policy trace references from a completed Command Execution. Scheduler neither
generates Reports nor reinterprets their content.

## Alert integration

The pure Alert Engine may consume a validated Rule-backed or Policy-backed
Canonical Report only for report-level completeness and traceability. It never
recreates Alert candidates from Report items because those items preserve
upstream Rule or Policy records that own the canonical evidence chain. A
complete Report resolves its report-completeness condition; an incomplete
Report produces one indeterminate report-level condition. Report generation,
rendering, and delivery remain separate.

## Operator presentation consumer

The Canonical Operator Presentation Model may consume validated Report and
Policy Report completeness and identities as explicit observations. It does
not rebuild sections, explanations, Rule outcomes, or Policy decisions. Report
remains the engineering-report authority; the operator model owns only the
cross-domain overview needed by replaceable interfaces.
