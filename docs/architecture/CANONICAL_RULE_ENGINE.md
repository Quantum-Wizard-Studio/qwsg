# Canonical Rule Engine

## Permanent architectural boundary

The Rule Engine is QWSG's deterministic condition-matching layer:

```text
Inventory -> Snapshot -> Compare -> Drift -> Health -> Rule
                                                    |
                         +--------------------------+-------------------+
                         |                          |                   |
                future Policy Engine     future Report Engine  future Automation
```

Compare determines what changed. Drift determines what kind of change occurred.
Health determines the engineering condition represented by canonical evidence.
Rule determines whether a predefined deterministic condition matches a
canonical Health Record. Policy may later decide how a matched rule is
governed. Automation may act only after separate Policy and authorization.

`internal/rule` is a pure, offline, AI-independent Go package. It accepts only
Rule Definition 1.0 values and a canonical `health.Result` 1.0 envelope. It
produces only Canonical Rule Evaluation Records 1.0. It performs no Inventory
collection, snapshot persistence, comparison, Drift classification, Health
evaluation, policy, compliance, risk scoring, monitoring, scheduling, daemon
work, alerting, notification, reporting, dashboard work, remediation, process,
script or plugin execution, networking, cloud access, host mutation, machine
learning, probabilistic evaluation, or arbitrary code execution.

## Rule definition model

Rule Definition 1.0 is explicit data, never executable content.

| Field | Contract |
| --- | --- |
| `id` | Required stable caller-owned canonical token; unique in an evaluation. |
| `contract_version` | Exact supported Rule contract, initially `1.0`. |
| `category` | Rule Taxonomy 1.0 category. |
| `scope` | Sorted optional Health IDs and Health categories. |
| `enabled` | Disabled Rules produce an explicit skipped result. |
| `input_requirements` | Sorted fixed Health fields required by the definition. |
| `condition` | Bounded typed operator tree. |
| `description` | Human-readable engineering description; not executable. |
| `metadata` | Bounded annotation data; never interpreted as logic. |

Rule IDs use lowercase alphanumeric segments separated by `.`, `_`, `:`, or
`-`. They are stable because they are explicit contract identity, not generated
from mutable descriptions or metadata. Evaluation IDs are engine-derived from
canonical Rule identity, Health evidence identity, outcome, explanation, and
versioned evaluation semantics.

Scope filters already evaluated Health Records. It does not infer category,
status, or health from Inventory, Change, or Drift data. Empty scope arrays
mean every Health Record in the validated envelope.

## Rule taxonomy 1.0

| Category | Responsibility |
| --- | --- |
| `health_status` | Match canonical Health status. |
| `health_category` | Match the preserved canonical Health category. |
| `health_evidence` | Match explicit Health evidence state, confidence, reason, or scope. |
| `composite` | Combine bounded canonical conditions. |
| `extension` | Reserved unsupported category for forward compatibility. |

An `extension` Rule is validly identified but not evaluated by Rule 1.0. An
unknown category is invalid. This distinction prevents unrecognized semantics
from being treated as a normal non-match.

## Deterministic operator model

The only addressable input fields are:

```text
health_id
category
status
evidence_state
confidence_basis_points
reason
scope.layer
scope.object_id
scope.path
```

The canonical operators are:

| Operator | Semantics |
| --- | --- |
| `eq`, `ne` | Typed scalar equality and inequality. |
| `gt`, `gte`, `lt`, `lte` | Integer comparison, only for confidence basis points. |
| `in` | Typed membership in an explicit finite list. |
| `exists` | The registered canonical field exists. |
| `status_matches` | Explicit Health status match. |
| `category_matches` | Explicit Health category match. |
| `and`, `or` | Ordered composition of at least two child conditions. |
| `not` | Negation of exactly one child condition. |

There is no implicit type conversion, regular expression, arithmetic, variable,
function, loop, callback, file, environment, command, network, plugin, or
extension execution. Conditions are limited to depth 8 and 64 nodes; an
evaluation accepts no more than 1024 Rules. Unsupported operators are
`unsupported_rule`; malformed shapes, incompatible types, or exceeded bounds
are `invalid_rule`.

## Canonical Rule Evaluation Record 1.0

The public `qwsg.rule-evaluation` schema, engine, and taxonomy begin at `1.0`.

| Field | Contract |
| --- | --- |
| `id` | Stable SHA-256-derived Evaluation ID. |
| `rule_id` | Canonical evaluated Rule identity. |
| `health_record_id` | Canonical Health Record identity when evidence was evaluated. |
| `outcome` | Canonical evaluation outcome below. |
| `evaluation_status` | `complete`, `skipped`, or `failed`. |
| `match_result` | `match`, `no_match`, or `indeterminate`. |
| `confidence_basis_points` | `10000` for a completed deterministic evaluation, otherwise `0`; never probability or risk. |
| `explanation` | Stable machine-readable reason token. |
| `evidence_references` | Bounded canonical Health ID references only. |
| `metadata` | Bounded Rule category and root-operator tokens. |
| `versions` | Evaluation schema, engine, taxonomy, Rule contract, Health schema, and Health taxonomy versions. |

