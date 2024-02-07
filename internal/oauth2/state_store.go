package oauth2

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type StateStore struct {
	store map[string]time.Time
}

func NewStateStore() StateStore {
	return StateStore{
		store: make(map[string]time.Time),
	}
}

func (s StateStore) GenerateState() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	state := base64.URLEncoding.EncodeToString(b)
	s.store[state] = time.Now().Add(5 * time.Minute)
	return state, nil
}

func (s StateStore) ValidateState(state string) bool {
	expiry, ok := s.store[state]
	if !ok {
		return false
	}
	if time.Now().After(expiry) {
		return false
	}
	return true
}

func (s StateStore) DeleteState(state string) {
	delete(s.store, state)
}
