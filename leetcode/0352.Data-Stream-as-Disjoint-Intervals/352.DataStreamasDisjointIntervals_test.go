package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	{
		summaryRanges := Constructor()
		summaryRanges.AddNum(1)
		assert.Equal(t, [][]int{{1, 1}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(3)
		assert.Equal(t, [][]int{{1, 1}, {3, 3}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(7)
		assert.Equal(t, [][]int{{1, 1}, {3, 3}, {7, 7}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(2)
		assert.Equal(t, [][]int{{1, 3}, {7, 7}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(6)
		assert.Equal(t, [][]int{{1, 3}, {6, 7}}, summaryRanges.GetIntervals())
	}
	{
		summaryRanges := Constructor()
		summaryRanges.AddNum(6)
		assert.Equal(t, [][]int{{6, 6}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(6)
		assert.Equal(t, [][]int{{6, 6}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(0)
		assert.Equal(t, [][]int{{0, 0}, {6, 6}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(4)
		assert.Equal(t, [][]int{{0, 0}, {4, 4}, {6, 6}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(8)
		assert.Equal(t, [][]int{{0, 0}, {4, 4}, {6, 6}, {8, 8}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(7)
		assert.Equal(t, [][]int{{0, 0}, {4, 4}, {6, 8}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(6)
		assert.Equal(t, [][]int{{0, 0}, {4, 4}, {6, 8}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(4)
		assert.Equal(t, [][]int{{0, 0}, {4, 4}, {6, 8}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(7)
		assert.Equal(t, [][]int{{0, 0}, {4, 4}, {6, 8}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(5)
		assert.Equal(t, [][]int{{0, 0}, {4, 8}}, summaryRanges.GetIntervals())
	}
	{
		summaryRanges := Constructor()
		summaryRanges.AddNum(1)
		assert.Equal(t, [][]int{{1, 1}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(3)
		assert.Equal(t, [][]int{{1, 1}, {3, 3}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(7)
		assert.Equal(t, [][]int{{1, 1}, {3, 3}, {7, 7}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(2)
		assert.Equal(t, [][]int{{1, 3}, {7, 7}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(6)
		assert.Equal(t, [][]int{{1, 3}, {6, 7}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(9)
		assert.Equal(t, [][]int{{1, 3}, {6, 7}, {9, 9}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(4)
		assert.Equal(t, [][]int{{1, 4}, {6, 7}, {9, 9}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(10)
		assert.Equal(t, [][]int{{1, 4}, {6, 7}, {9, 10}}, summaryRanges.GetIntervals())
		summaryRanges.AddNum(5)
		assert.Equal(t, [][]int{{1, 7}, {9, 10}}, summaryRanges.GetIntervals())
	}
}
