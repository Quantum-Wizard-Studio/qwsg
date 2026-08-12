# Canonical Professional Scheduler

## Purpose

Task 027 establishes QWSG's canonical scheduling capability. It evaluates
validated Schedule Definition 1.0 records from one immutable Effective
Configuration 1.0, produces versioned scheduler records, and optionally passes
due work through the existing Canonical Command resolver and Pipeline
Orchestrator.

The Scheduler does not define configuration, policy, command, pipeline, report,
or presentation semantics. It does not collect evidence and does not contain a
daemon. Its local adapter performs exactly one explicitly requested cycle.

## Architectural boundary

```text
Configuration Sources
        |
Effective Configuration 1.0
        |
Schedule Definition 1.0 + Scheduler State 1.0 + Clock Observation
        |
pure Scheduler Engine
        |
Evaluation + Events + Execution Requests + proposed State
        |
explicit one-cycle local adapter
        |
Canonical Command Definition -> Pipeline Orchestrator -> Command Execution
        |
Execution Results + durable Scheduler State
```

Schedule Definition describes intent. The engine calculates due occurrences and
plans work. An Execution Request describes planned canonical work but performs
none. Only the adapter resolves a known Command profile and invokes the existing
Pipeline. A future daemon may call the adapter, but recurring lifecycle,
monitoring, alerts, notifications, and service management are separate
architectural components.

## Canonical contracts

All Scheduler contracts are version `1.0`:

- `qwsg.scheduler-evaluation`: immutable evaluation envelope;
- Scheduler Evaluation Record: per-schedule decision and reason;
- `qwsg.scheduler-state`: restart-safe local state;
- `qwsg.scheduler-event`: bounded self-observability fact;
- `qwsg.scheduler-execution-request`: reserved canonical work;
- `qwsg.scheduler-execution-result`: execution outcome and traceability;
- `qwsg.scheduler-state-envelope`: persisted state plus SHA-256 integrity.

The engine and decision taxonomy versions are independently declared as `1.0`.
Unknown schemas or versions, malformed values, inconsistent identities, invalid
ordering, corrupt persisted bytes, and unsupported references fail closed.

Identities are SHA-256 content identities with type prefixes. Canonical JSON is
generated only after validation. Ordering is stable: schedules by identity,
candidates by retry status, descending priority, occurrence time, then schedule
identity; results by completion time and identity. Derived circular references
are excluded only from the enclosing evaluation identity and remain covered by
their own record identities.

## Evaluation semantics

Evaluation uses the half-open window `(last evaluation, observation]`. A first
observation establishes an interval anchor and never fabricates an immediate
run. Calendar schedules are evaluated in their configured IANA time zone.
Ambiguous local times obey the Schedule Definition's first- or second-occurrence
policy; nonexistent local times have no matching UTC instant and are skipped.
An unavailable time zone produces `indeterminate` and no request.

The decision taxonomy is:

- `disabled`: configuration disables the schedule;
- `not_due`: no occurrence lies in the evaluation window;
- `due`: a request was deterministically reserved;
- `skipped`: configured misfire or overlap behavior drops work;
- `queued`: one bounded replacement is retained;
- `delayed`: global concurrency has no capacity;
- `inapplicable`: existing Command 1.0 cannot represent the requested scope;
- `indeterminate`: safe evaluation is impossible.

For multiple missed occurrences, `skip` drops them, `run_once` coalesces them to
the latest occurrence, and `indeterminate` creates no work. Global concurrency
comes from Effective Configuration. `forbid` prevents overlap for one schedule;
at most one pending occurrence is retained, so no unbounded backlog exists.
`allow` permits multiple requests only within the same global bound.

An empty Schedule `check_ids` list means the entire referenced Command profile.
A non-empty list is inapplicable under Command Definition 1.0 because that
contract cannot express Check scoping. The Scheduler never silently broadens
the selection or invents another Command path.

## Clock and restart safety

Every cycle supplies wall time, a process-session identity, and monotonic elapsed
time. Wall and monotonic deltas that diverge beyond the fixed tolerance, move
backward, or regress produce a clock-discontinuity event and indeterminate
schedule decisions. No due work is emitted for that observation.

After a process-session change, reservations still marked active become
`interrupted` results. A retry is eligible only when the referenced configured
Retry Policy permits another finite attempt and its deterministic delay has
elapsed. Retry delays are bounded exponential delays without randomness. The
latest result for an occurrence prevents a completed retry from being planned
again.

## State, locking, and failure behavior

The engine is pure and returns a proposed next state. The local adapter requires
an explicit absolute state directory and injected clock, time-zone resolver,
lock provider, Command resolver, Pipeline executor, Effective Configuration, and
Command selection. It discovers no paths or services.

The file store uses a `0700` directory, `0600` files, strict JSON decoding,
payload integrity verification, temporary-file synchronization, atomic rename,
and directory synchronization. State and result retention are bounded. Corrupt,
tampered, unsupported, or configuration-mismatched state is rejected; it is not
silently reset.

The local file lock uses a non-blocking operating-system advisory lock and a
validated stable owner identity. Contention and invalid acquisition fail before
evaluation. The adapter persists reservations before invoking the Pipeline and
persists final results afterward. A crash between these writes is recovered as
an interrupted reservation on the next process session.

State, lock, time-zone, and Scheduler failures do not alter the manual Command
or Pipeline APIs. They stop only the explicitly requested scheduling cycle.

## Command, Pipeline, Policy, and Report integration

Each request references the exact Effective Configuration, scheduler evaluation,
schedule, occurrence, attempt, Command profile, retry policy, priority, timeout,
and Check scope. The adapter calls `command.ResolveProfile`; only the existing
`pipeline.Orchestrator` executes the resulting definition.

Execution Results record Command Execution identity and completeness, stage
contract names, and Policy Evaluation identities and outcomes present in the
typed Pipeline result. These values are traceability evidence only. Scheduler
does not re-evaluate Inventory, Compare, Drift, Health, Rule, Policy, or Report,
and a Policy Outcome is never interpreted as action authority.

The pure Alert Engine may consume validated Scheduler Evaluation Events as
scheduler-operational evidence. It preserves Scheduler evaluation, event,
schedule, request, reason, and timestamp references and does not recalculate
due work, interpret Command results, persist Scheduler state, or control the
one-cycle adapter. Scheduler state failure, clock discontinuity, and lock
contention have explicit Alert mappings; initialized, restart-recovered, and
execution-completed events can resolve the corresponding operational Alert
condition.

## Privacy, resource, and compatibility guarantees

Operation is local, offline, deterministic, presentation-independent, and
AI-independent. The contracts store identities, bounded status facts, and
canonical execution provenance rather than arbitrary payloads, credentials, or
host evidence. No network, shell, arbitrary executable, host mutation, remote
execution, notification, or remediation path exists.

Task 027 is additive. Existing Command profiles, Pipeline orchestration, manual
CLI behavior, Policy, Configuration, and Report contracts remain unchanged.
Schema `1.0` readers reject unsupported versions rather than guessing. Future
extensions require explicit new versions or additive validated contracts.

## Explicit exclusions

This Scheduler architecture implements no recurring loop, daemon, service
installation, automatic startup, monitor, Alert decision, notification
delivery, maintenance window, remediation, remote coordination, Dashboard,
Terminal UI, REST API,
configuration activation, secret backend, telemetry, AI, or machine learning.

Cycle Result 1.0 includes one immutable bounded Execution Trace per Request and
a public validator. The observational trace exposes only the Command Definition
and Command Execution already produced by Scheduler-owned Pipeline invocation,
or a fixed failure token; Scheduler retains all execution and state ownership.
