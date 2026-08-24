# Task History 058: Streamlined Engineering Governance, Forgejo Distribution and Practical Release Acceptance

## Task metadata

- Task ID: `058`
- Task slug: `streamlined-engineering-governance-forgejo-distribution-practical-release-acceptance`
- Status: `active — implementation validated; integration pending`
- Date generated: `2026-08-24` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/058_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder installed the approved prompt/history pair with
source identity
`cb74da0f378d868e3019c7abe00112b5a583a43a2d1e1b442e4d5d9167cf018e`
and rendered-prompt identity
`061212f36231634d2e0fcb3d8b2ad87207478874b5ee12ce7f60ea5f5d52d8e0`.
The Owner then authorized the complete bounded implementation, integration,
clean-fast-forward push, validation, and canonical closure cycle.

## Starting state

- Repository root: `/home/qws/web/qwsg.quantumwizard.hu/qwsg`; branch `main`.
- Initial `HEAD`, `origin/main`, and direct Forgejo `main`:
  `0e17a8598821071c4121e1df69b59cd0b62e5506`; ahead/behind `0/0`.
- Source version: `1.1.0-rc.5`.
- Task 058 was the sole active canonical task; index and tracked worktree were
  clean. The only unrelated untracked pathname was the excluded Owner-owned
  blueprint, which was not read, hashed, copied, modified, staged, or packaged.
- Task 057 remained complete with disclosed acceptance limitations. RC.1–RC.5,
  Tasks 055–057, findings, v1.0.0, and LICENSE were preserved.
- No candidate build, transfer, VPS access, credential operation, tag, Forgejo
  Release, upload, publication, deployment, or Task 059 existed or began.

The prompt's pre-install idle wording was recognized as the verified planning
baseline; the canonical Builder-installed Task 058 state was the expected
execution baseline and not a material variance.

## Snapshot

- Private implementation snapshot:
  `/tmp/qwsg-task058-implementation-20260824T.LFDqCZ`
- Mode: `0700`.
- Readable tracked-HEAD archive SHA-256:
  `5c3f8162464824af9f57ee9d1b088f7d088d38de37a600fc1e2e4cf7903ab510`.
- Builder prompt snapshot SHA-256:
  `061212f36231634d2e0fcb3d8b2ad87207478874b5ee12ce7f60ea5f5d52d8e0`.
- Builder history snapshot SHA-256:
  `6b4a0123a805e965c298badd6dc1c2873482fed821a416419c59224198d63f12`.
- `ROLLBACK.md` records literal bounded restoration and forbids broad reset,
  checkout, clean, wildcard deletion, history rewriting, and Owner-content
  access. `ABSENT_TARGETS.txt` records the two new release documents.
- Archive readability, absence claims, snapshot modes, and collision-aware
  rollback were verified before modification.

## Work performed

### Governance audit and correction

Tasks 055–057 showed that their task-specific Gate A–F definitions created most
micro-gating. Framework wording amplified that behavior through unconditional
validation-stop language, an ambiguous Builder handoff, absence of a task-level
execution envelope, and Git policy text that did not distinguish an explicit
task-scoped push authorization from an implicit push.

Reusable Engineering Framework `1.1.0` now requires a first-class `Authority
Envelope` with nine deterministic categories: authorized targets; routine
operations; correction/retest authority; repository integration; lifecycle
completion; permitted external actions; evidence/rollback; Owner-reserved
operations; and mandatory STOP conditions.

Builder approval now means `APPROVED AND AUTHORIZED FOR EXECUTION`. It
authorizes task start and the recorded envelope without a second routine gate.
The canonical task standard, agent rules, lifecycle, prompt workflow, Git
policy, Framework validator, Builder, draft generator, project-local skill, and
focused tests were updated consistently. Historical archived tasks remain
unchanged and valid. Prepared drafts remain unapproved and executable only
after Builder approval.

