package main

import "testing"

func TestVerifyEmail(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"regular email": {
			input: "username@gmail.com",
			want:  true,
		},
		"email without username": {
			input: "@gmail.com",
			want:  false,
		},
		"email without domain": {
			input: "username@",
			want:  false,
		},
		"email without mailserver": {
			input: "username@.com",
			want:  false,
		},
		"email without top-level domain": {
			input: "username@gmail",
			want:  false,
		},
		"email without top-level domain 2": {
			input: "username@gmail.",
			want:  false,
		},
		"empty email": {
			input: "",
			want:  false,
		},
		"multiple top-level domains": {
			input: "username@schools.city.edu",
			want:  true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := verifyEmail(test.input); got != test.want {
				t.Fatalf("|TEST: %20s| got: %5t, want: %5t", name, got, test.want)
			}
		})
	}
}
