# QWSG 1.2.0-rc.2 Guided Installation Candidate

This private development candidate adds the QWSG-owned guided terminal
installer, truthful phase-derived progress, English/Hungarian/German catalogs,
strict Ubuntu 24.04 amd64 gating, pre-mutation planning, setup/service/readiness
integration, and explicit manual/notify update policy.

It is not published. The official public release remains QWSG 1.1.0. Existing
native update and rollback commands remain available. Migration from 1.1.0 or
private 1.2.0-rc.1 is an explicit no-mutation compatibility path for existing
configuration and state schemas.

Notification provider values remain configurable through `qwsg setup` and the
protected credential command after the installer explains the optional
capability. Notify policy records preference but does not yet run a background
notifier. Automatic privileged updates are not implemented.
