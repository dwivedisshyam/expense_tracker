package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr/http"
)

func Test_UserValidate(t *testing.T) {
	tcs := []struct {
		name     string
		user     User
		expected error
	}{
		{
			name: "success",
			user: User{
				FirstName: "John",
				LastName:  "Carter",
				Email:     "john.carter@gmail.com",
				Password:  "123456",
			},
			expected: nil,
		},
		{
			name:     "fields missing",
			user:     User{},
			expected: http.ErrorMissingParam{Params: []string{"first_name", "last_name", "email", "password"}},
		},
	}

	for _, tt := range tcs {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.user.Validate()

			assert.Equal(t, tc.expected, actual)
		})
	}
}
