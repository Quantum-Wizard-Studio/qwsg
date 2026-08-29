# Current Engineering Task 073: QWSG 1.2.0 Final Release and Publication

## Task Metadata

- Task ID: `073`
- Task slug: `qwsg-1-2-0-final-release-publication`
- Status: `active`
- Date opened: `2026-08-29` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

QWSG 1.2.0 Final Release and Publication


## Objective

Complete the canonical QWSG 1.2.0 final release from the fully accepted 1.2.0-rc.7 implementation and the Project Owner-supplied Contabo and OVH real-host acceptance evidence. Produce, publish, retrieve, and independently verify the deterministic final artifact, then close the lifecycle only after every mandatory release invariant passes.


## Scope

- Canonically record the supplied real-host acceptance evidence and Task 073 release history.
- Prove product-code equivalence between accepted candidate source commit `07d30315cebba4bd213e76fe9fb9c32e82aa9b38` and the Task 072 final repository commit `79585da0b2fa95f8b7aaa3e14fc15182d690a3b7`, distinguishing lifecycle/history/documentation-only commits.
- Determine and document the canonical traceable promotion method from accepted `1.2.0-rc.7` to final `1.2.0`.
- Make only version and release-metadata/documentation changes required for final identity, including the minimum declarative compatibility metadata needed for a deterministic installed RC.2 to final 1.2.0 upgrade path.
- Run all canonical final-release validation and deterministic packaging gates, produce `qwsg-1.2.0-linux-amd64.tar.gz`, record provenance and checksum, create annotated tag `v1.2.0`, push the release commit and tag, create the canonical Forgejo Release, attach the artifact/checksum information, retrieve the published artifact independently, and verify it.
- Finish with exact local/remote/release/lifecycle evidence and archive Task 073 only on successful release.


## Out of Scope

- No features, architecture redesign, opportunistic refactoring, product-behavior change, or unrelated UX/documentation correction.
- No mutation of Contabo or OVH hosts.
- No invented acceptance evidence, secret exposure, force push, history rewrite, or release publication after a failed mandatory gate.
- No materially different source rebuild concealed as RC.7 promotion.
- No successor task creation.


## Authority Envelope

**Task targets and boundaries:** Framework 2.0 Standard Execution Authority applies to the release-only objective: Task 073 lifecycle/evidence, final-version identity and release metadata/documentation, declarative final-version compatibility routing if required, deterministic build and package outputs, task-scoped Git integration, annotated `v1.2.0` tag, Forgejo v1.2.0 Release publication, independent external artifact retrieval and verification, and lifecycle closure. Product behavior, architecture, unrelated improvements, and external host mutation are excluded.

**Permitted external actions:** Fetch and inspect canonical Forgejo refs; clean-fast-forward push the reviewed release and lifecycle commits; create and push the annotated `v1.2.0` tag; create the authorized Forgejo v1.2.0 Release; upload the exact validated final artifact and checksum/release information; retrieve the published artifact through the actual Forgejo release download endpoint using curl and/or wget; and verify remote refs and release availability. These publication actions are explicitly authorized by the Project Owner. No Contabo or OVH mutation is permitted.

**Owner-reserved decisions:** Any material product-code change, architectural redesign, unplanned migration mechanism, acceptance waiver, failed-gate override, destructive external action beyond the exact publication transaction, credential decision, force push, history rewrite, tag replacement/deletion, or scope expansion remains Owner-reserved.

**Task-specific STOP conditions:** Stop and report BLOCKED before tag or publication if product-code equivalence to accepted RC.7 cannot be proven, deterministic rollback is unavailable, any mandatory release gate fails unresolved, final RC.2 compatibility requires product redesign or an unplanned migration mechanism, candidate provenance or reproducibility is ambiguous, secrets/privacy/security are at risk, remote state diverges, or the exact final artifact cannot be independently retrieved and verified. After publication, use bounded containment and report truthfully if an external invariant fails; never hide or override failure.


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

