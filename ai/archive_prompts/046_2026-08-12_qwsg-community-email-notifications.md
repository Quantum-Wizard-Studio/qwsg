# Current Engineering Task 046: QWSG Community Email Notifications

## Task Metadata

- Task ID: `046`
- Task slug: `qwsg-community-email-notifications`
- Status: `complete`
- Date opened: `2026-08-12` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG Community Email Notifications


## Objective

Implement the first production local operator-notification capability by activating the existing Canonical Alert, Notification Delivery, Runtime, and Guardian contracts for standards-compatible SMTP email. A Community operator must be able to configure exactly one administrator recipient, keep SMTP credentials outside public `config.json`, validate readiness, safely store or update a private credential, send an explicit test message, and receive bounded privacy-safe messages for selected actionable Alert lifecycle transitions. Notification must remain optional, deterministic, restart-persistent, storm-resistant, and independent of QWS accounts, APIs, subscriptions, managed services, or Internet availability. Monitoring truth and local Community operation must continue when notification is disabled, incomplete, offline, or failing. Preserve a compatible future path for Pro to activate multiple recipients without changing the public configuration contract, while implementing no Pro entitlement behavior.


## Scope

- Re-audit and record the canonical Alert lifecycle, Notification Delivery Model 1.0, Runtime coordinator, Guardian checkpoint/service loop, Task 045 Configuration Contract and filesystem store, CLI/setup surface, state epochs, operator projection, packaging, systemd sandbox, privacy policy, and uninstall behavior before design or modification.
- Reuse `internal/alert` as the sole authority for attention events and `internal/notification` as the sole routing, deduplication, retry, idempotency, attempt/status, acknowledgement, evidence, and queue-state authority. Do not create a second event or delivery state machine.
- Extend the canonical Task 045 configuration model compatibly with typed, versioned notification/email settings sufficient for enabled state, one Community administrator recipient, SMTP host and port, transport-security mode, authentication mode, sender identity, opaque credential reference, bounded connection/delivery timeouts, retry policy, and explicit alert event/severity preferences.
- Replace the inert Task 045 notification-recipient extension with an explicit supported configuration representation only through a documented compatible contract amendment. Preserve deterministic normalization, precedence, provenance, strict unknown-field rejection, bounded values, canonical identities, and fail-before-runtime-mutation behavior.
- Enforce Community cardinality at every public activation boundary: configuration decode/normalize/resolve, setup/config mutation, readiness validation, CLI test, and Guardian startup must accept zero recipients only while notification is disabled and exactly one recipient while enabled; multiple Community recipients must fail closed. Do not infer, query, or implement a Pro entitlement.
- Preserve a future ordered recipient collection and provider-neutral endpoint mapping so a separately authorized Pro layer can validate an unlimited-by-entitlement recipient set subject to global safety/resource bounds without changing Community keys or stored meaning. Do not expose a bypass flag or hidden multi-recipient Community path.
- Implement the smallest private per-user credential mechanism consistent with Task 045 secret references, under a separate canonical credential location in the per-user QWSG configuration domain. Store only the SMTP authentication secret needed by the supported mode; public configuration stores an opaque reference, never the value.
- Require credential directories to be current-user-owned mode `0700` and credential files current-user-owned regular files mode `0600`; reject symlinks in every component, hard-link ambiguity where defensible, special files, wrong ownership, permissive modes, oversized content, NUL/control data, unsafe names, and race-detected replacement. Use bounded reads and same-directory atomic write, file sync, rename, and directory sync for supported mutation.
- Ensure secrets never appear in human or JSON config/setup/notification output, command arguments, environment variables, logs, diagnostics, Guardian checkpoint/queue, Runtime results, Notification requests/evidence, Console state, snapshots, tests, documentation, or release/package payloads. Diagnostics may expose only stable privacy-safe reason tokens and non-secret configuration field names.
- Implement one concrete SMTP provider behind the existing `notification.Provider` interface. Only this provider may initiate outbound connections, only for an explicit test or a due Guardian delivery request, and only to the validated configured SMTP host/port.
- Support standards-compatible SMTP without hard-coding a commercial provider. Define explicit transport modes at minimum for implicit TLS and required STARTTLS. Reject silent TLS downgrade. Plaintext SMTP must not be enabled by default; if investigation demonstrates a necessary local-only mode, it requires explicit opt-in, must reject credential authentication over plaintext, and must be documented as unsafe and excluded from production acceptance. System certificate verification and configured server-name validation must remain enabled; no insecure-skip-verify option is allowed.
- Support the smallest evidence-backed authentication modes, expected to be no authentication and username/password authentication through the private credential reference. Any supported authentication must occur only after authenticated TLS unless a no-auth local test mode is explicitly selected. Do not add OAuth, provider SDKs, arbitrary auth plugins, or credentials on process arguments.
- Bound DNS/connect, TLS handshake, SMTP commands, payload size, and total delivery by explicit validated timeouts and contexts. No inbound listener, arbitrary URL, HTTP client, shell, subprocess, unbounded goroutine, background worker, or uncontrolled network capability may be added.
- Create a concise deterministic email renderer from immutable Alert Records. Include QWSG identity, clearly bounded severity/category/event/reason summary, UTC event timestamp, and a recommendation to inspect details locally. Exclude raw inventory, addresses, host identifiers, paths, mount details, complete evidence, configuration values, credentials, and arbitrary unbounded error strings.
- Define the initial Community notification policy from existing Alert semantics rather than all degraded states. Route selected unsuppressed actionable lifecycle records such as activation, escalation, material change, and useful recovery only after an evidence-backed severity/event review. Preserve Alert-owned suppression, stale/partial/unknown semantics and document the exact default matrix. Do not synthesize emails directly from raw Health, Policy, Report, Console, or Guardian-cycle values.
- Use existing Alert Record identity plus Notification delivery identity as the deduplication boundary. Repeated identical Guardian conditions that emit no new Alert Record must produce no request; replay of an existing record/route/endpoint must reuse its delivery identity and must not duplicate successful or terminal delivery.
- Use the existing durable Guardian `NotificationQueueState` for restart continuity. Preserve immutable attempts and deterministic next-attempt scheduling. Select a small bounded retry policy with finite attempts, increasing backoff, and an absolute delivery window. No sleeping inside a Guardian cycle, jitter, tight loop, or retry outside the canonical planner.
- Classify connection refusal, timeout, temporary SMTP rejection, rate/availability conditions, TLS/certificate failure, authentication failure, invalid recipient/sender, and permanent rejection into existing retryable/terminal/indeterminate failure classes conservatively. Never claim delivery from mere socket write; acknowledgements mean only the existing provider-level semantics.
- Keep delivery failure separate from monitoring/evaluation truth. SMTP failure must produce bounded observable notification status/evidence and persist queue state without crashing Guardian, stopping future monitoring cycles, recursively creating a notification about notification failure, or claiming that an email was delivered.
- Determine and implement the smallest durable operator-visible notification readiness/last-attempt projection that does not expose secrets or overwrite canonical health truth. Console must remain usable during notification failure; detailed evidence remains local.
- Apply the Owner-approved deterministic preflight-first model to Task 046 notification readiness only. Before notification test or Guardian activation, classify each supported SMTP/configuration/credential/TLS-trust/network prerequisite as `satisfied`, `missing_required`, `missing_optional`, `unknown_requires_verification`, or `incompatible`. Prefer `unknown_requires_verification` over inference when a capability cannot be established reliably. Required failures must stop the affected notification test or activation safely, identify the exact requirement, and, only for explicitly supported platforms where the command is deterministic, print the exact recommended operator-run remediation command.
- Keep preflight detection and recommendation read-only. Task 046 must not execute package-manager commands, install dependencies, modify trust stores, enable infrastructure, change firewall/network policy, or otherwise perform remediation. Recommendations must be fixed, platform-gated, bounded data selected by validated detection—not constructed from untrusted host/configuration text—and must clearly state that the operator performs the action and reruns validation.
- Add the smallest CLI/setup integration consistent with Task 045: extend `setup` and typed `config get|set|show|validate` keys; add one bounded credential update operation that reads a secret from a terminal without echo or from a protected file descriptor/file only if its safety can be proven; and add `qwsg notification test` plus a readiness/status view if not already covered by config output.
- Never accept a secret directly as a CLI positional value or flag. Noninteractive credential provisioning must use a documented protected input channel, must refuse terminals where inappropriate, and must not echo or serialize the secret.
- Make test email explicitly operator-triggered, clearly labeled as a test, use the same validated configuration/credential/provider path as Guardian, and record no fake Alert/Health/Policy evidence or incident state. Its result must use bounded privacy-safe status and exit codes.
- Resolve complete notification configuration and credential readiness before Guardian lifecycle/checkpoint mutation when notification is enabled. Disabled notification and an absent SMTP configuration remain valid. Enabled but materially incomplete/unsafe configuration must fail the existing `config validate`/systemd `ExecCondition` gate without restart churn.
- Audit the systemd unit before change. Preserve non-root execution, `NoNewPrivileges`, filesystem protections, private state, resource limits, start limiting, and all unrelated sandbox directives. If outbound SMTP is already permitted, do not change the unit merely to signal capability. If a change is strictly required, limit it to the minimum outbound address-family/network allowance supported by configurable SMTP destinations and document why destination restriction must remain application-enforced.
- Keep `install.sh` noninteractive. It may point operators to setup/configuration, private credential provisioning, notification test, explicit Guardian activation, and local verification, but must not collect credentials, contact SMTP, enable notification, or start/restart the user service.
- Define and run controlled local SMTP acceptance without public-Internet delivery or adding an unreviewed runtime dependency. Test infrastructure may implement a bounded loopback SMTP endpoint in Go tests and isolated acceptance tooling; it must never ship as a product listener.
- Preserve local/offline Community monitoring when notification is disabled, SMTP is absent, DNS/network is unavailable, authentication fails, or delivery is exhausted. No QWS remote dependency may become required.
- Preserve existing QWSG observation, bootstrap, partial-evidence, scheduling, Console, setup/config, state-root, systemd/restart, install/uninstall, privacy, and no-root guarantees.
- Update only the minimum canonical architecture, configuration reference, setup, operations, troubleshooting, security/privacy, upgrade/uninstall, CLI, package guidance, roadmap/history, and Task 046 lifecycle documents needed to make the feature coherent and non-contradictory.
- Treat Task 046 as post-1.0 development suitable for a future 1.1.0 line, but do not change `VERSION`, create a release identity, rebuild/republish `v1.0.0`, tag, push, or publish without separate Owner authorization.
- Preserve byte-for-byte the annotated `v1.0.0` tag, its release-source commit `177535e44b2ce5ed9efd73ab0793ffe6881f0cd6`, published artifact SHA-256 `edfba7366adf2c1ce0a8ce56369bb0dc5ad11326c4e3d1e301625a5313292fa5`, LICENSE, Task 045 completed baseline, and unrelated Owner draft `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`.


