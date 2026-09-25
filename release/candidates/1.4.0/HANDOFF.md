# Task 093 frozen Community 1.4.0 candidate and offline handoff

**State: RELEASE CANDIDATE — UNSIGNED — UNPUBLISHED — NOT ACCEPTED.**
No tag, production signature, Forgejo Release, production index replacement or
real-host acceptance is performed or authorized by this record. C9 is MISSING.
Development source, candidate, signed, published and accepted release are distinct
states. Only the first two exist for 1.4.0 at Task 093 closure.

## Frozen candidate identity

| Input | Exact value |
| --- | --- |
| Version | `1.4.0` |
| Source commit | `996fb90d0f53d4ae6088cd884a654d9d0a32fff9` |
| SOURCE_DATE_EPOCH | `1790357147` |
| Embedded build timestamp | `2026-09-25T17:25:47Z` |
| Artifact | `qwsg-1.4.0-linux-amd64.tar.gz` |
| Size | `3680796` bytes |
| SHA-256 | `26cbf8d76fa27c652e11ee49f4570e05b3be6b969309e3cfd486977a27a6224d` |
| Go | `go1.26.5 linux/amd64` |
| Tar / gzip | GNU tar `1.35`, gzip `1.12` |
| Build platform | `linux-amd64`, `CGO_ENABLED=0` |
| Archive checksum sidecar | `qwsg-1.4.0-linux-amd64.tar.gz.sha256` |
| Package manifest copy | `MANIFEST.sha256` |
| Package provenance copy | `RELEASE.json` |

Manifest SHA-256: `1dfb3fe81d8767a83a11d9c1dd53b21d309c3d5d3a0fa2cba9dd2fbb0ee91001`.
RELEASE.json SHA-256: `df7cfb5d7dd404d22d26fac8420ea4cb17c0ab98fd8714fe89a32c2a901d9602`.
MANIFEST.sha256 is a verbatim package copy; verify its entries from the extracted
package root, not this metadata directory. The signed archive hash transitively
binds the complete manifest, provenance, binary and all packaged files.

Two clean exports of the source commit, separate Go caches and umasks 0022/0002
produced identical archive bytes. Archive checksum, every manifest member,
RELEASE.json, embedded version/commit/date and unchanged production anchor passed.
The fixed source commit precedes this metadata/closure commit to avoid circular
source/manifest/signing hashes. No package input changes after that source boundary.
The later test-only staging-directory mode correction is not a package input.
Never build a later checkout while stamping it with this frozen commit.

Payload retention: protected local task storage, logical identifier
`qwsg-task093-candidate`, with two independent outputs under `one/dist` and
`two/dist`. Keep both through Owner acceptance, signing/publication and the
rollback window. Payload is not committed or published; regeneration from this
source and controlled inputs must reproduce the exact hash before use. Loss of a
local temporary directory never authorizes substituting a different artifact.

## Deterministic regeneration

In a new private workspace with the recorded toolchain, export exactly the source
commit using `git archive 996fb90d0f53d4ae6088cd884a654d9d0a32fff9`. Extract into an empty source directory;
from that directory run the existing machinery with new output/cache directories:

```sh
env -u GOFLAGS BUILD_COMMIT=996fb90d0f53d4ae6088cd884a654d9d0a32fff9 SOURCE_DATE_EPOCH=1790357147 DIST_DIR=/ABSOLUTE/NEW/DIST GOCACHE=/ABSOLUTE/NEW/CACHE GOMODCACHE=/ABSOLUTE/VERIFIED/MODULE-CACHE ./scripts/build-release.sh
```

The path arguments are local build locations, not release identity inputs. Require
the frozen archive size/hash, manifest/provenance and version output above. Retain
this closure's unsigned candidate as data; from a trusted source build of the
existing `cmd/qwsg-release-index` run with a new no-clobber output filename:

```sh
./qwsg-release-index generate unsigned-candidate.json regenerated-signing-input.json
cmp regenerated-signing-input.json qwsg-release-index-1.4.0-signing-input.json
```

The URLs in the unsigned candidate are prospective immutable Forgejo destinations,
not availability claims. `active`, `stable`, `published_at` and `generated_at` are
frozen prospective signed-schema fields; they do not publish this candidate.
The existing schema has no draft state and is not extended. Empty signatures,
staging location and this state record distinguish it from production authority.
At later publication review, timestamps must still satisfy client clock/watermark
constraints. Do not silently refresh timestamps; changed bytes require a new
reviewed signing input and signature.

## Dell1 offline signing handoff

