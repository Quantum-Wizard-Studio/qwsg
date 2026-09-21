# Canonical System Inventory v1 Implementation

## Status and authority

This document describes the Task 014 implementation of the canonical Inventory Architecture in `ai/core/12_INVENTORY_ARCHITECTURE.md`. That architecture remains normative. The implementation is an internal pre-alpha contract and is not a supported public release.

## Runtime boundary

The existing Collector Registry is the only host-discovery entry point. `internal/app.Collect` executes the Registry once and uses the same structured Results to produce two synchronized views:

1. the established Inventory `1.0` category/item envelope for backward compatibility;
2. the additive `canonical_inventory` representation for all new platform consumers.

No canonical consumer performs direct Linux discovery. The canonical projection is validated before output and cannot silently repair invalid evidence.

The optional Task 016 Inventory Store is a post-validation consumer. It
persists the complete synchronized legacy/canonical envelope and never invokes
or modifies collectors. Its format, atomicity, integrity, permissions, and
retention contract is defined in
`INVENTORY_PERSISTENCE_AND_DIGITAL_TWIN.md`.

## Implemented collectors

| Collector | Capability | Primary bounded Linux evidence | Canonical layer | Dependency |
|---|---|---|---|---|
| `host` | `host` | hostname API and `/etc/machine-id` | `host` | none |
| `os` | `os` | `/etc/os-release` | `operating_system` | none |
| `kernel` | `kernel` | `uname(2)` | `operating_system` | none |
| `cpu` | `cpu` | `/proc/cpuinfo`, Go runtime | `hardware` | none |
| `memory` | `memory` | `/proc/meminfo` | `hardware` | none |
| `storage` | `storage` | `/sys/class/block` | `storage` | none |
| `filesystem` | `filesystem` | `/proc/self/mounts`, `statfs(2)` | `storage` | `storage` |
| `network` | `network` | `/sys/class/net` | `network` | none |
| `virtualization` | `virtualization` | bounded container markers, cgroup and DMI metadata | `host` | `host` |

Previously implemented capability, service, and allowlisted component collectors remain registered for Inventory 1.0 compatibility. Task 014 does not extend service or component monitoring.

Every descriptor declares Linux support, ordinary-user privilege, a finite timeout and maximum, a two-MiB normalized-output ceiling, compatibility `1.0`, and sensitivity classes. Registry execution retains deterministic dependency planning, availability checks, parent cancellation, per-collector timeout, panic isolation, and structured failure Results. The coordinator now also enforces the declared normalized-output limit.

## Canonical contract

`inventory.SystemInventory` contains:

- `schema_name: qwsg.inventory`, schema/profile identity, snapshot/request/subject identity, timing, status, and producer;
- deterministic `collector_results` with identity, capability, platform, timing, status, warnings, errors, and metadata;
- deterministic canonical `layers` with timing, status, contributing collectors, resources, issues, redactions, and metadata;
- resources with privacy-safe namespaced IDs, kind, lifecycle state, typed facts, relationships, labels, provenance, and metadata;
- snapshot-level structured issues, redactions, and bounded projection metadata.

Layers and resources are sorted by stable machine identifiers. Go JSON map serialization provides deterministic fact-key ordering, while timestamps and unique request IDs correctly differ between observations. `inventory.Validate` validates both the legacy envelope and canonical representation, including schema identity, timestamp ordering, collector/layer/resource uniqueness and order, required resource fields, prohibited-secret rejection, and relationship referential integrity.

## Privacy and evidence

Stable subject, block-device, mount, and interface identifiers are one-way namespaced SHA-256 derivatives truncated to 128 bits. Raw machine IDs, hostnames, interface names, network and hardware addresses, mount paths, raw device names, and service identities are not emitted. Redacted facts have no value and include locale-independent reason tokens. Collection reads bounded kernel interfaces and does not elevate privilege, recurse through the filesystem, access the network, persist data, or mutate the host.

Persistence is a separate explicitly invoked post-validation adapter. It stores
no raw collector evidence and rejects a prohibited-secret fact in either the
compatibility or canonical representation.

## Compatibility

The top-level schema remains `1.0`; existing `categories`, status calculation, exit codes (`0` complete, `1` failed, `2` partial), redactions, and producer fields remain present. `canonical_inventory` is an additive field derived loss-aware from those categories. New QWSG modules must consume the canonical representation. Legacy consumers may ignore the additive field under the existing `1.x` unknown-field rule.

## Verification boundary

Tests cover every registered collector contract, required collector set, parser fixtures, privacy IDs, dependencies, deterministic Registry and canonical ordering, cancellation, timeout, panic and output-limit isolation, partial/fatal aggregation, secret rejection, relationship integrity, legacy compatibility, and end-to-end canonical assembly. A live Linux run verifies all nine Task 014 collectors as available in the implementation environment; unrelated legacy service discovery may truthfully leave the overall snapshot partial when systemd access is unavailable.

## Stable protected service identity (Task 089)

The running systemd service collector uses the exact unit token (first column)
from bounded `systemctl list-units --type=service --state=running
--no-legend --plain --full` observation. `--full` prevents display ellipsization
from truncating the identity input. Case, instance suffixes and systemd escaping
are preserved; descriptions and enumeration positions are not identity inputs.
Malformed, duplicate or non-running rows fail collection without publishing raw
evidence. Service coverage remains running units of the system manager only.

The item ID is `systemd-unit-v1:` followed by the first 128 bits of
`SHA-256("qwsg:systemd-unit:v1:" + unitToken)` in lowercase hexadecimal, reusing
the existing collector `privacyID` primitive. The canonical resource ID adds
`services:`. Items and canonical resources are sorted by protected ID. The
`service_identity` fact remains redacted with no value. Identical unit tokens
retain identity across independent runs; ordering and description changes do
not replace a service. A renamed unit is a different identity, not a claim
about the identity of its executable or workload. Presence changes refer only
to the observed running set, not installation/removal of unit files.

These unkeyed deterministic identifiers are **pseudonyms, not anonymous IDs**.
They allow correlation, including across hosts for the same unit token, and
candidate-name dictionary guessing. This preserves the existing deterministic
privacy primitive's semantics; it is not encryption or resistance to guessing.
There is no new secret, registry, key lifecycle, persistence or network lookup.
Host, network, mount, device and service-name redactions remain intact.

Historical nonempty ordinal service inventories remain readable, unchanged,
under Inventory/Store 1.0. They cannot prove stable service identity. The
[Comparison compatibility boundary](SNAPSHOT_COMPARISON_ENGINE.md#service-identity-compatibility)
refuses exact comparison involving such evidence. No bulk migration occurs.
