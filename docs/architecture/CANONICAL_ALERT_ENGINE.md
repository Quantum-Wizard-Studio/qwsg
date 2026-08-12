# Canonical Professional Alert Engine

## Purpose and permanent boundary

Task 028 establishes QWSG's canonical Alert decision boundary. The Alert
Engine determines whether a meaningful alert lifecycle event exists. It does
not deliver, persist, display, schedule, or act on that decision.

```text
Health / Rule / Policy / Scheduler Event / Canonical Report
                              + Effective Configuration identity
                              + previous Alert State
                              + explicit acknowledgement and suppression facts
                              + explicit evaluation time
                                           |
                                           v
                                  pure Alert Engine
                                           |
                                           v
                       Canonical Alert Records + proposed Alert State
                                           |
                                           v
                           future persistence and delivery adapters
```

`internal/alert` is a pure offline Go package. It reads no clock, filesystem,
environment, process, network, host, secret store, or mutable global state. It
has no persistence, recipient, channel, retry, daemon, presentation,
remediation, or AI dependency. Equivalent validated inputs, including the same
explicit evaluation time, produce byte-identical canonical JSON.

## Public contracts

All initial contracts use version `1.0`:

- `qwsg.alert-evaluation-input`: explicit immutable evaluation envelope;
- `qwsg.alert-evaluation-result`: ordered Alert Records and proposed state;
- `qwsg.alert-record`: immutable Canonical Alert Record 1.0;
- `qwsg.alert-state`: lifecycle continuity supplied and returned as data;
- `qwsg.alert-acknowledgement`: immutable operator-awareness fact;
- `qwsg.alert-suppression-window`: bounded operational or maintenance control.

The engine, taxonomy, and model versions evolve independently. Unsupported
versions, malformed identities, inconsistent time, invalid upstream contracts,
ambiguous sources, unordered controls, tampering, and resource-limit excess
fail closed.

Effective Configuration 1.0 is optional validated context. Its identity is
bound to Alert State when supplied. Alert does not add absent alert,
maintenance, acknowledgement, recipient, or notification fields to the
Configuration contract.

## Canonical Alert Record 1.0

Every record contains:

- stable condition key, lifecycle identity, generation, and immutable record
  identity;
- lifecycle event, decision, current lifecycle state, severity, prior severity,
  and category;
- exact source schema, version, record identity, subject, and bounded evidence
  references;
- distinct observation, event, and evaluation timestamps;
- acknowledgement and matching suppression identities where applicable;
- distinct expiration and recovery times;
- a localization-ready reason token and exact semantic versions.

Records exist only for meaningful lifecycle decisions. A changed source record
that preserves the same condition, severity, controls, and lifecycle produces
no repeated Alert Record, although proposed state retains the latest canonical
source reference.

## Identity and correlation model

Alert uses three identities:

1. **Condition key** — SHA-256 identity of the Alert category and exact
   privacy-bounded subject. It remains stable across evidence updates and
   recurrences.
2. **Lifecycle identity** — identity of the condition key, recurrence
   generation, and first observation time. Recovery or expiration ends that
   generation.
3. **Alert Record identity** — SHA-256 identity of the complete immutable
   lifecycle event record.

Correlation is exact. Health scope, Rule identity plus its Health scope,
Policy identity plus preserved evidence, Scheduler schedule scope, and Report
contract scope are canonical keys. There is no prose similarity, fuzzy match,
heuristic grouping, probabilistic inference, or AI reasoning.

Resolved and expired generations remain in bounded state so recurrence creates
a new lifecycle identity without losing the stable condition relationship.

## Source ownership and deduplication

Every supplied upstream envelope is validated by its owning package before use.
Alert-owned evidence-reference arrays remain bounded to 64 entries. A Policy
Report candidate references the single canonical Policy Report identity rather
than copying every Report source identity. The validated Policy Report remains
the canonical envelope that preserves complete, ordered source traceability;
therefore valid reports with hundreds or more sources do not expand an Alert
record or fail Alert evaluation.

