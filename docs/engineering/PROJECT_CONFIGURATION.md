# Engineering Project Configuration Reference

## Files

`ai/config/engineering-project.conf` is a strict `key=value` UTF-8 LF file.
Blank lines and lines beginning with `#` are ignored. Values are data: shell
expansion, quoting, command substitution, and repeated keys are unsupported.

`ai/config/engineering-validations.tsv` contains one validation per line:

```text
identifier<TAB>executable<TAB>argument<TAB>argument...
```

Commands are invoked as argv without a shell.

## Required fields

| Field | Meaning |
|---|---|
| `framework_version` | exact semantic version matching `ai/framework/VERSION` |
| `project_name` | human-readable project name |
| `project_slug` | lowercase kebab-case identity |
| `repository_path` | project-relative root, currently `.` |
| `repository_marker` | required project identity file |
| `canonical_remote_name` | validated Git remote name |
| `canonical_remote_url` | exact canonical HTTPS URL |
| `primary_branch` | required active branch |
| `owner_communication_language` | owner-facing message language |
| `engineering_documentation_language` | engineering artifact language |
| `required_reading` | comma-separated safe repository-relative files |
| `validation_commands_file` | safe relative TSV command list |
| `prompts_dir` | active production prompt directory |
| `archive_prompts_dir` | immutable production prompt archive |
| `history_dir` | task history directory |
| `test_tasks_dir` | independent test-task namespace |
| `task_number_width` | three-digit compatibility width |
| `snapshot_location` | absolute snapshot parent |
| `rollback_policy` | mandatory rollback policy |

Safety-critical behavior is not configurable. `task_number_width=3` and
`rollback_policy=mandatory` are enforced for QWSG compatibility.

## Commands

```bash
ai/scripts/framework-check.sh
ai/scripts/framework-check.sh --show
ai/scripts/framework-check.sh --run-validations
```

The validator must run from the repository root. It verifies framework version,
all fields, required reading, lifecycle directories, repository marker,
validation command structure, current branch, and canonical remote.

Never store passwords, tokens, private keys, credential-bearing URLs, shell
source, or environment secrets in project configuration.
