package _008_Construct_Binary_Search_Tree_from_Preorder_Traversal

import (
	"reflect"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	type args struct {
		preorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "[1,3]",
			args: args{[]int{1, 3}},
			want: &TreeNode{Val: 1, Right: &TreeNode{Val: 3}},
		},
		{
			name: "[4,2]",
			args: args{[]int{4, 2}},
			want: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstFromPreorder(tt.args.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
