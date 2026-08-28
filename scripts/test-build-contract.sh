#!/usr/bin/env bash
set -euo pipefail

repo=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd -P)
version=$(tr -d '\r\n' < "$repo/VERSION")
[[ "$version" == 1.2.0-rc.5 ]]
[[ -z ${GOFLAGS+x} ]] || {
    printf '%s\n' 'build contract: run with GOFLAGS unset' >&2
    exit 1
}

work=$(mktemp -d /tmp/qwsg-build-contract.XXXXXX)
cleanup() { rm -rf -- "$work"; }
trap cleanup EXIT HUP INT TERM
chmod 0700 "$work"

payload="$work/source.tar"
(
    cd "$repo"
    git ls-files -c | while IFS= read -r path; do
        [[ -f "$path" && ! -L "$path" ]] && printf '%s\n' "$path"
    done | tar -cf "$payload" -T -
)

for label in one two explicit; do
    mkdir -m 0700 "$work/$label" "$work/$label/source" \
        "$work/$label/gocache" "$work/$label/gomodcache"
    tar -xf "$payload" -C "$work/$label/source"
    [[ ! -e "$work/$label/source/.git" ]]
done

build_export() {
    local label=$1
    shift
    (
        cd "$work/$label/source"
        env -u GOFLAGS GOCACHE="$work/$label/gocache" \
            GOMODCACHE="$work/$label/gomodcache" make "$@" build
    )
}

build_export one
build_export two
cmp "$work/one/source/build/qwsg" "$work/two/source/build/qwsg"

expected_default=$(printf 'QWSG %s\ncommit: unknown\nbuilt: unknown' "$version")
[[ "$($work/one/source/build/qwsg version)" == "$expected_default" ]]

explicit_commit=0123456789abcdef0123456789abcdef01234567
explicit_date=2026-08-21T00:00:00Z
build_export explicit "BUILD_COMMIT=$explicit_commit" "BUILD_DATE=$explicit_date"
mkdir -m 0700 "$work/checkout-gocache" "$work/checkout-gomodcache"
env -u GOFLAGS GOCACHE="$work/checkout-gocache" \
    GOMODCACHE="$work/checkout-gomodcache" make -C "$repo" \
    "BUILD_COMMIT=$explicit_commit" "BUILD_DATE=$explicit_date" build
cmp "$repo/build/qwsg" "$work/explicit/source/build/qwsg"
expected_explicit=$(printf 'QWSG %s\ncommit: %s\nbuilt: %s' \
    "$version" "$explicit_commit" "$explicit_date")
[[ "$($work/explicit/source/build/qwsg version)" == "$expected_explicit" ]]

for binary in "$repo/build/qwsg" "$work/one/source/build/qwsg" \
    "$work/two/source/build/qwsg" "$work/explicit/source/build/qwsg"; do
    if go version -m "$binary" | grep -Eq $'^[[:space:]]*build[[:space:]]+vcs(\.|=)'; then
        printf '%s\n' 'build contract: ambient Go VCS metadata is present' >&2
        exit 1
    fi
done

grep -F -- '-buildvcs=false' "$repo/Makefile" >/dev/null
grep -F -- '-buildvcs=false' "$repo/scripts/build-release.sh" >/dev/null
printf '%s\n' 'PASS: checkout/export build provenance contract'
