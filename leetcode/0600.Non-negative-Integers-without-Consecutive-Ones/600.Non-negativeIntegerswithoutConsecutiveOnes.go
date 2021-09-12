package _600_Non_negative_Integers_without_Consecutive_Ones

// https://leetcode-cn.com/problems/non-negative-integers-without-consecutive-ones/solution/suan-fa-xiao-ai-wo-lai-gei-ni-jie-shi-qi-4nh4/
func findIntegers(n int) int {
	ans, dp := 0, [31]int{1, 1}
	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	for i, preIs1 := 29, false; i >= 0; i-- {
		if val := 1 << i; val&n > 0 {
			ans += dp[i+1]
			if preIs1 {
				break
			}
			preIs1 = true
		} else {
			preIs1 = false
		}

		if i == 0 {
			ans++
		}
	}
	return ans
}
