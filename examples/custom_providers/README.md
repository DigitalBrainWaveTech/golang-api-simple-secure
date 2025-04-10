# Custom Providers Example

This example shows how to plug in your own `UserProvider` that mimics a database.

Features:
- A mocked user with email `developer@example.com`
- Hardcoded permissions: `deploy_code`, `read_logs`
- `RequirePermission("deploy_code")` middleware
- Routes:
  - `POST /login`
  - `GET /deploy`
  - `GET /me`
  - `POST /logout`

## Running

```bash
go run main.go
```

### Example usage:

```bash
curl -X POST http://localhost:8084/login -d '{"email":"developer@example.com","password":"hunter2"}' -H "Content-Type: application/json"
curl http://localhost:8084/deploy -H "Authorization: Bearer <your_token>"
```