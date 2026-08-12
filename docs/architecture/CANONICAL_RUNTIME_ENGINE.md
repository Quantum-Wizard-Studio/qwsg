# Canonical Runtime Engine

## Status

Runtime Model 1.0 is canonical from Task 030. `internal/runtime` coordinates
one explicitly requested bounded cycle; it is not a service or daemon.

## Deterministic cycle

The coordinator validates an explicit Execution Context and existing canonical
inputs, then performs this fixed order:

1. invoke the existing Scheduler Cycle exactly once;
2. validate its Cycle Result and immutable execution traces;
3. send the Scheduler Evaluation to Alert exactly once;
4. in Scheduler Request order, validate successful Command Executions and pass
   exact typed Health, Rule, Policy, and Policy Report values to Alert while
   threading Alert's proposed state;
5. pass only Alert Records returned by Alert to Notification planning;
6. invoke the Notification provider cycle at most once and only when requests
   exist;
7. return a canonical Runtime Result and proposed `idle` Runtime State.

Scheduler retains due-time, locking, persistence, Command resolution, Pipeline
execution, and completion ownership. Pipeline retains stage ordering and engine
invocation. Alert alone decides whether alerts exist. Notification alone owns
delivery planning and provider outcomes. Runtime re-evaluates none of them.

## Contracts, evidence, and failure

Runtime Model 1.0 defines Execution Context, Runtime State, Runtime Input,
Component Result, Runtime Event, Runtime Evidence, and Runtime Result contracts.
Content identities use SHA-256 over canonical JSON. Caller-supplied identity,
initiator, start, and deadline are explicit; time observations are injected,
UTC-normalized, and bounded by that window.

Outcomes are `completed`, `partial`, `failed`, `cancelled`, and `timed_out`.
Components are `completed`, `failed`, or `skipped`. Checks occur at component
boundaries and between traces. Completed evidence survives later failure.
Failure records contain fixed tokens, never raw errors, stage values, provider
payloads, destinations, secrets, credentials, or host paths.

Component failure tokens are a closed privacy-safe vocabulary, including
`scheduler_cycle_failed`, `alert_evaluation_failed`,
`notification_planning_failed`, `notification_delivery_failed`, `cancelled`,
and `timed_out`. The presentation boundary may translate these tokens but must
not expose the underlying Go error. A valid large Policy Report is passed to
Alert by canonical aggregate identity, so source cardinality alone cannot stop
Runtime before Notification planning.

## Scheduler trace seam

Scheduler Cycle Result 1.0 exposes one bounded immutable Execution Trace per
Request. It binds Request identity to the Command Definition, Command Execution,
and a fixed failure token. Scheduler's public validator checks schema, identity,
cardinality, ordering, and correlation. Runtime additionally validates every
successful execution against its canonical Command Plan before typed
projection. This observational seam adds no scheduling or execution decision.

## State and durability

Scheduler may persist state inside its existing adapter. Alert, Notification,
and Runtime return proposed states only; Runtime persists none of them. Runtime
always proposes `idle` on terminal return and makes no cross-engine atomicity or
durability claim.

## Bounds and exclusions

Owning engines retain their existing occurrence, record, queue, retry, and
evidence limits. Runtime additionally bounds its event/evidence envelopes and
performs no loop, retry, sleep, polling, collection, or transport of its own.

Runtime 1.0 includes no Linux daemon, systemd integration, service installation,
background worker, continuous monitoring, watchdog, durable Runtime/Alert/
Notification store, REST API, Dashboard, remote execution, clustering, fleet
management, remediation, AI, infrastructure mutation, or concrete provider.

## Runtime Service caller

The Canonical Runtime Service repeatedly invokes Runtime through the public
one-cycle `Run` boundary. It supplies a new explicit Execution Context and
mechanically forwards only Runtime-proposed states. Runtime retains all cycle
ordering, validation, component, result, and cancellation semantics. The
Service cannot call downstream engines or turn Runtime into a competing loop.
