#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/../.." && pwd -P)"
work="$(mktemp -d /tmp/qwsg-diversion-test.XXXXXX)"
trap 'rm -rf -- "$work"' EXIT
passes=0
today="$(date -u +%F)"

new_fixture() {
    local name="$1" root="$work/$1"
    mkdir -p "$root/ai/scripts" "$root/ai/prompts" "$root/ai/archive_prompts" \
        "$root/ai/history" "$root/ai/core" "$root/ai/test_tasks" "$root/bin"
    cp "$repo_root/ai/scripts/divert-task-to-test.sh" \
        "$repo_root/ai/scripts/task-builder.sh" \
        "$repo_root/ai/scripts/next-task.sh" "$root/ai/scripts/"
    cp "$repo_root/bin/job" "$root/bin/job"
    chmod +x "$root/ai/scripts/"*.sh "$root/bin/job"
    printf 'test\n' >"$root/VERSION"
    printf '# marker\n' >"$root/ai/core/00_PROJECT_PHILOSOPHY.md"
    printf '# marker\n' >"$root/ai/core/08_JOB_TEMPLATE.md"
    cat >"$root/ai/archive_prompts/014_${today}_completed-baseline.md" <<'EOF'
# Current Engineering Task 014: Completed Baseline
## Task Metadata
- Task ID: `014`
- Task slug: `completed-baseline`
- Status: `complete`
EOF
    cat >"$root/ai/history/014_${today}_completed-baseline.md" <<'EOF'
# Task History 014: Completed Baseline
## Task metadata
- Task ID: `014`
- Task slug: `completed-baseline`
- Status: `complete`
## Starting state
Recorded.
## Snapshot
Verified.
## Work performed
Complete.
## Verification
Passed.
## Rollback
Verified.
## Completion state
`complete`
EOF
    cat >"$root/ai/prompts/015_CURRENT_TASK.md" <<'EOF'
# Current Engineering Task 015: Interrupted
## Task Metadata
- Task ID: `015`
- Task slug: `interrupted-method`
- Status: `active`
EOF
    cat >"$root/ai/history/015_${today}_interrupted-method.md" <<'EOF'
# Task History 015: Interrupted
## Task metadata
- Task ID: `015`
- Task slug: `interrupted-method`
- Status: `active — failed attempts recorded`
## Work performed
Attempt 1 failed.
## Completion state
`active — incomplete`
EOF
    printf 'do not touch\n' >"$root/unrelated-untracked.txt"
    printf '%s\n' "$root"
}

new_input() {
    local name="$1" dir field
    dir="$work/input-$name"
    mkdir -p "$dir"
    printf 'replacement-task\n' >"$dir/slug"
    printf 'Replacement Task\n' >"$dir/title"
    printf 'Attila — Project Owner\n' >"$dir/authority"
    printf 'Hungarian\n' >"$dir/language"
    printf 'APPROVE\n' >"$dir/approval"
    for field in objective scope out-of-scope starting-state snapshot risk \
        planned-work rollback deliverables verification documentation completion; do
        printf 'Validated %s content.\n' "$field" >"$dir/$field"
    done
    printf '%s\n' "$dir"
}

divert() {
    ./ai/scripts/divert-task-to-test.sh \
        --authority 'Attila — Project Owner' \
        --reason 'Extended unsuccessful execution attempts; original method abandoned' \
        --disposition aborted-test \
        --release-production-id yes \
        --token DIVERT-TO-TEST
}

expect_success() {
    local root="$1"; shift
    (cd "$root" && "$@") >/dev/null ||
        { printf 'FAIL expected success: %s\n' "$*" >&2; exit 1; }
    passes=$((passes + 1))
}

expect_failure() {
    local root="$1"; shift
    if (cd "$root" && "$@") >/dev/null 2>&1; then
        printf 'FAIL expected failure: %s\n' "$*" >&2
        exit 1
    fi
    passes=$((passes + 1))
}

assert_original() {
    local root="$1"
    [[ -f "$root/ai/prompts/015_CURRENT_TASK.md" ]] ||
        { printf 'FAIL active prompt not restored\n' >&2; exit 1; }
    [[ -f "$root/ai/history/015_${today}_interrupted-method.md" ]] ||
        { printf 'FAIL active history not restored\n' >&2; exit 1; }
    [[ ! -e "$root/ai/test_tasks/001_TEST_TASK" ]] ||
        { printf 'FAIL partial test task remains\n' >&2; exit 1; }
    [[ "$(cat "$root/unrelated-untracked.txt")" == 'do not touch' ]] ||
        { printf 'FAIL unrelated file changed\n' >&2; exit 1; }
    passes=$((passes + 1))
}

# Incomplete tasks remain blocked from the normal archival path.
root="$(new_fixture normal-archive-blocked)"
expect_failure "$root" ./ai/scripts/next-task.sh --prepare --slug replacement-task

# Explicit authority, reason, and token are independently mandatory.
root="$(new_fixture missing-authority)"
expect_failure "$root" ./ai/scripts/divert-task-to-test.sh \
    --authority Agent --reason reason --disposition aborted-test \
    --release-production-id yes --token DIVERT-TO-TEST
