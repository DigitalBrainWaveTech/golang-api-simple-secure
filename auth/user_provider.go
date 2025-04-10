package auth

type UserProvider interface {
	GetUserByEmail(email string) (*User, error)
	ValidateCredentials(email, password string) (*User, error)
}
