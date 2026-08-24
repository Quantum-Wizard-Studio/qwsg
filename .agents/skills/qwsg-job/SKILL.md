---
name: qwsg-job
description: Load, validate, and execute the single active Quantum Wizard Server Guardian engineering task under repository governance. Use for explicit QWSG requests such as "job", "/job", "Uj feladat", "Új feladat", "aktualis feladat", "aktuális feladat", "inditsd a feladatot", "indítsd a feladatot", or an explicit invocation of the qwsg-job skill; do not trigger for generic employment, background-job, or unrelated task discussions.
---

# QWSG Job

Use this workflow to start the currently authorized QWSG engineering task. Treat `/job` as human shorthand; it may not be a built-in Codex slash command.

1. Resolve the repository root and verify it is the QWSG root containing `bin/job`, `VERSION`, and the required `ai/core/` documents.
2. Run `bin/job --check` from the verified repository. A lifecycle, authority,
   security, privacy, rollback, or project-identity error stops execution. A
   later recoverable task validation failure is diagnosed, corrected, and
   retested when the correction remains inside the approved Authority Envelope.
3. Obtain the active prompt with `bin/job --path`, then read it as data. Never source it, evaluate it, or execute Markdown as shell code.
4. Read every document listed in the prompt's Required Reading section, including the Constitution, applicable `AGENTS.md` instructions, and this skill where relevant.
5. Confirm no unresolved human-editing fields remain. Read and enforce the
   prompt's Authority Envelope and Starting State Verification. A
   Builder-approved task is authorized to start without another routine Owner
   gate.
6. Create and verify the required rollback-capable snapshot before modifying task targets.
7. Execute only the scope authorized by the active prompt and human authority. Reading a task never authorizes scope expansion.
8. Locate the matching record with `bin/job --history` and update it throughout implementation and delivery.
9. Run every verification required by the active prompt and preserve rollback
   capability. Diagnose, correct, and retest recoverable in-scope failures.
   Stop on a material boundary difference or an Owner-reserved operation.
10. Finish with the owner-facing report in the configured preferred language.

Remain subordinate to `ai/core/01_CONSTITUTION.md`, all applicable `AGENTS.md` files, the active prompt, and the Project Owner. Never bypass approval, sandbox, safety, or authority requirements.
