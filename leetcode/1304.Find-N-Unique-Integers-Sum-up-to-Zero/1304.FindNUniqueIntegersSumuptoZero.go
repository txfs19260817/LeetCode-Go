package _304_Find_N_Unique_Integers_Sum_up_to_Zero

func sumZero(n int) []int {
	ans := make([]int, 0, n)
	if n%2 == 1 {
		ans = append(ans, 0)
		n--
	}
	for i := 1; len(ans) < n; i++ {
		ans = append(ans, i)
		ans = append(ans, -i)
	}
	return ans
}
