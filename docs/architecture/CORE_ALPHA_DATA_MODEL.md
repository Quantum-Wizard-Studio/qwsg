# Core Alpha Inventory Data Model

## Authority

This document owns the canonical logical inventory model for Slice 1. It defines semantics, not a programming language, serialization library, database, or retention period. JSON is the required external representation; an implementation may use equivalent internal types.

## InventorySnapshot

| Field | Type | Rule |
| --- | --- | --- |
| `schema_version` | string | `major.minor`; required |
| `snapshot_id` | opaque string | unique, non-secret, generated locally |
| `request_id` | opaque string | correlates logs and invocation |
| `instance_id` | opaque string | privacy-safe stable local identity; not raw hardware ID |
| `observed_at` | UTC timestamp | start of collection with explicit offset |
| `completed_at` | UTC timestamp | envelope completion time |
| `fresh_until` | UTC timestamp | bounded validity claim, never inferred as health |
| `duration_ms` | integer | monotonic elapsed duration |
| `status` | enum | `complete`, `partial`, `failed`, `cancelled` |
| `categories` | array | one independently valid `CategoryResult` per requested category |
| `errors` | array | request/envelope-level `InventoryError` only |
| `redactions` | array | categories/counts, never removed values |
| `producer` | object | QWSG version and contract version; no environment dump |

## CategoryResult

Required fields are `category_id`, `contract_version`, `status`, `observed_at`, `completed_at`, `fresh_until`, `duration_ms`, `collector_id`, `privilege_used`, `source_summary`, `items`, `errors`, and `redactions`.

`status` is one of `available`, `unavailable`, `unsupported`, `permission_denied`, `timeout`, `error`, or `cancelled`. `privilege_used` is `ordinary_user` for Slice 1. An unavailable result contains no fabricated items. A category can be `available` with item-level unknown fields only when the known subset remains truthful.

## InventoryItem and values

Each item has stable `item_id`, optional privacy-safe `display_label`, `kind`, and typed `facts`. Every fact is a `FactValue`:

- `value`: string, integer, number, boolean, timestamp, duration, byte count, percentage, enum, object, or array;
- `unit`: canonical unit where applicable;
- `quality`: `observed`, `derived`, `reported`, `unknown`, or `redacted`;
- `sensitivity`: `public`, `operational`, `host_identifying`, `network_sensitive`, or `secret_prohibited`;
- `provenance`: source type, collector, observation time, and transformation identifier;
- `reason`: required for `unknown` or `redacted`.

Unknown is a value quality, not `null`, zero, empty string, false, or omission. Optional inapplicable fields may be absent; required but unobtainable fields use `quality: unknown` with a reason.

## Minimum category schemas

| Category | Minimum facts |
| --- | --- |
| `os` | distribution identity, version, kernel release, machine architecture, virtualization/container indication when detectable |
| `host` | privacy-safe display label, local instance identity basis, boot/session identity only if non-sensitive |
| `cpu` | visible logical capacity, architecture, model family only when safely available |
| `memory` | total/available memory, total/used swap, source semantics |
| `storage` | stable mount identity, filesystem type, total/available bytes, read-only flag; no file traversal |
| `network` | interface identity, state, address families and redacted addresses according to policy; no payloads |
| `services` | stable unit/service identity, manager type, active state; no control or full command lines |
| `components` | allowlisted component ID, version, detection source; no package installation or broad package inventory |
| `collector_capabilities` | collector ID, availability, missing prerequisite, ordinary-user access result |

## Provenance

`source_type` is one of `kernel_virtual_file`, `system_api`, `filesystem_metadata`, `service_manager`, or `allowlisted_command`. Provenance records a logical source label, not secret-bearing paths or raw output. Derived facts identify the transformation and input fact names. Observation and completion timestamps belong to every category so mixed-age data cannot masquerade as simultaneous.

## Errors

`InventoryError` contains `code`, `category_id` when applicable, `class`, `safe_message_key`, `retryable`, `occurred_at`, and redacted `details`. Classes are `capability`, `permission`, `timeout`, `invalid_evidence`, `resource_limit`, `cancelled`, `storage`, and `internal`. Human prose is rendered from `safe_message_key` and is not canonical data.

## Freshness and partial semantics

Freshness is a collection-validity deadline, not a health assertion. After `fresh_until`, consumers label facts stale and must not claim a current inventory. `partial` means at least one requested category is usable and at least one is not. `failed` means no requested category produced usable facts or envelope validation failed. Consumers display per-category status and do not collapse partial into complete.

## Validation invariants

- `observed_at <= completed_at <= fresh_until` for a successful fresh envelope.
- IDs and enum tokens are locale independent and bounded in length.
- Categories and items are unique by stable ID and deterministically ordered.
- Byte counts and durations are non-negative; percentages use a documented range.
- A redacted or unknown fact never carries the original value.
- `complete` requires every requested category to be `available`.
- Secret-prohibited values are rejected, not merely labeled.
- Unsupported major schema versions are rejected without overwriting prior valid state.

## Storage evolution

The serialized envelope is self-describing. Writers produce one supported major version. Readers reject future majors and accept additive minor fields only under documented forward-compatibility rules. Migration is explicit, testable, and non-destructive; no implementation silently resets incompatible inventory.
