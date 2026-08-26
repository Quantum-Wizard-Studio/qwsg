# Current Engineering Task 063: QWSG 1.1.0 Final Release

## Task Metadata

- Task ID: `063`
- Task slug: `qwsg-1-1-0-final-release`
- Status: `complete`
- Date opened: `2026-08-26` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG 1.1.0 Final Release


## Objective

Promote the externally accepted QWSG 1.1.0-rc.6 product lineage into the final QWSG 1.1.0 release through the shortest safe Framework 2.0 workflow. Make only final release-identity, documentation and deterministic release-plumbing changes unless source analysis proves another minimal non-behavioral requirement; validate and commit the exact release-source state; construct and freeze two byte-identical deterministic linux-amd64 outputs; reuse trustworthy Task 061 external behavioral acceptance unless a relevant behavior-changing mutation invalidates it; and, only after every mandatory release gate passes, create and verify tag v1.1.0, publish one final Forgejo Release with the exact frozen archive and checksum sidecar, verify public downloads and integrity, record the authoritative release identity, integrate evidence, and close Task 063 canonically.


## Scope

- Verify canonical idle, Framework 2.0.0, exact Git/ref state, Task 061 acceptance evidence, Task 062 framework-only changes, accepted RC.6 behavioral source ancestry and absence of unresolved product defects.
- Create and verify private pre-change and pre-publication rollback snapshots.
- Determine and implement the minimal final-release diff from `1.1.0-rc.6` to `1.1.0`, normally limited to `VERSION`, final release notes/acceptance/changelog/install references, deterministic release identity validation, release-plumbing tests and Task 063 lifecycle/evidence records.
- Permit iterative in-scope diagnosis, correction, refactoring and proportional retesting under Framework 2.0 Standard Execution Authority.
- Create one exact release-source commit descended from the current canonical main lineage. Prove that its behavior-bearing QWSG source, service unit, installer/uninstaller and runtime configuration are unchanged from accepted RC.6 except for an explicitly classified and validated release-identity/plumbing change.
- Construct two isolated deterministic final linux-amd64 outputs from separate `git archive` exports of the exact release-source commit using one explicit commit-derived `SOURCE_DATE_EPOCH` and full `BUILD_COMMIT`.
- Select and freeze one byte-identical final archive and checksum sidecar; verify archive safety, package layout, manifest, binary identity/provenance, documentation, modes, static linkage and source determinism.
- Reuse Task 061 practical acceptance for unchanged behavior. If a relevant Task 063 mutation affects runtime behavior, classify it and perform only the acceptance materially invalidated by that mutation; do not repeat Task 061 wholesale.
- Produce the release-readiness verdict. Conditional on every mandatory gate passing, create annotated tag `v1.1.0` pointing exactly to the release-source commit, push and verify it, create the final non-draft/non-prerelease Forgejo Release `QWSG 1.1.0`, attach exactly the frozen archive and sidecar, and verify versioned public downloads and hashes using clean `wget` and `curl -fLO` locations.
- Update canonical release evidence, release documentation, engineering history and Task 063 history; perform explicit-path commits, clean-fast-forward pushes, direct Forgejo verification and canonical idle closure.


## Out of Scope

- Do not change QWSG runtime/product behavior merely to perform the final release; do not invent new acceptance requirements or repeat Task 061 practical acceptance without a relevant invalidating mutation.
- Do not access the existing test VPS unless a behavior-changing Task 063 mutation truthfully invalidates specific external evidence and the installed Authority Envelope permits the narrowly affected acceptance. Routine final-release identity/plumbing work does not authorize VPS access.
- Do not expose or store credentials, tokens, private keys, private provider identities or protected addresses in chat, Git, logs, task evidence, command arguments where avoidable, artifacts or release metadata.
- Do not publish if any mandatory provenance, deterministic-build, package, security, validation, acceptance-reuse, tag or distribution gate fails.
- Do not force push, rewrite history, move an existing tag, replace frozen bytes, upload extra/unverified assets, use a mutable `latest` URL, deploy QWSG to production infrastructure, or announce through unrelated channels.
- Do not change product architecture or scope and do not create or prepare Task 064.


## Authority Envelope

