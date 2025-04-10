package main

import (
	"fmt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/handlers"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/jwt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
	"net/http"
	"strings"
)

func main() {
	secret := "secret"

	userProvider := &MockDBUserProvider{}
	authProvider := jwt.New(secret, userProvider)

	mux := http.NewServeMux()

	mux.Handle("/login", handlers.LoginHandler(authProvider))

	mux.Handle("/deploy", middleware.AuthMiddleware(authProvider)(
		middleware.RequirePermission("deploy_code")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "You are allowed to deploy code.")
		})),
	))

	http.ListenAndServe(":8084", mux)
}

type MockDBUserProvider struct{}

func (p *MockDBUserProvider) GetUserByEmail(email string) (*auth.User, error) {
	return p.find(email)
}

func (p *MockDBUserProvider) ValidateCredentials(email, password string) (*auth.User, error) {
	u, err := p.find(email)
	if err != nil {
		return nil, err
	}
	if u.Password != password {
		return nil, auth.ErrInvalidCredentials
	}
	return u, nil
}

func (p *MockDBUserProvider) find(email string) (*auth.User, error) {
	email = strings.ToLower(email)
	if email == "developer@example.com" {
		return &auth.User{
			ID:          "42",
			Email:       email,
			Password:    "hunter2",
			Roles:       []string{"dev"},
			Permissions: []string{"deploy_code", "read_logs"},
		}, nil
	}
	return nil, auth.ErrUserNotFound
}
