package apikey

import (
	"errors"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"net/http"
)

type Authenticator struct {
	expectedKey string
	userProvider auth.UserProvider
}

var _ auth.Authenticator = (*Authenticator)(nil)

func NewAPIKeyAuthenticator(expectedKey string) *Authenticator {
	return &Authenticator{
		expectedKey: expectedKey,
		userProvider:
	}
}

func NewAPIKeyAuthenticatorWithUserProvider(expectedKey string, userProvider auth.UserProvider) *Authenticator {
	return &Authenticator{
		expectedKey: expectedKey,
		userProvider: userProvider,
	}
}

func (a *Authenticator) AuthenticateRequest(r *http.Request) (*auth.User, error) {
	incomingKey := r.Header.Get("X-API-Key")
	if incomingKey != a.expectedKey {
		return nil, errors.New("invalid API key")
	}

	return &auth.User{
		ID:    "laravel-service",
		Email: "laravel@digitalbrainwave.internal",
		Role:  "system",
	}, nil
}

func (a *Authenticator) Login(email, password string) (*auth.Token, error) {
	return nil, errors.New("Login not supported for API key auth")
}

func (a *Authenticator) ValidateToken(tokenString string) (*auth.User, error) {
	return nil, errors.New("Token validation not used in API key auth")
}
