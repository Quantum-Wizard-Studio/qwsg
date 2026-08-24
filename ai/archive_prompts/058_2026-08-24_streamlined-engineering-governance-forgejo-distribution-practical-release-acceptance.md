# Current Engineering Task 058: Streamlined Engineering Governance, Forgejo Distribution and Practical Release Acceptance

## Task Metadata

- Task ID: `058`
- Task slug: `streamlined-engineering-governance-forgejo-distribution-practical-release-acceptance`
- Status: `complete`
- Date opened: `2026-08-24` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

Streamlined Engineering Governance, Forgejo Distribution and Practical Release Acceptance


## Objective

Replace unnecessary per-operation Owner gates with a reusable task-level authority-envelope model while preserving rollback, security, privacy, lifecycle and historical-evidence invariants; define and locally validate the practical Forgejo Release-asset distribution contract for ordinary `wget`/`curl` users and a future Smart Installer without publishing artifacts; and replace the historical 26-checkpoint model with one concise, evidence-based 12-step acceptance run authorized as a bounded whole. The result must make routine in-scope engineering autonomous, retain explicit Owner control over material scope and reserved operations, and preserve Task 057's disclosed limited-acceptance outcome unchanged.


## Scope

- Audit Tasks 055–057 and the reusable Engineering Framework to identify rules and task-definition patterns that caused repeated approval gates.
- Introduce a mandatory, explicit Authority Envelope in the reusable task definition. It must enumerate task-owned paths/components, routine repository operations, bounded correction/retest authority, permitted external systems/actions, lifecycle completion authority, and Owner-reserved operations.
- Within an approved envelope, authorize ordinary inspection, snapshot, implementation, tests, bounded diagnosis/correction/retest, documentation, explicit path staging, staged-diff review, task-scoped commit, push dry-run, clean fast-forward push, post-push validation, and canonical lifecycle closure without new artificial sub-gates.
- Preserve one active task, deterministic Builder identity, task history, snapshots, rollback, explicit path staging, validation, security/privacy checks, clean fast-forward rules, and immutable historical evidence.
- Make STOP semantics proportional: recoverable in-scope failures follow diagnose -> correct -> retest -> continue; boundary, safety, authority, destructive, credential, uncertain external-mutation, publication and other Owner-reserved conditions stop.
- Update Builder, templates, checks, tests, workflow documentation and the project-local job skill only as required to encode and validate the authority envelope consistently.
- Establish the release-distribution design around Forgejo Release assets and exact versioned archive/sidecar names. Verify the actual self-hosted Forgejo route and unauthenticated download behavior read-only before documenting any concrete project URL; do not invent a URL.
- Add practical download/checksum documentation and validation suitable for `wget`, `curl`, `sha256sum -c`, and later Smart Installer consumption, but do not create a tag, Release or asset.
- Replace the 26-checkpoint model prospectively with one 12-step practical acceptance workflow under one bounded Owner authorization. Preserve Task 057 protocol and ledger as immutable history.
- Define concise mandatory evidence, privacy-safe reconciliation rules, technical clean-host validity, defect stop rules, protected credential pauses and separate publication authority.
- Integrate, validate, commit, clean-fast-forward push and canonically close Task 058 when all completion criteria pass, as routine operations expressly covered by this envelope.


## Out of Scope

- Do not rewrite, relabel or discard Task 057 or earlier RC/finding chronology. Task 057 remains COMPLETE WITH DISCLOSED ACCEPTANCE LIMITATIONS; core RC.5 functional proof achieved; formal 26-checkpoint certification incomplete; NOT READY was procedural, not a demonstrated RC.5 product defect.
- Do not access, reset, reinstall or otherwise modify the disposable VPS; it remains clean for later acceptance.
- Do not construct or rebuild a candidate, transfer or execute artifacts externally, handle credentials, perform external acceptance, operate Guardian externally, or mutate external infrastructure.
- Do not create or move a tag, create a Forgejo Release, upload assets, publish, announce, deploy to production, or make a public stable URL live. These remain separately Owner-authorized.
- Do not implement the Smart Installer itself. Define only the stable distribution contract it can consume later.
- Do not weaken snapshot/rollback, explicit task identity, single-active lifecycle, targeted staging, validation, historical evidence, secret handling, clean fast-forward, or Owner architecture/scope authority.
- Do not grant blanket permission for destructive operations, privilege escalation, credentials, unplanned external mutation, architecture expansion, publication or production deployment.
- Do not read, hash, modify, copy, stage, package or otherwise access Owner-owned `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`; pathname metadata-only exclusion checks are permitted.
- Do not create Task 059.

