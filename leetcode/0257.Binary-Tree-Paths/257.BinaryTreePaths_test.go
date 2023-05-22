package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "[1,2,3,null,5]",
			args: args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3},
			}},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "[1]",
			args: args{&TreeNode{Val: 1}},
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, binaryTreePaths(tt.args.root))
		})
	}
}
