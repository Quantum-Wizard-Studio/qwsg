# Task History 076: Approved Task

## Task metadata

- Task ID: `076`
- Task slug: `release-index-metadata-authenticity-release-source-contract`
- Status: `complete — validated and canonically closed for Owner review`
- Date generated: `2026-08-29` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/archive_prompts/076_2026-08-29_release-index-metadata-authenticity-release-source-contract.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded, and the Project Owner separately directed canonical execution on 2026-08-30 UTC.

## Starting state

- Repository: `/home/qws/web/qwsg.quantumwizard.hu/qwsg`
- Branch: `main`
- HEAD and `origin/main`: `687285a4d525557c1f9b1ffb2444c7918e2f7cb4`
- Divergence: `0/0`; index clean; only the approved Task 076 prompt/history
  pair was initially untracked.
- Remote: `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg.git`
- `VERSION`: `1.2.0`
- Annotated tag object: `ac395b568b8e1f83c0ef85c9aa02f98c15402af0`
- Tag peel/release commit: `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`
- Frozen artifact: `3524214` bytes; SHA-256
  `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`
- Framework and active Task 076 lifecycle validation: PASS.

## Snapshot

- Builder snapshot: `/tmp/qwsg-task076-builder-install.i2shcT`; reviewed Owner
  input, generated pair, complete tracked bundle, identity and rollback
  evidence; directory `0700`, files `0600`, checksums/bundle PASS.
- Execution snapshot: `/tmp/qwsg-task076-execution.WsUfQ2`; exact intended
  before-images/absence, Git/lifecycle/framework/protected-release identities,
  modes/ACLs, complete bundle, checksums and literal bounded rollback;
  directory `0700`, evidence `0600`, checksums/bundle PASS before production
  modification.
- Closure snapshot: `/tmp/qwsg-task076-closure.CM0O9V`; committed active-task
  prompt/history, Git/lifecycle/framework and protected-release identities,
  complete repository bundle, bounded rollback, modes and checksums; directory
  `0700`, evidence `0600`, checksums and bundle verification PASS before prompt
  archival.

## Work performed

- Inventoried existing Forgejo API discovery, bounded HTTP client, acquisition,
  sidecar/archive/package verification, strict versions, declared migrations,
  rollback, Task 075 installed classifier and Task 074 trust boundaries.
- Added the dedicated read-only `internal/releasediscovery` subsystem rather
  than changing or duplicating the updater.
- Implemented strict duplicate-aware bounded `qwsg.release-index/1` parsing,
  semantic validation and safe failure categories.
- Defined deterministic typed JSON signature bytes excluding only signatures;
  implemented Ed25519 verification against a copied explicit trust-anchor set
  and immutable authenticated metadata results.
- Implemented source-neutral `ReleaseSource`, bounded static HTTPS retrieval,
  strict authority/path/redirect/media/body/validator/timeout/cancellation
  handling and privacy-minimized fixed anonymous requests.
- Implemented the canonical fetch -> parse -> authenticate -> evaluate pipeline.
  Evaluation accepts only Task 075 verified installed identity; remote minimum
  source/routes can narrow advice, while Task 075/local migration evidence
  remains authoritative.
- Added deterministic parser, adversarial, signature, immutability, collection,
  local TLS transport, privacy, withdrawal, prerelease, installed identity and
  migration interaction tests.
- Added the canonical architecture contract and synchronized discovery, native
  update, security/privacy, operator, Architecture, System Map, Roadmap and
  Engineering History documents.

## Verification

- Focused `go test ./internal/releasediscovery ./internal/update
  ./internal/installation`: PASS, including strict schema, deterministic
  signature vector, tamper/unknown-key/duplicate/trailing-data rejection,
  bounded transport, privacy, installed identity and migration authority.
- The first focused run could not create local TLS test listeners in the
  restricted sandbox; classified `ENVIRONMENTAL ISSUE`. The unchanged suite
  passed with loopback-only test-listener authority.
- `make fmt-check vet test`: PASS for formatting, `go vet ./...`, and every Go
  package.
- `go test -race ./...`: PASS for every Go package.
- `make engineering-test`: PASS: Framework 25, diversion 36, lifecycle 29 and
  Builder 49 assertions, plus deterministic checkout/export build-contract
  checks.
- `scripts/test-release-reproducibility.sh`: PASS across source-mode and umask
  variation. `make release-check` stopped at its intentional clean-output
  precondition because the immutable published final artifact already exists
  in `dist/`; classified expected behavior. The artifact was not removed or
  overwritten merely to exercise release construction.
- Active `bin/job --check`, Framework validation, `git diff --check`, staged
  diff check, exact scope/mode/ACL review and bounded secret-pattern scan:
  PASS.
- Protected evidence remains exact: annotated tag object
  `ac395b568b8e1f83c0ef85c9aa02f98c15402af0`, peel/release commit
  `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`, artifact size `3524214`, and
  SHA-256
  `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`.
- No external host was mutated or contacted. No Guardian scheduling, Update
  Awareness persistence, notification/deduplication, automatic installation,
  Pro/telemetry, publication or Task 077 implementation was added.

## Rollback

Verify snapshot checksums/bundle, exact repository/protected-release identity,
current target hashes and absence of later Owner edits. Restore only named
before-images and remove only Task 076-created paths with recorded prior
absence/current identity. Preserve modes/ACLs. Never reset, checkout, clean,
rewrite refs, alter `v1.2.0`, remove the frozen artifact, or touch an external
system. Re-run focused/full/framework/lifecycle/Git/reproducibility/release
identity checks.

## Completion state

`complete — implementation validated and canonically closed for Owner review`
