# Task History 095: Community 1.4.0 Real-System Clean Install Acceptance

## Task metadata

- Task ID: `095`
- Task slug: `community-1-4-0-real-system-clean-install-acceptance`
- Status: `active`
- Date generated: `2026-09-30` UTC
- Human authority: Owner-supplied Task 095 specification in the current session.
- Preferred owner communication language: Hungarian (validated project configuration).
- Related prompt: `ai/prompts/095_CURRENT_TASK.md`

## Starting state

Canonical main and fetched origin/main both c4dfe92c010195a7fb43162498815ec4748d64b4; divergence 0/0, clean worktree. Framework/job validated idle after complete archived 094. Git fetch required escalation because Git metadata is read-only in the sandbox; approved fetch succeeded. Required core reading and applicable skill/AGENTS read. Canonical task builder installed prompt/history transactionally; job validation PASS.

## Snapshot

Protected local baseline archive created before lifecycle mutation from git archive HEAD. SHA256 bf6dd250974f93fa9c60c1d7c5db07ee39acdeb766b58f6df65c2a2298ee6792. Archive listing/readability PASS. Retain through Owner acceptance and all rollback dependencies. Full payload and raw external evidence remain outside Git.

## Work performed

Phase 1 public release-authority reconfirmation PASS on 2026-09-30. Anonymous HTTPS downloads of canonical 1.4.0 artifact/index and historical 1.3.1 archive/sidecar succeed. Remote annotated tag 5396c8297007f92fe06a4587fe35796f6ba3abac peels to 996fb90d0f53d4ae6088cd884a654d9d0a32fff9. Historical tag efb7afcac7bd65a3419f6bf3264892e09de631f6 peels to unchanged 45009fe6169bff00842a4c4e9561bf339a5db81e.

Artifact URL: https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v1.4.0/qwsg-1.4.0-linux-amd64.tar.gz
Artifact: 3680796 bytes, SHA256 26cbf8d76fa27c652e11ee49f4570e05b3be6b969309e3cfd486977a27a6224d.
Canonical index URL: https://releases.quantumwizard.hu/qwsg/v1/release-index.json
Index: 1153 bytes, SHA256 fbb39b52e240ff3691090a48a1e7915e516da6fb2f6a1faa9ed5fc069b633c41; expected release media type.
Trusted retained verifier independently rehashed to d95604f9fb60a047300761fefa7dbbeded5b1ccb3afdb0386b3fb86f1cb3671d and verified canonical production signature. Checkpoint confirms production key qwsg-community-release-2026-01 and fingerprint 0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6. Exact stable/active release identity, source/platform/artifact authority and declared preserve-package-v1 migration metadata confirmed; migration not exercised.
Archive member safety and every MANIFEST.sha256 entry PASS. Embedded RELEASE.json: 1.4.0, exact source, linux-amd64, build 2026-09-25T17:25:47Z.
Historical public archive: 3600553 bytes, SHA256 0ab726bcde36182232ff89e3ea33e2d5ab77d8cad2b94ce56c9534a8147d9cd7; published sidecar matches. Forgejo historical asset list contains archive and sidecar only. A guessed historical signed-index companion URL returned 404; classified acceptance-path assumption, corrected by actual asset inventory, not product/release defect.
An initial archive safety assertion omitted the valid top-level directory member; corrected assertion then validated every member and manifest. No release bytes changed.
A redundant verifier build was stopped by sandbox cache permissions; existing independently hash-verified trusted frozen verifier used instead. No artifact rebuilt.

## Target preflight and continuation

Awaiting Owner identification of the designated OVH playground SSH host/alias and ordinary runtime user through the existing authorized access mechanism. No unambiguous target address was established from repository context. Do not infer that an existing acceptance key identifies the correct host. No remote target contacted or mutated. Next: read-only target inventory, supported-platform verification, exact minimal reset/install plan, Owner approval at destructive/credential boundary, then official public acquisition/authentication/install and existing real-system stability contracts.

## Verification

Repository baseline, snapshot, builder lifecycle and Phase 1 PASS. Phases 2-8 and acceptance items concerning target installation/operation remain unverified. No PASS claim for clean installation.

## Rollback

Local payload must be checksum/listing verified and extracted only into a separate directory; compare before exact task-owned restoration. No live-worktree extraction or broad reset. Target rollback/reset plan awaits inventory and explicit bounded authority.

## Safety and gates

No production VPS, unrelated infrastructure, historical release or signed metadata mutation. No 1.3.1 upgrade/bootstrap, rollback or successor work. Community 8 PASS / 0 PARTIAL / 1 MISSING; C9 remains MISSING. Clean-install portion is not yet proven. Remaining C9 requires separately authorized later historical 1.3.1 authenticated migration/state-preservation/rollback/recovery/Guardian/final health acceptance in addition to this clean-install proof.

## Completion state

PARTIAL; Phase 1 complete, awaiting explicit target identification. Prompt remains active for deterministic continuation. No closure/idle claim. Targeted evidence checkpoint commit/push is authorized; final Git identity is reported separately to avoid a self-referential commit hash.


## Owner-identified target and Phase 2 preflight (2026-09-30)

