# Guided Installer Architecture

QWSG 1.2.0-rc.2 places installation orchestration in the Go product rather
than in an independent shell program. Operators download the release archive
and checksum sidecar, verify the outer checksum and internal manifest, and run
`./bin/qwsg install --guided`. The existing fixed-destination `install.sh` is a
narrow privileged package-copy helper invoked only after read-only preflight,
a visible plan, and consent. Task 064 remains the archive provenance, update
transaction, and rollback contract.

`internal/installer` is presentation-neutral. It owns ordered phases,
deterministic weights, phase states, plans,
update-policy vocabulary, supported-platform detection, stable message IDs,
and catalogs. CLI rendering and process execution are adapters. Future concise,
answer-file, or fleet interfaces must consume this engine rather than
reimplement installation decisions.

Task 075 places installed-package identity in the single
`internal/installation` classifier. The wizard consumes its bounded states and
reasons. A complete safe package whose executable and installed
`qwsg.release/1` provenance agree may skip package copying; an exact declared
migration route is sent to `qwsg update`; legacy, unknown, partial,
inconsistent, or unsafe artifacts stop before mutation. Binary presence alone
never proves installation. See
`docs/architecture/INSTALLED_PACKAGE_CLASSIFICATION.md`.

## Truthful progress

Weights total 100: preflight 12, plan 10, package installation 20,
configuration 15, optional notification 10, update policy 8, activation 12,
readiness 10, and completion 3. Only completed phases contribute. Active,
failed, retried, and restored phases never masquerade as completed work. The
completion view begins at 97 percent and reaches 100 only after completion.

## Localization and terminal behavior

Language selection precedes localized interaction. English is canonical and
the safe fallback; Hungarian and German are complete initial catalogs. Logic
uses stable message identifiers. Catalog validation requires every canonical
identifier in each initial language.

Capable terminals receive a compact redrawn header, progress bar, stage, and
explanation. `TERM=dumb` and explicit `--line-mode` avoid cursor control.
Guided installation requires a terminal; expert assessment remains available
as `install --check --format json`.

## Mutation and failure boundary

Platform inspection and planning are read-only. Unsupported systems stop
before package mutation. Consent precedes `sudo`. Package replacement retains
Task 064 transaction and rollback semantics; configuration, credentials, and
persistent state remain outside package payloads. Failures name their stage and
do not count it complete. Notification is optional. Manual and notify update
policies are durable, but notify is a preference foundation—not a claim that a
background notifier exists. Automatic privileged updates are not enabled.
