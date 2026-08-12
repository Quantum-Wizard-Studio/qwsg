# QWSG 1.0.0 Engineering Acceptance

## Current gate

The final `1.0.0` source metadata and Owner-approved QWS Community / Free
License Version 1.0 are prepared locally. This record does not yet claim a final
artifact or publication decision. The exact release-source staging allowlist,
diff, modes, license hash and validation evidence must receive Project Owner
authorization before the release-source commit.

## Accepted product baseline

The final product behavior is unchanged from the Task 043-accepted RC.3
baseline. Tasks 038–042 provide reproducible archive, large Policy Report,
read-only Console refresh, bounded Runtime diagnostics, truthful lifecycle and
secure empty-HOME bootstrap evidence. Task 043 closes the genuine no-Go clean
host, physical reboot, recurring Guardian, restart, resource and uninstall
gates on Ubuntu 24.04 amd64 with systemd 255.

## Final identity contract

- Version: `1.0.0`.
- Expected archive: `dist/qwsg-1.0.0-linux-amd64.tar.gz`.
- Expected sidecar: `dist/qwsg-1.0.0-linux-amd64.tar.gz.sha256`.
- Embedded commit: the exact, separately authorized release-source commit;
  it must not use the historical RC.3 commit label for the reconciled source.
- Build time: controlled by the commit-tied `SOURCE_DATE_EPOCH` recorded after
  the release-source commit exists.
- License: exact Owner-approved QWS Community / Free License Version 1.0.

## Required post-commit evidence

After the release-source commit is separately authorized, the artifact must be
built twice in independent cache/output roots. Binary, internal manifest,
archive and sidecar must match byte-for-byte. Archive safety, manifest,
checksum, static version/commit identity, staged install/upgrade/rollback/
uninstall, systemd unit and applicable Task 039/041 artifact behavior must pass.

Actual commit hashes, artifact hashes, measurements and the final
`READY FOR QWSG 1.0.0 PUBLICATION` decision belong here only after those events
occur. Tag, push, Forgejo Release and public publication remain separate Owner
actions even after technical readiness.
