package _764_Largest_Plus_Sign

func orderOfLargestPlusSign(n int, mines [][]int) int {
	m := map[[2]int]bool{}
	for _, mine := range mines {
		m[[2]int{mine[0], mine[1]}] = true
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var ans, cnt int
	for r := 0; r < n; r++ {
		cnt = 0
		for c := 0; c < n; c++ {
			if m[[2]int{r, c}] {
				cnt = 0
			} else {
				cnt++
			}
			dp[r][c] = cnt
		}
		cnt = 0
		for c := n - 1; c >= 0; c-- {
			if m[[2]int{r, c}] {
				cnt = 0
			} else {
				cnt++
			}
			if dp[r][c] > cnt {
				dp[r][c] = cnt
			}
		}
	}
	for c := 0; c < n; c++ {
		cnt = 0
		for r := 0; r < n; r++ {
			if m[[2]int{r, c}] {
				cnt = 0
			} else {
				cnt++
			}
			if dp[r][c] > cnt {
				dp[r][c] = cnt
			}
		}
		cnt = 0
		for r := n - 1; r >= 0; r-- {
			if m[[2]int{r, c}] {
				cnt = 0
			} else {
				cnt++
			}
			if dp[r][c] > cnt {
				dp[r][c] = cnt
			}
			if ans < dp[r][c] {
				ans = dp[r][c]
			}
		}
	}
	return ans
}
