# QWSG 1.1.0-rc.4 Private Acceptance Notes

This source identity prepares a future private replacement candidate for a
separately authorized clean-host acceptance. It does not create an artifact,
tag, release, publication or acceptance claim.

## Exported-source build determinism

- The canonical ordinary `make build` path explicitly disables Go automatic
  VCS stamping, matching the existing release-builder policy.
- Ordinary builds work in a normal checkout and in a genuine source export
  without `.git` metadata or a caller-supplied `GOFLAGS` workaround.
- Checkout builds retain their Git-derived development commit default. Exported
  builds report the truthful `unknown` commit/date defaults unless explicit
  build identity is supplied.
- Release builds still require the exact explicit full source commit and
  commit-derived `SOURCE_DATE_EPOCH`; ambient Git metadata never substitutes
  for controlled release provenance.

RC.1, RC.2 and the failed-gate RC.3 candidate/evidence remain immutable.
`QWSG-053-F001` is historical and blocking until this correction is integrated
and a new candidate passes the required build and acceptance gates. RC.4 must
be constructed later from one exact clean commit under separate Owner authority.
