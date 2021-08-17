package _303_Range_Sum_Query___Immutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeSumQuery(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	rsq := Constructor(nums)
	cases := []struct {
		i, j int
		ans  int
	}{

		{
			0, 2,
			1,
		},

		{
			2, 5,
			-1,
		},

		{
			0, 5,
			-3,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.ans, rsq.SumRange(tc.i, tc.j))
	}
}
