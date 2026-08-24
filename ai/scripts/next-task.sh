#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"
project_root="$(cd -- "$script_dir/../.." && pwd -P)"
prompts_dir="$project_root/ai/prompts"
archive_dir="$project_root/ai/archive_prompts"
history_dir="$project_root/ai/history"

fail() { printf 'Hiba: %s\n' "$1" >&2; exit 1; }
usage() {
    cat <<'EOF'
Usage:
  ./ai/scripts/next-task.sh --check
  ./ai/scripts/next-task.sh --prepare --slug <next-task-slug>
  ./ai/scripts/next-task.sh

--check validates the current prompt/history lifecycle without changing files.
--prepare transactionally archives a completed task and prepares the next draft.
No-argument mode asks the owner for the next slug and then performs --prepare.
EOF
}

require_root() {
    [[ "$(pwd -P)" == "$project_root" ]] || fail "run from the project root: $project_root"
    [[ -f "$project_root/VERSION" && -f "$project_root/ai/core/08_JOB_TEMPLATE.md" ]] || fail 'QWSG project markers are missing'
    [[ -d "$prompts_dir" && -d "$archive_dir" && -d "$history_dir" ]] || fail 'lifecycle directories are missing'
    if [[ -e "$project_root/ai/framework/VERSION" ]]; then
        [[ -x "$project_root/ai/scripts/framework-check.sh" ]] ||
            fail 'framework validator is missing or not executable'
        "$project_root/ai/scripts/framework-check.sh" --quiet ||
            fail 'reusable framework validation failed'
    fi
}

