package leetcode

// Binary search
func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l+1 < r {
		mid := l + (r-l)/2
		if s := mid * mid; s == num {
			return true
		} else if s > num {
			r = mid
		} else {
			l = mid
		}
	}
	return l*l == num || r*r == num
}

// Newton y = x**2 - num
func isPerfectSquare2(num int) bool {
	for x0 := float64(num); ; {
		x1 := (x0 + float64(num)/x0) / 2
		if x0-x1 < 1e-6 {
			x := int(x0)
			return x*x == num
		}
		x0 = x1
	}
}
