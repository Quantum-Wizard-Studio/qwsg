# Current Engineering Task 081: QWSG Community 1.3.0 Release & Production Acceptance

## Task Metadata

- Task ID: `081`
- Task slug: `community-1-3-0-release-production-acceptance`
- Status: `complete`
- Date opened: `2026-09-06` UTC
- Human authority: Project Owner explicit Task 081 authorization and APPROVE in the current session
- Owner or lead-developer communication language: Hungarian

## Title

QWSG Community 1.3.0 Release & Production Acceptance


## Objective

Produce, publish and production-accept QWSG Community 1.3.0 through the complete authenticated operator-controlled update lifecycle: installed 1.2.0 discovers real official 1.3.0, delivers and deduplicates notification, explicitly upgrades, then authenticates current/equal 1.3.0. No unattended installation.


## Scope

Accepted Tasks 074–080 release/update architecture and Scheduler remediation; canonical version metadata, deterministic Linux AMD64 packaging, supported clean installation, 1.2.0 to 1.3.0 migration and rollback, EN/HU documentation, full release gates, canonical Forgejo commit/tag/Release/assets, offline-signed production index, real pre/post-install production acceptance, evidence and lifecycle closure.


## Out of Scope

Task 082 and Pro work; automatic artifact download from awareness or automatic installation; unrelated features, Release Center UI, server.quantumwizard.hu TLS, SMTP repair, Hestia maintenance, DNS redesign, GitHub release redesign or infrastructure tuning. Preserve official 1.2.0 artifact unchanged. Never handle private signing keys or passphrases.


## Authority Envelope

**Task targets and boundaries:** Complete Owner-authorized Task 081 release and production acceptance only; standard autonomous engineering authority applies.
**Permitted external actions:** Canonical Forgejo push, annotated v1.3.0 tag, Release and immutable assets; authenticated official production index retrieval and atomic publication; configured real update notification acceptance; explicit supported production 1.2.0 to 1.3.0 update and bounded health observation. GitHub remains read-only mirror/storefront.
**Owner-reserved decisions:** Offline production signing on Dell1 and interactive sudo credentials; only material new scope/architecture, out-of-plan destructive operations, security incidents or inability to satisfy the accepted release contract require a new decision. No routine approval or continuation gates.
**Task-specific STOP conditions:** Pause only for exact final offline signing input or technically unavoidable interactive root commands; resume automatically after Owner completion. Fail closed on release gates, provenance or authenticity inconsistency; diagnose and correct in-scope failures.


## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/core/17_EXECUTION_MODEL.md`
- `ai/core/18_BOUNDED_DIAGNOSTIC_RUNNER.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

Verify clean main, fetched HEAD equals origin/main and divergence 0/0; expected closure 7a9cd0ab09275b01a7c7a945e5781b95ce2a9126 and idle Task 080 lifecycle. Verify protected official 1.2.0 artifact SHA-256 44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11, source 348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2, official index authentication and production Guardian health before product modification.


## Snapshot Requirements

Create and verify protected canonical pre-change Git/source snapshot with hashes, archive readability, baseline and bounded restore procedure. Capture exact production service/state and rollback readiness before privileged mutation; retain through acceptance and rollback window.


## Risk Assessment

High: provenance mismatch, signing/publication errors, state loss and production resource regression. Mitigate with frozen deterministic artifacts, authenticated exact metadata, atomic rollback-capable publication, isolated migration/rollback proof, private evidence and bounded real acceptance. Never expose credentials.


## Planned Work

Establish baseline and snapshot; determine accepted release content; align version/provenance; prove clean install, migration, legacy oversized Scheduler safety and rollback; document EN/HU release notes; run all release gates and reproducibility; finalize source commit and deterministic artifact; publish canonical Forgejo v1.3.0; generate exact index signing bytes; obtain offline Owner signature; verify and atomically publish signed index; prove real 1.2.0 newer awareness plus notification/deduplication without download/install; explicitly update production; prove current/equal 1.3.0 and bounded Guardian health; archive evidence and close idle without Task 082.


## Rollback Plan

Verify snapshot checksums before isolated restoration; restore only reviewed exact task paths, never extract over live worktree or rewrite published history. Preserve official 1.2.0 assets. Prove deterministic isolated installer rollback; capture exact production rollback inputs before update; execute production rollback only for actual failed installation/acceptance. Index publication must preserve previous signed bytes and anti-rollback/future-index contract.


## Deliverables

qwsg-1.3.0-linux-amd64.tar.gz, derived exact size/SHA-256 and reproducibility evidence; consistent source/binary/RELEASE/package/index/Forgejo/tag identity; EN/HU notes; authenticated production stable 1.3.0 index with key qwsg-community-release-2026-01 and trust fingerprint 0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6; real discovery/notification/migration/health/authority acceptance and complete history.


## Verification

All mandatory gates: format, vet, all Go tests, full race, engineering/framework/lifecycle, release and release-authority reproducibility, installer/clean-install/1.2.0-to-1.3.0 migration/rollback, privacy/security, protected 1.2.0 hash, Scheduler/resource, awareness, notification/deduplication. Prove oversized legacy Scheduler cannot cause OOM/restart loop. Before update use real signed published index to prove installed 1.2.0 newer/available, successful notification and persistent deduplication across repeated checks, healthy Guardian and no automatic download/install. After explicit migration prove exact 1.3.0 identity/current-equal, preserved config/credentials/state/integration, zero-network local status, no telemetry/registration/installation ID/API key/listeners/leakage, bounded scheduler, active/running Result=success no restart/OOM and safe memory/tasks. Verify full public bytes/signature/key/fingerprint/artifact/source/tag/installed chain; HTTPS redirect, dedicated TLS, media type application/vnd.quantumwizard.qwsg-releases+json, Cache-Control no-cache, ETag/Last-Modified/304 and no unsafe Expires.


## Documentation Updates

Update canonical version/release/migration metadata, concise EN/HU operator notes covering authenticated awareness, automatic checks, notification/deduplication, operator-only installation, privacy/security and Scheduler resource fixes. Record exact commits/tag object/Release ID and URL/artifact and index hashes, gates, rollback, production acceptance, snapshots and final Git/lifecycle state in Task 081 history and milestone index.


## Completion Criteria

Complete only when all release gates pass, official Forgejo v1.3.0 and signed production index are published, real installed 1.2.0 discovers newer 1.3.0 and notification/deduplication passes, explicit migration and post-install authority/resource acceptance pass, documentation/evidence archived, closure pushed, HEAD equals origin/main divergence 0/0, clean worktree and idle lifecycle. Final report includes all exact release identities and acceptance results. Do not start Task 082.


## Owner Approval Requirements

Approved by Project Owner explicit Task 081 authorization and APPROVE in the current session through the Engineering Task Builder on 2026-09-06 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.

## Owner-authorized compatibility prerequisite

The Project Owner explicitly approved the bounded installed-1.2.0 migration
compatibility backport in the current session. Add only the explicit
1.2.0 to 1.3.0 route, preserve authentication/classification/provenance and all
configuration/state, test supported and rejected paths and notification
eligibility, capture exact production rollback before deployment, and maintain
matching installed binary/RELEASE identity. Never mutate the historical official
1.2.0 artifact/tag/Release. Document the remediated installed baseline separately.
Interactive sudo remains Owner-executed; after deployment verification continue
Task 081 automatically. Canonical 1.3.0 supersedes this temporary baseline.

## Owner-authorized 1.3.1 corrective release

The Owner explicitly approved QWSG Community 1.3.1 within Task 081 after
post-install F001/F002. Status remains IN PROGRESS until full corrected
production acceptance. Scope is only stale 304 installed/available relation,
notification deduplication preservation across version changes, directly needed
regressions, documentation, deterministic packaging/provenance, migration and
authenticated release publication/acceptance. Production installed 1.3.0 is the
real source. Immutable 1.3.0 artifact, tag, Release metadata and history must
not be rewritten. No Task 082 or unrelated feature work. Existing authority
continues; only offline Dell1 signing, interactive sudo or genuinely new
architectural/security decisions pause execution. Final completion uses the
canonical 1.3.1 corrective release with explicit historical 1.3.0 failure record.

## Final outcome

Complete with Owner-authorized corrective 1.3.1 production acceptance.
Historical 1.3.0 remains immutable; final evidence is in Task 081 history.
No Task 082 created.
