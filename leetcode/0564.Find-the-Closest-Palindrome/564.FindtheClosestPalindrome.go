package leetcode

import (
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	minDiff, ans := math.MaxInt, 0
	num, _ := strconv.Atoi(n)
	m := len(n)
	update := func(p int) {
		// abs(diff) > 0, as ans == n is not allowed
		if diff := abs(p - num); diff > 0 && (diff < minDiff || diff == minDiff && p < ans) {
			minDiff, ans = diff, p
		}
	}

	update(int(math.Pow10(m-1)) - 1)
	update(int(math.Pow10(m)) + 1)

	left, _ := strconv.Atoi(n[:(m+1)/2])
	for l := left - 1; l <= left+1; l++ {
		p, x := l, l
		if m%2 == 1 {
			x /= 10
		}
		for ; x > 0; x /= 10 {
			p = p*10 + x%10
		}
		update(p)
	}

	return strconv.Itoa(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
