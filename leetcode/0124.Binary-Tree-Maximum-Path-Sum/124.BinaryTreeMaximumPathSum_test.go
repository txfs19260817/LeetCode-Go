package leetcode

import "testing"

const null = -1 << 63

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "root = [1,2,3]",
			args: args{slice2TreeNode([]int{1, 2, 3})},
			want: 6,
		},
		{
			name: "root = [-10,9,20,null,null,15,7]",
			args: args{slice2TreeNode([]int{-10, 9, 20, null, null, 15, 7})},
			want: 42,
		},
		{
			name: "root = [-3]",
			args: args{slice2TreeNode([]int{-3})},
			want: -3,
		},
		{
			name: "root = [-10,-2]",
			args: args{slice2TreeNode([]int{-10, -2})},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func slice2TreeNode(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != null {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != null {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
