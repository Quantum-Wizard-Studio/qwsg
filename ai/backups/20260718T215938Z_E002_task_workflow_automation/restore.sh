#!/usr/bin/env bash
set -euo pipefail

expected_root='/home/qws/web/qwsg.quantumwizard.hu/qwsg'
baseline_commit='fa156697198a2230c938fb6780e9b8e42c860f57'

if [[ "$(pwd -P)" != "$expected_root" ]]; then
    printf 'Refusing rollback: run this script from exactly %s\n' "$expected_root" >&2
    exit 1
fi
if ! git cat-file -e "${baseline_commit}^{commit}" 2>/dev/null; then
    printf '%s\n' 'Refusing rollback: the recorded baseline commit is unavailable.' >&2
    exit 1
fi

printf '%s\n' 'QWSG E002 task-workflow rollback'
printf '%s\n' 'This restores only the listed core and original prompt files from the baseline.'
printf '%s\n' 'It removes only E002-created active/archive/script/history files and then empty new directories.'
printf '%s\n' 'The E002 audit snapshot remains available. No application or server files are touched.'
printf '%s' 'Type ROLLBACK-QWSG-E002 to continue: '
read -r confirmation
if [[ "$confirmation" != 'ROLLBACK-QWSG-E002' ]]; then
    printf '%s\n' 'Rollback cancelled.'
    exit 1
fi

tracked_baseline=(
    ai/core/02_PROJECT_STRUCTURE.md
    ai/core/06_ENGINEERING_STANDARDS.md
    ai/core/07_ENGINEERING_HISTORY.md
    ai/core/08_JOB_TEMPLATE.md
    ai/core/10_DOCUMENTATION_POLICY.md
    ai/prompts/README.md
    ai/prompts/001_PRODUCT_ARCHITECTURE.md
)
git restore --source="$baseline_commit" --worktree --staged -- "${tracked_baseline[@]}"

created_files=(
    ai/core/14_PROMPT_WORKFLOW.md
    ai/archive_prompts/001_2026-07-18_product-architecture-draft.md
    ai/prompts/002_CURRENT_TASK.md
    ai/scripts/next-task.sh
    ai/history/002_2026-07-18_task-workflow-automation.md
)
for path in "${created_files[@]}"; do
    if [[ -e "$path" ]]; then
        rm -- "$path"
    fi
done

for path in ai/archive_prompts ai/scripts; do
    if [[ -d "$path" ]]; then
        rmdir -- "$path" 2>/dev/null || printf 'Retained non-empty directory: %s\n' "$path"
    fi
done

printf '%s\n' 'Rollback complete. Review Git status before continuing.'
