# Configuration Activation and Setup Architecture

## Decision

Task 045 activates the existing Canonical Configuration Contract 1.0 through
one per-user filesystem adapter. The Source Record remains the sole persisted
semantic contract; setup, CLI, Guardian, automation, and a future UI are
adapters to it and may not create another configuration universe.

The default primary file is `$XDG_CONFIG_HOME/qwsg/config.json`, or
`$HOME/.config/qwsg/config.json` when `XDG_CONFIG_HOME` is unset. A nonempty
unsafe XDG path fails instead of being ignored. `--config ABSOLUTE_FILE`
explicitly replaces discovery. The legacy Guardian `--config-source` spelling
is a compatibility alias; supplying both selectors is rejected.

## Resolution and compatibility

Precedence is compiled defaults, the primary local file, an explicitly selected
file replacing discovery, then narrowly typed command-temporary overrides.
Temporary Guardian duration flags are Source Records and therefore affect both
reported Effective Configuration and actual recurrence. Equal-precedence
conflicts remain errors.

Unknown fields, malformed or trailing JSON, unsupported contract/model
versions, invalid identities, invalid references, and material bounds fail
closed. Configuration 1.0 is never silently migrated or rewritten. Additive
future behavior requires an explicitly compatible contract/version decision.

## Filesystem transaction

The configuration directory is an ordinary current-user-owned mode-`0700`
directory. The file is an ordinary current-user-owned mode-`0600` regular file.
Every existing path component is checked with `lstat`; symlink components and
special files are rejected. Reads are bounded and revalidate file identity.

Writes use a mode-`0600` same-directory temporary file, complete write, file
sync, atomic rename, and directory sync. Failures before rename preserve the
previous file. A directory-sync failure is reported and can expose only the old
complete file or the new complete file, never accepted partial JSON.

Configuration is operator intent. Runtime evidence remains below the existing
XDG state root. They are not interchangeable and uninstall preserves both.

## Secret and future-product boundary

Public configuration can contain only typed opaque secret references. Future
credential material belongs to a separate private credential-provider/store
boundary below the QWSG per-user configuration domain; Task 045 implements no
backend and accepts no secret value.

An inert optional extension can preserve a future ordered recipient collection
without activating it. Basic local email is a future Community capability with
exactly one administrator address and no QWS account, API key, subscription, or
QWS remote service. Pro has no entitlement-imposed recipient limit, subject to
global parser and resource bounds. Delivery and entitlement enforcement are
not Task 045 behavior.

## Guardian failure boundary

Guardian resolves and validates configuration before creating locks, state, or
checkpoint changes. Invalid configuration emits the bounded
`guardian_configuration_invalid` diagnostic and never claims health. The user
unit uses `ExecCondition=qwsg config validate --format json`; invalid input
skips process start instead of entering the restart policy. The sandbox and
resource limits are unchanged.
