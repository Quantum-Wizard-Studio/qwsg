# Community Email Notifications

Task 046 activates the existing Alert → Notification → Runtime → Guardian path for one optional local SMTP route. Alert Records remain the event authority; Notification delivery identity and the Guardian checkpoint queue provide deduplication, finite retry, and restart continuity.

The supported optional Configuration Model 1.0 extension is `notification.email` version `1.0`, with bounded fields `enabled`, `recipients`, `host`, `port`, `sender`, `security`, `auth`, `username`, `credential_ref`, and `timeout`. Enabled Community operation accepts exactly one address. The collection-shaped field preserves a future Pro multiple-recipient path without implementing entitlement.

Transport is `implicit_tls` or required `starttls`; certificate verification and downgrade protection cannot be disabled. Authentication is `none` or `password`. The password is stored separately at `<configuration directory>/credentials/<reference>` in a current-user-owned `0700` directory and `0600` regular file. Symlinks, special files, hard links, unsafe names, modes, and oversized values fail closed.

The default route covers unsuppressed `entered`, `escalated`, and `recovered` Alert Records at warning, critical, or emergency severity. Delivery uses three attempts within one hour, with one-minute and five-minute retry eligibility. Planning never sleeps. SMTP failure records bounded delivery evidence and never changes monitoring truth.

`qwsg notification preflight` reports the common Assessment Model 1.0 values
`satisfied`, `missing_required`, `missing_optional`,
`unknown_requires_verification`, or `incompatible`. SMTP owns its focused
detection while `internal/assessment` owns the shared classification. It is
read-only and never installs packages or changes infrastructure. Composite
readiness is available through `qwsg readiness`.

Messages contain bounded Alert metadata and local-action guidance only. Inventory, host/network identifiers, paths, mounts, full evidence, configuration, and credentials remain local.
