package _926_Flip_String_to_Monotone_Increasing

func minFlipsMonoIncr(s string) int {
	ans, dp := len(s), make([]int, len(s)+1)
	for i := range s {
		dp[i+1] = dp[i] + int(s[i]-'0')
	}

	for x := 0; x <= len(s); x++ {
		if times := dp[x] + len(s) - x - (dp[len(s)] - dp[x]); ans > times {
			ans = times
		}
	}
	return ans
}
