# Current Engineering Task 039: Guardian Runtime Integration and Console UX Hardening

## Task Metadata

- Task ID: `039`
- Task slug: `guardian-runtime-integration-and-console-ux-hardening`
- Status: `complete`
- Date opened: `2026-08-11` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Guardian Runtime Integration and Console UX Hardening


## Objective


Correct the proven live-host blockers that prevent the existing QWSG Version 1.0 Release Candidate from operating as a genuinely usable local Guardian. Make valid large canonical Policy Reports compatible with bounded Alert evaluation, make interactive Console refresh a read-only Current Operator State reload, preserve privacy-safe Runtime component failure causes through the operator projection, and prove that stale lifecycle evidence cannot continue to represent a dead Guardian as running.

Task 039 is complete only after automated regression coverage and a real local product acceptance using the locally built executable demonstrate at least two recurring full Guardian cycles, successful Alert evaluation for the real 366-or-more-source Policy Report, completed Runtime and notification planning reachability, truthful Current Operator State publication, a working separate Console and read-only `r` refresh, retained manual-observe concurrency protection, useful bounded diagnostics, and truthful graceful-stop and unexpected-termination behavior.

This task shall make the smallest compatible corrections inside the established Version 1.0 architecture. It shall not create new architecture, add optional features, weaken canonical validation, or start clean-VPS/reboot acceptance.



## Scope


- Correct the Runtime-to-Alert integration so a valid canonical Policy Report with 366 or more sources, and arbitrary larger valid reports within the Report contract, does not fail solely because Alert's per-reference bound is 64.
- Preserve `alert.MaxReferences` as the bound for Alert-owned reference arrays and controls. Do not solve the defect by increasing it.
- Represent a canonical Policy Report in Alert evidence by its validated aggregate envelope identity: the Alert source reference identifies the Policy Report contract and canonical Report ID, while its bounded evidence-reference list contains only the canonical aggregate identity needed for correlation. The Policy Report remains the canonical object containing the complete source list and full traceability.
- Ensure the aggregate-reference strategy is deterministic, stable across identical inputs, independent of host-specific values, privacy-safe, bounded regardless of Policy Report source count, and compatible with Alert condition identity, lifecycle, resolution, serialization and validation.
- Apply the same aggregate-envelope rule to the legacy canonical Report input only where the same existing compatibility defect and shared Alert path require it; do not change Report generation or semantics.
- Add realistic Alert and Runtime integration regressions at 64, 65, 366 and a larger valid source count. Prove stable identity/correlation, preserved Report traceability, bounded Alert references, no private payload projection, successful Alert evaluation, and progression into Notification planning when upstream canonical inputs are valid.
- Change interactive Console `r` to call a read-only local provider that loads Current Operator State, performs the existing path, integrity, schema, permission and ownership validations, requalifies freshness at the refresh time, and returns the resulting Operator Overview for rendering.
- Ensure Console refresh performs no Guardian-lock acquisition, `observe` execution, Inventory collection, Pipeline execution, state publication or other writer action. Bare/noninteractive Console loading and initial interactive loading shall use the same canonical read-only semantics where practical without redesign.
- Preserve explicit manual active observation as `qwsg observe`. Preserve the existing operation lock and canonical `guardian_active` refusal when manual observe is attempted while Guardian owns the lock.
- Project bounded, canonical Runtime component `FailureToken` values into operator attention when Runtime does not complete. At minimum distinguish `alert_evaluation_failed`, `notification_planning_failed`, `notification_delivery_failed`, `runtime_timeout`, `runtime_cancelled`, and other existing validated component failure tokens that are already present in Runtime Result.
- Keep the generic Runtime outcome/completeness fact where useful, but do not let it erase a more specific canonical component cause. Do not expose raw Go errors, paths, hostnames, IP addresses, secret values, source payloads or unbounded diagnostics.
- Add or reuse localized presentation tokens so Console Overview/Attention communicates practical English and Hungarian meanings and an existing bounded recommended action, for example that Alert evaluation failed and Guardian details should be inspected. Do not build a general diagnostics framework.
- Verify Task 036 bounded attention projection and correlation with realistic large live-shaped evidence. Correct only a reproduced defect that creates meaningless duplicate Runtime/component attention or obscures the actionable failure cause; retain the existing `MaxAttention`, correlation summary and omission disclosure.
- Verify lifecycle truthfulness across checkpoint `active`, runtime-service evidence, Current Operator State and freshness requalification after unexpected process disappearance. A current, externally correlated running observation may remain running only until its exclusive freshness boundary. At or after that boundary, Console must requalify it to stale/unavailable and must not claim a live Guardian.
- If current freshness behavior fully guarantees unexpected-termination truthfulness, preserve production semantics and add regression plus local acceptance evidence. If evidence proves a gap, implement only the minimum demotion/freshness correction in the existing Guardian, Current State or Presentation boundary; do not add a supervisor.
- Perform real local acceptance on the current development host with the locally built binary, repository-local private state/configuration, bounded short intervals and a uniquely recorded Guardian generation. Use the actual host Inventory and full canonical pipeline, not only synthetic fixtures.
- Update only directly affected tests, architecture/operator documentation, release notes/changelog, Task 039 history and lifecycle records. Preserve unrelated Owner changes and the Task 038 Release Candidate artifacts unless a correction directly requires rebuilding them for local acceptance.



