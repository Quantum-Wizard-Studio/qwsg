#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/../.." && pwd -P)"
work="$(mktemp -d /tmp/qwsg-task-builder-test.XXXXXX)"
trap 'rm -rf -- "$work"' EXIT
passes=0
today="$(date -u +%F)"

new_fixture() {
    local name="$1" root="$work/$1"
    mkdir -p "$root/ai/scripts" "$root/ai/prompts" "$root/ai/archive_prompts" "$root/ai/history" "$root/ai/core" "$root/bin"
    cp "$repo_root/ai/scripts/task-builder.sh" "$repo_root/ai/scripts/next-task.sh" "$root/ai/scripts/"
    cp "$repo_root/bin/job" "$root/bin/job"
    chmod +x "$root/ai/scripts/"*.sh "$root/bin/job"
    printf 'test\n' >"$root/VERSION"
    printf '# marker\n' >"$root/ai/core/00_PROJECT_PHILOSOPHY.md"
    printf '# marker\n' >"$root/ai/core/08_JOB_TEMPLATE.md"
    cat >"$root/ai/prompts/013_CURRENT_TASK.md" <<'EOF'
# Current Engineering Task 013: Builder
## Task Metadata
- Task ID: `013`
- Task slug: `engineering-task-builder`
- Status: `complete`
EOF
    cat >"$root/ai/history/013_${today}_engineering-task-builder.md" <<'EOF'
# Task History 013: Builder
## Task metadata
- Task ID: `013`
- Task slug: `engineering-task-builder`
- Status: `complete`
EOF
    printf '%s\n' "$root"
}

new_input() {
    local name="$1" approval="${2:-APPROVE}" dir
    dir="$work/input-$name"
    mkdir -p "$dir"
    printf 'next-engineering-task\n' >"$dir/slug"
    printf 'Next Engineering Task\n' >"$dir/title"
    printf 'Project Owner\n' >"$dir/authority"
    printf 'Hungarian\n' >"$dir/language"
    printf '%s\n' "$approval" >"$dir/approval"
    local field
    for field in objective scope out-of-scope starting-state snapshot risk planned-work rollback deliverables verification documentation completion; do
        printf 'First %s line.\nSecond %s line.\n' "$field" "$field" >"$dir/$field"
    done
    printf '%s\n' "$dir"
}

expect_success() {
    local root="$1"; shift
    (cd "$root" && "$@") >/dev/null || { printf 'FAIL expected success: %s\n' "$*" >&2; exit 1; }
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
    [[ -f "$root/ai/prompts/013_CURRENT_TASK.md" && ! -e "$root/ai/prompts/014_CURRENT_TASK.md" ]] || { printf 'FAIL prompt rollback\n' >&2; exit 1; }
    [[ ! -e "$root/ai/archive_prompts/013_${today}_engineering-task-builder.md" ]] || { printf 'FAIL archive rollback\n' >&2; exit 1; }
    [[ ! -e "$root/ai/history/014_${today}_next-engineering-task.md" ]] || { printf 'FAIL history rollback\n' >&2; exit 1; }
    ! find "$root/ai" -type f -name '.task-builder-*' -print -quit | grep -q . || { printf 'FAIL temporary file cleanup\n' >&2; exit 1; }
    passes=$((passes + 1))
}

root="$(new_fixture valid-input)"; input="$(new_input valid-input)"
expect_success "$root" ./ai/scripts/task-builder.sh --check-input "$input"
assert_original "$root"

root="$(new_fixture install)"; input="$(new_input install)"
expect_success "$root" ./ai/scripts/task-builder.sh --input-dir "$input"
expect_success "$root" ./bin/job --check
expect_success "$root" ./ai/scripts/next-task.sh --check
grep -Fq -- '- Status: `approved`' "$root/ai/prompts/014_CURRENT_TASK.md"
grep -Fq 'First objective line.' "$root/ai/prompts/014_CURRENT_TASK.md"
grep -Fq 'Second objective line.' "$root/ai/prompts/014_CURRENT_TASK.md"
grep -Fq 'Approved by Project Owner' "$root/ai/prompts/014_CURRENT_TASK.md"
grep -Fq -- '- Status: `approved — task not started`' "$root/ai/history/014_${today}_next-engineering-task.md"
passes=$((passes + 5))