## Authority Envelope

1. **Authorized paths/components/systems:** Task 058 may modify
   the reusable engineering framework, Builder/lifecycle tooling and tests,
   prospective release-acceptance standards, release/distribution
   documentation, release-plumbing validation, and its own lifecycle records.
   Work is confined to this repository and private rollback/validation storage.
2. **Routine operations:** inspect, snapshot, analyze, edit, test,
   diagnose, correct, retest, document, validate, target-stage reviewed Task 058
   paths, review staged content and modes, commit, validate the exact commit,
   push dry-run, clean fast-forward push, verify refs, and report.
3. **Correction/retest authority:** recoverable implementation, test,
   documentation, and validation failures inside this scope follow diagnose ->
   correct -> retest -> continue. Missing evidence is never converted to PASS.
4. **Repository integration:** explicit path staging, task-scoped commits, push
   dry-run, clean fast-forward push to `origin/main`, and direct read-only
   Forgejo ref verification are authorized after required validation.
5. **Lifecycle completion:** Task 058 may truthfully finalize its prompt/history,
   archive the prompt, integrate the lifecycle-only closure, push it as a clean
   fast-forward, and return the repository to canonical idle without another
   routine Owner gate.
6. **Permitted external actions:** read-only access to public/official Forgejo
   documentation and read-only verification of repository/Release route
   behavior. No VPS or other infrastructure access is permitted.
7. **Evidence and rollback:** a verified private mode-0700 snapshot, bounded
   restore instructions, exact diffs/path sets/modes, validation results,
   commit/ref identities, security/privacy checks, and chronological history are
   mandatory.
8. **Owner-reserved operations:** material scope or architecture expansion,
   destructive or irreversible work outside the recorded rollback, credentials,
   secrets, privilege escalation, unplanned infrastructure mutation, tags,
   Forgejo Releases, asset upload, publication, deployment, announcement, and
   Task 059.
9. **Mandatory STOP conditions:** unavailable reliable rollback; meaningful
   unresolved security/privacy uncertainty; a need for credentials, elevated
   authority, unplanned external mutation, or an Owner-reserved operation; or
   meaningful risk of damage outside this bounded task. Ambiguity never expands
   the envelope.


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

- Verify UTC date, ordinary user, exact repository root and `main`; canonical HTTPS origin; `HEAD == origin/main ==` direct Forgejo `main` at `0e17a8598821071c4121e1df69b59cd0b62e5506`; ahead/behind `0/0`; clean index and tracked tree; zero active prompts; canonical idle with Task 057 complete and archived.
- Verify the only unrelated state is the excluded Owner blueprint by pathname metadata only.
- Verify `VERSION=1.1.0-rc.5`, no Task 058 prompt/history exists, and no new RC.5 construction/publication state was created after Task 057.
- Verify Task 057's terminal classification and evidence exactly as preserved, plus Tasks 055/056, RC.1/RC.2/failed RC.3/failed RC.4, QWSG-055-F001, QWSG-053-F001, QWSG-051-F001, Task 049 F002/F003, v1.0.0 identities and LICENSE.
- Read the Constitution, agent/job template, lifecycle, prompt workflow, Git policy, Builder, framework/lifecycle/diversion/job scripts and tests, and Tasks 055–057 definitions/history needed to identify actual gating causes.
- Inspect official Forgejo documentation and the configured repository read-only to establish the supported release-asset route model. Any concrete `git.quantumwizard.hu` asset URL remains provisional until verified against a separately published test or real Release under Owner authority.
- Run Framework, Builder-input, lifecycle, diversion, active-job and test-task baseline validation before modification; record material differences.


## Snapshot Requirements

Before modifying task targets, create and verify a unique private mode-0700 snapshot under `/tmp` containing a readable exact tracked-HEAD archive, literal copies and absence records for every proposed target, Git/ref/status evidence, modes/ACLs where relevant, framework/Builder identities, protected-history identities and exact bounded rollback instructions. Exclude Owner content, credentials, private infrastructure identity, candidate bytes, caches and unrelated files. Verify archive readability, hashes, target absence claims, collision safety and restore instructions before implementation. Create additional bounded snapshots before any materially distinct in-scope migration when the original snapshot would not provide a reliable rollback.


## Risk Assessment

