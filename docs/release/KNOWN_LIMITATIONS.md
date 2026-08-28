# QWSG 1.2.0-rc.6 Known Limitations

The first native transition from QWSG 1.1.0 must be initiated with the fully
verified newer archive binary because 1.1.0 predates the `update` command.
Future supported updates use the installed command directly.

The guided installer currently explains the SMTP values and protected
credential workflow, but provider-specific values are entered through the
existing `qwsg config` and `qwsg notification credential` commands. The
`notify` update policy records operator preference; a background update
notifier and automatic privileged updater are not implemented.

- The supported production contract is limited to Ubuntu 24.04 LTS, systemd 255+, and linux-amd64.
- The Console is local terminal output; there is no Web Dashboard, REST API, listener, fleet or remote management.
- Community notification supports one administrator recipient through operator-controlled SMTP. QWS-managed delivery, multiple recipients, Telegram, Discord and webhooks are not included.
- QWSG reports and recommends; it provides no automatic remediation.
- SHA-256 checksums provide integrity, not publisher authenticity.
- Lifecycle email delivery is synchronous and reports SMTP acceptance, not end-recipient mailbox delivery. Duplicate suppression is bounded to one QWSG command process; externally repeated distinct operations remain distinct events.
- QWSG uses the proprietary source-available QWS Community / Free License Version 1.0; it is not an OSI open-source license.
