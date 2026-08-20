# QWSG 1.0 Supported Platform

Supported for QWSG 1.0: Ubuntu 24.04 LTS, Linux amd64/x86-64, systemd 255 or later, glibc-compatible userspace, an ordinary non-root user with a working systemd user manager, and a local filesystem providing atomic rename, advisory `flock`, owner/mode checks, and private `0700`/`0600` storage.

Other Linux distributions are experimental. Other CPU architectures, containers without a compatible user manager, non-systemd systems, remote/fleet operation, and privileged/root Guardian execution are unsupported.

`qwsg install --check` reports uncertainty as `unknown_requires_verification`
and proven support violations as `incompatible`. Unsupported distributions and
managed/control-panel stacks receive no speculative package command.

Task 043 completed the Owner-run physical reboot journey on a freshly reinstalled disposable supported host: explicit lingering, boot-before-login service recovery, recurring post-reboot state, Console freshness, controlled restart and uninstall all passed. Repository tests do not replace that retained real-host evidence.
