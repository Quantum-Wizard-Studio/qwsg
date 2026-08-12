# Canonical Runtime Service

## Status

Runtime Service Model 1.0 is canonical from Task 031. The implementation is
`internal/runtimeservice`. It is an explicitly invoked continuous process
contract, not an installed, supervised, or production-supported daemon.

## Ownership boundary

The Runtime Service owns recurrence and local process lifecycle only:

```text
Runtime Service
      |
      v
Runtime Engine (one cycle)
      |
      v
Existing canonical components
```

The Service calls one injected Runtime Runner and imports no Scheduler,
Command, Pipeline, Health, Rule, Policy, Report, Alert, or Notification engine
package. Runtime remains the sole owner of a complete execution cycle. The
Service does not inspect or reinterpret component results.

## Fixed-rate recurrence

Service Definition 1.0 provides a stable service identity, positive interval,
positive per-cycle timeout, and the `immediate` startup mode. The first Runtime
cycle uses the explicit service start as its nominal boundary. Every later
boundary is calculated from the preceding nominal boundary, never from cycle
completion time.

Calls are synchronous and sequential, so at most one Runtime cycle is active.
If a cycle spans nominal boundaries, the Service records one compressed missed-
interval event and advances to the first non-elapsed boundary. It never overlaps
cycles, performs catch-up bursts, retries Runtime, polls, or starts workers.

Clock and waiter contracts are injected. Equivalent explicit input, clock
observations, Runtime results, and sink outcomes produce the same ordered
records and identities. Tests use no ambient sleeps.

## Lifecycle and shutdown

Lifecycle states are `created`, `starting`, `running`, `stopping`, `stopped`,
and `failed`. Caller cancellation is checked before waiting and invocation and
is propagated through a bounded per-cycle context. Graceful shutdown cancels
active Runtime work and waits synchronously for Runtime's existing context
contract. The Service does not detach or abandon a non-compliant Runner and
therefore makes no bounded-shutdown claim when that dependency ignores context.

The separate signal adapter maps only SIGINT and SIGTERM to context
cancellation using the Go standard library. Registration is scoped to one
explicit call and released on return. It installs no unit, startup policy,
supervisor, watchdog, or package.

## In-memory state handoff

After a validated Runtime Result, the Service copies only exact Runtime-owned
`NextState`, `FinalAlertState`, and `FinalNotificationQueue` into the next
Runtime Input. Effective Configuration, Alert controls, Notification policy,
and other seed fields remain unchanged. Scheduler retains its own state and
persistence ownership.

This handoff is process-local and non-durable. Task 031 provides no restart
continuity, crash recovery, replay, durable queue, checkpoint, or transaction.
A supported Version 1.0 Agent still requires separately governed persistence,
recovery, configuration activation, operational integration, and support gates.

## Contracts, evidence, and privacy

Runtime Service Model 1.0 defines Definition, State, Input, Event, Evidence,
and Result contracts with strict validation, canonical JSON, content identity,
UTC times, fixed tokens, and bounded counters. Events describe lifecycle
occurrences. Evidence records refer only to their Event and at most the cycle
and Runtime Result identities.

The synchronous injected sink decides how evidence is consumed. Service retains
only counters and the last Runtime Result reference; it keeps no growing event
history and implements no storage. Sink refusal is a Service-owned terminal
failure. Records exclude Runtime payloads, Alert Records, Notification Queue
entries, configuration bodies, destinations, secret references, provider
payloads, report prose, host paths, credentials, raw errors, and signal details.

## Explicit exclusions

Runtime Service 1.0 adds no systemd/init integration, installation, enablement,
supervision, automatic restart, watchdog, socket activation, PID file,
privilege management, filesystem layout, log rotation, package, persistence,
monitoring/product-health logic, configuration reload, concrete provider,
interface, public listener, remediation, remote execution, clustering, fleet
management, licensing, billing, AI, infrastructure mutation, deployment, or
release support claim.
