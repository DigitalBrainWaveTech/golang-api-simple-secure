# JWT with Role-Based Access Control

This example builds on basic JWT to include user roles.  
It includes:

- A user with the `admin` role
- A `RequireRole("admin")` middleware
- Routes:
  - `POST /login`
  - `GET /admin` (requires `admin` role)
  - `GET /me`
  - `POST /logout`

## Running

```bash
go run main.go
```

### Test login and access

```bash
curl -X POST http://localhost:8082/login -d '{"email":"admin@example.com","password":"password123"}' -H "Content-Type: application/json"
curl http://localhost:8082/admin -H "Authorization: Bearer <your_token>"
```