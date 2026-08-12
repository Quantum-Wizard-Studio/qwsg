# QWSG 1.0.0-rc.3 Release Notes

This private technical Release Candidate refreshes the accepted QWSG 1.0 local
Guardian after clean-host acceptance exposed and Task 041 corrected the
first-use state bootstrap boundary. It retains the frozen product behavior,
supported platform and artifact structure from RC.2.

RC.3 includes the accepted Task 041 release-blocking corrections:

- a first observation safely creates the missing per-user QWSG state hierarchy,
  including clean accounts where `.local` and `.local/state` do not exist;
- private QWSG directories remain mode `0700` and private state files remain
  mode `0600`, while existing ancestors are not chmodded or taken over;
- symlink-component, wrong-owner and unsafe-mode rejection remain enforced;
- truthful partial Inventory can establish and publish the first baseline when
  optional component capability, including Go, is unavailable;
- bounded `state_bootstrap_failed` and `state_publication_failed` diagnostics
  distinguish first-use initialization and publication failures.

RC.3 also contains the accepted Task 039 large-Report Alert integration,
read-only Console refresh, bounded Runtime diagnostics, Attention correlation
and lifecycle-freshness behavior already released in RC.2.

RC.3 adds no feature, collector, provider, network interface, remediation,
Dashboard, API, fleet capability, licensing enforcement, runtime dependency or
AI. Clean Ubuntu 24.04 installation and reboot acceptance remain the Owner's
next separate gate. Public licensing, signing, Git tagging, upload and
publication require separate Owner authority.
