package main

import (
	"fmt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/jwt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/providers"
	"net/http"
)

func main() {
	secret := "secret"
	user := auth.User{
		ID:       "1",
		Email:    "user@example.com",
		Password: "password123",
	}

	provider := providers.NewStaticUserProvider(map[string]auth.User{
		"user@example.com": user,
	})
	authProvider := jwt.New(secret, provider)

	mux := http.NewServeMux()

	mux.Handle("/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := authProvider.Login(user.Email, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Fprintf(w, "Token: %s", token.Value)
	}))

	mux.Handle("/secure", middleware.AuthMiddleware(authProvider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUserFromContext(r.Context())
		fmt.Fprintf(w, "Welcome, %s", user.Email)
	})))

	http.ListenAndServe(":8081", mux)
}
