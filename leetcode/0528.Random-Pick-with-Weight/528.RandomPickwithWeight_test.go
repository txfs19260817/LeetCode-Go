package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution_PickIndex(t *testing.T) {
	w := []int{1, 3}
	s, m := Constructor(w), map[int]bool{}
	for i := 0; i < 10000; i++ {
		m[s.PickIndex()] = true
	}
	for i := range w {
		assert.True(t, m[i])
	}
}
