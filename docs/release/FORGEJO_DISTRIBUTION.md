# Forgejo Release Distribution Contract

## Status and boundary

QWSG release artifacts are distributed as immutable, version-specific Forgejo
Release assets. Tag creation, Forgejo Release creation, asset upload, and
publication remain separate Owner-reserved actions.

QWSG `1.1.0` is the first release verified under this contract. On 2026-08-26,
anonymous repository and Release access, independent clean-environment `wget`
and `curl -fLO` downloads, and sidecar verification all passed for:

```text
https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v1.1.0/qwsg-1.1.0-linux-amd64.tar.gz
https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v1.1.0/qwsg-1.1.0-linux-amd64.tar.gz.sha256
```

The archive size is `2951638` bytes, its SHA-256 is
`10a39d96b93b72a3f4799a76d769bc264afd6845a32a1ecc5531b062d6f42349`,
and the sidecar SHA-256 is
`b9414bba5a6d9bc100f7c391c11867ceda6c2139002272a0f46fbf55dc9d3cc1`.

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
