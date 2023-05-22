package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomFlipMatrix(t *testing.T) {
	s := Constructor(3, 1)
	candidates := [][]int{{0, 0}, {1, 0}, {2, 0}}
	candidatesCpy := candidates

	flip := s.Flip()
	assert.Contains(t, candidates, flip)
	candidates = remove(candidates, flip)
	flip = s.Flip()
	assert.Contains(t, candidates, flip)
	candidates = remove(candidates, flip)
	assert.Contains(t, candidates, s.Flip())

	s.Reset()
	assert.Contains(t, candidatesCpy, s.Flip())
}

func remove(arr [][]int, target []int) [][]int {
	for i, a := range arr {
		if a[0] == target[0] && a[1] == target[1] {
			arr[i] = arr[len(arr)-1]
			break
		}
	}
	return arr[:len(arr)-1]
}
