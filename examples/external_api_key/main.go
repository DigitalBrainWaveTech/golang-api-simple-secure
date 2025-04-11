package main

import (
	"fmt"
	"net/http"

	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/apikey"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
)

func main() {
	keyID := "external-001"
	rawKey := "super-secret-key"

	hashed := apikey.MustGenerateExternalAPIKey(rawKey)

	store := apikey.NewStaticKeyStore(map[string]*apikey.APIKey{
		keyID: {
			KeyID:       keyID,
			KeyHash:     hashed,
			Owner:       "partner@example.com",
			Roles:       []string{"partner"},
			Permissions: []string{"read_public"},
		},
	})

	authProvider := apikey.NewExternalAuthenticator(store)

	mux := http.NewServeMux()

	mux.Handle("/external-ping", middleware.AuthMiddleware(authProvider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUserFromContext(r.Context())
		fmt.Fprintf(w, "External Authenticated as: %s", user.Email)
	})))

	http.ListenAndServe(":8087", mux)
}
