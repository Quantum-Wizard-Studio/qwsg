# Release-index generation and publication runbook

This runbook prepares QWSG's deterministic `qwsg.release-index/1` transaction.
It does not authorize DNS, TLS, hosting, signing, publication, or production
acceptance. Every production execution requires the Task 078 gate recorded by
the Project Owner.

## Fixed authority

- Source: `community-release-index`
- Endpoint: `https://releases.quantumwizard.hu/qwsg/v1/release-index.json`
- Key ID: `qwsg-community-release-2026-01`
- Public key Base64: `r+iDDJJlGRzU/1bv7aSlVl63PcipILaGmdk7130drHQ=`
- Public-key SHA-256: `0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6`

The canonical Forgejo repository and immutable tag/Release/artifact evidence
are the only publication inputs. GitHub is never authority.

## Tool separation

`make release-authority-tools` creates:

- `qwsg-release-index`, for the ordinary isolated release-operator workspace;
- `qwsg-release-sign-offline.exe`, for transfer to and execution only on the
  approved Windows custodian workstation.

The runtime and publication host never receive the offline signer or private
key. The custodian workstation never receives repository credentials. Transfer
only canonical signing input in and a detached Base64 signature out. The
offline signer reads the encrypted OpenSSH private-key path and passphrase
interactively, refuses non-Ed25519 keys and existing output paths, emits a
mode-`0600` detached signature where supported, and zeroes secret byte buffers.
Windows custody ACLs remain an independently verified prerequisite.

## Deterministic transaction

1. Create a data-only unsigned candidate with exactly the Task 076 schema and
   canonical Forgejo provenance. It contains an empty `signatures` array and no
   credentials or private paths.
2. Run `qwsg-release-index generate CANDIDATE SIGNING_INPUT`. The tool strictly
   parses all fields, enforces the canonical Forgejo origin/path, rejects any
   signature, and writes no-clobber canonical signing bytes.
3. Compare two isolated generation runs byte-for-byte and record the SHA-256.
4. Only after explicit signing authorization, transfer `SIGNING_INPUT` to the
   approved custodian workstation and run
   `qwsg-release-sign-offline.exe sign SIGNING_INPUT SIGNATURE`. Never transfer
   the private key or passphrase away from custody.
5. Transfer only `SIGNATURE` back and run
   `qwsg-release-index assemble SIGNING_INPUT SIGNATURE SIGNED_INDEX`. Assembly
   accepts only the fixed production key ID and a strict 64-byte Ed25519
   signature encoding.
6. Run `qwsg-release-index verify SIGNED_INDEX CHECKPOINT`. Verification
   requires canonical bytes, exact Forgejo provenance, the bundled production
   anchor, and a valid signature. The generated
   `qwsg.release-publication-checkpoint/1` records exact endpoint, source, key,
   public fingerprint, index size/SHA-256, and
   `publication_authorized=false`.
7. Review the signed index and checkpoint, protected release identities,
   rollback object, staging name, hosting pre-state, TLS/DNS state, modes/ACLs,
   and secret scan. Obtain explicit Project Owner authorization for the exact
   first-publication action.
8. After authorization only, transfer the already verified signed index to a
   unique staging object, verify its digest, atomically publish without
   clobbering an unknown object, retrieve the exact HTTPS bytes, and authenticate
   them again. Publication failure preserves the prior object and stops.

No step permits rebuilding or changing QWSG 1.2.0, unsigned fallback,
cross-origin redirect, key replacement from network content, production
acceptance, or Task 079.

## Local verification

`make release-authority-check` rebuilds the Linux publication tool and Windows
offline signer twice with isolated caches, compares the binaries, generates
canonical signing input twice, compares the bytes, and verifies private output
modes. Go tests cover trust identity, malformed candidates, unsafe Forgejo
provenance, wrong key, wrong signature, encrypted OpenSSH Ed25519 signing, and
existing Task 075–077 regressions using non-production keys.
