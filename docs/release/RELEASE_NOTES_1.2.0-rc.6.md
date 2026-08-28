# QWSG 1.2.0-rc.6 Private Candidate Notes

RC.6 remediates the guided-installer readiness synchronization defect manually reproduced during RC.5 pre-release acceptance. After Guardian activation, guided install and guided setup now share one cancellable, configured bounded wait for fresh canonical Guardian evidence.

The waiter captures the pre-activation evidence identity and accepts only a different integrity-checked state whose canonical overview reports the Guardian running with current freshness. Preserved state from an earlier installation therefore cannot produce a false pass. Polling is cancellable and the maximum wait derives from the validated Guardian cycle timeout plus the existing five-second completion margin; no arbitrary fixed sleep determines correctness.

When fresh mandatory evidence arrives within the bound, the installer performs the complete operational assessment. Missing optional notification produces `overall=partial`, the 97%/100% completion UI and summary are rendered, and installation exits zero. If fresh mandatory evidence does not arrive, the installer preserves data, prints a localized readiness/resume diagnostic, returns exit 4 and does not print a false success summary.

RC.6 retains the Task 070 declarative fail-closed migration architecture and declares the required `1.2.0-rc.2 -> 1.2.0-rc.6` path with no configuration/state schema transformation. Real clean-host and production-like acceptance remains pending; RC.6 is private and not a final release.