root="$(new_fixture idle-install)"; input="$(new_input idle-install)"
mv "$root/ai/prompts/013_CURRENT_TASK.md" "$root/ai/archive_prompts/013_${today}_engineering-task-builder.md"
expect_success "$root" ./ai/scripts/task-builder.sh --input-dir "$input"
expect_success "$root" ./bin/job --check
[[ -f "$root/ai/archive_prompts/013_${today}_engineering-task-builder.md" && -f "$root/ai/prompts/014_CURRENT_TASK.md" ]] || { printf 'FAIL idle builder state\n' >&2; exit 1; }
passes=$((passes + 1))

root="$(new_fixture interactive)"
interactive_input="$work/interactive-input"
{
    printf 'interactive-task\nInteractive Task\nProject Owner\nHungarian\n'
    for field in objective scope out-of-scope starting-state snapshot risk planned-work rollback deliverables verification documentation completion; do
        printf 'Interactive %s line one.\nInteractive %s line two.\n.\n' "$field" "$field"
    done
    printf 'APPROVE\n'
} >"$interactive_input"
expect_success "$root" bash -c './ai/scripts/task-builder.sh <"$1"' _ "$interactive_input"
grep -Fq 'Interactive objective line one.' "$root/ai/prompts/014_CURRENT_TASK.md"
grep -Fq 'Interactive objective line two.' "$root/ai/prompts/014_CURRENT_TASK.md"
passes=$((passes + 2))

root="$(new_fixture no-approval)"; input="$(new_input no-approval NO)"
expect_failure "$root" ./ai/scripts/task-builder.sh --input-dir "$input"
assert_original "$root"

root="$(new_fixture missing-field)"; input="$(new_input missing-field)"; rm "$input/scope"
expect_failure "$root" ./ai/scripts/task-builder.sh --input-dir "$input"
assert_original "$root"

root="$(new_fixture bad-slug)"; input="$(new_input bad-slug)"; printf 'Bad Slug\n' >"$input/slug"
expect_failure "$root" ./ai/scripts/task-builder.sh --input-dir "$input"
assert_original "$root"

root="$(new_fixture marker)"; input="$(new_input marker)"; printf '[REQUIRES HUMAN EDITING]\n' >"$input/objective"
expect_failure "$root" ./ai/scripts/task-builder.sh --input-dir "$input"
assert_original "$root"

root="$(new_fixture active-current)"; input="$(new_input active-current)"; sed -i 's/Status: `complete`/Status: `active`/' "$root/ai/prompts/013_CURRENT_TASK.md"
expect_failure "$root" ./ai/scripts/task-builder.sh --input-dir "$input"
assert_original "$root"

for point in archive prompt history; do
    root="$(new_fixture "rollback-$point")"; input="$(new_input "rollback-$point")"
    expect_failure "$root" env QWSG_BUILDER_TEST_FAIL_AFTER="$point" ./ai/scripts/task-builder.sh --input-dir "$input"
    assert_original "$root"
done

root_a="$(new_fixture deterministic-a)"; root_b="$(new_fixture deterministic-b)"; input="$(new_input deterministic)"
expect_success "$root_a" ./ai/scripts/task-builder.sh --input-dir "$input"
expect_success "$root_b" ./ai/scripts/task-builder.sh --input-dir "$input"
cmp "$root_a/ai/prompts/014_CURRENT_TASK.md" "$root_b/ai/prompts/014_CURRENT_TASK.md"
cmp "$root_a/ai/history/014_${today}_next-engineering-task.md" "$root_b/ai/history/014_${today}_next-engineering-task.md"
passes=$((passes + 2))

printf 'PASS: %d task-builder assertions\n' "$passes"