## Out of Scope


- No Dashboard, REST API, network Console, fleet, remote management, cloud service, AI, remediation, licensing enforcement, updater, package repository, new persistence platform or new monitoring architecture.
- No new notification provider or transport. Provider-neutral Notification planning/delivery contracts may be exercised only to prove Runtime stage reachability and existing semantics.
- No redesign or reopening of Scheduler, Inventory, Snapshot, Comparison, Drift, Health, Rule, Policy, Report, Alert, Runtime, Current Operator State, Presentation or Console architecture beyond the minimum direct corrections listed in Scope.
- No weakening of manual-observe/Guardian concurrency protection, state integrity, privacy, deterministic identity, canonical validation, freshness, boundedness or fail-closed behavior.
- No arbitrary increase of `alert.MaxReferences`, `presentationmodel.MaxAttention` or other bounds as a substitute for correct aggregation/correlation.
- No duplication or truncation of the Policy Report's canonical full source traceability. Alert must refer to the aggregate; Report continues to own its validated source list.
- No raw error propagation or general-purpose diagnostics framework.
- No new process supervisor. Production supervision remains systemd.
- No change to system linger, real installed service activation, unrelated user services, external infrastructure, remote hosts or clean-VPS/reboot acceptance.
- No dependency installation, new third-party dependency, schema migration platform or unsupported-platform expansion.
- No staging, commit, tag, push, publication or release authorization. Do not install Task 039 as part of preparation; Builder installation is a later owner-authorized lifecycle action.



## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/config/engineering-project.conf`

## Starting State Verification


1. Start at the verified repository root and run `bin/job --check`. Require Task 038 to be the latest complete archived task, its unique matching history to be complete, `ai/prompts/` to be canonically empty before Builder installation, and Task 039 to be the sole active prompt with one matching history after installation.
2. Run the Reusable Engineering Framework validation and inspect the current Builder/lifecycle interfaces. Verify Task 039 metadata, authority, scope and exclusions exactly match this Owner definition; stop on unresolved markers, prompt/history mismatch, later task, collision or lifecycle temporary residue.
3. Record UTC time, effective user, working directory, `VERSION`, platform and architecture, Go version, systemd user-manager availability/version, branch, HEAD, configured remotes, upstream/ahead-behind state, full Git status, empty/nonempty index and all existing untracked/modified Owner paths. Do not fetch unless separately authorized and necessary.
4. Treat the existing dirty worktree as Owner content. Identify the exact Task 039 targets and compare them with the Task 038 completed snapshot/history. Stop before overlap if a target changed after the accepted Task 038 baseline and safe preservation cannot be proved.
5. Record relevant state-root/configuration paths by sanitized logical name, ownership, modes, ACL availability, symlink status, filesystem atomic-rename/advisory-lock support, free space and current QWSG/Guardian process or user-unit state. Do not expose private absolute paths or host identity in committed evidence.
6. Verify from source and tests that Alert currently limits Report and Policy Report source counts against `MaxReferences`, constructs Report-level evidence from every Report source, and uses bounded per-Alert source references; record exact symbols and current contract versions.
7. Verify Runtime's actual component failure tokens and stage ordering, including Alert failure conversion and Notification planning/delivery reachability. Verify Policy Report validation and full source traceability remain canonical outside Alert.
8. Verify that `localOverviewProvider.Refresh` currently resolves `observe`, acquires the one-shot Guardian lock and runs `observeOnce`, while initial Console loading reads Current Operator State and calls `RequalifyFreshness`. Record the reproduced `guardian_active` refresh failure without weakening the lock.
9. Verify how Runtime component failure tokens are validated and persisted and how Operator Presentation currently reduces a non-completed Runtime to `runtime_not_completed`. Inventory only existing canonical bounded tokens; do not infer raw error text as an interface.
10. Verify Guardian checkpoint `Active`, `ReportExit`, process-exit behavior, runtime-service evidence publication, Current Operator State `FreshUntil`, and `RequalifyFreshness`. Establish whether an ungraceful foreground death bypasses `report-exit` and leaves checkpoint `active=true`, and whether Console nevertheless demotes running evidence at the exclusive freshness deadline.
11. Preserve sanitized evidence of the live defect: a valid 366-source Policy Report, Alert validation failure classification, partial Runtime outcome, skipped Notification stages, degraded Guardian, generic operator diagnosis and Console refresh lock conflict. Do not commit raw Inventory, state files, usernames, paths, hostname/IP data or journal payload.
12. Run the existing focused and full baseline tests with bounded writable caches. Stop on a relevant pre-existing failure or a material environment difference; unrelated known failures must be disclosed and isolated rather than silently repaired.



