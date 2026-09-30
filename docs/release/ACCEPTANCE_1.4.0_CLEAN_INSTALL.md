# Community 1.4.0 clean-install acceptance — Task 095

## Scope and authority

Owner-authorized clean installation on the designated OVH playground, separate
from production and historical migration/recovery acceptance. Clean-install acceptance PASS on 2026-09-30; final repository delivery is
recorded in the matching Task 095 history.

## Release trust and public acquisition

Tag v1.4.0 points to source 996fb90d0f53d4ae6088cd884a654d9d0a32fff9.
Both release-operator reconfirmation and target acquisition independently use
the official public Forgejo artifact and canonical production index:

- [Official artifact](https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v1.4.0/qwsg-1.4.0-linux-amd64.tar.gz): 3680796 bytes,
  SHA256 26cbf8d76fa27c652e11ee49f4570e05b3be6b969309e3cfd486977a27a6224d.
- [Production index](https://releases.quantumwizard.hu/qwsg/v1/release-index.json): 1153 bytes,
  SHA256 fbb39b52e240ff3691090a48a1e7915e516da6fb2f6a1faa9ed5fc069b633c41.
- Ed25519 verification PASS with pinned qwsg-community-release-2026-01;
  fingerprint 0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6.

The separate trusted verifier matched its frozen identity before transfer and on
target. No locally built release binary supplied installation evidence. Target
artifact length/hash were bound to authenticated release semantics before safe
extraction. Sidecar, full manifest, source/platform/build identity all pass.
Published 1.4.0 and historical 1.3.1 remained immutable.

## Supported real-system baseline and installation

Ubuntu 24.04.4 LTS, amd64, kernel 6.8.0-138-generic, systemd 255, ordinary user,
ext4 local filesystem, functioning user manager and existing lingering. Preflight
found official installed 1.3.1 plus prior configuration and private evidence.
The Owner approved an exact hash-bound reset. It stopped/disabled only QWSG,
archived and relocated the five active QWSG paths without deleting private data.
Independent comparison verified 49 preserved files, ownership and modes. Normal
installed/config/state paths and enablement were absent before installation.
Historical archives/source/test trees and earlier private backups remained
outside product discovery. No unrelated configuration or security setting changed.

The target downloaded 1.4.0 through the public release path. Official read-only
installation assessment passed all required prerequisites. Documented expert
path installed the verified archive with sudo install.sh, then fresh ordinary-user
setup --accept-defaults. No replacement/migration flags used. All 20 installed
immutable artifacts match the authenticated package and root ownership, 0755/0644 modes.
Private configuration/state follow ordinary-user 0700/0600 requirements.
Default 5m interval, 2m timeout and 64MiB Go soft limit/128MiB cgroup/32 tasks/25% CPU
are retained. Guardian activation is separately explicit.

## Required behavior and bounded qualifications

Required OS, CPU, memory, storage, network and systemd/services collectors are
available. Capability metadata and ordinary-user privilege provenance remain
explicit. Optional components detects no Go executable at its fixed allowlisted
path; the prebuilt product requires no Go, and no dependency was installed to
hide that unavailable coverage. Inventory exit 2 means partial but usable.

Two successful manual observations establish fresh baseline/comparison. Native
snapshot info/load verifies evidence integrity; read-only operator view succeeds.
Guardian exclusion rejects a competing observe with guardian_active. Repeated
Guardian/scheduler envelope integrity and private modes pass. Privacy-safe
service identity continuity is checked across original and Guardian snapshots.

Installed release awareness authenticates the real production index and reports
verified_supported_installation/current/equal 1.4.0 with exact release artifact
authority. Community retains manual update policy and no automatic installation.
Documented version command passes; --version is unsupported by this release.

Installation/environment/configuration/Guardian readiness domains are ready.
Optional notification is disabled/missing_optional, so composite overall readiness
is partial as documented. Preflight sends no SMTP message; notification queue
remains empty. This is not an external mail-delivery acceptance claim.

## Stability evidence

PASS: 17:46:57 through 18:07:11 UTC, over 20 minutes and 15 samples.
Same Guardian PID, zero restarts; three successful Scheduler executions in the
window, four by final verification. Peak memory 24653824 bytes below 128MiB,
maximum 8 tasks below 32, all cgroup OOM/max events zero. Final scheduler state
91146 bytes below 8MiB and four results below retention 64. Six snapshots pass
native integrity checks. All 19 original service pseudonyms remain stable;
a transient fwupd service addition/removal is a real host observation, not an
identity change. Four distinct sampled completed cycles and operator IDs prove
continued publication. No journal panic/fatal/OOM/failure or QWSG listener.
Release awareness remains authenticated current/equal, zero failures; empty
notification queue. Private sampling evidence and final summary remain on target
outside Git. The prior preservation archive remains verified and unchanged.

## Gates and exclusions

Community remains 8 PASS / 0 PARTIAL / 1 MISSING. C9 remains MISSING. This task never
performs official 1.3.1->1.4.0 authenticated migration/bootstrap, state preservation
across migration, rollback/recovery acceptance, production mutation, release
publication or Task 096. These remaining C9 proofs require a fresh later session.

## Actual deferred finding

The prior 1.3.1 scheduler main state was 24873413 bytes, above the current 8MiB
limit, with historical temporary files. All prior bytes are preserved in the
approved private archive/quarantine. No historical-state investigation or
remediation is claimed; the clean install starts fresh.
