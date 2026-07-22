# ADR 0002: Versioned Inventory Envelope

- Status: Accepted for Core Alpha Slice 1
- Date: 2026-07-20

## Context

The Functional Specification requires versioned data, explicit unknown states, evidence provenance, and safe incompatible-schema behavior.

## Decision

All discovery output uses one versioned `InventorySnapshot` envelope with independently status-bearing categories, typed facts, provenance, sensitivity, freshness, errors, and redactions. Unknown is explicit and partial is not complete.

## Consequences

CLI, persistence, tests, and a future Console share one canonical truth. Unsupported major versions fail without overwrite; additive minor evolution remains possible.
