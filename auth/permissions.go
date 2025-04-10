package auth

func HasPermission(user *User, permission string) bool {
	if user == nil {
		return false
	}
	for _, p := range user.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

func HasAnyPermission(user *User, perms ...string) bool {
	if user == nil {
		return false
	}
	for _, up := range user.Permissions {
		for _, rp := range perms {
			if up == rp {
				return true
			}
		}
	}
	return false
}
