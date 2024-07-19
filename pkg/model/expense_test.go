package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr/http"
)

func Test_ExpenseValidate(t *testing.T) {
	tcs := []struct {
		name     string
		exp      Expense
		expected error
	}{
		{
			name: "success",
			exp: Expense{
				UserID:     "12l23",
				CategoryID: "45d3j",
				Title:      "expense",
				Amount:     24.5,
			},
			expected: nil,
		},
		{
			name:     "fields missing",
			exp:      Expense{},
			expected: http.ErrorMissingParam{Params: []string{"title", "amount", "category_id", "user_id"}},
		},
	}

	for _, tt := range tcs {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.exp.Validate()

			assert.Equal(t, tc.expected, actual)
		})
	}
}