- Verify repository root, ordinary user, Framework 2.0.0, canonical HTTPS origin and `main` branch.
- Require a clean worktree and index with `HEAD == origin/main == 79585da0b2fa95f8b7aaa3e14fc15182d690a3b7` and zero divergence after fetch.
- Require canonical idle after completed/archived Task 072 and next production Task ID 073.
- Verify accepted artifact identity `dist/qwsg-1.2.0-rc.7-linux-amd64.tar.gz` SHA-256 `fe5fdeb93efe0363376d5d4b85ac915b61817a3f7775c8158c5d62cb7db7c631`, accepted implementation source commit `07d30315cebba4bd213e76fe9fb9c32e82aa9b38`, and Task 072 final commit `79585da0b2fa95f8b7aaa3e14fc15182d690a3b7`.
- Verify relevant tag/Forgejo Release/destination absence or classify any collision before mutation.
- Treat the supplied Contabo and OVH facts as manually supplied Project Owner acceptance evidence; record them without claiming new host execution.


## Snapshot Requirements

Before any repository or publication mutation, create a unique private mode-0700 Task 073 snapshot under `/tmp`. Include a readable exact tracked-HEAD archive; literal lifecycle/release/documentation target before-images and absence records; Git/ref/status/version/mode/ACL evidence; accepted RC.7 artifact/source identities; Builder input identity; relevant release-tool and compatibility-registry identities; tag/Release collision evidence; checksum manifest; and exact guarded bounded restore/containment instructions. Exclude secrets, credentials, private host access material, caches, and unrelated content. Before tag/publication create and verify a second release-boundary snapshot containing the exact release commit/tree, frozen artifact identity, proposed tag identity, publication ledger, and containment actions. Preserve snapshots through delivery and verify integrity/readability, modes, exclusions, and rollback capability.


## Risk Assessment

- Critical: publishing source not product-code-equivalent to accepted RC.7; mitigate with path-classified Git lineage and tree/diff proof before finalization.
- Critical: releasing after a failed mandatory gate or with a non-reproducible artifact; mitigate with fail-closed gate ledger, isolated builds, byte comparison, archive integrity and provenance checks before tagging.
- Critical: missing RC.2 to final 1.2.0 declared upgrade route; inspect the registry and add only minimum declarative release metadata, otherwise stop BLOCKED.
- High: wrong tag target, remote divergence, or mismatched uploaded bytes; mitigate with exact hashes, annotated-tag inspection, clean-fast-forward push, direct remote-ref checks, and independent external retrieval.
- High: secret/privacy disclosure through evidence or publication; use redacted classifications and exclude credentials/private access material.
- Medium: lifecycle or rollback inconsistency; use two verified bounded snapshots, explicit staging, lifecycle validators, and canonical idle closure only after release success.


## Planned Work

1. Install and validate Task 073 through the canonical Builder; read all required governance and relevant release/acceptance records.
2. Create and verify the pre-change snapshot, record the exact clean baseline, fetch canonical refs, and prove accepted RC.7 artifact/source identities.
3. Audit commits from accepted implementation source through Task 072 closure and prove no product-code change; determine and document the promotion contract.
4. Inspect final-version metadata, packaged documentation, release tooling and migration registry; implement the smallest release-only diff for 1.2.0 and explicit RC.2-to-final compatibility when declarative metadata suffices.
5. Run all mandatory format, vet, full/race, Framework, Builder, lifecycle, packaging, reproducibility, cross-umask/source-mode, archive ownership/mode/executable, security/redaction, notification, migration/update/rollback and release-plumbing tests.
6. Build the final deterministic artifact in isolated outputs, verify reproducibility and MANIFEST, record the exact source commit, release commit, timestamp contract, filename and SHA-256, and freeze the bytes.
7. Create and verify the release-boundary snapshot; review/stage explicit paths, commit and clean-fast-forward push the validated release commit.
8. Create annotated `v1.2.0` at the exact canonical release commit only after all local gates pass; push and verify the tag.
9. Create the Forgejo v1.2.0 Release, upload the frozen artifact and checksum/release information, then independently download it through the release endpoint and verify non-empty bytes, exact SHA-256, readable tar, MANIFEST and package structure.
10. Finalize Task 073 evidence, commit/push lifecycle closure, archive the task to canonical idle, and verify local/remote/release state. If any mandatory invariant cannot be proven, stop and report BLOCKED without falsely completing the release.


