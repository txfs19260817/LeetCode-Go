package _057_Count_Nodes_Equal_to_Average_of_Subtree

import "testing"

const null = -1 << 63

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
			name: "[4,8,5,0,1,null,6]",
			args: args{slice2TreeNode([]int{4, 8, 5, 0, 1, null, 6})},
			want: 5,
		},
		{
			name: "[1,null,3,null,1,null,3]",
			args: args{slice2TreeNode([]int{1, null, 3, null, 1, null, 3})},
			want: 1,
		},
		{
			name: "[1]",
			args: args{slice2TreeNode([]int{1})},
			want: 1,
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
