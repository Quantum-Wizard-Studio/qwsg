# QWSG User CLI Demonstration

This walkthrough is the Task 017 Ubuntu 24.04 acceptance sequence.

## Build and isolated install

```bash
make build
stage="$(mktemp -d)"
make DESTDIR="$stage" PREFIX=/usr/local install
export PATH="$stage/usr/local/bin:$PATH"
```

The equivalent reviewed system installation is `sudo make install`; it may be
used only after the exact `/usr/local/bin/qwsg` target has been inspected and a
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
