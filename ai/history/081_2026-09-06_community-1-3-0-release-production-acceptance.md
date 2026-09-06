# Task History 081: QWSG Community 1.3.0 Release & Production Acceptance

## Task metadata

- Task ID: `081`
- Task slug: `community-1-3-0-release-production-acceptance`
- Status: `complete — corrective 1.3.1 released and production accepted`
- Date generated: `2026-09-06` UTC
- Human authority: Project Owner explicit Task 081 authorization and APPROVE in the current session
- Preferred owner communication language: Hungarian
- Related prompt: `ai/archive_prompts/081_2026-09-06_community-1-3-0-release-production-acceptance.md`

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

## Backport production activation accepted

The Owner ran the reviewed script and reported activation PASS. A separately
pasted shell `&&` produced a syntax error but both checksum verification and
the subsequent sudo script completed successfully. Independent installed
inspection confirms exact new binary/RELEASE hashes and source
`27f25ed11d9cb211571fabfd2abffddc0a806f38`. Guardian active/running,
Result=success, NRestarts=0, MemoryCurrent=23941120, MemoryPeak=48963584,
TasksCurrent=8, unchanged MemoryMax/TasksMax. Real official index check at
`2026-09-06T11:29:54Z` succeeded current/equal installed/available 1.2.0.
Activation PASS entails the script's exact backup hash checks and root-private
rollback creation. Private pre-activation copies are also retained. This is the
explicitly compatibility-remediated installed baseline, not the historical
published 1.2.0 package. Task 081 release preparation resumed automatically.

## 1.3.0 preparation

VERSION, permitted deterministic release identity, current installation examples,
README, changelog and EN/HU release notes aligned to 1.3.0. Added real-archive
acceptance covering canonical old artifact hash, clean install, actual package
verification/classification, migration transaction, exact rollback and preserved
configuration/credential/awareness/inventory/oversized legacy state fixtures.
Added a 50-cycle sparse oversized Scheduler regression: rejects before decode,
retains the source, executes no pipeline work and releases the lock. This proves
bounded fail-closed handling, not healthy scheduling of an oversized state;
operator documentation requires resolution of unsupported legacy state before
claiming scheduling health. Production Scheduler was already bounded by Task 080.

Pre-publication validation: format, vet, all Go packages and full race PASS.
Engineering/build-export/framework/diversion/lifecycle/Builder PASS; additional
Framework v2 (15) and bounded diagnostic runner (11) assertions PASS. The old
build-contract version assertion was a TEST OR ACCEPTANCE DEFECT and was
updated to the authorized 1.3.0 identity; its full rerun passed. Release plumbing,
umask/source-mode reproducibility and two-build release-authority/tool/input
reproducibility PASS. Final documentation-only changes receive a further release
reproducibility check before freezing.

Real preliminary pipeline archive clean install, canonical official 1.2.0
migration, exact old-package rollback, identity convergence and private-state
preservation PASS through `TestRealRelease130CleanMigrationRollback` with both
explicit archive inputs. The ordinary suite intentionally skips this test when
real artifacts are not supplied; that skip is not acceptance evidence. Final
frozen artifacts will be checked again with explicit inputs.

Isolated real Guardian acceptance used separate private configuration/state,
2-second Guardian interval, GOMEMLIMIT=64MiB, MemoryMax=128M, TasksMax=32 and a
180-second runtime ceiling. Observed active/running Result=success, zero
restarts, MemoryPeak=43503616 bytes; final Scheduler retained exactly 64 successful
results in 1636952 bytes. The deliberate RuntimeMaxSec expiry later reports
Result=timeout with zero restarts; it is the bounded test termination, not a
product crash. No production config/state/service mutation occurred in this test.
The 50-cycle oversized sparse-state regression also passes without deleting the
legacy file or invoking the pipeline. Privacy/security review confirms only
version/packaging/documentation plus test changes after the one-line route
backport; authentication, credentials, listeners and automatic installation
boundaries remain unchanged.

