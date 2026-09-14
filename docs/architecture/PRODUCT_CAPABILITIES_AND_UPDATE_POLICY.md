# Product capabilities and update policy

Task 083 establishes policy eligibility only. No automatic update execution,
new scheduler, Guardian installation action or production Pro licensing exists.
Community remains the complete local toolkit: observation, understanding,
comparison, reporting, discovery, authenticated verification and explicit update.
Pro adds automation eligibility around the same core.

## Authority boundary

`internal/productcapability` is the single product capability boundary. Its
immutable `Set` answers `Has(update.manual)` and `Has(update.automatic)`;
unknown capabilities and the zero Set grant nothing. `Resolve(nil)` supplies
the explicit built-in Community baseline: manual allowed, automatic absent.
A supplied declaration must use `qwsg.product-capabilities/1`, a known product
(`community` or `pro`), and unique known grants including manual. Community
cannot declare automatic; unknown product/schema/grants, duplicate grants,
missing manual or conflicting product/grants return an error and zero authority.
Pro does not implicitly grant automatic: its declaration must include it.

A declaration is trusted adapter input, **not a license proof**. The shipped
CLI composition `installationCapabilities` always resolves the Community
baseline. Only Go test fixtures inject a Pro declaration; there is no production
file, environment variable, flag, remote API or account that grants Pro.
A later authenticated entitlement adapter must validate its own authority and
supply the declaration through this seam. Adapter errors must propagate rather
than be converted to Pro authority. Commercial transport, offline activation,
revocation and licensing policy remain separate future work.

## Configuration and policy

The existing Configuration 1.0 extension remains canonical, with its existing
identity, precedence and atomic private storage behavior:

```json
{"id":"installer.update-policy","version":"1.0","required":false,"fields":{"policy":"manual"}}
```

| Configured value | Effective installation policy | Required capability | Notification intent |
|---|---|---|---|
| extension absent / `manual` | `manual` | `update.manual` | off |
| existing `notify` | `manual` | `update.manual` | on, existing SMTP prerequisites apply |
| `automatic` | `automatic`, policy eligibility only | `update.automatic` and manual | off |

`configuration.UpdatePolicy` extracts intent; `updatepolicy.Parse` owns the
strict extension shape. Unknown values, missing/extra fields, unsupported
versions, required declarations and duplicates refuse. Existing configurations
without this extension retain their bytes and canonical identities and resolve
to manual. No new model field, schema migration, credentials or file-permission
changes are introduced. Existing notify settings remain compatible.

`updatepolicy.Evaluate` joins explicit intent with validated product authority
and returns presentation-independent effective policy and capability state.
Unknown/zero authority or policy and automatic without capability refuse,
returning no usable state. CLI configuration resolution/set/validate use this
boundary; a configuration file cannot grant entitlement. `qwsg config get
update.policy` retains the configured legacy value. `qwsg update status` shows
effective manual/automatic policy, product, manual capability and automatic
policy permitted/unavailable without networking or installation. Configuration
errors use bounded diagnostics and never print declaration/credential content.
An injected Pro authority lost while configured automatic causes validation to
refuse; it never silently enables automatic behavior or rewrites configuration.

## Shared update security and future orchestration

The explicit update entry checks manual capability. Product policy does not
construct or substitute an `updateauthority.Candidate`. Task 082's authenticated
release and compatibility authority, locally implemented migration, exact
source/target/platform checks, package provenance, artifact identity/integrity,
deterministic transaction, privileged reauthentication, post-validation and
rollback remain mandatory and unchanged. Both editions use this same engine.
Tests validate Pro policy then present unsigned, tampered, unsupported and
source-mismatched candidates to that common authority gate; all refuse.

Task 084 is responsible for separately authorized automatic orchestration. It
must consume fresh capability/policy eligibility and invoke the same validated
update path with all service, configuration and rollback safeguards. A policy
State is not an execution token. This task adds no unattended apply entry point,
schedules, windows, channels, reboot or automatic rollback orchestration.
Task 084 has not been started. No release is published by Task 083.
