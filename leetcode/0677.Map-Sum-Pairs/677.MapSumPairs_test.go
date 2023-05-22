package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSumPairs(t *testing.T) {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	assert.Equal(t, 3, mapSum.Sum("ap"))
	mapSum.Insert("app", 2)
	mapSum.Insert("apple", 2)
	assert.Equal(t, 4, mapSum.Sum("ap"))
}
