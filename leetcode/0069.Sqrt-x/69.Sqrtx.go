package leetcode

func mySqrt(x int) int {
	l, r := 0, x
	for l+1 < r {
		mid := l + (r-l)/2
		s := mid * mid
		if s == x {
			return mid
		}
		if s > x {
			r = mid
		} else {
			l = mid
		}
	}
	if r*r <= x {
		return r
	}
	return l
}
