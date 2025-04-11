# Golang API: Simple Secure

A simple solution to adding basic security features to a server-to-server accessed
API within your own VPC network.

It uses the standard logging library "log", and is designed to extremely simple to 
integrate into a Golang application, with minimal dependencies.

## Features

- Basic authentication
- Logging
- API key validation for additional security
- Detailed request/response auditing
- Support for role-based access control (RBAC)

## Installation

```bash
go get github.com/DigitalBrainWaveTech/golang-api-simple-secure
```
## Examples

### API-Key Based

#### Simple API Key Auth Example (Server-to-Server)

This example demonstrates a minimal API key-based authentication flow, ideal for internal services communicating within a closed network.

For more details and implementation examples, refer to
the [documentation](examples/api_key_simple).

#### Internal API Key Example

A fast, HMAC-based API key authenticator designed for internal services and server-to-server communication.

For more details and implementation examples, refer to
the [documentation](examples/internal_api_key).

#### External API Key Example

A bcrypt-backed API key authenticator ideal for securely validating external partners and third-party integrations.

For more details and implementation examples, refer to
the [documentation](examples/external_api_key).

### JWT Based Auth

#### Basic JWT Authentication Example

This example demonstrates a minimal setup using static user authentication with JWT.

For more details and implementation examples, refer to
the [documentation](examples/basic_jwt).

#### JWT with Role-Based Access Control

This example builds on basic JWT to include user roles.

For more details and implementation examples, refer to
the [documentation](examples/roles_jwt).

#### Roles and Permissions Example

This example demonstrates using a `PermissionProvider` to enrich users with permissions based on their roles.

For more details and implementation examples, refer to
the [documentation](examples/roles_and_permissions).

### Custom Providers

#### Custom Providers Example

This example shows how to plug in your own `UserProvider` that mimics a database.

For more details and implementation examples, refer to
the [documentation](examples/custom_providers).
