package auth

import "net/http"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Roles        []string
	Permissions  []string
}

type Token struct {
	Value     string
	ExpiresAt int64
}

type Authenticator interface {
	AuthenticateRequest(r *http.Request) (*User, error)
	Login(email, password string) (*Token, error)
	ValidateToken(tokenString string) (*User, error)
}
