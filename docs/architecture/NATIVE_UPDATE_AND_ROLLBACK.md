# Native Update and Rollback Architecture

## Boundary

QWSG updates are a two-phase transaction. The installed ordinary-user binary
discovers, downloads, and completely verifies an immutable canonical Release in
a private staging directory. Only a verified plan may cross the narrow
privileged package-artifact boundary. Remote metadata and archive paths never
select privileged destinations.

The canonical operator workflow is:

```text
qwsg update check
qwsg update
qwsg update status
qwsg update rollback
```

`check` and staging do not change installed package artifacts. Explicit update
holds the common nonblocking mutation lease through authentication, Guardian
handoff, helper execution, validation, recovery and final evidence. An equal
online candidate is a no-op; an explicit equal archive remains refused.
Configuration, credentials and persistent Guardian/operator state remain outside
the package write set. Guardian enablement is unchanged.

The success order is PREPARE → MUTATE → VALIDATE → RECOVER GUARDIAN → COMMIT.
The privileged helper reauthenticates the exact source, signed authority and
staged package. The coordinator verifies every installed allowlisted artifact
against the authenticated staged package, then installed identity/configuration
and manager reload, before starting Guardian. Recovery requires active/running,
successful unit result and a nonzero MainPID. Previously inactive Guardian must
remain verified inactive. A successful start command alone is not success.

A failed update always exits nonzero, including `update_failed_recovered`.
Rollback execution, rollback validation and Guardian recovery are separate facts.
The helper's structured apply/rollback receipt is retained; an unknown receipt
cannot prove absence of mutation. Failed rollback or validation blocks restart.
`recovery_failed` means package restoration validated but runtime recovery failed.
`rollback_failed`, `rollback_validation_failed`, `recovery_incomplete` and
`evidence_failed` require operator intervention. Mutation contention retains
exit 3; other refusals/failures retain exit 1. Only fully verified success exits 0.

## Trust model

Task 082 uses authenticated release-index discovery for both check and install.
The production Ed25519 anchor authenticates the exact artifact digest, size,
source commit and compatibility declaration. Local archives require the same
signed authority. The privileged helper independently reauthenticates and
binds package provenance before mutation. The compiled capability controls
executable behavior; metadata cannot supply migration code. See
[Authenticated migration authority](AUTHENTICATED_MIGRATION_AUTHORITY.md) for
the `/2` contract, bootstrap boundary, offline archive procedure and acceptance.

Task 077 adds `internal/updateawareness`, a separate integrity-checked private
record for explicit read-only check transitions and network-free status. It
does not own or modify acquisition, installation, migration, transaction or
rollback state. See `docs/architecture/UPDATE_AWARENESS_STATE.md`.

## Version and migration model

Version ordering follows SemVer: major, minor, patch, then prerelease ordering;
a final version is newer than its prereleases. Normal update accepts only a
newer supported target and refuses equal, older, malformed, ambiguous, and
unsupported-major identities.

Historical `/1` migration plans are explicit source/target records. `/2` plans
are derived from authenticated exact-source declarations and locally implemented
`preserve-package-v1`; no future target is compiled into the client.
Each plan is deterministic, validates before mutation, journals completion, and
defines rollback behavior. The 1.1.0 to 1.2.0-rc.1 path preserves the existing
Configuration Source 1.0, Guardian Checkpoint 1.0, Scheduler State 1.0, and
Current Operator State 1.0–1.2 without mutation; its migration is an explicit
no-op compatibility decision. Unknown schemas fail closed.

## Rollback state

Before the first destination change, Apply copies and fsyncs the complete
allowlisted rollback set, verifies all copies, atomically writes/fsyncs the
`Prepared` transaction journal, and syncs backup directories and their ancestors.
Replacement files and destination directories are synced. A prepared journal
is rollback-capable even when application was interrupted before `Complete`.
Legacy complete journals remain readable; unprepared/incomplete journals refuse.
Package completion is not operator-visible transaction success.

Restore validates the entire source set, hashes, paths and types before changing
any destination. It then restores and validates all recorded bytes/modes or
required absence. Journal reads are bounded and reject unsafe file types.
The root helper retains the backup after rollback, including through validation
or Guardian failures, so repeated recovery remains possible. Only a fully
successful later update may discard the previous successful update's backup.
There is no new automatic retention policy; P4 is unchanged.

Private `update/update-result.json` records source/target, backup, phase, apply
receipt, package validation, rollback execution/validation, Guardian intent/result
and intervention requirement. Intent is durable before stop and mutation. An
interruption leaves incomplete evidence, which blocks another manual update.
`manual-attempt.json` records pre-mutation refusal/no-op phase and exit result;
`rollback-result.json` independently records explicit rollback. Atomic replacement
uses file fsync, rename and directory fsync; a write/sync failure cannot produce a
successful command. `current.json` becomes the normal rollback pointer only after
validation and verified Guardian recovery. Pending evidence supplies the pointer
if interrupted before that point. No autonomous crash replay is introduced.

`qwsg update status` displays update and rollback evidence without network access.
After an interruption, first ensure the original coordinator and any privileged
helper have exited; preserve private evidence and backup files. Then use
`qwsg update rollback`. It revalidates the source and installation, preserves the
original Guardian intent across retries, and verifies recovery independently.
An interruption before mutation can recover Guardian without package rollback.
Missing/corrupt rollback material fails closed; do not manually start Guardian or
delete evidence to bypass the refusal. Correct the reported storage/service
failure or restore an independently verified package through the supported
operator procedure before retrying. No arbitrary-disk-repair or power-loss
survival beyond supported local filesystem fsync guarantees is claimed.

## Task 081 compatibility prerequisite

The explicit `compat-1.2.0-to-1.3.0` route preserves the same configuration,
credential and persistent-state schemas. It permits no other new source/target
pair. The authenticated index must advertise this exact route, and the local
installed-package classifier must independently verify the installed identity.

For the first 1.3.0 acceptance, the Owner authorized this single route as a
compatibility backport to the production Task 080-based 1.2.0 instance. Its
binary and RELEASE.json share the remediation source commit and build date.
This installed compatibility baseline is distinct from the immutable official
1.2.0 archive, tag and Release, which remain unchanged. It is superseded by the
canonical 1.3.0 installation. No alternate public 1.2.0 release is created.
