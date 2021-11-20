package _668_Kth_Smallest_Number_in_Multiplication_Table

func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n
	for l+1 < r {
		mid := l + (r-l)/2
		if !enoughK(mid, m, n, k) {
			l = mid
		} else {
			r = mid
		}
	}
	if enoughK(l, m, n, k) {
		return l
	}
	return r
}

func enoughK(num, m, n, k int) bool {
	var cnt int
	for i := 1; i <= m; i++ {
		cnt += min(num/i, n)
	}
	return cnt >= k
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
