# QWSG Community Release Authority Decision Record

## Status

Task 078 production authority: **Owner-approved; production public key,
recovery, and first signature verified; signed index checkpoint complete;
first publication unauthorized**.

The Project Owner explicitly approved the endpoint, hosting and provenance
model, initial key ID, dedicated-workstation signing, private custody, separate
encrypted recovery, monotonic rotation, dual-signed overlap, and fail-closed
emergency behavior on 2026-08-30 UTC. The Owner also authorized the production
Ed25519 ceremony and approved the current-stage custody model below. The public
key identity was completed on the dedicated custodian workstation and verified
below. On 2026-08-31 UTC, the Project Owner completed separate recovery
verification on a second physically separate Windows computer. The recovery
gate is `PASS`; first production publication remains a separate explicit Owner
gate.

## Proposed production authority

| Decision | Exact proposed value |
| --- | --- |
| Product and channel | `qwsg`, `stable` |
| Source ID | `community-release-index` |
| Official endpoint | `https://releases.quantumwizard.hu/qwsg/v1/release-index.json` |
| HTTPS origin | `https://releases.quantumwizard.hu` only; default port 443; no alternate origin |
| Redirect policy | No redirect in normal operation; any redirect must remain the exact scheme, authority, port, and path and is bounded by the Task 076 client |
| DNS owner | Quantum Wizard Studio / Project Owner-controlled authoritative DNS for `quantumwizard.hu` |
| Serving boundary | Dedicated static, read-only HTTPS virtual host or object-serving boundary for public release metadata; no application session, account, API key, cookie, inbound client listener, or dynamic per-installation response |
| Canonical publication source | The canonical Forgejo repository and its Owner-approved immutable tag/Release/artifact evidence at `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg`; GitHub is a read-only mirror and never publication authority |
| Deployment method | A separately authorized release operator transfers one already signed, validated index plus public provenance from the isolated release-authority workspace to a staging object; server-side no-clobber/atomic rename publishes the exact bytes; post-publication retrieval must match the approved SHA-256 and authenticate under the bundled anchor |
| Availability assumption | Best-effort public HTTPS. Unavailability or stale evidence never affects Guardian health/readiness and never authorizes fallback to another origin or unsigned data |
| Cache policy | Authenticated `200`; conditional `304` only when the client retains a matching previously authenticated body; bounded validators; no cache response can create or replace trust |
| Initial key ID | Approved `qwsg-community-release-2026-01` |
| Initial Ed25519 public key | Raw 32-byte key, Base64: `r+iDDJJlGRzU/1bv7aSlVl63PcipILaGmdk7130drHQ=` |
| Initial public-key fingerprint | SHA-256 of the exact raw 32 bytes: `0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6` |
| Trust-anchor encoding | Repository/package asset containing schema `qwsg.release-trust/1`, key ID, algorithm `ed25519`, base64 of exactly 32 raw public-key bytes, lowercase SHA-256 fingerprint, monotonic epoch `1`, lifecycle state, and validity metadata; strict decoding and a compiled/package identity assertion must agree |
| Runtime distribution | Public anchor only, bundled in deterministic QWSG release packages and compiled/loaded through one canonical trust provider; private key is forbidden on runtime hosts |

DNS inspection on 2026-08-30 UTC found no address record for
`releases.quantumwizard.hu`. This is expected pre-activation evidence, not
authority to create DNS or choose another host. Existing unrelated hosts are
not inferred as the release authority.

The authorized ceremony was not executed in the repository environment. Safe
inspection established that this machine is an active QWSG runtime host.
Creating private material in its repository, `/tmp`, user home, or other local
filesystem would violate the approved custody boundary. The Owner has
identified a separate dedicated local Windows custodian workstation on which
the ceremony was completed. No key, passphrase, encrypted backup, or secret
intermediate was generated on the QWSG host.

## Initial production key ceremony evidence

