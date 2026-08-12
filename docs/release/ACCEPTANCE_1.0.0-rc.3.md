# QWSG 1.0.0-rc.3 Engineering Acceptance

## Release decision

`READY FOR QWSG 1.0 FINAL RELEASE`

This technical decision combines reproducible development-host artifact
validation with the Project Owner's completed acceptance on a freshly
reinstalled disposable Ubuntu 24.04 amd64 VPS. The previously outstanding
physical no-Go, reboot-persistence and real-uninstall gates are satisfied. This
decision does not authorize a public license, final `1.0.0` identity, commit,
Git tag, signing, push, upload, publication or announcement.

## Artifact identity and reproducibility

- Identity: `1.0.0-rc.3`, linux-amd64.
- Archive: `dist/qwsg-1.0.0-rc.3-linux-amd64.tar.gz`.
- Archive SHA-256:
  `bc4eca323cb07d23f0d6c884886655eb5549c03c136d8c773782d7813551585c`.
- Sidecar: `dist/qwsg-1.0.0-rc.3-linux-amd64.tar.gz.sha256`.
- Internal manifest SHA-256:
  `1b764edc8686f12618c73a8870ce714ce02378152784a94817d9ff2aabc333f4`.
- Fixed build inputs were commit `0a8a5c7e7224` and
  `SOURCE_DATE_EPOCH=1786406400` (`2026-08-11T00:00:00Z`).
- Two independent clean build/cache roots produced byte-identical binaries,
  manifests, archives and location-independent checksum sidecars.
- The archive has one safe root, stable metadata, no links or special files,
  the exact RC.3 release notes and a statically linked linux-amd64 executable.
- Historical RC.1 and RC.2 archives, sidecars and release documents remained
  byte-identical to the pre-task snapshot.

## Package and installation evidence

The internal manifest verified the binary, systemd user unit, installer,
uninstaller, configuration example, license, changelog and complete packaged
operator documentation. The extracted binary reports `1.0.0-rc.3`, commit
`0a8a5c7e7224` and the controlled build time.

Archive-only staged installation passed clean install, collision refusal,
explicit RC.2-to-RC.3 replacement, backup integrity, exact-backup rollback,
modified-artifact uninstall refusal, archive restoration and clean owned
uninstall. Unrelated staged Owner data remained untouched. The shipped unit
passed `systemd-analyze verify --user`. No real development-host installation,
service activation or linger change occurred.

The executable is statically linked and ran from an extracted archive outside
the repository with a minimal runtime PATH. The development host still has a
physical Go installation, so absence of that tree was not claimed during Task
042. Execution on the actual no-Go clean host was the remaining Owner
acceptance gate and is satisfied by the evidence below.

## Task 039 artifact behavior

- Two real host observations from the extracted binary completed the full
  pipeline with exactly 366 Policy and Report sources. Runtime did not become
  partial from Alert source-count compatibility.
- An isolated foreground Guardian recorded 51 successful Scheduler/Runtime
  cycles. A separate Console showed `Guardian: running` with current complete
  evidence, while explicit `observe` was refused with `guardian_active`.
- A direct PTY `r` action cleared and redrew the loaded Current State without
  `Refresh failed` or a competing observation.
- Attention remained correlated and bounded. No `alert_evaluation_failed` or
  source-count-induced `runtime_not_completed` evidence appeared.
- Graceful interrupt wrote `active=false`. A separately SIGKILLed Guardian
  intentionally retained historical `active=true`, but the Console demoted it
  to unavailable after the exclusive lifecycle freshness boundary.

## Task 041 artifact behavior

- From a genuinely empty synthetic HOME with no `.local` hierarchy, the first
  extracted-binary `observe` created the private state tree, one truthful
  partial baseline and Current Operator State without manual preparation.
- A second `observe` completed the normal eight-stage pipeline; a separate bare
  Console loaded the published state, and partial `check` published normally.
- Every QWSG directory was mode `0700`, every private state file was mode
  `0600`, and the existing HOME mode remained unchanged.
- Artifact-level symlink-component and permissive-root attempts failed closed
  with bounded `state_unsafe` and `state_permission` diagnostics and did not
  create or chmod protected descendants.
- Focused regressions retain `state_bootstrap_failed` and
  `state_publication_failed` category coverage without exposing raw errors.

## Owner clean-host acceptance

The Project Owner supplied the following sanitized acceptance evidence from a
fresh disposable Ubuntu 24.04 LTS amd64 VPS with systemd 255, an ordinary UID
1000 runtime user, a running user manager, no prior QWSG installation or state,
and no Go toolchain. Git happened to be installed but was not used by QWSG.
No hostname, address, private configuration, raw state or unbounded journal
content is retained in this record.

