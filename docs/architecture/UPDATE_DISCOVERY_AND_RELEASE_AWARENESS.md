# Update Discovery and Release Awareness Architecture

## Status and scope

This document defines the approved QWSG 1.3 architecture for public Community
release discovery and local release awareness. Task 074 is architecture and
analysis only: it adds no production behavior, network service, release
manifest, signing key, telemetry, registration, installation, or Pro feature.

The default policy is:

```text
automatic check         allowed
automatic notification  allowed when configured
automatic installation  forbidden
```

Discovery is a read-only ordinary-user operation. Installation remains an
explicit operator action which reuses the QWSG 1.2.0 updater. No central
service logs in to, controls, or mutates a monitored server.

## Repository-backed QWSG 1.2.0 inventory

| Concern | Actual 1.2.0 implementation | 1.3 disposition |
| --- | --- | --- |
| Running identity | `cmd/qwsg` embeds `version`, full build commit and UTC build date; `version` reports them. `installedVersion` executes `/usr/local/bin/qwsg version` and parses its first line. | Reuse embedded provenance, but classify installation identity before treating a path as a supported package. |
| Source identity | Repository `VERSION`; deterministic build injects version, commit and date. | Reuse unchanged as build input, not as runtime discovery state. |
| Release provenance | Package `RELEASE.json` uses `qwsg.release/1` with version, commit, build time and `linux-amd64`. | Reuse for downloaded-package verification; do not confuse it with the public discovery contract. |
| Package integrity | Release contains `MANIFEST.sha256`; `VerifyPackage` checks bounded single-root layout, regular members, required files, every manifest hash, modes, platform and embedded provenance. | Reuse unchanged after acquisition. |
| Public discovery | `internal/update.Discover` calls the canonical Forgejo releases API, reads at most 1 MiB/50 records, accepts exact tags/assets, HTTPS URLs, final/prerelease consistency, two assets and bounded sizes. | Retain as a compatibility source adapter; place it behind the new source-neutral manifest contract. |
| Artifact acquisition | Private mode-0700 staging; bounded HTTPS client, same-host redirects, declared sizes, archive and sidecar download, SHA-256 verification. | Reuse; add manifest-declared digest as an independently validated input. |
| Version ordering | Strict SemVer-like parser/comparator rejects whitespace, build metadata, leading zeroes and malformed identifiers; update supports major 1 only. | Reuse comparison logic; add explicit release-channel eligibility before comparison. |
| Update CLI | `qwsg update check`, `qwsg update`, `qwsg update status`, `qwsg update rollback`; `update` without a local archive discovers and installs after explicit invocation. | Preserve compatibility. Make `check` the read-only refresh, `status` a network-free state read, and document `update install` as a possible explicit alias only after a later CLI decision. Do not create a second updater. |
| Update policy | Configuration extension `installer.update-policy/1` stores `manual` or `notify`; guided installer localizes the choice and states that privileged automatic updates stay disabled. | Reuse vocabulary. Both policies may permit checks; only `notify` permits update-availability delivery. Neither permits installation. |
| Migration | Declarative exact routes in `internal/update`; plans validate schema compatibility and preservation. Unknown paths fail before mutation and are revalidated at the privileged boundary. | Manifest advertises route identifiers and minimum source as hints; the installed local registry remains authoritative for install eligibility. |
| Transaction/rollback | Fixed destination allowlist, package backup, integrity metadata, service-intent restoration, installed-version validation, automatic failure rollback and explicit `update rollback`. User configuration, credentials and state are excluded from package backup. | Reuse unchanged. Discovery never writes transaction state or crosses the privileged boundary. |
| Guardian recurrence/state | Runtime Service supplies local recurrence; Guardian persists one integrity-protected checkpoint containing Runtime, Alert and Notification queue state. Five-minute observe interval, two-minute cycle timeout and resource limits are local-monitoring contracts. | Add a separate, low-frequency discovery schedule and separate small update state. Do not put network work in every observe cycle or make it part of Guardian health. |
| Notifications | Canonical Alert/Notification provides provider-neutral durable queue/idempotency. Community SMTP is an existing adapter. `changenotification` supplies localized lifecycle messages and only process-local duplicate suppression. | Reuse transport, credential, locale and delivery result boundaries. Use persistent update transition identity; do not rely on `changenotification.Dispatcher.seen`. |
| SMTP diagnostics | SMTP returns bounded evidence such as `smtp_server_accepted` or `smtp_delivery_failed`; only some failures are categorized. | Later work may add safe categories. Discovery must neither expose raw SMTP text nor change delivery truth. |
| Readiness | Assessment distinguishes satisfied, missing, incompatible and unknown evidence. Optional `notification.external` can keep overall readiness partial. | Update availability is information, not readiness or health. Check failure must not create a false unhealthy verdict. |
| Configuration/localization | Canonical Configuration Source/Effective Configuration owns values and provenance; EN/HU/DE installer/change catalogs exist. | Extend the canonical configuration and catalogs in later implementation; no hard-coded user text in the discovery core. |
| Security/privacy | Ordinary-user local core, private atomic state, credential references, bounded subprocess/network use and sanitized output. | Preserve these boundaries and add an explicit minimal outbound request/privacy contract. |

