package _244_Shortest_Word_Distance_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestWordDistance(t *testing.T) {
	s := Constructor([]string{"a", "c", "b", "b", "a"})
	assert.Equal(t, 1, s.Shortest("a", "b"))
	assert.Equal(t, 1, s.Shortest("c", "b"))
	assert.Equal(t, 1, s.Shortest("b", "a"))
}