- **Task targets and boundaries:** Framework 2.0 Standard Execution Authority applies to the exact Task 063 objective and to QWSG final-release identity/documentation/plumbing, release-source integration, deterministic build/verification, Task 061 evidence reconciliation, release evidence and lifecycle files. Ordinary reversible inspect/edit/diagnose/classify/fix/test/retest/refactor/build/document/integrate/rollback cycles are authorized. Product behavior and architecture remain unchanged unless an unexpected difference is classified and remains within the narrow final-release objective.
- **Permitted external actions:** read-only canonical Git/Forgejo verification; clean-fast-forward pushes of reviewed Task 063 commits; and, only after every preceding mandatory release gate is explicitly PASS, creation and push of the single annotated tag `v1.1.0`, creation of the final non-draft/non-prerelease Forgejo Release `QWSG 1.1.0`, upload of exactly the frozen `qwsg-1.1.0-linux-amd64.tar.gz` and matching `.sha256` sidecar, and anonymous public `wget`/`curl -fLO` verification of those exact assets. Existing protected Forgejo authentication may be used without disclosure; any inherently Owner-only credential entry is an interaction under the same authority, not a new engineering approval gate. No VPS access is permitted unless a behavior-changing mutation invalidates narrowly identified Task 061 evidence.
- **Owner-reserved decisions:** material product/architecture/scope change; credential disclosure or replacement; unplanned infrastructure mutation; destructive or irreversible external action without rollback; force push/history rewrite; tag replacement/deletion; publication despite a failed gate; extra assets, signing/trust changes, deployment, announcement outside the Forgejo Release itself; and Task 064.
- **Task-specific STOP conditions:** STOP for a security/privacy regression; credential exposure risk; unavailable required rollback; material runtime/product behavior change needing an Owner product decision; untrustworthy source or acceptance lineage; unexplained or unresolved deterministic artifact mismatch; unsafe archive/package or provenance mismatch; inability to reuse Task 061 truthfully; failed corrected Guardian behavior or another critical accepted behavior; tag/release/asset identity collision; inability to verify the authoritative pushed source/tag/assets; a mandatory release gate that cannot be corrected safely within scope; or any work outside this envelope. Ordinary editing, build, test, packaging, checksum, documentation or integration failures enter diagnose -> classify -> correct -> retest and are not STOP conditions by themselves.


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

- Verify ordinary-user execution at the canonical repository root on branch `main`; Framework `2.0.0` VALID; canonical lifecycle idle after completed Task 062; zero active prompts; empty index; clean tracked worktree; and only the unrelated excluded Owner blueprint present and untouched.
- Fetch and prove `HEAD == origin/main ==` direct Forgejo `refs/heads/main` at Task 062 closure commit `1f26b7c16e542fdf72614269d1039bbf436d6bbf`, with no divergence.
- Verify accepted RC.6 behavioral source commit `25a30718bc92882e9773a5c405ad648c0eee1a81` is an ancestor; Task 061 evidence commit `6677e5e6c5fa515412302dde8fd9e76060b07fa4` and closure commit `d3ee69bc3bbbc012eaaff88100764404a5856313` are present; later commits through Task 062 change governance/evidence only and no behavior-bearing QWSG source or packaging.
- Verify `VERSION=1.1.0-rc.6`, Task 061 targeted practical acceptance PASS, QWSG-059-F001 externally corrected, RC.6 product defects NONE, READY FOR RELEASE DECISION, and the recorded candidate/source identities.
- Verify tag `v1.1.0`, a Forgejo Release for it, final 1.1.0 assets, `docs/release/RELEASE_NOTES_1.1.0.md`, and final acceptance evidence do not already exist. Treat any collision as a boundary requiring read-only diagnosis before mutation.
- Run the Framework, Builder, lifecycle, test-task, release-plumbing, Go test/race/vet/format, shell syntax, systemd static, Git and privacy baselines proportionally before modification.


## Snapshot Requirements

Before modifying tracked Task 063 targets, create one unique private mode-0700 snapshot under `/tmp`. Include a readable exact tracked-HEAD archive; literal before-images/hashes/modes of release identity, build/plumbing, documentation and lifecycle targets; Git/ref/tag/Release-absence evidence; Task 061 acceptance and candidate-identity evidence; exact product-diff lineage; the excluded Owner-path record without reading its content; and guarded collision-aware rollback instructions. Exclude credentials, candidate bytes, caches and unrelated content. Before tag/publication create a second private release-boundary snapshot containing the exact release-source commit/tree, validated frozen artifact identities, proposed tag identity, publication gate ledger and rollback/containment actions. Verify snapshot readability, hashes, permissions, exclusions and restoration constraints.


## Risk Assessment

