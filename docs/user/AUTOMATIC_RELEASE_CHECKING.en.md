# Automatic Release Checking

QWSG Guardian checks the official Community release index automatically when
release awareness is due. The default interval is 24 hours. On first use, a
missing awareness record is due after Guardian completes its first local
monitoring cycle. Later checks are due 24 hours after the last recorded attempt.
A restart before that deadline waits for the remaining time and does not repeat
the request.

The request is credential-free and goes only to the configured QWSG production
HTTPS release-index source. QWSG enforces the release media type, validates the
signed index with its bundled Ed25519 trust anchor, and applies rollback and
future-index protection. The request contains no installation ID, hostname,
inventory, account, API key, email address, or telemetry payload. As with any
HTTPS request, the destination and network can observe ordinary connection
metadata such as the source IP.

A failed network or authentication attempt is recorded only as a bounded local
failure category. Guardian monitoring continues, and there is no immediate
retry loop. Automatic checking never downloads an artifact, stages or installs
an update, or restarts QWSG.

When `update.policy` is `notify` and Community email is enabled and valid, a
Guardian check that proves an authenticated newer applicable stable release
sends one concise operator email. The message identifies both versions, the
stable channel and authenticated metadata, links the canonical release source,
and states that installation is not automatic. Successful SMTP acceptance is
persisted against the authenticated release version, artifact digest and
signing identity; the same release is therefore suppressed after later checks
and Guardian restarts. A different authenticated newer release is eligible
again.

Disabled email or `update.policy=manual` does not attempt delivery. Delivery
failure never makes Guardian unhealthy and is not recorded as success; retry is
possible only after a later scheduled release check, so there is no tight retry
loop. No SMTP credential or destination is stored in release-awareness state.

Use `qwsg update check` for an immediate authenticated manual refresh; it does
not send update notifications. Use `qwsg update status` to read the persisted
result without any network access or notification.
Installation remains the separate explicit `qwsg update` operation.
