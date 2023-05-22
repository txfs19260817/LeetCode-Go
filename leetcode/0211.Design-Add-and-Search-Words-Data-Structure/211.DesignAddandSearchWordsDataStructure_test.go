package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordDictionary(t *testing.T) {
	{
		wordDictionary := Constructor()
		wordDictionary.AddWord("bad")
		wordDictionary.AddWord("abcad")
		wordDictionary.AddWord("dad")
		wordDictionary.AddWord("mad")
		assert.True(t, wordDictionary.Search("mad"))
		assert.False(t, wordDictionary.Search("pad"))
		assert.True(t, wordDictionary.Search(".ad"))
		assert.False(t, wordDictionary.Search(".ae"))
		assert.True(t, wordDictionary.Search("b.."))
		assert.False(t, wordDictionary.Search("c.."))
		assert.True(t, wordDictionary.Search("d.d"))
		assert.False(t, wordDictionary.Search("d.c"))
		assert.True(t, wordDictionary.Search("..."))
		assert.False(t, wordDictionary.Search("......"))
	}
	{
		wordDictionary := Constructor()
		wordDictionary.AddWord("a")
		wordDictionary.AddWord("ab")
		assert.True(t, wordDictionary.Search("a"))
		assert.True(t, wordDictionary.Search("a."))
		assert.True(t, wordDictionary.Search("ab"))
		assert.False(t, wordDictionary.Search(".a"))
		assert.True(t, wordDictionary.Search(".b"))
		assert.False(t, wordDictionary.Search("ab."))
		assert.True(t, wordDictionary.Search("."))
		assert.True(t, wordDictionary.Search(".."))
	}
}
