package _236_Lowest_Common_Ancestor_of_a_Binary_Tree

import (
	"reflect"
	"testing"
)

var t0 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
var t1 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "root = [1,2], p = 1, q = 2",
			args: args{t0, t0, t0.Left},
			want: t0,
		},
		{
			name: "root = [1,2,3], p = 2, q = 3",
			args: args{t1, t1.Left, t1.Right},
			want: t1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
