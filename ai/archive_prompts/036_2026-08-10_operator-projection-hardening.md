# Current Engineering Task 036: Operator Projection Hardening

## Task Metadata

- Task ID: `036`
- Task slug: `operator-projection-hardening`
- Status: `complete`
- Date opened: `2026-08-10` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Operator Projection Hardening


## Objective


Repair the Version 1.0-blocking Task 035 product regression in the existing Canonical Operator Presentation Model and its application diagnostics without changing the successful canonical evaluation pipeline.

A valid large `observe` execution must project into one deterministic, privacy-safe, resource-bounded Operator Overview even when valid Health, Rule, and Policy evidence produces more candidate attention concerns than `MaxAttention`. The projection must retain the most important concerns, correlate duplicate Rule/Policy views where canonical identities prove they describe the same concern, explicitly disclose aggregation or omission, preserve bounded source traceability, and never discard a later critical concern because earlier lower-severity candidates filled the limit.

The existing `observe` workflow must distinguish privacy-safe pipeline/stage, operator-projection, and Current Operator State publication failures. A successful large projection must publish atomically and be consumable by a separately started bare `qwsg`. Recommendations must distinguish refreshable stale evidence from current partial or failed non-self-healing conditions and must not imply that blind repetition will repair a known deterministic failure.

Task 036 is a narrow integration hardening task. It shall not add an engine, architectural layer, pipeline, persistence system, remediation system, or general UX redesign.



## Scope


- Confirm and freeze the live-host regression fixture: complete Inventory and Snapshot, 366-record complete Comparison/Drift/Health/Rule/Policy/Report, sufficient Health evidence, 366 matched Rule results, 366 observe Policy results, complete Command execution, and projection failure caused only by more than 256 candidate attention items.
- Keep `MaxAttention` unchanged and keep projection bounded for arbitrarily large valid input.
- Introduce the smallest Presentation Model-owned candidate-reduction step before Overview validation. Candidate collection may be proportional only to already contract-bounded input; emitted Overview attention remains at or below `MaxAttention`.
- Correlate a Rule attention candidate with its downstream Policy evaluation only when the existing `PolicyEvaluation.RuleEvaluationID` and validated Rule/Policy provenance prove the relationship. Represent the downstream Policy decision as the authoritative operator concern and count the correlated Rule view as aggregated, while retaining a stable source reference to the Policy evaluation and envelope-level Rule/Policy sources. Do not infer correlation from text, position, or mutable display fields.
- Do not merge unrelated Health, Rule, Policy, Runtime, service, or alert concerns. Preserve direct critical and warning Health evidence as independent operator concerns unless an existing canonical identity proves equivalence.
- Canonically rank all remaining candidates before reduction: severity descending (`urgent` before `review`), then operator importance by evidence kind with direct safety/health and final policy decisions ahead of intermediate rule views, then stable title token, reason token, source kind, source contract/version, source record ID, and observation time. Equivalent input sets must yield identical retained items and Overview identity regardless of enumeration order.
- Retain the first `MaxAttention` candidates only after canonical ranking. Because ranking occurs globally, no later urgent or otherwise higher-ranked candidate may be displaced by an earlier lower-ranked candidate.
- Add a bounded, localization-neutral overflow/aggregation summary to the existing Overview contract containing total candidate concerns, individually represented concerns, correlated duplicate views, and omitted concerns. It must be absent or canonically zero when no reduction occurred, internally consistent, validation-bounded, included in deterministic identity/JSON behavior, and explicitly rendered by the Console when any candidate was aggregated or omitted. The Console wording must state that additional concerns were summarized and must not imply an exhaustive list.
- Perform a documented compatibility analysis for the additive overflow summary. Preserve Presentation Model schema 1.0 only if old stored Overview values remain valid and existing decoders can safely ignore the additive field; otherwise apply the smallest coordinated version change across Presentation Model and Current Operator State compatibility. Do not create parallel model versions or migrations unrelated to this field.
- Preserve source traceability for every individually represented item. Overflow reporting shall expose counts and stable tokens only, never raw host values, paths, arbitrary errors, secrets, or identifiers.
- Add typed internal error classification at the application boundary sufficient to distinguish canonical pipeline/stage execution failure, operator projection/validation failure, and Current Operator State publication failure. Keep existing specific Inventory Store and Current State safety diagnostics.
- Map the typed categories to deterministic localized or localization-ready safe tokens such as `evaluation_pipeline_failed`, `operator_projection_failed`, and `current_state_publication_failed`. Do not classify with string matching and do not print wrapped raw errors.
- Ensure the discovered overflow condition succeeds after reduction; it must not be reported as a projection failure merely because valid candidates exceed the output limit.
- Correct recommendation semantics minimally: stale evidence may recommend a fresh `observe`; current intrinsic partial/unsupported evidence must recommend inspection rather than promise refresh repair; an explicit failed operation or non-self-healing application failure must recommend inspection/correction and must not recommend blind repetition. Keep recommendation ownership in `internal/presentationmodel` and localization in the Console catalog.
- Preserve all canonical Command, Inventory, Snapshot, Comparison, Drift, Health, Rule, Policy, Report, Runtime, Runtime Service, Current Operator State atomicity, freshness, Console load-only startup, and Guardian semantics.
- Add deterministic real-scale tests with at least 366 Rule and 366 Policy records, shuffled equivalent inputs, later urgent candidates, correlated and uncorrelated concerns, overflow, validation, publication, subprocess consumption, privacy, and resource bounds.

