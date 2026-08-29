# QWSG 1.2.0 Final Acceptance and Release Record

## Evidence boundary

The Project Owner supplied and accepted the Contabo and OVH observations below after executing the real-host RC.7 acceptance. Task 073 records that evidence as Owner acceptance; it does not claim that the release agent contacted or mutated either host. Local release construction, publication and independent distribution checks are recorded separately as they are completed.

## Accepted implementation and promotion

- Accepted candidate: `1.2.0-rc.7`.
- Accepted artifact: `qwsg-1.2.0-rc.7-linux-amd64.tar.gz`.
- Accepted artifact SHA-256: `fe5fdeb93efe0363376d5d4b85ac915b61817a3f7775c8158c5d62cb7db7c631`.
- Accepted implementation source: `07d30315cebba4bd213e76fe9fb9c32e82aa9b38`.
- Task 072 final repository commit: `79585da0b2fa95f8b7aaa3e14fc15182d690a3b7`.
- Product-equivalence proof: the only intervening commit closes Task 072 and changes only its prompt location/status and history. No behavior-bearing or packaging source changed.
- Canonical promotion: preserve the accepted RC.7 product implementation; apply final `1.2.0` identity, release notes/documentation/plumbing, and the minimum declarative RC.2-to-final compatibility entry using the already accepted migration mechanism. No feature or unrelated behavior change is permitted.

## Contabo loaded-host acceptance supplied by the Project Owner

- Ubuntu 24.04.4 LTS amd64 with HestiaCP 1.10.4 and its existing web, mail, database and security services.
- Installed RC.2 baseline verified; RC.2 to RC.7 deterministic update PASS; RC.7 runtime stability PASS; RC.7 to RC.2 deterministic rollback PASS with exact RC.2 binary/version restoration; RC.2 to RC.7 re-update PASS.
- Reboot PASS: Guardian automatically started, `NRestarts=0`, and no OOM, kill, failure or restart followed reboot.
- Resource contract preserved: `MemoryMax=128 MiB`, `TasksMax=32`, and `GOMEMLIMIT=64MiB` active.
- Hestia coexistence PASS; all important Hestia services remained running.
- Local Exim STARTTLS verified the host certificate with credential-free `auth=none`. Test notification SMTP acceptance PASS and Project Owner manually confirmed receipt in the actual mailbox.
- Configuration lifecycle notification SMTP acceptance PASS and the Project Owner manually confirmed receipt of the localized Hungarian message. The notification exposed only the changed key, not its secret or value; operation and notification results remained separate.
- The prior loaded-host OOM blocker is closed.

## OVH clean/preserved-state acceptance supplied by the Project Owner

- Ubuntu 24.04 amd64 clean/preserved-state acceptance host.
- RC.7 artifact SHA-256 matched exactly and every packaged file passed `MANIFEST.sha256`.
- Canonical RC.6 uninstallation removed release-owned artifacts while preserving user configuration and state.
- Hungarian guided RC.7 installation passed package verification, installation, Guardian activation and readiness synchronization. Progress reached 87%, 97% and 100%; installed version was `1.2.0-rc.7`; update policy was `notify`; optional notification remained disabled.
- Readiness was Partial only because the optional capability was not configured; all mandatory Guardian requirements passed.
- Before reboot: `NRestarts=0`, `MemoryCurrent=69873664`, `MemoryPeak=70332416`, `MemoryMax=134217728`, `TasksMax=32`, and `GOMEMLIMIT=64MiB`.
- After reboot: Guardian automatically started; `ExecCondition` configuration validation succeeded; service remained active beyond three minutes; Tasks were 8/32; memory was approximately 44.6 MiB with approximately 82.8 MiB peak.
- Final evidence: `NRestarts=0`, `MemoryCurrent=46837760`, `MemoryPeak=86880256`, `MemoryMax=134217728`, `TasksMax=32`, `GOMEMLIMIT=64MiB`; journal contained no OOM, kill, failure or restart evidence.
- The prior guided-installer readiness synchronization blocker is closed.

## Final-version compatibility gate

Final 1.2.0 declares a direct deterministic compatibility path from an actual installed `1.2.0-rc.2` identity. The existing credential-free installed-state fixture must prove package replacement to final `1.2.0`, schema compatibility, configuration/credential/state preservation, final Guardian-unit installation and exact RC.2 rollback. RC.7-only coverage is not accepted as a substitute.

## Publication ledger

- Mandatory local release validation: `PASS` — focused final-version update/rollback, full Go, repository-wide race, vet, formatting, Framework/configured engineering suites, Builder/lifecycle/diversion/test-task, shell syntax, systemd static, release plumbing, build provenance, cross-umask/source-mode reproducibility, security/redaction and notification coverage passed. The first focused attempt selected a read-only default Go cache; it was classified as an environmental issue and passed unchanged with task-private writable caches.
- Deterministic twin-build identity: `PENDING`.
- Final release commit and epoch: `PENDING`.
- Frozen archive, sidecar, MANIFEST and binary identities: `PENDING`.
- Annotated `v1.2.0` local/remote tag: `PENDING`.
- Forgejo v1.2.0 Release and assets: `PENDING`.
- Independent external wget/curl retrieval, checksum, tar, MANIFEST and structure verification: `PENDING`.
- Final repository/lifecycle state: `PENDING`.

No `RELEASED` decision is valid while any mandatory ledger item remains pending or failed.
