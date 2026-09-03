# Current Engineering Task 079: Guardian Isolated Automatic Release Checking

## Task Metadata

- Task ID: `079`
- Task slug: `guardian-isolated-automatic-release-checking`
- Status: `complete`
- Date opened: `2026-09-03` UTC
- Human authority: Project Owner — explicitly authorized Task 079 and supplied the exact engineering scope with APPROVE on 2026-09-03 UTC.
- Owner or lead-developer communication language: English

## Title

Guardian Isolated Automatic Release Checking


## Objective

Extend QWSG Community so the Guardian periodically performs an isolated, read-only authenticated production release-awareness check through the Task 076–078 release authority and persists the result through qwsg.update-awareness/1. The default interval is 24 hours. The work is complete only after implementation, deterministic automated coverage, documentation, bounded production acceptance, full repository validation, archived evidence, pushed commits, clean synchronized Git state, and canonical idle lifecycle.


## Scope

- Integrate a non-overlapping, cancellation-aware, bounded automatic release-index check into the existing Guardian lifecycle.
- Reuse the existing production release source, Ed25519 trust anchor, installed classification, rollback/future-index protections, update-awareness manager, and persistence contract.
- Use a conservative 24-hour default interval and the existing configuration architecture if interval configuration fits naturally.
- Derive due/not-due startup behavior from persisted awareness timestamps so short restarts do not duplicate retrieval.
- Preserve manual `qwsg update check` and local-only `qwsg update status` semantics.
- Add focused and integration tests, resource/lifecycle checks, EN/HU user documentation where convention requires it, architecture/system records, chronological history, and production acceptance evidence.
- Perform the authorized outbound-only production acceptance against `https://releases.quantumwizard.hu/qwsg/v1/release-index.json` in an isolated Community-equivalent state boundary without mutating the active installed production state.
- Perform task-scoped targeted Git integration, clean fast-forward push to `origin/main`, and canonical idle closure.


## Out of Scope

- Automatic or manual artifact acquisition changes, artifact download, staging, update installation, privileged update operations, version mutation, update-triggered restart, or package mutation.
- Update notification, notification deduplication, email/SMTP changes, Task 080, Task 081, or Community 1.3 release work.
- Pro functionality, API keys, telemetry, registration, installation identifiers, custom identifying request headers, inbound listeners, or central infrastructure integration.
- Hestia, Nginx, Apache, DNS, TLS, release-hosting, system-package, release publication, tag, or public artifact mutations.
- Unrelated infrastructure, Guardian, configuration, or update redesign.


## Authority Envelope

1. **Task targets and boundaries:** Implement and document only Guardian-driven isolated automatic release awareness using the accepted Task 076–078 contracts, with the exact objective, scope, exclusions, tests, acceptance, Git integration, and lifecycle closure defined here. Preserve Community 1.2.0 compatibility and do not begin Task 080.
2. **Permitted external actions:** Read-only Git fetch/push verification and clean fast-forward task commits to the canonical repository; outbound credential-free HTTPS retrieval of only the authorized production release index for bounded acceptance; read-only Guardian/service health inspection. No production configuration or installed-state mutation is authorized.
3. **Owner-reserved decisions:** Any material architecture change, scope expansion beyond Task 079, privileged production mutation outside the authorized acceptance scope, irreversible/destructive operation, security-policy exception, publication/tag/release action, or resolution of a genuine security issue requiring product policy.
4. **Task-specific STOP conditions:** Stop only if a material architecture change or scope expansion is required; privileged production mutation beyond acceptance is required; an irreversible/destructive action becomes necessary; a genuine security issue requires Owner decision; required rollback cannot be established; or minimum root authentication is the sole remaining blocker, in which case return the exact bounded Owner command/script and resume after execution.


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

- Verify the repository root, Framework identity, lifecycle, current task identity, branch, complete status, remote URL, fetched `origin/main`, HEAD, divergence, permissions/ACLs, relevant services, installed version, production trust identity, and Task 078 closure baseline before target changes.
- Expected intake baseline: canonical idle after complete Task 078; clean `main`; `HEAD == origin/main == 7d16e1c7cc486940e8c9b11b498caecde4d3ec59`; divergence `0/0`; production endpoint `https://releases.quantumwizard.hu/qwsg/v1/release-index.json`; required media type `application/vnd.quantumwizard.qwsg-releases+json`; installed/current stable release `1.2.0`; key ID `qwsg-community-release-2026-01`; trust fingerprint `0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6`.
- Inspect the Guardian recurrence/composition, configuration, update-awareness state, release-discovery, CLI status/check, installation classifier, systemd resource limits, documentation conventions, and Task 074–078 records before design decisions.
- Diagnose recoverable in-scope variance. Stop on a material identity, authority, security, privacy, destructive, external-mutation, or rollback boundary.


