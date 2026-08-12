# QWSG 1.0 Upgrade, Rollback, and Uninstall

Stop the exact user unit before replacing artifacts. Verify the new archive, preserve a private backup of the old binary/unit and state, then use `./install.sh --replace --backup-dir ABSOLUTE_NEW_DIRECTORY`. Reload the user manager and start only if it was previously active.

QWSG 1.0 reads Current Operator State 1.0, 1.1 and 1.2, Scheduler State 1.0, Guardian Checkpoint 1.0 and Configuration Source 1.0. Unknown, corrupt, wrong-mode, wrong-owner, symlinked, or incompatible state fails closed and is not migrated or deleted.

Run `qwsg config validate` before and after upgrade. Source Record 1.0 remains
strict and is not silently migrated. Setup preserves unspecified valid values.
Back up configuration separately from runtime state.

Rollback restores only the recorded old binary and unit after stopping the Guardian. Preserve state. If the old binary rejects newer state, leave the service stopped and retain the data for review.

Before uninstall, explicitly run `systemctl --user disable --now qwsg-guardian.service` and remove only the copied per-user unit. Run the matching verified release archive's `sudo ./uninstall.sh`; it refuses modified artifacts. Configuration and private state are preserved. QWSG 1.0 provides no automatic purge command.
