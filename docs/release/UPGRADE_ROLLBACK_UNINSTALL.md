# QWSG Upgrade, Rollback, and Uninstall

## Native workflow (QWSG 1.2 and later)

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

Update discovery and download use the anonymous canonical Forgejo Release
source. QWSG verifies the sidecar, archive layout, manifest, required package
files, platform and embedded `RELEASE.json` provenance before stopping the
Guardian or requesting privilege. Configuration, credentials and persistent
state are never package replacement targets or rollback payloads.

QWSG 1.1.0 has no native update command. For the single transition from 1.1.0,
run the verified newer archive binary with its own archive identity:

```sh
./bin/qwsg update --archive /absolute/path/qwsg-1.2.0-rc.1-linux-amd64.tar.gz --version 1.2.0-rc.1
```

The matching `.sha256` file must be adjacent. This private-candidate form is for
controlled acceptance; ordinary published updates use `qwsg update`.

Stop the exact user unit before replacing artifacts. Verify the new archive, preserve a private backup of the old binary/unit and state, then use `./install.sh --replace --backup-dir ABSOLUTE_NEW_DIRECTORY`. Reload the user manager and start only if it was previously active.

QWSG 1.0 reads Current Operator State 1.0, 1.1 and 1.2, Scheduler State 1.0, Guardian Checkpoint 1.0 and Configuration Source 1.0. Unknown, corrupt, wrong-mode, wrong-owner, symlinked, or incompatible state fails closed and is not migrated or deleted.

Run `qwsg config validate` before and after upgrade. Source Record 1.0 remains
strict and is not silently migrated. Setup preserves unspecified valid values.
Back up configuration separately from runtime state.

Run `qwsg install --check` against a new archive before replacement and
`qwsg readiness` after upgrade or rollback. Neither command executes
remediation or changes a service.

Rollback restores only the recorded old binary and unit after stopping the Guardian. Preserve state. If the old binary rejects newer state, leave the service stopped and retain the data for review.

Before uninstall, explicitly run `systemctl --user disable --now qwsg-guardian.service` and remove only the copied per-user unit. Run the matching verified release archive's `sudo ./uninstall.sh`; it refuses modified artifacts. Configuration and private state are preserved. QWSG 1.0 provides no automatic purge command.
