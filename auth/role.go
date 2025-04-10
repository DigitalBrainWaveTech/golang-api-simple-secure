package auth

func HasRole(user *User, role string) bool {
	if user == nil {
		return false
	}
	for _, r := range user.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func HasAnyRole(user *User, roles ...string) bool {
	if user == nil {
		return false
	}
	for _, ur := range user.Roles {
		for _, rr := range roles {
			if ur == rr {
				return true
			}
		}
	}
	return false
}
