package _173_Binary_Search_Tree_Iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	bstIterator := Constructor(&TreeNode{
		Val:   7,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}},
	})
	assert.Equal(t, 3, bstIterator.Next())
	assert.Equal(t, 7, bstIterator.Next())
	assert.True(t, bstIterator.HasNext())
	assert.Equal(t, 9, bstIterator.Next())
	assert.Equal(t, 15, bstIterator.Next())
	assert.True(t, bstIterator.HasNext())
	assert.Equal(t, 20, bstIterator.Next())
	assert.False(t, bstIterator.HasNext())
}
