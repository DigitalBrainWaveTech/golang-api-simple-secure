package apikey

import "errors"

// StaticKeyStore is a simple in-memory implementation of the KeyStore interface.
// It looks up API keys by their KeyID (string).
type StaticKeyStore struct {
	keys map[string]*APIKey
}

// NewStaticKeyStore creates a new StaticKeyStore with a map of keys.
// The key of the map is the KeyID.
func NewStaticKeyStore(keys map[string]*APIKey) *StaticKeyStore {
	return &StaticKeyStore{
		keys: keys,
	}
}

// FindByKeyID returns the API key record by its KeyID.
// Returns an error if the key is not found.
func (s *StaticKeyStore) FindByKeyID(id string) (*APIKey, error) {
	key, ok := s.keys[id]
	if !ok {
		return nil, errors.New("API key not found")
	}
	return key, nil
}
