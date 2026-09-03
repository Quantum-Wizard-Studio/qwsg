# Current Engineering Task 080: Update Notification & Deduplication

## Task Metadata

- Task ID: `080`
- Task slug: `update-notification-deduplication`
- Status: `approved`
- Date opened: `2026-09-03` UTC
- Human authority: Project Owner — explicitly authorized Task 080 and supplied the exact engineering scope with APPROVE on 2026-09-03 UTC.
- Owner or lead-developer communication language: Hungarian

## Title

Update Notification & Deduplication


## Objective

Extend QWSG Community so an authenticated newer applicable stable release discovered by the accepted release-awareness system can produce a practical operator notification through the existing notification subsystem, while persistently deduplicating successful delivery across checks and restarts. Delivery failures remain retryable only on bounded Guardian scheduling, and no automatic artifact download, staging, installation, version mutation, or restart is introduced. Completion requires implementation, deterministic isolated acceptance, current-production acceptance, all mandated validation, documentation, archived evidence, pushed commits, clean synchronized Git state, and canonical idle lifecycle.


## Scope

- Integrate update-available notification into the Task 079 Guardian automatic release-check path only after the existing Task 076–079 authentication, trust, validation, applicability, and newer-version decisions succeed.
- Reuse the existing QWSG notification subsystem and existing SMTP/notification configuration without enabling or mutating it.
- Add persistent deterministic deduplication based only on the smallest robust authenticated release identity, written only after established successful-delivery semantics.
- Make failures non-fatal and retryable on the existing bounded Guardian schedule, without tight loops, overlap, busy polling, or a new general queue.
- Preserve manual `qwsg update check` as a non-notifying check and `qwsg update status` as zero-network and zero-notification.
- Add deterministic automated and isolated acceptance coverage for current/older/invalid/newer releases, first/same/restart/different release behavior, delivery success/failure/retry bounding, disabled delivery, concurrency, corruption, cleanup, and excluded side effects.
- Update architecture, operator EN/HU, security/privacy, changelog/system records, and task history.
- Perform read-only production acceptance against the official endpoint for installed/current 1.2.0, targeted Git integration, clean fast-forward push, and canonical idle closure.


## Out of Scope

- Automatic update installation; artifact download, acquisition, staging, or installer execution; privileged update operations; installed-version or system-package mutation; update-triggered restart.
- Community 1.3 publication, fake production release publication, tags, Releases, public artifact changes, Task 081, Pro functionality, API keys, telemetry, registration, installation identifiers, new identifying outbound headers, inbound listeners, or Release Center UI.
- SMTP credential workflow repair, automatic SMTP/email enablement, production SMTP credential mutation, or a new general-purpose notification queue.
- Server TLS remediation, unrelated infrastructure maintenance, or unrelated architecture/configuration redesign.


## Authority Envelope

1. **Task targets and boundaries:** Implement and document only authenticated newer-stable-release operator notification and persistent successful-delivery deduplication through the accepted Task 076–079 and existing notification architecture, with the objective, scope, exclusions, tests, acceptance, Git integration, and lifecycle closure defined here. Preserve Community 1.2.0 behavior and do not begin Task 081.
2. **Permitted external actions:** Read-only Git fetch and push verification plus clean fast-forward task commits to the canonical repository; outbound credential-free HTTPS retrieval of only the authorized official production release index for bounded current-release acceptance; read-only Guardian/service health inspection. Existing notification transport may be exercised only with deterministic isolated test doubles/fixtures. No real production notification delivery or production configuration/credential mutation is authorized.
3. **Owner-reserved decisions:** Any material architecture change, scope expansion, privileged or destructive production mutation, real external notification test, security-policy exception, publication/tag/release action, or resolution of a genuine security issue requiring product policy.
4. **Task-specific STOP conditions:** Stop only if material architecture or scope expansion is required; privileged/destructive production mutation is required; real external notification delivery becomes necessary; a genuine security issue requires Owner decision; rollback cannot be established; or root authentication is the sole blocker to a bounded acceptance operation, in which case return the exact minimum Owner command/script.


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

