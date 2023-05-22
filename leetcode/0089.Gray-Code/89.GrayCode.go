package leetcode

func grayCode(n int) []int {
	ans := make([]int, 0, 1<<n)
	for i := 0; i < 1<<n; i++ {
		ans = append(ans, i^(i>>1))
	}
	return ans
}