## Frozen 1.3.0 publication and signing checkpoint

All local release gates passed, including repeated final documentation
reproducibility. Final release source/implementation commit:
`2e2b723f6f9a2bcea368fbbddc8449cdbaac7fab`.
Annotated tag `v1.3.0`, tag object `75d00762f4db1cb08ead68eccf388549540d6ee7`.
Commit and tag were pushed to canonical Forgejo after a successful dry-run.
Build epoch `1788696490`, built `2026-09-06T12:08:10Z`.
Two builds from an isolated Git archive of this exact source matched:
- Artifact `qwsg-1.3.0-linux-amd64.tar.gz`.
- Size `3598177` bytes.
- SHA-256 `8e19f624ccd1a49f32f6127889390e3389aec6eda958c7fbaaaa657d7b7cf9a0`.
- Sidecar size `96` bytes.

The real-package clean-install/migration/rollback test passed again with the
final frozen archive. Forgejo Release ID `4` is final/non-draft/non-prerelease,
URL `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/tag/v1.3.0`,
published `2026-09-06T12:11:54Z`, with the exact archive and sidecar attached.
Anonymous curl retrieval matches exact size/hash and original sidecar bytes.
No official 1.2.0 object was modified. No GitHub Release was created.

After the usage-limit interruption, actual local and fetched Git state was
verified: clean, synchronized main at the release source, divergence 0/0,
active Task 081. Publication records and external artifact retrieval agree.
Canonical signing input generated twice identically:
`release/production/qwsg-release-index-1.3.0-signing-input.json`, 733 bytes,
SHA-256 `9e6a61ec6727ae3596283d8c204e1ed8b845f633f0e4a6dc09fe55f63959f0b9`.
Generated-at `2026-09-06T16:38:50Z`; final byte `}` without newline.
It identifies stable active 1.3.0, minimum source 1.2.0 and exact route
`compat-1.2.0-to-1.3.0`, canonical Forgejo source and immutable artifact.

Dell1 uses the already provisioned reviewed offline signer:
`qwsg-release-sign-offline.exe`, 5118464 bytes, SHA-256
`c3f7e9459a8fa23cf6f87daf46046d0cd9bd67c7682efd2a450bf2bf1f7c8b0d`.
Local retained tool identity was verified; no private material was accessed.
Awaiting only the Owner's detached 89-byte Base64 signature output. The
production index remains signed 1.2.0; installed production remains the
compatibility-remediated 1.2.0 baseline. Real newer discovery/notification,
production 1.3.0 update and post-install acceptance are still pending.

Independent anonymous wget retrieval also matched the final artifact SHA-256.
Signing-checkpoint documentation updated; source tag and frozen archive remain
unchanged by these subsequent audit/signing-input commits.

## Owner offline signature verified; index publication prepared

Owner returned signing PASS and private-material-exposure NONE, detached
signature `dHglwHbrorZpjHGFMJrhBEGZpBN565mCAQYLkADvjf0UUomPMEmcvcxTWZ8ldt6+d0AB4N8CckJlkJdZmSG6CQ==`.
The LF-terminated 89-byte signature file matches Owner SHA-256
`9e797e6b52cfd167397ed1864c24e7a612cfa13690501db59561d525bdee5b40`.
Assembly and production verification PASS against the frozen 733-byte input
and unchanged bundled production anchor. Signed index: 913 bytes, SHA-256
`71af8a7c3c31b342e6fa20cce1b70e50879cd25fa6649c55713a2d50cdb371f9`.
Verification checkpoint: 433 bytes, SHA-256
`3ac9c55d4ce24913211a1c5e79d8c8f087d4fada370224d7a68c688343403f12`.

