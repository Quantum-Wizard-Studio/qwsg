# Canonical Policy Engine

## Status and purpose

Task 025 establishes the Canonical Policy Engine as QWSG's sole governance
interpretation boundary. It consumes immutable Canonical Rule Evaluation
Records and produces deterministic Policy Evaluation Record 1.0 values through
versioned Policy Profile 1.0 definitions.

Policy answers how a technical Rule outcome is treated in an applicable
operating profile. It never changes whether the Rule matched, modifies Health
or other evidence, or performs an operational action.

The permanent flow is:

```text
Inventory → Snapshot → Compare → Drift → Health → Rule → Policy → Report
```

Future Scheduler, Alert, Dashboard, Terminal UI, REST API, remote-management,
and Automation components consume Policy Evaluation. They may not reinterpret
Rule outcomes or create competing governance semantics.

## Package and dependency boundary

`internal/policy` is a pure offline Go package. It imports only the standard
library and `internal/rule`. It performs no collection, host reads, shell or
process execution, persistence, scheduling, notification, network access,
presentation, remediation, telemetry, or AI work.

Only `internal/pipeline` orchestrates Policy. Command definitions select the
canonical `policy` stage; CLI and presentation packages contain no profile,
precedence, conflict, or outcome logic.

## Contracts

### Policy Profile 1.0

A Profile contains:

- a canonical ID and content-derived identity;
- exact contract and independently evolvable profile versions;
- bounded priority and deterministic enabled state;
- sorted parent-profile references;
- a Rule ID and Rule Outcome scope;
- sorted bounded statements;
- an explicit default outcome;
- bounded metadata.

Every statement has a stable ID, priority, Rule selector, Policy outcome,
localization-ready explanation token, and bounded metadata. Human prose is not
part of evaluation semantics.

Profile identity is SHA-256 over the complete canonical profile excluding only
the identity field. An altered profile cannot retain its old identity.
Unsupported contract versions, malformed identities, unsorted or duplicate
selectors, unknown parents, cycles, excessive inheritance depth, invalid
outcomes, and resource-limit violations fail closed.

### Policy Evaluation Record 1.0

One record is emitted for every source Rule Evaluation Record. It contains:

- stable content-derived Policy Evaluation identity;
- exact Rule Evaluation and Rule IDs;
- canonical Policy outcome and evaluation status;
- applied profile and statement IDs;
- an explanation token;
- the immutable Rule Evaluation evidence reference;
- source outcome metadata and complete contract versions.

Records never embed or reinterpret Health, Drift, Change, Inventory, or private
fact values. Traceability follows Policy → Rule → Health → Drift → Change.

### Policy Result 1.0

`qwsg.policy-evaluation/1.0` contains deterministically ordered records, the
sorted evaluated Profile IDs, engine and taxonomy versions, and explicit input
contract metadata. Canonical JSON is byte-stable for identical validated
inputs.

## Policy taxonomy

| Outcome | Meaning |
|---|---|
| `accepted` | The applicable profile explicitly accepts the Rule outcome. |
| `observe` | The outcome remains noteworthy for operator interpretation. |
| `suppressed` | The profile explicitly suppresses operational attention; evidence remains unchanged and visible. |
| `escalated` | The profile assigns elevated governance importance; no alert or action occurs. |
| `indeterminate` | Available policy or evidence cannot justify a determinate treatment. |
| `not_applicable` | No enabled profile applies to the source Rule Evaluation. |
| `conflict` | Equal-precedence applicable policies demand different outcomes. |

`suppressed` is not deletion. `escalated` is not an alert. `accepted` is not a
claim that the underlying system is healthy. `conflict` is a successful,
explicit governance result rather than arbitrary tie-breaking.

Evaluation status is independent: `complete` means Policy evaluation completed;
`skipped` is used only when no Policy applies. Technical Rule status and outcome
remain immutable source fields.