Expected primary targets are `internal/presentationmodel`, the narrow `cmd/qwsg` observe error boundary, `internal/operatorstate` only if additive Overview compatibility requires it, `internal/operatorconsole` rendering/localization, their tests, and directly affected documentation. No canonical evaluation engine is a target.



## Out of Scope


Task 036 shall not implement or redesign:

- Inventory collectors, Snapshot persistence, baseline selection, Comparison, Drift, Health, Rule, Policy, Report, Command, or Pipeline semantics;
- `observe` profile planning, bootstrap behavior, snapshot retention, or the order in which the canonical Snapshot stage persists a valid observation;
- a new attention, decision, aggregation, query, remediation, workflow, or diagnostic engine;
- general persistence, new Current Operator State history, incident history, audit history, trends, or a database;
- Operational Guardian Service, daemonization, monitoring, supervision, heartbeat, watchdog, Scheduler, Alert, Notification provider, transport, REST API, Dashboard, remote access, fleet operation, remediation, shell execution, or AI;
- changing `MaxAttention` upward as the fix, silently dropping overflow, printing all candidates, weakening Overview validation, or accepting invalid/untyped/uncorrelated evidence;
- changing Health taxonomy, Rule conditions, Policy outcomes, Report completeness, or suppressing valid upstream records merely to fit Presentation output;
- privileged collection, dependency or service installation, real user-state deletion, infrastructure mutation, packaging, deployment, release publication, staging, commit, push, fetch, branch, or tag operations.



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


### Task-creation state

- Verify the exact QWSG repository root, markers, Framework 1.x configuration, canonical HTTPS remote, `main` branch, HEAD, `0/0` relationship, complete Git status, empty index, ownership, modes, and ACLs.
- Require canonical idle with Task 035 as the unique latest completed archive/history pair, no active prompt, and no Task 036 prompt/history/archive collision.
- Run `bin/job --check`, `ai/scripts/next-task.sh --check`, `bin/job --check-test-tasks`, `ai/scripts/framework-check.sh`, Framework tests, Builder tests, configured engineering validations, full Go tests, race tests with writable temporary caches, vet, format, and Git diff checks.
- Preserve every pre-existing Owner-owned modified and untracked path exactly except this authorized preparation source.
- Confirm by source inspection and read-only stored-pair execution that all canonical stages complete, Health evidence is sufficient, Report is complete, Pipeline diagnostics are empty, and Current Operator State publication is not reached after projection failure.
- Confirm that `presentationmodel.builder` produces one Rule attention item per matched Rule result and one Policy attention item per observe Policy result, `addAttention` is unbounded before validation, `MaxAttention` is 256, and `Validate` returns `invalid operator overview` for the resulting 732 items.
- Confirm that `observeDiagnostic` preserves selected store/state errors but collapses all other pipeline, projection, and publication errors to `evaluation_failed`.
- Confirm Console startup only loads/requalifies Current Operator State and rendering is not the failure source.
- Confirm the old stale/partial recommendation is generated generically and cannot represent a known non-self-healing failure.
- Record the Release Minimalism decision: the defect is repaired inside existing Presentation Model, application diagnostic, Current State, and Console contracts; no new engine or layer is justified.
- Validate this source as a readable non-symlink UTF-8 regular file with no NUL, unfinished fence, placeholder, embedded approval protocol value, unresolved product choice, competing lifecycle, or ambiguous Builder field mapping.

