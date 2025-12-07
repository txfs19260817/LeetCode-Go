package leetcode

import (
	"reflect"
	"testing"
)

func Test_boundaryOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,null,2,3,4] => [1,3,4,2]",
			args: args{
				root: buildTreeFromLevelOrder([]*int{
					i(1), nil, i(2), i(3), i(4),
				}),
			},
			want: []int{1, 3, 4, 2},
		},
		{
			name: "[1,2,3,4,5,6,null,null,null,7,8,9,10] => [1,2,4,7,8,9,10,6,3]",
			args: args{
				root: buildTreeFromLevelOrder([]*int{
					i(1), i(2), i(3), i(4), i(5), i(6), nil,
					nil, nil, i(7), i(8), i(9), i(10),
				}),
			},
			want: []int{1, 2, 4, 7, 8, 9, 10, 6, 3},
		},
		{
			name: "[1,2,9,3,null,5,8,4,null,null,null,6,7] => [1,2,3,4,5,6,7,8,9]",
			args: args{
				root: buildTreeFromLevelOrder([]*int{
					i(1), i(2), i(9), i(3), nil, i(5), i(8),
					i(4), nil, nil, nil, i(6), i(7),
				}),
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := boundaryOfBinaryTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boundaryOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 小工具：快速创建 *int
func i(v int) *int {
	return &v
}

// 根据层序数组（带 nil）建树：
// 输入类似： []*int{i(1), nil, i(2), i(3), i(4)}
func buildTreeFromLevelOrder(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	idx := 1

	for len(queue) > 0 && idx < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// 左孩子
		if idx < len(vals) && vals[idx] != nil {
			left := &TreeNode{Val: *vals[idx]}
			node.Left = left
			queue = append(queue, left)
		}
		idx++

		// 右孩子
		if idx < len(vals) && vals[idx] != nil {
			right := &TreeNode{Val: *vals[idx]}
			node.Right = right
			queue = append(queue, right)
		}
		idx++
	}
	return root
}
