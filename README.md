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

### Basic JWT Authentication Example

This example demonstrates a minimal setup using static user authentication with JWT.

For more details and implementation examples, refer to
the [documentation](https://github.com/DigitalBrainWaveTech/golang-api-simple-secure/examples/basic_jwt/README.md).

### JWT with Role-Based Access Control

This example builds on basic JWT to include user roles.

For more details and implementation examples, refer to
the [documentation](https://github.com/DigitalBrainWaveTech/golang-api-simple-secure/examples/roles_jwt/README.md).

### Roles and Permissions Example

This example demonstrates using a `PermissionProvider` to enrich users with permissions based on their roles.

For more details and implementation examples, refer to
the [documentation](https://github.com/DigitalBrainWaveTech/golang-api-simple-secure/examples/roles_and_permissions/README.md).

### Custom Providers Example

This example shows how to plug in your own `UserProvider` that mimics a database.

For more details and implementation examples, refer to
the [documentation](https://github.com/DigitalBrainWaveTech/golang-api-simple-secure/examples/custom_providers/README.md).
