package _130_Minimum_Cost_Tree_From_Leaf_Values

func mctFromLeafValues(arr []int) int { // Monotonic stack
	var ans int
	stack := []int{1 << 31}
	for _, v := range arr {
		for v > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans += top * min(stack[len(stack)-1], v)
		}
		stack = append(stack, v)
	}
	for len(stack) > 2 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans += top * stack[len(stack)-1]
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mctFromLeafValues2(arr []int) int {
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}
	for l := 1; l < len(arr); l++ {
		for i := 0; i < len(arr)-l; i++ {
			j := i + l
			dp[i][j] = 1 << 31
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+maxIntSlice(arr[i:k+1])*maxIntSlice(arr[k+1:j+1])+dp[k+1][j])
			}
		}
	}
	return dp[0][len(arr)-1]
}

func maxIntSlice(v []int) (m int) {
	if len(v) == 0 {
		return 1
	}
	m = v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
		}
	}
	return
}
