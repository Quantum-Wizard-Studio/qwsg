# Task History 093: Community C9 Release Candidate Preparation

## Task metadata

- Task ID: `093`
- Task slug: `community-c9-release-candidate-preparation`
- Status: `complete`
- Date generated: `2026-09-25` UTC
- Human authority: Project Owner explicitly authorizes Task 093 in this session: release candidate preparation, lifecycle, source commit and push; no signing, publication or VPS mutation.
- Preferred owner communication language: Hungarian
- Related prompt: `ai/archive_prompts/093_2026-09-25_community-c9-release-candidate-preparation.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit session authority was recorded; source preparation and mandatory local validation completed. Candidate construction, exact offline handoff and local validation completed; closure archives only 093 without a successor.

## Starting state

See the verified baseline, snapshot and bounded restoration record below.

## Snapshot

See the verified baseline, snapshot and bounded restoration record below.

## Work performed

Prepared 1.4.0 identity, matching release notes/changelog, minimal build/test allowlists and a bounded actual-candidate package bootstrap test. No runtime code changed.

## Verification

Builder input, metadata, prompt/history identity, approval state, and lifecycle installation validated successfully.

## Rollback

See the verified baseline, snapshot and bounded restoration record below.

## Completion state

PASS: release candidate preparation complete. C9 remains MISSING; no signing, publication or real-system acceptance. Final integration follows reviewed source/closure commits, dry-run push, fast-forward push and clean synchronized idle verification.

## Verified baseline and release decision

Fetched main equals origin/main at 5b2d4acf36cf844988851e4b551bc7a9a3a7aa1c;
0/0 divergence and clean worktree. Framework 2.0 valid, idle after complete 092.
Canonical Builder installed 093 from explicit session authority. Required core,
backup/release/security/delivery/documentation policies, release machinery,
historical signing runbook and Tasks 082–092 inspected. Hungarian configured.
Snapshot logical identifier qwsg-task093-snapshot, protected configured local
storage; tracked baseline archive SHA256
94fc5e404228e753bc9fba184b61892f1b9171da2e9e4fd71271b16765805a82.
Archive readable; per-file SHA256 manifest retained outside Git. Retain through
Owner acceptance and signing/rollback dependencies. Bounded restoration follows
the prompt and snapshot RESTORE.md; no destructive restore performed.

1.4.0 is the smallest minor identity reflecting additive post-1.3.1 functionality;
see RELEASE_NOTES_1.4.0.md for authority and stable-index constraints. It remains
an unpublished candidate. No runtime mechanism, historic migration route or
production trust anchor is changed. Official 1.3.1 needs the independently
verified candidate binary as bootstrap coordinator; it cannot parse /2 itself.
Source boundary will be committed after mandatory validation; artifact/signing
identities will be recorded in a separate closure commit to avoid circular hashes.

## Source-boundary validation

- Focused updateauthority/releasepublication/releasediscovery/update/installation: PASS.
- Full make fmt-check vet test: PASS; full go test -race ./...: PASS.
- make engineering-test: PASS, checkout/export provenance identity plus 25 framework,
  36 diversion, 29 lifecycle and 49 Builder assertions.
- make release-check: PASS; frozen-client forward authority gate, identity/plumbing,
  archive equality across umask/source modes.
- make release-authority-check: PASS; two isolated Linux verifier/Windows signer
  builds and both /1 and /2 canonical signing inputs byte-identical.
- Framework v2 15 assertions, diagnostic runner 11 assertions, diverted-task audit:
  PASS. No tests weakened or skipped to obtain success. Optional real-package
  test awaits the committed candidate artifact and actual metadata.
- Historical production 1.3.1 index verifies under unchanged production key;
  official archive hash and v1.3.1 source commit match protected Owner values.
- Snapshot archive and all 617 regular-file member hashes reverified PASS.
- Exact nine-path source staging reviewed, modes unchanged, no private material,
  generated binary or historical production metadata staged. Whitespace PASS.

Diagnosed attempts: ENVIRONMENTAL ISSUE — sandbox blocked local TLS test socket
and .git writes; authorized elevated focused tests and staging passed. TEST OR
ACCEPTANCE DEFECT — existing build-contract check pinned 1.3.1; minimally updated
to 1.4.0, then remaining engineering/release gates passed. Full Go/race source
checks were not needlessly repeated for this shell identity assertion.

Source commit is a real candidate input checkpoint, not task completion. A separate
closure commit will retain concrete artifact/signing evidence without changing
package inputs. No tag, production signing, publication or host installation.

## Frozen candidate and exact handoff

Source commit: `996fb90d0f53d4ae6088cd884a654d9d0a32fff9`.
Source epoch: `1790357147`; embedded build: `2026-09-25T17:25:47Z`.
Archive: `qwsg-1.4.0-linux-amd64.tar.gz`, 3680796 bytes, SHA-256
`26cbf8d76fa27c652e11ee49f4570e05b3be6b969309e3cfd486977a27a6224d`.
Two independent source exports/caches and different umasks yield identical bytes.
Go 1.26.5, GNU tar 1.35 and gzip 1.12; existing deterministic build script.
Manifest, archive sidecar, RELEASE.json, exact embedded version/commit/date,
platform and production public trust asset verified. Payload retained in protected
local task storage outside Git; exact regeneration is documented.

`release/candidates/1.4.0/HANDOFF.md` is the exact signing/bootstrap handoff.
Its adjacent unsigned candidate, manifest/provenance copies, sidecar and checksum
list are trackable public metadata. Canonical signing input:
`qwsg-release-index-1.4.0-signing-input.json`, 973 bytes, SHA-256
`b44f83cb93f9068ab96a547e35e8544baaa45c6ef424281fd3cc63b95393554a`.
Generation repeated with the independently exported-source verifier matches.
Input binds target commit/version, artifact name/URL/size/hash and exact-source
compatibility. File is compact UTF-8, no final newline; signature is absent.
Payloads and signing inputs frozen read-only; no production checkpoint invented.

Unchanged production authority: `qwsg-community-release-2026-01`, raw public-key
SHA-256 `0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6`.
Expected later Owner return: `qwsg-release-index-1.4.0-signature.base64`, 89 bytes.
The verified existing Dell1 raw Ed25519 signer accepts canonical /2 bytes without
a runtime parser upgrade. Key/passphrase remain offline; no signature was created.
Prospective stable/active/timestamp/URL values are unsigned staged fields, not a
publication claim. No v1.4.0 tag or Forgejo Release was created.

## Bootstrap proof and limitations

Read the actual v1.3.1 source: its strict /1 parser and compiled routes cannot
implement the new path. Supported later bootstrap uses the independently
production-signature-verified candidate archive binary as coordinator, with
archive/checksum/full signed index. It probes the installed 1.3.1 identity and
its own privileged helper repeats authentication and package checks. No direct
old-client native update, install.sh overwrite, signature bypass or new route.
The exact declaration is 1.3.1 -> 1.4.0/linux-amd64/preserve-package-v1 with
Configuration/Guardian/Scheduler 1.0 and Operator State 1.0–1.2.

Actual unsigned candidate data is copied into a test-only signature domain for
`TestReleaseCandidateBootstrapPackage`. The original official 1.3.1 archive hash
and source identity are asserted before isolated staging. Candidate authorization,
manifest/provenance, binary version, exact package apply and rollback, unchanged
private fixture and all prior package bytes pass. Unknown capability, unsupported
schema, wrong source, altered signature and unsigned candidate refuse. Test key
is explicitly rejected by the production verifier. Full actual-candidate authority
package tests pass again under race after the fixture correction.

TEST OR ACCEPTANCE DEFECT: Go temporary staging parent inherited permissive
permissions; the real staging gate correctly refused. Set only that test directory
to 0700 and reran normal/race actual-package checks, both PASS. This test-only
change in the closure commit is excluded from every release package input; the
frozen source/artifact/signing bytes remain unchanged. No runtime fix or rebuild.

The existing frozen-client CLI gate also passes, but it is a post-082 capable
fixture, never falsely represented as official released 1.3.1. The new package
proof is isolated, not real sudo/systemd/VPS/Guardian acceptance. C9 must test the
explicitly supported bootstrap and recovery path later on an authorized real host.
Publishing /2 will make legacy /1-only awareness refuse until bootstrap; preserve
this fail-closed behavior. Historical service comparison constraints remain as
accepted in Task 089 and are documented in the handoff.

## Final acceptance and integration review

| Owner acceptance cases | Result |
| --- | --- |
| 1–4 baseline, version decision, historical immutability, distinct identity | PASS |
| 5–8 artifact, rebuild, manifest/checksum, metadata consistency | PASS |
| 9–10 authenticated bootstrap representation and unsupported-path refusal | PASS locally; real-system proof remains C9 |
| 11–12 deterministic signing input and bound integrity/migration fields | PASS |
| 13–16 no production signing/publication/VPS mutation; C9 stays MISSING | PASS |
| 17–18 lifecycle and deterministic repository integration | Closure and final synchronization verified before Owner delivery |

Production verifier independently rejects both unsigned forms without producing a
checkpoint. SHA256SUMS and both artifact hashes verified again after preparation.
Historical production paths, trust anchor, tags and 1.3.1 artifact remain unchanged.
No production key, secret, snapshot payload or binary is staged. Package input
comparison against the source commit and final scope/mode/whitespace review are
required before closure commit; no release runtime input may drift.

Documentation: matching release notes/changelog, frozen candidate handoff and
metadata, this independent history, archived prompt and one milestone. No product
scope or gate-status change. Community 8 PASS / 0 PARTIAL / 1 MISSING; Pro 9/2/3.
Deferred work is exactly Owner signing, signed-index assembly/verification,
separate publication and real C9 clean-install/bootstrap/preservation/rollback/
Guardian acceptance. No new unrelated finding, Pro work or Task 094.

Integration uses the source commit above plus a closure commit for immutable
candidate metadata and the corrected test fixture. Both use reviewed explicit
paths. Run canonical dry-run push, fast-forward push, fetch and verify HEAD equals
origin/main with divergence 0/0 and clean worktree. Final commit hash is reported
in the Owner delivery to avoid a self-reference in this record.
