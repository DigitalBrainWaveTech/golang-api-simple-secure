# Internal API Key Authentication Example

This example demonstrates an **internal API key** authentication mechanism using HMAC. It’s designed for **server-to-server** usage within a trusted network.

Internal API keys:
- Use a `keyID:rawKey` format in the `X-API-Key` header
- Are validated by computing an HMAC of the raw key using a shared secret
- Are fast and deterministic
- Are ideal for **microservices**, **internal jobs**, and **infrastructure**

## API Key Format

    X-API-Key: internal-001:internal-raw-key

- `internal-001` is the key ID
- `internal-raw-key` is the secret raw key
- A shared HMAC secret is used to validate it

## Running the Example

```bash
go run main.go
```

Server will run at:

    http://localhost:8086

## Example Request

```bash
curl -H "X-API-Key: internal-001:internal-raw-key" http://localhost:8086/internal-ping
```

Expected response:

    Internal Authenticated as: internal-service@digitalbrainwave

## Key Setup in Code

```go
hashed := apikey.GenerateHMACAPIKey("internal-raw-key", "my-internal-shared-secret")

store := apikey.NewStaticKeyStore(map[string]*apikey.APIKey{
    "internal-001": {
        KeyID:       "internal-001",
        KeyHash:     hashed,
        Owner:       "internal-service@digitalbrainwave",
        Roles:       []string{"internal"},
        Permissions: []string{"read_secure"},
    },
})
```

The raw key is hashed with a shared secret for verification.