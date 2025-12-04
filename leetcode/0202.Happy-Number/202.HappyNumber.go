package leetcode

func isHappy(n int) bool {
	step := func(x int) (sum int) {
		for ; x > 0; x /= 10 {
			sum += (x % 10) * (x % 10)
		}
		return
	}
	s, f := n, step(n)
	for f != 1 && s != f {
		s, f = step(s), step(step(f))
	}
	return f == 1
}
