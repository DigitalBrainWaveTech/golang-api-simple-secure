package basic

import (
	"encoding/base64"
	"log"
	"net/http"
)

func Middleware(next http.Handler, expectedToken string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		// base64 decode...
		decodedAuth, err := base64.StdEncoding.DecodeString(auth)
		if err != nil {
			log.Printf("error decoding authorization header: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if string(decodedAuth) != expectedToken {
			log.Printf("error: authorization header does not match expected token")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
