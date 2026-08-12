# Current Engineering Task 044: QWSG 1.0 Final Release & Git Reconciliation

## Task Metadata

- Task ID: `044`
- Task slug: `qwsg-1-0-final-release-git-reconciliation`
- Status: `active`
- Date opened: `2026-08-11` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

QWSG 1.0 Final Release & Git Reconciliation


## Objective

Reconcile the complete accepted Task 025–043 working-tree baseline with canonical Git source, prepare the smallest coherent `1.0.0` release identity and documentation transition, and establish a reproducible source-commit-to-artifact chain capable of reaching `READY FOR QWSG 1.0.0 PUBLICATION`. Preserve the accepted local Community Guardian without product reduction. Stop at the explicit Owner gates for final QWS Community / Free License legal text, reviewed staging and commit authority, and later tag, push and publication authority. Preparation and installation of this task grant none of those actions.



## Scope

- Audit every tracked modification, untracked path, ignored path, mode and relevant ACL against `HEAD` and `origin/main` at `0a8a5c7e722495b8c5eb425bca5b2d2413aaa175`; fetch only when separately permitted and never merge, rebase, reset or rewrite history.
- Produce a deterministic path-level classification manifest separating canonical release source, tests, canonical architecture/documentation, canonical engineering history/archive records, release packaging/scripts, generated release artifacts, local Builder inputs, backups/snapshots, compiled binaries and unrelated Owner-owned content.
- Treat accepted Go source and tests, repository governance, architecture and operator/release documentation, packaging, systemd unit, release scripts, and legitimate Task 025–043 prompt/history evidence as presumptive canonical source, subject to byte/path/privacy review.
- Treat `dist/`, `build/`, the root `qwsg` binary, `current-task-job.txt`, `current-task-maker.md`, temporary Builder inputs, generated caches, and backup/snapshot payloads as presumptively generated or local and excluded from the release commit. Do not delete, move or overwrite them merely because they are excluded.
- Review `ai/backups/` path by path: preserve already tracked historical metadata without rewriting history; do not add ignored payloads, archives, preserved working copies, host evidence or new untracked backup material unless a specific canonical-document requirement is proven and separately reviewed.
- Make only evidence-supported `.gitignore` corrections needed to prevent the root compiled binary, Builder input files, generated release artifacts, caches and noncanonical backup payloads from accidental staging. Ignore rules must not conceal canonical source, Task 025–044 lifecycle evidence or unrelated Owner content from the audit.
- Preserve all accepted Task 039 Runtime/Alert/Console behavior, Task 041 first-run bootstrap behavior, Task 042 reproducibility behavior and Task 043 clean-host/reboot acceptance evidence without reimplementation or semantic change.
- Document the QWS Community product policy: the accepted local inventory, Snapshot/Digital Twin, comparison, drift, health, rule, policy, report, Guardian, Console, state/history, diagnostics and install/upgrade/rollback/uninstall capabilities remain genuinely useful free local functionality and are not removed to manufacture a Pro boundary.
- Document future Pro/API-key direction only as nonbinding product architecture: central API/repository, fleet, Dashboard, managed notification/alerting, teams/roles, compliance services, remote management and managed backup may be future commercial services; their absence or failure must not disable the local Community Guardian or corrupt local evidence. Implement none of them.
- Identify the current `LICENSE` precisely as a temporary proprietary all-rights-reserved notice that grants no redistribution, modification or sublicensing permission and therefore is not sufficient as the final QWS Community / Free publication license.
- Do not invent, select or import MIT, Apache, GPL or any other license. Require the Project Owner to provide or explicitly approve final legal QWS Community / Free License text before changing `LICENSE`, building a public final artifact or declaring publication readiness. Do not describe the product as OSI open source unless the approved legal text is an OSI-approved license.
- Prepare the smallest coherent `1.0.0` metadata transition, expected to include `VERSION`, the command fallback/version tests where necessary, `CHANGELOG.md`, `scripts/build-release.sh` final-version acceptance, `docs/release/QUICK_START.md`, a new `docs/release/RELEASE_NOTES_1.0.0.md`, a new `docs/release/ACCEPTANCE_1.0.0.md`, and narrowly required support/release-policy references. Preserve RC.1–RC.3 evidence.
- Generalize the release builder only enough to accept the final `1.0.0` identity while retaining its strict `1.0.0-rc.N` validation, deterministic archive contract and existing release gates.
- Before any staging, produce and review an explicit allowlist of every path proposed for the canonical source commit plus an explicit exclusion list. Broad staging (`git add -A`, `git add .`, `git add --all`) is prohibited.
- Require an explicit Project Owner authorization after the final legal license text, classification manifest, proposed exact path allowlist, diff, modes, secret/privacy scan and validation summary are available and before the first `git add` or `git commit`. Without that authorization, stop with the task active and report the exact remaining gate.
- After that specific authorization only, stage exact allowlisted paths, review the staged diff and create one truthful canonical QWSG 1.0 release-source commit without artificial backdating, squashing, rebasing or rewriting published history. This first commit contains every artifact build input and a pre-build acceptance plan, but must not falsely claim post-build results. Record its exact hash and require a clean committed release-source tree apart from explicitly classified preserved local/ignored Owner content.
- Build the final `1.0.0` linux-amd64 archive from that exact committed source using the full Task 038/040/042 reproducible process, with `BUILD_COMMIT` derived from the real release-source commit and a controlled `SOURCE_DATE_EPOCH` derived from or explicitly tied to that commit. Prove two controlled byte-identical builds.
- Verify the expected final archive `dist/qwsg-1.0.0-linux-amd64.tar.gz`, matching sidecar, internal `MANIFEST.sha256`, static binary version, full embedded commit identity, safe archive layout, installer, uninstaller, user unit and complete release documentation.
- Re-run applicable Task 038, 039, 041, 042 and 043 gates, including artifact-level large Report behavior, read-only Console refresh, privacy-safe diagnostics, lifecycle truthfulness, empty-HOME bootstrap, truthful partial baseline, later full evaluation, staged install/upgrade/rollback/uninstall, service validation and retained clean-host/reboot evidence.
- After the artifact gates pass, finalize only non-artifact-input acceptance and lifecycle evidence with actual hashes. Require a second explicit Owner authorization before an evidence-only commit. Prove that this second commit changes no Go source, VERSION, LICENSE, CHANGELOG, packaged release document, packaging file, release script or other build input. This avoids the impossible circular claim of placing a commit-dependent artifact hash inside the same commit from which that artifact is built.
- Produce a final technical decision `READY FOR QWSG 1.0.0 PUBLICATION` only when source, license, release-source commit identity, reproducibility, artifact, evidence-only commit and acceptance gates all pass. This decision still does not authorize a tag, push, Forgejo Release, upload or announcement.
- Prepare but do not execute the annotated `v1.0.0` tag, push and Forgejo/public-release plan. The plan must present both defensible tag targets—release-source commit or later evidence-only commit—and recommend one with exact traceability consequences; the Owner selects and authorizes the target after reviewing both hashes. Require a new explicit Owner authorization before tag creation and explicit push/publication authority before each external action.



