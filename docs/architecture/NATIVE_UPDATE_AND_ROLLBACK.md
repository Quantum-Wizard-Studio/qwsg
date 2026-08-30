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

Discovery is restricted to the public canonical Forgejo repository and exact
version tags. QWSG accepts only strict semantic identities supported by its
update policy and immutable linux-amd64 archive/sidecar names. HTTPS retrieval
is bounded by time, redirects, response size, and private staging permissions.
Before installation QWSG verifies the sidecar, archive size and SHA-256, unique
single-root regular-file/directory-only layout, manifest, required files and
modes, binary platform, and embedded version/source provenance. SHA-256 protects
integrity after authoritative HTTPS discovery; signing remains a separate
Owner decision.

Task 076 adds the separate `qwsg.release-index/1` authenticity foundation in
`internal/releasediscovery`: strict metadata, deterministic Ed25519 signature
bytes, explicit trust anchors, and bounded source-neutral HTTPS retrieval. It
does not activate a production endpoint or key and does not replace this
updater. Artifact acquisition, package verification, privileged transaction,
migration execution and rollback remain here. See
`docs/architecture/RELEASE_INDEX_AND_SOURCE_CONTRACT.md`.

## Version and migration model

Version ordering follows SemVer: major, minor, patch, then prerelease ordering;
a final version is newer than its prereleases. Normal update accepts only a
newer supported target and refuses equal, older, malformed, ambiguous, and
unsupported-major identities.

Migration plans are explicit `(from schema/version, to schema/version)` records.
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
