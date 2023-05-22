package leetcode

import "testing"

const null = 1 << 60

var eg = slice2TreeNode([]int{1, 1, 3, 1, 1, 3, 4, 3, 1, 1, 1, 3, 8, 4, 8, 3, 3, 1, 6, 2, 1})

func Test_findSecondMinimumValue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,1,3,1,1,3,4,3,1,1,1,3,8,4,8,3,3,1,6,2,1]",
			args: args{eg},
			want: 2,
		},
		{
			name: "[1,1,1,1,5,1,4]",
			args: args{slice2TreeNode([]int{1, 1, 1, 1, 5, 1, 4})},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSecondMinimumValue(tt.args.root); got != tt.want {
				t.Errorf("findSecondMinimumValue() = %v, want %v", got, tt.want)
			}
			if got := findSecondMinimumValue2(tt.args.root); got != tt.want {
				t.Errorf("findSecondMinimumValue2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findSecondMinimumValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findSecondMinimumValue(eg)
	}
}

func Benchmark_findSecondMinimumValue2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findSecondMinimumValue2(eg)
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
