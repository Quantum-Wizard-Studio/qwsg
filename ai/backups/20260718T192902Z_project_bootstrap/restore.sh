#!/usr/bin/env bash
set -euo pipefail

expected_root='/home/qws/web/qwsg.quantumwizard.hu/qwsg'
snapshot_rel='ai/backups/20260718T192902Z_project_bootstrap'

if [[ "$(pwd -P)" != "$expected_root" ]]; then
    printf 'Refusing rollback: run this script from exactly %s\n' "$expected_root" >&2
    exit 1
fi

printf '%s\n' 'QWSG bootstrap rollback'
printf '%s\n' 'This removes only the explicitly listed bootstrap files and empty directories inside the QWSG root.'
printf '%s\n' 'It changes the unborn/current branch name from main back to master when safe.'
printf '%s\n' "It retains the audit snapshot at $snapshot_rel so this procedure remains available."
printf '%s' 'Type ROLLBACK-QWSG-BOOTSTRAP to continue: '
read -r confirmation
if [[ "$confirmation" != 'ROLLBACK-QWSG-BOOTSTRAP' ]]; then
    printf '%s\n' 'Rollback cancelled.'
    exit 1
fi

current_branch=''
has_bootstrap_commit='false'
if git symbolic-ref -q HEAD >/dev/null 2>&1; then
    current_branch="$(git symbolic-ref --short HEAD)"
fi
if [[ "$current_branch" == 'main' ]] && git rev-parse --verify HEAD >/dev/null 2>&1; then
    commit_count="$(git rev-list --count HEAD)"
    commit_subject="$(git log -1 --format=%s)"
    if [[ "$commit_count" != '1' || "$commit_subject" != 'chore: bootstrap QWSG project foundation' ]]; then
        printf '%s\n' 'Refusing rollback: main contains work beyond the single bootstrap commit.' >&2
        exit 1
    fi
    has_bootstrap_commit='true'
fi

files=(
    README.md LICENSE CHANGELOG.md VERSION .gitignore
    ai/README.md ai/projects/QWSG.md ai/history/000_project_bootstrap.md
    ai/core/00_PROJECT_PHILOSOPHY.md ai/core/01_CONSTITUTION.md
    ai/core/02_PROJECT_STRUCTURE.md ai/core/03_AGENTS.md
    ai/core/04_ARCHITECTURE.md ai/core/05_SYSTEM_MAP.md
    ai/core/06_ENGINEERING_STANDARDS.md ai/core/07_ENGINEERING_HISTORY.md
    ai/core/08_JOB_TEMPLATE.md ai/core/09_DELIVERY_POLICY.md
    ai/core/10_DOCUMENTATION_POLICY.md ai/core/11_SECURITY_POLICY.md
    ai/core/12_RELEASE_POLICY.md ai/core/13_ROADMAP.md
)

for path in "${files[@]}"; do
    if [[ -e "$path" ]]; then
        rm -- "$path"
    fi
done

directories=(
    ai/core ai/projects ai/prompts ai/history
    docs/architecture docs/installation docs/administration docs/development
    docs/security docs/releases docs/history docs
    installer agent console modules tests scripts tools build
)
for path in "${directories[@]}"; do
    if [[ -d "$path" ]]; then
        rmdir -- "$path" 2>/dev/null || printf 'Retained non-empty directory: %s\n' "$path"
    fi
done

if [[ "$current_branch" == 'main' ]]; then
    git symbolic-ref HEAD refs/heads/master
    if [[ "$has_bootstrap_commit" == 'true' ]]; then
        git update-ref -d refs/heads/main
    fi
fi
if [[ -e .git/index ]]; then
    rm -- .git/index
fi
chmod u-w .git

printf '%s\n' 'Rollback finished. The snapshot, non-empty directories, and unreachable Git objects were retained.'
