#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/../.." && pwd -P)"
work="$(mktemp -d /tmp/qw-engineering-framework-test.XXXXXX)"
trap 'rm -rf -- "$work"' EXIT
passes=0

new_fixture() {
    local name="$1" root="$work/$1" reading
    mkdir -p "$root/ai/scripts" "$root/ai/framework" "$root/ai/config" \
        "$root/ai/core" "$root/ai/prompts" "$root/ai/archive_prompts" \
        "$root/ai/history" "$root/ai/test_tasks" "$root/bin"
    cp "$repo_root/ai/scripts/framework-check.sh" "$root/ai/scripts/"
    chmod +x "$root/ai/scripts/framework-check.sh"
    printf '1.0.0\n' >"$root/ai/framework/VERSION"
    printf 'fixture\n' >"$root/VERSION"
    for reading in 00_PROJECT_PHILOSOPHY 01_CONSTITUTION 03_AGENTS \
        08_JOB_TEMPLATE 11_ENGINEERING_LIFECYCLE 14_PROMPT_WORKFLOW 16_GIT_POLICY; do
        printf '# fixture\n' >"$root/ai/core/$reading.md"
    done
    cat >"$root/bin/pass" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
exit 0
EOF
    chmod +x "$root/bin/pass"
    printf 'smoke\tbin/pass\n' >"$root/ai/config/engineering-validations.tsv"
    cat >"$root/ai/config/engineering-project.conf" <<'EOF'
framework_version=1.0.0
project_name=Synthetic Fixture
project_slug=synthetic-fixture
repository_path=.
repository_marker=VERSION
canonical_remote_name=origin
canonical_remote_url=https://example.invalid/synthetic/fixture.git
primary_branch=trunk
owner_communication_language=German
engineering_documentation_language=English
required_reading=ai/core/00_PROJECT_PHILOSOPHY.md,ai/core/01_CONSTITUTION.md,ai/core/03_AGENTS.md,ai/core/08_JOB_TEMPLATE.md,ai/core/11_ENGINEERING_LIFECYCLE.md,ai/core/14_PROMPT_WORKFLOW.md,ai/core/16_GIT_POLICY.md
validation_commands_file=ai/config/engineering-validations.tsv
prompts_dir=ai/prompts
archive_prompts_dir=ai/archive_prompts
history_dir=ai/history
test_tasks_dir=ai/test_tasks
task_number_width=3
snapshot_location=/tmp
rollback_policy=mandatory
EOF
    (
        cd "$root"
        git init -q -b trunk
        git remote add origin https://example.invalid/synthetic/fixture.git
    )
    printf '%s\n' "$root"
}

expect_success() {
    local root="$1"; shift
    (cd "$root" && "$@") >/dev/null ||
        { printf 'FAIL expected success: %s\n' "$*" >&2; exit 1; }
    passes=$((passes + 1))
}

expect_failure() {
    local root="$1"; shift
    if (cd "$root" && "$@") >/dev/null 2>&1; then
        printf 'FAIL expected failure: %s\n' "$*" >&2
        exit 1
    fi
    passes=$((passes + 1))
}

root="$(new_fixture valid)"
expect_success "$root" ./ai/scripts/framework-check.sh
expect_success "$root" ./ai/scripts/framework-check.sh --show
expect_success "$root" ./ai/scripts/framework-check.sh --run-validations
[[ ! -e "$root/ai/prompts/001_CURRENT_TASK.md" ]]
passes=$((passes + 1))

root="$(new_fixture missing-config)"
rm "$root/ai/config/engineering-project.conf"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture duplicate-key)"
printf 'project_slug=duplicate\n' >>"$root/ai/config/engineering-project.conf"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture unknown-key)"
printf 'disable_approval=yes\n' >>"$root/ai/config/engineering-project.conf"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture optional-rollback)"
sed -i 's/rollback_policy=mandatory/rollback_policy=optional/' \
    "$root/ai/config/engineering-project.conf"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture bad-width)"
sed -i 's/task_number_width=3/task_number_width=4/' \
    "$root/ai/config/engineering-project.conf"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture missing-reading)"
rm "$root/ai/core/16_GIT_POLICY.md"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture remote-mismatch)"
git -C "$root" remote set-url origin https://example.invalid/other.git
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture branch-mismatch)"
git -C "$root" branch -m other
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture unsafe-path)"
sed -i 's#prompts_dir=ai/prompts#prompts_dir=../prompts#' \
    "$root/ai/config/engineering-project.conf"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture unsafe-command)"
printf 'bad\t../bin/pass\n' >"$root/ai/config/engineering-validations.tsv"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture missing-command)"
printf 'bad\tbin/missing\n' >"$root/ai/config/engineering-validations.tsv"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture insecure-remote)"
sed -i 's#https://example.invalid/synthetic/fixture.git#ssh://example.invalid/fixture.git#' \
    "$root/ai/config/engineering-project.conf"
git -C "$root" remote set-url origin ssh://example.invalid/fixture.git
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture multiple-active)"
touch "$root/ai/prompts/001_CURRENT_TASK.md" "$root/ai/prompts/002_CURRENT_TASK.md"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture malformed-task)"
cat >"$root/ai/prompts/001_CURRENT_TASK.md" <<'EOF'
# Current Engineering Task 001: Malformed
## Task Metadata
[TODO]
EOF
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture missing-task-section)"
cp "$repo_root/ai/prompts/015_CURRENT_TASK.md" "$root/ai/prompts/001_CURRENT_TASK.md"
sed -i '/^## Snapshot Requirements$/d' "$root/ai/prompts/001_CURRENT_TASK.md"
expect_failure "$root" ./ai/scripts/framework-check.sh

root="$(new_fixture unsafe-task-git)"
cp "$repo_root/ai/prompts/015_CURRENT_TASK.md" "$root/ai/prompts/001_CURRENT_TASK.md"
sed -i '/^## Owner Approval Requirements$/i git add .' \
    "$root/ai/prompts/001_CURRENT_TASK.md"
expect_failure "$root" ./ai/scripts/framework-check.sh

! rg -n 'Quantum Wizard Server Guardian|qwsg\\.quantumwizard|/home/qws' \
    "$repo_root/ai/scripts/framework-check.sh"
passes=$((passes + 1))

printf 'PASS: %d framework assertions\n' "$passes"
