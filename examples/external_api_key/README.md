# External API Key Authentication Example

This example demonstrates how to authenticate third-party or partner clients using an **external API key** approach.

External API keys:
- Use a `keyID:rawKey` format in the `X-API-Key` header
- Are validated against a **bcrypt hash** stored in a key store
- Support per-key roles and permissions
- Are ideal for **external clients** or **untrusted networks**

## API Key Format

    X-API-Key: external-001:super-secret-key

- `external-001` is the key ID used to look up the stored hash
- `super-secret-key` is the raw secret (only known to the client)

## Running the Example

```bash
go run main.go
```

Server will run at:

    http://localhost:8087

## Example Request

```bash
curl -H "X-API-Key: external-001:super-secret-key" http://localhost:8087/external-ping
```

Expected response:

    External Authenticated as: partner@example.com

## Key Setup in Code

```go
hashed, _ := bcrypt.GenerateFromPassword([]byte("super-secret-key"), bcrypt.DefaultCost)

store := apikey.NewStaticKeyStore(map[string]*apikey.APIKey{
    "external-001": {
        KeyID:       "external-001",
        KeyHash:     string(hashed),
        Owner:       "partner@example.com",
        Roles:       []string{"partner"},
        Permissions: []string{"read_public"},
    },
})
```

Only the **hash** is stored — the plaintext key is **never persisted** after generation.