# QWSG 1.0 Troubleshooting

- `Guardian: unavailable`: lifecycle evidence is absent, invalid, incompatible, or stale. Inspect the exact user unit and its bounded journal.
- `Guardian: degraded`: the process is operating but its latest qualified cycle or evidence is incomplete or failed.
- `guardian_active`: a supervised Guardian owns the single-writer lock; use the Console instead of racing `qwsg observe`.
- Console `r` never starts `observe`; it reloads and freshness-requalifies the
  last integrity-checked Current State. A refresh failure therefore indicates
  a genuine state load/validation problem, not normal Guardian lock ownership.
- `alert_evaluation_failed`, `notification_planning_failed`,
  `notification_delivery_failed`, `runtime_timeout`, and `runtime_cancelled`
  are bounded component causes. Inspect Guardian details; QWSG intentionally
  withholds raw internal errors.
- `guardian_checkpoint_invalid` or `guardian_state_unsafe`: stop the unit and preserve the private state for diagnosis. QWSG refuses unsafe or incompatible data rather than rewriting it.
- `operator projection failed`: canonical evaluation succeeded but the bounded operator projection was invalid. Repeating the command is not claimed to repair it.
- `current state publication failed`: inspect ownership, mode, free space, and the private state directory without printing its contents.
- `state_bootstrap_failed`: QWSG could not safely initialize its private
  per-user state root. `state_publication_failed` means valid evaluation reached
  Current State publication but the bounded write did not complete. Neither
  token means that saved state was corrupt or unreadable.
- `configuration_invalid`, `configuration_path_unsafe`, or
  `configuration_permission_unsafe`: run `qwsg config validate`. The directory
  must be current-user-owned mode `0700`, the regular file mode `0600`, and no
  path component may be a symlink. Unknown, malformed, incompatible, and
  identity-mismatched Source Records fail closed.
- `guardian_configuration_invalid`: Guardian rejected configuration before
  changing checkpoint or lifecycle state. Its systemd condition prevents
  configuration faults from entering restart churn.

Raw state can contain host evidence. Do not paste it into public reports. QWSG emits privacy-safe categories instead of raw paths, identifiers, config values, or Go errors.