### Separately authorized Builder installation

- Repeat canonical idle, Task 035 baseline, Task 036 destination absence, source hash/content, Framework, repository/Git, permission/ACL, Builder-interface, and lifecycle checks immediately before installation.
- Create and verify one bounded external mode-0700 Builder-installation snapshot covering exact Task 036 lifecycle destinations and absence, this source and hash, Task 035 completed baseline, repository identity, complete Git state, ownership/modes/ACLs, checksum manifest, retention, and exact restore instructions.
- Map only the owner-authored Builder fields into separate UTF-8 regular input files in a unique mode-0700 temporary directory. Supply the explicit approval protocol value separately; it must not occur in this source.
- Run `task-builder.sh --check-input` only for read-only validation during preparation. Install only after a later explicit Owner instruction.
- After installation require Task 036 as the sole active approved task, Task 035 as latest completed baseline, an empty index, preserved unrelated status, and all Builder/lifecycle/Framework checks passing; stop without implementation.

### Implementation starting state

- Start only through a later explicit canonical `job` invocation; read the complete installed prompt, history, skill, and Required Reading as data.
- Reconfirm the real-scale failure with deterministic fixtures and verify that any live evidence inspected is read-only or confined to explicit temporary roots.
- Freeze existing Pipeline results and verify no engine change is needed.
- Specify the exact candidate identity, Rule-to-Policy correlation key, global ranking tuple, overflow accounting invariants, schema compatibility decision, diagnostic error types, and recommendation truth table before code changes.
- Prove tests can inject large shuffled inputs, urgent candidates, projection failures, publication failures, clocks, state/store roots, and subprocess environments without real home state, host mutation, network, privilege, service, or sleep.
- Create and verify the proportional implementation snapshot before modifying any target.
- Stop on any material need to alter canonical engine semantics, raise the output bound, silently lose overflow, disclose host-sensitive data, weaken atomic publication, or create a new architecture layer.

The valid preparation/install baseline is canonical idle after Task 035. The valid post-install state is one active approved Task 036. The valid completion state is canonical idle with Task 036 as the unique latest completed archive/history pair and no successor.



## Snapshot Requirements


Task preparation modifies only `current-task-job.txt`. Record and retain its pre-edit SHA-256, exact Git status, ownership, mode, ACL, and Task 035 idle lifecycle evidence; no Builder installation or implementation snapshot is created during preparation.

Before separately authorized installation, create one unique external mode-0700 snapshot under `/tmp` covering `current-task-job.txt`, exact Task 036 prompt/history/archive destinations and verified absence, Task 035 archived prompt/history, repository identity, complete Git status, empty index, ownership, modes, ACLs, checksums, collision guards, retention, and bounded restore instructions. Verify all payload checksums and absence records before Builder execution.

Before implementation, create one unique external rollback snapshot of every existing directly affected Presentation Model, CLI/application diagnostic, Current State compatibility, Console rendering/catalog, test, README, English/Hungarian guide, Product Architecture, Functional Specification, Roadmap, System Map, architecture, prompt, history, and archive target. Record verified absence for new paths. Preserve exact working-tree bytes, Git state, ownership, modes, ACLs, hashes, collision guards, restore preconditions, and removal identities for created paths.

Snapshots exclude broad repository archives, caches, real user state, live host payloads, secrets, credentials, processes, services, and unrelated content. Retain installation and implementation snapshots through validation and Owner acceptance.



## Risk Assessment


