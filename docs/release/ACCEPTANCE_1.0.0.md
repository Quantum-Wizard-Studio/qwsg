# QWSG 1.0.0 Engineering Acceptance

## Current gate

The Project Owner authorized the exact 140-path release-source allowlist, and
the canonical release-source commit now exists as
`177535e44b2ce5ed9efd73ab0793ffe6881f0cd6`. The final artifact was built from
that exact commit and passed the applicable technical gates below. This
post-build evidence update remains uncommitted pending the separate Owner gate
for an evidence-only commit. No tag, push or publication is authorized.

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

## Post-commit evidence

- Release-source commit: `177535e44b2ce5ed9efd73ab0793ffe6881f0cd6`.
- Commit subject: `release: establish QWSG 1.0.0 release-source baseline`.
- Commit scope: exactly the Owner-approved 140 paths; none of the 94 excluded
  paths entered the commit.
- Controlled build epoch: `1786511170` (`2026-08-12T05:06:10Z`).
- Final archive: `qwsg-1.0.0-linux-amd64.tar.gz`.
- Final archive SHA-256:
  `edfba7366adf2c1ce0a8ce56369bb0dc5ad11326c4e3d1e301625a5313292fa5`.
- Two independent commit-archive builds produced byte-identical binary,
  internal manifest, archive and sidecar.
- The sidecar, internal `MANIFEST.sha256`, safe single-root archive layout,
  regular-file/directory type boundary, static linux-amd64 binary, packaged
  Owner-approved `LICENSE`, exact `1.0.0` version and full embedded commit
  identity all passed.
- Archive-only clean install, collision refusal, explicit replacement with
  backup verification and clean uninstall passed in an isolated destination.
- Empty-HOME extracted-binary acceptance passed: the first observation created
  a truthful private baseline, the second completed the full pipeline, a
  separate Console loaded the state, partial `check` published successfully,
  directories were `0700` and files were `0600`.
- Focused Task 039/041 tests, complete and race Go tests, vet, formatting,
  Framework, Builder, lifecycle, diverted-task, test-task, job, staged
  whitespace, privacy and exact-index gates passed. Static systemd unit
  verification passed with the already documented sandbox bus warning; Task
  043 remains the canonical real-host systemd/reboot evidence.

The release-source and artifact gates are technically satisfied. The evidence-
only commit and any `v1.0.0` tag, push, Forgejo Release or public publication
remain separate Owner-authorized actions. Therefore the current decision is
`READY FOR EVIDENCE-ONLY COMMIT REVIEW`, not yet the final publication
decision.
