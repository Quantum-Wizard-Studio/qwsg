# QWSG Upgrade, Rollback, and Uninstall

## Current source and release boundary

This is the authoritative Community operator procedure for accepted current
source through Task 091. [Community 1.0](../PRODUCT_1_0_SCOPE_FREEZE.md) is a
product maturity boundary, not retroactive version numbering. Released 1.3.1
remains immutable and does not contain Tasks 082–091. Current main is not a
published release. C9 requires a distinct new signed release and real supported
1.3.1 → next-release upgrade/bootstrap acceptance, including preservation,
rollback and Guardian recovery. That acceptance has not happened; no next
version or ready-to-run bootstrap command is selected here.

The old binary cannot acquire a new metadata parser through documentation or a
signed declaration. Obtain the separately authorized, authenticated bootstrap
procedure supplied with the future release. Do not bypass verification, edit
migration authority, or apply historical commands below to an unknown target.

## Explicit Community update

```sh
qwsg update check
qwsg update
qwsg update status
qwsg update rollback
```

When lifecycle notification is enabled and ready, update reports the previous
and resulting version and rollback reports the installed and restored version.
Operation success/failure and administrator-notification delivery are recorded
separately; SMTP failure never rewrites the package transaction result.

For clients containing Task 082, update discovery uses the authenticated
production release index. The signed declaration must match the installed
source and a compiled migration capability. Artifact name, size, SHA-256,
platform and source commit are bound to that authority, then normal archive,
manifest and package checks run before stopping Guardian or requesting
privilege. Configuration, credentials and persistent state remain outside
package replacement. See
[Authenticated migration authority](../architecture/AUTHENTICATED_MIGRATION_AUTHORITY.md).

For an explicit offline update, place `FILE.sha256` and the full production-signed
`FILE.release-index.json` beside the archive, then run:

```sh
qwsg update --archive FILE --version VERSION
```

The signed index must select `VERSION` and authorize the installed source and
known capability. Missing authority, a forged checksum, unsupported constraints
or any identity mismatch refuses. Awareness
`update_available_unsupported_source` means the release cannot use the client's
known capability; obtain an explicitly supported upgrade procedure.

QWSG 1.3's Task 076 foundation defines a source-neutral signed release index.
Task 077 makes `update check` one bounded authenticated awareness refresh and
`update status` a network-free local state read. Task 078 activates the exact
production endpoint and trust anchor, and Task 079 lets Guardian call that same
authenticated core when its 24-hour awareness interval is due. Neither manual
nor automatic awareness acquires or installs an artifact; explicit `qwsg
update` remains the only Community installation entry point. The Pro automatic foundation uses the same
authenticated engine but production Pro authority is not available.

## Deterministic compatibility and migration contract

Before any Guardian stop or package mutation, the updater classifies the
installed package from its complete safe layout, strict installed
`RELEASE.json` and exactly matching embedded binary identity. A binary or its
version output alone is unverified. The candidate's independently verified
`RELEASE.json` target must be newer, and one exact declared migration record
must exist. Task 082 accepts a signed exact-source declaration for a locally
known capability without a compiled future target. Historical `/1` metadata
still requires the local route. The record identifies its source/target pair,
configuration/Guardian/scheduler/operator-state schemas, whether schema
mutation is required, preservation rules, and the managed Guardian-unit
replacement boundary. Missing, malformed, partial, inconsistent, equal, older,
major-incompatible or undeclared identities fail closed; QWSG never guesses a
path or silently overwrites unknown artifacts.

QWSG 1.3.0 declares `compat-1.2.0-to-1.3.0`. Both sides use compatible Configuration 1.0, Guardian 1.0, Scheduler 1.0 and Operator State 1.0–1.2 contracts, so this path performs no configuration or state-schema transformation. Existing user configuration, protected notification credentials and persistent QWSG state remain byte-preserved outside package destinations. Only the verified binary, Guardian user unit and release-owned documentation are replaced.

## Validation and Guardian intent

