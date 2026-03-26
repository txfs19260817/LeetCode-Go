package leetcode

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_treeToDoublyList(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want []int
	}{
		{
			name: "nil tree",
			root: nil,
			want: nil,
		},
		{
			name: "single node",
			root: &Node{Val: 1},
			want: []int{1},
		},
		{
			name: "balanced bst",
			root: &Node{
				Val: 4,
				Left: &Node{
					Val:   2,
					Left:  &Node{Val: 1},
					Right: &Node{Val: 3},
				},
				Right: &Node{Val: 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "left skewed bst",
			root: &Node{
				Val: 3,
				Left: &Node{
					Val:  2,
					Left: &Node{Val: 1},
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertCircularList(t, treeToDoublyList(tt.root), tt.want)
		})
	}
}

func assertCircularList(t *testing.T, head *Node, want []int) {
	t.Helper()

	if len(want) == 0 {
		assert.Nil(t, head)
		return
	}

	if !assert.NotNil(t, head) {
		return
	}

	gotForward := make([]int, 0, len(want))
	cur := head
	for i := 0; i < len(want); i++ {
		if !assert.NotNilf(t, cur, "forward traversal hit nil at step %d", i) {
			return
		}
		gotForward = append(gotForward, cur.Val)
		cur = cur.Right
	}
	assert.Same(t, head, cur, "forward traversal did not loop back to head")
	assert.True(t, slices.Equal(gotForward, want), "forward traversal = %v, want %v", gotForward, want)

	gotBackward := make([]int, 0, len(want))
	cur = head.Left
	for i := len(want) - 1; i >= 0; i-- {
		if !assert.NotNilf(t, cur, "backward traversal hit nil at reverse step %d", i) {
			return
		}
		gotBackward = append(gotBackward, cur.Val)
		cur = cur.Left
	}
	wantBackward := slices.Clone(want)
	slices.Reverse(wantBackward)
	assert.Same(t, head.Left, cur, "backward traversal did not loop back to tail")
	assert.True(t, slices.Equal(gotBackward, wantBackward), "backward traversal = %v, want %v", gotBackward, wantBackward)
}
