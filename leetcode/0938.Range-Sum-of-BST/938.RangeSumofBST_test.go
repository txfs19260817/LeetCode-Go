package _938_Range_Sum_of_BST

import "testing"

const null = -1 << 63

func Test_rangeSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "root = [10,5,15,3,7,null,18], low = 7, high = 15",
			args: args{slice2TreeNode([]int{10, 5, 15, 3, 7, null, 18}), 7, 15},
			want: 32,
		},
		{
			name: "root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10",
			args: args{slice2TreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, null, 6}), 6, 10},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.args.root, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
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
