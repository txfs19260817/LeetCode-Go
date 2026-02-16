package waymo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func collect(next func() int, n int) []int {
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, next())
	}
	return out
}

func TestFib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "zero elements",
			n:    0,
			want: []int{},
		},
		{
			name: "first one",
			n:    1,
			want: []int{0},
		},
		{
			name: "first seven",
			n:    7,
			want: []int{0, 1, 1, 2, 3, 5, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Fib()
			assert.Equal(t, tt.want, collect(gen, tt.n))
		})
	}
}

func TestFibMaintainsState(t *testing.T) {
	gen := Fib()
	assert.Equal(t, []int{0, 1, 1, 2, 3}, collect(gen, 5))
	assert.Equal(t, []int{5, 8, 13}, collect(gen, 3))
}
