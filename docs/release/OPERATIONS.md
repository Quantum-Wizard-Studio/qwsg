# QWSG 1.2 Operations

## Native updates

Run `qwsg update check` without privilege. `qwsg update` performs anonymous
discovery, private staging and full verification before it stops Guardian or
invokes the narrow privileged helper. After replacement QWSG restores previous
enabled/active service intent and validates installed identity and
configuration. Inspect rollback availability with `qwsg update status`; use
`qwsg update rollback` to restore the prior verified package artifacts.

`qwsg update check` contacts only
`https://releases.quantumwizard.hu/qwsg/v1/release-index.json`, requires the
bundled `qwsg-community-release-2026-01` Ed25519 public identity, and fails
closed on transport, source, schema, signature, key, or artifact-metadata
failure. It neither downloads nor installs an artifact. `qwsg update status`
remains network-free and preserves the last authenticated success separately
from a later failed attempt.

The running Guardian performs the same authenticated awareness operation when
it is due, by default once every 24 hours. It waits until the first local
Guardian cycle has completed. The due time is the last recorded attempt plus
24 hours, so restarting Guardian shortly after a check does not repeat the
request. A failed attempt is non-destructive and is not retried until the next
interval; local monitoring continues. Automatic checking stores awareness
only: it does not download, stage, install, restart, or notify. Update
notification and deduplication are reserved for Task 080.

Use `qwsg` for the current read-only operator view and `qwsg observe` for an explicit full observation. The supervised Guardian runs the same canonical Runtime Service; it does not duplicate engine decisions.

Use `qwsg readiness` for the composite operational gate. Guardian core may be
`ready`, notification `not_ready`, and overall `partial`. A ready Guardian claim
requires fresh canonical evidence, not merely an installed or active unit.
Guided activation uses the same bounded fresh-evidence waiter as guided setup
and requires the post-activation evidence identity to differ from any preserved
pre-activation record.

Before activation or after a configuration change, run `qwsg config validate`
and `qwsg config show`. Guardian discovers that primary configuration and uses
its effective interval and timeout. Duration flags are temporary
highest-precedence configuration sources for bounded testing only.

Inside the interactive Console, `r` only reloads the integrity-checked Current
Operator State and requalifies freshness. It does not compete with the running
Guardian. An explicit `qwsg observe` remains lock-protected and fails safely
with `guardian_active` while Guardian owns the operation.

```sh
systemctl --user status qwsg-guardian.service
systemctl --user start qwsg-guardian.service
systemctl --user stop qwsg-guardian.service
systemctl --user restart qwsg-guardian.service
journalctl --user -u qwsg-guardian.service --since today
```

`Guardian: running`, `stopped`, `degraded`, or `unavailable` is derived from integrity-checked, freshness-bounded lifecycle evidence. A one-shot observation, PID, unit file, or old state never proves that the Guardian is running. Default private state is under the systemd user StateDirectory (`%S/qwsg`, normally `~/.local/state/qwsg`) with mode `0700`; records are mode `0600`.

Boot before login requires an administrator to authorize lingering for the exact runtime user (`loginctl enable-linger USER`). QWSG never changes this setting. Without lingering, the enabled user unit normally starts when the user's manager starts after login.
