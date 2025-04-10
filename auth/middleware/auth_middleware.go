package middleware

import (
	"context"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"net/http"
)

type contextKey string

const userKey contextKey = "user"

func AuthMiddleware(authenticator auth.Authenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := authenticator.AuthenticateRequest(r)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), userKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserFromContext(ctx context.Context) *auth.User {
	user, ok := ctx.Value(userKey).(*auth.User)
	if !ok {
		return nil
	}
	return user
}
