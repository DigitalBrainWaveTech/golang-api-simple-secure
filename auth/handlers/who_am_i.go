package handlers

import (
	"encoding/json"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
	"net/http"
)

func WhoAmIHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUserFromContext(r.Context())
		if user == nil {
			http.Error(w, "Not authenticated", http.StatusUnauthorized)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}
