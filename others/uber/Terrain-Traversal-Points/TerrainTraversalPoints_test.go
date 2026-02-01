package uber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTerrainTraversalPoints(t *testing.T) {
	tests := []struct {
		name    string
		terrain [][]int
		limits  []int
		want    []int
	}{
		{
			name: "Example 1",
			terrain: [][]int{
				{1, 4, 2, 8},
				{0, 4, 0, 8},
				{1, 2, 0, 8},
			},
			limits: []int{8, 2},
			want:   []int{9, 3},
		},
		{
			name: "Example 2",
			terrain: [][]int{
				{1, 2, 3},
				{2, 5, 7},
				{3, 5, 1},
			},
			limits: []int{5, 6, 2},
			want:   []int{5, 8, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MaxPoints(tt.terrain, tt.limits))
		})
	}
}
