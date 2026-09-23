<p align="center">
  <img src="assets/branding/qwsg-logo.png" alt="QWSG logo" width="240">
</p>

<h1 align="center">Quantum Wizard Server Guardian</h1>

<p align="center"><strong>QWSG</strong> — trustworthy, privacy-preserving Linux server monitoring for operators.</p>

QWSG is a local Linux Server Guardian that provides trustworthy,
privacy-preserving evidence about observed server state and change. The Community
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

The current Community source includes authenticated release awareness and a native,
rollback-capable, operator-controlled update path. `qwsg update` verifies and
applies a supported release only after explicit administrator invocation, and
`qwsg update rollback` restores the integrity-verified previous package while
preserving user configuration, credentials and state. Release discovery authenticates official metadata using the bundled Ed25519
trust anchor; `qwsg update status` reads local awareness without network access. The Guardian performs
one isolated authenticated release-awareness check every 24 hours when due;
unattended download and installation remain disabled.

Community 1.0 is the frozen [product maturity boundary](docs/PRODUCT_1_0_SCOPE_FREEZE.md),
not a change to historical version numbering. Released 1.3.1 is immutable and
predates the current authenticated migration and recovery improvements. Current
main is not a published release. C9 still requires a distinct new signed release
and real 1.3.1 → next-release acceptance; no next version is selected here.
See the [update and recovery procedure](docs/release/UPGRADE_ROLLBACK_UNINSTALL.md).

Health evaluates canonical change evidence; `healthy` for unchanged represented
state does not certify every aspect of the server. Rule and Policy evaluate
that bounded evidence; unobserved checks are not silently healthy.

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
qwsg update status
qwsg observe
qwsg console
```

Automation retains `qwsg setup --accept-defaults`, `qwsg setup --set
KEY=VALUE`, and `qwsg config ...`.

Guardian performs authenticated release awareness on a restart-safe 24-hour
schedule. With `update.policy=notify` and configured Community email it sends
at most one successfully accepted notice per authenticated release identity;
installation always remains an explicit operator action. `qwsg update status`
also shows effective policy and capability availability. Community defaults to
manual; the [Pro policy foundation](docs/architecture/PRODUCT_CAPABILITIES_AND_UPDATE_POLICY.md)
includes a Guardian trigger and verified handoff/recovery, exercised with trusted
Pro test authority. Production resolves Community authority; this does not
provide production Pro entitlement or unattended installation.

## Documentation

- [Installation, activation, reboot and uninstall](docs/installation/INSTALL.md)
- [Quick Start](docs/release/QUICK_START.md)
- [Setup](docs/release/SETUP_AND_CONFIGURATION.md)
- [Operations](docs/release/OPERATIONS.md)
- [Troubleshooting](docs/release/TROUBLESHOOTING.md)
- [Security and privacy](docs/release/SECURITY_AND_PRIVACY.md)
- [CLI guide](docs/user/CLI_AND_SNAPSHOT_EXPLORER.en.md)
- [Automatic release checking and notification](docs/user/AUTOMATIC_RELEASE_CHECKING.en.md)

Run `qwsg help` for the command reference. Developer and architecture material
lives under `docs/development/` and `docs/architecture/` rather than dominating
this operator entrypoint. QWSG uses the source-available QWS Community /
Free License Version 1.0; see `LICENSE`.
