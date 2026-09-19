# Automatic update orchestration core

Task 084 provides synchronous `internal/automaticupdate.Run`. Task 085 adds
Guardian triggering and a finite service handoff; see
[AUTOMATIC_UPDATE_TRIGGER.md](AUTOMATIC_UPDATE_TRIGGER.md).
The common core still owns the transaction. `cmd/qwsg` supplies canonical
configuration, capability, installed identity, awareness watermark, verifier and
transport dependencies. Production resolves Community authority; Pro remains a
trusted test composition source until a separately authorized entitlement adapter
exists. Configuration alone cannot grant automatic capability.

## One authenticated engine

The coordinator evaluates Task 083 `updatepolicy.Request` against the immutable
`productcapability.Set`; a display `State` cannot authorize execution. Missing
capability, manual, unknown and conflicting policy refuse before acquisition.
It calls Task 082 `updateauthority.Authorize`, requires the signed understood
`preserve-package-v1` migration, exact installed source, newer applicable target,
platform and awareness watermark, then uses `StageLocal` or `AcquireDigest` and
`Candidate.VerifyStaged`. The verifier and installed classifier are trusted
composition dependencies, never release-supplied code. Historical manual routes
remain available to the manual engine.

The production `automaticHost` invokes `privileged-apply-report`. This is the
same function as manual `privileged-apply`, with a bounded structured receipt
instead of discarded output. The helper independently re-stages under its own
private root, authenticates signatures and migration authority, reclassifies the
actual installed source, and verifies artifact digest, size and provenance before
calling the single `update.Apply`. Pro grants invocation permission only. It
cannot skip a helper check, add executable migration code, or run package scripts.
The `Host` interface is a trusted local OS/privilege adapter, not an untrusted
extension or policy plugin. Test adapters replace destructive host boundaries.

## Lifecycle and evidence

The deterministic successful trace is:

```text
idle -> policy_check -> candidate_check -> eligibility_check -> staging
     -> preflight -> backup -> apply -> post_update_validation -> success
```

Backup and apply are a single existing helper transaction: each allowlisted
managed destination's prior bytes, mode and hash are captured before changing
that destination. The coordinator does not introduce a second backup/installer.
Its `backup` and `apply` stages mark entry to that common transaction, not a claim
that all backups have completed before any write. Package configuration samples
are managed artifacts; private user configuration, credentials and runtime state
are outside the write set.

`Result` includes source/target, capability and policy decisions, terminal state,
ordered stages, original failure stage/category, mutation knowledge, conservative
mutation boundary and rollback attempted/result. Categories and booleans are
bounded audit data: no raw helper output, metadata, paths or secrets are stored.
`rollback_success` is a failed update with successful restoration, never success.
`rollback_failure` is an explicit degraded terminal outcome requiring attention;
callers must not hide it or automatically retry it as an ordinary refusal.

## Mutation boundary and recovery

The common transaction sets `MutationStarted` immediately before entering its
first destination-changing operation (creation or replacement). This deliberately
includes a write attempt whose failure cannot prove absence of side effects.
Private lock/staging/backup scratch is not installed-package mutation. Policy,
authority, package and preflight failures have no installation mutation and no
rollback. A helper rejection before `Apply` returns a known non-mutating receipt.

If the common apply fails after that boundary, its existing in-process restore
now reports whether rollback was attempted and successful, including validation
of every recorded restored file's bytes/mode or absence. The coordinator consumes
that receipt without running a duplicate restore. After apply, the canonical
installed package classifier and configuration validator accept the target;
recording the local rollback metadata is also part of success. Validation or
recording failure invokes the existing privileged rollback and validates the
source identity/configuration. Rollback failure is never swallowed.

A missing, malformed or inconsistent helper receipt means mutation is unknown
(`mutation_known=false`); the coordinator conservatively treats the mutation
boundary as crossed and attempts recovery. Failure to establish rollback is
reported as degraded. The finite local helper is allowed to finish even if the
caller cancels, avoiding a race between an orphaned privileged apply and rollback.
Recovery uses a separate context so caller cancellation does not abandon it.
Sudo is noninteractive; no password prompt or implicit privilege grant is added.
This is synchronous transaction recovery, not persistent crash recovery or
long-duration health monitoring.

## Invocation safety and service boundary

Task 086 shares `internal/updatemutation.Acquire` between manual authenticated
update, automatic `Run`, and manual rollback. It reuses the existing nonblocking
`flock` and retains the canonical private update directory's `automatic.lock`
name so existing automatic clients still coordinate on the same inode. Each
installed instance must use its canonical user/state directory consistently.
Directory/lock ownership and private modes are checked; symlinks are refused.
The file is never unlinked. Failure closes the descriptor just as success does.

The top-level coordinator owns a lease through synchronous helper completion,
post-update validation, rollback/recovery validation and rollback-record changes.
Manual acquisition precedes installed identity/rollback-record reads and Guardian
service changes. Automatic acquisition remains after capability/policy admission,
before candidate authorization/staging; production preflight rechecks source
identity under the lease. A process-wide atomic guard additionally preserves
automatic recursive/concurrent-call rejection across distinct host instances.
Separate file opens also reject same-process manual re-entry without waiting.

Contention returns `transaction_conflict` without package mutation: manual CLI
exit 3 includes guidance to retry after the active transaction finishes; automatic
results retain `MutationStarted=false`. Unsafe/unavailable locks fail closed with
ordinary manual failure or automatic `transaction_lock_unavailable`. There is no
retry loop. Update check, status and release awareness require no mutation lease;
the existing awareness-state lock remains independent.

The privileged apply/rollback/discard functions are root-only internal transaction
steps called synchronously by the owning coordinator, not independent public
orchestrators. They must not acquire recursively or call a top-level coordinator.
The common `update.Apply`/`Rollback` engine likewise executes inside that scope.
Lease ownership is coordination only: release authentication, migration authority,
capability/policy authorization, privileged helper validation and rollback checks
remain mandatory. No caller-controlled “already locked” flag or authorization
shortcut is introduced. Community remains manual by default.

This scope excludes Guardian handoff/restart redesign, cross-user/different-state
installation management and persistent recovery after coordinator/host loss.
The existing Task 085 handoff still stops/resumes around `Run`; this lease covers
the package transaction, not the entire service handoff. Existing older manual
clients do not acquire this lock; deploy consistent client code before relying
on common exclusion. Root administrative actions are outside cooperative locking.

The production adapter requires a verified inactive Guardian service and valid
installed configuration. Active, failed or unknown service state refuses before
mutation. Task 084 does not stop/start Guardian or change service enablement;
it reloads service definitions after package changes and validates the canonical
installed result. Future service coordination must hold the appropriate service
ownership boundary when invoking this core. Tests substitute systemd/sudo only
inside test-linked executables; shipped code has no root/key override flags.

Task 085 connects automatic authorization to the existing Guardian release-check
cadence through a separate decision/handoff layer. Maintenance windows, reboot
handling, delayed health observation and commercial entitlement remain deferred.
No release is published.
