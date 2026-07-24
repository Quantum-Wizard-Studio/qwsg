#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"
project_root="$(cd -- "$script_dir/../.." && pwd -P)"
prompts_dir="$project_root/ai/prompts"
archive_dir="$project_root/ai/archive_prompts"
history_dir="$project_root/ai/history"
test_root="$project_root/ai/test_tasks"

fail() {
    printf 'Error: %s\n' "$1" >&2
    exit 1
}

usage() {
    cat <<'EOF'
Usage:
  ./ai/scripts/divert-task-to-test.sh --check
  ./ai/scripts/divert-task-to-test.sh \
    --authority <owner> \
    --reason <reason> \
    --disposition aborted-test \
    --release-production-id yes \
    [--disposition-statement-file <regular UTF-8 file>] \
    --token DIVERT-TO-TEST

The diversion command transactionally removes one incomplete active production
task from the production lifecycle and preserves it under ai/test_tasks/. It
requires explicit Project Owner authority and never marks the task complete.
EOF
}

require_root() {
    [[ "$(pwd -P)" == "$project_root" ]] ||
        fail "run from the project root: $project_root"
    [[ -f "$project_root/VERSION" && -x "$project_root/bin/job" ]] ||
        fail 'QWSG project markers are missing'
    [[ -d "$prompts_dir" && -d "$archive_dir" && -d "$history_dir" ]] ||
        fail 'production lifecycle directories are missing'
    if [[ -e "$project_root/ai/framework/VERSION" ]]; then
        [[ -x "$project_root/ai/scripts/framework-check.sh" ]] ||
            fail 'framework validator is missing or not executable'
        "$project_root/ai/scripts/framework-check.sh" --quiet ||
            fail 'reusable framework validation failed'
    fi
}

