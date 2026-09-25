# Task History 093: Community C9 Release Candidate Preparation

## Task metadata

- Task ID: `093`
- Task slug: `community-c9-release-candidate-preparation`
- Status: `active`
- Date generated: `2026-09-25` UTC
- Human authority: Project Owner explicitly authorizes Task 093 in this session: release candidate preparation, lifecycle, source commit and push; no signing, publication or VPS mutation.
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/093_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit session authority was recorded; source preparation and mandatory local validation completed. Candidate construction and offline handoff remain in progress.

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

`active`

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