Current production filesystem object and fresh public HTTPS retrieval still
match the prior signed 918-byte 1.2.0 index. Prepared exact root-only script
`/tmp/qwsg-task081-index-publication/publish.py`, SHA-256
`c073b3e4725edc7c42754884eafbb32e4d5321db24a61505f35316fdae2f799b`.
Python syntax compilation PASS. It checks regular-file identities and exact
old/new hashes, acquires a bounded publication lock, saves private old/new bytes
under `/var/lib/qwsg-task081-release-index`, stages on the destination filesystem,
preserves existing owner/group/mode, rechecks old identity, atomically replaces
only release-index.json, fsyncs and verifies the installed object. No web config,
reload, TLS, DNS, release artifact or installed QWSG mutation is included.

Preserve both signed objects through the rollback window. Index rollback must
not weaken client generated-at protections: after clients observe 1.3.0, normal
recovery uses a newly generated/offline-signed later index. Restoring historical
bytes would intentionally fail closed for those clients and is not represented
as transparent recovery. Root execution remains Owner-required and pending.

## Published index and real pre-install acceptance PASS

Owner executed the bounded root publication; reported publication PASS with
index SHA-256 `71af8a7c3c31b342e6fa20cce1b70e50879cd25fa6649c55713a2d50cdb371f9`.
Independent public HTTPS retrieval matches all 913 signed bytes; production
verifier PASS. Required media type, no-cache, absent Expires, ETag,
Last-Modified, both conditional 304 forms and HTTP-to-HTTPS redirect PASS.
No hosting configuration was changed.

The real pre-install acceptance driver was compiled from the exact installed
backport source `27f25ed11d9cb211571fabfd2abffddc0a806f38` in a private isolated
checkout. It invokes the unchanged Guardian ReleaseCheckService, production
awareness manager, actual installed-package classifier and existing configured
SMTP provider against the actual published production index and actual
production awareness store. Only the driver interval is one minute, avoiding
waiting for the resident Guardian's normal 24-hour due interval. Production
configuration and resident interval remain unchanged. This is an explicit
bounded production acceptance invocation of Guardian's real core, not evidence
that the resident timer naturally became due or a fixture-index substitute.

At `2026-09-06T19:05:01Z`, installed compatibility-remediated 1.2.0 discovered
stable 1.3.0 as supported/newer/update_available with production key
`qwsg-community-release-2026-01`. The configured notify policy and SMTP
preflight passed. The real provider accepted the first update notification
(one delivery call) and persistent LastNotification was saved. At
`2026-09-06T19:06:01Z`, a new notifier instance plus repeated real authenticated
check produced zero delivery calls and retained deduplication. SMTP acceptance
is proven, not end-recipient mailbox receipt. Both bounded runs PASS.
No artifact or transaction appeared and installed version remained 1.2.0.
Configuration bytes were unchanged. `strace -f -e trace=network` confirms local
`qwsg update status` performs no network syscalls and reports update_available.

Resident Guardian after delivery: active/running, success, NRestarts=0,
MemoryCurrent=50155520, MemoryPeak=69976064, TasksCurrent=9,
MemoryMax=134217728, TasksMax=32; cgroup oom/oom_kill remain zero.

Pre-update rollback checkpoint `/tmp/qwsg-task081-production-update` contains
exact private binary/RELEASE/unit/config before-images and hashes, plus a live
runtime-state observation explicitly not represented as a coherent automatic
restore payload. Existing deterministic native package rollback is the supported
recovery contract. Backport binary/RELEASE still match frozen hashes; 1.3.0
archive reverified at its exact canonical hash. Retain all snapshots through
acceptance and rollback-window closure. The next Owner command is the ordinary
user's canonical `qwsg update --archive ... --version 1.3.0`; its existing narrow
helper requests interactive sudo, verifies package/migration again, preserves
service intent and records rollback. No further conceptual approval is needed.

## Production migration PASS; post-install acceptance BLOCKED

