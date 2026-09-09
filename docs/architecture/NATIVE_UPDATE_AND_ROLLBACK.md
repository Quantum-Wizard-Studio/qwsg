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

`check` and staging are unprivileged and non-mutating. `update` records current
package identity and Guardian enabled/active intent, verifies compatibility,
stops the Guardian only at the replacement boundary, applies fixed allowlisted
artifacts, reloads the user manager, restores service intent, and validates the
installed identity. An eligible post-mutation failure restores the exact prior
package transaction before returning failure. Configuration, credentials, and
persistent Guardian/operator state are user-owned and are never package backup
payloads.

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

Rollback metadata and prior package artifacts live in a private local root,
are integrity-bound, contain no credentials or user state, and identify the
from/to versions, artifact hashes, transaction state, and captured service
intent. Rollback refuses symlinks, unsafe ownership/modes, incomplete journals,
tampering, foreign destinations, or incompatible persistent state. Retention is
bounded; a newer successful transaction supersedes its predecessor only after
validation.

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
