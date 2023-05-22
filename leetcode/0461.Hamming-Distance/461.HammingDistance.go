package leetcode

func hammingDistance(x int, y int) int {
	var ans int
	for z := x ^ y; z != 0; z >>= 1 {
		if z&1 == 1 {
			ans++
		}
	}
	return ans
}
