# Canonical Professional Notification Delivery

## Purpose and status

This document defines Notification Delivery Model 1.0 implemented by Task 029.
It is the permanent provider-neutral boundary between immutable Canonical Alert
Records and replaceable delivery providers.

Notification answers how an existing Alert Record is routed and attempted. It
never answers whether an alert exists. That decision belongs exclusively to the
Canonical Alert Engine.

## Architecture boundary

```text
Canonical Alert Records
          |
          v
pure Notification planner + explicit Delivery Policy + previous Queue State
          |
          v
Delivery Plan + proposed Queue State
          |
          v
explicit one-cycle adapter -> injected Provider interface -> Provider Result
          |
          v
Attempt + Status + Acknowledgement + Evidence + proposed Queue State
```

`internal/notification` imports only `internal/alert` from the QWSG canonical
packages. It cannot accept or evaluate Health, Rule, Policy, Drift, Scheduler,
Report, Configuration, Inventory, host, or monitoring records. The additive
`alert.ValidateRecord` entry point validates one immutable Alert Record without
running Alert evaluation or changing Alert semantics.

The package performs no persistence. Every Queue State is an immutable value
supplied by the caller and every operation returns a proposed next value.

## Model 1.0 contracts

The versioned contracts include:

- Delivery Policy, Route, Endpoint Reference, and Provider Binding;
- provider-neutral Delivery Envelope and Delivery Request;
- Planning Input, Delivery Plan, and per-Alert eligibility result;
- Notification Queue State and Queue Entry;
- Delivery Attempt, Status, Acknowledgement, and Evidence Reference;
- Provider Descriptor, Provider Result, injected Registry, and Provider
  interface;
- one-cycle result containing attempts and proposed Queue State.

Every public envelope uses schema version `1.0`. Policy, queue, request, plan,
attempt, status, acknowledgement, and cycle identities are content-derived.
Canonical validation rejects tampering, unsupported versions, unknown JSON
fields, unordered or duplicate records, invalid references, future Alert
inputs, incompatible provider bindings, and resource excess.

## Alert obedience and routing

The planner first calls the Alert-owned standalone validator. It copies only
the Alert Record identity, lifecycle/condition identities, event, severity,
category, reason token, event time, and bounded evidence references into the
provider-neutral envelope.

An Alert Record with decision `suppressed` is explicitly ineligible and creates
no request. Records with `alert` or `lifecycle` decisions may match enabled
Notification-owned routes by exact severity, category, and Alert event. Empty
route filters mean all valid values. Routing never changes or recreates Alert
identity, severity, category, lifecycle, acknowledgement, suppression,
expiration, correlation, recovery, or evidence semantics.

Fan-out creates one stable delivery identity for each Alert Record, route, and
endpoint tuple. Re-evaluating the same tuple against supplied Queue State does
not create a duplicate delivery. The idempotency key remains stable across all
attempts; each attempt receives a separate deterministic request identity.

## Channels and providers

Model 1.0 recognizes provider-neutral channel kinds `email`, `webhook`,
`slack`, `discord`, `telegram`, and `sms`. A Provider Descriptor declares the
supported channel kinds. An injected Registry rejects nil, invalid, duplicate,
or excessive providers.

Endpoint records contain opaque destination and optional secret references,
never destination values, credentials, or resolved secrets. Provider bindings
select a provider identity by channel. Canonical requests contain no SMTP,
HTTP, Slack, Discord, Telegram, SMS, SDK, authentication, payload-format, or
template semantics.

Task 029 includes no concrete transport. A Provider implementation owns its
encoding and transport behavior and remains replaceable. Tests use deterministic
in-memory conformance providers and require no network or credentials.

## Deterministic planning and queue order

The pure planner receives one explicit UTC evaluation time. It reads no clock,
environment, filesystem, network, process, randomness, configuration source,
or global mutable state.

Delivery requests and queue entries are ordered by:

1. severity priority: emergency, critical, warning, informational;
2. Alert event time, oldest first;
3. stable request or delivery identity.

Eligibility records and reference collections use stable canonical ordering.
Equivalent explicit inputs produce byte-identical canonical JSON and identities.

## Attempts, retries, and deadlines

Retry Policy defines a finite total-attempt limit, one absolute delivery
window, and a strictly increasing positive backoff sequence with exactly one
entry for each possible retry. Model bounds permit at most ten attempts and a
24-hour delivery window. A policy may express the Core Alpha default of three
total attempts within fifteen minutes without making that default an ambient
configuration source.

After a retryable, rate-limited, or indeterminate provider result, the next
planner evaluation deterministically returns one of:

- `retry_scheduled` with an explicit next-attempt time;
- `queued` with the next attempt request when due;
- `exhausted` after the attempt or deadline bound.

There is no jitter, sleep, timer, worker, or automatic retry loop. Provider
invocation happens only through `ExecuteCycle`, once per request in the supplied
Plan. Each call receives a context bounded by the request deadline. The adapter
does not plan another attempt and performs no persistence.

## Status, acknowledgement, failure, and evidence

Canonical statuses distinguish queued, accepted, provider-reported delivered,
retryable failure, retry scheduled, terminal failure, exhausted, and
indeterminate outcomes. Failure classes distinguish retryable, rate-limited,
indeterminate, authentication, authorization, invalid destination, rejected
payload, unsupported provider, and terminal failure.

A Delivery Acknowledgement means only provider acceptance,
provider-reported delivery, or unknown provider outcome. It never means human
receipt, reading, operator Alert acknowledgement, remediation, or successful
human action.

Delivery evidence contains only Alert/Request/Provider identities, an optional
opaque provider reference, and bounded sorted evidence tokens. Raw provider
responses, payloads, destination values, credentials, secret values, Alert
prose, host paths, and arbitrary metadata are excluded. Provider failures
produce Notification records only; they never create an Alert or recursively
invoke a failed channel.

## Resource and compatibility bounds

Model 1.0 bounds Alert Records, routes, endpoints, provider bindings, Queue
entries, attempts, backoff steps, strings, and provider evidence. Invalid input
fails closed. A failure for one provider or endpoint remains isolated in that
delivery attempt and does not alter other planned requests.

Adding a concrete provider, provider-specific canonical field, durable Queue,
configuration integration, secret resolution, callback, daemon, service, or
new taxonomy requires a separately authorized compatibility and security
review.

## Explicit exclusions

This architecture implements no Alert decision, upstream engine evaluation,
incident persistence, durable Queue, database, broker, daemon, timer, recurring
worker, automatic retry execution, production Email/Webhook/Slack/Discord/
Telegram/SMS transport, credential or secret backend, rendering/template
engine, inbound callback, CLI notification command, REST API, Dashboard,
Console, monitoring, channel-health Alert, remediation, repair, remote agent,
licensing, billing, AI, infrastructure mutation, deployment, or release.

The Alert Engine decides. Notification Delivery plans and records bounded
delivery through replaceable providers. Operational hosting remains later work.

Runtime calls planning with Alert Records only and invokes at most one provider
cycle when requests exist. It does not select routes, retry attempts, classify
provider outcomes, or persist queues.

Task 046 is the separately authorized concrete Community SMTP adapter and
Guardian-hosting layer. It activates this unchanged provider interface,
delivery policy, and queue contract through the canonical configuration and
private credential boundaries documented in
`COMMUNITY_EMAIL_NOTIFICATIONS.md`. Other channels, callbacks, managed
delivery, and entitlement remain excluded.