Preflight requires authenticated source/target/migration authority, verified
package integrity/provenance and valid installed configuration. A common
nonblocking mutation lease excludes conflicting manual update, automatic work
and explicit rollback through handoff, helper execution, validation, recovery
and final evidence. Contention refuses without mutation (exit 3); it does not
queue or grant authority. The privileged helper independently reauthenticates.

Before any destination change, QWSG preserves and verifies the complete package
before-image set and durably records its prepared rollback journal. Success
requires installed package bytes/modes and identity/configuration validation,
then user-manager reload and verified Guardian recovery, then committed local
evidence. Guardian start alone proves neither update success nor long-term
server health. Run `qwsg config validate` and `qwsg readiness` afterward for
current configuration/readiness; optional SMTP can still be PARTIAL.

The previous Guardian running/stopped intent is preserved. A previously active
Guardian is stopped for mutation and must recover to verified active/running
state; a previously inactive Guardian stays inactive. Update and rollback do
not change enablement policy. Package validation and runtime recovery are
separate facts; failed package recovery never authorizes starting Guardian.

## Rollback and interrupted recovery

Inspect `qwsg update status` first. It reads local awareness and transaction
results without networking. Update success, rollback execution/validation and
Guardian recovery must be read separately:

| Observed outcome | Meaning and next action |
| --- | --- |
| `success` | The requested transaction validated and required Guardian intent was verified. Check current readiness; success is not a universal server-health certificate. |
| `update_failed_recovered` | Update failed and exits nonzero even though the prior package and Guardian intent recovered. Review the original failure before choosing another explicit update; do not report the attempted target as installed successfully. |
| `rollback_failed` / `rollback_validation_failed` | Restoration failed or did not validate. Preserve evidence and backup, leave Guardian stopped, correct the reported cause or obtain an independently verified recovery source, then retry explicit rollback. |
| `recovery_failed` | Package rollback validated but Guardian recovery failed. Preserve the backup and original intent, resolve the user-manager/runtime finding and retry `qwsg update rollback`; do not repeat package installation merely to hide this result. |
| Incomplete, ambiguous or invalid evidence; `recovery_incomplete` / `evidence_failed` | No completed recovery is proven. Preserve private evidence and installation, establish that prior processes have exited and inspect the reported cause before explicit recovery. Another manual update is blocked by unresolved transaction evidence. |

Before recovery after an interrupted privileged operation, the operator MUST
ensure that the previous privileged helper is no longer running, as well as
the coordinator. A lost terminal or coordinator exit is not proof of helper
exit. If this cannot be established, stop and obtain administrator assistance;
do not race recovery against it. QWSG does not automate privileged-process
termination or crash replay.

Preserve a private copy of local update evidence and the referenced rollback
backup. Once the previous operation is inactive and the cause is understood,
run `qwsg update rollback`. It can use the prepared transaction after interrupted
mutation even if normal completion was not recorded. Interrupted pre-mutation
handoff can recover original Guardian intent without a nonexistent package
backup. Recovery retry preserves original intent, including after a failed
Guardian start. Configuration, credentials and persistent operator state are
outside the package replacement/rollback write set.

Rollback prevalidates the entire backup source before restoring package-owned
artifacts and verifies restored bytes/modes or absence. Missing/corrupt sources
fail closed and require a separately verified recovery source; arbitrary state
or filesystem corruption is not automatically repairable. Backups remain
available through failed validation/recovery. A consumed successful explicit
rollback cannot silently select a different prior transaction on repetition.
Never delete evidence, guess a backup or manually start Guardian to bypass
failed package validation. If the restored binary rejects retained state,
leave it stopped and retain the data for review.

See [native transaction details](../architecture/NATIVE_UPDATE_AND_ROLLBACK.md).
SMTP notification delivery is a separate result and cannot turn a failed
transaction into success or a successful transaction into failure.

## Local observation evidence recovery

