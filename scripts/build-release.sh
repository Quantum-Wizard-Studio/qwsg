#!/bin/sh
set -eu

repo=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd -P)
version=$(tr -d '\r\n' < "$repo/VERSION")
case "$version" in
  1.0.0) :;;
  1.0.0-rc.*)
    rc_number=${version#1.0.0-rc.}
    case "$rc_number" in ''|*[!0-9]*) printf '%s\n' 'release build: VERSION has an invalid RC number' >&2; exit 1;; esac
    ;;
  *) printf '%s\n' 'release build: VERSION is not a QWSG 1.0 release identity' >&2; exit 1;;
esac
command -v go >/dev/null && command -v sha256sum >/dev/null && command -v tar >/dev/null || {
  printf '%s\n' 'release build: go, sha256sum and GNU tar are required' >&2; exit 1;
}
epoch=${SOURCE_DATE_EPOCH:-0}
case "$epoch" in ''|*[!0-9]*) printf '%s\n' 'release build: SOURCE_DATE_EPOCH must be an unsigned integer' >&2; exit 1;; esac
commit=${BUILD_COMMIT:-unknown}
case "$commit" in unknown) :;; *[!0-9a-fA-F]*|'') printf '%s\n' 'release build: BUILD_COMMIT must be hexadecimal or unknown' >&2; exit 1;; esac
build_date=$(date -u -d "@$epoch" '+%Y-%m-%dT%H:%M:%SZ')
out=${DIST_DIR:-"$repo/dist"}
name="qwsg-$version-linux-amd64"
work=$(mktemp -d /tmp/qwsg-release.XXXXXX)
trap 'rm -rf "$work"' EXIT HUP INT TERM
root="$work/$name"
mkdir -p "$root/bin" "$root/lib/systemd/user" "$root/docs" "$out"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOCACHE=${GOCACHE:-/tmp/qwsg-go-cache} GOMODCACHE=${GOMODCACHE:-/tmp/qwsg-go-modcache} \
  go build -trimpath -buildvcs=false -ldflags "-s -w -X main.version=$version -X main.buildCommit=$commit -X main.buildDate=$build_date" -o "$root/bin/qwsg" "$repo/cmd/qwsg"
cp "$repo/packaging/systemd/qwsg-guardian.service" "$root/lib/systemd/user/"
cp "$repo/packaging/release/install.sh" "$repo/packaging/release/uninstall.sh" "$root/"
cp "$repo/packaging/release/qwsg-config.json" "$root/"
cp "$repo/LICENSE" "$repo/CHANGELOG.md" "$root/"
release_notes="RELEASE_NOTES_$version"
for doc in QUICK_START SETUP_AND_CONFIGURATION OPERATIONS TROUBLESHOOTING UPGRADE_ROLLBACK_UNINSTALL SUPPORT SECURITY_AND_PRIVACY KNOWN_LIMITATIONS "$release_notes"; do cp "$repo/docs/release/$doc.md" "$root/docs/"; done
chmod 0755 "$root/bin/qwsg" "$root/install.sh" "$root/uninstall.sh"
find "$root" -type f ! -name MANIFEST.sha256 -printf '%P\n' | LC_ALL=C sort | while IFS= read -r file; do sha256sum "$root/$file"; done | sed "s#  $root/#  #" > "$root/MANIFEST.sha256"
find "$root" -exec touch -h -d "@$epoch" {} +
archive="$out/$name.tar.gz"
LC_ALL=C tar --sort=name --mtime="@$epoch" --owner=0 --group=0 --numeric-owner --format=ustar -C "$work" -cf - "$name" | gzip -n > "$archive"
(cd "$out" && sha256sum "$name.tar.gz") > "$archive.sha256"
printf '%s\n' "$archive"
