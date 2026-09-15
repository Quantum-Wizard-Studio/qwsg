# Guardian automatic update trigger

Task 085 reuses `guardian.ReleaseCheckService`: startup readiness, one sequential
check, the existing 24-hour cadence/35-second check budget, persisted awareness
attempt time and bounded waits. There is no second scheduler or timer. The
canonical Guardian callback retains Community discovery and optional notification,
then evaluates `AutomaticDecision` using Task 083 capability and effective policy.
Community/default and Pro/manual do not request orchestration. Configuration does
not grant Pro; production still uses Task 083's Community authority adapter.
A trusted Pro-capable test composition demonstrates the automatic path.

An authenticated current observation is a successful no-op. Failed, stale,
unsupported or unauthenticated checks cannot request a handoff. Awareness is only
a decision hint: the handoff reloads canonical policy/capability, installed
identity, signed metadata and watermark before stopping anything. Task 082
`Authorize` establishes current/no-update or supported-newer authority. Task 084
repeats authorization and owns staging, package integrity, source/platform/
provenance checks, privileged helper validation, apply, commit and rollback.

## Safe service handoff

The shipped Guardian has `NoNewPrivileges=true` and a systemd-controlled cgroup.
A forked or detached child would inherit privilege restrictions and be killed
when Guardian stops. Inspection found no existing helper/handoff primitive.
This is the task's explicit architectural-constraint exception: `systemd-run
--user` creates the finite `qwsg-automatic-update.service` transient job. No unit
file is installed and no updater daemon, independent cadence or licensing service
is added. An unavailable manager/systemd-run fails safely. No live service was
changed during implementation or acceptance.

Only the canonical managed Guardian generation/state root and default configuration
path may request the job; custom `guardian --config` instances refuse. The state
root is passed explicitly. The private pending receipt must be persisted before
handoff. The internal `guardian automatic-handoff` entry requires the transient
unit's cgroup and rechecks authority; it is not an operator automatic-update CLI.

The worker verifies the running Guardian InvocationID, requests a synchronous
stop and verifies inactivity. It acquires Guardian's existing instance lock to
exclude a second local Guardian during the transaction. Task 084's inactive check
is unchanged. After the transaction finishes, the worker releases the instance
lock and requests start using an independent recovery context, including stop
failure and rollback cases. The transient unit also has an ExecStopPost start
request as a service-manager fallback. Ordinary transaction failures do not kill
the Guardian scheduler; restart failure is explicit in evidence.

## Repetition, evidence and failures

The fixed transient unit name prevents concurrent handoff jobs. Task 084's
process guard and nonblocking `automatic.lock` still protect actual execution.
A successful update changes the installed identity; the next authenticated check
is current and cannot apply that release again. Repeated observations are never
package authority. A conflicting Task 084 transaction refuses without mutation.

The existing awareness private atomic writer persists two bounded JSON records
under the canonical update state directory:

- `automatic-trigger.json`: latest trigger time, capability/policy decision,
  candidate version and handoff/no-action outcome.
- `automatic-result.json`: last handoff receipt, including whether Task 084 was
  invoked and its unmodified transaction result. That result carries mutation
  knowledge and rollback outcome. A later no-op cannot erase this receipt.

Outcomes distinguish `not_authorized`, `no_update`, `candidate_rejected`,
`handoff_pending`, `handoff_requested`, `handoff_failed`, `orchestration_failed`,
`orchestration_rolled_back`, `rollback_failed` and `success`. Rollback success is
still a failed update. `restart_failed` is separate; it cannot turn a failed
rollback into success. No raw command errors, metadata or secrets are stored.
Failed evidence persistence prevents admission; post-transaction persistence
failure returns a nonzero worker result and an explicit journal diagnostic.

A prior `rollback_failed` receipt inhibits unattended retries (`rollback_blocked`);
an unfinished pending receipt inhibits them as `handoff_incomplete`. Malformed or
unsafe evidence fails closed (`evidence_invalid`). These decisions preserve the
original terminal receipt. After repairing and verifying the installation, an
operator may archive that exact receipt outside the active path before explicitly
restoring automation. No automated repair or clearing of severe evidence exists.
Task 084's lack of persistent crash recovery remains: unexpected coordinator/host
loss cannot be claimed as verified success or rollback, and requires review.

## Validation and deferred scope

Tests cover policy/capability defaults, current no-op, signed supported candidates,
authority/migration refusal, re-entry locking, stop/restart failures, rollback
success/degraded failure, private evidence and scheduler continuation. Separate
process acceptance runs the real release scheduler and Task 084 package transaction
with signed fixtures and isolated fake service/privilege boundaries. The Community
case uses the same candidate without handoff; the Pro case waits for scheduler
quiescence, applies once, persists success and verifies the subsequent no-op.
Existing frozen-client manual/helper security tests remain mandatory. These tests
do not claim a live systemd or production entitlement deployment.

Maintenance/weekday/blackout/timezone rules, fleet/staged/canary rollout, reboot,
remote orchestration, licensing, persistent crash recovery and long health windows
remain deferred. No release is published and Task 086 is not started.
