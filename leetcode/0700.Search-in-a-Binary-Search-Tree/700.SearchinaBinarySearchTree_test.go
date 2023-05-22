package leetcode

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
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
			name: `root = [4,2,7,1,3], val = 2`,
			args: args{&TreeNode{
				Val:   4,
				Left:  &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 7},
			}, 2},
			want: &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
		},
		{
			name: `root = [4,2,7,1,3], val = 5`,
			args: args{&TreeNode{
				Val:   4,
				Left:  &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 7},
			}, 5},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
			if got := searchBST2(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST2() = %v, want %v", got, tt.want)
			}
		})
	}
}
