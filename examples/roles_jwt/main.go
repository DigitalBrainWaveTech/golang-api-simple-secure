package main

import (
	"fmt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/handlers"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/jwt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/passwords"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/providers"
	"log"
	"net/http"
)

func main() {
	secret := "secret"
	users := map[string]auth.User{
		"admin@example.com": {
			ID:           "1",
			Email:        "admin@example.com",
			PasswordHash: passwords.MustHashPassword("password123"),
			Roles:        []string{"admin"},
		},
	}

	provider := providers.NewStaticUserProvider(users)
	authProvider := jwt.New(secret, provider)

	mux := http.NewServeMux()

	mux.Handle("/login", handlers.LoginHandler(authProvider))

	mux.Handle("/admin", middleware.AuthMiddleware(authProvider)(
		middleware.RequireRole("admin")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Welcome admin!")
		})),
	))

	log.Println("Server started - listening on :8082")
	http.ListenAndServe(":8082", mux)
}
