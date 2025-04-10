package handlers

import (
	"encoding/json"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"net/http"
)

func LoginHandler(authenticator auth.Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		token, err := authenticator.Login(req.Email, req.Password)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		json.NewEncoder(w).Encode(token)
	}
}
