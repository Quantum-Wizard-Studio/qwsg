# Canonical Operator Presentation Model

Inventory is now an explicit typed observation for limited current coverage. Stored Overviews are aged only by `RequalifyFreshness`; storage and interfaces cannot alter freshness, completeness, condition, or recommendations.

The `observe` application adapter supplies one validated, correlated Command, Inventory, Comparison, Drift, Health, Rule, Policy, and Policy Report input. The model applies its existing rules unchanged; the adapter neither derives condition nor invents Alert, Runtime, or Guardian evidence.

The Interactive Operator Console is its first interface consumer. It receives only validated Overview values; all condition, attention, Guardian, freshness, completeness, Alert, change, and recommendation meaning remains owned here.

## Status and purpose

Operator Presentation Model 1.0 is the permanent shared read model between
canonical QWSG records and replaceable user interfaces. Its implementation is
`internal/presentationmodel`.

```text
Canonical Engineering and Operational Data
                    |
                    v
Canonical Operator Presentation Model
                    |
                    v
Replaceable Interface
```

The model answers the first operator questions—condition, attention, change,
Alerts, Guardian state, evidence freshness/completeness, and safe next step—
without exposing internal engine vocabulary in the beginner summary. Stable
source references remain available for advanced drill-down and machine use.

## Ownership boundary

The projection consumes only validated existing contracts and never changes
their meaning. Comparison owns factual change, Drift owns classification,
Health owns engineering condition, Rule owns matching, Policy owns governance,
Report owns engineering reports, Alert owns Alert lifecycle, Runtime owns one
cycle, and Runtime Service owns process lifecycle. The presentation model owns
only their bounded cross-domain operator projection.

`command.Execution.View` remains Command stage metadata. Canonical Report and
Policy Report remain engineering-report contracts. Neither is expanded into a
cross-domain current-status model, and the operator model does not rebuild
their logic.

Canonical engines, Command, Report, Runtime, and Runtime Service do not import
the presentation model. Future terminal, REST, Web Dashboard, and other
adapters consume the same validated Overview instead of interpreting evidence
independently.

## Input and observation contract

Projection Input 1.0 contains one explicit UTC projection time, a positive
bounded freshness duration, and optional timestamped observations of Command,
Comparison, Drift, Health, Rule, Policy, Report, Policy Report, Runtime, and
Runtime Service contracts. Every supplied canonical value passes its owning
validator. Related Comparison/Drift/Health inputs must preserve their exact
identity chain.

Runtime Service State and terminal Service Result observations are mutually
exclusive. `running` is emitted only from an explicitly observed, validated
Service State whose lifecycle is `running`; it is never inferred from elapsed
time, a PID, systemd, process discovery, or a terminal result. Source
observation times cannot be after the projection time. The projection uses no
ambient clock.

## Overview contract

The Overview contains:

- overall condition: `healthy`, `degraded`, `critical`, `unknown`, or
  `unavailable`;
- attention: `none`, `review`, `urgent`, or `unknown`;
- Guardian state: `running`, `starting`, `stopping`, `stopped`, `failed`, or
  `not_observed`;
- freshness and completeness as separate dimensions;
- bounded counts for changes, Health conditions, and Alert lifecycle facts;
- category-level change summaries without raw before/after values;
- ordered attention items with localization tokens and canonical references;
- an optional attention summary that discloses correlated and omitted
  candidates when the bounded list is not exhaustive;
- an ordered closed set of read-only recommendation tokens;
- sorted source references for advanced inspection.

Missing evidence is `missing`/`unavailable`, never healthy. Evidence before
the exclusive freshness boundary is current; evidence at or beyond it is stale and makes the
Overview partial/degraded. Unsupported and partial upstream facts remain
explicit. Invalid contracts fail projection and produce no Overview.

## Precedence and recommendations

The closed precedence is:

1. urgent Alert, critical Health, or failed Guardian facts make the Overview
   critical and urgent;
2. review-level Alert/Health/Runtime/Guardian facts make it degraded and
   reviewable;
3. stale, partial, or unsupported evidence cannot produce healthy;
4. healthy requires explicit current, complete Health evidence and no higher
   precedence fact;
5. otherwise the result remains unknown or unavailable.

Recommendations are machine tokens only: `inspect_attention`,
`review_changes`, `run_fresh_check`, `inspect_evidence`, `inspect_failed_operation`,
`verify_guardian_operation`, and `no_action`. They direct read-only inspection
or a fresh check. They contain no shell command, generated advice, mutation, or
remediation authority.

## Determinism, localization, privacy, and bounds

Equivalent explicit inputs produce byte-identical canonical JSON and identity.
Strict decoding rejects unknown fields and trailing data. Collections, tokens,
references, freshness, and recommendations are bounded and canonically
ordered. Canonical user meaning is represented by stable tokens rather than
English or Hungarian prose.

Overview output excludes raw host values, before/after values, report prose,
configuration bodies, errors, filesystem paths, environment values, provider
payloads, endpoints, credentials, and secrets. It retains only bounded counts,
states, tokens, times, and canonical contract identities.

Model 1.1 ranks all attention candidates globally by severity, source
importance, stable tokens, and canonical source identity before applying the
256-item bound. Direct Health, Alert, Runtime, and Guardian evidence outranks
derived Policy and Rule views at equal severity. A Rule view is correlated
away only when a validated Policy record cites that exact Rule evaluation;
the downstream Policy view retains source traceability. The summary preserves
the pre-reduction total and exact correlated and omitted counts. A late
critical fact therefore cannot be displaced by earlier lower-severity input.
Attention facts with identical severity, localized meaning, source kind, and
contract are represented once after ranking; the summary counts the correlated
duplicates and the complete source set remains available in `Sources`.

Current partial or unsupported evidence recommends `inspect_evidence` rather
than promising that an identical observation will repair an intrinsic
limitation. `run_fresh_check` is reserved for missing or stale evidence.
Readers accept legacy Model 1.0 Overviews without summaries; new projections
use 1.1.

When Runtime does not complete, its validated component failure token is
retained as bounded attention evidence. Alert evaluation, Notification
planning, Notification delivery, cancellation, and timeout therefore remain
distinguishable. The generic `runtime_not_completed` token is reserved for a
non-completed Runtime result without a more specific canonical cause.

## Explicit exclusions

The model adds no renderer, CLI command, bare-`qwsg` behavior, Console, TUI,
HTTP/API, Dashboard, monitor, timer, goroutine, process probe, persistence,
restart recovery, configuration activation, provider, installation,
supervision, remediation, network operation, AI, packaging, deployment, or
release support claim.
