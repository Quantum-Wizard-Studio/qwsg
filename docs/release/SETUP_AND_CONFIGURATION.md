# QWSG Setup and Configuration

## Normal path

Run setup as the ordinary user who will run Guardian:

```sh
qwsg setup
qwsg config show
qwsg config validate
```

For deterministic automation, use `qwsg setup --accept-defaults`. Repeat setup
preserves a valid configuration unless an explicit `--set KEY=VALUE` is given.
Setup displays its destination and proposed effective values before an
interactive write. It never enables or starts a service.

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

The Task 045 mutable keys are `locale`, `time_zone`, `snapshot_retention`,
`guardian.interval`, and `guardian.cycle_timeout`. Durations use forms such as
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

## Secrets and future notifications

The general JSON file stores references only, never passwords, API keys,
tokens, or credentials. A future protected credential provider/store will
resolve those references without changing this public contract.

Task 045 sends no email. The model reserves a future ordered recipient
collection: Community basic local email will support exactly one administrator
recipient without a QWS account, API key, subscription, or QWS remote service;
Pro will have no entitlement recipient cap, subject to global safety bounds.

After validation, install and activate the shipped user unit explicitly. Setup
does not enable or start it.
