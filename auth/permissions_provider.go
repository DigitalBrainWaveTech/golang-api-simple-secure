package auth

type PermissionProvider interface {
	GetPermissionsForRoles(roles []string) ([]string, error)
}
