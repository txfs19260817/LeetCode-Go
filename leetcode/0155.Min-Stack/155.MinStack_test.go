package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	assert.Equal(t, -3, minStack.GetMin())
	minStack.Pop()
	assert.Equal(t, 0, minStack.Top())
	assert.Equal(t, -2, minStack.GetMin())
}
