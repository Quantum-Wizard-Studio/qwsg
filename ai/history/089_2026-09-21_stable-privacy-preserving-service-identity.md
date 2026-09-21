# Task History 089: Stable Privacy-Preserving Service Identity

## Task metadata

- Task ID: `089`
- Task slug: `stable-privacy-preserving-service-identity`
- Status: `complete`
- Date: `2026-09-21` UTC
- Agent: Codex
- Human authority: Project Owner explicitly authorized Task 089 implementation, focused integration, documentation, lifecycle completion, commit and push in this session.
- Preferred owner communication language: Hungarian, verified project configuration.
- Related prompt: `ai/archive_prompts/089_2026-09-21_stable-privacy-preserving-service-identity.md`

## Objective, scope and exclusions

Close only frozen Community C2: stable protected running-service identity,
canonical ordering, meaningful change evidence and honest historical comparison.
No service coverage expansion, new health rule, monitoring architecture, key
lifecycle, migration, database, dependency, network lookup, telemetry, production
mutation, release/version change, other gate implementation or successor task.

## Required reading and verified starting state

Read repository guidance and qwsg-job skill; all prompt-required philosophy,
Constitution, agent/task/lifecycle/prompt/Git/execution/diagnostic policies and
project configuration; backup, engineering, delivery and documentation policies;
Task 088 history and frozen product scope; canonical inventory architecture,
collector/privacy, inventory, comparison, Drift, store and pipeline code and
relevant tests and implementation/user documentation. Inspected installed
systemctl manual's full-output contract, without external dependencies.

Main and freshly fetched origin/main both
`975789e60e9a87a224665693964d987b52f02a10`; clean worktree, divergence 0/0,
canonical HTTPS origin, framework valid and lifecycle idle after complete 088.
Ownership/modes and ACL evidence remain in private external-to-Git capture.
Builder transaction installed only the Owner-authorized 089 prompt/history pair.

## Snapshot

Before any mutation, protected local external-to-Git baseline archive captured
600 tracked files with deterministic SHA-256 manifest. Archive readability and
all member hashes passed and were reverified. Archive SHA-256:
`57ef577014362727bdfe79a9b8a930e667ac274522de7f3260a59cc4eaedfb2c`.
Private payload, exact location, environment capture and restore notes remain in
session evidence outside Git. Retain until Owner acceptance.

## Rollback

Verify archive checksum and all manifest members. Extract only into a NEW empty
private directory, never over the live checkout. Compare the exact reviewed
Task 089 changed paths with baseline and review intervening drift before bounded
restoration. The new 089 lifecycle and test files are absent at baseline;
removal requires destructive authority. Preserve evidence; no broad reset/clean
or published-history rewrite. Revalidate focused tests, framework/lifecycle,
permissions and Git status afterward. No destructive restore was executed.

## Risk assessment and plan

Primary risks: retaining positional ambiguity, falsely precise transition
changes, raw-name disclosure, name truncation and overclaiming anonymity.
Reuse the existing namespaced SHA-256/128 primitive, pin its identity contract,
sort protected IDs, request full unit names, keep raw facts redacted, and refuse
exact comparison involving nonempty unsupported identities. No secret or
registry is needed. Existing canonical assembly and Change/Drift engines can
carry protected presence differences without new health logic.

Plan followed: baseline/snapshot, Builder lifecycle, focused inspection,
implementation, early acceptance, stable-source local validation, minimum docs,
C2 decision, idle closure and reviewed integration. No unrelated fixes.

## Work performed

- `internal/collector/collector.go`: item ID `systemd-unit-v1:` plus existing
  `privacyID("systemd-unit", exactUnitToken)`, using the first systemctl column.
  Preserve case, escaped bytes and instance names; descriptions never contribute.
  Sort IDs, check cancellation, reject malformed/duplicate/non-running rows
  without exposing raw evidence. Add only `--full` to existing running-service
  query to prevent display ellipsization. Source remains the system manager.
- `internal/inventory/canonical.go`: shared versioned service identity prefix;
  existing canonical `services:` namespace and resource sorting remain in use.
- `internal/comparison/engine.go`: explicit sentinel compatibility error for
  any nonempty services layer containing ordinal/unknown/malformed IDs. Refuse
  the whole comparison, in either direction and old-to-old, without fabricated
  precise records. Existing pipeline stops and surfaces the safe diagnostic.
  Empty sets have no ambiguous IDs. Store readers and historical bytes remain
  unchanged; two corrected snapshots resume ordinary reporting.
- `internal/collector/service_identity_test.go`: deterministic collector through
  canonical inventory, comparison, Drift, Health, fixture Rule and Report tests;
  stored historical evidence, pipeline refusal and resumed default Community
  report integration. Fixture Rule is test-only, not a new product rule.

## Acceptance evidence

