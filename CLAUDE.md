# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Official BitMart Cloud API Go SDK (`github.com/bitmartexchange/bitmart-go-sdk-api`, package `bitmart`). It wraps the BitMart Spot, Margin, and Futures REST APIs plus the Spot/Futures WebSocket streams. The implementation mirrors the online API reference at https://developer-pro.bitmart.com/en/spot (Spot) and https://developer-pro.bitmart.com/en/futures (Futures) — endpoint paths, params, and auth levels in the code should always match those docs.

Go 1.15, single external dependency: `github.com/gorilla/websocket`.

## Commands

```bash
go build ./...                       # build everything
go vet ./...                         # static checks
go test ./...                        # run all tests
go test -run TestGetSpotV3Ticker .   # run a single test (tests live in the root package)
go test -v -run TestGetSpotSymbol .  # verbose single test
```

Note: most tests in `*_test.go` hit the **live** BitMart API. They use `NewTestClient()` / `NewTestFuturesClient()` from `cloud_client_test.go`, which carry placeholder credentials (`"Your API KEY"` etc.). Public-data tests work as-is; SIGNED tests will return auth errors unless you edit `GetDefaultConfig` with real keys. Tests assert little — they mostly `PrintResponse` for manual inspection. There is no mock server.

## Architecture

The SDK is a flat package at the repo root (no internal sub-packages). Two client types, one shared request/signing core.

### REST: `CloudClient`
- `NewClient(Config)` → `*CloudClient`. `Config` (in `cloud_consts.go`) holds `Url`, `WsUrl`, `ApiKey`, `SecretKey`, `Memo`, `TimeoutSecond`, `CustomLogger`, and custom `Headers`.
- **Spot defaults to `API_URL_PRO`** (`api-cloud.bitmart.com`); **Futures requires `Url: API_URL_V2_PRO`** (`api-cloud-v2.bitmart.com`) — this must be set explicitly in the config or futures calls hit the wrong host.
- `cloud_client.go` is the heart: `Request()` builds the URL, attaches auth headers, sends, and fills a `CloudResponse{HttpStatus, Response string, Limit RateLimit}`. The raw JSON body is returned as a **string** in `.Response` — this SDK does **not** unmarshal responses into typed structs. Rate-limit headers (`X-BM-RateLimit-*`) are parsed into `.Limit`.
- GET/DELETE params are sorted and appended as a query string (`CreateQueryString`); POST params are JSON-marshalled into the body (`ParseRequestParams`).

### Auth levels (`Auth` enum: `NONE`, `KEYED`, `SIGNED`)
Set per-endpoint as the last arg to the request helpers, and reflected in each method's doc comment:
- `NONE` — public, no headers.
- `KEYED` — adds `X-BM-KEY` (API key only).
- `SIGNED` — adds key + `X-BM-TIMESTAMP` + `X-BM-SIGN`. The signature is `HmacSha256` (hex) over `PreHashString = timestamp + "#" + memo + "#" + body` (see `cloud_utils.go`). This `timestamp#memo#body` scheme is BitMart-specific — preserve it exactly.

### API method files (one per API domain)
Each file is a set of methods on `*CloudClient` that map 1:1 to BitMart endpoints. Endpoint path constants and WS channel constants all live in `cloud_consts.go`.
- `api_system.go` — system time / service status
- `api_account.go` — funding account, deposit/withdraw, fees, transfers
- `api_spot.go` — spot public market data + spot/margin trading; defines `Order`, `MarginOrder`, `BatchOrder` request structs
- `api_margin_loan.go` — isolated margin borrow/repay
- `api_contract.go` — all futures market data + trading; defines `ContractOrder`, `ContractPlanOrder`, `ContractTrailOrder`
- `api_broker.go` — broker rebate

Convention for methods: required params are positional args; optional params come in via `options ...map[string]interface{}` and are merged with `CreateParams(options...)`. Use the `NewParams`/`AddParams`/`CreateParams`/`AddToParams` helpers in `cloud_utils.go` rather than building maps by hand. Doc comments note the version (v1/v2/v3/v4) and auth level.

### WebSocket: `CloudWsClient`
- `cloud_ws_client.go` holds shared connection logic (connect, goroutines for `work`/`receive`/`finalize`, auto-reconnect, keepalive ticker, flate decompression of binary frames). `Send()` records every subscribe message in `reconnectUseMessage` so subscriptions are **replayed automatically on reconnect**.
- `cloud_spot_ws_client.go` (`NewSpotWSClient`) and `cloud_futures_ws_client.go` (`NewFuturesWSClient`) embed `CloudWsClient` and supply protocol differences:
  - **Spot** uses `{"op":"subscribe",...}` / login op `"login"` with args `[apiKey, timestamp, sign]`; ping is the literal string `"ping"`.
  - **Futures** uses `{"action":"subscribe",...}` / login action `"access"` with args `[apiKey, timestamp, sign, "web"]`; ping is `{"action":"ping"}`.
  - WS login sign uses the same HMAC scheme but with the fixed body `"bitmart.WebSocket"`.
- Pick the right URL constant: `SPOT_WS_URL` (public) / `SPOT_WS_USER` (private), `FUTURES_WS_URL` / `FUTURES_WS_USER`. Private streams require `.Login()` before subscribing.
- All incoming messages are delivered to the user-supplied `Callback func(message string)` as a raw string.

### Logging
`cloud_log.go` defines `CustomLogger` with levels `DEBUG`/`INFO`/`ERROR`. Pass `NewCustomLogger(DEBUG, os.Stdout)` (or any `io.Writer`, e.g. a file) via `Config.CustomLogger` to see full request/response dumps. Defaults to `INFO` to stdout.

### Examples
`examples/` contains a runnable `main()` per endpoint, organized as `examples/{spot,futures}/{domain}/{Endpoint}/`. These are the best reference for how each SDK method is called.

## Adding or updating an endpoint

When BitMart adds/changes an API, the change touches several files in lockstep:
1. Add the path constant to `cloud_consts.go` (grouped by domain, with the doc-section URL comment).
2. Add the method to the matching `api_*.go` file, following the positional-required / `options...`-optional convention and noting the auth level in the doc comment.
3. Add a request struct if the body is structured (see `Order` / `ContractOrder` patterns).
4. Add an `examples/.../<Endpoint>/<Endpoint>.go` and a `Test...` in the matching `*_test.go`.
5. Bump `VERSION` in `cloud_consts.go` and record the change in `CHANGELOG.md`.
