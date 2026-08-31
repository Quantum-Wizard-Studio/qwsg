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

## Frozen first-production signing input

Owner authorization on 2026-08-31 UTC froze the exact canonical input at
`release/production/qwsg-release-index-first-signing-input.json`.

- Size: `738` bytes
- SHA-256: `770759c7935c35c7d9c726837ceb4ce8237f6799aa38b1a449454023cf9c8b68`
- Last byte: `0x7d` (`}`); there is no trailing newline
- Generated-at: `2026-08-31T18:19:25Z`
- Release: unchanged QWSG `1.2.0`, commit
  `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`
- Artifact SHA-256:
  `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`

On Dell 1, place the already verified Task 078 offline signer and this one
non-secret input file in a dedicated local non-cloud-synchronized working
directory. The reviewed Windows amd64 signer build is `5118464` bytes with
SHA-256
`c3f7e9459a8fa23cf6f87daf46046d0cd9bd67c7682efd2a450bf2bf1f7c8b0d`.
Verify that already provisioned executable before use. In PowerShell, verify
the input before signing:

```powershell
(Get-Item -LiteralPath .\qwsg-release-index-first-signing-input.json).Length
(Get-FileHash -Algorithm SHA256 -LiteralPath .\qwsg-release-index-first-signing-input.json).Hash.ToLowerInvariant()
(Get-Item -LiteralPath .\qwsg-release-sign-offline.exe).Length
(Get-FileHash -Algorithm SHA256 -LiteralPath .\qwsg-release-sign-offline.exe).Hash.ToLowerInvariant()
```

The results must be exactly `738`,
`770759c7935c35c7d9c726837ceb4ce8237f6799aa38b1a449454023cf9c8b68`,
`5118464`, and
`c3f7e9459a8fa23cf6f87daf46046d0cd9bd67c7682efd2a450bf2bf1f7c8b0d`.
Then run:

```powershell
.\qwsg-release-sign-offline.exe sign .\qwsg-release-index-first-signing-input.json .\qwsg-release-index-first-signature.base64
```

The signer prompts locally for `OpenSSH private-key file:` and then
`Private-key passphrase:` with no passphrase echo. Enter the Dell 1 path to the
encrypted key whose public identity is `qwsg-community-release-2026-01`, then
the separately controlled passphrase. A successful operation exits zero,
prints no signature or secret to the console, and creates exactly one new file:
`qwsg-release-index-first-signature.base64`. The file is exactly 89 bytes: one
strict Base64 line encoding a 64-byte Ed25519 signature, followed by one LF.

Return only that detached-signature file, its complete 88-character Base64
line, its byte size, SHA-256, and these non-secret results:

```text
key_id=qwsg-community-release-2026-01
signing_input_sha256=770759c7935c35c7d9c726837ceb4ce8237f6799aa38b1a449454023cf9c8b68
signing=PASS
private_material_exposure=NONE
```

Do not return the private-key path, key file, passphrase, console transcript,
or other Dell 1 identity. Transfer the input to Dell 1 and the detached
signature back through Owner-controlled non-cloud transfer storage. Compare
the input hash after transfer and the signature hash before and after return.
Do not modify line endings, reserialize JSON, paste the input through an editor,
or sign a copy with different bytes.

## Local verification

`make release-authority-check` rebuilds the Linux publication tool and Windows
offline signer twice with isolated caches, compares the binaries, generates
canonical signing input twice, compares the bytes, and verifies private output
modes. Go tests cover trust identity, malformed candidates, unsafe Forgejo
provenance, wrong key, wrong signature, encrypted OpenSSH Ed25519 signing, and
existing Task 075–077 regressions using non-production keys.
