package main

import (
	"fmt"
	"net/http"

	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/can"
)

func main() {
	user := &auth.User{
		ID:          "user-1",
		Email:       "example@digitalbrainwave.com",
		Roles:       []string{"user"},
		Permissions: []string{"read_dashboard", "send_notification"},
	}

	mux := http.NewServeMux()

	mux.Handle("/dashboard", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !can.Do(user, "read_dashboard") {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		fmt.Fprintln(w, "Welcome to the dashboard!")
	}))

	mux.Handle("/notify", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		can.DoFunc(user, "send_notification", func() {
			fmt.Fprintln(w, "Notification sent.")
		})
	}))

	mux.Handle("/audit", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		can.DoAnyFunc(user, []string{"audit_logs", "admin"}, func() {
			fmt.Fprintln(w, "You can view audit logs.")
		})
	}))

	fmt.Println("Listening on :8088")
	http.ListenAndServe(":8088", mux)
}
