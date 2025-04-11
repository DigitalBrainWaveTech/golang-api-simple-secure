package apikey

type KeyStore interface {
	FindByKeyID(id string) (*APIKey, error)
}
