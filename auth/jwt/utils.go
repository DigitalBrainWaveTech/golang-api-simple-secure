package jwt

import (
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtExpiry = time.Hour * 24

func GenerateJWT(email string, role string, secret string) (*auth.Token, error) {
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(jwtExpiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &auth.Token{
		Value:     tokenStr,
		ExpiresAt: time.Now().Add(jwtExpiry).Unix(),
	}, nil
}

func ParseJWT(tokenStr, secret string) (*auth.User, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return &auth.User{
		Email: claims["email"].(string),
		Role:  claims["role"].(string),
	}, nil
}