Owner explicitly identified the OVH playground and ordinary runtime account in this session. Private target details are retained in the Owner conversation and protected local plan; not published in Git. Repository resumed from 28c2df6205a1363b7d9fbbf72461acf791225902 with matching origin/main, 0/0 divergence and clean worktree. Existing Task 095 lifecycle preserved.

Initial strict hostname host-key verification refused because no hostname entry existed. DNS resolves to the already trusted IP entry; independently scanned ED25519 key exactly matches that existing entry. SSH uses HostKeyAlias for that existing identity with StrictHostKeyChecking=yes, without accepting a new key or modifying known_hosts. Connectivity and hostname verification PASS. No private key or credentials read, copied or printed.

Supported platform: Ubuntu 24.04.4 LTS, x86_64/amd64, kernel 6.8.0-138-generic, systemd 255 (255.4-1ubuntu8.17), ordinary UID 1000 runtime account, working systemd user bus, existing lingering, ext4 local filesystem, 34 GiB available on a 38 GiB root filesystem. Existing noninteractive sudo works; no privilege exceptions added.

Existing installed QWSG is official 1.3.1, source 45009fe6169bff00842a4c4e9561bf339a5db81e, build 2026-09-06T22:50:33Z. Installed binary is byte-identical to preserved authenticated historical archive binary (SHA256 278a80119013f8bfc1ec908a410b9905d747fce4c1d35356cc38605d5267ed1a). Target historical archive hash equals official 0ab726bcde36182232ff89e3ea33e2d5ab77d8cad2b94ce56c9534a8147d9cd7; all extracted package manifest entries PASS. Historical binary supports `version`; `--version` is refused as unknown, not a 1.4.0 finding.

Guardian user unit active/running/enabled since 2026-09-07, zero restarts, ExecMainStatus 0. Instantaneous memory approximately 5.3 MiB, peak 7.5 MiB, 9 tasks. Packaged unit has GOMEMLIMIT=64MiB, MemoryMax=128M, TasksMax=32, CPUQuota=25%, NoNewPrivileges and private state. No separate QWSG system timer/unit; Guardian owns internal scheduling. Unrelated cache-clean timer exists and is excluded from reset.

Active installation paths: /usr/local/bin/qwsg (root 0755), /usr/local/lib/systemd/user/qwsg-guardian.service (root 0644), /usr/local/share/doc/qwsg (root 0755). Private configuration and state are current-user 0700 directories; records 0600. Configuration model 1.0 contains locale/time-zone/extensions patch only; no credential file discovered. State includes ten snapshots, store metadata, operator state model 1.2, Guardian checkpoint and locks, scheduler state and temporary scheduler files, authenticated update-awareness evidence. Awareness reports historical current; this is not fresh 1.4.0 authentication evidence.

Private state uses approximately 74 MiB. Scheduler main JSON is 24873413 bytes, beyond current 8 MiB accepted bound, with several partial temporary files. Preserve all as historical evidence; do not load or claim valid current scheduler continuity. Current clean-install scope resets this legacy state after explicit approval, rather than attempting migration or rollback acceptance.

Additional inactive remnants: several 1.2.0 RC extracted packages/archives, 1.3.0 and 1.3.1 packages/archives/sidecars, a historical source tree, two root-private pre-upgrade backup directories. Preserve these unchanged: they are outside installed/PATH/configuration/state discovery paths. No additional QWSG /etc, /var/lib or /var/log runtime directory, per-user cache/share directory, or user-unit override found. Enabled user-unit symlink points exactly to packaged global user unit. Active PrivateTmp runtime directory belongs to systemd and must be retired by normal unit stop, never broad filesystem removal.

## Phase 3 proposed checkpoint (NOT EXECUTED)

Exact bounded proposed script is prepared in protected local task evidence and presented to Owner. It guards target hostname/account, exact installed binary hash, non-symlink fixed paths and absent backup destination. It stops only QWSG Guardian; requires inactivity and no QWSG process; creates a new root-private 0700 backup directory; archives the five active QWSG paths plus exact enablement symlink; hashes/verifies archive readability before reset; disables only QWSG unit; moves binary, packaged unit, entire QWSG docs directory, private configuration and entire private state into separately named backup objects; reloads only current user manager; verifies every runtime target absent, no PATH binary/process/active unit, and unchanged archive digest. No deletion of private state or backup payload.

Using preserved moves avoids the historical uninstaller's narrower removal list leaving older release documentation/trust files behind. All existing bytes and ownership remain recoverable, with full archive and relocated objects. Recovery is a separately approved exact-path restore after digest/readability validation and collision checks; restore saved unit/enablement/config/state/artifacts, reload/start only the prior unit and reverify. This is recovery planning, not rollback acceptance.

The resulting baseline has no installed package, enabled/running QWSG, discoverable prior configuration, state, inventory, scheduler, awareness, operator evidence or identity. Quarantined backups and explicitly preserved inactive release/source remnants cannot seed the new install. All unrelated services, packages, users, SSH, firewall, mail, database/network settings and existing lingering remain untouched.

PARTIAL at mandatory Owner destructive/reset checkpoint. No remote mutation, reset, stop, installation, upgrade or rollback performed. Await explicit approval of exact script before any such action. Task 095 remains active; C9 remains MISSING and Community totals unchanged.
