#!/usr/bin/env bash
set -euo pipefail

expected_root='/home/qws/web/qwsg.quantumwizard.hu/qwsg'
baseline_commit='ac164d41bd8ef951cf3a4eecc307dd834b5d750d'

if [[ "$(pwd -P)" != "$expected_root" ]]; then
    printf 'Refusing rollback: run this script from exactly %s\n' "$expected_root" >&2
    exit 1
fi

if ! git cat-file -e "${baseline_commit}^{commit}" 2>/dev/null; then
    printf '%s\n' 'Refusing rollback: the recorded baseline commit is unavailable.' >&2
    exit 1
fi

printf '%s\n' 'QWSG engineering-standard update rollback'
printf '%s\n' 'This restores only the listed core documents from the recorded baseline.'
printf '%s\n' 'It removes only ai/history/001_engineering_standard_update.md if present.'
printf '%s\n' 'The audit snapshot is retained. No application or server files are touched.'
printf '%s' 'Type ROLLBACK-QWSG-STANDARD-UPDATE to continue: '
read -r confirmation
if [[ "$confirmation" != 'ROLLBACK-QWSG-STANDARD-UPDATE' ]]; then
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
if [[ -e ai/history/001_engineering_standard_update.md ]]; then
    rm -- ai/history/001_engineering_standard_update.md
fi

printf '%s\n' 'Rollback complete. Review git status before continuing.'
