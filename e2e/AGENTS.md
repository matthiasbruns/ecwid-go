# AGENTS.md — e2e

End-to-end tests running against a real Ecwid store.

## Gate

All tests require `ECWID_E2E=1` — they are skipped entirely otherwise.

## Secrets needed

| Variable | Description |
|----------|-------------|
| `ECWID_STORE_ID` | Ecwid store ID |
| `ECWID_TOKEN` | API access token with full scopes |

## Rules

- **Idempotent**: create → verify → delete. Leave the store clean.
- **Respect rate limits**: client is configured with `MaxRetries: 3`.
- **One file per domain**: `dictionaries_test.go`, `reports_test.go`, etc.
- **Read-only tests first**: dictionaries, reports don't mutate state.
- **CRUD tests**: create test data with unique names, clean up in `t.Cleanup`.

## Adding tests for new domains

When a new domain package lands (e.g., `ecwid/products/`):
1. Add `e2e/<domain>_test.go`
2. Add CRUD test using `testClient.<Domain>.*`
3. Update the checklist in issue #17
