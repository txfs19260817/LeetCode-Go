package _446_Arithmetic_Slices_II___Subsequence

func numberOfArithmeticSlices(nums []int) int {
	ans, dp := 0, make([]map[int]int, len(nums))
	for i, x := range nums {
		dp[i] = map[int]int{}
		for j, y := range nums[:i] {
			d := x - y
			cnt := dp[j][d]
			ans += cnt
			dp[i][d] += cnt + 1
		}
	}
	return ans
}
