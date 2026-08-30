# Current Engineering Task 077: Update Awareness State and Operator Check/Status

## Task Metadata

- Task ID: `077`
- Task slug: `update-awareness-state-operator-check-status`
- Status: `complete`
- Date opened: `2026-08-30` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

Update Awareness State and Operator Check/Status


## Objective

Implement persistent local Update Awareness State and safe operator-facing update check/status behavior for QWSG 1.3. Consume Task 075 verified installed-package identity and Task 076 authenticated read-only release discovery, preserve awareness across restarts, distinguish freshness from historical success, and provide network-free status without creating a second updater or crossing the installation boundary.


## Scope

Implement one canonical versioned Update Awareness State model and private atomic store. Represent source/channel, verified installed identity, current/available relationship, last attempt, last successful authenticated check, authenticated source/release evidence, compatibility, safe failures, withdrawal and freshness. Integrate `qwsg update check` as one explicit bounded authenticated read-only refresh and `qwsg update status` as a strictly network-free local-state read. Preserve the last authenticated success across later failures. Reconcile stored identity with Task 075 where necessary. Reuse existing canonical state-root, ownership, permission, symlink rejection, strict decoding, integrity, advisory lock, temporary-file synchronization, atomic rename and directory-sync patterns. Do not duplicate Guardian checkpoint, operator-state, rollback, configuration, scheduler, notification storage, or the updater.


## Out of Scope

No Guardian periodic scheduling; notification generation, transition tracking, delivery or deduplication; automatic/unattended installation; artifact acquisition or package installation; second updater; migration/transaction/rollback redesign; release publication or signing-key custody; production service deployment; Pro registration, API keys, telemetry, fleet management, QUWIP or Telegram; external-host/infrastructure mutation; QWSG 1.2.0 tag/release/artifact mutation; unrelated refactor; Task 078 or later work.


## Authority Envelope

**Task targets and boundaries:** Inspect repository state persistence, command routing, Task 075 classification, Task 076 discovery, configuration, localization, update policy, updater boundaries and documentation. Add one canonical Update Awareness State model/store; modify update check/status routing strictly as required; consume Task 075/076 contracts; add deterministic tests/fixtures; make only narrow refactors preventing duplicate logic; update canonical documentation/history; complete normal Git integration and lifecycle closure. Work remains ordinary-user, local, bounded, reversible, privacy-minimized and read-only with respect to installation.

**Permitted external actions:** Read-only dependency retrieval required by existing tests when separately permitted; local loopback HTTP/TLS fixtures; normal origin fetch/push. No external release service, VPS, mailbox, Forgejo release or monitored host mutation.

**Owner-reserved decisions:** Production endpoint activation; production trust-anchor introduction/rotation/revocation/recovery; schema changes outside qwsg.update-awareness/1; Guardian scheduling; notification transition/deduplication; installation automation; Pro/telemetry/registration/fleet/QUWIP/Telegram; external infrastructure or release mutation; scope expansion; Task 078 start.

**Task-specific STOP conditions:** Stop for material baseline mismatch; any need to weaken Task 075 identity or Task 076 authentication; incompatible private-state redesign; updater replacement/duplication; production endpoint/key decision; external mutation, telemetry, automatic installation or notification scheduling; unprovable rollback ownership; or changed protected QWSG 1.2.0 evidence.


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

