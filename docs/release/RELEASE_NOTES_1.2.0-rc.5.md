# QWSG 1.2.0-rc.5 Private Candidate Notes

RC.5 remediates release blocker `QWSG-069-F001`. The updater now selects an explicit declarative compatibility path from an installed QWSG `1.2.0-rc.2` state to `1.2.0-rc.5`; every unknown, malformed, equal, downgrade or undeclared transition still fails closed.

The RC.2 and RC.5 configuration, Guardian, scheduler and operator-state schemas are compatible. No configuration-schema transformation is required. User configuration, protected notification credentials and QWSG-managed state are outside the package replacement set and remain byte-preserved. The managed binary, Guardian user unit and packaged documentation are transactionally replaced. The privileged helper independently revalidates the declared path before mutation.

Preflight requires a valid installed version, verified candidate/package and successful validation of the installed configuration. Post-update validation requires the RC.5 binary identity and configuration validation; orchestration restores the prior package on later failure. Rollback metadata records source/target identity, commit, every replaced destination, prior existence/mode and SHA-256-protected backup. Explicit rollback restores the RC.2 package state and preserves user data.

Update and rollback continue to emit Task 068 lifecycle events through the existing transport-independent dispatcher. Operation SUCCESS/FAILED remains separate from notification ACCEPTED/FAILED/DISABLED/DUPLICATE; previous/resulting/restored versions, operation identity and administrator-action state are retained and secret-bearing reasons are redacted.

Known limitation: RC.5 supports the specifically declared historical paths, including RC.2 -> RC.5. It does not infer arbitrary candidate-to-candidate paths. Real OVH/Contabo update, rollback, reboot, coexistence and mailbox acceptance remain pending for the subsequent final-acceptance task. RC.5 is private and is not a final release.
