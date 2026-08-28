# Task History 068: Approved Task

## Task metadata

- Task ID: `068`
- Task slug: `qwsg-change-notifications-rc4`
- Status: `complete — private RC.4 ready for final acceptance`
- Date generated: `2026-08-28` UTC
- Human authority: Project Owner: Attila
- Preferred owner communication language: English
- Related prompt: `ai/prompts/068_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval, Framework validation and lifecycle checks passed; execution started under the approved envelope.

## Starting state

- Verified canonical root, Framework 2.0.0, HTTPS origin, branch `main`, and fetched `HEAD == origin/main == 1f876d130d3ab4f204b8c1758411ef1f95889991` with `0 0` divergence.
- Before Builder installation the tracked/index worktree was clean and lifecycle idle after Task 067; afterward only the expected Task 068 prompt/history were untracked.
- Verified private RC.3 SHA-256 `8543c3e09b48085b01c037d7db5106ea793374dc099b0b9be5f6cacb55af13ee`, source commit `b6eb357ad03a02b41ac93536fc3be91ecf929803`, deterministic packaging record, and absence of `v1.2.0*` tags.

## Snapshot

- Protected payload root `/tmp/qwsg-task068-execution.RhmHLa`, mode `0700`, retained through Owner acceptance and release rollback-window closure.
- Complete tracked HEAD archive plus separate Task 068 prompt/history before-images are mode `0600`, deterministically manifested and SHA-256 covered.
- Checksum verification, archive readability, isolated mode-`0700` extraction, notification-tree presence and prompt/history comparisons passed. Restore remains exact-path, collision-aware and never extracts over the live worktree.

## Work performed

- Inspected the provider-neutral notification engine, Community SMTP adapter, protected credentials, canonical configuration/extensions, guided installer/catalogs, native updater/rollback, Guardian/user-service activation, security and existing tests.
- Selected the narrow integration: a lifecycle-event composition/dispatch layer reusing the existing SMTP `Config`, `Provider`, credential store and locale setting. No second transport, credential store, daemon or continuous file monitor was introduced.
- Added privacy-safe EN/HU/DE lifecycle event rendering, explicit operation versus delivery results, deterministic in-process event-ID duplicate suppression, secret-keyword redaction and transport-independent tests.
- Extended the existing email extension with `lifecycle_enabled` exposed as `notification.lifecycle.enabled`, and reused existing recipient/preflight/test/credential paths.
- Integrated configuration changes, post-configuration installation success/failure, successful Guardian activation, update success/failure with old/new versions, and rollback success/failure with installed/restored versions. Delivery status is printed independently and never replaces the operation exit result.
- Advanced the private identity to RC.4 only after full validation, documented RC.3 supersession, and froze source commit `4f7dcc11b5ccc9f078755946995baebd31ad6870`.
- Built `qwsg-1.2.0-rc.4-linux-amd64.tar.gz` with epoch `1787907901`; archive SHA-256 is `adeb591605c0d37a5fc98d541125ca388cd4561703d0f0823bba931bc7d08684` and sidecar-file SHA-256 is `d68523a5bdea974b23629e91b9b0b94f2f5eb01c6da13d08b9c3e91af474e36b`.
- Independently rebuilt from a normalized Git export under `umask 0077`; it matched the primary archive byte-for-byte. Extraction under `umask 0002` and `0077`, internal manifest, embedded version/commit/time, numeric `0/0` ownership, all-directory `0755`, ordinary-file `0644`, exactly three intended executables, and forbidden writable-mode checks passed.

## Verification

- Builder input/install, Framework and active lifecycle validation: PASS.
- Focused `internal/changenotification`, `internal/smtpnotification`, and `cmd/qwsg` tests: PASS.
- First aggregate build-contract attempt could not copy the new untracked package because that canonical test intentionally exports indexed files. Classified `EXPECTED BEHAVIOR`; targeted staging made the complete candidate source visible without broad staging.
- RC.4 `make engineering-test test vet fmt-check release-check` and `go test -race ./...`: PASS, including deterministic cross-mode/umask packaging.
- Required event matrix (installation, update, rollback, version transition, configuration and Guardian activation), success/failure direction, disabled behavior, transport failure separation, redaction, EN/HU/DE and duplicate suppression: PASS with deterministic in-memory transport.
- No external SMTP delivery, VPS mutation, `v1.2.0` tag, Forgejo Release or publication occurred.
- Targeted staging, diff/mode/privacy review, source commit `4f7dcc11b5ccc9f078755946995baebd31ad6870`, evidence commit `28d9a1176141b0a8fc24b92ec60b9a5914306756`, push dry-run and clean fast-forward push passed. Fresh fetch proved `HEAD == origin/main == 28d9a1176141b0a8fc24b92ec60b9a5914306756` with `0 0` divergence before lifecycle archival.

## Rollback

Verify `/tmp/qwsg-task068-execution.RhmHLa/SHA256SUMS`, follow its bounded `ROLLBACK.md`, restore only reviewed task paths from isolation, and rerun focused/full/lifecycle checks. Never use broad reset/clean or modify release refs.

## Completion state

`complete — QWSG-managed change notifications integrated; private RC.4 is READY FOR FINAL ACCEPTANCE while real delivery and final release remain separately gated`
