# Current Engineering Task 054: QWSG Exported-Source Build Determinism Fix

## Task Metadata

- Task ID: `054`
- Task slug: `qwsg-exported-source-build-determinism-fix`
- Status: `approved`
- Date opened: `2026-08-21` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG Exported-Source Build Determinism Fix


## Objective

Correct `QWSG-053-F001` by making the canonical ordinary `make build` path
operate deterministically both in a normal Git checkout and in an independent
source export with no `.git` metadata. Align ordinary builds with the existing
release-builder VCS-stamping policy without weakening explicit release
provenance, changing runtime behavior, or silently replacing the failed private
RC.3 candidate. Integrate the narrow correction, advance the next private
candidate identity to `QWSG 1.1.0-rc.4`, and stop before candidate construction
at a separate Owner gate.


## Scope

- Confirm the current architecture: the Makefile derives an ordinary
  `BUILD_COMMIT` from Git when available or uses the truthful `unknown`
  sentinel, but its `go build` permits automatic ambient VCS stamping; the
  canonical release builder already uses `-buildvcs=false` and requires an
  explicit full commit plus commit-derived `SOURCE_DATE_EPOCH` for 1.1 RCs.
- Apply the smallest repository-consistent build-policy correction: make the
  canonical Makefile Go build invocation explicitly disable automatic VCS
  stamping. Do not require a caller-supplied `GOFLAGS` workaround and do not
  infer or invent provenance.
- Preserve the ordinary-build model: `VERSION` is embedded; `BUILD_COMMIT`
  defaults to a Git-derived development identity in a checkout and to
  `unknown` in a source export; `BUILD_DATE` defaults to `unknown`; explicit
  make-variable values are embedded exactly after bounded validation/testing.
- Preserve the release-build model: release artifacts continue to accept only
  controlled source, explicit exact full lowercase commit and explicit
  commit-derived epoch, use `-trimpath -buildvcs=false`, and expose the exact
  resulting version/commit/UTC time. No ambient `.git` metadata may alter
  release output or substitute for explicit provenance.
- Add focused automated coverage that creates independent clean source exports
  without `.git`, runs ordinary builds with no `GOFLAGS`, verifies successful
  truthful `unknown` defaults, verifies explicit version/commit/date injection,
  and compares outputs across controlled contexts where deterministic identity
  inputs match. Tests must prove ambient VCS metadata does not add or change Go
  VCS build settings.
- Verify ordinary builds in the real checkout and at least two independent
  exported roots, including exact `qwsg version` output and `go version -m`
  inspection. Keep all generated outputs/caches outside tracked source or in
  existing ignored build locations.
- Preserve and extend release-plumbing coverage so `make release-check`, exact
  full release commit requirements, commit-derived epoch behavior, collision
  refusal, twin release reproducibility and package provenance remain intact.
- Advance `VERSION` and matching private release metadata/plumbing from
  `1.1.0-rc.3` to `1.1.0-rc.4`. RC.3 source, candidate hashes, reproducibility
  evidence and `QWSG-053-F001` remain immutable historical failed-gate
  evidence. Never overwrite or relabel RC.3 bytes.
- Use explicit Owner gates: implementation may prepare the exact correction
  and tests locally; source integration requires separate authorization with a
  reviewed path allowlist, commit and clean fast-forward push. Task 054 closure
  is separate. RC.4 construction and clean-host acceptance require a later
  independently authorized acceptance task from one exact clean commit.


## Out of Scope

- No runtime product, Smart Install, assessment, guided setup, Guardian,
  notification, systemd runtime, installer/uninstaller, configuration,
  credential, Community/Pro or privilege-model behavior changes.
- No arbitrary environment inheritance, shell interpretation, network access,
  new dependency, sudo, package installation, lingering or privilege escalation.
- No `GOFLAGS` workaround as the product contract and no reliance on ambient
  `.git` metadata for release identity.
- No modification, deletion, relabeling, transfer or execution of the existing
  private RC.3 candidate. No RC.3 rebuild and no RC.4 artifact construction.
- No external VPS, SSH/SCP, SMTP credential, Guardian acceptance, external
  checkpoint, tag, Forgejo Release, upload, publication or announcement.
- No external resolution claim for `QWSG-051-F001`; local Task 054 work cannot
  prove the Guardian correction. No Task 055 creation.
