# QWSG 1.4.0 release candidate

Task 093 prepares an **unsigned, unpublished release candidate**. The semantic
version is 1.4.0; this is not a production-signed, published or accepted release.
Community C9 remains MISSING. No new product scope is introduced here.

## Version decision

VERSION was 1.3.1. Historical 1.3.0 introduced authenticated discovery and
notifications; 1.3.1 was its narrowly corrective patch. Accepted Tasks 082–092
add authenticated exact-source migration capabilities, stronger offline archive
authority, policy/capability and bounded orchestration foundations, stable service
identity and durable recovery. These include additive functionality beyond a
patch; existing supported configuration/state schemas remain preserved. The
smallest next minor identity is 1.4.0, within major 1.

The existing stable index rejects prerelease identities, and the production
update path selects stable only. Using 1.4.0 as a frozen candidate identity
preserves that contract without introducing a preview update mechanism. The
candidate state is recorded separately from its SemVer; no tag or production
publication is authorized. The historical 1.0.0 RC convention in the initial
release policy does not renumber the later 1.1/1.2/1.3 release lineage.

## Included accepted changes

- Task 082: release-index/2, exact source migration/1 declarations and the local
  preserve-package-v1 capability; signed archive companions and helper revalidation.
- Tasks 083–087: policy/capability foundations, bounded orchestration, Guardian
  handoff, common mutation exclusion and verified recovery. Production still
  resolves Community; no production Pro entitlement or automatic install claim.
- Task 088: frozen product maturity scope, separate from semantic release numbering.
- Task 089: stable privacy-preserving service identity.
- Task 090: recognized local evidence interruption recovery and integrity checks.
- Task 091: durable explicit update/rollback recovery and truthful outcome evidence.
- Task 092: operator readiness, optional SMTP and recovery guidance reconciliation.

## Bootstrap and compatibility

Official 1.3.1 predates Task 082 and cannot consume release-index/2. Its native
update command is not the supported bootstrap coordinator. After separate signing
and publication authorization, independently authenticate the signed index using
the established release-index verifier and bundled production public key; bind
archive name, length and SHA-256 to that index before executing extracted code.
Use the verified 1.4.0 archive binary as the explicit bootstrap coordinator with
its archive, checksum and full signed index companion. It classifies the actual
installed 1.3.1 package, checks exact signed compatibility and independently
reauthenticates in the privileged helper before mutation.

The sole declaration is 1.3.1 -> 1.4.0, linux-amd64, preserve-package-v1:
Configuration 1.0, Guardian 1.0, Scheduler 1.0, Operator State 1.0–1.2. No schema
conversion or new migration route is added. Unsupported source, capability,
schema, platform, signature or package identity fails closed. Do not run the
archive install.sh over an existing installation to bypass these checks.

The exact frozen artifact and Dell1 handoff are recorded after the source commit
under release/candidates/1.4.0/. Build only that source commit with the recorded
full BUILD_COMMIT, SOURCE_DATE_EPOCH and toolchain. Closure metadata is outside
the package input boundary. Never rebuild using a later documentation commit and
label it with the frozen source commit.

## Remaining release gate

C9 must prove clean installation, the actual candidate-coordinated bootstrap from
immutable official 1.3.1, preservation, explicit rollback and Guardian recovery on
a separately authorized real supported system after signing/publication. Optional
notification acceptance remains truthful and separate from core health. Task 093
performs no such acceptance. Historical 1.3.1 tag, artifact, hash, signatures and
source/release commit remain immutable.
