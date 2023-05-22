package leetcode

var pick int

func guess(n int) int {
	switch {
	case pick > n:
		return 1
	case pick < n:
		return -1
	default:
		return 0
	}
}

func guessNumber(n int) int {
	l, r := 1, n
	for l+1 < r {
		mid := l + (r-l)/2
		if guess(mid) == 0 {
			return mid
		}
		if guess(mid) > 0 {
			l = mid
		} else {
			r = mid
		}
	}
	if guess(r) == 0 {
		return r
	}
	return l
}
