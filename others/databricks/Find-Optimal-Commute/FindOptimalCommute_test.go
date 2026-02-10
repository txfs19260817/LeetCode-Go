package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOptimalCommute_Sample(t *testing.T) {
	grid := [][]string{
		{"3", "3", "S", "2", "X"},
		{"3", "1", "1", "2", "X"},
		{"3", "1", "1", "2", "2"},
		{"3", "1", "1", "1", "D"},
		{"3", "3", "3", "3", "4"},
		{"4", "4", "4", "4", "4"},
	}
	costs := []int{0, 1, 3, 2}
	times := []int{3, 2, 1, 1}
	assert.Equal(t, "Bike", FindOptimalCommute(grid, costs, times))
}

func TestFindOptimalCommute_TieBreakByCost(t *testing.T) {
	// Walk and Bike both reach D in 4 steps.
	// Walk: time=4*1=4, cost=4*1=4.  Bike: time=4*1=4, cost=4*2=8.
	// Same time → Walk wins (lower cost).
	grid := [][]string{
		{"1", "S", "2"},
		{"1", "X", "2"},
		{"1", "D", "2"},
	}
	costs := []int{1, 2, 0, 0}
	times := []int{1, 1, 0, 0}
	assert.Equal(t, "Walk", FindOptimalCommute(grid, costs, times))
}

func TestFindOptimalCommute_NoPath(t *testing.T) {
	grid := [][]string{
		{"S", "X"},
		{"X", "D"},
	}
	costs := []int{1, 1, 1, 1}
	times := []int{1, 1, 1, 1}
	assert.Equal(t, "", FindOptimalCommute(grid, costs, times))
}

func TestFindOptimalCommute_SingleMode(t *testing.T) {
	// Only Walk cells, direct path: S → 1 → 1 → D (3 steps).
	grid := [][]string{
		{"S", "1", "1", "D"},
	}
	costs := []int{2, 5, 5, 5}
	times := []int{3, 9, 9, 9}
	assert.Equal(t, "Walk", FindOptimalCommute(grid, costs, times))
}

func TestFindOptimalCommute_EmptyGrid(t *testing.T) {
	assert.Equal(t, "", FindOptimalCommute(nil, nil, nil))
}