## Out of Scope

- No QWS-managed email delivery, QWS cloud/API backend, QWS service API key, account registration, remote-service dependency, hosted relay, or external notification service.
- No Pro licensing, entitlement lookup/enforcement, plan detection, billing, subscription, or activation of multiple recipients. Future Pro supports multiple recipients; Community runtime remains exactly one.
- No notification channels other than SMTP email: no SMS, push, webhook, Slack, Discord, Telegram, chat, desktop notification, or generic arbitrary provider configuration.
- No fleet management, remote administration, browser dashboard, inbound network listener, callback receiver, remote command, remediation, arbitrary HTTP/network client, shell execution, subprocess, or plugin system.
- No broad alert-policy redesign, new health truth, severity reclassification for product convenience, automatic acknowledgement, maintenance UI, or direct email generation from raw observation evidence.
- No unbounded retries, retry daemon, queue broker/database, parallel fan-out, background worker, delivery receipt claim, mailbox polling, bounce processing, inbound DSN processing, or human-read acknowledgement semantics.
- No plaintext secret in `config.json`, environment, CLI arguments, output, logs, checkpoint, report, Console, test fixture, repository, package, or release artifact. No general-purpose enterprise secret manager, OS keyring integration, OAuth provider flow, or credential migration framework unless a blocker is reported and separately authorized.
- No insecure TLS skip, certificate-verification disablement, opportunistic downgrade presented as secure, provider-specific certificate pinning, or hard-coded commercial SMTP service.
- No implicit email during setup, install, validate, config show, Guardian activation, or ordinary manual observation. Only an explicit test command or eligible Guardian Alert delivery may connect.
- No automatic service enable/start/restart or credential collection in `install.sh` or `qwsg setup` without explicit operator confirmation and approved command semantics.
- No general smart installer, operating-system package inventory framework, package-manager abstraction, automatic dependency installation, infrastructure remediation, or execution of recommended commands. Full install-wide host assessment is reserved as accepted design direction for a separately authorized Smart Setup / Dependency Assessment task.
- No modification, retagging, rebuilding, replacement, republication, or new release derived from QWSG 1.0.0. No automatic public release/tag/publication.
- No Task 047 creation or implementation. A successor may be recommended only; it requires separate Owner authorization.
- No blanket Git staging, cleanup, commit, push, tag, or publication authority. Task preparation and later installation do not authorize implementation or source integration.
- No deletion, relocation, chmod, ownership change, staging, or substantive reading of the unrelated Owner draft beyond the minimum hash/metadata checks required to preserve it.


## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

- Verify the QWSG repository root, UTC date/user, branch `main`, exact HEAD and `origin/main` both `c2a676440c1641f09802fdf4667649f9fe13a0b6`, `0` ahead and `0` behind, canonical origin, empty index, and clean tracked working tree.
- Verify canonical idle: no active prompt; Task 045 is the unique latest complete archive/history pair; `bin/job --check`, Framework, lifecycle, diversion, test-task, and Builder validations pass.
- Record all visible and ignored local paths. Verify the only unrelated visible untracked path is `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`; record its hash, mode, owner, ACL, and size without reading or changing substantive content.
- Verify local and remote annotated `v1.0.0` object and peeled target remain unchanged; verify release-source commit `177535e44b2ce5ed9efd73ab0793ffe6881f0cd6`, LICENSE, retained artifact and sidecar hashes, and published release evidence where locally available.
- Read mandatory governance plus Task 028–030 and Task 045 archive/history; Canonical Alert, Notification, Runtime, Runtime Service, Guardian, Configuration, Scheduler, Security/Privacy, setup, install/uninstall, Console, system map, roadmap, and release documents; inspect every directly relevant implementation and test.
- Record the exact current notification architecture: Alert events/decisions/lifecycle/deduplication/recovery/stale/partial semantics; Notification routes/endpoints/provider bindings/delivery IDs/idempotency/retry/failure/evidence; Runtime inputs/results; Guardian queue persistence and configuration epochs; current inert empty provider/policy wiring; and all observable status paths.
- Record the exact Task 045 configuration and secret-reference extension points, mutable CLI registry, path/store security, precedence, systemd precondition, setup/installer boundaries, and compatibility constraints. Identify every contradiction before choosing schema keys.
- Audit the systemd sandbox and host capability relevant to outbound SMTP, including network namespace/address-family restrictions, DNS, TLS trust-store readability, resource limits, restart behavior, and whether any unit change is actually necessary.
- Run pre-change build, focused Alert/Notification/Runtime/Guardian/Configuration/CLI tests, full tests, bounded-cache race, vet, formatting, whitespace, security/static scans, packaging, systemd verification, Framework, lifecycle, diversion, Builder, and offline/no-network baselines. Stop on unexplained differences.


