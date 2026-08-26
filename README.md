# Quantum Wizard Server Guardian

QWSG is a local Linux Server Guardian for operators who need trustworthy,
privacy-preserving evidence about server health and change. Community runs as
an ordinary user, keeps evidence locally, and supports one administrator email
recipient through operator-controlled SMTP without a QWS account.

The supported platform is Ubuntu 24.04 LTS on amd64 with systemd 255+. Start
with the dedicated [installation guide](docs/installation/INSTALL.md). A
release archive exposes it as `INSTALL.md`; after installation it is at
`/usr/local/share/doc/qwsg/INSTALL.md`.

QWSG 1.2 adds a native, rollback-capable update path. Use `qwsg update check`
to inspect the canonical public Release source, `qwsg update` to verify and
apply a newer supported release, `qwsg update status` to inspect local rollback
availability, and `qwsg update rollback` to restore the integrity-verified
previous package while preserving user configuration, credentials and state.

## Normal journey

```text
verify archive -> ./bin/qwsg install --check -> sudo ./install.sh
-> qwsg setup -> qwsg readiness
```

Setup is resumable and guides configuration, optional external notification,
and explicit Guardian activation. QWSG never installs packages, invokes sudo,
or enables lingering for you. Readiness distinguishes working Guardian core
from external notification and requires fresh evidence—not process state alone.

## Principal commands

```sh
qwsg setup
qwsg setup --plan --format json
qwsg config show
qwsg notification preflight
qwsg notification test
qwsg readiness
qwsg update check
qwsg observe
qwsg console
```

Automation retains `qwsg setup --accept-defaults`, `qwsg setup --set
KEY=VALUE`, and `qwsg config ...`.

## Documentation

- [Installation, activation, reboot and uninstall](docs/installation/INSTALL.md)
- [Quick Start](docs/release/QUICK_START.md)
- [Setup](docs/release/SETUP_AND_CONFIGURATION.md)
- [Operations](docs/release/OPERATIONS.md)
- [Troubleshooting](docs/release/TROUBLESHOOTING.md)
- [Security and privacy](docs/release/SECURITY_AND_PRIVACY.md)
- [CLI guide](docs/user/CLI_AND_SNAPSHOT_EXPLORER.en.md)

Run `qwsg help` for the command reference. Developer and architecture material
lives under `docs/development/` and `docs/architecture/` rather than dominating
this operator entrypoint. QWSG uses the source-available QWS Community /
Free License Version 1.0; see `LICENSE`.
