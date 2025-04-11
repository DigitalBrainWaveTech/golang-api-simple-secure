package apikey

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// GenerateInternalAPIKey returns a deterministic HMAC hash of a raw API key using a shared secret.
// This is used for internal API key authentication.
func GenerateInternalAPIKey(rawKey, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(rawKey))
	return hex.EncodeToString(mac.Sum(nil))
}

func GenerateExternalAPIKey(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func MustGenerateExternalAPIKey(password string) string {
	key, err := GenerateExternalAPIKey(password)
	if err != nil {
		panic(err)
	}

	return key
}
