package apikey

import (
	"errors"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/providers"
	"net/http"
)

type SimpleAuthenticator struct {
	expectedKey  string
	userProvider auth.UserProvider
}

var _ auth.Authenticator = (*SimpleAuthenticator)(nil)

func NewSimpleAPIKeyAuthenticator(expectedKey string) *SimpleAuthenticator {
	return &SimpleAuthenticator{
		expectedKey: expectedKey,
		userProvider: providers.NewStaticUserProviderWithPermissions(map[string]auth.User{
			"admin": {
				ID:          "admin",
				Email:       "admin",
				Roles:       []string{"admin"},
				Permissions: []string{"read", "write"},
			},
		}, providers.NewMapPermissionProvider(map[string][]string{
			"admin": {"read", "write"},
		})),
	}
}

func NewAPIKeyAuthenticatorWithUserProvider(expectedKey string, userProvider auth.UserProvider) *SimpleAuthenticator {
	return &SimpleAuthenticator{
		expectedKey:  expectedKey,
		userProvider: userProvider,
	}
}

func (a *SimpleAuthenticator) AuthenticateRequest(r *http.Request) (*auth.User, error) {
	incomingKey := r.Header.Get("X-API-Key")
	if incomingKey != a.expectedKey {
		return nil, errors.New("invalid API key")
	}

	user, err := a.userProvider.GetUserByEmail("admin")
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *SimpleAuthenticator) Login(email, password string) (*auth.Token, error) {
	return nil, errors.New("Login not supported for API key auth")
}

func (a *SimpleAuthenticator) ValidateToken(tokenString string) (*auth.User, error) {
	return nil, errors.New("Token validation not used in API key auth")
}
