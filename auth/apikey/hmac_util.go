package apikey

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// GenerateHMACAPIKey returns a deterministic HMAC hash of a raw API key using a shared secret.
// This is used for internal API key authentication.
func GenerateHMACAPIKey(rawKey, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(rawKey))
	return hex.EncodeToString(mac.Sum(nil))
}
