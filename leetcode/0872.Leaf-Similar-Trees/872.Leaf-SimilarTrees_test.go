package leetcode

import "testing"

const null = -1 << 63

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]",
			args: args{slice2TreeNode([]int{3, 5, 1, 6, 2, 9, 8, null, null, 7, 4}), slice2TreeNode([]int{3, 5, 1, 6, 7, 4, 2, null, null, null, null, null, null, 9, 8})},
			want: true,
		},
		{
			name: "root1 = [1], root2 = [2]",
			args: args{&TreeNode{Val: 1}, &TreeNode{Val: 2}},
			want: false,
		},
		{
			name: "root1 = [1], root2 = [1]",
			args: args{&TreeNode{Val: 1}, &TreeNode{Val: 1}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
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
