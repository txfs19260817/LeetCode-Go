package _225_Implement_Stack_using_Queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyStack(t *testing.T) {
	s := Constructor()
	assert.True(t, s.Empty())
	s.Push(1)
	s.Push(2)
	assert.Equal(t, 2, s.Top())
	assert.Equal(t, 2, s.Pop())
	assert.False(t, s.Empty())
}
