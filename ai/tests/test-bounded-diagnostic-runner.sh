#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/../.." && pwd -P)"
tool="$repo_root/ai/scripts/bounded-diagnostic-runner.sh"
work="$(mktemp -d /tmp/framework-diagnostic-test.XXXXXX)"
trap 'rm -rf -- "$work"' EXIT
passes=0

make_runner() {
    local name="$1" body="$2" path
    path="$work/$name"
    printf '#!/usr/bin/env bash\nset -euo pipefail\n%s\n' "$body" >"$path"
    chmod 0700 "$path"
    printf '%s\n' "$path"
}

make_manifest() {
    local name="$1" runner="$2" mode="${3:-read-only}" path
    path="$work/$name"
    printf 'schema\tframework.bounded-diagnostic/1\nmode\t%s\ntargets\tfixture_state\nactions\tinspect,compare\ntimeout_seconds\t2\nmax_output_bytes\t4096\nrunner_sha256\t%s\n' \
        "$mode" "$(sha256sum "$runner" | awk '{print $1}')" >"$path"
    printf '%s\n' "$path"
}

expect_success() { "$@" >/dev/null || { printf 'FAIL expected success\n' >&2; exit 1; }; passes=$((passes + 1)); }
expect_failure() { if "$@" >/dev/null 2>&1; then printf 'FAIL expected failure\n' >&2; exit 1; fi; passes=$((passes + 1)); }

runner="$(make_runner valid "printf '%s\\n' diagnostic.network=PASS diagnostic.classification=ENVIRONMENTAL_ISSUE diagnostic.cleanup=PASS")"
manifest="$(make_manifest valid.manifest "$runner")"
output="$($tool --manifest "$manifest" --runner "$runner")"
grep -Fq 'diagnostic.mode=read-only' <<<"$output"
grep -Fq 'diagnostic.classification=ENVIRONMENTAL_ISSUE' <<<"$output"
grep -Fq 'diagnostic.targets=fixture_state' <<<"$output"
passes=$((passes + 3))
expect_success "$tool" --manifest "$manifest" --runner "$runner"

bad_manifest="$(make_manifest bad-hash.manifest "$runner")"
sed -i 's/[0-9a-f]\{64\}$/0000000000000000000000000000000000000000000000000000000000000000/' "$bad_manifest"
expect_failure "$tool" --manifest "$bad_manifest" --runner "$runner"

runner_missing="$(make_runner missing "printf '%s\\n' diagnostic.result=PASS diagnostic.cleanup=PASS")"
expect_failure "$tool" --manifest "$(make_manifest missing.manifest "$runner_missing")" --runner "$runner_missing"

runner_secret="$(make_runner secret "printf '%s\\n' diagnostic.secret=REDACTED diagnostic.classification=INCONCLUSIVE diagnostic.cleanup=PASS")"
expect_failure "$tool" --manifest "$(make_manifest secret.manifest "$runner_secret")" --runner "$runner_secret"

runner_duplicate="$(make_runner duplicate "printf '%s\\n' diagnostic.result=PASS diagnostic.result=PASS diagnostic.classification=EXPECTED_BEHAVIOR diagnostic.cleanup=PASS")"
expect_failure "$tool" --manifest "$(make_manifest duplicate.manifest "$runner_duplicate")" --runner "$runner_duplicate"

runner_timeout="$(make_runner timeout 'sleep 3')"
expect_failure "$tool" --manifest "$(make_manifest timeout.manifest "$runner_timeout")" --runner "$runner_timeout"

runner_mutation="$(make_runner mutation "printf '%s\\n' diagnostic.change=PASS diagnostic.classification=PRODUCT_FRAMEWORK_DEFECT diagnostic.cleanup=PASS")"
expect_success "$tool" --manifest "$(make_manifest mutation.manifest "$runner_mutation" bounded-mutation)" --runner "$runner_mutation"

[[ -z "$(find /tmp -maxdepth 1 -type f -name 'framework-bounded-diagnostic.*' -user "$(id -u)" -print -quit)" ]]
passes=$((passes + 1))
printf 'PASS: %d bounded diagnostic runner assertions\n' "$passes"
