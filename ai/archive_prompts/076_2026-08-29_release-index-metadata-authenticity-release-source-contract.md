# Current Engineering Task 076: Release Index, Metadata Authenticity and Release Source Contract

## Task Metadata

- Task ID: `076`
- Task slug: `release-index-metadata-authenticity-release-source-contract`
- Status: `complete`
- Date opened: `2026-08-29` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

Release Index, Metadata Authenticity and Release Source Contract


## Objective

Implement the next approved Task 074 architecture stage for QWSG 1.3: a strict, bounded `qwsg.release-index/1` parser and validator; an explicit authoritative-metadata verification boundary; and a source-neutral, read-only release-source contract with bounded HTTPS transport. Integrate Task 075's canonical installed-package classifier only where installed identity is required for deterministic release evaluation. Deliver reusable, tested foundations without adding Guardian scheduling, persistent Update Awareness State, notification behavior, installation automation, telemetry, registration, publication, or external-host mutation.


## Scope

Inspect the existing `internal/update` discovery, strict version comparison, HTTP client, Forgejo API handling, artifact acquisition/verification, migration registry, Task 074 architecture, and Task 075 installation classifier before changing production code.

Define and implement one canonical release-index domain contract consistent with `qwsg.release-index/1`. Strictly validate schema/product identity, generated and publication timestamps, stable channel vocabulary, release status, strict versions, source commit, release-notes reference, minimum-source advisory, migration identifiers, platform-specific artifact identity, HTTPS URL, artifact name, positive bounded size, lowercase SHA-256, uniqueness, count/size limits, and signature records. Reject unknown fields, duplicate JSON members, duplicate semantic identities, ambiguity, unsupported schema/product/platform/algorithm, unsafe URLs, malformed values, trailing data, excessive nesting or resource use, and conflicting active/withdrawn entries.

Define a deterministic signature-input contract over a typed canonical serialization excluding the signatures member. Implement metadata-authenticity verification using Ed25519, explicit key IDs, an immutable caller-supplied or compiled trust-anchor set, and a policy requiring at least one valid approved signature. Keep transport authentication, metadata authenticity, artifact SHA-256 integrity, internal package verification, and release provenance as separate evidence layers. Production private-key generation, custody, signing, rotation publication, revocation publication, and release-index publication are not part of this task. Do not claim cryptographic authenticity when no approved trust anchor and valid signature exist.

Define a source-neutral read-only `ReleaseSource` abstraction equivalent to `Fetch(channel, validators, deadline) -> not-modified or bounded manifest bytes plus safe source evidence`. Implement the minimum static HTTPS source adapter needed to exercise the contract and place existing Forgejo-specific discovery behind or adjacent to the abstraction only where this can be done without redesigning the updater. No HTML scraping is permitted. Source fallback must never silently widen the approved authority set.

Implement bounded transport behavior: HTTPS only; explicit approved host/path policy supplied by configuration or construction; normal certificate validation; constrained redirects that cannot change to an unapproved authority; overall request deadline; connection/TLS/header/idle bounds where supported; bounded response body and header-derived validators; bounded status handling; content-type policy; cancellation; no credential requirement; and safe failure categories for timeout, transport, HTTP status, size, media type, malformed index, unauthenticated metadata, and unsupported contract. Do not log response bodies, tokens, credentials, host inventory, or sensitive server state.

Implement deterministic release selection/evaluation sufficient for this stage: consume only Task 075 verified installed identity; select the explicitly requested channel; exclude withdrawn, wrong-platform and ineligible prerelease entries; choose the greatest eligible strict version; compare it with the installed version; treat manifest minimum-source and migration IDs as advisory; and consult the existing local migration registry for executable compatibility. Return a read-only structured result. Do not persist it, notify it, schedule it, download an artifact, or cross the privileged installation boundary.

Add deterministic unit and integration tests using local in-process HTTP/TLS fixtures and test-only Ed25519 keys. Cover valid stable metadata, valid signature, wrong/unknown key, altered payload, malformed signature, unsigned data, strict JSON failures, duplicate keys and identities, schema/product/channel/platform/status/version/timestamp/commit/URL/hash/size/count bounds, withdrawal, prerelease eligibility, HTTPS/authority/redirect/content-type/status/body/timeout/cancellation behavior, no fallback trust expansion, Task 075 verified identity consumption, legacy/unknown/inconsistent installed identity refusal, supported/unsupported migration advisory interaction, and existing Forgejo/update compatibility where affected.

