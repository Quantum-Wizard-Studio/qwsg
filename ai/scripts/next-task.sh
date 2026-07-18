#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"
project_root="$(cd -- "$script_dir/../.." && pwd -P)"
prompts_dir="$project_root/ai/prompts"
archive_dir="$project_root/ai/archive_prompts"
history_dir="$project_root/ai/history"
core_marker="$project_root/ai/core/08_JOB_TEMPLATE.md"

fail() {
    printf 'Hiba: %s\n' "$1" >&2
    exit 1
}

if [[ "$(pwd -P)" != "$project_root" ]]; then
    fail "a szkript kizárólag a projekt gyökeréből futtatható: $project_root"
fi
[[ -f "$core_marker" ]] || fail 'a QWSG projektazonosító dokumentum hiányzik'
[[ -d "$prompts_dir" ]] || fail 'az ai/prompts könyvtár hiányzik'
[[ -d "$archive_dir" ]] || fail 'az ai/archive_prompts könyvtár hiányzik'
[[ -d "$history_dir" ]] || fail 'az ai/history könyvtár hiányzik'

mapfile -d '' active_files < <(find "$prompts_dir" -maxdepth 1 -type f -name '*.md' -print0)
active_count="${#active_files[@]}"
if (( active_count > 1 )); then
    fail 'az ai/prompts könyvtárban egynél több Markdown prompt található'
fi

