package leetcode

import "testing"

func Test_averageOfSubtree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   8,
						Left:  &TreeNode{Val: 0},
						Right: &TreeNode{Val: 1},
					},
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 6},
					},
				},
			},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: 1,
		},
		{
			name: "Two nodes",
			args: args{
				root: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 2},
				},
			},
			want: 2,
		},
		{
			name: "All same values",
			args: args{
				root: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 5},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfSubtree(tt.args.root); got != tt.want {
				t.Errorf("averageOfSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
