package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateJWT(t *testing.T) {
	type Input struct {
		id        uuid.UUID
		secret    string
		expiresIn time.Duration
		issuer    string
		nowFunc   timeNowFunc
	}

	type Want struct {
		value string
		err   error
	}

	fixedUUID, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	expectedJWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0Iiwic3ViIjoiMTIzZTQ1NjctZTg5Yi0xMmQzLWE0NTYtNDI2NjE0MTc0MDAwIiwiZXhwIjoxNzA0MDY3MjAwLCJpYXQiOjE3MDQwNjcyMDB9.UPHcGrR__xxK5sHeY_kXhsLFwGWeh3oL-54CPbKZnRg"

	tests := map[string]struct {
		input Input
		want  Want
	}{
		"normal JWT": {
			input: Input{
				id:        fixedUUID,
				secret:    "secret",
				expiresIn: time.Duration(1),
				issuer:    "test",
				nowFunc:   mockTimeNow,
			},
			want: Want{
				value: expectedJWT,
				err:   nil,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if value, err := createJWT(
				test.input.id,
				test.input.secret,
				test.input.expiresIn,
				test.input.issuer,
				test.input.nowFunc,
			); value != test.want.value || err != test.want.err {
				t.Fatalf(
					"|TEST: %20s| got value: %5v, got err: %5v, want value: %5v, want err:%5v",
					name, value, err, test.want.value, test.want.err,
				)
			}
		})
	}

}

func TestCalcExpiry(t *testing.T) {
	testCases := map[string]struct {
		input    int
		want time.Duration
	}{
		"1 hour": {1, time.Duration(3600000000000)},
		"0 hours": {0, time.Duration(0)},
		"24 hours": {24, time.Duration(86400000000000)},
		"negative hours": {-5, time.Duration(0)},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			if got, _ := calcExpiry(test.input); got != test.want {
				t.Fatalf("|TEST: %20s| got: %5v, want: %5v", name, got, test.want)
			}
		})
	}
}

func mockTimeNow() time.Time {
	return time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
}
