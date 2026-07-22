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

Task 014 introduces no user configuration or collector-selection interface. Collection is one-shot, has finite per-collector timeouts and output limits, and performs no historical storage. Scheduling and persistence are separate future capabilities.

## Troubleshooting and upgrades

`permission_denied`, `unavailable`, `unsupported`, `timeout`, `cancelled`, `resource_limit`, and `error` describe evidence collection, not server health. Correct the host access boundary only when appropriate; never run as root merely to hide a partial result. Schema upgrades follow explicit versioned migrations. QWSG does not remove retained inventory because this version does not persist it.
