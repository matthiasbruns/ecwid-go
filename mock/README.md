# ecwid-mock

A local stand-in for an Ecwid store, so you can develop and test an embedded
Ecwid app without a real store. Ecwid renders apps in an iframe inside the store
admin and hands them context through a JS SDK — and provides **no** local tooling
and **no** webhook test tooling (no sender, no replay, no delivery log). This
module fills that gap:

- **Admin shell** — serves your app in an iframe with a correctly-built auth
  payload, exactly as Ecwid's admin does.
- **Simulated REST** — the app-storage endpoints the JS SDK calls over HTTP,
  plus built-in **store-profile and customer fixtures** so an app can be
  exercised without a real store, with an optional proxy to a real store for
  everything else.
- **Webhook trigger** — compose, sign, and deliver any of Ecwid's 42 event types,
  including deliberately-invalid signatures to prove your handler fails closed.

It pairs with two library packages you'll use in the app itself:
[`ecwid/appauth`](https://pkg.go.dev/github.com/matthiasbruns/ecwid-go/ecwid/appauth)
(decode the auth payload) and
[`ecwid/webhooks`](https://pkg.go.dev/github.com/matthiasbruns/ecwid-go/ecwid/webhooks)
(verify and dispatch webhooks).

> **Before you trust a single webhook or auth payload, read
> [Ecwid gotchas](#ecwid-gotchas).** Every item there is a real bug waiting to
> happen, and none is discoverable from Ecwid's current documentation.

## Quickstart

```bash
# Install (builds a binary named `mock`; rename or alias it as you like):
go install github.com/matthiasbruns/ecwid-go/mock@latest

# Or run from a checkout of this repo:
go run ./mock serve --app-url=http://localhost:3000
```

`--app-url` is the only required flag — the URL of the app you're developing,
already running locally. Then open the admin shell:

```
http://localhost:8080/
```

You'll see your app iframed with a store-context payload injected the same way
Ecwid injects it. The startup banner prints the store ID, auth mode, client ID,
and — when it generated one — the `client_secret` to configure your app with:

```
ecwid-mock
  admin:      http://localhost:8080/
  app URL:    http://localhost:3000
  store ID:   1003
  auth mode:  default
  client_id:  mock-app
  client_secret (generated): 3f9a...c1
  ^ configure your app with this secret; override with --client-secret
```

Health check: `curl localhost:8080/_mock/health` → `{"status":"ok"}`.

## The two auth modes — *which one do you actually have?*

This is the highest-value thing to get right, because the two modes are
completely different and Ecwid's current docs no longer distinguish them. Pick
with `--auth-mode`.

### `default` — Default User Auth (what every app gets)

The payload is **hex-encoded plaintext JSON in the URL fragment**:

```
https://your-app.example/iframe#7b2273746f72655f6964223a...
```

A fragment is **never sent to the server**, so your Go backend cannot see this
payload at all. It is read **client-side only**, via `EcwidApp.getPayload()`.
Decode it (in tests, or in tooling that receives the fragment out of band) with
[`appauth.DecodeHex`](https://pkg.go.dev/github.com/matthiasbruns/ecwid-go/ecwid/appauth#DecodeHex).

A hex payload is unauthenticated plaintext — no secret is involved. **Re-validate
its `access_token` against the Ecwid API before trusting anything it carries;**
do not take `store_id` or the tokens at face value.

### `enhanced` — Enhanced Security User Auth (opt-in, maybe unavailable)

The payload is **AES-128-CBC encrypted in the query string**, so a Go server
*can* act on it:

```
https://your-app.example/iframe?payload=<urlsafe-b64>&app_state=...&cache-killer=...
```

Decrypt it server-side with
[`appauth.Decrypt`](https://pkg.go.dev/github.com/matthiasbruns/ecwid-go/ecwid/appauth#Decrypt),
keyed by the first 16 bytes of your `client_secret`.

Two things will bite you:

- **`EcwidApp.getPayload()` returns `undefined` in enhanced mode.** The SDK
  hex-decodes the fragment; it never decrypts the query blob. Ecwid's own docs
  work around this by having the server decrypt the payload and inject it into
  the app's HTML. If you switch the mock to `enhanced` and your client-side code
  still calls `getPayload()`, it will get `undefined` — and you'll assume the
  mock is broken. It isn't; that's Ecwid's behavior.
- ⚠️ **Whether Ecwid still grants Enhanced Security User Auth to new apps is
  UNCONFIRMED.** The mode survives only in 2020-era documentation
  ([`Ecwid/ecwid-api-docs`](https://github.com/Ecwid/ecwid-api-docs)), gated
  behind emailing `ec.apps@lightspeedhq.com`. Do not architect around a mode you
  may not be able to get — verify with Ecwid first.

## Pointing the SDK at the mock

The SDK's `getAppStorage`/`setAppStorage` calls (and their public-config
variants) go to `app.ecwid.com` **by default**. Without redirection, your app's
storage calls bypass the mock entirely and hit the real Ecwid API. Redirect them
with either hook:

```js
// Cleanest: define this before the SDK loads. Return the mock's host.
window.getEcwidSdkApiDomain = function () {
  return "localhost:8080";
};
```

or rely on the undocumented `domain` payload field (SDK 1.3.1+), which overrides
the API host from within the payload itself. The admin shell surfaces both under
its **“Point the SDK's storage calls at the mock”** panel.

The mock authenticates storage calls with `Authorization: Bearer <access_token>`,
comparing against `--access-token` (constant-time). The shell injects a
deterministic `secret_mock_<storeId>_access` token into the payload, so set
`--access-token` to match if your app forwards the payload token as the bearer
(see the [worked example](#worked-example-ecwidsample-native-app)).

## Webhook testing

Ecwid ships no webhook test tooling, which is the headline reason this module
exists. Three ways in, all sharing one signing path (`webhooks.Sign`) so what the
mock sends is exactly what `webhooks.Verify` checks:

- **UI panel** — `http://localhost:8080/_mock/webhooks/ui`. Pick an event, edit
  the `entityId`/`data`, choose a signature mode, fire, and read the delivery
  verdict.
- **Control API** — `POST /_mock/webhooks/trigger`, the surface CI integration
  tests drive:

  ```bash
  curl -sX POST localhost:8080/_mock/webhooks/trigger \
    -H 'Content-Type: application/json' \
    -d '{"eventType":"order.created","signature":"valid"}'
  ```

  Returns the composed event plus the delivery result — status, latency,
  `delivered` verdict with a reason, and the response body. Set `--webhook-url`
  to your handler's endpoint first, or the call returns `409`.

- **Event catalog** — `GET /_mock/webhooks/events` lists all 42 event types with
  their group, `entityId` wire type, and fixture `data`.

### Prove your handler fails closed

The `signature` field takes three modes. Fire the bad ones and confirm your
handler **rejects** them:

| `signature` | What it sends | A correct handler… |
|-------------|---------------|--------------------|
| `valid` | HMAC under your `client_secret` | accepts (2xx) |
| `invalid` | a well-formed HMAC under the *wrong* key | rejects (401) |
| `missing` | no signature header at all | rejects (401) — **this is the bug Ecwid's own PHP sample ships** |

If your endpoint returns success for `invalid` or `missing`, it is not verifying.
`webhooks.NewHandler` gets this right; a hand-rolled handler often does not.

## Simulated REST fixtures — run without a real store

Beyond app storage, the mock serves a set of read/write REST endpoints from
in-memory **fixtures**. These are the **default** — no real store and no proxy
are needed — so a consuming app (and its integration tests) can be driven
hermetically. Responses match real Ecwid shapes and field names; they are built
from the [`ecwid/customers`](https://pkg.go.dev/github.com/matthiasbruns/ecwid-go/ecwid/customers)
and [`ecwid/profile`](https://pkg.go.dev/github.com/matthiasbruns/ecwid-go/ecwid/profile)
types, so the mock and the typed client agree by construction.

| Method & path | Scope | Behaviour |
|---------------|-------|-----------|
| `GET /api/v3/{storeId}/profile` | `read_store_profile` | The store profile. |
| `GET /api/v3/{storeId}/customers` | `read_customers` | Paged list (`items`/`count`/`total`/`offset`/`limit`); `?offset=`/`?limit=` window it, `?email=` filters to the case-insensitive match (the *find-by-email* path). |
| `GET /api/v3/{storeId}/customers/{id}` | `read_customers` | A single customer, or `404`. |
| `PUT /api/v3/{storeId}/customers/{id}` | `update_customers` | Merges the supplied fields (e.g. `acceptMarketing`) into the stored customer and returns `{"updateCount":1}`. The write is reflected in subsequent `GET`s (read-after-write). |

Requests carry the mock's `access_token` as `Bearer` (a wrong/missing token is
`401`), exactly like the storage routes. **These endpoints are always served
locally and never proxied**, even with proxying enabled.

**Seeded defaults.** On startup the configured store is seeded with a store
profile and three customers (`ada@example.com`, `grace@example.com`,
`alan@example.com`), so the happy path answers out of the box:

```bash
curl -H "Authorization: Bearer $TOKEN" localhost:8080/api/v3/1003/profile
curl -H "Authorization: Bearer $TOKEN" "localhost:8080/api/v3/1003/customers?email=grace@example.com"
```

### Scope enforcement — testing the unhappy path

By default every scope is granted. Narrow the grant with `--scopes` (comma
-separated) to make an endpoint return the same `403` a real store returns when
its token lacks the scope:

```bash
go run ./mock serve --app-url=http://localhost:3000 \
  --scopes=read_store_profile,read_customers      # update_customers now → 403
```

```json
{"errorMessage":"This method requires the 'update_customers' access scope, which the access token was not granted."}
```

### Seeding your own fixtures

A control API under `/_mock/fixtures/` lets an out-of-process consumer (e.g. an
E2E suite that boots the mock as a subprocess) install its own fixtures before
driving the endpoints above. It needs no bearer token. Add `?storeId=` to target
a specific store (default: the configured store), which is how a multi-tenant
consumer seeds several stores against one mock.

```bash
# Upsert one customer (or POST a JSON array to seed many). A missing id is assigned.
curl -X POST localhost:8080/_mock/fixtures/customers \
  -d '{"name":"Katherine Johnson","email":"katherine@example.com","acceptMarketing":false}'
# -> {"id":1004}

# Replace the store profile.
curl -X PUT localhost:8080/_mock/fixtures/profile \
  -d '{"settings":{"storeName":"Acme Test Store"}}'
```

The seeded customer is then findable through the simulated REST endpoints
(`GET /customers?email=…`), and its `acceptMarketing` can be flipped through the
real `PUT /customers/{id}` above.

## Proxy mode — forward the rest to a real store

The mock implements the app-storage, profile, and customer endpoints locally.
Every other REST route returns an informative `501` unless you enable proxying,
which forwards unimplemented routes to a **real** store:

```bash
go run ./mock serve --app-url=http://localhost:3000 \
  --proxy-store=12345 --proxy-token=secret_abc...
```

- **Read-only by default.** Only `GET`/`HEAD` are forwarded; write methods return
  `403`. This is deliberate — a proxied write mutates a **real** store.
- **`--proxy-readonly=false` is dangerous.** Proxied `POST`/`PUT`/`DELETE` hit
  the real store *and fire real webhooks from it*. The startup banner shouts this
  in a prominent warning block.
- **`/storage` is always served locally**, even with proxying on, so dev scratch
  state never lands in the real store's app storage.

## Flags

Precedence is **flags > env > defaults**.

| Flag | Env | Default | Purpose |
|------|-----|---------|---------|
| `--app-url` | `ECWID_MOCK_APP_URL` | *(required)* | URL of the app to iframe (must be `http`/`https`) |
| `--client-id` | `ECWID_MOCK_CLIENT_ID` | `mock-app` | App `client_id`; also `EcwidApp.init({app_id})` |
| `--client-secret` | `ECWID_MOCK_CLIENT_SECRET` | *(generated)* | Signs webhooks + encrypts payloads; **≥16 bytes** |
| `--store-id` | `ECWID_MOCK_STORE_ID` | `1003` | Store ID placed in the payload |
| `--auth-mode` | `ECWID_MOCK_AUTH_MODE` | `default` | `default` (hex fragment) \| `enhanced` (AES query) |
| `--access-token` | `ECWID_MOCK_ACCESS_TOKEN` | *(generated)* | `access_token` required as `Bearer` on simulated REST calls |
| `--scopes` | `ECWID_MOCK_SCOPES` | *(all granted)* | Comma-separated access scopes the token is granted; narrow it to test scope-denied (`403`) paths |
| `--webhook-url` | `ECWID_MOCK_WEBHOOK_URL` | *(optional)* | Where triggered webhooks POST |
| `--port` | `ECWID_MOCK_PORT` | `8080` | Listen port |
| `--proxy-store` | `ECWID_MOCK_PROXY_STORE` | *(optional)* | Store ID to forward unimplemented REST calls to |
| `--proxy-token` | `ECWID_MOCK_PROXY_TOKEN` | *(optional)* | Access token for the proxy store |
| `--proxy-readonly` | `ECWID_MOCK_PROXY_READONLY` | `true` | Only proxy `GET`/`HEAD`; block proxied mutations |

`--proxy-store` and `--proxy-token` must be set together. `--client-secret` is
validated at startup — anything under 16 bytes cannot derive the AES-128 key. A
generated `client_secret` is printed in the banner; a user-supplied one never is.

## What the mock does *not* simulate

It is a developer tool, not a store emulator. It intentionally omits:

- **No product / order / category fixtures.** The REST endpoints served locally
  are app storage plus the store-profile and customer fixtures above. Everything
  else is `501` (or proxied).
- **Scope enforcement is limited to the fixture endpoints.** The profile and
  customer routes honour their required scopes (see `--scopes`); storage and
  other routes are not scope-gated.
- **No webhook retry schedule.** A real failed delivery is retried 27 times over
  24h (see the gotchas). The mock fires once and reports the outcome — re-fire
  manually to test your retry handling.
- **No persistence.** App storage lives in memory and resets on every restart.
- **Not the real JS SDK.** The mock builds the payload and serves the storage
  API; the SDK itself (`ecwid-app.js`) is loaded by *your* app from Ecwid's CDN.

## Worked example: `Ecwid/sample-native-app`

[`Ecwid/sample-native-app`](https://github.com/Ecwid/sample-native-app) is a
pure client-side app (static `appProto.html` + `functions.js`, no backend). It
calls `EcwidApp.init({app_id: "sample-native-app"})`, reads the store context via
`EcwidApp.getPayload()`, and persists settings with
`EcwidApp.getAppStorage`/`setAppStorage`. That makes it a clean **default-mode**
demo.

```bash
# 1. Serve the sample app (any static file server works).
git clone https://github.com/Ecwid/sample-native-app.git
cd sample-native-app && python3 -m http.server 3000 &

# 2. Run the mock against it.
#    --client-id MUST equal the app's EcwidApp.init({app_id}).
#    --access-token is aligned with the payload token the shell injects
#    (secret_mock_<storeId>_access), so the SDK's storage calls authenticate.
go run ./mock serve \
  --app-url=http://localhost:3000/appProto.html \
  --client-id=sample-native-app \
  --store-id=1003 \
  --access-token=secret_mock_1003_access

# 3. Open http://localhost:8080/ — the app is iframed with a hex-fragment payload.
```

The sample app's flow, and how the mock answers each call:

1. `EcwidApp.getPayload()` reads the hex fragment → `{store_id: 1003, lang, access_token, ...}`.
2. `EcwidApp.getAppStorage('installed', cb)` → `GET /api/v3/1003/storage/installed`.
   On a fresh store the mock returns **404**, which the SDK maps to `null`, so the
   app runs its new-user path (`createUserData`).
3. `setAppStorage` / `setAppPublicConfig` → `PUT /api/v3/1003/storage/{key}` with
   the **raw value string** as the body. The mock stores it verbatim and returns
   `{"success": true}`.
4. `getAppStorage()` with no key → `GET /api/v3/1003/storage` returns the full
   entry array.

You can reproduce every one of those calls without a browser:

```bash
TOKEN=secret_mock_1003_access
# New user: 404 -> SDK sees null
curl -s -o /dev/null -w '%{http_code}\n' -H "Authorization: Bearer $TOKEN" \
  localhost:8080/api/v3/1003/storage/installed              # -> 404
# Save a private value (raw body, NOT {"value": ...})
curl -s -X PUT --data-raw 'true' -H "Authorization: Bearer $TOKEN" \
  localhost:8080/api/v3/1003/storage/installed             # -> {"success":true}
# Read it back
curl -s -H "Authorization: Bearer $TOKEN" \
  localhost:8080/api/v3/1003/storage/installed             # -> {"key":"installed","value":"true"}
```

> **Note on scope of verification.** The commands above — the shell render with a
> correct hex payload, the full storage lifecycle, and the webhook fail-closed
> loop (`valid`→accepted, `invalid`/`missing`→rejected against a real
> `webhooks.NewHandler` receiver) — were run and confirmed against the actual
> `sample-native-app` files. Driving `ecwid-app.js` itself inside a headless
> browser (to observe `getPayload`/`getAppStorage` from JS) was **not** performed
> here; the SDK is loaded by the app from Ecwid's CDN at runtime.

## Ecwid gotchas

Each of these is a real bug waiting to happen, and none is discoverable from
Ecwid's current documentation. Sources are linked; where the only source is the
legacy docs repo, that is called out.

### Webhooks

- **The signature does not cover the body.** It is
  `base64(HMAC-SHA256("<eventCreated>.<eventId>", client_secret))` — an HMAC over
  **only** those two fields. `storeId`, `eventType`, `entityId`, and all of `data`
  are **unauthenticated and fully replayable** (the signature never expires).
  Defend in depth: **dedupe on `eventId`**, **enforce a recency window on
  `eventCreated`**, and **re-fetch the entity via REST** rather than trusting
  `data`. `webhooks.Handler` wires up the first two (`Options.Deduper`,
  `Options.MaxAge`); the re-fetch is yours.
  ([Ecwid webhooks docs](https://docs.ecwid.com/api-reference/objects-events-formats/webhooks); the body-coverage gap is not stated there — it's observable from the algorithm.)

- **Ecwid's own PHP verification example fails open.** It reads
  `$signatureHeaderPresent` *before* assigning it, so a request carrying **no
  signature header** silently skips verification and is accepted. Do not port it.
  `webhooks.Verify` fails closed: an empty signature is rejected, never skipped.
  Prove it with `--signature=missing` in the trigger.

- **`entityId` changes JSON type by event family.** It is a **number** for
  order/product events (`"entityId":103878161`) and a **quoted string** for
  `application.*` events (`"entityId":"1003"`). A decoder that assumes one type
  breaks on the other. `webhooks.Event` normalizes both to a string.

- **Order events' `entityId` is the *internal* order ID**, which the REST order
  endpoints do **not** accept. The usable, human-readable ID is in the payload as
  `OrderData.OrderID`. Re-fetch order events by that, not by `EntityID`.

- **Success is not “any 2xx”.** Ecwid counts **`200`, `201`, `202`, `204`, `209`**
  as delivered. **`203` and `208` are failures** despite being 2xx, and **every
  3xx is a failure** — an endpoint that `301`-redirects HTTP→HTTPS *silently never
  receives the webhook*. Timeouts are **3s connect / 10s response**. A failed
  delivery is retried **27 times over 24h**; **two weeks of failure blocks your
  app's webhooks entirely.** `webhooks.SuccessCodes()` is the canonical set; the
  trigger classifies deliveries by exactly these rules.

- **Webhooks fire regardless of source.** An order created through the REST API
  still emits `order.created`. An integration that writes back to Ecwid in
  response to a webhook can **trigger itself in a loop** — guard sync integrations
  against re-entrancy.

- **Ecwid provides no webhook test tooling** — no sender, no replay, no delivery
  log. That absence is the entire reason this module exists.

### App storage

- **The `PUT` body is the raw value string**, not a `{"value": "..."}` wrapper.
  Send the value directly; the mock (and Ecwid) store it verbatim.

- **`GET` on a missing key returns `404`**, which the SDK maps to `null`. The mock
  returns a real 404 (never a 200 with an empty value) so your new-vs-returning
  user branch behaves identically to production.

### Auth payloads (enhanced mode)

- **The AES docs' example string is hex, not base64** — it is **not** a usable
  test vector. Decoding it as base64 yields garbage. (Legacy
  [`ecwid-api-docs`](https://github.com/Ecwid/ecwid-api-docs) `_add_to_cp.md`.)

- **Ecwid's official C# decrypt sample is wrong.** It passes the whole blob
  (IV included) as ciphertext and then strips leading bytes; it only *appears* to
  work because CBC self-synchronizes and corrupts just the first block. Slice the
  IV off properly (`appauth.Decrypt` does). Its base64-padding line also appends
  four `=` when the length is already a multiple of 4 — wrong; pad to the next
  multiple of 4, adding **nothing** when already aligned.

- **The two-mode distinction (`default` vs `enhanced`) is gone from the current
  docs.** It survives only in the legacy
  [`ecwid-api-docs`](https://github.com/Ecwid/ecwid-api-docs) repo (last pushed
  2020). If you only read `docs.ecwid.com`, you will not know enhanced mode
  exists — or that you almost certainly don't have it.

## For maintainers: reading Ecwid's docs

- `docs.ecwid.com` serves **raw Markdown** by appending `.md` to any URL, and
  indexes every page at [`llms.txt`](https://docs.ecwid.com/llms.txt).
- `developers.ecwid.com` `301`-redirects to `docs.ecwid.com`.
- The legacy [`Ecwid/ecwid-api-docs`](https://github.com/Ecwid/ecwid-api-docs)
  repo (last pushed 2020) is stale but is the **only** source documenting the two
  auth modes — that distinction was dropped from the current docs.

## Development

```bash
task mock:lint     # golangci-lint
task mock:test     # go test ./... -race
task mock:tidy     # go mod tidy
```

The mock ships a binary, not a library — everything lives under `internal/`, so
it has zero exported API surface. See [`AGENTS.md`](./AGENTS.md) for the route
namespace conventions and module layout.