## Snapshot Requirements

- Retain and verify the pre-intake snapshot `/tmp/qwsg-task079-intake.9Ly6eJ`, created before lifecycle mutation with mode 0700, a mode-0600 complete Git bundle at baseline `7d16e1c7cc486940e8c9b11b498caecde4d3ec59`, and a mode-0600 lifecycle archive. Recorded SHA-256 values are `6ff6a30a4b8be033015aba43c44a45989f5531a89a3f8b878db61bb3efd6ecbf` and `ff910ef9ecc4da0aa9690c8e4d6072c402e299af5a5f4fe933d2ad936b02004a` respectively.
- After Builder installation and before implementation target changes, create a second timestamped mode-0700 execution snapshot under `/tmp` containing a complete verified Git bundle, exact prompt/history before-images, relevant implementation/documentation targets, permissions/ACL evidence, manifest hashes, and bounded rollback instructions. Keep payloads outside Git and retain both snapshots through Owner acceptance.


## Risk Assessment

- High: a release check could accidentally cross into acquisition, installation, notification, telemetry, or production mutation. Mitigate with dependency boundaries, injected fakes, side-effect assertions, and acceptance artifact/state inspection.
- High: authentication/source/media/rollback/future-index protections could diverge from manual checking. Mitigate by sharing the Task 076–078 checking core and regression tests.
- Medium: recurrence could overlap, retry tightly, leak goroutines, ignore cancellation, or create restart storms. Mitigate with one sequential owner, persisted due-time calculation, bounded timeout, no immediate failure retry, race/lifecycle tests, and resource checks.
- Medium: failed checks could crash or stall Guardian monitoring. Mitigate with failure isolation and continuation tests.
- Medium: awareness-state compatibility or filesystem permissions could regress. Mitigate with existing schema/store reuse, atomic persistence tests, migration-free compatibility checks, and mode/ACL inspection.
- Low: documentation/localization drift. Mitigate with EN/HU convention inspection and synchronized updates.
- Rollback risk is low while changes remain source/history-only and bounded snapshots plus Git before-images are retained.


## Planned Work

1. Validate governance, Git/lifecycle identity, snapshot evidence, and read all required architecture, security, release, Guardian, configuration, update-awareness, and Task 074–078 records.
2. Inspect the existing Guardian service recurrence/composition and Task 076–078 checker seams; select the smallest single-core integration and document deterministic due semantics.
3. Add a default 24-hour automatic-check schedule, configuration support only if native to the existing contract, persisted startup due calculation, bounded timeout/cancellation, sequential non-overlap, and failure isolation from Guardian operation.
4. Reuse the production checker and awareness manager without adding any acquisition, installation, notification, credential, identifying-header, listener, telemetry, or registration path.
5. Add deterministic unit/integration coverage for due/not-due, restart suppression, authenticated current/newer results, all required failure classes, persistence/local-only status, side-effect isolation, non-overlap, Guardian continuation, and cleanup.
6. Update architecture, operator EN/HU documentation, examples/configuration references if applicable, changelog/system records, and Task 079 history.
7. Run focused, full, race, resource, engineering/framework/lifecycle, reproducibility, security, and relevant release tests; diagnose and correct in-scope failures.
8. Run bounded production acceptance against the exact authorized endpoint using isolated state and prove authenticated current 1.2.0, subsequent not-due behavior, zero-network status, healthy Guardian, and absence of artifact/install/notification/unrelated mutation.
9. Review exact diffs and privacy, stage only explicit task paths, commit and push clean fast-forward changes, finalize prompt/history, archive Task 079 into canonical idle state, push closure, and verify `HEAD == origin/main`, divergence `0/0`, clean worktree, lifecycle idle.


## Rollback Plan

