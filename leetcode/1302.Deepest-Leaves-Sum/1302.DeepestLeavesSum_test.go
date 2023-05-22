package leetcode

import "testing"

const null = -1 << 63

func Test_deepestLeavesSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "root = [1,2,3,4,5,null,6,7,null,null,null,null,8]",
			args: args{slice2TreeNode([]int{1, 2, 3, 4, 5, null, 6, 7, null, null, null, null, 8})},
			want: 15,
		},
		{
			name: "root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]",
			args: args{slice2TreeNode([]int{6, 7, 8, 2, 7, 1, 3, 9, null, 1, 4, null, null, null, 5})},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepestLeavesSum(tt.args.root); got != tt.want {
				t.Errorf("deepestLeavesSum() = %v, want %v", got, tt.want)
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