## Snapshot Requirements

### Mandatory Engineering Safety Rule

Before any Task 046-owned repository modification, create a unique external mode-`0700` snapshot under `/tmp` containing a verified manifest and exact payloads or absence records for every intended source, test, packaging, documentation, and lifecycle target. Record Git branch/HEAD/origin/upstream/index/status, file hashes, modes, owners, ACLs, tool versions, relevant process/service state, configuration/state roots, and immutable release identities. Include a precise collision-aware bounded restore procedure. Verify checksums, archive readability, target inventory, and rollback commands before the first modification. Stop if repository identity, ownership, targets, or restore safety is ambiguous.

- Record but do not copy the Owner draft: preserve only its path, hash, mode, owner, ACL, size, and explicit exclusion.
- Never capture live credentials in a snapshot. Automated tests and acceptance must use generated ephemeral non-production values inside isolated mode-`0700` roots and must record only redacted metadata/evidence.
- Before SMTP, staged-package, or service acceptance, create separate isolated config/state/HOME/install/trust/test-endpoint roots with explicit boundaries, ownership, modes, ports, process IDs, and cleanup rules. Do not target the real home, published artifact, public network, privileged ports, or production SMTP account.
- Preserve the snapshot through completion and record its verified location and restore procedure in Task 046 history. Snapshot creation grants no authority over excluded content.