- Authority overreach — critical: the envelope enumerates routine categories and exact reserved operations; ambiguity resolves to STOP, never assumed authority.
- Safety regression — critical: snapshot, bounded rollback, security/privacy, targeted staging, validation, single-task lifecycle and historical immutability remain mandatory and tested.
- Accidental publication — critical: tags, Releases, uploads, stable public activation and deployment remain Owner-reserved even when implementation/commit/push are enveloped.
- URL invention or instability — high: use official Forgejo route semantics, then verify the actual project endpoint before publishing concrete examples; never claim an unavailable asset.
- Smart Installer coupling — high: define a versioned archive/sidecar contract and checksum behavior without prematurely implementing installer networking or mutable-latest semantics.
- Acceptance under-testing — high: simplify duplicated ceremony, not product boundaries; preserve clean-host, integrity, install/setup/Guardian, real notification, reboot, restart, uninstall and same-candidate reinstall proof.
- Evidence gaming — critical: late reporting may be reconciled only from reliable independent evidence; missing proof remains missing and no defect is inferred solely from procedural absence.
- Credential/privacy exposure — critical: Owner-only protected entry, no secrets/private identities in chat, argv when avoidable, Git, history or artifacts.
- Historical corruption — critical: old protocols/findings remain immutable; new policy is prospective and documented as such.
- Owner content exposure — critical: excluded blueprint is never accessed beyond pathname metadata.


## Planned Work

1. Verify baseline, read governing material, create/verify the implementation snapshot, and produce a path-level authority/gating audit of Tasks 055–057 and current framework behavior.
2. Specify the Authority Envelope schema and STOP decision rule. Update the Constitution only if needed for top-level precedence; otherwise prefer narrower lifecycle, template, prompt, Git, skill and tooling changes.
3. Update Builder rendering/input validation, job template, lifecycle/prompt/Git guidance, local job skill and framework checks so every future approved task states its complete routine authority and reserved operations explicitly.
4. Add focused tests proving authority-envelope presence, deterministic rendering, backward-compatible historical validation, single-active lifecycle, targeted staging and failure semantics. Routine in-scope failures must be correctable/retestable; boundary failures must stop.
5. Define the practical Forgejo distribution architecture: immutable version tag + Forgejo Release + exact archive and sidecar assets + versioned direct-download endpoint. Read-only verify actual route/configuration where possible; document concrete URLs only when proven, otherwise document the verified template and publication prerequisites.
6. Add `wget`, `curl` and `sha256sum -c` user guidance plus a machine-consumable naming/URL/checksum contract for a future Smart Installer. Test documentation and release-plumbing assertions without publishing anything.
7. Add a prospective practical release-acceptance standard with one Owner authorization covering 12 sequential steps. Define mandatory evidence, Owner-only credential pause, recoverable reporting reconciliation, technical clean-host invalidation criteria, product-defect STOP and separate publication boundary.
8. Preserve historical Task 057 protocol/ledger without retroactive edits. Reference its limited outcome only as rationale and migration history.
9. Run focused and complete validation. Diagnose, correct and retest routine in-scope failures; stop only on the Authority Envelope's reserved/boundary conditions.
10. Review exact task paths/diffs/modes/security exclusions, target-stage only task paths, commit with a narrow subject, push dry-run and clean fast-forward push, verify direct Forgejo state, complete lifecycle, validate canonical idle and report. These routine repository operations are explicitly authorized by Task 058 once installed; no tag/Release/upload/publication is authorized.


## Rollback Plan

- Restore only exact Task 058-owned targets from the verified snapshot, after proving the target list and collision conditions. Remove only newly added paths whose prior absence and current Task 058 identity are proven.
- Never use broad reset, checkout, restore, clean, wildcard deletion, history rewrite, force push, tag mutation or Owner-content access as rollback.
- If an in-scope implementation/test failure occurs, preserve evidence, apply the smallest authorized correction and retest. Roll back only when correction is unsafe, no reliable path remains, or Owner directs it.
- If a commit exists but has not been pushed, prefer a bounded follow-up correction within the active task over history rewriting. Published commits are never rewritten.
- External distribution/acceptance rollback is not applicable because Task 058 performs no release, upload, VPS or candidate mutation. If external state changes unexpectedly, stop and obtain Owner direction.
- After rollback, rerun Git, lifecycle, Framework, Builder, security/privacy and target-integrity checks and report the exact restored state.


## Deliverables

- A reusable, tested Authority Envelope model embedded in canonical task definitions and Builder output, with explicit routine authority, bounded correction/retest behavior and Owner-reserved STOP conditions.
- Updated framework/lifecycle/Git/job guidance and tooling that remove artificial per-operation gates without weakening safety invariants.
- A prospective concise practical release-acceptance standard consisting of 12 product-relevant steps under one bounded Owner authorization, with privacy-safe evidence and technically meaningful stop/restart rules.
- A verified Forgejo distribution architecture and documentation contract for versioned archive/sidecar download using `wget` and `curl` plus `sha256sum -c`, ready for later separately authorized publication and future Smart Installer use.
- Focused and full validation, exact implementation history, snapshot/rollback evidence, task-scoped integration/push evidence and canonical lifecycle closure.
- No tag, Forgejo Release, artifact upload/publication, candidate construction, VPS access, credentials or Task 059.


