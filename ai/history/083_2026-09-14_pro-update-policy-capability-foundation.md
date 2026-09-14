# Task History 083: Pro Update Policy & Capability Foundation

## Task metadata

- Task ID: `083`
- Task slug: `pro-update-policy-capability-foundation`
- Status: `complete`
- Date: `2026-09-14` UTC
- Agent: Codex
- Human authority: Project Owner explicit Task 083 APPROVE; Task 082 accepted.
- Preferred owner communication language: Hungarian, verified from validated project configuration.
- Related prompt: `ai/archive_prompts/083_2026-09-14_pro-update-policy-capability-foundation.md`
- Dependency: completed Task 082.

## Objective, scope and exclusions

Deliver the canonical product capability and manual/automatic update-policy
foundation, compatible Community behavior, operator visibility, deterministic
Pro test authority, focused security tests and concise documentation. Reuse the
Task 082 common authenticated updater. No automatic execution, scheduler,
Guardian install, reboot, new rollback orchestration, commercial licensing,
network entitlement, deployment, publication or unrelated cleanup is authorized.
Task 084 must not be started or prepared. Published v1.3.1 remains immutable.

## Required reading and verified starting state

Read the project-local qwsg-job skill and applicable repository agent guidance;
all generated prompt Required Reading documents; backup, delivery, architecture,
engineering and security policies; Task 082 history and authenticated migration
contract; existing product edition principles and canonical configuration,
configuration store, installer policy, notifications and update implementation.

Before mutation, framework identity validation and `bin/job --check` passed.
Actual branch was main, HEAD and fetched origin/main both
`38ee760c267aef826c8fde97843853ad206d72ca`, divergence `0/0`, worktree clean,
lifecycle idle at completed Task 082. Canonical HTTPS origin matched project
configuration. Go was `go1.26.5 linux/amd64`; existing module cache was reused.
Ownership/modes were inspected; inherited group-writable source files remain
non-executable. No services, dependencies, credentials or infrastructure changed.
The approved builder transaction created exactly Task 083's prompt/history.

## Snapshot, rollback and risk

Before lifecycle or implementation mutation, a protected local task snapshot
captured the complete tracked baseline as `baseline.tar`, with a per-file
SHA-256 manifest for 565 regular members. Archive SHA-256:

`a2bc5221188a91c27c44654699cdf06a8eb6633cf729556049e3faae95df1fe6`

Payload and manifest are retained outside Git in the task-private snapshot
directory until Owner acceptance. Archive readability, archive hash and every
member hash were verified again after implementation. Only sanitized metadata
is committed. For rollback, verify both hash layers, extract only to a new empty
private directory, compare the exact Task 083 changed paths to baseline and
individually review bounded restoration/removal. Never extract over a live tree,
reset/clean broadly, rewrite published history or touch unrelated files. Preserve
history; re-run relevant tests and lifecycle checks after an authorized restore.

Main risk was granting automatic authority or weakening update validation.
Mitigations: private immutable capability Set, strict declaration validation,
no production Pro input, pure policy eligibility, unchanged authenticated update
engine and negative security tests. Existing schema/extension/storage reuse
limits compatibility and permission risk. No runtime data migration is needed.

## Work performed and decisions

- `internal/productcapability`: one deterministic query boundary for manual and
  automatic product authority. Nil declaration yields Community/manual; supplied
  malformed, unknown, unsupported, duplicate or conflicting declarations return
  zero authority and error. Pro requires explicit grants. No entitlement parser,
  transport, payment/account or production license proof is claimed.
- `internal/updatepolicy`: pure strict manual/automatic eligibility with an
  operator-visible typed result. A policy result never executes or authenticates
  an update. Unknown and unauthorized requests return zero state plus error.
- `internal/configuration`: reuse `installer.update-policy` version 1.0; absent
  means manual, existing notify means manual with notification intent. Strict
  known-extension validation rejects malformed values and ambiguous declarations.
  Configuration schema, old document bytes/identities and private storage remain
  unchanged. Entitlement is not operator configuration.