## Out of Scope

- No new product feature, product redesign, collector change or accepted Task 039/041 behavior change.
- No licensing enforcement, API-key validation, API communication, payment, entitlement service, remote control, Dashboard, fleet, central repository, managed notification, team/role, compliance or backup service implementation.
- No removal or artificial restriction of accepted local Community functionality and no dependency of local Guardian operation on a future Pro/API service.
- No invented legal terms, silent standard-license substitution or unsupported OSI/open-source claim.
- No blanket staging, broad cleanup, history rewrite, reset, rebase, squash, force push, tag movement or deletion.
- No staging, commit, tag, push, Forgejo Release, artifact publication or public announcement without the exact later Owner authorization required by the corresponding gate.
- No deletion, relocation, chmod, ownership change or silent ignore treatment of unrelated Owner-owned content.
- No clean-VPS rerun unless a new objective release gate is proven to require it and the Owner separately authorizes the host operation; Task 043 physical acceptance remains canonical evidence for unchanged behavior.
- No public-license or commercial-product promise beyond the Owner-selected Community policy and explicitly approved legal text.



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

- Require canonical idle with Task 043 as the unique latest complete archived baseline, zero active prompts and no Task 044 prompt/history/archive collision.
- Run Framework and lifecycle validation; verify project root, `main`, exact HEAD, canonical HTTPS origin, `0/0` upstream relationship, tag namespace, empty index and complete porcelain status.
- Confirm Task 043's exact decision is `READY FOR QWSG 1.0 FINAL RELEASE`, no known technical blocker is recorded and RC.3 SHA-256 remains `bc4eca323cb07d23f0d6c884886655eb5549c03c136d8c773782d7813551585c`.
- Read the Constitution, Agent rules, Job Template, lifecycle/prompt/Git/Release policies, Tasks 038–043 prompts and histories, RC.1–RC.3 notes/acceptance, release build script, packaging, install/uninstall/unit, VERSION, CHANGELOG, LICENSE and `.gitignore`.
- Record all tracked modifications, untracked and ignored paths with sizes, modes, owners and ACLs. Resolve repository-relative real paths and reject symlink or special-file ambiguity before classification or staging.
- Confirm the current LICENSE hash/content and record that it is temporary proprietary text, not the final Community legal grant.
- Record current release artifacts and hashes as generated historical evidence; do not treat their existence as authorization to commit them.
- Run pre-change build, focused/full/race/vet/format, Framework, Builder, lifecycle, diversion/test-task, job, Git whitespace/index and RC.3 checksum/manifest gates. Stop on any unexplained baseline deviation.



