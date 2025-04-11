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
