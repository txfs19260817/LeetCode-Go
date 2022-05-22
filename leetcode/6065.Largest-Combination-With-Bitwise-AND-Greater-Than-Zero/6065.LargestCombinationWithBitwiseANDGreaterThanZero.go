package _065_Largest_Combination_With_Bitwise_AND_Greater_Than_Zero

func largestCombination(candidates []int) int {
	var ans int
	for i := 0; i < 26; i++ { // 10^7
		var sum int
		for _, c := range candidates {
			if (1<<i)&c > 0 {
				sum++
			}
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
