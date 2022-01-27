package parenthesis

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid input",
			input:    "{}",
			expected: true,
		},
		{
			name:     "valid input mixed",
			input:    "{({[{}({})]})}",
			expected: true,
		},
		{
			name:     "invalid input",
			input:    "{}}",
			expected: false,
		},
	}

	for _, tt := range tests {
		if tt.expected != IsValid(tt.input) {
			t.Errorf("error")
		}
	}
}