Source precedence prevents competing alerts for one evidence chain:

1. a Policy Evaluation supersedes its referenced direct Rule Evaluation;
2. a Rule or Policy Evaluation supersedes its referenced direct Health Record;
3. a Canonical Report never recreates item-level Policy, Rule, or Health
   candidates and may create only a report-level incompleteness condition;
4. Scheduler Events create only scheduler-operational conditions and do not
   re-evaluate scheduling or Command/Pipeline results.

Conflicting sources for one condition key fail as ambiguous rather than using
input order as a tie-breaker. Effective Configuration supplies validated
identity and context only; it is not condition evidence.

## Deterministic source mappings

Alert severity is independent from Health status and Policy outcome.

| Canonical source fact | Alert category | Alert severity / result |
| --- | --- | --- |
| Health `informational` | `engineering_condition` | `informational` |
| Health `advisory` or `warning` | `engineering_condition` | `warning` |
| Health `critical` | `engineering_condition` | `critical` |
| Health unsupported or missing envelope evidence | `evidence_loss` | warning, indeterminate |
| Rule `matched` | `rule_match` | `warning` |
| Rule insufficient or unsupported | `evidence_loss` | warning, indeterminate |
| Rule invalid or evaluation error | `evidence_loss` | critical, indeterminate |
| Policy `escalated` | `policy_governance` | `emergency` |
| Policy `conflict` | `policy_governance` | critical, indeterminate |
| Policy `indeterminate` | `policy_governance` | warning, indeterminate |
| Scheduler state failure or clock discontinuity | `scheduler_operation` | `critical` |
| Scheduler lock contention | `scheduler_operation` | `warning` |
| incomplete Canonical Report | `report_completeness` | warning, indeterminate |

Health `healthy`, Rule `not_matched` or `disabled_rule`, and complete Reports
are explicit resolution evidence for their corresponding condition. Policy
`accepted`, `observe`, `suppressed`, and `not_applicable` resolve the Policy
governance alert condition only. Policy `suppressed` remains a governance fact;
it does not create or impersonate an operational suppression window.

Scheduler initialized, restart-recovered, and execution-completed events are
resolution evidence for the corresponding scheduler-operational condition.

## Lifecycle model

Lifecycle states are:

- `candidate`: internal validated candidate before the event decision;
- `active`: an alert condition currently exists;
- `acknowledged`: the condition exists and operator awareness is recorded;
- `suppressed`: the condition exists but a matching bounded control suppresses
  the decision from delivery eligibility;
- `indeterminate`: evidence cannot justify a determinate condition result;
- `expired`: evidence lifetime ended without canonical recovery evidence;
- `resolved`: canonical upstream evidence proves the condition ended.

Meaningful events are entry, escalation, de-escalation, acknowledgement,
suppression start/end, maintenance end, reminder, expiration, recovery, and an
indeterminate transition. Ordinary same-severity evidence changes are retained
without repeated alerts.

Severity order is `informational < warning < critical < emergency`.
Escalation and de-escalation compare this fixed order. Recovery is explicit
upstream resolution. Missing input never fabricates recovery.

After recovery or expiration, later qualifying evidence creates a new
generation. Expiration ends actionability because evidence aged past the
explicit TTL; it is not proof of health and carries no recovery timestamp.

## Acknowledgement

An acknowledgement is content-derived immutable evidence containing the exact
Alert lifecycle, actor, authority reference, UTC time, and optional bounded
localization-ready note token. It is accepted only for an existing nonterminal
lifecycle, cannot predate the condition, cannot be future-dated, and is unique
per lifecycle in an evaluation.

Acknowledgement changes no source condition, Alert severity, escalation,
de-escalation, suppression, expiration, or recovery. It grants no remediation
or delivery authority. Model 1.0 suppresses unchanged emergency reminders
while acknowledged but does not suppress escalation or recovery records.

