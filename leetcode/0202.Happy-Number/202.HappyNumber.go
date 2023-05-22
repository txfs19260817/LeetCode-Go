package leetcode

func isHappy(n int) bool {
	step := func(n int) int {
		var sum int
		for m := n; m > 0; m /= 10 {
			sum += (m % 10) * (m % 10)
		}
		return sum
	}
	s, f := n, step(n)
	for f != 1 && s != f {
		s = step(s)
		f = step(step(f))
	}
	return f == 1
}
