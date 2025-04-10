# Roles and Permissions Example

This example demonstrates using a `PermissionProvider` to enrich users with permissions based on their roles.

Features:
- User has a `manager` role
- Role maps to `view_reports` permission
- `RequirePermission("view_reports")` middleware
- Routes:
  - `POST /login`
  - `GET /reports`
  - `GET /me`
  - `POST /logout`

## Running

```bash
go run main.go
```

### Test:

```bash
curl -X POST http://localhost:8083/login -d '{"email":"manager@example.com","password":"password123"}' -H "Content-Type: application/json"
curl http://localhost:8083/reports -H "Authorization: Bearer <your_token>"
```