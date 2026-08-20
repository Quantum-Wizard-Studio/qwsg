# Task History 047: Approved Task

## Task metadata

- Task ID: `047`
- Task slug: `qwsg-smart-install-readiness-assessment-foundation`
- Status: `complete`
- Date generated: `2026-08-20` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/047_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded; implementation has not started.

The Project Owner started Task 047 through the canonical `job` workflow on
2026-08-20 UTC. Implementation remains limited to read-only assessment,
planning, guidance, and readiness; no remediation executor or Task 048 is
authorized.

## Starting state

Task 047 started at the verified repository root on `main`, with HEAD and
`origin/main` both `3aa332be188bbb62789c5262a23a7ef0bd3ee8b9`, ahead/behind
`0/0`, an empty index, and a clean tracked tree. The active Builder prompt and
history plus the preserved Owner draft were the only visible untracked task
state. Source SHA-256, Task 046 baseline, Framework/lifecycle/diversion/Builder,
v1.0.0 tag/target, LICENSE, release artifact, and Owner-draft preservation
checks passed. Focused/full/race/vet/format, engineering, shell/package, and
systemd-unit pre-change gates passed; system-bus access emitted the expected
sandbox warning without changing the unit or host.

## Snapshot

Implementation snapshot:
`/tmp/qwsg-task047-implementation-20260820T.OxcdzV`. The mode-0700 snapshot
contains a verified complete tracked-HEAD archive, exact active lifecycle and
Builder-source payloads, target absence records, Git/release/Owner metadata,
ACLs, checksums, archive-readability proof, and collision-aware bounded restore
instructions. The Owner draft was recorded by metadata/hash only and was not
copied.

## Work performed

Implemented the approved foundation without adding a host-mutation path:

- added `internal/assessment` Assessment Model and Registry 1.0 with the five
  canonical finding classifications, six requirement classes, four summary
  states, validation, deterministic ordering, platform gates, and inert
  structured remediation plans;
- added direct bounded OS/architecture/user evidence and a compiled
  fixed-argument runner for glibc, systemd version, user manager, unit presence,
  enabled state, and active state; no shell or caller-controlled command/argv
  exists;
- adapted Task 046 SMTP findings to the common classification without moving
  SMTP ownership, and composed disabled, invalid, configured-unverified,
  Guardian-queue verified, and delivery-failed notification states;
- added `qwsg install --check` for archive/installed preflight and `qwsg
  readiness` for installation, environment, configuration, notification,
  Guardian/service, and overall human/JSON readiness;
- distinguished unit presence, enabled, active, lingering, and fresh canonical
  Guardian evidence; service state alone never proves monitoring;
- kept `install.sh` a deterministic low-level copier and changed only its
  post-copy guidance; the bundled binary owns pre-install assessment;
- documented the dependency audit, exact Ubuntu 24.04 amd64/systemd 255+
  boundary, no-command uncertainty behavior, privacy/security model, deferred
  VPS profile, Community completeness, and future Pro Detect/Plan/Ask/Execute/
  Verify/Continue boundary.

No package-manager mapping was added because repository evidence did not prove
a safe missing-package/coexistence mapping. Exact supported recommendations are
limited to QWSG setup, release installation, and user-unit activation. No
Postfix/Exim/Sendmail/certificate package is recommended.

The first staged acceptance invocation (`attempt-failed`) ran the relative
checksum sidecar outside its archive directory, so the sidecar correctly could
not locate the archive and stopped before installation. The corrected method
verified it from its containing directory. A final assertion initially expected
an absent state directory although the fixture had pre-created that empty root;
the corrected assertion verified zero entries, proving readiness wrote no
state. Neither invocation indicated a product defect or mutated a real host.

## Verification

PASS: build; focused assessment/configuration/SMTP/Guardian/runner/CLI tests;
full `go test ./...`; repository-wide race detector; vet; formatting; Git
whitespace; Framework 21; diversion 36; lifecycle 28; Builder 38; diverted-task
audit; active-job validation; shell syntax; and static systemd unit verification.
The system-bus warning is the known sandbox limitation; no live user service was
changed or claimed.

Registry/model tests cover IDs, versions, ordering, all classifications and
summary precedence, unsupported platforms, old systemd, root/wrong architecture,
missing probes, exact recommendation mapping, unknown-platform command absence,
and hostile command metadata. Fixed-runner tests prove unknown IDs and supplied
malicious argv cannot select execution. CLI tests prove human/JSON behavior,
exit `4`, missing configuration, optional notification aggregation, and no
config/state creation.

Staged local acceptance at
`/tmp/qwsg-task047-staged-acceptance.Qgz3Fg` built a release-style archive in an
isolated output directory, verified sidecar/internal manifest, ran the bundled
preflight, installed into an isolated root, performed isolated setup and config
validation, reported truthful not-ready service/Guardian evidence, uninstalled,
preserved configuration, and left the pre-created state root empty. No package,
service, lingering, firewall, external network, control panel, clean external
host, or real SMTP provider was used or claimed.

Completion/archive/canonical-idle simulation passed at
`/tmp/qwsg-task047-completion-simulation.JTaA3B`. Snapshot manifests and archive
readability remain valid. Source scans found no package-manager execution,
generic shell/command runner, remediation executor, inbound listener, Pro/QWS
entitlement, VPS profile, or Task 048 implementation.

## Rollback

Verify the implementation snapshot manifest and archive. Restore only explicit
Task 047 targets after checking for later edits; remove only recorded created
paths with proven pre-task absence. Never use broad Git recovery or alter the
Owner draft, Task 046, LICENSE, v1.0.0, release artifacts, or published history.

## Completion state

`complete`

Task 047 is complete within its read-only foundation boundary. External
clean-host/control-panel/real-provider acceptance and any automatic
provisioning remain explicitly unperformed and require separate authority.
The implementation supports a later separately authorized 1.1.0 release line;
no version, tag, release artifact, publication, commit, push, or Task 048 was
created.
