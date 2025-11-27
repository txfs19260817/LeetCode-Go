package k

import (
	"reflect"
	"testing"
)

func TestBestDayToCross(t *testing.T) {
	tests := []struct {
		name      string
		altitudes []int
		snow      [][]int
		expected  []int
	}{
		{
			name:      "Example 1",
			altitudes: []int{0, 1, 2, 1},
			snow: [][]int{
				{1, 0, 1, 0},
				{0, 0, 0, 0},
				{1, 1, 0, 2},
			},
			expected: []int{2, 1},
		},
		{
			name:      "Immediate Crossing (Day 0)",
			altitudes: []int{0, 0, 0, 0},
			snow: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			expected: []int{0, 0},
		},
		{
			name:      "Impossible Crossing (Always too steep)",
			altitudes: []int{0, 5, 0},
			snow: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: []int{-1, -1},
		},
		{
			name:      "Snow Fills Gap",
			altitudes: []int{0, 2, 0},
			snow: [][]int{
				{1, 0, 1}, // Makes heights [1, 2, 1] -> traversable
			},
			expected: []int{0, 1}, // 1->2 is 1 climb. 2->1 is 0 climbs. Total 1.
		},
		{
			name:      "Melting Helps",
			altitudes: []int{0, 1, 1, 0},
			snow: [][]int{
				{0, 1, 0, 0}, // Day 0: Heights [0, 2, 1, 0]. Diff 0->2 invalid.
				{0, 0, 0, 0}, // Day 1: No melt yet. Heights [0, 2, 1, 0].
				{0, 0, 0, 0}, // Day 2: Pos 1 melts. Heights [0, 1, 1, 0]. Traversable.
			},
			expected: []int{2, 1}, // Day 2, 1 climb (0->1)
		},
		{
			name:      "Late Day Better (Melting reduces climbs)",
			altitudes: []int{0, 0},
			snow: [][]int{
				{0, 1}, // Day 0: Heights [0, 1]. Valid. 1 climb.
				{0, 0}, // Day 1: Heights [0, 1]. Valid. 1 climb.
				{0, 0}, // Day 2: Pos 1 melts. Heights [0, 0]. Valid. 0 climbs.
			},
			expected: []int{2, 0},
		},
		{
			name:      "Prefer Earliest Best Day",
			altitudes: []int{0, 1, 2},
			snow: [][]int{
				{0, 0, 0}, // Day 0: Heights [0, 1, 2]. Valid. 2 climbs.
				{0, 0, 0}, // Day 1: Heights [0, 1, 2]. Valid. 2 climbs.
			},
			expected: []int{0, 2},
		},
		{
			name:      "Complex Melting",
			altitudes: []int{0, 3, 0},
			snow: [][]int{
				{0, 0, 0}, // D0: [0,3,0] Invalid
				{1, 0, 1}, // D1: [1,3,1] Invalid (diff 2)
				{0, 0, 0}, // D2: [1,3,1] Melt? No, snow added on D1.
				{0, 0, 0}, // D3: [1,3,1] Melt? Yes, pos 0&2 melt? Wait.
				// Pos 1: D1 snow 0. D2 snow 0. D3 snow 0.
				// Pos 1 snow was 0 initially. Never added. So snow is 0. Can't melt below 0.
				// Pos 0&2: D1 snow 1. D2 snow 0. D3 snow 0. Melt at D3.
				// Heights: [0, 3, 0]. Still invalid.
			},
			expected: []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestDayToCross(tt.altitudes, tt.snow); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("BestDayToCross() = %v, want %v", got, tt.expected)
			}
		})
	}
}