## Snapshot Requirements


Before modifying any Task 039 target, create a UTC-stamped, mode-`0700`, rollback-capable snapshot outside Git under a unique `/tmp/qwsg-task039-*` directory. Capture the exact pre-change bytes and mode of every intended tracked or untracked target, record verified absence for new paths, and include sanitized Git baseline metadata, Task 039 prompt/history identity and the exact local acceptance unit/state/config/artifact identities. Do not copy unrelated Owner content.

Create a deterministic manifest with logical repository-relative names, type, size, mode and SHA-256 for every payload object. Verify checksums and archive readability after creation. Record an exact bounded restore procedure before implementation, including collision/identity checks, service-stop preconditions and post-restore validation. Retain the payload until Owner acceptance and rollback-window closure; do not commit snapshot payload, raw diff, private host evidence or absolute private paths.

Before local product acceptance, create a separate service/state safety record identifying the exact locally built binary hash, Guardian generation, process or temporary unit identity, state/config/store roots, pre-existing process/unit state, fresh-until bound and cleanup targets. Never target a broad home, repository, system or shared state directory.



## Risk Assessment


- **Alert semantic weakening — high:** an aggregate reference must change only evidence projection, not condition precedence, severity, completeness, lifecycle, resolution, validation or Notification decisions. Mitigate with equivalence tests across small and large Reports.
- **Traceability loss — high:** bounding Alert references could hide Report provenance. Mitigate by retaining the validated Policy Report ID/contract as the Alert source and proving the canonical Report still carries all source IDs and validates independently.
- **Identity or compatibility drift — high:** changing reference material can alter Alert condition/lifecycle identities or stored Alert state. Freeze deterministic aggregate correlation, test repeat identity, prior-state continuation and supported stored schema behavior; avoid schema change unless strictly required and documented.
- **Unbounded valid input — high:** canonical Reports can exceed the discovered 366 sources. Ensure Alert work/memory and persisted references stay bounded by aggregate identity, with regression above 366 and no source-count-proportional Alert evidence copy.
- **Privacy leakage — high:** live host evidence and diagnostic errors may contain sensitive data. Permit only validated canonical IDs/tokens and localized static meanings; sanitize acceptance evidence and inspect state/Console/journal output.
- **Refresh side effects — high:** reusing observation code can retain lock, collection or publication behavior. Introduce or reuse one read-only loader path and assert prohibited calls are absent through focused integration tests.
- **Concurrency regression — critical:** a refresh fix must not weaken explicit `qwsg observe` exclusion. Test a real running Guardian plus concurrent manual observe and require the canonical safe refusal.
- **False lifecycle claim — critical:** stale checkpoint or state can mislead the operator after SIGKILL. Bound running truth by Current State freshness, test both sides of the exclusive deadline and use minimal demotion only if current behavior fails.
- **Diagnostic explosion — medium:** projecting every component/event could create duplicate attention. Select failed component results only, preserve deterministic ordering/correlation and Task 036 bounds/summary.
- **Localization regression — medium:** add static EN/HU meanings and fallback tests without exposing internal tokens as untranslated user text where a catalog entry is required.
- **Live-host interference — critical:** use an exact process/generation and isolated local roots, inspect before signalling, never alter linger or unrelated services/state, and perform verified cleanup.
- **Acceptance flakiness — medium:** use bounded short intervals and wait for persisted correlated cycle evidence rather than sleeps alone; record deadlines and stop safely on timeout.
- **Dirty worktree/data loss — critical:** preserve unrelated Owner files, inspect every overlap, use the snapshot for exact rollback, and prohibit reset/clean/checkout/restore and broad deletion.



