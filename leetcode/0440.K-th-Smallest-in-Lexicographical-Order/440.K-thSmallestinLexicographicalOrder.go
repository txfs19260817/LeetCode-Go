package leetcode

func findKthNumber(n int, k int) int {
	getSteps := func(n, prefix int) (count int) {
		for cur, next := prefix, prefix+1; cur <= n; cur, next = cur*10, next*10 {
			count += min(n+1, next) - cur
		}
		return
	}
	cur, k := 1, k-1 // skip 0
	for k > 0 {
		steps := getSteps(n, cur)
		if k >= steps { // right
			k -= steps
			cur++
		} else { // down
			k--
			cur *= 10
		}
	}
	return cur
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
