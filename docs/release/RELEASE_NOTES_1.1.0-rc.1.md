# QWSG 1.1.0-rc.1 Private Acceptance Candidate

This identity is a private, non-published acceptance candidate for the
post-1.0 Community productization delivered by Tasks 045–048. It is not a
public release, tag, announcement, or authorization to publish QWSG 1.1.0.

The candidate adds the operator-facing capabilities already integrated on the
canonical source branch:

- canonical per-user setup and configuration with deterministic validation;
- one-recipient Community SMTP notification, protected credential storage,
  preflight, and controlled test delivery;
- read-only Smart Install assessment and evidence-backed composite readiness;
- resumable guided setup, explicit ordinary-user Guardian activation, bounded
  fresh-evidence verification, and actionable next-step guidance;
- operator-first README and installation documentation installed under
  `/usr/local/share/doc/qwsg/`.

The supported host remains Ubuntu 24.04 LTS on amd64 with systemd 255+ and an
ordinary non-root runtime user. QWSG does not install packages, invoke sudo,
enable lingering, provision SMTP infrastructure, or silently mutate the host.

Task 049 requires a reproducible candidate built twice from one exact clean,
integrated Git commit. External clean-host, real SMTP receipt, logout, physical
reboot, uninstall, and reinstall evidence remain mandatory acceptance gates.
Until those gates pass, the candidate is not release-ready.
