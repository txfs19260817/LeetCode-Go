package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	KthL := Constructor(1, []int{})
	assert.Equal(t, -3, KthL.Add(-3))
	assert.Equal(t, -2, KthL.Add(-2))
	assert.Equal(t, -2, KthL.Add(-4))
	assert.Equal(t, 0, KthL.Add(0))
	assert.Equal(t, 4, KthL.Add(4))
}