## Verification

- Baseline Git/ref/lifecycle/version and Task 057 terminal-evidence preservation.
- Framework and Builder schema/rendering tests require a complete Authority Envelope and preserve deterministic prompt/history identity.
- Tests prove one approved task envelope can cover snapshot, implementation, bounded fixes/retests, explicit staging, commit, clean fast-forward push and lifecycle closure, while architecture expansion, destructive uncertainty, security/privacy issues, credentials, external mutation, elevated authority, tags, Releases, publication and deployment remain STOP/Owner-reserved.
- Existing single-active-task, No Task Without History, approval, idle, diversion, rollback and targeted-staging invariants remain passing.
- Practical acceptance standard contains exactly 12 coherent steps and one-run authorization semantics; it retains clean host, package integrity, Smart Install, install, setup, Guardian/state/service/evidence, real notification, reboot, explicit restart, uninstall preservation and same-candidate reinstall/final readiness.
- Evidence policy accepts reliable out-of-order reconciliation but never invents missing PASS evidence; host reinstall is required only when technical clean-host validity is destroyed.
- Official Forgejo documentation supports direct release-asset downloads; actual project route is verified read-only or explicitly left as a publication prerequisite. Download examples use exact filenames, fail-on-HTTP-error behavior where appropriate, and checksum verification.
- Future Smart Installer contract uses immutable versioned asset identities, authenticated metadata/checksum expectations and no mutable or invented URL assumption.
- Full applicable Framework, Builder, lifecycle, diversion, active-job/test-task, shell syntax, Git whitespace, formatting, documentation-link, security/privacy/secret and rollback validation passes.
- If Go/product source is unchanged, run the repository-mandated governance and relevant release-plumbing suites; do not add unrelated product test churn merely for ceremony.
- Review exact unstaged/staged path sets, complete diff, modes, `git diff --cached --check`, commit/push identity, clean fast-forward and post-push canonical idle state.


## Documentation Updates

- Update the smallest necessary reusable framework documents, likely `ai/core/08_JOB_TEMPLATE.md`, `ai/core/11_ENGINEERING_LIFECYCLE.md`, `ai/core/14_PROMPT_WORKFLOW.md`, `ai/core/16_GIT_POLICY.md`, `ai/core/03_AGENTS.md` and `.agents/skills/qwsg-job/SKILL.md`; modify `ai/core/01_CONSTITUTION.md` only if the authority precedence cannot be expressed coherently below it.
- Update `ai/scripts/task-builder.sh`, `ai/scripts/framework-check.sh`, `ai/scripts/next-task.sh` and their focused tests for the Authority Envelope schema and deterministic rendering.
- Add a new prospective practical release-acceptance standard rather than rewriting historical Task 057 protocol/ledger.
- Update installation/release documentation with the verified Forgejo archive/sidecar download contract and `wget`/`curl`/checksum examples only to the degree supported without publication.
- Update release-plumbing validation only if required to validate the new prospective standard and download contract while leaving RC.1–RC.5 chronology immutable.
- Maintain the Task 058 prompt/history throughout implementation and archive it canonically on completion. Do not create Task 059.


## Completion Criteria

Task 058 is complete when the reusable Authority Envelope is implemented and enforced across templates, Builder, workflow guidance and tests; routine in-scope engineering through clean-fast-forward integration and lifecycle closure no longer requires artificial intermediate Owner gates; all meaningful boundary and Owner-reserved STOP conditions remain explicit and tested; the 12-step prospective acceptance model is documented and validated without rewriting Task 057; the Forgejo Release-asset distribution contract is grounded in official behavior and the actual repository architecture, with safe `wget`/`curl`/checksum guidance and future Smart Installer requirements but no publication; all mandated validation and rollback checks pass; changes are integrated and pushed cleanly under the installed task's authority; Task 058 is archived and the repository is canonical idle. Any unverified concrete download URL remains a disclosed publication prerequisite, not an invented deliverable. No candidate, VPS operation, credential handling, tag, Release, upload, publication, deployment or Task 059 occurs.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-24 UTC.

The structured task definition and Authority Envelope have been explicitly approved. The task is authorized to start and execute every routine operation inside that envelope without another Owner gate. Further scope changes and every Owner-reserved operation require explicit Project Owner approval.
