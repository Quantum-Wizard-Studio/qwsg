# Task History 063: Approved Task

## Task metadata

- Task ID: `063`
- Task slug: `qwsg-1-1-0-final-release`
- Status: `active`
- Date generated: `2026-08-26` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/063_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured Owner input. The Project Owner authorized full Task 063 execution under Framework 2.0 Standard Execution Authority on 2026-08-26 UTC. Implementation is active through final evidence integration and lifecycle closure.

## Starting state

- Ordinary user `attila`, canonical repository root, branch `main`.
- Framework `2.0.0` VALID; Task 063 was the sole approved active prompt/history
  pair; index and tracked worktree began clean at
  `1f26b7c16e542fdf72614269d1039bbf436d6bbf`.
- `HEAD`, `origin/main`, and direct Forgejo `refs/heads/main` matched that
  commit with `0/0` divergence. Remote/local tag `v1.1.0` was absent.
- `VERSION=1.1.0-rc.6`; final 1.1.0 notes, acceptance evidence and artifacts
  were absent. The unrelated Owner blueprint was excluded and unread.
- Accepted behavioral source
  `25a30718bc92882e9773a5c405ad648c0eee1a81` is an ancestor. The path diff
  through current HEAD contains only Task 061 evidence/lifecycle and Framework
  2 changes; behavior-bearing QWSG source and packaging differences are NONE.
- Task 061 evidence remains authoritative: practical acceptance PASS,
  QWSG-059-F001 externally corrected, no RC.6 product defect, controlled SMTP
  plus Owner receipt PASS, expected notification/overall partial semantics,
  and READY FOR RELEASE DECISION.
- Baseline execution first exposed a Framework v2 self-test fixture assumption:
  its incomplete-envelope negative case removed only a legacy label although
  the current Task 063 fixture uses the concise v2 envelope. Classified `TEST
  OR ACCEPTANCE DEFECT`; the test now removes the applicable label for either
  supported envelope form.

## Snapshot

- Execution snapshot: `/tmp/qwsg-task063-execution.1QW3IP`, mode `0700`.
- Manifest SHA-256:
  `f5d0076fcfabacd1177c0afaf4c2dc897ee4eccac05d821c1026c855fcf836ce`.
- Exact tracked-HEAD archive SHA-256:
  `35ba83361438f981a9fc60bb7252f80a50423088806f8c3df17d4aa730045d0c`.
- Contains exact HEAD payload, Task 063 lifecycle before-images, explicit
  pre-correction Framework test before-image/diff, accepted-lineage path
  evidence, release absence claims, exclusion record, hashes and bounded
  rollback instructions. Credentials, candidate bytes, caches and Owner
  content are excluded.
- The small Framework test-fixture correction occurred immediately before the
  execution snapshot while diagnosing the baseline suite. Exact pre-change
  content was recovered from immutable HEAD and included in the verified
  snapshot; rollback remained available throughout. This is recorded as a
  procedural ordering limitation, not missing rollback evidence.

## Work performed

Execution started. Corrected the concise/legacy Authority Envelope fixture
selection in `ai/tests/test-framework.sh`; no framework behavior changed.

- Changed final release identity from `1.1.0-rc.6` to `1.1.0`.
- Added final release notes, moved the completed Unreleased changelog content
  under the dated `1.1.0` heading, and changed canonical installation examples
  to the final immutable archive name.
- Extended `scripts/build-release.sh` so final `1.1.0` is accepted only with the
  same explicit epoch and full lowercase 40-character source commit required
  for the accepted RC.6 lineage. Updated positive, missing-metadata,
  short-commit and collision release-plumbing regressions.
- No `cmd/`, `internal/`, systemd unit, installer, uninstaller, release runtime
  configuration or operator behavior changed. Task 061 external product
  acceptance therefore remains valid and is reused without VPS access.

## Verification

Builder input, metadata, prompt/history identity, approval state, lifecycle
installation, canonical refs, accepted lineage, release absence and snapshot
integrity validated successfully. Full baseline validation resumes after the
test-fixture correction.

- Corrected baseline: Framework/configured validations PASS (Framework 25,
  Builder 49, v2 semantics 15, diagnostic runner 11, full Go, vet, formatting),
  repository-wide Go race PASS, lifecycle 29, diversion 36, test-task audit,
  RC.6 release plumbing, build contract, shell syntax, Git whitespace and
  systemd static verification PASS. The static systemd tool emitted the known
  sandbox bus warnings while returning success.
- Final release-source validation: final release plumbing, strict provenance
  negative checks, configured Framework/full Go/vet/format suite, build
  contract, shell syntax, Git whitespace and static unit verification PASS.
- Release-source change classification: release identity/documentation/plumbing
  only; runtime behavior mutation NONE; Task 061 acceptance reuse VALID.

## Release-source integration and deterministic construction

- Exact release-source commit:
  `305f4088e94b14d6cbb3114eb8cce4e32d847c16`, tree
  `696127b54f2c1e39289fa3e020a5b7e23de9694a`, epoch `1787752902`, UTC
  `2026-08-26T14:01:42Z`.
- Commit subject: `release: establish QWSG 1.1.0 source`. Dry-run and actual
  push clean-fast-forwarded `main`; direct Forgejo main matched the full commit.
