package main

import (
	"fmt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/jwt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
	"net/http"
)

func main() {
	secret := "super_secret_key"
	auth := jwt.New(secret)

	mux := http.NewServeMux()

	mux.Handle("/secure", middleware.AuthMiddleware(auth)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUserFromContext(r.Context())
		fmt.Fprintf(w, "Hello, %s!", user.Email)
	})))

	http.ListenAndServe(":8080", mux)
}
