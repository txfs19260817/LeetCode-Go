package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	{
		randomizedCollection := Constructor()
		assert.True(t, randomizedCollection.Insert(4))
		assert.True(t, randomizedCollection.Insert(3))
		assert.False(t, randomizedCollection.Insert(4))
		assert.True(t, randomizedCollection.Insert(2))
		assert.False(t, randomizedCollection.Insert(4))
		assert.True(t, randomizedCollection.Remove(4))
		assert.True(t, randomizedCollection.Remove(3))
		assert.True(t, randomizedCollection.Remove(4))
		assert.True(t, randomizedCollection.Remove(4))
	}
	{
		randomizedCollection := Constructor()
		assert.True(t, randomizedCollection.Insert(0))
		assert.True(t, randomizedCollection.Insert(1))
		assert.True(t, randomizedCollection.Remove(0))
		assert.True(t, randomizedCollection.Insert(2))
		assert.True(t, randomizedCollection.Remove(1))
		assert.Equal(t, randomizedCollection.GetRandom(), 2)
	}
	{
		randomizedCollection := Constructor()
		assert.True(t, randomizedCollection.Insert(1))
		assert.True(t, randomizedCollection.Remove(1))
		assert.True(t, randomizedCollection.Insert(1))
		assert.Equal(t, randomizedCollection.GetRandom(), 1)
	}
}
