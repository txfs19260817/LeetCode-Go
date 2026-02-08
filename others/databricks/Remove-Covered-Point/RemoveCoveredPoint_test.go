package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteCoveredPoint_Example1(t *testing.T) {
	// Flattened: [10,11,13,14,15,4,5,6,7]. idx=3 → point 14 → split [13,16]
	intervals := [][]int{{10, 12}, {13, 16}, {4, 8}}
	result := DeleteCoveredPoint(intervals, 3)
	assert.Equal(t, [][]int{{10, 12}, {13, 14}, {15, 16}, {4, 8}}, result)
}

func TestDeleteCoveredPoint_Example2(t *testing.T) {
	// Flattened: [4,5,6,7,13,14,15,10,11]. idx=0 → point 4 → shrink left
	intervals := [][]int{{4, 8}, {13, 16}, {10, 12}}
	result := DeleteCoveredPoint(intervals, 0)
	assert.Equal(t, [][]int{{5, 8}, {13, 16}, {10, 12}}, result)
}

func TestDeleteCoveredPoint_Example3(t *testing.T) {
	// Flattened: [2,3,4,5,8,9,15,16,17]. idx=3 → point 5 → shrink right
	intervals := [][]int{{2, 6}, {8, 10}, {15, 18}}
	result := DeleteCoveredPoint(intervals, 3)
	assert.Equal(t, [][]int{{2, 5}, {8, 10}, {15, 18}}, result)
}

func TestDeleteCoveredPoint_SingleElementInterval(t *testing.T) {
	// [[5,6]] has one point: 5. idx=0 → remove entirely
	intervals := [][]int{{5, 6}}
	result := DeleteCoveredPoint(intervals, 0)
	assert.Equal(t, [][]int{}, result)
}

func TestDeleteCoveredPoint_SplitLargeInterval(t *testing.T) {
	// [[1,10]] covers [1..9]. idx=4 → point 5 → split into [1,5] and [6,10]
	intervals := [][]int{{1, 10}}
	result := DeleteCoveredPoint(intervals, 4)
	assert.Equal(t, [][]int{{1, 5}, {6, 10}}, result)
}
