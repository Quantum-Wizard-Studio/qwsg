# Forgejo Release Distribution Contract

## Status and boundary

QWSG release artifacts are distributed as immutable, version-specific Forgejo
Release assets. This document defines the contract but does not claim that a
particular future asset currently exists. Tag creation, Forgejo Release
creation, asset upload, and publication remain separate Owner-reserved actions.

Anonymous read-only checks on 2026-08-24 did not expose the QWSG repository or
a Release endpoint. Therefore no operational
`git.quantumwizard.hu` asset URL is asserted by Task 058. The publication task
must verify the expanded URLs using unauthenticated `wget` and `curl` after the
tag, Release, and assets exist.

## Immutable asset model

For version `VERSION`, publication creates:

- immutable tag `vVERSION`;
- one Forgejo Release associated with that exact tag;
- archive `qwsg-VERSION-linux-amd64.tar.gz`; and
- sidecar `qwsg-VERSION-linux-amd64.tar.gz.sha256`.

Forgejo's standard versioned Release-asset route is:

```text
FORGEJO_BASE/OWNER/REPOSITORY/releases/download/TAG/ASSET
```

For QWSG the parameters are expected to be:

```text
FORGEJO_BASE=https://git.quantumwizard.hu
OWNER=Quantum_Wizard_Studio
REPOSITORY=qwsg
TAG=vVERSION
```

This is a URL contract/template, not evidence that an unpublished asset is
available. Do not add a mutable unverified `latest` URL to an installer or user
guide.

## Command-line workflow

After publication has verified the expanded `RELEASE_BASE`, users can run:

```sh
version=VERSION
release_base="FORGEJO_BASE/OWNER/REPOSITORY/releases/download/v${version}"
archive="qwsg-${version}-linux-amd64.tar.gz"

wget "${release_base}/${archive}"
wget "${release_base}/${archive}.sha256"
sha256sum -c "${archive}.sha256"
```

Or with curl:

```sh
version=VERSION
release_base="FORGEJO_BASE/OWNER/REPOSITORY/releases/download/v${version}"
archive="qwsg-${version}-linux-amd64.tar.gz"

curl -fLO "${release_base}/${archive}"
curl -fLO "${release_base}/${archive}.sha256"
sha256sum -c "${archive}.sha256"
```

`curl -f` is mandatory so an HTTP error page is not mistaken for an artifact.
The sidecar filename must reference only the archive basename so verification
works in the download directory. SHA-256 detects corruption after obtaining the
sidecar from the intended Release; publisher authentication or signing remains
a separate release-security decision.

## Publication verification

Before announcing a Release, its separately authorized publication workflow
must verify:

1. tag, Release, archive, and sidecar all identify the same version and source;
2. both assets are immutable regular downloads with expected sizes and hashes;
3. unauthenticated `wget` downloads both exact URLs;
4. unauthenticated `curl -fLO` downloads both exact URLs;
5. `sha256sum -c` passes in clean directories for both clients;
6. redirects, content disposition, TLS, and filenames behave predictably; and
7. no credentials or private repository access are required for public users.

Failure means the distribution is not ready; it does not justify inventing a
different URL or weakening access/security settings inside an acceptance run.

## Future Smart Installer contract

A future Smart Installer should consume an explicitly selected immutable
version, construct or receive the verified version-specific Release base, fetch
only the archive and sidecar, fail on HTTP errors, enforce size/type limits,
verify SHA-256 before extraction, and apply the existing archive path/type and
manifest checks. Version discovery, signing, trust roots, mirrors, and mutable
channels require separate designs. Manual workstation-mediated copying is not
part of the normal published-release experience.