- Do not read, copy, hash, stage, modify, move, delete or otherwise touch
  Owner-owned `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`; metadata-only
  exclusion checks are allowed.


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

- Verify UTC date, ordinary user, exact repository root and `main`; canonical
  HTTPS origin; `HEAD == origin/main ==` direct remote main at
  `3e7d2f9d543078b49d1afa522dc6bf3baba1c949`; ahead/behind `0/0`; empty index.
- Verify the only tracked changes are the unstaged Task 053 history and RC.3
  acceptance-ledger evidence authorized during the stopped Gate B/C/D, and the
  only unrelated untracked content is the excluded Owner file. Validate those
  records before Task 053 closure.
- Verify Task 053 remains active and stopped, `QWSG-053-F001` is OPEN/BLOCKING,
  the final verdict is `NOT READY FOR QWSG 1.1.0 RELEASE`, and no transfer or
  external checkpoint occurred.
- Verify and preserve the exact RC.3 source commit, epoch, archive/sidecar/
  manifest/binary hashes, twin-build and package evidence, and private failed-
  gate candidate locations without packaging or exposing them.
- Verify `VERSION=1.1.0-rc.3`; Makefile lacks explicit `-buildvcs=false` while
  `scripts/build-release.sh` contains it; reproduce the exported-source failure
  once in a new bounded mode-0700 export without using `GOFLAGS`.
- Verify immutable RC.1/RC.2/Task 049/Task 051/v1.0.0/LICENSE identities and
  Owner exclusion. Verify no RC.4 artifact, sidecar, tag or release exists.


## Snapshot Requirements

Before lifecycle or implementation changes create a unique private mode-0700
snapshot with a readable tracked-HEAD archive; Task 053 prompt/history and
RC.3 ledger including unstaged evidence; Task 053 Builder source/input;
Makefile, release scripts/tests and proposed Task 054 targets; Git/mode/ACL/tool
metadata; RC.1/RC.2/RC.3/F001/F002/F003/v1.0.0/LICENSE preservation hashes;
absence records for new targets; and literal bounded restore instructions.
Record Owner content only as excluded stat metadata. Do not include candidate
bytes, credentials or private external identity in lifecycle snapshots.


## Risk Assessment

- **False provenance — critical:** disabling automatic VCS stamping must not
  cause ordinary or release builds to claim a commit that was not explicitly
  derived or supplied; `unknown` remains the truthful export default.
- **Release determinism regression — critical:** release output must retain
  exact explicit full-commit/epoch provenance and twin byte reproducibility.
- **Candidate collision — critical:** changed source and embedded commit require
  RC.4; RC.3 bytes and evidence are immutable and must never be overwritten.
- **Incomplete exported-source test — high:** tests must use genuine exports
  without `.git` and no ambient `GOFLAGS`, not a simulated checkout path.
- **Scope creep — high:** a Makefile/build-test/release-metadata correction must
  not alter runtime or external acceptance behavior.
- **Lifecycle evidence loss — critical:** Task 053 must close truthfully as NOT
  READY with F001 and all passed/failed Gate B/C/D evidence intact.
- **Owner-content exposure — critical:** the excluded Owner path is never read,
  hashed, copied, staged or packaged.


## Planned Work

1. Under separately authorized lifecycle installation, validate and close Task
   053 as complete with disclosed limitations — `NOT READY FOR QWSG 1.1.0
   RELEASE`; preserve `QWSG-053-F001` OPEN/BLOCKING and all RC.3 evidence;
   archive Task 053 and install exactly one approved Task 054 pair.
2. Revalidate the source/build baseline, reproduce the no-`.git` ordinary-build
   failure in an isolated mode-0700 export, and create the rollback snapshot.
3. Add explicit automatic-VCS-stamping disablement to the canonical Makefile
   Go build path without changing its truthful linker-variable model.
4. Add focused build-contract tests for checkout/export success, no-`GOFLAGS`,
   truthful defaults, explicit identity injection, ambient VCS independence and
   deterministic controlled output.
5. Advance source identity and matching private release notes/plumbing to
   `1.1.0-rc.4`, preserving explicit RC.1/RC.2/RC.3 historical checks and
   collision guards.
6. Run build, focused/full/race, vet, formatting, exported-source, deterministic
   twin-build, release-plumbing, shell/security/exclusion, Framework, Builder,
   lifecycle, diversion, job/test-task, whitespace and preservation suites.
