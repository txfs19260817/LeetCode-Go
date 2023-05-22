package leetcode

func matrixBlockSum(mat [][]int, k int) [][]int {
	// get prefix mat
	ans, prefix := make([][]int, len(mat)), make([][]int, len(mat))
	for i, line := range mat {
		prefix[i] = make([]int, len(line))
		ans[i] = make([]int, len(line))
	}
	prefix[0][0] = mat[0][0]
	for i := 1; i < len(mat[0]); i++ {
		prefix[0][i] = prefix[0][i-1] + mat[0][i]
	}
	for i := 1; i < len(mat); i++ {
		prefix[i][0] = prefix[i-1][0] + mat[i][0]
	}
	for i := 1; i < len(mat); i++ {
		for j := 1; j < len(mat[0]); j++ {
			prefix[i][j] = mat[i][j] + prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1]
		}
	}
	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans[i]); j++ {
			bottom, right := min(i+k, len(mat)-1), min(j+k, len(mat[0])-1)
			ans[i][j] = prefix[bottom][right]
			if i-k-1 >= 0 {
				ans[i][j] -= prefix[i-k-1][right]
			}
			if j-k-1 >= 0 {
				ans[i][j] -= prefix[bottom][j-k-1]
			}
			if i-k-1 >= 0 && j-k-1 >= 0 {
				ans[i][j] += prefix[i-k-1][j-k-1]
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