current_id=''
current_slug=''
active_path=''
if (( active_count == 1 )); then
    active_path="${active_files[0]}"
    active_name="${active_path##*/}"
    if [[ ! "$active_name" =~ ^([0-9]{3})_CURRENT_TASK\.md$ ]]; then
        fail "hibás aktív promptfájlnév: $active_name"
    fi
    current_id="${BASH_REMATCH[1]}"
    mapfile -t slug_lines < <(sed -nE 's/^- Task slug: `([a-z0-9]+(-[a-z0-9]+)*)`$/\1/p' "$active_path")
    if (( ${#slug_lines[@]} != 1 )); then
        fail 'az aktív prompt pontosan egy érvényes Task slug metaadatot igényel'
    fi
    current_slug="${slug_lines[0]}"
fi

printf '%s' 'Add meg a következő feladat rövid slugját: '
IFS= read -r raw_slug || fail 'nem érkezett feladatslug'
sanitized_slug="$(printf '%s' "$raw_slug" | LC_ALL=C tr '[:upper:]' '[:lower:]' | sed -E 's/[^a-z0-9]+/-/g; s/^-+//; s/-+$//')"
[[ -n "$sanitized_slug" ]] || fail 'a slug tisztítás után üres lett'

today="$(date -u +%F)"
if [[ -n "$current_id" ]]; then
    current_number=$((10#$current_id))
    next_number=$((current_number + 1))
else
    max_number=0
    while IFS= read -r -d '' record; do
        record_name="${record##*/}"
        if [[ "$record_name" =~ ^([0-9]{3})_ ]]; then
            record_number=$((10#${BASH_REMATCH[1]}))
            (( record_number > max_number )) && max_number="$record_number"
        fi
    done < <(find "$archive_dir" "$history_dir" -maxdepth 1 -type f -name '*.md' -print0)
    next_number=$((max_number + 1))
fi
(( next_number <= 999 )) || fail 'a háromjegyű feladatazonosító-tartomány elfogyott'
printf -v next_id '%03d' "$next_number"

new_active="$prompts_dir/${next_id}_CURRENT_TASK.md"
new_history="$history_dir/${next_id}_${today}_${sanitized_slug}.md"
[[ ! -e "$new_active" ]] || fail "a következő aktív prompt már létezik: $new_active"
[[ ! -e "$new_history" ]] || fail "a következő történeti fájl már létezik: $new_history"

archive_path=''
if [[ -n "$current_id" ]]; then
    archive_path="$archive_dir/${current_id}_${today}_${current_slug}.md"
    [[ ! -e "$archive_path" ]] || fail "az archív célfájl már létezik: $archive_path"
fi

temp_prompt="$(mktemp "$prompts_dir/.next-task-prompt.XXXXXX")"
temp_history="$(mktemp "$history_dir/.next-task-history.XXXXXX")"
archive_moved='false'
new_active_moved='false'
cleanup() {
    exit_status="$?"
    [[ ! -e "$temp_prompt" ]] || rm -- "$temp_prompt"
    [[ ! -e "$temp_history" ]] || rm -- "$temp_history"
    if (( exit_status != 0 )); then
        if [[ "$new_active_moved" == 'true' && -e "$new_active" ]]; then
            rm -- "$new_active"
        fi
        if [[ "$archive_moved" == 'true' && -e "$archive_path" && ! -e "$active_path" ]]; then
            mv -T --no-clobber -- "$archive_path" "$active_path"
        fi
    fi
    trap - EXIT
    exit "$exit_status"
}
trap cleanup EXIT

cat >"$temp_prompt" <<EOF
# Current Engineering Task $next_id: [REQUIRES HUMAN EDITING]

## Task Metadata

- Task ID: \`$next_id\`
- Task slug: \`$sanitized_slug\`
- Status: \`draft — requires human editing and approval\`
- Date opened: \`$today\` UTC
- Human authority: [REQUIRES HUMAN EDITING]
- Owner or lead-developer communication language: [REQUIRES HUMAN EDITING]

## Title

[REQUIRES HUMAN EDITING]

## Objective

[REQUIRES HUMAN EDITING: define one outcome and its acceptance intent.]

## Scope

[REQUIRES HUMAN EDITING: list exact authorized files, systems, and actions.]

## Out of Scope

[REQUIRES HUMAN EDITING: list forbidden, deferred, and next-milestone work.]

## Required Reading

- \`ai/core/00_PROJECT_PHILOSOPHY.md\`
- \`ai/core/01_CONSTITUTION.md\`
- \`ai/core/03_AGENTS.md\`
- \`ai/core/08_JOB_TEMPLATE.md\`
- [REQUIRES HUMAN EDITING: add task-specific records.]

## Starting State Verification

[REQUIRES HUMAN EDITING: record environment, Git, files, ownership, permissions, ACLs, and variances.]

## Snapshot Requirements

[REQUIRES HUMAN EDITING: define timestamped snapshot, integrity check, retention, and restore guard.]

## Risk Assessment

[REQUIRES HUMAN EDITING: rate security, stability, data, compatibility, permission, and rollback risks.]

## Planned Work

[REQUIRES HUMAN EDITING: define the smallest safe sequence and decision points.]

## Rollback Plan

[REQUIRES HUMAN EDITING: define exact bounded restoration and verification.]

## Deliverables

[REQUIRES HUMAN EDITING]

## Verification

[REQUIRES HUMAN EDITING: provide checks and expected results.]

## Documentation Updates

[REQUIRES HUMAN EDITING]

## Completion Criteria

[REQUIRES HUMAN EDITING: define objective pass/fail conditions.]

## Owner Approval Requirements

This task must not begin until the owner explicitly approves the edited prompt. Scope expansion and destructive actions require separate approval.
EOF

cat >"$temp_history" <<EOF
# Task History $next_id: [REQUIRES HUMAN EDITING]

## Task ID

\`$next_id\`

## Task title

[REQUIRES HUMAN EDITING]

## Date

\`$today\` UTC

## Status

\`pending — task not started\`

## Starting state

[PENDING]

## Snapshot location

[PENDING]

## Work performed

[PENDING]

## Files changed

[PENDING]

## Decisions

[PENDING]

## Verification evidence

[PENDING]

## Problems encountered

[PENDING]

## Rollback procedure

[PENDING]

## Git commit hash

[PENDING]

## Open questions

[PENDING]

## Recommended next task

[PENDING]
EOF

[[ -s "$temp_prompt" && -s "$temp_history" ]] || fail 'az ideiglenes sablonok létrehozása sikertelen'

if [[ -n "$active_path" ]]; then
    mv -T --no-clobber -- "$active_path" "$archive_path"
    [[ -e "$archive_path" && ! -e "$active_path" ]] || fail 'az aktív prompt archiválása sikertelen'
    archive_moved='true'
fi
mv -T --no-clobber -- "$temp_prompt" "$new_active"
[[ -e "$new_active" && ! -e "$temp_prompt" ]] || fail 'az új aktív prompt létrehozása sikertelen'
new_active_moved='true'
mv -T --no-clobber -- "$temp_history" "$new_history"
[[ -e "$new_history" && ! -e "$temp_history" ]] || fail 'az új történeti fájl létrehozása sikertelen'

printf '%s\n' 'A QWSG feladatrotáció sikeresen befejeződött.'
if [[ -n "$archive_path" ]]; then
    printf 'Archivált prompt: %s\n' "$archive_path"
else
    printf '%s\n' 'Inicializálás történt: nem volt archiválandó aktív prompt.'
fi
printf 'Új aktív prompt: %s\n' "$new_active"
printf 'Új, függőben lévő történeti fájl: %s\n' "$new_history"
printf 'Következő feladatazonosító: %s\n' "$next_id"
printf '%s\n' 'A feladatot csak a sablon mezőinek kitöltése és tulajdonosi jóváhagyás után kezdd el.'
