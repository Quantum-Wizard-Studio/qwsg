# Inventory Persistence and Digital Twin Foundation

## Status and authority

This document defines the Task 016 file-backed persistence profile for the
canonical Inventory Architecture in `ai/core/12_INVENTORY_ARCHITECTURE.md`.
Inventory remains observed evidence. Persistence does not turn it into desired
state, a health verdict, monitoring, or authority to mutate a host.

## Boundary

The flow is:

```text
collectors -> assemble -> redact -> validate -> Inventory Store -> validate -> consumer
```

`internal/inventorystore` imports the Inventory model. Collectors do not import
or invoke persistence. A store accepts only an object envelope that passes
`inventory.Validate`, contains canonical inventory, and contains no
`secret_prohibited` fact in either compatibility or canonical views.

Persistence is invoked explicitly and remains one-shot. There is no daemon,
scheduler, polling loop, listener, database, upload, comparison, policy,
health, alert, or notification behavior.

## Storage layout

An operator supplies a clean absolute store root:

```text
<store-root>/                         mode 0700
├── store.json                        mode 0600
└── snapshots/                        mode 0700
    └── <UTC>_<snapshot-hash>.json    mode 0600
```

The root, every existing path component, metadata file, snapshot directory, and
snapshot file must be a real safe file type. Symlink roots or components,
group/other-accessible store objects, traversal names, unknown files, hidden
transaction artifacts, and inconsistent metadata fail closed.

`store.json` records `format_name: qwsg.digital-twin`, `format_version: 1.0`,
and the immutable retention limit. Snapshot names use completion time with
nanosecond precision and the first 64 bits of SHA-256 over the opaque snapshot
ID. They disclose neither raw subject identity nor the snapshot ID.

## Persisted Digital Twin envelope

Each UTF-8 JSON snapshot contains persistence format identity, UTC creation
time, snapshot and privacy-safe subject identity, Inventory schema version,
SHA-256 of the deterministic compact JSON encoding of the embedded payload, and
the complete Inventory 1.0 envelope with its synchronized
`canonical_inventory`.

The checksum detects payload modification. It is not a digital signature,
message authentication code, provenance proof, or defense against an attacker
who can rewrite both payload and checksum. Authentication and signing require a
separate approved key-management design.

Loading rejects malformed JSON, duplicate object keys, trailing values, unknown
fields, unsupported versions, envelope/payload identity differences, integrity
mismatch, invalid Inventory, unsafe permissions, unsafe file types, and
traversal. Stored data is never silently repaired or migrated.

## Atomicity and locking

Save validates before creating store data. It writes a restrictive
same-directory temporary file, flushes and closes it, installs it with an atomic
no-clobber hard link, synchronizes the directory, and removes the temporary
name. A store-local exclusive lock rejects concurrent writers.

Retention temporarily renames the oldest snapshot. Failed new installation
restores it. Failed retirement removes the new snapshot and restores the old
one where the bounded filesystem transaction permits. Every rollback failure
is reported rather than hidden.

The profile requires a local filesystem supporting hard links, atomic
same-filesystem rename, file synchronization, and directory synchronization.
Unsupported filesystems fail explicitly; no weaker fallback is used.

## Retention

Retention is fixed when the store is created, ranges from 1 through 1000, and
defaults to 10 for the CLI. Opening a store with a different value fails. A
successful save keeps at most that number of visible snapshots, removes only
the deterministically oldest snapshot, and always retains at least one valid
snapshot. There is no timed cleanup or background maintenance.

Changing retention, stale-lock recovery, migration, export, and operator-data
deletion are future explicit operations. QWSG does not guess that
abandoned-looking state is safe to delete.

## Compatibility and CLI

`qwsg inventory` retains its original one-shot stdout JSON and exit semantics.
Persistence adds:

```text
qwsg inventory save --store /absolute/private/path [--retention N]
qwsg inventory load --store /absolute/private/path [--retention N]
qwsg inventory load --store /absolute/private/path [--retention N] --snapshot <name>
```

Save collects once, persists, emits the same Inventory JSON, and returns its
established status code. Load performs no collection, revalidates and emits the
stored Inventory, and returns its status code. A partial usable snapshot remains
exit code `2`; persistence does not relabel it complete.

## Recovery and deferred work

Store errors do not authorize deletion. Preserve the directory and record safe
error evidence. A stale `.write.lock` or `.retire-*` artifact requires review
and is not removed automatically. Task rollback removes only temporary test
fixtures and never an operator-selected store.

Comparison, drift analysis, health evaluation, monitoring, scheduling, alerts,
notifications, API/Console access, databases, signing, encryption-at-rest,
migrations, and remote synchronization are deferred.
