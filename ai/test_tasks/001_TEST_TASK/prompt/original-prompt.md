# Current Engineering Task 015: Forgejo Server Installation v1

## Task Metadata

- Task ID: `015`
- Task slug: `forgejo-installation`
- Status: `approved`
- Date opened: `2026-07-22` UTC
- Human authority: Attila — Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Forgejo Server Installation v1


## Objective

Install and configure a production-ready Forgejo Git server on the existing HestiaCP VPS using the subdomain:

git.quantumwizard.hu

The installation shall become the official source control platform for all Quantum Wizard Studio projects.

Existing Environment

Operating System:

Ubuntu 24.04 LTS

Control Panel:

HestiaCP 1.9.7

Web stack:

Nginx + Apache

Database:

MariaDB 11.4

Target domain:

git.quantumwizard.hu

The domain already exists and serves the default Hestia page.

Installation Requirements

Implement a native production installation.

Do NOT use:

Docker
Snap
SQLite

Use:

native Forgejo binary
dedicated system user
MariaDB
systemd service
reverse proxy through Hestia Nginx
Let's Encrypt HTTPS
Security

Configure:

registration disabled
private repositories by default
only administrator may create users
SSH enabled
HTTPS enabled
Git LFS enabled
Packages enabled
Actions enabled

The Forgejo process must not be exposed directly to the Internet.

Only Nginx may expose HTTPS.

Organization

Create the default organization:

Quantum Wizard Studio

The organization will host future repositories such as:

QWSG
QUWIP
QuantumWizard.hu
AlexTamas.hu
Infrastructure
AI
Initial Repository

Create:

QWSG

Private repository.

Git Migration

Configure the local repository:

origin

Push:

main
tags

including

v0.1.0


## Scope

Install and configure a production-ready, self-hosted Forgejo Git service for Quantum Wizard Studio on the existing HestiaCP VPS.

The scope includes:

- Create a verified pre-change system snapshot and documented rollback procedure.
- Install Forgejo as a native Linux binary, without Docker, Snap, or SQLite.
- Create a dedicated, non-login Forgejo system user and secure filesystem layout.
- Create a dedicated MariaDB database and database user with minimum required privileges.
- Configure Forgejo as a systemd-managed service.
- Bind the Forgejo web service exclusively to 127.0.0.1 on an available internal port.
- Configure git.quantumwizard.hu through a HestiaCP-compatible Nginx reverse proxy without directly editing generated configuration files that Hestia may overwrite.
- Enable and verify HTTPS using the existing HestiaCP and Let's Encrypt infrastructure.
- Configure the canonical external URL as https://git.quantumwizard.hu/.
- Configure secure session, cookie, proxy-header, and trusted-proxy behavior.
- Disable public user registration.
- Set newly created repositories to private by default.
- Restrict user creation and administration to the Forgejo administrator.
- Configure Git operations over HTTPS.
- Assess SSH Git access against the existing server SSH and HestiaCP configuration, then implement it only through a documented, non-conflicting design. Do not replace or disrupt the server's existing SSH service.
- Enable Git LFS and verify its storage path and basic operation.
- Enable the package registry where safely supported.
- Prepare Forgejo Actions support, but do not deploy an Actions runner unless explicitly approved by the Project Owner.
- Create the initial administrator account through a secure method that does not expose or commit credentials.
- Create the organization with the display name “Quantum Wizard Studio” and a URL-safe organization identifier selected according to Forgejo naming rules.
- Create a private repository named QWSG under the organization.
- Configure the existing local QWSG repository with the new Forgejo repository as its origin.
- Push the main branch and the approved v0.1.0 annotated tag.
- Verify clone, fetch, pull, push, branch, and tag operations.
- Verify service restart and persistence after reboot without performing an uncontrolled reboot.
- Document the installed version, filesystem layout, service configuration, database configuration, reverse proxy configuration, SSH design, backup requirements, upgrade procedure, validation results, and rollback procedure.
- Record all implementation evidence in the official QWSG engineering history and delivery report.

The implementation must preserve the existing HestiaCP, Nginx, Apache, MariaDB, mail, SSH, and hosted-domain services.


