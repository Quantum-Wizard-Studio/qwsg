# QWSG 1.0.0-rc.2 Engineering Acceptance

## Release decision

`READY FOR CLEAN-HOST ACCEPTANCE`

This decision covers the reproducible linux-amd64 archive and development-host
artifact validation. Clean Ubuntu 24.04 installation, reboot persistence and
final release acceptance remain Owner-run gates. It does not authorize a public
license, Git tag, signing, upload, publication or announcement.

## Artifact identity and reproducibility

- Identity: `1.0.0-rc.2`, linux-amd64.
- Archive: `dist/qwsg-1.0.0-rc.2-linux-amd64.tar.gz`.
- Archive SHA-256: `0694fb7f382ea1b373aaf0b6f0171a3fc580c491c1af3b8657a3bb8697ed897b`.
- Sidecar: `dist/qwsg-1.0.0-rc.2-linux-amd64.tar.gz.sha256`.
- Internal manifest SHA-256:
  `f2d70d8e444d0b4021007f9e4dc7052629a0877fce586310e4ce8d04ad2aa206`.
- Fixed build inputs were commit `0a8a5c7e7224` and
  `SOURCE_DATE_EPOCH=1786406400` (`2026-08-11T00:00:00Z`).
- Two clean output directories produced byte-identical binaries, manifests,
  archives and checksum sidecars.
- The archive has one safe root, stable metadata, no links or special files,
  the exact RC.2 release notes and no RC.1 release-notes substitution.
- Historical RC.1 archive, sidecar and release-document bytes remained
  unchanged.

## Package and installation evidence

The internal manifest verified the statically linked QWSG executable, systemd
user unit, installer, uninstaller, configuration example, license, changelog
and complete release documentation. The extracted binary reports version
`1.0.0-rc.2`, commit `0a8a5c7e7224` and the controlled build time. Supported
installation and execution require no Go, Git, repository checkout or
development tooling.

Archive-only staged installation passed clean install, collision refusal,
explicit replacement with backup, backup integrity and clean owned uninstall.
A deliberately modified packaged file correctly caused uninstall refusal. The
test staging tree was then restored from the archive before clean uninstall;
unrelated staged user data remained untouched. No real development-host
installation destination was used.

## Task 039 artifact behavior

- The extracted final binary completed recurring cycles against an isolated
  copy of real canonical host evidence with Policy Reports above the discovered
  366-source boundary. A bounded endurance sample recorded 50 consecutive
  successful Scheduler results with a maximum of 456 Policy sources.
- A separate positive lifecycle recorded 13 successful Scheduler results,
  including consecutive completed Runtime cycles with up to 456 Policy
  sources. No `alert_evaluation_failed` or source-count-induced
  `runtime_not_completed` evidence appeared.
- A separate interactive Console displayed a running Guardian and current,
  complete evidence. Pressing `r` redrew the loaded Current State without a
  refresh failure or `guardian_active` conflict; explicit `observe` remained
  safely refused while Guardian owned the lock.
- Attention projected 708 candidates into one represented item with 707
  correlated and none omitted, avoiding duplicate internal-token noise.
- Graceful termination produced `active=false` checkpoint evidence and a
  stopped Console state. After forced termination, the last active checkpoint
  remained historical but the Console demoted the Guardian to unavailable once
  lifecycle evidence crossed its freshness boundary.
- An intentionally short 400 ms Runtime deadline separately exercised timeout
  behavior; it is not counted as positive Runtime-completion evidence.

## Lifecycle, resource and validation evidence

The shipped systemd user unit passed host verification and an isolated link,
start, recurrence, separate Console observation, restart and graceful-stop
sequence. The service remained unprivileged, linger stayed disabled, and the
temporary unit and private state were removed. During the positive foreground
run, observed resident memory was about 12 MiB with 10 tasks/threads and seven
open descriptors. The systemd run stayed within its configured resource limits
and did not restart unexpectedly.

Focused Task 039 regressions, complete Go tests, race tests, vet, formatting,
Framework, Builder, lifecycle, diverted-task, repository, manifest, checksum,
privacy and Git-diff checks passed. No product semantics were changed by Task
040; only coherent release identity, metadata and packaging references changed.

## Clean-host transfer and remaining Owner gates

Transfer exactly these two files to the clean Ubuntu 24.04 x86-64 host:

- `dist/qwsg-1.0.0-rc.2-linux-amd64.tar.gz`
- `dist/qwsg-1.0.0-rc.2-linux-amd64.tar.gz.sha256`

Verify the sidecar before extraction, then follow the packaged Quick Start.
Clean-host installation, Guardian operation, reboot persistence and uninstall
remain unexecuted here. SHA-256 supplies integrity, not publisher
authentication. The temporary proprietary license remains unchanged and must
be confirmed or replaced before public publication.
