package day2

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		fn       func(string) int
		input    string
		expected int
	}{
		{
			expected: 8,
			input:    `test.txt`,
			fn:       part1,
		},
		{
			expected: 2076,
			input:    `input.txt`,
			fn:       part1,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)))
	}
}
