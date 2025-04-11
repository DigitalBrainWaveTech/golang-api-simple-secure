package main

import (
	"fmt"
	"net/http"

	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/apikey"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	keyID := "external-001"
	rawKey := "super-secret-key"

	hashed, _ := bcrypt.GenerateFromPassword([]byte(rawKey), bcrypt.DefaultCost)

	store := apikey.NewStaticKeyStore(map[string]*apikey.APIKey{
		keyID: {
			KeyID:       keyID,
			KeyHash:     string(hashed),
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
