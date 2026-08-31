#!/bin/sh
set -eu

repo=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd -P)
version=$(tr -d '\r\n' < "$repo/VERSION")
case "$version" in
  1.0.0|1.1.0|1.2.0) :;;
  1.0.0-rc.*|1.1.0-rc.*|1.2.0-rc.*)
    rc_number=${version##*-rc.}
    case "$rc_number" in ''|*[!0-9]*) printf '%s\n' 'release build: VERSION has an invalid RC number' >&2; exit 1;; esac
    ;;
  *) printf '%s\n' 'release build: VERSION is not an approved QWSG release identity' >&2; exit 1;;
esac
release_notes="$repo/docs/release/RELEASE_NOTES_$version.md"
test -f "$release_notes" || { printf '%s\n' 'release build: matching release notes are missing' >&2; exit 1; }
validate_only=${QWSG_RELEASE_VALIDATE_ONLY:-0}
case "$validate_only" in
  0) :;;
  1) printf '%s\n' "release build: identity $version is valid"; exit 0;;
  *) printf '%s\n' 'release build: QWSG_RELEASE_VALIDATE_ONLY must be 0 or 1' >&2; exit 1;;
esac
command -v go >/dev/null && command -v sha256sum >/dev/null && command -v tar >/dev/null || {
  printf '%s\n' 'release build: go, sha256sum and GNU tar are required' >&2; exit 1;
}
case "$version" in
  1.1.0|1.1.0-rc.*|1.2.0|1.2.0-rc.*)
    test "${SOURCE_DATE_EPOCH+x}" = x || { printf '%s\n' 'release build: QWSG 1.1+ requires explicit SOURCE_DATE_EPOCH' >&2; exit 1; }
    test "${BUILD_COMMIT+x}" = x || { printf '%s\n' 'release build: QWSG 1.1+ requires explicit BUILD_COMMIT' >&2; exit 1; }
    ;;
esac
epoch=${SOURCE_DATE_EPOCH:-0}
case "$epoch" in ''|*[!0-9]*) printf '%s\n' 'release build: SOURCE_DATE_EPOCH must be an unsigned integer' >&2; exit 1;; esac
commit=${BUILD_COMMIT:-unknown}
case "$commit" in unknown) :;; *[!0-9a-fA-F]*|'') printf '%s\n' 'release build: BUILD_COMMIT must be hexadecimal or unknown' >&2; exit 1;; esac
case "$version" in
  1.1.0|1.1.0-rc.*|1.2.0|1.2.0-rc.*)
    test "${#commit}" -eq 40 || { printf '%s\n' 'release build: QWSG 1.1+ requires the full 40-character commit' >&2; exit 1; }
    case "$commit" in *[!0-9a-f]*) printf '%s\n' 'release build: QWSG 1.1+ commit must be lowercase hexadecimal' >&2; exit 1;; esac
    ;;
esac
build_date=$(date -u -d "@$epoch" '+%Y-%m-%dT%H:%M:%SZ')
out=${DIST_DIR:-"$repo/dist"}
name="qwsg-$version-linux-amd64"
archive="$out/$name.tar.gz"
test ! -e "$archive" && test ! -L "$archive" && test ! -e "$archive.sha256" && test ! -L "$archive.sha256" || {
  printf '%s\n' 'release build: output archive or sidecar already exists' >&2; exit 1;
}
work=$(mktemp -d /tmp/qwsg-release.XXXXXX)
trap 'rm -rf "$work"' EXIT HUP INT TERM
root="$work/$name"
mkdir -p "$root/bin" "$root/lib/systemd/user" "$root/docs" "$out"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOCACHE=${GOCACHE:-/tmp/qwsg-go-cache} GOMODCACHE=${GOMODCACHE:-/tmp/qwsg-go-modcache} \
  go build -trimpath -buildvcs=false -ldflags "-s -w -X main.version=$version -X main.buildCommit=$commit -X main.buildDate=$build_date" -o "$root/bin/qwsg" "$repo/cmd/qwsg"
cp "$repo/packaging/systemd/qwsg-guardian.service" "$root/lib/systemd/user/"
cp "$repo/packaging/release/install.sh" "$repo/packaging/release/uninstall.sh" "$root/"
cp "$repo/packaging/release/qwsg-config.json" "$root/"
cp "$repo/internal/releasediscovery/trust/production.json" "$root/qwsg-release-trust.json"
cp "$repo/LICENSE" "$repo/CHANGELOG.md" "$root/"
printf '{"Schema":"qwsg.release/1","Version":"%s","Commit":"%s","Built":"%s","Platform":"linux-amd64"}\n' "$version" "$commit" "$build_date" > "$root/RELEASE.json"
cp "$repo/README.md" "$root/README.md"
cp "$repo/docs/installation/INSTALL.md" "$root/INSTALL.md"
release_notes_name="RELEASE_NOTES_$version"
for doc in QUICK_START SETUP_AND_CONFIGURATION OPERATIONS TROUBLESHOOTING UPGRADE_ROLLBACK_UNINSTALL CHANGE_NOTIFICATIONS SUPPORT SECURITY_AND_PRIVACY KNOWN_LIMITATIONS RELEASE_INDEX_PUBLICATION "$release_notes_name"; do cp "$repo/docs/release/$doc.md" "$root/docs/"; done
find "$root" -type f ! -name MANIFEST.sha256 -printf '%P\n' | LC_ALL=C sort | while IFS= read -r file; do sha256sum "$root/$file"; done | sed "s#  $root/#  #" > "$root/MANIFEST.sha256"
find "$root" -type d -exec chmod 0755 {} +
find "$root" -type f -exec chmod 0644 {} +
chmod 0755 "$root/bin/qwsg" "$root/install.sh" "$root/uninstall.sh"
find "$root" -exec touch -h -d "@$epoch" {} +
LC_ALL=C tar --sort=name --mtime="@$epoch" --owner=0 --group=0 --numeric-owner --format=ustar -C "$work" -cf - "$name" | gzip -n > "$archive"
(cd "$out" && sha256sum "$name.tar.gz") > "$archive.sha256"
printf '%s\n' "$archive"
