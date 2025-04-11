package main

import (
	"fmt"
	"net/http"

	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/apikey"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
)

func main() {
	sharedSecret := "my-internal-shared-secret"
	keyID := "internal-001"
	rawKey := "internal-raw-key"

	hashed := apikey.GenerateInternalAPIKey(rawKey, sharedSecret)

	store := apikey.NewStaticKeyStore(map[string]*apikey.APIKey{
		keyID: {
			KeyID:       keyID,
			KeyHash:     hashed,
			Owner:       "internal-service@digitalbrainwave",
			Roles:       []string{"internal"},
			Permissions: []string{"read_secure"},
		},
	})

	authProvider := apikey.NewInternalAuthenticator(store, sharedSecret)

	mux := http.NewServeMux()

	mux.Handle("/internal-ping", middleware.AuthMiddleware(authProvider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUserFromContext(r.Context())
		fmt.Fprintf(w, "Internal Authenticated as: %s", user.Email)
	})))

	http.ListenAndServe(":8086", mux)
}
