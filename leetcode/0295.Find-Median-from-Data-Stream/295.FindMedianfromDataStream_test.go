package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	assert.Equal(t, 1.5, m.FindMedian())
	m.AddNum(3)
	assert.Equal(t, 2.0, m.FindMedian())
}