Owner ran the canonical explicit updater and reported:
`QWSG updated safely: 1.2.0 -> 1.3.0`, native rollback available, lifecycle
administrator notification ACCEPTED. Independent comparison proves installed
binary, RELEASE.json and Guardian unit are byte-identical to the frozen official
1.3.0 archive. Version/source/build all match release commit
`2e2b723f6f9a2bcea368fbbddc8449cdbaac7fab`. Configuration and credential-file
hashes match private before-images. Guardian active/running, Result=success,
NRestarts=0, initial MemoryCurrent=27222016, MemoryPeak=33779712,
TasksCurrent=8, unchanged MemoryMax/TasksMax; active since
`2026-09-06T19:09:17Z`. Active Scheduler remains bounded at 64 results,
approximately 1.64 MB; the separately retained legacy backup is not active state.
Native update created current.json rollback metadata.

At `2026-09-06T19:10:11Z`, the post-install real official-index check exposed
QWSG-081-F001 (PRODUCT/FRAMEWORK DEFECT, release acceptance blocker): installed
and available are both 1.3.0, but status remains update_available and relation
newer. Manager.Check reuses HTTP validators across installed version changes;
NewSuccess handles 304 by cloning the prior evaluation and replacing only its
installed identity, without recomputing relation/compatibility. Therefore the
required current/equal acceptance fails despite authentic metadata.

QWSG-081-F002 (PRODUCT/FRAMEWORK DEFECT, release acceptance blocker): NewSuccess
preserves LastNotification only when installed identity is unchanged. The
pre-update successful persistent notification record was present in the private
snapshot but absent after this check. Combined with stale newer classification,
a subsequent notifier evaluation could duplicate an already delivered release
notification. No second update-availability email was intentionally triggered.

Both defects reproduce deterministically against the exact frozen release
source in `/tmp/qwsg-task081-postinstall-finding/source`. The initial added
notification regression had a temporary-directory mode mismatch (TEST OR
ACCEPTANCE DEFECT); fixing the fixture to mode 0700 made both intended product
regressions fail on original source. A minimal isolated correction rejects 304
across identity changes, omits conditional validators after identity change,
and preserves release-notification identity across the version transition.
Both new regressions and existing awareness/notification suites pass with that
local correction. It has not changed repository runtime code, the installed
binary or any public artifact. Full successor validation remains necessary.

The official v1.3.0 tag/archive/signature chain is already public and immutable.
Replacing its bytes, retagging, or installing a private 1.3.0 correction would
violate the requested final canonical provenance chain. Task 081 section 20
therefore requires an Owner decision on a distinct corrective release identity
(recommended 1.3.1) and revised completion scope. No Task 082 is started.
The healthy monitoring service is left running; no unnecessary destructive
production rollback or awareness-state deletion has been performed. Production
acceptance is incomplete, not PASS. Pre-update deduplication evidence and
post-update defective awareness bytes are preserved privately for bounded
recovery with the authorized correction.

## Owner-approved corrective 1.3.1 scope — IN PROGRESS

The Owner explicitly approved the corrective 1.3.1 release within Task 081;
only F001/F002 and directly required validation, docs, metadata, migration and
production acceptance are authorized. Immutable 1.3.0 tag/artifact/Release and
production history remain unchanged. Actual installed 1.3.0 is the upgrade
source; no Task 082 is opened. The former release-identity decision is resolved.

Corrective baseline: fetched synchronized clean main
`524b107fb237f3adb477ec26980f60a2006719b2`, valid active Task 081.
Snapshot `/tmp/qwsg-task081-131-baseline.lrtkzh9p`: protected complete Git bundle,
tracked-source archive, hashes and bounded RESTORE.md; bundle, checksums and
archive readability PASS before changes. Retain through acceptance/rollback
window. Production state/configuration snapshots from the 1.3.0 transition and
post-install findings remain preserved privately.

Correction omits HTTP validators for a changed installed identity or legacy
inconsistent cached relation, refuses unsafe 304 reuse, and obtains a new fully
authenticated evaluation. Failed refreshes retain historical authenticated
observation and rollback watermark without claiming a valid current relation:
status is unknown until fresh evaluation succeeds. Successful notification
identity survives success/failure version transitions. Eligibility independently
checks actual version ordering so a legacy stale equal record cannot notify.
No network fallback, unsigned metadata, classification/provenance bypass or
automatic acquisition/installation is added.

