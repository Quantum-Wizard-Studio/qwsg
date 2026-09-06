# Community 1.3.1 production acceptance — Task 081

The corrective release completes Task 081 following the real 1.2.0 → 1.3.0
upgrade. Historical 1.3.0 was released, but its post-install acceptance exposed
two regressions; it was not retroactively declared accepted or rewritten.
The Owner-authorized 1.3.1 addresses those regressions only.

## Authenticated release identity

- Source: `45009fe6169bff00842a4c4e9561bf339a5db81e`
- Annotated tag: `v1.3.1`, object `efb7afcac7bd65a3419f6bf3264892e09de631f6`
- Forgejo Release: [5](https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/tag/v1.3.1)
- Artifact: `qwsg-1.3.1-linux-amd64.tar.gz`, 3600553 bytes
- Artifact SHA-256: `0ab726bcde36182232ff89e3ea33e2d5ab77d8cad2b94ce56c9534a8147d9cd7`
- Signed production index: 913 bytes; SHA-256 `2eefd328c5ded8807101bdd69d5fcee0ed4f4f2e49b10cdd02b55da6e7aaf879`
- Key ID: `qwsg-community-release-2026-01`
- Trust fingerprint: `0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6`

Two isolated builds from the exact source commit were byte-identical. Public
curl/wget artifact retrieval, sidecar, canonical Ed25519 verification and installed
binary/RELEASE/service-unit comparison passed. Forgejo release metadata and
historical 1.2.0/1.3.0 artifact hashes remain unchanged; download counters are
naturally mutable. Production preserves HTTPS redirect, required media type,
no-cache, validators and empty safe conditional 304 without Expires.

## Real operator-controlled lifecycle

Initial 1.2.0 acceptance used the explicitly authorized installed compatibility
remediation, not a byte-identical copy of historical official 1.2.0. The real
signed 1.3.0 was detected and notified once, with repeated evaluation suppressed;
the Owner installed canonical 1.3.0. That superseded the temporary backport.

For the corrective migration, the final 1.3.1 Guardian core ran a bounded
pre-install driver against actual production 1.3.0, public signed 1.3.1 and the
configured SMTP provider/store. One-minute repeated evaluation sent once, then
zero times with a new notifier. This is not evidence of the resident 1.3.0
24-hour timer firing or supporting a migration compiled only in 1.3.1.
SMTP acceptance means provider acceptance, not independently observed inbox receipt.
No awareness check downloaded or installed an artifact.

The Owner invoked the verified archive's native updater for 1.3.0 → 1.3.1.
The native transaction records rollback to 1.3.0. Private configuration and
credential files remained byte-identical. Deterministic real-archive clean
installation, migration and exact rollback passed before release; healthy
production was not rolled back merely to repeat that proof.

## Separate regression proofs

1. Fresh installed-identity evaluation produced current/equal 1.3.1. Two real
   conditional HTTP 304 responses through the final production checker also
   retained current/equal. The earlier isolated real-index test corrected the
   legacy stale 1.3.0 state. Unsupported/unsafe reuse remains fail-closed.
2. The exact successful 1.3.1 notification record survived installation/service
   restart and subsequent checks. Freshly opened stores and notifier instances
   retained that identity and made zero delivery calls. Failed-refresh and
   rollback preservation are additionally covered by deterministic regression tests.

`update status` showed current/equal with zero network syscalls (trace contained
only process signals). GOMEMLIMIT=64MiB, MemoryMax=128MiB and TasksMax=32 remain
in force; no QWSG inbound listener exists. Scheduler retains 64 results below
8MiB; the separately named historical oversized backup is not the active state.
Final bounded production observation metrics and closure are recorded in
[Task 081 history](../../ai/history/081_2026-09-06_community-1-3-0-release-production-acceptance.md).

Formatting, vet, all Go tests, full race suite, engineering/framework/lifecycle,
release and release-authority reproducibility, installer/migration/rollback,
Scheduler bounds, awareness, notification and privacy/security gates passed.
No automatic installation, telemetry, registration, credential disclosure or new
listener was introduced. Full backups remain private through the rollback window.
