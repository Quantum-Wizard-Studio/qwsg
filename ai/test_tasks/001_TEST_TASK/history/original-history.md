# Task History 015: Approved Task

## Task metadata

- Task ID: `015`
- Task slug: `forgejo-installation`
- Status: `active — production snapshot verified`
- Date generated: `2026-07-22` UTC
- Human authority: Attila — Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/015_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded. Execution started after reconciling the approved release-tag baseline.

## Starting state

The verified repository root is on `main` at `b4316050436bf8be4062f0e1d4ba7c371c334223`. The Project Owner explicitly authorized creation of the missing annotated `v0.1.0` tag on that exact commit. A verified pre-tag Git bundle was created outside the repository before the tag was added. The resulting tag is an annotated tag object, points exactly to the approved commit, and has the approved message. No remote is configured. Existing untracked engineering records and snapshots were preserved. The production baseline confirms Ubuntu 24.04 on x86_64, HestiaCP 1.9.7, Nginx and Apache, MariaDB, operational HTTP/HTTPS, no Forgejo target paths, and no listener on TCP port 3000.

## Snapshot

The valid production engineering snapshot is `/tmp/qwsg-task015-20260722T022100Z`. It contains the Git bundle and ref/status evidence, target-domain Hestia configuration, matching web configuration evidence, systemd inventory, MariaDB database/user metadata without authentication data, firewall and listener state, target filesystem state, public TLS certificate metadata, DNS and HTTP(S) baselines, relevant versions, scope, retention, and bounded restore documentation. All `SHA256SUMS` entries passed, the Git bundle verified as complete, and every archive passed a listing/readability test. A targeted filename and content scan found no private key, password, token, API key, or credential pattern. The partial failed snapshot at `/tmp/qwsg-task015-20260722T021500Z` remains untouched and is not valid rollback evidence.

## Work performed

Reconciled and verified the approved annotated release tag. Created and independently verified the complete production snapshot before installation changes. The owner-executed Phase 2–5 installer created the native service, database, filesystem layout, and initial configuration, but the first systemd start entered a restart loop. Forgejo reported that it could not save a missing JWT secret to `/etc/forgejo/app.ini` because the systemd `ProtectSystem=full` sandbox correctly made `/etc` read-only. Investigation against the Forgejo 15.0.5 release configuration found that the installer had placed `JWT_SECRET` under `[lfs]`; Forgejo 15.0.5 requires `LFS_JWT_SECRET` or `LFS_JWT_SECRET_URI` under `[server]`. The first scoped repair correctly removed the ineffective key and installed the server LFS secret URI, but it did not produce a successful startup and must not be treated as successful. The subsequent full journal identifies the next exact fatal gate at `setting/oauth2.go:149`: effective `oauth2.JWT_SECRET` is absent and Forgejo attempts to persist it to the read-only custom configuration. The prior repair incorrectly conditioned OAuth2 secret generation on an explicitly configured HS signing algorithm. The owner-provided doctor attempt ran as root and was rejected before configuration diagnosis, as expected; the next validation must run as the `forgejo` service identity. The systemd sandbox and `/etc/forgejo` directory permissions remain intentionally unchanged. A second, narrowly scoped and idempotent repair was prepared at `/tmp/qwsg-task015-repair-oauth2-jwt.sh` to install only the protected OAuth2 JWT file reference, validate as `forgejo`, restart, and verify localhost-only binding.

## Verification

Builder input, metadata, prompt/history identity, approval state, and lifecycle installation validated successfully. Pre-tag bundle verification, release-tag object and target verification, production snapshot SHA-256 verification, bundle verification, archive readability, secret scan, baseline listener inspection, and target-path absence checks passed. Forgejo 15.0.5 binary signature and version verification passed before installation. Initial service startup and the first repair attempt both failed at explicit configuration secret-persistence gates; no successful service or repair state is claimed. The second repair script passed Bash syntax validation and a review confirming that it does not print secret values. OAuth2 repair execution, service validation, and Forgejo-identity doctor validation remain pending owner-side root execution.

## Rollback

Verify the production snapshot checksums before restore. Restore only exact Task 015 targets after impact review. Never extract the repository bundle over the live worktree; clone it separately and reconcile refs. Remove only identities, database objects, directories, service files, binary, proxy integration, and remote configuration proven to have been created exclusively by Task 015. Validate Hestia, Nginx, Apache, MariaDB, mail, SSH, firewall, hosted domains, and Git state after rollback. Private keys are deliberately excluded from the snapshot.

## Completion state

`active — OAuth2 JWT startup gate diagnosed; second scoped repair pending`
