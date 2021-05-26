package _469_Find_All_The_Lonely_Nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const null = -1 << 63

func Test_getLonelyNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "root = [1,2,3,null,4]",
			args: args{slice2TreeNode([]int{1, 2, 3, null, 4})},
			want: []int{4},
		},
		{
			name: "root = [7,1,4,6,null,5,3,null,null,null,null,null,2]",
			args: args{slice2TreeNode([]int{7, 1, 4, 6, null, 5, 3, null, null, null, null, null, 2})},
			want: []int{6, 2},
		},
		{
			name: "root = [11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]",
			args: args{slice2TreeNode([]int{11, 99, 88, 77, null, null, 66, 55, null, null, 44, 33, null, null, 22})},
			want: []int{77, 55, 33, 66, 44, 22},
		},
		{
			name: "root = [197]",
			args: args{&TreeNode{Val: 197}},
			want: nil,
		},
		{
			name: "root = [31,null,78,null,28]",
			args: args{slice2TreeNode([]int{31, null, 78, null, 28})},
			want: []int{78, 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, getLonelyNodes(tt.args.root))
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
