package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToInt64(t *testing.T) {
	tcs := []struct {
		name    string
		s       string
		exp     int64
		hasEerr bool
	}{
		{"success", "23", 23, false},
		{"non-string", "3N", 0, true},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ToInt64(tc.s)

			assert.Equal(t, tc.exp, actual)

			if tc.hasEerr && err == nil {
				t.Fail()
			}
		})
	}
}