## Planned Work


### Phase 1 — Contract confirmation and regression design

1. Load Task 039 through `job`, read all required governance and component documents, verify the installed prompt/history, record the exact starting state, and create/verify the implementation and local-acceptance snapshots.
2. Trace Policy Report identity and source ownership through Report, Alert, Runtime, Notification, Current State and Presentation. Freeze the narrow contract: Report owns full sources; Alert owns one aggregate Report source reference with bounded canonical aggregate identity.
3. Reproduce the 64/65/366/larger-source boundary in deterministic fixtures and reproduce Console `r` contention with Guardian lock held. Add failing regressions before or with production corrections.
4. Trace all existing Runtime component failure tokens and lifecycle freshness transitions. Define a closed projection mapping only for already canonical tokens, with safe fallback for an existing validated but unmapped token.

### Phase 2 — Minimal integration corrections

5. Remove Report source-count coupling from Alert input validation while retaining full Report validation. Change Report-level Alert source construction to aggregate-envelope correlation whose reference count is constant and within `MaxReferences` for every valid Report size.
6. Prove Alert evaluation and Runtime advance through Notification planning for a complete large Policy Report. Preserve existing behavior for incomplete Reports, Alert precedence, lifecycle continuity, resolution and Notification semantics.
7. Replace Console refresh's active-observation provider with a Current Operator State loader that performs the same safe load and freshness requalification as initial Console startup. Factor only enough shared code to prevent semantic drift; keep render/session architecture unchanged.
8. Extend Operator Presentation consumption of Runtime Result so specific failed component tokens become bounded attention with localized static meaning and actionable existing recommendation. Preserve generic outcome/completeness evidence without redundant attention spam.
9. Verify Task 036 attention selection/correlation at live-shaped size. Make no presentation change unless a deterministic regression demonstrates duplicated/obscured release-blocking output.
10. Test graceful, abnormal-current and abnormal-stale lifecycle cases. If the existing exclusive `FreshUntil` demotion passes, add regression evidence only. If it fails, implement the smallest correction at the existing freshness/lifecycle boundary and document why.

### Phase 3 — Automated and real local acceptance

11. Run focused Alert, Runtime, Presentation, Current State, Console, Guardian and command tests; then full build, tests, race, vet, formatting, static/privacy, Framework, lifecycle and rollback checks.
12. Build the repository-local executable and record its hash/version. Use isolated private local roots and a unique Guardian generation. Start `./qwsg guardian run` in a controlled foreground or PTY process and wait on persisted evidence for at least two completed real cycles.
13. In a separate process run `./qwsg`; verify Guardian is `running` when the latest Runtime cycle completed, evidence is current, the actual 366-or-more-source Policy Report did not cause partial Runtime, Notification planning was reached, Current Operator State is correlated and server condition remains evidence-derived without false certainty.
14. In the interactive Console press `r` and verify state reload/redraw succeeds, freshness is requalified, no observe starts, no Inventory/State publication occurs, no Guardian lock conflict appears and no generic refresh failure is shown. Inspect Guardian detail/Attention for understandable bounded cause text.
15. While Guardian still holds the operation lock, run `./qwsg observe` separately and require the existing canonical Guardian-active refusal with no state corruption or second pipeline.
16. Stop Guardian gracefully through the supported signal path. Verify checkpoint inactive/stopped lifecycle evidence and a later Console does not claim running.
17. Restart with isolated roots, wait for current running evidence, terminate the exact verified Guardian process unexpectedly, and verify Console behavior before and at/after `FreshUntil`: it may retain current evidence before expiry but must show stale/unavailable rather than running at expiry. Record checkpoint behavior without treating stale `active=true` as a live-process oracle.
18. Clean up only Task 039 acceptance processes and exact isolated roots after identity checks. Prove no process/unit/state/config residue and no change to real service, real operator state, linger or unrelated infrastructure.