## Risk Assessment

- **Credential disclosure — critical:** separate private store, strict ownership/modes/path checks, protected input, redaction, bounded errors, repository/package scans, and output/state assertions must prove no secret crosses the boundary.
- **Notification storm or duplicate mail — critical:** Alert lifecycle records plus canonical delivery/idempotency identities, durable queue state, finite retry policy, and restart/replay tests prevent per-cycle or repeated-success delivery.
- **Monitoring interruption — critical:** provider failure remains downstream delivery evidence; Guardian continues scheduling/evaluation and persists bounded failure state without crash, recursive alert, or false monitoring result.
- **TLS downgrade or destination abuse — critical:** validated host/port, required TLS modes, system trust verification, no insecure option, bounded contexts, no arbitrary URL/protocol, and exact network call-site audits.
- **Community entitlement bypass — high:** enforce exactly one enabled recipient at every public/config/runtime activation boundary while retaining only a documented future Pro cardinality policy hook.
- **Configuration compatibility break — high:** extend one Task 045 contract with explicit versioning/migration semantics, preserve precedence/provenance/unknown-field rules, and test existing files/defaults unchanged.
- **Queue/checkpoint corruption — critical:** retain canonical validation, atomic Guardian checkpoint publication, bounded attempts/entries, and rollback; never write secrets or raw SMTP errors into durable state.
- **Ambiguous SMTP outcomes — high:** conservative failure classification and existing indeterminate acknowledgement semantics prevent false delivery claims and unsafe automatic duplication.
- **Privacy leakage in mail — high:** fixed bounded templates from Alert metadata only, no raw evidence/configuration/host inventory, deterministic subject/body tests, and documentation of email exposure.
- **Systemd sandbox regression — high:** prove existing outbound behavior first; preserve every unrelated protection and modify only a demonstrably necessary network directive.
- **CLI automation leaks — critical:** forbid secret arguments/environment and design protected terminal/file-descriptor input; test process listings, human/JSON output, errors, and logs.
- **Test endpoint escape or flaky acceptance — high:** loopback-only controlled SMTP, deterministic clocks/timeouts, isolated trust material, no public Internet, explicit teardown, and bounded resources.
- **Owner/release damage — critical:** hash/exclude Owner draft and immutable release identities before/after; no release, tag, LICENSE, artifact, or publication mutation.
- **Scope growth — high:** one SMTP provider, one recipient, one credential type, existing state engines, and closed CLI surface; defer Pro/QWS/other channels/UI/general secret management.


## Planned Work

