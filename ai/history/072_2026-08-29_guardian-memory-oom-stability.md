# Task History 072: Approved Task

## Task metadata

- Task ID: `072`
- Task slug: `guardian-memory-oom-stability`
- Status: `complete — private RC.7 READY FOR FINAL ACCEPTANCE`
- Date generated: `2026-08-29` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/072_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Task 072 completed under the approved Authority Envelope and produced private RC.7.

## Starting state

- Repository root `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, branch `main`, HEAD and fetched `origin/main` both `776ee47a10530266ce020526c813cab57c89fe62` (`0` ahead, `0` behind).
- Framework 2.0.0 validation and `bin/job --check` passed. The only worktree additions after Builder activation were the canonical Task 072 prompt/history pair.
- RC.6 artifact `dist/qwsg-1.2.0-rc.6-linux-amd64.tar.gz` matched Owner-supplied SHA-256 `7f29f4cfe361412680c439eba7d8da5ae7d949b8702ce25c460baf991de99e33`; mode/ownership were `0660`, numeric `1000:65534`.
- Owner-supplied Contabo evidence was accepted as the external failure baseline: repeated unit-scoped cgroup OOM kills under `MemoryMax=128M`, including immediate post-start failures, with historical RC.2 occurrence and misleading post-restart `systemctl` peak evidence. No external acceptance host was contacted or mutated.

## Snapshot

- Protected local snapshot: `/tmp/qwsg-task072-snapshot-20260829T000000Z` (created `2026-08-29T11:33:26Z`, mode `0700`; payloads mode `0600`).
- `repository.bundle` is a complete-history bundle at pre-task HEAD `776ee47a10530266ce020526c813cab57c89fe62`, SHA-256 `803376001d4a26be68d64d21931b0ae01528b612d71df7e79c1961733cd3df05`; `git bundle verify` passed.
- The accepted RC.6 artifact was captured with its verified Owner-supplied SHA-256. `MANIFEST.sha256` verification and artifact archive listing passed. `SNAPSHOT.md` records scope, exclusions, retention, and guarded restore procedure.

## Work performed

- Read the active prompt, all Required Reading, and the Engineering Backup Policy.
- Captured and verified repository, lifecycle, release-artifact, remote, ownership, and permission baselines.
- Began code-path inspection. The Guardian retains one scheduler cycle trace in `guardian.CapturingScheduler` for end-of-cycle publication; live observation loads the previous persisted Inventory before collecting the next one; Inventory/Snapshot stage values share typed data but persistence and deterministic identity validation perform repeated JSON encodings. These are candidate transient-amplification paths, not yet a final root-cause conclusion.
- Local one-shot inventory on the engineering host produced a 306,520-byte JSON document at 13,312 KiB maximum process RSS. The local user-bus service collector was unavailable in the sandbox, so this measurement is a clean lower-cardinality baseline rather than loaded-host proof.
- A deterministic 367-service qualified-observation fixture reproduced the loaded-workload allocation shape. One isolated cycle allocated 112,176,808 bytes cumulatively and peaked at 54,136 KiB process RSS. Runtime GC traces showed approximately 36 MiB transient managed heap and approximately 18 MiB live heap. Five cycles under the stricter diagnostic `GOMEMLIMIT=32MiB` completed at 43,204 KiB maximum RSS with no upward live-heap trend.
- Root cause: `PRODUCT/FRAMEWORK DEFECT`. Go had no cgroup-relative soft budget, while `CapturingScheduler` retained the prior full Scheduler execution trace until replacement. The cgroup also accounts collector subprocesses, non-heap mappings, and charged file cache, so process RSS and post-restart `systemctl MemoryPeak` do not describe the complete OOM boundary. Evidence did not justify increasing the 128 MiB hard ceiling.
- Added `CapturingScheduler.Take`; publication consumes and clears the full trace before the next collection. Added regression coverage proving release of the execution graph.
- Retained `MemoryMax=128M`, `TasksMax=32`, `CPUQuota=25%`, hardening, restart semantics, and `oom-kill` reporting. Added `GOMEMLIMIT=64MiB`, above the observed managed-heap working region and leaving 64 MiB for non-Go/cgroup-accounted memory. This controls GC pacing; it does not suppress or recover from OOM.
- Advanced the private identity to `1.2.0-rc.7` and declared `compat-1.2.0-rc.2-to-1.2.0-rc.7`. Fixture migration proved exact RC.2 binary/unit rollback plus configuration, credential, and state preservation.
- Candidate source commit: `07d30315cebba4bd213e76fe9fb9c32e82aa9b38`; controlled epoch `1788004085` (`2026-08-29T11:48:05Z`). Private artifact: `dist/qwsg-1.2.0-rc.7-linux-amd64.tar.gz`; SHA-256 `fe5fdeb93efe0363376d5d4b85ac915b61817a3f7775c8158c5d62cb7db7c631`.

## Verification

- Builder input, metadata, prompt/history identity, approval state, lifecycle installation, Framework validation, remote convergence, snapshot hashes, bundle readability, and RC.6 archive readability passed.
- Focused Guardian/update/CLI tests, full Go tests, full race tests, vet, formatting, Framework (25), Builder (49), Framework v2 (15), bounded diagnostic runner (11), lifecycle, test-task, build-contract, release-plumbing, migration/rollback, notification, security/redaction, and systemd unit verification passed.
- Reproducibility passed across normalized/group-writable source modes and umasks. Independent RC.7 rebuilds matched byte-for-byte. Archive/sidecar, gzip/tar integrity, manifest, numeric `0/0` ownership, canonical modes, and executable allowlist passed.
- No OVH/Contabo mutation, final tag, public artifact, Forgejo Release, or publication occurred. Real-host acceptance remains a separate Owner-controlled manual gate.

## Rollback

- Exact pre-task tracked state is recoverable by cloning the verified bundle into an empty directory. The current worktree must only be rolled back through an exact reviewed Task 072 path list; broad reset/checkout/clean operations remain forbidden. The two Builder-created lifecycle paths are removed only as explicit bounded targets when returning to the canonical pre-Task-072 idle state.

## Completion state

`complete — root cause established; bounded remediation implemented; private RC.7 READY FOR FINAL ACCEPTANCE`

Candidate source commit is recorded above. The lifecycle-closing repository commit and `origin/main` convergence are verified after final integration.