Expected repository `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, branch `main`, `HEAD == origin/main == aea9f73b7938d9715ae84e8f276d92ab6c040eb5`, divergence `0/0`, clean index/work tree except the Builder-installed Task 077 pair, and canonical idle before installation with Task 076 latest complete. Verify root, remote, lifecycle, framework, ownership/modes/ACLs, VERSION 1.2.0, Task 075/076 records, state-root/update commands, annotated v1.2.0 tag object `ac395b568b8e1f83c0ef85c9aa02f98c15402af0`, peel/release commit `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`, artifact size `3524214` and SHA-256 `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11` before production modification. Stop on unauthorized material mismatch.


## Snapshot Requirements

Retain and validate the Builder installation snapshot. Before production modification create a unique external mode-0700 execution snapshot under `/tmp` recording Git/branch/origin/divergence/index/work-tree/lifecycle/framework/version/protected-release identities, exact target before-images or absence, ownership/modes/ACLs, checksums, complete readable Git bundle and literal bounded rollback. Evidence files are mode 0600 and exclude credentials/private host data/runtime secrets/caches. Validate checksums, bundle, readability and rollback before change. Create a separately verified closure snapshot before prompt archival.


## Risk Assessment

Stale data could appear current; failed checks could erase authenticated success; 304 could be trusted without authenticated cached data; corrupt state could be accepted; installed identity could change after update/rollback; withdrawal could remain actionable or erase useful history; unavailability could be mistaken for no update; state/output could leak sensitive data; unsafe ownership/modes/links/filesystem could weaken storage; concurrency could publish partial state; check could cross acquisition/installation; CLI could bypass/duplicate updater; awareness could affect Guardian health/readiness; ambiguous transitions could weaken later notifications. Mitigate with strict bounded versioned state, integrity, private storage, locking, atomic publication, last-success preservation, explicit timestamps/freshness, safe reason tokens, identity reconciliation and subsystem boundaries.


## Planned Work

1. Verify baseline and lifecycle.
2. Create and validate execution snapshot.
3. Inventory update commands, state-root/private stores, locking/integrity/localization, Task 075 and Task 076.
4. Define qwsg.update-awareness/1, evidence precedence, freshness, failures, withdrawal and identity reconciliation.
5. Implement one bounded canonical state model and private atomic integrity-conscious store.
6. Define deterministic successful/failed/unavailable/withdrawn/current/newer/unsupported/identity-changed transitions.
7. Integrate explicit read-only `qwsg update check` and network-free `qwsg update status`.
8. Preserve last authenticated success across failed attempts; never download/install.
9. Add deterministic model/store/CLI/integration/localization/failure-injection tests.
10. Prove updater/migration/transaction/rollback/Guardian/readiness/notification remain unchanged.
11. Synchronize documentation/history and run complete validation/security/privacy review.
12. Create closure snapshot, archive prompt, close lifecycle, synchronize origin and stop for Owner review.


## Rollback Plan

Use only verified Task 077 snapshots and exact target evidence. Restore recorded before-images and modes/ACLs; remove only Task 077-created paths whose prior absence and current identity are proven; preserve later Owner edits. Never broad reset/checkout/restore/clean, wildcard-delete, rewrite history/refs, mutate tags/artifacts or touch external systems. Re-run focused/full/framework/lifecycle/Git/reproducibility/protected-release checks after rollback. Stop if ownership cannot be proven.


## Deliverables

Canonical qwsg.update-awareness/1 model; private atomic integrity-conscious store; documented state/freshness/last-attempt/last-success/withdrawal/unavailability semantics; Task 075 identity reconciliation; Task 076 authenticated-result consumption; safe explicit update check; network-free update status; deterministic unit/integration/CLI tests; synchronized architecture/operator/security/roadmap/engineering/task documentation; exact completion, snapshot and rollback evidence.


## Verification

Run active/idle lifecycle and Framework validation; Builder/diversion/lifecycle assertions; gofmt; go vet ./...; go test ./...; go test -race ./...; focused Task 075/076/077 tests; deterministic build/reproducibility; git diff --check; staged scope/privacy/secret/mode/ACL review; final cleanliness/origin convergence. Test never-checked, current, supported update, locally unsupported newer release, failed first/later checks preserving success, stale status, authenticated 304 cache handling, withdrawal, source unavailable, corrupt state, installed identity change, concurrency, pre-rename preservation, post-rename recovery, no-network status, no acquisition/install check, and no Guardian/readiness/notification mutation. Prove protected QWSG 1.2.0 evidence unchanged and no external VPS required.


## Documentation Updates

Update the release-awareness and release-index contracts, native update/rollback architecture, architecture/system map, security/privacy, operator check/status guidance, roadmap, engineering history and Task 077 history. Document schema, evidence/integrity, freshness/last-success, withdrawal/source failure, installed-identity reconciliation, check/status network boundaries, corruption/recovery, future Guardian/notification relationship, and separation of local identity, remote authenticity, artifact integrity, migration compatibility, awareness and installation.


## Completion Criteria

Complete only when one canonical bounded strict integrity-conscious atomic awareness model/store exists; last attempt and authenticated success are distinct; failures preserve valid success; stale/unavailable/withdrawn/identity-change states are truthful; check consumes Task 076 and remains read-only; status is network-free; neither downloads/installs; no second updater, Guardian schedule, notification logic, telemetry or Pro feature exists; all validation passes; snapshots/rollback are valid; repository is clean/synchronized/canonical idle; QWSG 1.2.0 evidence is unchanged; and the Owner report recommends but does not start the next task.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-30 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
