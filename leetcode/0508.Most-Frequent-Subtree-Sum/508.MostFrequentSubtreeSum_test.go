package _508_Most_Frequent_Subtree_Sum

import (
	"testing"
)

const null = -1 << 63

func Test_findFrequentTreeSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[5,2,-3]",
			args: args{slice2TreeNode([]int{5, 2, -3})},
			want: []int{2, -3, 4},
		},
		{
			name: "[5,2,-5]",
			args: args{slice2TreeNode([]int{5, 2, -5})},
			want: []int{2},
		},
		{
			name: "[5,14,null,1]",
			args: args{slice2TreeNode([]int{5, 14, null, 1})},
			want: []int{1, 15, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFrequentTreeSum(tt.args.root); !sameSlice(got, tt.want) {
				t.Errorf("findFrequentTreeSum() = %v, want %v", got, tt.want)
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

func sameSlice(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[int]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}
