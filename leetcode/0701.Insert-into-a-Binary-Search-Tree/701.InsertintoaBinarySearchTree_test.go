package _701_Insert_into_a_Binary_Search_Tree

import (
	"reflect"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "root = [4,2,7,1,3], val = 5",
			args: args{&TreeNode{
				Val:   4,
				Left:  &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 7},
			}, 5},
			want: &TreeNode{
				Val:   4,
				Left:  &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
