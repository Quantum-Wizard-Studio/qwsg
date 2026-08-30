<p align="center">
  <img src="assets/branding/qwsg-logo.png" alt="QWSG logo" width="240">
</p>

<h1 align="center">Quantum Wizard Server Guardian</h1>

<p align="center"><strong>QWSG</strong> — trustworthy, privacy-preserving Linux server monitoring for operators.</p>

QWSG is a local Linux Server Guardian that provides trustworthy,
privacy-preserving evidence about server health and change. The Community
edition runs as an ordinary user, keeps evidence locally, and supports one
administrator email recipient through operator-controlled SMTP without a QWS
account.

The canonical source repository is hosted by
[Quantum Wizard Studio on Forgejo](https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg).
Public mirrors are read-only distribution points.

The supported platform is Ubuntu 24.04 LTS on amd64 with systemd 255+. Start
with the dedicated [installation guide](docs/installation/INSTALL.md). A
release archive exposes it as `INSTALL.md`; after installation it is at
`/usr/local/share/doc/qwsg/INSTALL.md`.

QWSG 1.2.0 is the current stable baseline and includes a native,
rollback-capable, operator-controlled update path. `qwsg update` verifies and
applies a supported release only after explicit administrator invocation, and
`qwsg update rollback` restores the integrity-verified previous package while
preserving user configuration, credentials and state. QWSG 1.3 update
discovery and local awareness work is under development; unattended
installation remains disabled.

## Normal journey

```text
verify archive -> ./bin/qwsg install --guided -> localized plan and consent
-> narrow package installation -> configuration -> Guardian readiness
```

The wizard supports English, Magyar, and Deutsch, derives progress from actual
phase state, and invokes the archive's fixed package helper through `sudo` only
after consent. It never enables lingering automatically. The concise expert
path remains `install --check`, `sudo ./install.sh`, `qwsg setup`, then
`qwsg readiness`. Readiness distinguishes working Guardian core from optional
external notification and requires fresh evidence—not process state alone.

## Principal commands

```sh
qwsg install --guided
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
