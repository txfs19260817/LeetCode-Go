package _647_Palindromic_Substrings

func countSubstrings(s string) int {
	var ans int
	expansion := func(s string, l, r int) (count int) {
		for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			count++
		}
		return count
	}
	for i := range s {
		ans += expansion(s, i, i) + expansion(s, i, i+1)
	}
	return ans
}

func countSubstrings2(s string) int {
	ans, dp := 0, make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i+1 < 3 || dp[i+1][j-1]) {
				dp[i][j] = true
				ans++
			}
		}
	}
	return ans
}
