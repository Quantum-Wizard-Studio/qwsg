# QWSG 1.2.0 Release Notes

QWSG 1.2.0 promotes the externally accepted RC.7 implementation unchanged to the final release identity. It provides guided installation and setup, deterministic release discovery and update/rollback, localized lifecycle notifications, synchronized Guardian readiness, and bounded Guardian operation on the supported Ubuntu 24.04 LTS amd64 platform.

The final release retains `MemoryMax=128M`, `TasksMax=32`, `CPUQuota=25%`, the hardened user service, and `GOMEMLIMIT=64MiB`. The measured memory remediation releases completed scheduler execution graphs after publication and gives the Go runtime a cgroup-aware soft pacing limit while preserving independently enforced systemd limits and observable OOM behavior.

The update registry explicitly declares a fail-closed `1.2.0-rc.2 -> 1.2.0` compatibility path. It performs package replacement with the final Guardian unit while preserving configuration, protected SMTP credentials, Guardian/Scheduler/operator state schemas, service intent and deterministic rollback metadata. Unsupported routes remain rejected before mutation.

Project Owner-supplied final acceptance covered deterministic RC.2 update, exact rollback and re-update; guided installation on a preserved-state host; reboot autostart; readiness synchronization; bounded resources without restart or OOM; coexistence with the production-like Hestia stack; STARTTLS SMTP acceptance; and manual mailbox receipt of test and localized lifecycle notifications. See `ACCEPTANCE_1.2.0.md` for the canonical evidence boundary and final publication ledger.

QWSG remains a local, non-root Community Guardian. Optional email delivery depends on operator configuration; a working core with that optional capability disabled or unconfigured truthfully reports Partial readiness.
