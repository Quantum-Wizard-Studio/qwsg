# Smart Install and Readiness Assessment

## Boundary and bootstrap

Smart Install detects host facts, produces a deterministic plan, and guides the
operator. It never installs packages, changes repositories or services, enables
lingering, alters a firewall, contacts SMTP, or executes a recommendation.
Future provisioning must remain `Detect -> Plan -> Ask -> Execute -> Verify ->
Continue`; Task 047 implements only detection, planning, guidance, and
revalidation.

The release archive already contains the canonical Go engine. Run it before
`/usr/local/bin/qwsg` exists:

```sh
./bin/qwsg install --check
sudo ./install.sh
qwsg setup
qwsg readiness
```

`install.sh` stays a noninteractive artifact copier and does not duplicate
registry truth or run rich assessment as root. `qwsg notification preflight`
remains the focused SMTP view.

## Findings and readiness

Findings use exactly `satisfied`, `missing_required`, `missing_optional`,
`unknown_requires_verification`, and `incompatible`. Domain summaries are the
separate values `ready`, `partial`, `not_ready`, and `unknown`. Mandatory
unknown/missing/incompatible evidence prevents a ready claim. Missing optional
or unknown recommended capability affects only that capability. Guardian core
may be ready while external notification is not ready; overall is then partial.

`internal/assessment` owns Registry and Assessment Model 1.1. Entries contain a
stable ID, purpose token, requirement class, disposition, capability, bounded
probe ID, platform/version constraint, privacy class, and optional structured
remediation with display text, elevation, compatibility guard, and mandatory
revalidation. Model 1.1 additionally owns command-free actionable guidance:
explanation, blocking effect, verification and operator actions, privilege
requirement, manual-verification flag, safety notes, and revalidation action.
The same registry-selected plan drives human and JSON output. CLI code owns
only localization-ready presentation text, not finding or remediation logic.

| Requirement class | Current meaning |
| --- | --- |
| runtime dependency | supported OS/architecture and systemd 255+ |
| install-time dependency | release verification/copy and unit placement |
| optional feature dependency | outbound SMTP and TLS trust |
| environment capability | ordinary user and working user manager |
| configuration requirement | private valid operator configuration |
| recommended non-blocking | filesystem verification and lingering |

The Go release binary needs no Go runtime or external Go module. Shell,
archive extraction, `sha256sum`, `install`, `cp`, `mkdir`, `rm`, `rmdir`, and
`awk` are release-tooling concerns, not automatically Guardian dependencies.

## Platform, probes, and remediation

Supported production remediation is limited to Ubuntu 24.04 LTS, amd64,
systemd 255+, glibc-compatible userspace, an ordinary user, a working user
manager, and documented local filesystem semantics. Other distributions get no
inferred package mapping. Package presence alone does not prove compatibility,
especially on control-panel or managed-stack hosts. No Postfix, Exim, Sendmail,
relay, or certificate-package recommendation is generated.

Commands appear only for exact Registry mappings on the recognized
platform. Current mappings guide QWSG setup and user-unit placement/activation;
package-manager mappings remain absent until separately proven and tested.
Recommendations are displayed as inert data and never executed.

The user-manager probe validates the effective UID's canonical
`/run/user/<uid>` directory before supplying only that derived
`XDG_RUNTIME_DIR` to fixed `systemctl --user is-system-running` arguments. It
does not inherit arbitrary session environment or expose raw stderr. Stable,
transient, unavailable, unsafe, timeout, bounded-output, and unrecognized
states remain separate evidence tokens; ambiguous evidence receives no repair
command.

Filesystem assessment inspects the effective user's QWSG configuration and
state path ancestors without creating files. Ext-family, XFS, and Btrfs local
filesystem types are accepted for the documented Unix ownership, private-mode,
advisory-lock, and atomic-rename contract. Remote, overlay, pseudo, unknown,
inaccessible, symlink, or wrong-owner evidence remains an actionable manual
verification state. Default assessment never performs a behavioral write test.

Direct Go/runtime/filesystem evidence is preferred. External probes map a
compiled ID to one absolute executable and fixed arguments through the bounded
runner: controlled environment, two-second timeout, output bound, process-group
cancellation, no shell, and no caller-controlled executable or arguments.
Default assessment performs no write probe or network connection.

Output excludes credentials, references, destinations, host identifiers, raw
paths, raw output/errors, and inventory payloads. JSON is deterministic apart
from its explicit assessment timestamp and is suitable for future GUI use.

## Existing truth and deferred work

Setup remains responsible for operator values. Addresses, SMTP choices,
credentials, API keys, and VPS purpose cannot be invented. Notification uses
Task 046 preflight; only explicit `notification test` verifies SMTP acceptance.
Unit installed/enabled/active, lingering, and Guardian monitoring are distinct.
Only fresh integrity-checked Current Operator State proves Guardian readiness.

Task 047 persists no VPS-purpose/profile field. A future versioned profile
extension may be designed when role semantics and migration are authorized.
Community receives the complete assessment and guidance. Future Pro may add a
separately authorized executor, not more truthful local detection.