### Phase 4 — Documentation, delivery and lifecycle closure

19. Update only directly affected architecture, operator, release/changelog and test traceability documents. Record exact behavior, bounds, diagnostic mapping, freshness semantics and sanitized local evidence in Task 039 history.
20. Rebuild and rerun mandatory acceptance after documentation/code finalization. Review full diff, permissions, ACLs, secrets/private-host evidence, snapshot integrity, rollback validity, index, branch/HEAD and unrelated Owner paths.
21. Mark Task 039 prompt/history complete only if every completion criterion passes. Archive it transactionally and verify canonical idle with Task 039 as the latest complete task. Do not generate a successor, stage, commit, tag, push or publish.



## Rollback Plan


Stop rollback if the active Task 039 acceptance process, target bytes, snapshot manifest, expected owner/mode or locally built binary identity differs from recorded facts. First stop only the exact Task 039 Guardian process after verifying its generation and executable identity. Do not signal another QWSG process or alter an installed/shared user unit.

Restore only Task 039 target paths captured in the verified snapshot. For pre-existing paths, atomically restore exact captured bytes and modes after collision checks. Remove a task-created path only when the manifest recorded its pre-task absence and its current content/hash proves Task 039 ownership. Never use wildcard deletion, Git reset, Git clean, Git checkout, Git restore or extraction over the live worktree.

Preserve canonical Inventory snapshots, Scheduler state, Guardian checkpoint, Current Operator State and failure evidence until the rollback/report decision is complete. For isolated Task 039 acceptance roots, remove only exact manifest-listed paths after all associated processes are stopped and evidence is retained; leave real QWSG state untouched.

After rollback, run focused and full tests, build, format/vet checks, `bin/job --check`, Framework/lifecycle validations, Git status/index/diff checks, snapshot SHA-256 verification and local process/residue checks. Confirm the original Alert limit behavior, Console provider behavior and documented baseline are restored consistently. Retain the snapshot and sanitized failure report until Owner acceptance.



## Deliverables


- a bounded Runtime-to-Alert aggregate Policy Report evidence-reference correction supporting 366 and arbitrary larger valid source lists without increasing `MaxReferences`;
- preserved full canonical Policy Report source traceability and stable deterministic Alert/Runtime correlation;
- regression fixtures and tests at 64, 65, 366 and above 366 sources, including Runtime progression into Notification planning;
- a read-only Console refresh path that loads, validates and freshness-requalifies Current Operator State without lock acquisition, observation, collection, Pipeline execution or publication;
- preserved canonical `guardian_active` refusal for explicit manual `qwsg observe` during Guardian ownership;
- bounded privacy-safe operator attention for existing Runtime component failure tokens, with understandable EN/HU Console meanings and practical recommendation;
- Task 036 attention/correlation sanity evidence at live-shaped scale and only any proven minimal correction;
- automated and real-host evidence that unexpected termination cannot produce a running claim at or beyond the Current State freshness boundary;
- real local acceptance evidence for at least two recurring cycles, large Policy Report Alert success, completed Runtime, Current State publication, separate Console, refresh, manual-observe exclusion, graceful stop and unexpected termination;
- directly affected architecture/operator/release documentation, changelog entry, completed Task 039 history/archive and canonical idle closure.



## Verification


