# Task History 069: Approved Task

## Task metadata

- Task ID: `069`
- Task slug: `qwsg-1-2-0-final-acceptance-release`
- Status: `complete — BLOCKED by QWSG-069-F001; RC.4 not released`
- Date generated: `2026-08-28` UTC
- Human authority: Project Owner: Attila
- Preferred owner communication language: English in the approved prompt; validated project configuration currently declares Hungarian for owner-facing progress
- Related prompt: `ai/prompts/069_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from the complete Owner-approved Task 069 specification. Exact `APPROVE`, Builder input validation, transactional installation, Framework validation and lifecycle consistency passed. Execution started under the approved Framework 2.0 Authority Envelope.

## Starting state

- Verified the exact QWSG root, Framework 2.0.0, canonical HTTPS origin and `main` branch. Before Builder installation, `HEAD` was `48e4d47b6987afcaa5986c86c4d321fd2d6c7f94`, tracked/index state was clean and lifecycle was idle after complete/archived Task 068.
- After an authorized fresh fetch, `HEAD == origin/main == 48e4d47b6987afcaa5986c86c4d321fd2d6c7f94` with `0 0` divergence. Only the expected Builder-created Task 069 prompt/history are untracked.
- Verified repository version `1.2.0-rc.4`; artifact `dist/qwsg-1.2.0-rc.4-linux-amd64.tar.gz` size `3516915`, SHA-256 `adeb591605c0d37a5fc98d541125ca388cd4561703d0f0823bba931bc7d08684`, sidecar SHA-256 `d68523a5bdea974b23629e91b9b0b94f2f5eb01c6da13d08b9c3e91af474e36b`, embedded source `4f7dcc11b5ccc9f078755946995baebd31ad6870`, build time `2026-08-28T09:05:01Z`, and CLI identity `QWSG 1.2.0-rc.4`.
- Archive listing/readability, safe intended member types, internal manifest availability, numeric ownership `0/0`, all directories `0755`, exactly `bin/qwsg`, `install.sh`, and `uninstall.sh` executable `0755`, and all other regular files `0644` passed. No local `v1.2.0*` tag exists.
- The first sidecar check was invoked from the repository root although its canonical relative filename is scoped to `dist/`, and the first metadata read assumed a noncanonical archive root. Classified `TEST OR ACCEPTANCE DEFECT`; corrected read-only commands passed immediately. Candidate bytes were not changed.
- The designated Contabo identity resolved to an address already present in protected `known_hosts`; strict host-key verification and the existing acceptance key succeeded. Read-only preflight proved user `attila`, Linux x86_64, installed QWSG `1.2.0-rc.2` from commit `c260dc18c2004473ec55496d16e66718fd128865`, and active `qwsg-guardian.service`. No host mutation occurred.
- Contabo QWSG baseline: configuration validation, readiness and update status PASS; Guardian is enabled/active/running with one matching process, 8 tasks of 32, `MemoryMax=128 MiB`, current memory about 51.3 MiB and peak about 114.1 MiB. Binary mode `0755`, user unit `0644`, configuration/state directories `0700`.
- Contabo coexistence baseline reported active HestiaCP, Nginx, Apache, MariaDB, PostgreSQL, Exim, Dovecot, BIND/named, Fail2Ban, ClamAV daemon and twelve PHP-FPM services. SpamAssassin, ClamAV freshclam, UFW and nftables units reported inactive; quota command reported no enabled mount entries. These are pre-existing observations, not QWSG regressions. Roundcube/webmail, SSL/domain behavior, firewall rules and quota functionality still require their safe canonical checks.
- Initial credential-directory probe used an incorrect data-path assumption and reported absent; source inspection classified this as `TEST OR ACCEPTANCE DEFECT`. The corrected configuration-relative credential directory is also absent, notification preflight returns success in the existing optional/disabled state, and no protected credential is currently available. No secret, recipient or SMTP/provider value was read. Real lifecycle-notification acceptance therefore requires a safe canonical non-secret configuration path or the Owner-only evidence boundary; it is not yet PASS.
- Owner supplied the designated OVH identity and ordinary SSH user. DNS matched the supplied endpoint, but the live ED25519 key differed from the protected stale record. Strict host-key verification correctly stopped before authentication or command execution. Classified `ENVIRONMENTAL ISSUE — MATERIAL HOST IDENTITY DIFFERENCE`; no entry was changed until the Owner independently verified the new fingerprint through the trusted OVH console. Private host/network identity and fingerprint values are intentionally omitted from canonical history.
- Owner independently verified the new ED25519 fingerprint through the trusted OVH console and explicitly authorized replacement of only the stale entry. The live key was fetched separately, independently fingerprinted, and matched exactly before the IP entry was replaced; `StrictHostKeyChecking=yes` remained mandatory. `ssh-keygen -R` also rewrote its conventional `known_hosts.old` backup as a side effect; the unrelated previous three-line backup was immediately reconstructed byte-for-byte from the protected before-image while retaining the verified active entry.
- The subsequent strict connection reached the verified OVH endpoint but authentication for user `ubuntu` failed with `Permission denied (publickey,password)`: the existing QWSG acceptance public key is not currently authorized for that account. Classified `ENVIRONMENTAL ISSUE — ACCESS BOUNDARY`. No remote command executed and no OVH host mutation occurred. Owner-side installation of the non-secret acceptance public key, or another already-authorized non-secret access mechanism, is required before sterile-host preflight.
- Owner installed the public acceptance key with verified `0700` `.ssh` and `0600` `authorized_keys`; strict authentication then succeeded as ordinary UID 1000. Read-only preflight proved Ubuntu 24.04 x86_64, noninteractive sudo availability, lingering enabled, sufficient disk space and no pending reboot marker.
- Mandatory clean-host baseline FAILED before QWSG mutation: `/usr/local/bin/qwsg`, the installed user unit, per-user configuration and state all exist; multiple prior QWSG acceptance directories remain. Installed version is QWSG `1.1.0`, commit `305f4088e94b14d6cbb3114eb8cce4e32d847c16`, configuration validation passes, but readiness is nonzero and the enabled Guardian unit is `failed` with `Result=oom-kill`, `NRestarts=136`, `MainPID=0`, `MemoryMax=128 MiB`, `TasksMax=32`.
- The first broad process-count probe matched its own remote diagnostic shell command and reported one; classified `TEST OR ACCEPTANCE DEFECT`. Exact systemd `MainPID=0` and the captured command line prove no Guardian runtime process. This does not change the material non-sterile baseline.
- Classified `ENVIRONMENTAL ISSUE — MATERIAL APPROVED-BASELINE DIFFERENCE`. Per the Task 069 STOP rule, no candidate transfer, installation, removal, cleanup, service mutation or Contabo mutation followed. The OVH host must be returned through an independently authorized sterile-host mechanism (normally rebuild/reinstall), not simplified by Task 069 cleanup, before Phase 3 can start.

## Snapshot

- Protected snapshot root: `/tmp/qwsg-task069-execution.DcQUOx`, mode `0700`, retained through Owner acceptance and the QWSG 1.2.0 release rollback-window closure.
- Captured complete tracked `HEAD` archive, exact Builder-created Task 069 prompt/history before-images, and immutable RC.4 artifact/sidecar. Payload files and checksum record are mode `0600`; `README.md`, `ROLLBACK.md`, and deterministic `SHA256SUMS` record purpose, scope, exclusions, collision policy, retention and exact bounded restore rules.
- `sha256sum -c` passed for every snapshot object, archive listing/readability passed, and extraction into the empty protected `restore-rehearsal/` directory reproduced version `1.2.0-rc.4`. Prompt/history comparisons and candidate sidecar verification passed. No archive was extracted over the live worktree.
- External host mutations remain gated on their own canonical QWSG pre-state backup/snapshot and state-verified rollback evidence.

## Work performed

- Mapped the complete Owner specification into the current concise Framework 2.0 Builder contract without adding product features or reducing the acceptance matrix.
- Read the active prompt and all required governance/configuration documents plus the engineering backup policy; began relevant release, installer, updater, rollback, notification and Tasks 063-068 evidence review.
- Performed only local read-only candidate inspection and a strict-host-key Contabo read-only identity/version/service preflight after snapshot creation. RC.4 remains immutable and neither acceptance host has been mutated.
- Completed the mandatory local pre-acceptance regression matrix. The release-check collision guard was handled by moving only the two verified RC.4 files into protected snapshot holding, running the checks, and restoring them through an EXIT trap; the restored artifact retained its exact approved SHA-256.
- Owner rebuilt the designated OVH host to clean Ubuntu 24.04 after the first preflight correctly found an old QWSG 1.1.0 installation. The reinstall changed the SSH key; Owner independently verified the new fingerprint through the OVH console, the local keyscan matched, and only the stale OVH entry was replaced while strict checking remained enabled.
- On the verified sterile OVH host, privately transferred exactly the RC.4 archive/sidecar, independently verified outer hash, archive integrity, internal manifest, provenance and canonical extracted modes, and passed the read-only install assessment.
- Executed the actual PTY guided-install path as the ordinary user with its narrow sudo package helper. Version/configuration/readiness, installed files/modes, enabled active Guardian, restart/PID transition, 128 MiB/32-task limits and absence of unrelated system paths passed.
- Documented uninstall removed only unchanged release-owned artifacts and preserved the exact configuration hash and five state files. Guided reinstall restored the package/service/configuration state. Its immediate wizard readiness check briefly returned exit 4 before the first canonical cycle; the documented follow-up readiness check passed with satisfied canonical evidence and stable active service, classified `EXPECTED BEHAVIOR`.
- Enabled exact-user lingering, rebooted the OVH host, and proved host return, boot-time Guardian activation, RC.4 identity, valid configuration, readiness, enabled/active service, zero restarts and resource limits. After the later release blocker, disabled/stopped Guardian, uninstalled RC.4, disabled lingering, and removed only Task 069-created QWSG config/state/acceptance paths. Final OVH checks proved no QWSG process, unit, system path or top-level acceptance path: sterile state restored.
- Created protected Contabo pre-update snapshot `/tmp/qwsg-task069-contabo-preupdate.I5SwQ1`, mode `0700`, containing RC.2 binary/unit, complete QWSG config/state/user-systemd state, manifest and checksums. Every object passed SHA-256 verification; payload remains outside Git through the release rollback window.
- Privately transferred and verified RC.4 on Contabo. The existing RC.2 configuration had no notification recipient or credential. Owner supplied the administrator recipient; the existing local Exim route, STARTTLS certificate, routing and credential-free `auth=none` path were verified without creating a credential.
- Candidate configuration-event delivery and an explicit notification test were SMTP-accepted. Because the canonical update requires sudo unavailable to the engineering SSH session, a protected mode-0700 Owner-run script performed configuration plus update in one bounded interactive-sudo sequence without storing a password.
- The real update returned exit 1: `Update refused: no deterministic compatible migration path.` Its failure notification independently reported `Admin notification: ACCEPTED`. SMTP acceptance was proven; actual mailbox receipt was not claimed because the underlying mandatory update gate had already failed.
- The first pre-sudo configuration rehearsal exposed that installed RC.2 rejects RC.4's new lifecycle field and briefly restarted. Each attempt was immediately restored from the exact pre-update configuration snapshot; the empty failed-update state directory was removed, config/readiness passed, and restart counts stabilized. Final Contabo binary, unit and configuration hashes exactly matched the pre-task RC.2 snapshot.
- Post-blocker coexistence retained the baseline service states: active HestiaCP, Nginx, Apache, MariaDB, PostgreSQL, Exim, Dovecot, BIND/named, Fail2Ban, ClamAV daemon and twelve PHP-FPM services; baseline-inactive SpamAssassin, freshclam, UFW/nftables units and zero quota mounts remained unchanged. The server hostname's HTTPS endpoint was unreachable both as a safe post-check and lacked a matching pre-proof, so web/webmail/SSL remained `ENVIRONMENTAL ISSUE / INCONCLUSIVE`, not a QWSG regression. No package replacement occurred.
- Preserved the protected Owner update log/exit in the local Contabo snapshot, removed the exact Task 069 Contabo acceptance directory, and verified RC.2 config/readiness plus active enabled Guardian and resource limits.
- Targeted staging, staged diff/mode/privacy review, `git diff --cached --check`, commit `3f7af8361496340b3d9b602b19ccf1566a05c53e`, push dry-run and clean fast-forward push passed. The completed prompt was then moved to its exact Task 069 archive path for canonical idle closure.

## Verification

- Protected Builder input directory mode `0700`, files `0600`; `task-builder.sh --check-input`: `Structured owner input: VALID`.
- Builder installation: `APPROVED AND AUTHORIZED FOR EXECUTION`; `ai/prompts/069_CURRENT_TASK.md`, matching history, Framework and active lifecycle validation: PASS.
- Fresh Git fetch and exact baseline comparison: PASS.
- RC.4 external checksum, internal metadata, CLI, ownership and canonical permission checks: PASS.
- Snapshot checksum, readability, isolated restore rehearsal and before-image comparison: PASS.
- `make engineering-test test vet fmt-check`: PASS, including build contract, Framework (25), diversion (36), lifecycle (29), Builder (49), full Go tests, vet and formatting.
- Framework v2 semantic suite (15), bounded diagnostic runner suite (11), shell syntax across canonical shell entry points, and repository-wide `go test -race ./...`: PASS.
- `make release-check`: release plumbing and deterministic reproduction across different umasks/source modes PASS. RC.4 was restored byte-identically with SHA-256 `adeb591605c0d37a5fc98d541125ca388cd4561703d0f0823bba931bc7d08684`.
- Active lifecycle, diverted-task audit, lifecycle consistency, Git whitespace and expected-only untracked Task 069 path checks: PASS.
- OVH clean distributed acquisition/install/restart/uninstall/reinstall/reboot/readiness/sterile rollback: PASS.
- Contabo real RC.2 -> RC.4 canonical update: FAIL, release-blocking before package replacement.
- Failure notification operation/transport separation: PASS (`operation FAILED`, `transport ACCEPTED`); mailbox receipt: NOT CLAIMED.
- Source diagnosis: `internal/update.PlanMigration` lacks RC.2 -> RC.4 and tests cover only transitions ending at RC.2. Classification: `PRODUCT/FRAMEWORK DEFECT`, finding `QWSG-069-F001`.
- Final `v1.2.0` metadata/tag/artifact/Forgejo Release/wget-curl gates: NOT EXECUTED after mandatory STOP.

## Rollback

Verify `/tmp/qwsg-task069-execution.DcQUOx/SHA256SUMS` and follow its `ROLLBACK.md`. Restore only explicit reviewed targets from isolated extraction; never use broad Git reset/checkout/restore/clean. OVH and Contabo require separate canonical QWSG host backups before their first mutation and state verification after any rollback.

Contabo is restored to the protected `/tmp/qwsg-task069-contabo-preupdate.I5SwQ1` RC.2 binary/unit/config identity with active stable Guardian. OVH is restored to QWSG-free sterile state. Both Task 069 remote acceptance directories were removed after local protected evidence retention.

## Completion state

`complete — BLOCKED: QWSG-069-F001 proves RC.4 lacks the mandatory RC.2 -> RC.4 deterministic migration path; no final release, tag or publication`

## Final decision

1. Decision: `BLOCKED`.
2. Failed gate: real Contabo canonical update from actual installed QWSG `1.2.0-rc.2` to immutable RC.4.
3. Exact evidence: update exit `1`, `Update refused: no deterministic compatible migration path`; source and tests omit the transition.
4. Candidate: QWSG `1.2.0-rc.4`, SHA-256 `adeb591605c0d37a5fc98d541125ca388cd4561703d0f0823bba931bc7d08684`.
5. Affected environment: Contabo production-like acceptance host; failure occurred before package replacement.
6. Host mutation: temporary QWSG notification configuration and acceptance files only; exact pre-update QWSG identity restored and task-created paths removed. OVH acceptance mutations were fully rolled back to sterile state.
7. Repository: no candidate/product repair and no final metadata/tag/release/publication; BLOCKED evidence commit `3f7af8361496340b3d9b602b19ccf1566a05c53e` was pushed cleanly.
8. Lifecycle: Task 069 prompt archived and closed truthfully as BLOCKED in canonical idle state.
9. Smallest remediation: add explicit compatible RC.2 -> new-candidate migration plus deterministic tests, freeze a new candidate, and repeat the complete mandatory final acceptance/release matrix.
