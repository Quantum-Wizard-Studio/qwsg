# Task History 070: Approved Task

## Task metadata

- Task ID: `070`
- Task slug: `deterministic-update-migration-rc5`
- Status: `complete — READY FOR FINAL ACCEPTANCE`
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
- Created reviewed candidate-source commit `3f2f7822473d3d97c38c171f006e8b263250321c`. Built private `qwsg-1.2.0-rc.5-linux-amd64.tar.gz` with controlled epoch `1787943678` (`2026-08-28T19:01:18Z`).
- The first independent export rebuild was invoked from the live repository instead of the exported module root and Go rejected the external package path before creating output. Classified `TEST OR ACCEPTANCE DEFECT`; rerunning the identical normalized export from its own root succeeded without changing candidate source or bytes.
- Independently rebuilt from a deliberately group-writable Git export under `umask 0077`; archive and sidecar matched byte-for-byte. Audited integrity, manifest, provenance, version, numeric `0/0` ownership, canonical `0755` directories, exactly binary/install/uninstall executable, all other files `0644`, and extraction behavior.

## Verification

- Protected Builder input directory mode `0700`, files `0600`; `task-builder.sh --check-input`: `Structured owner input: VALID`.
- Builder installation: `APPROVED AND AUTHORIZED FOR EXECUTION`; active Task 070 prompt/history identity, Framework and lifecycle validation: PASS.
- Fresh Git fetch, exact baseline/remote comparison, RC.4 hash, absence of final tag, snapshot integrity/readability and isolated restore rehearsal: PASS.
- Focused `go test ./internal/update ./internal/changenotification ./cmd/qwsg`: PASS after the fixture-contract correction.
- `make engineering-test test vet fmt-check`: PASS, including build provenance, Framework (25), diversion (36), lifecycle (29), Builder (49), all Go packages, vet and formatting.
- Framework v2 semantic (15), bounded diagnostic runner (11), active lifecycle and diverted-task audit: PASS.
- Repository-wide `go test -race ./...`: PASS.
- `make release-check`: RC.5 plumbing PASS and Task 067 cross-umask/source-mode byte reproducibility regression PASS.
- Candidate archive SHA-256: `ca3123f313c2fb95138fceb0fc2af17c5a8295b28e7b13e60f5d7fd57ed19faf`; sidecar-file SHA-256: `17e09c6c54a1f502b9dcd745b931f7ec3c8db4bd4a26088d72916dad12187b76`.
- Primary versus independent normalized/group-writable export archive and sidecar `cmp`: PASS. `gzip -t`, internal `sha256sum -c MANIFEST.sha256`, embedded CLI identity/provenance and tar ownership/mode inspection: PASS.
- Final focused update/change-notification/CLI tests after RC.2/RC.5 event-direction assertions: PASS.

## Rollback

Verify `/tmp/qwsg-task070-execution.ulaK8c/SHA256SUMS` and follow its `ROLLBACK.md`. Restore only explicit reviewed Task 070 paths from isolated extraction; never use broad reset/checkout/restore/clean. Rerun focused tests, Framework, Git and lifecycle checks after any restore.

## Completion state

`complete — QWSG-069-F001 remediated; private RC.5 READY FOR FINAL ACCEPTANCE`

## Final decision

1. Decision: `COMPLETE`.
2. Remediation: explicit reusable declarative compatibility registry plus preflight and privileged-boundary revalidation; no version bypass or guessed path.
3. Supported required path: `QWSG 1.2.0-rc.2 -> QWSG 1.2.0-rc.5`; undeclared paths fail closed before Guardian/package mutation.
4. Preservation/security: compatible schemas require no transformation; configuration, protected credential and persistent state fixtures remain byte-identical and secret values are not emitted.
5. Guardian/systemd: managed user unit is transactionally replaced/restored with fixed `/usr/local/bin/qwsg` path and preserved `MemoryMax=128M`/`TasksMax=32`; enabled/active orchestration remains unchanged and validated by existing suites. Real service acceptance is intentionally deferred.
6. Rollback: protected transaction metadata records versions/commit/destinations/modes/hashes; isolated RC.5 -> RC.2 restored binary/unit and preserved data byte-for-byte; tamper/failure tests remain fail-safe.
7. Notifications: Task 068 dispatcher remains transport-independent; update/rollback SUCCESS/FAILED, RC.2/RC.5 direction, operation identity/action requirement, EN/HU/DE, duplicate behavior, redaction and transport-result separation pass deterministic tests.
8. Fixture/regression: credential-free `internal/update/testdata/rc2-installed-state.json` reproduces the compatibility-relevant RC.2 state without live-host/private data.
9. Candidate: `qwsg-1.2.0-rc.5-linux-amd64.tar.gz`, SHA-256 `ca3123f313c2fb95138fceb0fc2af17c5a8295b28e7b13e60f5d7fd57ed19faf`, source `3f2f7822473d3d97c38c171f006e8b263250321c`; reproducibility PASS.
10. Remaining limitation: real OVH/Contabo update/rollback/re-update/reboot/coexistence and mailbox acceptance remains mandatory in the subsequent final-acceptance task. No final tag, public artifact, Forgejo Release, VPS mutation or external SMTP delivery occurred.
11. Classification: `READY FOR FINAL ACCEPTANCE`.
