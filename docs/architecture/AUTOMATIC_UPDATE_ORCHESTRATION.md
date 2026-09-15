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

A process-wide atomic guard rejects recursive/concurrent automatic calls even
with separate host instances. A nonblocking `flock` on the canonical private
update directory's `automatic.lock` rejects other processes using that same
installation identity. The descriptor remains held through commit or rollback;
its file is not unlinked. Unsafe directory ownership/mode and symlink lock files
refuse. All callers must use the canonical installation directory. This guard
covers automatic transactions; it is not distributed locking or a new global
manual-operation scheduling system.

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
