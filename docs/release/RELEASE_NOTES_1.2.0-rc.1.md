# QWSG 1.2.0-rc.1 Native Update Foundation

This private acceptance candidate introduces QWSG's native update and rollback
foundation. It discovers immutable public releases from the canonical Forgejo
repository, stages downloads privately, verifies checksum, archive safety,
manifest and release provenance before mutation, and applies only fixed
package destinations through a rollback-capable transaction.

The updater preserves configuration, protected credentials, compatible
Guardian/operator state, and Guardian enablement/active intent. Package
replacement failures restore prior artifacts automatically where deterministic;
`qwsg update rollback` restores the integrity-verified prior package explicitly.
The 1.1.0 to 1.2.0-rc.1 migration is an explicit compatible no-op for existing
Configuration Source, Guardian Checkpoint, Scheduler State, and Current Operator
State contracts.

QWSG 1.1.0 predates the native command. Its one-time bootstrap runs the fully
verified 1.2.0-rc.1 archive binary against the preserved 1.1.0 installation.
Subsequent supported releases use the installed `qwsg update` workflow.

This candidate is private and unpublished. It does not change or replace the
official QWSG 1.1.0 Release.