1. Verify the exact idle/Git/release/Owner baseline, complete the notification/configuration/systemd/privacy audit, run pre-change gates, and create the verified bounded snapshot.
2. Write the evidence-backed Task 046 architecture decision mapping Community email onto existing Alert Records, Notification policy/queue/provider contracts, Runtime, Guardian checkpoint, Task 045 configuration, and private credential reference. Freeze the default event/severity/recovery matrix and retry/failure taxonomy before transport coding.
3. Extend Configuration Contract 1.0 compatibly with typed email-notification settings and one-recipient Community validation. Update normalization, provenance, identity, strict decoding, defaults, public key registry, setup planning, readiness validation, and old-config compatibility tests. Remove or migrate the inert recipient extension without silent semantic loss.
4. Add a narrow private credential-store package with canonical XDG-derived location and opaque reference resolution, secure creation/update/load, current-user ownership and private mode checks, symlink/special/hard-link/race rejection, bounded data, atomic persistence, and redaction-safe typed errors.
5. Implement a deterministic privacy-safe email renderer and a concrete SMTP `notification.Provider` adapter using bounded contexts, required TLS, system certificate verification, exact configured destination, minimal auth, conservative SMTP response classification, and no provider-specific dependency unless separately justified.
6. Construct the canonical Community Notification Policy from validated effective configuration and wire the SMTP provider registry into Guardian/Runtime. Persist retries in the existing checkpoint queue, preserve configuration-epoch safety, and keep disabled notification on the empty-policy/no-provider path.
7. Extend `setup` and typed `config` operations with email settings, add protected credential provisioning and `notification test`/readiness UX, stable human and JSON results, explicit one-recipient errors, and comprehensive secret non-rendering. Keep the test path separate from Alert and Guardian evidence.
8. Expose bounded local readiness and last-attempt diagnostics through the smallest existing operator projection or dedicated local status contract without changing monitoring truth or disclosing raw provider errors.
9. Add a typed deterministic notification preflight result with the five Owner-approved classifications. Inspect only Task 046-relevant requirements, stop notification test/activation on required or incompatible findings, preserve unknown results explicitly, and provide exact allowlisted remediation commands only where platform support is proven. Never execute remediation.
10. Audit and, only if required, minimally update the systemd user unit. Keep install noninteractive and update next-step guidance without implicit network, credential, dependency-remediation, or service actions.
11. Add focused unit, integration, adversarial, restart-persistence, concurrency/resource, privacy, configuration, CLI, Runtime/Guardian, preflight-classification, and controlled loopback SMTP tests. Use deterministic injected clocks and endpoints; cover TLS/auth success and failure without public Internet.
12. Build a staged package and execute isolated clean-account acceptance through install, setup/configure, notification preflight, protected credential creation, revalidation, controlled test email, Guardian actionable transition/deduplication/retry/recovery, Console/status, restart continuity, offline failure, and uninstall preservation. Distinguish sandbox, host-local, external clean-host, and real-provider evidence truthfully.
13. Update the minimum canonical documentation and Task 046 history, preserve the accepted full Smart Setup / Dependency Assessment direction for a future separately authorized task, run all final gates, verify snapshot/rollback and immutable identities, archive Task 046, and return to canonical idle without creating a successor.


## Rollback Plan

- Before rollback, stop only Task 046-owned isolated test processes/services by exact recorded identity; do not stop unrelated QWSG or Owner processes.
- Verify the Task 046 snapshot manifest, archive readability, target inventory, hashes, and restore instructions. Compare every current target to the snapshot and refuse overwrite if post-snapshot Owner/later-task changes are possible.
- Restore only exact pre-existing Task 046 targets from the verified snapshot and remove only exact Task 046-created paths whose pre-task absence and continued ownership are proven. Never use broad `git reset`, checkout, restore, clean, wildcard deletion, repository-wide extraction, or recursive deletion of an unresolved path.
- Preserve all pre-task configuration, runtime state, release evidence, Task 045 history, ignored Builder inputs/backups/artifacts, and the Owner draft. Do not attempt to roll back or mutate `v1.0.0`.
- Clean only recorded isolated temporary roots after validating their absolute bounded paths and ensuring no real credential or external SMTP dependency exists. Preserve failure evidence needed for diagnosis.
- Rerun focused/full/race/vet/format, configuration/credential/notification/Guardian/systemd/package/security/privacy/offline, Framework/lifecycle/diversion/Builder, Git, permission/ACL, immutable-release, and Owner-draft checks after rollback. Record the result in Task 046 history.


## Deliverables

- A documented Community Email Notification architecture that maps exactly onto Canonical Alert, Notification Delivery, Runtime, Guardian, and Configuration contracts.
- A compatible typed canonical email-notification configuration with deterministic defaults/provenance/validation and enforced exactly-one Community recipient when enabled.
- A private per-user SMTP credential store/resolver with secure atomic mutation and zero secret rendering.
- One bounded standards-compatible SMTP provider with required TLS modes, certificate verification, minimal authentication, deterministic privacy-safe messages, and conservative failure classification.
- Guardian/Runtime activation using existing queue persistence, deduplication, idempotency, retry, and Alert lifecycle semantics while monitoring remains independent of delivery.
- Coherent setup/config/credential/readiness/test CLI behavior with human and JSON output and no secret arguments or output.
- A typed, read-only notification readiness preflight using the five Owner-approved classifications, safe stopping behavior, and platform-gated operator-run remediation recommendations without automatic execution.
- Focused and full automated tests, controlled loopback SMTP integration, staged-package acceptance, restart/offline/deduplication evidence, and truthful clean-host/provider test classification.
- Updated canonical operator, configuration, security/privacy, troubleshooting, packaging, architecture, roadmap/history, and Task 046 lifecycle documentation.
- Verified implementation snapshot and collision-aware rollback evidence; canonical idle completion with no Task 047.