## Out of Scope

The following items are explicitly out of scope:

- Installing Forgejo with Docker, Podman, Snap, Kubernetes, or another container platform.
- Using SQLite as the production database.
- Replacing or reconfiguring the existing HestiaCP installation.
- Replacing the existing Nginx, Apache, MariaDB, mail, firewall, or SSH architecture.
- Editing Hestia-generated configuration files in a way that will be overwritten by a domain rebuild.
- Exposing the Forgejo application port directly to the public Internet.
- Opening new public firewall ports without explicit Project Owner approval.
- Replacing the server’s existing SSH daemon or changing its primary SSH port.
- Enabling SSH Git access through a design that conflicts with existing server users, HestiaCP, SFTP restrictions, or administrative access.
- Installing or activating a Forgejo Actions runner.
- Allowing untrusted workflow execution on the server.
- Implementing CI/CD pipelines for QWSG or any other project.
- Migrating QUWIP, QuantumWizard.hu, AlexTamas.hu, AI, Infrastructure, or other repositories.
- Creating additional repositories beyond the initial private QWSG repository.
- Making the QWSG repository or the Forgejo instance publicly writable.
- Enabling public user registration.
- Creating non-administrator user accounts.
- Configuring external OAuth, LDAP, Active Directory, SSO, or social login providers.
- Configuring SMTP, email notifications, or password-reset delivery unless required for safe initial administration and separately approved.
- Installing object storage, S3-compatible storage, or an external artifact store.
- Migrating repository data to another server.
- Configuring high availability, clustering, replication, load balancing, or database failover.
- Implementing a full disaster-recovery platform.
- Automating Forgejo upgrades.
- Customizing the Forgejo theme, branding, logo, templates, or frontend appearance.
- Developing Forgejo plugins, extensions, webhooks, bots, or API integrations.
- Publishing a public release page or release binaries.
- Creating a new QWSG engineering feature beyond what is strictly required to document and validate this infrastructure task.
- Modifying the QWSG runtime, collector framework, inventory model, or application behavior.
- Starting or preparing the next QWSG product-development task.
- Creating artificial historical commits or rewriting the existing Git history.
- Force-pushing, rebasing, squashing, or deleting existing branches or tags.
- Deleting the pre-installation snapshot or any rollback evidence.
- Performing an uncontrolled server reboot.
- Removing or disabling existing hosted domains or production services.


## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`

## Starting State Verification

The implementation starts from the following verified baseline:

- QWSG Engineering Lifecycle v1 is fully operational.
- Foundation milestone commit completed:
  b4316050436bf8be4062f0e1d4ba7c371c334223
- Release candidate:
  v0.1.0 (approved, not yet published remotely).
- No Git remote repository is currently configured.
- The local QWSG repository exists only on the development VPS.
- There is no existing Forgejo, Gitea, GitLab, or other self-hosted Git service installed.
- No service is currently listening on TCP port 3000.
- The subdomain https://git.quantumwizard.hu exists and currently serves the default HestiaCP placeholder page.
- Target platform:
  - Ubuntu 24.04.4 LTS
  - HestiaCP 1.9.7
  - Nginx + Apache
  - MariaDB 11.4
  - x86_64
  - KVM virtual machine
- Approximately 68 GB of free disk space is available.
- Approximately 9.7 GB RAM is available.
- Existing production services (websites, mail, SSH, databases and hosted domains) are operational and must remain unaffected.
- The QWSG engineering backup policy has already been implemented.
- Engineering snapshot and rollback procedures are available.
- Canonical idle lifecycle state is valid:
  - no active CURRENT_TASK
  - previous task archived
  - validators pass.
- No Task 016 or later has been created.
- The objective of this task is infrastructure deployment only; no Guardian runtime functionality shall be modified.


## Snapshot Requirements

Before making any system modification, create a complete engineering snapshot of the current server state.

The snapshot shall include, at minimum:

