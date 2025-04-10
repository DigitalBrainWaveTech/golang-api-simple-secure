package apikey

import (
	"errors"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/providers"
	"net/http"
)

type Authenticator struct {
	expectedKey  string
	userProvider auth.UserProvider
}

var _ auth.Authenticator = (*Authenticator)(nil)

func NewAPIKeyAuthenticator(expectedKey string) *Authenticator {
	return &Authenticator{
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

func NewAPIKeyAuthenticatorWithUserProvider(expectedKey string, userProvider auth.UserProvider) *Authenticator {
	return &Authenticator{
		expectedKey:  expectedKey,
		userProvider: userProvider,
	}
}

func (a *Authenticator) AuthenticateRequest(r *http.Request) (*auth.User, error) {
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

func (a *Authenticator) Login(email, password string) (*auth.Token, error) {
	return nil, errors.New("Login not supported for API key auth")
}

func (a *Authenticator) ValidateToken(tokenString string) (*auth.User, error) {
	return nil, errors.New("Token validation not used in API key auth")
}
