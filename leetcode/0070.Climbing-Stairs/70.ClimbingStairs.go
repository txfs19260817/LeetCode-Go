package _070_Climbing_Stairs

func climbStairs(n int) int {
	a, b := 1, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func climbStairs1(n int) int {
	if n <= 3 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a
		b += a
	}
	return b
}
