package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plaintext password with a secure cost factor
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// MustHashPassword hashes a plaintext password and panics on error
func MustHashPassword(password string) string {
	hash, err := HashPassword(password)
	if err != nil {
		panic(err)
	}

	return hash
}

// CheckPasswordHash compares plaintext password to a stored, hashed one
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