There is no release channel configuration, versioned public manifest schema,
metadata signature, persistent check result, last-success timestamp, durable
update-availability deduplication, withdrawal state, or Guardian update-check
schedule in 1.2.0. The current `update check` is network-only and its output is
not a durable status record. The current Forgejo response and TLS connection
establish the discovery authority operationally; the downloaded SHA-256
sidecar establishes byte integrity but cannot independently prove publisher
authority.

## Ownership and non-coupling boundaries

The dedicated `Update Discovery` subsystem owns only source retrieval,
manifest validation, channel selection, release comparison, compatibility
advisory evaluation, and proposed Update Awareness State. It must not own:

- installed-package classification, file copying, privilege escalation,
  Guardian stop/start, migration execution, rollback or uninstall;
- host Inventory, Drift, Health, Rule, Policy, Report or readiness evaluation;
- SMTP, credentials, localized rendering or provider retry semantics;
- Guardian's primary local monitoring schedule or checkpoint truth;
- release publication, signing-key custody, account registration, telemetry,
  central inventory, license enforcement or fleet control.

The intended flow is:

```text
public HTTPS source -> source adapter -> canonical manifest validation
                    -> channel/version/compatibility evaluation
                    -> proposed Update Awareness State -> atomic state store
                                                   |
                                                   +-> status presentation
                                                   `-> transition event -> existing notification delivery