- Key ID: `qwsg-community-release-2026-01`
- Tool: `OpenSSH_for_Windows_9.5p1`
- Method: `ed25519-openssh-bcrypt-kdf-rounds-100`
- Environment: `dedicated-local-non-cloud-synchronized-windows`
- Primary custody: `PASS`
- Private-material exposure: `NONE`
- Raw public-key length: independently verified as exactly 32 bytes
- Raw public-key Base64: `r+iDDJJlGRzU/1bv7aSlVl63PcipILaGmdk7130drHQ=`
- Raw public-key SHA-256: independently verified as
  `0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6`
- Separate recovery computer: `PASS`
- Recovery-copy creation and custody: `PASS`
- Recovery private-key decryption: `PASS`
- Recovery public-key match: `PASS`
- Recovery verification: `PASS`
- Private-material exposure: `NONE`
- First production publication: `UNAUTHORIZED`

Only the public key and privacy-safe ceremony result are recorded. The
repository contains no private key, passphrase, encrypted private-key copy,
private-file digest, custody path, or recovery-medium identity.

The encrypted recovery copy is stored on a second physically separate Windows
computer in a dedicated local non-cloud-synchronized directory with restricted
filesystem ACLs. This is recorded only as a custody property; no private path,
private material, passphrase, or host identifier is recorded.

## Operational roles and private-key custody

- The Project Owner approves the endpoint, initial public identity, rotation,
  revocation, first publication, production acceptance, and every release.
- A release operator prepares deterministic unsigned index bytes from canonical
  Forgejo release evidence and performs the authorized publication transaction.
- The Project Owner is the initial production-key custodian. Production
  Ed25519 key generation and signing occur only on the Owner's dedicated local
  Windows workstation in a dedicated directory outside every cloud-synchronized
  folder. The encrypted OpenSSH private key is protected by a strong unique
  passphrase and filesystem ACLs restricted to the custodian and `SYSTEM`.
- The private key never enters Git, Forgejo, GitHub, QWSG runtime hosts, release
  hosting, CI, chat, logs, caches, cloud-sync services, snapshots, published
  artifacts, or task records. The passphrase is controlled separately and is
  never stored beside any private-key copy.
- One encrypted recovery copy is mandatory before first production publication.
  It is stored separately from the primary workstation on a second physically
  separate Windows computer in a dedicated local non-cloud-synchronized
  directory with restricted filesystem ACLs. Offline encrypted removable media
  remains a preferred future defense-in-depth option, not an additional gate
  for the verified current custody model.
- Before first publication, verify that the recovery copy decrypts under
  custodian control and derives the identical public key/fingerprint. Record
  only the public fingerprint, ceremony date, tool/method, and privacy-safe
  PASS/FAIL results. Never record private bytes, passphrases, private filesystem
  locations, or recovery-media identifiers in Git or task evidence.
- Retirement requires a replacement/revocation plan, verified recovery path,
  explicit Owner approval, and independently witnessed destruction of working
  and backup copies. Public keys, signed transition history, and audit evidence
  remain available.

## Signing, generation, and publication transaction

1. Freeze an Owner-approved Forgejo release candidate; never reinterpret the
   GitHub mirror as canonical.
2. Generate typed canonical `qwsg.release-index/1` unsigned bytes from exact
   Forgejo tag, commit, notes, artifact name, size, SHA-256, compatibility, and
   withdrawal inputs. Generation must be reproducible from a data-only input.
3. Validate schema, ordering, bounds, canonicalization, URLs, compatibility,
   artifact metadata, rollback/epoch policy, and byte identity on two isolated
   runs.
4. Transfer only the canonical unsigned bytes and their digest into the offline
   signing environment. Sign with the active Ed25519 key without exposing it.
5. Reassemble the signed document deterministically; independently validate its
   canonical bytes, signature, key ID, fingerprint, and Forgejo provenance.
6. Create a publication checkpoint. Upload to a unique staging name, verify the
   exact digest, then publish through a no-clobber atomic operation. Do not
   overwrite an unrecognized current object.
7. Retrieve through the exact public HTTPS endpoint and verify TLS/origin,
   headers, size, digest, strict parse, signature, trust selection, release
   metadata, and cached behavior. Retain privacy-safe hashes and result tokens.