root="$(new_fixture missing-reason)"
expect_failure "$root" ./ai/scripts/divert-task-to-test.sh \
    --authority 'Attila — Project Owner' --reason '' --disposition aborted-test \
    --release-production-id yes --token DIVERT-TO-TEST
root="$(new_fixture missing-token)"
expect_failure "$root" ./ai/scripts/divert-task-to-test.sh \
    --authority 'Attila — Project Owner' --reason reason --disposition aborted-test \
    --release-production-id yes --token WRONG

# Successful diversion preserves evidence, releases 015, and ignores test tasks.
root="$(new_fixture successful)"
expect_success "$root" divert
expect_success "$root" ./bin/job --check
expect_success "$root" ./ai/scripts/next-task.sh --check
expect_success "$root" ./bin/job --check-test-tasks
[[ -f "$root/ai/test_tasks/001_TEST_TASK/prompt/original-prompt.md" ]]
[[ -f "$root/ai/test_tasks/001_TEST_TASK/history/original-history.md" ]]
[[ ! -e "$root/ai/prompts/015_CURRENT_TASK.md" ]]
[[ ! -e "$root/ai/history/015_${today}_interrupted-method.md" ]]
[[ "$(cat "$root/unrelated-untracked.txt")" == 'do not touch' ]]
passes=$((passes + 5))

# A clean builder input can reuse released production ID 015.
input="$(new_input reuse-015)"
expect_success "$root" ./ai/scripts/task-builder.sh --check-input "$input"
expect_success "$root" ./ai/scripts/task-builder.sh --input-dir "$input"
[[ -f "$root/ai/prompts/015_CURRENT_TASK.md" ]]
grep -Fxq -- '- Task ID: `015`' "$root/ai/prompts/015_CURRENT_TASK.md"
[[ ! -e "$root/ai/prompts/016_CURRENT_TASK.md" ]]
passes=$((passes + 3))

# Complete active tasks cannot be mislabeled as incomplete diversions.
root="$(new_fixture complete-rejected)"
sed -i 's/Status: `active`/Status: `complete`/' "$root/ai/prompts/015_CURRENT_TASK.md"
sed -i 's/Status: `active — failed attempts recorded`/Status: `complete`/' \
    "$root/ai/history/015_${today}_interrupted-method.md"
expect_failure "$root" divert

# Existing test destinations are never overwritten and numbering advances.
root="$(new_fixture next-test-id)"
expect_success "$root" divert
first_hash="$(sha256sum "$root/ai/test_tasks/001_TEST_TASK/prompt/original-prompt.md")"
cat >"$root/ai/prompts/016_CURRENT_TASK.md" <<'EOF'
# Current Engineering Task 016: Second Interrupted Method
## Task Metadata
- Task ID: `016`
- Task slug: `second-interrupted-method`
- Status: `active`
EOF
cat >"$root/ai/history/016_${today}_second-interrupted-method.md" <<'EOF'
# Task History 016: Second Interrupted Method
## Task metadata
- Task ID: `016`
- Task slug: `second-interrupted-method`
- Status: `active — incomplete`
EOF
expect_success "$root" divert
[[ -d "$root/ai/test_tasks/002_TEST_TASK" ]]
[[ "$(sha256sum "$root/ai/test_tasks/001_TEST_TASK/prompt/original-prompt.md")" == "$first_hash" ]]
expect_success "$root" ./bin/job --check-test-tasks
passes=$((passes + 2))

# Duplicate original production references are rejected deterministically.
root="$(new_fixture duplicate-original)"
expect_success "$root" divert
cp -a "$root/ai/test_tasks/001_TEST_TASK" "$root/ai/test_tasks/002_TEST_TASK"
sed -i 's/001_TEST_TASK/002_TEST_TASK/g' \
    "$root/ai/test_tasks/002_TEST_TASK/reports/DISPOSITION.md" \
    "$root/ai/test_tasks/002_TEST_TASK/DIVERSION_MANIFEST.md"
(
    cd "$root/ai/test_tasks/002_TEST_TASK"
    sha256sum prompt/original-prompt.md history/original-history.md \
        reports/DISPOSITION.md DIVERSION_MANIFEST.md >SHA256SUMS
)
expect_failure "$root" ./bin/job --check-test-tasks

# Audit rejects missing metadata and corrupted evidence.
root="$(new_fixture audit-metadata)"
expect_success "$root" divert
sed -i '/Disposition Authority/d' \
    "$root/ai/test_tasks/001_TEST_TASK/reports/DISPOSITION.md"
expect_failure "$root" ./bin/job --check-test-tasks

root="$(new_fixture audit-hash)"
expect_success "$root" divert
printf 'corruption\n' >>"$root/ai/test_tasks/001_TEST_TASK/history/original-history.md"
expect_failure "$root" ./bin/job --check-test-tasks

# Every injected transaction failure restores the exact active state.
for point in prompt history install; do
    root="$(new_fixture "rollback-$point")"
    expect_failure "$root" env QWSG_DIVERSION_TEST_FAIL_AFTER="$point" \
        ./ai/scripts/divert-task-to-test.sh \
        --authority 'Attila — Project Owner' \
        --reason 'Controlled rollback test' \
        --disposition aborted-test \
        --release-production-id yes \
        --token DIVERT-TO-TEST
    assert_original "$root"
done

printf 'PASS: %d diversion assertions\n' "$passes"
