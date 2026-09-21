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
├── .write.flock                      mode 0600, permanent kernel-lock inode
├── .write.lock                       mode 0600, versioned legacy-writer guard
└── snapshots/                        mode 0700
    └── <UTC>_<snapshot-hash>.json    mode 0600
```

The root, every existing path component, metadata file, snapshot directory, and
snapshot file must be a real safe file type. Symlink roots or components,
group/other-accessible store objects, traversal names, unknown files and
inconsistent metadata fail closed. Recognized interrupted transaction artifacts
follow the bounded recovery contract below.

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

Save validates before creating store data. Under an exclusive nonblocking Linux
`flock`, it writes a private same-directory `.tmp-inventory-<decimal>` file,
flushes and closes it, installs an atomic no-clobber hard link to the final
snapshot name, then synchronizes the directory. **The durable commit boundary
is successful file fsync, final hard link and directory fsync.** A valid final
name observed after interruption is re-synchronized before recovery discards any
retired evidence. A sync failure returns an error and preserves the installed
object; it does not assert durability or remove a possibly committed snapshot.
The temporary name never becomes evidence by itself.

The permanent `.write.flock` inode is never unlinked. The kernel releases its
lock on descriptor close or process exit. Save, List, Load and LoadLatest all
hold this lock through recovery and their operation; a collision returns an
explicit busy error. Readers can therefore complete interrupted housekeeping.
No timeout, process-ID guess or age-based lock theft is used.

A permanent `.write.lock` guard containing `qwsg.inventory-kernel-lock/1` plus
newline is installed atomically under the kernel lock. It prevents older
O_EXCL writers from entering this store. An empty legacy lock or unrecognized
marker has unknowable ownership: preserve it and fail explicitly. Normal new
writer interruption leaves recognizable guard bytes and an unlocked kernel
inode. Existing format 1.0 snapshots/metadata remain readable without rewriting;
older binaries cannot write a converted store without an operator-controlled
quiescent downgrade of the locking state.

The profile requires a local Linux filesystem supporting flock, hard links,
atomic same-filesystem rename, file fsync and directory fsync. Unsupported
operations fail explicitly. This is an interruption contract for the supported
single transaction, not repair of arbitrary filesystem/hardware corruption or
protection from a same-user process modifying store files outside the protocol.

## Retention

Retention is fixed when the store is created, ranges from 1 through 1000, and
defaults to 10 for the CLI. Opening a store with a different value fails. A
successful save keeps at most that number of visible snapshots, removes only
the deterministically oldest snapshot, and always retains at least one valid
snapshot. There is no timed cleanup or background maintenance.

Changing retention, migration, export and operator-data deletion remain
separate explicit operations. Task 090 adds only the recovery below.

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

Task 017 adds a user-facing consumer boundary without changing this store
contract. `inventory list` obtains deterministic names from `Store.List` and
validates every displayed entry through `Store.Load`. `inventory info` and
`inventory load` validate the latest or an explicitly named snapshot through
the same load boundary. Task 090 makes these reads finish recognized interrupted
transactions under lock; they never rewrite snapshot payloads or migrate data.

JSON remains the compatibility default for collect, save, and load. The
separate human renderer exposes status, timestamps, counts, and privacy-safe
metadata without presenting stored evidence as current state or a health
verdict. An operator may explicitly provide the store and format through
command options or the session-scoped `QWSG_STORE` and `QWSG_FORMAT`
environment variables; there is no automatic global store discovery.

Task 018 adds a read-only Comparison Engine above this store. It selects the
previous/latest or an exact pair, loads both through the existing integrity and
Inventory validation path, and emits canonical Change Records. The store format
and Inventory schema remain unchanged.

## Interruption recovery (Task 090, C3)

Recovery follows detect, classify, verify, recover and continue under the lock.
At capacity N, the existing transaction renames the oldest snapshot to
`.retire-<original-name>` and fsyncs the directory before installing its successor.
Before mutating a retirement, recovery checks all visible and retired envelopes,
canonical validity, identity/filename, checksum, private file types and counts.

| Observed state | Deterministic action |
| --- | --- |
| No retirement; recognized private temporary files | Never promote them; synchronize directories and remove abandoned temporary names. |
| One valid retirement, N-1 valid visible snapshots, no original-name conflict | Restore the original name and fsync; the new commit did not appear. |
| One valid retirement, N valid visible snapshots, no original-name conflict | Synchronize observed committed names, remove retirement, fsync; do not revert the new commit because cleanup was interrupted. |
| Multiple retirements, conflicting original, other counts, malformed evidence, unknown/unsafe artifacts | Preserve evidence and return an explicit integrity/recovery error. |
| Missing metadata with existing snapshot artifacts | Refuse to invent retention/format metadata. |

An interrupted cleanup can be retried; repeated recovery does not rewrite valid
payloads. A normal failed pre-install write leaves prior evidence recoverable by
the next Save/List/Load/LoadLatest. Retirement only deletes previously committed
old evidence after the new installation boundary; errors never trigger deletion
of the new object as rollback.

Recovery scans only the store root (at most 6 entries) and its snapshot directory
(at most N+2 entries), using one extra entry for overflow detection. No recursive
scan or history/journal is created. Excess artifacts require explicit review;
they are not erased to satisfy a limit. Snapshot reads consume at most 16 MiB+1,
metadata 4096+1 and lock marker 128+1 bytes; oversize input is rejected. Pre-stat
checks are supplementary, never the only bound. Files are opened without
following symlinks; special files and non-private modes are refused.

For a legacy-lock error, stop and verify all older QWSG writers are inactive,
preserve a private copy of the store, then move the **legacy** `.write.lock`
outside the store and retry the same inventory command. Never remove the
permanent `.write.flock` inode or the recognized versioned guard to bypass an
active writer. Other ambiguous states require offline inspection of the
preserved evidence; do not guess an authoritative object or reset the store.

Task 090 also bounds Guardian checkpoint input to its existing 4 MiB limit plus
one detection byte and checks private regular-file type. Its existing
fsync/atomic-rename/directory-fsync replacement and generation recovery remain
unchanged. Checkpoint temporary names are never loaded as authoritative.
Updater backup/transaction durability and rollback are C7, outside this change.

Independent snapshot comparison, drift analysis, health evaluation, monitoring,
scheduling and notifications have their own contracts. Databases, signing,
encryption-at-rest, migrations and remote synchronization are outside this
file-backed interruption contract.