## Rollback Plan

- Before publication, restore only explicitly reviewed Task 073 repository targets from the verified snapshot when their identity and collision guards still match; remove only Task 073-created local outputs proven absent before the task. Never use broad reset, checkout, clean, wildcard deletion, or history rewrite.
- After the release commit is pushed but before publication, prefer a new bounded corrective commit if needed; never rewrite accepted history.
- Once tag/Release/publication exists, do not silently move/delete published identities. Stop, preserve evidence, apply only explicitly authorized bounded containment, and request Owner direction for tag/Release replacement or withdrawal.
- After any rollback or containment, rerun relevant Framework/lifecycle/Git/product/release checks and retain the failed-attempt evidence and snapshots.


## Deliverables

- Canonical real-host final acceptance and release documentation plus complete Task 073 history.
- Written product-code-equivalence and RC.7-to-final promotion proof.
- Final `1.2.0` version/release metadata and explicit deterministic RC.2-to-final upgrade compatibility route where the existing registry requires it.
- Reproducible `qwsg-1.2.0-linux-amd64.tar.gz` with exact SHA-256, source/release commit provenance, verified MANIFEST, archive modes/ownership and executable structure.
- Annotated local and remote `v1.2.0` tag targeting the exact canonical final release commit.
- Published Forgejo v1.2.0 Release with attached artifact and checksum/release information.
- Independent external download and integrity evidence.
- Clean synchronized repository and canonical idle lifecycle after successful archival, or an explicit BLOCKED report preserving truthful state.


## Verification

- `gofmt` verification, `go vet`, full `go test`, required race tests, ordinary build and exact version output.
- Framework, Builder, lifecycle, active-job/test-task and configured engineering validations.
- Release packaging/plumbing, isolated deterministic/reproducible build comparison, cross-umask/source-mode checks, MANIFEST verification, canonical archive modes, numeric ownership and executable semantics.
- Security/privacy/redaction and notification regression suites.
- Migration/update/rollback regressions, including an actual deterministic installed RC.2 fixture route to final `1.2.0`; no assumption that RC.7 coverage is sufficient.
- Exact accepted-source lineage and path-classified diff proving no post-RC.7 product-code change.
- Clean explicit staged diff, whitespace/mode/ACL review, no secrets/private host material, and clean-fast-forward push evidence.
- Annotated tag object and target verified locally and remotely; Forgejo Release existence and assets verified.
- External curl/wget download uses a separate path and actual release endpoint, succeeds with non-empty bytes, matches canonical SHA-256, opens as tar, passes extracted `MANIFEST.sha256`, and contains expected executable/package structure.
- Final `HEAD == origin/main`, zero divergence, clean worktree/index, canonical tag/release availability, and canonical idle lifecycle.


## Documentation Updates

Update only the canonical final acceptance/release record, release notes/changelog/version/install or packaged documentation required to identify final 1.2.0 accurately, relevant compatibility documentation, concise Engineering History milestone, and Task 073 prompt/history. Record supplied host evidence as Project Owner-provided acceptance evidence and clearly distinguish it from local/external checks performed during Task 073.


## Completion Criteria

Task 073 is complete only when product-code equivalence to accepted RC.7 is proven; final 1.2.0 metadata and any required declarative RC.2 compatibility route are correct; all mandatory local release gates pass; one exact reproducible final artifact is frozen with recorded provenance and SHA-256; the exact release commit is pushed; annotated `v1.2.0` targets it locally and remotely; the Forgejo v1.2.0 Release exists with the exact artifact/checksum information; an independently downloaded release artifact passes checksum, tar, MANIFEST and structure verification; the repository is clean and synchronized; Task 073 history is complete; and lifecycle is archived/idle. The final decision must be exactly RELEASED or BLOCKED and report every Owner-required identity and status. Any unresolved mandatory invariant produces BLOCKED, no false completion, and no public release after a pre-publication gate failure. No product feature, behavior change, unrelated refactor, host mutation, or successor task is permitted.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-29 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
