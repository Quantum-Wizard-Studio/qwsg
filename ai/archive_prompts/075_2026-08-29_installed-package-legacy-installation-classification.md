# Current Engineering Task 075: Installed Package and Legacy Installation Classification

## Task Metadata

- Task ID: `075`
- Task slug: `installed-package-legacy-installation-classification`
- Status: `complete`
- Date opened: `2026-08-29` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

Installed Package and Legacy Installation Classification


## Objective

Implement the canonical QWSG installed-package and legacy-installation classification model required by the accepted Task 074 architecture. Determine installation state from verifiable local QWSG package evidence rather than binary presence, eliminate the server.quantumwizard.hu 0.0.1-prealpha guided-installer defect, preserve supported installation/update paths, and fail safely for legacy, unknown, incomplete or inconsistent artifacts.


## Scope

Inspect current installer, updater, package verification, provenance, version, migration, configuration, state and guided-installation code. Implement one deterministic fail-closed classifier using the strongest trustworthy existing evidence. Distinguish no installation, verified supported installation, supported upgrade source, legacy installation, unknown/unverified installation and inconsistent/incomplete installation. Return safe structured reasons, integrate only the existing installer/update decision paths that use weak evidence, add deterministic fixtures/tests including the 0.0.1-prealpha regression, preserve update/migration/rollback, and update required architecture/user/engineering/task documentation.


## Out of Scope

No updater redesign or second updater; no network release discovery, qwsg.release-index/1, Guardian update scheduling, update notification/deduplication, unattended updates, external-host mutation, infrastructure/deployment/publication, QWSG 1.2.0 release/tag/Forgejo asset mutation, unrelated refactor or feature, or automatic Task 076 start.


## Authority Envelope

**Task targets and boundaries:** Inspect required repository components; modify only canonical installation classification, strictly necessary installer/update/guided-install consumers, tests/fixtures, narrow deduplication refactors and synchronized documentation. Preserve existing updater, migration, rollback and released 1.2.0 identities.
**Permitted external actions:** Read-only canonical Git synchronization and public release-identity verification only; no VPS, production service, infrastructure or release mutation.
**Owner-reserved decisions:** Updater redesign, release discovery/manifest, Guardian scheduling, update notification, unattended updates, external-host work, publication, release/tag/assets, Task 076 scope and all material scope expansion.
**Task-specific STOP conditions:** Stop on baseline/release identity mismatch, unavailable rollback, credential/privacy risk, need to overwrite unknown installation silently, unsafe external mutation, or unresolved evidence ambiguity that cannot fail closed.


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

Verify exact QWSG root, branch main, HEAD and origin/main both 19571803edb8d2f4bf47688ce277cc574a216d9b, divergence 0/0, clean tree/index and canonical idle after Owner-accepted Task 074. Verify Framework/lifecycle and protected QWSG 1.2.0 VERSION, annotated v1.2.0 tag, release commit, frozen artifact/checksum and public release evidence before production modification. Diagnose and stop safely on an unauthorized material difference.


## Snapshot Requirements

Retain verified Builder-installation snapshot /tmp/qwsg-task075-builder-install.Jqyr6l. Before production modification create a unique external mode-0700 Task 075 execution snapshot with exact working-tree payloads or absence records for every intended target, Git/lifecycle/release identities, ownership/modes/ACLs, tool identity, checksums, readable archive/bundle and literal bounded rollback instructions. Exclude credentials, runtime/private host data, caches and unrelated content; retain evidence through Owner review.


## Risk Assessment

High classification/data-loss risk: arbitrary or spoofed binaries, stale configuration/state and incomplete layouts must not establish trust or be overwritten; mitigate with explicit evidence precedence and fail-closed states. Medium compatibility risk: valid 1.2.0 and supported upgrade sources must remain recognized; mitigate with existing provenance/migration contracts and regression tests. Medium divergence risk: one classifier and shared consumers only. Medium privacy risk: expose bounded reason tokens, never paths, configuration, credentials or host identity. Low rollback risk: exact snapshots, targeted edits and preserved release identities.


## Planned Work

Verify baseline and snapshot; inventory every installed/version/support decision; identify weak binary-presence checks; define canonical states/evidence precedence; implement one classifier using existing provenance/package mechanisms; integrate installer/update/guided-install consumers without duplication; add clean/current/older/legacy/arbitrary/version-only/partial/inconsistent/upgrade/unsupported/malformed/stale-state/mixed-artifact tests and the server.quantumwizard.hu regression; verify update/migration/rollback; update documentation; run complete validation; integrate and canonically close without starting Task 076.


## Rollback Plan

Verify snapshot checksums, repository/release identities, exact target list and absence of later Owner edits. Restore only pre-existing Task 075 targets from exact isolated snapshot payloads with recorded metadata; remove only Task 075-created files whose prior absence and current identity are proven. Never use broad reset, checkout, restore, clean, wildcard deletion, history rewrite or release-ref mutation. Re-run focused/full/framework/lifecycle/Git/release checks after any authorized rollback.


## Deliverables

Canonical installed-package/legacy classifier and safe structured evidence; documented states/evidence precedence; existing installer/update/guided-install integration; deterministic boundary fixtures/tests and 0.0.1-prealpha regression; preserved supported migration/update/rollback behavior; synchronized architecture/user/roadmap/engineering/task documentation; exact Owner completion report covering all nineteen requested items.


## Verification

Run active/idle lifecycle, Framework, Builder and diversion assertions; focused classifier/installer/update/migration/rollback tests; gofmt verification, go vet ./..., go test ./..., go test -race ./...; deterministic build/reproducibility and applicable release checks; git diff --check; targeted staged scope/privacy/secret/mode/ACL review; explicit clean/current/legacy/unknown/inconsistent/upgrade/guided regression/no-overwrite proofs; final clean HEAD/origin relation; unchanged v1.2.0 tag, commit, artifact, checksum and published Release evidence. No external VPS mutation is required.


## Documentation Updates

Update canonical installation classification architecture, evidence hierarchy, legacy and unknown/inconsistent handling, update-discovery and migration relationships, fail-closed security/privacy behavior and server.quantumwizard.hu regression protection. Update required user/operator guidance, Architecture Governance/System Map/Roadmap/Engineering History and detailed Task 075 history according to repository evidence; keep engineering artifacts English and user-facing content localized.


## Completion Criteria

One canonical evidence-based classifier exists and all installer/update decision consumers use it; binary presence alone cannot establish verified installation; 0.0.1-prealpha regression passes; valid/current/supported-upgrade installations remain usable; unknown/inconsistent/unsafe states fail safely without silent overwrite; required tests/validation/docs/rollback pass; repository closes clean and synchronized; QWSG 1.2.0 release identities remain untouched; no network discovery or Task 076 work is implemented or started.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-29 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