## Verification

- Verify baseline and final repository identity, branch/upstream, exact expected HEAD relationship, canonical origin, index, tracked/untracked/ignored inventory, permissions/ACLs, lifecycle state, and exact Task 046 path ownership. No broad staging or Git mutation is authorized by implementation.
- Run `make build`, focused package tests, `go test ./...`, `go test -race ./...` with bounded writable caches, `go vet ./...`, format checks, `git diff --check`, security/static scans, package/release-script validations applicable to post-1.0 source, and `systemd-analyze verify --user`.
- Run Framework validation and all engineering suites, including Framework, diversion, lifecycle, test-task, and Builder assertions; validate the active prompt/history throughout and simulate archive to canonical idle before actual completion.
- Configuration tests: disabled/default/old Task 045 files; enabled valid one recipient; zero/multiple recipients; malformed/unknown/duplicate fields; invalid host/port/sender/address/security/auth/timeout/retry combinations; strict precedence/provenance/identity; explicit override; migration compatibility; fail before Guardian mutation.
- Preflight tests: exact `satisfied`, `missing_required`, `missing_optional`, `unknown_requires_verification`, and `incompatible` results; required-stop semantics; unknown rather than guessed classification; deterministic supported-platform remediation text; unsupported-platform absence of guessed commands; hostile host/configuration text cannot enter commands; no command execution, package mutation, trust-store mutation, infrastructure mutation, or network action during detection.
- Credential tests: first create/update/read; `0700`/`0600`; current ownership; symlink in each component; special file; hard link; permissive mode; wrong owner where host permits; oversized/empty/control content; unsafe reference; atomic failure/interruption preserving prior value; concurrent replacement; bounded errors; absence from config, args, env, output, state, logs, snapshots, fixtures, binaries, package, and Git diff.
- SMTP tests against a controlled loopback endpoint: implicit TLS and required STARTTLS success; certificate hostname/trust failure; no downgrade; connection refusal; DNS/connect/handshake/command/total timeout; authentication success/failure; recipient/sender rejection; transient and permanent response mapping; partial/indeterminate outcome; bounded payload; no inbound product listener; no public network dependency.
- Notification semantics tests: disabled/no configuration/no network; explicit test clearly labeled and no fake incident state; selected unsuppressed actionable activation; escalation/material change; repeated identical condition; replay after success; retry schedule/backoff/window/exhaustion; restart queue persistence; recovery if included; suppressed/stale/partial/unknown conditions; notification failure not recursively alerting; monitoring cycles continue after every failure class.
- Community boundary tests: every supported CLI/config/setup/runtime path rejects multiple recipients; exactly one enabled recipient works without QWS account/API/subscription/remote service; future multi-recipient representation remains a documented inactive Pro extension with independent global safety bounds and no Community bypass.
- Privacy tests compare exact deterministic subject/body and prove absence of raw inventory, IP/MAC/host identifiers, paths, mounts, full evidence, config values, secret references where sensitive, credentials, and raw errors. Human/JSON/Console/log output must remain bounded and redacted.
- Guardian/systemd tests prove valid startup, invalid enabled configuration stopped by `ExecCondition` without restart churn, delivery failure does not crash or stop scheduling, queue survives restart, resource limits remain, non-root operation remains, and unrelated sandbox directives are unchanged.
- Regression tests cover Task 045 setup/config commands and atomic store; QWSG 1.0 observation/bootstrap/partial evidence; Alert/Notification/Runtime identities; Guardian scheduling/checkpoint/reboot behavior; Console usability; install/uninstall preservation; and no-network local Community operation.
- Staged-package acceptance must use an isolated clean HOME/XDG config/state/install root and controlled SMTP endpoint. Exercise install → setup/configure → protected credential → validate → test email → activate Guardian → actionable transition → deduplication/retry/recovery → inspect status/Console → restart → offline behavior → uninstall preservation. Record whether user systemd was genuinely used.
- Label evidence exactly: unit/integration, controlled SMTP, staged-package, actual external clean-host, and actual real-provider. Do not claim external clean-host or provider delivery unless performed. A real provider is not required for Task 046 acceptance absent separate Owner credentials/authority.
- Verify the snapshot and rollback simulation, Task 046 history completeness, archive/no-active-prompt canonical idle simulation, and exact preservation of Owner draft, LICENSE, `v1.0.0`, release-source commit, published artifact/checksum, and Forgejo release identity.


