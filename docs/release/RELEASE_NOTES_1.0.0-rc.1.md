# QWSG 1.0.0-rc.1 Release Notes

This private technical Release Candidate integrates QWSG's canonical Inventory-to-Report evaluation, operator projection/current state, local Console, Scheduler/Alert/Notification contracts, Runtime Service, and continuously supervised local Guardian.

The RC adds deterministic linux-amd64 assembly, SHA-256 manifests, a fixed-prefix safe installer/uninstaller, explicit support/security/upgrade boundaries, and bounded Console redraw. It adds no new engine semantics, provider transport, network service, remediation, fleet feature, licensing enforcement, or AI.

Task 039 hardens the local RC after live-host acceptance: large Policy Reports
now use bounded aggregate Alert evidence without losing Report traceability;
Console `r` reloads Current Operator State without competing with Guardian;
privacy-safe Runtime causes reach the operator; repeated Attention meaning is
correlated; and stale lifecycle evidence cannot preserve a running claim after
unexpected termination.

Public licensing, signing, Git tagging and external publication require separate Owner authority.
