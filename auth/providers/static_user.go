package providers

import (
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
)

type StaticUserProvider struct {
	users              map[string]auth.User
	permissionProvider auth.PermissionProvider
}

var _ auth.UserProvider = (*StaticUserProvider)(nil)

func NewStaticUserProvider(users map[string]auth.User) *StaticUserProvider {
	return &StaticUserProvider{
		users:              users,
		permissionProvider: nil,
	}
}

func NewStaticUserProviderWithPermissions(users map[string]auth.User, permissions auth.PermissionProvider) *StaticUserProvider {
	return &StaticUserProvider{
		users:              users,
		permissionProvider: permissions,
	}
}

func (p *StaticUserProvider) GetUserByEmail(email string) (*auth.User, error) {
	user, exists := p.users[email]
	if !exists {
		return nil, auth.ErrUserNotFound
	}

	// Enrich permissions if not already present
	if len(user.Permissions) == 0 && p.permissionProvider != nil {
		perms, _ := p.permissionProvider.GetPermissionsForRoles(user.Roles)
		user.Permissions = perms
	}

	return &user, nil
}

func (p *StaticUserProvider) ValidateCredentials(email, password string) (*auth.User, error) {
	user, exists := p.users[email]
	if !exists {
		return nil, auth.ErrUserNotFound
	}

	if !auth.CheckPasswordHash(password, user.PasswordHash) {
		return nil, auth.ErrInvalidCredentials
	}

	// Enrich permissions from role mapping if provider is present
	if len(user.Permissions) == 0 && p.permissionProvider != nil {
		perms, _ := p.permissionProvider.GetPermissionsForRoles(user.Roles)
		user.Permissions = perms
	}

	return &user, nil
}
