package _232_Implement_Queue_using_Stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyQueue(t *testing.T) {
	q := Constructor()
	assert.True(t, q.Empty())
	for i := 1; i <= 4; i++ {
		q.Push(i)
	}
	assert.False(t, q.Empty())
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 1, q.Pop())
	q.Push(5)
	for i := 2; i <= 5; i++ {
		assert.False(t, q.Empty())
		assert.Equal(t, i, q.Peek())
		assert.Equal(t, i, q.Pop())
	}
	assert.True(t, q.Empty())
}