- Exact input: `qwsg-release-index-1.4.0-signing-input.json` in this directory.
- Input size: **973 bytes**; compact UTF-8 JSON, final byte `0x7d`, no newline.
- Input SHA-256: **`b44f83cb93f9068ab96a547e35e8544baaa45c6ef424281fd3cc63b95393554a`**.
- Algorithm: Ed25519; key ID `qwsg-community-release-2026-01`.
- Raw public key Base64: `r+iDDJJlGRzU/1bv7aSlVl63PcipILaGmdk7130drHQ=`.
- Raw public-key SHA-256: `0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6`.
- Expected return: **`qwsg-release-index-1.4.0-signature.base64`**, 89 bytes:
  88 strict Base64 characters encoding 64 signature bytes, then one LF.

Next action belongs to the Owner-controlled signing phase. Copy only this exact
non-secret input through Owner-controlled non-cloud transfer storage to Dell1's
existing dedicated local, non-cloud-synchronized custody working directory beside
the already provisioned, verified offline signer. Private key and passphrase stay
there. This task does not perform signing or request private material.

Use the existing approved Dell1 executable: `qwsg-release-sign-offline.exe`,
5118464 bytes, SHA-256
`c3f7e9459a8fa23cf6f87daf46046d0cd9bd67c7682efd2a450bf2bf1f7c8b0d`.
Its unchanged raw Ed25519 signing operation signs these canonical /2 bytes; it
requires no schema-specific client/parser update. Do not substitute a newly built
executable merely because the filename matches.

On Dell1, verify the input size/hash and signer size/hash before the separately
Owner-authorized signing operation:

```powershell
(Get-Item -LiteralPath .\qwsg-release-index-1.4.0-signing-input.json).Length
(Get-FileHash -Algorithm SHA256 -LiteralPath .\qwsg-release-index-1.4.0-signing-input.json).Hash.ToLowerInvariant()
(Get-Item -LiteralPath .\qwsg-release-sign-offline.exe).Length
(Get-FileHash -Algorithm SHA256 -LiteralPath .\qwsg-release-sign-offline.exe).Hash.ToLowerInvariant()
.\qwsg-release-sign-offline.exe sign .\qwsg-release-index-1.4.0-signing-input.json .\qwsg-release-index-1.4.0-signature.base64
```

Enter the encrypted production key path and passphrase only at the local prompts.
Return only the detached signature file, its size/SHA-256, key ID, signing-input
SHA-256, signing PASS/FAIL and private-material-exposure NONE. Compare signature
hash before and after return. Never return key paths, keys, passphrases or console
transcripts. Do not edit JSON, line endings, dates, compatibility or signatures in
an editor. The input and every bound release field must remain byte-identical.

After return, in a separately authorized phase, the trusted release-index tool
assembles and verifies the exact input/signature before any publication:

```sh
./qwsg-release-index assemble qwsg-release-index-1.4.0-signing-input.json qwsg-release-index-1.4.0-signature.base64 qwsg-release-index-1.4.0-signed.json
./qwsg-release-index verify qwsg-release-index-1.4.0-signed.json qwsg-release-index-1.4.0-checkpoint.json
```

No expected production signature or signed-index hash can be asserted before that
return. The generated checkpoint must retain `publication_authorized=false`;
separate Owner publication approval remains necessary.

## Tool identity and prepublication immutable review

The Linux `qwsg-release-index` verifier/generator built from the frozen source with
Go 1.26.5, `go build -trimpath -buildvcs=false`, no linker overrides:
**5864474 bytes**, SHA-256
`d95604f9fb60a047300761fefa7dbbeded5b1ccb3afdb0386b3fb86f1cb3671d`.
Build it in the independently trusted operator workspace; no runtime trust/key
substitution flags are introduced. It authenticates the old production 1.3.1
index and must reject the unsigned 1.4.0 input and any test-key signature.

Before publication verify the exact source commit, future tag target `v1.4.0`,
version, epoch, archive name/size/hash, manifest/provenance, notes and artifact URLs,
stable channel/status, generated/published timestamps, exact 1.3.1 compatibility
and schemas, unchanged public anchor, signing input hash, returned signature,
canonical assembled index hash and checkpoint. Preserve the old published object
and review atomic publication/rollback per the established runbook. No field may
change after signing. A defect returns to a newly reviewed candidate; never retag,
replace or rebuild-as-history the protected 1.3.1 release.

Historical 1.3.1 remains source/tag `45009fe6169bff00842a4c4e9561bf339a5db81e`,
closure `fa570723c4d6b0a1f6758a9c2ea9ddf610716765`, archive SHA-256
`0ab726bcde36182232ff89e3ea33e2d5ab77d8cad2b94ce56c9534a8147d9cd7`.

## Supported 1.3.1 bootstrap (NOT executed by Task 093)

Official 1.3.1 predates Task 082: its strict parser accepts only /1 and its
compiled routes end at 1.3.0 -> 1.3.1. Remote data cannot add the new parser or
capability. Do not claim native `qwsg update` from that binary can reach 1.4.0.
Publishing /2 at the existing endpoint makes legacy 1.3.1 checks refuse until
bootstrap; it must not trigger unsigned fallback or a rewritten historical index.

