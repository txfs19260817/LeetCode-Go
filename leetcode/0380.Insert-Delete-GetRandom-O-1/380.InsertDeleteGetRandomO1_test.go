package _380_Insert_Delete_GetRandom_O_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	{
		randomizedSet := Constructor()
		assert.True(t, randomizedSet.Insert(0))
		assert.True(t, randomizedSet.Insert(1))
		assert.True(t, randomizedSet.Remove(0))
		assert.True(t, randomizedSet.Insert(2))
		assert.True(t, randomizedSet.Remove(1))
		assert.Equal(t, randomizedSet.GetRandom(), 2)
	}
	{
		randomizedSet := Constructor()
		assert.False(t, randomizedSet.Remove(0))
		assert.True(t, randomizedSet.Insert(0))
		assert.False(t, randomizedSet.Insert(0))
		assert.Equal(t, randomizedSet.GetRandom(), 0)
		assert.True(t, randomizedSet.Remove(0))
		assert.True(t, randomizedSet.Insert(0))
	}
}
