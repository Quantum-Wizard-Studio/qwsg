# QWSG 1.0 Operations

Use `qwsg` for the current read-only operator view and `qwsg observe` for an explicit full observation. The supervised Guardian runs the same canonical Runtime Service; it does not duplicate engine decisions.

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
