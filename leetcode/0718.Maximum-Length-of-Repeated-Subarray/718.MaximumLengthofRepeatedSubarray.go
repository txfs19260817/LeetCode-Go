package _718_Maximum_Length_of_Repeated_Subarray

func findLength(A []int, B []int) int {
	ans, dp := 0, make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i] == B[j] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}
