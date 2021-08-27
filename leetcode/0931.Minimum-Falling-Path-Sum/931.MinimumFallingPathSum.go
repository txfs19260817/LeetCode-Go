package _931_Minimum_Falling_Path_Sum

func minFallingPathSum(matrix [][]int) int {
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			cost := matrix[i-1][j]
			if j-1 >= 0 {
				cost = min(cost, matrix[i-1][j-1])
			}
			if j+1 < len(matrix[i-1]) {
				cost = min(cost, matrix[i-1][j+1])
			}
			matrix[i][j] += cost
		}

	}
	ans := matrix[len(matrix)-1][0]
	for _, n := range matrix[len(matrix)-1] {
		ans = min(ans, n)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
