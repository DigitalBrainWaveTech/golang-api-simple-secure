package main

import (
	"fmt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/apikey"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
	"net/http"
)

func main() {
	apiKey := "super-secret-api-key"
	authProvider := apikey.NewAPIKeyAuthenticator(apiKey)

	mux := http.NewServeMux()

	mux.Handle("/ping", middleware.AuthMiddleware(authProvider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUserFromContext(r.Context())
		fmt.Fprintf(w, "Hello %s! Secure ping received.", user.Email)
	})))

	http.ListenAndServe(":8085", mux)
}
