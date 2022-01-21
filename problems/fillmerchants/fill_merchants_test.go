package fillmerchants

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerchants(t *testing.T) {
	tests := []struct {
		name        string
		input       []*merchant
		inputAmount int
		expected    []*merchant
	}{
		{
			name: "valid - even",
			input: []*merchant{
				{key: "a", maxCapacity: 20},
				{key: "c", maxCapacity: 10},
				{key: "d", maxCapacity: 100},
				{key: "b", maxCapacity: 100},
			},
			inputAmount: 100,
			expected: []*merchant{
				{key: "a", maxCapacity: 20, amount: 20},
				{key: "c", maxCapacity: 10, amount: 10},
				{key: "d", maxCapacity: 100, amount: 35},
				{key: "b", maxCapacity: 100, amount: 35},
			},
		},
		{
			name: "valid - uneven",
			input: []*merchant{
				{key: "a", maxCapacity: 20},
				{key: "c", maxCapacity: 10},
				{key: "d", maxCapacity: 100},
				{key: "b", maxCapacity: 100},
			},
			inputAmount: 101,
			expected: []*merchant{
				{key: "a", maxCapacity: 20, amount: 20},
				{key: "c", maxCapacity: 10, amount: 10},
				{key: "d", maxCapacity: 100, amount: 35},
				{key: "b", maxCapacity: 100, amount: 36},
			},
		},
		{
			name: "valid - sean (from the subscriptions team)",
			input: []*merchant{
				{key: "a", maxCapacity: 20},
				{key: "c", maxCapacity: 10},
				{key: "d", maxCapacity: 110},
				{key: "b", maxCapacity: 110},
				{key: "e", maxCapacity: 100},
				{key: "f", maxCapacity: 100},
			},
			inputAmount: 200,
			expected: []*merchant{
				{key: "a", maxCapacity: 20, amount: 20},
				{key: "c", maxCapacity: 10, amount: 10},
				{key: "d", maxCapacity: 110, amount: 42},
				{key: "b", maxCapacity: 110, amount: 44},
				{key: "e", maxCapacity: 100, amount: 42},
				{key: "f", maxCapacity: 100, amount: 42},
			},
		},
	}

	for _, tt := range tests {
		result := split(tt.input, tt.inputAmount)
		for i := range result {
			fmt.Println(result[i])
			resultByKey := sortByKey(result)
			expectedByKey := sortByKey(tt.expected)
			assert.Equal(t, expectedByKey[i].amount, resultByKey[i].amount)
		}
	}
}