- Before implementation, enumerate exact Task 079 target paths in the execution snapshot. Restore only those paths from the verified execution archive or baseline Git bundle after confirming current hashes and preserving any unrelated user changes.
- For uncommitted source/document changes, apply exact before-images only to the enumerated Task 079 targets, then run formatting, focused tests, framework/lifecycle checks, and Git status.
- For lifecycle intake rollback before implementation commit, remove only Builder-created `ai/prompts/079_CURRENT_TASK.md` and its exact matching Task 079 history after collision/hash verification, restoring canonical idle validation. Never use broad reset, checkout, clean, wildcard deletion, or extraction over the live worktree.
- After a pushed commit, prefer a new explicit revert commit limited to Task 079 paths; do not rewrite published history. No production installed state should require rollback because acceptance is read-only and isolated.


## Deliverables

- Production-ready Guardian automatic authenticated release-awareness integration with default 24-hour cadence, deterministic restart-safe due behavior, bounded cancellation-aware retrieval, non-overlap, and failure isolation.
- Existing qwsg.update-awareness/1 persistence and shared authenticated release-checking core retained.
- Deterministic automated tests covering every required success, failure, isolation, lifecycle, and resource case.
- Updated English engineering documentation and HU/EN user-facing documentation according to project convention.
- Bounded production acceptance evidence for the exact Task 078 authority without release fabrication or production mutation.
- Complete Task 079 prompt/history/archive evidence, implementation and closure commits pushed, synchronized clean Git state, and canonical idle lifecycle.


## Verification

- `bin/job --check`, Framework validation, lifecycle consistency, exact task/history identity, and snapshot integrity pass.
- Formatting, vet, full tests, `go test -race ./...`, canonical engineering/framework tests, lifecycle checks, and relevant reproducibility/security/release tests pass.
- Automated evidence proves due after interval, not due before interval, restart suppression, authenticated current/newer classifications, network failure, timeout/cancellation, invalid media type/signature/key/index, rollback/future rejection, awareness persistence, status zero-network, no artifact/install/notification effects, non-overlap, Guardian continuation, and cleanup.
- Configuration/default behavior is backward-compatible; Guardian MemoryMax/TasksMax/GOMEMLIMIT behavior remains accepted.
- Production acceptance authenticates the exact endpoint/media type/key/fingerprint, classifies installed 1.2.0 against stable 1.2.0, persists valid isolated awareness state, proves the next cycle not due, proves local status performs no network, and confirms no artifact, install, notification, or unrelated configuration mutation.
- Review exact diff, modes, ownership/ACLs, privacy/security boundaries, documentation consistency, rollback validity, and excluded paths.
- Final prompt/history are complete, Task 079 is archived, commits are pushed, `HEAD == origin/main`, divergence is `0/0`, worktree is clean, and `bin/job --check` reports canonical idle after Task 079.


## Documentation Updates

- Update `docs/architecture/UPDATE_DISCOVERY_AND_RELEASE_AWARENESS.md` and the relevant Guardian/update-awareness architecture/system-map records with the implemented scheduling and isolation semantics.
- Update EN/HU operator documentation, CLI/configuration references, examples, and README/CHANGELOG where project conventions require them.
- Document default 24 hours, due calculation, restart behavior, failure behavior, privacy, awareness-versus-installation, Task 079-versus-Task 080 notification boundary, and manual check/status semantics.
- Maintain `ai/history/079_2026-09-03_guardian-isolated-automatic-release-checking.md`, the active/archive prompt, and concise core engineering milestone records required by convention.
- Record actual changed paths, verification evidence, production acceptance, rollback, commits, push, final lifecycle/Git state, and any disclosed limitations in English.


## Completion Criteria

- Implementation, required tests, race/resource checks, documentation, bounded production acceptance, repository validation, history/evidence, targeted commits/push, and lifecycle closure all pass.
- No automatic installation, artifact acquisition/download/staging, notification, telemetry, registration, identifying request metadata, inbound service, or unrelated production mutation exists.
- Manual `qwsg update check` remains authenticated and `qwsg update status` remains local-only.
- Final state is clean synchronized `main` with `HEAD == origin/main`, divergence `0/0`, and canonical idle lifecycle after archived complete Task 079.
- Completion result must be `TASK 079 — ACCEPTED / COMPLETE` with implementation commit(s), closure commit, production acceptance result, test summary, and final repository/lifecycle state. Do not start Task 080.


## Owner Approval Requirements

Approved by Project Owner — explicitly authorized Task 079 and supplied the exact engineering scope with APPROVE on 2026-09-03 UTC. through the Engineering Task Builder on 2026-09-03 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
