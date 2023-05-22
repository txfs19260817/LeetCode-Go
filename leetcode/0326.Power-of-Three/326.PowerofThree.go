package leetcode

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0 // 3^19 < 2^31 - 1
}
