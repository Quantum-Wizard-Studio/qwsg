# Smart Setup and Guided Guardian Activation

Smart Setup is an orchestration adapter over canonical Configuration,
Notification, Assessment, systemd probe, and Current Operator State owners. It
adds no progress store. Each run derives a versioned `qwsg.setup-flow` plan
from current evidence, so interruption and rerun are naturally resumable.

The ordered phases are environment, configuration, notification, and Guardian
activation/evidence. Human prompts and `qwsg setup --plan --format json` consume
the same presentation-independent plan. Existing `--accept-defaults`, `--set`,
`config`, notification, install-check, and readiness commands remain
deterministic automation interfaces.

Activation is a separately confirmed ordinary-user action limited to absolute
`/usr/bin/systemctl --user daemon-reload` and `enable --now
qwsg-guardian.service`. There is no shell, arbitrary unit, sudo, loginctl,
package action, or generic executor. Lingering remains independently assessed
guidance. Unit state never substitutes for fresh integrity-checked Current
Operator State; operators re-run `qwsg readiness` until qualified evidence is
available or the bounded Guardian cycle fails.

Direct hidden credential entry is deferred until a dependency-free no-echo
terminal boundary can be proven. The existing current-user mode-0600
`notification credential set --from-file` path remains the safe guided and
automation mechanism. Notification tests remain explicit and never create a
Guardian incident.

The setup plan's stable tokens and canonical readiness summaries are suitable
for a future GUI. A future Pro executor must remain Detect -> Plan -> Ask ->
Execute -> Verify -> Continue. VPS profile and branding remain deferred
presentation/product extensions.