New deterministic regressions cover changed identity, 1.3.0 legacy stale state,
current/equal 304 reuse, unsolicited unsafe 304 refusal, failed-refresh restart,
metadata rollback after identity change, dedup persistence through upgrade and
rollback, and suppression of legacy equal-state notification. Existing focused
awareness/notification/update tests plus new regressions PASS.
The explicit sole new route is `compat-1.3.0-to-1.3.1`. The verified 1.3.1 archive
binary will orchestrate the existing native transaction because immutable
installed 1.3.0 lacks this route. Real-package acceptance now covers both
historical 1.2.0 to 1.3.0 and corrective 1.3.0 to 1.3.1 with exact predecessor
hashes, preservation, clean-install convergence and rollback. VERSION and
current EN/HU packaging/docs aligned to 1.3.1; historical 1.3.0 notes unchanged.

Corrective continuation after usage interruption verified staged scope, unchanged
1.3.0 tag object and artifact hash, and existing preliminary 1.3.1 archive. The
prior background sessions no longer exposed completion results; only those
unproven gates were repeated with durable logs. Full `make fmt-check vet test
engineering-test release-check release-authority-check` and `go test -race ./...`
now completed exit 0. Logs retained under `/tmp/qwsg-task081-131-candidate`.
Additional Framework v2 and bounded diagnostic tests had already passed.
Real preliminary archive `TestRealRelease131CleanMigrationRollback` PASS.

Using the preliminary 1.3.1 binary on an isolated copy of actual production
1.3.0 stale awareness, real HTTPS checks at `2026-09-06T22:48:00Z` and
`2026-09-06T22:48:58Z` both report current installed/available 1.3.0. The first
refresh corrects the legacy stale relation; the subsequent conditional check
retains the correct equal relation. The public index remained the unchanged
signed 1.3.0 index throughout; production state was not modified by this test.
Scope/privacy review and staged whitespace checks PASS. No historical 1.3.0
release notes, production signing objects, tags or artifact bytes are modified.

## Canonical 1.3.1 release published; offline signing pending

Final corrective source commit `45009fe6169bff00842a4c4e9561bf339a5db81e`,
build epoch `1788735033`. Two deterministic builds from its isolated Git source
archive matched byte-for-byte:
`qwsg-1.3.1-linux-amd64.tar.gz`, `3600553` bytes, SHA-256
`0ab726bcde36182232ff89e3ea33e2d5ab77d8cad2b94ce56c9534a8147d9cd7`.
Final real-archive clean/migration/rollback acceptance PASS (explicit 1.3.0
and 1.3.1 artifact inputs), with the immutable predecessor hash verified.

Final binary's isolated Guardian ran under MemoryMax=128M, TasksMax=32,
GOMEMLIMIT=64MiB and a 180-second runtime ceiling. Retained 63 successful
executions plus one pipeline cancellation at the deliberate runtime-ceiling
shutdown, 64 total results in 1612282 bytes, NRestarts=0. Unit Result=timeout
and journal prove the intentional deadline shutdown, not a restart/OOM loop.
The completed transient unit no longer exposed MemoryPeak, so no measured peak
is claimed for this run; actual post-install production resource measurement
remains required. Historical 1.3.0 measured resource evidence remains intact.

A bounded recovery driver compiled from the final corrective source restored
only the lost successful 1.3.0 notification record from the integrity-verified
pre-upgrade awareness before-image, verifying its exact release/artifact/key
identity. It then used the real production discoverer to repair current/equal
installed/available 1.3.0 while preserving that delivery record. Both checks
PASS; no SMTP delivery, installed binary or configuration mutation occurred.
Pre-recovery awareness SHA-256
`52702c856b82deabc3a20b3eced5aee51d6fff0e645a0491525e54ed824e2a1f`
and exact before-image retained under `/tmp/qwsg-task081-131-driver`.
This evidence-backed recovery does not fabricate delivery history.

