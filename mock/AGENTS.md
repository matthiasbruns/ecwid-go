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
| `/api/v3/{storeId}/...` | Simulated Ecwid REST: app storage, store-profile + customer fixtures; proxy or `501` fallback for unimplemented routes | #5, #7 |

**REST fallback (`internal/server/proxy.go`).** Unimplemented REST routes return an
informative `501` naming the endpoint and the `--proxy-store`/`--proxy-token`
remedy. With proxying configured they forward to
`https://app.ecwid.com/api/v3/{proxyStore}/…`, rewriting the store ID and
swapping the `Authorization` bearer for the proxy token. Proxying uses a plain
`http.Client` (the `ecwid/internal/api` transport is JSON-oriented and its retry
transport is unexported — reusing it would widen that module's exported surface,
which the "CRUCIAL RULE" forbids). `--proxy-readonly` (default **true**) blocks
proxied mutations with `403`. **`/storage` is always served locally**, even when
proxying, so dev scratch state never lands in the real store's app storage —
enforced in the fallback handler and by ServeMux route specificity.

**Fixtures (`internal/server/fixtures.go`).** `GET /profile`, `GET /customers`
(paged + `?email=` filter), `GET /customers/{id}`, and `PUT /customers/{id}`
(field-merge, read-after-write) are served from an in-memory `fixtureStore`
keyed by store ID — the **default**, no proxy needed. Responses use the
`ecwid/customers` and `ecwid/profile` structs so shapes/field-names match the
typed client. Each route is gated on its Ecwid scope via `Config.HasScope`
(missing scope → the real `403` shape; empty `--scopes` grants all). The
configured store is seeded with a default profile + customers at `New()`;
`POST/PUT /_mock/fixtures/...` (control plane) lets an out-of-process consumer
install its own. These routes are **always local, never proxied**.
| `/_mock/...` | The mock's own control API (health, webhook trigger, fixtures, …) | #6 |

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
    ├── webhook/                # webhook trigger: compose, sign, deliver, classify
    │   ├── catalog.go          # the 42 event fixtures + per-family metadata
    │   ├── trigger.go          # Trigger: compose (ecwid/webhooks), sign, POST
    │   ├── classify.go         # Ecwid success/failure rules, with reasons
    │   ├── http.go             # control API handlers (trigger, events)
    │   ├── ui.go               # trigger panel handler
    │   └── templates/panel.html
    └── server/                 # ServeMux, request logging, graceful lifecycle
        ├── server.go           # New, routes, Run (graceful shutdown)
        ├── middleware.go       # structured slog request logging
        ├── health.go           # GET /_mock/health
        ├── storage.go          # app-storage REST endpoints (the JS SDK's only HTTP calls)
        ├── fixtures.go         # profile + customer fixtures: store, seed data, REST handlers, scope gate
        └── fixtures_control.go # POST/PUT /_mock/fixtures/... to seed/override fixtures
```

## Webhook trigger (control plane)

The mock's headline feature — Ecwid ships no webhook test tooling at all. All
under `/_mock/` (the control-plane namespace):

| Route | Purpose |
|-------|---------|
| `POST /_mock/webhooks/trigger` | Compose, sign, and deliver a webhook; returns the delivery result (status, latency, success/failure classification with reason, response body). **The primary surface** — CI integration tests drive this; the UI is a client of it. |
| `GET /_mock/webhooks/events` | The 42 event types with their group, `entityId` wire type, and fixture `data`. |
| `GET /_mock/webhooks/ui` | The developer-facing trigger panel (stdlib `html/template` + `go:embed`, no JS framework). |

Events are composed with `ecwid/webhooks` (types, `Event` marshaling, `Sign`,
`SuccessCodes`) so the mock exercises the exact code a real integration runs —
including the `entityId` quirk (a number for order/product families, a quoted
string for `application.*`) and the no-`data`-key families. The 24h/27-attempt
retry schedule is deliberately **not** implemented; report the outcome and let
the user re-fire.

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
| `--scopes` | `ECWID_MOCK_SCOPES` | *(all granted)* | Comma-separated granted scopes; gates the profile/customer fixtures. Empty = all granted; narrow it to test `403` paths |
| `--port` | `ECWID_MOCK_PORT` | `8080` | Listen port |
| `--proxy-store` / `--proxy-token` | `ECWID_MOCK_PROXY_*` | *(optional)* | Forward unimplemented REST to a real store (both required together) |
| `--proxy-readonly` | `ECWID_MOCK_PROXY_READONLY` | `true` | Only proxy `GET`/`HEAD`; mutations → `403`. Off = proxied writes hit the real store |

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
