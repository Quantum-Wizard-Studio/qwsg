# Current Engineering Task 094: QWSG 1.4.0 Signed Release Prepublication Verification and Controlled Publication

## Task Metadata

- Task ID: `094`
- Task slug: `signed-release-prepublication-verification-and-controlled-publication`
- Status: `complete`
- Date opened: `2026-09-28` UTC
- Human authority: Project Owner explicitly authorizes Task 094 in this session: independently verify the exact frozen signed 1.4.0 candidate, publish only if every release-authority check passes, record evidence and deterministic lifecycle/commit/push. No real-system acceptance.
- Owner or lead-developer communication language: Hungarian

## Title

QWSG 1.4.0 Signed Release Prepublication Verification and Controlled Publication


## Objective

Final immutable prepublication verification of the already signed Community 1.4.0 release, followed by controlled publication only if all release-authority conditions pass. One fresh session, current capable model, no expanded scope.


## Scope

Frozen source 996fb90d0f53d4ae6088cd884a654d9d0a32fff9; release/candidates/1.4.0; existing release machinery and runbooks; exact v1.4.0 tag, Forgejo Release/assets and canonical production endpoint; Task 094 evidence and lifecycle. VERSION 1.4.0; epoch 1790357147; build timestamp 2026-09-25T17:25:47Z. Artifact qwsg-1.4.0-linux-amd64.tar.gz, 3680796 bytes, SHA256 26cbf8d76fa27c652e11ee49f4570e05b3be6b969309e3cfd486977a27a6224d. Preserve historical 1.3.1.


## Out of Scope

No playground or production VPS mutation, clean installation acceptance, real 1.3.1 upgrade/bootstrap/rollback/Guardian acceptance, C9 closure, Community 1.0 completion, Pro P1/P3/P4/P5, new features, refactoring, packaging expansion or follow-on task.


## Authority Envelope

- **Task targets and boundaries:** Independently verify and conditionally publish only the exact frozen signed 1.4.0 candidate; evidence, lifecycle and deterministic Git integration.
- **Permitted external actions:** Fetch/push canonical main; read normal external release paths; publish exact authorized v1.4.0 tag, Forgejo Release/assets and production release-index through existing atomic machinery only after every prepublication condition passes.
- **Owner-reserved decisions:** Changed signed input or artifact, new signing, alternate version, unsupported migration/trust changes, real-system acceptance, scope expansion.
- **Task-specific STOP conditions:** Any frozen identity mismatch; invalid signature, trust anchor, compatibility, clock/watermark, provenance or atomicity; unexpected repository state beyond signing return; source/runtime changes required; unavailable safe rollback. Never repair signed bytes and publish under this authority.


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

Expected main and origin/main 07505afe91b324478e73b6ad996fa85644341662, divergence 0/0, idle after complete archived 093. Inspect signing return before snapshot; only expected signature, signed index and checkpoint may differ. Verify trusted frozen source and all authority bytes independently.


## Snapshot Requirements

After controlled signing-return classification, protected /tmp snapshot of tracked baseline and returned public authority files; deterministic hashes and readable archive; preserve until Owner acceptance and rollback dependencies expire. Separate verified previous production object backup before any index replacement.


## Risk Assessment

High release-authority risk: require exact hashes/signature/key/provenance/compatibility and current clock/watermark. High partial-publication risk: existing atomic staging and prior-object backup; external verification and established rollback. Preserve historical objects and anti-rollback semantics.


## Planned Work

Validate baseline/lifecycle and signing return; snapshot; verify complete immutable release authority; focused security negative checks; inspect established publication/rollback machinery and destinations; conditionally publish exact release set and tag; retrieve via public paths and reverify; record truthful result, lifecycle and deterministic commit/push. At budget thresholds 50/25/10 percent converge, limit to safety/evidence, and start no new publication mutation at 10 percent.


## Rollback Plan

