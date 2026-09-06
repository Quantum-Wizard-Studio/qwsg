# QWSG Community 1.3.0

QWSG Community 1.3.0 adds authenticated update awareness to the local Guardian.
Official release metadata is verified against the bundled Ed25519 Community
trust anchor, and installed binary/package provenance is checked independently.

- `qwsg update check` explicitly refreshes awareness. `qwsg update status`
  reads only local state and performs no network request.
- Guardian checks the official release index every 24 hours when due, in an
  isolated bounded task. A failed release check does not stop monitoring.
- With the existing notify policy and Community SMTP configured, an
  authenticated supported newer release becomes eligible for a localized email.
  Successful delivery is persisted and deduplicated across repeated checks and
  Guardian restarts; failed deliveries remain bounded and retryable.
- Awareness checking downloads metadata only. Artifact acquisition and
  installation still require the operator to invoke `qwsg update` explicitly.
- Scheduler history retains at most 64 results. State larger than 8 MiB is
  rejected before decoding, avoiding legacy-history allocation amplification.
  Guardian retains its 128 MiB memory and 32-task limits and 64 MiB Go memory target.

The supported upgrade is 1.2.0 to 1.3.0. Configuration, SMTP credential
references/storage, scheduling, supported local evidence and service intent
are preserved by the existing update/rollback workflow. Unknown or oversized
Scheduler state fails closed without being deleted; retain a private backup
and resolve it before relying on healthy scheduled execution. Never remove
legacy state merely to make installation appear successful.

For the first production acceptance, the installed 1.2.0 instance received an
Owner-authorized compatibility-route backport. It is distinct from the
unchanged historical official 1.2.0 archive. Canonical 1.3.0 supersedes that
private compatibility baseline.

Community requires no registration, API key or installation ID. There is no
telemetry, inventory upload to release infrastructure, new inbound listener,
automatic artifact download from awareness, or automatic update installation.
Existing explicit SMTP notification policy remains operator controlled.
