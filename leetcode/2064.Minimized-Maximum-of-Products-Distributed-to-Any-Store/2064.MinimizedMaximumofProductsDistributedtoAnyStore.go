package leetcode

import "sort"

func minimizedMaximum(n int, quantities []int) int {
	return sort.Search(1e5, func(mid int) bool {
		var cnt int
		for _, quantity := range quantities {
			cnt += (quantity + mid) / (mid + 1) // ceil, left starts from 1
		}
		return cnt <= n
	}) + 1
}

func minimizedMaximum2(n int, quantities []int) int {
	l, r := 1, 100000
	for l < r {
		mid, cnt := l+(r-l)/2, 0
		for _, quantity := range quantities {
			cnt += (quantity + mid - 1) / mid
		}
		if cnt <= n {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
