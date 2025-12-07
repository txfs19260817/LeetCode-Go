package leetcode

import "testing"

func Test_deleteTreeNodes(t *testing.T) {
	tests := []struct {
		name   string
		nodes  int
		parent []int
		value  []int
		want   int
	}{
		{
			name:   "Example1",
			nodes:  7,
			parent: []int{-1, 0, 0, 1, 2, 2, 2},
			value:  []int{1, -2, 4, 0, -2, -1, -1},
			want:   2,
		}, {
			name:   "Example2",
			nodes:  7,
			parent: []int{-1, 0, 0, 1, 2, 2, 2},
			value:  []int{1, -2, 4, 0, -2, -1, -2},
			want:   6,
		},
		{
			name:   "Single node non-zero",
			nodes:  1,
			parent: []int{-1},
			value:  []int{5},
			want:   1,
		},
		{
			name:   "Single node zero",
			nodes:  1,
			parent: []int{-1},
			value:  []int{0},
			want:   0,
		},
		{
			name:   "All zero subtree deletes entire tree",
			nodes:  3,
			parent: []int{-1, 0, 0},
			value:  []int{1, -1, 0}, // sums to 0
			want:   0,
		},
		{
			name:   "No node removed",
			nodes:  4,
			parent: []int{-1, 0, 0, 1},
			value:  []int{3, 1, 1, 2},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteTreeNodes(tt.nodes, tt.parent, tt.value); got != tt.want {
				t.Errorf("deleteTreeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
