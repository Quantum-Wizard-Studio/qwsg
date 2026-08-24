# Canonical Engineering Git Policy

## Purpose

This policy defines the reusable Git safety boundary for engineering work. It
is mandatory for QWSG and for every project adopting the Reusable Engineering
Framework.

## Repository verification

Before changes, record the project root, current branch, HEAD, configured
remotes, local-to-remote relationship, tags relevant to the task, and complete
status. Compare the current remote and primary branch with the validated project
configuration. Stop on an unexplained difference.

QWSG uses:

- remote name `origin`;
- HTTPS remote
  `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg.git`;
- primary branch `main`.

HTTPS is the canonical QWSG Git transport. Forgejo SSH on port 2222 is not
required for normal QWSG engineering.

## Synchronization

Fetch before relying on remote state when network access and task authority
permit it. A fetch updates remote-tracking evidence but does not authorize
merge, rebase, reset, checkout, or working-tree changes. Record ahead/behind
counts and stop on unexpected divergence.

## Targeted staging

Stage only explicit reviewed paths belonging to the active task. Broad staging
such as `git add .`, `git add -A`, or `git add --all` is prohibited. Existing
untracked paths belong to their owner and must not be staged, moved, deleted,
ignored, or rewritten without explicit scope.

Before commit, review:

- unstaged and staged path lists;
- staged diff and `git diff --cached --check`;
- permissions and executable bits;
- secret and private-host evidence;
- generated or backup payload exclusions;
- task-specific tests and documentation.

## Commit and push

A commit requires completed mandatory validation, a truthful task-scoped
message, and an explicit reviewed path set. Do not create artificial historical
commits.

Before a real push, run a dry-run push against the configured remote and branch.
Push only the reviewed commit and record the command, result, commit hash, and
remote relationship. Automatic or implicit push is prohibited. An approved
Authority Envelope may explicitly authorize task-scoped targeted staging,
commit, push dry-run, clean fast-forward push, post-push verification, and
lifecycle closure as one routine engineering cycle; those operations then
require no additional micro-gate. Force push, history rewriting, tags,
Releases, publication, and deployment remain separately reserved unless
explicitly authorized.

## Branch and tag safety

Do not force-push, rewrite published history, rebase or squash published work,
delete branches or tags, or move an existing tag without explicit Project Owner
authority and a dedicated recovery plan. Temporary remote verification branches
must use unique names, be verified, and be removed with deletion verification.

## Prohibited recovery shortcuts

Broad reset, restore, checkout, or clean operations are not rollback. Use the
task's bounded snapshot and exact target list. Destructive operations require
resolved targets, impact review, explicit authority, and post-action validation.

## Completion evidence

Delivery records include starting and final Git state, exact changed and staged
paths, validation results, commit status, push status, unresolved differences,
and confirmation that unrelated untracked content remained untouched.
