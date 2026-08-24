# QWSG Installation Guide

This is the canonical installation guide. Return to the [operator
README](../../README.md) for the product overview. In an archive this file is
`INSTALL.md`; installed documentation is under `/usr/local/share/doc/qwsg/`.

## Supported host and prerequisites

QWSG supports Ubuntu 24.04 LTS, amd64, systemd 255+, an ordinary non-root user,
a working systemd user manager, and local filesystem semantics. The prebuilt
binary needs no Go runtime. Verification needs `sha256sum`; installation uses
the archive's deterministic `install.sh`. Unsupported hosts receive no guessed
commands.

## Verify, assess, and install

```sh
sha256sum -c qwsg-1.1.0-rc.5-linux-amd64.tar.gz.sha256
tar -xzf qwsg-1.1.0-rc.5-linux-amd64.tar.gz
cd qwsg-1.1.0-rc.5-linux-amd64
sha256sum -c MANIFEST.sha256
./bin/qwsg install --check
sudo ./install.sh
```

Smart Install is read-only. For every supported actionable finding it explains
the impact, verification procedure, operator action, privilege boundary,
safety notes, and mandatory revalidation. It prints an exact remediation
command only when the detected cause and supported-platform mapping prove that
command safe. Apply any host change yourself, then rerun:

```sh
qwsg install --check
```
`install.sh` copies immutable artifacts only; it never configures SMTP,
lingering, a user, or a service.

## Guided setup and notification

As the ordinary Guardian user:

```sh
qwsg setup
qwsg readiness
```

Setup resumes from valid existing evidence. Community supports one recipient.
For password automation use a private current-user mode-0600 input file:

```sh
qwsg notification credential set --from-file PRIVATE_FILE
qwsg notification preflight
qwsg notification test
```

A controlled test—not configuration acceptance—verifies notification.

## Guardian activation and readiness

Guided setup separately confirms ordinary-user activation and never invokes
sudo. Before contacting systemd, it creates or validates the packaged
Guardian's canonical state directory as a real current-user-owned mode-0700
non-symlink path. Unsafe existing state is never removed, followed, replaced,
chmodded or chowned; activation stops with a state-preparation diagnostic. It
then validates the effective user's canonical runtime directory, checks the
same user manager as readiness, and performs only the fixed user-unit reload
and enable/start sequence. A failure names the fixed stage, preserves
configuration, and gives a bounded QWSG assessment/resume action. Manual
operation remains:

```sh
systemctl --user daemon-reload
systemctl --user enable --now qwsg-guardian.service
qwsg readiness
```

READY requires fresh integrity-checked Guardian evidence. Allow the bounded
first cycle to complete and rerun readiness when instructed. Working core with
unverified notification is PARTIAL.

## Logout, reboot, uninstall, and rollback

Boot before login additionally requires administrator-authorized lingering for
the exact runtime user. QWSG reports but never changes it. After restart or
reboot, run `qwsg readiness` and require fresh evidence again.

```sh
systemctl --user disable --now qwsg-guardian.service
sudo ./uninstall.sh
```

The uninstaller removes only unchanged release-owned artifacts and preserves
user configuration, credentials, and state. Keep the verified archive for
uninstall/rollback. Replacement uses the documented `--replace` with a new
private `--backup-dir`; never recursively delete `/usr/local` or user data.

## Troubleshooting and instruction locations

- Run `qwsg install --check`, then `qwsg readiness`.
- Run `qwsg config validate` for configuration failures.
- Run `qwsg notification preflight` before a controlled test.
- Run `systemctl --user status qwsg-guardian.service` for unit state.
- Read archive `docs/TROUBLESHOOTING.md` or installed
  `/usr/local/share/doc/qwsg/TROUBLESHOOTING.md`.

Repository entrypoints are [README.md](../../README.md) and this guide. Archive
entrypoints are `README.md` and `INSTALL.md`. Installed entrypoints are
`/usr/local/share/doc/qwsg/README.md` and
`/usr/local/share/doc/qwsg/INSTALL.md`.
