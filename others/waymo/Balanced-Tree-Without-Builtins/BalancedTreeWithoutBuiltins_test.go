package waymo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTreeExampleFlow(t *testing.T) {
	tree := NewAVLTree()
	for _, key := range []int{10, 20, 30, 40, 50, 25} {
		tree.Insert(key)
	}
	assert.Equal(t, []int{10, 20, 25, 30, 40, 50}, tree.InOrder())

	tree.Delete(40)
	assert.False(t, tree.Contains(40))
	assert.True(t, tree.Contains(25))
	assert.Equal(t, []int{10, 20, 25, 30, 50}, tree.InOrder())
}

func TestAVLTreeInsertDeleteTable(t *testing.T) {
	tests := []struct {
		name   string
		insert []int
		delete []int
		want   []int
	}{
		{
			name:   "empty",
			insert: []int{},
			delete: []int{},
			want:   []int{},
		},
		{
			name:   "duplicates ignored",
			insert: []int{5, 5, 5, 3, 7, 7},
			delete: []int{},
			want:   []int{3, 5, 7},
		},
		{
			name:   "delete non-existing keys",
			insert: []int{8, 4, 12, 2, 6, 10, 14},
			delete: []int{100, 1},
			want:   []int{2, 4, 6, 8, 10, 12, 14},
		},
		{
			name:   "delete root and internal nodes",
			insert: []int{9, 5, 10, 0, 6, 11, -1, 1, 2},
			delete: []int{10, 9, 1},
			want:   []int{-1, 0, 2, 5, 6, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewAVLTree()
			for _, key := range tt.insert {
				tree.Insert(key)
			}
			for _, key := range tt.delete {
				tree.Delete(key)
			}
			assert.Equal(t, tt.want, tree.InOrder())
			assert.True(t, isBalanced(tree.root))
		})
	}
}

func TestAVLTreeStaysBalancedOnSortedInput(t *testing.T) {
	tree := NewAVLTree()
	for i := 1; i <= 200; i++ {
		tree.Insert(i)
	}
	assert.True(t, isBalanced(tree.root))
	assert.Equal(t, 200, len(tree.InOrder()))
}

func isBalanced(node *avlNode) bool {
	_, ok := checkBalance(node)
	return ok
}

func checkBalance(node *avlNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	lh, lok := checkBalance(node.left)
	if !lok {
		return 0, false
	}
	rh, rok := checkBalance(node.right)
	if !rok {
		return 0, false
	}
	diff := lh - rh
	if diff < -1 || diff > 1 {
		return 0, false
	}
	if node.height != max(lh, rh)+1 {
		return 0, false
	}
	return max(lh, rh) + 1, true
}