Annotated `v1.3.1` tag object `efb7afcac7bd65a3419f6bf3264892e09de631f6`
points to the exact source commit. Commit and tag pushed after dry-run.
Forgejo Release ID `5`, final/non-draft/non-prerelease, published
`2026-09-06T22:56:55Z`, URL
`https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/tag/v1.3.1`.
Exact artifact and sidecar attached. Independent anonymous curl and wget
retrievals match final size/hash; downloaded sidecar matches original bytes.
No 1.3.0 tag, artifact, Release metadata or signing/history object was mutated.

Canonical 1.3.1 signing input generated twice identically:
`release/production/qwsg-release-index-1.3.1-signing-input.json`, 733 bytes,
SHA-256 `7b1fac2eb0c05e98e4f6f3ddc9fe5aad91581722070d70cb763555e286ee1959`.
Generated-at `2026-09-06T22:58:30Z`; no trailing newline. It advertises exact
Forgejo stable active 1.3.1, minimum source 1.3.0 and migration route
`compat-1.3.0-to-1.3.1`. Offline signer remains exactly 5118464 bytes, SHA-256
`c3f7e9459a8fa23cf6f87daf46046d0cd9bd67c7682efd2a450bf2bf1f7c8b0d`.
Only the detached Owner signature is pending at this boundary. Production index
and installed version remain 1.3.0. Task 081 remains IN PROGRESS; full signed
1.3.1 publication and production upgrade acceptance are not yet claimed.

## Owner offline 1.3.1 signature verified; root publication prepared

Owner returned the detached public signature; no private material was handled.
Signature file: 89 bytes, SHA-256
`c2585b68312a63dd73e22c06353d44829b9e64acd5b2609fbdd6b13156ce2160`.
Canonical authority assembly and Ed25519 verification PASS against the frozen
signing input and accepted production trust anchor/key ID.
Signed index: 913 bytes, SHA-256
`2eefd328c5ded8807101bdd69d5fcee0ed4f4f2e49b10cdd02b55da6e7aaf879`.

Bounded root publication script:
`/tmp/qwsg-task081-131-index-publication/publish.py`, 2538 bytes, SHA-256
`d7857aa378dfe89579bae3271d1b66f2e2a4be71bcb9f60c6e57961522cefc37`.
Syntax validation PASS. Requires exact previous signed 1.3.0 index hash
`71af8a7c3c31b342e6fa20cce1b70e50879cd25fa6649c55713a2d50cdb371f9`.
Preserves ownership/mode, uses locked same-directory atomic replacement and
fsync, retains exact before/after objects and recovery instructions under
`/var/lib/qwsg-task081-131-release-index`. Restoring an older index after newer
client observation remains fail-closed; normal recovery requires a later signed
index. Historical releases and web-server configuration remain unchanged.
Publication is pending Owner interactive root execution. Task remains IN PROGRESS.

## Signed 1.3.1 publication and real pre-install acceptance PASS

Owner executed root publication successfully. Independent public HTTPS retrieval
matches signed index SHA-256
`2eefd328c5ded8807101bdd69d5fcee0ed4f4f2e49b10cdd02b55da6e7aaf879`;
canonical authority verification PASS. Required media type, no-cache,
ETag/Last-Modified, empty conditional 304 without Expires and HTTP-to-HTTPS
redirect PASS. No server configuration changes.

Final-source 1.3.1 bounded acceptance driver used the actual installed 1.3.0,
real production index, configuration, credential provider and persistent store.
First supported newer 1.3.1 evaluation: SMTP calls=1, provider accepted.
Second evaluation after one minute with a fresh notifier: SMTP calls=0.
Persistent successful notification identity and unchanged configuration verified.
No automatic artifact staging/download/installation; installed version remained
1.3.0. This is the final corrective Guardian core running as a bounded driver,
not a claim that immutable resident 1.3.0 supports the new compiled migration
route or that its ordinary 24-hour timer fired. Existing updater is invoked from
the verified canonical 1.3.1 archive for the operator-controlled migration.

