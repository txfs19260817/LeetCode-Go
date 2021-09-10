package _650_2_Keys_Keyboard

func minSteps(n int) int { // Math
	var ans int
	for d := 2; n > 1; d++ {
		for n%d == 0 {
			ans += d
			n /= d
		}
	}
	return ans
}

func minSteps2(n int) int { // DP
	if n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = i
		for j := 2; j <= i/2; j++ {
			if i%j == 0 && dp[j] < dp[i] {
				dp[i] = dp[j] + i/j
			}
		}
	}
	return dp[n]
}
