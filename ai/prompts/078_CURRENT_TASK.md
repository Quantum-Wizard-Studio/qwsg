# Current Engineering Task 078: Community Release Authority & Production Manifest Endpoint

## Task Metadata

- Task ID: `078`
- Task slug: `community-release-authority`
- Status: `active`
- Date opened: `2026-08-30` UTC
- Human authority: Project Owner — approved continuation on 2026-08-30 UTC
- Owner or lead-developer communication language: English

## Title

Community Release Authority & Production Manifest Endpoint

## Objective

After explicit Project Owner approval, connect the existing Task 075 installation classifier, Task 076 authenticated `qwsg.release-index/1` discovery contract, and Task 077 persistent update-awareness/operator-command path to one explicitly authorized Community production release authority. Define, approve, document, implement, publish, and verify the exact official HTTPS release-index endpoint, its canonical hosting/publication ownership, the initial Ed25519 signing-key identity and custody model, the bundled public trust anchor, rotation and emergency-revocation procedures, deterministic index generation/signing in the canonical Forgejo release lifecycle, authenticated production `qwsg update check`, and fail-closed behavior.

Task 078 ends only when the Owner-approved authority design is implemented and deterministic local verification plus one explicitly authorized real production acceptance test pass. It does not install updates, schedule checks, create a second updater, or begin Task 079.

## Scope

- Produce an Owner-reviewable authority decision record before activation. It must name the exact official `https://` URL, DNS/hosting owner and serving boundary, canonical Forgejo publication source, deployment method, availability assumptions, initial Ed25519 public-key fingerprint/key ID, trust-anchor encoding/distribution path, and operational roles. Do not leave aliases such as “the production endpoint” unresolved in the accepted design.
- Inspect and reuse Task 075 installation classification, Task 076 release-index schema/canonicalization/signature/transport/compatibility/artifact-integrity contract, and Task 077 update-awareness/check/status behavior. Extend one canonical path only where production authority requires it.
- Implement deterministic release-index generation, canonical serialization, Ed25519 signing, verification, reproducibility checks, and publication plumbing as part of the canonical Forgejo-controlled release lifecycle. Preserve the separation between metadata authenticity and artifact integrity.
- Establish secure offline or dedicated release-authority private-key generation, storage, least-privilege access, encrypted backup, recovery testing, compromise response, and destruction/retirement rules. Private key material must never enter Git, ordinary runtime hosts, logs, command-line evidence, build artifacts, CI caches, or task records.
- Bundle and document the Owner-approved public trust anchor with QWSG, including stable key identity/fingerprint semantics and strict trust selection. Define backward-compatible rotation, overlap, rollback protection, and emergency revocation without accepting an unsigned network-provided replacement trust root.
- Activate `qwsg update check` against only the exact approved production authority and endpoint. Community use must require no registration, account, API key, mandatory telemetry, inbound listener, or communication other than outbound HTTPS.
- Add strict fail-closed handling for TLS/transport, redirect/origin policy, timeout/size/resource bounds, HTTP semantics, schema/version/canonicalization, signature/key trust, rollback/freshness, compatibility, withdrawal, cache/304, and artifact-metadata validation failures. Preserve the last valid authenticated awareness evidence according to Task 077 while clearly recording the failed attempt.
- Add deterministic unit, integration, fixture, failure-injection, release-lifecycle, security, privacy, reproducibility, and CLI tests, plus one narrowly scoped real production acceptance test only after the endpoint, trust anchor, publication, and acceptance action are explicitly authorized.
- Update directly affected architecture, release, security, operator, lifecycle, and engineering records; perform task-scoped Git integration and canonical lifecycle closure after all gates pass.

## Out of Scope

- Guardian periodic scheduling or coupling update-source availability to Guardian health/readiness.
- Notification generation, transition tracking, delivery, or deduplication.
- Automatic or unattended acquisition/installation, migration, replacement, rollback redesign, or a second updater.
- Pro licensing, registration, accounts, API keys, fleet management, central or mandatory telemetry, QUWIP, Telegram, or unrelated integrations.
- Forgejo HTML scraping, GitHub release authority, GitHub-origin publication, or treating the public GitHub mirror as canonical. The Quantum Wizard Studio Forgejo repository remains canonical; GitHub remains read-only.
- Inbound communication to Community installations, non-HTTPS production transport, or mandatory runtime credentials.
- Modifying, rebuilding, resigning, retagging, replacing, or republishing the accepted QWSG 1.2.0 release or its artifacts.
- Unrelated infrastructure, release-policy, updater, installer, Guardian, notification, or product work.
- Task 079 or any later milestone.