extract_one() {
    local pattern="$1" file="$2" label="$3" values=()
    mapfile -t values < <(sed -nE "$pattern" "$file")
    (( ${#values[@]} == 1 )) ||
        fail "$label must occur exactly once in ${file#$project_root/}"
    printf '%s\n' "${values[0]}"
}

check_test_tasks() {
    [[ -d "$test_root" ]] || {
        printf 'Test-task audit: VALID (no diverted test tasks)\n'
        return 0
    }

    local dir name expected=1 test_id original_id original_slug authority reason
    local prompt history disposition checksum
    declare -A seen_original_ids=()
    mapfile -t test_dirs < <(
        find -P "$test_root" -mindepth 1 -maxdepth 1 -type d \
            -name '[0-9][0-9][0-9]_TEST_TASK' -print | LC_ALL=C sort
    )
    if find -P "$test_root" -mindepth 1 -maxdepth 1 \
        ! -name '[0-9][0-9][0-9]_TEST_TASK' -print -quit | grep -q .; then
        fail 'unexpected entry exists in ai/test_tasks'
    fi

    for dir in "${test_dirs[@]}"; do
        name="${dir##*/}"
        printf -v test_id '%03d_TEST_TASK' "$expected"
        [[ "$name" == "$test_id" ]] ||
            fail "test-task numbering is not contiguous at $name; expected $test_id"
        expected=$((expected + 1))

        prompt="$dir/prompt/original-prompt.md"
        history="$dir/history/original-history.md"
        disposition="$dir/reports/DISPOSITION.md"
        checksum="$dir/SHA256SUMS"
        [[ -f "$prompt" && ! -L "$prompt" ]] ||
            fail "$name lacks a regular preserved prompt"
        [[ -f "$history" && ! -L "$history" ]] ||
            fail "$name lacks a regular preserved history"
        [[ -f "$disposition" && ! -L "$disposition" ]] ||
            fail "$name lacks a regular disposition record"
        [[ -f "$dir/DIVERSION_MANIFEST.md" && ! -L "$dir/DIVERSION_MANIFEST.md" ]] ||
            fail "$name lacks a diversion manifest"
        [[ -f "$checksum" && ! -L "$checksum" ]] ||
            fail "$name lacks SHA256SUMS"

        grep -Fxq -- "- Test Task ID: \`$name\`" "$disposition" ||
            fail "$name disposition has an invalid Test Task ID"
        original_id="$(extract_one 's/^- Original Task ID: `([0-9]{3})`$/\1/p' \
            "$disposition" "$name Original Task ID")"
        original_slug="$(extract_one 's/^- Original Task Slug: `([a-z0-9]+(-[a-z0-9]+)*)`$/\1/p' \
            "$disposition" "$name Original Task Slug")"
        authority="$(extract_one 's/^- Disposition Authority: `([^`]+)`$/\1/p' \
            "$disposition" "$name Disposition Authority")"
        reason="$(extract_one 's/^- Reason: `([^`]+)`$/\1/p' \
            "$disposition" "$name Reason")"
        [[ -n "$authority" && "$authority" == *'Project Owner'* ]] ||
            fail "$name lacks Project Owner disposition authority"
        [[ -n "$reason" ]] || fail "$name has an empty diversion reason"
        grep -Fxq -- '- Status: `aborted-test`' "$disposition" ||
            fail "$name disposition status is not aborted-test"
        grep -Fxq -- '- Result: `incomplete`' "$disposition" ||
            fail "$name disposition result is not incomplete"
        grep -Fxq -- '- Production Sequence Consumed: `no`' "$disposition" ||
            fail "$name does not release its production sequence number"
        grep -Fxq -- '- Retry Allowed: `yes`' "$disposition" ||
            fail "$name lacks explicit retry metadata"
        grep -Fxq -- "- Task ID: \`$original_id\`" "$prompt" ||
            fail "$name preserved prompt Task ID differs from disposition"
        grep -Fxq -- "- Task slug: \`$original_slug\`" "$prompt" ||
            fail "$name preserved prompt slug differs from disposition"
        grep -Fxq -- "- Task ID: \`$original_id\`" "$history" ||
            fail "$name preserved history Task ID differs from disposition"
        grep -Fxq -- "- Task slug: \`$original_slug\`" "$history" ||
            fail "$name preserved history slug differs from disposition"
        [[ -z "${seen_original_ids[$original_id]:-}" ]] ||
            fail "duplicate diverted Original Task ID: $original_id"
        seen_original_ids["$original_id"]="$name"
        (cd "$dir" && sha256sum -c SHA256SUMS >/dev/null) ||
            fail "$name evidence checksum validation failed"
    done

    printf 'Test-task audit: VALID (%d diverted task(s))\n' "${#test_dirs[@]}"
}

load_active() {
    mapfile -d '' active_files < <(
        find -P "$prompts_dir" -maxdepth 1 -type f -name '*.md' -print0
    )
    (( ${#active_files[@]} == 1 )) ||
        fail "diversion requires exactly one active production prompt; found ${#active_files[@]}"
    active_prompt="${active_files[0]}"
    active_name="${active_prompt##*/}"
    [[ "$active_name" =~ ^([0-9]{3})_CURRENT_TASK\.md$ ]] ||
        fail "invalid active prompt filename: $active_name"
    original_id="${BASH_REMATCH[1]}"
    prompt_id="$(extract_one 's/^- Task ID: `([0-9]{3})`$/\1/p' \
        "$active_prompt" 'prompt Task ID')"
    original_slug="$(extract_one 's/^- Task slug: `([a-z0-9]+(-[a-z0-9]+)*)`$/\1/p' \
        "$active_prompt" 'prompt Task slug')"
    prompt_status="$(extract_one 's/^- Status: `?([^`]+)`?$/\1/p' \
        "$active_prompt" 'prompt Status')"
    [[ "$prompt_id" == "$original_id" ]] ||
        fail 'active prompt filename and Task ID differ'
    [[ "$prompt_status" != complete* ]] ||
        fail 'a complete production task must use the normal archival lifecycle'

    mapfile -d '' histories < <(
        find -P "$history_dir" -maxdepth 1 -type f \
            -name "${original_id}_????-??-??_*.md" -print0
    )
    (( ${#histories[@]} == 1 )) ||
        fail "exactly one active-task history is required; found ${#histories[@]}"
    active_history="${histories[0]}"
    history_id="$(extract_one 's/^- Task ID: `([0-9]{3})`$/\1/p' \
        "$active_history" 'history Task ID')"
    history_slug="$(extract_one 's/^- Task slug: `([a-z0-9]+(-[a-z0-9]+)*)`$/\1/p' \
        "$active_history" 'history Task slug')"
    history_status="$(extract_one 's/^- Status: `?([^`]+)`?$/\1/p' \
        "$active_history" 'history Status')"
    [[ "$history_id" == "$original_id" && "$history_slug" == "$original_slug" ]] ||
        fail 'active prompt and history identity differ'
    [[ "$history_status" != complete* ]] ||
        fail 'a complete production history cannot be diverted as incomplete'
    if find -P "$archive_dir" -maxdepth 1 -type f \
        -name "${original_id}_????-??-??_*.md" -print -quit | grep -q .; then
        fail "production archive already contains Task $original_id"
    fi
}

next_test_identity() {
    mkdir -p "$test_root"
    local number=1 candidate
    while (( number <= 999 )); do
        printf -v candidate '%03d_TEST_TASK' "$number"
        if [[ ! -e "$test_root/$candidate" ]]; then
            test_id="$candidate"
            test_dir="$test_root/$candidate"
            return 0
        fi
        number=$((number + 1))
    done
    fail 'three-digit test-task ID range exhausted'
}

divert_active() {
    local authority="$1" reason="$2" disposition="$3" release="$4" token="$5"
    local statement_file="$6"
    [[ -n "$authority" && "$authority" == *'Project Owner'* ]] ||
        fail 'authority must identify the Project Owner'
    [[ -n "$reason" ]] || fail 'diversion reason must not be empty'
    [[ "$disposition" == aborted-test ]] ||
        fail 'disposition must be exactly aborted-test'
    [[ "$release" == yes ]] ||
        fail 'release-production-id must be exactly yes'
    [[ "$token" == DIVERT-TO-TEST ]] ||
        fail 'explicit override token must be exactly DIVERT-TO-TEST'
    if [[ -n "$statement_file" ]]; then
        [[ -f "$statement_file" && ! -L "$statement_file" && -s "$statement_file" ]] ||
            fail 'disposition statement must be a nonempty regular file'
        iconv -f UTF-8 -t UTF-8 "$statement_file" >/dev/null 2>&1 ||
            fail 'disposition statement is not valid UTF-8'
        if od -An -t x1 "$statement_file" | grep -Eq '(^| )00( |$)'; then
            fail 'disposition statement contains a NUL byte'
        fi
    fi

    load_active
    check_test_tasks >/dev/null
    next_test_identity

    local disposition_date timestamp temp_dir final_installed=false
    local prompt_moved=false history_moved=false status
    local prompt_before history_before prompt_after history_after
    disposition_date="$(date -u +%F)"
    timestamp="$(date -u +%Y%m%dT%H%M%SZ)"
    temp_dir="$(mktemp -d "$test_root/.divert-task.XXXXXX")"
    mkdir -p "$temp_dir/prompt" "$temp_dir/history" "$temp_dir/reports"
    chmod 2770 "$temp_dir" "$temp_dir/prompt" "$temp_dir/history" "$temp_dir/reports"
    prompt_before="$(sha256sum "$active_prompt" | awk '{print $1}')"
    history_before="$(sha256sum "$active_history" | awk '{print $1}')"

    cleanup() {
        status="$?"
        if (( status != 0 )); then
            if [[ "$final_installed" == true && -d "$test_dir" ]]; then
                [[ -e "$active_prompt" || ! -e "$test_dir/prompt/original-prompt.md" ]] ||
                    mv -T --no-clobber -- "$test_dir/prompt/original-prompt.md" "$active_prompt"
                [[ -e "$active_history" || ! -e "$test_dir/history/original-history.md" ]] ||
                    mv -T --no-clobber -- "$test_dir/history/original-history.md" "$active_history"
                rm -rf -- "$test_dir"
            elif [[ -d "$temp_dir" ]]; then
                [[ "$prompt_moved" != true || -e "$active_prompt" ||
                    ! -e "$temp_dir/prompt/original-prompt.md" ]] ||
                    mv -T --no-clobber -- "$temp_dir/prompt/original-prompt.md" "$active_prompt"
                [[ "$history_moved" != true || -e "$active_history" ||
                    ! -e "$temp_dir/history/original-history.md" ]] ||
                    mv -T --no-clobber -- "$temp_dir/history/original-history.md" "$active_history"
                rm -rf -- "$temp_dir"
            fi
        fi
        trap - EXIT
        exit "$status"
    }
    trap cleanup EXIT

    mv -T --no-clobber -- "$active_prompt" "$temp_dir/prompt/original-prompt.md"
    prompt_moved=true
    [[ "${QWSG_DIVERSION_TEST_FAIL_AFTER:-}" != prompt ]] ||
        fail 'injected failure after prompt move'
    mv -T --no-clobber -- "$active_history" "$temp_dir/history/original-history.md"
    history_moved=true
    [[ "${QWSG_DIVERSION_TEST_FAIL_AFTER:-}" != history ]] ||
        fail 'injected failure after history move'

    cat >"$temp_dir/reports/DISPOSITION.md" <<EOF
# Aborted Test Task Disposition

- Test Task ID: \`$test_id\`
- Original Task ID: \`$original_id\`
- Original Task Slug: \`$original_slug\`
- Status: \`aborted-test\`
- Result: \`incomplete\`
- Disposition Authority: \`$authority\`
- Disposition Date: \`$disposition_date\`
- Reason: \`$reason\`
- Production Sequence Consumed: \`no\`
- Retry Allowed: \`yes\`
EOF

    if [[ -n "$statement_file" ]]; then
        printf '\n' >>"$temp_dir/reports/DISPOSITION.md"
        cat -- "$statement_file" >>"$temp_dir/reports/DISPOSITION.md"
        printf '\n' >>"$temp_dir/reports/DISPOSITION.md"
    else
        cat >>"$temp_dir/reports/DISPOSITION.md" <<EOF

The original production Task $original_id (\`$original_slug\`) was stopped by
$authority for the recorded reason.

The task was not completed, and no successful completion state is claimed.

The preserved task is reclassified as an aborted test task so that it no longer
blocks the production Engineering Lifecycle.

The normal production Task ID $original_id is released for reuse by a new clean
task.
EOF
    fi

    prompt_after="$(sha256sum "$temp_dir/prompt/original-prompt.md" | awk '{print $1}')"
    history_after="$(sha256sum "$temp_dir/history/original-history.md" | awk '{print $1}')"
    [[ "$prompt_before" == "$prompt_after" && "$history_before" == "$history_after" ]] ||
        fail 'preserved lifecycle evidence changed during diversion'

    cat >"$temp_dir/DIVERSION_MANIFEST.md" <<EOF
# Diversion Manifest

- Transaction UTC: \`$timestamp\`
- Test Task ID: \`$test_id\`
- Original production prompt:
  \`ai/prompts/$active_name\`
- Preserved prompt:
  \`ai/test_tasks/$test_id/prompt/original-prompt.md\`
- Original production history:
  \`ai/history/${active_history##*/}\`
- Preserved history:
  \`ai/test_tasks/$test_id/history/original-history.md\`
- Disposition: \`aborted-test\`
- Completion claimed: \`no\`
- Production ID released: \`yes\`
- Owner authority: \`$authority\`
- Reason: \`$reason\`
- Prompt SHA-256 before move: \`$prompt_before\`
- Prompt SHA-256 after move: \`$prompt_after\`
- History SHA-256 before move: \`$history_before\`
- History SHA-256 after move: \`$history_after\`

The original prompt and history were moved without content modification. Their
SHA-256 hashes are recorded in \`SHA256SUMS\`. External snapshots and repair
artifacts referenced by the history remain at their recorded paths and were not
deleted or modified by this transaction.
EOF

    (
        cd "$temp_dir"
        sha256sum prompt/original-prompt.md history/original-history.md \
            reports/DISPOSITION.md DIVERSION_MANIFEST.md >SHA256SUMS
        sha256sum -c SHA256SUMS >/dev/null
    )
    chmod 0660 "$temp_dir/prompt/original-prompt.md" \
        "$temp_dir/history/original-history.md" \
        "$temp_dir/reports/DISPOSITION.md" \
        "$temp_dir/DIVERSION_MANIFEST.md" "$temp_dir/SHA256SUMS"

    mv -T --no-clobber -- "$temp_dir" "$test_dir"
    final_installed=true
    [[ "${QWSG_DIVERSION_TEST_FAIL_AFTER:-}" != install ]] ||
        fail 'injected failure after test-task install'

    "$project_root/bin/job" --check >/dev/null ||
        fail 'production idle lifecycle validation failed'
    "$project_root/ai/scripts/next-task.sh" --check >/dev/null ||
        fail 'production lifecycle consistency validation failed'
    check_test_tasks >/dev/null || fail 'test-task audit failed'

    trap - EXIT
    printf 'Diverted Test Task ID: %s\nOriginal production Task ID: %s\nOriginal slug: %s\nDisposition: aborted-test\nResult: incomplete\nProduction Sequence Consumed: no\nTest task: %s\nProduction lifecycle: IDLE AND VALID\nNext production Task ID: %s\n' \
        "$test_id" "$original_id" "$original_slug" \
        "${test_dir#$project_root/}" "$original_id"
}

require_root
if [[ "${1:-}" == --check ]]; then
    (( $# == 1 )) || fail '--check accepts no other arguments'
    check_test_tasks
    exit 0
fi

authority='' reason='' disposition='' release='' token='' statement_file=''
while (( $# > 0 )); do
    case "$1" in
        --authority) (( $# >= 2 )) || fail '--authority requires a value'; authority="$2"; shift 2 ;;
        --reason) (( $# >= 2 )) || fail '--reason requires a value'; reason="$2"; shift 2 ;;
        --disposition) (( $# >= 2 )) || fail '--disposition requires a value'; disposition="$2"; shift 2 ;;
        --release-production-id) (( $# >= 2 )) || fail '--release-production-id requires a value'; release="$2"; shift 2 ;;
        --disposition-statement-file) (( $# >= 2 )) || fail '--disposition-statement-file requires a value'; statement_file="$2"; shift 2 ;;
        --token) (( $# >= 2 )) || fail '--token requires a value'; token="$2"; shift 2 ;;
        --help|-h) usage; exit 0 ;;
        *) usage >&2; fail "unknown option: $1" ;;
    esac
done

divert_active "$authority" "$reason" "$disposition" "$release" "$token" "$statement_file"
