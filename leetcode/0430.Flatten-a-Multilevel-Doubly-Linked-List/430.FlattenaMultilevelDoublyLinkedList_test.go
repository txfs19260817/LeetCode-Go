package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "head = [1,2,null,3]",
			args: args{&Node{Val: 1, Next: &Node{Val: 2}, Child: &Node{Val: 3}}},
			want: &Node{Val: 1, Next: &Node{Val: 3, Next: &Node{Val: 2}}},
		},
		{
			name: "head = []",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := flatten(tt.args.root)
			if root == nil {
				assert.Equal(t, root, tt.want)
			} else {
				for root != nil {
					assert.Equal(t, root.Val, tt.want.Val)
					root, tt.want = root.Next, tt.want.Next
				}
			}
		})
	}
}
