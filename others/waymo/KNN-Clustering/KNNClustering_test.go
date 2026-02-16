package waymo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignClusters(t *testing.T) {
	tests := []struct {
		name    string
		train   []LabeledPoint
		queries []Point
		k       int
		want    []int
	}{
		{
			name: "example",
			train: []LabeledPoint{
				{Point: Point{X: 0, Y: 0}, Label: 0},
				{Point: Point{X: 0, Y: 1}, Label: 0},
				{Point: Point{X: 10, Y: 10}, Label: 1},
				{Point: Point{X: 10, Y: 11}, Label: 1},
			},
			queries: []Point{
				{X: 0.2, Y: 0.1},
				{X: 9.8, Y: 10.2},
				{X: 5, Y: 5},
			},
			k:    3,
			want: []int{0, 1, 0},
		},
		{
			name: "k greater than training size",
			train: []LabeledPoint{
				{Point: Point{X: 0, Y: 0}, Label: 1},
				{Point: Point{X: 10, Y: 10}, Label: 2},
			},
			queries: []Point{
				{X: 1, Y: 1},
			},
			k:    10,
			want: []int{1},
		},
		{
			name: "tie broken by label id",
			train: []LabeledPoint{
				{Point: Point{X: -1, Y: 0}, Label: 5},
				{Point: Point{X: 1, Y: 0}, Label: 3},
			},
			queries: []Point{
				{X: 0, Y: 0},
			},
			k:    2,
			want: []int{3},
		},
		{
			name:  "invalid k or empty train",
			train: []LabeledPoint{},
			queries: []Point{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
			},
			k:    3,
			want: []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AssignClusters(tt.train, tt.queries, tt.k))
		})
	}
}
