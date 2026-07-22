# Canonical System Inventory Developer Guide

## Build and test

```bash
make build
make test
make vet
make fmt-check
```

Run one read-only observation with:

```bash
build/qwsg inventory
```

Exit status `0` means every requested category was available, `2` means a valid partial inventory, and `1` means no usable inventory or a fatal error. A partial result is expected on hosts where an optional legacy collector such as systemd service discovery is unavailable.

## Package responsibilities

- `internal/collector`: Collector contract, Registry orchestration, Linux evidence acquisition, privacy transformations, and legacy Category contributions.
- `internal/inventory`: Inventory 1.0 compatibility model, canonical domain model, deterministic assembly, and validation.
- `internal/app`: one-shot coordination and synchronized legacy/canonical aggregation.
- `cmd/qwsg`: CLI serialization and exit policy only; it does not discover host state.

Add a collector by implementing `collector.Collector`, declaring a complete finite descriptor, returning one structured Result, and registering it through the Registry. Collectors must use bounded read-only evidence, observe cancellation, avoid user-facing prose, redact before returning, and never call other collectors. Add parser fixtures, error-state tests, Registry/integration coverage, privacy review, and documentation with every capability.

## Compatibility rules

Do not add a second host model or direct discovery path. New consumers use `Snapshot.Canonical`; compatibility code alone may use legacy `Snapshot.Categories`. A contract change follows `ai/core/12_INVENTORY_ARCHITECTURE.md` versioning rules. Removing or renaming an existing Inventory 1.0 field requires a separately authorized major-version migration.

## Test fixtures and troubleshooting

Parser tests use synthetic evidence and must never contain real host identifiers. When a collector is unavailable, inspect its structured `health_status` and safe error code. `permission_denied`, `timeout`, `cancelled`, `resource_limit`, and `error` are explicit collector outcomes; they are not host-health verdicts. Never work around missing evidence with privilege escalation or unbounded scanning.