- Current Git repository state.
- Current HestiaCP web configuration related to git.quantumwizard.hu.
- Existing Nginx and Apache configuration relevant to the target domain.
- Existing systemd service inventory.
- Existing MariaDB databases and users (metadata only; never export passwords or secrets).
- Current firewall configuration.
- Current DNS-related configuration where applicable.
- Current filesystem state of all directories that will be modified.
- Existing SSL certificates and certificate metadata relevant to the target domain.
- Current package and binary versions relevant to the installation.

The snapshot shall:

- generate a manifest of captured artifacts;
- calculate SHA-256 hashes for snapshot archives where applicable;
- include a documented restore procedure;
- remain outside the Git repository unless explicitly permitted by the Engineering Backup Policy;
- contain no plaintext credentials, private keys, API keys, access tokens, or other secrets.

No installation or configuration changes may begin until the snapshot has completed successfully and its integrity has been verified.


## Risk Assessment

This task modifies production server infrastructure and therefore carries elevated operational risk.

The primary risks are:

- Loss of access to the Git service due to incorrect reverse proxy configuration.
- Service startup failure caused by invalid Forgejo configuration.
- Database connection failures resulting from incorrect MariaDB configuration.
- Permission or ownership errors affecting repository storage.
- HTTPS misconfiguration causing unavailable or insecure access.
- SSH Git configuration conflicting with the existing SSH service or HestiaCP administration.
- HestiaCP configuration overwrite during future domain rebuilds.
- Accidental exposure of the Forgejo application port to the public Internet.
- Accidental exposure of repositories through incorrect visibility settings.
- Leakage of credentials, secrets, SSH keys, or administrator information.
- Interruption of existing hosted websites or production services.
- Introduction of configuration changes that cannot be safely rolled back.

Risk mitigation requirements:

- Create and verify a complete engineering snapshot before any modification.
- Apply changes incrementally and validate each stage before continuing.
- Preserve all existing production services.
- Never replace or modify existing SSH infrastructure without explicit justification and verification.
- Keep the Forgejo application bound to localhost and expose it only through the approved reverse proxy.
- Validate every configuration before restarting affected services.
- Maintain a documented rollback path for every modified component.
- Do not proceed past a failed validation until the issue has been resolved.
- Record all implementation decisions and deviations in the engineering history and final delivery report.


## Planned Work

The implementation shall be performed in the following controlled phases:

### Phase 1 — Engineering Preparation

- Verify the starting state.
- Create and verify the engineering snapshot.
- Confirm rollback capability.
- Verify required software versions.
- Verify port availability.
- Verify target domain readiness.

### Phase 2 — System Preparation

- Create the dedicated Forgejo system user and group.
- Create the required directory structure.
- Apply secure ownership and filesystem permissions.
- Prepare runtime and data directories.

### Phase 3 — Database

- Create the dedicated MariaDB database.
- Create the dedicated database user.
- Apply minimum required privileges.
- Verify database connectivity.

### Phase 4 — Forgejo Installation

- Download the official Forgejo release.
- Verify integrity where possible.
- Install the native binary.
- Create the initial configuration.
- Configure filesystem paths.
- Configure logging.
- Configure security options.

### Phase 5 — Service Configuration

- Create the systemd service.
- Configure automatic startup.
- Verify startup and shutdown.
- Validate service health.

### Phase 6 — Reverse Proxy

- Configure Hestia-compatible reverse proxy integration.
- Configure HTTPS.
- Configure trusted proxy headers.
- Validate external access.
- Ensure the internal Forgejo service remains bound to localhost only.

### Phase 7 — Security Hardening

- Disable public registration.
- Configure private repositories as the default.
- Restrict administrative functions.
- Configure secure session and cookie settings.
- Configure Git LFS.
- Enable Packages.
- Prepare Actions support without deploying a runner.
- Review all security-related configuration.

### Phase 8 — Initial Platform Configuration

- Create the administrator account.
- Create the "Quantum Wizard Studio" organization.
- Create the private "QWSG" repository.
- Configure repository defaults.

### Phase 9 — Repository Migration

- Configure the local Git repository remote.
- Push the main branch.
- Push the approved v0.1.0 annotated tag.
- Verify clone, fetch, pull, push, branch and tag operations over HTTPS.
- Evaluate and document SSH Git access if implemented.