## Selection, precedence, and conflicts

All supplied enabled profiles whose scopes match a Rule Evaluation participate.
Empty selector lists are explicit wildcards. Selection uses only canonical Rule
ID and Rule Outcome fields.

Precedence is lexicographic and transparent:

1. higher Profile priority;
2. higher statement priority within equal Profile priority;
3. all candidates at the winning precedence are retained.

Lower-precedence candidates cannot override the winner. Equal-precedence
candidates with the same outcome merge their traceability. Equal-precedence
candidates with different outcomes produce `conflict`. Profile ID or input
order never breaks a semantic tie.

When an applicable Profile has no matching statement, its explicit default
outcome participates below every valid statement in that Profile. Missing or
incomplete technical evidence therefore cannot silently become accepted unless
the Profile explicitly says so.

## Inheritance and composition

Profiles may extend sorted parent IDs. Parent statements are composed into the
child's evaluation context; child Profile priority remains the governing
Profile precedence. The graph must be complete, acyclic, and at most eight
levels deep. Evaluation deduplicates final traceability IDs and remains stable
regardless of input Profile order.

Inheritance is configuration reuse only. It cannot mutate a parent, hide a
conflict, execute an action, or change Rule evaluation.

## Report and command integration

Task 025 adds `policy` to the Command 1.0 canonical stage registry between
`rule` and `report`. Stored `report` workflows now execute:

```text
Compare → Drift → Health → Rule → Policy → Report
```

The default `qwsg.canonical.observation` Profile maps matched observation Rules
to `observe`, non-matches to `accepted`, and every other Rule outcome to the
explicit `indeterminate` default.

`internal/report.GeneratePolicy` creates the additive
`qwsg.policy-report/1.0` contract directly from validated Policy Evaluations.
The earlier Rule-backed `qwsg.report/1.0` API remains available and unchanged
for compatibility, but canonical pipeline Report execution uses Policy-backed
reports. Presentation renders completed execution data and never evaluates
Policy.

Task 027 Scheduler consumes Policy only indirectly through completed canonical
Command Executions. Its Execution Result may preserve Policy Evaluation identity
and outcome as immutable traceability evidence; scheduling never re-evaluates
Policy or treats an outcome as execution, alert, or remediation authority.

## Alert integration

The pure Canonical Alert Engine consumes validated Policy Evaluation Records as
immutable governance evidence. Alert owns the explicit versioned mapping from
Policy outcomes to Alert decisions. `escalated` maps to an emergency Alert
condition; `conflict` and `indeterminate` remain explicit indeterminate Alert
conditions. `accepted`, `observe`, `suppressed`, and `not_applicable` resolve
only the Policy-governance Alert condition.

Policy `suppressed` remains governance evidence and never becomes an
operational suppression window, evidence deletion, notification action, or
remediation authority. When Policy and its referenced Rule are supplied in the
same Alert evaluation, Policy is the downstream canonical source and prevents
a duplicate direct Rule alert.

## Determinism, safety, and limits

- Maximum Profiles: 64.
- Maximum statements per Profile: 256.
- Maximum Policy Evaluation Records: 4096.
- Maximum inheritance depth: 8.
- Profile and statement priorities are bounded to `-1000..1000`.
- Metadata maps are bounded to 16 fields per Profile or statement.
- IDs, selectors, Profiles, records, sources, and Report sections have canonical
  ordering.
- Exact contract versions fail closed; future semantics require versioned
  contracts or documented compatible extensions.
- Inputs are never mutated, private upstream values are not copied, and no
  hidden time, host, network, environment, or randomness source exists.

## Explicit non-goals

Task 025 implements no Scheduler, Alert Engine, monitoring, daemon, notification,
email, webhook, Dashboard, Terminal UI, REST API, remote execution, remediation,
Configuration UI, host mutation, AI, or machine learning. Policy records may be
consumed by those separately authorized future components, but never perform
their work.
