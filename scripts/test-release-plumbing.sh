#!/bin/sh
set -eu

repo=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd -P)
test "$(tr -d '\r\n' < "$repo/VERSION")" = 1.1.0-rc.5
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.5.md"
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.4.md"
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.3.md"
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.2.md"
test ! -e "$repo/dist/qwsg-1.1.0-rc.5-linux-amd64.tar.gz"
test "$(QWSG_RELEASE_VALIDATE_ONLY=1 "$repo/scripts/build-release.sh")" = \
  'release build: identity 1.1.0-rc.5 is valid'
grep -F '1.1 RC requires explicit SOURCE_DATE_EPOCH' "$repo/scripts/build-release.sh" >/dev/null
grep -F '1.1 RC requires the full 40-character commit' "$repo/scripts/build-release.sh" >/dev/null
grep -F 'output archive or sidecar already exists' "$repo/scripts/build-release.sh" >/dev/null

work=$(mktemp -d /tmp/qwsg-release-check.XXXXXX)
trap 'rm -rf "$work"' EXIT HUP INT TERM
if "$repo/scripts/build-release.sh" >"$work/missing-metadata" 2>&1; then exit 1; fi
grep -F '1.1 RC requires explicit SOURCE_DATE_EPOCH' "$work/missing-metadata" >/dev/null
if SOURCE_DATE_EPOCH=0 BUILD_COMMIT=0123456789abcdef "$repo/scripts/build-release.sh" >"$work/short-commit" 2>&1; then exit 1; fi
grep -F '1.1 RC requires the full 40-character commit' "$work/short-commit" >/dev/null
mkdir "$work/dist"
touch "$work/dist/qwsg-1.1.0-rc.5-linux-amd64.tar.gz"
if SOURCE_DATE_EPOCH=0 BUILD_COMMIT=0123456789abcdef0123456789abcdef01234567 DIST_DIR="$work/dist" \
  "$repo/scripts/build-release.sh" >"$work/collision" 2>&1; then exit 1; fi
grep -F 'output archive or sidecar already exists' "$work/collision" >/dev/null
grep -F 'qwsg-1.1.0-rc.5-linux-amd64.tar.gz' "$repo/docs/installation/INSTALL.md" >/dev/null
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
grep -F -- '- Candidate: `BUILT PRIVATELY; NOT TRANSFERRED`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
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
grep -F -- '- Transfer: `NOT STARTED`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F -- '- External acceptance: `NOT STARTED`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F -- '- Final verdict: `PENDING`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
test "$(grep -c '| NOT EXECUTED |' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md")" -ge 26
grep -F 'QWSG-055-F001 remains `OPEN, BLOCKING`' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'QWSG-051-F001' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'QWSG-053-F001' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'QWSG-049-F002' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'QWSG-049-F003' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F 'git.quantumwizard.hu' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F '`wget` and `curl` examples' "$repo/docs/release/ACCEPTANCE_1.1.0-rc.5.md" >/dev/null
grep -F -- '-buildvcs=false' "$repo/Makefile" >/dev/null
grep -F '1.1.0-rc.4' "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.4.md" >/dev/null
grep -F '1.1.0-rc.5' "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.5.md" >/dev/null

printf '%s\n' 'PASS: QWSG 1.1.0-rc.5 release plumbing'
