package can

import (
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth/permissions"
)

// Do checks if the user has the given permission.
func Do(user *auth.User, permission string) bool {
	return permissions.HasPermission(user, permission)
}

// DoAny checks if the user has *any* of the given permissions.
func DoAny(user *auth.User, perms ...string) bool {
	return permissions.HasAnyPermission(user, perms...)
}

// DoAll checks if the user has *all* of the given permissions.
func DoAll(user *auth.User, permissions ...string) bool {
	if user == nil {
		return false
	}

	permSet := make(map[string]struct{}, len(user.Permissions))
	for _, p := range user.Permissions {
		permSet[p] = struct{}{}
	}

	for _, p := range permissions {
		if _, ok := permSet[p]; !ok {
			return false
		}
	}

	return true
}

// DoFunc runs a callback only if the user has the given permission.
func DoFunc(user *auth.User, permission string, fn func()) {
	if Do(user, permission) {
		fn()
	}
}

// DoAnyFunc runs a callback if the user has *any* of the given permissions.
func DoAnyFunc(user *auth.User, permissions []string, fn func()) {
	if DoAny(user, permissions...) {
		fn()
	}
}

// DoAllFunc runs a callback if the user has *all* of the given permissions.
func DoAllFunc(user *auth.User, permissions []string, fn func()) {
	if DoAll(user, permissions...) {
		fn()
	}
}
