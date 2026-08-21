# QWSG 1.1.0-rc.3 Private Acceptance Notes

This source identity prepares a future private replacement candidate for a
separately authorized clean-host acceptance. No RC.3 artifact, tag, release,
publication, or acceptance claim is created by this source change.

## Guided Guardian activation correction

- Smart Install/readiness and guided activation now share one effective-UID
  runtime-directory validator for the canonical `/run/user/<uid>` context.
- The bounded runner still replaces the ambient environment and permits only
  one canonical trusted `XDG_RUNTIME_DIR` entry for fixed operations. It does
  not inherit `DBUS_SESSION_BUS_ADDRESS`, `HOME`, or caller values.
- Guided activation checks the same user manager before the fixed
  `daemon-reload` and `enable --now qwsg-guardian.service` sequence.
- Activation failures identify the fixed stage and a privacy-safe bounded
  cause, preserve configuration, and direct the operator to QWSG assessment
  and resumable setup without exposing raw command output or environment.
- Readiness remains independent and still requires enabled/active state plus
  fresh integrity-checked Guardian evidence.

QWSG `1.1.0-rc.1`, `1.1.0-rc.2`, Task 049 and Task 051 evidence remain
immutable historical records. RC.3 clean-host acceptance must restart at
Checkpoint 01 under separate Owner authority.
