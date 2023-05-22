package leetcode

func constructArray(n int, k int) []int {
	l, r, ans := 1, k+1, make([]int, 0, n)
	for l <= r {
		ans = append(ans, l)
		l++
		if l <= r {
			ans = append(ans, r)
			r--
		}
	}
	for i := k + 2; i <= n; i++ {
		ans = append(ans, i)
	}
	return ans
}
