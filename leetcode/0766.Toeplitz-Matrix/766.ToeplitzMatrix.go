package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	for i, row := range matrix {
		for j, e := range row {
			if i >= 1 && j >= 1 && e != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