### Phase 10 — Validation

- Verify all services.
- Verify HTTPS operation.
- Verify repository management.
- Verify Git LFS.
- Verify service restart.
- Verify persistence after reboot simulation.
- Execute all required engineering validation procedures.

### Phase 11 — Documentation

- Record every implementation step.
- Document all modified files.
- Document service configuration.
- Document database configuration.
- Document reverse proxy configuration.
- Document backup considerations.
- Document upgrade recommendations.
- Produce the final Engineering Delivery Report.


## Rollback Plan

Rollback shall be possible at every implementation phase.

If any validation fails, the implementation shall stop immediately and either:

- correct the issue before continuing; or
- restore the affected component to its verified pre-change state.

Rollback shall include, where applicable:

- Removal of the Forgejo binary.
- Removal of the Forgejo systemd service.
- Restoration of the previous reverse proxy configuration.
- Removal of all newly created runtime directories.
- Restoration of modified configuration files from the engineering snapshot.
- Removal of the dedicated Forgejo database and database user if they were created exclusively for this task.
- Restoration of filesystem ownership and permissions where modified.
- Removal of newly created repository configuration from the local QWSG repository if migration has not completed successfully.
- Verification that existing HestiaCP websites, mail services, SSH access, databases, and hosted domains remain fully operational.

Rollback requirements:

- No rollback step may delete unrelated production data.
- Rollback shall preserve all existing Git history.
- Rollback shall never modify repositories unrelated to this task.
- Every rollback action shall be recorded in the engineering history.
- After rollback, the server shall be validated to confirm that it matches the verified starting state and that no residual Forgejo components remain unless explicitly approved by the Project Owner.


## Deliverables

The completed implementation shall deliver all of the following:

### Infrastructure

- Production-ready native Forgejo installation.
- Dedicated Forgejo system user.
- Dedicated MariaDB database.
- Secure filesystem layout.
- Functional systemd service.
- Hestia-compatible reverse proxy configuration.
- Fully operational HTTPS access.
- Secure localhost-only application binding.

### Platform Configuration

- Administrator account.
- "Quantum Wizard Studio" organization.
- Private "QWSG" repository.
- Git LFS enabled.
- Package Registry enabled where supported.
- Actions prepared but not deployed.
- Secure default platform configuration.

### Repository Integration

- Local QWSG repository connected to the new remote.
- Main branch successfully published.
- Approved v0.1.0 annotated tag published.
- Successful verification of clone, fetch, pull, push, branch, and tag operations.

### Validation

- Service health verification.
- HTTPS verification.
- Reverse proxy verification.
- Database connectivity verification.
- Repository operation verification.
- Git LFS verification.
- Restart verification.
- Persistence verification.
- Security configuration verification.

### Documentation

- Engineering snapshot manifest.
- Rollback documentation.
- Installation record.
- Configuration inventory.
- Service inventory.
- Database inventory.
- Reverse proxy documentation.
- Security configuration summary.
- Backup recommendations.
- Upgrade recommendations.
- Validation results.
- Final Engineering Delivery Report.
- Updated QWSG Engineering History.


## Verification

The implementation shall not be considered complete until all verification steps have passed successfully.

The minimum required verification includes:

### Engineering

- Starting state documented.
- Engineering snapshot successfully created.
- Rollback procedure verified.
- No unresolved implementation errors.

### Installation

- Forgejo service installed.
- Systemd service enabled.
- Service starts without errors.
- Service stops cleanly.
- Service restarts successfully.

### Network

- Internal Forgejo service reachable only through localhost.
- Reverse proxy functioning correctly.
- HTTPS accessible.
- HTTP redirects correctly to HTTPS.
- No unintended public ports exposed.

### Database

- Database connection successful.
- Repository metadata stored correctly.
- No database permission errors.

### Platform

- Administrator authentication successful.
- Public registration disabled.
- Private repositories created by default.
- Organization successfully created.
- QWSG repository successfully created.

### Git Operations

