#!/usr/bin/env bash
set -euo pipefail

expected_root='/home/qws/web/qwsg.quantumwizard.hu/qwsg'
baseline_commit='4b4a0ac97fe3ac21ef1fae4beec70870cd53cad5'
snapshot='ai/backups/20260718T233811Z_task003_product_definition'

if [[ "$(pwd -P)" != "$expected_root" ]]; then
    printf 'Refusing rollback: run this script from exactly %s\n' "$expected_root" >&2
    exit 1
fi
if ! git cat-file -e "${baseline_commit}^{commit}" 2>/dev/null; then
    printf '%s\n' 'Refusing rollback: the recorded baseline commit is unavailable.' >&2
    exit 1
fi

printf '%s\n' 'QWSG Task 003 Product Definition rollback'
printf '%s\n' 'This removes only the Product Definition and corrected Task 003 history path.'
printf '%s\n' 'It restores only named reference documents and the captured uncommitted Prompt 003/history state.'
printf '%s\n' 'The snapshot and pre-existing prompt rotation archive remain available.'
printf '%s' 'Type ROLLBACK-QWSG-TASK-003 to continue: '
read -r confirmation
if [[ "$confirmation" != 'ROLLBACK-QWSG-TASK-003' ]]; then
    printf '%s\n' 'Rollback cancelled.'
    exit 1
fi

tracked_references=(
    README.md
    ai/projects/QWSG.md
    ai/core/07_ENGINEERING_HISTORY.md
    ai/core/13_ROADMAP.md
)
git restore --source="$baseline_commit" --worktree --staged -- "${tracked_references[@]}"

for path in docs/PRODUCT_DEFINITION.md ai/history/003_2026-07-18_product-definition.md; do
    if [[ -e "$path" ]]; then
        rm -- "$path"
    fi
done

cp -p -- "$snapshot/current-state/003_CURRENT_TASK.md" ai/prompts/003_CURRENT_TASK.md
cp -p -- "$snapshot/current-state/003_2026-07-18_product-architecture.md" ai/history/003_2026-07-18_product-architecture.md

printf '%s\n' 'Rollback complete. Review Git status and document consistency before continuing.'