7. Report the exact implementation allowlist and stop at a separate Owner
   source-integration gate. Do not build RC.4 candidate bytes.


## Rollback Plan

- Restore only literal Task 054-owned paths from the verified private snapshot
  after Owner authorization. Remove only new files proven absent before Task
  054. Never use broad reset, checkout or clean operations.
- If build tests fail, preserve bounded privacy-safe logs and exact inputs;
  revert only the narrow Makefile/test/metadata edits or stop for Owner review.
  Do not set ambient GOFLAGS or alter Git metadata to manufacture a PASS.
- Never delete or mutate the existing RC.3 private roots or historical records;
  never touch Owner content, tags, releases, credentials or external systems.
- After rollback rerun build-contract, release-plumbing, governance, Git and
  preservation checks and report exact remaining state.


## Deliverables

- Canonical ordinary build policy that succeeds in checkout and clean exported
  source trees without ambient VCS dependence or GOFLAGS workaround.
- Focused automated exported-source/determinism/provenance regression tests.
- Truthful ordinary-build default and explicit identity evidence from binary
  version output and Go build metadata.
- Preserved reproducible release-builder provenance and release-check evidence.
- `VERSION=1.1.0-rc.4` with matching private notes/plumbing and immutable RC.3
  failed-gate evidence.
- Complete local/regression/security/governance/preservation evidence and an
  exact source-integration allowlist; no candidate artifact.


## Verification

- `make build` succeeds in the normal checkout with no GOFLAGS workaround and
  embeds the expected development version/commit/date model.
- `make build` succeeds independently in at least two Git exports containing no
  `.git`; defaults report commit/date `unknown` truthfully.
- Explicit `BUILD_COMMIT` and `BUILD_DATE` values are embedded exactly in an
  export; tests use fixed safe values and reject/inhibit ambient VCS metadata.
- Controlled ordinary outputs with identical source/toolchain/identity inputs
  are byte-identical; `go version -m` contains no unintended `vcs.*` settings.
- The canonical release builder still requires explicit full lowercase commit
  and `SOURCE_DATE_EPOCH`, embeds their exact version/commit/UTC identity, uses
  disabled VCS stamping, refuses collisions, and produces twin byte-identical
  binary/manifest/archive/sidecar output.
- `make release-check`, package/layout/manifest/LICENSE/docs/static/exclusion
  checks and RC.1/RC.2/RC.3 preservation checks pass for RC.4 plumbing without
  constructing a repository or acceptance candidate.
- Full tests, repository race tests, vet, formatting, focused build tests,
  shell syntax/static checks, secret/privacy scan, Git whitespace, Framework
  21, Builder 38, lifecycle 28, diversion 36, active-job and test-task suites
  pass.
- No runtime source or external state changes; no RC.4 artifact/tag/release;
  Task 051 F001 remains historical OPEN/BLOCKING pending external acceptance.


## Documentation Updates

- Update `Makefile` only for the canonical ordinary build policy.
- Add or extend a focused build-contract test script under `scripts/` and wire
  it into the appropriate Make/validation target if needed.
- Update the narrow development build documentation only if required to state
  truthful checkout/export provenance; do not document GOFLAGS as a workaround.
- Update `VERSION`, add `docs/release/RELEASE_NOTES_1.1.0-rc.4.md`, and extend
  `scripts/test-release-plumbing.sh` for RC.4 while retaining RC.1–RC.3 history.
- Update Task 054 history during work. Preserve Task 053/RC.3 records unchanged
  after their authorized archival; do not create RC.4 acceptance records or
  Task 055 during Task 054.


## Completion Criteria

Task 054 completes only when the ordinary canonical Makefile build succeeds
without GOFLAGS in both a normal checkout and genuine no-`.git` exports;
truthful default and explicit provenance are proven; ambient VCS metadata cannot
alter controlled output; release provenance and twin reproducibility remain
strict; full regression/security/governance/preservation validation passes;
source version and private metadata consistently identify `1.1.0-rc.4`; the
correction is integrated from an exact reviewed allowlist under separate Owner
authority; and no RC.4 candidate, transfer, external acceptance, tag, release,
upload, publication or Task 055 exists. `QWSG-053-F001` may be locally corrected
by this evidence but its historical Task 053 record remains unchanged; external
Guardian blocker `QWSG-051-F001` remains unresolved until later RC.4 clean-host
acceptance.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-21 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