## Authority Envelope

**Task targets and boundaries:** After Owner approval, inspect and modify only canonical release-authority/discovery/awareness/release-lifecycle code, fixtures, packaging trust-anchor assets, narrowly required configuration, tests, scripts, and directly affected documentation/lifecycle records. Reuse Tasks 075/076/077; do not duplicate their classifier, discovery client, awareness store, commands, updater, or integrity mechanisms. Ordinary reversible local implementation, correction, retesting, documentation, targeted Git integration, and lifecycle closure are permitted only within this scope.

**Permitted external actions:** Read-only origin synchronization and ordinary approved Git integration; DNS/TLS/hosting inspection needed for the decision record; and, only after the Project Owner explicitly approves the exact endpoint, hosting/publishing plan, trust-anchor identity, private-key ceremony, production publication action, and acceptance protocol, the narrowly defined endpoint deployment/publication and one outbound-HTTPS real acceptance test. No GitHub write, release mutation, unrelated infrastructure mutation, runtime-host private-key placement, account creation, telemetry, or production installation is authorized.

**Owner-reserved decisions:** Approval of this Task 078 definition; selection and activation of the exact official endpoint and hosting boundary; acceptance of the initial Ed25519 public-key fingerprint/key ID and custody/backup/recovery roles; authorization of any private-key ceremony; approval of rotation/revocation policy; authorization of DNS/TLS/hosting mutation and first production index publication; approval of the exact real production acceptance target/protocol; release/tag/artifact decisions; scope expansion; and Task 079 start. No endpoint or trust anchor may be inferred, silently selected, generated as authoritative, or activated before these decisions are recorded.

**Task-specific STOP conditions:** Stop before implementation until this prepared task is explicitly approved. During execution stop at any unresolved endpoint/hosting/trust/key-custody decision; risk of private-key exposure; inability to establish exclusive authority and rollback; need to weaken Task 076 authenticity or Task 075/077 contracts; redirect/cross-origin ambiguity; unapproved DNS, hosting, Release, tag, artifact, or publication mutation; changed protected QWSG 1.2.0 evidence; inability to fail closed; or any need for scheduling, telemetry, automatic installation, a second updater, GitHub authority, unrelated infrastructure, or Task 079.

## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/12_RELEASE_POLICY.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/core/17_EXECUTION_MODEL.md`
- `ai/core/18_BOUNDED_DIAGNOSTIC_RUNNER.md`
- `ai/config/engineering-project.conf`
- `ai/archive_prompts/074_2026-08-29_update-discovery-release-awareness-architecture.md`
- `ai/history/074_2026-08-29_update-discovery-release-awareness-architecture.md`
- `ai/archive_prompts/075_2026-08-29_installed-package-legacy-installation-classification.md`
- `ai/history/075_2026-08-29_installed-package-legacy-installation-classification.md`
- `ai/archive_prompts/076_2026-08-29_release-index-metadata-authenticity-release-source-contract.md`
- `ai/history/076_2026-08-29_release-index-metadata-authenticity-release-source-contract.md`
- `ai/archive_prompts/077_2026-08-30_update-awareness-state-operator-check-status.md`
- `ai/history/077_2026-08-30_update-awareness-state-operator-check-status.md`
- Task 075/076/077 implementation, tests, fixtures, and directly affected architecture/release/security/operator documentation identified during inspection.

## Starting State Verification

