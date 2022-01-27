package matchparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "happy",
			input:    "()",
			expected: "()",
		},
		{
			name:     "happy - open",
			input:    "((a)",
			expected: "(a)",
		},
		{
			name:     "happy - closed",
			input:    "(a))()",
			expected: "(a)()",
		},
		{
			name:     "happy - closed",
			input:    "(a)(()",
			expected: "(a)()",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, matchParentheses(tt.input))
	}
}
