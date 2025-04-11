package can

import "github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"

// Do checks if the user has the given permission.
func Do(user *auth.User, permission string) bool {
	return auth.HasPermission(user, permission)
}

// DoAny checks if the user has *any* of the given permissions.
func DoAny(user *auth.User, permissions ...string) bool {
	return auth.HasAnyPermission(user, permissions...)
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
