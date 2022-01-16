package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	tests := []struct {
		name         string
		insert       []string
		search       []string
		delete       []string
		expectedSize int
	}{
		{
			name:         "valid",
			insert:       []string{"sparsh", "cindy", "isha"},
			expectedSize: 3,
		},
	}

	for _, tt := range tests {
		table := Init()
		for _, item := range tt.insert {
			table.Insert(item)
		}
		table.Print()
		assert.Equal(t, tt.expectedSize, table.size)
	}
}