## Rotation, revocation, rollback, and recovery

- Every trusted key has a monotonic epoch. A lower epoch can never replace a
  higher authenticated epoch.
- Normal rotation uses an overlap release signed by both the current and new
  keys. The new public anchor is shipped in a QWSG package whose metadata is
  authenticated by an already trusted key. Network content alone cannot add a
  trust root.
- Old clients continue to accept the overlap index under the old key; upgraded
  clients accept both during the bounded overlap. Retirement occurs only after
  supported-client compatibility evidence and explicit Owner approval.
- Emergency compromise response stops ordinary publication, preserves the last
  valid authenticated index, removes signing access, records the affected key
  ID/epoch, and uses a separately approved recovery release or bundled
  emergency anchor. An unsigned network revocation or replacement is never
  trusted.
- Loss without suspected compromise invokes the verified separate encrypted
  recovery-copy procedure. Suspected compromise never restores the key for
  ordinary signing. If the primary key is lost before recovery verification,
  production publication remains blocked and a new Owner-approved key identity
  must replace it.
- Production rollback restores only a previously authenticated exact object
  whose identity and authority are proven; it never publishes unsigned fallback
  metadata or rolls trust state back.

## Client fail-closed contract

`qwsg update check` may activate this source only after the decision gate is
complete. It must reject TLS/transport failure, any authority/path change,
unsafe redirect, timeout, oversized body, invalid HTTP/media/cache semantics,
malformed or non-canonical schema, unknown/retired/revoked key, invalid or
missing signature, key-ID/fingerprint mismatch, stale/rollback epoch,
incompatible/withdrawn/ambiguous release, and invalid artifact metadata. A
failed attempt is recorded with a privacy-safe reason token while the last
valid authenticated success remains intact. `qwsg update status` remains
network-free. No check downloads or installs an artifact.

Metadata authenticity proves that an approved release authority signed the
index. Artifact integrity remains a separate SHA-256, size, package-manifest,
and embedded-provenance verification step in the existing updater.

## Production acceptance proposal

After separate approval of publication and acceptance, use one isolated
Community-equivalent installation with verified Task 075 package identity and
no credentials. Permit only outbound HTTPS DNS/TCP needed for the exact endpoint
and canonical Forgejo artifact references. Run one authenticated `qwsg update
check`, then network-isolate the host and run `qwsg update status`. Verify the
expected version relationship, key ID/fingerprint, last-success preservation,
no download/install, no Guardian/readiness mutation, and no account, API key,
telemetry, cookie, or inbound listener. Run bounded negative probes against
local fixtures for wrong origin, altered signature, unknown key, rollback
epoch, stale cache, and invalid artifact metadata; do not attack or mutate the
production service.

## Remaining prerequisites before first production publication

The production authority decisions, key ceremony, public-anchor derivation,
separate recovery-copy custody, recovery decryption, and recovered
public-identity match are complete. The recovery gate is `PASS`.

Before first publication, Task 078 still requires:

1. implementation and verification of the exact bundled public trust anchor
   and production source through the existing Task 076/077 path;
2. deterministic generation, isolated signing-input handling, validation,
   no-clobber publication plumbing, fail-closed behavior, and their required
   tests and reproducibility evidence;
3. a separately verified publication checkpoint for the exact bytes, hashes,
   provenance, rollback target, DNS/TLS origin, and hosting transaction;
4. explicit Project Owner authorization for the exact DNS/TLS/hosting mutation
   and first production publication action; and
5. separate explicit Project Owner approval of the real production acceptance
   target/protocol and action before that acceptance is run.

The trust anchor and exact explicit-check source are implemented. Source
construction failure remains `source_authority_refused`; malformed, unsigned,
wrong-key, and wrong-signature content fails closed while the last valid
authenticated awareness evidence is preserved. Publication remains prohibited
until the exact signed index passes the generated checkpoint and the explicit
first-publication approval gate passes. No DNS, TLS, hosting, production
signing, publication, production acceptance, or Task 079 work is authorized.
