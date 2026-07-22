# ADR 0001: Read-only Collector Boundary

- Status: Accepted for Core Alpha Slice 1
- Date: 2026-07-20

## Context

Task 008 explicitly authorizes a read-only discovery slice, default non-root operation, truthful permission denial, and no product-initiated privilege escalation.

## Decision

Collectors return facts through a versioned contract and cannot mutate monitored state, escalate privilege, construct shell strings, persist data, or invoke one another. The coordinator alone manages bounded execution, cancellation, normalization, and aggregation.

## Consequences

Restricted environments produce partial inventory rather than unsafe elevation. Future privileged or mutating behavior needs a separate component, contract, threat model, and owner authorization.
