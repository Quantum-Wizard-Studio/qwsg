#!/bin/sh
set -eu

root=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd -P)
destdir=
while test "$#" -gt 0; do case "$1" in --destdir) test "$#" -ge 2 || exit 2; destdir=$2; shift 2;; *) printf 'uninstall: unknown option: %s\n' "$1" >&2; exit 2;; esac; done
(cd "$root" && sha256sum -c MANIFEST.sha256)
remove_one() {
  source=$1 destination=$2
  test ! -L "$destination" && test -f "$destination" || { printf 'uninstall: missing or unsafe owned artifact: %s\n' "$destination" >&2; exit 1; }
  test "$(sha256sum "$source" | awk '{print $1}')" = "$(sha256sum "$destination" | awk '{print $1}')" || { printf 'uninstall: modified artifact refused: %s\n' "$destination" >&2; exit 1; }
  rm -- "$destination"
}
remove_one "$root/bin/qwsg" "$destdir/usr/local/bin/qwsg"
remove_one "$root/lib/systemd/user/qwsg-guardian.service" "$destdir/usr/local/lib/systemd/user/qwsg-guardian.service"
for file in "$root"/docs/*.md "$root/README.md" "$root/INSTALL.md" "$root/LICENSE" "$root/CHANGELOG.md" "$root/qwsg-config.json"; do remove_one "$file" "$destdir/usr/local/share/doc/qwsg/$(basename "$file")"; done
rmdir "$destdir/usr/local/share/doc/qwsg" 2>/dev/null || true
printf '%s\n' 'QWSG release artifacts removed. User configuration and state were preserved.'
