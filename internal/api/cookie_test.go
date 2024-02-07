package api

import (
	"net/http"
	"testing"
	"time"
)

func TestIsCookieExpired(t *testing.T) {
	tests := map[string]struct {
		input http.Cookie
		want  bool
	}{
		"valid cookie": {
			input: http.Cookie{
				Expires: time.Now().Add(1 * time.Hour),
			},
			want: false,
		},
		"expired cookie": {
			input: http.Cookie{
				Expires: time.Now(),
			},
			want: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isCookieExpired(&test.input); got != test.want {
				t.Fatalf(
					"|TEST: %20s| got: %5v, want: %5v", name, got, test.want,
				)
			}
		})
	}
}
