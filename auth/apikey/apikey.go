package apikey

type APIKey struct {
	KeyID       string
	KeyHash     string // bcrypt for external, raw HMAC for internal
	Owner       string
	Roles       []string
	Permissions []string
}
