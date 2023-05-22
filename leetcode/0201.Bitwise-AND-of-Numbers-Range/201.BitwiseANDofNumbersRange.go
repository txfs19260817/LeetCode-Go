package leetcode

func rangeBitwiseAnd(m int, n int) int {
	var shift int // common prefix + remaining 0s
	for m < n {
		m >>= 1
		n >>= 1
		shift++
	}
	return m << shift
}
