# Current Engineering Task 074: Update Discovery and Release Awareness Architecture

## Task Metadata

- Task ID: `074`
- Task slug: `update-discovery-release-awareness-architecture`
- Status: `complete`
- Date opened: `2026-08-29` UTC
- Human authority: Attila — Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Update Discovery and Release Awareness Architecture


## Objective

Design and document the repository-backed QWSG 1.3 Update Discovery and Release Awareness architecture only, preserving read-only discovery and explicit operator-controlled installation.


## Scope

Inventory actual QWSG 1.2.0 version/release/update/migration/rollback/Guardian/notification/SMTP/readiness/configuration/localization/state/security implementations; document reuse, gaps, boundaries, versioned release contract and trust; design CLI, Guardian, notification, failure, privacy, compatibility and Community/Pro models; incorporate all three supplied field findings; update canonical architecture, roadmap/system map/index, Engineering History and Task 074 history; validate and prove v1.2.0 identities unchanged.


## Out of Scope

No production feature, network checker, release service, updater, unattended install, mandatory telemetry/registration, Pro implementation, infrastructure/runtime mutation, release/tag/artifact change, deployment/publication, unrelated refactor, or automatic later task.


## Authority Envelope

**Task targets and boundaries:** Repository inspection and English architecture/documentation/history changes for Task 074 only; production behavior, released 1.2.0 identities, external infrastructure and implementation tasks are excluded.
**Permitted external actions:** Read-only canonical Git fetch and public release/tag identity verification only; no remote mutation.
**Owner-reserved decisions:** Implementation tasks, signing policy, release service, Pro scope, publication, tags, Releases, infrastructure and rollout.
**Task-specific STOP conditions:** Release identity mismatch, unresolvable lifecycle/rollback failure, credential/privacy risk, unexpected divergence, or required external mutation.


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

Verify root, remote, main, exact HEAD/status/index and fresh origin/main relation; idle Task 073 before Builder and active Task 074 after; QWSG 1.2.0 version, annotated v1.2.0 tag, release commit 348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2, artifact name and stated SHA256; inspect repository implementation without assumptions.


## Snapshot Requirements

Retain verified mode-0700 Builder snapshot /tmp/qwsg-task074-builder-install.AvAmIG. Before documentation edits create a unique mode-0700 /tmp execution snapshot with exact target payloads/absence, Git/lifecycle/release identities, metadata, checksums, readable bundle/archive and bounded restore instructions; exclude secrets, runtime data, caches and unrelated content; retain through Owner review.


## Risk Assessment

Medium documentation/design/security/privacy/compatibility risk mitigated by repository evidence, versioned fail-closed contracts, data minimization, separation of authenticity and integrity, existing subsystem reuse and diff review. Low rollback/data-loss risk mitigated by exact external snapshots and bounded recovery. Low-to-medium localization/operational risk mitigated by existing localized notification architecture and Guardian failure isolation.


## Planned Work

Validate lifecycle/repository/release/snapshots; read governance and relevant 1.2.0 code/docs/tests; inventory and analyze gaps; design discovery/manifest/trust/state/failure/privacy/Guardian/notification/install boundaries; incorporate field findings and Community/Pro separation; update canonical documentation/history only; run full proportional validation, integrate, close lifecycle, and stop for Owner review.


## Rollback Plan

Preserve diff/evidence; verify snapshots, repository and exact identities and no later Owner edits; restore only exact pre-existing Task 074 targets and remove only created files with proven absence/current identity; never broad reset/checkout/restore/clean, wildcard delete, rewrite history or mutate tags; revalidate lifecycle/framework/docs/Git/release identities.


## Deliverables

Canonical QWSG 1.3 Update Discovery and Release Awareness architecture; applicable roadmap/system map/index updates; concise Engineering History and complete Task 074 history; Hungarian Owner report with exact Git states/files/decisions/reuse/gaps/field findings/security/privacy/manifest/Guardian/notification/Community-Pro/tests/snapshots/nonimplementation/immutability/recommended sequence.


## Verification

Run bin/job lifecycle/Builder checks, documentation/framework suites, gofmt verification, go vet ./..., go test ./..., canonical race/build checks where applicable, git diff --check, target/staged/permissions/ACL/privacy/secret/scope review; prove no production behavior changed and v1.2.0 tag target, release commit, artifact name and SHA256 remain unchanged.


## Documentation Updates

Create/update canonical Update Discovery architecture; update applicable Product Architecture, Roadmap and System Map/index; update concise Engineering History and detailed Task 074 history; keep engineering artifacts English and user-facing concepts localization-ready.


## Completion Criteria

All requested repository-backed architecture sections and field findings are complete; discovery is read-only and install operator-controlled; trust, isolation, deduplication, privacy, compatibility and Community/Pro boundaries are explicit; validation passes and rollback remains usable; no production/external mutation occurred; v1.2.0 identities are untouched; lifecycle closes to idle without starting implementation.


## Owner Approval Requirements

Approved by Attila — Project Owner through the Engineering Task Builder on 2026-08-29 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
