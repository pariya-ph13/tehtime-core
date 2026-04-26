# TehTime Core APIs

This document describes the available HTTP APIs and the htmx web endpoints provided by this service.

## Base URL

- Local development: `http://localhost:3000`

## APIs

- GET `/time`
  - Returns the current time.
  - Query params:
    - `calendar` (optional): `gregorian` or `solar`.
      - Default: `solar` (Jalali/Solar Hijri)
  - Examples:
    - `GET /time` → `{ "time": "1405-02-06T15:22:05Z", "calendar": "solar" }`
    - `GET /time?calendar=gregorian` → `{ "time": "2026-04-26T15:22:05Z", "calendar": "gregorian" }`

## Web (htmx) Endpoints

- GET `/`
  - Renders the full calendar page (default calendar = `solar`).
  - The page includes controls to switch calendar and navigate between months using htmx.

- GET `/calendar`
  - Returns the calendar month grid partial used by htmx swaps.
  - Query params:
    - `calendar`: `solar` (default) or `gregorian`
    - `year`: integer year (optional; defaults to current for chosen calendar)
    - `month`: integer month 1-12 (optional; defaults to current for chosen calendar)
  - Example:
    - `GET /calendar?calendar=gregorian&year=2026&month=4`

## Configuration

- Config file: `config.yaml` (read via viper).
- Relevant keys in `internal/config/config.go`:
  - `debug` (bool)
  - `http_server.address` (string, optional; server listens on `:3000` by default via fx lifecycle)

## Tech Stack

- Go (Fiber v2, Fx DI)
- Templates: `github.com/gofiber/template/html/v2`
- htmx frontend: loaded from CDN on the page
- Jalali conversion: `github.com/jalaali/go-jalaali`

## Run Locally

- Build: `go build ./...`
- Run: `go run main.go` (or your entrypoint)
- Open: `http://localhost:3000/`

## Notes

- Default calendar is Solar (Jalali). Pass `calendar=gregorian` where applicable to switch.
- The `/calendar` endpoint is designed for partial HTML swaps with htmx; use `/` for the initial page load.
