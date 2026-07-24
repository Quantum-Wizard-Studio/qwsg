# Task 018 Delivery Audit — Snapshot Comparison Engine

## Scope delivered

- Dedicated `internal/comparison` architectural layer.
- Canonical deterministic `qwsg.comparison` 1.0 result and Change Records.
- Previous/latest and explicit snapshot-pair CLI selection.
- Canonical JSON plus terminal-safe human grouping derived only from records.
- Architecture, schema, developer, user, and Ubuntu 24.04 demonstration docs.

## Contract and security findings

The engine compares canonical layer status, resource presence/kind, and typed
facts for every Inventory layer. It rejects invalid, non-canonical,
cross-subject, or contract-incompatible sources. Stable IDs, evidence-derived
timestamps, sorted unions, and ordered records eliminate runtime randomness.

Inventory 1.0, Canonical Inventory v1, and Inventory Store formats are
unchanged. The workflow remains non-root, read-only, local, dependency-free,
and one-shot. It adds no collection, network communication, daemon, scheduler,
database, host mutation, health judgement, alert, score, recommendation, or
privilege path.

## Verification evidence

Targeted tests cover identical, added, removed, modified, reversed,
metadata-insensitive, escaped-path, invalid-source, immutable-input, default
selection, explicit selection, deterministic JSON, human output, insufficient
history, and invalid selectors. Repository-wide tests, race checks, formatting,
vet, engineering-framework validation, lifecycle validation, and the Ubuntu
24.04 normal-user build and clean private-store acceptance demonstration pass.
Repeated canonical output was byte-identical. Project Owner acceptance and
lifecycle completion remain outstanding.

The post-commit binary was also installed into an isolated staging root with
`PATH=/usr/bin:/bin`; installation succeeded without Go and the installed
artifact was byte-identical. System installation remains an owner-operated
`sudo make install` step because this non-interactive session has no sudo
credential. The currently installed Task 017 binary is not claimed as Task 018.

## Rollback

The pre-implementation rollback snapshot is
`/tmp/qwsg-task018-implementation.20260724Tm7qzQr`. Rollback restores only
Task 018-owned tracked targets from that manifest and removes only Task
018-created absent targets after validation. Operator snapshot stores and
pre-existing untracked repository files are never rollback targets.
