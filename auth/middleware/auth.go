package middleware

import (
	"context"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"log"
	"net/http"
)

type contextKey string

const userKey contextKey = "user"

func AuthMiddleware(authenticator auth.Authenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := authenticator.AuthenticateRequest(r)
			if err != nil {
				log.Printf("Authentication failed: %v", err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			log.Printf("Authenticated user: %q", user.ID)
			ctx := SetUserInContext(r.Context(), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func SetUserInContext(ctx context.Context, user *auth.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUserFromContext(ctx context.Context) *auth.User {
	user, ok := ctx.Value(userKey).(*auth.User)
	if !ok {
		return nil
	}
	return user
}
