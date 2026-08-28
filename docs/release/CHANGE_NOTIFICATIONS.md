# QWSG-Managed Change Notifications

QWSG can notify the configured Community administrator about QWSG-managed
installation, update, rollback, version-transition, configuration and selected
Guardian activation changes. It does not continuously monitor arbitrary
external filesystem changes.

## Configuration

The capability reuses the existing Community SMTP configuration, one-recipient
limit, protected credential store, preflight and test command:

```sh
qwsg config set notification.email.enabled true
qwsg config set notification.lifecycle.enabled true
qwsg config set notification.email.recipient admin@example.invalid
qwsg notification credential set --from-file PRIVATE_0600_FILE
qwsg notification preflight
qwsg notification test
```

Set `notification.lifecycle.enabled` to `false` to disable change messages
without disabling Guardian alert email. Configure all non-secret SMTP fields
through `qwsg config`; passwords remain accepted only from a current-user-owned
mode-`0600` file and are never included in notifications or command output.

## Events and content

Messages use the configured locale (`en`, `hu`, or `de`) and include QWSG,
privacy-safe local hostname, event/result, UTC timestamp, operation ID,
administrator-action flag and applicable previous/new/restored versions.
Configuration messages identify the managed key, never its value.

Installation failures are reported only after canonical configuration and the
notification capability are available. Guardian messages are limited to the
explicit successful activation performed by guided installation to avoid noisy
runtime mail. Update and rollback messages show version direction.

## Result semantics and duplicate safety

The lifecycle operation result and notification delivery result are separate:

```text
QWSG update: SUCCESS
Admin notification: FAILED
```

SMTP failure never changes a successful installation/update/rollback into a
corrupted operation and never hides an underlying operation failure. QWSG
reports `ACCEPTED`, `FAILED`, `DISABLED`, or `DUPLICATE`; accepted means the
configured SMTP server accepted the message. One command process suppresses a
repeated deterministic event ID. QWSG does not retry lifecycle messages or
create notification loops.

Automated tests use deterministic in-memory transports and never contact an
external email service. Actual delivery remains a later acceptance task.