- Wrong source or tag lineage — critical: prove accepted RC.6 ancestry, exact release-source tree, commit-derived provenance and annotated-tag peel before publication.
- Accidental behavior change invalidates acceptance reuse — critical: compare behavior-bearing paths against accepted RC.6, classify every difference, and rerun only materially invalidated evidence; STOP on an unexpected product decision boundary.
- Nondeterministic or mutable artifacts — critical: isolated twin builds, byte comparison, private freeze, immutable hashes and no post-freeze rebuild/substitution.
- Release builder accepts weak final metadata — high: extend 1.1.0 final identity with the same explicit epoch/full lowercase commit requirements as RC.6 and add negative/positive plumbing tests.
- Unsafe archive or incorrect documentation — high: verify types, paths, unique root, manifest, modes, LICENSE, versioned notes, binary provenance and documentation consistency.
- Premature or incorrect publication — critical: tag/Release/assets are conditionally authorized only after a recorded all-PASS gate; collision and remote identity checks precede mutation.
- Credential/privacy exposure — critical: use protected authentication without retained secrets or private identities; redact evidence and avoid command-line secrets.
- Distribution mismatch — critical: upload only frozen bytes and verify clean anonymous versioned downloads with both wget and curl plus SHA-256/manifest checks.
- Git/lifecycle corruption — critical: explicit staging, reviewed commits, clean-fast-forward only, direct refs, snapshots and canonical idle validation.


## Planned Work

1. Validate canonical idle, Framework 2.0.0, Git/Forgejo equality, absent final-release objects, exact RC.6 source/evidence ancestry and Task 061 acceptance conclusions; create the verified pre-change snapshot.
2. Diff accepted RC.6 behavior-bearing paths through current HEAD and define the minimal final-release target set. Preserve all trustworthy Task 061 evidence not invalidated by a relevant mutation.
3. Change release identity from `1.1.0-rc.6` to `1.1.0`; add final release notes and acceptance/evidence structure; move changelog Unreleased material into 1.1.0 as appropriate; update installation/download references; and extend deterministic build/plumbing tests so final 1.1.0 requires explicit full source provenance exactly as RC.6 does.
4. Iterate diagnose/classify/correct/test within Task 063 until focused and proportional full validation pass. Do not create another task for ordinary correction.
5. Review and explicitly stage only the release-source target set, create and clean-fast-forward push one exact release-source commit, and verify its tree/ref identity. Derive one `SOURCE_DATE_EPOCH` from that commit.
6. Export the exact release-source commit twice with `git archive` into independent private roots and use isolated output/cache locations to build QWSG 1.1.0 twice. Require byte-identical binaries, manifests, archives and checksum sidecars.
7. Select exactly one output, freeze archive and sidecar read-only, record name/size/SHA-256, manifest and binary hashes, build epoch/time and full source commit, and forbid byte replacement.
8. Verify sidecar; safe unique single-root archive; only regular files/directories; manifest content; required LICENSE/docs/final notes; expected file modes; static linux-amd64 binary; exact `QWSG 1.1.0`, commit and build-time provenance; extracted Smart Install readiness; installer/uninstaller packaging contracts; and absence of ambient VCS metadata.
9. Reconcile Task 061 evidence. If the release-source diff is non-behavioral, reuse its external PASS and limitations without VPS access. If a relevant behavior change exists, run only the materially invalidated acceptance if authorized; otherwise STOP at the product-decision boundary.
10. Produce an explicit all-gates release ledger and verdict. If any mandatory gate is not PASS, do not tag or publish; diagnose/correct within scope or STOP only at a genuine boundary.
11. After all gates PASS, create annotated `v1.1.0` pointing exactly to the release-source commit, verify the local tag object/peel/message, push only that tag and verify Forgejo. Create final Forgejo Release `QWSG 1.1.0`, attach exactly the frozen archive and sidecar, and verify release state and asset identities.
12. Download both versioned assets anonymously using clean independent wget and curl locations; verify filenames, sizes, archive/sidecar hashes, `sha256sum -c`, internal manifest, layout and binary provenance against the frozen ledger.
13. Record final release evidence and authoritative URL/identities without secrets; explicitly stage and push the evidence/history commit; verify HEAD/origin/main/direct Forgejo and tag/Release/assets; archive Task 063 and return to canonical idle without creating Task 064.


## Rollback Plan

- Before the release-source commit, restore only literal Task 063 targets from the verified pre-change snapshot after identity and collision checks; remove only new files proven absent and Task 063-owned. Never use broad reset, checkout or clean.
- During ordinary iteration, use narrow patches or verified before-images and rerun affected checks.
- After a pushed release-source commit but before tag publication, use a forward correction commit if necessary and rebuild only before freeze; do not rewrite published history.
- After candidate freeze, any required byte/source correction invalidates the frozen output and returns to the in-task development/build loop only if no tag/Release exists. Record invalidation explicitly; never silently replace bytes.
- Once `v1.1.0` or the final Release is public, do not move/delete the tag, overwrite assets, hide failures or force push. Contain publication, preserve evidence and STOP for Owner direction if authoritative identity cannot be completed safely. The release-boundary snapshot provides exact proposed identities and containment steps, not authority to rewrite public history.
- Verify Framework/lifecycle/Git/product/release state after every rollback or containment action.


## Deliverables

