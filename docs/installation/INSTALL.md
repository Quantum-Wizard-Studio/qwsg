# QWSG Installation

The Version 1.0 operator path is the verified prebuilt archive described in
`docs/release/QUICK_START.md`; it needs neither repository knowledge nor Go.
The production prefix is deliberately fixed at `/usr/local`, keeping the
audited absolute systemd `ExecStart` and executable inseparable. Installation
copies artifacts only; service activation and lingering are explicit acts.

After installation, run `qwsg setup` as the intended ordinary Guardian user,
review with `qwsg config show`, and validate before explicitly enabling the user
unit. Installation never writes per-user configuration. See
`docs/release/SETUP_AND_CONFIGURATION.md`.

Repository `make install-service` is a development convenience and rejects a
nonstandard prefix. `DESTDIR` remains available for isolated staging.

## Supported acceptance environment

Task 017 verifies source builds and installation on Ubuntu 24.04 with GNU Make,
the Go version declared by `go.mod`, and GNU `install`. QWSG has no external Go
module or runtime dependency.

## Build

```bash
make build
build/qwsg version
```

The binary is written to `build/qwsg`. `BUILD_COMMIT` defaults to the current
short Git commit and `BUILD_DATE` defaults to the deterministic value
`unknown`. Release automation may provide controlled values:

```bash
make BUILD_COMMIT=abcdef123456 BUILD_DATE=2026-07-24T00:00:00Z build
```

## Isolated installation

Use `DESTDIR` to verify the layout without privilege:

```bash
stage="$(mktemp -d)"
make DESTDIR="$stage" PREFIX=/usr/local install
"$stage/usr/local/bin/qwsg" version
```

This creates only `usr/local/bin/qwsg` under the staging directory and installs
the binary with mode `0755`.

## System installation

The default destination is `/usr/local/bin/qwsg`. Inspect and back up an
existing target before replacement. Build as the normal user, then elevate only
the artifact-copy step:

```bash
make build
test ! -e /usr/local/bin/qwsg
sudo make install
qwsg version
```

`make install` never invokes Go and fails if the normal-user build artifact
`build/qwsg` is missing or not executable. The privileged environment therefore
does not need the Go compiler or the normal user's Go `PATH`.

If the target exists, record its owner, mode, and SHA-256 and copy it to a
private rollback location before installation. `make install` does not create a
service, user, group, configuration, state directory, scheduled job, or network
listener. QWSG runtime commands must be run as an ordinary user.

For another prefix:

```bash
make build
make PREFIX=/opt/qwsg install
/opt/qwsg/bin/qwsg version
```

## Upgrade, rollback, and removal

An installation replaces only the exact `$(DESTDIR)$(PREFIX)/bin/qwsg` target.
To roll back, restore the exact pre-install binary with its recorded owner and
mode. If the target was previously absent, compare its hash with the delivered
Task 017 binary and remove only that exact file:

```bash
sudo rm -- /usr/local/bin/qwsg
```

Do not recursively remove a prefix, an Inventory Store, or unrelated files.
There is no automatic updater or package-manager integration in Task 017.

## Guardian user service

Build as the ordinary runtime user, then install the binary and auditable user
unit (installation does not enable or start it):

```bash
make build
sudo make install-service
systemctl --user daemon-reload
systemctl --user enable --now qwsg-guardian.service
```

Use `systemctl --user start|stop|restart qwsg-guardian.service`,
`systemctl --user enable|disable qwsg-guardian.service`, and
`systemctl --user status qwsg-guardian.service`. Read service logs with
`journalctl --user-unit=qwsg-guardian.service`. Installed, enabled, active and
canonically observed are distinct states; bare `qwsg` reports only validated
canonical lifecycle evidence.

The service runs as the same non-root user that owns QWSG state. Starting at
boot before login additionally requires an administrator to deliberately
enable lingering for that exact user (`loginctl enable-linger USER`). Without
a verified user manager and lingering, boot operation is unavailable rather
than guaranteed. Never run the Guardian as root merely to bypass collector or
state permissions.

The default cadence is five minutes with a two-minute cycle timeout and a
ten-minute lifecycle freshness window. State is private below the systemd user
state directory. To use a canonical Source Record, add `--config-source` to a
local unit override; do not put secrets or raw provider destinations in the
unit. Stop and disable the unit before removing it. For rollback, restore the
previous binary/unit, run `systemctl --user daemon-reload`, and retain the
private state until compatibility has been assessed.