The transferred RC.3 archive matched the published sidecar and its complete
internal manifest. `sudo ./install.sh` installed the fixed `/usr/local`
artifacts and deliberately did not enable or start the service. The installed
binary reported version `1.0.0-rc.3`, commit `0a8a5c7e7224` and build time
`2026-08-11T00:00:00Z`.

Before first execution, `~/.local/state/qwsg` did not exist. The first
`qwsg observe` created the private state hierarchy without manual preparation,
completed successfully, emitted eight truthful partial Inventory records and
one partial Snapshot, established the baseline, and retained unknown condition
semantics. The partial components capability was expected because Go was
absent; QWSG itself required neither Go nor repository content.

The second observation completed the full canonical pipeline. Inventory and
Snapshot remained truthfully partial; Compare, Drift, Health, Rule, Policy and
Report each completed with 323 records. Overall execution was complete. A
separate Console loaded the published Current Operator State and reported a
degraded, current/partial view with two changes, zero Alerts, 645 correlated
concerns and zero omissions.

## Real Guardian, reboot and restart acceptance

The Owner loaded and enabled the shipped user unit explicitly. Before reboot it
was active/running with zero restarts, six tasks, about 2.7 MiB resident memory
and about 3.1 MiB observed peak. A separate Console reported the Guardian
running with current/complete evidence.

Boot-before-login support was explicitly enabled for the exact runtime user and
`Linger=yes` was verified. The disposable VPS was physically rebooted. No QWSG
process or service was manually started afterward. The user manager restored
the enabled unit with a new PID 835 and new InvocationID
`90720850ce3e4d498827350861708813`; the process was the installed
`/usr/local/bin/qwsg guardian run`. This generation change, the boot journal
start evidence and the separate running/current Console exclude stale
pre-reboot evidence.

After a complete recurring post-reboot cycle, the unit remained active/running
with zero restarts and the same valid post-reboot generation. Observed
`MemoryCurrent` was 13,828,096 bytes, `MemoryPeak` was 21,262,336 bytes and
`TasksCurrent` was 8, all below the shipped 128 MiB and 32-task limits. Current
Operator State, Guardian checkpoint and Scheduler state timestamps advanced;
all observed state files were mode `0600`. The Console reported running,
current/partial evidence with 11 changes, zero Alerts, 647 correlated concerns
and zero omissions.

A real controlled service restart produced new PID 1176 and InvocationID
`e2b15a8b697b4c94b7b78105cfe73101`, while retaining active/running state and
zero restarts. A separate Console again reported running/current/partial
evidence.

## Real uninstall acceptance

The Guardian was disabled and stopped and lingering was disabled before the
archive uninstaller ran. The owned binary and user unit were removed, the unit
became unavailable to systemd, and private runtime/state evidence was
intentionally preserved with mode `0600`. A shell may temporarily cache a
removed command path until `hash -r`; filesystem absence is authoritative and
this shell behavior is not an uninstall defect.

## Release-gate reconciliation

| Gate source | Clean-host evidence | Result |
| --- | --- | --- |
| Task 038 / Release Policy | Verified install, user service, linger, physical reboot, post-reboot recurrence, restart, bounded resources and uninstall | Satisfied |
| Task 039 | Separate Console observed truthful Guardian lifecycle; recurrence completed without Runtime/Alert compatibility failure | Satisfied |
| Task 041 | Genuine absent state root and no-Go first observation bootstrapped safely; partial evidence remained truthful | Satisfied |
| Task 042 | Exact RC.3 archive/sidecar/manifest/version was the tested payload | Satisfied |

No supplied result contradicts Tasks 038, 039, 041 or 042, and no technical
Version 1.0 release blocker remains.

## Console PTY observation

The Owner observed two Overview renderings in some real PTY interactions. The
tested Console contract emits one initial Overview and one clear/redraw after
an explicit refresh or accepted navigation action. Existing deterministic and
real PTY evidence verifies that boundary, and the clean-host observation did
not establish an unsolicited duplicate independent of input or any practical
impairment. The observation is therefore non-blocking cosmetic follow-up for
post-1.0 investigation, not authority to modify Console behavior in this
acceptance task.

## Validation and remaining Owner decisions

Focused Task 039/041 tests, complete Go tests, repository-wide race tests, vet,
formatting, Framework, Builder, lifecycle, diverted-task, test-task audit,
archive, checksum, manifest, staged installer/uninstaller, unit, privacy and
Git-diff gates passed. Two rejected test procedures are recorded in Task 042
history and changed neither artifact nor product behavior.

The packaged `KNOWN_LIMITATIONS.md` correctly described physical reboot as an
Owner gate when the artifact was built; that immutable historical payload is
not rewritten after acceptance. The gate is now closed by the evidence above.

The technical release is ready. SHA-256 supplies integrity, not publisher
authentication. The temporary proprietary license remains unchanged. Final
public licensing, a final `1.0.0` metadata refresh if authorized, commit
strategy, Git tag, signing, push and external publication remain separate
Project Owner decisions.
