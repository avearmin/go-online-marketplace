package oauth2

import (
	"sync"
	"testing"
	"time"
)

func TestValidateState(t *testing.T) {
	tests := map[string]struct {
		stateStore StateStore
		input      string
		want       bool
	}{
		"valid": {
			stateStore: StateStore{
				store: map[string]time.Time{"state": time.Now().Add(1 * time.Hour)},
				mu:    &sync.Mutex{},
			},
			input: "state",
			want:  true,
		},
		"state does not exist": {
			stateStore: StateStore{
				store: map[string]time.Time{},
				mu:    &sync.Mutex{},
			},
			input: "state",
			want:  false,
		},
		"state is expired": {
			stateStore: StateStore{
				store: map[string]time.Time{"state": time.Time{}}, // The zero value is always in the past
				mu:    &sync.Mutex{},
			},
			input: "state",
			want:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := test.stateStore.ValidateState(test.input); got != test.want {
				t.Fatalf("|TEST: %20s| got: %5t, want: %5t", name, got, test.want)
			}
		})
	}
}

func TestDeleteState(t *testing.T) {
	tests := map[string]struct {
		stateStore StateStore
		input      string
		want       bool
	}{
		"successful": {
			stateStore: StateStore{
				store: map[string]time.Time{"state": time.Now()},
				mu:    &sync.Mutex{},
			},
			input: "state",
			want:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.stateStore.DeleteState(test.input)
			if _, got := test.stateStore.store["state"]; got != test.want {
				t.Fatalf("|TEST: %20s| got: %5t, want: %5t", name, got, test.want)
			}
		})
	}
}
