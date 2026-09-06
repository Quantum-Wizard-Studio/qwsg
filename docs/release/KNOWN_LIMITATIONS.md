# QWSG 1.3.1 Known Limitations

The first native transition from QWSG 1.1.0 must be initiated with the fully
verified newer archive binary because 1.1.0 predates the `update` command.
Future supported updates use the installed command directly.

The guided installer currently explains the SMTP values and protected
credential workflow, but provider-specific values are entered through the
existing `qwsg config` and `qwsg notification credential` commands. The
`notify` policy enables Guardian notification for authenticated supported newer
releases when Community SMTP is ready. Automatic installation is not implemented.

- The supported production contract is limited to Ubuntu 24.04 LTS, systemd 255+, and linux-amd64.
- The Console is local terminal output; there is no Web Dashboard, REST API, listener, fleet or remote management.
- Community notification supports one administrator recipient through operator-controlled SMTP. QWS-managed delivery, multiple recipients, Telegram, Discord and webhooks are not included.
- QWSG reports and recommends; it provides no automatic remediation.
- SHA-256 checksums provide artifact integrity. Official awareness metadata is authenticated separately with the bundled Ed25519 trust anchor.
- Lifecycle email delivery is synchronous and reports SMTP acceptance, not end-recipient mailbox delivery. Duplicate suppression is bounded to one QWSG command process; externally repeated distinct operations remain distinct events.
- QWSG uses the proprietary source-available QWS Community / Free License Version 1.0; it is not an OSI open-source license.

- Scheduler state over 8 MiB is rejected before decoding and retained for operator review. Monitoring cannot be declared healthy until invalid legacy state is resolved.
- Update-availability notification deduplication is persistent; lifecycle-operation email deduplication has the narrower process scope described above.

- The first 1.3.0 to 1.3.1 correction uses the fully verified 1.3.1 archive binary as the native updater because the installed 1.3.0 has no compiled route to its successor. No public 1.3.0 backport or replacement artifact is created.
