package _284_Peeking_Iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	peekingIterator := Constructor(NewIterator([]int{1, 2, 3}))
	assert.Equal(t, 1, peekingIterator.next())
	assert.True(t, peekingIterator.hasNext())
	assert.Equal(t, 2, peekingIterator.peek())
	assert.Equal(t, 2, peekingIterator.next())
	assert.Equal(t, 3, peekingIterator.next())
	assert.False(t, peekingIterator.hasNext())
}
