# Core Alpha Slice 1 Development

## Status

Slice 1 is an internal pre-alpha implementation, not a supported release. The repository pins Go `1.26` in `go.mod`; development was verified with Go `1.26.5` on Linux/amd64. It uses only the Go standard library and needs no network module resolution.

## Layout and commands

- `cmd/qwsg`: CLI and exit behavior.
- `internal/app`: deterministic coordination and aggregation.
- `internal/collector`: read-only Linux collectors.
- `internal/inventory`: versioned envelope and validation.
- `internal/runner`: allowlisted, shell-free, bounded subprocess execution.

Run `make fmt-check`, `make vet`, `make test`, and `make build`. The ignored binary is `build/qwsg`. Commands are `qwsg help`, `qwsg version`, and `qwsg inventory`; unknown commands and options fail without interaction.

## Output contract

`inventory` writes one schema-versioned JSON envelope to stdout and diagnostics to stderr. Exit `0` means every category is available, `2` means a structurally valid partial inventory, and `1` means fatal failure with no valid inventory. Category states distinguish `available`, `unavailable`, `unsupported`, `permission_denied`, `timeout`, `error`, and `cancelled`.

The enabled categories are operating system, kernel, privacy-safe host identity, CPU, memory/swap, block storage, filesystems/mounts, network interfaces, virtualization, running systemd services where accessible, allowlisted Go runtime detection, and collector capabilities. The same Registry Results now also produce the authoritative `canonical_inventory` representation. Service names, hostnames, addresses, hardware addresses, interface names, device names, and mount paths are omitted, hashed, or redacted by default.

## Security and limitations

Collection is local, one-shot, non-root, read-only, and makes no network connection. Subprocesses use fixed absolute executables and arguments, a controlled environment, closed stdin, a two-second timeout, and one-MiB output bounds. No shell, daemon, listener, persistence, installer, privilege escalation, remediation, package inventory, or arbitrary path/command input exists.

Exact supported distributions and architectures remain gate `AG-001`. Service discovery is unavailable outside safely accessible systemd contexts. Component discovery is intentionally limited to an allowlisted Go version command. Container, namespace, filesystem, and permission coverage requires Task 010 hardening.
