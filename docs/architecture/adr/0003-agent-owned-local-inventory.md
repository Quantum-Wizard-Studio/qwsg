# ADR 0003: Agent-owned Local Inventory

- Status: Accepted for Core Alpha Slice 1
- Date: 2026-07-20

## Context

QWSG must remain useful without a Console or vendor cloud, and Console security is unresolved.

## Decision

The Agent owns local inventory acquisition, validation, and any authorized local persistence. Slice 1 exposes local CLI/JSON only. A future Console consumes a redacted Agent contract and cannot access collectors or privileged host interfaces directly.

## Consequences

Slice 1 can proceed without resolving network authentication or topology. Console loss cannot disable discovery, and later transport work remains isolated behind gate `AG-004`.
