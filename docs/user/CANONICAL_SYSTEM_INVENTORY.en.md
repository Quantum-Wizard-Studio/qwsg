# Canonical System Inventory

## Overview

Canonical System Inventory is QWSG's read-only description of the current Linux host. It records observed host, operating-system, kernel, CPU, memory, storage, filesystem, network, and virtualization facts. It does not score health, apply policy, monitor continuously, or change the server.

## Prerequisites and authority

Run QWSG as an ordinary user on Linux. No root privilege or network access is required. Some optional evidence can be unavailable because of host permissions or isolation; QWSG reports that state instead of elevating privilege.

## Usage

Build with `make build`, then run:

```bash
build/qwsg inventory
```

The command writes JSON to standard output. Exit code `0` means complete, `2` means partial but usable, and `1` means failed. The top-level Inventory 1.0 view remains available for compatibility. New integrations should consume `canonical_inventory`.

## Privacy and interpretation

Hostnames, network and hardware addresses, interface names, mount paths, raw device names, machine IDs, and service identities are omitted, redacted, or replaced by privacy-safe identifiers. A missing or redacted value does not mean empty, false, or healthy. Check each collector and layer status plus structured issues.

## Configuration and budgets

Collection remains one-shot and has finite per-collector timeouts and output
limits. To persist one validated result explicitly, provide an absolute private
directory:

```bash
build/qwsg inventory save --store /absolute/private/qwsg-inventory
build/qwsg inventory load --store /absolute/private/qwsg-inventory
```

The store defaults to retaining the latest 10 snapshots. Set a fixed value from
1 through 1000 with `--retention N` on creation and use the same value when
opening it. Directories must be private (`0700`), and files are written as
`0600`. Saving and loading emit the same Inventory JSON and preserve status
exit codes, including `2` for a partial but usable snapshot.

Persistence is manual. It introduces no monitoring, comparison, health
scoring, daemon, scheduler, alert, notification, network service, database, or
upload.

## Troubleshooting and upgrades

`permission_denied`, `unavailable`, `unsupported`, `timeout`, `cancelled`,
`resource_limit`, and `error` describe evidence collection, not server health.
Correct the host access boundary only when appropriate; never run as root
merely to hide a partial result.

Stored data is rejected on unsafe permissions or paths, unsupported versions,
malformed JSON, duplicate keys, integrity mismatch, or Inventory validation
failure. The checksum detects corruption but is not a cryptographic signature.
Do not manually delete stale lock or transaction files without reviewing the
store. Schema migration and authenticated storage are future capabilities.
