# QWSG 1.0 Quick Start

QWSG 1.0 is a local, non-root Linux Server Guardian. Verify the downloaded archive checksum, unpack it, then install the fixed `/usr/local` artifact set:

```sh
sha256sum -c qwsg-1.0.0-linux-amd64.tar.gz.sha256
tar -xzf qwsg-1.0.0-linux-amd64.tar.gz
cd qwsg-1.0.0-linux-amd64
sudo ./install.sh
qwsg version
qwsg observe
qwsg observe
```

The first observation can establish the local baseline. The second can evaluate change and health. `qwsg` opens the Console on a terminal and prints one read-only Overview otherwise. Unknown or degraded results are truthful when evidence is missing, partial, stale, or a cycle failed.

Install the user unit for the ordinary runtime user, then enable and start it explicitly:

```sh
mkdir -p ~/.config/systemd/user
cp /usr/local/lib/systemd/user/qwsg-guardian.service ~/.config/systemd/user/
systemctl --user daemon-reload
systemctl --user enable --now qwsg-guardian.service
```

Installation never enables the service and never changes user lingering.
