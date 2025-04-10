package providers

type MapPermissionProvider struct {
	rolePermissions map[string][]string
}

func NewMapPermissionProvider(rp map[string][]string) *MapPermissionProvider {
	return &MapPermissionProvider{rolePermissions: rp}
}

func (m *MapPermissionProvider) GetPermissionsForRoles(roles []string) ([]string, error) {
	permSet := make(map[string]struct{})
	for _, role := range roles {
		if perms, ok := m.rolePermissions[role]; ok {
			for _, p := range perms {
				permSet[p] = struct{}{}
			}
		}
	}

	var perms []string
	for p := range permSet {
		perms = append(perms, p)
	}

	return perms, nil
}
