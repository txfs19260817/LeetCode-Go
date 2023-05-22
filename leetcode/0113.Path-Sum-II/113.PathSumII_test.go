package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const null = -1 << 63

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22",
			args: args{slice2TreeNode([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1}), 22},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			name: "root = [1,2,3], targetSum = 5",
			args: args{slice2TreeNode([]int{1, 2, 3}), 5},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, pathSum(tt.args.root, tt.args.targetSum), tt.want)
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
