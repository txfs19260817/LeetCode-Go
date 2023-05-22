package leetcode

func getSum(a int, b int) int {
	for b != 0 {
		c := uint(a&b) << 1
		a ^= b
		b = int(c)
	}
	return a
}