## Snapshot Requirements

- Before any Task 044-owned repository edit, create a unique external mode-`0700` snapshot containing exact payloads or absence records for every proposed metadata, license, `.gitignore`, release-script, documentation and lifecycle target.
- Record the entire pre-task Git path inventory, tracked/untracked/ignored classification, index, branch/HEAD/origin/upstream/tags, file identities, hashes, modes, owners, ACLs, tool versions, disk capacity and active QWSG/release/Git processes.
- Include exact hashes or bounded manifests for canonical source candidates, local exclusions, RC.1–RC.3 artifacts and historical lifecycle evidence; do not duplicate large ignored backup payloads when a deterministic inventory/hash is sufficient.
- Create a second verified pre-staging snapshot after metadata/license work and successful validation, containing the exact proposed source-commit payload, classification manifest, allowlist/exclusion list and guarded restore instructions.
- Preserve both snapshots through completion. Rollback instructions must reject active processes, later Owner edits, unexpected target hashes, symlinks, ambiguous ownership and broad Git recovery.



## Risk Assessment

- **Uncontrolled staging or Owner-content capture — critical:** use an explicit path allowlist only; review staged/unstaged status, modes, privacy and exclusions before commit.
- **Loss of accepted uncommitted source — critical:** inventory and snapshot the full working tree before edits; never reset, checkout, clean, stash broadly or delete excluded paths.
- **Legally invalid public release — critical:** temporary LICENSE grants no Community rights; final Owner-approved legal text is mandatory before final artifact/publication readiness.
- **False source-to-binary identity — critical:** final artifact must be built only after the canonical source commit and embed that actual commit hash, not the historical RC.3 hash or an uncommitted tree identity.
- **Generated/transient contamination — critical:** exclude `dist/`, binaries, caches, Builder input and backup payloads while preserving them locally and retaining canonical lifecycle evidence.
- **Release behavior regression — critical:** retain Tasks 038/039/041/042/043 gates and stop if final metadata/build changes alter accepted product behavior.
- **Premature external action — critical:** separate Owner gates govern staging/commit and later tag/push/publication; technical readiness alone grants no external authority.
- **License/product-policy conflation — high:** keep Community capability policy, future Pro architecture and binding legal license text distinct.
- **Dirty-tree ambiguity after commit — high:** a canonical committed source can coexist with explicitly classified ignored/local Owner files, but every remaining path must be explained and must not affect build inputs.



## Planned Work