- Repository clone succeeds.
- Fetch succeeds.
- Pull succeeds.
- Push succeeds.
- Branch creation succeeds.
- Tag creation and retrieval succeed.
- v0.1.0 annotated tag verified.

### Features

- Git LFS operational.
- Package Registry operational if enabled.
- Actions configuration present but no runner active.

### Security

- Administrator-only management verified.
- Repository visibility verified.
- Session security verified.
- Cookie security verified.
- Trusted proxy configuration verified.
- No credentials exposed in logs or configuration.

### Service Reliability

- Service restart verified.
- Configuration reload verified.
- Persistence after simulated reboot verified.
- Existing HestiaCP services remain operational.

### Documentation

- Engineering History updated.
- Delivery Report completed.
- Configuration inventory completed.
- Validation evidence recorded.
- Rollback documentation completed.

Implementation is complete only after every verification item above has passed successfully.


## Documentation Updates

The implementation shall produce complete engineering documentation in accordance with the QWSG Engineering Lifecycle.

The following documentation shall be updated or created as part of this task:

### Engineering Records

- Engineering History
- Delivery Report
- Implementation Record
- Validation Record
- Rollback Record (if used)

### Technical Documentation

- Installed Forgejo version.
- Installation method.
- Filesystem layout.
- Directory ownership and permissions.
- System user configuration.
- MariaDB configuration.
- Reverse proxy configuration.
- HTTPS configuration.
- systemd service configuration.
- Git LFS configuration.
- Package Registry configuration.
- Actions preparation status.
- Repository configuration.
- Organization configuration.

### Operational Documentation

- Backup requirements.
- Restore procedure.
- Upgrade recommendations.
- Routine maintenance recommendations.
- Log file locations.
- Service management commands.
- Health verification procedure.
- Troubleshooting notes for known issues encountered during implementation.

### Validation Evidence

- Service status verification.
- HTTPS verification.
- Git operation verification.
- Repository verification.
- Security verification.
- Configuration validation results.
- Screenshots only where they provide unique implementation evidence.

### Engineering Requirements

- All documentation shall accurately reflect the final implemented system.
- Documentation shall describe the final state, not intermediate work.
- Secrets, passwords, API keys, tokens, private keys, and certificates shall never be recorded.
- Every deviation from the original implementation plan shall be documented with its technical justification.
- All documentation shall be sufficiently complete to allow a qualified engineer to understand, maintain, upgrade, or reproduce the installation without relying on undocumented knowledge.


## Completion Criteria

This task shall be considered complete only when all of the following conditions have been satisfied:

### Engineering

- All planned implementation phases have been completed or explicitly documented as intentionally omitted.
- No unresolved critical or high-severity issues remain.
- No unresolved validation failures remain.
- No temporary workarounds remain undocumented.

### Infrastructure

- Forgejo is fully operational as a native production service.
- All required supporting components are functioning correctly.
- The service starts automatically after system boot.
- Existing production services continue to operate without regression.

### Platform

- The administrator account is operational.
- The "Quantum Wizard Studio" organization exists.
- The private "QWSG" repository exists and is accessible.
- Public registration is disabled.
- Security configuration has been verified.

### Repository

- The local QWSG repository is connected to the Forgejo remote.
- The main branch has been successfully published.
- The approved v0.1.0 annotated tag has been successfully published.
- Standard Git operations have been successfully verified.

### Validation

- Every required verification step has passed.
- No outstanding warnings require immediate action.
- All implementation evidence has been recorded.

### Documentation

- Engineering History updated.
- Delivery Report completed.
- Configuration documentation completed.
- Validation documentation completed.
- Operational documentation completed.
- Backup and recovery documentation completed.

### Final Acceptance

The task is complete only after:

- all verification requirements have passed;
- all required documentation has been completed;
- the final implementation accurately reflects the documented architecture;
- the Project Owner has reviewed the Delivery Report and formally accepted the implementation.

Only after successful acceptance may this Engineering Task be archived and the next engineering task become active.


## Owner Approval Requirements

Approved by Attila — Project Owner through the Engineering Task Builder on 2026-07-22 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
