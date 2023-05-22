package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	rleIterator := Constructor([]int{3, 8, 0, 9, 2, 5})
	nexts := []int{2, 1, 1, 2}
	res := []int{8, 8, 5, -1}
	for i, n := range nexts {
		assert.Equal(t, res[i], rleIterator.Next(n))
	}
	rleIterator = Constructor([]int{811, 903, 310, 730, 899, 684, 472, 100, 434, 611})
	nexts = []int{358, 345, 154, 265, 73, 220, 138, 4, 170, 88}
	res = []int{903, 903, 730, 684, 684, 684, 684, 684, 684, 684}
	for i, n := range nexts {
		assert.Equal(t, res[i], rleIterator.Next(n))
	}
}
