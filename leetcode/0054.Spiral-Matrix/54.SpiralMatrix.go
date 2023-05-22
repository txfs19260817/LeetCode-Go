package leetcode

func spiralOrder(matrix [][]int) []int {
	top, left, bottom, right, count := 0, 0, len(matrix)-1, len(matrix[0])-1, len(matrix)*len(matrix[0])
	ans := make([]int, 0, count)
	for count > 0 {
		for i := left; i <= right && count > 0; i++ {
			ans = append(ans, matrix[top][i])
			count--
		}
		top++
		for i := top; i <= bottom && count > 0; i++ {
			ans = append(ans, matrix[i][right])
			count--
		}
		right--
		for i := right; i >= left && count > 0; i-- {
			ans = append(ans, matrix[bottom][i])
			count--
		}
		bottom--
		for i := bottom; i >= top && count > 0; i-- {
			ans = append(ans, matrix[i][left])
			count--
		}
		left++
	}
	return ans
}
