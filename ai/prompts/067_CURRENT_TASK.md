# Current Engineering Task 067: Deterministic Release Packaging & QWSG 1.2.0-rc.3

## Task Metadata

- Task ID: `067`
- Task slug: `deterministic-release-packaging-rc3`
- Status: `approved`
- Date opened: `2026-08-28` UTC
- Human authority: Project Owner: Attila
- Owner or lead-developer communication language: English

## Title

Deterministic Release Packaging & QWSG 1.2.0-rc.3


## Objective

Correct only the release-packaging determinism defect identified by Task 066, establish and enforce canonical archive permissions independent of ambient filesystem modes and caller umask, prove byte-for-byte artifact reproducibility while preserving intended executable semantics, and produce a private QWSG 1.2.0-rc.3 candidate suitable for later final acceptance.


## Scope

- Establish and record the mandatory starting-state evidence and rollback-capable pre-task snapshot before modifying task targets.
- Inspect the current release builder and its existing release tests before designing the remediation.
- Define an explicit canonical packaged-permission policy distinguishing directories, executable regular files, and non-executable regular files.
- Modify `scripts/build-release.sh`, or its canonical delegated release-building layer, to normalize package modes explicitly without losing required executable semantics.
- Preserve all existing deterministic ordering, timestamp, ownership, path, metadata, and compression behavior.
- Add automated regression coverage for caller umask differences, group-writable versus normalized regular-file modes, group-writable versus normalized directory modes, intended executable preservation, and SHA-256 byte identity.
- Update private candidate metadata from 1.2.0-rc.2 to 1.2.0-rc.3 only where canonically required after remediation tests pass.
- Build `qwsg-1.2.0-rc.3-linux-amd64.tar.gz`, calculate and record its SHA-256, independently inspect archive metadata and modes, repeat the build under materially different permission/umask conditions, compare hashes and bytes, and verify extracted permissions.
- Update only the canonical release/build documentation and Task 067 history needed to record the Task 066 defect, RC.2 rejection, canonical permission policy, reproducibility rules and regression protection, RC.3 creation, and exact RC.3 SHA-256.
- Run all focused and repository-wide required tests, perform task-scoped Git integration, and complete the canonical lifecycle when every completion criterion is proven.


## Out of Scope

- Do not create the final `v1.2.0` release tag, publish QWSG 1.2.0, create a public Forgejo final release, claim final release acceptance, or bypass Task 066 acceptance gates.
- Do not publish RC.3 publicly or mutate external release distribution systems.
- Do not perform the full OVH or Contabo/Hestia release-acceptance matrix and do not modify either VPS.
- Do not change unrelated product behavior, architecture, infrastructure, or repository permissions merely to make a build pass.
- Do not weaken release-security properties, expand privileges, derive archive ownership from the host, or rely on the developer machine's current umask as canonicalization.


## Authority Envelope

**Task targets and boundaries:** Framework 2.0 Standard Execution Authority applies to the release builder and delegated packaging layer, release regression tests and fixtures, private candidate version metadata, local/private RC.3 artifacts and checksums, narrowly relevant release/build documentation, Task 067 lifecycle records, rollback evidence, and task-scoped Git integration. The work is limited to correcting the Task 066 permission-determinism defect and producing a private RC.3 candidate. The agent may inspect, snapshot, reproduce, edit, build, test, correct, retest, document, stage explicit task paths, commit, perform push dry-run and clean fast-forward push, and close the lifecycle when all evidence passes.

**Permitted external actions:** Fetch the configured `origin` to verify the baseline and perform the canonical task-scoped clean fast-forward Git push after mandatory validation. No VPS mutation or release publication is permitted. Local protected snapshot and private build-artifact handling required for rollback and reproducibility evidence is permitted.

**Owner-reserved decisions:** Any final `v1.2.0` tag, final or public release, Forgejo Release or asset publication, public RC publication, acceptance waiver, change to the release identity beyond `1.2.0-rc.3`, material architecture or product-scope expansion, destructive external operation, force push, history rewrite, or modification of OVH/Contabo remains reserved to Project Owner Attila.