1. Validate the Task 043 idle baseline, read governance/release evidence, inventory the complete working tree and run pre-change gates.
2. Create and verify the initial bounded snapshot. Produce a path-level classification manifest and exact proposed canonical-source allowlist/exclusion list, including tracked modifications, all 109 currently observed untracked paths and ignored/local content.
3. Review `.gitignore` and make only narrow rules for proven generated/local paths. Re-run the inventory to prove canonical files remain visible and no Owner content was altered.
4. Record Community capability policy and future nonbinding Pro/API direction without implementing entitlement or reducing local behavior.
5. Present the current temporary LICENSE incompatibility. Obtain exact Owner-approved QWS Community / Free legal text before modifying LICENSE; if it is not supplied, stop at this gate without claiming final publication readiness.
6. With approved legal text, perform the bounded `1.0.0` metadata/release-document transition, including final notes and acceptance baseline, and narrowly generalize the release script for final identity.
7. Run complete source, security/privacy, documentation, lifecycle and release-precommit validation. Create and verify the pre-staging snapshot, exact path allowlist, exclusion list, diff and mode report.
8. Stop and request explicit Owner authorization before any staging or release-source commit. After authorization only, stage each reviewed path explicitly, verify cached diff/whitespace/modes/secrets and create the canonical release-source commit with a truthful message.
9. Verify the release-source commit contains exactly the accepted build baseline and that remaining working-tree paths are only classified local/ignored Owner content. Derive `BUILD_COMMIT` from that exact commit and use a controlled commit-tied `SOURCE_DATE_EPOCH`.
10. Build twice in independent clean cache/output roots; compare binary, internal manifest, archive and sidecar byte-for-byte. Produce the canonical `dist/qwsg-1.0.0-linux-amd64.tar.gz` and `.sha256` locally without installing or publishing them.
11. Run the full Task 038/039/041/042 artifact suite and reconcile Task 043 clean-host/reboot evidence for unchanged functionality. Verify version and full embedded commit identity match the canonical source commit.
12. Finalize `docs/release/ACCEPTANCE_1.0.0.md` and Task 044 history with artifact hashes, reproducibility and exact remaining Owner actions. Prove these post-build changes touch no artifact build input, then stop for explicit Owner authorization before exact-path staging and creation of a second evidence-only commit.
13. If every gate and both authorized commits pass, state `READY FOR QWSG 1.0.0 PUBLICATION`. Report both commit hashes and the recommended annotated `v1.0.0` target/message, exact push commands and Forgejo release payload. Execute no tag or external action without a new explicit Owner authorization. Archive Task 044 and return canonical idle only if its authorized objective and all completion gates are actually satisfied; otherwise keep the task truthfully active at the Owner gate.



## Rollback Plan

Before staging, restore only Task 044-owned metadata, license, ignore, release-document, script and lifecycle paths whose current hashes still match recorded Task 044 output, using the verified snapshot and no broad Git command. Never delete excluded local artifacts, Builder inputs, backups or Owner content. After staging but before commit, unstage only the exact Task 044 allowlisted paths through the approved bounded mechanism and verify worktree bytes remain intact. After a local commit, do not rewrite or discard it automatically; stop and request Owner direction, preserving the commit and snapshots. After any tag or external action, no rollback is implied: tag deletion/movement, force operations and remote changes require separate explicit authority and recovery planning. Re-run lifecycle, source inventory, hashes, privacy, index and repository validations after every rollback.



## Deliverables

- Complete deterministic Git reconciliation report and path-level classification manifest for tracked, untracked and ignored content.
- Exact canonical source allowlist and explicit local/generated/backup/Owner-content exclusion list; no blanket staging.
- Narrow `.gitignore` hardening where proven necessary, without hiding canonical evidence.
- Documented Community capability policy, future Pro/API direction and clear separation from binding legal terms.
- Exact assessment of the temporary proprietary LICENSE and, only after Owner approval, the final QWS Community / Free legal text installed verbatim.
- Coherent `1.0.0` metadata, changelog, Quick Start, final release notes and final acceptance documentation while preserving RC.1–RC.3 evidence.
- Two-commit reconciliation strategy: an explicitly authorized release-source commit whose hash becomes the artifact's embedded identity, followed after validation by a separately authorized evidence-only commit that changes no artifact input.
- Reproducible `dist/qwsg-1.0.0-linux-amd64.tar.gz`, matching sidecar and internal manifest built twice from the exact committed source.
- Complete validation and Task 038/039/041/042/043 gate reconciliation with a justified readiness decision.
- Prepared annotated `v1.0.0` tag/push/Forgejo publication plan with an explicit Owner choice between the release-source and evidence-only commit targets, not executed without separate Owner authority.
- Completed Task 044 lifecycle evidence, retained snapshots and either canonical idle closure or a truthful active Owner-gate state.



