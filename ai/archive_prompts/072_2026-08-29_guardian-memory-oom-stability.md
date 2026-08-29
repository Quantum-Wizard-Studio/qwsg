# Current Engineering Task 072: Guardian Memory Bound & OOM Stability Remediation

## Task Metadata

- Task ID: `072`
- Task slug: `guardian-memory-oom-stability`
- Status: `complete`
- Date opened: `2026-08-29` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

Guardian Memory Bound & OOM Stability Remediation


## Objective

Diagnose and remediate the reproducible QWSG Guardian cgroup OOM instability observed on the loaded Contabo acceptance host. Determine with evidence whether the failure is caused by QWSG memory retention or growth, collector/inventory workload characteristics, transient allocation pressure, systemd/cgroup accounting behavior, an incorrectly sized 128 MiB Guardian memory contract, or another deterministic cause. Implement the smallest safe and measurable remediation, preserve the established Guardian security/resource-control architecture, and produce a new private release candidate suitable for resuming final QWSG 1.2.0 acceptance.


## Scope

Work only on the Guardian memory/OOM stability problem and directly required implementation, tests, resource-contract documentation, release-candidate packaging, update/rollback compatibility, and acceptance guidance. Reproduce the failure in controlled tests where practical; inspect Guardian execution and collector lifecycle; establish meaningful memory measurements; determine the root cause before selecting remediation; add regression coverage for repeated Guardian cycles and loaded-host conditions; validate bounded long-running behavior; preserve deterministic update and rollback behavior from the currently accepted migration baseline; produce the next private release candidate only after all mandatory gates pass.


## Out of Scope

Do not add unrelated product features. Do not implement new notification transports. Do not redesign the Guardian architecture unless evidence proves a narrowly scoped architectural correction is necessary. Do not modify the Contabo or OVH hosts remotely. Do not publish a final v1.2.0 tag, public artifact, or Forgejo Release. Do not weaken security controls merely to make the test pass. Do not remove the Guardian memory limit. Do not arbitrarily change MemoryMax from 128 MiB to 256 MiB or another value without measured evidence and documented resource-contract justification. Do not hide, suppress, or automatically recover from OOM in a way that masks the underlying cause.


## Authority Envelope

**Task targets and boundaries:** Guardian memory/OOM diagnosis and the directly required implementation, tests, resource-contract documentation, private release-candidate packaging, deterministic update/rollback compatibility, lifecycle evidence, and acceptance guidance described by this task. No unrelated features or architecture redesign are authorized.

**Permitted external actions:** local repository work, local controlled fixtures and cgroup-capable testing where available, and private release-candidate construction. Read-only verification of canonical Git state is permitted. The Contabo and OVH acceptance hosts must not be mutated remotely; no public release, final tag, or Forgejo Release is authorized.

**Owner-reserved decisions:** manual OVH/Contabo acceptance, final v1.2.0 publication, any security-control weakening, and any resource-contract increase not quantitatively justified by task evidence.

**Task-specific STOP conditions:** stop and classify Task 072 BLOCKED if root cause cannot be established, bounded stability cannot be demonstrated, the only remediation is an unjustified resource-limit increase, migration/rollback fails, or any mandatory gate cannot pass. Do not produce the private candidate unless all mandatory gates pass.


## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/core/17_EXECUTION_MODEL.md`
- `ai/core/18_BOUNDED_DIAGNOSTIC_RUNNER.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

Canonical repository state after completed Task 071 is clean and lifecycle-idle. HEAD and origin/main are expected at 776ee47a10530266ce020526c813cab57c89fe62. Task 071 produced private RC.6, artifact dist/qwsg-1.2.0-rc.6-linux-amd64.tar.gz, SHA256 7f29f4cfe361412680c439eba7d8da5ae7d949b8702ce25c460baf991de99e33, candidate source commit 6355022fd08ae4461d79edfc86943418fc1958d0. RC.6 fixed guided-installer readiness synchronization and retains the explicit deterministic 1.2.0-rc.2 -> 1.2.0-rc.6 migration path. Manual OVH RC.6 clean/preserved-state install, 100% guided completion, readiness, and reboot Guardian activation passed. Manual Contabo migration from actual installed RC.2 to RC.6 also passed and mandatory readiness remained satisfied. Final acceptance is nevertheless blocked because the Guardian repeatedly suffers cgroup OOM kills on the loaded Contabo host. Historical evidence proves the same behavior existed under RC.2, so it is not currently classified as an RC.6 regression. The loaded Contabo host is Ubuntu 24.04.4 LTS x86_64 with HestiaCP and a full hosting/mail/database stack. Guardian systemd resource contract currently includes TasksMax=32 and MemoryMax=128M. Before RC.6 the Guardian repeatedly approached the memory ceiling and was OOM-killed. After RC.6 and after reboot the failure remains immediately reproducible. Current-boot evidence includes: Guardian started 2026-08-29 13:11:19 CEST and was killed by the OOM killer at 13:11:33; systemd restarted it at 13:11:38. After a later user-manager restart Guardian started at 13:15:08 and was OOM-killed at 13:15:20, restarted at 13:15:26, OOM-killed again at 13:15:40, and restarted at 13:15:45. systemd explicitly reports "A process of this unit has been killed by the OOM killer", main process status=9/KILL and result='oom-kill'. One reported cycle showed only 35.2M memory peak despite the cgroup OOM event, so apparent systemctl Memory/MemoryPeak values must not be treated as sufficient proof of the actual accounting/root cause. Earlier RC.2 evidence showed repeated OOM kills approximately every 5-10 minutes and restart counters reaching high values. Contabo must remain an external acceptance/reproduction host; this task must not mutate it remotely.


## Snapshot Requirements

Before modifying tracked files, create and verify the canonical task snapshot according to the Engineering Safety Rule and current framework. Preserve rollback capability for every tracked modification. Capture repository/lifecycle baseline, candidate/release baseline, Guardian resource-control definitions, collector execution architecture, inventory persistence behavior, and all code paths capable of retaining or accumulating memory across Guardian cycles. Any experimental instrumentation must be safe, bounded, removable or intentionally documented, and must not expose host-sensitive data or secrets.


## Risk Assessment

Medium-to-high release-blocking engineering risk. The main risks are misdiagnosing a cgroup OOM as a simple process RSS problem, masking a memory leak by increasing MemoryMax, changing the resource contract without evidence, introducing unbounded diagnostic instrumentation, destabilizing collector isolation, weakening Guardian limits, breaking RC.2 migration/update/rollback behavior, or producing an RC that passes synthetic tests but remains unstable on a loaded host.


## Planned Work

First reproduce or model the failure sufficiently to establish a measurable baseline. Inspect the Guardian long-running cycle, collector registry/execution, inventory construction/persistence, canonical evidence generation, retained references, buffers, serialization, subprocess/process behavior if any, goroutine lifecycle, timers/tickers, channels, caches, and repeated-cycle allocations. Inspect the systemd unit and cgroup-v2 memory controls and determine which memory categories can trigger the unit OOM. Where appropriate use deterministic Go memory/runtime profiling, bounded repeated-cycle tests, cgroup evidence, and before/after measurements. Distinguish steady-state memory, transient peak allocation, cumulative growth/leak, child/process accounting, page/cache effects, and systemd reporting limitations. Establish the root cause in the task evidence before selecting the remediation. If an implementation defect is found, fix that defect and retain the smallest justified MemoryMax. If the implementation is demonstrably bounded but 128 MiB is below the legitimate measured workload requirement, calculate and document a new resource limit with explicit safety headroom based on measured worst-case behavior rather than an arbitrary round-number increase. If both implementation and contract changes are required, justify each independently. Add regression tests that exercise repeated Guardian cycles and prove bounded memory behavior or enforce the measurable invariant appropriate to the root cause. Validate that no goroutine/process/ticker/channel accumulation occurs across cycles. Preserve TasksMax and other security constraints unless evidence specifically requires a separately justified change. Preserve Guardian restart/failure semantics and make OOM/resource-limit failures observable rather than silently hidden. Validate clean-host and loaded-host-compatible behavior in local/fixture testing without remotely mutating acceptance VPSs. Preserve notification behavior and EN/HU/DE contracts where affected. Preserve deterministic RC.2 update compatibility to the new candidate and deterministic rollback to RC.2, adapting the declared migration target from RC.6 to the new candidate as required by the established compatibility architecture. Build a new private RC only after the root-cause, regression, security, update/rollback, full Go, race, vet, formatting, Framework, Builder, lifecycle, release, manifest, mode, ownership, and reproducibility gates pass.


## Rollback Plan

Rollback must restore the exact pre-Task-072 repository state from the verified snapshot. Any resource-contract/unit changes must be independently reversible. The generated candidate must retain the existing deterministic update rollback guarantees, including restoration of the prior binary/unit/config/state contract. No rollback procedure may depend on unverified network resources or mutate the external OVH/Contabo acceptance hosts during this task.


## Deliverables

A root-cause analysis backed by measurements; implementation of the smallest justified remediation; regression tests for repeated Guardian cycles and memory/resource stability; any required systemd resource-contract update with measured justification; updated architecture/security/operations/release documentation where the effective contract changes; preserved RC.2 -> new-candidate deterministic migration and new-candidate -> RC.2 rollback; complete task history/evidence; a reproducible private next release candidate with artifact path, SHA256, candidate source commit, final repository commit, and explicit READY FOR FINAL ACCEPTANCE or BLOCKED classification.


## Verification

Verify the pre-change failure model and post-change behavior with deterministic or bounded tests. Record relevant memory metrics and explain what they measure. Prove that repeated Guardian cycles do not exhibit unbounded growth or lifecycle accumulation. If MemoryMax changes, prove why the previous 128 MiB contract is insufficient after implementation defects are excluded/fixed and document measured worst-case usage plus chosen headroom. Run focused tests, full Go tests, race tests, go vet, formatting checks, Framework tests, Builder tests, lifecycle tests, release tests, update/migration/rollback tests, notification regressions, security/redaction tests, release reproducibility across repeated builds/umasks/source modes, manifest verification, canonical permissions and numeric ownership validation. Verify RC.2 -> new candidate update compatibility and deterministic rollback without using the real Contabo host. Verify repository cleanliness and HEAD/origin/main convergence at completion.


## Documentation Updates

Update the canonical Task 072 history with the observed real-host evidence, root-cause analysis, measurements, remediation decision, before/after results, resource-contract decision, tests, migration/rollback evidence, release-candidate identity and limitations. Update Guardian/resource-limit and upgrade/rollback documentation if behavior or limits change. Explicitly state that real OVH/Contabo acceptance remains a separate Owner-controlled manual gate and was not performed by this task.


## Completion Criteria

Task 072 is complete only when the actual OOM root cause is identified with evidence; the smallest justified remediation is implemented; repeated Guardian operation is demonstrated bounded under representative stress; resource limits remain enforced and any changed limit is quantitatively justified; RC.2 -> new candidate migration and rollback are deterministic; all mandatory engineering/security/release/reproducibility gates pass; repository and lifecycle are clean; and the new private candidate is classified READY FOR FINAL ACCEPTANCE. If root cause cannot be established, bounded stability cannot be demonstrated, the only available remediation is an unjustified resource-limit increase, migration/rollback fails, or any mandatory gate fails, stop and classify Task 072 BLOCKED without producing or publishing a final release.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-29 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
