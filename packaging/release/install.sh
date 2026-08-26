#!/bin/sh
set -eu

root=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd -P)
destdir=
replace=false
backup=
while test "$#" -gt 0; do
  case "$1" in
    --destdir) test "$#" -ge 2 || exit 2; destdir=$2; shift 2;;
    --replace) replace=true; shift;;
    --backup-dir) test "$#" -ge 2 || exit 2; backup=$2; shift 2;;
    *) printf 'install: unknown option: %s\n' "$1" >&2; exit 2;;
  esac
done
test "$(uname -s)" = Linux && test "$(uname -m)" = x86_64 || { printf '%s\n' 'install: supported platform is Linux x86-64' >&2; exit 1; }
(cd "$root" && sha256sum -c MANIFEST.sha256)
if $replace; then test -n "$backup" && test ! -e "$backup" || { printf '%s\n' 'install: --replace requires a new --backup-dir' >&2; exit 1; }; mkdir -m 0700 -p "$backup"; fi
install_one() {
  source=$1 destination=$2 mode=$3
  if test -e "$destination" || test -L "$destination"; then
    $replace || { printf 'install: destination exists: %s\n' "$destination" >&2; exit 1; }
    test ! -L "$destination" && test -f "$destination" || { printf 'install: unsafe destination: %s\n' "$destination" >&2; exit 1; }
    saved="$backup${destination#$destdir}"
    mkdir -p "$(dirname "$saved")"; cp -p "$destination" "$saved"
  fi
  install -d -m 0755 "$(dirname "$destination")"
  install -m "$mode" "$source" "$destination"
}
install_one "$root/bin/qwsg" "$destdir/usr/local/bin/qwsg" 0755
install_one "$root/lib/systemd/user/qwsg-guardian.service" "$destdir/usr/local/lib/systemd/user/qwsg-guardian.service" 0644
for file in "$root"/docs/*.md "$root/README.md" "$root/INSTALL.md" "$root/LICENSE" "$root/CHANGELOG.md" "$root/qwsg-config.json" "$root/RELEASE.json"; do install_one "$file" "$destdir/usr/local/share/doc/qwsg/$(basename "$file")" 0644; done
printf '%s\n' 'QWSG artifacts installed. Service was not enabled or started.'
printf '%s\n' 'Instructions: /usr/local/share/doc/qwsg/README.md and /usr/local/share/doc/qwsg/INSTALL.md'
printf '%s\n' 'Next (as the intended non-root user): qwsg setup'
printf '%s\n' 'Readiness check: qwsg readiness'
