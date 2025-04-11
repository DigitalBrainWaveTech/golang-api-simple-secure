package apikey

import (
	"errors"
	"net/http"
	"strings"

	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
)

type InternalAuthenticator struct {
	store        KeyStore
	sharedSecret string
}

var _ auth.Authenticator = (*InternalAuthenticator)(nil)

func NewInternalAuthenticator(store KeyStore, sharedSecret string) *InternalAuthenticator {
	return &InternalAuthenticator{
		store:        store,
		sharedSecret: sharedSecret,
	}
}

func (a *InternalAuthenticator) AuthenticateRequest(r *http.Request) (*auth.User, error) {
	header := r.Header.Get("X-API-Key")
	if header == "" {
		return nil, errors.New("missing API key")
	}

	parts := strings.SplitN(header, ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid API key format")
	}

	keyID := parts[0]
	rawKey := parts[1]

	record, err := a.store.FindByKeyID(keyID)
	if err != nil || record == nil {
		return nil, auth.ErrInvalidCredentials
	}

	expected := GenerateHMACAPIKey(rawKey, a.sharedSecret)
	if expected != record.KeyHash {
		return nil, auth.ErrInvalidCredentials
	}

	return &auth.User{
		ID:          record.KeyID,
		Email:       record.Owner,
		Roles:       record.Roles,
		Permissions: record.Permissions,
	}, nil
}

func (a *InternalAuthenticator) Login(email, password string) (*auth.Token, error) {
	return nil, errors.New("login not supported")
}

func (a *InternalAuthenticator) ValidateToken(token string) (*auth.User, error) {
	return nil, errors.New("token validation not supported")
}
