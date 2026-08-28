#!/bin/sh
set -eu

repo=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd -P)
test "$(tr -d '\r\n' < "$repo/VERSION")" = 1.2.0-rc.5
test -f "$repo/docs/release/RELEASE_NOTES_1.2.0-rc.5.md"
test -f "$repo/docs/release/CHANGE_NOTIFICATIONS.md"
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0.md"
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.6.md"
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.5.md"
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.4.md"
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.3.md"
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.2.md"
test ! -e "$repo/dist/qwsg-1.2.0-rc.5-linux-amd64.tar.gz"
test "$(QWSG_RELEASE_VALIDATE_ONLY=1 "$repo/scripts/build-release.sh")" = \
  'release build: identity 1.2.0-rc.5 is valid'
grep -F 'QWSG 1.1+ requires explicit SOURCE_DATE_EPOCH' "$repo/scripts/build-release.sh" >/dev/null
grep -F 'QWSG 1.1+ requires the full 40-character commit' "$repo/scripts/build-release.sh" >/dev/null
grep -F 'output archive or sidecar already exists' "$repo/scripts/build-release.sh" >/dev/null

work=$(mktemp -d /tmp/qwsg-release-check.XXXXXX)
trap 'rm -rf "$work"' EXIT HUP INT TERM
if "$repo/scripts/build-release.sh" >"$work/missing-metadata" 2>&1; then exit 1; fi
grep -F 'QWSG 1.1+ requires explicit SOURCE_DATE_EPOCH' "$work/missing-metadata" >/dev/null
if SOURCE_DATE_EPOCH=0 BUILD_COMMIT=0123456789abcdef "$repo/scripts/build-release.sh" >"$work/short-commit" 2>&1; then exit 1; fi
grep -F 'QWSG 1.1+ requires the full 40-character commit' "$work/short-commit" >/dev/null
mkdir "$work/dist"
touch "$work/dist/qwsg-1.2.0-rc.5-linux-amd64.tar.gz"
if SOURCE_DATE_EPOCH=0 BUILD_COMMIT=0123456789abcdef0123456789abcdef01234567 DIST_DIR="$work/dist" \
  "$repo/scripts/build-release.sh" >"$work/collision" 2>&1; then exit 1; fi
grep -F 'output archive or sidecar already exists' "$work/collision" >/dev/null
grep -F 'qwsg-1.2.0-rc.5-linux-amd64.tar.gz' "$repo/docs/installation/INSTALL.md" >/dev/null
grep -F 'README.md' "$repo/docs/installation/INSTALL.md" >/dev/null
grep -F 'INSTALL.md' "$repo/README.md" >/dev/null

validate_protocol() {
protocol=$1
expected_checkpoints=$2
test "$(grep -c '^## Checkpoint [0-9][0-9] ' "$protocol")" -eq "$expected_checkpoints"
awk '
  function verify() {
    if (!checkpoint) return
    if (!(purpose && action && expected && pass && fail && safe && retain)) exit 1
  }
  /^## Checkpoint [0-9][0-9] / {
    verify(); checkpoint=1
    purpose=action=expected=pass=fail=safe=retain=0
  }
  checkpoint && /^- \*\*Purpose:/ { purpose=1 }
  checkpoint && /^- \*\*Action:/ { action=1 }
  checkpoint && /^- \*\*Expected evidence:/ { expected=1 }
  checkpoint && /^- \*\*PASS:/ { pass=1 }
  checkpoint && /^- \*\*FAIL\/finding:/ { fail=1 }
  checkpoint && /^- \*\*Safe continuation:/ { safe=1 }
  checkpoint && /^- \*\*Retain\/redact:/ { retain=1 }
  END { verify() }
' "$protocol"
}

validate_protocol "$repo/docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.1.md" 16
validate_protocol "$repo/docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.2.md" 16
validate_protocol "$repo/docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.3.md" 25
validate_protocol "$repo/docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.4.md" 25
validate_protocol "$repo/docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.5.md" 26
test -f "$repo/docs/release/ACCEPTANCE_1.1.0-rc.2.md"
grep -F 'QWSG-049-F002' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.2.md" >/dev/null
grep -F 'QWSG-049-F003' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.2.md" >/dev/null
grep -F 'READY FOR QWSG 1.1.0 RELEASE' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.2.md" >/dev/null
test -f "$repo/docs/release/ACCEPTANCE_1.1.0-rc.3.md"
grep -F 'QWSG-051-F001' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.3.md" >/dev/null
grep -F 'QWSG-049-F002' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.3.md" >/dev/null
grep -F 'QWSG-049-F003' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.3.md" >/dev/null
grep -F 'OPEN, BLOCKING' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.3.md" >/dev/null
grep -F 'READY FOR QWSG 1.1.0 RELEASE' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.3.md" >/dev/null
grep -F 'NOT READY FOR QWSG 1.1.0 RELEASE' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.3.md" >/dev/null
grep -F 'QWSG-053-F001' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.3.md" >/dev/null
grep -F 'c3ba763701b7ee0340d4928b21c23276dfdc083536b08814157366310629a0cc' \
  "$repo/docs/release/ACCEPTANCE_1.1.0-rc.3.md" >/dev/null
