# Practical Release Acceptance Standard

## Purpose

This prospective standard answers one operational question: can a real user
install, configure, operate, reboot, uninstall, and reinstall the exact QWSG
candidate on a fresh supported Linux host? It replaces artificial per-step
approval gates with one bounded Owner-authorized acceptance run. Historical
protocols and ledgers, including Task 057 and RC.5, retain their original
meaning and are never retroactively reclassified by this standard.

## One-run authority

One explicit Owner authorization may cover all twelve steps, including their
ordinary documented host mutations. The operator reports relevant evidence as
the run progresses, but no additional approval is required between routine
steps. Pause for an Owner-only protected credential action when needed, then
resume the same run.

Stop for a product, package, provenance, integrity, security, or mandatory
operational defect; material scope expansion; an unplanned destructive host
action; unavailable rollback; unresolved privacy uncertainty; or another
Owner-reserved operation. Do not use undocumented repair or a hidden developer
workaround to manufacture a PASS. A source correction creates new bytes and
therefore requires a new candidate identity.

## Twelve-step workflow

1. **Fresh supported host baseline.** Record the supported distribution,
   architecture, ordinary-user context, required service manager, and absence
   of QWSG-specific preparation. A previous test host is acceptable only after
   a reset that technically restores this baseline.
2. **Exact candidate receipt.** Download or privately receive only the
   authorized archive and checksum sidecar. Record the candidate version and
   source identity without private host identifiers.
3. **Integrity and package safety.** Verify regular non-symlink file types,
   archive size and SHA-256, `sha256sum -c`, safe archive paths/types/layout,
   complete internal manifest, required documentation, and LICENSE.
4. **Smart Install/readiness.** Run the documented read-only assessment and
   require supported-host readiness, including actionable guidance for any
   permitted operator prerequisite.
5. **Documented installation.** Install the exact candidate through its
   packaged workflow and verify installed identity, immutable files, unit, and
   documentation.
6. **Guided setup.** Run the documented guided setup as the ordinary Guardian
   user and verify safe configuration creation without manual product bypass.
7. **Guardian and state contract.** Use guided activation and independently
   verify a real, non-symlink, current-user-owned mode-0700 canonical state
   root; safe path components; distinct configuration/state roots; enabled and
   active service; fresh canonical Guardian evidence; and satisfied
   `filesystem.local_semantics`.
8. **Protected notification.** The Owner enters credentials only through the
   protected local boundary. Verify preflight and one real Owner-confirmed
   receipt while retaining no secret, provider, recipient, or private host
   identity in canonical evidence.
9. **Physical reboot.** Reboot the host and verify automatic Guardian recovery,
   enabled/active state, and fresh post-boot evidence.
10. **Explicit restart.** Perform the documented Guardian restart and verify
    bounded recovery and fresh evidence.
11. **Uninstall preservation.** Run the documented uninstaller and verify
    release-owned removal plus every promised configuration, credential, and
    state preservation boundary.
12. **Same-candidate reinstall.** Reinstall the identical verified candidate,
    resume preserved configuration safely, reactivate as documented, and
    require final operational readiness and fresh evidence.

## Mandatory evidence

Retain privacy-safe evidence for candidate identity and checksums; host
baseline; package/manifest results; each documented product command and outcome;
state-root safety; service enabled/active state; canonical evidence freshness;
one notification receipt classification; reboot and restart recovery;
uninstall preservation; same-candidate reinstall; unresolved findings; and the
final READY or NOT READY verdict.

Evidence supports the product decision rather than becoming a second product.
Reliable evidence may be reconciled after the corresponding action or in a
different reporting order when identity, continuity, and independence remain
provable. Late reporting alone does not invalidate a clean host. Missing
mandatory evidence remains missing and is never converted to PASS. Reinstall
the host only when product execution, manual repair, candidate mixing, or other
mutation has technically destroyed the required baseline—not merely because
console output arrived late.

## Verdict and publication boundary

`READY FOR RELEASE` requires all twelve mandatory steps and every applicable
security boundary to pass with no release blocker. Otherwise record `NOT READY
FOR RELEASE` with the exact defect or missing proof. Neither verdict authorizes
a Git tag, Forgejo Release, asset upload, publication, deployment, announcement,
or the next engineering task.