**Task-specific STOP conditions:** Stop for an unexpected baseline discrepancy that establishes an authority, identity, safety, or rollback boundary; unavailable or unverified rollback; credential, secret, privacy, or security exposure; inability to prove deterministic byte identity across required umask/source-mode variations; unexpected executable or writable archive modes that cannot be safely corrected in scope; a release-blocking engineering defect; an irreversible or unauthorized external mutation; or a canonical lifecycle requirement explicitly requiring Owner action. If deterministic reproducibility cannot be proven, record BLOCKED and do not declare RC.3 valid.


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

- Verify repository root `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, branch `main`, clean complete working tree, HEAD `90a5815058919d1d913e28a19662672ad5100b01`, and local `origin/main` at the same commit.
- Verify the configured canonical `origin`, current tags relevant to 1.2.0 candidates, ahead/behind relationship, Framework validation, and canonical idle lifecycle with Task 066 closed BLOCKED.
- Verify rejected private candidate metadata: source commit `c260dc18c2004473ec55496d16e66718fd128865`, artifact `qwsg-1.2.0-rc.2-linux-amd64.tar.gz`, rejected SHA-256 `a34be8b18f80d877c0ccfd69dc9d9e9f197fc35fa765cdf1d5c0d72e2cb0a554`, and normalized comparison SHA-256 `5b32df3b090658cfb9a08a7d670848c65af4d5d048dc053e3ad0973d11f0082a` where records remain available.
- Verify Task 066 recorded directories at 0775 and group-writable regular files including 0660 and 0664, and classify any material mismatch before implementation.


## Snapshot Requirements

- Before modifying task targets, create a protected local rollback-capable snapshot outside Git in accordance with `ai/core/15_ENGINEERING_BACKUP_POLICY.md`.
- Record UTC creation time, purpose, Git baseline, captured task-target scope, exclusions, retention through Owner acceptance, deterministic manifest and SHA-256 integrity evidence, archive readability, exact bounded restore procedure, collision behavior, and post-restore Git/test/lifecycle checks.
- Verify the snapshot checksum and perform a safe non-destructive restore rehearsal into an isolated temporary location. Do not extract over the live worktree for verification and do not commit payloads or sensitive host evidence.


## Risk Assessment

- Risk is high because release archives are candidate supply-chain artifacts and incorrect permission normalization can change bytes, remove legitimate executability, introduce writable or executable files, or invalidate acceptance evidence.
- Primary controls are protected pre-change rollback, inspection before design, explicit type-aware mode policy, isolated cross-mode/cross-umask builds, SHA-256 and byte comparison, archive listing and extraction inspection, focused regressions, full repository validation, targeted staging, and prohibition on publication.
- The implementation must avoid symlink traversal, host ownership leakage, privilege expansion, group/world-writable regular files, unexpected executables, and dependence on ambient umask.


## Planned Work

1. Read all required governance and relevant release/build documentation plus Task 066 history; establish and record exact starting evidence.
2. Create, verify, and rehearse the mandatory protected rollback snapshot.
3. Inspect `scripts/build-release.sh`, delegated packaging logic, release inputs, existing deterministic controls, current version metadata, and all existing release tests before choosing the smallest safe remediation.
4. Reproduce the permission leak and define the canonical archive policy for directories, executable regular files, and non-executable regular files.
5. Implement explicit mode normalization while preserving existing deterministic ordering, timestamps, ownership, paths, metadata, compression, and intended executable semantics.
6. Add automated SHA-256 regression tests covering different caller umasks, source file modes, source directory modes, executable preservation, forbidden writable modes, and unexpected executability.
7. Run focused tests and correct any in-scope defects. Only after they pass, update canonical private-candidate metadata to `1.2.0-rc.3`.
8. Build RC.3, independently inspect metadata and extracted modes, rebuild under materially different mode/umask conditions, and prove identical bytes and SHA-256.
9. Update narrowly relevant canonical documentation and Task 067 history with exact evidence, RC.2 rejection, policy, RC.3 identity, and limitations.
10. Run repository-wide mandatory tests and framework/lifecycle checks, review exact diffs and modes, stage explicit paths, commit and clean fast-forward push if permitted and safe, verify the final remote relationship, then complete and archive the lifecycle.


## Rollback Plan

- Before Git integration, restore only exact modified task-target files from the verified protected snapshot after checking its manifest and hashes; remove only explicitly identified Task 067-generated private build outputs when rollback requires it; then rerun Git status, focused release tests, Framework validation, and lifecycle validation.
- After a task commit but before push, use a new corrective commit or restore exact task paths from the snapshot and commit the bounded rollback; do not reset, rewrite, clean broadly, or disturb unrelated paths.
- After a clean fast-forward push, use a forward corrective commit under the same bounded authority when safe; tags and public releases are prohibited, so no publication rollback should be necessary.
- If exact restoration cannot be proven or rollback would affect material unrelated state, stop and request Owner direction.


## Deliverables

- Deterministic release-packaging implementation with an explicit canonical archive permission policy.
- Automated regression tests proving byte-for-byte identity across required caller umask and logically equivalent source-mode variations while preserving intended executable status.
- Security evidence that packaged regular files are not group/world writable, directories are canonical, executables are expected, and ownership is host-independent.
- Private `qwsg-1.2.0-rc.3-linux-amd64.tar.gz` candidate and recorded SHA-256, with independent archive and extraction inspection evidence.
- Repeated-build evidence showing identical RC.3 bytes and SHA-256 under materially different permission/umask conditions.
- Narrow release/build documentation updates and complete Task 067 history, Git evidence, and canonically closed lifecycle.


## Verification

- Run builder installation checks, `bin/job --check`, lifecycle consistency, reusable Framework validation, and all task-required governance checks.
- Demonstrate the pre-fix defect or preserve authoritative Task 066 evidence, then run automated tests that build logically equivalent release inputs with materially different caller umasks, group-writable versus normalized regular-file modes, group-writable versus normalized directory modes, and retained intended executable status.
- Compare resulting artifacts byte-for-byte and with SHA-256; the regression must fail if ambient mode differences leak into an archive.
- Run focused release/build tests, formatting/static checks, full tests, race tests, vet, and every repository-wide check required by repository scripts and policy.
- Inspect the RC.3 archive independently for stable path ordering, timestamps, ownership, metadata, compression expectations, and exact modes; reject group-writable regular files, world-writable entries, unexpected executables, privilege expansion, and build-host ownership.
- Extract into isolated directories under differing umasks and verify intended canonical file and directory modes without mutating external VPS hosts.
- Repeat RC.3 construction under at least one materially different permission/umask condition and require identical bytes and SHA-256.
- Review complete unstaged/staged diffs, path lists, executable bits, secrets/privacy, generated artifacts, and `git diff --cached --check`; use explicit targeted staging only.
- Verify final branch, commit, clean tree, local/remote relationship, lifecycle closure, and absence of any tag or publication action.


## Documentation Updates

- Update only canonical release/build documentation necessary to record the Task 066 packaging defect, canonical release permission policy, reproducibility requirements, automated regression protection, RC.2 rejection, RC.3 creation, exact RC.3 SHA-256, verification evidence, and remaining acceptance limitations.
- Update the matching Task 067 history throughout execution with starting state, snapshot and rollback proof, inspected design basis, changes, attempts and classifications, exact commands/results, artifact identity, Git integration, completion decision, and unresolved limitations.
- Do not rewrite unrelated architecture documentation or record secrets, private host data, raw snapshot payloads, or unsupported claims.


## Completion Criteria

Task 067 is COMPLETE only when the packaging defect is corrected; canonical archive modes are explicitly defined and enforced; cross-umask and equivalent source file/directory mode reproducibility pass byte-for-byte with SHA-256 evidence; legitimate executability is preserved; automated regression and repository-wide required tests pass; the private RC.3 artifact is created; repeated materially different builds have the same SHA-256; independent archive and extraction inspection confirms safe canonical permissions and host-independent ownership; exact RC.3 SHA-256 and RC.2 rejection are documented; task-scoped Git integration is verified; and the lifecycle is canonically closed. If deterministic reproducibility or another mandatory release condition cannot be proven, the final decision is BLOCKED, RC.3 is not declared valid, and the lifecycle/history truthfully records the blocker without publication.


## Owner Approval Requirements

Approved by Project Owner: Attila through the Engineering Task Builder on 2026-08-28 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
