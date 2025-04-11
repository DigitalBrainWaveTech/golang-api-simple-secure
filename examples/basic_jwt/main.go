package main

import (
	"fmt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/handlers"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/jwt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/passwords"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/users"
	"log"
	"net/http"
)

func main() {
	secret := "secret"
	user := auth.User{
		ID:           "1",
		Email:        "user@example.com",
		PasswordHash: passwords.MustHashPassword("password123"),
	}

	provider := users.NewStaticUserProvider(map[string]auth.User{
		"user@example.com": user,
	})
	authProvider := jwt.New(secret, provider)

	mux := http.NewServeMux()

	mux.Handle("/login", handlers.LoginHandler(authProvider))

	mux.Handle("/secure", middleware.AuthMiddleware(authProvider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUserFromContext(r.Context())
		fmt.Fprintf(w, "Welcome, %s", user.Email)
	})))

	log.Println("Server started - listening on :8081")
	http.ListenAndServe(":8081", mux)
}
