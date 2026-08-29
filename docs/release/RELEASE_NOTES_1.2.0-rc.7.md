# QWSG 1.2.0-rc.7 Private Candidate Notes

RC.7 remediates the Guardian cgroup OOM instability reproduced on the loaded Contabo acceptance host while retaining the `MemoryMax=128M` and `TasksMax=32` security contract. The Guardian now supplies Go with a measured `64MiB` soft memory limit, leaving half of the cgroup ceiling for non-Go memory, collector subprocesses, and cgroup-accounted file cache. It also releases the one-cycle scheduler execution graph immediately after canonical publication instead of overlapping it with the next collection.

The remediation is based on a bounded loaded-host model: one qualified 367-record observation allocated approximately 112 MiB cumulatively and reached about 54 MiB process RSS in isolation. Five repeated cycles with a 32 MiB diagnostic Go limit remained bounded at about 43 MiB RSS. The production 64 MiB soft limit therefore remains above the observed 36 MiB transient managed-heap high-water region while reserving 64 MiB of independently enforced cgroup headroom. The limit is a GC pacing input, not OOM suppression; systemd restart and `oom-kill` exit evidence remain unchanged.

RC.7 includes the RC.6 guided-readiness correction unchanged. No Configuration, Guardian checkpoint, Scheduler state, Operator State, or notification schema changes are introduced.

RC.7 retains the Task 070 declarative fail-closed migration architecture and declares the required `1.2.0-rc.2 -> 1.2.0-rc.7` path with no configuration/state schema transformation. Real clean-host and production-like acceptance remains pending; RC.7 is private and not a final release.