- CLI composition supplies only Community authority in production; test code
  injects Pro. Configuration resolution and setting enforce policy eligibility.
  Explicit update checks manual capability before existing update orchestration.
  `update status` reports effective policy/capability availability without network
  or installation. Known errors have bounded safe diagnostics. Notification
  intent consumes the same canonical parser.
- Task 082 authority/package/transaction/helper code is unchanged. Regression
  validates Pro automatic eligibility, then refuses unsigned, tampered,
  unsupported migration and source-mismatched candidates through the common
  authenticated gate. Community and Pro have no separate update engines.

## Verification evidence

- Focused capability/policy/configuration/CLI/authority tests: PASS.
- `make fmt-check vet test`: formatting and full vet PASS; all Go packages PASS
  after the environment-specific TLS rerun described below. Full CLI suite
  includes `TestOldBinaryForwardAuthenticatedUpdate`, including real isolated
  common apply/rollback and independent helper reauthentication.
- `go test -race ./internal/productcapability ./internal/updatepolicy
  ./internal/configuration ./internal/configurationstore ./internal/updateauthority
  ./internal/update ./cmd/qwsg`: PASS, including frozen-old-client regression.
- `ai/tests/test-framework-v2.sh`: 15 assertions PASS.
- `ai/tests/test-bounded-diagnostic-runner.sh`: 11 assertions PASS.
- Framework, active lifecycle and diverted-test validators: PASS.
- Snapshot checksum/member verification and Git diff whitespace: PASS.
- `make engineering-test`: checkout/export build provenance PASS; 25 framework,
  36 diversion, 29 lifecycle and 49 task-builder assertions PASS.

Tests cover Community defaults, automatic refusal with no configuration write,
Pro policy acceptance with no metadata fetch or update transaction state,
entitlement loss refusal without configuration mutation, legacy notify,
unknown policy, malformed schema/shape/conflict/capability, zero authority,
immutable grants, and release-security independence. Existing configuration
store permission/atomicity and full integration suites remain intact.

### Diagnosed attempts

- `TEST OR ACCEPTANCE DEFECT`: the initial closure record used narrative headings
  instead of the validator-required Work performed / Completion state headings.
  The closure transaction restored active state, headings were corrected and
  the unchanged lifecycle validators were rerun.

- `PRODUCT/FRAMEWORK DEFECT`: initial policy refusal reached the existing generic
  config diagnostic. Added bounded specific policy/capability diagnostic mapping;
  the same integration test passed without weakening its assertion.
- `ENVIRONMENTAL ISSUE`: sandbox denied `.git/index.lock`; approved scoped
  elevated staging succeeded. Initial export/build could not resolve the new
  packages before that staging and encountered sandbox DNS denial in its private
  cache. Reran the unchanged canonical engineering gate after targeted staging
  with approved sandbox elevation.
- `ENVIRONMENTAL ISSUE`: full Go suite's existing TLS source test could not open
  a local socket in sandbox. Approved `make test` outside sandbox passed all
  packages, preserving already valid cached evidence. No test was skipped,
  disabled or weakened to obtain PASS.

## Documentation and deferred work

Updated README visibility note and architecture index; added the concise
`PRODUCT_CAPABILITIES_AND_UPDATE_POLICY.md` contract. This record and archived
prompt provide task evidence. Production Pro entitlement authentication,
transport/backend and commercial policy remain future work. Task 084 owns
actual automatic orchestration and must retain all Task 082 security/service/
rollback gates. No Task 084 prompt/history or execution is created.

## Completion state and delivery

Exact task paths were reviewed for scope, non-executable modes, whitespace,
privacy and payload exclusion. Published release/tag content remains unchanged. Final commit identity and actual push/fetch relationship
are reported in the Owner handoff rather than embedded self-referentially here.
All required implementation gates passed. Task 083 is complete and its prompt
is archived without a successor; canonical idle validators pass. No source
changed after full validation. Final integration uses targeted staging, commit,
dry-run/fast-forward push and post-push fetch/integrity verification; actual
commit/remote identity is reported in the Owner handoff. Task 084 was NOT started.
No unresolved implementation defect remains; production entitlement and actual
automation are deliberately deferred scope, not shipped capabilities.