Inventory snapshots and Guardian checkpoints are separate from update rollback
backups. Recognized interrupted snapshot writes/retention recover deterministically
under writer exclusion; unknown, conflicting, unsafe or corrupt evidence is
preserved and rejected explicitly. For a legacy inventory-lock error, verify
all older writers have stopped and preserve a private store copy before moving
only the legacy lock outside the store. Never remove the permanent writer lock
to bypass an active writer. Follow the exact
[local-evidence recovery contract](../architecture/INVENTORY_PERSISTENCE_AND_DIGITAL_TWIN.md#interruption-recovery-task-090-c3)
for classification; do not reset the store or guess an authoritative snapshot.
These guarantees require the supported local filesystem's ownership, locking,
atomic rename and file/directory fsync semantics. They do not promise recovery
from arbitrary hardware or filesystem corruption.

## Historical compatibility and installation replacement

The following version-specific examples retain their historical scope. They
are not the pending C9 bootstrap procedure or commands for a Task 082 client,
which requires signed companion authority for offline update.

QWSG 1.1.0 has no native update command. For the single transition from 1.1.0,
run the verified newer archive binary with its own archive identity:

```sh
./bin/qwsg update --archive /absolute/path/qwsg-1.2.0-rc.1-linux-amd64.tar.gz --version 1.2.0-rc.1
```

The matching `.sha256` file must be adjacent. This private-candidate form is for
controlled acceptance; ordinary published updates use `qwsg update`.

Stop the exact user unit before replacing artifacts. Verify the new archive, preserve a private backup of the old binary/unit and state, then use `./install.sh --replace --backup-dir ABSOLUTE_NEW_DIRECTORY`. Reload the user manager and start only if it was previously active.

QWSG 1.3.0 reads Current Operator State 1.0, 1.1 and 1.2, Scheduler State 1.0, Guardian Checkpoint 1.0 and Configuration Source 1.0. Unknown, corrupt, wrong-mode, wrong-owner, symlinked, or incompatible state fails closed and is not migrated or deleted.

Run `qwsg config validate` before and after upgrade. Source Record 1.0 remains
strict and is not silently migrated. Setup preserves unspecified valid values.
Back up configuration separately from runtime state.

Run `qwsg install --check` against a new archive before replacement and
`qwsg readiness` after upgrade or rollback. Neither command executes
remediation or changes a service.

Current native rollback covers the complete recorded package-owned write set,
including release-owned documentation; use the recovery procedure above.

Before uninstall, explicitly run `systemctl --user disable --now qwsg-guardian.service` and remove only the copied per-user unit. Run the matching verified release archive's `sudo ./uninstall.sh`; it refuses modified artifacts. Configuration and private state are preserved. QWSG 1.0 provides no automatic purge command.

The first 1.3.0 production acceptance uses the explicitly Owner-authorized
compatibility-remediated installed 1.2.0 baseline. The historical official
1.2.0 artifact remains unchanged. Check Scheduler state size before upgrade:
1.3.0 rejects envelopes over 8 MiB before decoding and never deletes them.
Keep an exact private backup and resolve oversized legacy state before healthy
scheduling acceptance. Supported current state remains byte-preserved by the
package transaction; subsequent normal Scheduler execution retains 64 results.

## Historical corrective 1.3.0 to 1.3.1 update

This historical procedure applies to the immutable 1.3.1 binary, which predates
Task 082. It does not describe the new signed-companion requirement. Use the
verified 1.3.1 archive binary to orchestrate the existing transaction:

```sh
./bin/qwsg update --archive /absolute/path/qwsg-1.3.1-linux-amd64.tar.gz --version 1.3.1
```

The matching checksum sidecar must be adjacent. The installed historical 1.3.0
updater has no compiled route to 1.3.1; the candidate updater independently
classifies installed 1.3.0 and permits exactly `compat-1.3.0-to-1.3.1`. Its
privileged helper repeats verification. Rollback restores the exact installed
1.3.0 package. No configuration or state-schema conversion is required.

After replacement, `qwsg update check` obtains a fresh authenticated evaluation
when installed identity changed or the cached relation is stale. The successful
update-notification identity is preserved, including if this refresh initially
fails. Failed refresh reports unknown until a valid evaluation is available,
while retaining historical authenticity/anti-rollback evidence.
