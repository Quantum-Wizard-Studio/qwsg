# QWSG 1.2 Troubleshooting

## Native update failures

Acquisition or verification failures occur before installed files or Guardian
state change. A package replacement or post-install validation failure triggers
an automatic restoration attempt. Do not manually overwrite `/usr/local`;
preserve the private rollback root and use `qwsg update status` or
`qwsg update rollback`. Missing, tampered, symlinked or incomplete rollback
records fail closed.

- `missing_required`: resolve the blocker and rerun `qwsg install --check` or
  `qwsg readiness`.
- `unknown_requires_verification`: manually verify the named condition; QWSG
  deliberately does not guess a package command. Read the finding's
  Explanation, Verify, Operator action, Privileges, Safety, and Revalidate
  lines; JSON exposes the equivalent `guidance` object.
- `incompatible`: the observed host violates the supported boundary.
- `overall: partial`: core operation can be ready while optional notification
  remains unavailable or unverified.
- A guided installer fresh-evidence timeout exits 4 before the 100% summary,
  preserves package/configuration/state, and directs the operator to `qwsg
  readiness` and `qwsg setup`. It does not mean package installation failed.

- `systemd.user_manager`: perform the displayed ordinary-user verification.
  QWSG distinguishes a missing/unsafe runtime directory, transient manager,
  unreachable manager, timeout, oversized response, and unknown state. It does
  not prescribe package, PAM, lingering, or service changes from ambiguous
  evidence. After the verified host action, rerun `qwsg install --check`.
- `filesystem.local_semantics`: default assessment is read-only. If the
  filesystem cannot be proven, verify that the QWSG configuration and state
  locations support atomic rename, advisory `flock`, Unix ownership/modes, and
  private `0700/0600` storage, then rerun the assessment.

- Guided Guardian activation failures name one fixed stage: user runtime-context
  validation, user-manager reachability, systemd user-unit reload, or Guardian
  enable/start. Configuration remains preserved. Follow the displayed
  `qwsg install --check` or `qwsg readiness` evidence and resume with
  `qwsg setup`; do not substitute an inferred systemctl command. Timeout and
  output-limit causes are intentionally distinct from a fixed-operation
  failure, and raw command output is withheld.

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

- `notification_not_ready`: run `qwsg notification preflight`, correct required
  or incompatible findings, and rerun it. Unknown capability is not guessed.
- `smtp_delivery_failed`: verify destination, TLS trust, authentication, and
  network externally. Guardian monitoring continues.
- `credential_path_unsafe`: use the canonical current-user-owned `0700`
  directory and a `0600` regular credential file; links are rejected.