Before live mutation preserve prior production bytes and verify digest, destination, hosting metadata and unique staging path. Use established atomic publication rollback only; never delete/retag historical 1.3.1 or bypass watermark. On inconsistency preserve evidence, stop and report PARTIAL/BLOCKED. Local snapshot restoration is exact-path and reviewed; no broad reset/clean.


## Deliverables

Exact verified publication if all conditions pass; independent external verification; bounded Owner authority and unchanged publication_authorized=false cryptographic checkpoint plus separate trusted publication evidence; Task 094 history, chronological milestone and truthful lifecycle. Otherwise blocked evidence without publication.


## Verification

Independently verify all 30 prepublication items and 25 acceptance items in Owner Task 094: repository/lifecycle, frozen source/version/build/epoch, artifact name/size/hash/sidecar, manifest/provenance, reproducibility, signing input, detached signature, verifier, signed index, production signature/key/fingerprint, exact target/platform/stable/active/artifact URL, exact 1.3.1->1.4.0 preserve-package-v1 with Configuration/Guardian/Scheduler 1.0 and Operator State 1.0-1.2, timestamps/current clock/watermark, tag target, notes, endpoint, immutable history, rollback/atomicity.
Exact authority: input 973 bytes SHA256 b44f83cb93f9068ab96a547e35e8544baaa45c6ef424281fd3cc63b95393554a; signature 89 bytes SHA256 128dd2c15eabeb644135e184785440c1c67565508335c60037da5590d811601f; verifier 5864474 bytes SHA256 d95604f9fb60a047300761fefa7dbbeded5b1ccb3afdb0386b3fb86f1cb3671d; signed index 1153 bytes SHA256 fbb39b52e240ff3691090a48a1e7915e516da6fb2f6a1faa9ed5fc069b633c41; checkpoint 434 bytes SHA256 2477382647ca0f6d5bd13863171a443b0210771f0cea490b54976059525ecf83. Manifest SHA256 1dfb3fe81d8767a83a11d9c1dd53b21d309c3d5d3a0fa2cba9dd2fbb0ee91001; RELEASE.json SHA256 df7cfb5d7dd404d22d26fac8420ea4cb17c0ab98fd8714fe89a32c2a901d9602. Production key qwsg-community-release-2026-01 fingerprint 0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6.
Reject unsigned, malformed/test-key signature, changed payload/artifact, unsupported compatibility and wrong anchor. After publication externally retrieve and verify exact artifact/index/manifest/provenance/signature/source/tag/migration authority, unsupported refusal, historical 1.3.1 and isolated safe awareness. Upload success alone is insufficient. No broad repeated source tests for unchanged frozen bytes.


## Documentation Updates

Task 094 prompt/history and chronological engineering milestone; release publication evidence and metadata only as required by existing architecture. Preserve all frozen and historical bytes. Owner report uses requested Task 094 result/prepublication/publication/external/safety/repository/gates/deferred structure.


## Completion Criteria

PASS only after all prepublication and external publication checks, atomicity and rollback evidence, lifecycle closure and deterministic commit/push. Otherwise PARTIAL/BLOCKED with exact missing evidence; never claim completion. Community 8 PASS/0 PARTIAL/1 MISSING, Pro 9 PASS/2 PARTIAL/3 MISSING. C9 remains MISSING pending later real clean install, official 1.3.1 authenticated bootstrap, supported state preservation, rollback, Guardian recovery and health/readiness evidence. Stop after reporting.


## Owner Approval Requirements

Approved by Project Owner explicitly authorizes Task 094 in this session: independently verify the exact frozen signed 1.4.0 candidate, publish only if every release-authority check passes, record evidence and deterministic lifecycle/commit/push. No real-system acceptance. through the Engineering Task Builder on 2026-09-28 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.

## Task 094 completion

Completed 2026-09-30 after Owner-confirmed execution of the credential-bound
root publication operation. Independent canonical external index/signature,
artifact/provenance/tag/compatibility/history and safe isolated discovery checks
pass. See matching history and release/candidates/1.4.0/PUBLICATION.md.
No real-system acceptance, C9 closure or successor task. Archived into the
canonical idle state; final integration checks precede Owner delivery.
