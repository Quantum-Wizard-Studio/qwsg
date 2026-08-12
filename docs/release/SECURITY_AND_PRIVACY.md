# QWSG 1.0 Security and Privacy

The Guardian runs as the ordinary user. The systemd unit applies `NoNewPrivileges`, private temporary storage, read-only system/home protection with one private state exception, a restrictive umask, and CPU, task, and memory limits. QWSG opens no network listener, executes no remote command, performs no remediation, and stores only secret references—not secret values—in canonical configuration.

Installation privilege is limited to copying checked artifacts under `/usr/local`; it never runs the Guardian as root. State and snapshots can describe the host and must remain private. Checksums detect corruption but do not authenticate an untrusted distribution channel. The QWS Community / Free License governs use and redistribution; it does not replace artifact-authenticity or publisher-trust controls.

Per-user configuration is separate from state, current-user-owned, and exact
mode `0700/0600`. QWSG rejects symlinks, special files, wrong owners, and
permissive modes, and publishes changes atomically. General configuration has
only opaque secret references. Task 045 adds no credentials, network client,
SMTP transport, or entitlement lookup.
