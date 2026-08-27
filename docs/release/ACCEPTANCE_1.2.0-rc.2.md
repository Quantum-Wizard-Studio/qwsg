# QWSG 1.2.0-rc.2 Guided Installation Acceptance

## Scope and privacy

Task 065 used one Owner-authorized freshly reinstalled disposable Ubuntu
24.04.4 LTS amd64 VPS. The preserved official-1.1.0 update/rollback VPS was not
accessed. Evidence below excludes the private endpoint, SSH material, account
identity, passwords, SMTP values, hostnames, and raw private operational data.
No tag, Release, public upload, production deployment, or announcement occurred.

## Final candidate

- Version: `1.2.0-rc.2` (private, unpublished).
- Source commit: `c260dc18c2004473ec55496d16e66718fd128865`.
- Archive SHA-256:
  `a34be8b18f80d877c0ccfd69dc9d9e9f197fc35fa765cdf1d5c0d72e2cb0a554`.
- Sidecar SHA-256:
  `23218e15a85ab5ee644031bf5ad6469d33a247fc1e271f550de1d52d83e88e11`.
- Two isolated builds were byte-identical. Outer checksum, internal manifest,
  release provenance, embedded binary identity, modes, and packaged guidance
  passed before transfer and again on the target.

Earlier RC.2 builds from commits `0788b1a` and `6f7379d` are superseded private
development candidates. They were never published. The first exposed a living
packaged-documentation mismatch; the second enabled the initial fresh-install
journey and exposed completion-wording defects. Neither is the final candidate.

## Preflight and fresh installation

Read-only preflight proved Ubuntu 24.04, amd64, an ordinary non-root user,
systemd 255, a running user manager, local filesystem semantics, normal
interactive sudo, and absence of QWSG package, configuration, state and unit.
Candidate acquisition used a private mode-0700 staging root and frozen
mode-0400 archive/sidecar.

English, Hungarian and German PTY journeys each showed the QWSG Installation
Wizard identity, localized preflight and plan, 0 percent while preflight was
active, 12 percent only after it completed, exact intended package/operator
data/service destinations, and localized cancellation before mutation. A
capable `TERM=xterm` journey emitted bounded ANSI clear/home control; explicit
line mode and `TERM=dumb` emitted clean sequential output.

The complete fresh journey obtained sudo authentication only through the
Owner's direct VPS terminal. The credential was never passed to QWSG evidence,
automation, source, chat, arguments, environment, or logs. Package installation,
configuration initialization, notify-policy selection, Guardian activation and
readiness completed. The installed binary had exact candidate identity;
configuration and state were current-user private; Guardian was enabled and
active and produced fresh canonical evidence.

## Diagnose, correct and retest

The first complete journey found two ordinary product UX defects: optional
email wording promised immediate configuration although the action intentionally
provides contextual provider-value/protected-credential guidance for later
configuration; and completion reported only requested activation while omitting
the actual overall partial readiness classification. Task 065 corrected message
catalogs and made completion derive readiness from the canonical operational
assessment. Focused/full local validation passed and a new deterministic
candidate was installed with fixed-destination replacement and a private
root-owned backup.

A detached replacement session later appeared to vanish. Read-only diagnosis
proved exact corrected bytes already installed and the backup present: the
still-valid sudo authentication completed the command, after which tmux exited
normally. Classification: `EXPECTED BEHAVIOR`, with no product change.

The corrected Hungarian full journey produced deterministic percentages
`0,12,22,42,57,67,75,87,97,100`; explained SMTP meaning, expected host example,
provider source, fields and protected password command; retained notify policy;
reported Guardian `active`; and reported readiness truthfully as `partial`
because optional notification and lingering remained disabled/unverified.
Required installation, environment, configuration, service and canonical
Guardian evidence were satisfied. English and German catalog completeness and
English fallback passed deterministic catalog tests; their external localized
preflight/plan journeys passed.

## Security and update integration

- Installed binary mode: `0755`.
- Configuration directory/state-root modes: `0700`; configuration: `0600`.
- No credential directory or value was created for the skipped optional SMTP
  capability; guidance never echoed or accepted a secret.
- Replacement backup mode: root-private `0700`; ordinary-user content listing
  was refused as expected.
- `qwsg update status` truthfully reported no native update transaction.
- `qwsg update check` reported installed `1.2.0-rc.2`, official available
  `1.1.0`, relation `older`, and performed no downgrade or mutation.
- Automatic privileged updates remained disabled; notify is preference only.

## Result

`PASS — complete clean-host guided-install journey, correction loop, operational Guardian, truthful progress/readiness, localization, security boundaries and update integration externally validated`
