#!/usr/bin/env bash
set -euo pipefail

repo=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd -P)
work=$(mktemp -d /tmp/qwsg-release-reproducibility.XXXXXX)
cleanup() { rm -rf -- "$work"; }
trap cleanup EXIT HUP INT TERM
chmod 0700 "$work"

payload="$work/source.tar"
version=$(tr -d '\r\n' < "$repo/VERSION")
(
    cd "$repo"
    { git ls-files; printf '%s\n' "docs/release/RELEASE_NOTES_$version.md"; } |
        LC_ALL=C sort -u | tar -cf "$payload" -T -
)

for label in normalized group_writable; do
    mkdir -m 0700 "$work/$label" "$work/$label/source" "$work/$label/dist" \
        "$work/$label/gocache" "$work/$label/gomodcache"
    tar -xf "$payload" -C "$work/$label/source"
done

find "$work/normalized/source" -type d -exec chmod 0755 {} +
find "$work/normalized/source" -type f -exec chmod 0644 {} +
chmod 0755 "$work/normalized/source/scripts/build-release.sh" \
    "$work/normalized/source/packaging/release/install.sh" \
    "$work/normalized/source/packaging/release/uninstall.sh"

find "$work/group_writable/source" -type d -exec chmod 0775 {} +
find "$work/group_writable/source" -type f -exec chmod 0664 {} +
chmod 0775 "$work/group_writable/source/scripts/build-release.sh" \
    "$work/group_writable/source/packaging/release/install.sh" \
    "$work/group_writable/source/packaging/release/uninstall.sh"

commit=0123456789abcdef0123456789abcdef01234567
epoch=1787904000
(
    umask 0022
    cd "$work/normalized/source"
    env -u GOFLAGS SOURCE_DATE_EPOCH=$epoch BUILD_COMMIT=$commit \
        DIST_DIR="$work/normalized/dist" GOCACHE="$work/normalized/gocache" \
        GOMODCACHE="$work/normalized/gomodcache" ./scripts/build-release.sh >/dev/null
)
(
    umask 0002
    cd "$work/group_writable/source"
    env -u GOFLAGS SOURCE_DATE_EPOCH=$epoch BUILD_COMMIT=$commit \
        DIST_DIR="$work/group_writable/dist" GOCACHE="$work/group_writable/gocache" \
        GOMODCACHE="$work/group_writable/gomodcache" ./scripts/build-release.sh >/dev/null
)

archive="qwsg-$version-linux-amd64.tar.gz"
cmp "$work/normalized/dist/$archive" "$work/group_writable/dist/$archive"
test "$(sha256sum "$work/normalized/dist/$archive" | cut -d' ' -f1)" = \
    "$(sha256sum "$work/group_writable/dist/$archive" | cut -d' ' -f1)"

tar -tvzf "$work/normalized/dist/$archive" | awk '
    $1 ~ /^d/ { if ($1 != "drwxr-xr-x") exit 1; next }
    $1 ~ /^-/ {
        path=$NF
        expected="-rw-r--r--"
        if (path ~ /\/bin\/qwsg$/ || path ~ /\/install\.sh$/ || path ~ /\/uninstall\.sh$/)
            expected="-rwxr-xr-x"
        if ($1 != expected) exit 1
        next
    }
    { exit 1 }
'

printf '%s\n' 'PASS: release artifacts are reproducible across umask and source-mode variations'
