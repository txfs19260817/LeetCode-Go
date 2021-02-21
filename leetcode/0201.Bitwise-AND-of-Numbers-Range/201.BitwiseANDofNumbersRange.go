package _201_Bitwise_AND_of_Numbers_Range

func rangeBitwiseAnd(m int, n int) int {
	var shift int
	for m < n {
		m >>= 1
		n >>= 1
		shift++
	}
	return m << shift
}