Production Guardian: active/running, Result=success, NRestarts=0,
MemoryCurrent=46366720, MemoryPeak=70963200, TasksCurrent=9;
cgroup max/oom/oom_kill/oom_group_kill all zero.
Pre-update private configuration, binary, RELEASE, unit and hashes preserved at
`/tmp/qwsg-task081-131-production-update`; awareness retained as an observation.
Installed 1.3.0 and candidate 1.3.1 binary/RELEASE bytes match their respective
verified immutable archives exactly. Native deterministic rollback already
passed the final real-package acceptance; production rollback is reserved for
actual failure via `qwsg update rollback` and native transaction backup.
Operator update requiring interactive credentials is the next boundary.

## Final corrected production acceptance and closure

Owner native update: 1.3.0 → 1.3.1 PASS; rollback available and lifecycle admin
notification ACCEPTED. Installed binary, RELEASE.json and service unit match the
canonical 1.3.1 archive exactly. Source commit
`45009fe6169bff00842a4c4e9561bf339a5db81e`, built `2026-09-06T22:50:33Z`.
All previously captured private configuration/credential file hashes unchanged.
Native update transaction identifies Installed=1.3.1, Previous=1.3.0 with backup
reference; deterministic final real-package clean/migrate/exact-rollback test
already PASS. No unnecessary rollback of healthy production was executed.

F001 PASS: first post-install authenticated check reports current/equal 1.3.1.
Two additional real HTTP 304 responses observed through the final checker each
retain current/equal. F002 PASS: exact successful pre-install 1.3.1 delivery
record survives native update/service restart, checks and fresh store/notifier
instances; both post-install evaluations make zero delivery calls. This is
separate from the earlier evidence-backed restoration of the historical lost
1.3.0 record. Final installed identity supersedes both older production baselines.
`update status` syscall trace contains only process signals, zero network calls.

Production observed from service activation `2026-09-06T23:35:43Z` through
`2026-09-06T23:48:56Z` (over 13 minutes). Same main process, active/running,
Result=success, NRestarts=0. Final MemoryCurrent=32190464 bytes,
MemoryPeak=60149760 bytes, safely below MemoryMax=134217728;
TasksCurrent=8 below TasksMax=32; GOMEMLIMIT=64MiB confirmed.
Cgroup low/high/max/oom/oom_kill/oom_group_kill all zero.
Three post-install Scheduler results succeeded, retained total=64,
active state=1636882 bytes below 8MiB. Historical oversized backup remains a
separate inactive before-image. No panic/fatal/OOM journal events and no QWSG
inbound listeners. No unrelated production configuration changes, telemetry,
registration, automatic artifact download/install or credential disclosure.

Public release/index, production signature/trust, source/tag, exact artifact
and installed identity form one verified chain. Historical official 1.2.0 and
1.3.0 artifact hashes reconfirmed unchanged. Forgejo 1.3.0 and 1.3.1 release
metadata/assets match publication evidence; only ordinary download counters
changed. EN/HU acceptance documentation added, immutable release notes untouched.
All previously completed full/race/engineering/reproducibility/security gates
remain applicable; no product code changed after the final validated release.
Only documentation/lifecycle closure checks are required at this boundary.

Task 081 completes with corrective 1.3.1 RELEASED / ACCEPTED. Initial 1.3.0
remains an immutable published release whose post-install regressions motivated
the authorized correction; it is not retrospectively relabeled fully accepted.
Private snapshots/native rollback evidence retained through the rollback window.
Archive Task 081 without creating Task 082. The closure commit is the Git commit
containing this final record and prompt archival; final push/sync is verified
in the Owner handoff rather than embedding a self-referential commit hash.
