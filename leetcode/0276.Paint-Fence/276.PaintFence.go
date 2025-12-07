package leetcode

func numWays(n int, k int) int {
	if n == 1 {
		return k
	}

	f0, f1 := k, k*k // n=1, n=2
	for i := 3; i <= n; i++ {
		f0, f1 = f1, (k-1)*(f0+f1)
	}
	return f1
}
