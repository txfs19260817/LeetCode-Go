package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	start, end := 0, rows*cols-1
	for start+1 < end {
		mid := start + (end-start)/2
		if center := matrix[mid/cols][mid%cols]; center > target {
			end = mid
		} else if center < target {
			start = mid
		} else {
			return true
		}
	}
	if matrix[start/cols][start%cols] == target || matrix[end/cols][end%cols] == target {
		return true
	}
	return false
}
