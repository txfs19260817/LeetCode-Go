package leetcode

import "math/bits"

func rangeBitwiseAnd(m int, n int) int {
	var shift int // common prefix + remaining 0s
	for m < n {
		m >>= 1
		n >>= 1
		shift++
	}
	return m << shift
}

func rangeBitwiseAnd2(left int, right int) int {
	m := bits.Len(uint(left ^ right))
	return left &^ (1<<m - 1)
}
