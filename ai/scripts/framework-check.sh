#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"
project_root="$(cd -- "$script_dir/../.." && pwd -P)"
config_path="$project_root/ai/config/engineering-project.conf"
framework_version_path="$project_root/ai/framework/VERSION"
mode=check

fail() {
    printf 'Framework validation error: %s\n' "$1" >&2
    exit 1
}

usage() {
    cat <<'EOF'
Usage:
  ./ai/scripts/framework-check.sh [--quiet]
  ./ai/scripts/framework-check.sh --show
  ./ai/scripts/framework-check.sh --run-validations

The command parses project configuration strictly as data. Validation commands
use tab-separated argv fields and are executed without source, eval, or a shell.
EOF
}

case "${1:-}" in
    '') ;;
    --quiet) mode=quiet ;;
    --show) mode=show ;;
    --run-validations) mode=run ;;
    --help|-h) usage; exit 0 ;;
    *) usage >&2; fail "unknown option: $1" ;;
esac
(( $# <= 1 )) || fail 'only one option is accepted'

[[ "$(pwd -P)" == "$project_root" ]] ||
    fail "run from the project root: $project_root"
[[ -f "$config_path" && ! -L "$config_path" ]] ||
    fail 'project configuration is missing or not a regular file'
[[ -f "$framework_version_path" && ! -L "$framework_version_path" ]] ||
    fail 'framework VERSION is missing or not a regular file'
iconv -f UTF-8 -t UTF-8 "$config_path" >/dev/null 2>&1 ||
    fail 'project configuration is not valid UTF-8'
if od -An -t x1 "$config_path" | grep -Eq '(^| )00( |$)'; then
    fail 'project configuration contains a NUL byte'
fi
if LC_ALL=C grep -q $'\r' "$config_path"; then
    fail 'project configuration must use LF line endings'
fi

readonly required_keys=(
    framework_version project_name project_slug repository_path repository_marker
    canonical_remote_name canonical_remote_url primary_branch
    owner_communication_language engineering_documentation_language
    required_reading validation_commands_file prompts_dir archive_prompts_dir
    history_dir test_tasks_dir task_number_width snapshot_location rollback_policy
)
declare -A config=() allowed=()
for key in "${required_keys[@]}"; do allowed["$key"]=true; done

line_number=0
while IFS= read -r line || [[ -n "$line" ]]; do
    line_number=$((line_number + 1))
    [[ -z "$line" || "$line" == \#* ]] && continue
    [[ "$line" == *=* ]] ||
        fail "configuration line $line_number lacks '='"
    key="${line%%=*}"
    value="${line#*=}"
    [[ "$key" =~ ^[a-z][a-z0-9_]*$ ]] ||
        fail "invalid configuration key at line $line_number"
    [[ -n "${allowed[$key]:-}" ]] ||
        fail "unknown configuration key: $key"
    [[ -z "${config[$key]+set}" ]] ||
        fail "duplicate configuration key: $key"
    [[ -n "$value" ]] || fail "empty configuration value: $key"
    config["$key"]="$value"
done <"$config_path"

for key in "${required_keys[@]}"; do
    [[ -n "${config[$key]:-}" ]] || fail "missing configuration key: $key"
done

framework_version="$(tr -d '\n' <"$framework_version_path")"
[[ "$framework_version" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] ||
    fail 'framework VERSION must use semantic X.Y.Z form'
[[ "${config[framework_version]}" == "$framework_version" ]] ||
    fail 'configured framework_version differs from ai/framework/VERSION'
[[ "${config[project_slug]}" =~ ^[a-z0-9]+(-[a-z0-9]+)*$ ]] ||
    fail 'project_slug must use lowercase kebab-case'
[[ "${config[repository_path]}" == . ]] ||
    fail 'repository_path must be project-relative dot'
[[ "${config[canonical_remote_name]}" =~ ^[A-Za-z0-9._-]+$ ]] ||
    fail 'canonical_remote_name is unsafe'
[[ "${config[canonical_remote_url]}" == https://* ]] ||
    fail 'canonical_remote_url must use HTTPS'
[[ "${config[primary_branch]}" =~ ^[A-Za-z0-9._/-]+$ &&
    "${config[primary_branch]}" != *..* ]] ||
    fail 'primary_branch is unsafe'
[[ "${config[task_number_width]}" == 3 ]] ||
    fail 'task_number_width must remain 3 for QWSG compatibility'
[[ "${config[rollback_policy]}" == mandatory ]] ||
    fail 'rollback_policy must remain mandatory'
[[ "${config[engineering_documentation_language]}" == English ]] ||
    fail 'engineering_documentation_language must remain English'
[[ "${config[snapshot_location]}" == /* ]] ||
    fail 'snapshot_location must be absolute'

validate_relative_path() {
    local label="$1" value="$2"
    [[ "$value" != /* && "$value" != *..* && "$value" != *$'\n'* ]] ||
        fail "$label must be a safe repository-relative path"
}

for key in repository_marker validation_commands_file prompts_dir \
    archive_prompts_dir history_dir test_tasks_dir; do
    validate_relative_path "$key" "${config[$key]}"
done
[[ -f "$project_root/${config[repository_marker]}" ]] ||
    fail 'configured repository marker is missing'
for key in prompts_dir archive_prompts_dir history_dir test_tasks_dir; do
    [[ -d "$project_root/${config[$key]}" ]] ||
        fail "configured directory is missing: $key"
done

IFS=',' read -r -a reading_paths <<<"${config[required_reading]}"
(( ${#reading_paths[@]} > 0 )) || fail 'required_reading is empty'
for path in "${reading_paths[@]}"; do
    validate_relative_path required_reading "$path"
    [[ -f "$project_root/$path" ]] ||
        fail "required reading file is missing: $path"
done

validation_file="$project_root/${config[validation_commands_file]}"
[[ -f "$validation_file" && ! -L "$validation_file" ]] ||
    fail 'validation command file is missing or not regular'
iconv -f UTF-8 -t UTF-8 "$validation_file" >/dev/null 2>&1 ||
    fail 'validation command file is not valid UTF-8'

branch="$(git -C "$project_root" branch --show-current)" ||
    fail 'cannot determine current Git branch'
[[ "$branch" == "${config[primary_branch]}" ]] ||
    fail "current branch differs from primary_branch: $branch"
remote_url="$(git -C "$project_root" remote get-url "${config[canonical_remote_name]}")" ||
    fail 'configured canonical Git remote is missing'
[[ "$remote_url" == "${config[canonical_remote_url]}" ]] ||
    fail 'canonical Git remote URL differs from project configuration'

validate_command_file() {
    local validation_id executable args=() line count=0
    declare -A seen_ids=()
    while IFS=$'\t' read -r validation_id executable args_text || \
        [[ -n "$validation_id$executable${args_text:-}" ]]; do
        [[ -z "$validation_id" || "$validation_id" == \#* ]] && continue
        count=$((count + 1))
        [[ "$validation_id" =~ ^[a-z0-9]+(-[a-z0-9]+)*$ ]] ||
            fail "unsafe validation id: $validation_id"
        [[ -z "${seen_ids[$validation_id]:-}" ]] ||
            fail "duplicate validation id: $validation_id"
        seen_ids["$validation_id"]=true
        validate_relative_path "validation executable" "$executable"
        [[ -x "$project_root/$executable" ||
            ( "$executable" == make && -n "$(command -v make)" ) ]] ||
            fail "validation executable is unavailable: $executable"
        args=()
        [[ -z "${args_text:-}" ]] || IFS=$'\t' read -r -a args <<<"$args_text"
        for arg in "${args[@]}"; do
            [[ "$arg" != *$'\n'* && "$arg" != *$'\r'* ]] ||
                fail "unsafe argument in validation: $validation_id"
        done
        if [[ "$mode" == run ]]; then
            printf 'Validation %s: ' "$validation_id"
            if [[ "$executable" == make ]]; then
                (cd "$project_root" && command make "${args[@]}")
            else
                (cd "$project_root" && "$project_root/$executable" "${args[@]}")
            fi
        fi
    done <"$validation_file"
    (( count > 0 )) || fail 'validation command file contains no commands'
}

validate_command_file

validate_active_task_structure() {
    local prompt heading section='' trimmed status
    mapfile -d '' active_prompts < <(
        find -P "$project_root/${config[prompts_dir]}" -maxdepth 1 -type f \
            -name '*.md' -print0
    )
    (( ${#active_prompts[@]} <= 1 )) ||
        fail "multiple active production tasks exist: ${#active_prompts[@]}"
    (( ${#active_prompts[@]} == 1 )) || return 0
    prompt="${active_prompts[0]}"
    iconv -f UTF-8 -t UTF-8 "$prompt" >/dev/null 2>&1 ||
        fail 'active task is not valid UTF-8'
    status="$(sed -nE 's/^- Status: `?([^`]+)`?$/\1/p' "$prompt")"
    [[ -n "$status" ]] || fail 'active task lacks status metadata'
    for heading in 'Task Metadata' 'Title' 'Objective' 'Scope' 'Out of Scope' \
        'Authority Envelope' 'Required Reading' 'Starting State Verification' 'Snapshot Requirements' \
        'Risk Assessment' 'Planned Work' 'Rollback Plan' 'Deliverables' \
        'Verification' 'Documentation Updates' 'Completion Criteria' \
        'Owner Approval Requirements'; do
        [[ "$(grep -Fxc "## $heading" "$prompt")" == 1 ]] ||
            fail "active task requires exactly one section: $heading"
    done
    if [[ "$status" != draft* ]]; then
        if grep -Eq '\[REQUIRES HUMAN EDITING\]|\[TODO\]|\bTBD\b|@[A-Z0-9_]+@' "$prompt"; then
            fail 'active task contains an unresolved placeholder'
        fi
        grep -Fq 'through the Engineering Task Builder' "$prompt" ||
            fail 'active task lacks generated owner approval evidence'
        validate_authority_envelope "$prompt" 'active task Authority Envelope'
    fi

    while IFS= read -r line || [[ -n "$line" ]]; do
        if [[ "$line" =~ ^##[[:space:]]+(.+)$ ]]; then
            section="${BASH_REMATCH[1]}"
            continue
        fi
        trimmed="${line#"${line%%[![:space:]]*}"}"
        trimmed="${trimmed#- }"
        trimmed="${trimmed#\\* }"
        if [[ "$trimmed" == 'git add .' || "$trimmed" == 'git add -A' ||
            "$trimmed" == 'git add --all' ]]; then
            [[ "$section" == 'Out of Scope' ]] ||
                fail "unsafe broad staging command outside Out of Scope: $trimmed"
        fi
    done <"$prompt"
}

validate_authority_envelope() {
    local file="$1" label="$2" authority_label legacy=true concise=true
    for authority_label in 'Authorized paths/components/systems' \
        'Routine operations' 'Correction/retest authority' \
        'Repository integration' 'Lifecycle completion' \
        'Permitted external actions' 'Evidence and rollback' \
        'Owner-reserved operations' 'Mandatory STOP conditions'; do
        grep -Fq "**$authority_label:**" "$file" || legacy=false
    done
    for authority_label in 'Task targets and boundaries' \
        'Permitted external actions' 'Owner-reserved decisions' \
        'Task-specific STOP conditions'; do
        grep -Fq "**$authority_label:**" "$file" || concise=false
    done
    [[ "$legacy" == true || "$concise" == true ]] ||
        fail "$label must use either the Framework 1.1 legacy or Framework 2.0 concise categories"
}

validate_active_task_structure

case "$mode" in
    quiet) ;;
    show)
        for key in "${required_keys[@]}"; do
            printf '%s=%s\n' "$key" "${config[$key]}"
        done
        ;;
    run) printf 'Configured engineering validations: PASS\n' ;;
    check)
        printf 'Reusable Engineering Framework %s: VALID (%s, %s)\n' \
            "$framework_version" "${config[project_slug]}" "$branch"
        ;;
esac
