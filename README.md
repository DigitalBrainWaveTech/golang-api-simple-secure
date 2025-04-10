# Golang API: Simple Secure

A simple solution to adding basic security features to a server-to-server accessed
API within your own VPC network.

It uses the standard logging library "log", and is designed to extremely simple to 
integrate into a Golang application, with no dependencies other than the standard library.

## Features

- Basic authentication
- IP whitelisting
- IP blacklisting
- Rate limiting
- Logging
- API key validation for additional security
- Detailed request/response auditing
- Support for role-based access control (RBAC)
- Configurable request throttling based on endpoint

## Installation

```bash
go get github.com/DigitalBrainWaveTech/golang-api-simple-secure
```