- **Dropped critical evidence — high:** naive first-N truncation can hide a later urgent concern. Mitigate with full canonical global ranking before bounded emission and shuffled-order tests placing urgent items last.
- **False exhaustiveness — high:** a bounded list can look complete. Mitigate with validated explicit total/represented/correlated/omitted accounting and mandatory Console disclosure whenever reduction occurs.
- **Unsafe correlation — high:** merging unrelated Rule and Policy concerns can erase meaning. Correlate only through validated canonical `RuleEvaluationID`; never by text, order, or coincidental values.
- **Contract compatibility — high:** an Overview field or schema change can invalidate saved Current Operator State. Perform explicit additive compatibility analysis and coordinated fixture tests before selecting schema treatment.
- **False health or attention — high:** reduction must not change canonical Health/Rule/Policy outcomes. Derive overall severity before or from the globally ranked complete candidate set and retain urgent evidence.
- **Hidden failures — medium-high:** raw error collapse prevents action, while raw errors can leak data. Use typed boundary errors and fixed privacy-safe diagnostic tokens.
- **Misleading recommendation — medium-high:** blind retry can repeat deterministic failure. Separate refreshable stale evidence from intrinsic partial/unsupported and failed-operation states.
- **Resource amplification — medium:** valid large input can generate many candidates. Keep input processing proportional to existing bounded records, avoid output growth, and prove stable memory/output limits with oversized fixtures.
- **Dirty working tree collision — medium:** preserve Owner-owned work with exact target snapshots, targeted patches, empty index, and complete status review.

Overall risk is medium-high because operator prioritization and persisted Overview compatibility affect safety, but the fix is confined after successful canonical evaluation and introduces no host mutation or new authority.



## Planned Work


### Phase 1 — Freeze the regression and contract

- Complete starting verification and the implementation snapshot.
- Add a deterministic 366-record regression fixture reproducing 732 pre-correlation Rule/Policy candidates and the current `invalid operator overview` failure.
- Define candidate identity, validated Rule/Policy correlation, ranking, overflow accounting, compatibility, diagnostics, and recommendation truth table.

### Phase 2 — Bounded deterministic projection

- Refactor only the Presentation Model builder's attention collection/reduction path.
- Correlate validated downstream Policy concerns with their originating Rule concerns, rank the complete remaining candidate set globally, retain the highest-ranked bounded subset, and generate explicit overflow/aggregation accounting.
- Extend normalization, validation, canonical JSON/ID behavior, freshness requalification, and fixtures only as required by the additive summary.

### Phase 3 — Diagnostics and recommendations

- Add small typed application-boundary wrappers for pipeline execution, operator projection, and Current State publication failures.
- Map them to fixed privacy-safe operator tokens while preserving existing specific store/state diagnostics.
- Correct stale/current-partial/unsupported/failed-operation recommendation cases and localized Console wording without introducing remediation logic.

### Phase 4 — Product-scale acceptance

- Exercise baseline then full observe using 366+ correlated Rule/Policy results and at least one late urgent concern.
- Prove successful validated Overview creation, atomic Current Operator State publication, separate-process Console rendering, explicit summarized-count disclosure, non-unknown qualified condition where evidence permits, no false health, and Guardian not observed.
- Prove shuffled equivalent inputs yield identical retained concerns, accounting, Overview ID, Current State content, and Console output.

### Phase 5 — Document and close

- Update directly affected model, diagnostic, Console, product, and user documentation.
- Run all focused/full/race/vet/format, Framework, Builder, lifecycle, privacy, resource, compatibility, rollback, and Git validations.
- Finalize Task 036 history, archive without a successor, and verify canonical idle.



## Rollback Plan


Rollback is exact, file-bounded, identity-checked, and requires Owner confirmation before overwriting material post-snapshot work.

Record failure evidence and stop test processes. Verify snapshot manifests, hashes, repository identity, exact target list, lifecycle identity, and absence of later Owner edits. Restore only pre-existing target payloads with recorded modes and ACLs where authorized. Remove only Task 036-created paths whose pre-change absence and current Task 036 identity are both proven. Never use wildcard deletion, recursive cleanup, broad reset, clean, checkout, restore, or replacement from HEAD.

Remove temporary test state only by exact validated path under a Task 036 temporary root. Never delete or rewrite real Inventory Store or Current Operator State. If acceptance uses live collection, it must write only to explicit temporary roots and retain or remove those roots only according to recorded exact identities.

