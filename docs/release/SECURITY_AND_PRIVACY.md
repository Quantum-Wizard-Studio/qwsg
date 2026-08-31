# QWSG 1.2 Security and Privacy

The Guardian runs as the ordinary user. The systemd unit applies `NoNewPrivileges`, private temporary storage, read-only system/home protection with one private state exception, a restrictive umask, and CPU, task, and memory limits. QWSG opens no network listener, executes no remote command, performs no remediation, and stores only secret references—not secret values—in canonical configuration.

Installation privilege is limited to copying checked artifacts under `/usr/local`; it never runs the Guardian as root. State and snapshots can describe the host and must remain private. Checksums detect corruption but do not authenticate an untrusted distribution channel. The QWS Community / Free License governs use and redistribution; it does not replace artifact-authenticity or publisher-trust controls.

Native update discovery and downloads are anonymous and restricted to the
canonical public Forgejo repository. Downloads enter private user-owned staging
and must pass bounded redirect/size, sidecar, archive-layout/type, manifest,
required-file, platform and release-provenance checks before service or
privileged mutation. The privileged helper re-copies and re-verifies the
candidate in private root-owned staging and writes only fixed package
destinations. Configuration, credential and persistent state paths are excluded
from package transactions and rollback payloads.
The updater selects only an exact declared source/target compatibility record,
validates installed configuration before service or package mutation, and the
privileged helper revalidates the migration decision. Unsupported or malformed
identities fail closed. Rollback metadata and before-images are private,
integrity protected and contain only release-owned package destinations; secret
values are neither logged nor composed into lifecycle notifications.

The QWSG 1.3 release-index foundation keeps transport authentication,
publisher authenticity and artifact integrity separate. Its static source sends
only fixed product/media headers plus optional bounded HTTP cache validators;
it sends no installed version, hostname, inventory, identity, email, findings,
or state. It requires approved HTTPS authority/path policy, refuses insecure
TLS/client credentials/cookies/proxy credentials and bounds redirects,
timeouts, headers and body size. Ordinary HTTPS still reveals network-level
source information such as source IP; Community checks are credential-free and
privacy-minimized, not mathematically anonymous. Ed25519 verification requires
an explicitly approved public key and valid signature. Task 078 bundles only
the approved public key and activates only the exact credential-free HTTPS
source for explicit `update check`; it adds no scheduled or automatic network
behavior. Production private keys and passphrases remain exclusively in the
Owner's dedicated custody environments and are never accepted by QWSG runtime,
CI, hosting, Forgejo, GitHub, or repository tooling. The separately built
offline signer prompts locally for the encrypted OpenSSH key path and
passphrase, writes only a detached signature with no-clobber mode, and is used
only on the approved custodian workstation.
Guided activation accepts only a new integrity-checked canonical evidence
identity after service activation; a still-current preserved record cannot
impersonate the newly started Guardian. Waiting is cancellable and bounded by
validated configuration.

Smart Install assessment is read-only. Direct evidence is preferred; external
probes are compiled, absolute, allowlisted, fixed-argument, bounded, and
shell-free. Recommendations are inert registry data and are never executed.
The user-manager probe and guided activation receive only a shared, validated
effective-UID-derived `XDG_RUNTIME_DIR`; arbitrary caller environment,
duplicate entries, HOME and ambient DBus values are rejected. Activation is
limited to fixed absolute systemctl operations and emits only bounded
stage/cause diagnostics. Filesystem
semantics detection uses metadata and `statfs` without creating probe files.
Output uses stable reason tokens and minimal platform facts rather than raw
command output, identifiers, paths, destinations, or secrets.

Per-user configuration is separate from state, current-user-owned, and exact
mode `0700/0600`. QWSG rejects symlinks, special files, wrong owners, and
permissive modes, and publishes changes atomically. General configuration has
only opaque secret references. Optional Community email initiates outbound SMTP
only for due Alert delivery or an explicit test. TLS verification cannot be
disabled. Its separate credential store is `0700/0600`; messages contain
bounded Alert metadata rather than raw host evidence. Delivery failure never
disables local monitoring.

Update awareness persists only bounded source/channel tokens, verified
installed/release identities, Task 076 authenticity evidence, artifact digest
identity, timestamps, validators and safe failure categories. It stores no raw
manifest/body/header, URL, hostname, IP, credential, account, email, inventory,
Guardian finding or notification destination. Its strict digest-checked
single-link mode-`0600` record is atomically published under private mode-`0700`
storage. The local SHA-256 detects out-of-contract mutation but is not publisher
authenticity.
