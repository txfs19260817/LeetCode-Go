package _304_Find_N_Unique_Integers_Sum_up_to_Zero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumZero(t *testing.T) {
	sliceSum := func(nums []int) (sum int) {
		for _, num := range nums {
			sum += num
		}
		return
	}
	for n := 1; n <= 1000; n++ {
		got := sumZero(n)
		assert.Len(t, got, n)
		assert.Equal(t, 0, sliceSum(got))
	}
}
