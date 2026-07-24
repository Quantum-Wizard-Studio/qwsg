# QWSG User CLI Demonstration

This walkthrough covers the Task 017 CLI installation and the Task 018
Snapshot Comparison Engine acceptance sequence on Ubuntu 24.04.

## Build and isolated install

```bash
make build
stage="$(mktemp -d)"
make DESTDIR="$stage" PREFIX=/usr/local install
export PATH="$stage/usr/local/bin:$PATH"
```

The equivalent reviewed system installation is `make build` as the normal user,
followed by `sudo make install`. The install target copies the existing
artifact and never invokes Go in the privileged environment. It may be used
only after the exact `/usr/local/bin/qwsg` target has been inspected and a
bounded rollback is available.

## Explicit session configuration

```bash
demo_root="$(mktemp -d)"
export QWSG_STORE="$demo_root/inventory-store"
export QWSG_FORMAT=human
```

This explicitly selects a private temporary store while keeping the required
application commands concise. It is not automatic store discovery.

## Acceptance commands

```bash
qwsg version
qwsg help
qwsg inventory
qwsg inventory save
qwsg inventory list
qwsg inventory info
qwsg inventory load
qwsg compare
```

Expected observations:

- version prints QWSG version and controlled build metadata;
- help shows the command hierarchy and the JSON compatibility default;
- inventory prints a human summary and exits `0` or truthfully exits `2`;
- save performs one collection and creates one validated snapshot;
- list shows the snapshot name, completion time, schema, and status;
- info reports validated metadata and `Integrity: verified`;
- load reads the stored observation without collecting and clearly labels it.

For compatibility evidence:

```bash
QWSG_FORMAT=json qwsg inventory > inventory.json
QWSG_FORMAT=json qwsg inventory load > stored-inventory.json
```

Both JSON documents retain Inventory 1.0 plus `canonical_inventory`. A partial
but usable Inventory exits `2`; this is an accepted result, not command failure.

The temporary demonstration store may be removed only after confirming it was
created by this walkthrough and is not operator data.
## Snapshot comparison acceptance

Ensure the store contains at least two observations:

```sh
QWSG_FORMAT=json qwsg inventory save >/tmp/qwsg-observation-1.json
QWSG_FORMAT=json qwsg inventory save >/tmp/qwsg-observation-2.json
qwsg inventory list
```

Compare previous versus latest twice and verify deterministic JSON:

```sh
QWSG_FORMAT=json qwsg compare > /tmp/qwsg-compare-1.json
QWSG_FORMAT=json qwsg compare > /tmp/qwsg-compare-2.json
cmp /tmp/qwsg-compare-1.json /tmp/qwsg-compare-2.json
```

Copy two exact names from `qwsg inventory list`, oldest first:

```sh
QWSG_FORMAT=json qwsg compare --from SNAPSHOT_1.json --to SNAPSHOT_2.json
QWSG_FORMAT=human qwsg compare --from SNAPSHOT_1.json --to SNAPSHOT_2.json
```

Expected results:

- JSON has `schema_name: "qwsg.comparison"` and version `1.0`;
- repeat output for the same pair is byte-identical;
- human output groups Added, Removed, Modified, and Unchanged;
- identical semantic state has zero Added, Removed, and Modified counts;
- every displayed item is derived from a canonical Change Record;
- output contains no health, drift, alert, score, or recommendation judgement.

The comparison step is read-only. It needs no root privilege, compiler,
network access, daemon, scheduler, database, or collector modification.
