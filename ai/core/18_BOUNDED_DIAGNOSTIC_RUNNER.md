# Bounded Diagnostic Runner Standard

`ai/scripts/bounded-diagnostic-runner.sh` packages multiple diagnostics into one
reviewable operation. It verifies a data-only manifest and the SHA-256 identity
of an executable runner, imposes a timeout and output limit, sanitizes the
environment, accepts only privacy-safe machine-readable output, requires an
explicit failure classification, and removes its private capture file.

The manifest is UTF-8 tab-separated data with exactly these keys:

```text
schema	framework.bounded-diagnostic/1
mode	read-only
targets	service_state,operator_state
actions	inspect,compare
timeout_seconds	60
max_output_bytes	8192
runner_sha256	<64 lowercase hexadecimal characters>
```

`mode` is `read-only` or `bounded-mutation`. The declaration documents
authority; it is not a security sandbox. The runner must already be reviewed
against the task envelope. Targets and actions are comma-separated logical
identifiers, not private hostnames or shell code. Together with the verified
runner hash they make the operation set reviewable and fixed. External or
mutating use still requires the relevant task authority.

The diagnostic emits only unique `lowercase.dotted_key=value` records. Values
are short printable tokens and must not contain addresses, secrets, raw server
responses, or private identities. It must emit:

```text
diagnostic.classification=EXPECTED_BEHAVIOR
diagnostic.cleanup=PASS
```

The other permitted classifications are `PRODUCT_FRAMEWORK_DEFECT`,
`TEST_OR_ACCEPTANCE_DEFECT`, `ENVIRONMENTAL_ISSUE`, and `INCONCLUSIVE`.
Runner exit zero, valid bounded output, a classification, and cleanup PASS are
all required. The wrapper reports verified runner identity and execution mode;
it never evaluates manifest or runner output as shell code.

Prefer one such runner over a sequence of Owner copy/paste fragments. Each
check must answer a stated question, avoid credentials, and reuse prior evidence
not invalidated by a relevant mutation.
