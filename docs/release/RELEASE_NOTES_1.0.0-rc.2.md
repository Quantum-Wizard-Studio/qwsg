# QWSG 1.0.0-rc.2 Release Notes

This private technical Release Candidate refreshes the accepted QWSG 1.0 local
Guardian after development-host acceptance of Task 039. It retains the frozen
Inventory-to-Report pipeline, local Console, Runtime Service, systemd user
service, installer/uninstaller, supported platform and artifact structure from
RC.1.

RC.2 includes the accepted release-blocking corrections:

- large canonical Policy Reports use one bounded aggregate Alert evidence
  reference while the Report retains complete source traceability;
- Runtime reaches Notification planning instead of becoming partial solely
  because a valid Report has hundreds of sources;
- interactive Console `r` reloads and freshness-requalifies Current Operator
  State without starting `observe` or competing for the Guardian lock;
- bounded privacy-safe Runtime component causes reach the operator view;
- repeated Attention meaning is correlated without losing source traceability;
- graceful and unexpected process termination cannot preserve a running claim
  beyond validated lifecycle freshness.

RC.2 adds no product feature, provider, network interface, remediation,
Dashboard, API, fleet capability, licensing enforcement or AI. Clean Ubuntu
24.04 installation and reboot acceptance remain the Owner's next separate
gate. Public licensing, signing, Git tagging, upload and publication require
separate Owner authority.