Update canonical architecture, system-map, roadmap, security/privacy, release/operator guidance, engineering history, and Task 076 history as required by actual implementation. Use one implementation rather than parallel parsing, trust, or release-source subsystems.


## Out of Scope

Do not implement Guardian periodic update scheduling; persistent `qwsg.update-awareness/1` state; `last_attempt`/`last_success` state; update notification events, transitions, retries, or deduplication; unattended or automatic installation; a second updater; privileged mutation; release artifact download changes beyond interfaces strictly required by the contract; a central release service; a public release-index publisher; production private-key generation, storage, custody, or signing; production key rotation/revocation operations; Pro API keys, registration, fleet inventory, console, QUWIP, Telegram, licensing, or telemetry; HTML scraping; external-host mutation; VPS/runtime changes; Forgejo Release/tag/asset publication or alteration; QWSG 1.2.0 release/tag/artifact changes; unrelated refactoring; or Task 077 work.


## Authority Envelope

**Task targets and boundaries:** Inspect and modify only the existing update/discovery/version/migration integration necessary to add the canonical strict release-index model, authenticity verifier, source-neutral read-only source abstraction, bounded static HTTPS adapter, Task 075 installed-identity consumption, deterministic evaluation, tests/fixtures, and synchronized documentation. Preserve the existing updater, acquisition, package verification, migration, transaction and rollback authority. No awareness persistence, Guardian/notification integration, installation automation, publication, Pro work, or unrelated change.

**Permitted external actions:** Read-only canonical Git synchronization and, only if needed to verify an already documented public contract, bounded anonymous read-only access to the official QWSG Forgejo repository. Prefer local fixtures; no external host mutation, credential use, account registration, release publication, production signing operation, VPS access, service change, or infrastructure action.

**Owner-reserved decisions:** Exact production release-index endpoint activation; production Ed25519 public trust-anchor values and key IDs; private-key custody or signing; key introduction/rotation/revocation/recovery policy; release-index publication; fallback-source policy; enabling a new default network behavior or CLI behavior; Guardian scheduling; awareness persistence; notifications; automatic installation; Pro/central/telemetry capabilities; external acceptance; release/tag/artifact operations; and Task 077 scope.

**Task-specific STOP conditions:** Stop on baseline or protected-release identity mismatch; unavailable rollback; ambiguity that would require weakening strict parsing or authenticity; need to invent or embed an unapproved production trust anchor or endpoint; need to claim authenticity from SHA-256 or HTTPS alone contrary to the evidence model; credential/privacy exposure; authority fallback expansion; external mutation; updater redesign; or any required persistence, scheduling, notification, publication, privileged, or Pro expansion.


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

