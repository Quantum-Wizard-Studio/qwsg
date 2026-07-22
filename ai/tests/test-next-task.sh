#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/../.." && pwd -P)"
work="$(mktemp -d /tmp/qwsg-lifecycle-test.XXXXXX)"
trap 'rm -rf -- "$work"' EXIT
passes=0
today="$(date -u +%F)"

new_fixture() {
    local name="$1" root
    root="$work/$name"
    mkdir -p "$root/ai/scripts" "$root/ai/prompts" "$root/ai/archive_prompts" "$root/ai/history" "$root/ai/core" "$root/bin"
    cp "$repo_root/ai/scripts/next-task.sh" "$root/ai/scripts/next-task.sh"
    cp "$repo_root/bin/job" "$root/bin/job"
    chmod +x "$root/ai/scripts/next-task.sh" "$root/bin/job"
    printf 'test\n' >"$root/VERSION"
    printf '# marker\n' >"$root/ai/core/00_PROJECT_PHILOSOPHY.md"
    printf '# marker\n' >"$root/ai/core/08_JOB_TEMPLATE.md"
    cat >"$root/ai/prompts/010_CURRENT_TASK.md" <<'EOF'
# Current Engineering Task 010: Test
## Task Metadata
- Task ID: `010`
- Task slug: `automated-task-lifecycle`
- Status: `complete`
EOF
    cat >"$root/ai/history/010_${today}_automated-task-lifecycle.md" <<'EOF'
# Task History 010: Test
## Task metadata
- Task ID: `010`
- Task slug: `automated-task-lifecycle`
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
    printf '%s\n' "$root"
}

expect_success() { local root="$1"; shift; (cd "$root" && "$@") || { printf 'FAIL expected success: %s\n' "$*" >&2; exit 1; }; passes=$((passes+1)); }
expect_failure() { local root="$1"; shift; if (cd "$root" && "$@") >/dev/null 2>&1; then printf 'FAIL expected failure: %s\n' "$*" >&2; exit 1; fi; passes=$((passes+1)); }
assert_original() { local root="$1"; [[ -f "$root/ai/prompts/010_CURRENT_TASK.md" && ! -e "$root/ai/prompts/011_CURRENT_TASK.md" && ! -e "$root/ai/archive_prompts/010_${today}_automated-task-lifecycle.md" ]] || { printf 'FAIL rollback state: %s\n' "$root" >&2; exit 1; }; ! find "$root/ai" -type f -name '.next-task-*' -print -quit | grep -q . || { printf 'FAIL orphan temporary file\n' >&2; exit 1; }; passes=$((passes+1)); }

root="$(new_fixture check)"; expect_success "$root" ./ai/scripts/next-task.sh --check

root="$(new_fixture idle-check)"; mv "$root/ai/prompts/010_CURRENT_TASK.md" "$root/ai/archive_prompts/010_${today}_automated-task-lifecycle.md"
expect_success "$root" ./ai/scripts/next-task.sh --check
expect_success "$root" ./bin/job --check
expect_failure "$root" ./bin/job --path

root="$(new_fixture idle-prepare)"; mv "$root/ai/prompts/010_CURRENT_TASK.md" "$root/ai/archive_prompts/010_${today}_automated-task-lifecycle.md"
expect_success "$root" ./ai/scripts/next-task.sh --prepare --slug core-alpha-platform-hardening
expect_success "$root" ./bin/job --prepared-check

root="$(new_fixture noninteractive)"; expect_success "$root" ./ai/scripts/next-task.sh --prepare --slug core-alpha-platform-hardening
expect_success "$root" ./bin/job --prepared-check
grep -q 'Task ID: `011`' "$root/ai/prompts/011_CURRENT_TASK.md"; grep -q 'Task slug: `core-alpha-platform-hardening`' "$root/ai/history/011_${today}_core-alpha-platform-hardening.md"; passes=$((passes+1))

root="$(new_fixture interactive)"; expect_success "$root" bash -c "printf 'Interactive Slug\\n' | ./ai/scripts/next-task.sh"

root="$(new_fixture missing-prompt)"; rm "$root/ai/prompts/010_CURRENT_TASK.md"; expect_failure "$root" ./ai/scripts/next-task.sh --check
root="$(new_fixture malformed-name)"; mv "$root/ai/prompts/010_CURRENT_TASK.md" "$root/ai/prompts/bad.md"; expect_failure "$root" ./ai/scripts/next-task.sh --check
root="$(new_fixture missing-slug)"; sed -i '/Task slug/d' "$root/ai/prompts/010_CURRENT_TASK.md"; expect_failure "$root" ./ai/scripts/next-task.sh --check
root="$(new_fixture missing-history)"; rm "$root/ai/history/010_${today}_automated-task-lifecycle.md"; expect_failure "$root" ./ai/scripts/next-task.sh --check
root="$(new_fixture duplicate-history)"; cp "$root/ai/history/010_${today}_automated-task-lifecycle.md" "$root/ai/history/010_2026-07-22_duplicate.md"; expect_failure "$root" ./ai/scripts/next-task.sh --check
root="$(new_fixture pending-history)"; sed -i 's/Status: `complete`/Status: `pending — task not started`/' "$root/ai/history/010_${today}_automated-task-lifecycle.md"; expect_failure "$root" ./ai/scripts/next-task.sh --prepare --slug next-task
root="$(new_fixture archive-conflict)"; touch "$root/ai/archive_prompts/010_${today}_automated-task-lifecycle.md"; expect_failure "$root" ./ai/scripts/next-task.sh --prepare --slug next-task
root="$(new_fixture prompt-conflict)"; touch "$root/ai/prompts/011_CURRENT_TASK.md"; expect_failure "$root" ./ai/scripts/next-task.sh --prepare --slug next-task
root="$(new_fixture history-conflict)"; touch "$root/ai/history/011_${today}_next-task.md"; expect_failure "$root" ./ai/scripts/next-task.sh --prepare --slug next-task
root="$(new_fixture malformed-id)"; sed -i 's/Task ID: `010`/Task ID: `10`/' "$root/ai/prompts/010_CURRENT_TASK.md"; expect_failure "$root" ./ai/scripts/next-task.sh --check
root="$(new_fixture id-mismatch)"; sed -i 's/Task ID: `010`/Task ID: `009`/' "$root/ai/history/010_${today}_automated-task-lifecycle.md"; expect_failure "$root" ./ai/scripts/next-task.sh --check
root="$(new_fixture slug-mismatch)"; sed -i 's/automated-task-lifecycle/other-task/' "$root/ai/history/010_${today}_automated-task-lifecycle.md"; expect_failure "$root" ./ai/scripts/next-task.sh --check

for point in archive prompt history; do
    root="$(new_fixture "fail-$point")"
    expect_failure "$root" env QWSG_TEST_FAIL_AFTER="$point" ./ai/scripts/next-task.sh --prepare --slug next-task
    assert_original "$root"
done

printf 'PASS: %d lifecycle assertions\n' "$passes"
