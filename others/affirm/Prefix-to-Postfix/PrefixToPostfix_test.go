package affirm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix(t *testing.T) {
	tests := []struct {
		name   string
		prefix string
		want   string
	}{
		{"Example 1", "+12", "12+"},
		{"Example 2", "-*345", "34*5-"},
		{"Example 3", "*+AB-CD", "AB+CD-*"},
		{"Single operand", "A", "A"},
		{"Numbers and letters", "+A/9B", "A9B/+"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, PrefixToPostfix(tc.prefix))
		})
	}
}
