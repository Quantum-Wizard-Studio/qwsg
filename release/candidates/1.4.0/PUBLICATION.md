# Task 094 controlled signed 1.4.0 publication

## Authority and state

Owner session authority permits only the exact frozen Task 093 source, artifact,
canonical signing input and returned offline signature after all independent
prepublication checks pass. No changed signed fields, artifact rebuild, different
version, trust weakening or real-system acceptance is authorized.

Prepublication verification: PASS on 2026-09-28 UTC. Forgejo tag/Release/assets and their independent
external verification pass. Canonical index publication and its external proof
remain pending Owner root execution. Overall completion is not claimed.
The original cryptographic checkpoint remains byte-identical with
`publication_authorized=false`. The existing runbook separates cryptographic
verification from Owner authority and publication evidence; it does not define a
command that flips that field. Task 094 records these separately.

## Immutable evidence

| Object | Bytes | SHA-256 |
| --- | ---: | --- |
| Artifact | 3680796 | `26cbf8d76fa27c652e11ee49f4570e05b3be6b969309e3cfd486977a27a6224d` |
| Signing input | 973 | `b44f83cb93f9068ab96a547e35e8544baaa45c6ef424281fd3cc63b95393554a` |
| Returned signature | 89 | `128dd2c15eabeb644135e184785440c1c67565508335c60037da5590d811601f` |
| Trusted verifier | 5864474 | `d95604f9fb60a047300761fefa7dbbeded5b1ccb3afdb0386b3fb86f1cb3671d` |
| Signed index | 1153 | `fbb39b52e240ff3691090a48a1e7915e516da6fb2f6a1faa9ed5fc069b633c41` |
| Verification checkpoint | 434 | `2477382647ca0f6d5bd13863171a443b0210771f0cea490b54976059525ecf83` |

Source `996fb90d0f53d4ae6088cd884a654d9d0a32fff9`, VERSION `1.4.0`, epoch
`1790357147`, build `2026-09-25T17:25:47Z` independently match. Both retained
reproducible outputs have the exact artifact identity. Archive path/type safety,
all manifest entries, checksum sidecar, embedded version/source/build and exact
MANIFEST.sha256/RELEASE.json copies pass; their hashes are recorded in HANDOFF.md.
Frozen release notes remain verbatim, including their historical Task 093
preparation-state statements. The Forgejo body explicitly distinguishes Task 094
publication from those historical notes and from outstanding C9 acceptance.

Both retained trusted verifier copies match the frozen source/tool identity.
Independent assembly and production verification reproduce exact signed bytes
and checkpoint. Pinned production key ID and raw public-key fingerprint match
HANDOFF.md and the frozen trust contract. No untrusted build establishes authority.

## Prepublication checklist

- Baseline main/fetched origin/main, 0/0 divergence, Task 093 idle closure and
  expected signing-return files verified; protected snapshot and Builder lifecycle.
- Frozen source, VERSION, epoch, embedded identity, artifact name/size/hash,
  sidecar, manifest, provenance, reproducibility evidence and notes verified.
- Exact input/signature/verifier/index sizes and hashes; production signature,
  key ID/fingerprint and endpoint independently verified.
- Stable/active 1.4.0, exact source, linux-amd64 artifact URL/name/size/hash,
  minimum source 1.3.1 and sole exact 1.3.1 -> 1.4.0 declaration verified:
  preserve-package-v1, Configuration/Guardian/Scheduler 1.0, Operator State 1.0–1.2.
- Generated/published time remains 2026-09-25T17:25:47Z. Actual-clock authorization
  passes and exceeds prior authenticated 2026-09-06T22:58:30Z watermark. Code
  enforces 15-minute future skew and monotonic watermark, not maximum index age.
  A newer watermark still rejects this candidate; no timestamp was refreshed.
- Unsigned, malformed signature, test-key signature using production key ID,
  changed payload/artifact/compatibility and wrong anchor reject without a
  checkpoint. Unknown migration capability independently fails closed.
- Focused releasediscovery/releasepublication/updateauthority/updateawareness
  tests pass. Isolated actual signed-candidate authorization passes. A synthetic
  installed classifier is local test evidence, not installed-system acceptance.
- Historical public 1.3.1 Release 5, artifact, sidecar, signed index and remote
  tag target verified. Artifact digest 0ab726bcde36182232ff89e3ea33e2d5ab77d8cad2b94ce56c9534a8147d9cd7;
  index digest 2eefd328c5ded8807101bdd69d5fcee0ed4f4f2e49b10cdd02b55da6e7aaf879;
  source/tag 45009fe6169bff00842a4c4e9561bf339a5db81e. v1.4.0 absent before publication.
- Forgejo destination/access and all exact upload bytes verified. Index serving
  object equals public historical bytes; ordinary 0644 object / 0755 directory,
  no extended ACL; existing endpoint TLS/media type/no-cache intact.

## Publication and rollback plan

Use the established Task 081 Forgejo workflow: immutable annotated v1.4.0 tag
at the exact source, draft Release, exact assets, final Release, anonymous curl
and wget retrieval. Publish archive, sidecar, byte-identical signed archive
companion, manifest, provenance and verbatim release notes. No historical object
is deleted or overwritten.

Use the established Task 081 locked same-directory atomic index replacement,
adapted only to exact Task 094 identities, with additional public-artifact/tag
and clock guards. Preserve exact old/new objects, separate Owner authority,
rollback guidance and mutation receipt in protected root storage. Verify backup
before replacement, retain mode/owner/group, recheck prior hash, fsync and
independently retrieve production HTTPS bytes afterward. Root execution is an
Owner credential interaction; no passwordless sudo is available.

Prepared root operation SHA-256:
`c9137360280c995fbfdb1137ca24f93d597b277d595b59f628c540fc45083dec`.
Syntax and isolated atomic rehearsal pass: prior-object backup, replacement,
mode/receipt and repeated/unknown old-object refusal. Rehearsal is not publication.

If publication fails, preserve evidence and inspect exact current identity before
an established atomic restoration of the verified prior object. Never overwrite
an unknown object. After clients observe the new generated_at, historical bytes
fail closed; normal forward recovery requires a later Owner-reviewed/offline-signed
index. No anti-rollback bypass or unapproved signed-byte repair is authorized.
Partial Forgejo publication does not justify changing or deleting immutable tags
or historical assets. Do not claim PASS unless external proofs all succeed.

## Safety and gates

No playground VPS or QW production VPS mutation, real upgrade/rollback/Guardian
acceptance, Pro work or new product features. Community 8 PASS / 0 PARTIAL /
1 MISSING; Pro 9 PASS / 2 PARTIAL / 3 MISSING. C9 remains MISSING pending later
real supported-system clean install, authenticated official 1.3.1 bootstrap,
configuration/state preservation, rollback, Guardian recovery and required
health/readiness evidence. Task 094 authorizes no next task.

## Forgejo external proof

Release 6 is final/non-draft/non-prerelease, published 2026-09-28T17:19:19Z.
Annotated tag object 5396c8297007f92fe06a4587fe35796f6ba3abac peels to the
exact authorized source. All six assets independently retrieved anonymously with
both curl and wget match exact local sizes/hashes; both sidecar checks pass.
Full external archive manifest/provenance/notes/embedded identity and signed
companion production verification pass. Public body matches prepared metadata.
Historical Release 5 body/publication fields and asset IDs/bytes remain unchanged.

Canonical index replacement and post-publication signature/awareness verification
are pending; do not interpret Forgejo success as complete publication.
