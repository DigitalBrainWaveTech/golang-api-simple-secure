package permissions

type PermissionProvider interface {
	GetPermissionsForRoles(roles []string) ([]string, error)
}
