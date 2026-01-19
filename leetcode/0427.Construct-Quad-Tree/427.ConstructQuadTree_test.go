package leetcode

import (
	"testing"
)

func Test_construct(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Example 1: Mixed grid",
			args: args{
				grid: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: &Node{
				Val:    true, // Val is technically arbitrary for internal nodes, but the implementation uses true or false.
				IsLeaf: false,
				TopLeft: &Node{
					Val:    false,
					IsLeaf: true,
				},
				TopRight: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomRight: &Node{
					Val:    false,
					IsLeaf: true,
				},
			},
		},
		{
			name: "Example 2: All 1s",
			args: args{
				grid: [][]int{
					{1, 1},
					{1, 1},
				},
			},
			want: &Node{
				Val:    true,
				IsLeaf: true,
			},
		},
		{
			name: "Example 3: All 0s",
			args: args{
				grid: [][]int{
					{0, 0},
					{0, 0},
				},
			},
			want: &Node{
				Val:    false,
				IsLeaf: true,
			},
		},
		{
			name: "Example 4: 1x1 grid 1",
			args: args{
				grid: [][]int{{1}},
			},
			want: &Node{
				Val:    true,
				IsLeaf: true,
			},
		},
		{
			name: "Example 5: 1x1 grid 0",
			args: args{
				grid: [][]int{{0}},
			},
			want: &Node{
				Val:    false,
				IsLeaf: true,
			},
		},
		{
			name: "Example 6: Complex grid",
			args: args{
				grid: [][]int{
					{1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 1, 0, 0, 0, 0},
				},
			},
			want: &Node{
				Val:    true,
				IsLeaf: false,
				TopLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				TopRight: &Node{
					Val:    true,
					IsLeaf: false,
					TopLeft: &Node{
						Val:    false,
						IsLeaf: true,
					},
					TopRight: &Node{
						Val:    false,
						IsLeaf: true,
					},
					BottomLeft: &Node{
						Val:    true,
						IsLeaf: true,
					},
					BottomRight: &Node{
						Val:    true,
						IsLeaf: true,
					},
				},
				BottomLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomRight: &Node{
					Val:    false,
					IsLeaf: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct(tt.args.grid); !areEqual(got, tt.want) {
				t.Errorf("construct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func areEqual(n1, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	if n1.IsLeaf != n2.IsLeaf {
		return false
	}
	if n1.IsLeaf {
		return n1.Val == n2.Val
	}
	// For non-leaf nodes, ignore Val, check children
	return areEqual(n1.TopLeft, n2.TopLeft) &&
		areEqual(n1.TopRight, n2.TopRight) &&
		areEqual(n1.BottomLeft, n2.BottomLeft) &&
		areEqual(n1.BottomRight, n2.BottomRight)
}
