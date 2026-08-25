# QWSG 1.1.0-rc.6 Source Readiness Notes

This source identity corrects `QWSG-059-F001`, the release blocker established
by the RC.5 Practical Clean-Host Acceptance.

After notification or another configuration change, RC.5 reset the Guardian's
configuration-bound checkpoint state but retained Scheduler state owned by the
previous effective configuration. A rebooted or systemd-recovered Guardian
therefore failed every Scheduler cycle at the configuration-identity boundary,
published truthful degraded/partial canonical evidence, and could not regain
readiness.

The Scheduler cycle adapter now replaces only integrity-valid persisted state
whose configuration identity is superseded. It initializes private atomic
state for the active configuration before evaluation. Same-configuration
restart recovery, interrupted-request evidence, generation-correlated exit
demotion, corrupt-state failure, safe state paths, guided setup, SMTP handling,
and protected credentials remain unchanged.

Deterministic regression coverage proves configuration-change recovery, fresh
canonical work by the recovered generation, stale-generation exit rejection,
and truthful degradation for a genuine active-generation failure.

This source is prepared for a later separately authorized RC.6 candidate build
and external clean-host acceptance. No RC.6 candidate, tag, Forgejo Release,
asset, publication, deployment, or QWSG 1.1.0 release is created or authorized
by these notes.