## Documentation Updates

- Amend `docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md` and `docs/architecture/CANONICAL_NOTIFICATION_DELIVERY.md`; add one focused Community email/credential architecture document only if combining them would obscure permanent boundaries.
- Update `docs/architecture/CONFIGURATION_ACTIVATION.md`, operational Guardian/system map/roadmap/history references, and the minimum Task 046 archive/history records.
- Update the canonical Setup and Configuration guide/reference, Quick Start, Operations, Troubleshooting, Security and Privacy, Upgrade/Rollback/Uninstall, install guidance, and CLI documentation without duplicating contradictory procedures.
- Document exact configuration keys/defaults/precedence, one-recipient Community rule, future Pro multi-recipient path, SMTP TLS/auth support, credential location/permissions/update flow, test command, event/retry/recovery policy, readiness/status diagnostics, privacy payload, offline behavior, and uninstall preservation.
- Document the notification preflight classifications, which checks are Task 046-relevant, safe stop/rerun flow, exact supported remediation recommendations, the strict read-only/no-auto-install boundary, and the accepted future full Smart Setup / Dependency Assessment direction.
- Clearly state that QWS account/API/subscription/remote service and real-provider acceptance are not required; Pro unlimited recipients and QWS-managed delivery are not implemented. Mark Task 046 as post-1.0/future-1.1.0 development without changing public release identity.
- Record all decisions, validation evidence, acceptance classification, snapshot/rollback location, limitations, immutable-release verification, Owner-draft preservation, and canonical idle completion in Task 046 history.


## Completion Criteria

Task 046 is complete only when the exact idle/release baseline and existing notification machinery are verified; one compatible canonical configuration activates optional Community SMTP for exactly one administrator recipient; public configuration contains only an opaque credential reference; the private credential store passes ownership/mode/path/symlink/special/hard-link/bounded-read/atomic-write and non-disclosure gates; one TLS-verifying bounded SMTP provider operates only for explicit tests or due canonical requests; privacy-safe deterministic messages derive only from immutable Alert Records; existing Alert and Notification identities prevent storms and persist retry continuity across Guardian restart; selected actionable/recovery semantics and all stale/partial/unknown/suppressed cases are documented and tested; failed delivery is observable but never stops monitoring, recursively alerts, or makes remote service mandatory; setup/config/credential/readiness/test UX works in human and automation-safe modes; the read-only notification preflight deterministically emits all five Owner-approved classifications, stops safely for required/incompatible findings, reports uncertainty explicitly, and provides only proven platform-gated operator-run remediation without executing it; Community multiple-recipient bypasses fail; Pro-compatible collection semantics are documented but inactive; systemd protections remain or are minimally justified; staged-package controlled-SMTP acceptance and all focused/full/race/vet/format/security/privacy/offline/package/systemd/framework/lifecycle/rollback gates pass; evidence claims distinguish controlled, staged, external-host, and real-provider testing; Task 046 history is complete and archived; no Task 047 exists; canonical idle is restored; and the Owner draft plus immutable QWSG 1.0.0 tag/source/artifact/LICENSE/release remain unchanged. Any credential exposure, TLS downgrade, duplicate storm, Community cardinality bypass, monitoring interruption, automatic dependency/remediation execution, guessed preflight result, unsafe remediation text, false delivery claim, unsafe path/write, queue loss, privacy leak, unexplained regression, release mutation, or unperformed mandatory acceptance is a blocker and Task 046 must not be marked complete.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-12 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