## Verification

- Verify initial/final root, branch, HEAD, origin, upstream, tags, index, tracked/untracked/ignored inventory, modes, owners, ACLs and snapshots; prove unrelated Owner content is byte-identical.
- Review every proposed committed path and every remaining untracked/ignored path. Require zero unexplained paths that can influence the release build.
- Verify no broad staging command was used. Before commit, compare the exact staged path list with the approved allowlist and run `git diff --cached`, `git diff --cached --check`, mode review and secret/private-host scans.
- Verify final LICENSE is exactly Owner-approved, packaged by the build, internally manifested and described without an unsupported OSI/open-source claim.
- Verify `VERSION=1.0.0`; command output, fallback/tests, archive root/name, release notes, Quick Start, changelog and acceptance record agree exactly.
- Run `make build`, focused Task 039/041 tests, `go test ./...`, bounded-cache `go test -race ./...`, `go vet ./...`, `make fmt-check`, Framework, all Builder tests, lifecycle tests, diverted-task tests/audit, `bin/job`, Git whitespace and index gates at every required phase.
- Verify two independent final builds are byte-identical with fixed linux-amd64, CGO-disabled, trimpath, disabled VCS stamping, stable archive metadata and commit-tied build inputs.
- Verify sidecar from its correct directory, internal `MANIFEST.sha256`, safe single archive root, no links/special files, static binary, exact full commit identity, installer/uninstaller behavior and systemd unit.
- Repeat artifact-level large Policy Report/Alert completion, read-only Console refresh/no competing observe, bounded diagnostics/Attention, Guardian lifecycle/freshness, empty-HOME bootstrap, partial first baseline, later complete pipeline, state permissions and security-path regressions as practical in the engineering environment.
- Reconcile final behavior with the accepted Task 043 clean-host, physical reboot, recurrence, restart, resource and uninstall evidence. Do not claim a new external-host run.
- Confirm `dist/` and other generated/local outputs are not in the source commit or tag tree, while the final artifact hashes are recorded in non-artifact acceptance/history evidence according to policy.
- Require no tag, push, release upload or publication before the corresponding explicit Owner authorization; record commands and results only if later authorized and actually executed.



## Documentation Updates

- Update `.gitignore` only for the audited generated/local exclusions proven necessary.
- Update `VERSION`, `CHANGELOG.md`, `docs/release/QUICK_START.md`, narrowly necessary command version metadata/tests and `scripts/build-release.sh` for final `1.0.0` consistency.
- Add `docs/release/RELEASE_NOTES_1.0.0.md` and `docs/release/ACCEPTANCE_1.0.0.md`; preserve all RC release notes and acceptance evidence.
- Add a bounded repository reconciliation/classification record in the canonical release or Task 044 history location; do not store secrets, private host data or generated payload bytes.
- Update Community/product-policy documentation only enough to distinguish retained free local capabilities, future Pro/API direction and legal license terms.
- Replace `LICENSE` only with exact Owner-approved legal text received at the explicit gate.
- Complete `ai/history/044_2026-08-11_qwsg-1-0-final-release-git-reconciliation.md`, add one concise engineering milestone and archive the matching prompt only on genuine completion.



## Completion Criteria

Task 044 is complete only when it begins at canonical idle Task 043; every working-tree path is classified; accepted Task 025–043 source and canonical evidence are preserved; generated/local/backup/Owner content is safely excluded without loss; the final legal LICENSE is explicitly Owner-approved and internally consistent with the Community policy; final metadata is exactly `1.0.0`; a separately authorized exact-path release-source commit exists without history rewrite; the final artifact is built twice byte-identically from that exact commit and embeds its real identity; checksum, manifest, packaging, security and all applicable Task 038/039/041/042/043 gates pass; a separately authorized evidence-only commit records actual post-build results without changing artifact inputs; final acceptance states `READY FOR QWSG 1.0.0 PUBLICATION`; and no unexplained release input remains. Tag creation, push, Forgejo/public release and announcement require separate explicit Owner authorization and are not implied by completion. If legal text, either commit authority or a mandatory technical gate is absent, keep the task active at the exact Owner/blocker gate rather than claiming completion or publication readiness.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-11 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
