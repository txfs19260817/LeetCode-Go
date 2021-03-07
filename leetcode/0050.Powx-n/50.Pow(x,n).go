package _050_Powx_n

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	ans := myPow(x, n/2)
	ans *= ans
	switch n % 2 {
	case 1:
		ans *= x
	case -1:
		ans *= 1 / x
	}
	return ans
}

func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	ans, contrib, positive := 1.0, x, false
	if n > 0 {
		positive = true
	} else {
		n = -1 * n
	}
	for ; n > 0; n /= 2 {
		if n%2 == 1 {
			ans *= contrib
		}
		contrib *= contrib
	}
	if positive {
		return ans
	}
	return 1 / ans
}