- Validate Task 038 canonical idle before installation and Task 039 prompt/history identity after Builder installation. Re-run Framework, Builder, lifecycle, diverted-task and `bin/job --check` validations at the applicable phases.
- Verify starting/final Git root, branch, HEAD, remote configuration, upstream relationship, full status, index, exact changed paths, modes/ACLs and preservation of all unrelated modified/untracked Owner content. Run `git diff --check` and secret/private-host review.
- Run `make build`, `go test ./...`, `go test -race ./...` with bounded writable caches, `go vet ./...`, `make fmt-check`, plus focused tests for `internal/alert`, `internal/runtime`, `internal/presentationmodel`, `internal/operatorstate`, `internal/operatorconsole`, `internal/guardian`, `internal/app` and `cmd/qwsg`.
- Alert tests must use canonical valid Policy Reports with exactly 64, 65, 366 and at least one materially larger source count. Assert validation success, constant/bounded aggregate evidence references, stable IDs across repeat evaluation, full source retention in the Report, no payload/private data in Alert evidence, and unchanged incomplete/complete Report semantics.
- Runtime integration tests must prove a 366-or-more-source valid Policy Report yields successful Alert component evaluation and reaches Notification planning. With the existing empty provider-neutral policy it may produce no delivery work, but Alert size alone must not cause partial outcome or skip Notification stages.
- Negative tests must still reject invalid/tampered Report IDs, schema, ordering, duplicate sources, invalid source records, oversized Alert-owned evidence arrays and every existing unsafe control. Do not weaken `MaxReferences` validation for Alert records or Notification references.
- Compatibility tests must cover previous Alert state/lifecycle continuation and deterministic resolution across the aggregate-reference change. If persisted schema/identity compatibility would break, stop and choose a compatible narrow design rather than silently migrate.
- Console unit/integration tests must prove initial and refreshed loads use Current Operator State validation plus exclusive-deadline freshness requalification, retain the last valid view on genuine read failure, redraw once, and make zero calls to observe/Inventory/Pipeline/publication/Guardian locking.
- With a real Guardian lock held, interactive `r` must succeed from stored current state and explicit `./qwsg observe` must fail safely with the canonical Guardian-active behavior. Verify no second snapshot, scheduler cycle or Current State publication is caused by refresh.
- Presentation tests must cover every minimum Runtime failure token and any other currently canonical component token, deterministic ordering, safe fallback, EN/HU/fallback rendering, bounded attention, correlation/omission summary, recommendations and absence of raw error/path/host/secret content.
- Freshness tests must evaluate immediately before and exactly at `FreshUntil`. Before expiry the last validated current running evidence may remain; at expiry it must be stale and Guardian unavailable, completeness no stronger than partial, condition no false healthy/unknown certainty, and recommendations must direct a fresh check/Guardian verification.
- Guardian tests must cover graceful exit demotion/checkpoint inactive behavior, correlated abnormal exit when exit reporting exists, abrupt loss without exit reporting, stale checkpoint `active=true`, Current State aging and restart/checkpoint recovery. No supervisor or process-liveness polling is required.
- Real local acceptance must use the repository-built `./qwsg`, actual live Inventory and isolated private state/config/store roots. Wait for at least two persisted completed runtime cycles and verify the actual Policy Report source count is at least 366, Alert succeeded, Notification planning was reached, Runtime completed, state/checkpoint/scheduler updated and separate Console shows current truthful evidence.
- In the real interactive Console, press `r`; verify successful read-only refresh, one redraw, no `Refresh failed`, no `guardian_active`, no competing observe and useful bounded Guardian/Attention details.
- Run real concurrent manual `./qwsg observe` during Guardian lock ownership and require the existing safe refusal. Then graceful stop and later Console must not claim running.
- If practical and safe, restart and terminate the exact isolated Guardian unexpectedly. Observe the pre-expiry and post-expiry Console states and require no running claim at/after expiry. If host constraints prevent this one simulation, the task cannot claim complete unless equivalent real process evidence already exists and the Owner explicitly accepts the disclosed limitation; unit tests alone are insufficient.
- Verify no acceptance process, temporary unit, isolated state/config/store path or lock remains; real installed QWSG service/state, user-manager/linger and external infrastructure are unchanged.
- Verify documentation and behavior agree on aggregate traceability, read-only refresh, explicit observe, diagnostic tokens, graceful/abnormal lifecycle and exclusive freshness. Validate the implementation snapshot/checksums and exact rollback procedure remain usable.
- Mark complete only after the final built binary reruns the full local sequence. Archive Task 039, verify its prompt/history are complete and `bin/job --check` reports canonical idle with Task 039 latest. Do not stage, commit, tag, push or publish.



## Documentation Updates


