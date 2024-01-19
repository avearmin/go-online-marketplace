package oauth2

import (
	"crypto/rand"
	"encoding/base64"
)

type stateStore struct {
	store map[string]bool
}

func NewStateStore() *stateStore {
	return &stateStore{
		store: make(map[string]bool),
	}
}

func (s *stateStore) GenerateState() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	s.store[state] = true
	return state
}

func (s *stateStore) ValidateState(state string) bool {
	return s.store[state]
}

func (s *stateStore) DeleteState(state string) {
	delete(s.store, state)
}
