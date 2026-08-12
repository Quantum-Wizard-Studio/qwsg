# QWSG User CLI and Snapshot Explorer

## Start here

QWSG collects a one-shot, read-only Linux Inventory and can explicitly save and
browse validated snapshots. It does not monitor, compare, score health, run as
a daemon, or change the host.

```bash
qwsg help
qwsg version
qwsg help inventory
```

For persistent per-user Guardian configuration, run `qwsg setup`, then use
`qwsg config show|validate|get|set`. The canonical reference is
`docs/release/SETUP_AND_CONFIGURATION.md`; setup never starts the service.

Exit `0` means success or complete Inventory, `2` means partial but usable
Inventory, and `1` means a usage, validation, store, permission, corruption, or
runtime failure.

## Live Inventory

JSON remains the compatibility default:

```bash
qwsg inventory
```

It contains the complete Inventory 1.0 envelope and additive
`canonical_inventory`. For an administrator summary:

```bash
qwsg inventory --format human
```

Human output always shows status and observation time. It is not a health
verdict. Inspect JSON for structured collector issues and redactions.

## Choose a snapshot store

The store must be a clean absolute private path. Supply it on every command:

```bash
qwsg inventory save --store /absolute/private/qwsg-inventory
qwsg inventory list --store /absolute/private/qwsg-inventory
```

Or explicitly select it for a shell session:

```bash
export QWSG_STORE=/absolute/private/qwsg-inventory
export QWSG_FORMAT=human
```

The variables do not discover or create an implicit global store. Command-line
`--store` and `--format` values take precedence.

## Save and browse

```bash
qwsg inventory save
qwsg inventory list
qwsg inventory info
qwsg inventory load
```

`save` collects once and persists only a validated privacy-safe Inventory.
`list` validates every displayed snapshot and orders names from oldest to
newest. `info` and `load` use the latest snapshot by default. Select an exact
name returned by list:

```bash
qwsg inventory info --snapshot NAME
qwsg inventory load --snapshot NAME
```

Use `--format json` for machine-readable Inventory or explorer metadata. A
store retention value defaults to 10, is fixed at creation, and must be supplied
consistently with `--retention N` when non-default.

Stored timestamps describe when evidence was observed. List, info, and load do
not collect, repair, migrate, compare, delete, or declare data current.

## Fail-closed behavior

QWSG rejects relative or unsafe paths, symlinks, permissive store objects,
unexpected files, unsupported versions, malformed JSON, duplicate keys,
integrity mismatch, and invalid Inventory. Preserve a failed store for review.
Do not delete locks, transaction artifacts, or operator snapshots blindly.

Optional Community email commands are `qwsg notification preflight`,
`qwsg notification credential set --from-file FILE`, and the explicit
`qwsg notification test`. Community supports exactly one administrator
recipient; credential values are never accepted as command arguments or shown.