After rollback verify target checksums/presence, ownership, modes, ACLs, complete Git status, empty index, Framework/lifecycle/diverted-task checks, original Presentation limits and diagnostics, existing profile compatibility, full Go tests/race/vet/format, and Git diff checks. Retain snapshot, failed-work diff, restore log, and validation report.



## Deliverables


- deterministic bounded attention projection within the existing Canonical Operator Presentation Model;
- validated Rule-to-Policy correlation using existing canonical identities, without Rule or Policy semantic changes;
- explicit bounded overflow/aggregation accounting and non-exhaustive Console disclosure;
- global severity/importance ordering that preserves later urgent concerns and is independent of input enumeration;
- compatible Overview and Current Operator State handling for the additive summary;
- typed privacy-safe diagnostics separating pipeline, projection, and Current State publication failures;
- truthful recommendation behavior for stale, current-partial/unsupported, and failed non-self-healing states;
- realistic 366+ record unit, integration, shuffle, publication, subprocess, Console, privacy, and resource regression coverage;
- directly affected architecture, product, README, English/Hungarian user, Roadmap, System Map, Engineering History, prompt/history/archive, validation, and rollback evidence.

No engine, new layer, service, monitor, provider, API, Dashboard, remediation, dependency, installation, package, deployment, stage, commit, or push is a deliverable.



## Verification


- Prove the defect boundary with an exact test: Pipeline execution is complete with eight ordered complete typed stages, Health evidence sufficient, Report complete, and projection alone exceeds the old attention output contract.
- Prove Rule/Policy correlation accepts only validated `RuleEvaluationID` relationships, reduces each correlated pair once, retains Policy as the operator-facing final decision, and leaves uncorrelated concerns separate and traceable.
- Test at least 366 Rule and 366 Policy results, candidate count above 732 when late urgent and direct concerns are included, and output attention count at or below 256.
- Shuffle Health, Rule, Policy, and Report record enumeration repeatedly with a fixed seed and prove identical retained items, aggregation counts, recommendation order, Overview ID, canonical JSON, and Console output.
- Place urgent items after more than 256 review items and prove every globally higher-ranked urgent item is retained. Test deterministic tie-breaking at the exact bound.
- Validate overflow accounting invariants: total candidates equal represented plus correlated duplicate views plus omitted concerns according to the documented model; all counts are nonnegative/bounded; disclosure appears exactly when correlation or omission occurs; Overview never implies exhaustiveness during overflow.
- Prove every represented attention item has a valid bounded SourceReference and no overflow summary contains raw values, host identifiers, paths, environment contents, arbitrary errors, or secrets.
- Prove `presentationmodel.Validate`, deterministic identity, serialization, old-state compatibility, freshness requalification, `operatorstate.Normalize`, atomic Publish/Load, and Console render all accept the hardened Overview.
- Prove pipeline/stage, projection/validation, and Current State publication failures map to distinct fixed safe diagnostics; existing corrupt/incompatible/unsafe/permission diagnostics remain distinct; raw wrapped errors never reach terminal or persisted diagnostic text.
- Test recommendation truth table: stale otherwise-valid evidence recommends refresh; current partial/unsupported evidence recommends inspection without a repair promise; failed/non-self-healing operation recommends inspection/correction without blind retry; healthy/current complete evidence does not gain a false action.
- Product subprocess acceptance: process A establishes a fresh baseline; process B performs a full 366+ live-like observe whose pre-reduction candidates exceed 256 and publishes Current Operator State; process C starts bare `qwsg`, performs no collection, and displays the qualified state plus explicit summarized-attention disclosure. Assert no unavailable/unknown result when complete sufficient evidence supports a qualified condition, no false healthy claim, and Guardian remains not observed.
- Run focused tests for `internal/presentationmodel`, `cmd/qwsg`, `internal/operatorstate`, and `internal/operatorconsole`; `make build`; full `go test ./...`; repository-wide `go test -race ./...` with writable temporary caches; `make vet`; `make fmt-check`; Framework validation/tests; Builder tests; configured engineering validations; lifecycle checks; diverted-task audit; Git diff checks.
- Audit imports and source ownership: canonical engines remain unchanged, Console consumes only validated Overview, Current State remains presentation-independent storage, and no duplicate Pipeline, decision engine, remediation engine, monitoring, service, provider, or persistence layer was added.
- Verify exact changed targets, ownership, modes, ACLs, empty index, complete Git status, preservation of unrelated Owner content, snapshot checksums, bounded restore feasibility, and that nothing was installed, transmitted, staged, committed, pushed, packaged, deployed, or released.