explicit operator install -> existing acquire/package/migration/transaction/rollback path
```

## Installed identity and installer boundary

File presence is not installation proof. Future implementation must classify
the local state before discovery or installation advice:

| Classification | Required meaning | Update consequence |
| --- | --- | --- |
| `arbitrary_or_legacy_binary` | A path named `qwsg` exists but lacks supported package provenance or has a historical/pre-alpha identity. | Never infer installed/supported; report bounded operator guidance. |
| `supported_installed_package` | Expected package-owned files and supported embedded version/provenance are present. | May be compared with releases. |
| `verified_installation` | Supported package identity plus required installed artifact and configuration/state safety checks pass. | Eligible for compatibility planning. |
| `supported_upgrade_source` | Verified installation version has an exact locally declared route to the candidate. | Eligible for explicit installation. |
| `unsupported_or_unknown_installation` | Conflicting, malformed, incomplete, unverifiable or unsupported identity. | Discovery may report releases, but installation fails closed. |

The field incident on `server.quantumwizard.hu` proved this boundary: a
historical `/usr/local/bin/qwsg` reporting `QWSG 0.0.1-prealpha` caused the
guided 1.2.0 installer to classify the host as already installed and skip
package installation. Guardian could not activate. After the operator safely
backed up and removed the legacy binary, the final installer correctly selected
new installation. Task 075 fixes this boundary with the canonical evidence-based
installed-package classifier. Future update discovery must consume that
verified local result and must not fall back to binary presence. It must make
classification evidence-based and must never automatically delete or replace
unknown legacy material.

## Canonical release source abstraction

Clients consume one canonical `ReleaseSource` result, independent of transport:

```text
Fetch(channel, validators, deadline) -> NotModified | Manifest bytes + source evidence
```

The client owns strict decoding and policy. An adapter owns endpoint syntax and
transport translation only. This permits three sources without redesigning
comparison, state, Guardian or notification code:

1. **Static Quantum Wizard manifest — recommended Community authority.** A
   small immutable/versioned JSON document is simple to cache, mirror, audit
   and serve without an account. Release assets may remain on Forgejo.
2. **Forgejo Releases API — compatibility/fallback adapter.** It is already
   implemented and useful, but its provider-specific shape lacks QWSG
   compatibility, release-date/signature policy and withdrawal semantics.
   It must not be scraped from HTML.
3. **Future Quantum Wizard Release Service — optional adapter.** It may return
   the same canonical manifest contract and later add authenticated Pro APIs.
   Community public discovery remains credential-free and independent of Pro.

Source fallback is policy, not silent trust expansion. A source configuration
pins an approved authority set and priority. Failure of the primary source must
not cause acceptance from an unapproved host. Cache validators such as ETag may
reduce transfer, but a `304` is usable only with a previously authenticated,
validated local manifest.

## Proposed manifest contract

Media type: `application/vnd.quantumwizard.qwsg-releases+json`.
Schema identity: `qwsg.release-index/1`. JSON is UTF-8, bounded, decoded with
unknown fields rejected for schema 1, and contains one product plus a bounded
set of channels/releases. JSON member ordering is irrelevant; signature input
uses a specified canonical JSON serialization or detached digest file, never
the received byte layout by accident.

```json
{
  "schema": "qwsg.release-index/1",
  "product": "qwsg",
  "generated_at": "2026-08-29T18:00:00Z",
  "channels": [
    {
      "name": "stable",
      "releases": [
        {
          "version": "1.2.0",
          "published_at": "2026-08-29T16:53:34Z",
          "status": "active",
          "source_commit": "348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2",
          "release_notes_url": "https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/tag/v1.2.0",
          "minimum_source_version": "1.1.0",
          "migration_routes": ["compat-1.1.0-to-1.2.0", "compat-1.2.0-rc.2-to-1.2.0"],
          "artifacts": [
            {
              "platform": "linux-amd64",
              "name": "qwsg-1.2.0-linux-amd64.tar.gz",
              "url": "https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v1.2.0/qwsg-1.2.0-linux-amd64.tar.gz",
              "size": 3524214,
              "sha256": "44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11"
            }
          ]
        }
      ]
    }
  ],
  "signatures": [
    {"algorithm": "ed25519", "key_id": "community-release-2026", "value": "BASE64"}
  ]
}
```

Contract rules:

- `schema` and `product` prevent cross-contract/product substitution.
- `generated_at` is audit/cache information, not freshness proof by itself.
- channel names are stable tokens. QWSG 1.3 Community defaults to `stable`;
  prerelease/preview channels require explicit opt-in and never outrank stable
  by raw version alone.
- `status` is `active` or `withdrawn`. Withdrawn entries remain addressable for
  state reconciliation but are never new install candidates.
- `version` follows the existing strict version parser. A channel must not
  contain duplicate semantic identities.
- `published_at`, `source_commit` and `release_notes_url` are required,
  bounded, safe references. Release notes are never executed or used as trust.
- `minimum_source_version` is a coarse publisher statement.
  `migration_routes` identifies declared routes. Neither overrides the local
  migration registry; install eligibility requires both manifest consistency
  and a locally supported exact plan.
- `artifacts` is bounded and uniquely keyed by platform. Name, HTTPS URL,
  positive bounded size and lowercase SHA-256 are required. QWSG 1.3 accepts
  only `linux-amd64` and approved hosts/paths.
- signatures cover all fields except the `signatures` member under the defined
  canonicalization. Duplicate key IDs, unknown required algorithms, malformed
  signatures, ambiguity, unknown schema/product, excessive size/counts, unsafe
  URLs or conflicting releases fail closed.

The schema deliberately excludes hostname, server identity, entitlement,
telemetry, installed-service inventory and notification destination.

## Trust model

HTTPS with normal certificate validation, strict endpoint allowlisting,
bounded redirects and response limits protects transport and is sufficient for
an initial Community bootstrap only when the endpoint is controlled as an
authoritative Quantum Wizard publishing surface. SHA-256 then binds the
selected artifact bytes. SHA-256 alone does not authenticate the manifest.

The recommended incremental Community trust model is:

1. ship one or more release-manifest public keys/key IDs in QWSG;
2. publish the detached/in-document Ed25519 signature with the static manifest;
3. validate signature, schema, product, channel and constraints before state
   publication;
4. validate artifact size/SHA-256 during existing acquisition;
5. validate internal package manifest and embedded release provenance through
   the existing verifier before any migration or privileged action.

This is small enough for Community and separates publisher authenticity from
artifact integrity. Key introduction, rotation, revocation, threshold policy
and recovery are Owner-reserved and require a dedicated implementation task.
Until signing is approved and implemented, state must truthfully record
`transport_authenticated` rather than claim cryptographic manifest signing.
Forgejo API fallback must never be described as independently signed.

## Evaluation and compatibility

Evaluation order is deterministic: validate installed classification; validate
source evidence and manifest; select configured channel; exclude withdrawn or
ineligible prereleases/platforms; choose the greatest eligible strict version;
compare with installed; evaluate manifest compatibility advisory; ask the
existing local migration planner whether an exact route exists.

Discovery can report `update_available_unsupported_source` without offering
installation. It must not synthesize migration paths or equate
`minimum_source_version` with sufficient compatibility. Equal, older,
malformed, unsupported-major and ambiguous candidates remain distinct.

## Persistent Update Awareness State

Use a dedicated private atomic integrity-protected record under the canonical
QWSG state root, for example `update/awareness.json`, with schema
`qwsg.update-awareness/1`. It is not the Guardian checkpoint or rollback
transaction. Store only:

- configured source identity and channel;
- installed classification and version (no binary path is necessary);
- current state: `never_checked`, `current`, `update_available`,
  `update_available_unsupported_source`, `withdrawn`, or `unknown`;
- selected release version, publication time, manifest identity/digest and
  compatibility result;
- `last_attempt_at`, `last_success_at`, last bounded failure category and
  consecutive-failure count;
- last transition identity, last notification release/state identity, and
  notification outcome time/status;
- cache validators only when bound to the validated source/manifest.

Never persist raw responses, HTTP headers, IPs, credentials, tokens, hostnames,
email addresses, full release notes, SMTP errors or Guardian findings. Writes
reuse existing private-directory, size, integrity, validation, temporary-file,
fsync and atomic-rename patterns without coupling the record to Inventory or
Current Operator State.

`qwsg update check` performs a bounded network refresh and prints the new
result. `qwsg update status` performs no network access and renders the last
validated state, timestamps and staleness/failure separately. A later CLI task
may add JSON output. Existing bare `qwsg update` stays an explicit install
operation for compatibility; `qwsg update install` may become a clearer alias
only by delegating to the identical updater.

## Guardian integration

Guardian should add a separate optional update-discovery job, not another
Runtime/Pipeline engineering stage. Recommended Community defaults:

- enabled automatic checks with a 24-hour nominal interval and deterministic
  per-installation jitter up to 60 minutes to avoid synchronized load;
- minimum enforced interval of one hour and configurable interval/channel;
- one in-flight check, 30-second HTTP client timeout, 35-second whole-operation
  deadline, 1 MiB metadata limit, 50 release limit and existing redirect/host
  constraints;
- no retry loop inside a Guardian cycle. The next scheduled opportunity uses
  bounded exponential failure spacing capped at the configured interval;
- run only after local Guardian startup responsibilities and never hold the
  local observation lock or retain loaded Inventory graphs;
- persist success/failure independently. Check errors are diagnostics and
  stale awareness, never a Guardian unhealthy/degraded verdict by themselves;
- cancellation on shutdown and no inbound listener, remote shell, central root
  access, registration or mandatory telemetry.

The primary five-minute local monitoring cadence continues even if DNS, TLS,
Forgejo, the static manifest host or a future release service is unavailable.

## Notification transition model

Define a localized, privacy-bounded `update_available` lifecycle event with:
installed version, available version, channel, release date, compatibility
status, and `administrator_action_required=true`. It contains no artifact URL,
hostname requirement, credentials, tokens, email address, inventory or server
findings. Delivery uses the existing configured notification transport.

The durable idempotency key is derived from product, channel, installed
version, available version, manifest identity and transition kind. Notify only
after an atomically published transition into a new actionable identity:

| Observation | State/notification behavior |
| --- | --- |
| A newer release first appears | `current -> update_available`; notify once when policy is `notify`. |
| Administrator ignores it | Rechecks update timestamps but do not repeat the same event. |
| Another newer release appears | New release identity; notify once. |
| Server is updated | Installed-version change reevaluates to `current` or a still-newer release; clear obsolete actionable identity without a “success” claim unless the updater supplied it. |
| Source unavailable/malformed/untrusted | Preserve last validated release as stale, record bounded failure; no update-available replay and normally no outage email. |
| Release withdrawn | Transition the matching candidate to `withdrawn`; do not install or keep advertising it. A separately designed security advisory may notify once, but Task 074 does not invent that event. |
| Notification fails | Preserve update state and provider result. Existing provider retry policy may retry the same idempotency key; a later check does not create a new event. |

The existing lifecycle dispatcher is reusable for rendering/transport
composition, but its memory-only `seen` map is insufficient across Guardian
restart. The Update Awareness Store supplies durable transition/delivery
identity; canonical Notification queue semantics remain the transport retry
authority.

## Failure model

Network timeout/DNS/TLS errors, non-2xx status, source outage, oversized body,
malformed JSON, unknown schema, invalid signature, invalid channel/release,
unsafe URL, incompatible platform and ambiguous identities use stable safe
categories. User output never includes raw response bodies, certificate data,
proxy details or tokens. Last attempt and last successful validation are
reported separately. No failure replaces the last known-good manifest or makes
local monitoring unhealthy. Invalid authenticity, metadata or identity fails
closed and cannot reach acquisition.

SMTP field finding B is a notification diagnostics concern, not discovery.
The observed `smtp_delivery_failed` concealed an Exim sender-verification
failure. Future work may map internal errors to safe structured categories such
as `network_timeout`, `tls_validation`, `authentication`, `sender_rejected`,
`recipient_rejected`, `message_rejected`, and `outcome_unknown`; raw SMTP
responses, addresses and credentials remain private. These categories must not
promise mailbox delivery.

Field finding C confirms the readiness model: local configuration can be
validated locally; successful SMTP transaction can prove only server
acceptance; administrator evidence may record an externally verified mailbox
receipt at a time; otherwise end-to-end delivery is inherently unverifiable.
Therefore `notification.external = unknown_requires_verification` and overall
`partial` can remain correct even after a separate manual receipt was observed
unless a canonical bounded external-evidence contract records it. No automatic
update notification may turn acceptance into a false `delivered` or `ready`.

## Privacy model

Anonymous Community discovery requires no account or credential and sends only
an ordinary HTTPS request for a public manifest, plus protocol headers needed
for HTTP/TLS and optional cache validators. QWSG does not add hostname, MAC
address, IP inventory, installed services, user identity, email address,
Guardian findings or server state. The request may inherently reveal source IP
and ordinary network metadata to the server, CDN, DNS resolver and network
operators; it is privacy-minimized and unregistered, not mathematically
anonymous. Logs and retention at an authoritative service require a separately
published privacy policy.

## Community and future central/Pro boundary

QWSG 1.3 Community may implement credential-free public checks, local state,
local status and optional existing-channel notification. It has no account,
mandatory telemetry or central inventory. An optional central email release
announcement subscription is a separate human subscription and is not server
discovery state.

Future Pro may add opt-in API-key registration, server inventory, a central
console, staged update management and channel-neutral QUWIP, Telegram or other
notification adapters. Those services consume versioned public contracts; they
do not gain remote shell/root authority and cannot silently install updates.
Their identity, consent, tenant, privacy, revocation, audit and deletion models
require later Owner approval. None is part of the QWSG 1.3 Community
implementation boundary.

## Recommended implementation sequence

No item starts automatically after Task 074:

1. installation identity/classification contract and legacy-boundary tests;
2. release-index schema, strict parser, source abstraction and static manifest
   publication/signing decision;
3. private Update Awareness Store and network-free status contract;
4. refactor existing Forgejo discovery behind the source adapter while
   preserving updater behavior;
5. CLI check/status integration and compatibility/migration advisory;
6. Guardian low-frequency isolated scheduling and resource/failure tests;
7. persistent transition-to-notification integration and localization;
8. safe SMTP diagnostic categories and optional external-delivery evidence as
   separate tasks;
9. clean-host/privacy/outage/withdrawal/update/rollback acceptance before any
   QWSG 1.3 release decision.
