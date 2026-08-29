# QWSG Setup and Configuration

## Existing installation safety

Guided installation verifies the canonical installed package layout and
release provenance before deciding that QWSG is already installed. A file
named `/usr/local/bin/qwsg`, a PATH entry, old configuration, or state alone is
not sufficient. Legacy or unverified artifacts are not overwritten: back them
up, review their origin, and resolve them explicitly before retrying. Supported
older packages use the explicit `qwsg update` workflow rather than guided
installation.

## Normal path

Run setup as the ordinary user who will run Guardian:

```sh
qwsg install --guided
qwsg setup
qwsg config show
qwsg config validate
qwsg readiness
```

For new installations, `install --guided` selects English, Hungarian, or
German; verifies the supported platform; presents a pre-mutation plan; invokes
the narrow archive package helper only after consent; initializes safe defaults;
explains optional SMTP setup; records manual or notify update preference; and
offers Guardian activation and readiness verification. `--line-mode` disables
dashboard cursor control.

On a terminal, standalone setup guides the configuration write and separately asks before
activating the fixed QWSG user service. It is resumable from canonical state.
Use `qwsg setup --plan [--format human|json]` for a read-only deterministic
plan; it never prompts, writes, contacts SMTP, or changes a service.
Environment blockers shown by guided setup reuse the same Assessment Model 1.1
guidance as `qwsg install --check`; setup neither duplicates nor executes the
recommended operator action.

For deterministic automation, use `qwsg setup --accept-defaults`. Repeat setup
preserves a valid configuration unless an explicit `--set KEY=VALUE` is given.
Setup displays its destination and proposed effective values before an
interactive write. Noninteractive setup never enables or starts a service;
interactive setup does so only after the separate explicit activation prompt.

Assessment does not invent administrator addresses, SMTP providers,
credentials, sender identities, or server-purpose choices. No VPS profile is
persisted; a versioned profile extension is deferred until separately designed.

## Locations and precedence

The primary file is `$XDG_CONFIG_HOME/qwsg/config.json`, or
`$HOME/.config/qwsg/config.json` when XDG configuration home is unset. Use
`--config ABSOLUTE_FILE` for an explicit command-specific selection.

Resolution order is compiled defaults, the discovered primary local file, an
explicit file replacing discovery, then typed command-temporary overrides.

The configuration directory is mode `0700`; the file is mode `0600`; both are
owned by the current user. Symlinks, special files, unsafe paths, wrong owners,
and permissive modes are rejected. Writes are atomic and preserve the prior
valid file on a pre-rename failure.

## Command reference

```sh
qwsg config show [--format human|json]
qwsg config validate [--format human|json]
qwsg config get KEY [--format human|json]
qwsg config set KEY VALUE [--format human|json]
```

The base mutable keys are `locale`, `time_zone`, `snapshot_retention`,
`guardian.interval`, `guardian.cycle_timeout`, and `update.policy` (`manual` or
`notify`; automatic is rejected). Task 046 also supports
`notification.email.enabled`, `.recipient`, `.host`, `.port`, `.sender`,
`.security`, `.auth`, `.username`, `.credential_ref`, and `.timeout`.
`notification.lifecycle.enabled` independently enables QWSG-managed change
messages while reusing that same SMTP transport and recipient.
Community permits exactly one administrator recipient. Durations use forms such as
`30s`, `2m`, or `1h`. Unknown keys and arbitrary JSON paths are rejected.

The compiled defaults are locale `en`, time zone `UTC`, snapshot retention
`10`, Guardian interval `5m`, cycle timeout `2m`, and concurrency `1`.

The packaged `qwsg-config.json` is the exact default setup Source Record for
reference only. Installation places it with documentation and does not activate
it; use `qwsg setup` to create the private per-user copy.

## Strict validation and compatibility

Configuration Source Record 1.0 is strict. Unknown fields, malformed/trailing
JSON, unsupported versions, invalid identities or references, and out-of-bound
values fail closed. QWSG does not silently migrate, repair, or discard an
invalid file. Guardian validates before mutating state; the Console also
refuses normal operation while discovered configuration is invalid.

## Community email and secrets

The general JSON file stores references only. Configure non-secret email fields
while disabled, store a password from a protected current-user-owned `0600`
file with `qwsg notification credential set --from-file FILE`, enable email,
then run `qwsg notification preflight` and `qwsg notification test`. Passwords
are never command arguments or output. Pro multiple-recipient support remains
future work.

See `CHANGE_NOTIFICATIONS.md` for supported events, localization, redaction,
duplicate behavior and operation-versus-delivery result semantics.

After validation, guided setup may activate the shipped user unit only after
explicit confirmation. It uses the same validated effective-UID user-runtime
context as readiness and performs only fixed systemd user-manager operations.
Failures identify runtime validation, manager reachability, unit reload, or
enable/start without raw command output; configuration remains preserved and
the displayed QWSG assessment/resume action is safe to follow. Noninteractive
setup never enables or starts the unit.
