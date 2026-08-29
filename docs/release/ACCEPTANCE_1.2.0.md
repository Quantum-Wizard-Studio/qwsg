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
- Deterministic twin-build identity: `PASS` — two independent `git archive` exports of the exact release commit used isolated source, output and Go-cache roots and produced byte-identical archive and sidecar bytes.
- Final release commit and epoch: `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`; tree `e7aa50f4ed405598b0d7b9a9224d8dd183e9a703`; `SOURCE_DATE_EPOCH=1788022414`; `2026-08-29T16:53:34Z`.
- Frozen archive: `qwsg-1.2.0-linux-amd64.tar.gz`; size `3524214`; SHA-256 `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`.
- Frozen sidecar SHA-256: `22242a8d0702e7dbe33c9beee84b8ef497e6372bf08e26747b754a01382e81c6`; MANIFEST SHA-256: `76fdd7809eaeffe1d923338a29b248e9413e5d1517375a5c651359ae62b118b3`; binary SHA-256: `c922f51873ee9e9428da4b900aedec31f08f9e16f86b7aba680a5bf4e96755a6`.
- Package validation: `PASS` — 27 readable members, safe single-root layout, every MANIFEST entry verified, static linux-amd64 binary, exact version/full commit/build time, numeric `0/0` archive ownership, canonical directory/file modes, and only `bin/qwsg`, `install.sh`, and `uninstall.sh` executable.
- Annotated `v1.2.0`: `PASS` — tag object `ac395b568b8e1f83c0ef85c9aa02f98c15402af0` exists locally and remotely and peels exactly to `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`.
- Forgejo Release: `PASS` — Release ID `3`, title `QWSG 1.2.0`, final/non-draft/non-prerelease, exact target commit, exactly the versioned archive (`3524214` bytes) and sidecar (`96` bytes).
- Independent external retrieval: `PASS` — credential-free wget and `curl -fLO` each returned the two actual versioned Release assets into separate directories. Both archives were non-empty, matched SHA-256 `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`, passed their downloaded sidecars, opened as tar, passed every extracted MANIFEST entry and contained the expected executable/package structure and embedded provenance.
- Canonical archive URL: `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v1.2.0/qwsg-1.2.0-linux-amd64.tar.gz`.
- Final repository/lifecycle state: `PASS` — the release evidence commit was pushed, Task 073 was archived without a successor, canonical idle validation passed, and the lifecycle-closure commit converged on canonical `main`.

## Decision

`RELEASED` — every mandatory local, provenance, packaging, tag, publication and external distribution invariant passed. No product feature, architecture redesign, unrelated refactor or host mutation was introduced by final promotion.
