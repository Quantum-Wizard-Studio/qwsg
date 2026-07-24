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

The command preserves JSON as its compatibility default. Use
`--format human` for the terminal summary. Parsing, human rendering, JSON
serialization, and exit policy remain in `cmd/qwsg`; domain and collection
packages contain no user-facing prose.

Exit status `0` means every requested category was available, `2` means a valid partial inventory, and `1` means no usable inventory or a fatal error. A partial result is expected on hosts where an optional legacy collector such as systemd service discovery is unavailable.

## Package responsibilities

- `internal/collector`: Collector contract, Registry orchestration, Linux evidence acquisition, privacy transformations, and legacy Category contributions.
- `internal/inventory`: Inventory 1.0 compatibility model, canonical domain model, deterministic assembly, and validation.
- `internal/app`: one-shot coordination and synchronized legacy/canonical aggregation.
- `internal/inventorystore`: validated file-backed Digital Twin envelope,
  atomic save/load/list behavior, integrity checking, and bounded retention.
- `cmd/qwsg`: CLI serialization and exit policy only; it does not discover host state.

## Explicit persistence

Create or use a private store with an explicit absolute path:

```bash
build/qwsg inventory save --store /absolute/private/qwsg-inventory
build/qwsg inventory load --store /absolute/private/qwsg-inventory
```

The default retention is 10. A different value is fixed at store creation and
must be supplied consistently:

```bash
build/qwsg inventory save --store /absolute/private/qwsg-inventory --retention 5
build/qwsg inventory load --store /absolute/private/qwsg-inventory --retention 5
```

Use `--snapshot <filename>` with `inventory load` for an explicit stored
snapshot. The store accepts only one safe base filename returned by its
deterministic listing.

Snapshot Explorer adds:

```bash
build/qwsg inventory list --store /absolute/private/qwsg-inventory
build/qwsg inventory info --store /absolute/private/qwsg-inventory
```

List validates every displayed entry through the store load boundary. Info and
load accept `--snapshot`, otherwise they select the validated latest entry.
`QWSG_STORE` and `QWSG_FORMAT` are explicit session configuration equivalents;
command-line values take precedence. Human rendering escapes terminal control
characters and shows only status, times, schema identity, and aggregate counts.

Save performs one collection, validates and atomically persists it, emits the
same JSON, and retains its status exit code. Load performs no collection,
revalidates the envelope and Inventory, emits JSON, and returns the stored
status exit code. Exit `2` remains a truthful partial-but-usable result.

Store roots and directories require mode `0700`; metadata and snapshots require
`0600`. Symlinks, permissive modes, corruption, duplicate JSON keys,
unsupported versions, integrity mismatches, stale transaction artifacts, and
retention mismatch fail closed.

Add a collector by implementing `collector.Collector`, declaring a complete finite descriptor, returning one structured Result, and registering it through the Registry. Collectors must use bounded read-only evidence, observe cancellation, avoid user-facing prose, redact before returning, and never call other collectors. Add parser fixtures, error-state tests, Registry/integration coverage, privacy review, and documentation with every capability.

## Compatibility rules

Do not add a second host model or direct discovery path. New consumers use `Snapshot.Canonical`; compatibility code alone may use legacy `Snapshot.Categories`. A contract change follows `ai/core/12_INVENTORY_ARCHITECTURE.md` versioning rules. Removing or renaming an existing Inventory 1.0 field requires a separately authorized major-version migration.

## Test fixtures and troubleshooting

Parser tests use synthetic evidence and must never contain real host identifiers. When a collector is unavailable, inspect its structured `health_status` and safe error code. `permission_denied`, `timeout`, `cancelled`, `resource_limit`, and `error` are explicit collector outcomes; they are not host-health verdicts. Never work around missing evidence with privilege escalation or unbounded scanning.

Persistence tests use `t.TempDir` only. A store checksum detects corruption but
does not authenticate data. Do not delete a stale lock, retirement artifact, or
operator store automatically; preserve it and review the interrupted
transaction. There is no daemon, scheduler, database, comparison, monitoring,
alerting, or migration tool in this foundation.
