# QWSG Operator Console

Run `qwsg observe`. The first run establishes a private baseline and remains unknown; run it again later to perform the full operator evaluation. Then start `qwsg` in a separate process to view the qualified result. `qwsg check` remains the limited Inventory/Snapshot operation. `QWSG_STATE_DIR` selects a private absolute state root; `QWSG_STORE` may override the Inventory Store.

When an evaluation produces more attention facts than the bounded Console can
list, QWSG keeps the highest-severity and highest-importance facts and shows
how many additional facts were correlated or omitted. A current but partial
result recommends inspecting evidence instead of claiming that repetition will
repair it. Observation errors identify whether evaluation, projection, or
state publication failed without printing private host details.

Run `qwsg`. In a terminal it opens the read-only Console; in a pipe or script it prints one concise view and exits.

- `j` / `k`: move
- Enter: open
- `b`: back
- `r`: reload and requalify Current State without starting an observation
- `h`: help
- `q`: quit

Home explains server condition, attention, changes, Alerts, Guardian, evidence quality, and the recommended next step. IDs appear only in Details. Advanced commands and structured JSON remain available.

The condition covers implemented canonical comparison and engineering evidence only. It does not prove every possible operational concern. A successful one-shot observation is not continuous service evidence, so Guardian remains “not observed” until a real Runtime Service publishes current lifecycle evidence.

`r` is read-only: it does not acquire the Guardian operation lock, collect
Inventory, run Pipeline, or publish state. Use `qwsg observe` explicitly for a
manual active observation. Runtime failures are shown as bounded practical
reasons such as Alert evaluation, Notification planning or delivery, timeout,
or cancellation; private raw errors are never displayed.
