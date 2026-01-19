package leetcode

func kthSmallest(matrix [][]int, k int) int {
	// Binary search on value range (not index range).
	m, n := len(matrix), len(matrix[0])
	l, r := matrix[0][0], matrix[m-1][n-1]

	// Count how many elements are <= mid
	count := func(mid int) (c int) {
		for i, j := m-1, 0; i >= 0 && j < n; { // from bottom-left
			if matrix[i][j] > mid {
				i--
			} else {
				// All elements above matrix[i][j] in this column are <= mid.
				j++
				c += i + 1
			}
		}
		return
	}

	// Find the smallest value such that count(value) >= k.
	for l <= r {
		if mid := l + (r-l)/2; k <= count(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