test -f "$repo/docs/release/ACCEPTANCE_1.1.0-rc.4.md"
grep -F 'Candidate: `BUILT PRIVATELY; TRANSFERRED THROUGH OWNER-WORKSTATION FALLBACK`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.4.md" >/dev/null
grep -F 'QWSG-055-F001' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.4.md" >/dev/null
grep -F '7ad8a0f1be9bdbaa4403c3a816a6b474fd8e052934abd031047e4f82fe73a333' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.4.md" >/dev/null
grep -F 'QWSG-053-F001' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.4.md" >/dev/null
grep -F 'QWSG-051-F001' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.4.md" >/dev/null
grep -F 'QWSG-049-F002' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.4.md" >/dev/null
grep -F 'QWSG-049-F003' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.4.md" >/dev/null
grep -F 'READY FOR QWSG 1.1.0 RELEASE' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.4.md" >/dev/null
grep -F 'NOT READY FOR QWSG 1.1.0 RELEASE' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.4.md" >/dev/null
test -f "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md"
grep -F -- '- Candidate: `BUILT PRIVATELY; TRANSFERRED AND DESTINATION-VERIFIED`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F '1025d36d05b2f6f919f0ea4ec4a7029f67536000' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F '1787594463' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F '2026-08-24T18:01:03Z' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'qwsg-1.1.0-rc.5-linux-amd64.tar.gz' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F '2951350' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'cfe300c0f1f312d80120f74a9f24bed4a64387471bf2097ddc63d94f0fb2f7b0' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F '69f3eb4bf89dc126a7eafd08354eec37a941014171b3d1d70c6e6a4cf52e5eb0' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'ae51aca0bc4ddc61b0daea3a87f0acabcde5ec9fd8fadddc050f0786d6915e9e' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F '5484aab96d5c3748e81b065fdb11ec8c34385589bb07ee7ea1b2b35fdffa6b93' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'PASS; PRIVACY-SAFE LOCAL EVIDENCE RECORDED' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'LOCAL GATE B RETEST PASS; historical state remains OPEN/BLOCKING' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F -- '- Transfer: `PASS — OWNER-WORKSTATION FALLBACK; READ-ONLY RECOVERY VERIFIED`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F -- '- External acceptance: `TERMINATED INCOMPLETE BY OWNER PROCESS DECISION`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F -- '- Task classification: `COMPLETE WITH DISCLOSED ACCEPTANCE LIMITATIONS`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F -- '- Core result: `RC.5 CORE CLEAN-HOST FUNCTIONAL PROOF: ACHIEVED`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F -- '- Formal result: `FORMAL 26-CHECKPOINT RELEASE-READINESS ACCEPTANCE: INCOMPLETE`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F -- '- Final verdict: `NOT READY FOR QWSG 1.1.0 RELEASE — FORMAL CERTIFICATION TERMINATED INCOMPLETE BY OWNER DECISION`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
test "$(grep -c '| NOT EXECUTED |' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md")" -ge 23
grep -F 'No product defect is inferred solely from missing procedural' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'ADDITIVE EXTERNALLY VERIFIED CORRECTED IN RC.5; HISTORICAL OPEN/BLOCKING IMMUTABLE' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'QWSG-051-F001' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'QWSG-053-F001' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'QWSG-049-F002' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'QWSG-049-F003' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'git.quantumwizard.hu' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F '`wget` and `curl` examples' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
practical="$repo/docs/release/PRACTICAL_RELEASE_ACCEPTANCE.md"
distribution="$repo/docs/release/FORGEJO_DISTRIBUTION.md"
test -f "$practical"
test "$(grep -cE '^[0-9]+\. \*\*' "$practical")" -eq 12
grep -F 'one bounded Owner-authorized acceptance run' "$practical" >/dev/null
grep -F 'mandatory evidence remains missing and is never converted to PASS' "$practical" >/dev/null
grep -F 'reporting alone does not invalidate a clean host' "$practical" >/dev/null
test -f "$distribution"
grep -F 'FORGEJO_BASE/OWNER/REPOSITORY/releases/download/TAG/ASSET' "$distribution" >/dev/null
grep -F 'wget "${release_base}/${archive}"' "$distribution" >/dev/null
grep -F 'curl -fLO "${release_base}/${archive}"' "$distribution" >/dev/null
grep -F 'sha256sum -c "${archive}.sha256"' "$distribution" >/dev/null
grep -F 'first release verified under this contract' "$distribution" >/dev/null
grep -F 'Future Smart Installer contract' "$distribution" >/dev/null
grep -F -- '-buildvcs=false' "$repo/Makefile" >/dev/null
grep -F '1.1.0-rc.4' "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.4.md" >/dev/null
grep -F '1.1.0-rc.5' "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.5.md" >/dev/null
grep -F '1.1.0-rc.6' "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.6.md" >/dev/null
grep -F 'QWSG 1.1.0' "$repo/docs/release/RELEASE_NOTES_1.1.0.md" >/dev/null

printf '%s\n' 'PASS: QWSG 1.2.0-rc.5 release plumbing'
