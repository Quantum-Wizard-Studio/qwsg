# Change Record Schema 1.0

## Contract

`qwsg compare` emits the canonical `qwsg.comparison` JSON document. Version
`1.0` contains a deterministic comparison envelope and an ordered `changes`
array. This is the machine interface for system evolution; human output is a
rendering of this document and is not a second comparison implementation.

Each Change Record has:

| Field | Meaning |
| --- | --- |
| `id` | Stable SHA-256-derived identifier for the comparison, path, and change type. |
| `layer` | Canonical Inventory layer identifier. |
| `object_id` | Layer or resource identity owning the compared value. |
| `path` | RFC 6901-escaped canonical path to the compared semantic value. |
| `change_type` | `added`, `removed`, `modified`, or `unchanged`. |
| `previous` | Previous typed value, or `null` for `added`. |
| `current` | Current typed value, or `null` for `removed`. |
| `comparison_timestamp` | The `to` snapshot completion time. |
| `metadata` | Stable comparison classification such as object kind and fact name. |

A typed value preserves the Inventory fact's `type`, JSON `value`, `unit`,
`quality`, `sensitivity`, and `reason_code`. Layer status and resource kind use
the same typed representation. JSON `null`, zero, false, and an empty string
remain distinct values.

## Determinism and ordering

The comparison timestamp is evidence-derived and never uses wall-clock time.
Comparison IDs and record IDs contain no random input. Records are sorted by
layer, object identifier, path, and change type. Repeating a comparison over
the same validated snapshots and selectors therefore produces byte-identical
indented JSON.

`counts` exactly matches the four record groups. Identical snapshots have zero
`added`, `removed`, and `modified` records; their explicit `unchanged` records
remain available as comparison evidence.

## Compatibility

Schema additions require documented backward-compatible semantics. Removing or
renaming a field, changing identity or ordering, changing value semantics, or
reclassifying an existing record requires a new major schema version. Consumers
must reject unsupported schema versions rather than guess.

Change Records are factual. They never encode drift policy, health, severity,
alerting, scoring, recommendation, or remediation.
