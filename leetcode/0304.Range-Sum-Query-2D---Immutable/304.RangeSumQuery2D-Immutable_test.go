package _304_Range_Sum_Query_2D___Immutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeSumQuery2D(t *testing.T) {
	mat := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	rsq := Constructor(mat)
	cases := []struct {
		row1, col1, row2, col2 int
		ans                    int
	}{
		{
			2, 1, 4, 3,
			8,
		},
		{
			1, 1, 2, 2,
			11,
		},
		{
			1, 2, 2, 4,
			12,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.ans, rsq.SumRegion(tc.row1, tc.col1, tc.row2, tc.col2))
	}
}
