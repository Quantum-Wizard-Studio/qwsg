# QWSG 1.0 Known Limitations

- The supported production contract is limited to Ubuntu 24.04 LTS, systemd 255+, and linux-amd64.
- The Console is local terminal output; there is no Web Dashboard, REST API, listener, fleet or remote management.
- Notification delivery is provider-neutral architecture only; SMTP, Telegram, Discord and webhook transports are post-1.0.
- QWSG reports and recommends; it provides no automatic remediation.
- SHA-256 checksums provide integrity, not publisher authenticity.
- QWSG uses the proprietary source-available QWS Community / Free License Version 1.0; it is not an OSI open-source license.
