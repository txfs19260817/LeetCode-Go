package uber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       int
	}{
		{
			name:       "Example 1",
			expression: "add(add(1,3), sub(1,3))",
			want:       2,
		},
		{
			name:       "Example 2",
			expression: "sub(1,3)",
			want:       -2,
		},
		{
			name:       "Example 3",
			expression: "add(-1, 3)",
			want:       2,
		},
		{
			name:       "Complex Nested with Spaces",
			expression: " add(  add( -10 , sub(20, 5) ), sub( -5, add(3,2) ) ) ",
			want:       -5,
		},
		{
			name:       "Nested Sub and Add",
			expression: "sub(add(5, sub(-2, -3)), add(-1,4))",
			want:       3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Evaluate(tt.expression))
		})
	}
}