Expected canonical baseline after accepted Task 075: repository `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, branch `main`, `HEAD == origin/main == 687285a4d525557c1f9b1ffb2444c7918e2f7cb4`, divergence `0/0`, clean index/work tree, and canonical idle with Task 075 as the latest complete archived task. Independently verify repository root, remote, branch, HEAD, origin relationship, lifecycle, framework, ownership/modes/ACLs, `VERSION`, Task 074/075 records, annotated `v1.2.0` tag object and peel, frozen artifact size/checksum, and relevant source state before modification. Stop on an unauthorized material mismatch; diagnose recoverable lifecycle-only differences within the approved framework.


## Snapshot Requirements

Before any production modification, retain the verified Builder installation snapshot and create a unique external mode-0700 Task 076 execution snapshot under `/tmp`. Record exact HEAD/branch/origin/divergence/index/work-tree/lifecycle/framework identities; protected QWSG 1.2.0 tag, release commit, artifact size/checksum and relevant release evidence; exact before-images or absence records for every intended target; ownership, modes and ACLs; tool identity; checksums; a complete readable Git bundle; and literal bounded rollback instructions. Evidence files must be mode 0600, exclude credentials/runtime/private host data/caches, pass checksum/bundle/readability validation before production change, and remain available through Owner review. Create a separately verified closure snapshot before prompt archival.


## Risk Assessment

Critical authenticity risk: accepting attacker-controlled or ambiguous metadata could direct acquisition to a malicious artifact. Mitigate with strict bounded parsing, duplicate rejection, authority allowlists, deterministic signed representation, approved Ed25519 trust anchors, explicit signature policy, and fail-closed evidence states.

High canonicalization risk: inconsistent serialization could verify a different semantic document. Mitigate with one typed canonical encoder, complete normative test vectors, no maps in signed content where ambiguity is possible, exclusion of only the signatures field, and altered/duplicate/trailing-data tests.

High trust-bootstrap/key risk: inventing a production key or conflating TLS/SHA-256 with publisher signatures creates false authenticity. Mitigate by keeping production anchors Owner-reserved, supporting injected immutable trust sets, using test-only keys in fixtures, and reporting transport-authenticated versus signature-authenticated evidence truthfully.

High network-security/privacy risk: redirects, proxies, oversized bodies, unbounded waits, credentials, or detailed errors could leak data or consume resources. Mitigate with HTTPS/authority constraints, bounded deadlines/body/counts/validators, cancellation, safe error categories, anonymous minimal requests, and no host-state fields.

Medium compatibility risk: replacing existing Forgejo discovery or version/migration behavior could break explicit updates. Mitigate by inventorying and reusing existing `internal/update` behavior, isolating adapters, retaining strict version/migration authorities, and running existing updater/acquisition/rollback tests.

Medium installed-identity risk: falling back to binary output would undo Task 075. Mitigate by consuming only `internal/installation.Classify` results and testing every unsafe state.

Medium scope risk: release awareness persistence, Guardian scheduling and notification are adjacent but explicitly deferred. Mitigate with pure/read-only outputs and source-level exclusion review.

Low rollback risk: changes are repository-local and reversible through exact snapshots, targeted Git integration and preserved protected-release identities.


## Planned Work

1. Verify canonical repository, lifecycle, framework, protected QWSG 1.2.0 identities, exact origin relationship, ownership/modes/ACLs and clean baseline.
2. Create and validate the mandatory Builder/execution rollback evidence before production modification.
3. Inventory existing `internal/update` discovery/API/HTTP/version/artifact/migration paths and Task 075 classifier consumers; identify reusable code and any divergent parsing or trust assumptions.
4. Write the normative `qwsg.release-index/1` Go domain model, bounds, invariants, canonical signing representation, safe evidence and error categories before transport integration.
5. Implement strict duplicate-aware bounded JSON parsing and semantic validation with one canonical path.
6. Implement Ed25519 signature verification against explicit immutable approved trust anchors, with test-only keys and deterministic vectors. Do not create or publish production private keys.
7. Implement the source-neutral `ReleaseSource` contract and bounded static HTTPS adapter with authority, redirect, media-type, size, timeout, cancellation, validator and status controls.
8. Adapt or preserve existing Forgejo release discovery without HTML scraping or silent trust/fallback widening; avoid a second updater.
9. Implement pure deterministic evaluation using Task 075 verified installed identity, strict version/channel/platform/status selection, manifest advisories and the authoritative local migration registry.
10. Add boundary, adversarial, transport, authenticity, compatibility and regression tests, including legacy/unknown/inconsistent installed identity refusal.
11. Verify no awareness persistence, Guardian scheduling, notification, automatic installation, Pro/telemetry, external mutation or publication behavior was added.
12. Synchronize architecture, security/privacy, operator/release, roadmap, engineering and task documentation.
13. Run focused and complete canonical validation, reproducibility, scope/privacy/mode/ACL/Git checks and protected-release identity proofs; diagnose and correct in-scope failures.
14. Finalize history, create closure evidence, commit/push by clean fast-forward, archive Task 076 to canonical idle, report exact evidence, and stop for Project Owner review without starting Task 077.


## Rollback Plan

Verify the applicable snapshot checksums, complete bundle, repository/protected-release identity, exact target list, current Task 076 file identities, and absence of later Owner edits. Restore only recorded pre-existing Task 076 targets with their modes/ACLs and remove only Task 076-created paths whose prior absence and current Task 076 identity are proven. Never use broad reset, checkout, restore, clean, wildcard deletion, history rewrite, tag mutation, artifact deletion, or external-system changes. Re-run focused/full/framework/lifecycle/Git/reproducibility/protected-release checks after rollback. If exact restoration or ownership of a changed target cannot be proven, stop for Owner direction.


## Deliverables

- One canonical strict bounded `qwsg.release-index/1` model, parser and semantic validator.
- One deterministic canonical signature-input contract and Ed25519 metadata-authenticity verifier with explicit key-ID/trust-anchor policy and deterministic test vectors.
- One source-neutral read-only `ReleaseSource` abstraction and bounded static HTTPS adapter, with no HTML scraping or silent authority expansion.
- Safe structured source/authenticity/failure evidence that keeps transport authentication, metadata authenticity and artifact integrity separate.
- Pure deterministic release evaluation consuming Task 075 verified installed identity and the existing strict version/migration authorities.
- Compatibility preservation for the existing Forgejo discovery/updater/acquisition/migration/transaction/rollback path.
- Deterministic unit/integration/adversarial tests and local TLS fixtures covering parser, signature, source, transport, selection and installed-identity boundaries.
- Synchronized canonical architecture, security/privacy, operator/release, roadmap, engineering and Task 076 history documentation.
- Exact completion report and valid rollback/lifecycle evidence.


## Verification

Run every validation required by the canonical QWSG engineering framework. At minimum: active/idle lifecycle validation; framework tests; Builder tests; diversion/lifecycle assertions; gofmt verification; `go vet ./...`; `go test ./...`; `go test -race ./...`; deterministic build/reproducibility checks; `git diff --check`; targeted staged scope/privacy/secret/mode/ACL review; and final repository cleanliness/origin convergence.

Explicitly verify: valid signed stable index; deterministic canonical signature bytes/test vectors; altered payload rejection; unsigned, malformed, wrong-key, unknown-key and duplicate-signature handling; duplicate JSON key and semantic identity rejection; all schema/product/channel/status/version/time/commit/platform/artifact/URL/hash/size/count/body/nesting bounds; withdrawn and prerelease behavior; HTTPS-only approved authority; redirect confinement; timeout/cancellation; bounded validators/content type/status/body; safe failure categories; no response-body or sensitive-data leakage; no silent source fallback expansion; anonymous request contains no hostname, inventory, MAC, identity, email, findings or server state; verified Task 075 installed identity acceptance; no-install/legacy/unknown/inconsistent refusal; manifest compatibility remains advisory; local migration registry remains authoritative; existing update/acquisition/package/migration/transaction/rollback tests remain valid.

Prove by source/diff/test inspection that no awareness state, Guardian scheduling, notification/deduplication, automatic installation, registration, telemetry, Pro behavior, publisher/private-key operation, external mutation or Task 077 work was added. Reconfirm exact QWSG 1.2.0 `VERSION`, annotated tag object/peel, release commit, artifact size/SHA-256 and published release evidence. Do not remove or rebuild the frozen artifact merely to satisfy a clean-output construction precondition; classify that expected gate truthfully and run applicable constituents separately.


## Documentation Updates

Create or update the canonical release-index/release-source/authenticity architecture contract, including schema tables, normative bounds, signature-input rules, trust evidence, source abstraction, transport/failure/privacy model, selection/evaluation order, Task 075 relationship, migration advisory boundary, Forgejo compatibility and deferred awareness/Guardian/notification boundaries. Update `ai/core/04_ARCHITECTURE.md`, `ai/core/05_SYSTEM_MAP.md`, `ai/core/07_ENGINEERING_HISTORY.md`, `ai/core/13_ROADMAP.md`, `docs/architecture/UPDATE_DISCOVERY_AND_RELEASE_AWARENESS.md`, relevant native update/release/security/operator documentation, and the detailed Task 076 history as actual changes require. Preserve English engineering artifacts and localization readiness for any user-facing identifiers; do not claim deployment, production signing, publication or external acceptance not performed.


## Completion Criteria

Task 076 is complete only when one strict bounded release-index parser/validator, one deterministic authenticity verifier, one source-neutral release-source contract, bounded HTTPS transport, and pure installed-aware release evaluation exist and share existing QWSG authorities rather than duplicating them; valid signed metadata succeeds and malformed/ambiguous/unauthenticated/unapproved metadata fails closed; Task 075 unsafe installed states cannot become trusted version inputs; local migration planning remains authoritative; all mandated focused/full/race/framework/reproducibility/security/privacy/rollback/lifecycle checks pass; documentation and history are synchronized; Git closes clean and synchronized in canonical idle; QWSG 1.2.0 tag/release/artifact evidence remains exact; no excluded awareness persistence, Guardian scheduling, notification, automatic installation, publication, Pro/telemetry, external-host or Task 077 work occurred; and the full Owner report identifies exact commits/files/contracts/trust decisions/tests/snapshots/limitations and recommends—but does not start—the next separately reviewed task.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-29 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
