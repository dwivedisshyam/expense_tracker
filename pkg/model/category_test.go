package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr/http"
)

func Test_CategoryValidate(t *testing.T) {
	tcs := []struct {
		Name     string
		category Category
		expected error
	}{
		{
			Name: "success",
			category: Category{
				UserID: "533w",
				Name:   "Rent",
			},
			expected: nil,
		},
		{
			Name:     "fields missing",
			category: Category{},
			expected: http.ErrorMissingParam{Params: []string{"name", "user_id"}},
		},
	}

	for _, tt := range tcs {
		tc := tt
		t.Run(tc.Name, func(t *testing.T) {
			actual := tc.category.Validate()

			assert.Equal(t, tc.expected, actual)
		})
	}
}
