#!/usr/bin/env bash
set -euo pipefail

expected_root='/home/qws/web/qwsg.quantumwizard.hu/qwsg'
baseline_commit='cba6318815f59398dda9fcd28df498b5990963de'

if [[ "$(pwd -P)" != "$expected_root" ]]; then
    printf 'Refusing rollback: run this script from exactly %s\n' "$expected_root" >&2
    exit 1
fi

if ! git cat-file -e "${baseline_commit}^{commit}" 2>/dev/null; then
    printf '%s\n' 'Refusing rollback: the recorded baseline commit is unavailable.' >&2
    exit 1
fi

printf '%s\n' 'QWSG E001 engineering workflow rollback'
printf '%s\n' 'This restores only the listed core documents from the recorded baseline.'
printf '%s\n' 'It removes only the E001 history record and two prompt-workflow documents.'
printf '%s\n' 'The audit snapshot is retained. Application and server files are never touched.'
printf '%s' 'Type ROLLBACK-QWSG-E001 to continue: '
read -r confirmation
if [[ "$confirmation" != 'ROLLBACK-QWSG-E001' ]]; then
    printf '%s\n' 'Rollback cancelled.'
    exit 1
fi

documents=(
    ai/core/01_CONSTITUTION.md
    ai/core/03_AGENTS.md
    ai/core/06_ENGINEERING_STANDARDS.md
    ai/core/07_ENGINEERING_HISTORY.md
    ai/core/08_JOB_TEMPLATE.md
    ai/core/10_DOCUMENTATION_POLICY.md
)
git restore --source="$baseline_commit" --worktree --staged -- "${documents[@]}"

new_files=(
    ai/prompts/README.md
    ai/prompts/001_PRODUCT_ARCHITECTURE.md
    ai/history/E001_engineering_workflow_refinement.md
)
for path in "${new_files[@]}"; do
    if [[ -e "$path" ]]; then
        rm -- "$path"
    fi
done

printf '%s\n' 'Rollback complete. The E001 snapshot remains available; review git status.'
