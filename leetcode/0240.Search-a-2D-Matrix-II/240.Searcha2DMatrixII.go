package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	for i, j := len(matrix)-1, 0; i >= 0; i-- {
		for j < len(matrix[0]) {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {
				break
			} else {
				j++
			}
		}
	}
	return false
}
