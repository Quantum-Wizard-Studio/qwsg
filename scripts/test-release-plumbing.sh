#!/bin/sh
set -eu

repo=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd -P)
test "$(tr -d '\r\n' < "$repo/VERSION")" = 1.1.0-rc.3
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.3.md"
test -f "$repo/docs/release/RELEASE_NOTES_1.1.0-rc.2.md"
test ! -e "$repo/dist/qwsg-1.1.0-rc.3-linux-amd64.tar.gz"
test "$(QWSG_RELEASE_VALIDATE_ONLY=1 "$repo/scripts/build-release.sh")" = \
  'release build: identity 1.1.0-rc.3 is valid'
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
touch "$work/dist/qwsg-1.1.0-rc.3-linux-amd64.tar.gz"
if SOURCE_DATE_EPOCH=0 BUILD_COMMIT=0123456789abcdef0123456789abcdef01234567 DIST_DIR="$work/dist" \
  "$repo/scripts/build-release.sh" >"$work/collision" 2>&1; then exit 1; fi
grep -F 'output archive or sidecar already exists' "$work/collision" >/dev/null
grep -F 'qwsg-1.1.0-rc.3-linux-amd64.tar.gz' "$repo/docs/installation/INSTALL.md" >/dev/null
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

printf '%s\n' 'PASS: QWSG 1.1.0-rc.3 release plumbing'