All automated acceptance uses deterministic fixtures and explicit temporary state/store roots. It requires no real home state, host mutation, network, credentials, privilege, service, or sleep.



## Documentation Updates


Update only directly affected sections of:

- `docs/architecture/CANONICAL_OPERATOR_PRESENTATION_MODEL.md` for candidate correlation, global ordering, output bound, overflow accounting, traceability, compatibility, and recommendation semantics;
- `docs/architecture/CANONICAL_CURRENT_OPERATOR_STATE.md` only for persisted additive Overview compatibility and atomic publication behavior;
- `docs/architecture/INTERACTIVE_OPERATOR_CONSOLE.md` for summarized-attention disclosure and corrected recommendations;
- `docs/architecture/CANONICAL_COMMAND_ARCHITECTURE.md` only for privacy-safe `observe` failure categories, without Command or Pipeline changes;
- `docs/PRODUCT_ARCHITECTURE.md`, `docs/FUNCTIONAL_SPECIFICATION.md`, `README.md`, and directly affected English/Hungarian Console guidance for large valid observations and diagnostic meanings;
- `ai/core/04_ARCHITECTURE.md`, `ai/core/05_SYSTEM_MAP.md`, `ai/core/07_ENGINEERING_HISTORY.md`, and `ai/core/13_ROADMAP.md` to record the regression fix and Version 1.0 gate status;
- active prompt, independent Task 036 history, and completed archive.

Documentation must state that attention is a deterministic prioritized bounded operator view, disclose when additional concerns are aggregated or omitted, distinguish canonical evaluation success from projection/publication failure, and avoid claiming that a retry repairs a deterministic application defect. Record actual updates and justified omissions in Task 036 history.



## Completion Criteria


Task 036 is complete only when:

- the verified 366-record canonical evaluation remains unchanged and complete;
- more than 256 valid candidate concerns no longer cause `invalid operator overview`;
- `MaxAttention` is not raised, emitted attention is always bounded, and arbitrary valid input cannot grow Overview output beyond contract limits;
- globally ranked reduction always retains higher-severity/higher-importance concerns regardless of enumeration order, with deterministic tie-breaking and stable Overview identity;
- Rule/Policy duplicates are correlated only by validated canonical identity, downstream Policy information is retained, unrelated concerns are not merged, and represented sources remain traceable;
- aggregation/omission counts are validated, bounded, deterministic, persisted, and visibly rendered so the operator is never told the attention list is exhaustive when it is not;
- schema/Current State compatibility is explicitly proven or minimally and coherently versioned;
- pipeline/stage, projection/validation, and Current State publication failures have distinct privacy-safe operator diagnostics, while raw errors and host-sensitive content remain hidden;
- recommendation behavior distinguishes refreshable stale evidence from current partial/unsupported and failed non-self-healing conditions without adding remediation logic;
- the large subprocess acceptance completes baseline, full observe, projection, publication, and new-process Console consumption with truthful condition, attention, changes, Alerts, evidence, recommendation, overflow disclosure, and Guardian not observed;
- focused/full/race/vet/format, Framework, Builder, lifecycle, compatibility, privacy, resource, documentation, Git, snapshot, and rollback validations pass;
- no canonical engine semantics, service, host state, dependency, infrastructure, package, deployment, stage, commit, or push changed;
- Task 036 prompt/history are complete and archived without a successor, and `bin/job --check` reports canonical idle with Task 036 as the unique latest completed baseline.

A valid result is `complete`, `complete with disclosed limitations`, or `blocked`. Completion may not be claimed while bounded prioritization, overflow truthfulness, source traceability, diagnostic differentiation, recommendation correctness, real-scale publication, separate-process visibility, rollback, or lifecycle evidence remains unresolved.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-10 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
