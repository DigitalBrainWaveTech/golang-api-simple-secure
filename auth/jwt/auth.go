package jwt

import (
	"errors"
	"fmt"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"net/http"
	"strings"
)

type Authenticator struct {
	secret       string
	userProvider auth.UserProvider
}

var _ auth.Authenticator = (*Authenticator)(nil)

func New(secret string, provider auth.UserProvider) *Authenticator {
	return &Authenticator{
		secret:       secret,
		userProvider: provider,
	}
}

func (j *Authenticator) AuthenticateRequest(r *http.Request) (*auth.User, error) {
	tokenStr := extractBearerToken(r.Header.Get("Authorization"))
	if tokenStr == "" {
		return nil, errors.New("missing token")
	}
	return j.ValidateToken(tokenStr)
}

func (j *Authenticator) Login(email, password string) (*auth.Token, error) {
	user, err := j.userProvider.ValidateCredentials(email, password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials: %w", err)
	}
	return GenerateJWT(user.Email, user.Roles, j.secret)
}

func (j *Authenticator) ValidateToken(tokenStr string) (*auth.User, error) {
	return ParseJWT(tokenStr, j.secret)
}

func extractBearerToken(authHeader string) string {
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return ""
}
