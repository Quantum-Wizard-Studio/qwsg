#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"
project_root="$(cd -- "$script_dir/../.." && pwd -P)"
prompts_dir="$project_root/ai/prompts"
archive_dir="$project_root/ai/archive_prompts"
history_dir="$project_root/ai/history"

fail() { printf 'Error: %s\n' "$1" >&2; exit 1; }
usage() {
    cat <<'EOF'
Usage:
  ./ai/scripts/task-builder.sh
  ./ai/scripts/task-builder.sh --input-dir <directory>
  ./ai/scripts/task-builder.sh --check-input <directory>

Interactive mode collects structured owner input. Multi-line fields end with a line
containing only a single period. --input-dir reads the same fields from individual
UTF-8 text files and installs the generated approved task transactionally.
--check-input validates and renders nothing; it never changes lifecycle files.

Required input files:
  slug title authority language objective scope out-of-scope authority-envelope
  starting-state snapshot risk planned-work rollback deliverables verification
  documentation completion approval

The approval file must contain exactly APPROVE (apart from a final newline).
EOF
}

require_root() {
    [[ "$(pwd -P)" == "$project_root" ]] || fail "run from the project root: $project_root"
    [[ -f "$project_root/VERSION" && -x "$project_root/bin/job" ]] || fail 'QWSG project markers are missing'
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
    current_slug="$(extract_one 's/^- Task slug: `([a-z0-9]+(-[a-z0-9]+)*)`$/\1/p' "$active_path" 'prompt Task slug')"
    prompt_status="$(extract_one 's/^- Status: `?([^`]+)`?$/\1/p' "$active_path" 'prompt Status')"
    [[ "$prompt_id" == "$current_id" ]] || fail 'active prompt filename and Task ID differ'
    [[ "$prompt_status" == complete* ]] || fail "current prompt is not complete: $prompt_status"

    mapfile -d '' current_histories < <(find -P "$history_dir" -maxdepth 1 -type f -name "${current_id}_????-??-??_*.md" -print0)
    (( ${#current_histories[@]} == 1 )) || fail "exactly one current history is required; found ${#current_histories[@]}"
    current_history="${current_histories[0]}"
    history_status="$(extract_one 's/^- Status: `?([^`]+)`?$/\1/p' "$current_history" 'history Status')"
    [[ "$history_status" == complete* ]] || fail "current history is not complete: $history_status"
}

readonly fields=(slug title authority language objective scope out-of-scope authority-envelope starting-state snapshot risk planned-work rollback deliverables verification documentation completion approval)

validate_input_dir() {
    input_dir="$(cd -- "$1" 2>/dev/null && pwd -P)" || fail "input directory is unavailable: $1"
    local field path
    for field in "${fields[@]}"; do
        path="$input_dir/$field"
        [[ -f "$path" && ! -L "$path" ]] || fail "required input is not a regular file: $field"
        [[ -s "$path" ]] || fail "required input is empty: $field"
        grep -Fq '[REQUIRES HUMAN EDITING]' "$path" && fail "unresolved editing marker in input: $field" || true
    done
    read_single_value slug slug
    [[ "$slug" =~ ^[a-z0-9]+(-[a-z0-9]+)*$ ]] || fail 'slug must use lowercase kebab-case'
    read_single_value title title
    read_single_value authority authority
    read_single_value language language
    read_single_value approval approval
    [[ -n "$title" && -n "$authority" && -n "$language" ]] || fail 'title, authority, and language must be single non-empty lines'
    [[ "$approval" == APPROVE ]] || fail 'explicit owner approval is required: approval must be exactly APPROVE'
    local label
    for label in 'Authorized paths/components/systems' 'Routine operations' \
        'Correction/retest authority' 'Repository integration' \
        'Lifecycle completion' 'Permitted external actions' \
        'Evidence and rollback' 'Owner-reserved operations' \
        'Mandatory STOP conditions'; do
        grep -Fq "**$label:**" "$input_dir/authority-envelope" ||
            fail "authority-envelope lacks required category: $label"
    done
}

read_single_value() {
    local variable="$1" field="$2" values=()
    mapfile -t values <"$input_dir/$field"
    (( ${#values[@]} == 1 )) || fail "$field must contain exactly one line"
    values[0]="${values[0]%$'\r'}"
    printf -v "$variable" '%s' "${values[0]}"
}

read_line_field() {
    local label="$1" target="$2" value
    printf '%s: ' "$label"
    IFS= read -r value || fail "no input received for $target"
    [[ -n "$value" ]] || fail "$target must not be empty"
    printf '%s\n' "$value" >"$interactive_dir/$target"
}

read_multiline_field() {
    local label="$1" target="$2" line count=0
    printf '%s (finish with a single period):\n' "$label"
    : >"$interactive_dir/$target"
    while IFS= read -r line; do
        [[ "$line" != . ]] || break
        printf '%s\n' "$line" >>"$interactive_dir/$target"
        count=$((count + 1))
    done
    (( count > 0 )) || fail "$target must not be empty"
}

collect_interactive() {
    interactive_dir="$(mktemp -d /tmp/qwsg-task-builder-input.XXXXXX)"
    read_line_field 'Task slug (lowercase kebab-case)' slug
    read_line_field 'English task title' title
    read_line_field 'Human authority' authority
    read_line_field 'Preferred owner communication language' language
    local field label
    for field in objective scope out-of-scope authority-envelope starting-state snapshot risk planned-work rollback deliverables verification documentation completion; do
        label="${field//-/ }"
        read_multiline_field "$label" "$field"
    done
    printf '%s' 'Type APPROVE to approve and install the generated task: '
    IFS= read -r approval || fail 'no approval received'
    printf '%s\n' "$approval" >"$interactive_dir/approval"
    validate_input_dir "$interactive_dir"
}

append_section() {
    local heading="$1" source="$2"
    printf '\n## %s\n\n' "$heading" >>"$temp_prompt"
    cat -- "$input_dir/$source" >>"$temp_prompt"
    printf '\n' >>"$temp_prompt"
}

render_pair() {
    local next_number
    next_number=$((10#$current_id + 1)); (( next_number <= 999 )) || fail 'three-digit Task ID range exhausted'
    printf -v next_id '%03d' "$next_number"
    today="$(date -u +%F)"
    new_prompt="$prompts_dir/${next_id}_CURRENT_TASK.md"
    new_history="$history_dir/${next_id}_${today}_${slug}.md"
    archive_path="$archive_dir/${current_id}_${today}_${current_slug}.md"
    [[ ! -e "$new_prompt" && ! -e "$new_history" ]] || fail 'a lifecycle destination already exists'
    if [[ "$current_is_archived" == false ]]; then
        [[ ! -e "$archive_path" ]] || fail 'a lifecycle archive destination already exists'
    fi
    temp_prompt="$(mktemp "$prompts_dir/.task-builder-prompt.XXXXXX")"
    temp_history="$(mktemp "$history_dir/.task-builder-history.XXXXXX")"

    cat >"$temp_prompt" <<EOF
# Current Engineering Task $next_id: $title

## Task Metadata

- Task ID: \`$next_id\`
- Task slug: \`$slug\`
- Status: \`approved\`
- Date opened: \`$today\` UTC
- Human authority: $authority
- Owner or lead-developer communication language: $language
EOF
    append_section 'Title' title
    append_section 'Objective' objective
    append_section 'Scope' scope
    append_section 'Out of Scope' out-of-scope
    append_section 'Authority Envelope' authority-envelope
    cat >>"$temp_prompt" <<'EOF'

## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/config/engineering-project.conf`
EOF
    append_section 'Starting State Verification' starting-state
    append_section 'Snapshot Requirements' snapshot
    append_section 'Risk Assessment' risk
    append_section 'Planned Work' planned-work
    append_section 'Rollback Plan' rollback
    append_section 'Deliverables' deliverables
    append_section 'Verification' verification
    append_section 'Documentation Updates' documentation
    append_section 'Completion Criteria' completion
    cat >>"$temp_prompt" <<EOF

## Owner Approval Requirements

Approved by $authority through the Engineering Task Builder on $today UTC.

The structured task definition and Authority Envelope have been explicitly approved. The task is authorized to start and execute every routine operation inside that envelope without another Owner gate. Further scope changes and every Owner-reserved operation require explicit Project Owner approval.
EOF

    cat >"$temp_history" <<EOF
# Task History $next_id: Approved Task

## Task metadata

- Task ID: \`$next_id\`
- Task slug: \`$slug\`
- Status: \`approved — task not started\`
- Date generated: \`$today\` UTC
- Human authority: $authority
- Preferred owner communication language: $language
- Related prompt: \`ai/prompts/${next_id}_CURRENT_TASK.md\`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded; implementation has not started.

## Starting state

[PENDING UNTIL TASK START]

## Snapshot

[PENDING UNTIL TASK START]

## Work performed

None. Approved task not started.

## Verification

Builder input, metadata, prompt/history identity, approval state, and lifecycle installation validated successfully.

## Rollback

[PENDING UNTIL TASK START]

## Completion state

\`approved — task not started\`
EOF
    validate_rendered_pair
}

validate_rendered_pair() {
    [[ -s "$temp_prompt" && -s "$temp_history" ]] || fail 'generated document is empty'
    ! grep -Fq '[REQUIRES HUMAN EDITING]' "$temp_prompt" || fail 'generated prompt contains an editing marker'
    grep -Fxq -- "- Task ID: \`$next_id\`" "$temp_prompt" || fail 'generated prompt Task ID is invalid'
    grep -Fxq -- "- Task slug: \`$slug\`" "$temp_prompt" || fail 'generated prompt slug is invalid'
    grep -Fxq -- '- Status: `approved`' "$temp_prompt" || fail 'generated prompt approval state is invalid'
    grep -Fxq -- "- Task ID: \`$next_id\`" "$temp_history" || fail 'generated history Task ID is invalid'
    grep -Fxq -- "- Task slug: \`$slug\`" "$temp_history" || fail 'generated history slug is invalid'
    local heading
    for heading in 'Objective' 'Scope' 'Out of Scope' 'Authority Envelope' 'Required Reading' 'Starting State Verification' 'Snapshot Requirements' 'Risk Assessment' 'Planned Work' 'Rollback Plan' 'Deliverables' 'Verification' 'Documentation Updates' 'Completion Criteria' 'Owner Approval Requirements'; do
        grep -Fxq "## $heading" "$temp_prompt" || fail "generated prompt lacks section: $heading"
    done
}

install_pair() {
    local archive_moved=false prompt_moved=false history_moved=false
    cleanup_transaction() {
        local status="$?"
        [[ ! -e "$temp_prompt" ]] || rm -- "$temp_prompt"
        [[ ! -e "$temp_history" ]] || rm -- "$temp_history"
        if (( status != 0 )); then
            [[ "$history_moved" != true || ! -e "$new_history" ]] || rm -- "$new_history"
            [[ "$prompt_moved" != true || ! -e "$new_prompt" ]] || rm -- "$new_prompt"
            if [[ "$archive_moved" == true && -e "$archive_path" && ! -e "$active_path" ]]; then
                mv -T --no-clobber -- "$archive_path" "$active_path"
            fi
        fi
        [[ -z "${interactive_dir:-}" || ! -d "$interactive_dir" ]] || rm -rf -- "$interactive_dir"
        trap - EXIT
        exit "$status"
    }
    trap cleanup_transaction EXIT

    if [[ "$current_is_archived" == false ]]; then
        mv -T --no-clobber -- "$active_path" "$archive_path"; archive_moved=true
        [[ "${QWSG_BUILDER_TEST_FAIL_AFTER:-}" != archive ]] || fail 'injected failure after archive'
    fi
    mv -T --no-clobber -- "$temp_prompt" "$new_prompt"; prompt_moved=true
    [[ "${QWSG_BUILDER_TEST_FAIL_AFTER:-}" != prompt ]] || fail 'injected failure after prompt install'
    mv -T --no-clobber -- "$temp_history" "$new_history"; history_moved=true
    [[ "${QWSG_BUILDER_TEST_FAIL_AFTER:-}" != history ]] || fail 'injected failure after history install'
    "$project_root/bin/job" --check >/dev/null || fail 'installed task failed executable lifecycle validation'
    "$project_root/ai/scripts/next-task.sh" --check >/dev/null || fail 'installed task failed lifecycle consistency validation'
    trap - EXIT
    [[ -z "${interactive_dir:-}" || ! -d "$interactive_dir" ]] || rm -rf -- "$interactive_dir"
    printf 'Completed Task ID: %s\nArchived prompt: %s\nNew approved Task ID: %s\nNew prompt: %s\nNew history: %s\nLifecycle validation: PASS\nAPPROVED AND AUTHORIZED FOR EXECUTION\n' \
        "$current_id" "${archive_path#$project_root/}" "$next_id" "${new_prompt#$project_root/}" "${new_history#$project_root/}"
}

require_root
mode=install
case "${1:-}" in
    '') collect_interactive ;;
    --input-dir) [[ $# == 2 ]] || fail 'use --input-dir <directory>'; validate_input_dir "$2" ;;
    --check-input) [[ $# == 2 ]] || fail 'use --check-input <directory>'; validate_input_dir "$2"; printf 'Structured owner input: VALID\n'; exit 0 ;;
    --help|-h) usage; exit 0 ;;
    *) usage >&2; fail "unknown option: $1" ;;
esac
load_current
render_pair
install_pair