extract_one() {
    local pattern="$1" file="$2" label="$3"
    mapfile -t extracted < <(sed -nE "$pattern" "$file")
    (( ${#extracted[@]} == 1 )) || fail "$label must occur exactly once in ${file#$project_root/}"
    printf '%s\n' "${extracted[0]}"
}

load_current() {
    mapfile -d '' active_files < <(find -P "$prompts_dir" -maxdepth 1 -type f -name '*.md' -print0)
    (( ${#active_files[@]} <= 1 )) || fail "at most one active Markdown prompt is allowed; found ${#active_files[@]}"
    current_is_archived=false
    if (( ${#active_files[@]} == 0 )); then
        mapfile -t archived_files < <(find -P "$archive_dir" -maxdepth 1 -type f -name '[0-9][0-9][0-9]_????-??-??_*.md' -print | LC_ALL=C sort)
        (( ${#archived_files[@]} > 0 )) || fail 'no active or archived task prompt exists'
        active_path="${archived_files[${#archived_files[@]}-1]}"
        current_is_archived=true
    else
        active_path="${active_files[0]}"
    fi
    active_name="${active_path##*/}"
    if [[ "$current_is_archived" == true ]]; then
        [[ "$active_name" =~ ^([0-9]{3})_[0-9]{4}-[0-9]{2}-[0-9]{2}_[a-z0-9]+(-[a-z0-9]+)*\.md$ ]] || fail "invalid archived prompt filename: $active_name"
    else
        [[ "$active_name" =~ ^([0-9]{3})_CURRENT_TASK\.md$ ]] || fail "invalid active prompt filename: $active_name"
    fi
    current_id="${BASH_REMATCH[1]}"
    prompt_id="$(extract_one 's/^- Task ID: `([0-9]{3})`$/\1/p' "$active_path" 'prompt Task ID')"
    [[ "$prompt_id" == "$current_id" ]] || fail 'prompt filename and metadata Task IDs differ'
    current_slug="$(extract_one 's/^- Task slug: `([a-z0-9]+(-[a-z0-9]+)*)`$/\1/p' "$active_path" 'prompt Task slug')"
    prompt_status="$(extract_one 's/^- Status: `?([^`]+)`?$/\1/p' "$active_path" 'prompt Status')"

    mapfile -d '' history_files < <(find -P "$history_dir" -maxdepth 1 -type f -name "${current_id}_????-??-??_*.md" -print0)
    (( ${#history_files[@]} == 1 )) || fail "exactly one Task $current_id history is required; found ${#history_files[@]}"
    history_path="${history_files[0]}"
    history_name="${history_path##*/}"
    [[ "$history_name" =~ ^${current_id}_[0-9]{4}-[0-9]{2}-[0-9]{2}_${current_slug}\.md$ ]] || fail 'history filename does not match prompt Task ID and slug'
    history_id="$(extract_one 's/^- Task ID: `([0-9]{3})`$/\1/p' "$history_path" 'history Task ID')"
    history_slug="$(extract_one 's/^- Task slug: `([a-z0-9]+(-[a-z0-9]+)*)`$/\1/p' "$history_path" 'history Task slug')"
    history_status="$(extract_one 's/^- Status: `?([^`]+)`?$/\1/p' "$history_path" 'history Status')"
    [[ "$history_id" == "$current_id" && "$history_slug" == "$current_slug" ]] || fail 'prompt/history metadata mismatch'
    if find "$prompts_dir" "$history_dir" -maxdepth 1 -type f \( -name '.next-task-*' -o -name '*.lifecycle-tmp' \) -print -quit | grep -q .; then
        fail 'stale lifecycle temporary file exists'
    fi
}

validate_current() {
    load_current
    if [[ "$current_is_archived" == true ]]; then
        require_completed
        printf 'Lifecycle valid: no active task; latest completed Task %s, prompt=%s, history=%s\n' "$current_id" "${active_path#$project_root/}" "${history_path#$project_root/}"
    else
        printf 'Lifecycle valid: Task %s (%s), prompt=%s, history=%s\n' "$current_id" "$prompt_status" "${active_path#$project_root/}" "${history_path#$project_root/}"
    fi
}

require_completed() {
    [[ "$prompt_status" == complete* ]] || fail "current prompt is not complete: $prompt_status"
    [[ "$history_status" == complete* ]] || fail "current history is not finalized: $history_status"
    if grep -Eq '\[PENDING\]|\[REQUIRES HUMAN EDITING\]|pending — task not started|pending - task not started' "$history_path"; then
        fail 'current history still contains pending or human-editing fields'
    fi
    for heading in 'Starting state' 'Snapshot' 'Work performed' 'Verification' 'Rollback' 'Completion state'; do
        grep -Eiq "^## .*${heading}" "$history_path" || fail "current history lacks completion evidence section: $heading"
    done
}

sanitize_slug() {
    printf '%s' "$1" | LC_ALL=C tr '[:upper:]' '[:lower:]' | sed -E 's/[^a-z0-9]+/-/g; s/^-+//; s/-+$//'
}

prepare_next() {
    local raw_slug="$1" next_slug next_number next_id today new_active new_history archive_path
    load_current
    require_completed
    next_slug="$(sanitize_slug "$raw_slug")"
    [[ -n "$next_slug" ]] || fail 'next-task slug is empty after normalization'
    [[ "$next_slug" == "$raw_slug" ]] || fail 'non-interactive slug must already use lowercase kebab-case'
    next_number=$((10#$current_id + 1)); (( next_number <= 999 )) || fail 'three-digit Task ID range exhausted'
    printf -v next_id '%03d' "$next_number"
    today="$(date -u +%F)"
    new_active="$prompts_dir/${next_id}_CURRENT_TASK.md"
    new_history="$history_dir/${next_id}_${today}_${next_slug}.md"
    archive_path="$archive_dir/${current_id}_${today}_${current_slug}.md"
    [[ ! -e "$new_active" ]] || fail "next prompt conflicts: ${new_active#$project_root/}"
    [[ ! -e "$new_history" ]] || fail "next history conflicts: ${new_history#$project_root/}"
    if [[ "$current_is_archived" == false ]]; then
        [[ ! -e "$archive_path" ]] || fail "archive destination conflicts: ${archive_path#$project_root/}"
    fi

    local temp_prompt temp_history archive_moved=false prompt_moved=false history_moved=false
    temp_prompt="$(mktemp "$prompts_dir/.next-task-prompt.XXXXXX")"
    temp_history="$(mktemp "$history_dir/.next-task-history.XXXXXX")"
    cleanup() {
        local status="$?"
        [[ ! -e "$temp_prompt" ]] || rm -- "$temp_prompt"
        [[ ! -e "$temp_history" ]] || rm -- "$temp_history"
        if (( status != 0 )); then
            [[ "$history_moved" != true || ! -e "$new_history" ]] || rm -- "$new_history"
            [[ "$prompt_moved" != true || ! -e "$new_active" ]] || rm -- "$new_active"
            if [[ "$archive_moved" == true && -e "$archive_path" && ! -e "$active_path" ]]; then mv -T --no-clobber -- "$archive_path" "$active_path"; fi
        fi
        trap - EXIT
        exit "$status"
    }
    trap cleanup EXIT

    sed "s/@TASK_ID@/$next_id/g; s/@TASK_SLUG@/$next_slug/g; s/@DATE@/$today/g" >"$temp_prompt" <<'EOF'
# Current Engineering Task @TASK_ID@: [REQUIRES HUMAN EDITING]

## Task Metadata

- Task ID: `@TASK_ID@`
- Task slug: `@TASK_SLUG@`
- Status: `draft — requires human editing and approval`
- Date opened: `@DATE@` UTC
- Human authority: [REQUIRES HUMAN EDITING]
- Owner or lead-developer communication language: [REQUIRES HUMAN EDITING]

## Title

[REQUIRES HUMAN EDITING]

## Objective

[REQUIRES HUMAN EDITING]

## Scope

[REQUIRES HUMAN EDITING]

## Out of Scope

[REQUIRES HUMAN EDITING]

## Authority Envelope

[REQUIRES HUMAN EDITING]

## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

[REQUIRES HUMAN EDITING]

## Snapshot Requirements

[REQUIRES HUMAN EDITING]

## Risk Assessment

[REQUIRES HUMAN EDITING]

## Planned Work

[REQUIRES HUMAN EDITING]

## Rollback Plan

[REQUIRES HUMAN EDITING]

## Deliverables

[REQUIRES HUMAN EDITING]

## Verification

[REQUIRES HUMAN EDITING]

## Documentation Updates

[REQUIRES HUMAN EDITING]

## Completion Criteria

[REQUIRES HUMAN EDITING]

## Owner Approval Requirements

This prepared task is not approved and must not begin until the Project Owner edits and explicitly approves it.
EOF
    sed "s/@TASK_ID@/$next_id/g; s/@TASK_SLUG@/$next_slug/g; s/@DATE@/$today/g" >"$temp_history" <<'EOF'
# Task History @TASK_ID@: Prepared Draft

## Task metadata

- Task ID: `@TASK_ID@`
- Task slug: `@TASK_SLUG@`
- Status: `prepared — not approved, task not started`
- Date prepared: `@DATE@` UTC

## Lifecycle state

This record was transactionally prepared with its matching prompt. The task has not started, is not approved, and contains no implementation evidence.

## Starting state

[PENDING UNTIL OWNER APPROVAL]

## Snapshot

[PENDING UNTIL OWNER APPROVAL]

## Work performed

None. Task not started.

## Verification

Preparation metadata only; implementation verification has not started.

## Rollback

[PENDING UNTIL OWNER APPROVAL]

## Completion state

`prepared — READY FOR OWNER REVIEW`
EOF
    [[ -s "$temp_prompt" && -s "$temp_history" ]] || fail 'temporary scaffolds are empty'
    if [[ "$current_is_archived" == false ]]; then
        mv -T --no-clobber -- "$active_path" "$archive_path"; archive_moved=true
        [[ "${QWSG_TEST_FAIL_AFTER:-}" != archive ]] || fail 'injected failure after archive'
    fi
    mv -T --no-clobber -- "$temp_prompt" "$new_active"; prompt_moved=true
    [[ "${QWSG_TEST_FAIL_AFTER:-}" != prompt ]] || fail 'injected failure after prompt install'
    mv -T --no-clobber -- "$temp_history" "$new_history"; history_moved=true
    [[ "${QWSG_TEST_FAIL_AFTER:-}" != history ]] || fail 'injected failure after history install'
    "$project_root/bin/job" --prepared-check >/dev/null || fail 'post-rotation prepared-state validation failed'
    trap - EXIT
    printf 'Completed Task ID: %s\nArchived prompt: %s\nNew Task ID: %s\nNew prompt: %s\nNew history: %s\nLifecycle validation: PASS\nREADY FOR OWNER REVIEW\n' "$current_id" "${archive_path#$project_root/}" "$next_id" "${new_active#$project_root/}" "${new_history#$project_root/}"
}

require_root
case "${1:-}" in
    --check) (( $# == 1 )) || fail '--check accepts no other arguments'; validate_current ;;
    --prepare) [[ "${2:-}" == --slug && -n "${3:-}" && $# == 3 ]] || fail 'use --prepare --slug <next-task-slug>'; prepare_next "$3" ;;
    --help|-h) usage ;;
    '') printf '%s' 'Add meg a következő feladat rövid slugját: '; IFS= read -r slug || fail 'no slug received'; prepare_next "$(sanitize_slug "$slug")" ;;
    *) fail "unknown option: $1" ;;
esac