- One exact final release-source commit for QWSG 1.1.0, descended from the accepted RC.6 lineage, with only justified release-identity/documentation/plumbing changes and no unclassified behavior change.
- Two isolated byte-identical deterministic builds and one frozen `qwsg-1.1.0-linux-amd64.tar.gz` plus `qwsg-1.1.0-linux-amd64.tar.gz.sha256` with complete privacy-safe identity/provenance ledger.
- Final `docs/release/RELEASE_NOTES_1.1.0.md`, `docs/release/ACCEPTANCE_1.1.0.md`, changelog/install/distribution consistency and updated release-plumbing regression coverage.
- Explicit Task 061 evidence-reuse matrix and final release-readiness verdict with limitations preserved truthfully.
- Conditional all-PASS publication result: annotated tag `v1.1.0`, final Forgejo Release `QWSG 1.1.0`, exactly two verified frozen assets, and verified immutable public download URLs. If a mandatory gate fails, a truthful NOT RELEASED result and defect/boundary evidence instead.
- Task 063 history, snapshots, validation, commits, push/ref/tag/Release/download evidence and canonical lifecycle closure without Task 064.


## Verification

- Framework 2.0.0, Builder, lifecycle, one-active-task/test-task, repository configuration, shell syntax and Git safety validations PASS.
- Exact ancestry and path-diff proof shows accepted behavioral source `25a30718bc92882e9773a5c405ad648c0eee1a81`, Task 061 acceptance commit `6677e5e6c5fa515412302dde8fd9e76060b07fa4`, Task 061 closure `d3ee69bc3bbbc012eaaff88100764404a5856313`, Task 062 closure/current baseline `1f26b7c16e542fdf72614269d1039bbf436d6bbf`, and the final release-source commit form one trustworthy lineage.
- Every behavior-bearing diff from accepted RC.6 is absent or explicitly classified; Task 061 external acceptance reuse is valid and no stronger acceptance requirement is invented.
- Focused release-plumbing tests prove final 1.1.0 identity, matching notes, mandatory explicit epoch/full lowercase commit, collision refusal and legacy release identities. Full Go tests, race, vet, formatting, deterministic ordinary build, systemd static verification and security/privacy checks pass proportionally.
- Twin builds from independent exact commit exports have byte-identical binary, manifest, archive and sidecar. Frozen artifacts retain exact recorded sizes/hashes and read-only identity.
- Archive and package verification covers sidecar, safe paths/types/root/member uniqueness, manifest, documentation/LICENSE, modes, static linux-amd64 format, binary version/full commit/build time, installer/uninstaller and Smart Install readiness.
- Publication gate ledger has no missing evidence. Annotated tag object peels exactly to the release-source commit; pushed tag matches Forgejo; Release is final, not draft/prerelease, and contains exactly the two expected assets.
- Clean anonymous wget and curl downloads each reproduce frozen archive/sidecar sizes and hashes, pass sidecar and manifest verification, preserve filenames and prove binary provenance.
- Explicit staging excludes the Owner blueprint, credentials, caches and unrelated files. Commits/pushes are clean fast-forwards; final HEAD/origin/main/direct Forgejo agree; lifecycle is canonical idle with zero active prompts.


## Documentation Updates

- Add `docs/release/RELEASE_NOTES_1.1.0.md` and `docs/release/ACCEPTANCE_1.1.0.md`.
- Update `VERSION`, `CHANGELOG.md`, canonical installation/download/release documentation and `scripts/build-release.sh` / `scripts/test-release-plumbing.sh` only as required for final 1.1.0 identity and provenance.
- Update concise Engineering History and Task 063 prompt/history with source lineage, reused Task 061 evidence, deterministic/frozen artifact identities, release verdict, tag/Release/assets/download evidence, rollback and final state.
- Preserve Task 061 and all earlier acceptance histories byte-for-byte. Do not place credentials, private identities, raw authenticated responses or candidate bytes in Git.


## Completion Criteria

Task 063 completes successfully only when the exact accepted RC.6 behavioral lineage is reconciled; minimal final-release changes are implemented without an unclassified behavior change; proportional local validation passes; one exact release-source commit is clean-fast-forward integrated; two isolated builds are byte-identical; one final archive/sidecar pair is frozen and fully verified; Task 061 acceptance is truthfully reused or only materially invalidated evidence is rerun; the final verdict is READY FOR RELEASE; every conditional publication gate is PASS; annotated tag `v1.1.0` points exactly to the release-source commit; the final non-draft/non-prerelease Forgejo Release contains exactly the frozen archive and sidecar; anonymous wget/curl downloads reproduce and verify those identities; canonical release documentation/history are integrated; refs agree; and Task 063 archives to idle. If a mandatory gate cannot pass safely, publication does not occur and completion truthfully records NOT RELEASED plus the exact blocker. Completion does not authorize production deployment, unrelated announcement, tag/asset replacement or Task 064.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-26 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
