package apikey

import (
	"errors"
	"net/http"
	"strings"

	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"golang.org/x/crypto/bcrypt"
)

type ExternalAuthenticator struct {
	store KeyStore
}

var _ auth.Authenticator = (*ExternalAuthenticator)(nil)

func NewExternalAuthenticator(store KeyStore) *ExternalAuthenticator {
	return &ExternalAuthenticator{store: store}
}

func (a *ExternalAuthenticator) AuthenticateRequest(r *http.Request) (*auth.User, error) {
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

	if err := bcrypt.CompareHashAndPassword([]byte(record.KeyHash), []byte(rawKey)); err != nil {
		return nil, auth.ErrInvalidCredentials
	}

	return &auth.User{
		ID:          record.KeyID,
		Email:       record.Owner,
		Roles:       record.Roles,
		Permissions: record.Permissions,
	}, nil
}

func (a *ExternalAuthenticator) Login(email, password string) (*auth.Token, error) {
	return nil, errors.New("login not supported")
}

func (a *ExternalAuthenticator) ValidateToken(token string) (*auth.User, error) {
	return nil, errors.New("token validation not supported")
}
