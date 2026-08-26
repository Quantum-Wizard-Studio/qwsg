# QWSG 1.1.0-rc.6 Targeted Practical Acceptance

## Verdict

- Candidate: `QWSG 1.1.0-rc.6`
- Source commit: `25a30718bc92882e9773a5c405ad648c0eee1a81`
- Targeted practical product acceptance: **PASS**
- QWSG-059-F001: **EXTERNALLY CORRECTED**
- Product-source defects found in RC.6: **NONE**
- Technical recommendation: **READY FOR RELEASE DECISION**
- Publication state: **NOT RELEASED**

This verdict authorizes no tag, Forgejo Release, upload, publication,
deployment or announcement. Those remain separate Project Owner decisions.

## Frozen candidate identity

Two isolated canonical builds produced byte-identical archives and sidecars.
Exactly one candidate was selected and frozen before external use.

- Archive: `qwsg-1.1.0-rc.6-linux-amd64.tar.gz`
- Archive size: `2951540` bytes
- Archive SHA-256:
  `31161ad30bbb83b76e99744512355872dd765700bbe428a1ba0b02d157a7b567`
- Sidecar SHA-256:
  `a12869185065aaa070dba3a5577a8536b7bc5228bdeabe0f81218e775c6ef375`
- `MANIFEST.sha256` SHA-256:
  `d1b9fb1e24b2b4c8053f8efb2188a803db99b7b4ab98f6d505a0210de70ee7c9`
- Binary SHA-256:
  `dcfd2f30dedfaf67dadafc86c04635a9fe9c4ec0bd801711f2da1a0ae1182b5a`
- Controlled source epoch: `1787698860`
- Controlled build time: `2026-08-25T23:01:00Z`

Source, local twins, private transfer and VPS destination identities matched.
The archive had one safe unique 25-member root containing only regular files
and directories. All 18 manifest entries, required documentation, LICENSE,
modes, binary version and source provenance passed.

## External evidence

The existing disposable Ubuntu 24.04 amd64 VPS began with the exact RC.5
binary/unit, enabled active Guardian, lingering, safe configuration/state and
protected credential storage. Its readiness preserved the known
QWSG-059-F001 state. No RC.6 bytes were present before verified transfer.

RC.6 Smart Install passed on Ubuntu 24.04 amd64 with glibc 2.39, ordinary-user
execution, running systemd user manager, local filesystem semantics and ready
installation/environment dependencies. Documented replacement used a new
private rollback backup, preserved configuration/state/credentials and prior
enabled/active intent, and installed every RC.6-owned artifact byte-for-byte.

The supported CLI changed one non-secret value from Guardian interval `5m0s`
to `4m0s`, changing the effective configuration identity. No runtime state was
manually repaired. One controlled SMTP test was accepted and the Project Owner
confirmed actual mailbox receipt; no protected identity or credential entered
the evidence.

The physical boot changed. The user manager, default target and Guardian all
started before interactive login. The current invocation owned an active
checkpoint for the changed configuration, completed a current-generation
cycle, and published current complete Guardian-running canonical state.
Readiness reported canonical evidence satisfied and Guardian ready. This is
external correction proof for QWSG-059-F001.

The reboot did not naturally exercise `Restart=on-failure`. After a private
state snapshot, the one authorized systemd-managed main-process failure
injection caused exactly one automatic restart, a new process/invocation and
`NRestarts=1`. The recovered invocation owned the current configuration and
checkpoint, completed execution, retained current Guardian-running evidence,
and readiness continued to report canonical evidence satisfied, service
enabled/active and Guardian ready. No injection was repeated.

## Notification and composite readiness contract

`qwsg notification test` is an immediate controlled transport test. It does
not create a Guardian monitoring incident or persist a Guardian queue entry.
`notification.external=satisfied` is reserved for accepted/delivered Guardian
monitoring-queue evidence. An empty queue therefore truthfully remains
`unknown_requires_verification`; working Guardian core may be ready while
notification and overall are partial.

Accordingly:

- controlled SMTP acceptance: **PASS**;
- Owner-confirmed actual receipt: **PASS**;
- `notification.external=unknown_requires_verification`: **EXPECTED**;
- `notification=partial`, `overall=partial`: **EXPECTED**;
- persistent notification satisfaction or literal overall-ready: **NOT A
  CURRENT PRODUCT-CONTRACT REQUIREMENT**.

`QWSG-061-F001` is classified **EXPECTED PRODUCT BEHAVIOR WITH ACCEPTANCE
REQUIREMENT ADJUSTMENT**, not a product defect.

## Harness reconciliation and limitations

A literal shell convergence predicate later returned FAIL while the semantic
product assessments simultaneously reported current configuration/generation
ownership, completed Guardian execution, fresh canonical evidence, enabled and
active service, Guardian ready and readiness exit zero. The predicate combined
file-timing and literal JSON assumptions stronger than the supported semantic
readiness contract. It is classified an acceptance-harness limitation and is
not used to override the product-owned assessment.

At Project Owner direction external testing then terminated. No further VPS
restart, polling, diagnostics, notification, uninstall or reinstall was run.
The explicit restart and RC.6 external uninstall/same-candidate reinstall
checks are therefore **NOT EXECUTED BY OWNER TERMINATION**, not PASS. Their
absence is a disclosed non-product evidence limitation. Candidate packaging,
installer/uninstaller behavior and preservation contracts remain covered by
the unchanged deterministic release plumbing and local validation; no observed
RC.6 product defect or security/privacy regression requires another candidate.

## Final state

Latest authoritative product evidence:

- `guardian.canonical_evidence=satisfied`;
- `guardian.service_active=satisfied`;
- `guardian.service_enabled=satisfied`;
- `guardian_service=ready`;
- `filesystem.local_semantics=satisfied`;
- `systemd.lingering=satisfied`;
- readiness exit `0`;
- notification and overall partial under the documented contract.

RC.6 is technically suitable to proceed to the separately authorized QWSG
1.1.0 release decision. QWSG 1.1.0 is not released by this acceptance.
