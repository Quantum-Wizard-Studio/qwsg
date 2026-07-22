# Repository Agent Guidance

## QWSG job workflow

For this repository, `job` means load the current QWSG engineering task. Use the canonical project-local command `bin/job` and the reusable workflow in `.agents/skills/qwsg-job/SKILL.md`. At most one active prompt is allowed under `ai/prompts/`; when present, that prompt is the authoritative task scope. An empty prompt directory is valid only for the canonical idle state defined by the lifecycle policy. Neither the command nor the skill overrides the Project Constitution or the Project Owner's authority.
