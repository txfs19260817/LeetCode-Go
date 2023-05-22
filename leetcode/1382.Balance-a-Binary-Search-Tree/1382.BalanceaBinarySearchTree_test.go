package leetcode

import (
	"reflect"
	"testing"
)

func Test_balanceBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "[2,1,3]",
			args: args{&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			}},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balanceBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("balanceBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
