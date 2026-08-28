# QWSG 1.2.0-rc.3 Deterministic Packaging Candidate

This private candidate supersedes rejected private RC.2 solely to correct the
release archive permission-determinism defect recorded by Task 066. It retains
the RC.2 guided installer and product behavior while making archive modes
independent of caller umask and equivalent source-tree permission differences.

The canonical package permission policy is:

- directories: `0755`;
- intended executable regular files (`bin/qwsg`, `install.sh`, and
  `uninstall.sh`): `0755`; and
- every other regular file: `0644`.

The release builder applies the policy to the complete assembled tree after
manifest creation and before timestamp/archive normalization. Automated release
checks build equivalent inputs under different umasks and source modes, require
identical SHA-256 and bytes, and audit the exact archive-mode allowlist.

RC.3 is private and unpublished. It is not final QWSG 1.2.0 acceptance; the
external clean/full/update/rollback/reboot/notification/distribution matrix is
deferred to a later Owner-authorized task.

The frozen private candidate was built from commit
`b6eb357ad03a02b41ac93536fc3be91ecf929803` with controlled epoch `1787905053`.
Artifact `qwsg-1.2.0-rc.3-linux-amd64.tar.gz` has SHA-256
`8543c3e09b48085b01c037d7db5106ea793374dc099b0b9be5f6cacb55af13ee`.
