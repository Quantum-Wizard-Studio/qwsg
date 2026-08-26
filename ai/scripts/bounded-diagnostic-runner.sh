#!/usr/bin/env bash
set -euo pipefail

fail() { printf 'Bounded diagnostic error: %s\n' "$1" >&2; exit 1; }
usage() { printf 'Usage: %s --manifest <file> --runner <file>\n' "${0##*/}"; }

[[ ${1:-} == --manifest && -n ${2:-} && ${3:-} == --runner && -n ${4:-} && $# == 4 ]] || {
    usage >&2
    exit 2
}
manifest="$2"
runner="$4"
for file in "$manifest" "$runner"; do
    [[ -f "$file" && ! -L "$file" ]] || fail 'manifest and runner must be regular non-symlink files'
done
[[ -x "$runner" ]] || fail 'runner is not executable'
[[ "$(stat -c %u "$manifest")" == "$(id -u)" && "$(stat -c %u "$runner")" == "$(id -u)" ]] ||
    fail 'manifest and runner must be owned by the invoking user'
iconv -f UTF-8 -t UTF-8 "$manifest" >/dev/null 2>&1 || fail 'manifest is not valid UTF-8'

declare -A data=() allowed=(
    [schema]=1 [mode]=1 [targets]=1 [actions]=1 [timeout_seconds]=1
    [max_output_bytes]=1 [runner_sha256]=1
)
while IFS=$'\t' read -r key value extra || [[ -n ${key:-}${value:-}${extra:-} ]]; do
    [[ -n ${key:-} && -n ${value:-} && -z ${extra:-} ]] || fail 'manifest requires exactly two tab-separated fields per line'
    [[ -n ${allowed[$key]:-} ]] || fail "unknown manifest key: $key"
    [[ -z ${data[$key]+set} ]] || fail "duplicate manifest key: $key"
    data[$key]="$value"
done <"$manifest"
for key in schema mode targets actions timeout_seconds max_output_bytes runner_sha256; do
    [[ -n ${data[$key]:-} ]] || fail "missing manifest key: $key"
done
[[ ${data[schema]} == framework.bounded-diagnostic/1 ]] || fail 'unsupported manifest schema'
[[ ${data[mode]} == read-only || ${data[mode]} == bounded-mutation ]] || fail 'invalid diagnostic mode'
[[ ${data[targets]} =~ ^[a-z0-9_]+(,[a-z0-9_]+)*$ ]] || fail 'targets must be comma-separated logical identifiers'
[[ ${data[actions]} =~ ^[a-z0-9_]+(,[a-z0-9_]+)*$ ]] || fail 'actions must be comma-separated logical identifiers'
[[ ${data[timeout_seconds]} =~ ^[1-9][0-9]{0,2}$ && ${data[timeout_seconds]} -le 900 ]] || fail 'timeout must be 1..900 seconds'
[[ ${data[max_output_bytes]} =~ ^[1-9][0-9]{2,4}$ && ${data[max_output_bytes]} -le 65536 ]] || fail 'output limit must be 100..65536 bytes'
[[ ${data[runner_sha256]} =~ ^[0-9a-f]{64}$ ]] || fail 'invalid runner SHA-256'
actual_sha="$(sha256sum "$runner" | awk '{print $1}')"
[[ "$actual_sha" == "${data[runner_sha256]}" ]] || fail 'runner identity mismatch'

capture="$(mktemp /tmp/framework-bounded-diagnostic.XXXXXX)"
chmod 0600 "$capture"
trap 'rm -f -- "$capture"' EXIT HUP INT TERM
set +e
env -i PATH=/usr/local/bin:/usr/bin:/bin LANG=C.UTF-8 LC_ALL=C.UTF-8 \
    timeout --signal=TERM --kill-after=5 "${data[timeout_seconds]}" "$runner" >"$capture" 2>&1
status=$?
set -e
[[ $status -eq 0 ]] || fail "runner failed or timed out (status $status)"
[[ "$(stat -c %s "$capture")" -le ${data[max_output_bytes]} ]] || fail 'runner output exceeds declared limit'

declare -A seen=()
classification=false cleanup=false
while IFS= read -r line || [[ -n "$line" ]]; do
    [[ "$line" =~ ^([a-z][a-z0-9_]*([.][a-z0-9_]+)*)=([A-Za-z0-9_.:/,+\ -]{1,160})$ ]] || fail 'runner output is not privacy-safe key=value data'
    key="${BASH_REMATCH[1]}"; value="${BASH_REMATCH[3]}"
    [[ -z ${seen[$key]:-} ]] || fail "duplicate diagnostic key: $key"
    seen[$key]=1
    case "$key" in *password*|*secret*|*token*|*credential*|*recipient*|*sender*|*private_key*) fail 'sensitive diagnostic key is forbidden' ;; esac
    if [[ "$key" == diagnostic.classification ]]; then
        [[ "$value" =~ ^(PRODUCT_FRAMEWORK_DEFECT|TEST_OR_ACCEPTANCE_DEFECT|ENVIRONMENTAL_ISSUE|EXPECTED_BEHAVIOR|INCONCLUSIVE)$ ]] || fail 'invalid diagnostic classification'
        classification=true
    fi
    [[ "$key" != diagnostic.cleanup || "$value" != PASS ]] || cleanup=true
done <"$capture"
[[ "$classification" == true ]] || fail 'diagnostic classification is missing'
[[ "$cleanup" == true ]] || fail 'diagnostic cleanup PASS is missing'

printf 'diagnostic.runner_sha256=%s\n' "$actual_sha"
printf 'diagnostic.mode=%s\n' "${data[mode]}"
printf 'diagnostic.targets=%s\n' "${data[targets]}"
printf 'diagnostic.actions=%s\n' "${data[actions]}"
cat -- "$capture"
