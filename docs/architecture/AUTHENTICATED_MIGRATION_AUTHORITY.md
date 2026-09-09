# Authenticated migration authority

Task 082 replaces future-target enumeration with authenticated authorization of
one locally implemented capability. Community installation remains an explicit
operator action. No Pro scheduling or automatic installation is implemented.

## Problem and upgrade boundary

An immutable 1.3.0 client cannot contain a migration entry for a release that
had not been designed when it was built. Adding another version pair fixes
only that pair. The new client understands a stable package replacement
mechanism; a future signed release declares which exact installed sources may
use it. The future target is data, not a compiled migration-table entry.

This cannot retrofit new parsing code into the already published 1.3.0 or
1.3.1 binaries. Those immutable binaries retain their historical behavior.
A separately authorized installation of a client containing Task 082 is the
bootstrap prerequisite. Thereafter future compatible releases need no client
backport. Task 082 neither publishes a bootstrap release nor alters historical
release material.

## Contract: release-index/2 and migration/1

The existing endpoint, media type, Ed25519 key, canonical typed JSON encoding,
size/depth/count limits and offline custody workflow remain unchanged.
`qwsg.release-index/1` remains supported with exactly its original signing
bytes and local historical route checks. `qwsg.release-index/2` requires a
nonempty `compatibility` array on every release and an empty `migration_routes`
array. Mixing route and capability authority is invalid.

A release contains at most 32 declarations. Each declaration contains exactly
these fields, in this canonical signing order:

```json
{
  "schema": "qwsg.migration/1",
  "source_version": "1.3.1",
  "target_version": "1.8.0",
  "platform": "linux-amd64",
  "capability": "preserve-package-v1",
  "configuration_schema": "1.0",
  "guardian_schema": "1.0",
  "scheduler_schema": "1.0",
  "operator_state": "1.0-1.2"
}
```

These are illustrative future-release values, not a publication claim. The
release object field order is `version`, `published_at`, `status`,
`source_commit`, `release_notes_url`, `minimum_source_version`, `compatibility`,
`migration_routes`, `artifacts`. `/1` omits `compatibility`, preserving its
historical byte representation. All arrays retain input order. Generation
validates before producing compact UTF-8 JSON without a final newline; the
signature excludes only the root `signatures` field.

Source selectors are exact canonical major-1 versions, at least the containing
release's minimum and strictly older than its target. There are no implicit
ranges, precedence rules, wildcards, downgrade policies or migration chains.
Target and platform must match the containing release/artifact. Duplicate
source/platform selectors are rejected even when their content is identical;
conflicting declarations cannot be resolved by ordering. Unknown fields and
migration schemas refuse. Unknown well-formed capability/constraint values
remain visible as unsupported and cannot authorize installation.

The declaration is signed together with the target version, source commit,
artifact name, platform, byte length, URL and SHA-256. It cannot be detached
and reused to authorize a different release or artifact.

## Executable authority and subsystem boundaries

`internal/update` implements `preserve-package-v1`. This is the existing fixed
package replacement transaction: preserve Configuration 1.0, Guardian 1.0,
Scheduler 1.0 and Operator State 1.0–1.2; replace the binary, managed Guardian
unit and allowlisted release documentation. Configuration, credentials and
private state are outside its write set. It performs no schema conversion.

The authority cannot supply scripts, shell fragments, commands, parameters,
destination paths, extra migration binaries or new executable mechanisms.
Package install/uninstall scripts are verified package content and are never
invoked by this migration. A future schema transformation needs new local
implementation and a separate capability; old clients refuse it.

The signed package still contains the ordinary target QWSG binary and service
unit: trusting those release artifacts is the existing software-distribution
boundary. The migration declaration grants no additional executable authority.

- `installation` verifies installed layout and matching binary/provenance.
  Its target-specific historical classification remains unchanged. A future
  target cannot be locally classified as upgrade-supported without authority.
- `releasediscovery` remains read-only. `/2` evaluation verifies installed
  identity once and checks the selected declaration against local capability
  constraints; `/1` retains the historical classifier/route conjunction.
- `updateauthority` binds authenticated discovery to staging/package
  verification. Its `Candidate` has private authority fields. A mutable
  awareness record or caller-created `Evaluation` cannot construct one.
- The ordinary-user CLI performs preflight and service orchestration. Its root
  helper re-stages into private root-owned storage, reauthenticates the signed
  index, reclassifies the actual source, and repeats capability, digest and
  provenance checks before calling the existing transaction.

`Authorize` and `Candidate.VerifyStaged` provide the common semantic foundation
for a later policy-controlled Pro caller. They do not grant permission to skip
configuration preflight, service handling, privileged re-verification,
post-validation or rollback. No duplicate Community/Pro engine is introduced.

## Operator flow and failures

`qwsg update check` authenticates/evaluates metadata and records awareness only.
`qwsg update` fetches a fresh unconditional signed index from the production
endpoint, selects the greatest active stable release, validates authority,
downloads its exact signed artifact, and verifies its manifest/provenance.
It does not fall back to unsigned Forgejo discovery. Equal/older releases
produce a no-update result; unsupported newer releases refuse.

The generated-at clock-skew bound is 15 minutes, matching awareness. A known
validated awareness generation is also a lower bound for installation: older
metadata refuses. The awareness cache can restrict authority but never grant
it. There is no invented expiry for historical signed release material.

`qwsg update --archive FILE --version VERSION` requires adjacent
`FILE.sha256` and `FILE.release-index.json`. The latter is the full signed index
under the production key; its greatest eligible stable target must match
`VERSION`. This route works offline and repeats exactly the same capability,
artifact and provenance checks. A matching archive/checksum pair alone is
insufficient. There are no production test-key, trust-anchor or host-root
override flags. Private test-linked fixtures are excluded from shipped builds.

Unknown capability/constraints yield awareness
`update_available_unsupported_source` and the update error
`authenticated compatible migration capability unavailable`. Operators should
obtain an explicitly supported bootstrap/upgrade procedure; they must not edit
metadata or bypass verification. Malformed, missing, unauthenticated,
conflicting or mismatched authority fails before package mutation. The helper
also refuses changed authority or source after unprivileged validation.

## Rollback and acceptance

The existing deterministic transaction records prior bytes, modes and hashes
for managed destinations. Post-validation failure attempts package rollback and
restores prior service intent. `qwsg update rollback` restores the recorded
predecessor. State/configuration remain outside the package transaction; normal
runtime compatibility checks still fail closed for corrupt or unknown state.

`make forward-update-check` is a prerequisite of `make release-check`.
`TestOldBinaryForwardAuthenticatedUpdate` builds/freezes an old test-linked
client before selecting the future target, proves its historical migration
registry has no such route, then builds and signs a future package. Separate
processes invoke the old client's real `update check`, `update`, archive,
privileged helper and rollback dispatch. Only host boundaries (transport,
service manager, root location and elevation) are isolated; authentication,
classification, planning, staging, package verification, configuration
preflight, apply and rollback remain real code. The future binary is used only
as target payload and for normal post-install identity/configuration checks.

The gate covers successful network/archive transactions, discovery without
installation, rollback after injected post-apply failure, missing/altered
signatures, unsupported capability, source/provenance/digest mismatch, forged
local checksum and independent helper reauthentication. Unit tests add schema,
platform, target, downgrade, duplicate/conflicting authority and historical
migration regression. This deterministic gate does not claim a real production
systemd/sudo acceptance or publication.
