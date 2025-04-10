package providers

import "github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"

type StaticUserProvider struct {
	users map[string]auth.User
}

var _ auth.UserProvider = (*StaticUserProvider)(nil)

func NewStaticUserProvider(users map[string]auth.User) *StaticUserProvider {
	return &StaticUserProvider{users: users}
}

func (p *StaticUserProvider) GetUserByEmail(email string) (*auth.User, error) {
	user, exists := p.users[email]
	if !exists {
		return nil, auth.ErrUserNotFound
	}
	return &user, nil
}

func (p *StaticUserProvider) ValidateCredentials(email, password string) (*auth.User, error) {
	user, exists := p.users[email]
	if !exists {
		return nil, auth.ErrUserNotFound
	}

	if user.Password != password {
		return nil, auth.ErrInvalidCredentials
	}

	return &user, nil
}
