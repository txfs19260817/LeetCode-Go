package leetcode

import "testing"

const null = -1 << 31

func Test_widthOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,3,2,5,3,null,9]",
			args: args{slice2TreeNode([]int{1, 3, 2, 5, 3, null, 9})},
			want: 4,
		},
		{
			name: "[1,3,null,5,3]",
			args: args{slice2TreeNode([]int{1, 3, null, 5, 3})},
			want: 2,
		},
		{
			name: "[1,1,1,1,null,null,1,1,null,null,1]",
			args: args{slice2TreeNode([]int{1, 1, 1, 1, null, null, 1, 1, null, null, 1})},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := widthOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
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
