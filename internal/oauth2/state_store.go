package oauth2

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"sync"
	"time"
)

type StateStore struct {
	store map[string]time.Time
	mu    *sync.Mutex
}

func NewStateStore() StateStore {
	stateStore := StateStore{
		store: make(map[string]time.Time),
		mu:    &sync.Mutex{},
	}
	stateStore.startReapLoop()
	return stateStore
}

func (s StateStore) GenerateState() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	state := base64.URLEncoding.EncodeToString(b)
	s.mu.Lock()
	s.store[state] = time.Now().Add(5 * time.Minute)
	s.mu.Unlock()
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
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, state)
}

// the Google callback handler should delete states from the server once they're used,
// but we need to reap any expired states that might have accumulated for whatever reason
func (s StateStore) startReapLoop() {
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			log.Println("awakening state reaper")
			s.reap()
			log.Println("state reaper going back to sleep")
		}
	}()
}

func (s StateStore) reap() {
	for k, v := range s.store {
		if time.Now().After(v) {
			log.Printf("Reaping: %s", k)
			s.DeleteState(k)
		}
	}
}