Recoverable in-scope failures now follow diagnose -> correct -> retest ->
continue. Mandatory stops remain explicit for material scope/architecture,
unplanned destructive or external mutation, unavailable rollback, unresolved
security/privacy uncertainty, credentials, privilege escalation, meaningful
damage risk, tags, Releases, publication, deployment, and envelope-reserved
work.

### Practical acceptance

`docs/release/PRACTICAL_RELEASE_ACCEPTANCE.md` prospectively replaces artificial
per-step gates with one bounded Owner-authorized twelve-step run covering fresh
host, candidate receipt, integrity, Smart Install, installation, guided setup,
Guardian/state proof, one real notification, physical reboot, explicit restart,
uninstall preservation, and same-candidate reinstall/final readiness.

Reliable late or out-of-order evidence may be reconciled when identity,
continuity, and independence remain provable. Reporting order alone does not
destroy clean-host validity. Missing mandatory evidence remains missing and is
never a PASS. Task 057 and all historical protocols/ledgers were not rewritten.

### Forgejo distribution readiness

`docs/release/FORGEJO_DISTRIBUTION.md` defines immutable version tag + associated
Forgejo Release + archive + SHA-256 sidecar and the standard version-specific
`releases/download/TAG/ASSET` contract, with bounded `wget`, `curl -fLO`, and
`sha256sum -c` workflows plus a future Smart Installer contract.

Official Forgejo documentation supports the route model. Read-only anonymous
checks of the configured QWSG repository/Release endpoints returned `404`; no
operational public QWSG asset URL was therefore claimed. A later separately
authorized publication must create the tag/Release/assets and verify expanded
anonymous URLs, redirects, filenames, sizes, hashes, `wget`, `curl`, and
checksum behavior. No authentication or external mutation was attempted.

### Controlled validation corrections

- The first governance run found the diversion test fixture still used the old
  Builder field list. The fixture was updated with the required Authority
  Envelope and passed on retry.
- The first release-plumbing run found an overly long exact grep crossing a
  deliberate Markdown line wrap. The assertion was narrowed to stable semantic
  text and passed on retry.
- A draft-framework fixture initially used an incorrect test-only `sed`
  expression that corrupted headings. The fixture was corrected to append the
  allowed draft marker without altering structure and passed on retry.

These were recoverable in-scope test/orchestration issues, not product or
security defects.

## Verification

- Framework 1.1.0 validation: PASS.
- Framework suite: PASS, 24 assertions.
- Builder suite: PASS, 44 assertions.
- Lifecycle suite: PASS, 29 assertions.
- Diversion suite: PASS, 36 assertions.
- Active-job, lifecycle, and test-task audit: PASS.
- QWSG 1.1.0-rc.5 release plumbing including the twelve-step standard and
  Forgejo distribution contract: PASS.
- Checkout/export build provenance contract, including genuine no-`.git`
  exports: PASS.
- Canonical checkout build: PASS.
- Full `go test ./...`: PASS.
- Repository-wide `go test -race ./...`: PASS.
- `go vet ./...`: PASS.
- Go formatting: PASS.
- Bash and POSIX shell syntax: PASS.
- Git whitespace: PASS.
- Historical protocol assertions and source version: PASS.
- No RC.5 candidate archive or sidecar was constructed.
- Security/privacy/exclusion review: PASS; no secret, credential, private key,
  token, private host identity, snapshot, cache, candidate, Builder input, or
  Owner-owned content is intended for integration.

## Rollback

Rollback remains available from the verified private snapshot. Before push,
restore only exact tracked paths from `tracked-head.tar`, restore the initial
Task 058 lifecycle files from `task-lifecycle/`, and remove a new Task 058 file
only after proving its recorded prior absence and current identity. After a
published clean fast-forward, use a new task-scoped corrective commit rather
than rewriting history. Re-run Framework, Builder, lifecycle, release plumbing,
security, and Git-state checks after any rollback.

Rollback procedure validation: PASS; no rollback was required.

## Completion state

`active — implementation validated; integration and lifecycle closure pending`
