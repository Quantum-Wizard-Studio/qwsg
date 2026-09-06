# QWSG Community 1.3.1

This corrective release addresses two regressions discovered during the real
1.3.0 production upgrade acceptance. The published 1.3.0 tag, artifact and
release metadata remain immutable.

- After installation changes, release checking fetches and authenticates the
  full index before re-evaluating installed versus available versions. A 304
  response cannot retain an earlier newer/update_available classification.
  Previously saved stale 1.3.0 evaluations also trigger a full authenticated
  fetch; equal versions become current/equal.
- Successful update-notification identity survives version changes, failed
  checks and restarts. The same already-notified release is not delivered
  again solely because the installed version changed. A genuinely different
  authenticated supported newer release remains eligible under the configured
  notify/SMTP policy.
- A failed refresh after a version change reports unknown awareness while
  retaining historical authenticated evidence for anti-rollback checks and
  the successful notification record. It does not claim the old classification
  applies to the new installation.

The supported corrective migration is 1.3.0 to 1.3.1. Because the installed
1.3.0 predates this explicit route, invoke the verified 1.3.1 archive binary's
existing updater with that archive and version. Its narrow privileged helper
revalidates provenance and compatibility and records deterministic rollback.
Configuration, credentials, supported state and Guardian service intent remain
preserved. This is an operator-controlled action, never an automatic install.

The production acceptance recovery may restore a previously lost notification
record only from the retained, integrity-verified pre-update evidence. The
product does not invent a record for an email whose successful delivery cannot
be proved. Historical release bytes and index authenticity are not altered.

The authenticated Ed25519 release contract, artifact integrity/provenance,
Guardian resource limits, Scheduler bounds and privacy boundaries are unchanged.
There is no telemetry, registration, new listener or unattended privileged update.
