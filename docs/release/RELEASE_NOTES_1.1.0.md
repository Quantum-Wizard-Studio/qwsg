# QWSG 1.1.0 Release Notes

QWSG 1.1.0 is the final release of the product behavior externally accepted as
QWSG 1.1.0-rc.6. The final source changes are limited to release identity,
documentation and deterministic release plumbing; they do not change runtime
behavior.

## Highlights

- Read-only Smart Install assessment for supported Ubuntu 24.04 amd64 hosts.
- Guided, resumable ordinary-user setup and explicit Guardian activation.
- Safe configuration/state separation with strict ownership, modes and
  non-symlink filesystem handling.
- Optional one-recipient SMTP notification using protected local credentials.
- Operational readiness that distinguishes Guardian core readiness from
  optional external-notification verification.
- Pre-login Guardian startup with systemd user lingering and automatic
  `Restart=on-failure` recovery.

## Guardian recovery correction

The final release contains the RC.6 correction for `QWSG-059-F001`. After a
configuration identity changes, a boot-time or systemd-recovered Guardian now
replaces only integrity-valid Scheduler state owned by the superseded
configuration. The recovered generation can own its checkpoint, complete fresh
work and publish current canonical evidence. Same-configuration recovery,
generation isolation and truthful degradation for genuine failures remain
intact.

## Acceptance basis

Task 061 externally verified Ubuntu 24.04 amd64 Smart Install, documented RC.5
replacement, safe state and credential preservation, controlled SMTP delivery
with actual Owner-confirmed receipt, physical reboot, pre-login autostart,
automatic systemd recovery, current configuration/generation ownership and
fresh canonical Guardian evidence. No RC.6 product defect was found.

`notification.external=unknown_requires_verification` and overall partial are
valid when the Guardian core is ready but no Guardian monitoring notification
queue entry exists. A controlled `notification test` proves transport delivery
without manufacturing such an incident.

## Installation

Download the immutable versioned archive and checksum sidecar from the Forgejo
Release, verify the sidecar, extract the archive, run `./bin/qwsg install
--check`, then follow `INSTALL.md`. QWSG does not install dependencies, enter
credentials, enable lingering or perform privileged host remediation.

This release is QWS Community / Free License Version 1.0 software. See
`LICENSE`, `README.md`, `INSTALL.md`, and the packaged operator documentation.
