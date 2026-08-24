# QWSG 1.1.0-rc.5 Private Acceptance Notes

This source identity corrects the systemd user-state-directory compatibility
blocker recorded as `QWSG-055-F001` without weakening QWSG state-path security.

Guided activation now securely creates or validates the canonical Guardian
state root before contacting the systemd user manager. The directory must be a
real current-user-owned mode-0700 directory reached without symlink components.
An existing symlink, wrong file type, wrong ownership or unsafe mode fails
closed before daemon-reload or enable/start; QWSG never removes, follows,
replaces, chmods, chowns or migrates unsafe state.

The packaged user unit and its `StateDirectory=qwsg`, private mode, working
directory, explicit state environment, writable-path boundary and hardening
remain unchanged. Pre-creation ensures systemd encounters the valid directory
rather than creating its user-service compatibility symlink.

RC.1, RC.2, failed RC.3 and failed RC.4 evidence remain immutable. RC.4 is NOT
READY and cannot be relabeled or rebuilt as different bytes. QWSG-055-F001
remains historical OPEN/BLOCKING after local correction and requires a new
private deterministic RC.5 candidate plus fresh clean-host acceptance from
Checkpoint 01 on a fully reinstalled disposable VPS.

This source note does not authorize candidate construction, transfer, external
execution, credentials, tagging, Forgejo Release creation, upload or
publication.
