package main

import (
	"fmt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/apikey"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
	"log"
	"net/http"
)

func main() {
	apiKey := "super-secret-api-key"
	authProvider := apikey.NewSimpleAPIKeyAuthenticator(apiKey)

	mux := http.NewServeMux()

	mux.Handle("/ping", middleware.AuthMiddleware(authProvider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUserFromContext(r.Context())
		fmt.Fprintf(w, "Hello %s! Secure ping received.", user.Email)
	})))

	log.Println("Server started - listening on :8085")
	http.ListenAndServe(":8085", mux)
}