- Two separate `git archive` exports used distinct source roots, output roots
  and Go build caches under private
  `/tmp/qwsg-task063-final-build.TuJFQg`. Archive and sidecar bytes are
  byte-identical.
- Twin one was selected exactly once and its archive/sidecar frozen mode `0400`.
  Frozen identity: archive `qwsg-1.1.0-linux-amd64.tar.gz`, size `2951638`,
  SHA-256
  `10a39d96b93b72a3f4799a76d769bc264afd6845a32a1ecc5531b062d6f42349`;
  sidecar SHA-256
  `b9414bba5a6d9bc100f7c391c11867ceda6c2139002272a0f46fbf55dc9d3cc1`;
  manifest SHA-256
  `310d41f9a8c71599290fd1d25efb7a2da8fd210e34cbf2666e40189c988ebc3d`;
  binary SHA-256
  `e7b5a2234221baa32a9c3fa79e0758ea49e7ed1996c99e9e1ddbc19628a5e924`.
- Sidecar, safe unique 25-member root, regular/directory types, no symlinks,
  all 18 manifest entries, LICENSE/docs/final notes, modes, static amd64,
  version/full commit/build time and absence of ambient VCS metadata PASS.
- Extracted local Smart Install correctly found its current user manager
  unreachable and returned not-ready while satisfying filesystem/platform/
  runtime checks. Classified `ENVIRONMENTAL ISSUE`; Task 061's supported-host
  Smart Install PASS is reused because no relevant behavior changed. No
  candidate byte changed or rebuild occurred.
- Added `docs/release/ACCEPTANCE_1.1.0.md` with frozen identities, package
  proof, Task 061 reuse matrix, limitations and the conditional publication
  ledger. Pre-publication technical verdict: `READY FOR RELEASE`.

## Publication attempt and boundary

- Release-boundary snapshot:
  `/tmp/qwsg-task063-release-boundary.tkqAUW`, mode `0700`, manifest SHA-256
  `1dd5c74b029112b8a1cda57942ffd2cbd2291063a2e3fa2e13bcc7b3e59d3edf`;
  exact release-source archive SHA-256
  `b3d6a5d24d09f88fe0f671df32988a9a705c2ec03973f3eec916db4b43c4041c`.
- Annotated `v1.1.0` tag object
  `b14347636f6c9873a5acf759c950d900a39bf1a7` was created and pushed. Direct
  Forgejo verification shows it peels exactly to release-source commit
  `305f4088e94b14d6cbb3114eb8cce4e32d847c16`.
- Forgejo Release ID `2`, title `QWSG 1.1.0`, tag `v1.1.0`, final/non-draft/
  non-prerelease was created. It contains exactly the frozen archive (size
  `2951638`) and sidecar (size `96`); no other asset was uploaded.
- The API-provided canonical versioned asset URLs match the documented
  distribution route. Anonymous requests for both exact URLs return HTTP
  `404`; therefore wget/curl integrity verification cannot begin.
- Protected read-only diagnosis confirms `repository.private=true`. Exact
  classification: `ENVIRONMENTAL ISSUE` requiring an Owner-reserved external
  infrastructure decision. Changing canonical repository visibility is not
  authorized by Task 063 and would affect more than the Release assets.
- No tag move/deletion, Release mutation, asset replacement, visibility change,
  deployment or release-success claim occurred after the boundary. Temporary
  protected Forgejo authentication material was deleted and absence verified.

## Authorized public distribution and final verdict

- The Project Owner subsequently authorized the canonical QWSG repository-only
  visibility change required by the existing distribution contract. A first
  repository-scoped API attempt produced Forgejo's `internal` state because the
  owning organization was private; anonymous access remained `404`. A rotated
  credential was entered locally without placing it in task evidence. Forgejo
  denied the internal-to-public transition to the repository-admin credential,
  and Owner intervention completed the external visibility decision.
- Final authoritative anonymous checks confirm repository visibility `PUBLIC`,
  anonymous repository readability PASS, anonymous Release-page accessibility
  PASS, and Release API accessibility PASS.
- Annotated tag object
  `b14347636f6c9873a5acf759c950d900a39bf1a7` remains unchanged and peels to
  `305f4088e94b14d6cbb3114eb8cce4e32d847c16`. The Forgejo Release remains
  final, non-draft and non-prerelease with exactly two existing assets.
- Independent credential-free clean-environment `wget` and `curl -fLO`
  downloads both succeeded. Each archive is `2951638` bytes and SHA-256
  `10a39d96b93b72a3f4799a76d769bc264afd6845a32a1ecc5531b062d6f42349`;
  each sidecar SHA-256 is
  `b9414bba5a6d9bc100f7c391c11867ceda6c2139002272a0f46fbf55dc9d3cc1`;
  `sha256sum -c` passed for both clients. Downloaded pairs were byte-identical.
- No tag, Release metadata, asset count or release byte changed during the
  visibility work. Temporary authentication material and anonymous download
  directories were removed.
- Final verdict: **QWSG 1.1.0 RELEASED**.

## Rollback

Use `/tmp/qwsg-task063-execution.1QW3IP/ROLLBACK.txt` only after exact identity
and collision checks. Restore literal targets from the tracked-HEAD archive or
explicit before-images; never use broad reset/checkout/clean or touch excluded
Owner content.

## Completion state

`active — QWSG 1.1.0 released; exact frozen assets anonymously downloadable and verified; final evidence integration and canonical archive/idle closure remain`
