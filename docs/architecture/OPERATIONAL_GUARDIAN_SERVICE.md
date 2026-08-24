# Operational Guardian Service

QWSG 1.0 supports one local operational shape: a foreground `qwsg guardian
run` process supervised by the systemd user manager. The process remains the
invoking ordinary user. systemd owns start, stop, restart, enablement and
process supervision; the Canonical Runtime Service remains the sole owner of
fixed-rate recurrence and the Runtime Engine remains the sole owner of a
cycle.

The adapter resolves Configuration Source 1.0 records through the Canonical
Configuration Contract. Its built-in source enables one five-minute local
`observe` schedule. Runtime Service cadence is five minutes, cycle timeout is
two minutes, and lifecycle freshness is ten minutes. The per-user primary file
is automatically discovered. Duration overrides are temporary Configuration
Sources, so actual recurrence and Effective Configuration cannot disagree.
`--config` replaces discovery; `--config-source` is a compatibility alias.

Private state is rooted at `$QWSG_STATE_DIR`, then `$XDG_STATE_HOME/qwsg`, then
`$HOME/.local/state/qwsg`. It contains Scheduler state/lock, inventory
snapshots, Current Operator State, a nonblocking Guardian operation lock, and
one integrity-protected atomic checkpoint. The checkpoint retains only the
last validated Runtime, Alert and Notification proposed states and launch
correlation. Missing state is a first start; invalid, incompatible or
invalid state fails closed and is not overwritten. A deliberate valid
configuration identity change begins a fresh Runtime/Alert/Notification state
epoch while preserving inventory and Current Operator evidence.

Guided activation additionally requires that this resolved root match the
packaged systemd user service's `%S/qwsg` root. Before any user-manager probe,
daemon reload or enable/start operation, QWSG uses the canonical private-root
primitive to create or validate the directory as a real current-user-owned
mode-0700 path with no symlink component. Existing unsafe state fails closed
without removal, replacement, chmod, chown or migration. This precondition
ensures systemd encounters the real state directory instead of applying its
same-name configuration-directory compatibility symlink behavior.

Configuration resolves before lock or checkpoint mutation. The unit validation
condition skips process start for invalid input, avoiding restart churn without
weakening the sandbox.

Checkpoint `active` records recovery intent and launch correlation; it is not
a process-liveness oracle. A running operator claim requires fresh validated
Runtime Service lifecycle evidence.

Runtime Service evidence supplies the exact validated proposed Service State.
The operational publisher projects that state and completed typed Scheduler
traces through the existing Operator Presentation Model into Current Operator
State. The Console only reads and renders it. A stale operational observation
becomes `unavailable`; it never remains `running`. A systemd `ExecStopPost`
adapter accepts only an allowlisted termination category and the matching
`INVOCATION_ID`; it can demote an active claim but cannot promote one.

One-shot `qwsg observe` and the Guardian share the same nonblocking operation
lock. While the Guardian owns the state writers, a one-shot observation fails
with `guardian_active` instead of racing or replacing Guardian evidence.

The supported unit uses `Type=simple`, SIGTERM, bounded stop and restart
policies, `UMask=0077`, `NoNewPrivileges`, private temporary storage, a
read-only system/home view with only the QWSG state exception, `MemoryMax=128M`,
`TasksMax=32`, and `CPUQuota=25%`. There is no PID file, self-daemonization,
watchdog, network listener, privileged collector, or internal restart loop.

Alert records remain local with the default empty Notification policy. The
absence of a transport never invents delivery. Abrupt termination may leave a
claim only until correlated exit reporting or its fixed freshness deadline;
after that the Console reports `unavailable`.
