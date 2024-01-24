package oauth2

import (
	"crypto/rand"
	"encoding/base64"
)

type StateStore struct {
	store map[string]bool
}

func NewStateStore() StateStore {
	return StateStore{
		store: make(map[string]bool),
	}
}

func (s StateStore) GenerateState() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	state := base64.URLEncoding.EncodeToString(b)
	s.store[state] = true
	return state, nil
}

func (s StateStore) ValidateState(state string) bool {
	return s.store[state]
}

func (s StateStore) DeleteState(state string) {
	delete(s.store, state)
}