The supported procedure uses the independently authenticated **candidate archive
binary** as coordinator. Its installed-source classifier probes the installed
1.3.1 independently of the coordinator's own version. Its elevation command uses
its own executable path, so the new helper reauthenticates /2 and the artifact.
It uses the existing preserve-package-v1 transaction; install.sh is not invoked
over an existing installation. This is explicitly authorized bootstrap software
installation, not a capability retroactively added to 1.3.1.

After separate signing and publication approval, prepare the following in a
private operator-controlled transfer directory (no concurrent untrusted writers):

- The exact frozen archive and its checksum sidecar.
- The verified full production-signed index, copied byte-for-byte to
  `qwsg-1.4.0-linux-amd64.tar.gz.release-index.json`.
- The established release-index verifier built from the recorded trusted source
  commit in the release-operator workspace, supplied independently of the archive.
  Its identity must match the Tool identity section. A verifier downloaded from
  an unauthenticated candidate must never establish its own trust.

First run the trusted verifier; the checkpoint output must be new:

```sh
./qwsg-release-index verify qwsg-1.4.0-linux-amd64.tar.gz.release-index.json bootstrap-checkpoint.json
```

Require exit zero, canonical production signature under the pinned key, exact
release/source/platform/compatibility values and index digest from the separately
verified post-signing checkpoint. Inspect the authenticated release's artifact
name, size and SHA-256; they must exactly match this frozen handoff. The unsigned
candidate or signing input is not a signed companion and must fail this command.
Then, with the exact frozen sidecar from this record:

```sh
stat -c %s qwsg-1.4.0-linux-amd64.tar.gz
sha256sum qwsg-1.4.0-linux-amd64.tar.gz
sha256sum -c qwsg-1.4.0-linux-amd64.tar.gz.sha256
```

All three must agree with the authenticated index and the frozen values. Only
after signature AND artifact length/digest validation, inspect the authenticated
archive and extract into a new empty private directory, never over the installation.
All members must be ordinary files/directories below qwsg-1.4.0-linux-amd64/;
no links, absolute paths, traversal or unexpected members. Check MANIFEST.sha256
from within that extracted package; RELEASE.json and `bin/qwsg version` must
match the frozen version, full source commit, build timestamp and linux-amd64.
The embedded public trust file must match the unchanged production anchor.

After bounded real-host authority, pre-state snapshot and recovery review, as the
normal QWSG operator, use the verified extracted binary by its absolute path:

```sh
/ABSOLUTE/PRIVATE/EXTRACT/qwsg-1.4.0-linux-amd64/bin/qwsg update --archive /ABSOLUTE/PRIVATE/TRANSFER/qwsg-1.4.0-linux-amd64.tar.gz --version 1.4.0
```

The absolute locations are selected for the later acceptance host; they are not
runtime test overrides. Do not use sudo on the outer command; its normal bounded
privileged helper performs independent source/index/package checks. Keep the
verified bootstrap binary available for explicit recovery if necessary. If the
installed predecessor differs from official 1.3.1, the signed exact-source grant
must not be broadened; stop for a separately supported plan.

The single signed declaration binds 1.3.1 -> 1.4.0, linux-amd64,
preserve-package-v1, Configuration/Guardian/Scheduler 1.0 and Operator State
1.0–1.2 to the target source commit and artifact name/URL/size/SHA-256. It adds no
script, remote command, migration interpreter, trust key, schema conversion,
wildcard or downgrade authority. Unknown/unsupported source, capability,
constraint, bad signature, bad manifest, stale watermark or mismatched artifact
must refuse. The helper repeats authorization after staging under privilege.

### What C9 still must prove

- Separately authorized publication of the exact signed artifact/index/tag,
  external retrieval, provenance and signature/checksum verification.
- Clean installation on the supported real Linux amd64 environment.
- Official released 1.3.1 artifact/source identity as predecessor; authenticated
  candidate-coordinated bootstrap with real sudo/systemd and operator state.
- Preserved supported configuration, private credentials and state, including
  release-awareness/notification history and interrupted-recovery evidence.
- Correct handling of historical service identities: old nonempty ordinal service
  evidence remains preserved and cannot be compared as if it had new identities;
  two fresh supported observations restore normal comparison (Task 089).
- Explicit rollback of the complete package write set to exact 1.3.1, truthful
  package/recovery outcomes, and original active/inactive Guardian intent with
  observed recovery. Keep the new verified coordinator for recovery procedures;
  a restored old binary does not gain new Task 091 recovery semantics.
- Guardian fresh evidence/health/readiness and applicable notification behavior,
  including failure paths; no blanket SMTP delivery or unattended Pro claim.

Local package tests and frozen-client fixtures do not prove these real-system
conditions. Community remains 8 PASS / 0 PARTIAL / 1 MISSING; Pro 9 / 2 / 3.
