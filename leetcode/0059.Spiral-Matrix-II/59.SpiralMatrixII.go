package leetcode

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	top, left, bottom, right := 0, 0, n-1, n-1
	n *= n
	for k := 1; k <= n; {
		for i := left; i <= right && k <= n; i, k = i+1, k+1 {
			ans[top][i] = k
		}
		top++
		for i := top; i <= bottom && k <= n; i, k = i+1, k+1 {
			ans[i][right] = k
		}
		right--
		for i := right; i >= left && k <= n; i, k = i-1, k+1 {
			ans[bottom][i] = k
		}
		bottom--
		for i := bottom; i >= top && k <= n; i, k = i-1, k+1 {
			ans[i][left] = k
		}
		left++
	}
	return ans
}