## Suppression and maintenance

Suppression windows are explicit immutable inputs with content identities,
kind, exact bounded scope, start and end, actor, authority reference, reason
token, severity applicability, and emergency flag. Empty scope lists are
wildcards. Open-ended or longer-than-366-day windows are invalid.

Condition evaluation continues during suppression. A matching condition is
retained as `suppressed`, evidence is preserved, and the Alert Record identifies
every matching window. Suppression is never recovery or deletion.

`maintenance` is a suppression kind only. The Alert Engine does not create,
discover, activate, persist, cancel, schedule, or authorize maintenance. When
a supplied maintenance window ends and the condition remains, exactly one
maintenance-end current-status Alert Record is produced; suppressed historical
events are not replayed.

Emergency suppression requires an explicit `suppress_emergency` fact. A
wildcard suppression lacking it does not match emergency severity.

## Time, reminders, and expiration

All time is explicit UTC data. Observation time comes from the canonical
source. Event and evaluation time come from the supplied evaluation time.
Acknowledgement, suppression boundaries, expiration, and recovery retain their
own normative fields.

Evidence TTL is an explicit positive input bounded to 30 days. A condition
without resolution expires when evaluation reaches `last_observed + TTL`.
Missing evidence remains expired or indeterminate, never healthy.

Alert Model 1.0 defines one fixed emergency reminder interval of 24 hours.
Reminder eligibility is decided from retained `last_alerted_at` and explicit
evaluation time. Only an active, unacknowledged, unsuppressed emergency can
emit the reminder record. No timer, scheduler, delivery, or retry behavior is
implemented.

## State, ordering, and resource bounds

The engine returns proposed Alert State; callers remain responsible for future
authorized persistence. State is versioned, content-identified, validated, and
bounded to 4096 lifecycle conditions. One evaluation accepts at most 4096
condition candidates, 1024 acknowledgements, 1024 suppression windows, and 64
evidence references per source.

Conditions are ordered by condition key and generation. Alert Records are
ordered by condition key, generation, and record identity. Controls and source
references must be canonical sorted unique collections. Input slices are
copied before state construction.

Canonical JSON uses strict unknown-field rejection when decoded. Validation
recomputes State, Record, Result, acknowledgement, and suppression identities.

## Privacy, compatibility, and failure isolation

Alert copies only privacy-bounded source identities, scope-derived subject
identities, fixed taxonomy values, reason tokens, and evidence references. It
does not copy raw Inventory values, Report prose, arbitrary metadata, actor
notes, configuration values, credentials, secret material, host paths, or
delivery destinations.

Unknown contracts and taxonomies fail rather than being guessed. New source
types, severity/category meanings, lifecycle events, identity inputs,
precedence, mappings, reminder semantics, bounds, or time rules require an
explicit compatibility review and version change.

Alert failure does not alter upstream records or existing manual Command and
Pipeline operation. No Alert stage is added to Command Definition 1.0 or the
canonical Pipeline.

## Explicit exclusions

This architecture implements no alert persistence, incident database, daemon,
monitoring loop, timer, CLI command, Dashboard, Terminal UI, REST API, email,
SMS, Discord, Telegram, Slack, webhook, push, desktop notification, recipient
routing, rendering, delivery retry, delivery audit, channel health, network
operation, remote coordination, remediation, repair, shell/process/plugin
execution, host mutation, telemetry, licensing, AI, or machine learning.

The Alert Engine decides. Task 029's Canonical Notification Delivery consumes
only validated immutable Alert Records through the standalone `ValidateRecord`
boundary. It may plan and record delivery but cannot re-evaluate or change Alert
semantics. Persistence, concrete transports, daemon hosting, and operational
lifecycle remain separately authorized work.

Runtime is an Alert caller only: it supplies Scheduler evaluation once and then
exact typed Pipeline results in Request order while threading Alert's proposed
state. Runtime never authors, deduplicates, or reclassifies Alert Records.
