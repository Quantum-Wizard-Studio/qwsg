# Task History 079: Approved Task

## Task metadata

- Task ID: `079`
- Task slug: `guardian-isolated-automatic-release-checking`
- Status: `complete — implementation and production acceptance validated; canonical closure in progress`
- Date generated: `2026-09-03` UTC
- Human authority: Project Owner — explicitly authorized Task 079 and supplied the exact engineering scope with APPROVE on 2026-09-03 UTC.
- Preferred owner communication language: English
- Related prompt: `ai/prompts/079_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured Owner input. Exact `APPROVE` was received, the active pair passed lifecycle validation, and implementation began under Framework 2.0 Standard Execution Authority. Task 080 remains prohibited.

## Starting state

- Repository `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, branch `main`, canonical origin `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg.git`.
- Pre-intake `HEAD == origin/main == 7d16e1c7cc486940e8c9b11b498caecde4d3ec59`; divergence `0/0`; index and worktree clean; lifecycle canonical idle after complete Task 078.
- Fresh `git fetch origin main`, Framework 2.0 validation, Task 078 closure identity, and active Task 079 prompt/history validation: PASS.
- `VERSION` and the installed supported package both report `1.2.0`; installed commit is `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`.
- Production source, required media type, key ID, and fingerprint match the accepted Task 078 authority exactly.
- The active production `qwsg-guardian.service` remained `active/running`, result `success`, with 8 tasks during read-only acceptance inspection.

## Snapshot

- Pre-intake snapshot: `/tmp/qwsg-task079-intake.9Ly6eJ`, mode `0700`, evidence mode `0600`; complete verified Git bundle SHA-256 `6ff6a30a4b8be033015aba43c44a45989f5531a89a3f8b878db61bb3efd6ecbf`; readable lifecycle archive SHA-256 `ff910ef9ecc4da0aa9690c8e4d6072c402e299af5a5f4fe933d2ad936b02004a`.
- Execution snapshot: `/tmp/qwsg-task079-execution.l9jB5Z`, mode `0700`, evidence mode `0600`; complete verified Git bundle with the same baseline SHA-256; readable target archive SHA-256 `ab4a63db8d17538569a4ad66f9abab80eaedf1bd4dbfd869a8c67d5a51e1c5c9`; ACL evidence SHA-256 `1581155a5e241ec83e037428c0f4edb552b5034882434acafef4d8d7c30b2a88`.
- Both snapshots exclude credentials/private keys and remain retained through Owner acceptance.

## Work performed

- Added `internal/guardian.ReleaseCheckService`, the sole sequential automatic release-awareness scheduler. It derives due time from `qwsg.update-awareness/1` `last_attempt.at`, uses a 24-hour default and 35-second child deadline, treats missing state as initially due, suppresses network on unsafe state, waits a full interval after all attempts, and exits on Guardian cancellation.
- Added `StartupSink`, which opens the release-check gate only after the first local Runtime cycle has been published and the retained engineering graph released.
- Composed the side loop with the existing Guardian signal context and joined it on shutdown. It is not a Runtime/Pipeline stage and cannot change Guardian health or checkpoint truth.
- Extracted one `updateAwarenessManager` constructor shared by manual `update check` and automatic Guardian checking. The exact Task 078 source, trust, Task 075 classifier, Task 076 authentication/evaluation, Task 077 persistence, rollback/future-index protection, and failure categories remain single-source.
- Added deterministic due/not-due/restart, startup gate, sequential non-overlap, failure continuation, corrupt-state fail-closed, cancellation, and cleanup tests. Existing Task 076/077 tests retain authenticated current/newer, transport/timeout/media/parser/signature/key/rollback/future-index, persistence, zero-network status, and side-effect coverage.
- Updated architecture, system map, security/privacy, operations, upgrade, README, changelog, and new EN/HU automatic-release-checking operator guides. Configuration was intentionally not extended: the existing canonical schedule collection owns Command/Pipeline work and is not a safe semantic fit for this isolated outbound side job.

## Verification

Builder input, metadata, prompt/history identity, approval state, lifecycle installation, both snapshot bundles/archives/checksums/modes, and focused formatting: PASS.

- Focused `go test ./internal/guardian ./cmd/qwsg`: PASS.
- Focused race tests for Guardian, awareness, release discovery, and CLI: Guardian/awareness/CLI PASS; the first restricted run was blocked only because the sandbox denied the existing loopback TLS listener. Classified `ENVIRONMENTAL ISSUE`; the unrestricted full race suite subsequently passed every package.
- `make fmt-check vet test`: PASS for every package in an unrestricted loopback-capable test context.
- `go test -race ./...`: PASS for every package.
- `make engineering-test`: PASS after exact Task 079 staging satisfied the
  Git-index-only build-contract fixture; Framework 25, diversion 36, lifecycle
  29, Builder 49, and deterministic checkout/export build assertions passed.
- `scripts/test-release-reproducibility.sh`: PASS across source-mode and umask
  variations. `make release-authority-check`: PASS for byte-identical release
  authority tools and canonical signing input.
- `make release-check` stopped at the intentional clean-output precondition
  because the immutable published `dist/qwsg-1.2.0-linux-amd64.tar.gz` exists;
  classified `EXPECTED BEHAVIOR`. The artifact was not removed or overwritten,
  and the constituent reproducibility test passed independently.
- Restricted full test first attempt also lacked module-network and loopback permissions; classified `ENVIRONMENTAL ISSUE`, with no implementation correction made solely for that environment.

Bounded production acceptance used isolated private state at `/tmp/qwsg-task079-acceptance.uhWDHx` without mutating the active installed state:

- Exact production HTTPS/media-type/Ed25519 verification: PASS. Installed `1.2.0` classified against authenticated stable `1.2.0` as `current`.
- Valid isolated awareness record: 1,216 bytes, mode `0600`, SHA-256 `56f950b06edc01ab4b9729bc2aa339a74785529798de973e7beb6821d0c4239b` after acceptance evidence finalization.
- A brief isolated Guardian run completed local cycles and graceful TERM handling. The awareness record hash remained byte-identical to its pre-run value `56f950b06edc01ab4b9729bc2aa339a74785529798de973e7beb6821d0c4239b`, proving the subsequent cycle was not due and did not repeat retrieval.
- `strace -f -e trace=network` around isolated `qwsg update status` recorded zero network syscalls; output remained `current`, installed/available `1.2.0`.
- Isolated state contained only awareness, Guardian checkpoint/locks/scheduler, Current Operator State, and inventory evidence. No archive, sidecar, staging, transaction, notification queue, artifact download, installation, version mutation, restart, notification, or unrelated configuration change occurred.
- Active production Guardian inspection after acceptance: `ActiveState=active`, `SubState=running`, `Result=success`, `TasksCurrent=8`, `MemoryCurrent=107466752`; no production service/configuration mutation was performed.

## Rollback

Verify the execution snapshot hashes and exact current target identities. Restore only enumerated Task 079 before-images from the target archive, and remove only Task 079-created paths after confirming their recorded prior absence. Preserve unrelated work and metadata. After a pushed commit, use a new exact-path revert commit rather than rewriting history. Never use broad reset, checkout, clean, wildcard removal, or extraction over the live worktree. Re-run focused/full/race/framework/lifecycle/Git/privacy and production-state checks after rollback.

## Completion state

`complete — implementation and production acceptance validated; canonical closure in progress`