The result envelope identifies its exact Health input contract and pipeline.
Records are ordered by Rule ID, Health Record ID, and Evaluation ID. Canonical
serialization validates the entire result before producing byte-stable JSON.

## Canonical evaluation outcomes

| Outcome | Meaning |
| --- | --- |
| `matched` | A supported enabled Rule condition matched one Health Record. |
| `not_matched` | A supported enabled Rule condition deterministically did not match. |
| `insufficient_evidence` | No canonical Health Record exists in the Rule's scope. |
| `unsupported_rule` | Rule contract, category, or operator is recognized as unsupported. |
| `invalid_rule` | Rule identity, shape, type, ordering, bounds, or required data is invalid. |
| `evaluation_error` | The canonical Health input contract failed technical validation. |
| `disabled_rule` | The Rule is explicitly disabled and was skipped. |

Only `matched` and `not_matched` have a definite match result and a Health
evidence reference. Every other outcome is `indeterminate`, has confidence
zero, and cannot be reported as a normal non-match. A disabled Rule is not an
insufficient or failed Rule.

## Evaluation lifecycle

1. Bound and copy the Rule list.
2. Sort Rules by stable Rule ID and reject empty or duplicate identities.
3. Validate the complete Health contract.
4. For each Rule, classify unsupported, invalid, and disabled definitions
   before selecting evidence.
5. Select Health Records only through canonical scope fields.
6. Emit `insufficient_evidence` when the scope selects no Health Record.
7. Evaluate the fixed condition tree independently for each selected Health
   Record.
8. Derive stable Evaluation IDs and sort all records canonically.
9. Validate the complete output before release.

Equivalent Rule definitions and Health input produce byte-identical output and
do not mutate input. Evaluation contains no wall clock, randomness, locale,
environment, or mutable global state.

## Health to Rule pipeline

The only supported input boundary is:

```text
health.Evaluate(validated Drift evidence)
        |
        v
validated health.Result 1.0
        |
        v
rule.Evaluate(Rule Definition 1.0, health.Result 1.0)
        |
        v
validated Rule Evaluation Result 1.0
```

Rule preserves Health IDs as evidence. It does not read Drift or Change IDs,
comparison values, Inventory facts, or host state. Rule cannot change Health
status, evidence state, confidence, reason, scope, ordering, or identity.

## Explainability and evidence model

Every record exposes the Rule ID, evaluated Health ID when available, definite
or indeterminate match result, evaluation lifecycle status, stable explanation,
root operator, Rule category, evidence references, and all semantic versions.
No arbitrary source metadata or raw upstream values are copied. Confidence
states only whether a fixed deterministic condition completed; it is not
probability, risk, severity, priority, trust, or policy weight.

## Invalid, unsupported, and error handling

Unsupported versions and operators are not guessed. Invalid structures are not
partially evaluated. Missing scoped evidence is not an error. Invalid Health
input is a technical `evaluation_error`, never `not_matched`. Duplicate or empty
Rule IDs prevent a stable envelope and therefore return an API error rather
than ambiguous records.

## Privacy and security boundaries

Rule sees only already privacy-bounded Health fields. Evidence output contains
Health IDs and fixed reason tokens, not previous/current comparison values or
arbitrary Health metadata. Descriptions and Rule metadata are not copied into
evaluation records. There is no executable string, dynamic field path, host
access, privilege, persistence, side effect, or communication boundary.

## Versioning and compatibility strategy

Consumers must reject unsupported evaluation schema, engine, taxonomy, Rule
contract, Health schema, or Health taxonomy versions rather than guess.

Adding an optional field may use a schema minor version only when old consumers
can ignore it safely. Adding a field selector or operator requires explicit
operator-registry and compatibility review. New categories remain
`extension`/`unsupported_rule` until a Rule taxonomy version defines them.

Changing field types, operator meaning, result outcomes, explanation meaning,
identity inputs, ordering, bounds, scope semantics, evidence references,
confidence semantics, or match/error separation requires a major version.

## Future Policy Engine integration

Policy may select Rule sets, interpret matches, resolve conflicts, assign
governance meaning, and decide whether an action is authorized. Policy must
consume immutable Rule Evaluation Records and must not reimplement Rule
matching or modify its evidence.

## Future Report Engine integration

Report may render localized terminal, file, Web, export, or notification views.
It must preserve Evaluation ID, Rule ID, Health ID, outcome, status,
explanation, evidence, and versions. Rule performs no presentation or delivery.

## Future Automation integration

Automation may consume only separately authorized Policy outcomes derived from
canonical Rule Evaluation Records. A Rule match is never an instruction,
authorization, remediation, command, alert, or scheduled action.

Future Policy, Reporting, Alerting, Compliance, and Automation components must
consume canonical Rule Evaluation Records instead of implementing independent
rule-matching logic.
