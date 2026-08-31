# Release Index and Release Source Contract

Task 076 implements the read-only metadata foundation approved by Task 074.
The canonical implementation is `internal/releasediscovery`. It is not an
updater: it fetches bounded public metadata, parses one strict contract,
authenticates that metadata, and evaluates it against Task 075 installed
identity. It owns no artifact acquisition, state, notification, schedule,
installation, privilege, migration execution, rollback, publication, key
custody, registration, telemetry, or Pro behavior.

## `qwsg.release-index/1`

Media type: `application/vnd.quantumwizard.qwsg-releases+json`.

The root contains exactly `schema`, `product`, `generated_at`, `channels`, and
`signatures`. Signed content contains the first four fields in that exact typed
order; `signatures` is excluded. Go's typed `encoding/json` serialization is
the normative Task 076 signature representation: compact UTF-8 JSON, declared
struct field order, input array order, normal JSON string escaping, and no
maps. The implementation publishes deterministic test vectors. A publisher
must validate the semantic document before signing these bytes.

| Element | Task 076 rule |
|---|---|
| document | UTF-8; 1–1,048,576 bytes; no trailing value; no duplicate or unknown member; nesting at most 8 |
| schema/product | exactly `qwsg.release-index/1` / `qwsg` |
| time | canonical second-resolution UTC RFC3339 (`YYYY-MM-DDTHH:MM:SSZ`); publication cannot follow index generation |
| channels | 1–4 unique channels; `stable` contains only finals, `preview` only prereleases |
| releases | 1–50 per channel; strict unique major-1 version identity across the document; status `active` or `withdrawn` |
| provenance | 40-character lowercase source commit; bounded HTTPS release-notes URL without user info, query, or fragment |
| compatibility | strict major-1 minimum source not newer than the release; at most 32 unique bounded migration identifiers |
| artifacts | 1–4 unique platforms; Community 1.3 accepts `linux-amd64`; exact `qwsg-VERSION-PLATFORM.tar.gz` name; HTTPS matching basename; 1–128 MiB; lowercase SHA-256 |
| signatures | at most 8 unique key IDs; `ed25519`; strict base64 encoding of exactly 64 bytes |

All arrays are order-significant for signing. Duplicate semantic release,
platform, route, channel, or key identities fail closed. Unknown schema,
product, algorithm, channel, platform, unsafe URL, malformed value, excessive
count, ambiguity, and conflicting active/withdrawn identities are rejected.
Withdrawn records remain authenticated historical statements but are never
eligible candidates.

## Authenticity and integrity layers

`Verifier` copies an explicit immutable caller-provided map of approved Ed25519
public keys keyed by bounded key ID. An index is authenticated only when at
least one signature verifies under an approved key. Unknown keys may coexist
for rotation compatibility but do not contribute trust. Missing, altered,
wrong-key, malformed, or otherwise unapproved metadata yields
`unauthenticated_metadata`. Authenticated documents are deep-copied on entry
and return so caller mutation cannot retain stale authenticity evidence.

Task 078 activates one exact production public authority on top of this Task
076 contract: source ID `community-release-index`, endpoint
`https://releases.quantumwizard.hu/qwsg/v1/release-index.json`, and key ID
`qwsg-community-release-2026-01`. The bundled `qwsg.release-trust/1` record is
strictly decoded and must reproduce the compiled 32-byte public key and
fingerprint. It contains no private material. Test private keys remain
unmistakably non-production fixtures only.

Trust layers remain separate:

1. HTTPS authenticates the approved transport endpoint.
2. Ed25519 authenticates the complete release-index semantics.
3. The signed artifact size and SHA-256 bind the selected bytes.
4. Existing acquisition verifies the sidecar and archive bytes.
5. Existing package verification validates `MANIFEST.sha256`, layout, modes,
   platform, and embedded `qwsg.release/1` provenance.
6. The local declared migration registry determines executable compatibility.

SHA-256 does not authenticate the publisher. HTTPS evidence is never reported
as Ed25519 evidence, and manifest migration claims never create a local route.

## Source-neutral retrieval

The interface is:

```text
Fetch(context, {channel, validators})
    -> {not_modified | manifest bytes, source evidence} | safe failure
```

`StaticHTTPSource` is the initial adapter. Construction fixes one HTTPS URL,
one safe public source ID, its exact authority, and its exact path. It does not
implement fallback. Redirects are limited to three and must retain HTTPS, the
exact authority, and the exact path. Encoded/dot-segment paths,
URL credentials, query strings, fragments, insecure TLS, client certificates,
custom TLS dialers, cookie jars, proxy credentials, and non-HTTP transports are
refused or removed. TLS 1.2 is the minimum.

The adapter uses bounded connection/TLS/header/idle/overall timeouts, a 1 MiB
body limit, strict media type, safe conditional validators, cancellation, and
fixed anonymous request headers. It sends no installed version, channel query,
hostname, IP inventory, MAC address, email, account, Guardian finding, service
inventory, or server state. Ordinary HTTPS still exposes network-level source
information such as source IP to the destination and intervening network
infrastructure; QWSG does not describe that as mathematical anonymity.

Stable failures are `source_authority_refused`, `source_canceled`,
`source_timeout`, `source_transport_failed`, `source_http_status`,
`source_too_large`, and `source_media_type`. Errors never include response
bodies, credentials, URLs, host state, or raw transport diagnostics.

`304 Not Modified` returns only safe source/validator evidence. Task 077 accepts
it only when bound to a matching previously authenticated awareness observation;
validators never become authenticity evidence.

## Deterministic evaluation

`Discoverer.Check` enforces the only canonical order:

```text
fetch -> strict parse -> Ed25519 authenticate -> installed-aware evaluation
```

The evaluator calls Task 075's canonical classifier without a candidate and
requires `verified_supported_installation`. It selects the explicit channel,
filters active platform/prerelease-eligible releases, and chooses the greatest
strict version. For a newer candidate it calls the same classifier with that
candidate. Compatibility is `supported` only when:

- installed version meets the signed minimum-source advisory;
- Task 075 returns `supported_upgrade_source` from the local migration
  registry; and
- the signed index advertises that exact local migration ID.

Thus remote metadata can narrow advice but cannot authorize executable
compatibility. Legacy, absent, unknown, inconsistent, incomplete, binary-only,
or otherwise unverified local identity stops evaluation. Equal/older releases
remain truthful relations and require no migration conclusion.

Task 078 constructs this source and verifier for `qwsg update check`. The QWSG
1.2 updater, acquisition, package verification, transaction, and rollback
remain authoritative and unchanged. Task 077 consumes the result through the
same checker seam; it does not create another updater. `update status` remains
network-free.