Preparation baseline recorded before repository mutation: repository `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, branch `main`, `HEAD == origin/main == e2b95f9138e07bd26e7882bd290fb97e761fc2ed`, divergence `0/0`, clean index/work tree, canonical origin `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg.git`, and canonical idle with Task 077 latest complete. QWSG `VERSION` is 1.2.0. Protected annotated tag object `ac395b568b8e1f83c0ef85c9aa02f98c15402af0` peels to release commit `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`.

At approved task start, re-verify repository identity, user/date, branch, HEAD/origin/divergence, complete status/index, remotes, framework/lifecycle/diversion/Builder state, VERSION, Task 074–077 identities/completion, Task 075 classifier, Task 076 contract/fixtures, Task 077 state/commands, canonical updater/release scripts, ownership/modes/ACLs, current source-authority refusal, and protected QWSG 1.2.0 tag/release/artifact identities. Fetch only when permitted. Diagnose recoverable in-scope variance; stop on authority, security, lifecycle, release, rollback, or unexplained material divergence.

## Snapshot Requirements

Retain and verify the preparation snapshot `/tmp/qwsg-task078-preparation.nYKZti`, mode 0700 with mode-0600 evidence, complete readable Git bundle, lifecycle before-images/absence, repository/ref/status, Task 077 hashes, modes/ACLs, protected v1.2.0 identities, checksums, and bounded preparation rollback.

After approval and before implementation or external mutation, create a new unique mode-0700 execution snapshot under `/tmp`. Record the approved prompt/history identities, exact target before-images or absence, Git/ref/status/lifecycle/framework evidence, release/discovery/awareness contract hashes, trust-anchor and endpoint pre-state/absence, permissions/ACLs, complete Git bundle, authorized publication/DNS/hosting pre-state, and literal bounded restore instructions. Evidence files are mode 0600. Exclude private keys, passphrases, credentials, tokens, private host/account identifiers, caches, generated secrets, and sensitive values. Before public publication create a separately verified publication checkpoint; before lifecycle closure create a closure snapshot. Verify hashes, readability, bundle completeness, collision safety, exclusions, permissions, and rollback at each gate.

## Risk Assessment

- **Wrong or ambiguous authority — critical:** clients could trust an attacker or unintended mirror. Require an Owner-approved exact HTTPS origin/path, pinned bundled trust identity, strict redirect/origin policy, canonical Forgejo provenance, and fail-closed selection.
- **Private-key compromise or loss — critical:** compromise permits forged metadata; loss can halt releases. Under the Owner-approved current-stage custody model, require dedicated-workstation generation, a strong unique passphrase controlled separately, least-privilege ACLs, exclusion from every cloud-sync/repository/runtime/hosting/CI boundary, and one separately stored encrypted recovery copy with a completed recovery verification before first production publication. Use fingerprint-only public audit records and documented revoke/rotate/retire procedures.
- **Key leakage — critical:** never place secrets in Git, argv where avoidable, logs, CI caches, runtime hosts, snapshots, fixtures, chat, artifacts, or ordinary workspaces; deterministic tests use public non-production test keys only.
- **Bootstrap/rotation/revocation weakness — critical:** network replacement keys could be forged or rollbacked. Use bundled anchors, signed transition rules, monotonic key epochs, overlap, emergency procedure, and recovery compatibility.
- **Publication non-determinism or split authority — high:** mitigate with canonical serialization, reproducible generation, isolated signing inputs, byte/hash comparison, no-clobber publication, provenance, and post-publication retrieval verification.
- **Transport/cache/redirect ambiguity — high:** reject stale, cross-origin, downgraded, oversized, or unauthenticated responses through HTTPS-only bounded clients, strict status/cache semantics, and authenticated cached bodies.
- **Compatibility or artifact confusion — high:** preserve Task 075 compatibility and separate artifact-integrity verification; authenticity never substitutes for artifact hashing.
- **Availability coupled to runtime health — medium:** keep checks explicit and awareness-only; preserve prior authenticated success and never change Guardian health/readiness.
- **Unintended release/infrastructure mutation — critical:** protect v1.2.0 and require Owner checkpoints for endpoint, DNS/TLS/hosting, key ceremony, first publication, and production acceptance.
- **Rollback lockout — high:** require staged rotation, compatibility fixtures, reversible publication, retained public history, and tested recovery.

## Planned Work

1. Revalidate authority, starting state, lifecycle, Tasks 074–077, release policy, protected v1.2.0 evidence, and canonical Forgejo/GitHub boundary.
2. Create and verify the approved-task execution snapshot.
3. Inventory the single Task 075/076/077 path, release tooling, canonicalization/signature code, configuration, updater integration, packaging, and documentation; identify the smallest production-authority seam.
4. Prepare a concrete decision record with the exact endpoint, hosting/publishing architecture, initial public-key fingerprint/key ID, key ceremony/custody/backup/recovery, trust-anchor distribution, rotation, emergency revocation, publication transaction, rollback, and acceptance protocol. Present it to the Owner and stop at the reserved decision gate; record approval without private material.
5. After recorded Owner decisions, add the public trust anchor and strict authority configuration. Keep test private keys clearly non-production and deterministic; never generate/store the production private key on ordinary runtime hosts.
6. Extend deterministic release-index generation and canonical signing in the existing release lifecycle with isolated inputs, reproducible unsigned bytes, stable signatures, artifact-integrity references, validation, no-clobber publication, and provenance.
7. Implement production discovery through the existing Task 076 client and Task 077 `update check`; preserve network-free `update status`, last authenticated success, classifier/awareness semantics, and updater boundaries.
8. Implement fail-closed transport, redirect/origin, HTTP/cache, resource, schema, canonicalization, signature/key trust, epoch/rollback/rotation/revocation, compatibility, withdrawal, and artifact-metadata validation with privacy-safe reason tokens.
9. Document secure production key generation and primary custody on the Owner's dedicated local workstation outside cloud-synchronized folders; require a separately stored encrypted recovery copy and successful recovery audit before first production publication; document rotation overlap, compromise revocation, loss recovery, and retirement/destruction without putting private material in the repository.
10. Add deterministic unit/integration/CLI/release/reproducibility/security/privacy/fault tests using loopback fixtures and non-production test keys. Prove no registration, telemetry, inbound communication, unattended installation, Guardian coupling, scraping, GitHub authority, or second updater.
11. At the explicitly approved publication gate, verify the checkpoint, perform only approved endpoint/key/index publication actions, retrieve/authenticate exact public bytes, and preserve privacy-safe hashes/fingerprints/status evidence.
12. Run the explicitly approved real production protocol from an authorized isolated Community-equivalent installation: outbound HTTPS only; authenticated check; network-free status; expected relationship; safe negative fail-closed probes; no download/install or Guardian-health change.
13. Synchronize documentation/history, run full validation, review scope/secrets/modes/ACLs, perform targeted Git integration, create closure snapshot, close lifecycle to idle, and stop without Task 079.

## Rollback Plan

Before rollback, stop task-owned local processes; preserve failure evidence; verify snapshot manifest/bundle/checksums, repository root, lifecycle identity, exact target inventory, publication checkpoint, and absence of later Owner edits. Restore only named repository before-images/modes/ACLs and remove only Task 078-created paths whose prior absence and current identity are proven. Never broad reset, checkout, restore, clean, wildcard-delete, rewrite history, force-push, or modify Task 077/QWSG 1.2.0 evidence.

For production rollback, use only the Owner-approved endpoint transaction: retain the last authenticated valid index and public trust history; revert exact published index/object/config/DNS targets only when prior identity is proven; never destroy the sole usable backup or publish unsigned fallback metadata; never revoke/remove an anchor unless compatible recovery is proven and explicitly authorized. Suspected compromise stops ordinary rollback and invokes the approved emergency-revocation procedure.

After rollback verify repository/status/refs, lifecycle/framework, snapshot identities, permissions/ACLs, Tasks 075–077, packaged trust state, endpoint authentication/fail-closed behavior, release reproducibility, protected v1.2.0, absence of secrets, and external publication state. Retain evidence for Owner review.

## Deliverables

- Owner-approved authority decision record naming the exact production `qwsg.release-index/1` HTTPS endpoint and hosting/publication ownership.
- Initial production Ed25519 public-key identity/fingerprint, bundled public trust anchor, and strict selection; no private key in Git/runtime hosts.
- Operational private-key generation, custody, backup, recovery, rotation, emergency revocation, retirement, and audit procedure.
- Deterministic canonical index generation/signing/validation/publication integrated with the Forgejo-controlled release lifecycle.
- Authenticated production `qwsg update check` through Task 076/077 and preserved network-free `qwsg update status`.
- Fail-closed transport, schema, authenticity, compatibility, trust, rollback/rotation/revocation, cache, and resource behavior with authenticity/integrity separation.
- Deterministic fixtures/tests, reproducibility and publication evidence, and one explicitly authorized real production acceptance record.
- Updated architecture, release, security, operator, roadmap, engineering history, Task 078 history, snapshot/rollback, Git, and lifecycle evidence.

## Verification

- Run Framework, lifecycle, Builder, diversion/test-task, Git identity/status/divergence, permission/ACL, and snapshot/checksum/bundle validations.
- Run gofmt/fmt-check, `go vet ./...`, focused Task 075/076/077/078 tests, `go test ./...`, repository-wide race tests with isolated caches, ordinary build, deterministic build/release reproducibility, shell/static checks, and `git diff --check`.
- Prove deterministic canonical index bytes/signatures; reject byte/schema/signature/key-ID mutations, unknown keys, unsigned data, duplicate/ambiguous fields, non-canonical encodings, wrong channel/product, downgrade/rollback, invalid epochs/times, incompatibility, invalid artifact metadata, oversized/deep/trailing data, and malicious URL/redirect/origin cases.
- Test initial trust, rotation overlap, designed old/new compatibility, revoked/expired/retired/unknown keys, emergency revocation, recovery from lost publishing key, withdrawal, authenticated `200`/cached `304`, stale cache, failure preservation, and source unavailability. Network content never creates trust.
- Test HTTPS-only policy, bounded time/size/resources, redirects, status/content, atomic/no-clobber publication, interruption/recovery, concurrency, permissions, unsafe links, and privacy-safe diagnostics.
- Prove one Task 076/077 path, network-free status, no download/install, and no Guardian/readiness/notification/updater/telemetry/account/API-key/inbound behavior change.
- Verify packaging contains the exact approved public anchor/fingerprint but no private material. Scan source, staged diff, outputs, fixtures, snapshots, logs, and evidence; test keys are unmistakably non-production.
- Verify Forgejo remains canonical, GitHub read-only, no HTML scraping, and no GitHub trust/publication path.
- Verify production DNS/TLS/hosting origin, exact endpoint bytes/hash/signature/key ID, HTTP/cache behavior, provenance, and rollback target. Perform the separately authorized real acceptance with privacy-safe positive and negative evidence.
- Prove QWSG 1.2.0 tag/commit/release/artifact/evidence unchanged. Review changed/staged paths, modes/ACLs, secrets/privacy, docs, rollback, commit/push/refs, final clean synchronization, canonical idle, and no Task 079.

## Documentation Updates

Update the canonical release-index/release-awareness contract, release policy/lifecycle, native update/rollback architecture, architecture/system map, security/privacy model, installation/operator guidance, release runbook, roadmap, engineering history, and Task 078 history. Add the exact endpoint and authority only after Owner approval. Document hosting responsibility, Forgejo provenance/GitHub mirror boundary, deterministic publication, trust-anchor packaging, fingerprint verification, private-key ceremony/custody/backup/recovery, rotation/revocation/emergency handling, transport/cache/fail-closed semantics, authenticity versus integrity, Community privacy properties, production acceptance, rollback, and limitations. Never document private keys, passphrases, credentials, or private infrastructure identity.

## Completion Criteria

Complete only when the Owner has approved this task and every reserved production-authority decision; the exact official HTTPS endpoint and hosting boundary are operational; the initial Ed25519 public identity is verified/bundled while its private key remains outside Git/runtime hosts; deterministic generation/signing/publication is integrated with canonical Forgejo releases; rotation/revocation/backup/recovery are documented and safely tested; production check authenticates through Task 076/077; every authenticity/transport/schema/compatibility/trust failure closes safely; authenticity remains distinct from artifact integrity; all deterministic/full/security/reproducibility checks and the authorized real acceptance pass; QWSG 1.2.0 is unchanged; no exclusion or Task 079 work exists; rollback and documentation/history/Git evidence are complete; and the repository is clean, synchronized, and idle.

A valid result is `complete`, `complete with disclosed limitations`, or `blocked`. Completion is forbidden while endpoint authority, key custody, rotation/revocation, deterministic publication, authenticated behavior, fail-closed guarantees, real acceptance, rollback, or protected-release preservation is unresolved.

## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder approval contract on 2026-08-30 UTC. The Owner explicitly approved this complete prepared Task 078 definition and authorized implementation to begin. Framework 2.0 Standard Execution Authority applies within this task's recorded scope and exclusions.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another routine Owner gate. Further scope changes, exceptional external actions, and the Owner-reserved production decisions below require explicit Project Owner approval.

The Owner-reserved endpoint, hosting/publication, trust-anchor/key identity, key-ceremony/custody, rotation/revocation, public publication, and real acceptance decisions still require explicit recorded approval at their gates. Task-start approval does not select or activate those values and does not authorize Task 079.
