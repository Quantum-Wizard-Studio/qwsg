# Task History 074: Update Discovery and Release Awareness Architecture

## Task metadata

- Task ID: `074`
- Task slug: `update-discovery-release-awareness-architecture`
- Status: `complete — architecture approved for Owner review; no implementation`
- Date generated: `2026-08-29` UTC
- Human authority: Attila — Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/archive_prompts/074_2026-08-29_update-discovery-release-awareness-architecture.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured Owner input. Exact `APPROVE` was received, the active pair passed lifecycle validation, and documentation-only execution began under the approved Framework 2.0 envelope.

After completion validation, the reviewed documentation commit was pushed by
clean fast-forward and the complete prompt was archived without creating a
successor. The repository returned to canonical idle for Owner review.

## Starting state

- Pre-Builder root `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, branch `main`, clean tree/index, `HEAD == origin/main == ee286088e856f3c6a3341ee9eb36f74173683757`, divergence `0 0`, canonical idle after Task 073.
- A fresh post-Builder fetch reconfirmed the same HEAD/upstream relation. Only the expected Task 074 prompt/history were initially untracked.
- `VERSION` and tagged `VERSION` are `1.2.0`. Annotated tag object `ac395b568b8e1f83c0ef85c9aa02f98c15402af0` peels exactly to release commit `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`.
- Frozen `dist/qwsg-1.2.0-linux-amd64.tar.gz` is `3524214` bytes and matched SHA-256 `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11` before documentation edits.

## Snapshot

- Builder snapshot `/tmp/qwsg-task074-builder-install.AvAmIG`: mode `0700`, files `0600`; clean Git/idle identity, Task 074 absence, Task 073 reference payloads, complete tracked-HEAD bundle, checksums and bounded rollback; all checks passed.
- Execution snapshot `/tmp/qwsg-task074-execution.oJDIxu`: mode `0700`, files `0600`; exact target before-images/new-file absence, Git/upstream/tag/artifact identities, ownership/modes/ACLs, complete tracked-HEAD bundle, checksums and literal bounded rollback; all checks passed.
- Both snapshots exclude credentials, runtime data, caches and unrelated payloads and are retained through Owner review.
- Closure snapshot `/tmp/qwsg-task074-closure.E0UUJR`: mode `0700`, files `0600`; exact committed prompt/history, archive-destination absence, synchronized commit identity, complete tracked bundle, checksums and bounded rollback; all checks passed before prompt archival.

## Work performed

- Read the active prompt, `qwsg-job` skill, all required governance/configuration documents and relevant architecture, release and Task 064/068/070/073 evidence.
- Inventoried actual QWSG 1.2.0 version/provenance, bounded Forgejo API discovery, strict version comparison, acquisition and layered package verification, declared migrations, fixed-destination transaction/rollback, update CLI/policy, Guardian recurrence/checkpoint, Notification/SMTP, readiness, configuration, localization and security/privacy boundaries.
- Established that 1.2.0 already has direct Forgejo discovery and explicit installation; the design extends it instead of adding a parallel updater.
- Added `docs/architecture/UPDATE_DISCOVERY_AND_RELEASE_AWARENESS.md`: current-state reuse/gaps; non-coupling and installed-identity boundaries; source abstraction; static/Forgejo/future-service evaluation; proposed `qwsg.release-index/1`; authenticity/channel/version/compatibility contract; private awareness state; CLI, Guardian, notification, failure and privacy models; Community/Pro boundary; and staged implementation recommendations.
- Incorporated all three Owner-verified field findings without fixes: legacy pre-alpha path misclassification, opaque SMTP sender-verification failure, and truthful external-delivery unknown/partial readiness.
- Updated Architecture Governance, System Map, Product Architecture, Roadmap and concise Engineering History.

## Architecture decisions

1. Extend `internal/update`; keep installation explicit and reuse its package, migration, transaction and rollback boundaries.
2. Use a strict source-neutral release-index rather than HTML scraping. Prefer a static credential-free Community manifest; retain Forgejo API as an adapter and permit a future Release Service adapter.
3. Separate metadata authenticity from artifact integrity. Recommend embedded-key Ed25519 manifest verification; retain HTTPS/host constraints and SHA-256/internal package verification as distinct layers.
4. Keep local declared migration planning authoritative over manifest compatibility hints.
5. Store small integrity-protected Update Awareness State separately from Guardian, rollback, Inventory, Current Operator State, Health and readiness state.
6. Check on an isolated recommended 24-hour schedule with jitter and bounded resources; failure never makes Guardian unhealthy.
7. Notify once per durable actionable transition identity through existing localized provider delivery; process-memory-only deduplication is insufficient.
8. Keep Community credential-free, registration-free and telemetry-free while accurately disclosing ordinary HTTPS source-IP/network exposure. Central/Pro features remain opt-in future scope.

## Verification

Builder input, metadata, prompt/history identity, approval state, lifecycle installation, fresh origin convergence, annotated-tag peel, frozen artifact checksum and both snapshot integrity gates validated successfully.

- `make engineering-test test vet fmt-check`: PASS, including deterministic build-contract checks, 25 Framework, 36 diversion, 29 lifecycle, 49 Builder assertions, all Go packages, vet and complete Go formatting.
- Framework v2 semantic assertions: 15 PASS; bounded diagnostic runner assertions: 11 PASS; active lifecycle and diverted-task audit: PASS.
- `go test -race ./...`: PASS for every Go package.
- `scripts/test-release-reproducibility.sh`: PASS across source-mode and umask variations.
- `make release-check`: expected precondition refusal before construction because the published final artifact already exists in `dist/`; classified `EXPECTED BEHAVIOR`. Its constituent reproducibility test passed separately. Task 074 neither removes nor overwrites the frozen release merely to satisfy a clean-output release-construction precondition.
- `git diff --check`, documentation key-section/reference checks, target ownership/mode review, production-source exclusion, snapshot checksum/bundle checks and protected tag/artifact identity checks: PASS.
- The first completion-time snapshot checksum command ran from the repository rather than each snapshot root and therefore reported relative members missing. Classified `TEST OR ACCEPTANCE DEFECT`; rerunning from both recorded roots verified every member and both complete bundles without changing repository content.
- No `cmd`, `internal`, packaging, release script, `VERSION`, Makefile, runtime, infrastructure, release, tag or published artifact behavior/content was modified. No SMTP, monitored server, central service or other external runtime was contacted.
- Reviewed documentation commit `3f7aeff5ff8eaa3a7966471d9b78394a7e585d5f` (`docs: design QWSG 1.3 update discovery architecture`) passed targeted staging and push dry-run, then clean-fast-forwarded canonical `main`. The lifecycle-closure and final synchronized identities are reported in the Owner handoff because a commit cannot truthfully contain its own final identity.

## Implementation gaps and recommended sequence

Missing production capabilities are evidence-based installation classification; strict release-index parser and authority verification; source adapter; persistent awareness store and network-free awareness status; release channel configuration; Guardian low-frequency scheduling/failure isolation; durable transition notification/localization; and explicit withdrawal handling. Later, separately approved tasks should implement them in that order, followed by safe SMTP diagnostic categories/external-delivery evidence as independent concerns and full clean-host/privacy/outage/update/rollback acceptance. No later task was started.

## Rollback

Verify the applicable snapshot `SHA256SUMS`, repository identity, exact target identities and absence of later Owner edits. Restore only named pre-existing payloads with recorded metadata; remove the new architecture document only after proving its recorded absence and current Task 074 identity. Never reset, clean, checkout, broadly restore, rewrite history or mutate a tag. Re-run lifecycle, framework, documentation, Git and protected-release checks.

## Completion state

`complete — architecture approved for Owner review; no implementation; v1.2.0 release identities untouched`
