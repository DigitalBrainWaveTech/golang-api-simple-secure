# Basic JWT Authentication Example

This example demonstrates a minimal setup using static user authentication with JWT.  
It includes:

- A hardcoded user (`user@example.com` / `password123`)
- A `POST /login` route that returns a JWT
- A `GET /secure` route protected by the JWT
- Optional `/me` and `/logout` handlers

## Running

```bash
go run main.go
```

Then use curl or Postman to test:

### Login

```bash
curl -X POST http://localhost:8081/login -d '{"email":"user@example.com","password":"password123"}' -H "Content-Type: application/json"
```

### Access secure route

```bash
curl http://localhost:8081/secure -H "Authorization: Bearer <your_token>"
```