| Case | Result | Evidence |
| --- | --- | --- |
| Repeated alpha/beta/gamma | PASS | Independent collections match canonical layer bytes after fixing observation timestamps; singleton IDs match set membership; comparison unchanged. |
| Reorder gamma/alpha/beta | PASS | Sorted protected IDs and canonical bytes unchanged; no replacement records. |
| Beta disappears | PASS | Beta's protected resource removed; other identities remain; matching service Drift. |
| Beta appears | PASS | Beta's protected resource added with matching service Drift. |
| Same-count beta-to-delta replacement | PASS | Distinct removed beta and added delta resource references despite equal count. |
| Privacy | PASS | Raw fixture names/descriptions absent from inventory, comparison, Drift, Health, Rule, JSON/text Report and pipeline output; error output safe. |
| Historical boundary | PASS | Ordinal Store 1.0 round-trip readable; unchanged file bytes; old/new, new/old and old/old return sentinel with no change records; report pipeline stops explicitly, then resumes with two corrected snapshots. |

Additional regression evidence: fixed known identity vector; case, escaped
instance and long common-prefix names remain distinct; empty-to-present and
present-to-empty changes; canonical byte stability; repeated comparison byte
identity; unknown version, invalid hex/length/case refusal; malformed, duplicate
and non-running rows fail collection; cancellation preserved.

## Verification

- Focused collector/inventory/comparison/Drift/store/report tests: PASS early.
- Final `make fmt-check vet test`: PASS across the complete Go repository,
  including relevant existing Community/app/CLI/pipeline/privacy behavior.
- Focused `go test -race` for collector, inventory, comparison, Drift, pipeline,
  report and inventorystore: PASS. Final collector race repeated after --full.
- Final stable-source `make engineering-test`: PASS, including checkout/export
  byte/provenance build contract, framework identity and 25 framework,
  36 diversion, 29 lifecycle and 49 task-builder assertions.
- `ai/tests/test-framework-v2.sh`: PASS, 15 assertions.
- `ai/tests/test-bounded-diagnostic-runner.sh`: PASS, 11 assertions.
- Active bin/job, lifecycle and independent diverted-test audit: PASS.
- Gate consistency: PASS, exact 14 gates; only C2 status changes; Community
  5 PASS / 3 PARTIAL / 1 MISSING; inherited Pro 6 PASS / 5 PARTIAL / 3 MISSING.
- Local documentation links, whitespace and scope review: PASS.
- Snapshot checksum/readability and every one of 600 member hashes: PASS.
- No new concurrency primitive, dependency, network service, private data or
  production access. No security/update/privilege/Community-Pro code changed.

## Diagnosed validation attempts

TEST OR ACCEPTANCE DEFECT: first idle rotation was rolled back transactionally
because history headings used descriptive synonyms rather than validator-required
`Work performed` and `Verification`. Corrected headings without changing evidence;
revalidated closure before integration.

EXPECTED BEHAVIOR / invalidated acceptance input: the initial full-validation
run passed Go checks, but its build-contract export preceded the final --full
source correction. Export-versus-live-checkout byte comparison therefore failed.
The old evidence was not treated as final. Final Go/race checks passed; the full
engineering/build contract was restarted from stable identical source. No build
policy, comparison test or security gate was weakened to obtain success.

## Documentation updates

Minimum identity/privacy and historical comparison architecture sections; paired
English/Hungarian snapshot comparison guidance for the compatibility error and
two-new-snapshot recovery; canonical frozen C2 register and totals; concise
engineering milestone; Task 089 prompt/history only. Historical 088 evidence and
all other release gate statuses remain unchanged.

## Privacy and security semantics

Identifiers are lowercase 128-bit truncated SHA-256 derivatives of
`qwsg:systemd-unit:v1:` plus exact unit token. They are deterministic unkeyed
pseudonyms, not anonymous identifiers or encryption. Candidate-name guessing
and cross-host correlation are possible, as with the existing primitive.
Raw service facts remain redacted; host/network/mount/device redaction is
unchanged. No secret/key management, new state registry or historical migration.
Presence means membership in the running set, not installed unit-file existence
or executable/workload continuity. No broader health claim is made.

## Deferred findings and limitations

No newly discovered unrelated issue requires a separate finding. Existing
C3/C7/C8/C9 and Pro gates remain exactly as frozen. The explicit historical
comparison refusal is an intentional compatibility limit, documented above.
Full release construction and real production acceptance are outside Task 089;
local build-contract validation is not a new release or C9 acceptance claim.

## Completion state

All implementation, seven acceptance cases and mandatory local gates passed.
C2 is PASS; every other gate retains its baseline status. Both records are
complete and archived without a successor; idle closure is validated before
integration. No unresolved Task 089 blocker.

Integration uses only reviewed explicit Task 089 paths, staged diff/mode/privacy
review, commit, canonical dry-run and fast-forward push, fetch and clean main
verification with HEAD == origin/main and divergence 0/0. Actual final commit
identity and push outcome are supplied in the Owner handoff to avoid a
self-referential commit hash. No next task, tag, release or production action.
