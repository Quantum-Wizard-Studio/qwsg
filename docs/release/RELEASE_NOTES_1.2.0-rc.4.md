# QWSG 1.2.0-rc.4 Change Notification Candidate

This private candidate supersedes private RC.3 because the Project Owner added
administrator notification as a mandatory pre-acceptance capability. RC.3's
deterministic packaging fix is retained unchanged.

RC.4 integrates localized EN/HU/DE messages for QWSG-managed installation,
update, rollback, version transition, configuration change and guided Guardian
activation through the existing Community SMTP configuration, protected
credential and provider boundaries. Operation results and SMTP delivery results
remain separate, secrets are redacted, and repeated event IDs are suppressed
within one command process.

RC.4 is private and unpublished. The complete real-host and actual-email
acceptance matrix remains deferred to a later Owner-authorized task.

Frozen source commit: `4f7dcc11b5ccc9f078755946995baebd31ad6870`.
Controlled epoch: `1787907901` (`2026-08-28T09:05:01Z`). Artifact
`qwsg-1.2.0-rc.4-linux-amd64.tar.gz` has SHA-256
`adeb591605c0d37a5fc98d541125ca388cd4561703d0f0823bba931bc7d08684`.
