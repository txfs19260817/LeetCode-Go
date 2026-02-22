package waymo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKMeans(t *testing.T) {
	tests := []struct {
		name            string
		points          [][]float64
		k               int
		maxIter         int
		wantCentroids   [][]float64
		wantAssignments []int
	}{
		{
			name: "2d example",
			points: [][]float64{
				{1, 1}, {1.5, 2}, {3, 4}, {5, 7}, {3.5, 5}, {4.5, 5}, {3.5, 4.5},
			},
			k:       2,
			maxIter: 10,
			wantCentroids: [][]float64{
				{1.25, 1.5}, {3.9, 5.1},
			},
			wantAssignments: []int{0, 0, 1, 1, 1, 1, 1},
		},
		{
			name: "3d example",
			points: [][]float64{
				{0, 0, 0}, {0, 1, 0}, {9, 9, 9}, {10, 9, 9},
			},
			k:       2,
			maxIter: 10,
			wantCentroids: [][]float64{
				{0, 0.5, 0}, {9.5, 9, 9},
			},
			wantAssignments: []int{0, 0, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCentroids, gotAssignments := KMeans(tt.points, tt.k, tt.maxIter)
			assert.Equal(t, tt.wantAssignments, gotAssignments)
			assert.Equal(t, len(tt.wantCentroids), len(gotCentroids))
			for i := range tt.wantCentroids {
				for d := range tt.wantCentroids[i] {
					assert.InDelta(t, tt.wantCentroids[i][d], gotCentroids[i][d], 1e-6)
				}
			}
		})
	}
}
