# Core Alpha Security and Permission Model

## Scope and security objective

This model governs Slice 1 read-only discovery. The objective is to produce useful inventory while preventing privilege expansion, host mutation, secret exposure, unbounded execution, and false completeness. Console network security, Installer privileges, remediation, update signing, and production retention remain separate gates.

## Assets and threats

Protected assets are host stability, credentials and secrets, private operational metadata, integrity of inventory, local QWSG state, audit attribution, and operator trust. Threats include malicious filenames and symlinks, poisoned command output, hostile environment variables, terminal/markup injection, oversized output, hanging processes, TOCTOU replacement, compromised optional tools, sensitive process arguments, path traversal, and a caller attempting to convert discovery into arbitrary execution.

## Permission model

- Default and only Slice 1 runtime privilege is the invoking ordinary user or a future dedicated unprivileged QWSG identity.
- The Agent initiates no `sudo`, `su`, polkit, capability grant, setuid helper, privileged socket, or root daemon request.
- Privileged facts are optional and separately marked `permission_denied`; unrelated collection continues.
- Files and directories created by QWSG, if persistence is authorized, are QWSG-owned, non-world-readable, non-executable data with least necessary group access.
- The Console never inherits filesystem or process privilege by virtue of presentation access.

## Command execution policy

Use direct system APIs or bounded virtual files before subprocesses. Every allowed command is registered by immutable executable identity and fixed argument grammar. No shell, shell string, `eval`, glob expansion, PATH lookup, user-provided executable, arbitrary environment, stdin, or working directory is permitted. The runtime sets a minimal known environment, fixed locale for parsing, safe umask, finite timeout, process-group cancellation, stdout/stderr byte limits, and descriptor closure. Exit status and parse validation are mandatory; raw output is transient.

## Filesystem and symlink safety

Collectors access only documented sources. They open paths relative to validated roots where possible, reject traversal, distinguish symlinks, avoid following mutable links for sensitive sources, bound reads, verify file type and metadata, and never recursively inspect arbitrary trees. Storage inventory uses mount metadata and filesystem statistics rather than content traversal. No collector writes access-time-sensitive monitored content intentionally.

## Data minimization and redaction

Secrets, tokens, keys, password hashes, cookies, environment values, process arguments, file contents, application configuration, mail data, database data, and network payloads are prohibited. Hostnames, IP addresses, MAC addresses, mount sources, user names, paths, and service names are sensitivity-tagged and redacted or pseudonymized by default according to the data model. Redaction occurs before logging, persistence, and presentation. Unknown evidence is never replaced with a guessed value.

## Network boundary

Slice 1 makes no outbound network connection and exposes no listener. Network inventory comes from local metadata only. Console transport, TLS, authentication, authorization, recovery, and exposure are blocked by gate `AG-004`.

## Failure behavior

Timeout, permission denial, unavailable dependency, unsupported context, invalid evidence, output limit, cancellation, and internal failure are distinct. One failure does not broaden access, trigger a retry under higher privilege, or discard safe completed categories. Retries are bounded within the original authority and deadline. Schema or redaction failure prevents persistence and marks the request failed or partial.

## Logging and audit

Logs use structured fields, safe message identifiers, request and collector correlation, and redaction counts. Untrusted text cannot control log structure, terminal state, HTML, or audit identity. Read-only scheduled discovery is operational history; operator invocation is attributable. Any configuration change or future mutation requires a separate audit contract.

## Security verification

Tests must prove non-root execution, no host mutation, no shell invocation, fixed executable resolution, environment sanitation, timeout/process cleanup, output limits, path traversal rejection, symlink race resistance appropriate to the chosen runtime, terminal/markup escaping, seeded-secret exclusion, schema rejection, permission-denied truthfulness, and failure isolation. Platform integration tests run in disposable fixtures or controlled test hosts, never by mutating production.

## Security conclusions

Slice 1 can be implemented safely before Console, e-mail, retention, and update-signing decisions because it has no network listener, outbound transmission, lifecycle mutation, or privilege escalation. It cannot be described as production-supported until the platform matrix is approved and applicable implementation/security tests pass.