- Verify repository root, Framework identity, lifecycle, Task 080 identity, branch, canonical remote, fetched `origin/main`, HEAD, divergence, status, permissions/ACLs, relevant service health, installed version, notification configuration boundaries, production trust identity, and Task 079 closure before implementation target changes.
- Expected baseline: canonical idle after complete Task 079; clean `main`; `HEAD == origin/main == 1b3fb747e91a3f01256c22250e31fe3f99e242fd`; divergence `0/0`; production endpoint `https://releases.quantumwizard.hu/qwsg/v1/release-index.json`; installed/current stable `1.2.0`; Task 079 implementation `c673318070cac0fc2607c5ed2042cd471d0e067a`; closure `1b3fb747e91a3f01256c22250e31fe3f99e242fd`.
- Inspect Task 046/068 notification contracts, Task 074–079 release-awareness contracts, Guardian scheduling/composition, persistence/security conventions, localization, and relevant tests before design decisions.
- Diagnose recoverable in-scope variance. Stop only at a genuine authority, identity, security, privacy, destructive, external-mutation, or rollback boundary.


## Snapshot Requirements

- Retain and verify the pre-intake snapshot `/tmp/qwsg-task080-intake.kxiFwc`, mode 0700, containing a mode-0600 complete Git bundle at baseline `1b3fb747e91a3f01256c22250e31fe3f99e242fd` and a readable mode-0600 lifecycle archive. Recorded SHA-256 values are `c1c7800022979817f3c60fccd6f03505a8542ba1d20fdd90e3c960c9b3fdee83` and `8f2c7e0267dc102ae5eaaf4ea23d84b0fff6a9861f1656971cec694ce2393402`.
- After Builder installation and before implementation target changes, create a second timestamped mode-0700 execution snapshot under `/tmp` containing a complete verified Git bundle, exact prompt/history before-images, relevant implementation/documentation targets, permissions/ACL evidence, manifest hashes, and bounded rollback instructions. Keep payloads outside Git and retain snapshots through Owner acceptance.


## Risk Assessment

- High: unauthenticated, invalid, inapplicable, current, or older metadata could trigger a misleading notification. Mitigate by consuming only the accepted authenticated evaluation result and explicit eligibility tests.
- High: deduplication could suppress an undelivered notice or spam repeated checks/restarts. Mitigate with post-success atomic persistence, authenticated identity, bounded schedule semantics, restart/failure/concurrency tests, and fail-safe state handling.
- High: notification work could accidentally cross into download or installation. Mitigate with dependency boundaries and explicit no-artifact/no-installer assertions.
- Medium: delivery failure could crash, overlap, stall, or exhaust Guardian resources. Mitigate with sequential ownership, cancellation/time bounds, error isolation, race/resource tests, and no immediate retry loop.
- Medium: schema, permissions, corruption, configuration, or localization compatibility could regress. Mitigate by extending the smallest existing private versioned state contract where sound, atomic persistence, permission/corruption tests, disabled-channel tests, and synchronized EN/HU updates.
- Rollback risk is low while work remains source/history-only, with verified snapshots and exact-path rollback retained.


## Planned Work

1. Validate governance, Git/lifecycle identity, snapshots, service/configuration boundaries, and read required architecture/security/notification/release Task 046/068/074–079 records.
2. Inspect existing release evaluation, awareness persistence, Guardian scheduling, and notification delivery contracts; choose the smallest authenticated deterministic release identity and state integration.
3. Implement Guardian-only notification eligibility after successful authenticated newer stable discovery, concise localized content, disabled-channel behavior, sequential non-overlap, failure isolation, and post-success atomic deduplication.
4. Keep manual check non-notifying, status local-only/non-notifying, and all artifact/install/telemetry/listener paths absent.
5. Add deterministic unit/integration and isolated acceptance coverage for every specified eligibility, deduplication, restart, different-release, delivery failure/retry bound, disabled, corruption, concurrency, cleanup, and side-effect case.
6. Update architecture, security/privacy, operations/update operator EN/HU documentation, README/changelog/system records, and Task 080 history.
7. Run focused, full, race, engineering/framework/lifecycle, reproducibility, release-authority, protected-artifact, security/privacy, resource, and rollback checks; diagnose/correct/retest in scope.
8. Run bounded production acceptance against the official endpoint proving authenticated current 1.2.0, no notification/dedup record, zero-network status, healthy Guardian, and no artifact/install/configuration/credential mutation.
9. Review exact diffs, stage only explicit task paths, commit/push implementation, finalize and archive Task 080, push closure, and verify synchronized clean idle state without starting Task 081.


## Rollback Plan

- Before implementation, enumerate exact Task 080 target paths in the execution snapshot. Restore only those paths from the verified target archive or baseline Git bundle after confirming current identities and preserving unrelated work.
- For uncommitted source/document changes, apply exact before-images only to enumerated Task 080 targets and remove only Task 080-created paths after confirming recorded prior absence; rerun formatting, focused/full tests, Framework/lifecycle checks, and Git status.
- For lifecycle intake rollback before implementation commit, remove only Builder-created `ai/prompts/080_CURRENT_TASK.md` and its exact matching Task 080 history after collision/hash verification, restoring canonical idle validation.
- After a pushed commit, use a new exact-path revert commit rather than rewriting published history. Never use broad reset, checkout, clean, wildcard removal, or extraction over the live worktree. No production state should require rollback because acceptance is read-only and isolated.


