package main

import (
	"fmt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/handlers"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/jwt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/middleware"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/providers"
	"net/http"
)

func main() {
	secret := "secret"
	users := map[string]auth.User{
		"manager@example.com": {
			ID:       "1",
			Email:    "manager@example.com",
			Password: "password123",
			Roles:    []string{"manager"},
		},
	}

	rolePermissions := map[string][]string{
		"manager": {"view_reports"},
	}

	permissionProvider := providers.NewMapPermissionProvider(rolePermissions)
	provider := providers.NewStaticUserProviderWithPermissions(users, permissionProvider)
	authProvider := jwt.New(secret, provider)

	mux := http.NewServeMux()

	mux.Handle("/login", handlers.LoginHandler(authProvider))

	mux.Handle("/reports", middleware.AuthMiddleware(authProvider)(
		middleware.RequirePermission("view_reports")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "You are allowed to view reports.")
		})),
	))

	http.ListenAndServe(":8083", mux)
}
