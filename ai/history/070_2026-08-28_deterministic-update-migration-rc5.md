# Task History 070: Approved Task

## Task metadata

- Task ID: `070`
- Task slug: `deterministic-update-migration-rc5`
- Status: `active — implementation in progress`
- Date generated: `2026-08-28` UTC
- Human authority: Project Owner: Attila
- Preferred owner communication language: English
- Related prompt: `ai/prompts/070_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from the complete Owner-authorized Task 070 specification. Exact `APPROVE`, protected structured input validation, Framework validation and lifecycle consistency passed. Execution started immediately under the approved Framework 2.0 Authority Envelope.

## Starting state

- Verified ordinary user `attila`, exact QWSG root, Framework `2.0.0`, canonical HTTPS origin, branch `main`, and fresh fetched `HEAD == origin/main == a28f29b69712a919b34ecee6e6cd66f504d755d2` with `0 0` divergence.
- Before Builder installation the full worktree/index was clean and lifecycle canonically idle after Task 069 archived `complete — BLOCKED by QWSG-069-F001; RC.4 not released`. After installation only the expected Task 070 prompt/history were untracked.
- Verified Task 069's authoritative RC.2 refusal before package replacement and the local immutable RC.4 reference artifact SHA-256 `adeb591605c0d37a5fc98d541125ca388cd4561703d0f0823bba931bc7d08684`. No `v1.2.0*` tag exists.
- OVH and Contabo were neither contacted nor mutated; their Task 069 terminal states are accepted only from preserved history as required.

## Snapshot

- Protected snapshot root: `/tmp/qwsg-task070-execution.ulaK8c`, directory mode `0700`, files mode `0600`, retained through Owner acceptance and the QWSG 1.2.0 release rollback-window closure.
- Captured complete tracked `HEAD` as `tracked-baseline.tar` plus exact Builder-created Task 070 prompt/history before-images. README, bounded rollback instructions and deterministic `SHA256SUMS` record scope, exclusions, collision policy and retention.
- All payload checksums passed; archive listing/readability passed; isolated protected extraction reproduced repository version `1.2.0-rc.4`; prompt/history byte comparisons passed. Nothing was extracted over the live worktree.

## Work performed

- Mapped the complete Owner specification losslessly into the concise Framework 2.0 Builder contract and transactionally installed Task 070.
- Read the active prompt, every required governance/configuration document, engineering backup policy and relevant Task 064/067/068/069 updater, release, notification and blocker evidence.
- Replaced the migration planner's compound conditional with a small declarative path registry. The registry preserves earlier supported paths and adds only `compat-1.2.0-rc.2-to-1.2.0-rc.5`; undeclared RC.2 -> RC.4 and all unknown/malformed/equal/downgrade paths remain fail-closed.
- Extended each path contract with stable identity, source/target and configuration/Guardian/scheduler/operator-state schemas, no-mutation declaration, user configuration/credential/state preservation and managed Guardian-unit replacement semantics. RC.2 and RC.5 require no schema transformation.
- Added installed-configuration validation before Guardian stop/package mutation and independent migration-plan validation inside the privileged helper after candidate re-verification. Package replacement remains fixed to release-owned destinations.
- Added a credential-free deterministic RC.2 installed-state fixture and an isolated end-to-end transaction test. It selects the exact path, replaces the binary/unit with RC.5, proves configuration/credential/state byte preservation, validates transaction metadata, rolls back and proves the RC.2 binary/unit plus user data byte-for-byte.
- Advanced private source identity and release plumbing to `1.2.0-rc.5`; added RC.5 notes, changelog, installation commands, compatibility/rollback/security documentation and Task 069-preserving acceptance status.
- The first fixture run selected the correct migration but its snake-case JSON schema did not populate untagged Go fields. Classified `TEST OR ACCEPTANCE DEFECT`; explicit JSON tags corrected the fixture contract and the unchanged product path passed.

## Verification

- Protected Builder input directory mode `0700`, files `0600`; `task-builder.sh --check-input`: `Structured owner input: VALID`.
- Builder installation: `APPROVED AND AUTHORIZED FOR EXECUTION`; active Task 070 prompt/history identity, Framework and lifecycle validation: PASS.
- Fresh Git fetch, exact baseline/remote comparison, RC.4 hash, absence of final tag, snapshot integrity/readability and isolated restore rehearsal: PASS.
- Focused `go test ./internal/update ./internal/changenotification ./cmd/qwsg`: PASS after the fixture-contract correction.
- `make engineering-test test vet fmt-check`: PASS, including build provenance, Framework (25), diversion (36), lifecycle (29), Builder (49), all Go packages, vet and formatting.
- Framework v2 semantic (15), bounded diagnostic runner (11), active lifecycle and diverted-task audit: PASS.
- Repository-wide `go test -race ./...`: PASS.
- `make release-check`: RC.5 plumbing PASS and Task 067 cross-umask/source-mode byte reproducibility regression PASS.

## Rollback

Verify `/tmp/qwsg-task070-execution.ulaK8c/SHA256SUMS` and follow its `ROLLBACK.md`. Restore only explicit reviewed Task 070 paths from isolated extraction; never use broad reset/checkout/restore/clean. Rerun focused tests, Framework, Git and lifecycle checks after any restore.

## Completion state

`active — implementation in progress`
