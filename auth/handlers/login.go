package handlers

import (
	"encoding/json"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"log"
	"net/http"
)

func LoginHandler(authenticator auth.Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("Error decoding request body: %v", err)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		token, err := authenticator.Login(req.Email, req.Password)
		if err != nil {
			log.Printf("Error authenticating user: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		json.NewEncoder(w).Encode(token)
	}
}
