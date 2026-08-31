#!/usr/bin/env bash
set -euo pipefail

repo=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd -P)
work=$(mktemp -d /tmp/qwsg-release-authority-reproducibility.XXXXXX)
cleanup() { rm -rf -- "$work"; }
trap cleanup EXIT HUP INT TERM
chmod 0700 "$work"

for label in one two; do
    mkdir -m 0700 "$work/$label" "$work/$label/gocache"
    env -u GOFLAGS GOCACHE="$work/$label/gocache" GOMODCACHE="${GOMODCACHE:-/tmp/qwsg-go-modcache}" \
        go build -trimpath -buildvcs=false -o "$work/$label/qwsg-release-index" "$repo/cmd/qwsg-release-index"
    env -u GOFLAGS GOCACHE="$work/$label/gocache" GOMODCACHE="${GOMODCACHE:-/tmp/qwsg-go-modcache}" \
        GOOS=windows GOARCH=amd64 go build -trimpath -buildvcs=false \
        -o "$work/$label/qwsg-release-sign-offline.exe" "$repo/cmd/qwsg-release-sign-offline"
done

cmp "$work/one/qwsg-release-index" "$work/two/qwsg-release-index"
cmp "$work/one/qwsg-release-sign-offline.exe" "$work/two/qwsg-release-sign-offline.exe"

candidate="$repo/internal/releasepublication/testdata/unsigned-candidate.json"
"$work/one/qwsg-release-index" generate "$candidate" "$work/one/signing-input.json"
"$work/two/qwsg-release-index" generate "$candidate" "$work/two/signing-input.json"
cmp "$work/one/signing-input.json" "$work/two/signing-input.json"
test "$(stat -c %a "$work/one/signing-input.json")" = 600

printf '%s\n' 'PASS: release-authority tools and canonical signing input are reproducible'
