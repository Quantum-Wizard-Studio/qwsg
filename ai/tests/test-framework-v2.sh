#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/../.." && pwd -P)"
passes=0
assert_text() { rg -q --fixed-strings "$2" "$repo_root/$1" || { printf 'FAIL missing semantic contract: %s\n' "$2" >&2; exit 1; }; passes=$((passes + 1)); }

[[ "$(<"$repo_root/ai/framework/VERSION")" == 2.0.0 ]]; passes=$((passes + 1))
assert_text ai/core/17_EXECUTION_MODEL.md 'Standard Execution Authority'
assert_text ai/core/17_EXECUTION_MODEL.md 'PRODUCT/FRAMEWORK DEFECT'
assert_text ai/core/17_EXECUTION_MODEL.md 'TEST OR ACCEPTANCE DEFECT'
assert_text ai/core/17_EXECUTION_MODEL.md 'ENVIRONMENTAL ISSUE'
assert_text ai/core/17_EXECUTION_MODEL.md 'EXPECTED BEHAVIOR'
assert_text ai/core/17_EXECUTION_MODEL.md 'INCONCLUSIVE'
assert_text ai/core/17_EXECUTION_MODEL.md 'A failed check, failed diagnostic,'
assert_text ai/core/17_EXECUTION_MODEL.md 'Candidate bytes are immutable after freeze.'
assert_text ai/core/11_ENGINEERING_LIFECYCLE.md 'There is no arbitrary retry limit'
assert_text ai/core/08_JOB_TEMPLATE.md '**Task targets and boundaries:**'
assert_text ai/framework/MIGRATION_1.1.0_TO_2.0.0.md 'Existing nine-category Authority Envelopes remain valid'
assert_text ai/framework/RETROSPECTIVE_TASKS_057_061.md 'acceptance defect or expected behavior'
assert_text ai/core/18_BOUNDED_DIAGNOSTIC_RUNNER.md 'privacy-safe machine-readable output'

! rg -q 'must not be repeated more than three times|Framework 1\.1\.0 active prompt' \
    "$repo_root/ai/core" "$repo_root/ai/scripts"
passes=$((passes + 1))

printf 'PASS: %d Framework v2 semantic assertions\n' "$passes"
