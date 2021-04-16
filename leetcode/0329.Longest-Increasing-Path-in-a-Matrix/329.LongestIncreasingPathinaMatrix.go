package _329_Longest_Increasing_Path_in_a_Matrix

import "sort"

func longestIncreasingPath(matrix [][]int) int {
	ans, elements := 1, make([][]int, 0, len(matrix)*len(matrix[0]))
	for i, line := range matrix {
		for j, v := range line {
			elements = append(elements, []int{v, i, j})
		}
	}
	sort.Slice(elements, func(i, j int) bool {
		return elements[i][0] > elements[j][0]
	})
	dp, dirs := make([][]int, len(matrix)), [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}
	for _, e := range elements {
		v, i, j := e[0], e[1], e[2]
		for _, d := range dirs {
			if x, y := i+d[0], j+d[1]; x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0]) && matrix[x][y] > v && dp[i][j] < dp[x][y]+1 {
				dp[i][j] = dp[x][y] + 1
			}
		}
		if dp[i][j] > ans {
			ans = dp[i][j]
		}
	}
	return ans
}