## Deliverables

- Production-ready Guardian update-available notification after authenticated newer applicable stable discovery, using the existing notification subsystem and concise EN/HU content.
- Persistent deterministic restart-safe successful-delivery deduplication, bounded retry after failures, disabled-channel safety, non-overlap, and Guardian failure isolation.
- Preserved manual check and local-only status semantics, with no artifact acquisition or automatic installation path.
- Deterministic automated and isolated newer-release/delivery acceptance evidence covering all required cases.
- Bounded official current-release production acceptance with no notification, false dedup state, artifact/install action, or production configuration/credential mutation.
- Updated engineering/operator/security documentation and complete Task 080 history/archive evidence.
- Implementation and closure commits pushed with synchronized clean `main` and canonical idle lifecycle.


## Verification

- `bin/job --check`, Framework validation, lifecycle consistency, task/history identity, both snapshot integrity checks, and rollback validation pass.
- Formatting, vet, all Go tests, `go test -race ./...`, engineering/framework/lifecycle tests, release reproducibility, release-authority reproducibility, protected artifact verification, and relevant security/privacy/resource checks pass.
- Automated evidence proves current/older/invalid/untrusted produce no notification; authenticated newer produces one; same release and restart are deduplicated; a distinct authenticated newer identity is eligible; success persists state; failure does not; retry is schedule-bounded; disabled delivery does not attempt; Guardian survives; attempts do not overlap; corruption is safe.
- Automated evidence proves manual `update check` does not notify, `update status` performs zero network and zero notification, and no artifact download, staging, installation, telemetry, registration, listener, credential persistence, or identifying header exists.
- Production acceptance authenticates the exact official endpoint and current 1.2.0, proves no update notification or false dedup state, status zero-network, Guardian health, and no artifact/install/SMTP/configuration mutation.
- Review exact changed/staged paths, modes, ownership/ACLs, privacy/secret boundaries, localization/documentation consistency, Git diff/check, rollback, commit/push, final clean synchronization, and idle lifecycle.


## Documentation Updates

- Update `docs/architecture/UPDATE_DISCOVERY_AND_RELEASE_AWARENESS.md` plus relevant notification, Guardian, system-map, security/privacy, operations/update records.
- Update EN/HU operator guidance and README/CHANGELOG where convention requires.
- Document trigger trust requirements, authenticated identity deduplication, restart and delivery-failure behavior, disabled-channel behavior, manual check/status semantics, privacy, and the awareness/notification/installation boundary.
- Maintain `ai/history/080_2026-09-03_update-notification-deduplication.md`, active/archive prompt, and concise core milestone records required by convention.
- Record actual paths, decisions, tests, isolated and production acceptance, rollback, commits/push, final state, and limitations in English.


## Completion Criteria

- Implementation, authenticated eligibility, persistent restart-safe deduplication, correct failure/retry semantics, disabled-channel behavior, tests, race/resource checks, documentation, deterministic isolated newer-release acceptance, bounded official current-release acceptance, repository validation, history/evidence, targeted commits/push, and lifecycle closure all pass.
- No artifact download/staging, automatic installation, version/package mutation, update restart, telemetry, registration, API key, listener, identifying request metadata, real production notification, or production SMTP/configuration mutation occurs.
- Manual update check remains non-notifying and update status remains zero-network/zero-notification.
- Final state is clean synchronized `main`, `HEAD == origin/main`, divergence `0/0`, and canonical idle after archived complete Task 080. Result must report `TASK 080 — ACCEPTED / COMPLETE` with implementation commit(s), closure commit, notification/deduplication and production acceptance summaries, validation summary, and final repository/lifecycle state. Do not begin Task 081.


## Owner Approval Requirements

Approved by Project Owner — explicitly authorized Task 080 and supplied the exact engineering scope with APPROVE on 2026-09-03 UTC. through the Engineering Task Builder on 2026-09-03 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.

On 2026-09-03 UTC the Project Owner explicitly authorized a bounded scope
expansion to diagnose and remediate the production Guardian OOM blocker,
including directly related code/configuration correction, rollback-protected
production state repair, and one required Guardian restart. Unrelated tuning,
infrastructure, SMTP, release publication and Task 081 remain prohibited.