- Update `docs/architecture/CANONICAL_ALERT_ENGINE.md` for aggregate canonical Report evidence ownership, bounds, traceability and compatibility.
- Update `docs/architecture/CANONICAL_RUNTIME_ENGINE.md` for large-Report Alert integration and existing component failure-token propagation.
- Update `docs/architecture/CANONICAL_OPERATOR_PRESENTATION_MODEL.md` and `docs/architecture/CANONICAL_CURRENT_OPERATOR_STATE.md` for bounded specific Runtime failure attention and refresh-time freshness requalification.
- Update `docs/architecture/INTERACTIVE_OPERATOR_CONSOLE.md` and EN/HU Console user guidance for read-only `r`, explicit `qwsg observe`, failure wording and freshness behavior.
- Update `docs/architecture/OPERATIONAL_GUARDIAN_SERVICE.md` for verified graceful/abnormal lifecycle truthfulness and the role of checkpoint versus freshness evidence.
- Update `CHANGELOG.md`, applicable release notes/known limitations and only directly affected README/install/troubleshooting text so the RC correction and local acceptance status are accurate. Do not claim clean-VPS or reboot acceptance.
- Update focused test/architecture traceability where maintained, Task 039 history throughout execution, Task 039 prompt status at completion and the concise engineering milestone index if required by established policy.
- Record sanitized commands, counts, versions, hashes, timings, state transitions, tests, cleanup, rollback and limitations. Do not store raw live Inventory, journal payload, private absolute paths, usernames, hostname/IP data or secret values.



## Completion Criteria


Task 039 is complete only when all of the following are true:

- valid canonical Policy Reports with 64, 65, 366 and a larger tested source count pass Alert input validation and evaluation without increasing `MaxReferences` or copying the full source list into Alert evidence;
- Alert references the validated canonical Report envelope/aggregate identity with constant bounded evidence, while the Policy Report retains and validates its complete ordered source traceability;
- Alert condition identity, severity, precedence, lifecycle, resolution, completeness and Notification semantics remain deterministic and are not weakened;
- a real 366-or-more-source Guardian cycle reaches successful Alert evaluation, Notification planning and completed Runtime instead of `alert_evaluation_failed`/partial solely because of source count;
- at least two real recurring cycles complete on the current development host with the locally built binary and publish correlated Scheduler, checkpoint and Current Operator State evidence;
- a separate Console truthfully shows Guardian `running` when the latest Runtime cycle completed and evidence is current, while server condition remains derived from canonical evidence without false certainty;
- interactive `r` performs only safe Current Operator State load, validation, freshness requalification and render; it starts no observe, acquires no Guardian operation lock, publishes no state and shows no generic refresh failure during normal Guardian operation;
- explicit manual `./qwsg observe` still fails safely with canonical Guardian-active behavior while Guardian owns the operation lock;
- operator evidence distinguishes at least Alert evaluation, Notification planning, Notification delivery, timeout and cancellation failures through bounded canonical tokens and understandable localized meanings, without raw/private/unbounded diagnostics;
- Attention remains bounded, correlated and useful at live-host scale, with no release-blocking duplicate-token flood and truthful omission disclosure;
- graceful stop publishes truthful stopped evidence and clears active checkpoint state; a later Console does not claim running;
- after unexpected termination, current evidence is never stronger than its freshness contract and Console shows stale/unavailable rather than running at or after the exclusive freshness deadline, regardless of a stale checkpoint `active=true` value;
- focused/full/race/vet/format/build, privacy/security, compatibility, Framework/Builder/lifecycle, snapshot/rollback, Git and real local acceptance checks pass with exact evidence;
- all Task 039 acceptance processes and isolated artifacts are removed safely, unrelated Owner content and real host service/state/linger remain unchanged, and no dependency installation, stage, commit, tag, push or publication occurred;
- directly affected documentation and Task 039 history are complete, limitations are explicit, prompt/history are archived consistently, and canonical idle reports Task 039 as the latest complete task.

Successful Task 039 means the existing local QWSG Version 1.0 product is demonstrably usable and ready to proceed to a separately authorized clean Ubuntu VPS installation/reboot acceptance. A new engineering task is required before that acceptance only if Task 039 uncovers a distinct release blocker that cannot be corrected within these bounded interfaces; optional or excluded product work does not justify a successor.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-11 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
