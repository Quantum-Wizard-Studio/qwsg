# Task History 081: QWSG Community 1.3.0 Release & Production Acceptance

## Task metadata

- Task ID: `081`
- Task slug: `community-1-3-0-release-production-acceptance`
- Status: `active — Owner-authorized compatibility backport in progress`
- Date generated: `2026-09-06` UTC
- Human authority: Project Owner explicit Task 081 authorization and APPROVE in the current session
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/081_CURRENT_TASK.md`

## Starting state

Fetched canonical Forgejo origin; main HEAD and origin/main both
`7a9cd0ab09275b01a7c7a945e5781b95ce2a9126`, divergence 0/0, clean worktree,
valid idle lifecycle following complete Task 080. Framework 2.0 validation PASS.
Read all prompt Required Reading plus backup, release, delivery and documentation
policies, Task 078 publication runbook, relevant Task 078/080 history and native
update architecture. Builder transaction installed only Task 081 prompt/history.

Production binary and installed RELEASE.json agree: version 1.2.0, commit
`accdf93de51adc944536e7ec0d53012d24aed318`, built `2026-09-03T22:57:15Z`.
Local update status reports current installed/available 1.2.0 and successful
check at `2026-09-05T23:19:03Z`. Guardian active/running, Result=success,
NRestarts=0, MemoryCurrent=47030272, MemoryPeak=72060928,
MemoryMax=134217728, TasksCurrent=9, TasksMax=32; cgroup memory events all zero,
including oom and oom_kill. These are initial observations, not release acceptance.

## Snapshot

Protected pre-mutation snapshot `/tmp/qwsg-task081-baseline.9ypji7eh` contains
complete verified Git bundle, readable tracked-source archive, SHA256SUMS and
bounded RESTORE.md. Directory mode 0700 and original snapshot files mode 0600.
Checksum verification, bundle verification and archive listing PASS before
Builder mutation. Retain through Owner acceptance and release rollback window.
No private signing material, credentials or production runtime state captured.

## Work performed

Installed canonical approved Task 081 using task-builder.sh with the Owner's
explicit authorization. No product, version, installed package, production
configuration or public release mutation performed.

Retrieved the official public index and protected 1.2.0 artifact after the
sandbox DNS failure was resolved by the permitted escalated retrieval.
Artifact SHA-256 remains
`44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`.
Public index is 918 bytes, SHA-256
`f9f95bf28d463a8403841d9cc56d817c248f1e0a01e3e65a5a9e1afc16d39704`.
Existing publication verifier authenticated the retrieved bytes with production
key `qwsg-community-release-2026-01` and trust fingerprint
`0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6`.

## Verification and release-blocking finding

Classification: PRODUCT/FRAMEWORK DEFECT in the accepted pre-release baseline.
Source inspection at the exact installed commit proves:

1. `internal/update/migration.go` has no 1.2.0 to 1.3.0 route and returns
   `no deterministic migration path` for that pair.
2. `internal/releasediscovery/evaluate.go` requires both the compiled local
   route and matching authenticated remote migration ID; index metadata alone
   cannot authorize a new route. Newer 1.3.0 therefore remains compatibility
   unsupported in the existing installed binary.
3. `internal/updateawareness/state.go` maps this to
   `update_available_unsupported_source`, not `update_available`.
4. `internal/updatenotification/service.go` Eligible requires exactly
   `UpdateAvailable`; the existing installed binary cannot deliver the requested
   pre-install notification for 1.3.0 even with a valid real signed index.
5. `cmd/qwsg/update_commands.go` also refuses the operator installation when
   its compiled PlanMigration fails, before stopping Guardian or applying files.

Adding the route only to the new 1.3.0 package cannot change the already
installed 1.2.0 checker/notifier/updater. Synthetic fixtures or weakening index
validation would not satisfy real pre-install acceptance. No production probe
with a fake release was attempted, and no missing acceptance is claimed PASS.

A bounded prerequisite option is an explicitly recorded compatibility backport
of the 1.2.0 to 1.3.0 route into the installed Task 080-based 1.2.0 binary, with
matching RELEASE provenance, focused validation and exact rollback capture.
The official immutable 1.2.0 archive would remain unchanged. This changes the
required existing-production pre-install baseline and needs the Owner decision
per Task 081 section 20 before that additional production deployment is made.
Release preparation must not silently substitute this changed baseline.

## Rollback

Verify original snapshot SHA256SUMS and bundle; extract only into a new isolated
directory. Restore only reviewed exact Task 081 paths; preserve unrelated work.
Never extract over the live tree or rewrite published history. No production
rollback is needed because production has not been changed. Before any future
production mutation capture its separate service/state/artifact rollback scope.

## Completion state

`incomplete — compatibility backport validation and production activation pending`

No 1.3.0 artifact, commit, tag, Release, signing input or publication is claimed.
Full release validation and production acceptance remain pending. Task 081
remains active; Task 082 has not been started.

## Authorized compatibility backport execution

The Owner explicitly approved the bounded prerequisite in the next session
message. The former decision boundary is resolved. Only the explicit
`compat-1.2.0-to-1.3.0` table entry changes runtime code; VERSION remains 1.2.0.
Regression coverage exercises actual installed-package classification, local
migration policy, deterministic test-key signature verification, awareness
classification and notification eligibility for supported/current/unsupported,
unsigned, untrusted-key and mismatched-provenance inputs. Filesystem comparison
proves no artifact acquisition or package mutation from this check path.

Production before-images are retained privately at
`/tmp/qwsg-task081-backport.jzpwq1c7`, with hashes:
- Binary: `dc86ba9eb19b83fccb3d1e233a3b21e482ec70d1041d53bc90fdb11d4ed96a93`
- RELEASE.json: `19f2c49997d7b183ca080411ac9dfe7c16d2f238e5c86c0e40944b901a5fd82b`

Latest pre-mutation Guardian: active/running, success, zero restarts,
MemoryCurrent=32002048, MemoryPeak=72060928, MemoryMax=134217728,
TasksCurrent=9, TasksMax=32. Source snapshot remains the verified pre-task bundle.
Initial restricted validation hit the existing HTTPS fixture socket denial:
ENVIRONMENTAL ISSUE; unrestricted identical validation is required. New
compatibility integration tests and focused update/installation/awareness/
notification tests passed in the restricted run.

Backport local gates PASS: `make fmt-check vet test engineering-test` with
unrestricted local HTTPS fixtures, complete `go test -race ./...`, build/export
provenance contract, framework (25), diversion (36), lifecycle (29) and Builder
(49) assertions. No test failure remained. The only runtime diff is the single
migration-table entry. Production activation requires interactive sudo
(`sudo -n true` returned password required); exact bounded deployment is being
prepared from the committed source. No activation has occurred.

Compatibility implementation committed and pushed after dry-run as
`27f25ed11d9cb211571fabfd2abffddc0a806f38`; synchronized main divergence 0/0.
Two local backport builds from this source reproduced identical bytes:
- Version `1.2.0`, built `2026-09-06T11:13:16Z`.
- Binary SHA-256 `161b952d83130b157a89560ce03caf1dc3165c276525b6029b85a650e7a914f5`.
- Matching RELEASE.json SHA-256 `d76f9edb3c87cc8b4007bef0f511a2888954ad05217a7a2d536ba14a7db021ca`.
- Independent fresh official 1.2.0 artifact retrieval still matches its protected hash.

Owner activation script: `/tmp/qwsg-task081-backport.jzpwq1c7/activate.sh`,
SHA-256 `a7356da74c9ac59715d8a1cc44828d3dfbc49aa5edfb2d27a9466e71d2b2ce79`.
It checks exact old/new hashes and service health, captures root-private copies
at `/var/lib/qwsg-task081-compatibility-backport`, stops only Guardian, replaces
only binary and RELEASE.json with preserved ownership/modes, then restarts and
checks health. Errors after the stop invoke the exact hash-guarded rollback.
Explicit rollback, only if acceptance requires it:
`sudo /var/lib/qwsg-task081-compatibility-backport/rollback.sh`.
The rollback refuses identities outside the old/backported pair, so cannot
silently overwrite a future canonical 1.3.0 installation. Both activation and
extracted rollback scripts passed bash syntax checks. Configuration, credentials,
index, unit file and user state are outside the mutation targets. Backups are
retained through acceptance and rollback-window closure. Owner execution is
pending; no production activation PASS is claimed.
