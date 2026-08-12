# QWSG 1.0.0-rc.1 Engineering Acceptance

## Release decision

`READY FOR QWSG 1.0 RELEASE`

This technical Release Candidate decision does not authorize a public license,
Git tag, signing, upload, publication, or announcement.

## Frozen capability and gate matrix

| Capability | Version 1.0 disposition |
| --- | --- |
| Inventory through Report canonical evaluation | Implemented and verified |
| Configuration, Scheduler, Alert, Notification contracts | Implemented; concrete providers are POST-1.0 |
| Runtime, Runtime Service, Current State, Operator Model | Implemented with bounded durable contracts |
| local Terminal Console and `observe` | Implemented; no network interface |
| systemd-supervised local Guardian | Implemented and verified non-root |
| Dashboard/API, fleet/remote, remediation, AI, commercial systems | POST-1.0 |
| broader distro/architecture support | SHOULD/LATER; not advertised |
| public license, signing, tag and publication | separate Owner authority |

No technical Version 1.0 blocker remains. Physical reboot is an Owner-run
pre-publication procedure because no disposable reboot host was used.

## Artifact evidence

- Identity: `1.0.0-rc.1`, linux-amd64.
- Archive SHA-256: `5b9751048e54eae584b89a4339877813ab0604297a1338ae4516b5657a8716c8`.
- Two assemblies with identical VERSION, commit and epoch were byte-identical.
- Manifest/checksum, safe relative paths, modes, stable metadata and archive
  ownership passed.
- Archive-only staged install, collision refusal, explicit backup replacement,
  old/new rollback, state compatibility, prefix refusal and owned uninstall passed.

## Product and lifecycle evidence

- Full, race, vet, format, Framework, Builder, lifecycle, diversion and Git
  checks passed.
- The shipped unit passed `systemd-analyze verify --user`.
- A unique isolated user service completed 78 recurrence results, published
  current state, and a separate process displayed `Guardian: running`.
- Graceful stop displayed stopped evidence; restart restored running evidence;
  SIGKILL produced one bounded restart with a new process identity.
- Competing Guardian and explicit `observe` were refused. Invalid configuration
  published no false current state. Foreground SIGINT checkpointed gracefully.
- Tests cover Current State 1.0/1.1/1.2, Scheduler 1.0, Guardian Checkpoint 1.0,
  incompatible/corrupt integrity, permissions, symlinks and stale evidence.
- Real PTY evidence showed one initial Overview and one clear/redraw for the
  accepted quit action; noninteractive output stayed a single render.

## Resource and security evidence

After warm-up and more than 50 boundaries, observed peak memory was 8.5 MiB,
tasks/threads 8, open descriptors 7, and restart count 0 during endurance.
Scheduler results were 78 against the 4096 cap; snapshots stayed at retention
10. Current State was 1.6 KiB, checkpoint 1.4 KiB, Scheduler state about 50 KiB.
No monotonic task/FD growth or restart loop was observed.

The ordinary-user unit retains `NoNewPrivileges`, private temporary storage,
read-only system/home protection, one private writable state root, restrictive
umask and 128 MiB/32-task/25%-CPU limits. Audit found no listener, remote
execution, remediation, secret value, raw host identity, arbitrary shell, or
root-runtime requirement. Installation confines privilege to exact artifact
copying and never activates service or lingering.

## Known limitations and Owner gates

Support is Ubuntu 24.04 LTS, systemd 255+, linux-amd64, a working user manager,
and documented filesystem guarantees. Linger remained disabled and unchanged.
The Owner must run the documented disposable-host reboot procedure before a
public boot-persistence claim. SHA-256 is integrity, not authentication. The
temporary proprietary license must be confirmed or replaced before publication.
