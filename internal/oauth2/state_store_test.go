package oauth2

import "testing"

func TestValidateState(t *testing.T) {
	tests := map[string]struct {
		stateStore StateStore
		input      string
		want       bool
	}{
		"valid": {
			stateStore: StateStore{
				store: map[string]bool{"state": true},
			},
			input: "state",
			want:  true,
		},
		"not valid": {
			stateStore: StateStore{
				store: map[string]bool{},
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
				store: map[string]bool{"state": true},
			},
			input: "state",
			want:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.stateStore.DeleteState(test.input)
			if got := test.stateStore.store["state"]; got != test.want {
				t.Fatalf("|TEST: %20s| got: %5t, want: %5t", name, got, test.want)
			}
		})
	}
}
