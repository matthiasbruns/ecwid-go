# mock/ — Ecwid local mock server

## Overview

`ecwid-mock` is an **app developer's** tool: a local stand-in for an Ecwid store
so a Go developer can exercise an embedded app without a real store. Ecwid
renders apps in an iframe inside the store admin and injects context via a JS
SDK; it provides **no** local tooling and **no** webhook test tooling (no test
sender, no replay, no delivery log). This module fills that gap.

It is a standalone Go module wired into `go.work`, alongside `config/`,
`ecwid/`, `cli/`, and `e2e/`.

## ⚠️ Route namespace convention — the contract other issues build on

The HTTP surface is partitioned into three namespaces. **Do not cross these
lines.** Later issues (#4 shell, #5 storage, #6 trigger, #7 proxy) plug into
exactly one namespace each.

| Prefix | Purpose | Owner issue |
|--------|---------|-------------|
| `/` | Admin shell — the developer-facing UI that hosts the app iframe | #4 |
| `/api/v3/{storeId}/...` | Simulated Ecwid REST; proxy or `501` fallback for unimplemented routes | #5, #7 |
| `/_mock/...` | The mock's own control API (health, webhook trigger, …) | #6 |

The `/_mock/` prefix is reserved so the mock's control plane can **never collide
with a real Ecwid REST route** (real routes live under `/api/v3/`). Register
control-plane endpoints only under `/_mock/`.

Routing uses `net/http.ServeMux` with Go 1.22+ method+pattern syntax
(`"GET /_mock/health"`, `"POST /_mock/webhooks/trigger"`). Register routes in
`server.routes()` (`internal/server/server.go`), keeping one handler file per
concern under `internal/server/`.

## Structure

```
mock/
├── main.go                     # entry point, slog setup, version injection
├── cmd/
│   ├── root.go                 # cobra root (ecwid-mock), --version
│   └── serve.go                # `ecwid-mock serve` + flags, config load, banner
└── internal/
    ├── config/                 # mock Config: flags > env > defaults, validation
    │   └── config.go
    └── server/                 # ServeMux, request logging, graceful lifecycle
        ├── server.go           # New, routes, Run (graceful shutdown)
        ├── middleware.go       # structured slog request logging
        ├── health.go           # GET /_mock/health
        └── storage.go          # app-storage REST endpoints (the JS SDK's only HTTP calls)
```

Everything lives under `internal/` — this module ships a binary, not a library,
so it has **zero exported API surface** to keep as a compatibility contract.

## Configuration

The mock's config is **local to this module** and deliberately distinct from
`config.Config` (which carries `StoreID`/`Token` for real API calls). Do **not**
widen `config.Config` with these developer-tool fields — that module is a
published compatibility contract.

Precedence is **flags > env > defaults**, matching `config/`'s behavior.

| Flag | Env | Default | Purpose |
|------|-----|---------|---------|
| `--app-url` | `ECWID_MOCK_APP_URL` | *(required)* | URL of the app to iframe |
| `--client-id` | `ECWID_MOCK_CLIENT_ID` | `mock-app` | App `client_id`; also `EcwidApp.init({app_id})` |
| `--client-secret` | `ECWID_MOCK_CLIENT_SECRET` | *(generated)* | Signs webhooks + encrypts payloads; **≥16 bytes** |
| `--store-id` | `ECWID_MOCK_STORE_ID` | `1003` | Store ID in the payload |
| `--auth-mode` | `ECWID_MOCK_AUTH_MODE` | `default` | `default` (hex fragment) \| `enhanced` (AES query) |
| `--webhook-url` | `ECWID_MOCK_WEBHOOK_URL` | *(optional)* | Where triggered webhooks POST |
| `--access-token` | `ECWID_MOCK_ACCESS_TOKEN` | *(generated)* | `access_token` issued in the payload; required as `Bearer` on REST calls |
| `--port` | `ECWID_MOCK_PORT` | `8080` | Listen port |
| `--proxy-store` / `--proxy-token` | `ECWID_MOCK_PROXY_*` | *(optional)* | Forward unimplemented REST to a real store |

`--client-secret` **must be ≥16 bytes** — `appauth.Encrypt` derives an AES-128
key from `client_secret[:16]`. It is validated at startup (referencing
`appauth.ErrShortSecret`), never at request time. When no secret is supplied one
is generated (≥16 bytes) and **printed in the startup banner** so the developer
can configure their app to match; a **user-supplied** secret is never printed.

## Dependencies

`config` + `ecwid` (local `replace` during dev, matching `ecwid/go.mod`) +
**cobra**. Cobra is the repo's one blessed external dep. **No other external
deps** — `net/http`, `crypto/rand`, `log/slog`, stdlib only.

## Security

- **NEVER** log `--client-secret`, `access_token`, or any credential at any
  level. Request logging records only method, path, status, bytes, duration, and
  remote address — no bodies, headers, or query strings.
- Redact secrets for diagnostics with `config.RedactedClientSecret()` /
  `RedactedProxyToken()`, which reuse `config.RedactedToken()`.

## Commands

```bash
task mock:lint     # golangci-lint
task mock:test     # go test ./... -race
task mock:tidy     # go mod tidy

# Run it:
go run ./mock serve --app-url=http://localhost:3000
curl localhost:8080/_mock/health   # -> 200 {"status":"ok"}
```

## Rules

- All commands use `RunE` (return errors, don't `os.Exit`).
- `http.Server` sets `ReadHeaderTimeout` (gosec/golangci require it).
- Graceful shutdown on SIGINT/SIGTERM via `signal.NotifyContext`.
- Version is injected at build time via `-ldflags "-X main.version=..."`.
- Conventional commits, signed. See root `AGENTS.md`.
