# Simple API Key Auth Example (Server-to-Server)

This example demonstrates a minimal API key-based authentication flow, ideal for internal services communicating within a closed network.

- No sessions, JWTs, or passwords
- Uses `X-API-Key` header for validation
- Only one route: `GET /ping`

## Running

```bash
go run main.go
```

## Request

```bash
curl -H "X-API-Key: super-secret-api-key" http://localhost:8085/ping
```

Expected response:

```
Hello internal-service@digitalbrainwave.internal! Secure ping received.
```