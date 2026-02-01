package uber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNAryTreePathSum(t *testing.T) {
	tests := []struct {
		name             string
		root             *Node
		expectedTotalSum int
		expectedMaxSum   int
		expectedPath     []int
	}{
		{
			name: "Example Tree",
			//       1
			//     /   \
			//    2     3
			//        /   \
			//       4     5
			root: &Node{
				Val: 1,
				Children: []*Node{
					{Val: 2},
					{
						Val: 3,
						Children: []*Node{
							{Val: 4},
							{Val: 5},
						},
					},
				},
			},
			expectedTotalSum: 15,
			expectedMaxSum:   9,
			expectedPath:     []int{1, 3, 5},
		},
		{
			name: "Tree with negatives",
			//       1
			//     /   \
			//   -5     2
			//        /   \
			//       3    -1
			root: &Node{
				Val: 1,
				Children: []*Node{
					{Val: -5},
					{
						Val: 2,
						Children: []*Node{
							{Val: 3},
							{Val: -1},
						},
					},
				},
			},
			expectedTotalSum: 0,
			expectedMaxSum:   6,
			expectedPath:     []int{1, 2, 3},
		},
		{
			name:             "Single Node",
			root:             &Node{Val: 10},
			expectedTotalSum: 10,
			expectedMaxSum:   10,
			expectedPath:     []int{10},
		},
		{
			name:             "Nil Root",
			root:             nil,
			expectedTotalSum: 0,
			expectedMaxSum:   0,
			expectedPath:     []int(nil),
		},
		{
			name: "Deep linear tree",
			// 1 -> 2 -> 3
			root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 2,
						Children: []*Node{
							{Val: 3},
						},
					},
				},
			},
			expectedTotalSum: 6,
			expectedMaxSum:   6,
			expectedPath:     []int{1, 2, 3},
		},
		{
			name: "Branching with negative leaf impacting path",
			//      10
			//     /  \
			//    5    5
			//   /      \
			// -10       1
			root: &Node{
				Val: 10,
				Children: []*Node{
					{
						Val: 5,
						Children: []*Node{
							{Val: -10},
						},
					},
					{
						Val: 5,
						Children: []*Node{
							{Val: 1},
						},
					},
				},
			},
			expectedTotalSum: 11,
			expectedMaxSum:   16,
			expectedPath:     []int{10, 5, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedTotalSum, SumAllValues(tt.root), "SumAllValues mismatch")
			assert.Equal(t, tt.expectedMaxSum, MaxPathSum(tt.root), "MaxPathSum mismatch")

			sum, path := MaxPathValues(tt.root)
			assert.Equal(t, tt.expectedMaxSum, sum, "MaxPathValues sum mismatch")
			assert.Equal(t, tt.expectedPath, path, "MaxPathValues path mismatch")
		})
	}
}
