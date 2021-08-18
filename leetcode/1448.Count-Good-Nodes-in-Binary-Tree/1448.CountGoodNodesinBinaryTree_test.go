package _448_Count_Good_Nodes_in_Binary_Tree

import "testing"

const null = -1 << 63

func Test_goodNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "root = [3,1,4,3,null,1,5]",
			args: args{slice2TreeNode([]int{3,1,4,3,null,1,5})},
			want: 4,
		},
		{
			name: "root = [3,3,null,4,2]",
			args: args{slice2TreeNode([]int{3,3,null,4,2})},
			want: 3,
		},
		{
			name: "root = [1]",
			args: args{slice2TreeNode([]int{1})},
			want: 1,
		},
		{
			name: "root = [2,null,4,10,8,null,null,4]",
			args: args{slice2TreeNode([]int{2,null,4,10,8,null,null,4})},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodNodes(tt.args.root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
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
