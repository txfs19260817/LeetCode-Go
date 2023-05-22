package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := Constructor()

	trie.Insert("abc")
	// ["abc"]

	assert.True(t, trie.Search("abc"), "Search abc in [abc]")
	// ["abc"]

	assert.False(t, trie.Search("ab"), "Search abcd in [a]")
	assert.False(t, trie.Search("abcd"), "Search abcd in [abc]")
	// ["abc"]

	assert.True(t, trie.StartsWith("abc"), "Search startWith abc in [abc]")
	// ["abc"]

	assert.False(t, trie.StartsWith("bc"), "Search startWith bc in [abc]")
	// ["abc"]
}
